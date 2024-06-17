package item
type GetAdditionalQueryParameterType int

const (
    ACCOUNT_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    ISSUEDPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
    LOGINNAME_GETADDITIONALQUERYPARAMETERTYPE
    MANAGEMENTPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
    MARKERS_GETADDITIONALQUERYPARAMETERTYPE
    STATISTICS_GETADDITIONALQUERYPARAMETERTYPE
    SUPPORTEDGROUPTYPES_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"account", "audit", "issuedPermissions", "loginName", "managementPermissions", "markers", "statistics", "supportedGroupTypes"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "account":
            result = ACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "issuedPermissions":
            result = ISSUEDPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
        case "loginName":
            result = LOGINNAME_GETADDITIONALQUERYPARAMETERTYPE
        case "managementPermissions":
            result = MANAGEMENTPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_GETADDITIONALQUERYPARAMETERTYPE
        case "statistics":
            result = STATISTICS_GETADDITIONALQUERYPARAMETERTYPE
        case "supportedGroupTypes":
            result = SUPPORTEDGROUPTYPES_GETADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetAdditionalQueryParameterType(values []GetAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
