package record
import (
    "errors"
)
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
    PARENT_POSTADDITIONALQUERYPARAMETERTYPE
    PASSWORDMETADATA_POSTADDITIONALQUERYPARAMETERTYPE
    SECRET_POSTADDITIONALQUERYPARAMETERTYPE
    SHARESUMMARY_POSTADDITIONALQUERYPARAMETERTYPE
    SHARES_POSTADDITIONALQUERYPARAMETERTYPE
    TILE_POSTADDITIONALQUERYPARAMETERTYPE
    VAULTHOLDER_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "deleteTile", "parent", "passwordMetadata", "secret", "shareSummary", "shares", "tile", "vaultholder"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "parent":
            result = PARENT_POSTADDITIONALQUERYPARAMETERTYPE
        case "passwordMetadata":
            result = PASSWORDMETADATA_POSTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_POSTADDITIONALQUERYPARAMETERTYPE
        case "shareSummary":
            result = SHARESUMMARY_POSTADDITIONALQUERYPARAMETERTYPE
        case "shares":
            result = SHARES_POSTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "vaultholder":
            result = VAULTHOLDER_POSTADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PostAdditionalQueryParameterType value: " + v)
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
