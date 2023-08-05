# Author: echel0n <echel0n@sickrage.ca>
# URL: https://sickrage.ca
# Git: https://git.sickrage.ca/SiCKRAGE/sickrage
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

import cookielib
import datetime
import httplib
import json
import urllib
import urllib2

import MultipartPostHandler

import sickrage


class SabNZBd(object):
    @staticmethod
    def sendNZB(nzb):
        """
        Sends an NZB to SABnzbd via the API.

        :param nzb: The NZBSearchResult object to send to SAB
        """

        # set up a dict with the URL params in it
        params = {}
        if sickrage.app.config.sab_username is not None:
            params['ma_username'] = sickrage.app.config.sab_username
        if sickrage.app.config.sab_password is not None:
            params['ma_password'] = sickrage.app.config.sab_password
        if sickrage.app.config.sab_apikey is not None:
            params['apikey'] = sickrage.app.config.sab_apikey
        category = sickrage.app.config.sab_category
        if nzb.show.is_anime:
            category = sickrage.app.config.sab_category_anime

        # if it aired more than 7 days ago, override with the backlog category IDs
        for curEp in nzb.episodes:
            if datetime.date.today() - curEp.airdate > datetime.timedelta(days=7):
                category = sickrage.app.config.sab_category_backlog
                if nzb.show.is_anime:
                    category = sickrage.app.config.sab_category_anime_backlog

        if category is not None:
            params['cat'] = category

        # use high priority if specified (recently aired episode)
        if nzb.priority == 1:
            if sickrage.app.config.sab_forced == 1:
                params['priority'] = 2
            else:
                params['priority'] = 1

        try:
            f = None

            # if it's a normal result we just pass SAB the URL
            if nzb.resultType == "nzb":
                # for newzbin results send the ID to sab specifically
                if nzb.provider.id == 'newzbin':
                    id = nzb.provider.getIDFromURL(nzb.url)
                    if not id:
                        sickrage.app.log.error(
                            "Unable to send NZB to sab, can't find ID in URL " + str(nzb.url))
                        return False
                    params['mode'] = 'addid'
                    params['name'] = id
                else:
                    params['mode'] = 'addurl'
                    params['name'] = nzb.url

                url = sickrage.app.config.sab_host + "api?" + urllib.urlencode(params)
                sickrage.app.log.info("Sending NZB to SABnzbd")
                sickrage.app.log.debug("URL: " + url)

                # if we have the URL to an NZB then we've built up the SAB API URL already so just call it
                if nzb.resultType == "nzb":
                    f = urllib.urlopen(url)

            elif nzb.resultType == "nzbdata":
                params['mode'] = 'addfile'
                multiPartParams = {"nzbfile": (nzb.name + ".nzb", nzb.extraInfo[0])}

                url = sickrage.app.config.sab_host + "api?" + urllib.urlencode(params)
                sickrage.app.log.info("Sending NZB to SABnzbd")
                sickrage.app.log.debug("URL: " + url)

                cookies = cookielib.CookieJar()
                opener = urllib2.build_opener(urllib2.HTTPCookieProcessor(cookies),
                                              MultipartPostHandler.MultipartPostHandler)
                req = urllib2.Request(url, multiPartParams, headers={'User-Agent': sickrage.app.user_agent})
                f = opener.open(req)

        except (EOFError, IOError) as e:
            sickrage.app.log.error("Unable to connect to SAB: {}".format(e))
            return False
        except httplib.InvalidURL as e:
            sickrage.app.log.error("Invalid SAB host, check your config: {}".format(e))
            return False

        # this means we couldn't open the connection or something just as bad
        if f is None:
            sickrage.app.log.error("No data returned from SABnzbd, NZB not sent")
            return False

        # if we opened the URL connection then read the result from SAB
        try:
            result = f.readlines()
        except Exception as e:
            sickrage.app.log.error("Error trying to get result from SAB, NZB not sent: {}".format(e))
            return False

        # SAB shouldn't return a blank result, this most likely (but not always) means that it timed out and didn't recieve the NZB
        if len(result) == 0:
            sickrage.app.log.warning("No data returned from SABnzbd, NZB not sent")
            return False

        # massage the result a little bit
        sabText = result[0].strip()

        sickrage.app.log.debug("Result text from SAB: " + sabText)

        # do some crude parsing of the result text to determine what SAB said
        if sabText == "ok":
            sickrage.app.log.debug("NZB sent to SAB successfully")
            return True
        elif sabText == "Missing authentication":
            sickrage.app.log.warning("Incorrect username/password sent to SAB, NZB not sent")
            return False
        else:
            sickrage.app.log.warning("Unknown failure sending NZB to sab. Return text is: " + sabText)
            return False

    @staticmethod
    def _checkSabResponse(f):
        """
        Check response from SAB

        :param f: Response from SAV
        :return: a list of (Boolean, string) which is True if SAB is not reporting an error
        """
        try:
            result = f.readlines()
        except Exception as e:
            sickrage.app.log.error("Error trying to get result from SAB{}".format(e))
            return False, "Error from SAB"

        if len(result) == 0:
            sickrage.app.log.warning("No data returned from SABnzbd, NZB not sent")
            return False, "No data from SAB"

        sabText = result[0].strip()
        sabJson = {}
        try:
            sabJson = json.loads(sabText)
        except ValueError as e:
            pass

        if sabText == "Missing authentication":
            sickrage.app.log.warning("Incorrect username/password sent to SAB")
            return False, "Incorrect username/password sent to SAB"
        elif 'error' in sabJson:
            sickrage.app.log.warning(sabJson['error'])
            return False, sabJson['error']
        else:
            return True, sabText

    @staticmethod
    def _sabURLOpenSimple(url):
        """
        Open a connection to SAB

        :param url: URL where SAB is at
        :return: (boolean, string) list, True if connection can be made
        """
        try:
            f = urllib.urlopen(url)
        except (EOFError, IOError) as e:
            sickrage.app.log.error("Unable to connect to SAB: {}".format(e))
            return False, "Unable to connect"
        except httplib.InvalidURL as e:
            sickrage.app.log.error("Invalid SAB host, check your config: {}".format(e))
            return False, "Invalid SAB host"

        if f is None:
            sickrage.app.log.warning("No data returned from SABnzbd")
            return False, "No data returned from SABnzbd"
        else:
            return True, f

    @staticmethod
    def getSabAccesMethod(host=None, username=None, password=None, apikey=None):
        """
        Find out how we should connect to SAB

        :param host: hostname where SAB lives
        :param username: username to use
        :param password: password to use
        :param apikey: apikey to use
        :return: (boolean, string) with True if method was successful
        """
        url = host + "api?mode=auth"

        result, f = SabNZBd._sabURLOpenSimple(url)
        if not result:
            return False, f

        result, sabText = SabNZBd._checkSabResponse(f)
        if not result:
            return False, sabText

        return True, sabText

    @staticmethod
    def test_authentication(host=None, username=None, password=None, apikey=None):
        """
        Sends a simple API request to SAB to determine if the given connection information is connect

        :param host: The host where SAB is running (incl port)
        :param username: The username to use for the HTTP request
        :param password: The password to use for the HTTP request
        :param apikey: The API key to provide to SAB
        :return: A tuple containing the success boolean and a message
        """

        # build up the URL parameters
        params = {'mode': 'queue', 'output': 'json', 'ma_username': username, 'ma_password': password, 'apikey': apikey}
        url = host + "api?" + urllib.urlencode(params)

        # send the test request
        sickrage.app.log.debug("SABnzbd test URL: " + url)
        result, f = SabNZBd._sabURLOpenSimple(url)
        if not result:
            return False, f

        # check the result and determine if it's good or not
        result, sabText = SabNZBd._checkSabResponse(f)
        if not result:
            return False, sabText

        return True, "Success"
