package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i9fc7d9da036a6eb05de51f913bd66283fecb6f679f8d7b90a69fac4544978dc6 "github.com/topicuskeyhub/sdk-go/account/item"
)

// WithAccountItemRequestBuilder builds and executes requests for operations under \account\{accountid}
type WithAccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithAccountItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithAccountItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithAccountItemRequestBuilderGetQueryParameters returns the account identified by the id.
type WithAccountItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i9fc7d9da036a6eb05de51f913bd66283fecb6f679f8d7b90a69fac4544978dc6.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// WithAccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithAccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithAccountItemRequestBuilderGetQueryParameters
}
// Attributevalue the attributevalue property
// returns a *ItemAttributevalueRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Attributevalue()(*ItemAttributevalueRequestBuilder) {
    return NewItemAttributevalueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithAccountItemRequestBuilderInternal instantiates a new WithAccountItemRequestBuilder and sets the default values.
func NewWithAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccountItemRequestBuilder) {
    m := &WithAccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}{?additional*}", pathParameters),
    }
    return m
}
// NewWithAccountItemRequestBuilder instantiates a new WithAccountItemRequestBuilder and sets the default values.
func NewWithAccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete permanently removes the account. This cannot be undone.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithAccountItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithAccountItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the account identified by the id.
// returns a AuthAccountable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *WithAccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithAccountItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthAccountable), nil
}
// Group the group property
// returns a *ItemGroupRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Group()(*ItemGroupRequestBuilder) {
    return NewItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Identity the identity property
// returns a *ItemIdentityRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Identity()(*ItemIdentityRequestBuilder) {
    return NewItemIdentityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizationalunit the organizationalunit property
// returns a *ItemOrganizationalunitRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Organizationalunit()(*ItemOrganizationalunitRequestBuilder) {
    return NewItemOrganizationalunitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Status the status property
// returns a *ItemStatusRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Status()(*ItemStatusRequestBuilder) {
    return NewItemStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation permanently removes the account. This cannot be undone.
// returns a *RequestInformation when successful
func (m *WithAccountItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithAccountItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=77")
    return requestInfo, nil
}
// ToGetRequestInformation returns the account identified by the id.
// returns a *RequestInformation when successful
func (m *WithAccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithAccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Validate the validate property
// returns a *ItemValidateRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Validate()(*ItemValidateRequestBuilder) {
    return NewItemValidateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vault the vault property
// returns a *ItemVaultRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) Vault()(*ItemVaultRequestBuilder) {
    return NewItemVaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithAccountItemRequestBuilder when successful
func (m *WithAccountItemRequestBuilder) WithUrl(rawUrl string)(*WithAccountItemRequestBuilder) {
    return NewWithAccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
