# cerbero - a multi-platform build system for Open Source software
# Copyright (C) 2012 Andoni Morales Alastruey <ylatuya@gmail.com>
#
# This library is free software; you can redistribute it and/or
# modify it under the terms of the GNU Library General Public
# License as published by the Free Software Foundation; either
# version 2 of the License, or (at your option) any later version.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# Library General Public License for more details.
#
# You should have received a copy of the GNU Library General Public
# License along with this library; if not, write to the
# Free Software Foundation, Inc., 59 Temple Place - Suite 330,
# Boston, MA 02111-1307, USA.

import os
import shutil
import tarfile
import urllib.request, urllib.parse, urllib.error
import urllib.parse
import logging
import uuid
#from cerbero.config import Platform
from cam.utils import git, console ,shell#svn, shell, _
from cam.errors import FatalError

class Source (object):
    '''
    Base class for sources handlers

    '''

    patches = []

    def __init__(self,config):
        self.config   = config
        self.name     = config.name
        self.version  = config.version
        self.location = config.location
        self.cache_dir = config.cache_dir
        self.patches  = config.patches


class CustomSource (Source):

    def fetch(self):
        pass

    def extract(self):
        pass


ARCHIVE_FORMAT ={
    'tar':'tar',
    'tar.gz':'gztar',
    'tar.bz2': 'bztar',
    'zip':'zip',
    'tar.xz':'xztar'
}


class Tarball (Source):
    '''
    Source handler for tarballs

    @cvar url: dowload URL for the tarball
    @type url: str
    '''

    url = None
    tarball_name = None
    tarball_dirname = None
    def __init__(self, config):
        '''
        origin:
          tarball: http://github.com/xx/release/v1.2.3.tar.gz
          format: tar.gz|tar.bz2|tar.xz|zip (optional)

        '''
        Source.__init__(self,config)
        self.url = config.origin.tarball
        #self.format = ARCHIVE_FORMAT[config.origin.format]
        self.location = config.location
        self.dir = None
        if hasattr(config.origin,'dir'):
            self.dir = config.origin.dir


        format = getattr(config.origin,'format',None)
        if format is None:
            for ext in ARCHIVE_FORMAT.keys():
                if self.url.endswith('.'+ext):
                    format = ext
                    break
        if format not in ARCHIVE_FORMAT:
            raise FatalError('can guess the format of tarball %s'%self.url)

        self.format = ARCHIVE_FORMAT[format]

        self.tarball_name = '%s-%s.%s'%( config.name,config.version,format )

        self.download_path = os.path.join(config.cache_dir, self.tarball_name)

        # URL-encode spaces and other special characters in the URL's path
        split = list(urllib.parse.urlsplit(self.url))
        split[2] = urllib.parse.quote(split[2])
        self.url = urllib.parse.urlunsplit(split)

    def fetch(self, redownload=False):
        if not os.path.exists(self.config.cache_dir):
            os.makedirs(self.config.cache_dir)

        cached_file = os.path.join(self.config.cache_dir,self.tarball_name)
        if not redownload and os.path.isfile(cached_file):
            print('Use cached tarball %s to instead of %s' %
                     (cached_file, self.url))
            return
        #DO DOWNLOAD HTERE
        ## Enable certificate checking Linux for now
        ## FIXME: Add more platforms here after testing
        #cc = self.config.platform == Platform.LINUX
        shell.download(self.url, self.download_path, overwrite=redownload)


    def extract(self):
        extract_dir = self.location

        if self.dir:
            extract_dir = os.path.join('%s@%s'%(self.location,uuid.uuid1()))

        if os.path.exists(self.location):
            print("to del ",self.location)
            #shutil.rmtree(self.location)

        if os.path.exists(extract_dir):
            print("to del as extract_dir",self.location)
            #shutil.rmtree(extract_dir)

        try:
            shutil.unpack_archive(self.download_path, 
            extract_dir,self.format)
        except (IOError, tarfile.ReadError):
            logging.info('Corrupted or partial tarball, redownloading...')
            shell.download(self.url, self.download_path, overwrite=True)
            shutil.unpack_archive(self.download_path, 
            extract_dir,self.format)

        if self.dir:
            #os.rename(os.path.join(extract_dir, self.dir),self.location)
            shutil.move(os.path.join(extract_dir, self.dir),self.location)

        if extract_dir != self.location:
            if os.path.isdir(extract_dir):
                shutil.rmtree(extract_dir)

        git.init_directory(self.location)
        for patch in self.patches:
            if not os.path.isabs(patch):
                patch = os.path.abspath(os.path.join(self.config.__directory__, patch))
            git.apply_patch(patch, self.location)
        #    if self.strip == 1:
        #        git.apply_patch(patch, self.build_dir)
        #    else:
        #        shell.apply_patch(patch, self.build_dir, self.strip)


