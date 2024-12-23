package certificate

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    iea9c82f767f3e4dc2d923fd086941fbc5112feb51486be0a7e2c29c8ddd3bb17 "github.com/topicuskeyhub/sdk-go/certificate/item"
)

// WithCertificateItemRequestBuilder builds and executes requests for operations under \certificate\{certificateid}
type WithCertificateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithCertificateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCertificateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithCertificateItemRequestBuilderGetQueryParameters returns the global trusted certificate identified by the id.
type WithCertificateItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []iea9c82f767f3e4dc2d923fd086941fbc5112feb51486be0a7e2c29c8ddd3bb17.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithCertificateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCertificateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithCertificateItemRequestBuilderGetQueryParameters
}
// WithCertificateItemRequestBuilderPutQueryParameters updates the global trusted certificate identified by the id.
type WithCertificateItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []iea9c82f767f3e4dc2d923fd086941fbc5112feb51486be0a7e2c29c8ddd3bb17.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithCertificateItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCertificateItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithCertificateItemRequestBuilderPutQueryParameters
}
// NewWithCertificateItemRequestBuilderInternal instantiates a new WithCertificateItemRequestBuilder and sets the default values.
func NewWithCertificateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCertificateItemRequestBuilder) {
    m := &WithCertificateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificate/{certificateid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithCertificateItemRequestBuilder instantiates a new WithCertificateItemRequestBuilder and sets the default values.
func NewWithCertificateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCertificateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithCertificateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete permanently removes the global trusted certificate identified by the id. This cannot be undone.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithCertificateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithCertificateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get returns the global trusted certificate identified by the id.
// returns a CertificateCertificateable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithCertificateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithCertificateItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateCertificateCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable), nil
}
// Put updates the global trusted certificate identified by the id.
// returns a CertificateCertificateable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithCertificateItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable, requestConfiguration *WithCertificateItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateCertificateCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable), nil
}
// ToDeleteRequestInformation permanently removes the global trusted certificate identified by the id. This cannot be undone.
// returns a *RequestInformation when successful
func (m *WithCertificateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithCertificateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=75")
    return requestInfo, nil
}
// ToGetRequestInformation returns the global trusted certificate identified by the id.
// returns a *RequestInformation when successful
func (m *WithCertificateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithCertificateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the global trusted certificate identified by the id.
// returns a *RequestInformation when successful
func (m *WithCertificateItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CertificateCertificateable, requestConfiguration *WithCertificateItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=75")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=75", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithCertificateItemRequestBuilder when successful
func (m *WithCertificateItemRequestBuilder) WithUrl(rawUrl string)(*WithCertificateItemRequestBuilder) {
    return NewWithCertificateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
