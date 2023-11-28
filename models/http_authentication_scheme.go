package models
import (
    "errors"
)
// 
type HttpAuthenticationScheme int

const (
    NONE_HTTPAUTHENTICATIONSCHEME HttpAuthenticationScheme = iota
    BASIC_HTTPAUTHENTICATIONSCHEME
    BEARER_HTTPAUTHENTICATIONSCHEME
    CUSTOM_HTTPAUTHENTICATIONSCHEME
)

func (i HttpAuthenticationScheme) String() string {
    return []string{"NONE", "BASIC", "BEARER", "CUSTOM"}[i]
}
func ParseHttpAuthenticationScheme(v string) (any, error) {
    result := NONE_HTTPAUTHENTICATIONSCHEME
    switch v {
        case "NONE":
            result = NONE_HTTPAUTHENTICATIONSCHEME
        case "BASIC":
            result = BASIC_HTTPAUTHENTICATIONSCHEME
        case "BEARER":
            result = BEARER_HTTPAUTHENTICATIONSCHEME
        case "CUSTOM":
            result = CUSTOM_HTTPAUTHENTICATIONSCHEME
        default:
            return 0, errors.New("Unknown HttpAuthenticationScheme value: " + v)
    }
    return &result, nil
}
func SerializeHttpAuthenticationScheme(values []HttpAuthenticationScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HttpAuthenticationScheme) isMultiValue() bool {
    return false
}
