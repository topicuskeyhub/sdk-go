package webhook

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// WebhookRequestBuilder builds and executes requests for operations under \webhook
type WebhookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WebhookRequestBuilderGetQueryParameters queries over all webhooks. The various query parameters can be used to filter the response.
type WebhookRequestBuilderGetQueryParameters struct {
    // Filter webhooks on the given accounts, specified by id.
    Account []int64 `uriparametername:"account"`
    // Only return active or inactive webhooks.
    // Deprecated: This property is deprecated, use ActiveAsBooleanEnum instead
    Active []string `uriparametername:"active"`
    // Only return active or inactive webhooks.
    ActiveAsBooleanEnum []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.BooleanEnum `uriparametername:"active"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter webhooks on the given groups performing technical administration for the client for the webhooks, specified by id.
    AppAdminGroup []int64 `uriparametername:"appAdminGroup"`
    // Filter the webhooks by groups owning the client, specified by id.
    AppOwnerGroup []int64 `uriparametername:"appOwnerGroup"`
    // Filter webhooks on the given clients, specified by id.
    Client []int64 `uriparametername:"client"`
    // Filter webhooks on the given groups performing content administration for the system for the webhooks, specified by id.
    ContentAdminGroup []int64 `uriparametername:"contentAdminGroup"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter webhooks on the given directories, specified by id.
    Directory []int64 `uriparametername:"directory"`
    // Filter the webhooks by groups owning the internal directory, specified by id.
    DirectoryOwnerGroup []int64 `uriparametername:"directoryOwnerGroup"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return webhooks that use certificates that are expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"expiredCertificate"`
    // Only return webhooks that are or are not global.
    Global []bool `uriparametername:"global"`
    // Filter webhooks on the given groups, specified by id.
    Group []int64 `uriparametername:"group"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter webhooks on (part of) the name, uuid or URL.
    NameContains []string `uriparametername:"nameContains"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter webhooks on the given service accounts, specified by id.
    ServiceAccount []int64 `uriparametername:"serviceAccount"`
    // Filter webhooks on the given service accounts, specified by id. This parameter supports composition with all parameters from the provisioning group resource.
    ServiceAccountNotNull []int64 `uriparametername:"serviceAccountNotNull"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter webhooks on the given systems, specified by id.
    System []int64 `uriparametername:"system"`
    // Filter webhooks on the given groups performing technical administration for the system for the webhooks, specified by id.
    SystemAdminGroup []int64 `uriparametername:"systemAdminGroup"`
    // Filter the webhooks by groups owning the provisioned system, specified by id.
    SystemOwnerGroup []int64 `uriparametername:"systemOwnerGroup"`
    // Filter results on the given TLS mode(s).
    // Deprecated: This property is deprecated, use TlsAsTLSLevel instead
    Tls []string `uriparametername:"tls"`
    // Filter results on the given TLS mode(s).
    TlsAsTLSLevel []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.TLSLevel `uriparametername:"tls"`
    // Only return webhooks that trigger on one of the given type(s).
    // Deprecated: This property is deprecated, use TypeAsAuditAuditRecordType instead
    Type []string `uriparametername:"type"`
    // Only return webhooks that trigger on one of the given type(s).
    TypeAsAuditAuditRecordType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordType `uriparametername:"type"`
    // Filter results on one or more URLs.
    Url []string `uriparametername:"url"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// WebhookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WebhookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WebhookRequestBuilderGetQueryParameters
}
// WebhookRequestBuilderPostQueryParameters creates one or more new webhooks and returns the newly created webhooks.
type WebhookRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// WebhookRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WebhookRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WebhookRequestBuilderPostQueryParameters
}
// ByWebhookid gets an item from the github.com/topicuskeyhub/sdk-go.webhook.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithWebhookItemRequestBuilder when successful
func (m *WebhookRequestBuilder) ByWebhookid(webhookid string)(*WithWebhookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if webhookid != "" {
        urlTplParams["webhookid"] = webhookid
    }
    return NewWithWebhookItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByWebhookidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.webhook.item collection
// returns a *WithWebhookItemRequestBuilder when successful
func (m *WebhookRequestBuilder) ByWebhookidInt64(webhookid int64)(*WithWebhookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["webhookid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(webhookid, 10)
    return NewWithWebhookItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWebhookRequestBuilderInternal instantiates a new WebhookRequestBuilder and sets the default values.
func NewWebhookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WebhookRequestBuilder) {
    m := &WebhookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhook{?account*,active*,additional*,any*,appAdminGroup*,appOwnerGroup*,client*,contentAdminGroup*,createdAfter*,createdBefore*,directory*,directoryOwnerGroup*,exclude*,expiredCertificate*,global*,group*,id*,modifiedSince*,nameContains*,q*,serviceAccount*,serviceAccountNotNull*,sort*,system*,systemAdminGroup*,systemOwnerGroup*,tls*,type*,url*,uuid*}", pathParameters),
    }
    return m
}
// NewWebhookRequestBuilder instantiates a new WebhookRequestBuilder and sets the default values.
func NewWebhookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WebhookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWebhookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all webhooks. The various query parameters can be used to filter the response.
// returns a WebhookWebhookLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WebhookRequestBuilder) Get(ctx context.Context, requestConfiguration *WebhookRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateWebhookWebhookLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable), nil
}
// Post creates one or more new webhooks and returns the newly created webhooks.
// returns a WebhookWebhookLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WebhookRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, requestConfiguration *WebhookRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateWebhookWebhookLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all webhooks. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *WebhookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WebhookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new webhooks and returns the newly created webhooks.
// returns a *RequestInformation when successful
func (m *WebhookRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, requestConfiguration *WebhookRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WebhookRequestBuilder when successful
func (m *WebhookRequestBuilder) WithUrl(rawUrl string)(*WebhookRequestBuilder) {
    return NewWebhookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
