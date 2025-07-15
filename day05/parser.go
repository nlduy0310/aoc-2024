package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	fileName string
}

func (p *Parser) String() string {

	return fmt.Sprintf("Parser[fileName=%s]", p.fileName)
}

func NewParser(fileName string) (*Parser, error) {

	_, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	return &Parser{fileName: fileName}, nil
}

func (p *Parser) Parse() ([]RulePair, []*Update, error) {

	file, err := os.Open(p.fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rulePairs := make([]RulePair, 0)
	updates := make([]*Update, 0)

	parsePhase := 0
scanLoop:
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			parsePhase += 1
			continue
		}

		switch parsePhase {
		case 0:
			{
				rulePair, err := NewRulePairFromString(line)
				if err != nil {
					return nil, nil, fmt.Errorf("error when parsing input file: %s", err.Error())
				}
				rulePairs = append(rulePairs, rulePair)
			}
		case 1:
			{
				update, err := NewUpdateFromString(line)
				if err != nil {
					return nil, nil, fmt.Errorf("error when parsing input file: %s", err.Error())
				}
				updates = append(updates, update)
			}
		default:
			break scanLoop
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rulePairs, updates, nil
}
