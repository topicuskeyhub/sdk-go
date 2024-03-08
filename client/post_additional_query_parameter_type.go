package client
import (
    "errors"
)
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    SECRET_POSTADDITIONALQUERYPARAMETERTYPE
    TILE_POSTADDITIONALQUERYPARAMETERTYPE
    VAULTRECORDCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "deleteTile", "groupclients", "groups", "secret", "tile", "vaultRecordCount"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "groupclients":
            result = GROUPCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_POSTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "vaultRecordCount":
            result = VAULTRECORDCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
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
