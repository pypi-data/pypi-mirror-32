# -*- coding: utf-8 -*-
# Copyright (c) 2016  Red Hat, Inc.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
#
# Written by Jan Kaluza <jkaluza@redhat.com>

import itertools

from freshmaker import conf

try:
    from inspect import signature
except ImportError:
    from funcsigs import signature


class BaseEvent(object):

    _parsers = {}

    def __init__(self, msg_id, manual=False, dry_run=False):
        """
        A base class to abstract events from different fedmsg messages.
        :param msg_id: the id of the msg (e.g. 2016-SomeGUID)
        :param manual: True if the event was trigerred manually by Freshmaker
            REST API.
        :param dry_run: True if the event should be handled in DRY_RUN mode.
        """
        self.msg_id = msg_id
        self.manual = manual
        self.dry_run = dry_run

        # Moksha calls `consumer.validate` on messages that it receives, and
        # even though we have validation turned off in the config there's still
        # a step that tries to access `msg['body']`, `msg['topic']` and
        # `msg.get('topic')`.
        # These are here just so that the `validate` method won't raise an
        # exception when we push our fake messages through.
        # Note that, our fake message pushing has worked for a while... but the
        # *latest* version of fedmsg has some code that exercises the bug.  I
        # didn't hit this until I went to test in jenkins.
        self.body = {}
        self.topic = None

    @classmethod
    def register_parser(cls, parser_class):
        """
        Registers a parser for BaseEvent which is used to parse
        fedmsg in `from_fedmsg(...)` method.
        """
        BaseEvent._parsers[parser_class.name] = parser_class()

    @classmethod
    def get_parsed_topics(cls):
        """
        Returns the list of topics this class is parsing using the
        registered parsers.
        """
        topic_suffixes = []
        for parser in BaseEvent._parsers.values():
            topic_suffixes.extend(parser.topic_suffixes)
        return ['{}.{}'.format(pref.rstrip('.'), cat)
                for pref, cat
                in itertools.product(
                    conf.messaging_topic_prefix,
                    topic_suffixes)]

    def __repr__(self):
        init_sig = signature(self.__init__)

        args_strs = (
            "{}={!r}".format(name, getattr(self, name))
            if param.default != param.empty
            else repr(getattr(self, name, {}))
            for name, param in init_sig.parameters.items())

        return "{}({})".format(type(self).__name__, ', '.join(args_strs))

    def __getitem__(self, key):
        """ Used to trick moksha into thinking we are a dict. """
        return getattr(self, key)

    def __setitem__(self, key, value):
        """ Used to trick moksha into thinking we are a dict. """
        return setattr(self, key, value)

    def get(self, key, value=None):
        """ Used to trick moksha into thinking we are a dict. """
        return getattr(self, key, value)

    def __json__(self):
        return dict(msg_id=self.msg_id, topic=self.topic, body=self.body)

    @staticmethod
    def from_fedmsg(topic, msg):
        """
        Takes a fedmsg topic and message and converts it to a BaseEvent
        object.
        :param topic: the topic of the fedmsg message
        :param msg: the message contents from the fedmsg message
        :return: an object of BaseEvent descent if the message is a type
        that the app looks for, otherwise None is returned
        """
        for parser in BaseEvent._parsers.values():
            if not parser.can_parse(topic, msg):
                continue

            return parser.parse(topic, msg)

        return None

    @property
    def search_key(self):
        """
        Returns the searchable key which is used to query for particular
        events using the JSON API.
        """
        return self.msg_id


class MBSModuleStateChangeEvent(BaseEvent):
    """ A class that inherits from BaseEvent to provide an event
    object for a module event generated by module-build-service
    :param msg_id: the id of the msg (e.g. 2016-SomeGUID)
    :param module_build_id: the id of the module build
    :param module_build_state: the state of the module build
    """
    def __init__(self, msg_id, module, stream, build_id, build_state, **kwargs):
        super(MBSModuleStateChangeEvent, self).__init__(msg_id, **kwargs)
        self.module = module
        self.stream = stream
        self.build_id = build_id
        self.build_state = build_state

    @property
    def search_key(self):
        return str(self.build_id)


class GitModuleMetadataChangeEvent(BaseEvent):
    """
    Provides an event object for "Module metadata in dist-git updated".
    :param scm_url: SCM URL of a updated module.
    :param branch: Branch of updated module.
    """
    def __init__(self, msg_id, module, branch, rev, **kwargs):
        super(GitModuleMetadataChangeEvent, self).__init__(msg_id, **kwargs)
        self.module = module
        self.branch = branch
        self.rev = rev

    @property
    def search_key(self):
        return "%s/%s?#%s" % (self.module, self.branch, self.rev)


