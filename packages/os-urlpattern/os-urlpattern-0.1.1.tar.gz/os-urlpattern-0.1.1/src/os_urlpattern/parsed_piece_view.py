from .definition import DIGIT_AND_ASCII_RULE_SET, BasePatternRule
from .parse_utils import ParsedPiece, mix


class ParsedPieceView(object):
    __slot__ = ('_parsed_piece', '_view')

    def __init__(self, parsed_piece):
        self._parsed_piece = parsed_piece
        self._parsed_pieces = None
        self._view = None

    @property
    def parsed_piece(self):
        return self._parsed_piece

    def __eq__(self, o):
        if not isinstance(o, ParsedPieceView):
            return False
        return self.view == o.view

    def __hash__(self):
        return hash(self.view)

    @property
    def view(self):
        if self._view is None:
            self._view = u' '.join([p.fuzzy_rule for p in self.parsed_pieces])
        return self._view

    @property
    def parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces

        self._parsed_pieces = [ParsedPiece([piece], [rule]) for piece, rule in zip(
            self._parsed_piece.pieces, self._parsed_piece.rules)]
        return self._parsed_pieces


class PieceView(ParsedPieceView):

    def __init__(self, parsed_piece):
        super(PieceView, self).__init__(parsed_piece)
        self._view = self._parsed_piece.piece


class LengthView(ParsedPieceView):
    def __init__(self, parsed_piece):
        super(LengthView, self).__init__(parsed_piece)
        self._view = self._parsed_piece.piece_length


class BaseView(ParsedPieceView):
    pass


class MixedView(ParsedPieceView):

    @property
    def parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces

        if len(self._parsed_piece.rules) <= 1:
            self._parsed_pieces = [self._parsed_piece]
        else:
            mixed_pieces, mixed_rules = mix(
                self._parsed_piece.pieces, self._parsed_piece.rules)

            self._parsed_pieces = [ParsedPiece(
                [piece], [rule]) for piece, rule in zip(mixed_pieces, mixed_rules)]
        return self._parsed_pieces


class LastDotSplitFuzzyView(ParsedPieceView):

    @property
    def parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces
        rules = self._parsed_piece.rules
        dot_idx = None
        part_num = len(rules)
        for idx, rule in enumerate(rules[::-1]):
            if idx > 2:
                break
            if rule == BasePatternRule.DOT:
                dot_idx = part_num - idx - 1
                break
        self._parsed_pieces = [ParsedPiece([self._parsed_piece.piece],
                                           [self._parsed_piece.fuzzy_rule])]
        if dot_idx is not None:
            skip = False
            for rule in self._parsed_piece.rules[dot_idx + 1:]:
                if rule not in DIGIT_AND_ASCII_RULE_SET:
                    skip = True
                    break
            if not skip:
                pieces = []
                rules = []
                pieces.append(u''.join(self._parsed_piece.pieces[0:dot_idx]))
                pieces.append(self._parsed_piece.pieces[dot_idx])
                rules.append(
                    u''.join(sorted(set(self._parsed_piece.rules[0:dot_idx]))))
                rules.append(self._parsed_piece.rules[dot_idx])
                mixed_pieces, mixed_rules = mix(
                    self._parsed_piece.pieces[dot_idx + 1:],
                    self._parsed_piece.rules[dot_idx + 1:])
                pieces.extend(mixed_pieces)
                rules.extend(mixed_rules)
                self._parsed_pieces = [ParsedPiece(
                    [piece], [rule]) for piece, rule in zip(pieces, rules)]
        return self._parsed_pieces


class FuzzyView(ParsedPieceView):
    def __init__(self, parsed_piece):
        super(FuzzyView, self).__init__(parsed_piece)
        self._view = self._parsed_piece.fuzzy_rule

    @property
    def parsed_pieces(self):
        if self._parsed_pieces:
            return self._parsed_pieces
        self._parsed_pieces = [ParsedPiece([self._parsed_piece.piece],
                                           [self._parsed_piece.fuzzy_rule])]
        return self._parsed_pieces
