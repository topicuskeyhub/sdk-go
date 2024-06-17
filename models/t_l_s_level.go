package models
type TLSLevel int

const (
    VERIFIED_PINNED_TLSLEVEL TLSLevel = iota
    VERIFIED_TLSLEVEL
    SECURE_PINNED_TLSLEVEL
    SECURE_TLSLEVEL
    ENCRYPTED_TLSLEVEL
    UNSECURE_TLSLEVEL
)

func (i TLSLevel) String() string {
    return []string{"VERIFIED_PINNED", "VERIFIED", "SECURE_PINNED", "SECURE", "ENCRYPTED", "UNSECURE"}[i]
}
func ParseTLSLevel(v string) (any, error) {
    result := VERIFIED_PINNED_TLSLEVEL
    switch v {
        case "VERIFIED_PINNED":
            result = VERIFIED_PINNED_TLSLEVEL
        case "VERIFIED":
            result = VERIFIED_TLSLEVEL
        case "SECURE_PINNED":
            result = SECURE_PINNED_TLSLEVEL
        case "SECURE":
            result = SECURE_TLSLEVEL
        case "ENCRYPTED":
            result = ENCRYPTED_TLSLEVEL
        case "UNSECURE":
            result = UNSECURE_TLSLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTLSLevel(values []TLSLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TLSLevel) isMultiValue() bool {
    return false
}
