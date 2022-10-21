package stuff

import (
	"fmt"
	"log"
	"os"

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
	//if scr.YOffset != 0 { // only clear the screen if the offset is not 0
	scr.Screen.Clear() // this is slow
	setFrame(scr.Screen, tcell.StyleDefault, scr.YOffset)
	//}
	_, h := scr.Screen.Size()
	txtHeight := h - info.topWidth - info.bottomWidth // height of the text area
	drawLineNumbers(scr)                              // draw the line numbers
	temp := scr.Lines[scr.YOffset : scr.YOffset+txtHeight]
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			scr.Screen.SetContent(j+info.variableWidth, i+info.topWidth, rune(temp[i][j]), nil, tcell.StyleDefault)
		}
	}
	//debug
	//draw offset at top left
	offset := fmt.Sprintf("offset: %d", scr.YOffset)
	for i := 0; i < len(offset); i++ {
		scr.Screen.SetContent(i+20, 0, rune(offset[i]), nil, tcell.StyleDefault)
	}
}
func drawLineNumbers(scr Bruh) {
	_, h := scr.Screen.Size()
	tempLine := 0
	for i := info.topWidth; i < h-(info.bottomWidth); i++ {
		temp := fmt.Sprintf("%d", i-info.topWidth+scr.YOffset)
		for j := 0; j < len(temp); j++ {
			scr.Screen.SetContent(j+info.leftWidth, i, rune(temp[j]), nil, tcell.StyleDefault)
		}
		if tempLine < len(temp) {
			tempLine = len(temp)
		}

	}
	info.variableWidth = tempLine + info.leftWidth + 1

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
func FileExplorer(scr tcell.Screen) tcell.Screen {

	//first we have to get the current directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//then we have to get the files in the directory
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	//then we have to draw the files
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(files[i].Name()); j++ {
			if j > info.leftWidth-1 {
				break
			}
			scr.SetContent(j+1, i+info.topWidth, rune(files[i].Name()[j]), nil, tcell.StyleDefault)
		}
	}
	//draw border box
	_, h := scr.Size()
	//draw top border
	for i := 1; i < info.leftWidth-1; i++ {
		scr.SetContent(i, info.topWidth-1, '─', nil, tcell.StyleDefault)
	}
	// draw right border
	for i := info.topWidth; i < h-1; i++ {
		scr.SetContent(info.leftWidth-1, i, '│', nil, tcell.StyleDefault)
	}
	scr.SetContent(info.leftWidth-1, info.topWidth-1, '┐', nil, tcell.StyleDefault)
	scr.SetContent(0, info.topWidth-1, '├', nil, tcell.StyleDefault)
	scr.SetContent(info.leftWidth-1, h-1, '┴', nil, tcell.StyleDefault)

	return scr

}
func RefreshScreen(scr Bruh) {
	scr.Screen.Clear()
	setFrame(scr.Screen, tcell.StyleDefault, scr.YOffset)
	SetText(scr)
	SetCursor(scr)
	scr.Screen.Show()

}

func SetCursor(scr Bruh) {
	x := scr.XCursor + info.variableWidth
	y := scr.YCursor + info.topWidth - scr.YOffset
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
	//draw file explorer
	info.leftWidth = x2 / 6
	s = FileExplorer(s)
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
