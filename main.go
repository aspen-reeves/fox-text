package main

import (
	"os"
	"tcell"
)

func run(s tcell.Screen, lines []string) {
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		printLines(s, lines)
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
	data, err := openFile()
	if err != nil {
		panic(err)
	}
	s := initScreen()
	lines := fileConvert(data)
	run(s, lines)

}
