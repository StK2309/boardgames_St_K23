package board

import (
	"fmt"
	"strings"
)

// String liefert eine menschenlesbare Darstellung des Spielfelds.
func (b *Board) String() string {
	if len(b.rows) == 0 {
		return ""
	}

	rows, cols := len(b.rows), len(b.rows[0])
	line := "+" + strings.Repeat("---+", cols)
	out := ""

	for r := 0; r < rows; r++ {
		out += fmt.Sprintf("%s\t%s\n", line, line) // obere Linie

		cells := "|" // Zellen und Zahlen
		numbers := "|"
		for c := 0; c < cols; c++ {
			cell := b.rows[r][c]
			if cell == "" {
				cell = " "
			}
			cells += fmt.Sprintf(" %s |", cell)

			// Nur Zahlen rechts anzeigen, wenn Zelle leer ist
			if b.rows[r][c] == " " || b.rows[r][c] == "" {
				numbers += fmt.Sprintf(" %d |", r*cols+c+1)
			} else {
				numbers += "   |" // sonst leeres Feld
			}
		}
		out += fmt.Sprintf("%s\t%s\n", cells, numbers)

	}
	out += fmt.Sprintf("%s\t%s\n", line, line) // untere Linie
	return out
}

/* if len(b.rows) == 0 {
		return ""
	}

	rows, cols := len(b.rows), len(b.rows[0])
	line := "+" + strings.Repeat("---+", cols)

	out := ""

	// obere Linie einmal
	out += fmt.Sprintf("%s\t%s\n", line, line)

	for r := 0; r < rows; r++ {

		// linke Seite: Board-Zellen
		cells := "|"
		// rechte Seite: Nummern
		numbers := "|"

		for c := 0; c < cols; c++ {
			cell := b.rows[r][c]
			if cell == "" {
				cell = " "
			}

			cells += fmt.Sprintf(" %s |", cell)

			// Nummer nur anzeigen, wenn Zelle leer ist
			if b.rows[r][c] == "" || b.rows[r][c] == " " {
				numbers += fmt.Sprintf(" %d |", r*cols+c+1)
			} else {
				numbers += "   |"
			}
		}

		// Zellenzeile
		out += fmt.Sprintf("%s\t%s\n", cells, numbers)
		// untere Linie für diese Zeile
		out += fmt.Sprintf("%s\t%s\n", line, line)
	}

	return out
} */
