package service

//Request request from pubsub
type Request struct {
	Payload   interface{}
	Attribute Attribute
}

// Attribute attribute hybrid
type Attribute struct {
	Type        string
	Code        string
	TransSerial string
}

//Action interfacing all data
type Action interface {
	Do(req *Request) error
}
