package error

import (
	"fmt"
)

var ErrorMessages map[string]string

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorReponse struct {
	Error *Error `json:"Error,omitempty"`
}

func (r *Error) Error() string {
	return fmt.Sprintf("code:%s;message:%s", r.Code, r.Message)
}

var INTERNAL_ERROR = &Error{Code: "INTERNAL_ERROR", Message: "Internal Error"}
var ErrInvalidCredential = &Error{Code: "BAD_CREDENTIAL", Message: "Invalid Credential"}
var ErrInvalidToken = &Error{Code: "BAD_CREDENTIAL", Message: "Invalid Credential"}
var ErrInvalidRequest = &Error{Code: "BAD_INPUT", Message: "Invalid Request"}
var NOT_FOUND_USER = &Error{Code: "NOT_FOUND_USER", Message: "User not found"}
var NOT_FOUND_CATEGORY = &Error{Code: "NOT_FOUND_CATEGORY", Message: "Category not found"}
var NOT_ENOUGH_STOCK = &Error{Code: "NOT_ENOUGH_STOCK", Message: "Not enough stock for product"}
var VARIANT_NOT_FOUND = &Error{Code: "VARIANT_NOT_FOUND", Message: "Variant not found"}
var PRODUCT_NOT_FOUND = &Error{Code: "PRODUCT_NOT_FOUND", Message: "Product not found"}
