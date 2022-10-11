package main

import (
	"fox-text/stuff"
	"os"
	"tcell"
)

func checkInput(s tcell.Screen, ev tcell.Event, lines []string, begin int, x int, y int) ([]string, int, int, int) {
	_, h := s.Size()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			s.Fini()
			os.Exit(0)
		case tcell.KeyCtrlC:
			s.Fini()
			os.Exit(0)
		case tcell.KeyCtrlS:
			stuff.SaveFile(lines)
		case tcell.KeyEnter:
			lines = stuff.LineEnter(lines, x, y)
			//set cursor
			y++
			x = 1

		case tcell.KeyBackspace, tcell.KeyBackspace2:
			lines = stuff.Backspace(lines, x, y)
			if x > 1 {
				x--
			}

		case tcell.KeyDelete:
			lines = stuff.Delete(lines, x, y)

		case tcell.KeyUp:

			if y > 1 {
				y--
			} else if y == 1 && begin > 0 {
				begin--
			}

		case tcell.KeyDown:
			if y < len(lines) {
				y++
			} else if y >= h {
				begin++
			}

		case tcell.KeyLeft:
			if x > 1 {
				x--
			}

		case tcell.KeyRight:
			if x < len(lines[y-1])+1 {
				x++
			}
		case tcell.KeyRune:
			lines[y-1] = lines[y-1][:(x-1)] + string(ev.Rune()) + lines[y-1][(x-1):]
			x++
		}
	}
	if x > len(lines[y-1]) {
		x = len(lines[y-1]) + 1
	}

	return lines, x, y, begin
}

func run(s tcell.Screen, lines []string) {
	//w, h := s.Size()
	begin := 0
	//set cursor
	x, y := 1, 1

	for {
		stuff.SetText(s, lines, begin)
		stuff.SetCursor(s, x, y)
		s.Show()
		lines, x, y, begin = checkInput(s, s.PollEvent(), lines, begin, x, y)
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
