package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    icdefaf12ff73f4650b7516ad3b7d930785b9f916d240c5c66cd0f1d083edc2b5 "github.com/topicuskeyhub/sdk-go/account/item/attributevalue/item"
)

// ItemAttributevalueWithAttributevalueItemRequestBuilder builds and executes requests for operations under \account\{accountid}\attributevalue\{attributevalueid}
type ItemAttributevalueWithAttributevalueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAttributevalueWithAttributevalueItemRequestBuilderGetQueryParameters returns the account attribute value identified by the id.
type ItemAttributevalueWithAttributevalueItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []icdefaf12ff73f4650b7516ad3b7d930785b9f916d240c5c66cd0f1d083edc2b5.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAttributevalueWithAttributevalueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributevalueWithAttributevalueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributevalueWithAttributevalueItemRequestBuilderGetQueryParameters
}
// NewItemAttributevalueWithAttributevalueItemRequestBuilderInternal instantiates a new ItemAttributevalueWithAttributevalueItemRequestBuilder and sets the default values.
func NewItemAttributevalueWithAttributevalueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributevalueWithAttributevalueItemRequestBuilder) {
    m := &ItemAttributevalueWithAttributevalueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/attributevalue/{attributevalueid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemAttributevalueWithAttributevalueItemRequestBuilder instantiates a new ItemAttributevalueWithAttributevalueItemRequestBuilder and sets the default values.
func NewItemAttributevalueWithAttributevalueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributevalueWithAttributevalueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributevalueWithAttributevalueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the account attribute value identified by the id.
// returns a IdentityAccountAttributeValueable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributevalueWithAttributevalueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAttributevalueWithAttributevalueItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeValueable), nil
}
// ToGetRequestInformation returns the account attribute value identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAttributevalueWithAttributevalueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAttributevalueWithAttributevalueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=76")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAttributevalueWithAttributevalueItemRequestBuilder when successful
func (m *ItemAttributevalueWithAttributevalueItemRequestBuilder) WithUrl(rawUrl string)(*ItemAttributevalueWithAttributevalueItemRequestBuilder) {
    return NewItemAttributevalueWithAttributevalueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
