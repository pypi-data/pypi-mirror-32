#!/usr/bin/env python3
import functools
import pathlib
import re
import sys
from glob import glob

from setuptools import Command, find_packages, setup

try:
    from pip.req import parse_requirements
except ImportError:  # pip >= 10.0.0
    from pip._internal.req import parse_requirements

WORK_DIR = pathlib.Path(__file__).parent

# Check python version
MINIMAL_PY_VERSION = (3, 6)
if sys.version_info < MINIMAL_PY_VERSION:
    raise RuntimeError('PyMTProxy works only with Python {}+'.format('.'.join(map(str, MINIMAL_PY_VERSION))))


@functools.lru_cache()
def get_version():
    """
    Read version

    :return: str
    """
    txt = (WORK_DIR / 'mtproxy' / '__init__.py').read_text('utf-8')
    try:
        return re.findall(r"^__version__ = '([^']+)'\r?$", txt, re.M)[0]
    except IndexError:
        raise RuntimeError('Unable to determine version.')


def get_description():
    """
    Read full description from 'README.rst'

    :return: description
    :rtype: str
    """
    with open('README.rst', 'r', encoding='utf-8') as f:
        return f.read()


@functools.lru_cache()
def get_requirements(filename=None):
    """
    Read requirements from 'requirements txt'

    :return: requirements
    :rtype: list
    """
    if filename is None:
        filename = 'requirements.txt'

    file = WORK_DIR / filename

    install_reqs = parse_requirements(str(file), session='hack')
    return [str(ir.req) for ir in install_reqs]


class UploadCommand(Command):
    command_name = 'upload'
    description = 'upload the distribution with the Python package index'

    user_options = [
        ('set-tag', 't', 'set Git tag')
    ]

    def initialize_options(self):
        self.path = WORK_DIR / 'dist'
        self.lib_version = get_version()
        self.set_tag = False

    def finalize_options(self):
        opts = self.distribution.get_option_dict(self.get_command_name())

        self.set_tag = 'set_tag' in opts

    def find_target(self):
        if not self.path.is_dir():
            return

        targets = []
        for file in glob(str(self.path / f"PyMTProxy-{self.lib_version}[-.]*")):
            targets.append(str((self.path / file).absolute()))
        return targets

    def upload(self, targets):
        import subprocess
        errno = subprocess.call(['twine', 'upload'] + targets)
        return errno

    def create_tag(self):
        import subprocess
        errno = subprocess.call(['git', 'tag', f"v{self.lib_version}"])
        return errno

    def run(self):
        targets = self.find_target()
        if targets:
            errno = self.upload(targets)
            if not errno and self.set_tag:
                errno = self.create_tag()
            raise SystemExit(errno)
        else:
            raise FileNotFoundError('No files to be uploaded!')


setup(
    name='PyMTProxy',
    version=get_version(),
    packages=find_packages(exclude=('tests', 'tests.*', 'examples.*', 'docs',)),
    # url='https://github.com/_/_',
    license='MIT',
    requires_python='>=3.6',
    author='Alex Root Junior',
    author_email='support@illemius.xyz',
    maintainer=', '.join((
        'Alex Root Junior <jroot.junior@gmail.com>',
    )),
    maintainer_email='support@illemius.xyz',
    description='MTProto Proxy server',
    long_description=get_description(),
    classifiers=[
        'Development Status :: 1 - Planning',
        'Topic :: Internet :: Proxy Servers',
        'Environment :: Console',
        'Framework :: AsyncIO',
        'Intended Audience :: System Administrators',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.6',
    ],
    install_requires=get_requirements(),
    tests_require=get_requirements('dev_requirements.txt'),
    extras_require={
        'dev': get_requirements('dev_requirements.txt')
    },
    cmdclass={
        UploadCommand.command_name: UploadCommand
    }
)
