package models
type ProvisioningLDAPSshPublicKeySupport int

const (
    DISABLED_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT ProvisioningLDAPSshPublicKeySupport = iota
    SSH_PUBLIC_KEY_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
    ALT_SECURITY_IDENTITIES_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
)

func (i ProvisioningLDAPSshPublicKeySupport) String() string {
    return []string{"DISABLED", "SSH_PUBLIC_KEY", "ALT_SECURITY_IDENTITIES"}[i]
}
func ParseProvisioningLDAPSshPublicKeySupport(v string) (any, error) {
    result := DISABLED_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
    switch v {
        case "DISABLED":
            result = DISABLED_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
        case "SSH_PUBLIC_KEY":
            result = SSH_PUBLIC_KEY_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
        case "ALT_SECURITY_IDENTITIES":
            result = ALT_SECURITY_IDENTITIES_PROVISIONINGLDAPSSHPUBLICKEYSUPPORT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProvisioningLDAPSshPublicKeySupport(values []ProvisioningLDAPSshPublicKeySupport) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningLDAPSshPublicKeySupport) isMultiValue() bool {
    return false
}
