package models
type DirectoryLDAPDirectoryPasswordRecovery int

const (
    DISABLED_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY DirectoryLDAPDirectoryPasswordRecovery = iota
    VERIFY_2FA_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
    VERIFY_MAIL_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
    VERIFY_MAIL_AND_2FA_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
)

func (i DirectoryLDAPDirectoryPasswordRecovery) String() string {
    return []string{"DISABLED", "VERIFY_2FA", "VERIFY_MAIL", "VERIFY_MAIL_AND_2FA"}[i]
}
func ParseDirectoryLDAPDirectoryPasswordRecovery(v string) (any, error) {
    result := DISABLED_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
    switch v {
        case "DISABLED":
            result = DISABLED_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
        case "VERIFY_2FA":
            result = VERIFY_2FA_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
        case "VERIFY_MAIL":
            result = VERIFY_MAIL_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
        case "VERIFY_MAIL_AND_2FA":
            result = VERIFY_MAIL_AND_2FA_DIRECTORYLDAPDIRECTORYPASSWORDRECOVERY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDirectoryLDAPDirectoryPasswordRecovery(values []DirectoryLDAPDirectoryPasswordRecovery) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryLDAPDirectoryPasswordRecovery) isMultiValue() bool {
    return false
}
