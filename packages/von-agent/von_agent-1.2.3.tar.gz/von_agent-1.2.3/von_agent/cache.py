"""
Copyright 2017-2018 Government of Canada - Public Services and Procurement Canada - buyandsell.gc.ca

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""


import json
import logging

from threading import RLock
from time import time
from typing import Awaitable, Callable, Tuple, Union
from von_agent.error import BadRevStateTime, CacheIndex
from von_agent.tails import Tails
from von_agent.util import SchemaKey, schema_key


class SchemaCache:
    """
    Retain schemata and fetch by schema key (origin_did, name, version) or by sequence number.
    Note that schema key is isomorphic to schema_id, but since schema_id is a str and indy-sdk
    stores sequence number as a str in some cases, it is more defensive to index by schema key
    than schema_id.

    A lock shares access to critical sections as relying code specifies them (e.g., check and get/set).
    Note that this one lock applies across all instances - the design of this class intends it to be a singleton.
    """

    lock = RLock()

    def __init__(self) -> None:
        """
        Initialize schema cache data.
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.__init__: >>>')

        self._schema_key2schema = {}
        self._seq_no2schema_key = {}

        logger.debug('SchemaCache.__init__: <<<')

    def __setitem__(self, index: Union[SchemaKey, int], schema: dict) -> dict:
        """
        Put schema into cache and return it.

        :param index: schema key or sequence number
        :param schema: schema to put into cache
        :return: input schema
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.__setitem__: >>> index: {}, schema: {}'.format(index, schema))

        if isinstance(index, SchemaKey):
            self._schema_key2schema[index] = schema
            self._seq_no2schema_key[schema['seqNo']] = index
        elif isinstance(index, int):
            s_key = schema_key(schema['id'])
            self._schema_key2schema[s_key] = schema
            self._seq_no2schema_key[index] = s_key
        else:
            logger.debug(
                'SchemaCache.__setitem__: <!< Bad index {} must be a schema key or a sequence number'.format(index))
            raise CacheIndex('Bad index {} must be a schema key or a sequence number'.format(index))

        logger.debug('SchemaCache.__setitem__: <<< {}'.format(schema))
        return schema

    def contains(self, index: Union[SchemaKey, int]) -> bool:
        """
        Return whether the cache contains a schema for the input key or sequence number.

        :param index: schema key or sequence number
        :return: whether the cache contains a schema for the input index
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.contains: >>> index: {}'.format(index))

        rv = None
        if isinstance(index, SchemaKey):
            rv = (index in self._schema_key2schema)
        elif isinstance(index, int):
            rv = (index in self._seq_no2schema_key)
        else:
            rv = False

        logger.debug('SchemaCache.contains: <<< {}'.format(rv))
        return rv

    def index(self) -> dict:
        """
        Return dict mapping content sequence numbers to schema keys.

        :return: dict mapping sequence numbers to schema keys
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.index: >>>')

        rv = self._seq_no2schema_key
        logger.debug('SchemaCache.index: <<< {}'.format(rv))
        return rv

    def __getitem__(self, index: Union[SchemaKey, int]) -> dict:
        """
        Get schema by key or sequence number, or raise CacheIndex for no such schema.

        Raise CacheIndex for no such index in schema store.

        :param index: schema key or sequence number
        :return: corresponding schema or None
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.__getitem__: >>> index: {}'.format(index))

        rv = None
        if isinstance(index, SchemaKey):
            rv = self._schema_key2schema[index]
        elif isinstance(index, int):
            try:
                rv = self._schema_key2schema[self._seq_no2schema_key[index]]
            except KeyError:
                logger.debug('SchemaCache.__getitem__: <!< index {} not present'.format(index))
                raise CacheIndex('{}'.format(index))
        else:
            logger.debug('SchemaCache.__getitem__: <!< index {} must be int or SchemaKey'.format(index))
            raise CacheIndex('{} must be int or SchemaKey'.format(index))

        logger.debug('SchemaCache.__getitem__: <<< {}'.format(rv))
        return rv

    def schema_key_for(self, seq_no: int) -> SchemaKey:
        """
        Get schema key for schema by sequence number if known, None for no such schema in cache.

        :param seq_no: sequence number
        :return: corresponding schema key or None
        """

        logger = logging.getLogger(__name__)
        logger.debug('SchemaCache.schema_key_for: >>> seq_no: {}'.format(seq_no))

        rv = self._seq_no2schema_key.get(seq_no, None)

        logger.debug('SchemaCache.schema_key_for: <<< {}'.format(rv))
        return rv

    def dict(self) -> dict:
        """
        Return flat dict with schemata cached.

        :return: flat dict, indexed by sequence number plus schema key data
        """

        return {
            '{}; {}'.format(seq_no, tuple(self._seq_no2schema_key[seq_no])):
                self._schema_key2schema[self._seq_no2schema_key[seq_no]] for seq_no in self._seq_no2schema_key}

    def __str__(self) -> str:
        """
        Return string pretty-print.

        :return: string pretty-print
        """

        return 'SchemaCache({})'.format(self.dict())


