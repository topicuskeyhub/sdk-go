// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package accessprofileclient

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// AccessprofileclientRequestBuilder builds and executes requests for operations under \accessprofileclient
type AccessprofileclientRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessprofileclientRequestBuilderGetQueryParameters queries over all client links for an access profile. The various query parameters can be used to filter the response.
type AccessprofileclientRequestBuilderGetQueryParameters struct {
    // Filter client access profile links on the given access profiles, specified by id. This parameter supports composition with all parameters from the access profile resource.
    AccessProfile []int64 `uriparametername:"accessProfile"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter client access profile links on the given clients, specified by id. This parameter supports composition with all parameters from the client resource.
    Client []int64 `uriparametername:"client"`
    // Only return access profile client links for which one of the given groups is the owner of its client application, specified by id. This parameter supports composition with all parameters from the groups resource.
    ClientOwnerGroup []int64 `uriparametername:"clientOwnerGroup"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter access profile client links for which the given groups are owner, specified by id. This parameter supports composition with all parameters from the access profile resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
}
// AccessprofileclientRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessprofileclientRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessprofileclientRequestBuilderGetQueryParameters
}
// ByAccessprofileclientid gets an item from the github.com/topicuskeyhub/sdk-go.accessprofileclient.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithAccessprofileclientItemRequestBuilder when successful
func (m *AccessprofileclientRequestBuilder) ByAccessprofileclientid(accessprofileclientid string)(*WithAccessprofileclientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessprofileclientid != "" {
        urlTplParams["accessprofileclientid"] = accessprofileclientid
    }
    return NewWithAccessprofileclientItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccessprofileclientidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.accessprofileclient.item collection
// returns a *WithAccessprofileclientItemRequestBuilder when successful
func (m *AccessprofileclientRequestBuilder) ByAccessprofileclientidInt64(accessprofileclientid int64)(*WithAccessprofileclientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["accessprofileclientid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(accessprofileclientid, 10)
    return NewWithAccessprofileclientItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccessprofileclientRequestBuilderInternal instantiates a new AccessprofileclientRequestBuilder and sets the default values.
func NewAccessprofileclientRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessprofileclientRequestBuilder) {
    m := &AccessprofileclientRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessprofileclient{?accessProfile*,additional*,any*,client*,clientOwnerGroup*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,ownedBy*,q*,sort*}", pathParameters),
    }
    return m
}
// NewAccessprofileclientRequestBuilder instantiates a new AccessprofileclientRequestBuilder and sets the default values.
func NewAccessprofileclientRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessprofileclientRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessprofileclientRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all client links for an access profile. The various query parameters can be used to filter the response.
// returns a ProfileAccessProfileClientLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *AccessprofileclientRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessprofileclientRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileClientLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProfileAccessProfileClientLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileClientLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all client links for an access profile. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *AccessprofileclientRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessprofileclientRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=79")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessprofileclientRequestBuilder when successful
func (m *AccessprofileclientRequestBuilder) WithUrl(rawUrl string)(*AccessprofileclientRequestBuilder) {
    return NewAccessprofileclientRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
