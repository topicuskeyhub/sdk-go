package launchpadtile

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// DisplayRequestBuilder builds and executes requests for operations under \launchpadtile\display
type DisplayRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DisplayRequestBuilderGetQueryParameters query tiles for display on the launchpad. This returns a normalized view over the tiles, per group. A single tile can be returned multiple times if access is given via multiple groups.
type DisplayRequestBuilderGetQueryParameters struct {
    Filter *string `uriparametername:"filter"`
    Group *int64 `uriparametername:"group"`
    ModifiedSince *string `uriparametername:"modifiedSince"`
}
// DisplayRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DisplayRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DisplayRequestBuilderGetQueryParameters
}
// NewDisplayRequestBuilderInternal instantiates a new DisplayRequestBuilder and sets the default values.
func NewDisplayRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DisplayRequestBuilder) {
    m := &DisplayRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/launchpadtile/display{?filter*,group*,modifiedSince*}", pathParameters),
    }
    return m
}
// NewDisplayRequestBuilder instantiates a new DisplayRequestBuilder and sets the default values.
func NewDisplayRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DisplayRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDisplayRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query tiles for display on the launchpad. This returns a normalized view over the tiles, per group. A single tile can be returned multiple times if access is given via multiple groups.
// returns a LaunchpadDisplayedLaunchpadTilesable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *DisplayRequestBuilder) Get(ctx context.Context, requestConfiguration *DisplayRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.LaunchpadDisplayedLaunchpadTilesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateLaunchpadDisplayedLaunchpadTilesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.LaunchpadDisplayedLaunchpadTilesable), nil
}
// ToGetRequestInformation query tiles for display on the launchpad. This returns a normalized view over the tiles, per group. A single tile can be returned multiple times if access is given via multiple groups.
// returns a *RequestInformation when successful
func (m *DisplayRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DisplayRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=73")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DisplayRequestBuilder when successful
func (m *DisplayRequestBuilder) WithUrl(rawUrl string)(*DisplayRequestBuilder) {
    return NewDisplayRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
