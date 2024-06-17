package client

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i80f198c78d6aecba6f4d022fc637d5f703bf68a2f46aaa5a67f160996e26d1a0 "github.com/topicuskeyhub/sdk-go/client/item/permission/item"
)

// ItemPermissionWithPermissionItemRequestBuilder builds and executes requests for operations under \client\{clientid}\permission\{permissionid}
type ItemPermissionWithPermissionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPermissionWithPermissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionWithPermissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPermissionWithPermissionItemRequestBuilderGetQueryParameters returns the permission identified by the id.
type ItemPermissionWithPermissionItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i80f198c78d6aecba6f4d022fc637d5f703bf68a2f46aaa5a67f160996e26d1a0.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemPermissionWithPermissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionWithPermissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPermissionWithPermissionItemRequestBuilderGetQueryParameters
}
// NewItemPermissionWithPermissionItemRequestBuilderInternal instantiates a new ItemPermissionWithPermissionItemRequestBuilder and sets the default values.
func NewItemPermissionWithPermissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionWithPermissionItemRequestBuilder) {
    m := &ItemPermissionWithPermissionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/{clientid}/permission/{permissionid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemPermissionWithPermissionItemRequestBuilder instantiates a new ItemPermissionWithPermissionItemRequestBuilder and sets the default values.
func NewItemPermissionWithPermissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionWithPermissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPermissionWithPermissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the permission identified by the id.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemPermissionWithPermissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPermissionWithPermissionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the permission identified by the id.
// returns a ClientOAuth2ClientPermissionable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemPermissionWithPermissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPermissionWithPermissionItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateClientOAuth2ClientPermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ClientOAuth2ClientPermissionable), nil
}
// ToDeleteRequestInformation deletes the permission identified by the id.
// returns a *RequestInformation when successful
func (m *ItemPermissionWithPermissionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionWithPermissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// ToGetRequestInformation returns the permission identified by the id.
// returns a *RequestInformation when successful
func (m *ItemPermissionWithPermissionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionWithPermissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPermissionWithPermissionItemRequestBuilder when successful
func (m *ItemPermissionWithPermissionItemRequestBuilder) WithUrl(rawUrl string)(*ItemPermissionWithPermissionItemRequestBuilder) {
    return NewItemPermissionWithPermissionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
