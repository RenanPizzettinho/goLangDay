package main

import "fmt"

func errHandler(err error) {

	if err != nil {

		fmt.Println(err)

	}

}
