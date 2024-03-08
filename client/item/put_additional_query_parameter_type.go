package item
import (
    "errors"
)
type PutAdditionalQueryParameterType int

const (
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    SECRET_PUTADDITIONALQUERYPARAMETERTYPE
    TILE_PUTADDITIONALQUERYPARAMETERTYPE
    VAULTRECORDCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"audit", "deleteTile", "groupclients", "groups", "secret", "tile", "vaultRecordCount"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "groupclients":
            result = GROUPCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_PUTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "vaultRecordCount":
            result = VAULTRECORDCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
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
