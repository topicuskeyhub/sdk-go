// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package audit

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
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
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Only return records after a given date.
    After []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"after"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records before a given date.
    Before []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"before"`
    // Perform a full text search on the contents of audit records.
    Containing []string `uriparametername:"containing"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Include or do not include records that are considered daily use.
    IncludeDaily []bool `uriparametername:"includeDaily"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Only return audit records targeting the given group either via onGroup1 or onGroup2, specified by id.
    OnGroup []int64 `uriparametername:"onGroup"`
    // Only return audit records targeting the given group specifically as onGroup1, specified by id.
    OnGroup1 []int64 `uriparametername:"onGroup1"`
    // Only return audit records performed by or targeting the given account, specified by id.
    OnOrByAccount []int64 `uriparametername:"onOrByAccount"`
    // Only return audit records performed by or targeting the given client, specified by id.
    OnOrByClient []int64 `uriparametername:"onOrByClient"`
    // Only return audit records targeting the given service account, specified by id.
    OnServiceAccount []int64 `uriparametername:"onServiceAccount"`
    // Only return audit records that have the performed-by-name set or not set.
    PerformedByNameSet []bool `uriparametername:"performedByNameSet"`
    // Only return audit records that are performed by the given account or client, specified by uuid.
    PerformedByUuid []string `uriparametername:"performedByUuid"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter audit records on the given type(s).
    // Deprecated: This property is deprecated, use TypeAsAuditAuditRecordType instead
    Type []string `uriparametername:"type"`
    // Filter audit records on the given type(s).
    TypeAsAuditAuditRecordType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordType `uriparametername:"type"`
    // Only return audit records that are meant to be displayed to the given account, specified by id.
    VisibleFor []int64 `uriparametername:"visibleFor"`
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
// ByAuditid gets an item from the github.com/topicuskeyhub/sdk-go.audit.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithAuditItemRequestBuilder when successful
func (m *AuditRequestBuilder) ByAuditid(auditid string)(*WithAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if auditid != "" {
        urlTplParams["auditid"] = auditid
    }
    return NewWithAuditItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAuditidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.audit.item collection
// returns a *WithAuditItemRequestBuilder when successful
func (m *AuditRequestBuilder) ByAuditidInt64(auditid int64)(*WithAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["auditid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(auditid, 10)
    return NewWithAuditItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuditRequestBuilderInternal instantiates a new AuditRequestBuilder and sets the default values.
func NewAuditRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditRequestBuilder) {
    m := &AuditRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit{?additional*,after*,any*,before*,containing*,createdAfter*,createdBefore*,exclude*,id*,includeDaily*,modifiedSince*,onGroup*,onGroup1*,onOrByAccount*,onOrByClient*,onServiceAccount*,performedByNameSet*,performedByUuid*,q*,sort*,type*,visibleFor*}", pathParameters),
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
// returns a *ExportRequestBuilder when successful
func (m *AuditRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all audit records in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a AuditAuditRecordLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *AuditRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuditAuditRecordLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *AuditRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuditRequestBuilder when successful
func (m *AuditRequestBuilder) WithUrl(rawUrl string)(*AuditRequestBuilder) {
    return NewAuditRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
