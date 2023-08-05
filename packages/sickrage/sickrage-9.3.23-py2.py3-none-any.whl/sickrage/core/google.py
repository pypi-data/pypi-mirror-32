# Author: echel0n <echel0n@sickrage.ca>
# URL: https://sickrage.ca
#
# This file is part of SickRage.
#
# SickRage is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# SickRage is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with SickRage.  If not, see <http://www.gnu.org/licenses/>.

from __future__ import unicode_literals

from oauth2client.client import OAuth2WebServerFlow, OAuth2Credentials

from sickrage.core.websession import WebSession


class GoogleAuth(object):
    def __init__(self):
        self.client_id = '48901323822-ebum0n1ago1bo2dku4mqm9l6kl2j60uv.apps.googleusercontent.com'
        self.client_secret = 'vFQy_bojwJ1f2X0hYD3wPu7U'
        self.scopes = ['https://www.googleapis.com/auth/drive.file',
                       'email',
                       'profile']

        self.credentials = None

        self.flow = OAuth2WebServerFlow(self.client_id, self.client_secret, ' '.join(self.scopes))

    def get_user_code(self):
        return self.flow.step1_get_device_and_user_codes()

    def get_credentials(self, flow_info):
        self.credentials = self.flow.step2_exchange(device_flow_info=flow_info)

        return self.credentials

    def refresh_credentials(self):
        if isinstance(self.credentials, OAuth2Credentials):
            self.credentials.refresh(WebSession())

    def logout(self):
        self.credentials = None
