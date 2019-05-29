package controller

import (
	"github.com/pratikjagrut/api-status-operator/pkg/controller/apistatus"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, apistatus.Add)
}
