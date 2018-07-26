package pactconsumer

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/nurali-techie/play-go-pact/hello-world-example/consumer"
	"github.com/pact-foundation/pact-go/dsl"
)

type request = dsl.Request

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var pact dsl.Pact
var term = dsl.Term
var name string

func TestMain(m *testing.M) {
	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	code := m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	pact.WritePact()
	pact.Teardown()

	os.Exit(code)
}

func setup() {
	pact = createPact()
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 "ping",
		Provider:                 "pong",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		LogLevel:                 "DEBUG",
		DisableToolValidityCheck: true,
	}
}

func TestPactConsumer_Hello(t *testing.T) {
	name := "John"

	testHello := func() error {
		host := fmt.Sprintf("http://localhost:%d", pact.Server.Port)
		msg, err := consumer.HelloCaller(host, name)
		if err != nil {
			return err
		}
		if !strings.Contains(msg, "Hello") {
			return errors.New("Invalid response")
		}
		return nil
	}

	pact.
		AddInteraction().
		WithRequest(request{
			Method: "GET",
			Path:   term("/", "/"),
			Query: dsl.MapMatcher{
				"name": term(name, "[a-zA-Z]+"),
			},
		}).
		WillRespondWith(dsl.Response{Status: 200, Body: "Hello John!"})

	err := pact.Verify(testHello)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}
