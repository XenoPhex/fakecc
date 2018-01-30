package main

import (
	"fmt"
	"net/http"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Command struct {
	Port        int    `short:"p" long:"port" description:"server port" default:"8080"`
	CCV2Version string `long:"ccv2" description:"Cloud Controller V2 version number" default:"2.101.0"`
	CCV3Version string `long:"ccv3" description:"Cloud Controller V3 version number" default:"3.36.0"`
	UAAVersion  string `long:"uaa" description:"UAA version number" default:"4.8.3"`

	DisableRouting         bool `long:"disable-routing" description:"Hide 'routing' from the root response"`
	DisableV3              bool `long:"disable-ccv3" description:"return a 404 from the root response"`
	DisableNetworkPolicyV0 bool `long:"disable-network-policy-v0" description:"Hide 'network_policy_v0' from the root response"`
	DisableNetworkPolicyV1 bool `long:"disable-network-policy-v1" description:"Hide 'network_policy_v1' from the root response"`

	serverURL string
}

func (opt *Command) Execute(args []string) error {
	opt.serverURL = fmt.Sprintf("http://localhost:%d", opt.Port)

	fmt.Println("API URL:", opt.serverURL)

	http.HandleFunc("/login", opt.Login)
	http.HandleFunc("/v2/info", opt.V2Info)
	if opt.DisableV3 {
		http.HandleFunc("/", opt.NotFound)
	} else {
		http.HandleFunc("/", opt.Root)
		http.HandleFunc("/v3", opt.V3)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", opt.Port), nil)
}

func main() {
	cmd := new(Command)
	parser := flags.NewParser(cmd, flags.Default)

	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Execute(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
