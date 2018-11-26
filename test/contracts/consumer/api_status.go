package consumer

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/fabric8-services/build-tool-detector/test/contracts/model"
	"github.com/pact-foundation/pact-go/dsl"
)

// BuildToolDetectorAPIStatus defines contract of /api/status(/) endpoint
func BuildToolDetectorAPIStatus(t *testing.T, pact *dsl.Pact) {

	log.Printf("\n\nInvoking BuildToolDetectorAPIStatus now\n")

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Build tool detector service is up and running.").
		UponReceiving("A request to get status").
		WithRequest(dsl.Request{
			Method: "GET",
			Path: dsl.Term(
				"/api/status/",
				"/api/status[/]?",
			),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/vnd.status+json")},
			Body:    dsl.Match(model.APIStatusMessage{}),
		})

	// Pass in test case
	var test = func() error {
		u := fmt.Sprintf("http://localhost:%d/api/status", pact.Server.Port)
		req, err := http.NewRequest("GET", u, nil)
		client := http.DefaultClient

		if err != nil {
			return err
		}
		_, err = client.Do(req)
		if err != nil {
			return err
		}

		u = fmt.Sprintf("http://localhost:%d/api/status/", pact.Server.Port)
		req, err = http.NewRequest("GET", u, nil)

		if err != nil {
			return err
		}
		_, err = client.Do(req)
		if err != nil {
			return err
		}
		return err
	}
	log.Printf("Pact interactions after BuildToolDetectorAPIStatus: %d", len(pact.Interactions))
	// Verify
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
