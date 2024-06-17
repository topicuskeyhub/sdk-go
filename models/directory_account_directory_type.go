package models
type DirectoryAccountDirectoryType int

const (
    MAINTENANCE_DIRECTORYACCOUNTDIRECTORYTYPE DirectoryAccountDirectoryType = iota
    LDAP_DIRECTORYACCOUNTDIRECTORYTYPE
    INTERNAL_DIRECTORYACCOUNTDIRECTORYTYPE
    OIDC_DIRECTORYACCOUNTDIRECTORYTYPE
    PENDING_ACCOUNTS_DIRECTORYACCOUNTDIRECTORYTYPE
)

func (i DirectoryAccountDirectoryType) String() string {
    return []string{"MAINTENANCE", "LDAP", "INTERNAL", "OIDC", "PENDING_ACCOUNTS"}[i]
}
func ParseDirectoryAccountDirectoryType(v string) (any, error) {
    result := MAINTENANCE_DIRECTORYACCOUNTDIRECTORYTYPE
    switch v {
        case "MAINTENANCE":
            result = MAINTENANCE_DIRECTORYACCOUNTDIRECTORYTYPE
        case "LDAP":
            result = LDAP_DIRECTORYACCOUNTDIRECTORYTYPE
        case "INTERNAL":
            result = INTERNAL_DIRECTORYACCOUNTDIRECTORYTYPE
        case "OIDC":
            result = OIDC_DIRECTORYACCOUNTDIRECTORYTYPE
        case "PENDING_ACCOUNTS":
            result = PENDING_ACCOUNTS_DIRECTORYACCOUNTDIRECTORYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDirectoryAccountDirectoryType(values []DirectoryAccountDirectoryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryAccountDirectoryType) isMultiValue() bool {
    return false
}
