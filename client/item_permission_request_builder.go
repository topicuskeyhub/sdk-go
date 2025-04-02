package client

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i1dd7e63ea07df257d2f3b0a697e0b82b6f546fa79219313f15a549ec152a7c52 "github.com/topicuskeyhub/sdk-go/client/item/permission"
)

// ItemPermissionRequestBuilder builds and executes requests for operations under \client\{clientid}\permission
type ItemPermissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPermissionRequestBuilderGetQueryParameters query for all permissions given to a OAuth2 client in Topicus KeyHub. The various query parameters can be used to filter the response.
type ItemPermissionRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i1dd7e63ea07df257d2f3b0a697e0b82b6f546fa79219313f15a549ec152a7c52.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter permissions on the clients to which they are given, specified by id. This parameter supports composition with all parameters from the client resource.
    Client []int64 `uriparametername:"client"`
    // Filter permissions on the administrator groups of the clients to which they are given, specified by id. This parameter supports composition with all parameters from the client resource.
    ClientAdministratorGroup []int64 `uriparametername:"clientAdministratorGroup"`
    // Filter permissions on the owner groups of the clients to which they are given, specified by id. This parameter supports composition with all parameters from the client resource.
    ClientOwnerGroup []int64 `uriparametername:"clientOwnerGroup"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter permissions on the groups to which they apply, specified by id. This parameter supports composition with all parameters from the group resource.
    ForGroup []int64 `uriparametername:"forGroup"`
    // Filter permissions on the systems to which they apply, specified by id. This parameter supports composition with all parameters from the system resource.
    ForSystem []int64 `uriparametername:"forSystem"`
    // Filter permissions on the content administrator groups of the systems to which they apply, specified by id. This parameter supports composition with all parameters from the group resource.
    ForSystemContentAdministratorGroup []int64 `uriparametername:"forSystemContentAdministratorGroup"`
    // Filter permissions on the owner groups of the systems to which they apply, specified by id. This parameter supports composition with all parameters from the group resource.
    ForSystemOwnerGroup []int64 `uriparametername:"forSystemOwnerGroup"`
    // Filter permissions on the technical administrator groups of the systems to which they apply, specified by id. This parameter supports composition with all parameters from the group resource.
    ForSystemTechnicalAdministratorGroup []int64 `uriparametername:"forSystemTechnicalAdministratorGroup"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter permissions on the permission type(s).
    // Deprecated: This property is deprecated, use ValueAsClientOAuth2ClientPermissionType instead
    Value []string `uriparametername:"value"`
    // Filter permissions on the permission type(s).
    ValueAsClientOAuth2ClientPermissionType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionType `uriparametername:"value"`
    // Filter the clients for the permissions with permissions for the given groups, either directly or via provisionedsystem ownership, specified by id.
    WithPermissionForOwningGroup []int64 `uriparametername:"withPermissionForOwningGroup"`
    // Filter the clients for the permissions with active requests for permissions for the given groups, either directly or via provisionedsystem ownership, specified by id.
    WithRequestedPermissionForOwningGroup []int64 `uriparametername:"withRequestedPermissionForOwningGroup"`
}
// ItemPermissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPermissionRequestBuilderGetQueryParameters
}
// ByPermissionid gets an item from the github.com/topicuskeyhub/sdk-go.client.item.permission.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemPermissionWithPermissionItemRequestBuilder when successful
func (m *ItemPermissionRequestBuilder) ByPermissionid(permissionid string)(*ItemPermissionWithPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionid != "" {
        urlTplParams["permissionid"] = permissionid
    }
    return NewItemPermissionWithPermissionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPermissionidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.client.item.permission.item collection
// returns a *ItemPermissionWithPermissionItemRequestBuilder when successful
func (m *ItemPermissionRequestBuilder) ByPermissionidInt64(permissionid int64)(*ItemPermissionWithPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["permissionid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(permissionid, 10)
    return NewItemPermissionWithPermissionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPermissionRequestBuilderInternal instantiates a new ItemPermissionRequestBuilder and sets the default values.
func NewItemPermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionRequestBuilder) {
    m := &ItemPermissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/{clientid}/permission{?additional*,any*,client*,clientAdministratorGroup*,clientOwnerGroup*,createdAfter*,createdBefore*,exclude*,forGroup*,forSystem*,forSystemContentAdministratorGroup*,forSystemOwnerGroup*,forSystemTechnicalAdministratorGroup*,id*,modifiedSince*,q*,sort*,value*,withPermissionForOwningGroup*,withRequestedPermissionForOwningGroup*}", pathParameters),
    }
    return m
}
// NewItemPermissionRequestBuilder instantiates a new ItemPermissionRequestBuilder and sets the default values.
func NewItemPermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all permissions given to a OAuth2 client in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ClientOAuth2ClientPermissionLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemPermissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPermissionRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientOAuth2ClientPermissionLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionLinkableWrapperable), nil
}
// ToGetRequestInformation query for all permissions given to a OAuth2 client in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemPermissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPermissionRequestBuilder when successful
func (m *ItemPermissionRequestBuilder) WithUrl(rawUrl string)(*ItemPermissionRequestBuilder) {
    return NewItemPermissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
