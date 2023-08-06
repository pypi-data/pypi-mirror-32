import re
from io import StringIO
from pathlib import Path
from subprocess import run, CalledProcessError, PIPE, STDOUT

from foliant.preprocessors.base import BasePreprocessor


class Preprocessor(BasePreprocessor):
    defaults = {
        'recursive': True,
        'cache_dir': Path('.includescache'),
        'aliases': {}
    }
    tags = 'include',

    _heading_pattern = re.compile(
        r'^(?P<hashes>\#+)\s*(?P<title>[^\#].+\S+)\s*$',
        flags=re.MULTILINE
    )
    _image_pattern = re.compile(r'\!\[(?P<caption>.*)\]\((?P<path>((?!:\/\/).)+)\)')
    _tag_body_pattern = re.compile(
        r'(\$(?P<repo>[^\#^\$]+)(#(?P<revision>[^\$]+))?\$)?' +
        r'(?P<path>[^\#]+)' +
        r'(\#(?P<from_heading>[^:]*)(:(?P<to_heading>.+))?)?'
    )

    @staticmethod
    def _find_file(file_name: str, lookup_dir: Path) -> Path or None:
        '''Find a file in a directory by name. Check subdirectories recursively.

        :param file_name: Name of the file
        :lookup_dir: Starting directory

        :returns: Path to the found file or None if the file was not found
        :raises: FileNotFoundError
        '''

        self.logger.debug('Trying to find the file {file_name} inside the directory {lookup_dir}')

        result = None

        for item in lookup_dir.rglob('*'):
            if item.name == file_name:
                result = item
                break
        else:
            raise FileNotFoundError(file_name)

        self.logger.debug('File found: {result}')

        return result

    def _sync_repo(self, repo_url: str, revision: str or None = None) -> Path:
        '''Clone a Git repository to the cache dir. If it has been cloned before, update it.

        :param repo_url: Repository URL
        :param revision: Revision: branch, commit hash, or tag

        :returns: Path to the cloned repository
        '''

        repo_name = repo_url.split('/')[-1].rsplit('.', maxsplit=1)[0]
        repo_path = self._cache_path / repo_name

        self.logger.debug(f'Synchronizing with repo; URL: {repo_url}, revision: {revision}')

        try:
            self.logger.debug(f'Cloning repo {repo_url} to {repo_path}')

            run(
                f'git clone {repo_url} {repo_path}',
                shell=True,
                check=True,
                stdout=PIPE,
                stderr=STDOUT
            )

        except CalledProcessError as exception:
            if repo_path.exists():
                self.logger.debug('Repo already cloned; pulling from remote')

                run('git pull', cwd=repo_path, shell=True, check=True, stdout=PIPE, stderr=STDOUT)

            else:
                self.logger.error(str(exception))

        if revision:
            run(
                f'git checkout {revision}',
                cwd=repo_path,
                shell=True,
                check=True,
                stdout=PIPE,
                stderr=STDOUT
            )

        return repo_path

    def _shift_headings(self, content: str, shift: int) -> str:
        '''Shift Markdown headings in a string by a given value. The shift can
        be positive or negative.

        :param content: Markdown content
        :param shift: Heading shift

        :returns: Markdown content with headings shifted by ``shift``
        '''

        def _sub(heading):
            new_heading_level = len(heading.group('hashes')) + shift

            self.logger.debug(f'Shift heading level to {new_heading_level}')

            return f'{"#" * new_heading_level} {heading.group("title")}'

        return self._heading_pattern.sub(_sub, content)

    def _find_top_heading_level(self, content: str) -> int:
        '''Find the highest level heading (i.e. having the least '#'s)
        in a Markdown string.

        :param content: Markdown content

        :returns: Maximum heading level detected; if no heading is found, 0 is returned
        '''

        result = float('inf')

        for heading in self._heading_pattern.findall(content):
            heading_level = len(self._heading_pattern.match(heading).group('hashes'))

            if heading_level < result:
                result = heading_level

            self.logger.debug(f'Maximum heading level: {result}')

        return result if result < float('inf') else 0

    def _cut_from_heading_to_heading(
            self,
            content: str,
            from_heading: str,
            to_heading: str or None = None,
            options={}
        ) -> str:
        '''Cut part of Markdown string between two headings, set internal heading level,
        and remove top heading.

        If only the starting heading is defined, cut to the next heading
        of the same level.

        Heading shift and top heading elimination are optional.

        :param content: Markdown content
        :param from_heading: Starting heading
        :param to_heading: Ending heading (will not be incuded in the output)
        :param options: ``sethead``, ``nohead``

        :returns: Part of the Markdown content between headings with internal headings adjusted
        '''

        self.logger.debug(f'Cutting from heading: {from_heading}, to heading: {to_heading}, options: {options}')

        from_heading_pattern = re.compile(rf'^\#+\s*{from_heading}\s*$', flags=re.MULTILINE)

        if not from_heading_pattern.findall(content):
            return ''

        from_heading_line = from_heading_pattern.findall(content)[0]
        from_heading_level = len(self._heading_pattern.match(from_heading_line).group('hashes'))

        self.logger.debug(f'From heading level: {from_heading_level}')

        result = from_heading_pattern.split(content)[1]

        if to_heading:
            to_heading_pattern = re.compile(rf'^\#+\s*{to_heading}\s*$', flags=re.MULTILINE)

        else:
            to_heading_pattern = re.compile(
                rf'^\#{{1,{from_heading_level}}}[^\#]+?$',
                flags=re.MULTILINE
            )

        result = to_heading_pattern.split(result)[0]

        if not options.get('nohead'):
            result = from_heading_line + result

        if options.get('sethead'):
            if options['sethead'] > 0:
                result = self._shift_headings(
                    result,
                    options['sethead'] - from_heading_level
                )

        return result

    def _cut_to_heading(
            self,
            content: str,
            to_heading: str or None = None,
            options={}
        ) -> str:
        '''Cut part of Markdown string from the start to a certain heading,
        set internal heading level, and remove top heading.

        If not heading is defined, the whole string is returned.

        Heading shift and top heading elimination are optional.

        :param content: Markdown content
        :param to_heading: Ending heading (will not be incuded in the output)
        :param options: ``sethead``, ``nohead``

        :returns: Part of the Markdown content from the start to ``to_heading``,
            with internal headings adjusted
        '''

        self.logger.debug(f'Cutting to heading: {to_heading}, options: {options}')

        content_buffer = StringIO(content)

        first_line = content_buffer.readline()

        if self._heading_pattern.fullmatch(first_line):
            from_heading_line = first_line
            from_heading_level = len(self._heading_pattern.match(from_heading_line).group('hashes'))
            result = content_buffer.read()

        else:
            from_heading_line = ''
            from_heading_level = self._find_top_heading_level(content)
            result = content

        self.logger.debug(f'From heading level: {from_heading_level}')

        if to_heading:
            to_heading_pattern = re.compile(rf'^\#+\s*{to_heading}\s*$', flags=re.MULTILINE)
            result = to_heading_pattern.split(result)[0]

        if not options.get('nohead'):
            result = from_heading_line + result

        if options.get('sethead'):
            if options['sethead'] > 0:
                result = self._shift_headings(
                    result,
                    options['sethead'] - from_heading_level
                )

        return result

    def _adjust_image_paths(self, content: str, md_file_path: Path) -> str:
        '''Locate images referenced in a Markdown string and replace their paths
        with the absolute ones.

        :param content: Markdown content
        :param md_file_path: Path to the Markdown file containing the content

        :returns: Markdown content with absolute image paths
        '''

        def _sub(image):
            image_caption = image.group('caption')
            image_path = md_file_path.parent / Path(image.group('path'))

            self.logger.debug(
                f'Updating image reference; user specified path: {image.group("path")}, ' +
                f'absolute path: {image_path}, caption: {image_caption}'
            )

            return f'![{image_caption}]({image_path.absolute().as_posix()})'

        return self._image_pattern.sub(_sub, content)

    def _get_local_file_path(self, user_specified_path: str) -> Path:
        '''Resolve user specified path to a local file.

        :param user_specified_path: User specified string that represents
            the path to a local file

        :returns: Local file path relative to the temporary working directory,
            if the file is located inside it, or to the project source directory,
            if the file is not located inside the temporary working directory
        '''

        self.logger.debug(f'Current directory: {self._current_dir}')

        try:
            path_relative_to_current_dir = (self._current_dir/user_specified_path).relative_to(self._current_dir)

        except ValueError:
            path_relative_to_current_dir = Path(user_specified_path)

        self.logger.debug(f'Path relative to the current directory: {path_relative_to_current_dir}')

        if self.working_dir.resolve() in (self._current_dir/user_specified_path).resolve().parents:
            resolved_file_path = self._current_dir/path_relative_to_current_dir

        else:
            resolved_file_path = (
                self.config['src_dir']/
                self._current_dir.relative_to(self.working_dir)/
                path_relative_to_current_dir
            )

        self.logger.debug(f'Resolved path: {resolved_file_path}')

        return resolved_file_path

    def _process_include(
            self,
            file_path: Path,
            from_heading: str or None = None,
            to_heading: str or None = None,
            options={}
        ) -> str:
        '''Replace a local include statement with the file content. Necessary
        adjustments are applied to the content: cut between certain headings,
        strip the top heading, set heading level.

        :param file_path: Path to the included file
        :param from_heading: Include starting from this heading
        :param to_heading: Include up to this heading (not including the heading itself)
        :param options: ``sethead``, ``nohead``

        :returns: Included file content
        '''

        self.logger.debug(
            f'File path: {file_path}, from heading: {from_heading}, ' +
            f'to heading: {to_heading}, options: {options}'
        )

        if file_path.name.startswith('^'):
            file_path = self._find_file(file_path.name[1:], file_path.parent)

        with open(file_path, encoding='utf8') as incl_file:
            incl_content = incl_file.read()

            if from_heading:
                incl_content = self._cut_from_heading_to_heading(
                    incl_content,
                    from_heading,
                    to_heading,
                    options
                )

            else:
                incl_content = self._cut_to_heading(
                    incl_content,
                    to_heading,
                    options
                )

            incl_content = self._adjust_image_paths(incl_content, file_path)

        return incl_content

    def process_includes(self, content: str) -> str:
        '''Replace all include statements with the respective file contents.

        :param content: Markdown content

        :returns: Markdown content with resolved includes
        '''

        def _sub(include):
            body = self._tag_body_pattern.match(include.group('body').strip())
            options = self.get_options(include.group('options'))

            self.logger.debug(f'Processing include statement; body: {body}, options: {options}')

            if body.group('repo'):
                repo = body.group('repo')
                repo_url = self.options['aliases'].get(repo) or repo
                repo_path = self._sync_repo(repo_url, body.group('revision'))

                self.logger.debug(f'File in Git repository referenced; URL: {repo_url}, path: {repo_path}')

                return self._process_include(
                    repo_path/body.group('path'),
                    body.group('from_heading'),
                    body.group('to_heading'),
                    options
                )

            else:
                self.logger.debug('Local file referenced')

                return self._process_include(
                    self._get_local_file_path(body.group('path')),
                    body.group('from_heading'),
                    body.group('to_heading'),
                    options
                )

        result = self.pattern.sub(_sub, content)

        if self.options['recursive'] and self.pattern.search(result):
            return self.process_includes(result)
        else:
            return result

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        self._cache_path = self.project_path / self.options['cache_dir']
        self._current_dir = self.working_dir


        self.logger = self.logger.getChild('includes')

        self.logger.debug(f'Preprocessor inited: {self.__dict__}')


    def apply(self):
        self.logger.info('Applying preprocessor')

        for markdown_file_path in self.working_dir.rglob('*.md'):
            with open(markdown_file_path, encoding='utf8') as markdown_file:
                content = markdown_file.read()

            with open(markdown_file_path, 'w', encoding='utf8') as markdown_file:
                self._current_dir = markdown_file_path.parent
                markdown_file.write(self.process_includes(content))

        self.logger.info('Preprocessor applied')
