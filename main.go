package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var strDestination = flag.String("destination", "", "Webhook destination (https://example.com/webhooks)")
var strBody = flag.String("body", "", "Webhook payload")
var strMethod = flag.String("method", "", "Webhook method (defaults to 'POST')")
var strAuth = flag.String("basic-auth", "", "Optional basic authentication in a 'user:pass' format")

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    webhook-dispatcher --destination https://my.webhookrelay.com/v1/webhooks/.... --body hello\n")
		flag.PrintDefaults()
	}

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if *strDestination == "" && os.Getenv("DESTINATION") == "" {
		fmt.Println("both --destination flag and DESTINATION env variable cannot be empty, usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := http.DefaultClient

	var (
		user string
		pass string
	)
	if *strAuth != "" {
		parts := strings.SplitN(*strAuth, ":", 2)
		if len(parts) > 0 {
			user = parts[0]
			pass = parts[1]
		}
	}

	method := http.MethodPost
	switch *strMethod {
	case http.MethodGet, http.MethodHead, http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch:
		method = *strMethod
	default:
		// something invalid
	}

	var destination string
	if *strDestination != "" {
		destination = *strDestination
	}
	if os.Getenv("DESTINATION") != "" {
		destination = os.Getenv("DESTINATION")
	}

	req, err := http.NewRequest(method, destination, bytes.NewBufferString(*strBody))
	if err != nil {
		fmt.Printf("failed to create request: %s \n", err)
		os.Exit(1)
	}

	if user != "" && pass != "" {
		req.SetBasicAuth(user, pass)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("request failed: %s \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("response body read failed: %s \n", err)
		os.Exit(1)
	}

	var wr webhookResponse

	wr.Body = string(bodyBts)
	wr.StatusCode = resp.StatusCode

	encoded, err := json.Marshal(&wr)
	if err != nil {
		fmt.Printf("failed to encode response: %s \n", err)
		os.Exit(1)
	}

	fmt.Println(string(encoded))
	os.Exit(0)
}

type webhookResponse struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}
