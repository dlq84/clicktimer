// clicktimer
package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "clicktimer"
var winWidth, winHeight int32 = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	var peepArray []sdl.Event = make([]sdl.Event, 1)
	id, err := window.GetID()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get window ID: %s\n", err)
		return 3
	}
	peepArray[0] = &sdl.UserEvent{
		Type:      sdl.USEREVENT,
		Timestamp: sdl.GetTicks(),
		WindowID:  id,
		Code:      1,
		Data1:     nil,
		Data2:     nil,
	}

	running = true
	lastPushTime := sdl.GetTicks()

	for running {
		fmt.Println("running, wait for yellow!")
		drawBlue(renderer)
		min := 2000
		max := min + 4000
		delay := uint32(rand.Intn(max-min) + min)
		lastPushTime = sdl.GetTicks()
		started := false
		var startTick uint32
		waitForClicks := true

		for waitForClicks {
			if !started && lastPushTime+delay < sdl.GetTicks() {
				started = true
				startTick = sdl.GetTicks()
				drawYellow(renderer)
			}
			for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
					started = false
					break
				case *sdl.MouseButtonEvent:
					if t.State == sdl.PRESSED {
						if started {
							reactionTime := t.Timestamp - startTick
							if reactionTime == 69 {
								fmt.Println("Nice")
							} else if reactionTime == 420 {
								fmt.Println("Blaze it!")
							} else if reactionTime < 100 {
								fmt.Println("Godlike!", reactionTime, "ms")
							} else if reactionTime < 150 {
								fmt.Println("Epic!", reactionTime, "ms")
							} else if reactionTime < 200 {
								fmt.Println("Decent!", reactionTime, "ms")
							} else if reactionTime < 300 {
								fmt.Println("Average!", reactionTime, "ms")
							} else {
								fmt.Println("You're literally garbage!", reactionTime, "ms")
							}
							drawGreen(renderer)
							sdl.Delay(2000)
							started = false
							waitForClicks = false
							break
						} else {
							drawRed(renderer)
							fmt.Println("FAIL! Wait for yellow!")
							sdl.Delay(2000)
							started = false
							waitForClicks = false
							break
						}
					}
				}
			}
			sdl.Delay(1)
		}
	}
	return 0
}

func drawRed(renderer *sdl.Renderer) {
	renderer.Clear()
	rect := sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&rect)
	renderer.Present()
}

func drawGreen(renderer *sdl.Renderer) {
	renderer.Clear()
	rect := sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight}
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)
	renderer.Present()
}

func drawBlue(renderer *sdl.Renderer) {
	renderer.Clear()
	rect := sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight}
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(&rect)
	renderer.Present()
}

func drawYellow(renderer *sdl.Renderer) {
	renderer.Clear()
	rect := sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight}
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.FillRect(&rect)
	renderer.Present()
}

func main() {
	os.Exit(run())
}
