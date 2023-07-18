package models
import (
    "errors"
)
// 
type ProvisioningProvisionedSystemType int

const (
    LDAP_PROVISIONINGPROVISIONEDSYSTEMTYPE ProvisioningProvisionedSystemType = iota
    ACTIVE_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
    AZURE_TENANT_PROVISIONINGPROVISIONEDSYSTEMTYPE
    SOURCE_LDAP_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
    SOURCE_AZURE_OIDC_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
    SOURCE_AZURE_SYNC_LDAP_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
    INTERNAL_LDAP_PROVISIONINGPROVISIONEDSYSTEMTYPE
)

func (i ProvisioningProvisionedSystemType) String() string {
    return []string{"LDAP", "ACTIVE_DIRECTORY", "AZURE_TENANT", "SOURCE_LDAP_DIRECTORY", "SOURCE_AZURE_OIDC_DIRECTORY", "SOURCE_AZURE_SYNC_LDAP_DIRECTORY", "INTERNAL_LDAP"}[i]
}
func ParseProvisioningProvisionedSystemType(v string) (any, error) {
    result := LDAP_PROVISIONINGPROVISIONEDSYSTEMTYPE
    switch v {
        case "LDAP":
            result = LDAP_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "ACTIVE_DIRECTORY":
            result = ACTIVE_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "AZURE_TENANT":
            result = AZURE_TENANT_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "SOURCE_LDAP_DIRECTORY":
            result = SOURCE_LDAP_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "SOURCE_AZURE_OIDC_DIRECTORY":
            result = SOURCE_AZURE_OIDC_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "SOURCE_AZURE_SYNC_LDAP_DIRECTORY":
            result = SOURCE_AZURE_SYNC_LDAP_DIRECTORY_PROVISIONINGPROVISIONEDSYSTEMTYPE
        case "INTERNAL_LDAP":
            result = INTERNAL_LDAP_PROVISIONINGPROVISIONEDSYSTEMTYPE
        default:
            return 0, errors.New("Unknown ProvisioningProvisionedSystemType value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningProvisionedSystemType(values []ProvisioningProvisionedSystemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
