package main

import (
	"fmt"
	"net/http"
)

func (opt Command) Root(w http.ResponseWriter, r *http.Request) {
	links := map[string]interface{}{
		"self": map[string]string{
			"href": opt.serverURL,
		},
		"cloud_controller_v2": map[string]interface{}{
			"href": fmt.Sprintf("%s/v2", opt.serverURL),
			"meta": map[string]string{
				"version": opt.CCV2Version,
			},
		},
		"cloud_controller_v3": map[string]interface{}{
			"href": fmt.Sprintf("%s/v3", opt.serverURL),
			"meta": map[string]string{
				"version": opt.CCV3Version,
			},
		},
		"logging": map[string]interface{}{
			"href": "wss://localhost:433",
		},
		"uaa": map[string]interface{}{
			"href": opt.serverURL,
		},
	}

	if !opt.DisableNetworkPolicyV0 {
		links["network_policy_v0"] = map[string]interface{}{
			"href": fmt.Sprintf("%s/networking/v0/external", opt.serverURL),
		}
	}

	if !opt.DisableNetworkPolicyV1 {
		links["network_policy_v1"] = map[string]interface{}{
			"href": fmt.Sprintf("%s/networking/v1/external", opt.serverURL),
		}
	}

	if !opt.DisableRouting {
		links["routing"] = map[string]interface{}{
			"href": fmt.Sprintf("%s/routing", opt.serverURL),
		}
	}

	response := map[string]interface{}{"links": links}

	fmt.Fprintf(w, ConvertToJSON(response))
}
