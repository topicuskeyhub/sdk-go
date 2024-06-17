package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ProvisioningTokenpwdRequestBuilder builds and executes requests for operations under \account\provisioning\tokenpwd
type ProvisioningTokenpwdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProvisioningTokenpwdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProvisioningTokenpwdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProvisioningTokenpwdRequestBuilderInternal instantiates a new ProvisioningTokenpwdRequestBuilder and sets the default values.
func NewProvisioningTokenpwdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningTokenpwdRequestBuilder) {
    m := &ProvisioningTokenpwdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/provisioning/tokenpwd", pathParameters),
    }
    return m
}
// NewProvisioningTokenpwdRequestBuilder instantiates a new ProvisioningTokenpwdRequestBuilder and sets the default values.
func NewProvisioningTokenpwdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningTokenpwdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvisioningTokenpwdRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the user's rotating password as it is set for the current user.
// returns a ProvisioningTokenPasswordable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ProvisioningTokenpwdRequestBuilder) Get(ctx context.Context, requestConfiguration *ProvisioningTokenpwdRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningTokenPasswordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningTokenPasswordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningTokenPasswordable), nil
}
// ToGetRequestInformation returns the user's rotating password as it is set for the current user.
// returns a *RequestInformation when successful
func (m *ProvisioningTokenpwdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProvisioningTokenpwdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProvisioningTokenpwdRequestBuilder when successful
func (m *ProvisioningTokenpwdRequestBuilder) WithUrl(rawUrl string)(*ProvisioningTokenpwdRequestBuilder) {
    return NewProvisioningTokenpwdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
