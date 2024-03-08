package models
import (
    "errors"
)
type VaultVaultRecordWarningPeriod int

const (
    AT_EXPIRATION_VAULTVAULTRECORDWARNINGPERIOD VaultVaultRecordWarningPeriod = iota
    TWO_WEEKS_VAULTVAULTRECORDWARNINGPERIOD
    ONE_MONTH_VAULTVAULTRECORDWARNINGPERIOD
    TWO_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
    THREE_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
    SIX_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
    NEVER_VAULTVAULTRECORDWARNINGPERIOD
)

func (i VaultVaultRecordWarningPeriod) String() string {
    return []string{"AT_EXPIRATION", "TWO_WEEKS", "ONE_MONTH", "TWO_MONTHS", "THREE_MONTHS", "SIX_MONTHS", "NEVER"}[i]
}
func ParseVaultVaultRecordWarningPeriod(v string) (any, error) {
    result := AT_EXPIRATION_VAULTVAULTRECORDWARNINGPERIOD
    switch v {
        case "AT_EXPIRATION":
            result = AT_EXPIRATION_VAULTVAULTRECORDWARNINGPERIOD
        case "TWO_WEEKS":
            result = TWO_WEEKS_VAULTVAULTRECORDWARNINGPERIOD
        case "ONE_MONTH":
            result = ONE_MONTH_VAULTVAULTRECORDWARNINGPERIOD
        case "TWO_MONTHS":
            result = TWO_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
        case "THREE_MONTHS":
            result = THREE_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
        case "SIX_MONTHS":
            result = SIX_MONTHS_VAULTVAULTRECORDWARNINGPERIOD
        case "NEVER":
            result = NEVER_VAULTVAULTRECORDWARNINGPERIOD
        default:
            return 0, errors.New("Unknown VaultVaultRecordWarningPeriod value: " + v)
    }
    return &result, nil
}
func SerializeVaultVaultRecordWarningPeriod(values []VaultVaultRecordWarningPeriod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VaultVaultRecordWarningPeriod) isMultiValue() bool {
    return false
}
