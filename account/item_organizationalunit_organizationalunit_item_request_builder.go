package account

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemOrganizationalunitOrganizationalunitItemRequestBuilder builds and executes requests for operations under \account\{accountid}\organizationalunit\{id}
type ItemOrganizationalunitOrganizationalunitItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetQueryParameters returns the single organizational unit membership for the account.
type ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
}
// ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetQueryParameters
}
// NewItemOrganizationalunitOrganizationalunitItemRequestBuilderInternal instantiates a new OrganizationalunitItemRequestBuilder and sets the default values.
func NewItemOrganizationalunitOrganizationalunitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitOrganizationalunitItemRequestBuilder) {
    m := &ItemOrganizationalunitOrganizationalunitItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/{accountid}/organizationalunit/{id}{?additional*}", pathParameters),
    }
    return m
}
// NewItemOrganizationalunitOrganizationalunitItemRequestBuilder instantiates a new OrganizationalunitItemRequestBuilder and sets the default values.
func NewItemOrganizationalunitOrganizationalunitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationalunitOrganizationalunitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationalunitOrganizationalunitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the single organizational unit membership for the account.
func (m *ItemOrganizationalunitOrganizationalunitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationAccountOrganizationalUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
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
func (m *ItemOrganizationalunitOrganizationalunitItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationalunitOrganizationalunitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=65")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
