package system

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// SystemRequestBuilder builds and executes requests for operations under \system
type SystemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SystemRequestBuilderGetQueryParameters query for all provisioned systems in Topicus KeyHub. The various query parameters can be used to filter the response.
type SystemRequestBuilderGetQueryParameters struct {
    // Filter the systems by provisioned accounts, specified by id. This parameter supports composition with all parameters from the provisioned account resource.
    Account []int64 `uriparametername:"account"`
    // Only return active or inactive systems.
    // Deprecated: This property is deprecated, use ActiveAsBooleanEnum instead
    Active []string `uriparametername:"active"`
    // Only return active or inactive systems.
    ActiveAsBooleanEnum []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.BooleanEnum `uriparametername:"active"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the systems on groups that perform content administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    ContentAdministrator []int64 `uriparametername:"contentAdministrator"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return systems for which one of the certificates used is expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"expiredCertificate"`
    // Filter the systems on one or more UUIDs as read from the system.
    ExternalUuid []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID `uriparametername:"externalUuid"`
    // Filter systems on the Azure OIDC directory for which it is a provisioned source directory, specified by id. This parameter supports composition with all parameters from the directory resource.
    ForAzureOIDCDirectory []int64 `uriparametername:"forAzureOIDCDirectory"`
    // Filter provisioned namespaces on their base system, specified by id. This parameter supports composition with all parameters from the provisioned system resource.
    ForBaseSystem []int64 `uriparametername:"forBaseSystem"`
    // Filter systems on the clients for which it is an internal provisioned LDAP, specified by id. This parameter supports composition with all parameters from the client resource.
    ForClient []int64 `uriparametername:"forClient"`
    // Filter systems on the LDAP directory for which it is a provisioned source directory, specified by id. This parameter supports composition with all parameters from the directory resource.
    ForLDAPDirectory []int64 `uriparametername:"forLDAPDirectory"`
    // Filter the systems by groups for which a provisioning group references a group on the returned systems, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Filter systems on the specified groupDN, for those system types that have such a property.
    GroupDN []string `uriparametername:"groupDN"`
    // Filter systems on the groups on a system, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter systems on the exact name.
    Name []string `uriparametername:"name"`
    // Search systems on (part of) the name or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Only return systems for which the name does not start with the given prefix.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Only return systems for which the name starts with the given prefix.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Filter the systems on not having a provisioning group for a group on system on the returned systems, specified by id.
    NotInGroup []int64 `uriparametername:"notInGroup"`
    // Filter provisioned LDAPs on the numbering used, specified by id. This parameter supports composition with all parameters from the numberseq resource.
    Numbering []int64 `uriparametername:"numbering"`
    // Filter systems on organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter systems on organizational units, specified by id. This parameter is automatically set and primarily used for security permission enforcement.
    OrganizationalUnitForEnforcement []int64 `uriparametername:"organizationalUnitForEnforcement"`
    // Filter the systems for which the given group is owner, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter the systems by provisioning groups which reference a group on the returned systems, specified by id. This parameter supports composition with all parameters from the provisioning group resource.
    ProvisioningGroup []int64 `uriparametername:"provisioningGroup"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Only return systems for which self-service for existing groups on system is enabled.
    SelfServiceExistingGroups []bool `uriparametername:"selfServiceExistingGroups"`
    // Only return systems for which self-service for new groups on system is enabled.
    SelfServiceNewGroups []bool `uriparametername:"selfServiceNewGroups"`
    // Only return systems for which self-service for service accounts is enabled.
    SelfServiceServiceAccounts []bool `uriparametername:"selfServiceServiceAccounts"`
    // Filter systems on the specified serviceAccountDN, for those system types that have such a property.
    ServiceAccountDN []string `uriparametername:"serviceAccountDN"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter the systems on groups that perform technical administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64 `uriparametername:"technicalAdministrator"`
    // Filter the systems on the TLS setting used.
    // Deprecated: This property is deprecated, use TlsAsTLSLevel instead
    Tls []string `uriparametername:"tls"`
    // Filter the systems on the TLS setting used.
    TlsAsTLSLevel []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.TLSLevel `uriparametername:"tls"`
    // Only return systems of the given type(s).
    Type []string `uriparametername:"type"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// SystemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SystemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SystemRequestBuilderGetQueryParameters
}
// SystemRequestBuilderPostQueryParameters creates one or more new provisioned systems and returns the newly created systems.
type SystemRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// SystemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SystemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SystemRequestBuilderPostQueryParameters
}
// BySystemid gets an item from the github.com/topicuskeyhub/sdk-go.system.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithSystemItemRequestBuilder when successful
func (m *SystemRequestBuilder) BySystemid(systemid string)(*WithSystemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if systemid != "" {
        urlTplParams["systemid"] = systemid
    }
    return NewWithSystemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// BySystemidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.system.item collection
// returns a *WithSystemItemRequestBuilder when successful
func (m *SystemRequestBuilder) BySystemidInt64(systemid int64)(*WithSystemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["systemid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(systemid, 10)
    return NewWithSystemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSystemRequestBuilderInternal instantiates a new SystemRequestBuilder and sets the default values.
func NewSystemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemRequestBuilder) {
    m := &SystemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system{?account*,active*,additional*,any*,contentAdministrator*,createdAfter*,createdBefore*,exclude*,expiredCertificate*,externalUuid*,forAzureOIDCDirectory*,forBaseSystem*,forClient*,forLDAPDirectory*,group*,groupDN*,groupOnSystem*,id*,modifiedSince*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,notInGroup*,numbering*,organizationalUnit*,organizationalUnitForEnforcement*,ownedBy*,provisioningGroup*,q*,selfServiceExistingGroups*,selfServiceNewGroups*,selfServiceServiceAccounts*,serviceAccountDN*,sort*,technicalAdministrator*,tls*,type*,uuid*}", pathParameters),
    }
    return m
}
// NewSystemRequestBuilder instantiates a new SystemRequestBuilder and sets the default values.
func NewSystemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSystemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all provisioned systems in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ProvisioningProvisionedSystemLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *SystemRequestBuilder) Get(ctx context.Context, requestConfiguration *SystemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable), nil
}
// Post creates one or more new provisioned systems and returns the newly created systems.
// returns a ProvisioningProvisionedSystemLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *SystemRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, requestConfiguration *SystemRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable), nil
}
// ToGetRequestInformation query for all provisioned systems in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *SystemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SystemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=69")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new provisioned systems and returns the newly created systems.
// returns a *RequestInformation when successful
func (m *SystemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, requestConfiguration *SystemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/system{?additional*}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=69")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=69", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SystemRequestBuilder when successful
func (m *SystemRequestBuilder) WithUrl(rawUrl string)(*SystemRequestBuilder) {
    return NewSystemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
