package mines

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

type Game struct {
	l log.Logger
	b Board
}

type Board struct {
	l     *log.Logger
	board [][]Tile
	h     int
	w     int
}

type Tile struct {
	denom   int
	flagged bool
	opened  bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newTile() Tile {
	t := Tile{}
	t.denom = 0
	t.flagged = false
	t.opened = false
	return t
}

type Cordinate struct {
	x int
	y int
}

func randCord(x int, y int) Cordinate {
	c := Cordinate{
		rand.Intn(x),
		rand.Intn(y),
	}

	return c
}

func MakeCord(x, y int) Cordinate {
	return Cordinate{x, y}
}

func hasCord(cordinates []Cordinate, c Cordinate) bool {
	for _, value := range cordinates {
		if c == value {
			return true
		}
	}
	return false
}

func genMineLocations(x int, y int, n int) []Cordinate {
	count := 0
	tries := 0
	cordinates := make([]Cordinate, n)
	for count < n {
		c := randCord(x, y)
		if !hasCord(cordinates, c) {
			cordinates = append(cordinates, c)
			count += 1
		}
		tries++
		if tries > 100 {
			// fugma
			fmt.Println("100 tries")
			return cordinates
		}
	}

	return cordinates
}

func (b Board) at(c Cordinate) *Tile {
	return &b.board[c.x][c.y]
}

func newBoard(x int, y int, m int) (Board, error) {
	b := Board{}
	b.h = x
	b.w = y
	b.board = make([][]Tile, x)
	for i := 0; i < x; i++ {
		//b.board[i] = make([]Tile, x)
		for j := 0; j < y; j++ {
			b.board[i] = append(b.board[i], newTile())
		}
	}

	// add in mines
	mines := genMineLocations(x, y, m)

	for _, cordinate := range mines {
		b.at(cordinate).denom = 9
	}
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			c := Cordinate{i, j}
			adjacents := b.adjacentTiles(c)
			mineCount := 0
			for _, cord := range adjacents {
				if b.at(cord).denom == 9 {
					mineCount += 1
				}
			}
			t := b.at(c)
			if t.denom != 9 {
				t.denom = mineCount
			}
		}
	}

	return b, nil
}

func (b Board) flagTile(c Cordinate) error {
	if b.at(c).opened {
		err := errors.New("This tile is opened already, cannot flag")
		// log it heugh
		return err
	}
	b.at(c).flagged = !b.at(c).flagged
	return nil
}

func (g Game) OpenTile(c Cordinate) bool {
	if g.b.at(c).opened {
		// already did this one
		return true
	}

	g.b.at(c).opened = true
	switch g.b.at(c).denom {
	case 0:
		{
			adjacents := g.b.adjacentTiles(c)
			for _, cord := range adjacents {
				e := g.OpenTile(cord)
				if !e {
					//this is fucked hit a mine
					panic("fugma opened a mine automatically")
				}
			}
		}
	case 9:
		{
			// bomb go boom
			return false
		}
		// open all adjacent tiles ?
	}

	return true
}

func (b Board) adjacentTiles(c Cordinate) []Cordinate {
	var xTop int
	var xBottom int
	var yLeft int
	var yRight int

	if (c.x - 1) >= 0 {
		xTop = c.x - 1
	} else {
		xTop = 0
	}

	if (c.x + 1) < b.h {
		xBottom = c.x + 1
	} else {
		xBottom = b.h - 1
	}

	if (c.y - 1) >= 0 {
		yLeft = c.y - 1
	} else {
		yLeft = 0
	}

	if (c.y + 1) < b.w {
		yRight = c.y + 1
	} else {
		yRight = b.w - 1
	}
	var cords []Cordinate
	for i := xTop; i <= xBottom; i++ {
		for j := yLeft; j <= yRight; j++ {
			cords = append(cords, Cordinate{i, j})
		}
	}

	return cords
}

func (b *Board) printBoard() {
	for x := 0; x < b.h; x++ {
		for y := 0; y < b.w; y++ {
			if b.at(Cordinate{x, y}).opened {
				fmt.Printf("| %d |", b.board[x][y].denom)
			} else {
				fmt.Printf("|   |")
			}
		}
		fmt.Println()
	}
}

func (g Game) Run() {
	/*	err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()

		termbox.SetOutputMode(termbox.OutputRGB)
	*/
	var err error
	g.b, err = newBoard(8, 6, 5)
	if err != nil {
		panic("fugma")
	}

	for true {

	}
	/*	eventQueue := make(chan termbox.Event)
		go func() {
			for {
				eventQueue <- termbox.PollEvent()
			}
		}()

		for {
			select {
			case ev := <-eventQueue:
				if ev.Type == termbox.EventKey {
					switch {

					case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
						return
					}
				}
			default:
				Render(&g)
			}
		}
	*/

}
