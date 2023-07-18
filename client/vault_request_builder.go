package client

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// VaultRequestBuilder builds and executes requests for operations under \client\vault
type VaultRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewVaultRequestBuilderInternal instantiates a new VaultRequestBuilder and sets the default values.
func NewVaultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VaultRequestBuilder) {
    m := &VaultRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/vault", pathParameters),
    }
    return m
}
// NewVaultRequestBuilder instantiates a new VaultRequestBuilder and sets the default values.
func NewVaultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VaultRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVaultRequestBuilderInternal(urlParams, requestAdapter)
}
// Unlock the unlock property
func (m *VaultRequestBuilder) Unlock()(*VaultUnlockRequestBuilder) {
    return NewVaultUnlockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
