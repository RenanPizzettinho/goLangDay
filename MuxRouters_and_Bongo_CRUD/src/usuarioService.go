package main

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/go-bongo/bongo"
)

func usuarioConnection() *bongo.Collection {

	return dbConnection().Collection("Usuario")

}

func usuarioFindAll() []Usuario{

	usuario := &Usuario{}
	usuarios := []Usuario{}

	set := usuarioConnection().Find(bson.M{})

	for set.Next(usuario){

		usuarios = append(usuarios, *usuario)

	}

	return usuarios

}

func usuarioById(id string) *Usuario{

	usuario := &Usuario{}
	err := usuarioConnection().FindById(bson.ObjectIdHex(id), usuario)
	errHandler(err)

	return usuario

}

func usuarioDelete(id string) {

	usuario := usuarioById(id)
	err := usuarioConnection().DeleteDocument(usuario)
	errHandler(err)

}


func usuarioSave(usuario *Usuario)  {

	usuarioConnection().Save(usuario)

}
