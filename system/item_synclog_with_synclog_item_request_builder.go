package system

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    ib6cb7080b13110e23dfc9f9630ecc76d39e0b39b8d3f713b41b2ff989e91459a "github.com/topicuskeyhub/sdk-go/system/item/synclog/item"
)

// ItemSynclogWithSynclogItemRequestBuilder builds and executes requests for operations under \system\{systemid}\synclog\{synclogid}
type ItemSynclogWithSynclogItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynclogWithSynclogItemRequestBuilderGetQueryParameters returns the sync log for a provisioned system.
type ItemSynclogWithSynclogItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []ib6cb7080b13110e23dfc9f9630ecc76d39e0b39b8d3f713b41b2ff989e91459a.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemSynclogWithSynclogItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynclogWithSynclogItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSynclogWithSynclogItemRequestBuilderGetQueryParameters
}
// NewItemSynclogWithSynclogItemRequestBuilderInternal instantiates a new ItemSynclogWithSynclogItemRequestBuilder and sets the default values.
func NewItemSynclogWithSynclogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynclogWithSynclogItemRequestBuilder) {
    m := &ItemSynclogWithSynclogItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}/synclog/{synclogid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemSynclogWithSynclogItemRequestBuilder instantiates a new ItemSynclogWithSynclogItemRequestBuilder and sets the default values.
func NewItemSynclogWithSynclogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynclogWithSynclogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynclogWithSynclogItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the sync log for a provisioned system.
// returns a ProvisioningProvisionedSystemSyncLogable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemSynclogWithSynclogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynclogWithSynclogItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemSyncLogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemSyncLogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemSyncLogable), nil
}
// ToGetRequestInformation returns the sync log for a provisioned system.
// returns a *RequestInformation when successful
func (m *ItemSynclogWithSynclogItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynclogWithSynclogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=73")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSynclogWithSynclogItemRequestBuilder when successful
func (m *ItemSynclogWithSynclogItemRequestBuilder) WithUrl(rawUrl string)(*ItemSynclogWithSynclogItemRequestBuilder) {
    return NewItemSynclogWithSynclogItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
