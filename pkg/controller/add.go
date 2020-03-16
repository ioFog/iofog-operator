package controller

import (
	"github.com/eclipse-iofog/iofog-operator/v2/pkg/controller/app"
	"github.com/eclipse-iofog/iofog-operator/v2/pkg/controller/kog"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, app.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, kog.Add)
}
