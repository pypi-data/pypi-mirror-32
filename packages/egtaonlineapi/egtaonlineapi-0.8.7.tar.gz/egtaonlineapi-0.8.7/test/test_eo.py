"""Test cli"""
import asyncio
import contextlib
import io
import json
import sys
import traceback
from unittest import mock

import pytest

from egtaonline import __main__ as main
from egtaonline import api
from egtaonline import mockserver


# TODO async fixtures may be possible with python 3.6, but it's not possible
# with async_generator


async def run(*args):
    """Run a command line and return if it ran successfully"""
    try:
        await main.amain(*args)
    except SystemExit as ex:
        return not int(str(ex))
    except Exception: # pylint: disable=broad-except
        traceback.print_exc()
        return False
    return True


def stdin(inp):
    """Patch stdin with input"""
    return mock.patch.object(sys, 'stdin', io.StringIO(inp))


# This is a hack to allow "writing" to the underlying buffer of a stringio
class _StringBytes(io.BytesIO):
    """A wrapper for bytes io that allows getting the string

    This is necessary because for zip files, the result needs to be written to
    a byte stream."""
    def close(self):
        pass

    def getvalue(self):
        return super().getvalue().decode('utf8')


@contextlib.contextmanager
def stdout():
    """Patch stdout and return stringio"""
    buff = _StringBytes()
    with mock.patch.object(sys, 'stdout', io.TextIOWrapper(buff)):
        yield buff


def stderr():
    """Patch stderr and return stringio"""
    return mock.patch.object(sys, 'stderr', io.StringIO())


@pytest.mark.asyncio
async def test_help():
    """Test getting help by itself"""
    with stderr() as err:
        assert await run('-h'), err.getvalue()


@pytest.mark.asyncio
@pytest.mark.parametrize('cmd', ['sim', 'game', 'sched', 'sims'])
async def test_cmd_help(cmd):
    """Test getting help from commands"""
    with stderr() as err:
        assert await run(cmd, '-h'), err.getvalue()


@pytest.mark.asyncio
async def test_sim():
    """Test sim functionality"""
    async with mockserver.server() as server:
        with stdout() as out, stderr() as err:
            assert await run('sim'), err.getvalue()
        assert not out.getvalue()

        server.create_simulator('sim', '1')
        with stdout() as out, stderr() as err:
            assert await run('sim'), err.getvalue()

        # get by id
        sim = json.loads(out.getvalue())
        with stderr() as err:
            assert await run('sim', str(sim['id'])), err.getvalue()

        # get by name
        with stdout() as out, stderr() as err:
            assert await run(
                'sim', sim['name'], '-n', sim['version']), err.getvalue()
        assert sim['id'] == json.loads(out.getvalue())['id']

        assert not await run('sim', '--', '-1')

        # add role
        with stderr() as err:
            assert await run('sim', str(sim['id']), '-rr'), err.getvalue()

        # add strategy
        with stderr() as err:
            assert await run(
                'sim', str(sim['id']), '-rr', '-ss'), err.getvalue()

        # add strategies
        with stdin(json.dumps({'r': ['q'], 'a': ['b']})), stderr() as err:
            assert await run('sim', str(sim['id']), '-j-'), err.getvalue()

        # remove strategy
        with stderr() as err:
            assert await run(
                'sim', str(sim['id']), '-drr', '-sq'), err.getvalue()

        # remove role
        with stderr() as err:
            assert await run('sim', str(sim['id']), '-dra'), err.getvalue()

        # remove strategies
        with stdin(json.dumps({'r': ['s']})), stderr() as err:
            assert await run('sim', str(sim['id']), '-dj-'), err.getvalue()

        # get zip
        with stdout() as out, stderr() as err:
            assert await run('sim', str(sim['id']), '-z'), err.getvalue()


