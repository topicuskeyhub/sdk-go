package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ErrorReport 
type ErrorReport struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The applicationError property
    applicationError *string
    // The applicationErrorParameters property
    applicationErrorParameters ErrorReport_applicationErrorParametersable
    // The code property
    code *int32
    // The errorDetails property
    errorDetails []string
    // The exception property
    exception *string
    // The message property
    message *string
    // The reason property
    reason *string
    // The stacktrace property
    stacktrace []string
}
// NewErrorReport instantiates a new ErrorReport and sets the default values.
func NewErrorReport()(*ErrorReport) {
    m := &ErrorReport{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateErrorReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateErrorReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewErrorReport(), nil
}
// Error the primary error message.
func (m *ErrorReport) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ErrorReport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationError gets the applicationError property value. The applicationError property
func (m *ErrorReport) GetApplicationError()(*string) {
    return m.applicationError
}
// GetApplicationErrorParameters gets the applicationErrorParameters property value. The applicationErrorParameters property
func (m *ErrorReport) GetApplicationErrorParameters()(ErrorReport_applicationErrorParametersable) {
    return m.applicationErrorParameters
}
// GetCode gets the code property value. The code property
func (m *ErrorReport) GetCode()(*int32) {
    return m.code
}
// GetErrorDetails gets the errorDetails property value. The errorDetails property
func (m *ErrorReport) GetErrorDetails()([]string) {
    return m.errorDetails
}
// GetException gets the exception property value. The exception property
func (m *ErrorReport) GetException()(*string) {
    return m.exception
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ErrorReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationError(val)
        }
        return nil
    }
    res["applicationErrorParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateErrorReport_applicationErrorParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationErrorParameters(val.(ErrorReport_applicationErrorParametersable))
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["errorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetErrorDetails(res)
        }
        return nil
    }
    res["exception"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetException(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["stacktrace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetStacktrace(res)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *ErrorReport) GetMessage()(*string) {
    return m.message
}
// GetReason gets the reason property value. The reason property
func (m *ErrorReport) GetReason()(*string) {
    return m.reason
}
// GetStacktrace gets the stacktrace property value. The stacktrace property
func (m *ErrorReport) GetStacktrace()([]string) {
    return m.stacktrace
}
// Serialize serializes information the current object
func (m *ErrorReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationError", m.GetApplicationError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("applicationErrorParameters", m.GetApplicationErrorParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    if m.GetErrorDetails() != nil {
        err := writer.WriteCollectionOfStringValues("errorDetails", m.GetErrorDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("exception", m.GetException())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetStacktrace() != nil {
        err := writer.WriteCollectionOfStringValues("stacktrace", m.GetStacktrace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ErrorReport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationError sets the applicationError property value. The applicationError property
func (m *ErrorReport) SetApplicationError(value *string)() {
    m.applicationError = value
}
// SetApplicationErrorParameters sets the applicationErrorParameters property value. The applicationErrorParameters property
func (m *ErrorReport) SetApplicationErrorParameters(value ErrorReport_applicationErrorParametersable)() {
    m.applicationErrorParameters = value
}
// SetCode sets the code property value. The code property
func (m *ErrorReport) SetCode(value *int32)() {
    m.code = value
}
// SetErrorDetails sets the errorDetails property value. The errorDetails property
func (m *ErrorReport) SetErrorDetails(value []string)() {
    m.errorDetails = value
}
// SetException sets the exception property value. The exception property
func (m *ErrorReport) SetException(value *string)() {
    m.exception = value
}
// SetMessage sets the message property value. The message property
func (m *ErrorReport) SetMessage(value *string)() {
    m.message = value
}
// SetReason sets the reason property value. The reason property
func (m *ErrorReport) SetReason(value *string)() {
    m.reason = value
}
// SetStacktrace sets the stacktrace property value. The stacktrace property
func (m *ErrorReport) SetStacktrace(value []string)() {
    m.stacktrace = value
}
// ErrorReportable 
type ErrorReportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationError()(*string)
    GetApplicationErrorParameters()(ErrorReport_applicationErrorParametersable)
    GetCode()(*int32)
    GetErrorDetails()([]string)
    GetException()(*string)
    GetMessage()(*string)
    GetReason()(*string)
    GetStacktrace()([]string)
    SetApplicationError(value *string)()
    SetApplicationErrorParameters(value ErrorReport_applicationErrorParametersable)()
    SetCode(value *int32)()
    SetErrorDetails(value []string)()
    SetException(value *string)()
    SetMessage(value *string)()
    SetReason(value *string)()
    SetStacktrace(value []string)()
}
