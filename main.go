package main

import (
	"github.com/jeyog/seemeetcan-server/app"
)

func init() {

}

func main() {
	app := app.GetInstance()
	app.Run()
}
