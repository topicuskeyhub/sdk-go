package client

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i585bf5187d12ada5efb2d7ce82e83e52ab243f26acb1346d550881470764671d "github.com/topicuskeyhub/sdk-go/client/item"
)

// WithClientItemRequestBuilder builds and executes requests for operations under \client\{clientid}
type WithClientItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithClientItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithClientItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithClientItemRequestBuilderGetQueryParameters returns the client identified by the id.
type WithClientItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i585bf5187d12ada5efb2d7ce82e83e52ab243f26acb1346d550881470764671d.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithClientItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithClientItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithClientItemRequestBuilderGetQueryParameters
}
// WithClientItemRequestBuilderPutQueryParameters updates the client identified by the id.
type WithClientItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []i585bf5187d12ada5efb2d7ce82e83e52ab243f26acb1346d550881470764671d.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithClientItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithClientItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithClientItemRequestBuilderPutQueryParameters
}
// NewWithClientItemRequestBuilderInternal instantiates a new WithClientItemRequestBuilder and sets the default values.
func NewWithClientItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClientItemRequestBuilder) {
    m := &WithClientItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/{clientid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithClientItemRequestBuilder instantiates a new WithClientItemRequestBuilder and sets the default values.
func NewWithClientItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClientItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithClientItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete permanently removes the client identified by the id. This cannot be undone.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithClientItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithClientItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the client identified by the id.
// returns a ClientClientApplicationable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithClientItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithClientItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientClientApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable), nil
}
// Permission the permission property
// returns a *ItemPermissionRequestBuilder when successful
func (m *WithClientItemRequestBuilder) Permission()(*ItemPermissionRequestBuilder) {
    return NewItemPermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the client identified by the id.
// returns a ClientClientApplicationable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithClientItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable, requestConfiguration *WithClientItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientClientApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable), nil
}
// ToDeleteRequestInformation permanently removes the client identified by the id. This cannot be undone.
// returns a *RequestInformation when successful
func (m *WithClientItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithClientItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    return requestInfo, nil
}
// ToGetRequestInformation returns the client identified by the id.
// returns a *RequestInformation when successful
func (m *WithClientItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithClientItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the client identified by the id.
// returns a *RequestInformation when successful
func (m *WithClientItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientClientApplicationable, requestConfiguration *WithClientItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Vault the vault property
// returns a *ItemVaultRequestBuilder when successful
func (m *WithClientItemRequestBuilder) Vault()(*ItemVaultRequestBuilder) {
    return NewItemVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithClientItemRequestBuilder when successful
func (m *WithClientItemRequestBuilder) WithUrl(rawUrl string)(*WithClientItemRequestBuilder) {
    return NewWithClientItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
