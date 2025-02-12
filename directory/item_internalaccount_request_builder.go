package directory

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
    i71c97f9512df5822f155a22bf755dad1f40cdfff8dd134c193ec96e5c2241965 "github.com/topicuskeyhub/sdk-go/directory/item/internalaccount"
)

// ItemInternalaccountRequestBuilder builds and executes requests for operations under \directory\{directoryid}\internalaccount
type ItemInternalaccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInternalaccountRequestBuilderGetQueryParameters queries over all internal accounts within a directory. The various query parameters can be used to filter the response.
type ItemInternalaccountRequestBuilderGetQueryParameters struct {
    // Filter the internal accounts by KeyHub accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64 `uriparametername:"account"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []i71c97f9512df5822f155a22bf755dad1f40cdfff8dd134c193ec96e5c2241965.GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the internal accounts by directory, specified by id. This parameter supports composition with all parameters from the account directory resource.
    Directory []int64 `uriparametername:"directory"`
    // Filter the internal accounts by the owner group of the directory, specified by id. This parameter supports composition with all parameters from the account resource.
    DirectoryOwnerGroup []int64 `uriparametername:"directoryOwnerGroup"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter the internal accounts by organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnit []int64 `uriparametername:"organizationalUnit"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Only return internal accounts with the given status.
    // Deprecated: This property is deprecated, use StatusAsAuthInternalAccountStatus instead
    Status []string `uriparametername:"status"`
    // Only return internal accounts with the given status.
    StatusAsAuthInternalAccountStatus []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountStatus `uriparametername:"status"`
}
// ItemInternalaccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInternalaccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInternalaccountRequestBuilderGetQueryParameters
}
// ItemInternalaccountRequestBuilderPostQueryParameters creates one or more new internal accounts within a directory and returns the newly created accounts.
type ItemInternalaccountRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []i71c97f9512df5822f155a22bf755dad1f40cdfff8dd134c193ec96e5c2241965.PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ItemInternalaccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInternalaccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInternalaccountRequestBuilderPostQueryParameters
}
// ByInternalaccountid gets an item from the github.com/topicuskeyhub/sdk-go.directory.item.internalaccount.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemInternalaccountWithInternalaccountItemRequestBuilder when successful
func (m *ItemInternalaccountRequestBuilder) ByInternalaccountid(internalaccountid string)(*ItemInternalaccountWithInternalaccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if internalaccountid != "" {
        urlTplParams["internalaccountid"] = internalaccountid
    }
    return NewItemInternalaccountWithInternalaccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByInternalaccountidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.directory.item.internalaccount.item collection
// returns a *ItemInternalaccountWithInternalaccountItemRequestBuilder when successful
func (m *ItemInternalaccountRequestBuilder) ByInternalaccountidInt64(internalaccountid int64)(*ItemInternalaccountWithInternalaccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["internalaccountid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(internalaccountid, 10)
    return NewItemInternalaccountWithInternalaccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInternalaccountRequestBuilderInternal instantiates a new ItemInternalaccountRequestBuilder and sets the default values.
func NewItemInternalaccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInternalaccountRequestBuilder) {
    m := &ItemInternalaccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/{directoryid}/internalaccount{?account*,additional*,any*,createdAfter*,createdBefore*,directory*,directoryOwnerGroup*,exclude*,id*,modifiedSince*,organizationalUnit*,q*,sort*,status*}", pathParameters),
    }
    return m
}
// NewItemInternalaccountRequestBuilder instantiates a new ItemInternalaccountRequestBuilder and sets the default values.
func NewItemInternalaccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInternalaccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInternalaccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get queries over all internal accounts within a directory. The various query parameters can be used to filter the response.
// returns a AuthInternalAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemInternalaccountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInternalaccountRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthInternalAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable), nil
}
// Post creates one or more new internal accounts within a directory and returns the newly created accounts.
// returns a AuthInternalAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemInternalaccountRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable, requestConfiguration *ItemInternalaccountRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateAuthInternalAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable), nil
}
// ToGetRequestInformation queries over all internal accounts within a directory. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ItemInternalaccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInternalaccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=76")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new internal accounts within a directory and returns the newly created accounts.
// returns a *RequestInformation when successful
func (m *ItemInternalaccountRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.AuthInternalAccountLinkableWrapperable, requestConfiguration *ItemInternalaccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=76")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=76", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInternalaccountRequestBuilder when successful
func (m *ItemInternalaccountRequestBuilder) WithUrl(rawUrl string)(*ItemInternalaccountRequestBuilder) {
    return NewItemInternalaccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
