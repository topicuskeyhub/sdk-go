package system

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemGroupRequestBuilder builds and executes requests for operations under \system\{id}\group
type ItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupRequestBuilderGetQueryParameters query for all groups on systems within a provisioned system in Topicus KeyHub. The various query parameters can be used to filter the response.
type ItemGroupRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
    // Filter the groups on system on groups that perform technical administration for the systems they belong to, specified by id. This parameter supports composition with all parameters from the group resource.
    AdminnedBy []int64
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter the groups on system on groups that perform content administration for the systems they belong to, specified by id. This parameter supports composition with all parameters from the group resource.
    ContentAdminnedBy []int64
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return groups on system that have a provisioning group with one of the given groups, specified by id. This parameter supports composition with all parameters from the groups resource.
    Group []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Search groups on (part of) the display name or the name in system.
    NameContains []string
    // Filter groups on system on the exact name in system.
    NameInSystem []string
    // Only return groups on system that do not have a provisioning group with one of the given groups, specified by id.
    NotLinkedToGroup []int64
    // Only return groups on system that do are not linked to the given service account, specified by id.
    NotLinkedToServiceAccount []int64
    // Filter the groups on system on groups that are owner for them, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64
    // Filter records on a complex CQL query.
    Q []string
    // Only return groups that are the on one of the given systems, specified by id. This parameter supports composition with all parameters from the systems resource.
    System []int64
    // Filter the groups on system on groups that are tier 2 owner for them, specified by id. A tier 2 owner is an owner of a group on system that is linked to a service account this group on system is also linked to.
    Tier2OwnedBy []int64
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
// ItemGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.system.item.group.item collection
func (m *ItemGroupRequestBuilder) ById(id string)(*ItemGroupGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewItemGroupGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    m := &ItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{id}/group{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,adminnedBy*,contentAdminnedBy*,group*,nameContains*,nameInSystem*,notLinkedToGroup*,notLinkedToServiceAccount*,ownedBy*,system*,tier2OwnedBy*}", pathParameters),
    }
    return m
}
// NewItemGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all groups on systems within a provisioned system in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningGroupOnSystemLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable), nil
}
// Post creates one or more new groups on systems and returns the newly created groups on systems. The groups will also be provisioned to the targeted provisioned system. By default, for every group on system, a provisioning group will be created granting the owner access to the newly created group. It is also possible to specify the provisioning groups to be created via the 'provgroups' addionalObjects property.
func (m *ItemGroupRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, requestConfiguration *ItemGroupRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningGroupOnSystemLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable), nil
}
// ToGetRequestInformation query for all groups on systems within a provisioned system in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new groups on systems and returns the newly created groups on systems. The groups will also be provisioned to the targeted provisioned system. By default, for every group on system, a provisioning group will be created granting the owner access to the newly created group. It is also possible to specify the provisioning groups to be created via the 'provgroups' addionalObjects property.
func (m *ItemGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, requestConfiguration *ItemGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
