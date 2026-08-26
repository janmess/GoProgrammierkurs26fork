package linkedlist

// Übung 10 – Verkettete Liste
// Schwierigkeit: ★★★★★
//
// Lernziele:
//   - Pointer als Verbindungen zwischen Structs
//   - nil als Ende einer Datenstruktur
//   - Eine kleine dynamische Datenstruktur selbst implementieren
//
// Ein Node sieht im Speicher konzeptionell so aus:
//
//   +---------+---------+      +---------+---------+
//   | Value 3 | Next ---+----->| Value 8 | Next --|----> nil
//   +---------+---------+      +---------+---------+
//
// `Next *Node` enthält also nicht den nächsten Node selbst, sondern dessen Adresse.

type Node struct {
	Value int
	Next  *Node
}

// Length zählt die Nodes ab head.
func Length(head *Node) int {
	// TODO:
	// current := head
	// for current != nil { ...; current = current.Next }
	current := head
	num := 0
	for current != nil {
		num++
		current = current.Next
	}
	return num
}

// Sum addiert alle Werte der Liste.
func Sum(head *Node) int {
	// TODO
	sum := 0
	current := head
	for current != nil {
		sum += current.Value
		current = current.Next
	}
	return sum
}

// Find gibt einen Pointer auf den ERSTEN Node mit dem gesuchten Wert zurück.
// Wenn der Wert nicht existiert, soll nil zurückgegeben werden.
func Find(head *Node, value int) *Node {
	// TODO
	current := head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Append hängt einen neuen Node ans Ende an und gibt den (möglicherweise neuen)
// Listenanfang zurück.
//
// Sonderfall:
// Ist head == nil, besteht die neue Liste nur aus dem neuen Node.
func Append(head *Node, value int) *Node {
	// TODO
	if head == nil {
		return &Node{
			Value: value,
			Next:  nil,
		}
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: value, Next: nil}
	return head
}

// Zusatzaufgabe (ohne Test):
// Implementiere anschließend:
//
//     func Prepend(head *Node, value int) *Node
//
// Diese Funktion ist überraschend kurz.
