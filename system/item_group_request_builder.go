package system

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    ie7728c2cafb9c86407f4f10f0899b38257a083c3d482e69f20e2744d87991d14 "github.com/topicuskeyhub/sdk-go/system/item/group"
)

// ItemGroupRequestBuilder builds and executes requests for operations under \system\{systemid}\group
type ItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupRequestBuilderGetQueryParameters query for all groups on systems within a provisioned system in Topicus KeyHub. The various query parameters can be used to filter the response.
type ItemGroupRequestBuilderGetQueryParameters struct {
    // Only return groups on system that have an access profile provisioning link with one of the given access profiles, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []ie7728c2cafb9c86407f4f10f0899b38257a083c3d482e69f20e2744d87991d14.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Filter the groups on system on groups that perform technical administration for the systems they belong to, specified by id. This parameter supports composition with all parameters from the group resource.
    AdminnedBy []int64 `uriparametername:"adminnedBy"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the groups on system on groups that perform content administration for the systems they belong to, specified by id. This parameter supports composition with all parameters from the group resource.
    ContentAdminnedBy []int64 `uriparametername:"contentAdminnedBy"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return groups on system that have a provisioning group with one of the given groups, specified by id. This parameter supports composition with all parameters from the groups resource.
    Group []int64 `uriparametername:"group"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Search groups on (part of) the display name or the name in system.
    NameContains []string `uriparametername:"nameContains"`
    // Filter groups on system on the exact name in system.
    NameInSystem []string `uriparametername:"nameInSystem"`
    // Only return groups on system that do not have a provisioning group with one of the given groups, specified by id.
    NotLinkedToGroup []int64 `uriparametername:"notLinkedToGroup"`
    // Only return groups on system that do are not linked to the given service account, specified by id.
    NotLinkedToServiceAccount []int64 `uriparametername:"notLinkedToServiceAccount"`
    // Filter groups-on-system on organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter groups-on-system on organizational units, specified by id.
    OrganizationalUnitForEnforcement []int64 `uriparametername:"organizationalUnitForEnforcement"`
    // Filter the groups on system on groups that are owner for them, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Only return groups that are the on one of the given systems, specified by id. This parameter supports composition with all parameters from the systems resource.
    System []int64 `uriparametername:"system"`
    // Constrain groups-on-system on the system the current user is content admin or owner of, specified by id.
    SystemForEnforcement []int64 `uriparametername:"systemForEnforcement"`
    // Filter the groups on system on groups that own the systems they belong to, specified by id. This parameter supports composition with all parameters from the group resource.
    SystemOwnedBy []int64 `uriparametername:"systemOwnedBy"`
    // Filter the groups on system on groups that are tier 2 owner for them, specified by id. A tier 2 owner is an owner of a group on system that is linked to a service account this group on system is also linked to.
    Tier2OwnedBy []int64 `uriparametername:"tier2OwnedBy"`
    // Filter groups on system on the type.
    // Deprecated: This property is deprecated, use TypeAsProvisioningGroupOnSystemType instead
    Type []string `uriparametername:"type"`
    // Filter groups on system on the type.
    TypeAsProvisioningGroupOnSystemType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemType `uriparametername:"type"`
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
// ItemGroupRequestBuilderPostQueryParameters creates one or more new groups on systems and returns the newly created groups on systems. The groups will also be provisioned to the targeted provisioned system. By default, for every group on system, a provisioning group will be created granting the owner access to the newly created group. It is also possible to specify the provisioning groups to be created via the 'provgroups' addionalObjects property.
type ItemGroupRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []ie7728c2cafb9c86407f4f10f0899b38257a083c3d482e69f20e2744d87991d14.PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGroupRequestBuilderPostQueryParameters
}
// ByGroupid gets an item from the github.com/topicuskeyhub/sdk-go.system.item.group.item collection
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
// ByGroupidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.system.item.group.item collection
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}/group{?accessProfile*,additional*,adminnedBy*,any*,contentAdminnedBy*,createdAfter*,createdBefore*,exclude*,group*,id*,modifiedSince*,nameContains*,nameInSystem*,notLinkedToGroup*,notLinkedToServiceAccount*,organizationalUnit*,organizationalUnitForEnforcement*,ownedBy*,q*,sort*,system*,systemForEnforcement*,systemOwnedBy*,tier2OwnedBy*,type*}", pathParameters),
    }
    return m
}
// NewItemGroupRequestBuilder instantiates a new ItemGroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all groups on systems within a provisioned system in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ProvisioningGroupOnSystemLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// returns a ProvisioningGroupOnSystemLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemGroupRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, requestConfiguration *ItemGroupRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new groups on systems and returns the newly created groups on systems. The groups will also be provisioned to the targeted provisioned system. By default, for every group on system, a provisioning group will be created granting the owner access to the newly created group. It is also possible to specify the provisioning groups to be created via the 'provgroups' addionalObjects property.
// returns a *RequestInformation when successful
func (m *ItemGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningGroupOnSystemLinkableWrapperable, requestConfiguration *ItemGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=78", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGroupRequestBuilder when successful
func (m *ItemGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGroupRequestBuilder) {
    return NewItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
