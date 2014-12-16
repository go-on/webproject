package pseudoroutes

import (
	"gopkg.in/go-on/method.v1"
	"gopkg.in/go-on/router.v2/route"
)

var (
	Id_        = "uuid"
	Ressource_ = "ressource"
	Item       = route.New("/:"+Ressource_+"/:"+Id_, method.GET, method.PATCH, method.DELETE)
	List       = route.New("/:"+Ressource_+"/", method.POST, method.GET)
)

func Mount(mountPoint string) {
	route.Mount(mountPoint, Item, List)
}
