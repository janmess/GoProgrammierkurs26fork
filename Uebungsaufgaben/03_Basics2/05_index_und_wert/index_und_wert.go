package index_und_wert

import "fmt"

// Index und Wert verwenden
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func LabelValues(values []string) []string {
	var result []string

	// TODO: Erzeuge für jeden Eintrag "Index: Wert".
	for index, i := range values {
		result = append(result, fmt.Sprintf("%d: %s", index, i))
	}

	return result
}
