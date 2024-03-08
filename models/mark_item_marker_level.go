package models
import (
    "errors"
)
type MarkItemMarkerLevel int

const (
    INFO_MARKITEMMARKERLEVEL MarkItemMarkerLevel = iota
    WARNING_MARKITEMMARKERLEVEL
)

func (i MarkItemMarkerLevel) String() string {
    return []string{"INFO", "WARNING"}[i]
}
func ParseMarkItemMarkerLevel(v string) (any, error) {
    result := INFO_MARKITEMMARKERLEVEL
    switch v {
        case "INFO":
            result = INFO_MARKITEMMARKERLEVEL
        case "WARNING":
            result = WARNING_MARKITEMMARKERLEVEL
        default:
            return 0, errors.New("Unknown MarkItemMarkerLevel value: " + v)
    }
    return &result, nil
}
func SerializeMarkItemMarkerLevel(values []MarkItemMarkerLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MarkItemMarkerLevel) isMultiValue() bool {
    return false
}
