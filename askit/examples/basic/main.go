package main

import (
	"fmt"

	"github.com/neatplatform/craft/askit"
)

func main() {
	asker := askit.NewAsker()

	info := struct {
		ID      int    `ask:"any, your identification number"`
		Name    string `ask:"any, your full name"`
		Contact struct {
			Email string `ask:"email, your email address"`
		}
		Secrets struct {
			Token string `ask:"secret, your access token"`
		}
	}{
		ID:   1,
		Name: "Jane Doe",
	}

	err := askit.Ask(&info, asker)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", info)
}
