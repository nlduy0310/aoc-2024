package antenna

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day08/location"
	"github.com/nlduy0310/aoc-2024/utils"
)

type Antenna struct {
	Name     string
	Location location.Location
}

func NewAntenna(name string, location location.Location) Antenna {

	utils.Assert(len(name) == 1, "an attenna's name must be exactly one character long")

	return Antenna{
		Name:     name,
		Location: location,
	}
}

func (antenna *Antenna) String() string {

	return fmt.Sprintf(
		"Antenna["+
			"name=%s, "+
			"location=%s"+
			"]",
		antenna.Name,
		antenna.Location.String(),
	)

}
