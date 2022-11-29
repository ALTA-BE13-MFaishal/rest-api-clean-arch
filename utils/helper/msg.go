package helper

import (
	"errors"
	"strings"
)

func ServiceErrorMsg(errData error) error {
	if strings.Contains(errData.Error(), "table") {
		return errors.New("Failed. Error on process. Please contact your administrator.")
	} else if strings.Contains(errData.Error(), "found") {
		return errors.New("Failed. Data not found. Please check input again.")
	} else {
		return errors.New("Failed. Other Error. Please contact your administrator.")
	}
}
