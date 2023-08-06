from .definition import DIGIT_AND_ASCII_RULE_SET, BasePatternRule
from .parse_utils import ParsedPiece, mix


class PiecePatternNodeMix(object):
    def __init__(self, node):
        self._node = node

    @property
    def parsed_piece(self):
        return self._node.parsed_piece

    @property
    def pattern(self):
        return self._node.pattern

    @property
    def piece(self):
        return self._node.piece

    @property
    def count(self):
        return self._node.count


class ClusterNode(PiecePatternNodeMix):
    __slots__ = ('_cluster_name')

    def __init__(self, node, cluster_name=''):
        super(ClusterNode, self).__init__(node)
        self._cluster_name = cluster_name

    @property
    def node(self):
        return self._node

    @property
    def cluster_name(self):
        return self._cluster_name

    def set_pattern(self, pattern, cluster_name):
        self._node.set_pattern(pattern)
        self._cluster_name = cluster_name


class ClusterNodeView(PiecePatternNodeMix):

    __slot__ = ()

    def __init__(self, cluster_node):
        super(ClusterNodeView, self).__init__(cluster_node)
        self._parsed_pieces = None

    @property
    def cluster_node(self):
        return self._node

    def view(self):
        return ' '.join([p.fuzzy_rule for p in self.view_parsed_pieces()])

    def view_parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces

        self._parsed_pieces = [ParsedPiece([piece], [rule])
                               for piece, rule in zip(self.parsed_piece.pieces, self.parsed_piece.rules)]
        return self._parsed_pieces

    @property
    def cluster_name(self):
        return self._node.cluster_name

    def set_pattern(self, pattern, cluster_name):
        self._node.set_pattern(pattern, cluster_name)


class PieceView(ClusterNodeView):
    def view(self):
        return self.piece


class LengthView(ClusterNodeView):
    def view(self):
        return self.parsed_piece.piece_length


class BaseView(ClusterNodeView):
    pass


class MixedView(ClusterNodeView):

    def view_parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces

        if len(self.parsed_piece.rules) <= 1:
            self._parsed_pieces = [self.parsed_piece]
        else:
            mixed_pieces, mixed_rules = mix(
                self.parsed_piece.pieces, self.parsed_piece.rules)

            self._parsed_pieces = [ParsedPiece(
                [piece], [rule]) for piece, rule in zip(mixed_pieces, mixed_rules)]
        return self._parsed_pieces


class LastDotSplitFuzzyView(ClusterNodeView):

    def view_parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces
        rules = self.parsed_piece.rules
        dot_idx = None
        part_num = len(rules)
        for idx, rule in enumerate(rules[::-1]):
            if idx > 2:
                break
            if rule == BasePatternRule.DOT:
                dot_idx = part_num - idx - 1
                break
        self._parsed_pieces = [ParsedPiece([self.parsed_piece.piece],
                                           [self.parsed_piece.fuzzy_rule])]
        if dot_idx is not None:
            skip = False
            for rule in self.parsed_piece.rules[dot_idx + 1:]:
                if rule not in DIGIT_AND_ASCII_RULE_SET:
                    skip = True
                    break
            if not skip:
                pieces = []
                rules = []
                pieces.append(''.join(self.parsed_piece.pieces[0:dot_idx]))
                pieces.append(self.parsed_piece.pieces[dot_idx])
                rules.append(
                    ''.join(sorted(set(self.parsed_piece.rules[0:dot_idx]))))
                rules.append(self.parsed_piece.rules[dot_idx])
                mixed_pieces, mixed_rules = mix(
                    self.parsed_piece.pieces[dot_idx + 1:], self.parsed_piece.rules[dot_idx + 1:])
                pieces.extend(mixed_pieces)
                rules.extend(mixed_rules)
                self._parsed_pieces = [ParsedPiece(
                    [piece], [rule]) for piece, rule in zip(pieces, rules)]
        return self._parsed_pieces


class FuzzyView(ClusterNodeView):
    def view(self):
        return self.parsed_piece.fuzzy_rule

    def view_parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces
        self._parsed_pieces = [ParsedPiece([self.parsed_piece.piece],
                                           [self.parsed_piece.fuzzy_rule])]
        return self._parsed_pieces
