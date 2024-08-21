package system

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    ic1e8631b3e2165398d89fdc41450d544d009fe723c5916e33ffb2b76f2df69fa "github.com/topicuskeyhub/sdk-go/system/item/account/item"
)

// ItemAccountWithAccountItemRequestBuilder builds and executes requests for operations under \system\{systemid}\account\{accountid}
type ItemAccountWithAccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAccountWithAccountItemRequestBuilderGetQueryParameters returns the provisioned account identified by the id.
type ItemAccountWithAccountItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []ic1e8631b3e2165398d89fdc41450d544d009fe723c5916e33ffb2b76f2df69fa.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAccountWithAccountItemRequestBuilderGetQueryParameters
}
// NewItemAccountWithAccountItemRequestBuilderInternal instantiates a new ItemAccountWithAccountItemRequestBuilder and sets the default values.
func NewItemAccountWithAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountWithAccountItemRequestBuilder) {
    m := &ItemAccountWithAccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}/account/{accountid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemAccountWithAccountItemRequestBuilder instantiates a new ItemAccountWithAccountItemRequestBuilder and sets the default values.
func NewItemAccountWithAccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountWithAccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAccountWithAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the provisioned account identified by the id.
// returns a ProvisioningProvisionedAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountWithAccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedAccountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedAccountable), nil
}
// ToGetRequestInformation returns the provisioned account identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAccountWithAccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAccountWithAccountItemRequestBuilder when successful
func (m *ItemAccountWithAccountItemRequestBuilder) WithUrl(rawUrl string)(*ItemAccountWithAccountItemRequestBuilder) {
    return NewItemAccountWithAccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
