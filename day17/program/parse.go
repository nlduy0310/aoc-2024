package program

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2024/day17/programstate"
	"github.com/nlduy0310/aoc-2024/day17/register"
	"github.com/nlduy0310/aoc-2024/utils"
)

var programLinePattern = regexp.MustCompile(`^Program:\s+([0-7,]+)$`)
var requiredRegisterNames = []string{"A", "B", "C"}

func ParseFromFile(file string) (*Program, error) {
	lines, err := utils.ReadLines(file)
	if err != nil {
		return nil, fmt.Errorf("can not parse from file \"%s\": %w", file, err)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("can not parse from empty file: \"%s\"", file)
	}

	i := 0
	registers := make([]*register.Register, 0)
	foundRequiredRegisters := make([]bool, len(requiredRegisterNames))
	for ; len(lines[i]) > 0; i++ {
		register, err := register.ParseLine(lines[i])
		if err != nil {
			return nil, fmt.Errorf("can not parse register at line %d: %w", i, err)
		}
		registers = append(registers, register)

		if idx := slices.Index(requiredRegisterNames, register.Name); idx != -1 {
			foundRequiredRegisters[idx] = true
		}
	}
	for idx, found := range foundRequiredRegisters {
		if !found {
			return nil, fmt.Errorf("can not find register \"%s\" in file \"%s\"", requiredRegisterNames[idx], file)
		}
	}

	i++
	commandSequence := make([]int, 0)
	programLineSubmatchs := programLinePattern.FindStringSubmatch(lines[i])
	if programLineSubmatchs == nil {
		return nil, fmt.Errorf("program line is invalid: \"%s\"", lines[i])
	}
	matchStr := programLineSubmatchs[1]
	literalTokens := strings.Split(matchStr, ",")
	for _, literalToken := range literalTokens {
		literalValue, err := strconv.Atoi(literalToken)
		if err != nil {
			return nil, fmt.Errorf("program line is invalid: \"%s\" is not an integer", literalToken)
		}
		commandSequence = append(commandSequence, literalValue)
	}

	initialState := programstate.Init(registers)
	return &Program{
		commandSequence:    commandSequence,
		instructionPointer: 0,
		programState:       initialState,
	}, nil
}
