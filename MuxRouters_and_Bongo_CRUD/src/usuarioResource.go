package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func usuarioResourceList(w http.ResponseWriter, _ *http.Request) {

	all := usuarioFindAll()

	json.NewEncoder(w).Encode(all)

}

func usuarioResourceFindById(w http.ResponseWriter, r *http.Request) {

	byId := usuarioById(mux.Vars(r)["id"])

	json.NewEncoder(w).Encode(byId)

}

func usuarioResourcePost(_ http.ResponseWriter, r *http.Request) {

	usuario := usuarioFromJson(r)
	usuarioSave(usuario)

}

func usuarioResourcePut(w http.ResponseWriter, r *http.Request) {

	byId := usuarioById(mux.Vars(r)["id"])
	usuario := usuarioFromJson(r)

	if byId.Id == usuario.Id {

		usuarioSave(usuario)

	} else {

		w.WriteHeader(http.StatusBadRequest)

	}

}

func usuarioResourceDelete(w http.ResponseWriter, r *http.Request) {

	usuarioDelete(mux.Vars(r)["id"])

	w.WriteHeader(http.StatusNoContent)

}

func usuarioFromJson(r *http.Request) *Usuario {

	body, err := ioutil.ReadAll(r.Body)
	errHandler(err)
	usuario := &Usuario{}
	json.Unmarshal(body, usuario)

	return usuario

}
