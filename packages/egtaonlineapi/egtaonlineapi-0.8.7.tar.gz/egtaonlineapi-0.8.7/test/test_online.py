"""Test against egta online"""
import asyncio
import collections

import pytest

from egtaonline import api
from egtaonline import mockserver


# TODO in python3.6 async fixtures may be able to be used, but async_generator
# doesn't work


class _fdict(dict):
    """Frozen dictionary"""
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.__hash = None

    def __hash__(self):
        if self.__hash is None:
            self.__hash = hash(frozenset(self.items()))
        return self.__hash


def describe_structure(obj, illegal=(), nums=False):
    """Compute an object that represents the recursive structure"""
    if isinstance(obj, dict): # pylint: disable=no-else-return
        return _fdict((k, describe_structure(v, illegal, nums))
                      for k, v in obj.items()
                      if k not in illegal)
    elif isinstance(obj, list):
        counts = collections.Counter(
            describe_structure(o, illegal, nums) for o in obj)
        return _fdict(counts)
    # NaNs are represented as None
    elif nums and isinstance(obj, (int, float, type(None))):
        return float
    else:
        return type(obj)


def assert_dicts_types(actual, expected, illegal=(), nums=False):
    """Test that dicts have the same type structure"""
    assert (describe_structure(actual, illegal, nums) ==
            describe_structure(expected, illegal, nums))


_ILLEGAL_KEYS = {'created_at', 'updated_at', 'simulator_instance_id'}


def assert_dicts_equal(actual, expected, illegal=()):
    """Test that dicts are equal"""
    assert actual.keys() == expected.keys(), \
        "keys weren't equal"
    assert ({k: v for k, v in actual.items()  # pragma: no branch
             if k not in _ILLEGAL_KEYS and k not in illegal} ==
            {k: v for k, v in expected.items()
             if k not in _ILLEGAL_KEYS and k not in illegal})


async def sched_complete(sched, sleep=0.001):
    """Wait until a scheduler is complete"""
    while (await sched.get_info())['active'] and not all(  # pragma: no branch
            p['requirement'] <= p['current_count'] for p
            in (await sched.get_requirements())['scheduling_requirements']):
        await asyncio.sleep(sleep)  # pragma: no cover


@pytest.mark.asyncio
@pytest.mark.egta
async def test_parity(): # pylint: disable=too-many-locals
    """Test that egta matches mock server"""
    # Get egta data
    async with api.api() as egta:
        true_sim = await egta.get_simulator(183)
        true_sched = await egta.get_scheduler(4997)
        true_game = await egta.get_game(2536)
        reqs = await true_sched.get_requirements()

        async def get_prof_info(prof):
            """Get all info about a profile"""
            return (await prof.get_structure(),
                    await prof.get_summary(),
                    await prof.get_observations(),
                    await prof.get_full_data())

        prof_info = await asyncio.gather(*[
            get_prof_info(prof)
            for prof in reqs['scheduling_requirements']])
        game_info = await true_game.get_structure()
        game_summ = await true_game.get_summary()

    # Replicate in mock
    async with mockserver.server() as server, api.api() as egta:
        for i in range(true_sim['id']):
            server.create_simulator('sim', str(i))
        mock_sim = await egta.get_simulator(server.create_simulator(
            true_sim['name'], true_sim['version'], true_sim['email'],
            true_sim['configuration']))
        await mock_sim.add_strategies(true_sim['role_configuration'])

        info = await mock_sim.get_info()
        assert_dicts_types(true_sim, info)
        assert_dicts_equal(true_sim, info)

        await asyncio.gather(*[
            mock_sim.create_generic_scheduler(str(i), False, 0, 0, 0, 0)
            for i in range(true_sched['id'])])
        mock_sched = await mock_sim.create_generic_scheduler(
            true_sched['name'], true_sched['active'],
            true_sched['process_memory'], true_sched['size'],
            true_sched['time_per_observation'],
            true_sched['observations_per_simulation'], true_sched['nodes'],
            dict(reqs['configuration']))

        info = await mock_sched.get_info()
        assert_dicts_types(true_sched, info)
        assert_dicts_equal(true_sched, info)

        game_sched = await mock_sim.create_generic_scheduler(
            'temp', True, 0, true_sched['size'], 0, 0, 1,
            dict(reqs['configuration']))

        counts = prof_info[0][0]['role_configuration']
        await mock_sched.add_roles(counts)
        await game_sched.add_roles(counts)

        await mock_sched.activate()
        for prof, (info, summ, obs, full) in zip(
                reqs['scheduling_requirements'], prof_info):
            sprof = await game_sched.add_profile(
                info['assignment'], prof['current_count'])
            mprof = await mock_sched.add_profile(
                info['assignment'], prof['requirement'])
            await sched_complete(game_sched)
            await sched_complete(mock_sched)
            assert ((await sprof.get_structure())['observations_count'] ==
                    prof['current_count'])
            assert sprof['id'] == mprof['id']

            struct = await mprof.get_structure()
            assert_dicts_types(info, struct)
            assert_dicts_equal(info, struct, {'id'})

            assert_dicts_types(summ, (await mprof.get_summary()), (), True)
            assert_dicts_types(obs, (await mprof.get_observations()),
                               {'extended_features', 'features'},
                               True)
            assert_dicts_types(full, (await mprof.get_full_data()),
                               {'extended_features', 'features', 'e', 'f'},
                               True)

        mock_reqs = await mock_sched.get_requirements()
        assert_dicts_types(mock_reqs, reqs)

        await asyncio.gather(*[
            mock_sim.create_game(str(i), 0) for i in range(true_game['id'])])
        mock_game = await mock_sim.create_game(
            true_game['name'], true_game['size'],
            dict(game_summ['configuration']))
        info = await mock_game.get_structure()
        assert_dicts_types(game_info, info)
        assert_dicts_equal(game_info, info)

        symgrps = [
            (grp['name'], grp['count'], grp['strategies'])
            for grp in game_summ['roles']]
        await mock_game.add_symgroups(symgrps)

        # Schedule next profiles
        await asyncio.gather(*[
            game_sched.add_profile(
                prof['symmetry_groups'], prof['observations_count'])
            for prof in game_summ['profiles']])
        await sched_complete(game_sched)

        mock_summ = await mock_game.get_summary()
        assert_dicts_types(game_summ, mock_summ, (), True)
        # TODO Assert full_data and observations


