package audit

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// AuditRequestBuilder builds and executes requests for operations under \audit
type AuditRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuditRequestBuilderGetQueryParameters query for all audit records in Topicus KeyHub. The various query parameters can be used to filter the response.
type AuditRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
    // Only return records after a given date.
    After []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Only return records before a given date.
    Before []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Perform a full text search on the contents of audit records.
    Containing []string
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Filter the results on the given ids.
    Id []int64
    // Include or do not include records that are considered daily use.
    IncludeDaily []bool
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return audit records targeting the given group either via onGroup1 or onGroup2, specified by id.
    OnGroup []int64
    // Only return audit records targeting the given group specifically as onGroup1, specified by id.
    OnGroup1 []int64
    // Only return audit records performed by or targeting the given account, specified by id.
    OnOrByAccount []int64
    // Only return audit records performed by or targeting the given client, specified by id.
    OnOrByClient []int64
    // Only return audit records targeting the given service account, specified by id.
    OnServiceAccount []int64
    // Only return audit records that have the performed-by-name set or not set.
    PerformedByNameSet []bool
    // Only return audit records that are performed by the given account or client, specified by uuid.
    PerformedByUuid []string
    // Filter records on a complex CQL query.
    Q []string
    // Filter audit records on the given type(s).
    Type []string
    // Only return audit records that are meant to be displayed to the given account, specified by id.
    VisibleFor []int64
}
// AuditRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.audit.item collection
func (m *AuditRequestBuilder) ById(id string)(*AuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewAuditItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuditRequestBuilderInternal instantiates a new AuditRequestBuilder and sets the default values.
func NewAuditRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditRequestBuilder) {
    m := &AuditRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,after*,before*,containing*,includeDaily*,onGroup*,onGroup1*,onOrByAccount*,onOrByClient*,onServiceAccount*,performedByNameSet*,performedByUuid*,type*,visibleFor*}", pathParameters),
    }
    return m
}
// NewAuditRequestBuilder instantiates a new AuditRequestBuilder and sets the default values.
func NewAuditRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *AuditRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all audit records in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *AuditRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuditAuditRecordLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordLinkableWrapperable), nil
}
// ToGetRequestInformation query for all audit records in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *AuditRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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