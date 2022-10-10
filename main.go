package main

import (
	"fox-text/stuff"
	"os"
	"tcell"
)

func run(s tcell.Screen, lines []string) {
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		stuff.PrintLines(s, lines)
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}

		}
	}
}

func main() {
	data, err := stuff.OpenFile()
	if err != nil {
		panic(err)
	}
	s := stuff.InitScreen()
	lines := stuff.FileConvert(data)
	run(s, lines)

}
