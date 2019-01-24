package sdk

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Hatch1fy/errors"
)

const (
	// Host of Enforcer Ninja service
	Host = "https://enforcer.ninja/"
	// APIPath is the API url route path
	APIPath = "api"

	// GetUserIDEndpoint is the endpoint for getting a user's current user ID
	GetUserIDEndpoint = "users/id"
	// GetUsageEndpoint is the endpoint for getting a user's current usage
	// Note: This endpoint accepts a "view" query parameter with the following values:
	//	- day
	//	- week
	//	- month
	GetUsageEndpoint = "usage"
	// GetLimitEndpoint is the endpoint for getting user limits
	// Note: This endpoint expects a user id parameter to replace the string tag
	GetLimitEndpoint = "usage/limit/%s"
	// VerifyEmailEndpoint is the endpoint for verifying an email
	// Note: This endpoint expects an email parameter to replace the string tag
	VerifyEmailEndpoint = "verify/email/%s"
)

const (
	// ErrInvalidView is returned when an unexpected view is provided
	// Note: Please see the Views constant block within view.go for the available options
	ErrInvalidView = errors.Error("invalid view provided")
	// ErrEmailDoesNotExist is returned when an email address being verified does not exist
	ErrEmailDoesNotExist = errors.Error("email address does not exist")
)

// New will return a new instance of SDK
func New(apiKey string) (sp *SDK, err error) {
	var s SDK
	// Set api key
	s.apiKey = apiKey
	// Set user ID
	if s.userID, err = s.getUserID(); err != nil {
		return
	}

	// Reference pointer to created SDK
	sp = &s
	return
}

// SDK manages the Enforcer Ninja SDK
type SDK struct {
	hc http.Client

	apiKey string
	userID string
}

func (s *SDK) request(endpoint string, query url.Values, response interface{}) (err error) {
	var u *url.URL
	// Create a new URL from the provided endpoint
	if u, err = getURL(endpoint, query); err != nil {
		return
	}

	var req *http.Request
	// Create a new HTTP request for the created URL
	if req, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}

	// Set API key for user authentication
	req.Header.Set("X-Api-Key", s.apiKey)

	var resp *http.Response
	// Make request
	if resp, err = s.hc.Do(req); err != nil {
		return
	}
	// Defer the closing of our response body
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		// Succesful no content response, return early
		return
	}

	// Process api response passing response body and provided response value
	return processAPIResponse(resp.Body, response)
}

func (s *SDK) getUserID() (userID string, err error) {
	err = s.request(GetUserIDEndpoint, nil, &userID)
	return
}

// UserID will get the current authenticated user's ID
// Note: This method does not make an HTTP request
func (s *SDK) UserID() (userID string) {
	return s.userID
}

// GetUsage will get the current usage of the SDK's user (api key)
func (s *SDK) GetUsage(v View) (usage []UsageEntry, err error) {
	if err = v.Validate(); err != nil {
		// View is invalid, return
		return
	}

	err = s.request(GetUsageEndpoint, getViewQuery(v), &usage)
	return
}

// GetLimit will get the current limit of the SDK's user (api key)
func (s *SDK) GetLimit() (limit uint64, err error) {
	endpoint := fmt.Sprintf(GetLimitEndpoint, s.userID)
	err = s.request(endpoint, nil, &limit)
	return
}

// VerifyEmail will verify an email
func (s *SDK) VerifyEmail(email string) (err error) {
	endpoint := fmt.Sprintf(VerifyEmailEndpoint, email)
	return s.request(endpoint, nil, nil)
}
