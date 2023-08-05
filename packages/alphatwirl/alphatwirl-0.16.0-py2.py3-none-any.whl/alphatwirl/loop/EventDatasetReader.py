# Tai Sakuma <tai.sakuma@gmail.com>
import copy
import itertools
from operator import itemgetter
from collections import OrderedDict

from .EventLoop import EventLoop

##__________________________________________________________________||
class EventDatasetReader(object):
    """This class manages objects involved in reading events in data sets.

    On receiving a data set, this class calls the function
    split_into_build_events(), which splits the data set into chunks,
    creates the function build_events() for each chunk, and returns a
    list of the functions. Then, for each build_events(), This class
    creates a copy of the reader, creates an event loop, and send it
    to the event loop runner.

    At the end, this class receives results from the event loop runner
    and have the collector collect them.

    """
    def __init__(self, eventLoopRunner, reader, collector,
                 split_into_build_events):

        self.eventLoopRunner = eventLoopRunner
        self.reader = reader
        self.collector = collector
        self.split_into_build_events = split_into_build_events

        self.EventLoop = EventLoop

        self.dataset_nreaders = [ ]

        self.runids = [ ]
        self.runid_dataset_map = { }
        self.dataset_runid_reader_map = OrderedDict()

    def __repr__(self):
        name_value_pairs = (
            ('eventLoopRunner',         self.eventLoopRunner),
            ('reader',                  self.reader),
            ('collector',               self.collector),
            ('split_into_build_events', self.split_into_build_events),
        )
        return '{}({})'.format(
            self.__class__.__name__,
            ', '.join(['{}={!r}'.format(n, v) for n, v in name_value_pairs]),
        )

    def begin(self):
        self.eventLoopRunner.begin()
        self.dataset_nreaders = [ ]

        self.runids = [ ]
        self.runid_dataset_map = { }
        self.dataset_runid_reader_map = OrderedDict()

    def read(self, dataset):
        build_events_list = self.split_into_build_events(dataset)
        self.dataset_nreaders.append((dataset, len(build_events_list)))
        eventLoops = [ ]
        for build_events in build_events_list:
            reader = copy.deepcopy(self.reader)
            eventLoop = self.EventLoop(build_events, reader, dataset.name)
            eventLoops.append(eventLoop)
        runids = self.eventLoopRunner.run_multiple(eventLoops)
        self.runids.extend(runids)
        self.runid_dataset_map.update({i: dataset.name for i in runids})
        self.dataset_runid_reader_map[dataset.name] = OrderedDict([(i, None) for i in runids])

    def end(self):

        runids_towait = self.runids[:]
        while runids_towait:
            runid, reader = self.eventLoopRunner.receive_one()
            self._merge_imp_1(runid, reader)
            runids_towait.remove(runid)

        dataset_readers_list = [(d, list(rr.values())) for d, rr in self.dataset_runid_reader_map.items()]

        return self.collector.collect(dataset_readers_list)

    def _merge_imp_1(self, runid, reader):
        dataset = self.runid_dataset_map[runid]
        runid_reader_map = self.dataset_runid_reader_map[dataset]
        self._merge_imp_2(runid_reader_map, runid, reader)

    def _merge_imp_2(self, map_, id_, data):

        ## merge to prev
        idx = list(map_.keys()).index(id_)
        idx_prev = idx - 1
        if idx_prev < 0:
            map_[id_] = data
        else:
            data_prev = list(map_.values())[idx_prev]
            id_prev = list(map_.keys())[idx_prev]
            if data_prev is None:
                map_[id_] = data
            else:
                if hasattr(data_prev, 'merge'):
                    data_prev.merge(data)
                map_.pop(id_)
                id_ = id_prev
                data = data_prev

        ## merge next
        idx = list(map_.keys()).index(id_)
        idx_next = idx + 1
        if idx_next < len(map_):
            data_next = list(map_.values())[idx_next]
            if data_next is not None:
                if hasattr(data, 'merge'):
                    data.merge(data_next)
                map_.pop(list(map_.keys())[idx_next])

##__________________________________________________________________||
