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
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
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

# We have name conflict between two modules here:
#  - "odcs" module provided by python2-odcs-client
#  - "odcs" submodule in freshmaker.handlers.odcs
#
# Unfortunatelly we want to use "odcs" provided by python2-odcs-client
# in freshmaker.handlers __init__.py. We cannot  "import odcs" there, because
# it would import freshmaker.handlers.odcs, so instead, we import it here
# and in freshmaker.handler do "from freshmaker.odcsclient import ODCS".

from odcs.client.odcs import AuthMech, ODCS
from odcs.common.types import COMPOSE_STATES  # noqa

from freshmaker import conf


def create_odcs_client():
    """
    Create instance of ODCS according to configured authentication mechasnim
    """
    if conf.odcs_auth_mech == 'kerberos':
        return ODCS(conf.odcs_server_url,
                    auth_mech=AuthMech.Kerberos,
                    verify_ssl=conf.odcs_verify_ssl)
    elif conf.odcs_auth_mech == 'openidc':
        if not conf.odcs_openidc_token:
            raise ValueError('Missing OpenIDC token in configuration.')
        return ODCS(conf.odcs_server_url,
                    auth_mech=AuthMech.OpenIDC,
                    openidc_token=conf.odcs_openidc_token,
                    verify_ssl=conf.odcs_verify_ssl)
    else:
        raise ValueError(
            'Authentication mechanism {0} is not supported yet.'.format(
                conf.odcs_auth_mech))
