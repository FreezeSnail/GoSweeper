package mines

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
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

func (g *Game) GetMap() [][]Tile {
	return g.b.board
}

func (g *Game) EncodeMap(rw http.ResponseWriter) {
	e := json.NewEncoder(rw)
	err := e.Encode(g.b.board)
	if err != nil {
		g.l.Println("Error encoding game list")
	}
}

type Tile struct {
	Denom   int  `json:"Denom"`
	Flagged bool `json: "Flagged"`
	Opened  bool `json: "Opened"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newTile() Tile {
	t := Tile{}
	t.Denom = 0
	t.Flagged = false
	t.Opened = false
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

func (g *Game) newBoard(x int, y int, m int) error {
	g.b = Board{}
	g.b.h = x
	g.b.w = y
	g.b.board = make([][]Tile, x)
	for i := 0; i < x; i++ {
		//b.board[i] = make([]Tile, x)
		for j := 0; j < y; j++ {
			g.b.board[i] = append(g.b.board[i], newTile())
		}
	}

	// add in mines
	mines := genMineLocations(x, y, m)

	for _, cordinate := range mines {
		g.b.at(cordinate).Denom = 9
	}
	for i := 0; i < g.b.h; i++ {
		for j := 0; j < g.b.w; j++ {
			c := Cordinate{i, j}
			adjacents := g.b.adjacentTiles(c)
			mineCount := 0
			for _, cord := range adjacents {
				if g.b.at(cord).Denom == 9 {
					mineCount += 1
				}
			}
			t := g.b.at(c)
			if t.Denom != 9 {
				t.Denom = mineCount
			}
		}
	}

	return nil
}

func (b Board) flagTile(c Cordinate) error {
	if b.at(c).Opened {
		err := errors.New("This tile is Opened already, cannot flag")
		// log it heugh
		return err
	}
	b.at(c).Flagged = !b.at(c).Flagged
	return nil
}

func (g *Game) OpenTile(c Cordinate) bool {
	if g.b.at(c).Opened {
		// already did this one
		return true
	}

	g.b.at(c).Opened = true
	switch g.b.at(c).Denom {
	case 0:
		{
			adjacents := g.b.adjacentTiles(c)
			for _, cord := range adjacents {
				e := g.OpenTile(cord)
				if !e {
					//this is fucked hit a mine
					panic("fugma Opened a mine automatically")
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
			if b.at(Cordinate{x, y}).Opened {
				fmt.Printf("| %d |", b.board[x][y].Denom)
			} else {
				fmt.Printf("|   |")
			}
		}
		fmt.Println()
	}
}

func (g *Game) Run() {
	/*	err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()

		termbox.SetOutputMode(termbox.OutputRGB)
	*/
	var err error
	err = g.newBoard(8, 6, 5)
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
