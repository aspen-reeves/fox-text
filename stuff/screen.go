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
	w, h := scr.Screen.Size()
	scr.Screen.Clear()
	setFrame(scr.Screen, 0, 0, w-1, h-1, tcell.StyleDefault)
	for i := 0; i < h-2; i++ {
		if i+scr.YOffset < len(scr.Lines) {
			DrawText(scr.Screen, 1, i+1, w-2, h-2, tcell.StyleDefault, scr.Lines[i+scr.YOffset])
		}
	}
}
func SetCursor(scr Bruh) {

	scr.Screen.ShowCursor(scr.XCursor+1, scr.YCursor+1)
}

// SetFrame draws the border of a frame.
func setFrame(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	for x := x1; x <= x2; x++ {
		s.SetContent(x, y1, '─', nil, style)
		s.SetContent(x, y2, '─', nil, style)
	}
	for y := y1; y <= y2; y++ {
		s.SetContent(x1, y, '│', nil, style)
		s.SetContent(x2, y, '│', nil, style)
	}
	s.SetContent(x1, y1, '┌', nil, style)
	s.SetContent(x2, y1, '┐', nil, style)
	s.SetContent(x1, y2, '└', nil, style)
	s.SetContent(x2, y2, '┘', nil, style)
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
	w, h := s.Size()
	setFrame(s, 0, 0, w-1, h-1, defStyle)
	return s
}
