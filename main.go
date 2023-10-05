// Package main is a http client that makes eportal api requests.
package main

// import modules
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// global variables
var api_username string
var api_password string
var eportal_url string

// endpoint handlers
func listServers() {
	var body string = setupRequest(eportal_url + "/admin/api/servers")
	fmt.Print(body)
}

func listKeys() {
	var body string = setupRequest(eportal_url + "/admin/api/keys")
	fmt.Print(body)
}

func listFeeds() {
	var body string = setupRequest(eportal_url + "/admin/api/feeds")
	fmt.Print(body)
}

func listPatchsets() {
	var body string = setupRequest(eportal_url + "/admin/api/patchsets")
	fmt.Print(body)
}

func listUsers() {
	var body string = setupRequest(eportal_url + "/admin/api/users")
	fmt.Print(body)
}

// api request wrapper
func setupRequest(uri string) (body string) {
	// create http client
	client := http.Client{Timeout: 5 * time.Second}

	// construct http request
	req, err := http.NewRequest(http.MethodGet, uri, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	// add basic auth headers
	req.SetBasicAuth(api_username, api_password)

	// make request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read the http body
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// return the body of the request as a json string
	return string(resBody)
}

func main() {
	// parse cli arguments
	var query string
	flag.StringVar(&query, "query", "", "endpoint")
	flag.Parse()

	// check we have one argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: eportal-go --query=<endpoint>")
		os.Exit(1)
	}

	// find user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Cannot find $HOME")
		os.Exit(1)
	}

	// read credentials file
	content, err := os.ReadFile(home + "/.eportal.ini")
	if err != nil {
		fmt.Println("Please create an ~/.eportal.ini credentials file")
		os.Exit(1)
	}

	// split creds file on newline
	lines := strings.Split(string(content), "\n")
	for i := 0; i < 3; i++ {
		// remove whitespace around equals sign
		lines[i] = strings.ReplaceAll(lines[i], " ", "")

		// split into key,value pairs
		line := strings.Split(lines[i], "=")
		switch line[0] {
		case "username":
			api_username = line[1]
		case "password":
			api_password = line[1]
		case "url":
			eportal_url = line[1]
		default:
			fmt.Println("Unable to parse ~/.eportal.ini")
			os.Exit(1)
		}
	}

	// call handlers for valid endpoints
	switch query {
	case "servers":
		listServers()
	case "keys":
		listKeys()
	case "feeds":
		listFeeds()
	case "users":
		listUsers()
	case "patches", "patchsets":
		listPatchsets()
	default:
		fmt.Println("Please use a valid endpoint e.g. servers, patches, users, feeds or keys")
		os.Exit(1)
	}
}
