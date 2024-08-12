package profileprovisioning

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i618a27d18e51be04be77ddb70414e3922299f7d6146f157d537212a1f9ab3bf1 "github.com/topicuskeyhub/sdk-go/profileprovisioning/item"
)

// WithProfileprovisioningItemRequestBuilder builds and executes requests for operations under \profileprovisioning\{profileprovisioningid}
type WithProfileprovisioningItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithProfileprovisioningItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithProfileprovisioningItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithProfileprovisioningItemRequestBuilderGetQueryParameters returns the access profile provisioning link identified by the id.
type WithProfileprovisioningItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i618a27d18e51be04be77ddb70414e3922299f7d6146f157d537212a1f9ab3bf1.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithProfileprovisioningItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithProfileprovisioningItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithProfileprovisioningItemRequestBuilderGetQueryParameters
}
// NewWithProfileprovisioningItemRequestBuilderInternal instantiates a new WithProfileprovisioningItemRequestBuilder and sets the default values.
func NewWithProfileprovisioningItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProfileprovisioningItemRequestBuilder) {
    m := &WithProfileprovisioningItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profileprovisioning/{profileprovisioningid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithProfileprovisioningItemRequestBuilder instantiates a new WithProfileprovisioningItemRequestBuilder and sets the default values.
func NewWithProfileprovisioningItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProfileprovisioningItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithProfileprovisioningItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the access profile provisioning link identified by the id.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithProfileprovisioningItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithProfileprovisioningItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the access profile provisioning link identified by the id.
// returns a ProfileAccessProfileProvisioningable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithProfileprovisioningItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithProfileprovisioningItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileProvisioningable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProfileAccessProfileProvisioningFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileProvisioningable), nil
}
// ToDeleteRequestInformation deletes the access profile provisioning link identified by the id.
// returns a *RequestInformation when successful
func (m *WithProfileprovisioningItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithProfileprovisioningItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=72")
    return requestInfo, nil
}
// ToGetRequestInformation returns the access profile provisioning link identified by the id.
// returns a *RequestInformation when successful
func (m *WithProfileprovisioningItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithProfileprovisioningItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=72")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithProfileprovisioningItemRequestBuilder when successful
func (m *WithProfileprovisioningItemRequestBuilder) WithUrl(rawUrl string)(*WithProfileprovisioningItemRequestBuilder) {
    return NewWithProfileprovisioningItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
