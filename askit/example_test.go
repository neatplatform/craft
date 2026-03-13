package askit_test

import "github.com/neatplatform/craft/askit"

func ExampleAsk() {
	asker := askit.NewAsker()

	info := struct {
		Name  string `ask:"any, your full name"`
		Email string `ask:"email, your email address"`
		Token string `ask:"secret, your access token"`
	}{}

	err := askit.Ask(&info, asker)
	if err != nil {
		panic(err)
	}
}
