package models
import (
    "errors"
)
type AuthPermittedOperation int

const (
    CREATE_AUTHPERMITTEDOPERATION AuthPermittedOperation = iota
    READ_AUTHPERMITTEDOPERATION
    UPDATE_AUTHPERMITTEDOPERATION
    DELETE_AUTHPERMITTEDOPERATION
)

func (i AuthPermittedOperation) String() string {
    return []string{"CREATE", "READ", "UPDATE", "DELETE"}[i]
}
func ParseAuthPermittedOperation(v string) (any, error) {
    result := CREATE_AUTHPERMITTEDOPERATION
    switch v {
        case "CREATE":
            result = CREATE_AUTHPERMITTEDOPERATION
        case "READ":
            result = READ_AUTHPERMITTEDOPERATION
        case "UPDATE":
            result = UPDATE_AUTHPERMITTEDOPERATION
        case "DELETE":
            result = DELETE_AUTHPERMITTEDOPERATION
        default:
            return 0, errors.New("Unknown AuthPermittedOperation value: " + v)
    }
    return &result, nil
}
func SerializeAuthPermittedOperation(values []AuthPermittedOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthPermittedOperation) isMultiValue() bool {
    return false
}
