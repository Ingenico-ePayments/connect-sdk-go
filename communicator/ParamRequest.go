package communicator

// ParamRequest represents a set of request parameters.
type ParamRequest interface {

	// ToRequestParameters converts this set of request parameters to a slice of RequestParam
	ToRequestParameters() RequestParams
}
