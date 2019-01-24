# Enforcer Ninja SDK (Go) [![GoDoc](https://godoc.org/github.com/hatchify/enforcer-ninja-go?status.svg)](https://godoc.org/github.com/hatchify/enforcer-ninja-go) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)
Enforcer Ninja SDK is a library for the go language to aid in interfacing with the Enforcer Ninja API

## Usage
Usage examples are available below. For more in-depth examples and documentation, be sure to visit our [documentation](https://godoc.org/github.com/hatchify/enforcer-ninja-go).

### Initializing SDK
```go
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
```

### Getting usage
```go
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
```

### Getting limit
```go
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
```

### Verifying e-mail address
```go
func ExampleSDK_VerifyEmail() {
	var err error
	if err = testSDK.VerifyEmail("open-source@hatchify.co"); err != nil {
		log.Fatalf("Error verifying email: %v", err)
	}
}
```

