package todo

import "errors"

var ErrTaskAlreadyExist = errors.New("task already exist")
var ErrTaskNotFound = errors.New("task not found")
