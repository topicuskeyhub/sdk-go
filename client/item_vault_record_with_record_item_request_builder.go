package client

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    ie098eec106de559f1ed4075ed2f377f52f0742f4b8f92e1f2a7dcbbef2274eb5 "github.com/topicuskeyhub/sdk-go/client/item/vault/record/item"
)

// ItemVaultRecordWithRecordItemRequestBuilder builds and executes requests for operations under \client\{clientid}\vault\record\{recordid}
type ItemVaultRecordWithRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVaultRecordWithRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordWithRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemVaultRecordWithRecordItemRequestBuilderGetQueryParameters returns the vault record identified by the id. When the 'secret' additonal object is requested, the 'topicus-Vault-session' header must be specified.
type ItemVaultRecordWithRecordItemRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []ie098eec106de559f1ed4075ed2f377f52f0742f4b8f92e1f2a7dcbbef2274eb5.GetAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultRecordWithRecordItemRequestBuilderGetQueryParameters
}
// ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters updates the vault record identified by the id. To update the secrets, the 'secret' additional object must be used, in addition to specifying the 'topicus-Vault-session' header. When updating a TOTP-secret, make sure to set 'writeTotp' field.
type ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPutAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPutAdditionalQueryParameterType []ie098eec106de559f1ed4075ed2f377f52f0742f4b8f92e1f2a7dcbbef2274eb5.PutAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters
}
// NewItemVaultRecordWithRecordItemRequestBuilderInternal instantiates a new ItemVaultRecordWithRecordItemRequestBuilder and sets the default values.
func NewItemVaultRecordWithRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordWithRecordItemRequestBuilder) {
    m := &ItemVaultRecordWithRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/client/{clientid}/vault/record/{recordid}{?additional*}", pathParameters),
    }
    return m
}
// NewItemVaultRecordWithRecordItemRequestBuilder instantiates a new ItemVaultRecordWithRecordItemRequestBuilder and sets the default values.
func NewItemVaultRecordWithRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordWithRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVaultRecordWithRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the vault record identified by the id. This cannot be undone.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemVaultRecordWithRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns the vault record identified by the id. When the 'secret' additonal object is requested, the 'topicus-Vault-session' header must be specified.
// returns a VaultVaultRecordable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemVaultRecordWithRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultVaultRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable), nil
}
// Move the move property
// returns a *ItemVaultRecordItemMoveRequestBuilder when successful
func (m *ItemVaultRecordWithRecordItemRequestBuilder) Move()(*ItemVaultRecordItemMoveRequestBuilder) {
    return NewItemVaultRecordItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put updates the vault record identified by the id. To update the secrets, the 'secret' additional object must be used, in addition to specifying the 'topicus-Vault-session' header. When updating a TOTP-secret, make sure to set 'writeTotp' field.
// returns a VaultVaultRecordable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemVaultRecordWithRecordItemRequestBuilder) Put(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultVaultRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable), nil
}
// ToDeleteRequestInformation removes the vault record identified by the id. This cannot be undone.
// returns a *RequestInformation when successful
func (m *ItemVaultRecordWithRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=78")
    return requestInfo, nil
}
// ToGetRequestInformation returns the vault record identified by the id. When the 'secret' additonal object is requested, the 'topicus-Vault-session' header must be specified.
// returns a *RequestInformation when successful
func (m *ItemVaultRecordWithRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the vault record identified by the id. To update the secrets, the 'secret' additional object must be used, in addition to specifying the 'topicus-Vault-session' header. When updating a TOTP-secret, make sure to set 'writeTotp' field.
// returns a *RequestInformation when successful
func (m *ItemVaultRecordWithRecordItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordable, requestConfiguration *ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemVaultRecordWithRecordItemRequestBuilder when successful
func (m *ItemVaultRecordWithRecordItemRequestBuilder) WithUrl(rawUrl string)(*ItemVaultRecordWithRecordItemRequestBuilder) {
    return NewItemVaultRecordWithRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
