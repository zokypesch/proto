package service

import "log"

// PaketData struct of service
type PaketData struct{}

// NewPublishDataService for new service of publish Data
func NewPublishDataService() Action {
	return &PaketData{}
}

// Do action pulsa service
func (paket *PaketData) Do(req *Request) error {
	log.Printf("Do action isi paket data %s", req.Attribute.Code)

	payload := req.Payload.(map[string]interface{})

	v, ok := payload["phone"]

	if !ok {
		log.Println("Phone not found")
	}

	log.Println("check", v.(string), req.Attribute.Code)

	return nil
}
