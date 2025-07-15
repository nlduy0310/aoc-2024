package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Update struct {
	Pages []int
}

func (u *Update) String() string {

	return fmt.Sprintf("Update[pages=%v]", u.Pages)
}

func NewUpdate(pages []int) (*Update, error) {

	if pages == nil {
		return nil, fmt.Errorf("pages can't be nil")
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("pages can't be empty")
	}

	return &Update{Pages: safeCopyList(pages)}, nil
}

func NewUpdateFromString(str string) (*Update, error) {

	tokens := strings.Split(str, ",")
	if len(tokens) == 0 {
		return nil, fmt.Errorf("invalid update string: '%s'. expected tokens separated by ',', got 0", str)
	}

	pages := make([]int, 0)

	for _, token := range tokens {
		page, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid update string '%s'. '%s' is not a valid number", str, token)
		}
		pages = append(pages, page)
	}

	ret, err := NewUpdate(pages)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
