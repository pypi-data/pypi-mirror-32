from manga_py.provider import Provider
from .helpers.std import Std


class AllHentaiRu(Provider, Std):

    def get_archive_name(self):
        name = self.re.search('/.+/([^/]+/[^/]+)/?', self.chapter)
        return name.group(1).replace('/', '-')

    def get_chapter_index(self):
        name = self.re.search('/.+/(?:vol)?([^/]+/[^/]+)/?', self.chapter)
        return name.group(1).replace('/', '-')

    def get_main_content(self):
        return self._get_content('{}/{}?mature=1&mtr=1')

    def get_manga_name(self) -> str:
        return self._get_name(r'\.ru/([^/]+)')

    def get_chapters(self):
        return self._elements('.expandable .cTable tr > td > a')

    def get_files(self):
        _uri = self.http().normalize_uri(self.chapter)
        content = self.http_get(_uri)
        result = self.re.search(r'var pictures.+?(\[\{.+\}\])', content, self.re.M)
        if not result:
            return []
        content = result.group(1).replace("'", '"')
        content = self.re.sub('(\w*):([^/])', r'"\1":\2', content)
        return [i['url'] for i in self.json.loads(content)]

    def get_cover(self):
        return self._cover_from_content('.picture-fotorama > img')

    def save_file(self, idx=None, callback=None, url=None, in_arc_name=None):
        _path = None
        try:
            _path = super().save_file(idx, callback, url, in_arc_name)
        except AttributeError:
            pass
        if _path is None:
            for i in ['a', 'b', 'c']:
                try:
                    _path, idx, _url = self._save_file_params_helper(url, idx)
                    _url = self.re.sub(r'//\w\.', '//%s.' % i, url)

                    self.http().download_file(_url, _path)
                    self._archive.add_file(_path, in_arc_name)

                    callable(callback) and callback()
                    self.after_file_save(_path, idx)
                    break
                except AttributeError:
                    pass
        return _path


main = AllHentaiRu