class GitCache (Source):
    '''
    Base class for source handlers using a Git repository
    '''

    remotes = None
    commit = None

    def __init__(self,config):
        Source.__init__(self,config)
        #if self.remotes is None:
        #    self.remotes = {}
        #if not 'origin' in self.remotes:
        #    self.remotes['origin'] = '%s/%s.git' % \
        #                             (self.config.git_root, self.name)
        self.repo_dir = os.path.join(self.cache_dir,
        '%s-%s@git'%(self.name,self.version)).replace('\\','/')
        
        self._previous_env = None

        self.url = config.origin.git
        self.commit = config.origin.commit

    def _git_env_setup(self):
        # When running git commands, which is the host git, we need to make
        # sure it is run in an environment which doesn't pick up the libraries
        # we build in cerbero
        env = os.environ.copy()
        self._previous_env = os.environ.copy()
        #env["LD_LIBRARY_PATH"] = self.config._pre_environ.get("LD_LIBRARY_PATH", "")
        os.environ = env

    def _git_env_restore(self):
        if self._previous_env is not None:
            os.environ = self._previous_env
            self._previous_env = None

    def fetch(self, checkout=True):
        self._git_env_setup()
        if not os.path.exists(self.repo_dir):
            git.init(self.repo_dir)


        git.add_remote(self.repo_dir, 'origin', self.url)
        # fetch remote branches
        git.fetch(self.repo_dir, fail=False)
        if checkout:
            commit = self.commit
            git.checkout(self.repo_dir, commit)
            git.submodules_update(self.repo_dir, src_dir=None, fail=False)
        self._git_env_restore()


    def built_version(self):
        return '%s+git~%s' % (self.version, git.get_hash(self.repo_dir, self.commit))


class LocalTarball (GitCache):
    '''
    Source handler for cerbero's local sources, a local git repository with
    the release tarball and a set of patches
    '''

    BRANCH_PREFIX = 'sdk'

    def __init__(self):
        GitCache.__init__(self)
        self.commit = "%s/%s-%s" % ('origin',
                                    self.BRANCH_PREFIX, self.version)
        self.platform_patches_dir = os.path.join(self.repo_dir,
                                                 self.config.platform)
        self.package_name = self.package_name
        self.unpack_dir = self.config.sources

    def extract(self):
        if not os.path.exists(self.build_dir):
            os.mkdir(self.build_dir)
        self._find_tarball()
        shell.unpack(self.tarball_path, self.unpack_dir)
        # apply common patches
        self._apply_patches(self.repo_dir)
        # apply platform patches
        self._apply_patches(self.platform_patches_dir)

    def _find_tarball(self):
        tarball = [x for x in os.listdir(self.repo_dir) if
                   x.startswith(self.package_name)]
        if len(tarball) != 1:
            raise FatalError(_("The local repository %s do not have a "
                             "valid tarball") % self.repo_dir)
        self.tarball_path = os.path.join(self.repo_dir, tarball[0])

    def _apply_patches(self, patches_dir):
        if not os.path.isdir(patches_dir):
            # FIXME: Add logs
            return

        # list patches in this directory
        patches = [os.path.join(patches_dir, x) for x in
                   os.listdir(patches_dir) if x.endswith('.patch')]
        # apply patches
        for patch in patches:
            shell.apply_patch(self.build_dir, patch)


