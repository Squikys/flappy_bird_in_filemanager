package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
)

var n int
var y_axis int = 0
var x_axis int = 0

func refresh() {

	for i := 0; i < 9; i++ {
		for j := 0; j < 16; j++ {
			t := strconv.Itoa(n)

			if i == y_axis && j == x_axis {
				player, err := os.Create("player.png")
				if err != nil {
					panic(err)
				}
				defer player.Close()
				f, err2 := os.Open("bird.png")
				if err2 != nil {
					panic(err2)
				}
				player.ReadFrom(f)
				player.Close()
				err4 := os.Rename("player.png", "game/"+t+"p.png")
				if err4 != nil {
					panic(err4)
				}

				fmt.Println("player", err)
			} else {
				file, err := os.Create("game/" + t)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer file.Close()
				fmt.Println(n)
				file.Close()

			}
			//create a new file
			n++

		}
	}
	n = 0

}
func nav() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Key press => " + string(r))
		if string(r) == "d" && x_axis < 15 {
			x_axis++
		} else if string(r) == "a" && x_axis > 0 {
			x_axis--
		} else if string(r) == "w" && y_axis > 0 {
			y_axis--
		} else if string(r) == "s" && y_axis < 8 {
			y_axis++
		}

	}
}
func main() {
	for {

		go refresh()
		dir, err := os.ReadDir("game")
		time.Sleep(1 * time.Second)
		//err2 := exec.Command("wscript.exe", "f5.vbs").Run()
		//fmt.Print(err2)

		for _, d := range dir {
			os.Remove("game/" + d.Name())
		}
		fmt.Print(err)
		go nav()

	}

}
