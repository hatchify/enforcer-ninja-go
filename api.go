package sdk

import (
	"encoding/json"
	"io"

	"github.com/Hatch1fy/errors"
)

// processAPIResponse is a helper function to parse the API response body.
// Under the hood, the following occurs:
//	- The provided data value will be associated with the Data field within the apiResponse struct
//	- The response will be parsed as an apiResponse struct
//	- If the apiResponse struct contains errors, they will be returned as an errors.ErrorList
func processAPIResponse(r io.Reader, data interface{}) (err error) {
	var resp apiResponse
	resp.Data = data

	if err = json.NewDecoder(r).Decode(&resp); err != nil {
		// Error encountered while decoding as JSON
		return
	}

	if len(resp.Errors) == 0 {
		// No errors encountered, we can move on!
		return
	}

	var errs errors.ErrorList
	// Iterate through response error messages
	for _, msg := range resp.Errors {
		// Push error message to errors list
		errs.Push(errors.Error(msg))
	}

	// Return errors list as an error
	return errs.Err()
}

type apiResponse struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}
