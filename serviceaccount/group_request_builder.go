package serviceaccount

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GroupRequestBuilder builds and executes requests for operations under \serviceaccount\group
type GroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount/group", pathParameters),
    }
    return m
}
// NewGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
func (m *GroupRequestBuilder) Export()(*GroupExportRequestBuilder) {
    return NewGroupExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
