package profile

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAttributesRequestBuilder builds and executes requests for operations under \profile\{accessprofile-id}\attributes
type ItemAttributesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAttributesRequestBuilderInternal instantiates a new ItemAttributesRequestBuilder and sets the default values.
func NewItemAttributesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributesRequestBuilder) {
    m := &ItemAttributesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/profile/{accessprofile%2Did}/attributes", pathParameters),
    }
    return m
}
// NewItemAttributesRequestBuilder instantiates a new ItemAttributesRequestBuilder and sets the default values.
func NewItemAttributesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttributesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttributesRequestBuilderInternal(urlParams, requestAdapter)
}
// Sync the sync property
// returns a *ItemAttributesSyncRequestBuilder when successful
func (m *ItemAttributesRequestBuilder) Sync()(*ItemAttributesSyncRequestBuilder) {
    return NewItemAttributesSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
