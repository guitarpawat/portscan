package model

// Json interface is the object that can be marshal to JSON format
// and send it between website and server by using RESTful API
type Json interface {
	Marshal() ([]byte, error)
}
