package profileprovisioning

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ProfileprovisioningRequestBuilder builds and executes requests for operations under \profileprovisioning
type ProfileprovisioningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProfileprovisioningRequestBuilderGetQueryParameters query for all access profile provisioning links in Topicus KeyHub. The various query parameters can be used to filter the response.
type ProfileprovisioningRequestBuilderGetQueryParameters struct {
    // Filter access profile provisioning links on the given access profiles, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter access profile provisioning links on the given group on systems, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Only return access profile provisioning links for which one of the given groups is the owner of its group on system, specified by id. This parameter supports composition with all parameters from the groups resource.
    GroupOnSystemOwnerGroup []int64 `uriparametername:"groupOnSystemOwnerGroup"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter access profile provisioning links for which the given groups are owner, specified by id. This parameter supports composition with all parameters from the access profile resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Only return access profile provisioning links that are provisionined on one of the given systems, specified by id. This parameter supports composition with all parameters from the systems resource.
    System []int64 `uriparametername:"system"`
    // Only return access profile provisioning links for which one of the given groups is technical administrator for the system of its group on system, specified by id. This parameter supports composition with all parameters from the groups resource.
    SystemAdminGroup []int64 `uriparametername:"systemAdminGroup"`
    // Only return access profile provisioning links for which one of the given groups is content administrator for the system of its group on system, specified by id. This parameter supports composition with all parameters from the groups resource.
    SystemContentAdminGroup []int64 `uriparametername:"systemContentAdminGroup"`
    // Only return access profile provisioning links for which one of the given groups is the owner of its provisioned system, specified by id. This parameter supports composition with all parameters from the groups resource.
    SystemOwnerGroup []int64 `uriparametername:"systemOwnerGroup"`
}
// ProfileprovisioningRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProfileprovisioningRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProfileprovisioningRequestBuilderGetQueryParameters
}
// ByProfileprovisioningid gets an item from the github.com/topicuskeyhub/sdk-go.profileprovisioning.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithProfileprovisioningItemRequestBuilder when successful
func (m *ProfileprovisioningRequestBuilder) ByProfileprovisioningid(profileprovisioningid string)(*WithProfileprovisioningItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if profileprovisioningid != "" {
        urlTplParams["profileprovisioningid"] = profileprovisioningid
    }
    return NewWithProfileprovisioningItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByProfileprovisioningidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.profileprovisioning.item collection
// returns a *WithProfileprovisioningItemRequestBuilder when successful
func (m *ProfileprovisioningRequestBuilder) ByProfileprovisioningidInt64(profileprovisioningid int64)(*WithProfileprovisioningItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["profileprovisioningid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(profileprovisioningid, 10)
    return NewWithProfileprovisioningItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProfileprovisioningRequestBuilderInternal instantiates a new ProfileprovisioningRequestBuilder and sets the default values.
func NewProfileprovisioningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProfileprovisioningRequestBuilder) {
    m := &ProfileprovisioningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profileprovisioning{?accessProfile*,additional*,any*,createdAfter*,createdBefore*,exclude*,groupOnSystem*,groupOnSystemOwnerGroup*,id*,modifiedSince*,ownedBy*,q*,sort*,system*,systemAdminGroup*,systemContentAdminGroup*,systemOwnerGroup*}", pathParameters),
    }
    return m
}
// NewProfileprovisioningRequestBuilder instantiates a new ProfileprovisioningRequestBuilder and sets the default values.
func NewProfileprovisioningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProfileprovisioningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProfileprovisioningRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all access profile provisioning links in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ProfileAccessProfileProvisioningLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ProfileprovisioningRequestBuilder) Get(ctx context.Context, requestConfiguration *ProfileprovisioningRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileProvisioningLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProfileAccessProfileProvisioningLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileProvisioningLinkableWrapperable), nil
}
// ToGetRequestInformation query for all access profile provisioning links in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ProfileprovisioningRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProfileprovisioningRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProfileprovisioningRequestBuilder when successful
func (m *ProfileprovisioningRequestBuilder) WithUrl(rawUrl string)(*ProfileprovisioningRequestBuilder) {
    return NewProfileprovisioningRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
