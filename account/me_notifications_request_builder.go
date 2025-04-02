package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// MeNotificationsRequestBuilder builds and executes requests for operations under \account\me\notifications
type MeNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeNotificationsRequestBuilderGetQueryParameters returns the notifications to be displayed on the dashboard of the user. At most 50 notifications are returned.
type MeNotificationsRequestBuilderGetQueryParameters struct {
    Lang *string `uriparametername:"lang"`
    RequestMailKey *string `uriparametername:"requestMailKey"`
    RequestsToRetainIds []int64 `uriparametername:"requestsToRetainIds"`
}
// MeNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeNotificationsRequestBuilderGetQueryParameters
}
// NewMeNotificationsRequestBuilderInternal instantiates a new MeNotificationsRequestBuilder and sets the default values.
func NewMeNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeNotificationsRequestBuilder) {
    m := &MeNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/me/notifications{?lang*,requestMailKey*,requestsToRetainIds*}", pathParameters),
    }
    return m
}
// NewMeNotificationsRequestBuilder instantiates a new MeNotificationsRequestBuilder and sets the default values.
func NewMeNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the notifications to be displayed on the dashboard of the user. At most 50 notifications are returned.
// returns a NotificationNotificationsable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *MeNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *MeNotificationsRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.NotificationNotificationsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateNotificationNotificationsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.NotificationNotificationsable), nil
}
// ToGetRequestInformation returns the notifications to be displayed on the dashboard of the user. At most 50 notifications are returned.
// returns a *RequestInformation when successful
func (m *MeNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MeNotificationsRequestBuilder when successful
func (m *MeNotificationsRequestBuilder) WithUrl(rawUrl string)(*MeNotificationsRequestBuilder) {
    return NewMeNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
