package controller

import (
	commonController "promise/common/controller"
)

// StudentRootController is the root controller of the student.
type StudentRootController struct {
	commonController.PromiseRootController
}

// GetResourceName return the name of the resource this controller deal with.
func (c *StudentRootController) GetResourceName() string {
	return "student"
}
