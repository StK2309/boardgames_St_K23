package board

// RowContains prüft, ob die angegebene Zeile den gesuchten Wert enthält.
func (b *Board) RowContains(i int, value string) bool {
	if i < 0 || i >= len(b.rows) {
		return false
	}
	for _, cell := range b.rows[i] {
		if cell == value {
			return true
		}
	}
	return false
}

// RowContainsChain prüft, ob die angegebene Zeile eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) RowContainsChain(rowIndex int, value string, length int) bool {
	if rowIndex < 0 || rowIndex >= len(b.rows) {
		return false
	}
	if length <= 0 {
		return true
	}
	row := b.rows[rowIndex]
	count := 0
	for _, cell := range row {
		if cell == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// RowContainsOnly prüft, ob die angegebene Zeile ausschließlich den gesuchten Wert enthält. Einzeiler!
func (b *Board) RowContainsOnly(rowIndex int, value string) bool {
	return b.RowContainsChain(rowIndex, value, len(b.rows))
}

// ColContains prüft, ob die angegebene Spalte den gesuchten Wert enthält.
func (b *Board) ColContains(colIndex int, value string) bool {
	if colIndex < 0 {
		return false
	}
	for _, row := range b.rows {
		if colIndex >= len(row) {
			continue
		}
		if row[colIndex] == value {
			return true
		}
	}
	return false
}

// ColContainsChain prüft, ob die angegebene Spalte eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) ColContainsChain(colIndex int, value string, length int) bool {
	if colIndex < 0 {
		return false
	}
	if length <= 0 {
		return true
	}
	count := 0
	for _, row := range b.rows {
		if colIndex >= len(row) {
			continue
		}
		if row[colIndex] == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// ColContainsOnly prüft, ob die angegebene Spalte ausschließlich den gesuchten Wert enthält.
func (b *Board) ColContainsOnly(colIndex int, value string) bool {
	return b.ColContainsChain(colIndex, value, len(b.rows))
}

// DiagDownRightContains prüft, ob die angegebene Diagonale von oben links nach unten rechts den gesuchten Wert enthält.
func (b *Board) DiagDownRightContains(startCol int, value string) bool {
	if startCol < 0 {
		return false
	}

	row := 0
	col := startCol

	for row < len(b.rows) {
		if col >= len(b.rows[row]) {
			break
		}
		if b.rows[row][col] == value {
			return true
		}
		row++
		col++
	}

	return false
}

// DiagDownRightContainsChain prüft, ob die angegebene Diagonale von oben links nach unten rechts eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownRightContainsChain(startCol int, value string, length int) bool {
	if startCol < 0 {
		return false
	}
	if length <= 0 {
		return true
	}

	row := 0
	col := startCol
	count := 0

	for row < len(b.rows) {
		if col >= len(b.rows[row]) {
			break
		}

		if b.rows[row][col] == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
		row++
		col++
	}
	return false
}

// DiagDownRightContainsOnly prüft, ob die angegebene Diagonale von oben links nach unten rechts ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownRightContainsOnly(startCol int, value string) bool {
	if startCol < 0 || startCol >= len(b.rows) {
		return false
	}
	row := 0
	col := startCol
	length := 0

	for row < len(b.rows) && col < len(b.rows[row]) {
		length++
		row++
		col++
	}
	return b.DiagDownRightContainsChain(startCol, value, length)
}

// DiagDownLeftContains prüft, ob die angegebene Diagonale von oben rechts nach unten links den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContains(startCol int, value string) bool {
	if startCol < 0 {
		return false
	}

	row := 0
	col := startCol

	for row < len(b.rows) && col >= 0 {
		if col >= len(b.rows[row]) {
			break
		}
		if b.rows[row][col] == value {
			return true
		}
		row++
		col--
	}

	return false
}

// DiagDownLeftContainsChain prüft, ob die angegebene Diagonale von oben rechts nach unten links eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownLeftContainsChain(startCol int, value string, length int) bool {
	if startCol < 0 {
		return false
	}
	if length <= 0 {
		return true
	}

	row := 0
	col := startCol
	count := 0

	for row < len(b.rows) && col >= 0 {
		if col >= len(b.rows[row]) {
			break
		}

		if b.rows[row][col] == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
		row++
		col--
	}
	return false
}

// DiagDownLeftContainsOnly prüft, ob die angegebene Diagonale von oben rechts nach unten links ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContainsOnly(startCol int, value string) bool {
	if startCol < 0 {
		return false
	}
	if len(b.rows) == 0 || startCol >= len(b.rows[0]) {
		return value == " "
	}

	row := 0
	col := startCol
	length := 0

	for row < len(b.rows) && col >= 0 {
		if col >= len(b.rows[row]) {
			row++
			col--
			continue
		}
		length++
		row++
		col--
	}
	if length == 0 {
		return value == " "
	}
	return b.DiagDownLeftContainsChain(startCol, value, length)
}

// BoardContains prüft, ob das Spielfeld den gesuchten Wert enthält.
func (b *Board) BoardContains(value string) bool {
	for _, row := range b.rows {
		for _, cell := range row {
			if cell == value {
				return true
			}
		}
	}
	return false
}

// BoardContainsChain prüft, ob das Spielfeld eine ununterbrochene Kette des gesuchten Werts enthält.
// Dabei werden alle Zeilen, Spalten und Diagonalen überprüft.
func (b *Board) BoardContainsChain(value string, length int) bool {
	if length <= 0 {
		return true
	}
	rows := len(b.rows)
	if rows == 0 {
		return false
	}

	cols := len(b.rows[0])
	if cols == 0 {
		return false
	}

	// Zeile prüfen
	for r := 0; r < rows; r++ {
		if b.RowContainsChain(r, value, length) {
			return true
		}
	}

	// Spalten prüfen
	for c := 0; c < cols; c++ {
		if b.ColContainsChain(c, value, length) {
			return true
		}
	}

	// Diagonale von links oben -> rechts unten
	for c := 0; c < cols; c++ {
		if b.DiagDownRightContainsChain(c, value, length) {
			return true
		}
	}

	for c := 0; c < cols; c++ {
		if b.DiagDownLeftContainsChain(c, value, length) {
			return true
		}
	}
	return false
}

// BoardContainsOnly prüft, ob das Spielfeld ausschließlich den gesuchten Wert enthält.
func (b *Board) BoardContainsOnly(value string) bool {
	for _, row := range b.rows {
		for _, cell := range row {
			if cell != value {
				return false
			}
		}
	}
	return true
}
