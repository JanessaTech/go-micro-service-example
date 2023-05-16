package routegroup

import (
	multiplegroupsindifferentfile "hi-supergirl/go-micro-service-example/Dive-gin/route-group/MultipleGroupsInDifferentFile"
	multiplegroupsinonefile "hi-supergirl/go-micro-service-example/Dive-gin/route-group/MultipleGroupsInOneFile"
)

func CreateRoutesAndRunServerInOneFile() {
	multiplegroupsinonefile.CreateRoutesAndRunServer()
}

func CreateRoutesAndRunServerInDifferentFile() {
	multiplegroupsindifferentfile.CreateRoutesAndRunServer()
}
