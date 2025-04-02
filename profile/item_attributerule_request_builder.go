package profile

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i783381d01b16caed4b91a8b77c9298c8551b8f016070080557d8a27462668795 "github.com/topicuskeyhub/sdk-go/profile/item/attributerule"
)

// ItemAttributeruleRequestBuilder builds and executes requests for operations under \profile\{accessprofile-id}\attributerule
type ItemAttributeruleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAttributeruleRequestBuilderGetQueryParameters query for all account attribute rules in Topicus KeyHub. The various query parameters can be used to filter the response.
type ItemAttributeruleRequestBuilderGetQueryParameters struct {
    // Filter account attribute rules on the access profile, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Filter account attribute rules on the owning group of the access profile, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfileOwningGroup []int64 `uriparametername:"accessProfileOwningGroup"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i783381d01b16caed4b91a8b77c9298c8551b8f016070080557d8a27462668795.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the attribute rules by attribute definition, specified by id. This parameter supports composition with all parameters from the attribute definition resource.
    Attribute []int64 `uriparametername:"attribute"`
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
    // Filter account attribute rules on organizational units of the owning group of the access profile, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
}
// ItemAttributeruleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributeruleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributeruleRequestBuilderGetQueryParameters
}
// ItemAttributeruleRequestBuilderPostQueryParameters creates one or more new account attribute rules and returns the newly created account attribute rules.
type ItemAttributeruleRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []i783381d01b16caed4b91a8b77c9298c8551b8f016070080557d8a27462668795.PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAttributeruleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAttributeruleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAttributeruleRequestBuilderPostQueryParameters
}
// ByAttributeruleid gets an item from the github.com/topicuskeyhub/sdk-go.profile.item.attributerule.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemAttributeruleWithAttributeruleItemRequestBuilder when successful
func (m *ItemAttributeruleRequestBuilder) ByAttributeruleid(attributeruleid string)(*ItemAttributeruleWithAttributeruleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attributeruleid != "" {
        urlTplParams["attributeruleid"] = attributeruleid
    }
    return NewItemAttributeruleWithAttributeruleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAttributeruleidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.profile.item.attributerule.item collection
// returns a *ItemAttributeruleWithAttributeruleItemRequestBuilder when successful
func (m *ItemAttributeruleRequestBuilder) ByAttributeruleidInt64(attributeruleid int64)(*ItemAttributeruleWithAttributeruleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["attributeruleid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(attributeruleid, 10)
    return NewItemAttributeruleWithAttributeruleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAttributeruleRequestBuilderInternal instantiates a new ItemAttributeruleRequestBuilder and sets the default values.
func NewItemAttributeruleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributeruleRequestBuilder) {
    m := &ItemAttributeruleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profile/{accessprofile%2Did}/attributerule{?accessProfile*,accessProfileOwningGroup*,additional*,any*,attribute*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,organizationalUnit*,q*,sort*}", pathParameters),
    }
    return m
}
// NewItemAttributeruleRequestBuilder instantiates a new ItemAttributeruleRequestBuilder and sets the default values.
func NewItemAttributeruleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributeruleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributeruleRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all account attribute rules in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a IdentityAccountAttributeRuleLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributeruleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAttributeruleRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeRuleLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable), nil
}
// Post creates one or more new account attribute rules and returns the newly created account attribute rules.
// returns a IdentityAccountAttributeRuleLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAttributeruleRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable, requestConfiguration *ItemAttributeruleRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateIdentityAccountAttributeRuleLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable), nil
}
// ToGetRequestInformation query for all account attribute rules in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemAttributeruleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAttributeruleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new account attribute rules and returns the newly created account attribute rules.
// returns a *RequestInformation when successful
func (m *ItemAttributeruleRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.IdentityAccountAttributeRuleLinkableWrapperable, requestConfiguration *ItemAttributeruleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAttributeruleRequestBuilder when successful
func (m *ItemAttributeruleRequestBuilder) WithUrl(rawUrl string)(*ItemAttributeruleRequestBuilder) {
    return NewItemAttributeruleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
