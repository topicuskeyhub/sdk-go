package group

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    icf209cccf4491e65048f97716247b282df93bf60a03ad3acaf1a2fae47ff7753 "github.com/topicuskeyhub/sdk-go/group/item"
)

// WithGroupItemRequestBuilder builds and executes requests for operations under \group\{groupid}
type WithGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithGroupItemRequestBuilderGetQueryParameters returns the group identified by the id.
type WithGroupItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []icf209cccf4491e65048f97716247b282df93bf60a03ad3acaf1a2fae47ff7753.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithGroupItemRequestBuilderGetQueryParameters
}
// WithGroupItemRequestBuilderPutQueryParameters updates the group identified by the id.
type WithGroupItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []icf209cccf4491e65048f97716247b282df93bf60a03ad3acaf1a2fae47ff7753.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithGroupItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithGroupItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithGroupItemRequestBuilderPutQueryParameters
}
// Account the account property
// returns a *ItemAccountRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Account()(*ItemAccountRequestBuilder) {
    return NewItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Audit the audit property
// returns a *ItemAuditRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Audit()(*ItemAuditRequestBuilder) {
    return NewItemAuditRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classification the classification property
// returns a *ItemClassificationRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Classification()(*ItemClassificationRequestBuilder) {
    return NewItemClassificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithGroupItemRequestBuilderInternal instantiates a new WithGroupItemRequestBuilder and sets the default values.
func NewWithGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithGroupItemRequestBuilder) {
    m := &WithGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithGroupItemRequestBuilder instantiates a new WithGroupItemRequestBuilder and sets the default values.
func NewWithGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the group identified by the id.
// returns a GroupGroupable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithGroupItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable), nil
}
// Mail the mail property
// returns a *ItemMailRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Mail()(*ItemMailRequestBuilder) {
    return NewItemMailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the group identified by the id.
// returns a GroupGroupable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithGroupItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable, requestConfiguration *WithGroupItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable), nil
}
// Requestaudit the requestaudit property
// returns a *ItemRequestauditRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Requestaudit()(*ItemRequestauditRequestBuilder) {
    return NewItemRequestauditRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns the group identified by the id.
// returns a *RequestInformation when successful
func (m *WithGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the group identified by the id.
// returns a *RequestInformation when successful
func (m *WithGroupItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupable, requestConfiguration *WithGroupItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Vault the vault property
// returns a *ItemVaultRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) Vault()(*ItemVaultRequestBuilder) {
    return NewItemVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithGroupItemRequestBuilder when successful
func (m *WithGroupItemRequestBuilder) WithUrl(rawUrl string)(*WithGroupItemRequestBuilder) {
    return NewWithGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
