package stuff

import "tcell"

func LineEnter(scr Bruh) Bruh {
	//convert cursor of screen to cursor of string

	// make temp array of strings
	temp := make([]string, len(scr.Lines)+1)
	for i := 0; i < len(temp); i++ {
		if i < scr.YCursor {
			temp[i] = scr.Lines[i]
		} else if i == scr.YCursor {
			temp[i] = scr.Lines[i][:scr.XCursor]
		} else if i == (scr.YCursor + 1) {
			temp[i] = scr.Lines[i-1][scr.XCursor:]
		} else {
			temp[i] = scr.Lines[i-1]
		}
	}
	scr.YCursor++
	scr.XCursor = 0
	scr.Lines = temp
	return scr
}
func Backspace(scr Bruh) Bruh {
	if scr.XCursor > 0 {

		temp1 := scr.Lines[scr.YCursor][:scr.XCursor-1]

		temp2 := scr.Lines[scr.YCursor][scr.XCursor:]
		scr.Lines[scr.YCursor] = temp1 + temp2
		scr.XCursor--
	} else if scr.XCursor == 0 && scr.YCursor > 0 {
		temp := len(scr.Lines[scr.YCursor])
		scr.Lines[scr.YCursor-1] = scr.Lines[scr.YCursor-1] + scr.Lines[scr.YCursor]
		for i := scr.YCursor; i < len(scr.Lines); i++ {
			if i == len(scr.Lines)-1 {
				scr.Lines[i] = ""
			} else {
				scr.Lines[i] = scr.Lines[i+1]
			}
		}
		scr.XCursor = len(scr.Lines[scr.YCursor-1]) - temp
		scr.YCursor--

	}
	return scr
}
func Insert(scr Bruh, ev *tcell.EventKey) Bruh {
	scr.Lines[scr.YCursor] = scr.Lines[scr.YCursor][:scr.XCursor] + string(ev.Rune()) + scr.Lines[scr.YCursor][scr.XCursor:]
	scr.XCursor++
	return scr
}
func Delete(scr Bruh) Bruh {
	if scr.XCursor < len(scr.Lines[scr.YCursor]) {
		//we are not erasing what the cursor is on, but what is after it
		temp1 := scr.Lines[scr.YCursor][:scr.XCursor]
		temp2 := scr.Lines[scr.YCursor][scr.XCursor:]
		scr.Lines[scr.YCursor] = temp1 + temp2
	}
	return scr
}
