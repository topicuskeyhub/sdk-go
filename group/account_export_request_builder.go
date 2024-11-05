package group

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// AccountExportRequestBuilder builds and executes requests for operations under \group\account\export
type AccountExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccountExportRequestBuilderPostQueryParameters prepares an export of groupaccounts using the filtering on groups specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
type AccountExportRequestBuilderPostQueryParameters struct {
    // Search groups by connected, owned or administrated clients, linked systems, webhooks or service accounts that have been modified since the given instance.
    AccessModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"accessModifiedSince"`
    // Only return groups that have or have not enabled access profile administration.
    AccessProfileAdministration []bool `uriparametername:"accessProfileAdministration"`
    // Search groups by name of connected, owned or administrated clients, linked systems, webhooks or service accounts.
    AccessQuicksearch []string `uriparametername:"accessQuicksearch"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return groups that have or have not enabled application administration.
    ApplicationAdministration []bool `uriparametername:"applicationAdministration"`
    // Only return groups that are or are not due for a periodic audit.
    AuditDue []bool `uriparametername:"auditDue"`
    // Only return groups that have been audited since the given instant.
    AuditedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"auditedSince"`
    // Filter groups on the status of its periodic auditing. These sets are disjunct. A group always has only one status. A group that is over due for auditing will not be returned when only AUDITED is requested.
    // Deprecated: This property is deprecated, use AuditingStatusAsGroupGroupAuditingStatus instead
    AuditingStatus []string `uriparametername:"auditingStatus"`
    // Filter groups on the status of its periodic auditing. These sets are disjunct. A group always has only one status. A group that is over due for auditing will not be returned when only AUDITED is requested.
    AuditingStatusAsGroupGroupAuditingStatus []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAuditingStatus `uriparametername:"auditingStatus"`
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
    // Only return groups that have or do not have an audit that is currently under review.
    HasAuditUnderReview []bool `uriparametername:"hasAuditUnderReview"`
    // Only return groups that have or do not have given permissions to OAuth2 clients.
    HasClientPermissions []bool `uriparametername:"hasClientPermissions"`
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
    // Deprecated: This property is deprecated, use MeetsClassificationCriteriaAsGroupGroupClassificationConformance instead
    MeetsClassificationCriteria []string `uriparametername:"meetsClassificationCriteria"`
    // Only return groups that do or do not meet the various criteria set in the classification (if any).
    MeetsClassificationCriteriaAsGroupGroupClassificationConformance []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupClassificationConformance `uriparametername:"meetsClassificationCriteria"`
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
    // Only return groups that are, or are not, nested under another group.
    NestedGroup []bool `uriparametername:"nestedGroup"`
    // Only return groups that are nested under the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    NestedUnder []int64 `uriparametername:"nestedUnder"`
    // Only return groups that are not nested under the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    NotNestedUnder []int64 `uriparametername:"notNestedUnder"`
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
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
    // Filter the groups by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64 `uriparametername:"vault"`
    // Only return groups with the given vault recovery availability.
    // Deprecated: This property is deprecated, use VaultRecoveryAsGroupVaultRecoveryAvailability instead
    VaultRecovery []string `uriparametername:"vaultRecovery"`
    // Only return groups with the given vault recovery availability.
    VaultRecoveryAsGroupVaultRecoveryAvailability []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupVaultRecoveryAvailability `uriparametername:"vaultRecovery"`
    // Only return groups for which activation is or is not required when accessing the vault.
    VaultRequiresActivation []bool `uriparametername:"vaultRequiresActivation"`
    // Filter groups on whether they are public or only visible to members. Possible values are ALL ("return all groups (public and hidden)"), PUBLIC ("return only public groups") and PUBLIC_OR_MEMBER ("return public groups and hidden groups of which the current subject is a member"). Default is PUBLIC_OR_MEMBER. This option does not influence permissions, i.e. filtering on ALL can lead to errors if the current subject does not have the correct permissions on one or more groups.
    // Deprecated: This property is deprecated, use VisibilityAsGroupGroupVisibility instead
    Visibility []string `uriparametername:"visibility"`
    // Filter groups on whether they are public or only visible to members. Possible values are ALL ("return all groups (public and hidden)"), PUBLIC ("return only public groups") and PUBLIC_OR_MEMBER ("return public groups and hidden groups of which the current subject is a member"). Default is PUBLIC_OR_MEMBER. This option does not influence permissions, i.e. filtering on ALL can lead to errors if the current subject does not have the correct permissions on one or more groups.
    VisibilityAsGroupGroupVisibility []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupVisibility `uriparametername:"visibility"`
}
// AccountExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccountExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccountExportRequestBuilderPostQueryParameters
}
// NewAccountExportRequestBuilderInternal instantiates a new AccountExportRequestBuilder and sets the default values.
func NewAccountExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountExportRequestBuilder) {
    m := &AccountExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/account/export{?accessModifiedSince*,accessProfileAdministration*,accessQuicksearch*,any*,applicationAdministration*,auditDue*,auditOverDue*,auditRequested*,auditedSince*,auditingStatus*,auditsReviewedBy*,authorizedBy*,classification*,containsAccount*,containsAllAccounts*,containsClient*,createdAfter*,createdBefore*,delegationGivenTo*,doesNotContainAccount*,doesNotContainClient*,exclude*,groupAuditConfig*,hasAnyAuthorizingGroupSet*,hasAuditUnderReview*,hasAuditing*,hasClientPermissions*,hasClients*,hasMoreThanOneManager*,hasSystems*,hasVault*,hasWebhooks*,id*,isManager*,keyHubAdmin*,meetsClassificationCriteria*,membershipAuthorizedBy*,modifiedSince*,myGroups*,name*,nameContains*,nameDoesNotStartWith*,nameOrDescriptionContains*,nameStartsWith*,nestedGroup*,nestedUnder*,notNestedUnder*,numberOfAccountsGreaterOrEqual*,organizationalUnit*,ownsClients*,ownsDirectories*,ownsGroupOnSystems*,ownsSystems*,privateGroup*,provisioningAuthorizedBy*,q*,rotatingPasswordRequired*,singleManaged*,sort*,uuid*,vault*,vaultRecovery*,vaultRequiresActivation*,visibility*}", pathParameters),
    }
    return m
}
// NewAccountExportRequestBuilder instantiates a new AccountExportRequestBuilder and sets the default values.
func NewAccountExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccountExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post prepares an export of groupaccounts using the filtering on groups specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *AccountExportRequestBuilder) Post(ctx context.Context, requestConfiguration *AccountExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation prepares an export of groupaccounts using the filtering on groups specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a *RequestInformation when successful
func (m *AccountExportRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccountExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccountExportRequestBuilder when successful
func (m *AccountExportRequestBuilder) WithUrl(rawUrl string)(*AccountExportRequestBuilder) {
    return NewAccountExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
