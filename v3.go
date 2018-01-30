package main

import (
	"fmt"
	"net/http"
)

func (opt *Command) V3(w http.ResponseWriter, r *http.Request) {
	links := map[string]interface{}{
		"self": map[string]string{
			"href": fmt.Sprintf("%s/v3", opt.serverURL),
		},
	}

	resources := []string{"tasks", "apps", "builds", "packages", "isolation_segments", "organizations", "spaces", "processes", "droplets", "service_instances"}

	for _, resource := range resources {
		links[resource] = map[string]interface{}{
			"href": fmt.Sprintf("%s/v3/%s", opt.serverURL, resource),
		}
	}

	response := map[string]interface{}{"links": links}

	fmt.Fprintf(w, ConvertToJSON(response))
}