@pytest.mark.asyncio
@pytest.mark.egta
async def test_gets():
    """Test that get functions work"""
    async with api.api() as egta:
        with pytest.raises(AssertionError):
            await egta.get_simulator_fullname('this name is impossible I hope')

        async def test_sim_name(sim):
            """Verify sim name is accurate"""
            assert sim['id'] == (await egta.get_simulator_fullname(
                '{}-{}'.format(sim['name'], sim['version'])))['id']

        await asyncio.gather(*[
            test_sim_name(sim) for sim
            in await egta.get_simulators()])

        sched = next(iter(await egta.get_generic_schedulers()))
        assert (await egta.get_scheduler(sched['id']))['id'] == sched['id']
        assert ((await egta.get_scheduler_name(sched['name']))['id'] ==
                sched['id'])
        with pytest.raises(AssertionError):
            await egta.get_scheduler_name('this name is impossible I hope')

        game = next(iter(await egta.get_games()))
        assert (await egta.get_game(game['id']))['id'] == game['id']
        assert (await egta.get_game_name(game['name']))['id'] == game['id']
        with pytest.raises(AssertionError):
            await egta.get_game_name('this name is impossible I hope')

        folds = egta.get_simulations()
        fold = await folds.__anext__()
        assert (await egta.get_simulation(
            fold['folder']))['folder_number'] == fold['folder']
        for sort in ['job', 'folder', 'profile', 'simulator', 'state']:
            assert 'folder' in await egta.get_simulations(
                column=sort).__anext__()
        with pytest.raises(StopAsyncIteration):
            await egta.get_simulations(page_start=10**9).__anext__()

        for sch in await egta.get_generic_schedulers():  # pragma: no branch
            sched = await sch.get_requirements()
            if sched['scheduling_requirements']:  # pragma: no branch
                break
        prof = sched['scheduling_requirements'][0]
        assert (await egta.get_profile(prof['id']))['id'] == prof['id']


