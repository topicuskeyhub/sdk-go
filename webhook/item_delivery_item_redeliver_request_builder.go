package webhook

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemDeliveryItemRedeliverRequestBuilder builds and executes requests for operations under \webhook\{webhookid}\delivery\{deliveryid}\redeliver
type ItemDeliveryItemRedeliverRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeliveryItemRedeliverRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeliveryItemRedeliverRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeliveryItemRedeliverRequestBuilderInternal instantiates a new ItemDeliveryItemRedeliverRequestBuilder and sets the default values.
func NewItemDeliveryItemRedeliverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryItemRedeliverRequestBuilder) {
    m := &ItemDeliveryItemRedeliverRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhook/{webhookid}/delivery/{deliveryid}/redeliver", pathParameters),
    }
    return m
}
// NewItemDeliveryItemRedeliverRequestBuilder instantiates a new ItemDeliveryItemRedeliverRequestBuilder and sets the default values.
func NewItemDeliveryItemRedeliverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryItemRedeliverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeliveryItemRedeliverRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform a redelivery of the webhook payload.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemDeliveryItemRedeliverRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemDeliveryItemRedeliverRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation perform a redelivery of the webhook payload.
// returns a *RequestInformation when successful
func (m *ItemDeliveryItemRedeliverRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemDeliveryItemRedeliverRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDeliveryItemRedeliverRequestBuilder when successful
func (m *ItemDeliveryItemRedeliverRequestBuilder) WithUrl(rawUrl string)(*ItemDeliveryItemRedeliverRequestBuilder) {
    return NewItemDeliveryItemRedeliverRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
