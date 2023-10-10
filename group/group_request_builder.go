package group

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// GroupRequestBuilder builds and executes requests for operations under \group
type GroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupRequestBuilderGetQueryParameters query for all groups in Topicus KeyHub. The various query parameters can be used to filter the response.
type GroupRequestBuilderGetQueryParameters struct {
    // Search groups by connected, owned or administrated clients, linked systems, webhooks or service accounts that have been modified since the given instance.
    AccessModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"accessModifiedSince"`
    // Search groups by name of connected, owned or administrated clients, linked systems, webhooks or service accounts.
    AccessQuicksearch []string `uriparametername:"accessQuicksearch"`
    // Request additional information to be returned for every record.
    Additional []string `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return groups that have or have not enabled application administration.
    ApplicationAdministration []bool `uriparametername:"applicationAdministration"`
    // Only return groups that are or are not due for a periodic audit.
    AuditDue []bool `uriparametername:"auditDue"`
    // Only return groups that have been audited since the given instant.
    AuditedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"auditedSince"`
    // Filter groups on the status of its periodic auditing. These sets are disjunct. A group always has only one status. A group that is over due for auditing will not be returned when only AUDITED is requested.
    AuditingStatus []string `uriparametername:"auditingStatus"`
    // Only returns the auditor group or all other groups.
    Auditor []bool `uriparametername:"auditor"`
    // Only return groups that are or are not over due for a periodic audit.
    AuditOverDue []bool `uriparametername:"auditOverDue"`
    // Only return groups for which an audit is or is not requested.
    AuditRequested []bool `uriparametername:"auditRequested"`
    // Filter groups for which audits are reviewed by one of the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    AuditsReviewedBy []int64 `uriparametername:"auditsReviewedBy"`
    // Filter groups for which the given groups perform some form of authorization, specified by id.
    AuthorizedBy []int64 `uriparametername:"authorizedBy"`
    // Filter groups by classification, specified by id. This parameter supports composition with all parameters from the group classification resource.
    Classification []int64 `uriparametername:"classification"`
    // Filter the groups by membership of accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    ContainsAccount []int64 `uriparametername:"containsAccount"`
    // Filter the groups by membership of all accounts, specified by id.
    ContainsAllAccounts []int64 `uriparametername:"containsAllAccounts"`
    // Filter the groups with clients connected, specified by id. This parameter supports composition with all parameters from the client resource.
    ContainsClient []int64 `uriparametername:"containsClient"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter groups that have delegated management to one of the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    DelegationGivenTo []int64 `uriparametername:"delegationGivenTo"`
    // Filter the groups by accounts not being member, specified by id.
    DoesNotContainAccount []int64 `uriparametername:"doesNotContainAccount"`
    // Filter the groups with clients not connected, specified by id.
    DoesNotContainClient []int64 `uriparametername:"doesNotContainClient"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter groups by auditing configuration.
    GroupAuditConfig []int64 `uriparametername:"groupAuditConfig"`
    // Only return groups that have, or do not have, any authorizing group set.
    HasAnyAuthorizingGroupSet []bool `uriparametername:"hasAnyAuthorizingGroupSet"`
    // Only return groups that have or do not have periodic auditing setup.
    HasAuditing []bool `uriparametername:"hasAuditing"`
    // Only return groups that have or do not have connected clients.
    HasClients []bool `uriparametername:"hasClients"`
    // Only return groups that have or do not have 2 or more managers.
    HasMoreThanOneManager []bool `uriparametername:"hasMoreThanOneManager"`
    // Only return groups that have or do not have dynamic provisioning of linked systems.
    HasSystems []bool `uriparametername:"hasSystems"`
    // Only return groups that have or do not have a vault.
    HasVault []bool `uriparametername:"hasVault"`
    // Only return groups that have or do not have webhooks on audit records.
    HasWebhooks []bool `uriparametername:"hasWebhooks"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return groups for which the given account is manager, specified by id.
    IsManager []int64 `uriparametername:"isManager"`
    // Only return the KeyHub administrator group or all other groups.
    KeyHubAdmin []bool `uriparametername:"keyHubAdmin"`
    // Only return groups that do or do not meet the various criteria set in the classification (if any).
    MeetsClassificationCriteria []string `uriparametername:"meetsClassificationCriteria"`
    // Filter groups for which membership is authorized by one of the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    MembershipAuthorizedBy []int64 `uriparametername:"membershipAuthorizedBy"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter my groups for the given accounts, specified by id. This contains the groups the account is member of or for which the account has delegated management privileges.
    MyGroups []int64 `uriparametername:"myGroups"`
    // Filter groups on the exact name.
    Name []string `uriparametername:"name"`
    // Filter groups on (part of) the name or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Filter groups with which the name does not start with the given values.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Filter groups on (part of) the name, description or uuid.
    NameOrDescriptionContains []string `uriparametername:"nameOrDescriptionContains"`
    // Filter groups on the start of the name.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Only return groups that are, or are not, nested under another group
    NestedGroup []bool `uriparametername:"nestedGroup"`
    // Only return groups that are nested under the given groups
    NestedUnder []int64 `uriparametername:"nestedUnder"`
    // Only return groups that have at least the given number of members.
    NumberOfAccountsGreaterOrEqual []int64 `uriparametername:"numberOfAccountsGreaterOrEqual"`
    // Filter groups on organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Only return groups that own or do not own clients.
    OwnsClients []bool `uriparametername:"ownsClients"`
    // Only return groups that own or do not own internal directories.
    OwnsDirectories []bool `uriparametername:"ownsDirectories"`
    // Only return groups that own or do not own groups on linked systems.
    OwnsGroupOnSystems []bool `uriparametername:"ownsGroupOnSystems"`
    // Only return groups that own or do not own provisioned systems.
    OwnsSystems []bool `uriparametername:"ownsSystems"`
    // Only return groups that are, or are not, marked as private group.
    PrivateGroup []bool `uriparametername:"privateGroup"`
    // Filter groups for which provisioning is authorized by one of the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    ProvisioningAuthorizedBy []int64 `uriparametername:"provisioningAuthorizedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Only return groups that do, or do not, require a rotating password to activate.
    RotatingPasswordRequired []bool `uriparametername:"rotatingPasswordRequired"`
    // Only return groups that allow or do not allow management by a single manager.
    SingleManaged []bool `uriparametername:"singleManaged"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
    // Filter the groups by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64 `uriparametername:"vault"`
    // Only return groups with the given vault recovery availability.
    VaultRecovery []string `uriparametername:"vaultRecovery"`
    // Filter groups on whether they are public or only visible to members. Possible values are ALL ("return all groups (public and hidden)"), PUBLIC ("return only public groups") and PUBLIC_OR_MEMBER ("return public groups and hidden groups of which the current subject is a member"). Default is PUBLIC_OR_MEMBER. This option does not influence permissions, i.e. filtering on ALL can lead to errors if the current subject does not have the correct permissions on one or more groups.
    Visibility []string `uriparametername:"visibility"`
}
// GroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupRequestBuilderGetQueryParameters
}
// GroupRequestBuilderPostQueryParameters creates one or more new groups and returns the newly created groups. It is required to specify the first admin via the admins additional object.
type GroupRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string `uriparametername:"additional"`
}
// GroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupRequestBuilderPostQueryParameters
}
// Account the account property
func (m *GroupRequestBuilder) Account()(*AccountRequestBuilder) {
    return NewAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Auditstats the auditstats property
func (m *GroupRequestBuilder) Auditstats()(*AuditstatsRequestBuilder) {
    return NewAuditstatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByGroupid gets an item from the github.com/topicuskeyhub/sdk-go.group.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *GroupRequestBuilder) ByGroupid(groupid string)(*WithGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupid != "" {
        urlTplParams["groupid"] = groupid
    }
    return NewWithGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByGroupidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.group.item collection
func (m *GroupRequestBuilder) ByGroupidInt64(groupid int64)(*WithGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["groupid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(groupid, 10)
    return NewWithGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,accessModifiedSince*,accessQuicksearch*,applicationAdministration*,auditDue*,auditOverDue*,auditRequested*,auditedSince*,auditingStatus*,auditor*,auditsReviewedBy*,authorizedBy*,classification*,containsAccount*,containsAllAccounts*,containsClient*,delegationGivenTo*,doesNotContainAccount*,doesNotContainClient*,groupAuditConfig*,hasAnyAuthorizingGroupSet*,hasAuditing*,hasClients*,hasMoreThanOneManager*,hasSystems*,hasVault*,hasWebhooks*,isManager*,keyHubAdmin*,meetsClassificationCriteria*,membershipAuthorizedBy*,myGroups*,name*,nameContains*,nameDoesNotStartWith*,nameOrDescriptionContains*,nameStartsWith*,nestedGroup*,nestedUnder*,numberOfAccountsGreaterOrEqual*,organizationalUnit*,ownsClients*,ownsDirectories*,ownsGroupOnSystems*,ownsSystems*,privateGroup*,provisioningAuthorizedBy*,rotatingPasswordRequired*,singleManaged*,uuid*,vault*,vaultRecovery*,visibility*}", pathParameters),
    }
    return m
}
// NewGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *GroupRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all groups in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *GroupRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable), nil
}
// Post creates one or more new groups and returns the newly created groups. It is required to specify the first admin via the admins additional object.
func (m *GroupRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable, requestConfiguration *GroupRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable), nil
}
// Segments the segments property
func (m *GroupRequestBuilder) Segments()(*SegmentsRequestBuilder) {
    return NewSegmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation query for all groups in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *GroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=67")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new groups and returns the newly created groups. It is required to specify the first admin via the admins additional object.
func (m *GroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupLinkableWrapperable, requestConfiguration *GroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=67")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=67", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GroupRequestBuilder) WithUrl(rawUrl string)(*GroupRequestBuilder) {
    return NewGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
