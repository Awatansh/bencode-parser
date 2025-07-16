package utils

import "fmt"


func BencodeError struct{
	message string
}

func (e *BencodeError) Error() string{
	return fmt.Sprintf("BencodeError: %s",e.message)
}

func NewBencodeError(format string,args ...interface{}) *BencodeError{
	return &BencodeError{message: fmt.Sprintf(format,args...)}
}
