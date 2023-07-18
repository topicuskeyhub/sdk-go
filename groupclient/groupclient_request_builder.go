package groupclient

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// GroupclientRequestBuilder builds and executes requests for operations under \groupclient
type GroupclientRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupclientRequestBuilderGetQueryParameters queries over all client links for a group. The various query parameters can be used to filter the response.
type GroupclientRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter client group links on the groups performing technical administration of the client, specified by id. This parameter supports composition with all parameters from the group resource.
    AppAdminGroup []int64
    // Filter client group links on the groups with ownership of the client, specified by id. This parameter supports composition with all parameters from the group resource.
    AppOwnerGroup []int64
    // Filter client group links on the given clients, specified by id. This parameter supports composition with all parameters from the client resource.
    Client []int64
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Filter client group links on the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter records on a complex CQL query.
    Q []string
}
// GroupclientRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupclientRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupclientRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.groupclient.item collection
func (m *GroupclientRequestBuilder) ById(id string)(*GroupclientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewGroupclientItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupclientRequestBuilderInternal instantiates a new GroupclientRequestBuilder and sets the default values.
func NewGroupclientRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupclientRequestBuilder) {
    m := &GroupclientRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groupclient{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,appAdminGroup*,appOwnerGroup*,client*,group*}", pathParameters),
    }
    return m
}
// NewGroupclientRequestBuilder instantiates a new GroupclientRequestBuilder and sets the default values.
func NewGroupclientRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupclientRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupclientRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all client links for a group. The various query parameters can be used to filter the response.
func (m *GroupclientRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupclientRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupClientLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupClientLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupClientLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all client links for a group. The various query parameters can be used to filter the response.
func (m *GroupclientRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupclientRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
