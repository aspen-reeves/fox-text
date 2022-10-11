package stuff

func LineEnter(lines []string, x int, y int) []string {
	//convert cursor of screen to cursor of string
	x--
	y--
	// make temp array of strings
	temp := make([]string, len(lines)+1)
	for i := 0; i < len(temp); i++ {
		if i < y {
			temp[i] = lines[i]
		} else if i == y {
			temp[i] = lines[i][:x]
		} else if i == (y + 1) {
			temp[i] = lines[i-1][x:]
		} else {
			temp[i] = lines[i-1]
		}
	}
	return temp
}
func Backspace(lines []string, x int, y int) []string {
	if x > 1 {
		x--
		y--
		temp1 := lines[y][:x-1]

		temp2 := lines[y][x:]
		lines[y] = temp1 + temp2
	} else if x == 1 && y > 1 {
		x--
		y--
		lines[y-1] = lines[y-1] + lines[y]
		for i := y; i < len(lines)-1; i++ {
			lines[i] = lines[i+1]
		}

	}
	return lines
}
func Delete(lines []string, x int, y int) []string {
	if x < len(lines[y-1]) {
		x++ //we are not erasing what the cursor is on, but what is after it
		temp1 := lines[y-1][:x-1]
		temp2 := lines[y-1][x:]
		lines[y-1] = temp1 + temp2
	}
	return lines
}
