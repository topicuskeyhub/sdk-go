package group

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemVaultRecordExportRequestBuilder builds and executes requests for operations under \group\{groupid}\vault\record\export
type ItemVaultRecordExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVaultRecordExportRequestBuilderPostQueryParameters prepares an export of vault records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
type ItemVaultRecordExportRequestBuilderPostQueryParameters struct {
    // Filter records that accessible by the given account(s), specified by id.
    AccessibleByAccount []int64 `uriparametername:"accessibleByAccount"`
    // Filter records that accessible by the given account(s) and the account is manager of the group, specified by id.
    AccessibleByAccountAsManager []int64 `uriparametername:"accessibleByAccountAsManager"`
    // Filter records that accessible by the given client(s), specified by id.
    AccessibleByClient []int64 `uriparametername:"accessibleByClient"`
    // Filter the records by account, specified by id. This parameter supports composition with all parameters from the account resource.
    Account []int64 `uriparametername:"account"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the records by client, specified by id. This parameter supports composition with all parameters from the client resource.
    Client []int64 `uriparametername:"client"`
    // Filter records on the color.
    // Deprecated: This property is deprecated, use ColorAsVaultVaultRecordColor instead
    Color []string `uriparametername:"color"`
    // Filter records on the color.
    ColorAsVaultVaultRecordColor []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordColor `uriparametername:"color"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return records that show an expiration warning at or before the given date.
    ExpireWarningBeforeOrAt []i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly `uriparametername:"expireWarningBeforeOrAt"`
    // Filter records on the exact filename.
    Filename []string `uriparametername:"filename"`
    // Filter the records by group, specified by id. This parameter supports composition with all parameters from the group resource.
    Group []int64 `uriparametername:"group"`
    // Filter the records by organizational unit of the group, specified by id. This parameter supports composition with all parameters from the client resource.
    GroupOrganizationalUnit []int64 `uriparametername:"groupOrganizationalUnit"`
    // Only return records that have or do not have an expiration policy set.
    HasNoPolicy []bool `uriparametername:"hasNoPolicy"`
    // Filter records that are shared (have a parent record)
    HasParent []bool `uriparametername:"hasParent"`
    // Only return records that have or do not have expired.
    HasValidPolicy []bool `uriparametername:"hasValidPolicy"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on the exact name.
    Name []string `uriparametername:"name"`
    // Search records on (part of) the UUID, name, filename, URI or username.
    NameContains []string `uriparametername:"nameContains"`
    // Filter records that are shared from the specified parent record, specified by id.
    Parent []int64 `uriparametername:"parent"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter records by secret using composition. Use secret.type to filter on type.
    Secret []int64 `uriparametername:"secret"`
    // Only return records for which the sharing period has expired at or before the given date.
    ShareExpiresBeforeOrAt []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"shareExpiresBeforeOrAt"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter records on the exact URI.
    Url []string `uriparametername:"url"`
    // Filter records on the exact username.
    Username []string `uriparametername:"username"`
    // Filter records on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
    // Filter the records by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64 `uriparametername:"vault"`
}
// ItemVaultRecordExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultRecordExportRequestBuilderPostQueryParameters
}
// NewItemVaultRecordExportRequestBuilderInternal instantiates a new ItemVaultRecordExportRequestBuilder and sets the default values.
func NewItemVaultRecordExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordExportRequestBuilder) {
    m := &ItemVaultRecordExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/group/{groupid}/vault/record/export{?accessibleByAccount*,accessibleByAccountAsManager*,accessibleByClient*,account*,any*,client*,color*,createdAfter*,createdBefore*,exclude*,expireWarningBeforeOrAt*,filename*,group*,groupOrganizationalUnit*,hasNoPolicy*,hasParent*,hasValidPolicy*,id*,modifiedSince*,name*,nameContains*,parent*,q*,secret*,shareExpiresBeforeOrAt*,sort*,url*,username*,uuid*,vault*}", pathParameters),
    }
    return m
}
// NewItemVaultRecordExportRequestBuilder instantiates a new ItemVaultRecordExportRequestBuilder and sets the default values.
func NewItemVaultRecordExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVaultRecordExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post prepares an export of vault records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ItemVaultRecordExportRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemVaultRecordExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation prepares an export of vault records using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a *RequestInformation when successful
func (m *ItemVaultRecordExportRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemVaultRecordExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVaultRecordExportRequestBuilder when successful
func (m *ItemVaultRecordExportRequestBuilder) WithUrl(rawUrl string)(*ItemVaultRecordExportRequestBuilder) {
    return NewItemVaultRecordExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
