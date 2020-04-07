package tools

import "strings"

// RepositoryResult struct to set common result of repo
type RepositoryResult struct {
	Data  interface{}
	Error error
}

// DeliveryResult struct to set common result of delivery
type DeliveryResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  bool        `json:"status"`
	Error   interface{} `json:"error,omitempty"`
}

// ErrorDeliveryValidatorResult to set error delivery validator result
func ErrorDeliveryValidatorResult(code int, err interface{}) DeliveryResult {
	var (
		errorResult interface{}
		mapError    = make(map[string]string, 0)
	)

	// initiate variable
	errors := err.(string)
	errList := strings.Split(errors, "\n")

	// loop each error
	for _, e := range errList {
		errSplit := strings.Split(e, "|")

		// append error
		mapError[errSplit[0]] = errSplit[1]
	}

	// set error
	errorResult = mapError

	return DeliveryResult{
		Code:   code,
		Error:  errorResult,
		Status: false,
	}
}

// ErrorDeliveryResult to set success delivery result
func ErrorDeliveryResult(code int, error interface{}) DeliveryResult {
	return DeliveryResult{
		Code:   code,
		Error:  error,
		Status: false,
	}
}

// SuccessDeliveryResult to set success delivery result
func SuccessDeliveryResult(code int, data interface{}, message string) DeliveryResult {
	return DeliveryResult{
		Code:    code,
		Data:    data,
		Message: message,
		Status:  true,
	}
}
