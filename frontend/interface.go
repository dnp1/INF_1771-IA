package frontend

import (
	//"fmt"
	e "github.com/daniloanp/IA/environment"
	walk "github.com/daniloanp/IA/pathThroughMap"
	fights "github.com/daniloanp/IA/templeFights"
	allegro "github.com/dradtke/go-allegro/allegro"
	font "github.com/dradtke/go-allegro/allegro/font"
	prim "github.com/dradtke/go-allegro/allegro/primitives"
	"log"
	"strconv"
	"time"
)

const (
	FPS      int = 30
	WIDTH        = 1000
	HEIGHT       = 560
	INTERVAL     = 1700
)

func InitAllegro(env *e.Environment, paths [][]*walk.Square, fts []*fights.GameState) {
	var (
		display    *allegro.Display
		eventQueue *allegro.EventQueue
		running    bool = true
		err        error
	)
	var (
		temples = env.Temples
		saints  = env.Saints
		_       = saints
	)

	allegro.Run(func() {
		// Create a 640x480 window and give it a title.
		allegro.SetNewDisplayFlags(allegro.WINDOWED)
		var width, height = int(WIDTH), int(HEIGHT)
		if display, err = allegro.CreateDisplay(width, height); err == nil {
			defer display.Destroy()
			display.SetWindowTitle("Saint Seiya: The 12 Zodiac Temples")
		} else {
			panic(err)
		}

		// Create an event queue. All of the event sources we care about should
		// register themselves to this queue.
		if eventQueue, err = allegro.CreateEventQueue(); err == nil {
			defer eventQueue.Destroy()

		} else {
			panic(err)
		}

		// Calculate the timeout value based on the desired FPS.
		timeout := float64(1) / float64(FPS)

		// Register event sources.
		eventQueue.Register(display)

		// Set the screen to black.
		allegro.ClearToColor(allegro.MapRGB(40, 40, 40))
		var visited = make(map[e.Point]bool)

		var paint = func(inx int) {
			sq_width := float32(height / 42)
			sq_height := float32(height / 42)

			for row, line := range env.Map {
				for column, ID := range line {
					var x, y = float32(column) * sq_width, float32(row) * sq_height

					var color allegro.Color
					if visited[e.Point{row, column}] {
						color = allegro.MapRGB(200, 30, 10)
					} else {
						switch ID {
						case "M":
							color = allegro.MapRGB(40, 40, 40)
						case "P":
							color = allegro.MapRGB(100, 100, 100)
						case "R":
							color = allegro.MapRGB(200, 200, 200)
						case "S":
							// row == env.Start.Row && column == env.Start.Column:
							color = allegro.MapRGB(200, 30, 10)
						case "E": // row == env.End.Row && column == env.End.Column:
							color = allegro.MapRGB(30, 200, 10)
						case "_":
							color = allegro.MapRGB(170, 200, 56)
						}
					}

					//prim.DrawFilledRectangle(prim.Point{x, y}, prim.Point{x + 20, y + 20}, color)
					if ID != "M" {
						prim.DrawRectangle(prim.Point{x, y}, prim.Point{x + sq_width, y + sq_height}, allegro.MapRGB(250, 250, 250), 1)
						prim.DrawFilledRoundedRectangle(prim.Point{x + 1, y + 1}, prim.Point{x + sq_width - 1, y + sq_height - 1}, 3, 3, color)
					} else {
						prim.DrawRectangle(prim.Point{x, y}, prim.Point{x + sq_width, y + sq_height}, allegro.MapRGB(100, 100, 100), 2)
						prim.DrawFilledRoundedRectangle(prim.Point{x + 2, y + 2}, prim.Point{x + sq_width - 2, y + sq_height - 2}, 3, 3, color)
					}
				}

			}
			var line = 50
			sans, err := font.Builtin()
			for i := int(1); i <= inx; i++ {
				font.DrawText(sans, allegro.MapRGB(250, 250, 250), HEIGHT+10, float32(line*(i-1)+3), font.ALIGN_LEFT, temples[i-1].Name)
				var luta = fts[len(fts)-(i+1)]
				saints := ""
				for _, s := range fights.StringfyFighters(luta.Fighters, luta.Context.Saints) {
					saints += "," + s
				}
				font.DrawText(sans, allegro.MapRGB(250, 250, 250), HEIGHT+19, float32(line*(i-1)+3)+15, font.ALIGN_LEFT, saints)
				font.DrawText(sans, allegro.MapRGB(250, 250, 250), HEIGHT+10, float32(line*(i-1)+3)+30, font.ALIGN_LEFT, "\tTempo de Luta: "+strconv.FormatFloat(luta.CostToMe(), byte('g'), 3, 64))

				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		paint(0)

		allegro.FlipDisplay()

		// Main loop.
		var event allegro.Event
		var iTime = time.Now()
		var currPath = 0
		for {
			if e, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout), &event); found {
				switch e.(type) {
				case allegro.DisplayCloseEvent:
					running = false
					break
					// Handle other events here.
				}
			}

			if time.Now().Sub(iTime).Nanoseconds()/(1000*1000) > INTERVAL {
				iTime = time.Now()

				if currPath < len(paths) {
					for _, sq := range paths[currPath] {
						visited[sq.Position] = true
					}
					paint(currPath)
					allegro.FlipDisplay()
					currPath++
				}
			}
			if !running {
				return
			}
		}
	})
}
