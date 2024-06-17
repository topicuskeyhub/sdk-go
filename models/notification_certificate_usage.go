package models
type NotificationCertificateUsage int

const (
    LDAP_DIRECTORY_NOTIFICATIONCERTIFICATEUSAGE NotificationCertificateUsage = iota
    PROVISIONED_AD_NOTIFICATIONCERTIFICATEUSAGE
    PROVISIONED_LDAP_NOTIFICATIONCERTIFICATEUSAGE
    WEBHOOK_NOTIFICATIONCERTIFICATEUSAGE
    LDAP_CLIENT_NOTIFICATIONCERTIFICATEUSAGE
    LDAP_SERVER_SETTING_NOTIFICATIONCERTIFICATEUSAGE
    GLOBAL_TRUSTED_CERTIFICATE_NOTIFICATIONCERTIFICATEUSAGE
    IDP_NOTIFICATIONCERTIFICATEUSAGE
    KEYHUB_NOTIFICATIONCERTIFICATEUSAGE
)

func (i NotificationCertificateUsage) String() string {
    return []string{"LDAP_DIRECTORY", "PROVISIONED_AD", "PROVISIONED_LDAP", "WEBHOOK", "LDAP_CLIENT", "LDAP_SERVER_SETTING", "GLOBAL_TRUSTED_CERTIFICATE", "IDP", "KEYHUB"}[i]
}
func ParseNotificationCertificateUsage(v string) (any, error) {
    result := LDAP_DIRECTORY_NOTIFICATIONCERTIFICATEUSAGE
    switch v {
        case "LDAP_DIRECTORY":
            result = LDAP_DIRECTORY_NOTIFICATIONCERTIFICATEUSAGE
        case "PROVISIONED_AD":
            result = PROVISIONED_AD_NOTIFICATIONCERTIFICATEUSAGE
        case "PROVISIONED_LDAP":
            result = PROVISIONED_LDAP_NOTIFICATIONCERTIFICATEUSAGE
        case "WEBHOOK":
            result = WEBHOOK_NOTIFICATIONCERTIFICATEUSAGE
        case "LDAP_CLIENT":
            result = LDAP_CLIENT_NOTIFICATIONCERTIFICATEUSAGE
        case "LDAP_SERVER_SETTING":
            result = LDAP_SERVER_SETTING_NOTIFICATIONCERTIFICATEUSAGE
        case "GLOBAL_TRUSTED_CERTIFICATE":
            result = GLOBAL_TRUSTED_CERTIFICATE_NOTIFICATIONCERTIFICATEUSAGE
        case "IDP":
            result = IDP_NOTIFICATIONCERTIFICATEUSAGE
        case "KEYHUB":
            result = KEYHUB_NOTIFICATIONCERTIFICATEUSAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNotificationCertificateUsage(values []NotificationCertificateUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NotificationCertificateUsage) isMultiValue() bool {
    return false
}
