package account

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    iaa007f6ced75fe943fc0c292ccaa44a9d3a88b1ccb028ddf0872e8b66543634f "github.com/topicuskeyhub/sdk-go/account/item/attributevalue"
)

// ItemAttributevalueRequestBuilder builds and executes requests for operations under \account\{accountid}\attributevalue
type ItemAttributevalueRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAttributevalueRequestBuilderGetQueryParameters query for account attribute values for a specific account. The various query parameters can be used to filter the response.
type ItemAttributevalueRequestBuilderGetQueryParameters struct {
    // Filter the attribute values by account, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64 `uriparametername:"account"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []iaa007f6ced75fe943fc0c292ccaa44a9d3a88b1ccb028ddf0872e8b66543634f.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the attribute values by attribute definition, specified by id. This parameter supports composition with all parameters from the attribute definition resource.
    Attribute []int64 `uriparametername:"attribute"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the attribute values by the account's directory, specified by id. This parameter supports composition with all parameters from the account directory resource.
    Directory []int64 `uriparametername:"directory"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return attribute values that are, or are not, the latest value for its source.
    LatestForSource []bool `uriparametername:"latestForSource"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter attribute values on the given source(s).
    // Deprecated: This property is deprecated, use SourceAsIdentityAccountAttributeSource instead
    Source []string `uriparametername:"source"`
    // Filter attribute values on the given source(s).
    SourceAsIdentityAccountAttributeSource []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeSource `uriparametername:"source"`
}
// ItemAttributevalueRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributevalueRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributevalueRequestBuilderGetQueryParameters
}
// ByAttributevalueid gets an item from the github.com/topicuskeyhub/sdk-go.account.item.attributevalue.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemAttributevalueWithAttributevalueItemRequestBuilder when successful
func (m *ItemAttributevalueRequestBuilder) ByAttributevalueid(attributevalueid string)(*ItemAttributevalueWithAttributevalueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attributevalueid != "" {
        urlTplParams["attributevalueid"] = attributevalueid
    }
    return NewItemAttributevalueWithAttributevalueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAttributevalueidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.account.item.attributevalue.item collection
// returns a *ItemAttributevalueWithAttributevalueItemRequestBuilder when successful
func (m *ItemAttributevalueRequestBuilder) ByAttributevalueidInt64(attributevalueid int64)(*ItemAttributevalueWithAttributevalueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["attributevalueid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(attributevalueid, 10)
    return NewItemAttributevalueWithAttributevalueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAttributevalueRequestBuilderInternal instantiates a new ItemAttributevalueRequestBuilder and sets the default values.
func NewItemAttributevalueRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributevalueRequestBuilder) {
    m := &ItemAttributevalueRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/attributevalue{?account*,additional*,any*,attribute*,createdAfter*,createdBefore*,directory*,exclude*,id*,latestForSource*,modifiedSince*,q*,sort*,source*}", pathParameters),
    }
    return m
}
// NewItemAttributevalueRequestBuilder instantiates a new ItemAttributevalueRequestBuilder and sets the default values.
func NewItemAttributevalueRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributevalueRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributevalueRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for account attribute values for a specific account. The various query parameters can be used to filter the response.
// returns a IdentityAccountAttributeValueLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributevalueRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAttributevalueRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeValueLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeValueLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeValueLinkableWrapperable), nil
}
// ToGetRequestInformation query for account attribute values for a specific account. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemAttributevalueRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAttributevalueRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=75")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAttributevalueRequestBuilder when successful
func (m *ItemAttributevalueRequestBuilder) WithUrl(rawUrl string)(*ItemAttributevalueRequestBuilder) {
    return NewItemAttributevalueRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
