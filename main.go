package main

import (
	// "example.com/m/loaders"
	//"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

// Go langs entry point
func main() {
	pixelgl.Run(run)
}

var (
	currentMap     *Map
	//conwayManager  *ConwayManager
	randomInstance *rand.Rand
)

func run() {

	window := createWindow()
	rand.Seed(time.Now().UnixNano())

	//INIT globals
	randomInstance = rand.New(rand.NewSource(time.Now().UnixNano()))

	camera := new(Camera)
	camera.Init()

	player := new(Player)
	player.SetSprite(GetSprite("img_person_0.png"))
	player.Init()

	camera.chaseObject = player

	last := time.Now()
	//frameTime := time.Now()
	//conwayManager = new(ConwayManager)
	//conwayManager.Init()

	currentMap = new(Map)
	currentMap.CreateGrid(pixel.V(64, 64))
	generate(currentMap)

	for !window.Closed() {
		// Clear window from the last frame.
		window.Clear(colornames.Yellow)
		// Update delta time
		dt := time.Since(last).Seconds()
		last = time.Now()

		player.Update(window, dt)
		camera.Update(window)

		if window.JustPressed(pixelgl.KeyEscape) {
			window.SetClosed(true)
		}

		if window.JustPressed(pixelgl.KeyR) {
			currentMap = new(Map)
			currentMap.CreateGrid(pixel.V(32, 32))
			generate(currentMap)

			currentMap.Draw(window)
			window.Update()

			time.Sleep(time.Second)

			for i := 0; i < 5; i++ {
				window.Clear(colornames.Yellow)
				for x := 0; x < int(currentMap.size.X); x++ {
					for y := 0; y < int(currentMap.size.Y); y++ {
						currentMap.grid[x][y].Tick()
					}
					currentMap.Draw(window)
					window.Update()
				}
			}
		}

		currentMap.Draw(window)
		camera.Draw(window, dt)
		player.Draw(window)

		// print out FPS
		//fmt.Println(float64(1) / time.Since(frameTime).Seconds())
		//frameTime = time.Now()

		window.Update()
	}
}

//https://github.com/faiface/pixel