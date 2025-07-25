// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package profile

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemAttributesSyncRequestBuilder builds and executes requests for operations under \profile\{accessprofile-id}\attributes\sync
type ItemAttributesSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type ItemAttributesSyncRequestBuilderPostQueryParameters struct {
    Account *int64 `uriparametername:"account"`
    // Deprecated: This property is deprecated, use OperationAsIdentityAccountAttributeUpdateOperation instead
    Operation *string `uriparametername:"operation"`
    OperationAsIdentityAccountAttributeUpdateOperation *ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeUpdateOperation `uriparametername:"operation"`
}
// ItemAttributesSyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributesSyncRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributesSyncRequestBuilderPostQueryParameters
}
// NewItemAttributesSyncRequestBuilderInternal instantiates a new ItemAttributesSyncRequestBuilder and sets the default values.
func NewItemAttributesSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributesSyncRequestBuilder) {
    m := &ItemAttributesSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profile/{accessprofile%2Did}/attributes/sync{?account*,operation*}", pathParameters),
    }
    return m
}
// NewItemAttributesSyncRequestBuilder instantiates a new ItemAttributesSyncRequestBuilder and sets the default values.
func NewItemAttributesSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributesSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributesSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributesSyncRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAttributesSyncRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// returns a *RequestInformation when successful
func (m *ItemAttributesSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAttributesSyncRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json, application/vnd.topicus.keyhub+xml, application/vnd.topicus.keyhub+json;version=79")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAttributesSyncRequestBuilder when successful
func (m *ItemAttributesSyncRequestBuilder) WithUrl(rawUrl string)(*ItemAttributesSyncRequestBuilder) {
    return NewItemAttributesSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
