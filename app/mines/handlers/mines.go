package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	mines "github.com/freezesnail/goSweeper/app/mines/minesGame"
	"github.com/go-chi/chi"
)

/*my handler holds a game instance pointer that wont work.
I need some store of all the games like in a map i guess id -> map and I
could pull the right game obj from the map in middleware I suppose
*/
type Mines struct {
	l         *log.Logger
	gs        map[int]*mines.Game
	lastIndex int
}

func NewMines(l *log.Logger, gs map[int]*mines.Game) *Mines {
	return &Mines{l, gs, 0}
}

func getGameID(r *http.Request) int {
	// parse the product id from the url
	i := chi.URLParam(r, "id")

	// convert the id into an integer and return
	id, err := strconv.Atoi(i)
	if err != nil {
		// should never happen
		fmt.Println(i)
		panic(err)
	}

	return id
}
