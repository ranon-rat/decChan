package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/node/src/db"
)

func GiveInfo(w http.ResponseWriter, r *http.Request) {

	dateString := r.URL.Query().Get("date")
	if dateString == "" {
		http.Error(w, "jodete esto no es", http.StatusBadRequest)
		return
	}
	date, err := strconv.Atoi(dateString)
	if err != nil {
		http.Error(w, "jodete maldito maricon me enviaste algo que no es un numero", http.StatusBadRequest)
		return
	}

	posts, deletes := db.GetPostsSince(date), db.GetDeleteSince(date)
	json.NewEncoder(w).Encode(core.Blocks{BlocksPosts: posts,
		BlocksDeletion: deletes,
	})
}
