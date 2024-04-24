package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i9de8e0b8ede20a2b36f229cd6abda61c462ebf686756469c4f0b1a809efc09d9 "github.com/topicuskeyhub/sdk-go/directory/item/internalaccount/item"
)

// ItemInternalaccountWithInternalaccountItemRequestBuilder builds and executes requests for operations under \directory\{directoryid}\internalaccount\{internalaccountid}
type ItemInternalaccountWithInternalaccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInternalaccountWithInternalaccountItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInternalaccountWithInternalaccountItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInternalaccountWithInternalaccountItemRequestBuilderGetQueryParameters returns the single internal accounts within a directory.
type ItemInternalaccountWithInternalaccountItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i9de8e0b8ede20a2b36f229cd6abda61c462ebf686756469c4f0b1a809efc09d9.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemInternalaccountWithInternalaccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInternalaccountWithInternalaccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInternalaccountWithInternalaccountItemRequestBuilderGetQueryParameters
}
// ItemInternalaccountWithInternalaccountItemRequestBuilderPutQueryParameters updates the internal accounts within a directory identified by the id.
type ItemInternalaccountWithInternalaccountItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []i9de8e0b8ede20a2b36f229cd6abda61c462ebf686756469c4f0b1a809efc09d9.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemInternalaccountWithInternalaccountItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInternalaccountWithInternalaccountItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInternalaccountWithInternalaccountItemRequestBuilderPutQueryParameters
}
// NewItemInternalaccountWithInternalaccountItemRequestBuilderInternal instantiates a new ItemInternalaccountWithInternalaccountItemRequestBuilder and sets the default values.
func NewItemInternalaccountWithInternalaccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInternalaccountWithInternalaccountItemRequestBuilder) {
    m := &ItemInternalaccountWithInternalaccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/{directoryid}/internalaccount/{internalaccountid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemInternalaccountWithInternalaccountItemRequestBuilder instantiates a new ItemInternalaccountWithInternalaccountItemRequestBuilder and sets the default values.
func NewItemInternalaccountWithInternalaccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInternalaccountWithInternalaccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInternalaccountWithInternalaccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the specified internal accounts from a directory, including the associated KeyHub account.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the single internal accounts within a directory.
// returns a AuthInternalAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthInternalAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable), nil
}
// Move the move property
// returns a *ItemInternalaccountItemMoveRequestBuilder when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) Move()(*ItemInternalaccountItemMoveRequestBuilder) {
    return NewItemInternalaccountItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the internal accounts within a directory identified by the id.
// returns a AuthInternalAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthInternalAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable), nil
}
// Status the status property
// returns a *ItemInternalaccountItemStatusRequestBuilder when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) Status()(*ItemInternalaccountItemStatusRequestBuilder) {
    return NewItemInternalaccountItemStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation removes the specified internal accounts from a directory, including the associated KeyHub account.
// returns a *RequestInformation when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/directory/{directoryid}/internalaccount/{internalaccountid}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    return requestInfo, nil
}
// ToGetRequestInformation returns the single internal accounts within a directory.
// returns a *RequestInformation when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    return requestInfo, nil
}
// ToPutRequestInformation updates the internal accounts within a directory identified by the id.
// returns a *RequestInformation when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountable, requestConfiguration *ItemInternalaccountWithInternalaccountItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=70", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInternalaccountWithInternalaccountItemRequestBuilder when successful
func (m *ItemInternalaccountWithInternalaccountItemRequestBuilder) WithUrl(rawUrl string)(*ItemInternalaccountWithInternalaccountItemRequestBuilder) {
    return NewItemInternalaccountWithInternalaccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
