package divegin

import routegroup "hi-supergirl/go-micro-service-example/Dive-gin/route-group"

func RouteGroup_InOneFile() {
	routegroup.CreateRoutesAndRunServerInOneFile()
}

func RouteGroup_InDifferentFile() {
	routegroup.CreateRoutesAndRunServerInDifferentFile()
}
