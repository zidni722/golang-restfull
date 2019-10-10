package utils

type UnAuthenticatedError struct {
	Message string
}

func (e *UnAuthenticatedError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "UnAuthenticated"
}
