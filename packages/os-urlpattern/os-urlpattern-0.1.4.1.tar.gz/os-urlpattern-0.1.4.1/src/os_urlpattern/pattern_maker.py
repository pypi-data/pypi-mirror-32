from .compat import itervalues
from .parse_utils import PieceParser, digest, parse_url
from .pattern_tree import PatternTree
from .piece_pattern_tree import PiecePatternTree
from .utils import load_obj


class PatternMaker(object):
    def __init__(self, config):
        self._config = config
        self._parser = PieceParser()
        self._makers = {}

    def load(self, url):
        url_meta, pieces = parse_url(url)
        parsed_pieces = [self._parser.parse(piece) for piece in pieces]
        sid = digest(url_meta, [p.fuzzy_rule for p in parsed_pieces])
        if sid not in self._makers:
            self._makers[sid] = Maker(self._config, url_meta)
        return self._makers[sid].load(parsed_pieces)

    def process(self):
        for maker in itervalues(self._makers):
            yield maker.make()


class Maker(object):
    def __init__(self, config, url_meta):
        self._config = config
        self._url_meta = url_meta
        self._piece_pattern_tree = PiecePatternTree()
        self._cluster = self._config.get('make', 'cluster_algorithm')

    def load(self, parsed_pieces, count=1, uniq_path=True):
        return self._piece_pattern_tree.add_from_parsed_pieces(
            parsed_pieces, count, uniq_path)

    def _path_dump_and_load(self, src, dest, index=0):
        for path in src.dump_paths():
            if path:
                dest.load_path(path[index:])

    def cluster(self):
        for clusterd_tree in self._cluster(self._config, self._url_meta,
                                           self._piece_pattern_tree):
            yield clusterd_tree

    def make(self):
        pattern_tree = PatternTree(self._url_meta)
        for clustered_tree in self.cluster():
            self._path_dump_and_load(clustered_tree, pattern_tree, 1)
        return pattern_tree
