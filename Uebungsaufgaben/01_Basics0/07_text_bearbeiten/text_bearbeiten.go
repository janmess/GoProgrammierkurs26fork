package text_bearbeiten

// Text mit strings bearbeiten
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import "strings"

func ChangeCase(text string) (string, string) {
	// TODO: Erzeuge Groß- und Kleinschreibung.
	x := strings.ToUpper(text)
	y := strings.ToLower(text)
	return x, y
}
