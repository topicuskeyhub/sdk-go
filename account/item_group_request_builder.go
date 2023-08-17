package account

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemGroupRequestBuilder builds and executes requests for operations under \account\{accountid}\group
type ItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupRequestBuilderGetQueryParameters queries over all groups the account is member of. The various query parameters can be used to filter the response.
type ItemGroupRequestBuilderGetQueryParameters struct {
    // Filter group memberships on the given accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Only return group memberships that can or cannot be activated.
    ApplicableForActivation []bool
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return group memberships that have an expiration set that is before the given date.
    ExpiredAt []i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Filter group memberships on the given groups, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64
    // Filter group memberships on (part of) the name or uuid of the group or the folder (if any).
    GroupNameContains []string
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return group memberships that are or are not nested.
    Nested []bool
    // Filter records on a complex CQL query.
    Q []string
    // Filter group memberships on the given rights.
    Rights []string
    // Filter group memberships on the given signatures, specified by id. This parameter supports composition with all parameters from the signature filter.
    Signature []int64
    // Only return group memberships that have or do not have access to the group's vault.
    VaultAccess []bool
}
// ItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGroupRequestBuilderGetQueryParameters
}
// ByAccountgroupid gets an item from the github.com/topicuskeyhub/sdk-go.account.item.group.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemGroupRequestBuilder) ByAccountgroupid(accountgroupid string)(*ItemGroupWithAccountgroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accountgroupid != "" {
        urlTplParams["accountgroupid"] = accountgroupid
    }
    return NewItemGroupWithAccountgroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccountgroupidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.account.item.group.item collection
func (m *ItemGroupRequestBuilder) ByAccountgroupidInt64(accountgroupid int64)(*ItemGroupWithAccountgroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["accountgroupid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(accountgroupid, 10)
    return NewItemGroupWithAccountgroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    m := &ItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/group{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,account*,applicableForActivation*,expiredAt*,group*,groupNameContains*,nested*,rights*,signature*,vaultAccess*}", pathParameters),
    }
    return m
}
// NewItemGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all groups the account is member of. The various query parameters can be used to filter the response.
func (m *ItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupAccountGroupLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupAccountGroupLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupAccountGroupLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all groups the account is member of. The various query parameters can be used to filter the response.
func (m *ItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
