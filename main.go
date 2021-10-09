package main

import (
	"savi8sant8s/api/route"
)

func main(){
	
	var route = new(route.AppRoute)
	route.PrepareRoutes()
	route.Run()
}