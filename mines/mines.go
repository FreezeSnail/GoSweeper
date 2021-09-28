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
	board []Tile
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

func (b Board) flattenCordinate(c Cordinate) int {
	index := b.w*c.x + c.y
	return index
}

func buildCord(y int, i int) Cordinate {
	return Cordinate{(i % y), (i / y)}
}

func randCord(x int, y int) Cordinate {
	c := Cordinate{
		rand.Intn(x),
		rand.Intn(y),
	}

	return c
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

	fmt.Printf("%v", cordinates)
	return cordinates
}

func newBoard(x int, y int, m int) (Board, error) {
	b := Board{}
	b.h = x
	b.w = y

	for i := 0; i < x*y; i++ {
		b.board = append(b.board, newTile())
	}

	// add in mines
	mines := genMineLocations(x, y, m)

	for _, cordinate := range mines {
		index := b.flattenCordinate(cordinate)
		b.board[index].denom = 9
	}

	return b, nil
}

func (b Board) flagTile(i int) error {
	if b.board[i].opened {
		err := errors.New("This tile is opened already, cannot flag")
		// log it heugh
		return err
	}
	b.board[i].flagged = !b.board[i].flagged
	return nil
}

func (b Board) openTile(c Cordinate) {
	i := b.flattenCordinate(c)
	if b.board[i].opened {
		// already did this one
		return
	}
	if b.board[i].denom == 9 {
		// bomb
	}
	b.board[i].opened = true
	if b.board[i].denom == 0 {
		start, end := b.adjacentTiles(c)
		for i := start; start <= end; i++ {
			b.openTile(buildCord(b.w, i))
		}
		// open all adjacent tiles ?
	}

}

func (b Board) adjacentTiles(c Cordinate) (int, int) {
	var xTop int
	var xBottom int
	var yLeft int
	var yRight int

	if (c.x - 3) >= 0 {
		xTop = c.x - 3
	} else {
		xTop = 0
	}

	if (c.x + 3) < b.h {
		xBottom = c.x + 3
	} else {
		xBottom = b.h - 1
	}

	if (c.y - 3) >= 0 {
		yLeft = c.y - 3
	} else {
		yLeft = 0
	}

	if (c.y + 3) < b.w {
		yRight = c.y + 3
	} else {
		yRight = b.w - 1
	}

	start := b.flattenCordinate(Cordinate{xTop, yLeft})
	end := b.flattenCordinate(Cordinate{xBottom, yRight})

	return start, end
}

func (b *Board) printBoard() {
	fmt.Println("x: ", b.h, "y, ", b.w)
	for x := 0; x < b.h; x++ {
		for y := 0; y < b.w; y++ {
			i := y*y + x
			if b.board[i].opened {
				fmt.Printf("| %d |", b.board[i].denom)
			} else {
				fmt.Printf("|   |")
			}
		}
		fmt.Println()
	}
}

func Run() {
	game := Game{}
	var err error
	game.b, err = newBoard(5, 6, 4)
	if err != nil {
		panic("ded")
	}
	fmt.Println("board made")
	game.b.printBoard()

}
