package accessprofileclient

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i675f71201e095fa7e14d58b140b98e958cafab227e2ec1db39dd1902dbca0b04 "github.com/topicuskeyhub/sdk-go/accessprofileclient/item"
)

// WithAccessprofileclientItemRequestBuilder builds and executes requests for operations under \accessprofileclient\{accessprofileclientid}
type WithAccessprofileclientItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithAccessprofileclientItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithAccessprofileclientItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithAccessprofileclientItemRequestBuilderGetQueryParameters returns the single client link for the access profile.
type WithAccessprofileclientItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i675f71201e095fa7e14d58b140b98e958cafab227e2ec1db39dd1902dbca0b04.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithAccessprofileclientItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithAccessprofileclientItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithAccessprofileclientItemRequestBuilderGetQueryParameters
}
// NewWithAccessprofileclientItemRequestBuilderInternal instantiates a new WithAccessprofileclientItemRequestBuilder and sets the default values.
func NewWithAccessprofileclientItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccessprofileclientItemRequestBuilder) {
    m := &WithAccessprofileclientItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessprofileclient/{accessprofileclientid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithAccessprofileclientItemRequestBuilder instantiates a new WithAccessprofileclientItemRequestBuilder and sets the default values.
func NewWithAccessprofileclientItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccessprofileclientItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithAccessprofileclientItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the specified client link, effectively removing access to the client for members of the access profile.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithAccessprofileclientItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithAccessprofileclientItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the single client link for the access profile.
// returns a ProfileAccessProfileClientable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithAccessprofileclientItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithAccessprofileclientItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileClientable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProfileAccessProfileClientFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProfileAccessProfileClientable), nil
}
// ToDeleteRequestInformation removes the specified client link, effectively removing access to the client for members of the access profile.
// returns a *RequestInformation when successful
func (m *WithAccessprofileclientItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithAccessprofileclientItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// ToGetRequestInformation returns the single client link for the access profile.
// returns a *RequestInformation when successful
func (m *WithAccessprofileclientItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithAccessprofileclientItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WithAccessprofileclientItemRequestBuilder when successful
func (m *WithAccessprofileclientItemRequestBuilder) WithUrl(rawUrl string)(*WithAccessprofileclientItemRequestBuilder) {
    return NewWithAccessprofileclientItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
