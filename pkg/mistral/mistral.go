package mistral

import (
	"github.com/gage-technologies/mistral-go"
)

func CreateMistralClient(apiKey string) *mistral.MistralClient {
	return mistral.NewMistralClientDefault(apiKey)
}
