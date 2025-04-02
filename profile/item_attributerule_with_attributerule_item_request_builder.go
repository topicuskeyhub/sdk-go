package profile

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i7c76b17c0ce5210fbfe0d58e16f93682753e56559458acdb5f85fb73d460a4cd "github.com/topicuskeyhub/sdk-go/profile/item/attributerule/item"
)

// ItemAttributeruleWithAttributeruleItemRequestBuilder builds and executes requests for operations under \profile\{accessprofile-id}\attributerule\{attributeruleid}
type ItemAttributeruleWithAttributeruleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAttributeruleWithAttributeruleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributeruleWithAttributeruleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAttributeruleWithAttributeruleItemRequestBuilderGetQueryParameters returns the account attribute rule identified by the id.
type ItemAttributeruleWithAttributeruleItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i7c76b17c0ce5210fbfe0d58e16f93682753e56559458acdb5f85fb73d460a4cd.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAttributeruleWithAttributeruleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributeruleWithAttributeruleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributeruleWithAttributeruleItemRequestBuilderGetQueryParameters
}
// ItemAttributeruleWithAttributeruleItemRequestBuilderPutQueryParameters updates the account attribute rule identified by the id.
type ItemAttributeruleWithAttributeruleItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []i7c76b17c0ce5210fbfe0d58e16f93682753e56559458acdb5f85fb73d460a4cd.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAttributeruleWithAttributeruleItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributeruleWithAttributeruleItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributeruleWithAttributeruleItemRequestBuilderPutQueryParameters
}
// NewItemAttributeruleWithAttributeruleItemRequestBuilderInternal instantiates a new ItemAttributeruleWithAttributeruleItemRequestBuilder and sets the default values.
func NewItemAttributeruleWithAttributeruleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributeruleWithAttributeruleItemRequestBuilder) {
    m := &ItemAttributeruleWithAttributeruleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profile/{accessprofile%2Did}/attributerule/{attributeruleid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemAttributeruleWithAttributeruleItemRequestBuilder instantiates a new ItemAttributeruleWithAttributeruleItemRequestBuilder and sets the default values.
func NewItemAttributeruleWithAttributeruleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributeruleWithAttributeruleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributeruleWithAttributeruleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the account attribute rule identified by the id.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the account attribute rule identified by the id.
// returns a IdentityAccountAttributeRuleable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable), nil
}
// Put updates the account attribute rule identified by the id.
// returns a IdentityAccountAttributeRuleable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable), nil
}
// ToDeleteRequestInformation deletes the account attribute rule identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// ToGetRequestInformation returns the account attribute rule identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// ToPutRequestInformation updates the account attribute rule identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleable, requestConfiguration *ItemAttributeruleWithAttributeruleItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=77", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAttributeruleWithAttributeruleItemRequestBuilder when successful
func (m *ItemAttributeruleWithAttributeruleItemRequestBuilder) WithUrl(rawUrl string)(*ItemAttributeruleWithAttributeruleItemRequestBuilder) {
    return NewItemAttributeruleWithAttributeruleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
