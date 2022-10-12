package stuff

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

//TODO color

type Bruh struct {
	Lines   []string
	XCursor int
	YCursor int
	XOffset int
	YOffset int
	Screen  tcell.Screen
}
type borderInfo struct {
	topWidth      int
	bottomWidth   int
	leftWidth     int
	variableWidth int
	rightWidth    int
}

var info borderInfo = borderInfo{ // set the border width
	topWidth:      5,
	bottomWidth:   1,
	leftWidth:     1,
	rightWidth:    1,
	variableWidth: 1, // should start at leftWidth, added by the line numbers
}

var asciiFox []string = []string{
	// split into runes so go doesn't complain
	" \\     |\\__/|    / ",
	"  \\   /     \\   /  ",
	"   \\ /_.~ ~,_\\ /   ",
	"    \\   \\@/   /    ",
}

// SetText draws all the data on the screen
func SetText(scr Bruh) {
	if scr.YOffset != 0 { // only clear the screen if the offset is not 0
		scr.Screen.Clear() // this is slow
		setFrame(scr.Screen, tcell.StyleDefault, scr.YOffset)
	}
	//setFrame(scr.Screen, tcell.StyleDefault, scr.YOffset) // maybe i shouldnt draw the border every time
	_, h := scr.Screen.Size()

	temp := scr.Lines[scr.YOffset:]
	for i := 0; i < len(temp); i++ {
		//we will make line numbers here
		lineNum := fmt.Sprintf("%d", i+scr.YOffset+1)
		for j := 0; j < len(temp[i]); j++ {
			if j < len(lineNum) {
				scr.Screen.SetContent(j+info.leftWidth, i+info.topWidth, rune(lineNum[j]), nil, tcell.StyleDefault)

			}
			info.variableWidth = len(lineNum) + info.leftWidth + 1
			scr.Screen.SetContent(j+info.variableWidth, i+info.topWidth, rune(temp[i][j]), nil, tcell.StyleDefault)
		}
		if i >= (h - info.bottomWidth) {
			break
		}
	}
}

// this will be a simpler function to draw the screen, to increase performance
func InsertLine(scr *Bruh) {
	i := scr.YCursor
	temp := scr.Lines[i]
	for j := 0; j < len(temp); j++ {
		scr.Screen.SetContent(j+info.variableWidth+1, i+info.topWidth, rune(temp[j]), nil, tcell.StyleDefault)
	}

}

// function to convert string cursor to absolute cursor
func cursorStrToAbs(scr Bruh) (int, int) {
	x := scr.XCursor + info.variableWidth
	y := scr.YCursor + info.topWidth

	return x, y
}

func SetCursor(scr Bruh) {
	x := scr.XCursor + info.variableWidth
	y := scr.YCursor + info.topWidth
	_, h := scr.Screen.Size()
	if y >= h-info.bottomWidth {
		y = h - info.bottomWidth
	}

	scr.Screen.ShowCursor(x, y)
}

// SetFrame draws the border of a frame.
func setFrame(s tcell.Screen, style tcell.Style, offset int) {
	x1 := 0
	y1 := 0
	x2, y2 := s.Size()
	x2--
	y2--
	for y := y1; y <= y2; y++ {
		s.SetContent(x1, y, '│', nil, style) // right
		s.SetContent(x2, y, '│', nil, style) // right
	}
	// write line numbers
	/*for i := 0; i < y2; i++ {
		temp := fmt.Sprintf("%d", i+offset+1)
		for j := 0; j < len(temp); j++ {
			s.SetContent(j+info.leftWidth, i+info.topWidth, rune(temp[j]), nil, style)
		}
		if len(temp) > info.variableWidth {
			info.variableWidth = len(temp) + info.leftWidth + 1
		}

	}*/

	for x := x1; x <= x2; x++ {
		s.SetContent(x, y1, '─', nil, style) // top
	}
	// write the fox to the center top of the screen
	for i := 0; i < len(asciiFox); i++ {
		for j := 0; j < len(asciiFox[i]); j++ {
			s.SetContent(j+x2/2-len(asciiFox[i])/2, i+y1, rune(asciiFox[i][j]), nil, style)
		}
	}

	for x := x1; x <= x2; x++ {
		s.SetContent(x, y2, '─', nil, style) // bottom
	}
	s.SetContent(x1, y1, '┌', nil, style) // top-left
	s.SetContent(x2, y1, '┐', nil, style) // top-right
	s.SetContent(x1, y2, '└', nil, style) // bottom-left
	s.SetContent(x2, y2, '┘', nil, style) // bottom-right
	//debug
	//output w and h of the screen
	temp := fmt.Sprintf("w: %d, h: %d", x2, y2)
	for i := 0; i < len(temp); i++ {
		s.SetContent(i, 0, rune(temp[i]), nil, style)
	}
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
	setFrame(s, defStyle, 0)
	return s
}
