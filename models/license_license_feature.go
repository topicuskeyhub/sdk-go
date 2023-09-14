package models
import (
    "errors"
)
// 
type LicenseLicenseFeature int

const (
    CLUSTER_HA_LICENSELICENSEFEATURE LicenseLicenseFeature = iota
    CLUSTER_HA_ENTERPRISE_LICENSELICENSEFEATURE
    WORKFLOWS_LICENSELICENSEFEATURE
    NESTED_GROUPS_LICENSELICENSEFEATURE
    ENTERPRISE_ORGANISATION_LICENSELICENSEFEATURE
    COMPLIANCE_PLUS_LICENSELICENSEFEATURE
    OFFLINE_MODE_LICENSELICENSEFEATURE
    SERVICE_ACCOUNTS_LICENSELICENSEFEATURE
)

func (i LicenseLicenseFeature) String() string {
    return []string{"CLUSTER_HA", "CLUSTER_HA_ENTERPRISE", "WORKFLOWS", "NESTED_GROUPS", "ENTERPRISE_ORGANISATION", "COMPLIANCE_PLUS", "OFFLINE_MODE", "SERVICE_ACCOUNTS"}[i]
}
func ParseLicenseLicenseFeature(v string) (any, error) {
    result := CLUSTER_HA_LICENSELICENSEFEATURE
    switch v {
        case "CLUSTER_HA":
            result = CLUSTER_HA_LICENSELICENSEFEATURE
        case "CLUSTER_HA_ENTERPRISE":
            result = CLUSTER_HA_ENTERPRISE_LICENSELICENSEFEATURE
        case "WORKFLOWS":
            result = WORKFLOWS_LICENSELICENSEFEATURE
        case "NESTED_GROUPS":
            result = NESTED_GROUPS_LICENSELICENSEFEATURE
        case "ENTERPRISE_ORGANISATION":
            result = ENTERPRISE_ORGANISATION_LICENSELICENSEFEATURE
        case "COMPLIANCE_PLUS":
            result = COMPLIANCE_PLUS_LICENSELICENSEFEATURE
        case "OFFLINE_MODE":
            result = OFFLINE_MODE_LICENSELICENSEFEATURE
        case "SERVICE_ACCOUNTS":
            result = SERVICE_ACCOUNTS_LICENSELICENSEFEATURE
        default:
            return 0, errors.New("Unknown LicenseLicenseFeature value: " + v)
    }
    return &result, nil
}
func SerializeLicenseLicenseFeature(values []LicenseLicenseFeature) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LicenseLicenseFeature) isMultiValue() bool {
    return false
}
