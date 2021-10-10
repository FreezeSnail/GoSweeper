package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mine "github.com/freezesnail/goSweeper/app/mines/minesGame"
)

func (m *Mines) Open(rw http.ResponseWriter, r *http.Request) {
	cords := mux.Vars(r)
	id := getGameID(r)
	x, _ := strconv.Atoi(cords["x"])
	y, _ := strconv.Atoi(cords["y"])
	cord := mine.MakeCord(x, y)
	m.gs[id].OpenTile(cord)
}

func (m *Mines) NewGame(rw http.ResponseWriter, r *http.Request) {
	i := m.lastIndex + 1
	m.gs[i] = &mine.Game{}
	m.l.Println("Launching new game: ", i)
	m.lastIndex = i
	go m.gs[i].Run()
}
