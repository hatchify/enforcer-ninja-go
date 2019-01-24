package sdk

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

var (
	testSDK    *SDK
	testAPIKey string
)

func TestMain(m *testing.M) {
	var (
		cfg Config
		err error
	)

	if _, err = toml.DecodeFile("./test.toml", &cfg); err != nil {
		log.Fatalf("Error loading test configuration: %v", err)
	}

	if testSDK, err = New(cfg.APIKey); err != nil {
		log.Fatalf("Error initializing SDK: %v", err)
	}

	testAPIKey = cfg.APIKey

	var statusCode int
	statusCode = m.Run()
	os.Exit(statusCode)
}

func TestGetUsage(t *testing.T) {
	var (
		usage []UsageEntry
		err   error
	)

	if usage, err = testSDK.GetUsage(ViewWeek); err != nil {
		t.Fatalf("error getting usage: %v", err)
	}

	if len(usage) != 7 {
		t.Fatalf("invalid number of days!")
	}
}

func TestGetLimit(t *testing.T) {
	var (
		limit uint64
		err   error
	)

	if limit, err = testSDK.GetLimit(); err != nil {
		t.Fatalf("error getting limit: %v", err)
	}

	if limit == 0 {
		t.Fatalf("invalid limit, expected a non-zero value")
	}
}

func TestVerifyEmail(t *testing.T) {
	var err error
	if err = testSDK.VerifyEmail("open-source@hatchify.co"); err != nil {
		t.Fatalf("error verifying email: %v", err)
	}

	if err = testSDK.VerifyEmail("foobar.bademail@hatchify.co"); err != ErrEmailDoesNotExist {
		t.Fatalf("unexpected error encountered, expected \"%s\" and received \"%s\"", ErrEmailDoesNotExist, err)
	}
}

func ExampleNew() {
	var (
		sdk *SDK
		err error
	)

	if sdk, err = New(testAPIKey); err != nil {
		log.Fatal(err)
	}

	// You can now begin making Enforcer Ninja requests!
	fmt.Println(sdk)
	return
}

func ExampleSDK_GetUsage() {
	var (
		usage []UsageEntry
		err   error
	)

	if usage, err = testSDK.GetUsage(ViewWeek); err != nil {
		log.Fatalf("Error getting usage: %v", err)
	}

	for _, entry := range usage {
		fmt.Printf("Entry: %+v\n", entry)
	}
}

func ExampleSDK_GetLimit() {
	var (
		limit uint64
		err   error
	)

	if limit, err = testSDK.GetLimit(); err != nil {
		log.Fatalf("Error getting limit: %v", err)
	}

	fmt.Printf("Limit: %d", limit)
}
