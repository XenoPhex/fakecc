package main

import (
	"fmt"
	"net/http"
)

func (opt *Command) V2Info(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"name":                         "",
		"build":                        "",
		"support":                      opt.serverURL,
		"version":                      0,
		"description":                  "This is a fake cloud foundry server",
		"authorization_endpoint":       opt.serverURL,
		"token_endpoint":               opt.serverURL,
		"min_cli_version":              "6.23.0",
		"min_recommended_cli_version":  "latest",
		"api_version":                  opt.CCV2Version,
		"app_ssh_endpoint":             "localhost:2222",
		"app_ssh_host_key_fingerprint": "some-finger-print",
		"app_ssh_oauth_client":         "ssh-proxy",
		"doppler_logging_endpoint":     "wss://localhost:443",
	}

	if !opt.DisableRouting {
		response["routing_endpoint"] = fmt.Sprintf("%s/routing", opt.serverURL)
	}

	fmt.Fprintf(w, ConvertToJSON(response))
}
