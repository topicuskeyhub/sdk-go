package account

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// AccountRequestBuilder builds and executes requests for operations under \account
type AccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccountRequestBuilderGetQueryParameters query for all accounts in Topicus KeyHub. This will query over all directories and return both active and inactive accounts. The various query parameters can be used to filter the response.
type AccountRequestBuilderGetQueryParameters struct {
    // Only return accounts with the given activation code(s).
    ActivationCode []string
    // Only return accounts that have an activation deadline that expires at or after the given instant.
    ActivationDeadlineAtOrAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return active or inactive accounts.
    Active []bool
    // Only return accounts that have been active since the given instant.
    ActiveSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Request additional information to be returned for every record.
    Additional []string
    // Only return accounts that are or are not KeyHub Administrators.
    Admin []bool
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the accounts by directory, specified by id. This parameter supports composition with all parameters from the account directory resource.
    Directory []int64
    // Only return internal accounts for directories with the given owner(s), specified by id. This parameter supports composition with all parameters from the group resource.
    DirectoryOwnedBy []int64
    // Only return the accounts with the given e-mail addresses.
    Email []string
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Filter the accounts by membership of groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Only return accounts that have or don't have an e-mail address.
    HasEmail []bool
    // Only return accouts that have or don't have their vaults setup.
    HasVault []bool
    // Only return accounts that have access to the vault for at least one of the given groups, specified by id.
    HasVaultSecretForGroup []int64
    // Filter the results on the given ids.
    Id []int64
    // Filter accounts by exact match on the id in the directory.
    IdInDirectory []string
    // Only return accounts that have not been active since the given instant.
    InactiveSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return accounts with the given license role(s).
    LicenseRole []string
    // Only return the keyhub maintenance user.
    MaintenanceAdmin []bool
    // Only return managers for at least one of the given groups, specified by id.
    ManagerForGroup []int64
    // Filter accounts on direct connections to organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    MemberOfOrganizationalUnit []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return accounts for which the username does not start with one of the given values.
    NameDoesNotStartWith []string
    // Only return accounts for which the username starts with one of the given values.
    NameStartsWith []string
    // Filter the accounts that are not a member of the given groups, specified by id. 
    NotInGroup []int64
    // Filter records on a complex CQL query.
    Q []string
    // Only return account that do or do not require reregistration.
    ReregistrationRequired []bool
    // Only return accouts that have or don't have rotating password enabled.
    RotatingPasswordEnabled []bool
    // Only return accounts that match the given 2FA status. Note that {@code APP} only matches the accounts who only have a KeyHub app linked, and similar for {@code TOTP} and {@code WEBAUTHN}. {@code MIXED} matches those who have both one or more WebAuthn keys, and either a KeyHub app or generic Totp app linked.
    TwoFactorStatus []string
    // Filter accounts on the exact username.
    Username []string
    // Search accounts on (part of) the username, display name or uuid.
    UsernameContains []string
    // Filter results on one or more UUIDs.
    Uuid []string
    // Only return accounts that are valid or invalid in the directory.
    ValidInDirectory []bool
    // Filter accounts on their validity. NOT_APPLICABLE will be seen as VALID.
    Validity []string
    // Only return accounts that are the owner of one of the given vaults, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64
}
// AccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccountRequestBuilderGetQueryParameters
}
// Auditstats the auditstats property
func (m *AccountRequestBuilder) Auditstats()(*AuditstatsRequestBuilder) {
    return NewAuditstatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccountid gets an item from the github.com/topicuskeyhub/sdk-go.account.item collection
func (m *AccountRequestBuilder) ByAccountid(accountid string)(*WithAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accountid != "" {
        urlTplParams["accountid"] = accountid
    }
    return NewWithAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccountRequestBuilderInternal instantiates a new AccountRequestBuilder and sets the default values.
func NewAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountRequestBuilder) {
    m := &AccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,activationCode*,activationDeadlineAtOrAfter*,active*,activeSince*,admin*,directory*,directoryOwnedBy*,email*,group*,hasEmail*,hasVault*,hasVaultSecretForGroup*,idInDirectory*,inactiveSince*,licenseRole*,maintenanceAdmin*,managerForGroup*,memberOfOrganizationalUnit*,nameDoesNotStartWith*,nameStartsWith*,notInGroup*,reregistrationRequired*,rotatingPasswordEnabled*,twoFactorStatus*,username*,usernameContains*,uuid*,validInDirectory*,validity*,vault*}", pathParameters),
    }
    return m
}
// NewAccountRequestBuilder instantiates a new AccountRequestBuilder and sets the default values.
func NewAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *AccountRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all accounts in Topicus KeyHub. This will query over all directories and return both active and inactive accounts. The various query parameters can be used to filter the response.
func (m *AccountRequestBuilder) Get(ctx context.Context, requestConfiguration *AccountRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountLinkableWrapperable), nil
}
// Me the me property
func (m *AccountRequestBuilder) Me()(*MeRequestBuilder) {
    return NewMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Provisioning the provisioning property
func (m *AccountRequestBuilder) Provisioning()(*ProvisioningRequestBuilder) {
    return NewProvisioningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Status the status property
func (m *AccountRequestBuilder) Status()(*StatusRequestBuilder) {
    return NewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation query for all accounts in Topicus KeyHub. This will query over all directories and return both active and inactive accounts. The various query parameters can be used to filter the response.
func (m *AccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Vault the vault property
func (m *AccountRequestBuilder) Vault()(*VaultRequestBuilder) {
    return NewVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
