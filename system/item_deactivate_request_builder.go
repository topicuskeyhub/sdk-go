package system

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemDeactivateRequestBuilder builds and executes requests for operations under \system\{systemid}\deactivate
type ItemDeactivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeactivateRequestBuilderPostQueryParameters deactivates the provisioned system specified by the given id. The action runs in the background.
type ItemDeactivateRequestBuilderPostQueryParameters struct {
    // Deprecated: This property is deprecated, use ActionAsProvisioningDeprovisionAction instead
    Action *string `uriparametername:"action"`
    ActionAsProvisioningDeprovisionAction *ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningDeprovisionAction `uriparametername:"action"`
}
// ItemDeactivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeactivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDeactivateRequestBuilderPostQueryParameters
}
// NewItemDeactivateRequestBuilderInternal instantiates a new ItemDeactivateRequestBuilder and sets the default values.
func NewItemDeactivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeactivateRequestBuilder) {
    m := &ItemDeactivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}/deactivate{?action*}", pathParameters),
    }
    return m
}
// NewItemDeactivateRequestBuilder instantiates a new ItemDeactivateRequestBuilder and sets the default values.
func NewItemDeactivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeactivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeactivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post deactivates the provisioned system specified by the given id. The action runs in the background.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemDeactivateRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemDeactivateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation deactivates the provisioned system specified by the given id. The action runs in the background.
// returns a *RequestInformation when successful
func (m *ItemDeactivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemDeactivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemDeactivateRequestBuilder when successful
func (m *ItemDeactivateRequestBuilder) WithUrl(rawUrl string)(*ItemDeactivateRequestBuilder) {
    return NewItemDeactivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
