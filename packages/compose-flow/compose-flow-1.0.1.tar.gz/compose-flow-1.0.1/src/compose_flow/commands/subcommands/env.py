"""
Env subcommand
"""
import io
import logging
import os
import shlex
import sys
import tempfile

import sh

from .base import BaseSubcommand

from compose_flow import docker, errors


class Env(BaseSubcommand):
    """
    Subcommand for managing environment
    """
    def __init__(self, *args, **kwargs):
        super().__init__(*args, *kwargs)

        self._config = None

    @classmethod
    def fill_subparser(cls, parser, subparser):
        subparser.add_argument('action')
        subparser.add_argument('path', nargs='*')

    def cat(self) -> str:
        """
        Prints the loaded config to stdout
        """
        if self.env_name not in docker.get_configs():
            return f'docker config named {self.env_name} not in swarm'

        print(self.render())

    @property
    def data(self) -> dict:
        """
        Returns the loaded config as a dictionary
        """
        data = {}

        env = self.load()
        for line in env.splitlines():
            # skip empty lines
            if line.strip() == '':
                continue

            # skip commented lines
            if line.startswith('#'):
                continue

            try:
                key, value = line.split('=', 1)
            except ValueError as exc:
                print(f'unable to split line={line}')
                raise

            data[key] = value

        # TODO: this is hacky because of the back-and-forth relationship
        # between data() and load() ... gotta fix this.
        docker_image = data.get('DOCKER_IMAGE')
        if 'VERSION' in data and docker_image and ':' in docker_image:
            data['DOCKER_IMAGE'] = f'{docker_image.split(":", 1)[0]}:{data["VERSION"]}'

        return data

    def edit(self) -> None:
        with tempfile.NamedTemporaryFile('w') as fh:
            path = fh.name

            self.render_buf(fh)

            fh.flush()

            editor = os.environ.get('EDITOR', os.environ.get('VISUAL', 'vi'))

            command = shlex.split(f'{editor} {path}')

            # os.execve(command[0], command, os.environ)
            proc = getattr(sh, command[0])
            proc(*command[1:], _env=os.environ, _fg=True)

            self.push(path)

    def is_dirty_working_copy_okay(self, exc):
        return self.workflow.args.action in ('cat', 'push')

    def is_env_error_okay(self, exc):
        return self.workflow.args.action in ('push',)

    def is_write_profile_error_okay(self, exc):
        return self.workflow.args.action in ('push',)

    def load(self) -> str:
        """
        Loads an environment from the docker swarm config
        """
        if self._config:
            return self._config

        self._config = docker.get_config(self.env_name)

        # inject the version from tag-version command into the loaded environment
        tag_version = 'unknown'
        try:
            tag_version_command = getattr(sh, 'tag-version')
        except Exception as exc:
            print(f'Warning: unable to find tag-version ({exc})\n', file=sys.stderr)
        else:
            try:
                tag_version = tag_version_command().stdout.decode('utf8').strip()
            except Exception as exc:
                # check if the subcommand is okay with a dirty working copy
                if not self.workflow.subcommand.is_dirty_working_copy_okay(exc):
                    raise errors.TagVersionError(f'Warning: unable to run tag-version ({exc})\n')

        data = self.data

        version_var = 'VERSION'
        data[version_var] = tag_version
        self._config = self.render(data)

        return self._config

    @property
    def logger(self):
        return logging.getLogger(f'{__name__}.{self.__class__.__name__}')

    def push(self, path:str=None) -> None:
        """
        Saves an environment into the swarm
        """
        path = path or self.args.path
        if not path:
            return self.print_subcommand_help(__doc__, error='path needed to load')

        docker.load_config(self.env_name, path)

    def render(self, data:dict=None) -> str:
        """
        Returns a rendered file in .env file format
        """
        buf = io.StringIO()

        self.render_buf(buf, data=data)

        return buf.getvalue()

    def render_buf(self, buf, data: dict=None):
        data = data or self.data
        for k, v in data.items():
            buf.write(f'{k}={v}\n')

    def rm(self) -> None:
        """
        Removes an environment from the swarm
        """
        docker.remove_config(self.env_name)

    def write_tag(self) -> None:
        """
        Writes the projects tag version into the environment

        This currently writes the tag to the `DOCKER_IMAGE` variable
        """
        data = self.data

        image_base = data['DOCKER_IMAGE'].rsplit(':', 1)[0]
        data['DOCKER_IMAGE'] = f'{image_base}:{data["VERSION"]}'

        with tempfile.NamedTemporaryFile('w+') as fh:
            fh.write(self.render(data))
            fh.flush()

            fh.seek(0, 0)

            self.push(path=fh.name)
