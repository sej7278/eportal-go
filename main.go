// Package main is a http client that makes eportal api requests.
package main

// import modules
import (
	"encoding/json"
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
var jsonarg bool

// endpoint handlers
func listServers() {
	var body string = setupRequest(eportal_url + "/admin/api/servers")

	// output raw json
	if jsonarg {
		fmt.Println(body)
		return
	}

	type Servers struct {
		Count  int `json:"count"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Result []struct {
			Checkin       any    `json:"checkin"`
			Distro        any    `json:"distro"`
			DistroVersion any    `json:"distro_version"`
			Euname        any    `json:"euname"`
			Feed          string `json:"feed"`
			Hostname      string `json:"hostname"`
			ID            string `json:"id"`
			IP            string `json:"ip"`
			KcareVersion  any    `json:"kcare_version"`
			KernelID      any    `json:"kernel_id"`
			Key           string `json:"key"`
			Machine       any    `json:"machine"`
			PatchLevel    any    `json:"patch_level"`
			PatchType     any    `json:"patch_type"`
			Patchset      any    `json:"patchset"`
			Processor     any    `json:"processor"`
			Registered    string `json:"registered"`
			Release       any    `json:"release"`
			Tags          any    `json:"tags"`
			Updated       any    `json:"updated"`
			Uptime        any    `json:"uptime"`
			Version       any    `json:"version"`
			Virt          any    `json:"virt"`
		} `json:"result"`
	}
}

func listKeys() {
	// make api request
	var body string = setupRequest(eportal_url + "/admin/api/keys")

	// output raw json
	if jsonarg {
		fmt.Println(body)
		return
	}

	// define struct
	type Keys struct {
		Result []struct {
			Feed        string   `json:"feed"`
			Key         string   `json:"key"`
			Description string   `json:"note"`
			Products    []string `json:"products"`
			Limit       int      `json:"server_limit"`
		} `json:"result"`
	}

	// decode the json to a struct
	var keys Keys
	json.Unmarshal([]byte(body), &keys)

	// loop through the result array
	for _, result := range keys.Result {
		fmt.Println("Feed:", result.Feed)
		fmt.Println("Key:", result.Key)
		fmt.Println("Description:", result.Description)
		fmt.Println("Limit:", result.Limit)
		fmt.Println("Products:")

		// loop through the products array
		for _, product := range result.Products {
			fmt.Println("  *", product)
		}

		// newline before next result
		fmt.Println()
	}
}

func listFeeds() {
	var body string = setupRequest(eportal_url + "/admin/api/feeds")

	// output raw json
	if jsonarg {
		fmt.Println(body)
		return
	}

	type Feeds struct {
		Result []struct {
			Auto        bool   `json:"auto"`
			Channel     string `json:"channel"`
			DeployAfter int    `json:"deploy_after"`
			Name        string `json:"name"`
		} `json:"result"`
	}
}

func listPatchsets() {
	var body string = setupRequest(eportal_url + "/admin/api/patchsets")

	// output raw json
	if jsonarg {
		fmt.Println(body)
		return
	}

	type Patchsets struct {
		Result []struct {
			Patchset string `json:"patchset"`
			Status   string `json:"status"`
		} `json:"result"`
	}
}

func listUsers() {
	var body string = setupRequest(eportal_url + "/admin/api/users")

	// output raw json
	if jsonarg {
		fmt.Println(body)
		return
	}

	type Users struct {
		Result []struct {
			Description any    `json:"description"`
			ID          int    `json:"id"`
			Readonly    bool   `json:"readonly"`
			Username    string `json:"username"`
		} `json:"result"`
	}
}

func loadCreds() {
	// find user home directory
	home, _ := os.UserHomeDir()

	// read credentials file
	content, err := os.ReadFile(home + "/.eportal.ini")
	if err != nil {
		log.Fatal(err)
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
			log.Fatal("Unable to parse ~/.eportal.ini")
		}
	}
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
	var serversarg, keysarg, feedsarg, usersarg, patchsetsarg bool
	flag.BoolVar(&serversarg, "servers", false, "--servers")
	flag.BoolVar(&keysarg, "keys", false, "--keys")
	flag.BoolVar(&feedsarg, "feeds", false, "--feeds")
	flag.BoolVar(&usersarg, "users", false, "--users")
	flag.BoolVar(&patchsetsarg, "patchsets", false, "--patchsets")
	flag.BoolVar(&jsonarg, "json", false, "--json")
	flag.Parse()

	// check we have at least 1 argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: eportal-go --<servers|keys|feeds|users|patchsets> [--json]")
		os.Exit(1)
	}

	// load credentials from file
	loadCreds()

	// call handlers for valid endpoints
	if serversarg {
		listServers()
	}

	if keysarg {
		listKeys()
	}

	if feedsarg {
		listFeeds()
	}

	if usersarg {
		listUsers()
	}

	if patchsetsarg {
		listPatchsets()
	}
}
