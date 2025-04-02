package audit

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ExportRequestBuilder builds and executes requests for operations under \audit\export
type ExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExportRequestBuilderPostQueryParameters prepares an export of audit records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
type ExportRequestBuilderPostQueryParameters struct {
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
// ExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExportRequestBuilderPostQueryParameters
}
// NewExportRequestBuilderInternal instantiates a new ExportRequestBuilder and sets the default values.
func NewExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportRequestBuilder) {
    m := &ExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit/export{?after*,any*,before*,containing*,createdAfter*,createdBefore*,exclude*,id*,includeDaily*,modifiedSince*,onGroup*,onGroup1*,onOrByAccount*,onOrByClient*,onServiceAccount*,performedByNameSet*,performedByUuid*,q*,sort*,type*,visibleFor*}", pathParameters),
    }
    return m
}
// NewExportRequestBuilder instantiates a new ExportRequestBuilder and sets the default values.
func NewExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post prepares an export of audit records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ExportRequestBuilder) Post(ctx context.Context, requestConfiguration *ExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation prepares an export of audit records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a *RequestInformation when successful
func (m *ExportRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ExportRequestBuilder when successful
func (m *ExportRequestBuilder) WithUrl(rawUrl string)(*ExportRequestBuilder) {
    return NewExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
