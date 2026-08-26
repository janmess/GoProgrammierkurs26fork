package summe_bilden

// Summe bilden
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func SumTo(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}
	// TODO: Addiere 1 bis n.

	return sum
}
