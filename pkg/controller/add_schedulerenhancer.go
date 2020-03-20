package controller

import (
	"github.com/CharlyF/schedulerenhancer/pkg/controller/schedulerenhancer"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, schedulerenhancer.Add)
}
