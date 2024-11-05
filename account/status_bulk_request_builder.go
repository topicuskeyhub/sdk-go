package account

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// StatusBulkRequestBuilder builds and executes requests for operations under \account\status\bulk
type StatusBulkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StatusBulkRequestBuilderPutQueryParameters updates multiple accounts. The given accounts will be updated according to the selected action, for example, have their license role set to the given license role, provided this matches the current license's conditions.
type StatusBulkRequestBuilderPutQueryParameters struct {
    // Filter the accounts by membership of access profiles, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Only return accounts with the given activation code(s).
    ActivationCode []string `uriparametername:"activationCode"`
    // Only return accounts that have an activation deadline that expires at or after the given instant.
    ActivationDeadlineAtOrAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"activationDeadlineAtOrAfter"`
    // Only return active or inactive accounts.
    Active []bool `uriparametername:"active"`
    // Only return accounts that have been active since the given instant.
    ActiveSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"activeSince"`
    // Only return accounts that are or are not KeyHub Administrators.
    Admin []bool `uriparametername:"admin"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return the accounts for which the base organizational unit of the directory is equal to or an ancestor of the given unit(s), specified by id.
    BaseOrganizationalUnitAncestorOf []int64 `uriparametername:"baseOrganizationalUnitAncestorOf"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the accounts by directory, specified by id. This parameter supports composition with all parameters from the account directory resource.
    Directory []int64 `uriparametername:"directory"`
    // Only return internal accounts for directories with the given owner(s), specified by id. This parameter supports composition with all parameters from the group resource.
    DirectoryOwnedBy []int64 `uriparametername:"directoryOwnedBy"`
    // Only return the accounts with the given e-mail addresses.
    Email []string `uriparametername:"email"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the accounts by their registered external ids. This parameter supports composition with all parameters from the 'external id' resource.
    ExternalId []int64 `uriparametername:"externalId"`
    // Filter the accounts by membership of groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Only return accounts that have or don't have an e-mail address.
    HasEmail []bool `uriparametername:"hasEmail"`
    // Only return accouts that have or don't have their vaults setup.
    HasVault []bool `uriparametername:"hasVault"`
    // Only return accounts that have access to the vault for at least one of the given groups, specified by id.
    HasVaultSecretForGroup []int64 `uriparametername:"hasVaultSecretForGroup"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Filter accounts by exact match on the id in the directory.
    IdInDirectory []string `uriparametername:"idInDirectory"`
    // Only return accounts that have not been active since the given instant.
    InactiveSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"inactiveSince"`
    // Only return accounts with the given license role(s).
    // Deprecated: This property is deprecated, use LicenseRoleAsAuthAccountLicenseRole instead
    LicenseRole []string `uriparametername:"licenseRole"`
    // Only return accounts with the given license role(s).
    LicenseRoleAsAuthAccountLicenseRole []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountLicenseRole `uriparametername:"licenseRole"`
    // Only return the keyhub maintenance user.
    MaintenanceAdmin []bool `uriparametername:"maintenanceAdmin"`
    // Only return managers for at least one of the given groups, specified by id.
    ManagerForGroup []int64 `uriparametername:"managerForGroup"`
    // Filter accounts on direct connections to organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    MemberOfOrganizationalUnit []int64 `uriparametername:"memberOfOrganizationalUnit"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Only return accounts for which the username does not start with one of the given values.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Only return accounts for which the username starts with one of the given values.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Filter the accounts that are not a member of the given access profiles, specified by id. 
    NotInAccessProfile []int64 `uriparametername:"notInAccessProfile"`
    // Filter the accounts that are not a member of the given groups, specified by id. 
    NotInGroup []int64 `uriparametername:"notInGroup"`
    // Filter the accounts on their 'owning' clients, specified by id. An 'owning client' is a client which can potentially modify the account. Typically this is the client which created the account (if applicable), but it could be a client which has been registered as such after the fact.
    OwningClient []int64 `uriparametername:"owningClient"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Only return account that do or do not require reregistration.
    ReregistrationRequired []bool `uriparametername:"reregistrationRequired"`
    // Only return accouts that have or don't have rotating password enabled.
    RotatingPasswordEnabled []bool `uriparametername:"rotatingPasswordEnabled"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Only return accounts that match the given 2FA status. Note that {@code APP} only matches the accounts who only have a KeyHub app linked, and similar for {@code TOTP} and {@code WEBAUTHN}. {@code MIXED} matches those who have both one or more WebAuthn keys, and either a KeyHub app or generic Totp app linked.
    // Deprecated: This property is deprecated, use TwoFactorStatusAsAuthTwoFactorAuthenticationStatus instead
    TwoFactorStatus []string `uriparametername:"twoFactorStatus"`
    // Only return accounts that match the given 2FA status. Note that {@code APP} only matches the accounts who only have a KeyHub app linked, and similar for {@code TOTP} and {@code WEBAUTHN}. {@code MIXED} matches those who have both one or more WebAuthn keys, and either a KeyHub app or generic Totp app linked.
    TwoFactorStatusAsAuthTwoFactorAuthenticationStatus []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthTwoFactorAuthenticationStatus `uriparametername:"twoFactorStatus"`
    // Filter accounts on the exact username.
    Username []string `uriparametername:"username"`
    // Search accounts on (part of) the username, display name or uuid.
    UsernameContains []string `uriparametername:"usernameContains"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
    // Only return accounts that are valid or invalid in the directory.
    ValidInDirectory []bool `uriparametername:"validInDirectory"`
    // Filter accounts on their validity. NOT_APPLICABLE will be seen as VALID.
    // Deprecated: This property is deprecated, use ValidityAsAuthAccountValidity instead
    Validity []string `uriparametername:"validity"`
    // Filter accounts on their validity. NOT_APPLICABLE will be seen as VALID.
    ValidityAsAuthAccountValidity []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountValidity `uriparametername:"validity"`
    // Only return accounts that are the owner of one of the given vaults, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64 `uriparametername:"vault"`
}
// StatusBulkRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StatusBulkRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *StatusBulkRequestBuilderPutQueryParameters
}
// NewStatusBulkRequestBuilderInternal instantiates a new StatusBulkRequestBuilder and sets the default values.
func NewStatusBulkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusBulkRequestBuilder) {
    m := &StatusBulkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/status/bulk{?accessProfile*,activationCode*,activationDeadlineAtOrAfter*,active*,activeSince*,admin*,any*,baseOrganizationalUnitAncestorOf*,createdAfter*,createdBefore*,directory*,directoryOwnedBy*,email*,exclude*,externalId*,group*,hasEmail*,hasVault*,hasVaultSecretForGroup*,id*,idInDirectory*,inactiveSince*,licenseRole*,maintenanceAdmin*,managerForGroup*,memberOfOrganizationalUnit*,modifiedSince*,nameDoesNotStartWith*,nameStartsWith*,notInAccessProfile*,notInGroup*,owningClient*,q*,reregistrationRequired*,rotatingPasswordEnabled*,sort*,twoFactorStatus*,username*,usernameContains*,uuid*,validInDirectory*,validity*,vault*}", pathParameters),
    }
    return m
}
// NewStatusBulkRequestBuilder instantiates a new StatusBulkRequestBuilder and sets the default values.
func NewStatusBulkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusBulkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatusBulkRequestBuilderInternal(urlParams, requestAdapter)
}
// Put updates multiple accounts. The given accounts will be updated according to the selected action, for example, have their license role set to the given license role, provided this matches the current license's conditions.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *StatusBulkRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountBulkUpdateable, requestConfiguration *StatusBulkRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToPutRequestInformation updates multiple accounts. The given accounts will be updated according to the selected action, for example, have their license role set to the given license role, provided this matches the current license's conditions.
// returns a *RequestInformation when successful
func (m *StatusBulkRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountBulkUpdateable, requestConfiguration *StatusBulkRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=74", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StatusBulkRequestBuilder when successful
func (m *StatusBulkRequestBuilder) WithUrl(rawUrl string)(*StatusBulkRequestBuilder) {
    return NewStatusBulkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
