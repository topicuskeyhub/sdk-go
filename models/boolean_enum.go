package models
import (
    "errors"
)
type BooleanEnum int

const (
    TRUE_BOOLEANENUM BooleanEnum = iota
    FALSE_BOOLEANENUM
    BOTH_BOOLEANENUM
)

func (i BooleanEnum) String() string {
    return []string{"TRUE", "FALSE", "BOTH"}[i]
}
func ParseBooleanEnum(v string) (any, error) {
    result := TRUE_BOOLEANENUM
    switch v {
        case "TRUE":
            result = TRUE_BOOLEANENUM
        case "FALSE":
            result = FALSE_BOOLEANENUM
        case "BOTH":
            result = BOTH_BOOLEANENUM
        default:
            return 0, errors.New("Unknown BooleanEnum value: " + v)
    }
    return &result, nil
}
func SerializeBooleanEnum(values []BooleanEnum) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BooleanEnum) isMultiValue() bool {
    return false
}
