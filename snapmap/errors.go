package snapmap

import "errors"

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError interface{}) error {
	if httpError != nil {
		return httpError
	}
	if apiError == "" || apiError == nil {
		return nil
	}
	return errors.New(apiError.(string))
}
