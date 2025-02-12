package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// VaultUnlockRequestBuilder builds and executes requests for operations under \account\vault\unlock
type VaultUnlockRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VaultUnlockRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VaultUnlockRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVaultUnlockRequestBuilderInternal instantiates a new VaultUnlockRequestBuilder and sets the default values.
func NewVaultUnlockRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VaultUnlockRequestBuilder) {
    m := &VaultUnlockRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/vault/unlock", pathParameters),
    }
    return m
}
// NewVaultUnlockRequestBuilder instantiates a new VaultUnlockRequestBuilder and sets the default values.
func NewVaultUnlockRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VaultUnlockRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVaultUnlockRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unlocks the vault of the authenticated user. Use the session password in the response via the topicus-Vault-session header to access secrets. The prefered way to unlock a vault is via the authentication by passing 'authVault="access"' to the OAuth2 authorization endpoint.
// returns a VaultVaultUnlockResponseable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *VaultUnlockRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultUnlockable, requestConfiguration *VaultUnlockRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultUnlockResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultVaultUnlockResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultUnlockResponseable), nil
}
// ToPostRequestInformation unlocks the vault of the authenticated user. Use the session password in the response via the topicus-Vault-session header to access secrets. The prefered way to unlock a vault is via the authentication by passing 'authVault="access"' to the OAuth2 authorization endpoint.
// returns a *RequestInformation when successful
func (m *VaultUnlockRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultUnlockable, requestConfiguration *VaultUnlockRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=76")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=76", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VaultUnlockRequestBuilder when successful
func (m *VaultUnlockRequestBuilder) WithUrl(rawUrl string)(*VaultUnlockRequestBuilder) {
    return NewVaultUnlockRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
