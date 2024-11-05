package certificate

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// CertificateRequestBuilder builds and executes requests for operations under \certificate
type CertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateRequestBuilderGetQueryParameters query for all global trusted certificates in Topicus KeyHub. The various query parameters can be used to filter the response.
type CertificateRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Filter certificates on the exact alias.
    Alias []string `uriparametername:"alias"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return certificates that are expired at the given instant.
    Expired []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"expired"`
    // Always true, even if not specified.
    Global []bool `uriparametername:"global"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// CertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateRequestBuilderGetQueryParameters
}
// CertificateRequestBuilderPostQueryParameters creates one or more new global trusted certificates and returns the newly created certificates.
type CertificateRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// CertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateRequestBuilderPostQueryParameters
}
// ByCertificateid gets an item from the github.com/topicuskeyhub/sdk-go.certificate.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithCertificateItemRequestBuilder when successful
func (m *CertificateRequestBuilder) ByCertificateid(certificateid string)(*WithCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if certificateid != "" {
        urlTplParams["certificateid"] = certificateid
    }
    return NewWithCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCertificateidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.certificate.item collection
// returns a *WithCertificateItemRequestBuilder when successful
func (m *CertificateRequestBuilder) ByCertificateidInt64(certificateid int64)(*WithCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["certificateid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(certificateid, 10)
    return NewWithCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCertificateRequestBuilderInternal instantiates a new CertificateRequestBuilder and sets the default values.
func NewCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateRequestBuilder) {
    m := &CertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificate{?additional*,alias*,any*,createdAfter*,createdBefore*,exclude*,expired*,global*,id*,modifiedSince*,q*,sort*,uuid*}", pathParameters),
    }
    return m
}
// NewCertificateRequestBuilder instantiates a new CertificateRequestBuilder and sets the default values.
func NewCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all global trusted certificates in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a CertificateCertificateLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *CertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateCertificateCertificateLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable), nil
}
// Post creates one or more new global trusted certificates and returns the newly created certificates.
// returns a CertificateCertificateLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *CertificateRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable, requestConfiguration *CertificateRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateCertificateCertificateLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable), nil
}
// ToGetRequestInformation query for all global trusted certificates in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *CertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new global trusted certificates and returns the newly created certificates.
// returns a *RequestInformation when successful
func (m *CertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateLinkableWrapperable, requestConfiguration *CertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=74", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CertificateRequestBuilder when successful
func (m *CertificateRequestBuilder) WithUrl(rawUrl string)(*CertificateRequestBuilder) {
    return NewCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
