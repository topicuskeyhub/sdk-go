package models
import (
    "errors"
)
type RequestModificationRequestReportConclusion int

const (
    OK_REQUESTMODIFICATIONREQUESTREPORTCONCLUSION RequestModificationRequestReportConclusion = iota
    CANNOT_BE_EXECUTED_REQUESTMODIFICATIONREQUESTREPORTCONCLUSION
)

func (i RequestModificationRequestReportConclusion) String() string {
    return []string{"OK", "CANNOT_BE_EXECUTED"}[i]
}
func ParseRequestModificationRequestReportConclusion(v string) (any, error) {
    result := OK_REQUESTMODIFICATIONREQUESTREPORTCONCLUSION
    switch v {
        case "OK":
            result = OK_REQUESTMODIFICATIONREQUESTREPORTCONCLUSION
        case "CANNOT_BE_EXECUTED":
            result = CANNOT_BE_EXECUTED_REQUESTMODIFICATIONREQUESTREPORTCONCLUSION
        default:
            return 0, errors.New("Unknown RequestModificationRequestReportConclusion value: " + v)
    }
    return &result, nil
}
func SerializeRequestModificationRequestReportConclusion(values []RequestModificationRequestReportConclusion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestModificationRequestReportConclusion) isMultiValue() bool {
    return false
}