@pytest.mark.asyncio
async def test_game(tmpdir): # pylint: disable=too-many-statements
    """Test game functionality"""
    conf = str(tmpdir.join('conf.json'))
    with open(conf, 'w') as fil:
        json.dump({}, fil)

    async with mockserver.server() as server:
        with stdout() as out, stderr() as err:
            assert await run('game'), err.getvalue()
        assert not out.getvalue()

        sim_id = server.create_simulator('sim', '1')
        game_spec = {
            'players': {
                'r': 2,
            },
            'strategies': {
                'r': ['s0', 's1'],
            },
        }
        with stdin(json.dumps(game_spec['strategies'])), stderr() as err:
            assert await run('sim', str(sim_id), '-j-'), err.getvalue()

        # get canon game
        with stdin(json.dumps(game_spec)), stdout() as out, \
                stderr() as err:
            assert await run(
                'game', str(sim_id), '-j-', '--fetch-conf',
                conf), err.getvalue()
        game = json.loads(out.getvalue())

        # verify its now listed with games
        with stdout() as out, stderr() as err:
            assert await run('game'), err.getvalue()
        game2 = json.loads(out.getvalue())
        assert game == game2

        # get game structure
        with stdout() as out, stderr() as err:
            assert await run('game', str(game['id'])), err.getvalue()
        struct = json.loads(out.getvalue())

        with stdin(json.dumps(game_spec)), stdout() as out, \
                stderr() as err:
            assert await run(
                'game', str(sim_id), '-j-', '--fetch-conf',
                conf), err.getvalue()
        assert struct == json.loads(out.getvalue())

        # get game summary
        with stdout() as out, stderr() as err:
            assert await run(
                'game', str(game['id']), '--summary'), err.getvalue()
        summ = json.loads(out.getvalue())

        with stdin(json.dumps(game_spec)), stdout() as out, \
                stderr() as err:
            assert await run(
                'game', str(sim_id), '-j-', '--fetch-conf',
                conf, '--summary'), err.getvalue()
        assert summ == json.loads(out.getvalue())

        # get observations
        with stderr() as err:
            assert await run(
                'game', str(game['id']), '--observations'), err.getvalue()
        obs = json.loads(out.getvalue())

        with stdin(json.dumps(game_spec)), stdout() as out, \
                stderr() as err:
            assert await run(
                'game', str(sim_id), '-j-', '--fetch-conf',
                conf, '--observations'), err.getvalue()
        assert obs == json.loads(out.getvalue())

        # get full data
        with stderr() as err:
            assert await run(
                'game', str(game['id']), '--full'), err.getvalue()
        full = json.loads(out.getvalue())

        with stdin(json.dumps(game_spec)), stdout() as out, \
                stderr() as err:
            assert await run(
                'game', str(sim_id), '-j-', '--fetch-conf',
                conf, '--full'), err.getvalue()
        assert full == json.loads(out.getvalue())

        # test name works
        with stdout() as out, stderr() as err:
            assert await run('game', game['name'], '-n'), err.getvalue()
        assert game['id'] == json.loads(out.getvalue())['id']

        # remove strategy
        with stderr() as err:
            assert await run(
                'game', str(game['id']), '-drr', '-ss0'), err.getvalue()

        # remove strategys
        with stdin(json.dumps({'r': ['s1']})), stderr() as err:
            assert await run(
                'game', str(game['id']), '-dj-'), err.getvalue()

        # remove role
        with stderr() as err:
            assert await run('game', str(game['id']), '-drr'), err.getvalue()

        # add role
        assert not await run('game', str(game['id']), '-rr')
        with stderr() as err:
            assert await run(
                'game', str(game['id']), '-rr', '-c2'), err.getvalue()

        # add strategies
        with stdin(json.dumps({'r': ['s1']})), stderr() as err:
            assert await run(
                'game', str(game['id']), '-j-'), err.getvalue()

        # add strategy
        with stderr() as err:
            assert await run(
                'game', str(game['id']), '-rr', '-ss0'), err.getvalue()


