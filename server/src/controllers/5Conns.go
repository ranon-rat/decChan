package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"

	"github.com/ranon-rat/decChan/core"
)

// i should delete this later but its actually handly
func GetRandomConns() []core.ConnIP {
	// at least it should return 5 if we have more than 5 connections
	rand5 := make(map[core.ConnIP]bool) // i want non repeated
	keys := reflect.ValueOf(listConns).MapKeys()

	for len(rand5) < 5 && len(rand5) < len(listConns) {
		val := keys[rand.Intn(len(keys))].Interface().(core.ConnIP)

		rand5[val] = true
	}
	rando := []core.ConnIP{}
	for v := range rand5 {
		rando = append(rando, v)
	}
	return rando
}
func Conns5(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetRandomConns())

}
