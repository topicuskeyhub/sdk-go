package group

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i7f644e1f5d438a9c6c1025e5e998db5736b5258fd876fe6a8bb687fd0219316a "github.com/topicuskeyhub/sdk-go/group/item/audit/item"
)

// ItemAuditWithAuditItemRequestBuilder builds and executes requests for operations under \group\{groupid}\audit\{auditid}
type ItemAuditWithAuditItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuditWithAuditItemRequestBuilderGetQueryParameters returns the single audits for the group.
type ItemAuditWithAuditItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i7f644e1f5d438a9c6c1025e5e998db5736b5258fd876fe6a8bb687fd0219316a.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAuditWithAuditItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuditWithAuditItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuditWithAuditItemRequestBuilderGetQueryParameters
}
// ItemAuditWithAuditItemRequestBuilderPutQueryParameters updates the audits for a group identified by the id. Only audits in draft can be updated.
type ItemAuditWithAuditItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []i7f644e1f5d438a9c6c1025e5e998db5736b5258fd876fe6a8bb687fd0219316a.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAuditWithAuditItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuditWithAuditItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuditWithAuditItemRequestBuilderPutQueryParameters
}
// NewItemAuditWithAuditItemRequestBuilderInternal instantiates a new ItemAuditWithAuditItemRequestBuilder and sets the default values.
func NewItemAuditWithAuditItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditWithAuditItemRequestBuilder) {
    m := &ItemAuditWithAuditItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}/audit/{auditid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemAuditWithAuditItemRequestBuilder instantiates a new ItemAuditWithAuditItemRequestBuilder and sets the default values.
func NewItemAuditWithAuditItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditWithAuditItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditWithAuditItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the single audits for the group.
// returns a AuditGroupAuditable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAuditWithAuditItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuditWithAuditItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuditGroupAuditFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable), nil
}
// Put updates the audits for a group identified by the id. Only audits in draft can be updated.
// returns a AuditGroupAuditable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAuditWithAuditItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable, requestConfiguration *ItemAuditWithAuditItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuditGroupAuditFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable), nil
}
// ToGetRequestInformation returns the single audits for the group.
// returns a *RequestInformation when successful
func (m *ItemAuditWithAuditItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuditWithAuditItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the audits for a group identified by the id. Only audits in draft can be updated.
// returns a *RequestInformation when successful
func (m *ItemAuditWithAuditItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditGroupAuditable, requestConfiguration *ItemAuditWithAuditItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
// returns a *ItemAuditWithAuditItemRequestBuilder when successful
func (m *ItemAuditWithAuditItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuditWithAuditItemRequestBuilder) {
    return NewItemAuditWithAuditItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
