package record
import (
    "errors"
)
type GetAdditionalQueryParameterType int

const (
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    DELETETILE_GETADDITIONALQUERYPARAMETERTYPE
    PARENT_GETADDITIONALQUERYPARAMETERTYPE
    PASSWORDMETADATA_GETADDITIONALQUERYPARAMETERTYPE
    SECRET_GETADDITIONALQUERYPARAMETERTYPE
    SHARESUMMARY_GETADDITIONALQUERYPARAMETERTYPE
    SHARES_GETADDITIONALQUERYPARAMETERTYPE
    TILE_GETADDITIONALQUERYPARAMETERTYPE
    VAULTHOLDER_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"audit", "deleteTile", "parent", "passwordMetadata", "secret", "shareSummary", "shares", "tile", "vaultholder"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_GETADDITIONALQUERYPARAMETERTYPE
        case "parent":
            result = PARENT_GETADDITIONALQUERYPARAMETERTYPE
        case "passwordMetadata":
            result = PASSWORDMETADATA_GETADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_GETADDITIONALQUERYPARAMETERTYPE
        case "shareSummary":
            result = SHARESUMMARY_GETADDITIONALQUERYPARAMETERTYPE
        case "shares":
            result = SHARES_GETADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_GETADDITIONALQUERYPARAMETERTYPE
        case "vaultholder":
            result = VAULTHOLDER_GETADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetAdditionalQueryParameterType value: " + v)
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
