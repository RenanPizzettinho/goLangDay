package main

import "github.com/go-bongo/bongo"

type Usuario struct {
	bongo.DocumentBase `bson:",inline"`
	Nome    string
	Login string
	Senha   string
}
