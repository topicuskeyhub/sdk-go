package item
import (
    "errors"
)
type PutAdditionalQueryParameterType int

const (
    ACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    ISSUEDPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
    LOGINNAME_PUTADDITIONALQUERYPARAMETERTYPE
    MANAGEMENTPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
    MARKERS_PUTADDITIONALQUERYPARAMETERTYPE
    STATISTICS_PUTADDITIONALQUERYPARAMETERTYPE
    SUPPORTEDGROUPTYPES_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"account", "audit", "issuedPermissions", "loginName", "managementPermissions", "markers", "statistics", "supportedGroupTypes"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "account":
            result = ACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "issuedPermissions":
            result = ISSUEDPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
        case "loginName":
            result = LOGINNAME_PUTADDITIONALQUERYPARAMETERTYPE
        case "managementPermissions":
            result = MANAGEMENTPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_PUTADDITIONALQUERYPARAMETERTYPE
        case "statistics":
            result = STATISTICS_PUTADDITIONALQUERYPARAMETERTYPE
        case "supportedGroupTypes":
            result = SUPPORTEDGROUPTYPES_PUTADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PutAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializePutAdditionalQueryParameterType(values []PutAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PutAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
