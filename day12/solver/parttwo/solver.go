package parttwo

import (
	"github.com/nlduy0310/aoc-2024/day12/garden"
	"github.com/nlduy0310/aoc-2024/day12/position"
	"github.com/nlduy0310/aoc-2024/day12/region"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Solver struct {
	inputFile string
	garden    garden.Garden
	regions   []region.Region
}

func MustInitSolver(inputFile string) Solver {

	garden := garden.MustParseFromFile(inputFile)

	return Solver{
		inputFile: inputFile,
		garden:    garden,
		regions:   getRegions(garden),
	}
}

func getRegions(g garden.Garden) []region.Region {

	ret := make([]region.Region, 0)

	visited := make([][]bool, g.Height)
	for r := range g.Height {
		visited[r] = utils.SliceInit(g.Width, false)
	}

	for rowIdx := range g.Height {
		for colIdx := range g.Width {
			if visited[rowIdx][colIdx] {
				continue
			}

			currentRegion := getRegionAt(g, position.NewPosition(rowIdx, colIdx))
			ret = append(ret, currentRegion)

			for _, regionPosition := range currentRegion.GetPositions() {
				visited[regionPosition.Row][regionPosition.Col] = true
			}
		}
	}

	return ret
}

func getRegionAt(g garden.Garden, origin position.Position) region.Region {

	plant, err := g.GetPlantAt(origin)
	utils.PanicIf(err)

	ret := region.NewRegion(*plant)

	positions := make(map[position.Position]struct{})
	positions[origin] = struct{}{}

	currentPositions := []position.Position{origin}
	for len(currentPositions) > 0 {
		tmp := make([]position.Position, 0)
		for _, currentPosition := range currentPositions {
			surroundingPositions := g.GetSurroundingPositionsFiltered(
				currentPosition,
				func(p position.Position) bool {
					pl, err := g.GetPlantAt(p)
					_, ok := positions[p]
					return !ok && err == nil && *pl == *plant
				},
			)
			tmp = append(tmp, surroundingPositions...)
			for _, surroundingPosition := range surroundingPositions {
				positions[surroundingPosition] = struct{}{}
			}
		}
		currentPositions = tmp
	}

	for k := range positions {
		ret.AddPosition(k)
	}

	return ret
}

func (s Solver) Solve() int {

	ret := 0

	for _, region := range s.regions {
		ret += region.GetArea() * region.GetSidesCount()
	}

	return ret
}
