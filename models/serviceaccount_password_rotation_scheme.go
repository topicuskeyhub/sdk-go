package models
import (
    "errors"
)
// 
type ServiceaccountPasswordRotationScheme int

const (
    MANUAL_SERVICEACCOUNTPASSWORDROTATIONSCHEME ServiceaccountPasswordRotationScheme = iota
    MANUAL_STORED_IN_VAULT_SERVICEACCOUNTPASSWORDROTATIONSCHEME
    DAILY_SERVICEACCOUNTPASSWORDROTATIONSCHEME
)

func (i ServiceaccountPasswordRotationScheme) String() string {
    return []string{"MANUAL", "MANUAL_STORED_IN_VAULT", "DAILY"}[i]
}
func ParseServiceaccountPasswordRotationScheme(v string) (any, error) {
    result := MANUAL_SERVICEACCOUNTPASSWORDROTATIONSCHEME
    switch v {
        case "MANUAL":
            result = MANUAL_SERVICEACCOUNTPASSWORDROTATIONSCHEME
        case "MANUAL_STORED_IN_VAULT":
            result = MANUAL_STORED_IN_VAULT_SERVICEACCOUNTPASSWORDROTATIONSCHEME
        case "DAILY":
            result = DAILY_SERVICEACCOUNTPASSWORDROTATIONSCHEME
        default:
            return 0, errors.New("Unknown ServiceaccountPasswordRotationScheme value: " + v)
    }
    return &result, nil
}
func SerializeServiceaccountPasswordRotationScheme(values []ServiceaccountPasswordRotationScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServiceaccountPasswordRotationScheme) isMultiValue() bool {
    return false
}
