"""Module for command line api"""
import argparse
import asyncio
import contextlib
import itertools
import json
import logging
import sys

import requests

import egtaonline
from egtaonline import api


async def amain(*argv): # pylint: disable=too-many-statements
    """Entry point for async cli with args"""
    parser = argparse.ArgumentParser(
        description="""Command line access to egta online apis. This works well
        in concert with `jq`.""")
    parser.add_argument(
        '--verbose', '-v', action='count', default=0, help="""Sets the
        verbosity of commands. Output is send to standard error.""")
    parser.add_argument(
        '--version', '-V', action='version',
        version='%(prog)s {}'.format(egtaonline.__version__))

    parser_auth = parser.add_mutually_exclusive_group()
    parser_auth.add_argument(
        '--auth-string', '-a', metavar='<auth-string>', help="""The string
        authorization token to connect to egta online.""")
    parser_auth.add_argument(
        '--auth-file', '-f', metavar='<auth-file>', type=argparse.FileType(),
        help="""Filename that just contains the string of the auth token.""")

    subparsers = parser.add_subparsers(title='Subcommands', dest='command')
    subparsers.required = True

    parser_sim = subparsers.add_parser(
        'sim', help="""Get information on or modify a simulator.""",
        description="""Operate on EGTA simulators. By defualt this will return
        information about each simulator on its own line.""")
    parser_sim.add_argument(
        'sim_id', metavar='sim-id', nargs='?', help="""Get information from a
        specific simulator instead of all of them.""")
    parser_sim.add_argument(
        '--sim-version', '-n', metavar='<version-name>', help="""If this is
        specified then the `sim-id` is treated as the name of the simulator and
        the supplied argument is treated as the version.""")
    parser_sim.add_argument(
        '--json', '-j', metavar='json-file', nargs='?',
        type=argparse.FileType('r'), help="""Modify simulator using the
        specified json file. By default this will add all of the roles and
        strategies in the json file. If `delete` is specified, this will remove
        only the strategies specified in the file. `-` can be used to read from
        stdin.""")
    parser_sim.add_argument(
        '--role', '-r', metavar='<role>', help="""Modify `role` of the
        simulator.  By default this will add `role` to the simulator. If
        `delete` is specified this will remove `role` instead. If `strategy` is
        specified see strategy.""")
    parser_sim.add_argument(
        '--strategy', '-s', metavar='<strategy>', help="""Modify `strategy` of
        the simulator. This requires that `role` is also specified. By default
        this adds `strategy` to `role`. If `delete` is specified, then this
        removes the strategy instead.""")
    parser_sim.add_argument(
        '--delete', '-d', action='store_true', help="""Triggers removal of
        roles or strategies instead of addition""")
    parser_sim.add_argument(
        '--zip', '-z', action='store_true', help="""Download simulator zip file
        to stdout.""")

    parser_game = subparsers.add_parser(
        'game', help="""Get information on, create, destroy, or modify a
        game.""", description="""Operate on EGTA Online games. Buy default this
        will return information about each game on its own line.""")
    parser_game.add_argument(
        'game_id', nargs='?', metavar='game-id', help="""Get data from a
        specific game instead of all of them.""")
    parser_game.add_argument(
        '--name', '-n', action='store_true', help="""If specified then get the
        game via its string name not its id number. This is much slower than
        accessing via id number.""")
    parser_game.add_argument(
        '--fetch-conf', metavar='<configuration>', type=argparse.FileType('r'),
        help="""If specified then interpret `game_id` as a simulator id and use
        the file specified as a game configuration. Fetch data of the
        appropriate granularity from the canonical game defined by that
        simulator and this configuration.  Games need specified roles and
        players, these will be pulled from `--json` which must have two top
        level entries, `players` and `strategies` which list the number of
        players per role and the strategies per role respectively`. The
        configuration will not be updated with simulator defaults, so it may be
        helpful to call `sim` first to get those.""")
    parser_game.add_argument(
        '--json', '-j', metavar='json-file', nargs='?',
        type=argparse.FileType('r'), help="""Modify game using the specified
        json file.  By default this will add all of the roles and strategies in
        the json file. If `delete` is specified, this will remove only the
        strategies specified in the file.""")
    parser_game.add_argument(
        '--role', '-r', metavar='<role>', help="""Modify `role` of the game. By
        default this will add `role` to the game. If `delete` is specified this
        will remove `role` instead. If `strategy` is specified see
        strategy.""")
    parser_game.add_argument(
        '--count', '-c', metavar='<count>', help="""If adding a role, the count
        of the role must be specified.""")
    parser_game.add_argument(
        '--strategy', '-s', metavar='<strategy>', help="""Modify `strategy` of
        the game. This requires that `role` is also specified. By default this
        adds `strategy` to `role`. If `delete` is specified, then this removes
        the strategy instead.""")
    parser_game.add_argument(
        '--delete', '-d', action='store_true', help="""Triggers removal of
        roles or strategies instead of addition""")
    parser_gran = (parser_game.add_argument_group('data granularity')
                   .add_mutually_exclusive_group())
    parser_gran.add_argument(
        '--structure', action='store_true', help="""Return game information but
        no profile information. (default)""")
    parser_gran.add_argument(
        '--summary', action='store_true', help="""Return profiles with
        aggregated payoffs.""")
    parser_gran.add_argument(
        '--observations', action='store_true', help="""Return profiles with
        observation data.""")
    parser_gran.add_argument(
        '--full', action='store_true', help="""Return all collected payoff
        data.""")

    parser_sched = subparsers.add_parser(
        'sched', help="""Get information on, create, destroy, or modify a
        scheduler.""", description="""Operate on EGTA Online schedulers. By
        default this will return information about each generic scheduler on
        its own line.""")
    parser_sched.add_argument(
        '--running', action='store_true', help="""If specified, filter
        schedulers by ones that are currently running simulations. To be
        running simulations, a scheduler has to be active, and have at least
        one profile whose current count is less than the requirement for that
        scheduler.""")
    parser_sched.add_argument(
        'sched_id', nargs='?', metavar='sched-id', help="""Get information
        about a specific scheduler. This scheduler doesn't need to be generic,
        but to operate on it, it must be.""")
    parser_sched.add_argument(
        '--name', '-n', action='store_true', help="""If specified then get the
        scheduler via its string name not its id number. This is much slower
        than accessing via id number, and only works for generic
        schedulers.""")
    parser_act = (parser_sched.add_argument_group('scheduler action')
                  .add_mutually_exclusive_group())
    parser_act.add_argument(
        '--requirements', '-r', action='store_true', help="""Get scheuler
        requirements instead of just information.""")
    parser_act.add_argument(
        '--deactivate', action='store_true', help="""Deactivate the specified
        scheduler.""")
    parser_act.add_argument(
        '--delete', '-d', action='store_true', help="""If specified with a
        scheduler, delete it.""")

    # TODO Specifying a flag to change from folder to job id is awkward
    parser_sims = subparsers.add_parser(
        'sims', help="""Get information about pending, scheduled, queued,
        complete, or failed simulations.""", description="""Get information
        about EGTA Online simulations. These are the actual scheduled
        simulations instead of the simulators that generate them. If no folder
        is specified, each simulation comes out on a different line, and can be
        easily filtered with `head` and `jq`.""")
    parser_sims.add_argument(
        'folder', metavar='folder-id', nargs='?', type=int, help="""Get
        information from a specific simulation instead of all of them. This
        should be the folder number of the simulation of interest.""")
    parser_sims.add_argument(
        '-j', '--job', action='store_true', help="""Fetch the simulation with a
        given PBS job id instead of its folder number.""")
    parser_sims.add_argument(
        '--page', '-p', metavar='<start-page>', default=1, type=int, help="""The
        page to start scanning at. (default: %(default)d)""")
    parser_sims.add_argument(
        '--ascending', '-a', action='store_true', help="""Return results in
        ascending order instead of descending.""")
    parser_sims.add_argument(
        '--sort-column', '-s', choices=('job', 'folder', 'profile', 'state'),
        default='job', help="""Column to order results by.  (default:
        %(default)s)""")
    parser_sims.add_argument(
        '--search', default='', metavar='<search-string>', help="""The string
        to filter results by. See egtaonline for examples of what this can
        be.""")
    parser_sims.add_argument(
        '--state',
        choices=['canceled', 'complete', 'failed', 'processing', 'queued',
                 'running'],
        help="""Only select jobs with a specific state.""")
    parser_sims.add_argument(
        '--profile', metavar='<profile-substring>', help="""Limit results to
        simulations whose profiles contain the supplied substring.""")
    parser_sims.add_argument(
        '--simulator', metavar='<sim-substring>', help="""Limit results to
        simulations where the simulator fullname contains the supplied
        substring.""")

    args = parser.parse_args(argv)
    if args.auth_string is None and args.auth_file is not None:
        args.auth_string = args.auth_file.read().strip()
    logging.basicConfig(stream=sys.stderr,
                        level=50 - 10 * min(args.verbose, 4))

    async with api.api(args.auth_string) as eoapi:
        if args.command == 'sim':
            return await _sim(eoapi, args)
        elif args.command == 'game':
            return await _game(eoapi, args)
        elif args.command == 'sched':
            return await _sched(eoapi, args)
        elif args.command == 'sims':
            return await _sims(eoapi, args)
        else:
            assert False  # pragma: no cover


