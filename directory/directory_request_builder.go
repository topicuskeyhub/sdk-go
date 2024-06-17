package directory

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// DirectoryRequestBuilder builds and executes requests for operations under \directory
type DirectoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRequestBuilderGetQueryParameters query for all directories in Topicus KeyHub. The various query parameters can be used to filter the response.
type DirectoryRequestBuilderGetQueryParameters struct {
    // Only return directories that are or are not active. Defaults to true.
    // Deprecated: This property is deprecated, use ActiveAsBooleanEnum instead
    Active []string `uriparametername:"active"`
    // Only return directories that are or are not active. Defaults to true.
    ActiveAsBooleanEnum []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.BooleanEnum `uriparametername:"active"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Filter the directories on the given base organizational unit, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    BaseOrganizationalUnit []int64 `uriparametername:"baseOrganizationalUnit"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Only return directories that are or are not the default directory.
    DefaultDirectory []bool `uriparametername:"defaultDirectory"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Only return directories that use certificates that are expired at the given instant.
    ExpiredCertificate []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"expiredCertificate"`
    // Filter the directories on the given helpdesk groups, specified by id. This parameter supports composition with all parameters from the group resource.
    HelpdeskGroup []int64 `uriparametername:"helpdeskGroup"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return directories that are not used for source directory provisioning for the given types of linked systems.
    // Deprecated: This property is deprecated, use IsNotProvisionedDirectoryAsProvisioningProvisionedSystemType instead
    IsNotProvisionedDirectory []string `uriparametername:"isNotProvisionedDirectory"`
    // Only return directories that are not used for source directory provisioning for the given types of linked systems.
    IsNotProvisionedDirectoryAsProvisioningProvisionedSystemType []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemType `uriparametername:"isNotProvisionedDirectory"`
    // Only return directories that are or are not the built-in maintenance directory.
    MaintenanceDirectory []bool `uriparametername:"maintenanceDirectory"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter directories on the exact name.
    Name []string `uriparametername:"name"`
    // Filter directories on (part of) the name or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Filter directories where the name does not start with the given values.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Filter directories on the start of the name.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Only return OIDC directories for the given vendor(s).
    // Deprecated: This property is deprecated, use OidcVenderAsDirectoryOIDCVendor instead
    OidcVender []string `uriparametername:"oidcVender"`
    // Only return OIDC directories for the given vendor(s).
    OidcVenderAsDirectoryOIDCVendor []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryOIDCVendor `uriparametername:"oidcVender"`
    // Only return internal directories for the given owner(s), specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64 `uriparametername:"ownedBy"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter LDAP directories on TLS setting(s) used.
    // Deprecated: This property is deprecated, use TlsAsTLSLevel instead
    Tls []string `uriparametername:"tls"`
    // Filter LDAP directories on TLS setting(s) used.
    TlsAsTLSLevel []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.TLSLevel `uriparametername:"tls"`
    // Only return directories of the given type(s).
    Type []string `uriparametername:"type"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// DirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters
}
// DirectoryRequestBuilderPostQueryParameters creates one or more new directories and returns the newly created directories.
type DirectoryRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// DirectoryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderPostQueryParameters
}
// ByDirectoryid gets an item from the github.com/topicuskeyhub/sdk-go.directory.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithDirectoryItemRequestBuilder when successful
func (m *DirectoryRequestBuilder) ByDirectoryid(directoryid string)(*WithDirectoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryid != "" {
        urlTplParams["directoryid"] = directoryid
    }
    return NewWithDirectoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDirectoryidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.directory.item collection
// returns a *WithDirectoryItemRequestBuilder when successful
func (m *DirectoryRequestBuilder) ByDirectoryidInt64(directoryid int64)(*WithDirectoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["directoryid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(directoryid, 10)
    return NewWithDirectoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory{?active*,additional*,any*,baseOrganizationalUnit*,createdAfter*,createdBefore*,defaultDirectory*,exclude*,expiredCertificate*,helpdeskGroup*,id*,isNotProvisionedDirectory*,maintenanceDirectory*,modifiedSince*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,oidcVender*,ownedBy*,q*,sort*,tls*,type*,uuid*}", pathParameters),
    }
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all directories in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a DirectoryAccountDirectoryLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *DirectoryRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateDirectoryAccountDirectoryLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable), nil
}
// Post creates one or more new directories and returns the newly created directories.
// returns a DirectoryAccountDirectoryLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *DirectoryRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable, requestConfiguration *DirectoryRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateDirectoryAccountDirectoryLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable), nil
}
// ToGetRequestInformation query for all directories in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *DirectoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new directories and returns the newly created directories.
// returns a *RequestInformation when successful
func (m *DirectoryRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.DirectoryAccountDirectoryLinkableWrapperable, requestConfiguration *DirectoryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=71")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=71", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DirectoryRequestBuilder when successful
func (m *DirectoryRequestBuilder) WithUrl(rawUrl string)(*DirectoryRequestBuilder) {
    return NewDirectoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
