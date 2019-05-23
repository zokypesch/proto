package service

import "log"

// Pulsa struct of service
type Pulsa struct{}

// NewPublishPulsaService for new service of publish pulsa
func NewPublishPulsaService() Action {
	return &Pulsa{}
}

// Do action pulsa service
func (pulsa *Pulsa) Do(req *Request) error {
	log.Printf("Do action isi pulsa %s", req.Attribute.Code)

	payload := req.Payload.(map[string]interface{})

	v, ok := payload["phone"]

	if !ok {
		log.Println("Phone not found")
	}

	log.Println("check", v.(string), req.Attribute.Code)

	return nil
}
