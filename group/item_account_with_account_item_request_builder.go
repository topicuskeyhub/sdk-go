package group

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    icbcff0fbb4a7494c430c35e03d2bfa34fee5927d258ab696a6147683554f81f1 "github.com/topicuskeyhub/sdk-go/group/item/account/item"
)

// ItemAccountWithAccountItemRequestBuilder builds and executes requests for operations under \group\{groupid}\account\{accountid}
type ItemAccountWithAccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAccountWithAccountItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountWithAccountItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAccountWithAccountItemRequestBuilderGetQueryParameters returns the single membership for the group. This also contains information such as the rights and provisioning information.
type ItemAccountWithAccountItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []icbcff0fbb4a7494c430c35e03d2bfa34fee5927d258ab696a6147683554f81f1.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAccountWithAccountItemRequestBuilderGetQueryParameters
}
// ItemAccountWithAccountItemRequestBuilderPutQueryParameters updates the memberships for a group identified by the id.
type ItemAccountWithAccountItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []icbcff0fbb4a7494c430c35e03d2bfa34fee5927d258ab696a6147683554f81f1.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemAccountWithAccountItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAccountWithAccountItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAccountWithAccountItemRequestBuilderPutQueryParameters
}
// NewItemAccountWithAccountItemRequestBuilderInternal instantiates a new ItemAccountWithAccountItemRequestBuilder and sets the default values.
func NewItemAccountWithAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountWithAccountItemRequestBuilder) {
    m := &ItemAccountWithAccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}/account/{accountid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemAccountWithAccountItemRequestBuilder instantiates a new ItemAccountWithAccountItemRequestBuilder and sets the default values.
func NewItemAccountWithAccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAccountWithAccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAccountWithAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the specified membership, effectively removing the account from the group.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountWithAccountItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the single membership for the group. This also contains information such as the rights and provisioning information.
// returns a GroupGroupAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountWithAccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable), nil
}
// Nesting the nesting property
// returns a *ItemAccountItemNestingRequestBuilder when successful
func (m *ItemAccountWithAccountItemRequestBuilder) Nesting()(*ItemAccountItemNestingRequestBuilder) {
    return NewItemAccountItemNestingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the memberships for a group identified by the id.
// returns a GroupGroupAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemAccountWithAccountItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable, requestConfiguration *ItemAccountWithAccountItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupGroupAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable), nil
}
// ToDeleteRequestInformation removes the specified membership, effectively removing the account from the group.
// returns a *RequestInformation when successful
func (m *ItemAccountWithAccountItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// ToGetRequestInformation returns the single membership for the group. This also contains information such as the rights and provisioning information.
// returns a *RequestInformation when successful
func (m *ItemAccountWithAccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAccountWithAccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the memberships for a group identified by the id.
// returns a *RequestInformation when successful
func (m *ItemAccountWithAccountItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupGroupAccountable, requestConfiguration *ItemAccountWithAccountItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAccountWithAccountItemRequestBuilder when successful
func (m *ItemAccountWithAccountItemRequestBuilder) WithUrl(rawUrl string)(*ItemAccountWithAccountItemRequestBuilder) {
    return NewItemAccountWithAccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
