package main

import (
	"fmt"
	"net/http"
)

func (opt *Command) Login(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"app": map[string]string{
			"version": opt.UAAVersion,
		},
		"commit_id":      "888888888",
		"entityID":       "localhost",
		"idpDefinitions": map[string]string{},
		"links": map[string]string{
			"login":    opt.serverURL,
			"passwd":   fmt.Sprintf("%s/forgot-password", opt.serverURL),
			"register": fmt.Sprintf("%s/sign-up", opt.serverURL),
			"uaa":      opt.serverURL,
		},
		"prompts": map[string]interface{}{
			"password": []string{
				"password",
				"Password",
			},
			"username": []string{
				"text",
				"Email",
			},
		},
		"timestamp": "2018-01-01T00:00:00-0000",
		"zone_name": "uaa",
	}

	fmt.Fprintf(w, ConvertToJSON(response))
}
