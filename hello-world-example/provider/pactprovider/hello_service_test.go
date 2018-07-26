package pactprovider

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/nurali-techie/play-go-pact/hello-world-example/provider"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

func TestPactProvider_Hello(t *testing.T) {
	go startInstrumentedProvider(port)

	pact := createPact()

	// Verify the Provider with local Pact Files
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/ping-pong.json", pactDir))},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func createPact() dsl.Pact {
	// Create Pact connecting to local Daemon
	return dsl.Pact{
		Consumer: "ping",
		Provider: "pong",
		LogDir:   logDir,
		PactDir:  pactDir,
	}
}

func startInstrumentedProvider(port int) {
	http.HandleFunc("/", provider.HelloHandler)
	host := fmt.Sprintf(":%d", port)
	http.ListenAndServe(host, nil)
}
