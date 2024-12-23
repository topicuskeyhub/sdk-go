package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// SessionWithSessionItemRequestBuilder builds and executes requests for operations under \account\session\{sessionId}
type SessionWithSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SessionWithSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SessionWithSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSessionWithSessionItemRequestBuilderInternal instantiates a new SessionWithSessionItemRequestBuilder and sets the default values.
func NewSessionWithSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SessionWithSessionItemRequestBuilder) {
    m := &SessionWithSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/session/{sessionId}", pathParameters),
    }
    return m
}
// NewSessionWithSessionItemRequestBuilder instantiates a new SessionWithSessionItemRequestBuilder and sets the default values.
func NewSessionWithSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SessionWithSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSessionWithSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete ends the given session for the current user. Access tokens issued for the session will be revoked.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *SessionWithSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SessionWithSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation ends the given session for the current user. Access tokens issued for the session will be revoked.
// returns a *RequestInformation when successful
func (m *SessionWithSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SessionWithSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=75")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SessionWithSessionItemRequestBuilder when successful
func (m *SessionWithSessionItemRequestBuilder) WithUrl(rawUrl string)(*SessionWithSessionItemRequestBuilder) {
    return NewSessionWithSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
