package client

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ClientRequestBuilder builds and executes requests for operations under \client
type ClientRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClientRequestBuilderGetQueryParameters query for all clients in Topicus KeyHub. The various query parameters can be used to filter the response.
type ClientRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter clients on the client ids.
    ClientId []string
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return LDAP clients for which the client certificate is expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the SSO applications connected to groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return LDAP clients that are used as provisioned internal LDAP.
    IsProvisionedInternalLDAP []bool
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter clients on the exact name.
    Name []string
    // Search clients on (part of) the name, client id or uuid.
    NameContains []string
    // Only return clients for which the name does not start with one of the given values.
    NameDoesNotStartWith []string
    // Only return clients for which the name starts with one of the given values.
    NameStartsWith []string
    // Filter the SSO applications not connected to groups, specified by id.
    NotInGroup []int64
    // Filter the clients for which the given group is owner, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64
    // Filter records on a complex CQL query.
    Q []string
    // Filter the clients by the secret shared in a vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    SharedSecret []int64
    // Filter the clients for which the given group is technical administrator, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64
    // Only return clients of the given type(s).
    Type []string
    // Only return OAuth2 clients that use or do not use the client credentials grant.
    UseClientCredentials []bool
    // Filter results on one or more UUIDs.
    Uuid []string
    // Filter the OAuth 2.0 clients by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64
    // Only return OAuth2 clients that have specific permissions, specified by id. This parameter supports composition with all parameters from the client permission resource.
    WithPermission []int64
    // Filter the client applications with permissions for the given groups, either directly or via provisionedsystem ownership, specified by id. 
    WithPermissionForOwningGroup []int64
    // Filter the client applications with active requests for permissions for the given groups, either directly or via provisionedsystem ownership, specified by id. 
    WithRequestedPermissionForOwningGroup []int64
}
// ClientRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClientRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClientRequestBuilderGetQueryParameters
}
// ClientRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClientRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByClientid gets an item from the github.com/topicuskeyhub/sdk-go.client.item collection
func (m *ClientRequestBuilder) ByClientid(clientid string)(*WithClientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if clientid != "" {
        urlTplParams["clientid"] = clientid
    }
    return NewWithClientItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClientRequestBuilderInternal instantiates a new ClientRequestBuilder and sets the default values.
func NewClientRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientRequestBuilder) {
    m := &ClientRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,clientId*,expiredCertificate*,group*,isProvisionedInternalLDAP*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,notInGroup*,ownedBy*,sharedSecret*,technicalAdministrator*,type*,useClientCredentials*,uuid*,vault*,withPermission*,withPermissionForOwningGroup*,withRequestedPermissionForOwningGroup*}", pathParameters),
    }
    return m
}
// NewClientRequestBuilder instantiates a new ClientRequestBuilder and sets the default values.
func NewClientRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClientRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all clients in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ClientRequestBuilder) Get(ctx context.Context, requestConfiguration *ClientRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientClientApplicationLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable), nil
}
// Me the me property
func (m *ClientRequestBuilder) Me()(*MeRequestBuilder) {
    return NewMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates one or more new clients and returns the newly created clients.
func (m *ClientRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, requestConfiguration *ClientRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientClientApplicationLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable), nil
}
// ToGetRequestInformation query for all clients in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ClientRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClientRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new clients and returns the newly created clients.
func (m *ClientRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, requestConfiguration *ClientRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Vault the vault property
func (m *ClientRequestBuilder) Vault()(*VaultRequestBuilder) {
    return NewVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
