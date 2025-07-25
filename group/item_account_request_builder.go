// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package group

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i0bdcc256f19e350e2ebf5365f1da8e8f159a5387b2eb9c7690e07c6c164452af "github.com/topicuskeyhub/sdk-go/group/item/account"
)

// ItemAccountRequestBuilder builds and executes requests for operations under \group\{groupid}\account
type ItemAccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAccountRequestBuilderGetQueryParameters queries over all members of a group. The various query parameters can be used to filter the response.
type ItemAccountRequestBuilderGetQueryParameters struct {
    // Filter group memberships on the given accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64 `uriparametername:"account"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i0bdcc256f19e350e2ebf5365f1da8e8f159a5387b2eb9c7690e07c6c164452af.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return group memberships that can or cannot be activated by the current user.
    ApplicableForActivation []bool `uriparametername:"applicableForActivation"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return group memberships that have an expiration set that is before the given date.
    ExpiredAt []i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly `uriparametername:"expiredAt"`
    // Filter group memberships on the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Filter group memberships on (part of) the name or uuid of the group or the folder (if any).
    GroupNameContains []string `uriparametername:"groupNameContains"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Only return group memberships that are or are not nested.
    Nested []bool `uriparametername:"nested"`
    // Filter group memberships on the given organizational unit of the group, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter group memberships on the given rights.
    // Deprecated: This property is deprecated, use RightsAsGroupGroupRights instead
    Rights []string `uriparametername:"rights"`
    // Filter group memberships on the given rights.
    RightsAsGroupGroupRights []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupRights `uriparametername:"rights"`
    // Filter group memberships on the given signatures, specified by id. This parameter supports composition with all parameters from the signature filter.
    Signature []int64 `uriparametername:"signature"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Only return group memberships that have or do not have access to the group's vault.
    VaultAccess []bool `uriparametername:"vaultAccess"`
}
// ItemAccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAccountRequestBuilderGetQueryParameters
}
// ItemAccountRequestBuilderPostQueryParameters creates one or more new memberships for a group and returns the newly created memberships.
type ItemAccountRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []i0bdcc256f19e350e2ebf5365f1da8e8f159a5387b2eb9c7690e07c6c164452af.PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAccountRequestBuilderPostQueryParameters
}
// ByAccountid gets an item from the github.com/topicuskeyhub/sdk-go.group.item.account.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemAccountWithAccountItemRequestBuilder when successful
func (m *ItemAccountRequestBuilder) ByAccountid(accountid string)(*ItemAccountWithAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accountid != "" {
        urlTplParams["accountid"] = accountid
    }
    return NewItemAccountWithAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccountidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.group.item.account.item collection
// returns a *ItemAccountWithAccountItemRequestBuilder when successful
func (m *ItemAccountRequestBuilder) ByAccountidInt64(accountid int64)(*ItemAccountWithAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["accountid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(accountid, 10)
    return NewItemAccountWithAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAccountRequestBuilderInternal instantiates a new ItemAccountRequestBuilder and sets the default values.
func NewItemAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountRequestBuilder) {
    m := &ItemAccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}/account{?account*,additional*,any*,applicableForActivation*,createdAfter*,createdBefore*,exclude*,expiredAt*,group*,groupNameContains*,id*,modifiedSince*,nested*,organizationalUnit*,q*,rights*,signature*,sort*,vaultAccess*}", pathParameters),
    }
    return m
}
// NewItemAccountRequestBuilder instantiates a new ItemAccountRequestBuilder and sets the default values.
func NewItemAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all members of a group. The various query parameters can be used to filter the response.
// returns a GroupGroupAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAccountRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable), nil
}
// Post creates one or more new memberships for a group and returns the newly created memberships.
// returns a GroupGroupAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable, requestConfiguration *ItemAccountRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all members of a group. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemAccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new memberships for a group and returns the newly created memberships.
// returns a *RequestInformation when successful
func (m *ItemAccountRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountLinkableWrapperable, requestConfiguration *ItemAccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=79")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=79", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAccountRequestBuilder when successful
func (m *ItemAccountRequestBuilder) WithUrl(rawUrl string)(*ItemAccountRequestBuilder) {
    return NewItemAccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
