package system

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i761ed4799fc1102dad078d6dc96accd5c28848aa30264ee692c85538984a4fbf "github.com/topicuskeyhub/sdk-go/system/item"
)

// WithSystemItemRequestBuilder builds and executes requests for operations under \system\{systemid}
type WithSystemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithSystemItemRequestBuilderGetQueryParameters returns the provisioned system identified by the id.
type WithSystemItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i761ed4799fc1102dad078d6dc96accd5c28848aa30264ee692c85538984a4fbf.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithSystemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithSystemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithSystemItemRequestBuilderGetQueryParameters
}
// WithSystemItemRequestBuilderPutQueryParameters updates the provisioned system identified by the id.
type WithSystemItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []i761ed4799fc1102dad078d6dc96accd5c28848aa30264ee692c85538984a4fbf.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithSystemItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithSystemItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithSystemItemRequestBuilderPutQueryParameters
}
// Account the account property
// returns a *ItemAccountRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) Account()(*ItemAccountRequestBuilder) {
    return NewItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithSystemItemRequestBuilderInternal instantiates a new WithSystemItemRequestBuilder and sets the default values.
func NewWithSystemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithSystemItemRequestBuilder) {
    m := &WithSystemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithSystemItemRequestBuilder instantiates a new WithSystemItemRequestBuilder and sets the default values.
func NewWithSystemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithSystemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithSystemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Deactivate the deactivate property
// returns a *ItemDeactivateRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) Deactivate()(*ItemDeactivateRequestBuilder) {
    return NewItemDeactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the provisioned system identified by the id.
// returns a ProvisioningProvisionedSystemable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithSystemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithSystemItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable), nil
}
// Group the group property
// returns a *ItemGroupRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) Group()(*ItemGroupRequestBuilder) {
    return NewItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the provisioned system identified by the id.
// returns a ProvisioningProvisionedSystemable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithSystemItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable, requestConfiguration *WithSystemItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable), nil
}
// Sync the sync property
// returns a *ItemSyncRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) Sync()(*ItemSyncRequestBuilder) {
    return NewItemSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Synclog the synclog property
// returns a *ItemSynclogRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) Synclog()(*ItemSynclogRequestBuilder) {
    return NewItemSynclogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns the provisioned system identified by the id.
// returns a *RequestInformation when successful
func (m *WithSystemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithSystemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// ToPutRequestInformation updates the provisioned system identified by the id.
// returns a *RequestInformation when successful
func (m *WithSystemItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemable, requestConfiguration *WithSystemItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=78", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithSystemItemRequestBuilder when successful
func (m *WithSystemItemRequestBuilder) WithUrl(rawUrl string)(*WithSystemItemRequestBuilder) {
    return NewWithSystemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
