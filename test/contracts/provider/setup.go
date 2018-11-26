package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type providerStateInfo struct {
	// Consumer name
	Consumer string `json:"consumer"`
	// State
	State string `json:"state"`
	// States
	States []string `json:"states"`
}

type ProviderInitialState struct {
	//Tokens loginusers.Tokens
}

// Setup starts a setup service for a provider - should be replaced by a provider setup endpoint
func Setup(setupHost string, setupPort int, authURL string, userName string, userPassword string) *ProviderInitialState {
	log.SetOutput(os.Stdout)

	/*
		loginUsersConfig := config.DefaultConfig()
		loginUsersConfig.Auth.ServerAddress = authURL
		// Log user in to get tokens
		userTokens, err := loginusers.OAuth2(userName, userPassword, loginUsersConfig)
		if err != nil {
			log.Fatalf("Unable to login user: %s", err)
			return nil
		}
	*/
	go setupEndpoint(setupHost, setupPort)

	return &ProviderInitialState{
		//Tokens: *userTokens,
	}
}

func setupEndpoint(setupHost string, setupPort int) {
	http.HandleFunc("/pact/setup", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf(">>> ERROR: Unable to read request body.\n %q", err)
			return
		}

		var providerState providerStateInfo
		json.Unmarshal(body, &providerState)

		switch providerState.State {
		case "Build tool detector service is up and running.":
			log.Printf(">>>> %s\n", providerState.State)
		default:
			errorMessage(w, fmt.Sprintf("State '%s' not impemented.", providerState.State))
			return
		}
		fmt.Fprintf(w, "Provider states has ben set up.\n")
	})

	var setupURL = fmt.Sprintf("%s:%d", setupHost, setupPort)
	log.Printf(">>> Starting ProviderSetup and listening at %s\n", setupURL)
	log.Fatal(http.ListenAndServe(setupURL, nil))
}

func errorMessage(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(500)
	fmt.Fprintf(w, `{"error": "%s"}`, errorMessage)
}
