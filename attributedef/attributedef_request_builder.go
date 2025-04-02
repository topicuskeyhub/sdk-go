package attributedef

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// AttributedefRequestBuilder builds and executes requests for operations under \attributedef
type AttributedefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttributedefRequestBuilderGetQueryParameters query for all account attribute definitions in Topicus KeyHub. The various query parameters can be used to filter the response.
type AttributedefRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter account attribute definitions on the given format(s).
    // Deprecated: This property is deprecated, use FormatAsIdentityAccountAttributeFormat instead
    Format []string `uriparametername:"format"`
    // Filter account attribute definitions on the given format(s).
    FormatAsIdentityAccountAttributeFormat []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeFormat `uriparametername:"format"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Only return account attributes for which the given access profiles do not have an attribute rule, specified by id.
    NoRuleForAccessProfile []int64 `uriparametername:"noRuleForAccessProfile"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter account attribute definitions on the given system definition(s).
    // Deprecated: This property is deprecated, use SystemDefinitionAsIdentityAccountAttributeSystemDefinition instead
    SystemDefinition []string `uriparametername:"systemDefinition"`
    // Filter account attribute definitions on the given system definition(s).
    SystemDefinitionAsIdentityAccountAttributeSystemDefinition []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeSystemDefinition `uriparametername:"systemDefinition"`
}
// AttributedefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttributedefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttributedefRequestBuilderGetQueryParameters
}
// AttributedefRequestBuilderPostQueryParameters creates one or more new account attribute definitions and returns the newly created account attribute definitions.
type AttributedefRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// AttributedefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttributedefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttributedefRequestBuilderPostQueryParameters
}
// ByAttributedefid gets an item from the github.com/topicuskeyhub/sdk-go.attributedef.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithAttributedefItemRequestBuilder when successful
func (m *AttributedefRequestBuilder) ByAttributedefid(attributedefid string)(*WithAttributedefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attributedefid != "" {
        urlTplParams["attributedefid"] = attributedefid
    }
    return NewWithAttributedefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAttributedefidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.attributedef.item collection
// returns a *WithAttributedefItemRequestBuilder when successful
func (m *AttributedefRequestBuilder) ByAttributedefidInt64(attributedefid int64)(*WithAttributedefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["attributedefid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(attributedefid, 10)
    return NewWithAttributedefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAttributedefRequestBuilderInternal instantiates a new AttributedefRequestBuilder and sets the default values.
func NewAttributedefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttributedefRequestBuilder) {
    m := &AttributedefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/attributedef{?additional*,any*,createdAfter*,createdBefore*,exclude*,format*,id*,modifiedSince*,noRuleForAccessProfile*,q*,sort*,systemDefinition*}", pathParameters),
    }
    return m
}
// NewAttributedefRequestBuilder instantiates a new AttributedefRequestBuilder and sets the default values.
func NewAttributedefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttributedefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttributedefRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all account attribute definitions in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a IdentityAccountAttributeDefinitionLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *AttributedefRequestBuilder) Get(ctx context.Context, requestConfiguration *AttributedefRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeDefinitionLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable), nil
}
// Post creates one or more new account attribute definitions and returns the newly created account attribute definitions.
// returns a IdentityAccountAttributeDefinitionLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *AttributedefRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable, requestConfiguration *AttributedefRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeDefinitionLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable), nil
}
// ToGetRequestInformation query for all account attribute definitions in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *AttributedefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttributedefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new account attribute definitions and returns the newly created account attribute definitions.
// returns a *RequestInformation when successful
func (m *AttributedefRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeDefinitionLinkableWrapperable, requestConfiguration *AttributedefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttributedefRequestBuilder when successful
func (m *AttributedefRequestBuilder) WithUrl(rawUrl string)(*AttributedefRequestBuilder) {
    return NewAttributedefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
