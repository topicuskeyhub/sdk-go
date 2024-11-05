package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    ibd6697d4627ce6bb8b76244c42dd0126edb76adae9eeda114b25174002278b6d "github.com/topicuskeyhub/sdk-go/account/item/organizationalunit/item"
)

// ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder builds and executes requests for operations under \account\{accountid}\organizationalunit\{organizationalunitid}
type ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetQueryParameters returns the single organizational unit membership for the account.
type ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []ibd6697d4627ce6bb8b76244c42dd0126edb76adae9eeda114b25174002278b6d.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetQueryParameters
}
// NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilderInternal instantiates a new ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder and sets the default values.
func NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) {
    m := &ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/organizationalunit/{organizationalunitid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilder instantiates a new ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder and sets the default values.
func NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the single organizational unit membership for the account.
// returns a OrganizationAccountOrganizationalUnitable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationAccountOrganizationalUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateOrganizationAccountOrganizationalUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationAccountOrganizationalUnitable), nil
}
// ToGetRequestInformation returns the single organizational unit membership for the account.
// returns a *RequestInformation when successful
func (m *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=74")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder when successful
func (m *ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationalunitWithOrganizationalunitItemRequestBuilder) {
    return NewItemOrganizationalunitWithOrganizationalunitItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
