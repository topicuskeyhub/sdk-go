package webhook

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
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
    Account []int64
    // Only return active or inactive webhooks.
    Active []string
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter webhooks on the given groups performing technical adminstration for the system or client for the webhooks, specified by id.
    AppAdminGroup []int64
    // Filter webhooks on the given groups with ownership for the client or directory for the webhooks, specified by id.
    AppOwnerGroup []int64
    // Filter webhooks on the given clients, specified by id.
    Client []int64
    // Filter webhooks on the given groups performing content adminstration for the system for the webhooks, specified by id.
    ContentAdminGroup []int64
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter webhooks on the given directories, specified by id.
    Directory []int64
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return webhooks that use certificates that are expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return webhooks that are or are not global.
    Global []bool
    // Filter webhooks on the given groups, specified by id.
    Group []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter webhooks on (part of) the name, uuid or URL.
    NameContains []string
    // Filter records on a complex CQL query.
    Q []string
    // Filter webhooks on the given service accounts, specified by id.
    ServiceAccount []int64
    // Filter webhooks on the given service accounts, specified by id. This parameter supports composition with all parameters from the provisioning group resource.
    ServiceAccountNotNull []int64
    // Filter webhooks on the given systems, specified by id.
    System []int64
    // Filter results on the given TLS mode(s).
    Tls []string
    // Only return webhooks that trigger on one of the given type(s).
    Type []string
    // Filter results on one or more URLs.
    Url []string
    // Filter results on one or more UUIDs.
    Uuid []string
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
// WebhookRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WebhookRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.webhook.item collection
func (m *WebhookRequestBuilder) ById(id string)(*WebhookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewWebhookItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWebhookRequestBuilderInternal instantiates a new WebhookRequestBuilder and sets the default values.
func NewWebhookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WebhookRequestBuilder) {
    m := &WebhookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhook{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,account*,active*,appAdminGroup*,appOwnerGroup*,client*,contentAdminGroup*,directory*,expiredCertificate*,global*,group*,nameContains*,serviceAccount*,serviceAccountNotNull*,system*,tls*,type*,url*,uuid*}", pathParameters),
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
func (m *WebhookRequestBuilder) Get(ctx context.Context, requestConfiguration *WebhookRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *WebhookRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, requestConfiguration *WebhookRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *WebhookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WebhookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=65")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new webhooks and returns the newly created webhooks.
func (m *WebhookRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.WebhookWebhookLinkableWrapperable, requestConfiguration *WebhookRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=65")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=65", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
