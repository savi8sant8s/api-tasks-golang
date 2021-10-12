package main

import "savi8sant8s/api/route"

func main(){
	app := new(route.App)
	app.Prepare()
	app.Run()
}