class RevRegUpdateFrame:
    """
    Revocation registry delta or state update, plus metadata, in revocation cache (which indexes on rev reg id).
    Keeps track of last query time, last-asked-for time, timestamp on distributed ledger, and rev reg update.
    The last query time is purely for cache management.

    Holder-Prover agents use deltas to create proof; verifier agents use updates to verify them.

    Necessarily for each cached delta frame, timestamp <= ask_for <= qtime.
    """

    def __init__(self, ask_for: int, timestamp: int, rr_update: dict):
        """
        Initialize a new revocation registry update frame for revocation cache.

        :param ask_for: the time (epoch sec) of interest
        :param timestamp: the timestamp (epoch sec) corresponding to the revocation delta on the ledger
        :param rr_update: the indy-sdk revocation registry delta or state update
        """

        self._qtime = int(time())
        self._ask_for = ask_for
        self._timestamp = timestamp
        self._rr_update = rr_update

    @property
    def qtime(self) -> int:
        """
        Accessor for the latest query time resolving to current frame.

        :return: latest latest query time resolving to current frame
        """

        return self._qtime

    @property
    def ask_for(self) -> int:
        """
        Accessor for the latest cached time of interest associated with the rev reg update.

        :return: latest time of interest requested regarding current frame's rev reg update
        """

        return self._ask_for

    @property
    def timestamp(self) -> int:
        """
        Accessor for timestamp on the distributed ledger for the rev reg update.

        :return: timestamp on distributed ledger for current frame's rev reg update
        """

        return self._timestamp

    @property
    def rr_update(self) -> dict:
        """
        Accessor for rev reg update.

        :return: current frame's rev reg update
        """

        return self._rr_update

    def __repr__(self):
        """
        Return canonical representation of the item.
        """

        return 'RevRegUpdateFrame({}, {}, {})'.format(self.ask_for, self.timestamp, self.rr_update)

    def __str__(self):
        """
        Return representation of the item showing query time.
        """

        return 'RevRegUpdateFrame<qtime={}, ask_for={}, timestamp={}, rr_update={}>'.format(
            self.qtime,
            self.ask_for,
            self.timestamp,
            self.rr_update)


