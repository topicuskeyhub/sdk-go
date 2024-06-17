package system
type PostAdditionalQueryParameterType int

const (
    ACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    ISSUEDPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
    LOGINNAME_POSTADDITIONALQUERYPARAMETERTYPE
    MANAGEMENTPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
    MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
    STATISTICS_POSTADDITIONALQUERYPARAMETERTYPE
    SUPPORTEDGROUPTYPES_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"account", "audit", "issuedPermissions", "loginName", "managementPermissions", "markers", "statistics", "supportedGroupTypes"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "account":
            result = ACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "issuedPermissions":
            result = ISSUEDPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
        case "loginName":
            result = LOGINNAME_POSTADDITIONALQUERYPARAMETERTYPE
        case "managementPermissions":
            result = MANAGEMENTPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
        case "statistics":
            result = STATISTICS_POSTADDITIONALQUERYPARAMETERTYPE
        case "supportedGroupTypes":
            result = SUPPORTEDGROUPTYPES_POSTADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePostAdditionalQueryParameterType(values []PostAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PostAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