@pytest.mark.asyncio
async def test_sched():
    """Test scheduler functionality"""
    async with mockserver.server() as server:
        # verify no schedulers
        with stdout() as out, stderr() as err:
            assert await run('sched'), err.getvalue()
        assert not out.getvalue()

        # create one
        sim_id = server.create_simulator('sim', '1')
        async with api.api() as egta:
            await egta.create_generic_scheduler(
                sim_id, 'sched', True, 1, 2, 1, 1)

        with stdout() as out, stderr() as err:
            assert await run('sched'), err.getvalue()
        sched = json.loads(out.getvalue())

        with stderr() as err:
            assert await run('sched', str(sched['id'])), err.getvalue()

        with stderr() as err:
            assert await run('sched', str(sched['id']), '-r'), err.getvalue()

        with stdout() as out, stderr() as err:
            assert await run('sched', sched['name'], '-n'), err.getvalue()
        assert sched['id'] == json.loads(out.getvalue())['id']

        # deactivate scheduler
        with stderr() as err:
            assert await run(
                'sched', str(sched['id']), '--deactivate'), err.getvalue()

        # verify deactivated
        with stdout() as out, stderr() as err:
            assert await run('sched', str(sched['id'])), err.getvalue()
        assert not json.loads(out.getvalue())['active']

        # delete scheduler
        with stderr() as err:
            assert await run('sched', str(sched['id']), '-d'), err.getvalue()

        # assert no schedulers
        with stdout() as out, stderr() as err:
            assert await run('sched'), err.getvalue()
        assert not out.getvalue()


@pytest.mark.asyncio
async def test_sched_running():
    """Test running scheduler functionality"""
    async with mockserver.server() as server:
        # verify no schedulers
        with stdout() as out, stderr() as err:
            assert await run('sched'), err.getvalue()
        assert not out.getvalue()

        # create one
        sim_id = server.create_simulator('sim', '1', delay_dist=lambda: 1)
        async with api.api() as egta:
            sim = await egta.get_simulator(sim_id)
            await sim.add_strategies({'r': ['s0', 's1']})
            sched = await egta.create_generic_scheduler(
                sim_id, 'sched', True, 1, 2, 1, 1)
            await sched.add_role('r', 2)

            with stdout() as out, stderr() as err:
                assert await run('sched', '--running'), err.getvalue()
            assert not out.getvalue()

            await sched.add_profile('r: 1 s0, 1 s1', 1)

            with stdout() as out, stderr() as err:
                assert await run('sched', '--running'), err.getvalue()
            lines = out.getvalue()[:-1].split('\n')
            assert len(lines) == 1
            running = json.loads(lines[0])
            assert running['active']

            # Wait to complete
            await asyncio.sleep(1.5)
            with stdout() as out, stderr() as err:
                assert await run('sched', '--running'), err.getvalue()
            assert not out.getvalue()


@pytest.mark.asyncio
async def test_sims():
    """Test getting simulations"""
    async with mockserver.server() as server:
        with stdout() as out, stderr() as err:
            assert await run('sched'), err.getvalue()
        assert not out.getvalue()

        sim_id = server.create_simulator('sim', '1')
        async with api.api() as egta:
            sim = await egta.get_simulator(sim_id)
            await sim.add_strategies({'r': ['s0', 's1']})
            sched = await egta.create_generic_scheduler(
                sim_id, 'sched', True, 1, 2, 1, 1)
            await sched.add_role('r', 2)

            # This fails because we don't implement search, so this is as if no
            # job exists
            assert not await run('sims', '-j', '0')

            await sched.add_profile('r: 1 s0, 1 s1', 1)

            # This works because there's only one simulation so far and we
            # don't implement search
            with stdout() as out, stderr() as err:
                assert await run('sims', '-j', '0'), err.getvalue()

            await sched.add_profile('r: 2 s0', 2)

            # This fails because we don't implement search and now there are
            # several simulations
            assert not await run('sims', '-j', '0')

        with stdout() as out, stderr() as err:
            assert await run('sims'), err.getvalue()
        sims = [json.loads(line) for line in out.getvalue()[:-1].split('\n')]
        assert len(sims) == 3
        with stderr() as err:
            assert await run('sims', str(sims[0]['folder'])), err.getvalue()


@pytest.mark.asyncio
async def test_authfile():
    """Test supplying auth file"""
    async with mockserver.server():
        with stdin(''):
            assert await run('-f-', 'sim')
