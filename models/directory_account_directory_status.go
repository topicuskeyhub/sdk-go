package models
import (
    "errors"
)
// 
type DirectoryAccountDirectoryStatus int

const (
    INIT_REQUIRED_DIRECTORYACCOUNTDIRECTORYSTATUS DirectoryAccountDirectoryStatus = iota
    ONLINE_DIRECTORYACCOUNTDIRECTORYSTATUS
    DEGRADED_DIRECTORYACCOUNTDIRECTORYSTATUS
    NON_REDUNDANT_DIRECTORYACCOUNTDIRECTORYSTATUS
    OFFLINE_DIRECTORYACCOUNTDIRECTORYSTATUS
    DISABLED_DIRECTORYACCOUNTDIRECTORYSTATUS
)

func (i DirectoryAccountDirectoryStatus) String() string {
    return []string{"INIT_REQUIRED", "ONLINE", "DEGRADED", "NON_REDUNDANT", "OFFLINE", "DISABLED"}[i]
}
func ParseDirectoryAccountDirectoryStatus(v string) (any, error) {
    result := INIT_REQUIRED_DIRECTORYACCOUNTDIRECTORYSTATUS
    switch v {
        case "INIT_REQUIRED":
            result = INIT_REQUIRED_DIRECTORYACCOUNTDIRECTORYSTATUS
        case "ONLINE":
            result = ONLINE_DIRECTORYACCOUNTDIRECTORYSTATUS
        case "DEGRADED":
            result = DEGRADED_DIRECTORYACCOUNTDIRECTORYSTATUS
        case "NON_REDUNDANT":
            result = NON_REDUNDANT_DIRECTORYACCOUNTDIRECTORYSTATUS
        case "OFFLINE":
            result = OFFLINE_DIRECTORYACCOUNTDIRECTORYSTATUS
        case "DISABLED":
            result = DISABLED_DIRECTORYACCOUNTDIRECTORYSTATUS
        default:
            return 0, errors.New("Unknown DirectoryAccountDirectoryStatus value: " + v)
    }
    return &result, nil
}
func SerializeDirectoryAccountDirectoryStatus(values []DirectoryAccountDirectoryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
