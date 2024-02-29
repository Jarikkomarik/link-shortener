package myError

type InvalidURLError struct {
}

func (invalidURLError InvalidURLError) Error() string { //implementing error
	return "Invalid url passed."
}
