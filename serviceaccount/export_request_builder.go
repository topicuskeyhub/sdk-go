package serviceaccount

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ExportRequestBuilder builds and executes requests for operations under \serviceaccount\export
type ExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExportRequestBuilderPostQueryParameters prepares an export of service accounts using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
type ExportRequestBuilderPostQueryParameters struct {
    // Only return active or inactive service accounts.
    // Deprecated: This property is deprecated, use ActiveAsBooleanEnum instead
    Active []string `uriparametername:"active"`
    // Only return active or inactive service accounts.
    ActiveAsBooleanEnum []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.BooleanEnum `uriparametername:"active"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the service accounts by groups on systems, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Filter the service accounts on group on systems owned by any of the given groups, specified by id.
    GroupOnSystemOwners []int64 `uriparametername:"groupOnSystemOwners"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter service accounts on the exact name.
    Name []string `uriparametername:"name"`
    // Search service accounts on (part of) the name, username or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Only return service accounts for which the name does not start with the given prefix.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Only return service accounts for which the name starts with the given prefix.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Filter service accounts on their organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnitForEnforcement []int64 `uriparametername:"organizationalUnitForEnforcement"`
    // Filter the service accounts by the password shared in the vault, specified by id. This parameter supports composition with all parameters from the vault record resource.
    Password []int64 `uriparametername:"password"`
    // Only return service accounts with the given password rotation scheme.
    // Deprecated: This property is deprecated, use PasswordRotationAsServiceaccountPasswordRotationScheme instead
    PasswordRotation []string `uriparametername:"passwordRotation"`
    // Only return service accounts with the given password rotation scheme.
    PasswordRotationAsServiceaccountPasswordRotationScheme []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountPasswordRotationScheme `uriparametername:"passwordRotation"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter the service accounts on active requests for group on systems owned by any of the given groups, specified by id.
    RequestedGroupOnSystemOwners []int64 `uriparametername:"requestedGroupOnSystemOwners"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter the service accounts by provisioned systems, specified by id. This parameter supports composition with all parameters from the provisioned system resource.
    System []int64 `uriparametername:"system"`
    // Filter the service accounts on the content administration group of a provisioned system, specified by id.
    SystemContentAdministrators []int64 `uriparametername:"systemContentAdministrators"`
    // Filter the service accounts on the owning group of its provisioned system, specified by id.
    SystemOwners []int64 `uriparametername:"systemOwners"`
    // Filter the service accounts on the technical administrator group of its provisioned system, specified by id.
    SystemTechnicalAdministrators []int64 `uriparametername:"systemTechnicalAdministrators"`
    // Filter the service accounts on groups that perform technical administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64 `uriparametername:"technicalAdministrator"`
    // Filter service accounts on the exact username.
    Username []string `uriparametername:"username"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// ExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExportRequestBuilderPostQueryParameters
}
// NewExportRequestBuilderInternal instantiates a new ExportRequestBuilder and sets the default values.
func NewExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportRequestBuilder) {
    m := &ExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount/export{?active*,any*,createdAfter*,createdBefore*,exclude*,groupOnSystem*,groupOnSystemOwners*,id*,modifiedSince*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,organizationalUnitForEnforcement*,password*,passwordRotation*,q*,requestedGroupOnSystemOwners*,sort*,system*,systemContentAdministrators*,systemOwners*,systemTechnicalAdministrators*,technicalAdministrator*,username*,uuid*}", pathParameters),
    }
    return m
}
// NewExportRequestBuilder instantiates a new ExportRequestBuilder and sets the default values.
func NewExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post prepares an export of service accounts using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ExportRequestBuilder) Post(ctx context.Context, requestConfiguration *ExportRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation prepares an export of service accounts using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a *RequestInformation when successful
func (m *ExportRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExportRequestBuilder when successful
func (m *ExportRequestBuilder) WithUrl(rawUrl string)(*ExportRequestBuilder) {
    return NewExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
