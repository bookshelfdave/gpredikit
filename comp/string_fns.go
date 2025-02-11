package comp

// This assumes that the input is always double quoted, as they will be from Antlr
func StripFirstAndLast(s string) string {
	runeInput := []rune(s)
	if len(runeInput) >= 2 {
		runeInput = runeInput[1 : len(runeInput)-1]
	}
	return string(runeInput)
}
