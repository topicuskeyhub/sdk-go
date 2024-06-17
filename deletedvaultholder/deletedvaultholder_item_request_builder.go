package deletedvaultholder

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i6988fee5c526fbb1d3ba4158663dc55e95163db39ce1ea0400f734d0b22775fc "github.com/topicuskeyhub/sdk-go/deletedvaultholder/item"
)

// DeletedvaultholderItemRequestBuilder builds and executes requests for operations under \deletedvaultholder\{deletedvaultholder-id}
type DeletedvaultholderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedvaultholderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedvaultholderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedvaultholderItemRequestBuilderGetQueryParameters returns the deleted vault identified by the id.
type DeletedvaultholderItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i6988fee5c526fbb1d3ba4158663dc55e95163db39ce1ea0400f734d0b22775fc.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// DeletedvaultholderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedvaultholderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedvaultholderItemRequestBuilderGetQueryParameters
}
// NewDeletedvaultholderItemRequestBuilderInternal instantiates a new DeletedvaultholderItemRequestBuilder and sets the default values.
func NewDeletedvaultholderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedvaultholderItemRequestBuilder) {
    m := &DeletedvaultholderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deletedvaultholder/{deletedvaultholder%2Did}{?additional*}", pathParameters),
    }
    return m
}
// NewDeletedvaultholderItemRequestBuilder instantiates a new DeletedvaultholderItemRequestBuilder and sets the default values.
func NewDeletedvaultholderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedvaultholderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedvaultholderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete permanently removes the deleted vault identified by the id. This cannot be undone.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *DeletedvaultholderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedvaultholderItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns the deleted vault identified by the id.
// returns a VaultDeletedVaultHolderable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *DeletedvaultholderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedvaultholderItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultDeletedVaultHolderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultDeletedVaultHolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultDeletedVaultHolderable), nil
}
// Recover the recover property
// returns a *ItemRecoverRequestBuilder when successful
func (m *DeletedvaultholderItemRequestBuilder) Recover()(*ItemRecoverRequestBuilder) {
    return NewItemRecoverRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation permanently removes the deleted vault identified by the id. This cannot be undone.
// returns a *RequestInformation when successful
func (m *DeletedvaultholderItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedvaultholderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// ToGetRequestInformation returns the deleted vault identified by the id.
// returns a *RequestInformation when successful
func (m *DeletedvaultholderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedvaultholderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// Vault the vault property
// returns a *ItemVaultRequestBuilder when successful
func (m *DeletedvaultholderItemRequestBuilder) Vault()(*ItemVaultRequestBuilder) {
    return NewItemVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedvaultholderItemRequestBuilder when successful
func (m *DeletedvaultholderItemRequestBuilder) WithUrl(rawUrl string)(*DeletedvaultholderItemRequestBuilder) {
    return NewDeletedvaultholderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}