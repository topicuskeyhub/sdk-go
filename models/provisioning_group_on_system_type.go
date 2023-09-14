package models
import (
    "errors"
)
// 
type ProvisioningGroupOnSystemType int

const (
    POSIX_GROUP_PROVISIONINGGROUPONSYSTEMTYPE ProvisioningGroupOnSystemType = iota
    GROUP_OF_NAMES_PROVISIONINGGROUPONSYSTEMTYPE
    GROUP_OF_UNIQUE_NAMES_PROVISIONINGGROUPONSYSTEMTYPE
    GROUP_PROVISIONINGGROUPONSYSTEMTYPE
    AZURE_ROLE_PROVISIONINGGROUPONSYSTEMTYPE
    AZURE_UNIFIED_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
    AZURE_SECURITY_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
)

func (i ProvisioningGroupOnSystemType) String() string {
    return []string{"POSIX_GROUP", "GROUP_OF_NAMES", "GROUP_OF_UNIQUE_NAMES", "GROUP", "AZURE_ROLE", "AZURE_UNIFIED_GROUP", "AZURE_SECURITY_GROUP"}[i]
}
func ParseProvisioningGroupOnSystemType(v string) (any, error) {
    result := POSIX_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
    switch v {
        case "POSIX_GROUP":
            result = POSIX_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
        case "GROUP_OF_NAMES":
            result = GROUP_OF_NAMES_PROVISIONINGGROUPONSYSTEMTYPE
        case "GROUP_OF_UNIQUE_NAMES":
            result = GROUP_OF_UNIQUE_NAMES_PROVISIONINGGROUPONSYSTEMTYPE
        case "GROUP":
            result = GROUP_PROVISIONINGGROUPONSYSTEMTYPE
        case "AZURE_ROLE":
            result = AZURE_ROLE_PROVISIONINGGROUPONSYSTEMTYPE
        case "AZURE_UNIFIED_GROUP":
            result = AZURE_UNIFIED_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
        case "AZURE_SECURITY_GROUP":
            result = AZURE_SECURITY_GROUP_PROVISIONINGGROUPONSYSTEMTYPE
        default:
            return 0, errors.New("Unknown ProvisioningGroupOnSystemType value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningGroupOnSystemType(values []ProvisioningGroupOnSystemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningGroupOnSystemType) isMultiValue() bool {
    return false
}
