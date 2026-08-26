package werte_mit_range

// Werte mit range durchlaufen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func CopyValues(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		result = append(result, num)
	}
	// TODO: Kopiere alle Werte mit range.

	return result
}