class RevoCacheEntry:
    """
    Revocation cache entry housing:
    * a revocation registry definition
    * a Tails structure
    * a list of revocation delta frames.
    """

    def __init__(self, rev_reg_def: dict, tails: Tails = None):
        """
        Initialize with revocation registry definition, optional tails file.
        Set revocation delta frames lists for rev reg deltas and rev reg states empty.

        :param rev_reg_def: revocation registry definition
        :param tails: current tails file object
        """

        logger = logging.getLogger(__name__)
        logger.debug('RevoCacheEntry.__init__: >>> rev_reg_def: {}, tails: {}'.format(rev_reg_def, tails))

        self._rev_reg_def = rev_reg_def or None
        self._tails = tails or None
        self._rr_delta_frames = []  # for holder-prover, creating proof
        self._rr_state_frames = []  # for verifier, verifying proof

        logger.debug('RevoCacheEntry.__init__: <<<')

    @property
    def rev_reg_def(self) -> dict:
        """
        Return rev reg def from cache entry.
        """

        return self._rev_reg_def

    @property
    def tails(self) -> Tails:
        """
        Return current tails file from cache entry.
        """

        return self._tails

    @property
    def rr_delta_frames(self) -> list:
        """
        Return current revocation delta frame list.
        """

        return self._rr_delta_frames

    @property
    def rr_state_frames(self) -> list:
        """
        Return current revocation state frame list.
        """

        return self._rr_state_frames

    async def _get_update(self, rr_builder: Callable, ask_for: int, delta: bool) -> (str, int):
        """
        Get rev reg delta/state json, and its timestamp on the distributed ledger,
        from cached rev reg delta/state frames list or distributed ledger,
        updating cache as necessary.

        Raise BadRevStateTime if caller asks for a delta/state in the future.

        Issuer agents cannot revoke retroactively.
        Hence, for any new request against asked-for time ask_for:
        * if the cache has a frame f on f.timestamp <= ask_for <= f.ask_for,
          > return its rev reg delta/state
        * if the cache has a frame f on f.timestamp < ask_for,
          > check the distributed ledger for a delta to/state for the rev reg since e.timestamp;
            - if there is one, merge it to a new delta/state and add new frame to cache; return rev reg delta/state
            - otherwise, update the ask_for time in the frame and return the rev reg delta/state
        * otherwise, there is no cache frame f on f.timestamp < ask_for:
          > create new frame and add it to cache; return rev reg delta/state.

        On return of any previously existing rev reg delta/state frame, always update its query time beforehand.

        :param rr_builder: callback to build rev reg delta/state if need be (specify holder-prover agent's
            _build_rr_delta_json() or verifier agent's _build_rr_state_json() as needed)
        :param ask_for: time (epoch seconds) of interest; upper-bounds returned revocation delta/state timestamp
        :param delta: true to operate on rev reg deltas, false for updates
        :return: rev reg delta/state json and ledger timestamp (epoch seconds)
        """

        now = int(time())
        if ask_for > now:
            raise BadRevStateTime('Cannot query a rev reg {} in the future ({} > {})'.format(
                'delta' if delta else 'state',
                ask_for,
                now))

        cache_frame = None
        rr_update_json = None
        rr_frames = self._rr_delta_frames if delta else self._rr_state_frames

        frames = [frame for frame in rr_frames if frame.timestamp <= ask_for <= frame.ask_for]
        if frames:
            assert len(frames) == 1
            cache_frame = frames[0]
        else:
            frames = [frame for frame in rr_frames if frame.timestamp < ask_for]  # frame.ask_for < ask_for
            if frames:
                latest_cached = max(frames, key=lambda frame: frame.timestamp)
                if delta:
                    (rr_update_json, timestamp) = await rr_builder(
                        self.rev_reg_def['id'],
                        to=ask_for,
                        fro=latest_cached.timestamp,
                        fro_delta=latest_cached.rr_update)
                else:
                    (rr_update_json, timestamp) = await rr_builder(self.rev_reg_def['id'], ask_for)
                if timestamp == latest_cached.timestamp:
                    latest_cached._ask_for = ask_for  # this timestamp now known good through more recent ask_for
                    cache_frame = latest_cached
            else:
                (rr_update_json, timestamp) = await rr_builder(self.rev_reg_def['id'], ask_for)

        if cache_frame is None:
            cache_frame = RevRegUpdateFrame(ask_for, timestamp, json.loads(rr_update_json))
            rr_frames.append(cache_frame)

            mark = 4096**0.5  # max rev reg size = 4096; heuristic: hover max around sqrt(4096) = 64
            if len(rr_frames) > int(mark * 1.25):
                rr_frames.sort(key=lambda x: -x.qtime)  # order by descending query time
                del rr_frames[int(mark * 0.75):]  # retain most recent, grow again from here
        else:
            cache_frame._qtime = int(time())

        return (json.dumps(cache_frame.rr_update), cache_frame.timestamp)

    async def get_delta_json(
            self,
            rr_delta_builder: Callable[['HolderProver', str, int, int, dict], Awaitable[Tuple[str, int]]],
            ask_for: int) -> (str, int):
        """
        Get rev reg delta json, and its timestamp on the distributed ledger,
        from cached rev reg delta frames list or distributed ledger,
        updating cache as necessary.

        Raise BadRevStateTime if caller asks for a delta to the future.

        Issuer agents cannot revoke retroactively.
        Hence, for any new request against asked-for time ask_for:
        * if the cache has a frame f on f.timestamp <= ask_for <= f.ask_for,
          > return its rev reg delta
        * if the cache has a frame f on f.timestamp < ask_for,
          > check the distributed ledger for a delta to the rev reg delta since e.timestamp;
            - if there is one, merge it to a new delta and add new frame to cache; return rev reg delta
            - otherwise, update the ask_for time in the frame and return the rev reg delta
        * otherwise, there is no cache frame f on f.timestamp < ask_for:
          > create new frame and add it to cache; return rev reg delta.

        On return of any previously existing rev reg delta frame, always update its query time beforehand.

        :param rr_delta_builder: callback to build rev reg delta if need be (specify agent instance's _build_rr_delta())
        :param ask_for: time (epoch seconds) of interest; upper-bounds returned revocation delta timestamp
        :return: rev reg delta json and ledger timestamp (epoch seconds)
        """

        logger = logging.getLogger(__name__)
        logger.debug('RevoCacheEntry.get_delta_json: >>> rr_delta_builder: (method), ask_for: {}'.format(ask_for))

        rv = await self._get_update(rr_delta_builder, ask_for, True)
        logger.debug('RevoCacheEntry.get_delta_json: <<< {}'.format(rv))
        return rv

    async def get_state_json(
            self,
            rr_state_builder: Callable[['Verifier', str, int], Awaitable[Tuple[str, int]]],
            ask_for: int) -> (str, int):
        """
        Get rev reg state json, and its timestamp on the distributed ledger,
        from cached rev reg state frames list or distributed ledger,
        updating cache as necessary.

        Raise BadRevStateTime if caller asks for a state in the future.

        On return of any previously existing rev reg state frame, always update its query time beforehand.

        :param rr_state_builder: callback to build rev reg state if need be (specify agent instance's _build_rr_state())
        :param ask_for: time (epoch seconds) of interest; upper-bounds returned revocation state timestamp
        :return: rev reg state json and ledger timestamp (epoch seconds)
        """

        logger = logging.getLogger(__name__)
        logger.debug('RevoCacheEntry.get_state_json: >>> rr_state_builder: (method), ask_for: {}'.format(ask_for))

        rv = await self._get_update(rr_state_builder, ask_for, False)
        logger.debug('RevoCacheEntry.get_state_json: <<< {}'.format(rv))
        return rv


SCHEMA_CACHE = SchemaCache()
CRED_DEF_CACHE = type('CredDefCache', (dict,), {'lock': RLock()})()
REVO_CACHE = type('RevoCache', (dict,), {'lock': RLock()})()
