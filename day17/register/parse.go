package register

import (
	"fmt"
	"regexp"
	"strconv"
)

var registerLinePattern = regexp.MustCompile(`^Register\s+([A-Z]):\s+(\d+)$`)

func ParseLine(line string) (*Register, error) {
	noMatchError := fmt.Errorf("line \"%s\" doesn't match register line pattern", line)
	submatchs := registerLinePattern.FindStringSubmatch(line)
	if submatchs == nil {
		return nil, noMatchError
	}
	if len(submatchs) != 3 {
		panic(fmt.Sprintf("invalid number of submatches: expected 3, got %d", len(submatchs)))
	}

	registerName := submatchs[1]
	registerValStr := submatchs[2]
	registerVal, err := strconv.Atoi(registerValStr)
	if err != nil {
		panic(fmt.Sprintf("register value string \"%s\" can not be parsed", registerValStr))
	}

	return NewRegister(registerName, registerVal), nil
}