class Git (GitCache):
    '''
    Source handler for git repositories
    '''

    def __init__(self,config):
        GitCache.__init__(self,config)
        #if self.commit is None:
        #    self.commit = 'origin/sdk-%s' % self.version
        ## For forced commits in the config
        #self.commit = self.config.recipe_commit(self.name) or self.commit

    def extract(self):
        if os.path.exists(self.location):
            try:
                commit_hash = git.get_hash(self.repo_dir, self.commit)
                checkout_hash = git.get_hash(self.build_dir, 'HEAD')
                if commit_hash == checkout_hash and not self.patches:
                    return False
            except Exception:
                pass
            shutil.rmtree(self.location)
        if not os.path.exists(self.location):
            os.mkdir(self.location)

        # checkout the current version
        git.local_checkout(self.location, self.repo_dir, self.commit)

        #for patch in self.patches:
        #    if not os.path.isabs(patch):
        #        patch = self.relative_path(patch)
#
        #    if self.strip == 1:
        #        git.apply_patch(patch, self.build_dir)
        #    else:
        #        shell.apply_patch(patch, self.build_dir, self.strip)

        return True


class GitExtractedTarball(Git):
    '''
    Source handle for git repositories with an extracted tarball

    Git doesn't conserve timestamps, which are reset after clonning the repo.
    This can confuse the autotools build system, producing innecessary calls
    to autoconf, aclocal, autoheaders or automake.
    For instance after doing './configure && make', 'configure' is called
    again if 'configure.ac' is newer than 'configure'.
    '''

    matches = ['.m4', '.in', 'configure']
    _files = {}

    def extract(self):
        if not Git.extract(self):
            return False
        for match in self.matches:
            self._files[match] = []
        self._find_files(self.build_dir)
        self._files['.in'] = [x for x in self._files['.in'] if
                os.path.join(self.build_dir, 'm4') not in x]
        self._fix_ts()

    def _fix_ts(self):
        for match in self.matches:
            for path in self._files[match]:
                shell.touch(path)

    def _find_files(self, directory):
        for path in os.listdir(directory):
            full_path = os.path.join(directory, path)
            if os.path.isdir(full_path):
                self._find_files(full_path)
            if path == 'configure.in':
                continue
            for match in self.matches:
                if path.endswith(match):
                    self._files[match].append(full_path)


class Svn(Source):
    '''
    Source handler for svn repositories
    '''

    url = None
    revision = 'HEAD'

    def __init__(self):
        Source.__init__(self)
        # For forced revision in the config
        self.revision = self.config.recipe_commit(self.name) or self.revision

    def fetch(self):
        if os.path.exists(self.repo_dir):
            shutil.rmtree(self.repo_dir)

        cached_dir = os.path.join(self.config.cached_sources, self.package_name)
        if os.path.isdir(os.path.join(cached_dir, ".svn")):
            m.action(_('Copying cached repo from %s to %s instead of %s') %
                     (cached_dir, self.repo_dir, self.url))
            shell.copy_dir(cached_dir, self.repo_dir)
            return

        os.makedirs(self.repo_dir)
        svn.checkout(self.url, self.repo_dir)
        svn.update(self.repo_dir, self.revision)

    def extract(self):
        if os.path.exists(self.build_dir):
            shutil.rmtree(self.build_dir)

        shutil.copytree(self.repo_dir, self.build_dir)

        for patch in self.patches:
            if not os.path.isabs(patch):
                patch = self.relative_path(patch)
            shell.apply_patch(patch, self.build_dir, self.strip)

    def built_version(self):
        return '%s+svn~%s' % (self.version, svn.revision(self.repo_dir))


class SourceType (object):

    CUSTOM = CustomSource
    TARBALL = Tarball
    LOCAL_TARBALL = LocalTarball
    GIT = Git
    GIT_TARBALL = GitExtractedTarball
    SVN = Svn
