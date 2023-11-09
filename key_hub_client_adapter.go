// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package sdkgo

// not-generated
//go:generate bash generate.sh

import (
	"context"
	"fmt"
	nethttp "net/http"
	"net/url"
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

	response, err := oauth2Conf.DeviceAuth(ctx, oauth2.SetAuthURLParam("authVault", "access"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("please enter code %s at %s\n", response.UserCode, response.VerificationURIComplete)
	token, err := oauth2Conf.DeviceAccessToken(ctx, response)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

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
