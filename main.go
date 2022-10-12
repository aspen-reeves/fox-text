package main

//   |\__/|
//  /     \
// /_.~ ~,_\
//    \@/
// welcom to fox-text
import (
	"vulpes-text/stuff"

	"github.com/gdamore/tcell/v2"
)

func run(s tcell.Screen, lines []string) {
	var scr stuff.Bruh
	scr = stuff.Bruh{
		Lines:   lines,
		XCursor: 0,
		YCursor: 0,
		XOffset: 0,
		YOffset: 0,
		Screen:  s,
	}
	stuff.SetText(scr)
	for {
		stuff.SetText(scr)
		stuff.SetCursor(scr)
		s.Show()
		scr = stuff.CheckInput(scr)
	}
}

func main() {
	data, err := stuff.OpenFile()
	if err != nil {
		panic(err)
	}
	s := stuff.InitScreen()
	lines := stuff.ByteToStr(data)
	run(s, lines)

}
