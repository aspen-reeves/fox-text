package stuff

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func CheckInput(scr Bruh) Bruh {
	ev := scr.Screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:

		switch ev.Key() {
		case tcell.KeyEscape:
			scr.Screen.Fini()
			os.Exit(0)
		case tcell.KeyCtrlC:
			scr.Screen.Fini()
			os.Exit(0)
		case tcell.KeyCtrlS:
			SaveFile(scr.Lines)
		case tcell.KeyEnter:
			scr = LineEnter(scr)

		case tcell.KeyBackspace, tcell.KeyBackspace2:
			scr = Backspace(scr)

		case tcell.KeyDelete:
			scr = Delete(scr)

		case tcell.KeyUp:

			scr = KeyUp(scr)

		case tcell.KeyDown:
			scr = KeyDown(scr)

		case tcell.KeyLeft:
			if scr.XCursor > 0 {
				scr.XCursor--

			} else if scr.XCursor == 0 && scr.YCursor != 0 {
				scr.XCursor = len(scr.Lines[scr.YCursor-1])
				scr.YCursor--
			}

		case tcell.KeyRight:
			if scr.XCursor < len(scr.Lines[scr.YCursor]) {
				scr.XCursor++
			} else if scr.XCursor == len(scr.Lines[scr.YCursor]) && scr.YCursor != len(scr.Lines)-1 {
				scr.XCursor = 0
				scr.YCursor++
			}
		case tcell.KeyRune:
			scr = Insert(scr, ev)
		}
	}

	return scr
}
func KeyDown(scr Bruh) Bruh {
	_, h := scr.Screen.Size()

	if scr.YCursor < len(scr.Lines)-1 {
		scr.YCursor++
		if scr.YCursor > h-info.bottomWidth-info.topWidth-1 {
			scr.YOffset++
		}
	}

	if scr.XCursor > len(scr.Lines[scr.YCursor]) {
		scr.XCursor = len(scr.Lines[scr.YCursor])
	}
	//check if offset is too big
	if scr.YOffset > len(scr.Lines)-h+info.bottomWidth+info.topWidth+1 {
		scr.YOffset = len(scr.Lines) - h + info.bottomWidth + info.topWidth + 1
	}

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
