package models
import (
    "errors"
)
type Month int

const (
    JANUARY_MONTH Month = iota
    FEBRUARY_MONTH
    MARCH_MONTH
    APRIL_MONTH
    MAY_MONTH
    JUNE_MONTH
    JULY_MONTH
    AUGUST_MONTH
    SEPTEMBER_MONTH
    OCTOBER_MONTH
    NOVEMBER_MONTH
    DECEMBER_MONTH
)

func (i Month) String() string {
    return []string{"JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"}[i]
}
func ParseMonth(v string) (any, error) {
    result := JANUARY_MONTH
    switch v {
        case "JANUARY":
            result = JANUARY_MONTH
        case "FEBRUARY":
            result = FEBRUARY_MONTH
        case "MARCH":
            result = MARCH_MONTH
        case "APRIL":
            result = APRIL_MONTH
        case "MAY":
            result = MAY_MONTH
        case "JUNE":
            result = JUNE_MONTH
        case "JULY":
            result = JULY_MONTH
        case "AUGUST":
            result = AUGUST_MONTH
        case "SEPTEMBER":
            result = SEPTEMBER_MONTH
        case "OCTOBER":
            result = OCTOBER_MONTH
        case "NOVEMBER":
            result = NOVEMBER_MONTH
        case "DECEMBER":
            result = DECEMBER_MONTH
        default:
            return 0, errors.New("Unknown Month value: " + v)
    }
    return &result, nil
}
func SerializeMonth(values []Month) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Month) isMultiValue() bool {
    return false
}
