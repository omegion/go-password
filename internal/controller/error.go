package controller

type CharsetEmptyError struct{}

func (e CharsetEmptyError) Error() string {
	return "Charset is empty."
}