async def _sim(eoapi, args): # pylint: disable=too-many-branches
    """Do stuff with simulators"""
    if args.sim_id is None:  # Get all simulators
        sims = await eoapi.get_simulators()
        try:
            for sim in sims:
                json.dump(sim, sys.stdout)
                sys.stdout.write('\n')
        except (BrokenPipeError, KeyboardInterrupt):  # pragma: no cover
            pass  # Don't care if stream breaks or is killed

    else:  # Operate on a single simulator
        # Get simulator
        if args.sim_version is None:
            sim = await eoapi.get_simulator(int(args.sim_id))
        else:
            sim = await eoapi.get_simulator_fullname('{}-{}'.format(
                args.sim_id, args.sim_version))

        # Operate
        if args.zip:
            info = await sim.get_info()
            url = 'https://' + eoapi.domain + info['source']['url']
            resp = requests.get(url)
            resp.raise_for_status()
            sys.stdout.buffer.write(resp.content)

        elif args.json is not None:  # Load from json
            role_strat = json.load(args.json)
            if args.delete:
                await sim.remove_strategies(role_strat)
            else:
                await sim.add_strategies(role_strat)

        elif args.role is not None:  # Operate on single role or strat
            if args.strategy is not None:  # Operate on strategy
                if args.delete:
                    await sim.remove_strategy(args.role, args.strategy)
                else:
                    await sim.add_strategy(args.role, args.strategy)
            else:  # Operate on role
                if args.delete:
                    await sim.remove_role(args.role)
                else:
                    await sim.add_role(args.role)

        else:  # Return information instead
            json.dump(await sim.get_info(), sys.stdout)
            sys.stdout.write('\n')


