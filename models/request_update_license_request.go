package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestUpdateLicenseRequest struct {
    RequestModificationRequest
    // The licenseKey property
    licenseKey *string
}
// NewRequestUpdateLicenseRequest instantiates a new RequestUpdateLicenseRequest and sets the default values.
func NewRequestUpdateLicenseRequest()(*RequestUpdateLicenseRequest) {
    m := &RequestUpdateLicenseRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.UpdateLicenseRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestUpdateLicenseRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestUpdateLicenseRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestUpdateLicenseRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestUpdateLicenseRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["licenseKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseKey(val)
        }
        return nil
    }
    return res
}
// GetLicenseKey gets the licenseKey property value. The licenseKey property
// returns a *string when successful
func (m *RequestUpdateLicenseRequest) GetLicenseKey()(*string) {
    return m.licenseKey
}
// Serialize serializes information the current object
func (m *RequestUpdateLicenseRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("licenseKey", m.GetLicenseKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLicenseKey sets the licenseKey property value. The licenseKey property
func (m *RequestUpdateLicenseRequest) SetLicenseKey(value *string)() {
    m.licenseKey = value
}
type RequestUpdateLicenseRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetLicenseKey()(*string)
    SetLicenseKey(value *string)()
}
