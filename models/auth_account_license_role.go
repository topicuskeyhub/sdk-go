package models
type AuthAccountLicenseRole int

const (
    BUSINESS_AUTHACCOUNTLICENSEROLE AuthAccountLicenseRole = iota
    PRO_AUTHACCOUNTLICENSEROLE
)

func (i AuthAccountLicenseRole) String() string {
    return []string{"BUSINESS", "PRO"}[i]
}
func ParseAuthAccountLicenseRole(v string) (any, error) {
    result := BUSINESS_AUTHACCOUNTLICENSEROLE
    switch v {
        case "BUSINESS":
            result = BUSINESS_AUTHACCOUNTLICENSEROLE
        case "PRO":
            result = PRO_AUTHACCOUNTLICENSEROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthAccountLicenseRole(values []AuthAccountLicenseRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthAccountLicenseRole) isMultiValue() bool {
    return false
}
