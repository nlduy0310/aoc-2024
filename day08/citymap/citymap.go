package citymap

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2024/day08/antenna"
	"github.com/nlduy0310/aoc-2024/day08/location"
	"github.com/nlduy0310/aoc-2024/utils"
)

type CityMap struct {
	Width          int
	Height         int
	AntennasByName map[string][]antenna.Antenna
}

func (m *CityMap) tryGetAntennaAt(l location.Location) *antenna.Antenna {

	for _, antennas := range m.AntennasByName {
		for _, antenna := range antennas {
			if antenna.Location == l {
				return &antenna
			}
		}
	}

	return nil
}

func (m *CityMap) TryProject(origin, target location.Location, scale int) (bool, *location.Location) {

	dr := target.Row - origin.Row
	dc := target.Col - origin.Col

	projectedLocation := location.NewLocation(
		origin.Row+dr*scale,
		origin.Col+dc*scale,
	)

	if m.Contains(projectedLocation) {
		return true, &projectedLocation
	} else {
		return false, nil
	}
}

func (m *CityMap) TryTranslate(origin, direction location.Location, scale int) (bool, *location.Location) {

	projectedLocation := location.NewLocation(
		origin.Row+direction.Row*scale,
		origin.Col+direction.Col*scale,
	)

	if m.Contains(projectedLocation) {
		return true, &projectedLocation
	} else {
		return false, nil
	}
}

func (m *CityMap) Contains(l location.Location) bool {

	return utils.IsInRangeInclusive(l.Row, 0, m.Height-1) && utils.IsInRangeInclusive(l.Col, 0, m.Width-1)
}

func (m *CityMap) String() string {

	return fmt.Sprintf(
		"CityMap["+
			"width=%d, "+
			"height=%d, "+
			"antennas=%v"+
			"]",
		m.Width,
		m.Height,
		m.AntennasByName,
	)
}

func (m *CityMap) PrettyString() string {

	emptyLocationString := "."
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("CityMap: width=%d, height=%d\n", m.Width, m.Height))

	for row := range m.Height {
		for col := range m.Width {
			if possibleAntenna := m.tryGetAntennaAt(location.NewLocation(row, col)); possibleAntenna != nil {
				builder.WriteString(possibleAntenna.Name)
			} else {
				builder.WriteString(emptyLocationString)
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
