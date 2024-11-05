package audit

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i427c306d268448021ad9db50f2b9e9fa108d09f04da12eda12e575d7adf99567 "github.com/topicuskeyhub/sdk-go/audit/item"
)

// WithAuditItemRequestBuilder builds and executes requests for operations under \audit\{auditid}
type WithAuditItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithAuditItemRequestBuilderGetQueryParameters returns the audit record identified by the id.
type WithAuditItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i427c306d268448021ad9db50f2b9e9fa108d09f04da12eda12e575d7adf99567.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithAuditItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithAuditItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithAuditItemRequestBuilderGetQueryParameters
}
// NewWithAuditItemRequestBuilderInternal instantiates a new WithAuditItemRequestBuilder and sets the default values.
func NewWithAuditItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAuditItemRequestBuilder) {
    m := &WithAuditItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit/{auditid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithAuditItemRequestBuilder instantiates a new WithAuditItemRequestBuilder and sets the default values.
func NewWithAuditItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAuditItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithAuditItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the audit record identified by the id.
// returns a AuditAuditRecordable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithAuditItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithAuditItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuditAuditRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordable), nil
}
// ToGetRequestInformation returns the audit record identified by the id.
// returns a *RequestInformation when successful
func (m *WithAuditItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithAuditItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithAuditItemRequestBuilder when successful
func (m *WithAuditItemRequestBuilder) WithUrl(rawUrl string)(*WithAuditItemRequestBuilder) {
    return NewWithAuditItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
