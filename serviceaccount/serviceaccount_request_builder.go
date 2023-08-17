package serviceaccount

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ServiceaccountRequestBuilder builds and executes requests for operations under \serviceaccount
type ServiceaccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceaccountRequestBuilderGetQueryParameters query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
type ServiceaccountRequestBuilderGetQueryParameters struct {
    // Only return active or inactive service accounts.
    Active []string
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
    // Filter the service accounts by groups on systems, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64
    // Filter the service accounts on group on systems owned by any of the given groups, specified by id.
    GroupOnSystemOwners []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter service accounts on the exact name.
    Name []string
    // Search service accounts on (part of) the name, username or uuid.
    NameContains []string
    // Only return service accounts for which the name does not start with the given prefix.
    NameDoesNotStartWith []string
    // Only return service accounts for which the name starts with the given prefix.
    NameStartsWith []string
    // Filter the service accounts by the password shared in the vault, specified by id. This parameter supports composition with all parameters from the vault record resource.
    Password []int64
    // Only return service accounts with the given password rotation scheme.
    PasswordRotation []string
    // Filter records on a complex CQL query.
    Q []string
    // Filter the service accounts on active requests for group on systems owned by any of the given groups, specified by id.
    RequestedGroupOnSystemOwners []int64
    // Filter the service accounts by provisioned systems, specified by id. This parameter supports composition with all parameters from the provisioned system resource.
    System []int64
    // Filter the service accounts on groups that perform technical administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64
    // Filter service accounts on the exact username.
    Username []string
    // Filter results on one or more UUIDs.
    Uuid []string
}
// ServiceaccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceaccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceaccountRequestBuilderGetQueryParameters
}
// ServiceaccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceaccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Auditstats the auditstats property
func (m *ServiceaccountRequestBuilder) Auditstats()(*AuditstatsRequestBuilder) {
    return NewAuditstatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccountid gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ServiceaccountRequestBuilder) ByAccountid(accountid string)(*WithAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accountid != "" {
        urlTplParams["accountid"] = accountid
    }
    return NewWithAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAccountidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item collection
func (m *ServiceaccountRequestBuilder) ByAccountidInt64(accountid int64)(*WithAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["accountid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(accountid, 10)
    return NewWithAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceaccountRequestBuilderInternal instantiates a new ServiceaccountRequestBuilder and sets the default values.
func NewServiceaccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceaccountRequestBuilder) {
    m := &ServiceaccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,active*,groupOnSystem*,groupOnSystemOwners*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,password*,passwordRotation*,requestedGroupOnSystemOwners*,system*,technicalAdministrator*,username*,uuid*}", pathParameters),
    }
    return m
}
// NewServiceaccountRequestBuilder instantiates a new ServiceaccountRequestBuilder and sets the default values.
func NewServiceaccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceaccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceaccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *ServiceaccountRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ServiceaccountRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceaccountRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateServiceaccountServiceAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable), nil
}
// Group the group property
func (m *ServiceaccountRequestBuilder) Group()(*GroupRequestBuilder) {
    return NewGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates one or more new service accounts and returns the newly created service accounts.
func (m *ServiceaccountRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, requestConfiguration *ServiceaccountRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateServiceaccountServiceAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable), nil
}
// ToGetRequestInformation query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *ServiceaccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceaccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new service accounts and returns the newly created service accounts.
func (m *ServiceaccountRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, requestConfiguration *ServiceaccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
