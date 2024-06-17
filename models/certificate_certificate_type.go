package models
type CertificateCertificateType int

const (
    PUBLIC_CERTIFICATE_CERTIFICATECERTIFICATETYPE CertificateCertificateType = iota
    PRIVATE_KEY_PAIR_CERTIFICATECERTIFICATETYPE
)

func (i CertificateCertificateType) String() string {
    return []string{"PUBLIC_CERTIFICATE", "PRIVATE_KEY_PAIR"}[i]
}
func ParseCertificateCertificateType(v string) (any, error) {
    result := PUBLIC_CERTIFICATE_CERTIFICATECERTIFICATETYPE
    switch v {
        case "PUBLIC_CERTIFICATE":
            result = PUBLIC_CERTIFICATE_CERTIFICATECERTIFICATETYPE
        case "PRIVATE_KEY_PAIR":
            result = PRIVATE_KEY_PAIR_CERTIFICATECERTIFICATETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCertificateCertificateType(values []CertificateCertificateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CertificateCertificateType) isMultiValue() bool {
    return false
}
