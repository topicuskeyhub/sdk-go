package webhook

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i7a3a7e28e53d5febda97bdd9c0aa6853b66b70a46e27b5f8e14390be27ff7166 "github.com/topicuskeyhub/sdk-go/webhook/item/delivery"
)

// ItemDeliveryRequestBuilder builds and executes requests for operations under \webhook\{webhookid}\delivery
type ItemDeliveryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeliveryRequestBuilderGetQueryParameters queries over all deliveries for a webhook. The various query parameters can be used to filter the response.
type ItemDeliveryRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i7a3a7e28e53d5febda97bdd9c0aa6853b66b70a46e27b5f8e14390be27ff7166.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter deliveries on the given webhooks, specified by id. This parameter supports composition with all parameters from the webhook resource.
    Webhook []int64 `uriparametername:"webhook"`
}
// ItemDeliveryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeliveryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDeliveryRequestBuilderGetQueryParameters
}
// ByDeliveryid gets an item from the github.com/topicuskeyhub/sdk-go.webhook.item.delivery.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemDeliveryWithDeliveryItemRequestBuilder when successful
func (m *ItemDeliveryRequestBuilder) ByDeliveryid(deliveryid string)(*ItemDeliveryWithDeliveryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deliveryid != "" {
        urlTplParams["deliveryid"] = deliveryid
    }
    return NewItemDeliveryWithDeliveryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDeliveryidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.webhook.item.delivery.item collection
// returns a *ItemDeliveryWithDeliveryItemRequestBuilder when successful
func (m *ItemDeliveryRequestBuilder) ByDeliveryidInt64(deliveryid int64)(*ItemDeliveryWithDeliveryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["deliveryid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(deliveryid, 10)
    return NewItemDeliveryWithDeliveryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDeliveryRequestBuilderInternal instantiates a new ItemDeliveryRequestBuilder and sets the default values.
func NewItemDeliveryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryRequestBuilder) {
    m := &ItemDeliveryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhook/{webhookid}/delivery{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,sort*,webhook*}", pathParameters),
    }
    return m
}
// NewItemDeliveryRequestBuilder instantiates a new ItemDeliveryRequestBuilder and sets the default values.
func NewItemDeliveryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeliveryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeliveryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all deliveries for a webhook. The various query parameters can be used to filter the response.
// returns a WebhookWebhookDeliveryLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemDeliveryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDeliveryRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookDeliveryLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateWebhookWebhookDeliveryLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookDeliveryLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all deliveries for a webhook. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemDeliveryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDeliveryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDeliveryRequestBuilder when successful
func (m *ItemDeliveryRequestBuilder) WithUrl(rawUrl string)(*ItemDeliveryRequestBuilder) {
    return NewItemDeliveryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
