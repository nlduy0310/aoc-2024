package direction

import (
	"strings"

	"github.com/nlduy0310/aoc-2024/utils"
)

func sliceString(directions []Direction, separator string) string {

	return strings.Join(
		utils.SliceMap(
			directions,
			func(d Direction) string {
				return d.String()
			},
		),
		separator,
	)
}
