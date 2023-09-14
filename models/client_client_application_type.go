package models
import (
    "errors"
)
// 
type ClientClientApplicationType int

const (
    OAUTH2_CLIENTCLIENTAPPLICATIONTYPE ClientClientApplicationType = iota
    SAML2_CLIENTCLIENTAPPLICATIONTYPE
    LDAP_CLIENTCLIENTAPPLICATIONTYPE
)

func (i ClientClientApplicationType) String() string {
    return []string{"OAUTH2", "SAML2", "LDAP"}[i]
}
func ParseClientClientApplicationType(v string) (any, error) {
    result := OAUTH2_CLIENTCLIENTAPPLICATIONTYPE
    switch v {
        case "OAUTH2":
            result = OAUTH2_CLIENTCLIENTAPPLICATIONTYPE
        case "SAML2":
            result = SAML2_CLIENTCLIENTAPPLICATIONTYPE
        case "LDAP":
            result = LDAP_CLIENTCLIENTAPPLICATIONTYPE
        default:
            return 0, errors.New("Unknown ClientClientApplicationType value: " + v)
    }
    return &result, nil
}
func SerializeClientClientApplicationType(values []ClientClientApplicationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClientClientApplicationType) isMultiValue() bool {
    return false
}
