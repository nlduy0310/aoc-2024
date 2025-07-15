package main

import (
	"fmt"
	"strconv"
	"strings"
)

type RulePair struct {
	Before int
	After  int
}

func (rp RulePair) String() string {

	return fmt.Sprintf("RulePair[before=%d, after=%d]", rp.Before, rp.After)
}

func NewRulePair(before, after int) RulePair {

	return RulePair{Before: before, After: after}
}

func NewRulePairFromString(str string) (RulePair, error) {

	parts := strings.Split(str, "|")
	if len(parts) != 2 {
		return RulePair{}, fmt.Errorf("invalid rule string: '%s'. expected 2 tokens separated by '|', got %d", str, len(parts))
	}

	before, err := strconv.Atoi(parts[0])
	if err != nil {
		return RulePair{}, fmt.Errorf("invalid rule string: '%s'. '%s' is not a valid integer", str, parts[0])
	}

	after, err := strconv.Atoi(parts[1])
	if err != nil {
		return RulePair{}, fmt.Errorf("invalid rule string: '%s'. '%s' is not a valid integer", str, parts[1])
	}

	return NewRulePair(before, after), nil
}
