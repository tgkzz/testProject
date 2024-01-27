package models

const (
	SuccessMsg = "success"
	ErrorMsg   = "error"

	SuccessGetOperation     = "successfully get person"
	SuccessCreatedOperation = "successfully created new person"
	SuccessDeleteOperation  = "successfully deleted person"
	SuccessPatchOperation   = "successfully update person"
)

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}
