# Topicus KeyHub SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/topicuskeyhub/sdk-go/)](https://pkg.go.dev/github.com/topicuskeyhub/sdk-go/)

> **Note:** The version number of the SDK must match the version of your Topicus KeyHub installation. An older version of the SDK might work on a newer version of Topicus KeyHub via the provided backwards compatibility layer, but a newer version of the SDK will not work on an older version of Topicus KeyHub.
>
> **Note:** The GO SDK is current considered in alpha status. Its API might change in future versions, it is largely untested and at the moment only supports the client credentials grant.

## 1. Installation

```Shell
go get github.com/topicuskeyhub/sdk-go
```

## 2. Getting started

### 2.1 Create a client application in Topicus KeyHub

Register a OAuth 2.0 client application in Topicus KeyHub as described in [section 16.1 of the manual](https://files.topicus-keyhub.com/manual/manual-en-GB.html#sec-oauth2-client).
[Section 16.5](https://files.topicus-keyhub.com/manual/manual-en-GB.html#sec-oauth2-permissions) describes the permissions you can assign to this application.

### 2.2 Create a KeyHubClient object and make an API call

```Golang
package main

import (
	"context"
	"log"
	"net/http"

	keyhub "github.com/topicuskeyhub/sdk-go"
)

func main() {
	issuer := "<the issuer of your Topicus KeyHub installation>"
	clientid := "<the client id of your application registration>"
	clientsecret := "<the client sercret of your application registration>"

	adapter, err := keyhub.NewKeyHubRequestAdapter(http.DefaultClient, issuer, clientid, clientsecret)
	if err != nil {
		log.Fatalf("Cannot create request adapter: %v\n", err)
	}

	client := keyhub.NewKeyHubClient(adapter)
	me, err := client.Client().Me().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting /client/me: %v\n", err)
	}
	log.Printf("me: %s", *me.GetName())
}
```

