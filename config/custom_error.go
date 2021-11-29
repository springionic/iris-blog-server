// Package error
// created by lilei at 2021/11/23
package config

type CustomError interface {
	error
	Code() uint
	Message() string
	HTTPCode() int
}

type ServerError struct {
	code     uint
	msg      string
	httpCode int
}

type ClientError struct {
	code     uint
	msg      string
	httpCode int
}

type BusinessError struct {
	code     uint
	msg      string
	httpCode int
}

func (s ServerError) Error() string {
	return s.msg
}

func (s ServerError) Code() uint {
	return s.code
}

func (s ServerError) Message() string {
	return s.msg
}

func (s ServerError) HTTPCode() int {
	return s.httpCode
}

func (c ClientError) Error() string {
	return c.msg
}

func (c ClientError) Code() uint {
	return c.code
}

func (c ClientError) Message() string {
	return c.msg
}

func (c ClientError) HTTPCode() int {
	return c.httpCode
}

func (b BusinessError) Error() string {
	return b.msg
}

func (b BusinessError) Code() uint {
	return b.code
}

func (b BusinessError) Message() string {
	return b.msg
}

func (b BusinessError) HTTPCode() int {
	return b.httpCode
}
