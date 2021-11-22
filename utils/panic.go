package utils

// PanicError panics error if it exists
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
