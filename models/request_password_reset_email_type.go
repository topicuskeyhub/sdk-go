package models
type RequestPasswordResetEmailType int

const (
    NONE_REQUESTPASSWORDRESETEMAILTYPE RequestPasswordResetEmailType = iota
    REACTIVATION_REQUESTPASSWORDRESETEMAILTYPE
    EXTERNAL_REQUESTPASSWORDRESETEMAILTYPE
)

func (i RequestPasswordResetEmailType) String() string {
    return []string{"NONE", "REACTIVATION", "EXTERNAL"}[i]
}
func ParseRequestPasswordResetEmailType(v string) (any, error) {
    result := NONE_REQUESTPASSWORDRESETEMAILTYPE
    switch v {
        case "NONE":
            result = NONE_REQUESTPASSWORDRESETEMAILTYPE
        case "REACTIVATION":
            result = REACTIVATION_REQUESTPASSWORDRESETEMAILTYPE
        case "EXTERNAL":
            result = EXTERNAL_REQUESTPASSWORDRESETEMAILTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRequestPasswordResetEmailType(values []RequestPasswordResetEmailType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestPasswordResetEmailType) isMultiValue() bool {
    return false
}
