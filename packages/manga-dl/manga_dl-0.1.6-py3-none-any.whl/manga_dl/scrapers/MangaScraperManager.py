"""LICENSE
Copyright 2015 Hermann Krumrey <hermann@krumreyh.com>

This file is part of manga-dl.

manga-dl is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

manga-dl is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with manga-dl.  If not, see <http://www.gnu.org/licenses/>.
LICENSE"""

from manga_dl.scrapers.GenericMangaScraper import GenericMangaScraper
from manga_dl.scrapers.MangaFoxScraper import MangaFoxScraper


class MangaScraperManager(object):
    """
    Class that acts as a negotiator for the various manga scrapers
    """

    scrapers = [MangaFoxScraper]
    """
    A list of scrapers that are implemented
    """

    @staticmethod
    def get_scraper_for(manga_url: str) -> type(GenericMangaScraper):
        """
        Returns the correct scraper for a specified manga URL
        :param manga_url: the URL of the Manga series
        :return: The correct scraper, or None if none was found
        """

        for scraper in MangaScraperManager.scrapers:
            if scraper.url_match(manga_url):
                return scraper
        return None

    @staticmethod
    def get_series_name_from_url(manga_url: str) -> str:
        """
        Tryes to figure out the name of a manga series from its URL

        :param manga_url: The URL to check
        :return:          The series name,
                          or an emtpy string if no applicable parser exists
        """
        scraper = MangaScraperManager.get_scraper_for(manga_url)
        return scraper.get_series_name(manga_url) if scraper is not None \
            else ""
