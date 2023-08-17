package deletedvaultholder

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemVaultRecordRequestBuilder builds and executes requests for operations under \deletedvaultholder\{vaultid}\vault\record
type ItemVaultRecordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVaultRecordRequestBuilderGetQueryParameters query for vault records in a specific vault. The various query parameters can be used to filter the response. It is not possible to access secrets with a query. Secrets can only be read one by one.
type ItemVaultRecordRequestBuilderGetQueryParameters struct {
    // Filter records that accessible by the given account(s), specified by id.
    AccessibleByAccount []int64
    // Filter records that accessible by the given account(s) and the account is manager of the group, specified by id.
    AccessibleByAccountAsManager []int64
    // Filter records that accessible by the given client(s), specified by id.
    AccessibleByClient []int64
    // 
    Account []int64
    // Request additional information to be returned for every record.
    Additional []string
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // 
    Client []int64
    // Filter records on the color.
    Color []string
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Only return records that show an expiration warning at or before the given date.
    ExpireWarningBeforeOrAt []i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Filter records on the exact filename.
    Filename []string
    // 
    Group []int64
    // Only return records that have or do not have an expiration policy set.
    HasNoPolicy []bool
    // Filter records that are shared (have a parent record)
    HasParent []bool
    // Only return records that have or do not have expired.
    HasValidPolicy []bool
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter records on the exact name.
    Name []string
    // Search records on (part of) the UUID, name, filename, URI or username.
    NameContains []string
    // Filter records that are shared from the specified parent record, specified by id.
    Parent []int64
    // Filter records on a complex CQL query.
    Q []string
    // Filter records by secret using composition. Use secret.type to filter on type.
    Secret []int64
    // Only return records for which the sharing period has expired at or before the given date.
    ShareExpiresBeforeOrAt []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter records on the exact URI.
    Url []string
    // Filter records on the exact username.
    Username []string
    // Filter records on one or more UUIDs.
    Uuid []string
    // Filter the records by vault, specified by id. This parameter supports composition with all parameters from the vault resource.
    Vault []int64
}
// ItemVaultRecordRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemVaultRecordRequestBuilderGetQueryParameters
}
// ItemVaultRecordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemVaultRecordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.deletedvaultholder.item.vault.record.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemVaultRecordRequestBuilder) ById(id string)(*ItemVaultRecordRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewItemVaultRecordRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInt64 gets an item from the github.com/topicuskeyhub/sdk-go.deletedvaultholder.item.vault.record.item collection
func (m *ItemVaultRecordRequestBuilder) ByIdInt64(id int64)(*ItemVaultRecordRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewItemVaultRecordRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemVaultRecordRequestBuilderInternal instantiates a new RecordRequestBuilder and sets the default values.
func NewItemVaultRecordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordRequestBuilder) {
    m := &ItemVaultRecordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deletedvaultholder/{vaultid}/vault/record{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,accessibleByAccount*,accessibleByAccountAsManager*,accessibleByClient*,account*,client*,color*,expireWarningBeforeOrAt*,filename*,group*,hasNoPolicy*,hasParent*,hasValidPolicy*,name*,nameContains*,parent*,secret*,shareExpiresBeforeOrAt*,url*,username*,uuid*,vault*}", pathParameters),
    }
    return m
}
// NewItemVaultRecordRequestBuilder instantiates a new RecordRequestBuilder and sets the default values.
func NewItemVaultRecordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVaultRecordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVaultRecordRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for vault records in a specific vault. The various query parameters can be used to filter the response. It is not possible to access secrets with a query. Secrets can only be read one by one.
func (m *ItemVaultRecordRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemVaultRecordRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultVaultRecordLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable), nil
}
// Post creates one or more new vault records and returns the newly created records. Secrets are specified via the additional object secret. It is required to specify the 'topicus-Vault-session' header. When updating a TOTP-secret, make sure to set 'writeTotp' field.
func (m *ItemVaultRecordRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable, requestConfiguration *ItemVaultRecordRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateVaultVaultRecordLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable), nil
}
// ToGetRequestInformation query for vault records in a specific vault. The various query parameters can be used to filter the response. It is not possible to access secrets with a query. Secrets can only be read one by one.
func (m *ItemVaultRecordRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemVaultRecordRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new vault records and returns the newly created records. Secrets are specified via the additional object secret. It is required to specify the 'topicus-Vault-session' header. When updating a TOTP-secret, make sure to set 'writeTotp' field.
func (m *ItemVaultRecordRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.VaultVaultRecordLinkableWrapperable, requestConfiguration *ItemVaultRecordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=65")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=65", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
