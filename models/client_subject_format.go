package models
import (
    "errors"
)
// 
type ClientSubjectFormat int

const (
    ID_CLIENTSUBJECTFORMAT ClientSubjectFormat = iota
    UPN_CLIENTSUBJECTFORMAT
    USERNAME_CLIENTSUBJECTFORMAT
    EMAIL_CLIENTSUBJECTFORMAT
)

func (i ClientSubjectFormat) String() string {
    return []string{"ID", "UPN", "USERNAME", "EMAIL"}[i]
}
func ParseClientSubjectFormat(v string) (any, error) {
    result := ID_CLIENTSUBJECTFORMAT
    switch v {
        case "ID":
            result = ID_CLIENTSUBJECTFORMAT
        case "UPN":
            result = UPN_CLIENTSUBJECTFORMAT
        case "USERNAME":
            result = USERNAME_CLIENTSUBJECTFORMAT
        case "EMAIL":
            result = EMAIL_CLIENTSUBJECTFORMAT
        default:
            return 0, errors.New("Unknown ClientSubjectFormat value: " + v)
    }
    return &result, nil
}
func SerializeClientSubjectFormat(values []ClientSubjectFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClientSubjectFormat) isMultiValue() bool {
    return false
}