@pytest.mark.asyncio
@pytest.mark.egta
async def test_modify_simulator():
    """Test that we can modify a simulator"""
    # This is very dangerous because we're just finding and modifying a random
    # simulator. However, adding and removing a random role shouldn't really
    # affect anything, so this should be fine
    async with api.api() as egta:
        role = '__unique_role__'
        strat1 = '__unique_strategy_1__'
        strat2 = '__unique_strategy_2__'
        # Old sims seem frozen so > 100
        for sim in await egta.get_simulators():  # pragma: no branch
            if sim['id'] > 100:
                break
        else: # pragma: no cover
            sim = None
            raise AssertionError('no simulators')
        try:
            await sim.add_strategies({role: [strat1]})
            assert ((await sim.get_info())['role_configuration'][role] ==
                    [strat1])

            await sim.add_strategy(role, strat2)
            expected = [strat1, strat2]
            assert ((await sim.get_info())['role_configuration'][role] ==
                    expected)

            await sim.remove_strategies({role: [strat1]})
            assert ((await sim.get_info())['role_configuration'][role] ==
                    [strat2])
        finally:
            await sim.remove_role(role)

        assert role not in (await sim.get_info())['role_configuration']


@pytest.mark.asyncio
@pytest.mark.egta
async def test_modify_scheduler():
    """Test that we can modify a scheduler"""
    async with api.api() as egta:
        for sim in await egta.get_simulators():  # pragma: no branch
            if next(iter(sim['role_configuration'].values()), None): # pragma: no branch pylint: disable=line-too-long
                break
        else: # pragma: no cover
            sim = None
            raise AssertionError('no simulators')
        sched = game = None
        try:
            sched = await sim.create_generic_scheduler(
                '__unique_scheduler__', False, 0, 1, 0, 0)
            await sched.activate()
            await sched.deactivate()

            role = next(iter(sim['role_configuration']))
            strat = sim['role_configuration'][role][0]
            symgrps = [{'role': role, 'strategy': strat, 'count': 1}]
            assignment = api.symgrps_to_assignment(symgrps)

            await sched.add_role(role, 1)

            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert not reqs

            prof = await sched.add_profile(symgrps, 1)
            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert len(reqs) == 1
            assert reqs[0]['requirement'] == 1

            await sched.remove_profile(prof['id'])
            assert (await sched.add_profile(assignment, 2))['id'] == prof['id']
            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert len(reqs) == 1
            assert reqs[0]['requirement'] == 2

            await sched.remove_profile(prof['id'])
            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert not reqs

            assert (await sched.add_profile(symgrps, 3))['id'] == prof['id']
            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert len(reqs) == 1
            assert reqs[0]['requirement'] == 3

            await sched.remove_all_profiles()
            reqs = (await sched.get_requirements())['scheduling_requirements']
            assert not reqs

            await sched.remove_role(role)

            game = await sched.create_game()

        finally:
            if sched is not None:  # pragma: no branch
                await sched.destroy_scheduler()
            if game is not None:  # pragma: no branch
                await game.destroy_game()


@pytest.mark.asyncio
@pytest.mark.egta
async def test_modify_game():
    """Test that we can modify a game"""
    async with api.api() as egta:  # pragma: no branch
        for sim in await egta.get_simulators():  # pragma: no branch
            if next(iter(sim['role_configuration'].values()), None): # pragma: no branch pylint: disable=line-too-long
                break
        else: # pragma: no cover
            sim = None # pytest detects that it's defined
            raise AssertionError('no simulators')
        game = None
        try:
            game = await sim.create_game('__unique_game__', 1)

            summ = await game.get_summary()
            assert not summ['roles']
            assert not summ['profiles']

            role = next(iter(sim['role_configuration']))
            strat = sim['role_configuration'][role][0]
            await game.add_role(role, 1)
            summ = await game.get_summary()
            assert len(summ['roles']) == 1
            assert summ['roles'][0]['name'] == role
            assert summ['roles'][0]['count'] == 1
            assert not summ['roles'][0]['strategies']

            await game.add_strategies({role: [strat]})
            summ = await game.get_summary()
            assert len(summ['roles']) == 1
            assert summ['roles'][0]['name'] == role
            assert summ['roles'][0]['count'] == 1
            assert summ['roles'][0]['strategies'] == [strat]

            await game.remove_strategies({role: [strat]})
            summ = await game.get_summary()
            assert len(summ['roles']) == 1
            assert summ['roles'][0]['name'] == role
            assert summ['roles'][0]['count'] == 1
            assert not summ['roles'][0]['strategies']

            await game.remove_role(role)
            summ = await game.get_summary()
            assert not summ['roles']
            assert not summ['profiles']

        finally:
            if game is not None:  # pragma: no branch
                await game.destroy_game()