async def _game(eoapi, args): # pylint: disable=too-many-statements,too-many-branches
    """Do stuff with games"""
    if args.game_id is None:  # Get all games
        games = await eoapi.get_games()
        try:
            for game in games:
                json.dump(game, sys.stdout)
                sys.stdout.write('\n')
        except (BrokenPipeError, KeyboardInterrupt):  # pragma: no cover
            pass  # Don't care if stream breaks or is killed

    elif args.fetch_conf:  # fetch game data
        desc = json.load(args.json)
        conf = json.load(args.fetch_conf)
        sim_id = int(args.game_id)
        symgrps = {}
        for role, players in desc['players'].items():
            symgrps.setdefault(role, [role, None, None])[1] = players
        for role, strats in desc['strategies'].items():
            symgrps.setdefault(role, [role, None, None])[2] = strats

        game = await eoapi.get_canon_game(
            sim_id, list(symgrps.values()), conf)

        if args.summary:
            dump = await game.get_summary()
        elif args.observations:
            dump = await game.get_observations()
        elif args.full:
            dump = await game.get_full_data()
        else:
            dump = await game.get_structure()
        json.dump(dump, sys.stdout)
        sys.stdout.write('\n')

    else:  # Operate on specific game
        # Get game
        if args.name:
            game = await eoapi.get_game_name(args.game_id)
        else:
            game = await eoapi.get_game(int(args.game_id))

        # Operate
        if args.json is not None:  # Load from json
            role_strat = json.load(args.json)
            if args.delete:
                await game.remove_strategies(role_strat)
            else:
                await game.add_strategies(role_strat)

        elif args.role is not None:  # Operate on single role or strat
            if args.strategy is not None:  # Operate on strategy
                if args.delete:
                    await game.remove_strategy(
                        args.role, args.strategy)
                else:
                    await game.add_strategy(args.role, args.strategy)
            else:  # Operate on role
                if args.delete:
                    await game.remove_role(args.role)
                elif args.count:
                    await game.add_role(args.role, args.count)
                else:
                    raise ValueError(
                        'If adding a role, count must be specified')

        else:  # Return information instead
            if args.summary:
                dump = await game.get_summary()
            elif args.observations:
                dump = await game.get_observations()
            elif args.full:
                dump = await game.get_full_data()
            else:
                dump = await game.get_structure()
            json.dump(dump, sys.stdout)
            sys.stdout.write('\n')


