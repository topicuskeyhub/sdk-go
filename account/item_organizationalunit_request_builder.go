package account

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i0711eb651f0b33ca1c048798c211cc32777a4acc749ac7f651ed1a9fa9a94b23 "github.com/topicuskeyhub/sdk-go/account/item/organizationalunit"
)

// ItemOrganizationalunitRequestBuilder builds and executes requests for operations under \account\{accountid}\organizationalunit
type ItemOrganizationalunitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationalunitRequestBuilderGetQueryParameters queries over all organizational units the account is member of. The various query parameters can be used to filter the response.
type ItemOrganizationalunitRequestBuilderGetQueryParameters struct {
    // Filter the organizational units-accounts by accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64 `uriparametername:"account"`
    // Filter the organizational units-accounts by groups owning the internal directory of the accounts, specified by id.
    AccountDirectoryOwnedBy []int64 `uriparametername:"accountDirectoryOwnedBy"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i0711eb651f0b33ca1c048798c211cc32777a4acc749ac7f651ed1a9fa9a94b23.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
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
    // Filter the organizational units-accounts by organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter the organizational units-accounts by groups owning the organizational unit, specified by id.
    OrganizationalUnitOwnedBy []int64 `uriparametername:"organizationalUnitOwnedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
}
// ItemOrganizationalunitRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationalunitRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOrganizationalunitRequestBuilderGetQueryParameters
}
// ByOrganizationalunitid gets an item from the github.com/topicuskeyhub/sdk-go.account.item.organizationalunit.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder when successful
func (m *ItemOrganizationalunitRequestBuilder) ByOrganizationalunitid(organizationalunitid string)(*ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if organizationalunitid != "" {
        urlTplParams["organizationalunitid"] = organizationalunitid
    }
    return NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByOrganizationalunitidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.account.item.organizationalunit.item collection
// returns a *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder when successful
func (m *ItemOrganizationalunitRequestBuilder) ByOrganizationalunitidInt64(organizationalunitid int64)(*ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["organizationalunitid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(organizationalunitid, 10)
    return NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOrganizationalunitRequestBuilderInternal instantiates a new ItemOrganizationalunitRequestBuilder and sets the default values.
func NewItemOrganizationalunitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitRequestBuilder) {
    m := &ItemOrganizationalunitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/organizationalunit{?account*,accountDirectoryOwnedBy*,additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,organizationalUnit*,organizationalUnitOwnedBy*,q*,sort*}", pathParameters),
    }
    return m
}
// NewItemOrganizationalunitRequestBuilder instantiates a new ItemOrganizationalunitRequestBuilder and sets the default values.
func NewItemOrganizationalunitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationalunitRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all organizational units the account is member of. The various query parameters can be used to filter the response.
// returns a OrganizationAccountOrganizationalUnitLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemOrganizationalunitRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationalunitRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationAccountOrganizationalUnitLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateOrganizationAccountOrganizationalUnitLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationAccountOrganizationalUnitLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all organizational units the account is member of. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemOrganizationalunitRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationalunitRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=75")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOrganizationalunitRequestBuilder when successful
func (m *ItemOrganizationalunitRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationalunitRequestBuilder) {
    return NewItemOrganizationalunitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
