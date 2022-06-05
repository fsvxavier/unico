package models

import "errors"

var (
	//ErrInternalServerError ...
	ErrInternalServerError = errors.New("Internal Server Error")

	//ErrNotFound ...
	ErrNotFound = errors.New("Your requested Item is not found")

	//ErrNoRows ...
	ErrNoRows = errors.New("sql: no rows in result set")

	//ErrConflict ...
	ErrConflict = errors.New("Your Item already exist")

	//ErrBadParamInput ...
	ErrBadParamInput = errors.New("Given Param is not valid")

	//ErrorSizeAccessKey ...
	ErrorSizeAccessKey = errors.New("Error when returning the Access Key invalid size")

	//ErrorAccessKeyDifferent ...
	ErrorAccessKeyDifferent = errors.New("Access key does not match invoice data")
)