async def _sched(eoapi, args):
    """Do stuff with schedulers"""
    if args.sched_id is None:  # Get all schedulers
        scheds = await eoapi.get_generic_schedulers()
        try:
            if args.running:
                async def print_running(sched):
                    """Print a scheduler if it's running"""
                    reqs = await sched.get_requirements()
                    if any(r['current_count'] < r['requirement'] for r
                           in reqs['scheduling_requirements']):
                        json.dump(sched, sys.stdout)
                        sys.stdout.write('\n')

                await asyncio.gather(*[
                    print_running(s) for s in scheds if s['active']])
            else:
                for sched in scheds:
                    json.dump(sched, sys.stdout)
                    sys.stdout.write('\n')
        except (BrokenPipeError, KeyboardInterrupt):  # pragma: no cover
            pass  # Don't care if stream breaks or is killed

    else:  # Get a single scheduler
        # Get scheduler
        if args.name:
            sched = await eoapi.get_scheduler_name(args.sched_id)
        else:
            sched = await eoapi.get_scheduler(int(args.sched_id))

        # Resolve
        if args.deactivate:
            await sched.deactivate()
        elif args.delete:
            await sched.destroy_scheduler()
        elif args.requirements:
            json.dump(await sched.get_requirements(), sys.stdout)
            sys.stdout.write('\n')
        else:
            json.dump(await sched.get_info(), sys.stdout)
            sys.stdout.write('\n')


async def _sims(eoapi, args):
    """Do stuff with simulations"""
    if args.folder is not None:  # Get info on one simulation
        if args.job:
            itr = eoapi.get_simulations(search='job="{:d}"'.format(args.folder))
            try:
                sim = await itr.__anext__()
            except StopAsyncIteration:
                raise ValueError('No simulation with job id {:d}'.format(
                    args.folder))
            with contextlib.suppress(StopAsyncIteration):
                await itr.__anext__()
                raise ValueError((
                    'Somehow there were multiple simulations with the same '
                    'job id {:d}').format(args.folder))
        else:
            sim = await eoapi.get_simulation(args.folder)
        json.dump(sim, sys.stdout)
        sys.stdout.write('\n')

    else:  # Stream simulations
        search = ' '.join(itertools.chain(
            ['{}="{}"'.format(key, val) for key, val in [
                ('state', args.state), ('simulator', args.simulator),
                ('profile', args.profile)] if val is not None],
            [args.search]))
        try:
            async for sim in eoapi.get_simulations(
                    page_start=args.page, asc=args.ascending,
                    column=args.sort_column, search=search):
                json.dump(sim, sys.stdout)
                sys.stdout.write('\n')
        except (BrokenPipeError, KeyboardInterrupt):  # pragma: no cover
            pass  # Don't care if stream breaks or is killed


def main():  # pragma: no cover
    """Entry point for cli"""
    loop = asyncio.get_event_loop()
    task = asyncio.ensure_future(amain(*sys.argv[1:]))
    try:
        loop.run_until_complete(task)
    except KeyboardInterrupt:
        task.cancel()
        loop.run_forever()
        raise
    finally:
        with contextlib.suppress(asyncio.CancelledError, SystemExit):
            task.exception()


if __name__ == '__main__':  # pragma: no cover
    main()
