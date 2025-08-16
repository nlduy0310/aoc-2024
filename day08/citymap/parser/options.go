package parser

type MapParserOption = func(*MapParser)

func UseEmptyRunes(emptyRunes []rune) MapParserOption {

	return func(parser *MapParser) {
		parser.emptyLocationRunes = emptyRunes
	}
}

func UseAntennaRunes(antennaRunes []rune) MapParserOption {

	return func(parser *MapParser) {
		parser.antennaLocationRunes = antennaRunes
	}
}
