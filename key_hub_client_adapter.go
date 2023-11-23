// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package sdkgo

// not-generated
//go:generate bash generate.sh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	nethttp "net/http"
	"net/url"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type KeyHubAccessTokenProvider struct {
	tokenSource  *oauth2.TokenSource
	vaultSession string
}

func NewKeyHubAccessTokenProvider(tokenSource *oauth2.TokenSource) *KeyHubAccessTokenProvider {
	result := KeyHubAccessTokenProvider{
		tokenSource: tokenSource,
	}
	return &result
}

func (p *KeyHubAccessTokenProvider) GetAuthorizationToken(ctx context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	token, err := (*p.tokenSource).Token()
	if err != nil {
		return "", err
	}
	if p.vaultSession == "" {
		vault := token.Extra("vaultSession")
		if vault != nil {
			p.vaultSession = fmt.Sprintf("%v", vault)
		}
	}
	return token.AccessToken, nil
}

func (p *KeyHubAccessTokenProvider) GetAllowedHostsValidator() *auth.AllowedHostsValidator {
	validator := auth.NewAllowedHostsValidator(nil)
	return &validator
}

func (p *KeyHubAccessTokenProvider) Intercept(pipeline http.Pipeline, middlewareIndex int, req *nethttp.Request) (*nethttp.Response, error) {
	if p.vaultSession != "" {
		req.Header.Set("topicus-Vault-session", fmt.Sprintf("%v", p.vaultSession))
	}
	return pipeline.Next(req, middlewareIndex)
}

func NewKeyHubRequestAdapter(client *nethttp.Client, issuer string, clientID string, clientSecret string) (*http.NetHttpRequestAdapter, error) {
	if client == nil {
		client = nethttp.DefaultClient
	}
	if client.Timeout == 0 {
		client.Timeout = time.Second * 20
	}
	ctx := oidc.ClientContext(context.Background(), client)
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	oauth2Conf := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{oidc.ScopeOpenID},
		TokenURL:     provider.Endpoint().TokenURL + "?authVault=access",
	}
	tokenSource := oauth2Conf.TokenSource(ctx)
	accessTokenProvider := NewKeyHubAccessTokenProvider(&tokenSource)
	middlewares := []http.Middleware{
		accessTokenProvider,
		http.NewParametersNameDecodingHandler(),
		http.NewUserAgentHandler(),
	}
	client.Transport = http.NewCustomTransportWithParentTransport(client.Transport, middlewares...)
	authProvider := auth.NewBaseBearerTokenAuthenticationProvider(accessTokenProvider)
	adapter, err := http.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(authProvider, nil, nil, client)
	if err != nil {
		return nil, err
	}
	adapter.SetBaseUrl(issuer + "/keyhub/rest/v1")
	return adapter, nil
}

func NewKeyHubRequestAdapterForDeviceCode(client *nethttp.Client, issuer string, clientID string, clientSecret string, scopes []string) (*http.NetHttpRequestAdapter, error) {
	if client == nil {
		client = nethttp.DefaultClient
	}
	if client.Timeout == 0 {
		client.Timeout = time.Second * 20
	}
	ctx := oidc.ClientContext(context.Background(), client)
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	oauth2Conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     provider.Endpoint(),
	}
	oauth2Conf.Endpoint.DeviceAuthURL = issuer + "/login/oauth2/authorizedevice"

	response, err := DeviceAuth(ctx, &oauth2Conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("please enter code %s at %s\n", response.UserCode, response.VerificationURIComplete)
	token, err := oauth2Conf.DeviceAccessToken(ctx, response)
	if err != nil {
		panic(err)
	}

	tokenSource := oauth2Conf.TokenSource(ctx, token)
	accessTokenProvider := NewKeyHubAccessTokenProvider(&tokenSource)
	middlewares := []http.Middleware{
		accessTokenProvider,
		http.NewParametersNameDecodingHandler(),
		http.NewUserAgentHandler(),
	}
	client.Transport = http.NewCustomTransportWithParentTransport(client.Transport, middlewares...)
	authProvider := auth.NewBaseBearerTokenAuthenticationProvider(accessTokenProvider)
	adapter, err := http.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(authProvider, nil, nil, client)
	if err != nil {
		return nil, err
	}
	adapter.SetBaseUrl(issuer + "/keyhub/rest/v1")
	return adapter, nil
}

// copied from https://github.com/golang/oauth2/blob/master/deviceauth.go until https://github.com/golang/oauth2/issues/685 is fixed
// DeviceAuth returns a device auth struct which contains a device code
// and authorization information provided for users to enter on another device.
func DeviceAuth(ctx context.Context, c *oauth2.Config, opts ...oauth2.AuthCodeOption) (*oauth2.DeviceAuthResponse, error) {
	// https://datatracker.ietf.org/doc/html/rfc8628#section-3.1
	v := url.Values{
		"client_id": {c.ClientID},
	}
	if len(c.Scopes) > 0 {
		v.Set("scope", strings.Join(c.Scopes, " "))
	}
	v.Set("authVault", "access")
	return retrieveDeviceAuth(ctx, c, v)
}

func retrieveDeviceAuth(ctx context.Context, c *oauth2.Config, v url.Values) (*oauth2.DeviceAuthResponse, error) {
	if c.Endpoint.DeviceAuthURL == "" {
		return nil, errors.New("endpoint missing DeviceAuthURL")
	}

	req, err := nethttp.NewRequest("POST", c.Endpoint.DeviceAuthURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	if c.ClientSecret != "" {
		req.SetBasicAuth(url.QueryEscape(c.ClientID), url.QueryEscape(c.ClientSecret))
	}

	t := time.Now()
	r, err := ContextClient(ctx).Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot auth device: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, &oauth2.RetrieveError{
			Response: r,
			Body:     body,
		}
	}

	da := &oauth2.DeviceAuthResponse{}
	err = json.Unmarshal(body, &da)
	if err != nil {
		return nil, fmt.Errorf("unmarshal %s", err)
	}

	if !da.Expiry.IsZero() {
		// Make a small adjustment to account for time taken by the request
		da.Expiry = da.Expiry.Add(-time.Since(t))
	}

	return da, nil
}

func ContextClient(ctx context.Context) *nethttp.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(oauth2.HTTPClient).(*nethttp.Client); ok {
			return hc
		}
	}
	return nethttp.DefaultClient
}
