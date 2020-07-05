package resource

import (
	"imgo/src/response"
	"net/http"
)

// CodeMessageMapping data
var CodeMessageMapping = response.CodeMessageMapping{
	"not_allow": response.CodeStatus{
		Code:   "000.000.000",
		Status: http.StatusBadRequest,
	},
	"invalid_full_name": response.CodeStatus{
		Code:   "001.001.001",
		Status: http.StatusBadRequest,
	},
	"invalid_username": response.CodeStatus{
		Code:   "001.001.002",
		Status: http.StatusBadRequest,
	},
	"invalid_email": response.CodeStatus{
		Code:   "001.001.003",
		Status: http.StatusBadRequest,
	},
	"invalid_password": response.CodeStatus{
		Code:   "001.001.004",
		Status: http.StatusBadRequest,
	},
}
