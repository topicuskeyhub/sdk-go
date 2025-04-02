package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i415e0cb9955fb6d146dae1cb4b4b6290c18ac7fb45e80eaf79a6f5778cd7917a "github.com/topicuskeyhub/sdk-go/account/item/group/item"
)

// ItemGroupWithGroupItemRequestBuilder builds and executes requests for operations under \account\{accountid}\group\{groupid}
type ItemGroupWithGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupWithGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupWithGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGroupWithGroupItemRequestBuilderGetQueryParameters returns the single group membership for the account. This also contains information such as the rights and provisioning information.
type ItemGroupWithGroupItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i415e0cb9955fb6d146dae1cb4b4b6290c18ac7fb45e80eaf79a6f5778cd7917a.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemGroupWithGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupWithGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGroupWithGroupItemRequestBuilderGetQueryParameters
}
// NewItemGroupWithGroupItemRequestBuilderInternal instantiates a new ItemGroupWithGroupItemRequestBuilder and sets the default values.
func NewItemGroupWithGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupWithGroupItemRequestBuilder) {
    m := &ItemGroupWithGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/group/{groupid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemGroupWithGroupItemRequestBuilder instantiates a new ItemGroupWithGroupItemRequestBuilder and sets the default values.
func NewItemGroupWithGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupWithGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupWithGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the specified group membership, effectively removing the account from the group.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemGroupWithGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGroupWithGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the single group membership for the account. This also contains information such as the rights and provisioning information.
// returns a GroupAccountGroupable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemGroupWithGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupWithGroupItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupAccountGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateGroupAccountGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.GroupAccountGroupable), nil
}
// Status the status property
// returns a *ItemGroupItemStatusRequestBuilder when successful
func (m *ItemGroupWithGroupItemRequestBuilder) Status()(*ItemGroupItemStatusRequestBuilder) {
    return NewItemGroupItemStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation removes the specified group membership, effectively removing the account from the group.
// returns a *RequestInformation when successful
func (m *ItemGroupWithGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGroupWithGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// ToGetRequestInformation returns the single group membership for the account. This also contains information such as the rights and provisioning information.
// returns a *RequestInformation when successful
func (m *ItemGroupWithGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGroupWithGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGroupWithGroupItemRequestBuilder when successful
func (m *ItemGroupWithGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemGroupWithGroupItemRequestBuilder) {
    return NewItemGroupWithGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
