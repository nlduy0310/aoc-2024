package partone

import (
	"github.com/nlduy0310/aoc-2024/day08/citymap"
	citymap_parser "github.com/nlduy0310/aoc-2024/day08/citymap/parser"
	"github.com/nlduy0310/aoc-2024/day08/location"
)

type Solver struct {
	inputFile string
	cityMap   citymap.CityMap
}

func MustInitPartOneSolver(inputFile string) *Solver {

	mapParser := citymap_parser.NewMapParser()
	cityMap := mapParser.MustParse(inputFile)

	return &Solver{
		inputFile: inputFile,
		cityMap:   *cityMap,
	}
}

func (s *Solver) Solve() int {

	uniqueLocations := make(map[location.Location]bool)

	for _, antennas := range s.cityMap.AntennasByName {
		n := len(antennas)

		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				var ok bool
				var projectedLocation *location.Location
				ok, projectedLocation = s.cityMap.TryProject(antennas[i].Location, antennas[j].Location, 2)
				if ok {
					uniqueLocations[*projectedLocation] = true
				}
				ok, projectedLocation = s.cityMap.TryProject(antennas[j].Location, antennas[i].Location, 2)
				if ok {
					uniqueLocations[*projectedLocation] = true
				}
			}
		}
	}

	return len(uniqueLocations)
}
