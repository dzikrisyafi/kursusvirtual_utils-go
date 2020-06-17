package controller_utils

import (
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

func GetIDInt64(paramID string, idName string) (int64, rest_errors.RestErr) {
	ParamID, idErr := strconv.ParseInt(paramID, 10, 64)
	if idErr != nil {
		return 0, rest_errors.NewBadRequestError(idName + " should be a number")
	}

	return ParamID, nil
}

func GetIDInt(paramID string, idName string) (int, rest_errors.RestErr) {
	ParamID, idErr := strconv.Atoi(paramID)
	if idErr != nil {
		return 0, rest_errors.NewBadRequestError(idName + " should be a number")
	}

	return ParamID, nil
}
