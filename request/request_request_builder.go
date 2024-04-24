package request

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// RequestRequestBuilder builds and executes requests for operations under \request
type RequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RequestRequestBuilderGetQueryParameters queries over all modification requests. The various query parameters can be used to filter the response.
type RequestRequestBuilderGetQueryParameters struct {
    // Filter requests on the given access profiles, specified by id. This parameter supports composition with all parameters from the access profile unit resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Only return UpdateGroupMembershipRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    AccountToUpdate []int64 `uriparametername:"accountToUpdate"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Only return RevokeAdminRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    AdminToRevoke []int64 `uriparametername:"adminToRevoke"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter requests on the given applications, specified by id. This parameter supports composition with all parameters from the client resource.
    Application []int64 `uriparametername:"application"`
    // Only return SetupAuthorizingGroupRequest for the given type of authorization.
    // Deprecated: This property is deprecated, use AuthorizationTypeAsRequestAuthorizingGroupType instead
    AuthorizationType []string `uriparametername:"authorizationType"`
    // Only return SetupAuthorizingGroupRequest for the given type of authorization.
    AuthorizationTypeAsRequestAuthorizingGroupType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestAuthorizingGroupType `uriparametername:"authorizationType"`
    // Only return SetupAuthorizingGroupRequest(s) for the given authorization type.
    // Deprecated: This property is deprecated, use AuthorizingGroupTypeAsRequestAuthorizingGroupType instead
    AuthorizingGroupType []string `uriparametername:"authorizingGroupType"`
    // Only return SetupAuthorizingGroupRequest(s) for the given authorization type.
    AuthorizingGroupTypeAsRequestAuthorizingGroupType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestAuthorizingGroupType `uriparametername:"authorizingGroupType"`
    // Only return GrantClientPermissionRequest for the given permission type(s).
    // Deprecated: This property is deprecated, use ClientPermissionAsClientOAuth2ClientPermissionType instead
    ClientPermission []string `uriparametername:"clientPermission"`
    // Only return GrantClientPermissionRequest for the given permission type(s).
    ClientPermissionAsClientOAuth2ClientPermissionType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionType `uriparametername:"clientPermission"`
    // Only return SetupAuthorizingGroupRequest(s) that either connect or disconnect additional authorization.
    ConnectAuthorization []bool `uriparametername:"connectAuthorization"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Only return VerifyInternalAccountRequest for the given internal directory(ies), specified by id. This parameter supports composition with all parameters from the directory resource.
    Directory []int64 `uriparametername:"directory"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter requests on the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Filter requests on the given group on system, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Filter requests for multiple groups if they contain any of the given groups.
    GroupsOverlapWith []int64 `uriparametername:"groupsOverlapWith"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return VerifyInternalAccountRequest for the given internal account(s), specified by id. This parameter supports composition with all parameters from the internal account resource.
    InternalAccount []int64 `uriparametername:"internalAccount"`
    // Filter requests on the exact mail keys.
    MailKey []string `uriparametername:"mailKey"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Only return AddGroupAdminRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    NewAdmin []int64 `uriparametername:"newAdmin"`
    // Only return CreateGroupRequest with an exact match on the group name.
    NewGroupName []string `uriparametername:"newGroupName"`
    // Only return CreateGroupOnSystemRequest with an exact match on the name in system.
    NewGroupOnSystemName []string `uriparametername:"newGroupOnSystemName"`
    // Only return CreateProvisionedNamespaceRequest with an exact match on the namespace name.
    NewNamespaceName []string `uriparametername:"newNamespaceName"`
    // Only return CreateServiceAccountRequest with an exact match on the username.
    NewServiceAccountName []string `uriparametername:"newServiceAccountName"`
    // Filter requests on the given organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Only return CreateGroupOnSystemRequest with the given owner group(s), specified by id. This parameter supports composition with all parameters from the grouop resource.
    OwnerGroup []int64 `uriparametername:"ownerGroup"`
    // Only return requests that were preceded by the given request(s), specified by id. This parameter supports composition with all parameters from the request resource.
    PrecededBy []int64 `uriparametername:"precededBy"`
    // Only return modification requests that can be processed by the given account, specified by id. This includes requests requested by the given account that are not yet processed or have been declined in the past 2 days.
    ProcessedBy []int64 `uriparametername:"processedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Only return requests requested by the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    RequestedBy []int64 `uriparametername:"requestedBy"`
    // Only return requests that are requested for the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    RequestingGroup []int64 `uriparametername:"requestingGroup"`
    // Filter requests on the given service accounts, specified by id. This parameter supports composition with all parameters from the service account resource.
    ServiceAccount []int64 `uriparametername:"serviceAccount"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter requests on the given status(es).
    // Deprecated: This property is deprecated, use StatusAsRequestModificationRequestStatus instead
    Status []string `uriparametername:"status"`
    // Filter requests on the given status(es).
    StatusAsRequestModificationRequestStatus []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestStatus `uriparametername:"status"`
    // Filter requests on the given systems, specified by id. This parameter supports composition with all parameters from the system resource.
    System []int64 `uriparametername:"system"`
    // Only return requests that did or did not trigger a follow up request.
    TriggeredFollowUpRequest []bool `uriparametername:"triggeredFollowUpRequest"`
    // Only return requests of the given type(s).
    Type []string `uriparametername:"type"`
    // Only return UpdateGroupMembershipRequest(s) of the given update type.
    // Deprecated: This property is deprecated, use UpdateGroupMembershipTypeAsRequestUpdateGroupMembershipType instead
    UpdateGroupMembershipType []string `uriparametername:"updateGroupMembershipType"`
    // Only return UpdateGroupMembershipRequest(s) of the given update type.
    UpdateGroupMembershipTypeAsRequestUpdateGroupMembershipType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestUpdateGroupMembershipType `uriparametername:"updateGroupMembershipType"`
    // Only return requests the were processed by the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    WasProcessedBy []int64 `uriparametername:"wasProcessedBy"`
}
// RequestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RequestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RequestRequestBuilderGetQueryParameters
}
// RequestRequestBuilderPostQueryParameters creates one or more new modification requests and returns the newly created requests.
type RequestRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// RequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RequestRequestBuilderPostQueryParameters
}
// ByRequestid gets an item from the github.com/topicuskeyhub/sdk-go.request.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithRequestItemRequestBuilder when successful
func (m *RequestRequestBuilder) ByRequestid(requestid string)(*WithRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if requestid != "" {
        urlTplParams["requestid"] = requestid
    }
    return NewWithRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRequestidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.request.item collection
// returns a *WithRequestItemRequestBuilder when successful
func (m *RequestRequestBuilder) ByRequestidInt64(requestid int64)(*WithRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["requestid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(requestid, 10)
    return NewWithRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRequestRequestBuilderInternal instantiates a new RequestRequestBuilder and sets the default values.
func NewRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RequestRequestBuilder) {
    m := &RequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/request{?accessProfile*,accountToUpdate*,additional*,adminToRevoke*,any*,application*,authorizationType*,authorizingGroupType*,clientPermission*,connectAuthorization*,createdAfter*,createdBefore*,directory*,exclude*,group*,groupOnSystem*,groupsOverlapWith*,id*,internalAccount*,mailKey*,modifiedSince*,newAdmin*,newGroupName*,newGroupOnSystemName*,newNamespaceName*,newServiceAccountName*,organizationalUnit*,ownerGroup*,precededBy*,processedBy*,q*,requestedBy*,requestingGroup*,serviceAccount*,sort*,status*,system*,triggeredFollowUpRequest*,type*,updateGroupMembershipType*,wasProcessedBy*}", pathParameters),
    }
    return m
}
// NewRequestRequestBuilder instantiates a new RequestRequestBuilder and sets the default values.
func NewRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all modification requests. The various query parameters can be used to filter the response.
// returns a RequestModificationRequestLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *RequestRequestBuilder) Get(ctx context.Context, requestConfiguration *RequestRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateRequestModificationRequestLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable), nil
}
// Post creates one or more new modification requests and returns the newly created requests.
// returns a RequestModificationRequestLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *RequestRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, requestConfiguration *RequestRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateRequestModificationRequestLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable), nil
}
// Report the report property
// returns a *ReportRequestBuilder when successful
func (m *RequestRequestBuilder) Report()(*ReportRequestBuilder) {
    return NewReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation queries over all modification requests. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *RequestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RequestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new modification requests and returns the newly created requests.
// returns a *RequestInformation when successful
func (m *RequestRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, requestConfiguration *RequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/request{?additional*}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=70", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RequestRequestBuilder when successful
func (m *RequestRequestBuilder) WithUrl(rawUrl string)(*RequestRequestBuilder) {
    return NewRequestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
