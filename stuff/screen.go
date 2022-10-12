package stuff

import (
	"log"
	"tcell"
)

type Bruh struct {
	Lines   []string
	XCursor int
	YCursor int
	XOffset int
	YOffset int
	Screen  tcell.Screen
}
type borderInfo struct {
	topWidth     int
	bottomWidth  int
	leftWidth    int
	rightWidth   int
	lineNumWidth int
}

var info borderInfo = borderInfo{ // set the border width
	topWidth:     1,
	bottomWidth:  1,
	leftWidth:    1,
	rightWidth:   1,
	lineNumWidth: 4,
}

// DrawText draws text on the screen.
func DrawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

// SetText draws all the data on the screen
func SetText(scr Bruh) {
	//w, h := scr.Screen.Size()
	scr.Screen.Clear()
	setFrame(scr.Screen, tcell.StyleDefault)
	for i := scr.YOffset; i < len(scr.Lines); i++ {
		for j := 0; j < len(scr.Lines[i]); j++ {

			// left

			scr.Screen.SetContent(j+info.leftWidth, i+info.topWidth, rune(scr.Lines[i][j]), nil, tcell.StyleDefault)
		}

	}

}
func SetCursor(scr Bruh) {

	scr.Screen.ShowCursor(scr.XCursor+info.leftWidth, scr.YCursor+info.rightWidth)
}

// SetFrame draws the border of a frame.
func setFrame(s tcell.Screen, style tcell.Style) {
	x1 := 0
	y1 := 0
	x2, y2 := s.Size()
	x2--
	y2--

	for y := y1; y <= y2; y++ {
		s.SetContent(x1, y, '│', nil, style) // right
		s.SetContent(x2, y, '│', nil, style) // right
	}
	for x := x1; x <= x2; x++ {
		s.SetContent(x, y1, '─', nil, style) // top
		s.SetContent(x, y2, '─', nil, style) // bottom
	}
	s.SetContent(x1, y1, '┌', nil, style) // top-left
	s.SetContent(x2, y1, '┐', nil, style) // top-right
	s.SetContent(x1, y2, '└', nil, style) // bottom-left
	s.SetContent(x2, y2, '┘', nil, style) // bottom-right
}

// initScreen initializes the screen
func InitScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()
	setFrame(s, defStyle)
	return s
}
