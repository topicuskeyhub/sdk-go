package models
type DirectoryDirectoryRotatingPassword int

const (
    DEFAULT_OFF_DIRECTORYDIRECTORYROTATINGPASSWORD DirectoryDirectoryRotatingPassword = iota
    DEFAULT_ON_DIRECTORYDIRECTORYROTATINGPASSWORD
    ALWAYS_ON_DIRECTORYDIRECTORYROTATINGPASSWORD
)

func (i DirectoryDirectoryRotatingPassword) String() string {
    return []string{"DEFAULT_OFF", "DEFAULT_ON", "ALWAYS_ON"}[i]
}
func ParseDirectoryDirectoryRotatingPassword(v string) (any, error) {
    result := DEFAULT_OFF_DIRECTORYDIRECTORYROTATINGPASSWORD
    switch v {
        case "DEFAULT_OFF":
            result = DEFAULT_OFF_DIRECTORYDIRECTORYROTATINGPASSWORD
        case "DEFAULT_ON":
            result = DEFAULT_ON_DIRECTORYDIRECTORYROTATINGPASSWORD
        case "ALWAYS_ON":
            result = ALWAYS_ON_DIRECTORYDIRECTORYROTATINGPASSWORD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDirectoryDirectoryRotatingPassword(values []DirectoryDirectoryRotatingPassword) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryDirectoryRotatingPassword) isMultiValue() bool {
    return false
}
