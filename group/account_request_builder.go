package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AccountRequestBuilder builds and executes requests for operations under \group\account
type AccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAccountRequestBuilderInternal instantiates a new AccountRequestBuilder and sets the default values.
func NewAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountRequestBuilder) {
    m := &AccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/account", pathParameters),
    }
    return m
}
// NewAccountRequestBuilder instantiates a new AccountRequestBuilder and sets the default values.
func NewAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *AccountRequestBuilder) Export()(*AccountExportRequestBuilder) {
    return NewAccountExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
