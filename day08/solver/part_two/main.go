package parttwo

import (
	"github.com/nlduy0310/aoc-2024/day08/citymap"
	citymap_parser "github.com/nlduy0310/aoc-2024/day08/citymap/parser"
	"github.com/nlduy0310/aoc-2024/day08/location"
	maths_utils "github.com/nlduy0310/aoc-2024/utils/maths"
)

type Solver struct {
	inputFile string
	cityMap   citymap.CityMap
}

func MustInitPartTwoSolver(inputFile string) Solver {

	parser := citymap_parser.NewMapParser()
	cityMap := parser.MustParse(inputFile)

	return Solver{
		inputFile: inputFile,
		cityMap:   *cityMap,
	}
}

func (s *Solver) Solve() int {

	uniqueLocations := make(map[location.Location]struct{})

	for _, antennas := range s.cityMap.AntennasByName {
		n := len(antennas)

		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				origin, target := antennas[i].Location, antennas[j].Location
				rowDiff, colDiff := target.Row-origin.Row, target.Col-origin.Col
				reducedRowDiff, reducedColDiff := maths_utils.Reduce(rowDiff, colDiff)
				translateVector := location.NewLocation(reducedRowDiff, reducedColDiff)

				for x := 0; ; x++ {
					ok, projectedLocation := s.cityMap.TryTranslate(origin, translateVector, x)
					if ok {
						uniqueLocations[*projectedLocation] = struct{}{}
					} else {
						break
					}
				}
				for x := -1; ; x-- {
					ok, projectedLocation := s.cityMap.TryTranslate(origin, translateVector, x)
					if ok {
						uniqueLocations[*projectedLocation] = struct{}{}
					} else {
						break
					}
				}
			}
		}
	}

	return len(uniqueLocations)
}
