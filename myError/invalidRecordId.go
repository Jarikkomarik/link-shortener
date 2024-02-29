package myError

type InvalidRecordId struct {
}

func (invalidRecordId InvalidRecordId) Error() string { //implementing error
	return "Invalid url id."
}
