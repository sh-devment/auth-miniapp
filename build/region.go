package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

// regionDef is everything that differs between the ru and com deployments
// apart from env config: the language of the page, the bot's greeting and the
// login errors a user can be shown. The values live in region-ru.go and
// region-com.go. Technical text (logs, the reasons passed through from
// auth-center or the network) is English in both and stays out.
//
// Today only com is deployed (the matrix in .github/workflows/deploy.yml). The
// ru profile is kept filled and checked at every start all the same, so that
// turning ru on is a matter of the deploy config, not of code.
type regionDef struct {
	Lang string // <html lang>

	// What a /start sends back. Plain text, no parse_mode — see webhook.go.
	Greeting string

	Text pageText
}

// pageText holds the strings of web/index.html. Fields, not a map: a typo in
// the template fails at render instead of silently printing nothing.
type pageText struct {
	Title    string // <title>
	SignedAs string // caption above the name
	Checking string // the frame until the answer arrives

	// Errors handleAuth writes for a broken login link. They reach the page as
	// the reason under "Error", so they are the user's language, not English.
	NoSession   string
	UnknownLink string

	JS jsText // strings set by script.js
}

// jsText is rendered into the page as JSON (window.TEXT) and read by script.js.
type jsText struct {
	Next       string `json:"next"` // the button onward
	Back       string `json:"back"` // the button that retries
	Failed     string `json:"failed"`
	NoAnswer   string `json:"noAnswer"` // %d = seconds
	NoSDK      string `json:"noSDK"`
	NoInitData string `json:"noInitData"`
}

var regions = map[string]*regionDef{
	"ru":  &regionRU,
	"com": &regionCOM,
}

// Active region, set once by initRegion before anything serves.
var (
	regionName string
	region     *regionDef
)

// initRegion picks the region by REGION (ru | com, set by the deploy matrix).
// Empty means local dev and falls back to ru; anything else is a config error
// and stops the process rather than serving the wrong zone — the health check
// then fails and deploy.sh rolls back. A string forgotten in either region is
// fatal too, whichever region this instance serves.
func initRegion() {
	for name, r := range regions {
		if err := checkFilled(reflect.ValueOf(*r), "region"); err != nil {
			log.Fatalf("region %s: %v", name, err)
		}
	}

	regionName = os.Getenv("REGION")
	if regionName == "" {
		regionName = "ru"
	}
	r, ok := regions[regionName]
	if !ok {
		log.Fatalf("unknown REGION %q (want ru or com)", regionName)
	}
	region = r
	log.Printf("region: %s (lang=%s)", regionName, region.Lang)
}

// checkFilled walks structs and fails on the first empty string.
func checkFilled(v reflect.Value, path string) error {
	switch v.Kind() {
	case reflect.String:
		if v.Len() == 0 {
			return fmt.Errorf("%s is empty", path)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if err := checkFilled(v.Field(i), path+"."+v.Type().Field(i).Name); err != nil {
				return err
			}
		}
	}
	return nil
}
