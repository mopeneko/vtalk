package interfaces

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	Param(string) string
}
