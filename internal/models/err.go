package models

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrAtoi                = constError("given parameter is not the number")
	ErrInvalidFilter       = constError("invalid filter data")
	ErrEmptyNameOrSurname  = constError("name and surname must be not empty")
	ErrInvalidUpdateParams = constError("json may be empty or filled in incorrectly")
	ErrBadStatusCode       = constError("bad statusCode")
	ErrNoRowsAffected      = constError("no rows affected")
	ErrServiceUnavailable  = constError("unable to connect to api")
	ErrSqlNoRows           = constError("The result query gave no object to show")
	ErrEmptyResult         = constError("empty result")
	ErrInvalidPagination   = constError("invalid pagination data")
)
