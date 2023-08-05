import requests
from ...factories import ZipfFromText


class ZipfFromUrl(ZipfFromText):
    def __init__(self, options=None):
        super().__init__(options)
        self._request_interface = lambda r: r.text

    def set_interface(self, request_interface):
        """Sets the interface with which read the text content of the file"""
        self._request_interface = request_interface

    def _download_file(self, url):
        r = requests.get(url)
        return self._request_interface(r)

    def run(self, url):
        """Loads and extracts a zipf from the given url"""
        return super().run(self._download_file(url))

    def enrich(self, url, zipf):
        """Loads and enriches a given zipf from the given url"""
        return super().enrich(self._download_file(url), zipf)
