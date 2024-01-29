package utils

import (
	"log"
)

var (
	LoadDotenvError    = "00"
	CreateRequestError = "10"
	SendRequestError   = "11"
	ReadResponseError  = "20"
	UnmarshalJSONError = "30"
	MarshalJSONError   = "31"
)

func ErrorManager(code string, err error) {
	if err == nil {
		return
	}
	switch code {
	case LoadDotenvError:
		log.Fatalf("Error loading .env file: %v", err)
	case CreateRequestError:
		log.Fatalf("Error creating request: %v", err)
	case SendRequestError:
		log.Fatalf("Error sending request: %v", err)
	case ReadResponseError:
		log.Fatalf("Error reading response: %v", err)
	case UnmarshalJSONError:
		log.Fatalf("Error unmarshaling JSON: %v", err)
	case MarshalJSONError:
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	log.Fatalf("Unknown Error: %v", err)
}
