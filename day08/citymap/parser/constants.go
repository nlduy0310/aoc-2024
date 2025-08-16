package parser

import (
	rune_utils "github.com/nlduy0310/aoc-2024/utils/runes"
)

var defaultEmptyLocationRunes []rune = []rune{'.'}
var defaultAntennaLocationRunes []rune = append(rune_utils.Digits, rune_utils.Letters...)
