package item
import (
    "errors"
)
type PutAdditionalQueryParameterType int

const (
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
    PARENT_PUTADDITIONALQUERYPARAMETERTYPE
    PASSWORDMETADATA_PUTADDITIONALQUERYPARAMETERTYPE
    SECRET_PUTADDITIONALQUERYPARAMETERTYPE
    SHARESUMMARY_PUTADDITIONALQUERYPARAMETERTYPE
    SHARES_PUTADDITIONALQUERYPARAMETERTYPE
    TILE_PUTADDITIONALQUERYPARAMETERTYPE
    VAULTHOLDER_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"audit", "deleteTile", "parent", "passwordMetadata", "secret", "shareSummary", "shares", "tile", "vaultholder"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "parent":
            result = PARENT_PUTADDITIONALQUERYPARAMETERTYPE
        case "passwordMetadata":
            result = PASSWORDMETADATA_PUTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_PUTADDITIONALQUERYPARAMETERTYPE
        case "shareSummary":
            result = SHARESUMMARY_PUTADDITIONALQUERYPARAMETERTYPE
        case "shares":
            result = SHARES_PUTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "vaultholder":
            result = VAULTHOLDER_PUTADDITIONALQUERYPARAMETERTYPE
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
