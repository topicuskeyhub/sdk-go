package models
type AuthWebAuthnPRFSupported int

const (
    UNKNOWN_AUTHWEBAUTHNPRFSUPPORTED AuthWebAuthnPRFSupported = iota
    SUPPORTED_NOT_LINKED_AUTHWEBAUTHNPRFSUPPORTED
    SUPPORTED_AUTHWEBAUTHNPRFSUPPORTED
    NOT_SUPPORTED_AUTHWEBAUTHNPRFSUPPORTED
)

func (i AuthWebAuthnPRFSupported) String() string {
    return []string{"UNKNOWN", "SUPPORTED_NOT_LINKED", "SUPPORTED", "NOT_SUPPORTED"}[i]
}
func ParseAuthWebAuthnPRFSupported(v string) (any, error) {
    result := UNKNOWN_AUTHWEBAUTHNPRFSUPPORTED
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_AUTHWEBAUTHNPRFSUPPORTED
        case "SUPPORTED_NOT_LINKED":
            result = SUPPORTED_NOT_LINKED_AUTHWEBAUTHNPRFSUPPORTED
        case "SUPPORTED":
            result = SUPPORTED_AUTHWEBAUTHNPRFSUPPORTED
        case "NOT_SUPPORTED":
            result = NOT_SUPPORTED_AUTHWEBAUTHNPRFSUPPORTED
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthWebAuthnPRFSupported(values []AuthWebAuthnPRFSupported) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthWebAuthnPRFSupported) isMultiValue() bool {
    return false
}
