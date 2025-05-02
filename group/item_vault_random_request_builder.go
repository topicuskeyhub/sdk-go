package group

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemVaultRandomRequestBuilder builds and executes requests for operations under \group\{groupid}\vault\random
type ItemVaultRandomRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVaultRandomRequestBuilderGetQueryParameters returns a random generated value of the given length, default 24.
type ItemVaultRandomRequestBuilderGetQueryParameters struct {
    Length *int32 `uriparametername:"length"`
}
// ItemVaultRandomRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRandomRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultRandomRequestBuilderGetQueryParameters
}
// NewItemVaultRandomRequestBuilderInternal instantiates a new ItemVaultRandomRequestBuilder and sets the default values.
func NewItemVaultRandomRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRandomRequestBuilder) {
    m := &ItemVaultRandomRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}/vault/random{?length*}", pathParameters),
    }
    return m
}
// NewItemVaultRandomRequestBuilder instantiates a new ItemVaultRandomRequestBuilder and sets the default values.
func NewItemVaultRandomRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRandomRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVaultRandomRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a random generated value of the given length, default 24.
// returns a SimpleStringValueable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemVaultRandomRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemVaultRandomRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.SimpleStringValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateSimpleStringValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.SimpleStringValueable), nil
}
// ToGetRequestInformation returns a random generated value of the given length, default 24.
// returns a *RequestInformation when successful
func (m *ItemVaultRandomRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemVaultRandomRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVaultRandomRequestBuilder when successful
func (m *ItemVaultRandomRequestBuilder) WithUrl(rawUrl string)(*ItemVaultRandomRequestBuilder) {
    return NewItemVaultRandomRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
