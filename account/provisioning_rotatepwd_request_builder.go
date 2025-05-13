package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ProvisioningRotatepwdRequestBuilder builds and executes requests for operations under \account\provisioning\rotatepwd
type ProvisioningRotatepwdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProvisioningRotatepwdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProvisioningRotatepwdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProvisioningRotatepwdRequestBuilderInternal instantiates a new ProvisioningRotatepwdRequestBuilder and sets the default values.
func NewProvisioningRotatepwdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningRotatepwdRequestBuilder) {
    m := &ProvisioningRotatepwdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/provisioning/rotatepwd", pathParameters),
    }
    return m
}
// NewProvisioningRotatepwdRequestBuilder instantiates a new ProvisioningRotatepwdRequestBuilder and sets the default values.
func NewProvisioningRotatepwdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningRotatepwdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvisioningRotatepwdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post manually rotates the user's rotating password. This is normally done once a day automatically.
// returns a ProvisioningAccountProvisioningStatusReportable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ProvisioningRotatepwdRequestBuilder) Post(ctx context.Context, requestConfiguration *ProvisioningRotatepwdRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningAccountProvisioningStatusReportable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningAccountProvisioningStatusReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningAccountProvisioningStatusReportable), nil
}
// ToPostRequestInformation manually rotates the user's rotating password. This is normally done once a day automatically.
// returns a *RequestInformation when successful
func (m *ProvisioningRotatepwdRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ProvisioningRotatepwdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProvisioningRotatepwdRequestBuilder when successful
func (m *ProvisioningRotatepwdRequestBuilder) WithUrl(rawUrl string)(*ProvisioningRotatepwdRequestBuilder) {
    return NewProvisioningRotatepwdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
