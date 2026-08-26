package status_aus_packages

// Status aus mehreren Packages
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildStatus(name string, number int) string {
	// TODO: Name bearbeiten, Zahl umwandeln und Status formatieren.
	a := strings.ToUpper(name)
	b := strconv.Itoa(number)
	text := fmt.Sprintf("%s - Aufgabe %s", a, b)
	return text
}
