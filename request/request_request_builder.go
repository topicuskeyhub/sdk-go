package request

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// RequestRequestBuilder builds and executes requests for operations under \request
type RequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RequestRequestBuilderGetQueryParameters queries over all modification requests. The various query parameters can be used to filter the response.
type RequestRequestBuilderGetQueryParameters struct {
    // Only return UpdateGroupMembershipRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    AccountToUpdate []int64
    // Request additional information to be returned for every record.
    Additional []string
    // Only return RevokeAdminRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    AdminToRevoke []int64
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter requests on the given applications, specified by id. This parameter supports composition with all parameters from the client resource.
    Application []int64
    // Only return SetupAuthorizingGroupRequest for the given type of authorization.
    AuthorizationType []string
    // Only return SetupAuthorizingGroupRequest(s) for the given authorization type.
    AuthorizingGroupType []string
    // Only return GrantClientPermissionRequest for the given permission type(s).
    ClientPermission []string
    // Only return SetupAuthorizingGroupRequest(s) that either connect or disconnect additional authorization.
    ConnectAuthorization []bool
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return VerifyInternalAccountRequest for the given internal directory(ies), specified by id. This parameter supports composition with all parameters from the directory resource.
    Directory []int64
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Filter requests on the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Filter requests on the given group on system, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return VerifyInternalAccountRequest for the given internal account(s), specified by id. This parameter supports composition with all parameters from the internal account resource.
    InternalAccount []int64
    // Filter requests on the exact mail keys.
    MailKey []string
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return AddGroupAdminRequest for the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    NewAdmin []int64
    // Only return CreateGroupRequest with an exact match on the group name.
    NewGroupName []string
    // Only return CreateGroupOnSystemRequest with an exact match on the name in system.
    NewGroupOnSystemName []string
    // Only return CreateServiceAccountRequest with an exact match on the username.
    NewServiceAccountName []string
    // Filter requests on the given organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64
    // Only return requests that were preceded by the given request(s), specified by id. This parameter supports composition with all parameters from the request resource.
    PrecededBy []int64
    // Only return modification requests that can be processed by the given account, specified by id. This includes requests requested by the given account that are not yet processed or have been declined in the past 2 days.
    ProcessedBy []int64
    // Filter records on a complex CQL query.
    Q []string
    // Only return requests requested by the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    RequestedBy []int64
    // Only return requests that are requested for the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    RequestingGroup []int64
    // Filter requests on the given service accounts, specified by id. This parameter supports composition with all parameters from the service account resource.
    ServiceAccount []int64
    // Filter requests on the given status(es).
    Status []string
    // Filter requests on the given systems, specified by id. This parameter supports composition with all parameters from the system resource.
    System []int64
    // Only return requests that did or did not trigger a follow up request.
    TriggeredFollowUpRequest []bool
    // Only return requests of the given type(s).
    Type []string
    // Only return UpdateGroupMembershipRequest(s) of the given update type.
    UpdateGroupMembershipType []string
    // Only return requests the were processed by the given account(s), specified by id. This parameter supports composition with all parameters from the account resource.
    WasProcessedBy []int64
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
// RequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.request.item collection
func (m *RequestRequestBuilder) ById(id string)(*RequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRequestRequestBuilderInternal instantiates a new RequestRequestBuilder and sets the default values.
func NewRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RequestRequestBuilder) {
    m := &RequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/request{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,accountToUpdate*,adminToRevoke*,application*,authorizationType*,authorizingGroupType*,clientPermission*,connectAuthorization*,directory*,group*,groupOnSystem*,internalAccount*,mailKey*,newAdmin*,newGroupName*,newGroupOnSystemName*,newServiceAccountName*,organizationalUnit*,precededBy*,processedBy*,requestedBy*,requestingGroup*,serviceAccount*,status*,system*,triggeredFollowUpRequest*,type*,updateGroupMembershipType*,wasProcessedBy*}", pathParameters),
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
func (m *RequestRequestBuilder) Get(ctx context.Context, requestConfiguration *RequestRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *RequestRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, requestConfiguration *RequestRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// ToGetRequestInformation queries over all modification requests. The various query parameters can be used to filter the response.
func (m *RequestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RequestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new modification requests and returns the newly created requests.
func (m *RequestRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.RequestModificationRequestLinkableWrapperable, requestConfiguration *RequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
