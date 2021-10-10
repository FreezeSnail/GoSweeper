package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Mines) ListGames(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Pulling list of Ids")
	rw.Header().Add("Content-Type", "application/json")

	keys := make([]int, len(m.gs))

	i := 0
	for k := range m.gs {
		keys[i] = k
		i++
	}

	encoder := json.NewEncoder(rw)
	err := encoder.Encode(keys)
	if err != nil {
		m.l.Println("Error encoding game list")
	}

}
