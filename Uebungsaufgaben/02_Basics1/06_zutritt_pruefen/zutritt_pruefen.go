package zutritt_pruefen

// Zutritt prüfen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func CanEnter(age int, hasID bool) bool {
	if age > 18 && hasID {
		return true
	}
	return false
}
