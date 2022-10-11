package main

import (
	"fox-text/stuff"
	"os"
)

/*

.___                              __
|   | _____ ______   ____________/  |_  ______
|   |/     \\____ \ /  _ \_  __ \   __\/  ___/
|   |  Y Y  \  |_> >  <_> )  | \/|  |  \___ \
|___|__|_|  /   __/ \____/|__|   |__| /____  >
          \/|__|                           \/
*/
import (
	"tcell"
)

func checkInput(scr stuff.Bruh) stuff.Bruh {
	ev := scr.Screen.PollEvent()
	_, h := scr.Screen.Size()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		/*
			///// /   / ///// /   / /////
			/      / /  /     //  /   /
			/////  / /  ///// / / /   /
			/      / /  /     /  //   /
			/////   /   ///// /   /   /
		*/
		switch ev.Key() {
		case tcell.KeyEscape:
			scr.Screen.Fini()
			os.Exit(0)
		case tcell.KeyCtrlC:
			scr.Screen.Fini()
			os.Exit(0)
		case tcell.KeyCtrlS:
			stuff.SaveFile(scr.Lines)
		case tcell.KeyEnter:
			scr = stuff.LineEnter(scr)

		case tcell.KeyBackspace, tcell.KeyBackspace2:
			scr = stuff.Backspace(scr)

		case tcell.KeyDelete:
			scr = stuff.Delete(scr)

		case tcell.KeyUp:

			if scr.YCursor > 1 {
				scr.YCursor--
			} else if scr.YCursor == 1 && scr.YOffset > 0 {
				scr.YOffset--
			}
			if scr.XCursor > len(scr.Lines[scr.YCursor]) {
				scr.XCursor = len(scr.Lines[scr.YCursor])
			}

		case tcell.KeyDown:
			if scr.YCursor < len(scr.Lines)-1 {
				scr.YCursor++
			} else if scr.YCursor >= h {
				scr.YOffset++
			}
			if scr.XCursor > len(scr.Lines[scr.YCursor]) {
				scr.XCursor = len(scr.Lines[scr.YCursor])
			}

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
			scr = stuff.Insert(scr, ev)
		}
	}

	return scr
}

/*
			 ,////,
			 /// 6|
			 //  _|
			_/_,-'
	   _.-/'/   \   ,/;,

,-' /'  \_   \ / _/
`\ /     _/\  ` /

	|     /,  `\_/
	|     \'

	pb  /\_        /`      /\
	  /' /_``--.__/\  `,. /  \
	 |_/`  `-._     `\/  `\   `.
	           `-.__/'     `\   |
	                         `\  \
	                           `\ \
	                             \_\__
	                              \___)
*/
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
	for {
		stuff.SetText(scr)
		stuff.SetCursor(scr)
		s.Show()
		scr = checkInput(scr)
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
