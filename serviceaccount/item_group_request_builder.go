package serviceaccount

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i1f919641dec4c0f74b6fa0e15a8a8518fc70cef715c4c61e18f80a2cfac70831 "github.com/topicuskeyhub/sdk-go/serviceaccount/item/group"
)

// ItemGroupRequestBuilder builds and executes requests for operations under \serviceaccount\{serviceaccountid}\group
type ItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupRequestBuilderGetQueryParameters query for all groups for a service system in Topicus KeyHub. The various query parameters can be used to filter the response.
type ItemGroupRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i1f919641dec4c0f74b6fa0e15a8a8518fc70cef715c4c61e18f80a2cfac70831.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the service account groups on the groups on systems, specified by id. This parameter supports composition with all parameters from the groups on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter the service account groups on the service accounts, specified by id. This parameter supports composition with all parameters from the service account resource.
    ServiceAccount []int64 `uriparametername:"serviceAccount"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
}
// ItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGroupRequestBuilderGetQueryParameters
}
// ByGroupid gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item.group.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemGroupWithGroupItemRequestBuilder when successful
func (m *ItemGroupRequestBuilder) ByGroupid(groupid string)(*ItemGroupWithGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupid != "" {
        urlTplParams["groupid"] = groupid
    }
    return NewItemGroupWithGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByGroupidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item.group.item collection
// returns a *ItemGroupWithGroupItemRequestBuilder when successful
func (m *ItemGroupRequestBuilder) ByGroupidInt64(groupid int64)(*ItemGroupWithGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["groupid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(groupid, 10)
    return NewItemGroupWithGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGroupRequestBuilderInternal instantiates a new ItemGroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    m := &ItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount/{serviceaccountid}/group{?additional*,any*,createdAfter*,createdBefore*,exclude*,groupOnSystem*,id*,modifiedSince*,q*,serviceAccount*,sort*}", pathParameters),
    }
    return m
}
// NewItemGroupRequestBuilder instantiates a new ItemGroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all groups for a service system in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ServiceaccountServiceAccountGroupLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountGroupLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateServiceaccountServiceAccountGroupLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountGroupLinkableWrapperable), nil
}
// ToGetRequestInformation query for all groups for a service system in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=72")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGroupRequestBuilder when successful
func (m *ItemGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGroupRequestBuilder) {
    return NewItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
