package frontend

//import (
//	allegro "github.com/dradtke/go-allegro/allegro"
//	prim "github.com/dradtke/go-allegro/allegro/primitives"
//)

//const (
//	FPS int = 30
//)

//func initAllegro(env *Environment) {
//	var (
//		display    *allegro.Display
//		eventQueue *allegro.EventQueue
//		running    bool = true
//		err        error
//	)

//	allegro.Run(func() {
//		// Create a 640x480 window and give it a title.
//		allegro.SetNewDisplayFlags(allegro.WINDOWED)
//		var width, height = int(840), int(840)
//		if display, err = allegro.CreateDisplay(width, height); err == nil {
//			defer display.Destroy()
//			display.SetWindowTitle("Saint Seiya: The 12 Zodiac Temples")
//		} else {
//			panic(err)
//		}

//		// Create an event queue. All of the event sources we care about should
//		// register themselves to this queue.
//		if eventQueue, err = allegro.CreateEventQueue(); err == nil {
//			defer eventQueue.Destroy()

//		} else {
//			panic(err)
//		}

//		// Calculate the timeout value based on the desired FPS.
//		timeout := float64(1) / float64(FPS)

//		// Register event sources.
//		eventQueue.Register(display)

//		// Set the screen to black.
//		allegro.ClearToColor(allegro.MapRGB(40, 40, 40))

//		sq_width := float32(width / 42)
//		sq_height := float32(height / 42)

//		for row, line := range env.Map {
//			for column, ID := range line {
//				var x, y = float32(column) * sq_width, float32(row) * sq_height

//				var color allegro.Color

//				switch ID {
//				case "M":
//					color = allegro.MapRGB(40, 40, 40)
//				case "P":
//					color = allegro.MapRGB(100, 100, 100)
//				case "R":
//					color = allegro.MapRGB(200, 200, 200)
//				case "S":
//					// row == env.Start.Row && column == env.Start.Column:
//					color = allegro.MapRGB(200, 30, 10)
//				case "E": // row == env.End.Row && column == env.End.Column:
//					color = allegro.MapRGB(30, 200, 10)
//				case "_":
//					color = allegro.MapRGB(170, 200, 56)
//				}

//				//prim.DrawFilledRectangle(prim.Point{x, y}, prim.Point{x + 20, y + 20}, color)
//				if ID != "M" {
//					prim.DrawRectangle(prim.Point{x, y}, prim.Point{x + sq_width, y + sq_height}, allegro.MapRGB(250, 250, 250), 1)
//					prim.DrawFilledRoundedRectangle(prim.Point{x + 1, y + 1}, prim.Point{x + sq_width - 1, y + sq_height - 1}, 3, 3, color)
//				} else {
//					prim.DrawRectangle(prim.Point{x, y}, prim.Point{x + sq_width, y + sq_height}, allegro.MapRGB(100, 100, 100), 2)
//					prim.DrawFilledRoundedRectangle(prim.Point{x + 2, y + 2}, prim.Point{x + sq_width - 2, y + sq_height - 2}, 3, 3, color)
//				}

//			}
//		}

//		allegro.FlipDisplay()

//		// Main loop.
//		var event allegro.Event
//		for {
//			if e, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout), &event); found {
//				switch e.(type) {
//				case allegro.DisplayCloseEvent:
//					running = false
//					break
//					// Handle other events here.
//				}
//			}

//			if !running {
//				return
//			}
//		}
//	})
//}
