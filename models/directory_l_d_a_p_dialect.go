package models
type DirectoryLDAPDialect int

const (
    ACTIVE_DIRECTORY_DIRECTORYLDAPDIALECT DirectoryLDAPDialect = iota
    OPENLDAP_DIRECTORYLDAPDIALECT
)

func (i DirectoryLDAPDialect) String() string {
    return []string{"ACTIVE_DIRECTORY", "OPENLDAP"}[i]
}
func ParseDirectoryLDAPDialect(v string) (any, error) {
    result := ACTIVE_DIRECTORY_DIRECTORYLDAPDIALECT
    switch v {
        case "ACTIVE_DIRECTORY":
            result = ACTIVE_DIRECTORY_DIRECTORYLDAPDIALECT
        case "OPENLDAP":
            result = OPENLDAP_DIRECTORYLDAPDIALECT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDirectoryLDAPDialect(values []DirectoryLDAPDialect) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryLDAPDialect) isMultiValue() bool {
    return false
}
