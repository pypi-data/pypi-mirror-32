import pygit2
import re
import subprocess

from git_deps.errors import InvalidCommitish
from git_deps.utils import abort


class GitUtils(object):
    @classmethod
    def abbreviate_sha1(cls, sha1):
        """Uniquely abbreviates the given SHA1."""

        # For now we invoke git-rev-parse(1), but hopefully eventually
        # we will be able to do this via pygit2.
        cmd = ['git', 'rev-parse', '--short', sha1]
        # cls.logger.debug(" ".join(cmd))
        out = subprocess.check_output(cmd).strip()
        # cls.logger.debug(out)
        return out

    @classmethod
    def describe(cls, sha1):
        """Returns a human-readable representation of the given SHA1."""

        # For now we invoke git-describe(1), but eventually we will be
        # able to do this via pygit2, since libgit2 already provides
        # an API for this:
        #   https://github.com/libgit2/pygit2/pull/459#issuecomment-68866929
        #   https://github.com/libgit2/libgit2/pull/2592
        cmd = [
            'git', 'describe',
            '--all',       # look for tags and branches
            '--long',      # remotes/github/master-0-g2b6d591
            # '--contains',
            # '--abbrev',
            sha1
        ]
        # cls.logger.debug(" ".join(cmd))
        out = None
        try:
            out = subprocess.check_output(cmd, stderr=subprocess.STDOUT)
        except subprocess.CalledProcessError as e:
            if e.output.find('No tags can describe') != -1:
                return ''
            raise

        out = out.strip()
        out = re.sub(r'^(heads|tags|remotes)/', '', out)
        # We already have the abbreviated SHA1 from abbreviate_sha1()
        out = re.sub(r'-g[0-9a-f]{7,}$', '', out)
        # cls.logger.debug(out)
        return out

    @classmethod
    def oneline(cls, commit):
        return commit.message.split('\n', 1)[0]

    @classmethod
    def commit_summary(cls, commit):
        return "%s %s" % (commit.hex[:8], cls.oneline(commit))

    @classmethod
    def refs_to(cls, sha1, repo):
        """Returns all refs pointing to the given SHA1."""
        matching = []
        for refname in repo.listall_references():
            symref = repo.lookup_reference(refname)
            dref = symref.resolve()
            oid = dref.target
            commit = repo.get(oid)
            if commit.hex == sha1:
                matching.append(symref.shorthand)

        return matching

    @classmethod
    def rev_list(cls, rev_range):
        cmd = ['git', 'rev-list', rev_range]
        return subprocess.check_output(cmd).strip().split('\n')

    @classmethod
    def ref_commit(cls, repo, rev):
        try:
            commit = repo.revparse_single(rev)
        except (KeyError, ValueError):
            raise InvalidCommitish(rev)

        if isinstance(commit, pygit2.Tag):
            commit = commit.get_object()

        return commit

    @classmethod
    def get_repo(cls, path='.'):
        try:
            repo_path = pygit2.discover_repository(path)
        except KeyError:
            abort("Couldn't find a repository in the current directory.")

        return pygit2.Repository(repo_path)
