# -*- coding: utf-8 -*-
# Copyright (c) 2017  Red Hat, Inc.
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

from freshmaker.parsers import BaseParser
from freshmaker.events import FreshmakerManualRebuildEvent


class FreshmakerManualRebuildParser(BaseParser):
    """Parser parsing odcs.compose.state.change"""

    name = "FreshmakerManualRebuildEvent"
    topic_suffixes = ["freshmaker.manual.rebuild"]

    def can_parse(self, topic, msg):
        return any([topic.endswith(s) for s in self.topic_suffixes])

    def parse(self, topic, msg):
        msg_id = msg.get('msg_id')
        inner_msg = msg.get('msg')
        errata_id = inner_msg.get('errata_id')
        dry_run = inner_msg.get('dry_run', False)

        return FreshmakerManualRebuildEvent(
            msg_id, errata_id=errata_id, dry_run=dry_run)
