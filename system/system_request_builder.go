package system

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
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
    Account []int64
    // Only return active or inactive systems.
    Active []string
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter the systems on groups that perform content administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    ContentAdministrator []int64
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return systems for which one of the certificates used is expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the systems on one or more UUIDs as read from the system.
    ExternalUuid []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Filter systems on the Azure OIDC directory for which it is a provisioned source directory, specified by id. This parameter supports composition with all parameters from the directory resource.
    ForAzureOIDCDirectory []int64
    // Filter systems on the clients for which it is an internal provisioned LDAP, specified by id. This parameter supports composition with all parameters from the client resource.
    ForClient []int64
    // Filter systems on the LDAP directory for which it is a provisioned source directory, specified by id. This parameter supports composition with all parameters from the directory resource.
    ForLDAPDirectory []int64
    // Filter the systems by groups for which a provisioning group references a group on the returned systems, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Filter systems on the groups on a system, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter systems on the exact name.
    Name []string
    // Search systems on (part of) the name or uuid.
    NameContains []string
    // Only return systems for which the name does not start with the given prefix.
    NameDoesNotStartWith []string
    // Only return systems for which the name starts with the given prefix.
    NameStartsWith []string
    // Filter the systems on not having a provisioning group for a group on system on the returned systems, specified by id.
    NotInGroup []int64
    // Filter provisioned LDAPs on the numbering used, specified by id. This parameter supports composition with all parameters from the numberseq resource.
    Numbering []int64
    // Filter systems on organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64
    // Filter systems on organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnitForEnforcement []int64
    // Filter the systems for which the given group is owner, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64
    // Filter the systems by provisioning groups which reference a group on the returned systems, specified by id. This parameter supports composition with all parameters from the provisioning group resource.
    ProvisioningGroup []int64
    // Filter records on a complex CQL query.
    Q []string
    // Only return systems for which self-service for existing groups on system is enabled.
    SelfServiceExistingGroups []bool
    // Only return systems for which self-service for new groups on system is enabled.
    SelfServiceNewGroups []bool
    // Only return systems for which self-service for service accounts is enabled.
    SelfServiceServiceAccounts []bool
    // Filter the systems on groups that perform technical administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64
    // Filter the systems on the TLS setting used.
    Tls []string
    // Only return systems of the given type(s).
    Type []string
    // Filter results on one or more UUIDs.
    Uuid []string
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
// SystemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SystemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.system.item collection
func (m *SystemRequestBuilder) ById(id string)(*SystemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewSystemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSystemRequestBuilderInternal instantiates a new SystemRequestBuilder and sets the default values.
func NewSystemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemRequestBuilder) {
    m := &SystemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,account*,active*,contentAdministrator*,expiredCertificate*,externalUuid*,forAzureOIDCDirectory*,forClient*,forLDAPDirectory*,group*,groupOnSystem*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,notInGroup*,numbering*,organizationalUnit*,organizationalUnitForEnforcement*,ownedBy*,provisioningGroup*,selfServiceExistingGroups*,selfServiceNewGroups*,selfServiceServiceAccounts*,technicalAdministrator*,tls*,type*,uuid*}", pathParameters),
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
func (m *SystemRequestBuilder) Get(ctx context.Context, requestConfiguration *SystemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *SystemRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, requestConfiguration *SystemRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *SystemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SystemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new provisioned systems and returns the newly created systems.
func (m *SystemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemLinkableWrapperable, requestConfiguration *SystemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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