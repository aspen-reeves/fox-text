package stuff

func KeyDown(scr Bruh) Bruh {
	_, h := scr.Screen.Size()

	if scr.YCursor < len(scr.Lines)-1 {
		scr.YCursor++
		if scr.YCursor > h-info.bottomWidth-info.topWidth-2 {
			scr.YOffset++
		}
	}

	if scr.XCursor > len(scr.Lines[scr.YCursor]) {
		scr.XCursor = len(scr.Lines[scr.YCursor])
	}
	//check if offset is too big

	return scr
}
func KeyUp(scr Bruh) Bruh {
	if scr.YCursor > 0 {
		scr.YCursor--
	}
	if scr.YCursor == 0 && scr.YOffset > 0 {
		scr.YOffset--
	}
	if scr.XCursor > len(scr.Lines[scr.YCursor]) {
		scr.XCursor = len(scr.Lines[scr.YCursor])
	}
	return scr
}
