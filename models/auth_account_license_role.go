package models
type AuthAccountLicenseRole int

const (
    IGA_AUTHACCOUNTLICENSEROLE AuthAccountLicenseRole = iota
    BUSINESS_AUTHACCOUNTLICENSEROLE
    PRO_AUTHACCOUNTLICENSEROLE
)

func (i AuthAccountLicenseRole) String() string {
    return []string{"IGA", "BUSINESS", "PRO"}[i]
}
func ParseAuthAccountLicenseRole(v string) (any, error) {
    result := IGA_AUTHACCOUNTLICENSEROLE
    switch v {
        case "IGA":
            result = IGA_AUTHACCOUNTLICENSEROLE
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
