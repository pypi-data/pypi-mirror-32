import os
import pytest
import shutil

from mock import Mock
from bumper.utils import PyPI

from workspace.scm import stat_repo, all_branches, commit_logs
from test_stubs import temp_dir, temp_git_repo, temp_remote_git_repo


@pytest.mark.parametrize('command,exception', [('diff', None), ('log', SystemExit), ('status', None)])
def test_sanity(wst, command, exception):
    with temp_dir():
        if exception:
            with pytest.raises(exception):
                wst(command)
        else:
            wst(command)

    with temp_git_repo():
        wst(command)


def test_bump(wst, monkeypatch):
    with temp_dir():
        with pytest.raises(SystemExit):
            wst('bump')

    with temp_remote_git_repo():
        # No requirements.txt
        if os.path.exists('requirements.txt'):
            os.unlink('requirements.txt')
        with pytest.raises(SystemExit):
            wst('bump')

        # All requirements are up to date
        with open('requirements.txt', 'w') as fp:
            fp.write('localconfig\nrequests')
        assert ({}, None, []) == wst('bump')

        # All requirements are outdated
        with open('requirements.txt', 'w') as fp:
            fp.write('# Comment for localconfig\nlocalconfig==0.0.1\n# Comment for requests\nrequests<0.1')
        msgs, commit_msg, bumps = wst('bump')
        file, msg = list(msgs.items())[0]
        assert 'requirements.txt' == file

        version = PyPI.latest_package_version('localconfig')
        expected_msg = 'Require localconfig==%s' % version
        assert expected_msg == msg[:len(expected_msg)]
        assert expected_msg == commit_msg[:len(expected_msg)]

        with open('requirements.txt') as fp:
            requirements = fp.read()
            assert '# Comment for localconfig\nlocalconfig==%s\n# Comment for requests\nrequests<0.1\n' % version == requirements


def test_cleanrun(wst):
    with temp_dir():
        wst('clean')


def test_commit(wst):
    with temp_dir():
        with pytest.raises(SystemExit):
            wst('commit')

    with temp_git_repo():
        with pytest.raises(SystemExit):
            wst('commit "no files to commit"')

        with open('new_file', 'w') as fp:
            fp.write('Hello World')
        assert 'new_file' in stat_repo(return_output=True)

        wst('commit "Add new file" --branch master')

        assert 'working tree clean' in stat_repo(return_output=True)
        assert 'Hello World' == open('new_file').read()

        with open('new_file', 'w') as fp:
            fp.write('New World')

        wst('commit "Update file"')

        assert ['update-file@master', 'master'] == all_branches()

        wst('commit --move release')

        assert ['update-file@master', 'master', 'release'] == all_branches()

        wst('commit --discard')

        assert ['master', 'release'] == all_branches()

        wst('checkout release')

        wst('commit --discard')

        assert ['release', 'master'] == all_branches()

        logs = commit_logs()
        assert 'new file' in logs
        assert 1 == len(list(filter(None, logs.split('commit'))))


def test_test(wst):
    with temp_dir():
        with pytest.raises(SystemExit):
            wst('test')

    if 'PYTESTARGS' in os.environ:
        del os.environ['PYTESTARGS']

    with temp_git_repo(name='foo'):
        with pytest.raises(SystemExit):
            wst('test')
        wst('setup --product')

        with open('foo/__init__.py', 'w') as fp:
            fp.write('hello = "world"')

        pass_test = 'from foo import hello\n\n\ndef test_pass():\n  assert hello == "world"'
        fail_test = 'def test_fail():\n  assert False'

        with open('tests/test_pass.py', 'w') as fp:
            fp.write(pass_test)
        commands = wst('test')
        assert 'test' in commands
        assert 'tox' in commands['test']

        with open('tests/test_fail.py', 'w') as fp:
            fp.write(pass_test + '\n' + fail_test)
        with pytest.raises(SystemExit):
            wst('test')

        output = wst('test tests/test_pass.py')
        assert 'test' in output

        os.utime('requirements.txt', None)
        assert 'test' in wst('test -k test_pass')

        with pytest.raises(SystemExit):
            wst('test style')
        with open('tests/test_fail.py', 'w') as fp:
            fp.write(pass_test + '\n\n\n' + fail_test)
        assert 'style' in wst('test style')

        os.unlink('tests/test_fail.py')
        assert 'cover' in wst('test cover')
        assert os.path.exists('coverage.xml')
        assert os.path.exists('htmlcov/index.html')

    shutil.rmtree(os.path.expanduser('~/.virtualenvs/foo'), ignore_errors=True)


def test_push_without_repo(wst):
    with temp_dir():
        with pytest.raises(SystemExit):
            wst('push')


def test_push(wst, monkeypatch):
    push_repo = Mock()
    remove_branch = Mock()
    monkeypatch.setattr('workspace.commands.push.push_repo', push_repo)
    monkeypatch.setattr('workspace.commands.push.remove_branch', remove_branch)

    with temp_remote_git_repo():
        wst('push')
        push_repo.assert_called_with(branch='master', force=False, remote='origin')

        with open('new_file', 'w') as fp:
            fp.write('Hello World')
        assert 'new_file' in stat_repo(return_output=True)

        wst('commit "Add new file"')

        wst('push')

        push_repo.assert_called_with(branch='add-new@master', force=False, remote='origin')

        assert ['add-new@master', 'master'] == all_branches()

        wst('push --merge')

        push_repo.assert_called_with(branch='master', force=False, remote='origin')
        remove_branch.assert_called_with('add-new@master', force=True, remote=True)

        # Remove local branch
        from workspace.scm import remove_branch as rb
        rb('add-new@master', force=True, remote=False)

        assert ['master'] == all_branches()

        assert "ahead of 'origin/master' by 1 commit." in stat_repo(return_output=True)
