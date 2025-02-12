package webhook

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i4195494863fbff0657373f0bf8749078b5ee02941417b30d38cce9d41f12ddb3 "github.com/topicuskeyhub/sdk-go/webhook/item/delivery/item"
)

// ItemDeliveryWithDeliveryItemRequestBuilder builds and executes requests for operations under \webhook\{webhookid}\delivery\{deliveryid}
type ItemDeliveryWithDeliveryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeliveryWithDeliveryItemRequestBuilderGetQueryParameters returns the single webhook delivery.
type ItemDeliveryWithDeliveryItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i4195494863fbff0657373f0bf8749078b5ee02941417b30d38cce9d41f12ddb3.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemDeliveryWithDeliveryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeliveryWithDeliveryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDeliveryWithDeliveryItemRequestBuilderGetQueryParameters
}
// NewItemDeliveryWithDeliveryItemRequestBuilderInternal instantiates a new ItemDeliveryWithDeliveryItemRequestBuilder and sets the default values.
func NewItemDeliveryWithDeliveryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryWithDeliveryItemRequestBuilder) {
    m := &ItemDeliveryWithDeliveryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhook/{webhookid}/delivery/{deliveryid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemDeliveryWithDeliveryItemRequestBuilder instantiates a new ItemDeliveryWithDeliveryItemRequestBuilder and sets the default values.
func NewItemDeliveryWithDeliveryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryWithDeliveryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeliveryWithDeliveryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the single webhook delivery.
// returns a WebhookWebhookDeliveryable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemDeliveryWithDeliveryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDeliveryWithDeliveryItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookDeliveryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateWebhookWebhookDeliveryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookDeliveryable), nil
}
// Redeliver the redeliver property
// returns a *ItemDeliveryItemRedeliverRequestBuilder when successful
func (m *ItemDeliveryWithDeliveryItemRequestBuilder) Redeliver()(*ItemDeliveryItemRedeliverRequestBuilder) {
    return NewItemDeliveryItemRedeliverRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns the single webhook delivery.
// returns a *RequestInformation when successful
func (m *ItemDeliveryWithDeliveryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDeliveryWithDeliveryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDeliveryWithDeliveryItemRequestBuilder when successful
func (m *ItemDeliveryWithDeliveryItemRequestBuilder) WithUrl(rawUrl string)(*ItemDeliveryWithDeliveryItemRequestBuilder) {
    return NewItemDeliveryWithDeliveryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
