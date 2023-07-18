package organizationalunit

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// OrganizationalunitRequestBuilder builds and executes requests for operations under \organizationalunit
type OrganizationalunitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OrganizationalunitRequestBuilderGetQueryParameters query for all organizational units in Topicus KeyHub. The various query parameters can be used to filter the response.
type OrganizationalunitRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string
    // Only return the organizational units that are equal to or an ancestor of the given unit(s), specified by id.
    AncestorOfOrEqualTo []int64
    // Return all or no records. This can be useful when composing parameters.
    Any []bool
    // Filter the organizational units by accounts, specified by id. This parameter supports composition with all parameters from the account resource.
    ConnectedToAccount []int64
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only return the organizational units that are equal to or a descendant of the given unit(s), specified by id.
    DescendantOfOrEqualTo []int64
    // Filter the results to exclude the given ids.
    Exclude []int64
    // Filter the results on the given ids.
    Id []int64
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filter organizational units on the exact name.
    Name []string
    // Filter organizational units on (part of) the name or uuid.
    NameContains []string
    // Filter the organizational units for which the given group is owner, specified by id. This parameter supports composition with all parameters from the group resource.
    OwnedBy []int64
    // Filter the organizational units by parent, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    Parent []int64
    // Filter records on a complex CQL query.
    Q []string
    // Only return organizational units that are or are not the root of the organizational tree.
    Root []bool
    // Filter results on one or more UUIDs.
    Uuid []string
}
// OrganizationalunitRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationalunitRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationalunitRequestBuilderGetQueryParameters
}
// OrganizationalunitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationalunitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/topicuskeyhub/sdk-go.organizationalunit.item collection
func (m *OrganizationalunitRequestBuilder) ById(id string)(*OrganizationalunitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewOrganizationalunitItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOrganizationalunitRequestBuilderInternal instantiates a new OrganizationalunitRequestBuilder and sets the default values.
func NewOrganizationalunitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationalunitRequestBuilder) {
    m := &OrganizationalunitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizationalunit{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,ancestorOfOrEqualTo*,connectedToAccount*,descendantOfOrEqualTo*,name*,nameContains*,ownedBy*,parent*,root*,uuid*}", pathParameters),
    }
    return m
}
// NewOrganizationalunitRequestBuilder instantiates a new OrganizationalunitRequestBuilder and sets the default values.
func NewOrganizationalunitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationalunitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationalunitRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for all organizational units in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *OrganizationalunitRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationalunitRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateOrganizationOrganizationalUnitLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable), nil
}
// Post creates one or more new organizational units and returns the newly created units.
func (m *OrganizationalunitRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable, requestConfiguration *OrganizationalunitRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateOrganizationOrganizationalUnitLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable), nil
}
// ToGetRequestInformation query for all organizational units in Topicus KeyHub. The various query parameters can be used to filter the response.
func (m *OrganizationalunitRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationalunitRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates one or more new organizational units and returns the newly created units.
func (m *OrganizationalunitRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.OrganizationOrganizationalUnitLinkableWrapperable, requestConfiguration *OrganizationalunitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
