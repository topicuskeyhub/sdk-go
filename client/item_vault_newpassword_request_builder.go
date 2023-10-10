package client

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemVaultNewpasswordRequestBuilder builds and executes requests for operations under \client\{clientid}\vault\newpassword
type ItemVaultNewpasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVaultNewpasswordRequestBuilderGetQueryParameters returns a random generated password using the given strategy, or the default strategy if none given.
type ItemVaultNewpasswordRequestBuilderGetQueryParameters struct {
    // 
    Strategy *string `uriparametername:"strategy"`
}
// ItemVaultNewpasswordRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultNewpasswordRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultNewpasswordRequestBuilderGetQueryParameters
}
// NewItemVaultNewpasswordRequestBuilderInternal instantiates a new NewpasswordRequestBuilder and sets the default values.
func NewItemVaultNewpasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultNewpasswordRequestBuilder) {
    m := &ItemVaultNewpasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/{clientid}/vault/newpassword{?strategy*}", pathParameters),
    }
    return m
}
// NewItemVaultNewpasswordRequestBuilder instantiates a new NewpasswordRequestBuilder and sets the default values.
func NewItemVaultNewpasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultNewpasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVaultNewpasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a random generated password using the given strategy, or the default strategy if none given.
func (m *ItemVaultNewpasswordRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemVaultNewpasswordRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.SimpleStringValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// ToGetRequestInformation returns a random generated password using the given strategy, or the default strategy if none given.
func (m *ItemVaultNewpasswordRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemVaultNewpasswordRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=66")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemVaultNewpasswordRequestBuilder) WithUrl(rawUrl string)(*ItemVaultNewpasswordRequestBuilder) {
    return NewItemVaultNewpasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