class GitRPMSpecChangeEvent(BaseEvent):
    """
    Provides an event object for "RPM spec file in dist-git updated".

    :param rpm: RPM name, also known as the name of component or source package.
    :param branch: Branch of updated RPM spec.
    :param rev: revision.
    """
    def __init__(self, msg_id, rpm, branch, rev, **kwargs):
        super(GitRPMSpecChangeEvent, self).__init__(msg_id, **kwargs)
        self.rpm = rpm
        self.branch = branch
        self.rev = rev

    @property
    def search_key(self):
        return "%s/%s?#%s" % (self.rpm, self.branch, self.rev)


class TestingEvent(BaseEvent):
    """
    Event useds in unit-tests.
    """
    def __init__(self, msg_id, **kwargs):
        super(TestingEvent, self).__init__(msg_id, **kwargs)


class GitDockerfileChangeEvent(BaseEvent):
    """Represent the message omitted when Dockerfile is changed in a push"""

    def __init__(self, msg_id, container, branch, rev, **kwargs):
        super(GitDockerfileChangeEvent, self).__init__(msg_id, **kwargs)
        self.container = container
        self.branch = branch
        self.rev = rev

    @property
    def search_key(self):
        return "%s/%s?#%s" % (self.container, self.branch, self.rev)


class BodhiUpdateCompleteStableEvent(BaseEvent):
    """Event when RPMs are available in Fedora master mirrors

    Refer to an example in datagrepper:

    https://apps.fedoraproject.org/datagrepper/raw?delta=572800& \
        topic=org.fedoraproject.prod.bodhi.update.complete.stable
    """

    def __init__(self, msg_id, update_id, builds, release, **kwargs):
        """Initiate event with data from message got from fedmsg

        Not complete data is required, only part of attributes that are useful
        for rebuild are stored in this event.

        :param str update_id: the Bodhi update ID got from message.
        :param list builds: a list of maps, each of them contains build NVRs
            that are useful for getting RPMs for the rebuild.
        :param dist release: a map contains release information, e.g. name and
            branch. Refer to the example given above to see all available
            attributes in a message.
        """
        super(BodhiUpdateCompleteStableEvent, self).__init__(msg_id, **kwargs)
        self.update_id = update_id
        self.builds = builds
        self.release = release

    @property
    def search_key(self):
        return str(self.update_id)


class KojiTaskStateChangeEvent(BaseEvent):
    """
    Provides an event object for "the state of task changed in koji"
    """
    def __init__(self, msg_id, task_id, task_state, **kwargs):
        super(KojiTaskStateChangeEvent, self).__init__(msg_id, **kwargs)
        self.task_id = task_id
        self.task_state = task_state


class ErrataBaseEvent(BaseEvent):
    def __init__(self, msg_id, advisory, **kwargs):
        """
        Creates new ErrataBaseEvent.

        :param str msg_id: Message id.
        :param ErrataAdvisory advisory: Errata advisory associated with event.
        """
        super(ErrataBaseEvent, self).__init__(msg_id, **kwargs)
        self.advisory = advisory

    @property
    def search_key(self):
        return str(self.advisory.errata_id)


class ErrataAdvisoryStateChangedEvent(ErrataBaseEvent):
    """
    Represents change of Errata Advisory status.
    """


class ErrataAdvisoryRPMsSignedEvent(ErrataBaseEvent):
    """
    Event when all RPMs in Errata advisory are signed.
    """


class BrewSignRPMEvent(BaseEvent):
    """
    Represents the message sent by Brew when RPM is signed.
    """
    def __init__(self, msg_id, nvr, **kwargs):
        super(BrewSignRPMEvent, self).__init__(msg_id, **kwargs)
        self.nvr = nvr

    @property
    def search_key(self):
        return str(self.nvr)


class BrewContainerTaskStateChangeEvent(BaseEvent):
    """
    Represents the message sent by Brew when a container task state is changed.
    """
    def __init__(self, msg_id, container, branch, target, task_id, old_state,
                 new_state, **kwargs):
        super(BrewContainerTaskStateChangeEvent, self).__init__(msg_id, **kwargs)
        self.container = container
        self.branch = branch
        self.target = target
        self.task_id = task_id
        self.old_state = old_state
        self.new_state = new_state

    @property
    def search_key(self):
        return str(self.task_id)


class ODCSComposeStateChangeEvent(BaseEvent):
    """Represent a compose' state change event from ODCS"""

    def __init__(self, msg_id, compose, **kwargs):
        super(ODCSComposeStateChangeEvent, self).__init__(msg_id, **kwargs)
        self.compose = compose


class FreshmakerManualRebuildEvent(BaseEvent):
    def __init__(self, msg_id, errata_id=None, dry_run=False):
        super(FreshmakerManualRebuildEvent, self).__init__(
            msg_id, dry_run=dry_run)
        self.errata_id = errata_id
