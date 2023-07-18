package provisioninggroup

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ProvisioninggroupRequestBuilder builds and executes requests for operations under \provisioninggroup
type ProvisioninggroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProvisioninggroupRequestBuilderGetQueryParameters query for all provisioning groups in Topicus KeyHub. The various query parameters can be used to filter the response.
type ProvisioninggroupRequestBuilderGetQueryParameters struct {
    // Only return provisioning groups that do or do not require activation.
    ActivationRequired []bool
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return provisioning groups for one of the given groups, specified by id. This parameter supports composition with all parameters from the groups resource.
    Group []int64
    // Only return provisioning groups for one of the given groups on system, specified by id. This parameter supports composition with all parameters from the groups on system resource.
    GroupOnSystem []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter records on a complex CQL query.
    Q []string
    // Only return provisioning groups that are provisionined the on one of the given systems, specified by id. This parameter supports composition with all parameters from the systems resource.
    System []int64
    // Only return provisioning groups for which one of the given groups is technical administrator for its group on system, specified by id. This parameter supports composition with all parameters from the groups resource.
    SystemAdminGroup []int64
    // Only return provisioning groups for which one of the given groups is content administrator for its group on system, specified by id. This parameter supports composition with all parameters from the groups resource.
    SystemContentAdminGroup []int64
}
// ProvisioninggroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProvisioninggroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProvisioninggroupRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.provisioninggroup.item collection
func (m *ProvisioninggroupRequestBuilder) ById(id string)(*ProvisioninggroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewProvisioninggroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProvisioninggroupRequestBuilderInternal instantiates a new ProvisioninggroupRequestBuilder and sets the default values.
func NewProvisioninggroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioninggroupRequestBuilder) {
    m := &ProvisioninggroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/provisioninggroup{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,activationRequired*,group*,groupOnSystem*,system*,systemAdminGroup*,systemContentAdminGroup*}", pathParameters),
    }
    return m
}
// NewProvisioninggroupRequestBuilder instantiates a new ProvisioninggroupRequestBuilder and sets the default values.
func NewProvisioninggroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioninggroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvisioninggroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all provisioning groups in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ProvisioninggroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ProvisioninggroupRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupProvisioningGroupLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupProvisioningGroupLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupProvisioningGroupLinkableWrapperable), nil
}
// ToGetRequestInformation query for all provisioning groups in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ProvisioninggroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProvisioninggroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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