package client

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
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
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter clients on the client ids.
    ClientId []string `uriparametername:"clientId"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return LDAP clients for which the client certificate is expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"expiredCertificate"`
    // Filter the SSO applications connected to groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return LDAP clients that are used as provisioned internal LDAP.
    IsProvisionedInternalLDAP []bool `uriparametername:"isProvisionedInternalLDAP"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter clients on the exact name.
    Name []string `uriparametername:"name"`
    // Search clients on (part of) the name, client id or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Only return clients for which the name does not start with one of the given values.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Only return clients for which the name starts with one of the given values.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Filter the SSO applications not connected to groups, specified by id.
    NotInGroup []int64 `uriparametername:"notInGroup"`
    // Filter the clients for which the given group is owner, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter the clients by the secret shared in a vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    SharedSecret []int64 `uriparametername:"sharedSecret"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter the clients for which the given group is technical administrator, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64 `uriparametername:"technicalAdministrator"`
    // Only return clients of the given type(s).
    Type []string `uriparametername:"type"`
    // Only return OAuth2 clients that use or do not use the client credentials grant.
    UseClientCredentials []bool `uriparametername:"useClientCredentials"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
    // Filter the OAuth 2.0 clients by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64 `uriparametername:"vault"`
    // Only return OAuth2 clients that have specific permissions, specified by id. This parameter supports composition with all parameters from the client permission resource.
    WithPermission []int64 `uriparametername:"withPermission"`
    // Filter the client applications with permissions for the given groups, either directly or via provisionedsystem ownership, specified by id. 
    WithPermissionForOwningGroup []int64 `uriparametername:"withPermissionForOwningGroup"`
    // Filter the client applications with active requests for permissions for the given groups, either directly or via provisionedsystem ownership, specified by id. 
    WithRequestedPermissionForOwningGroup []int64 `uriparametername:"withRequestedPermissionForOwningGroup"`
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
// ClientRequestBuilderPostQueryParameters creates one or more new clients and returns the newly created clients.
type ClientRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ClientRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClientRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClientRequestBuilderPostQueryParameters
}
// ByClientid gets an item from the github.com/topicuskeyhub/sdk-go.client.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithClientItemRequestBuilder when successful
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
// ByClientidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.client.item collection
// returns a *WithClientItemRequestBuilder when successful
func (m *ClientRequestBuilder) ByClientidInt64(clientid int64)(*WithClientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["clientid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(clientid, 10)
    return NewWithClientItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClientRequestBuilderInternal instantiates a new ClientRequestBuilder and sets the default values.
func NewClientRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientRequestBuilder) {
    m := &ClientRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client{?additional*,any*,clientId*,createdAfter*,createdBefore*,exclude*,expiredCertificate*,group*,id*,isProvisionedInternalLDAP*,modifiedSince*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,notInGroup*,ownedBy*,q*,sharedSecret*,sort*,technicalAdministrator*,type*,useClientCredentials*,uuid*,vault*,withPermission*,withPermissionForOwningGroup*,withRequestedPermissionForOwningGroup*}", pathParameters),
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
// returns a ClientClientApplicationLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ClientRequestBuilder) Get(ctx context.Context, requestConfiguration *ClientRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// returns a *MeRequestBuilder when successful
func (m *ClientRequestBuilder) Me()(*MeRequestBuilder) {
    return NewMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates one or more new clients and returns the newly created clients.
// returns a ClientClientApplicationLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ClientRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, requestConfiguration *ClientRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *ClientRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClientRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new clients and returns the newly created clients.
// returns a *RequestInformation when successful
func (m *ClientRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationLinkableWrapperable, requestConfiguration *ClientRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=71", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Vault the vault property
// returns a *VaultRequestBuilder when successful
func (m *ClientRequestBuilder) Vault()(*VaultRequestBuilder) {
    return NewVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ClientRequestBuilder when successful
func (m *ClientRequestBuilder) WithUrl(rawUrl string)(*ClientRequestBuilder) {
    return NewClientRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
