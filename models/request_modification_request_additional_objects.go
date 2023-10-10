package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestModificationRequest_additionalObjects 
type RequestModificationRequest_additionalObjects struct {
    // The acceptParams property
    acceptParams RequestAcceptModificationRequestParametersable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The resetStatus property
    resetStatus RequestPasswordResetRequestStatusable
}
// NewRequestModificationRequest_additionalObjects instantiates a new requestModificationRequest_additionalObjects and sets the default values.
func NewRequestModificationRequest_additionalObjects()(*RequestModificationRequest_additionalObjects) {
    m := &RequestModificationRequest_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequestModificationRequest_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestModificationRequest_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequest_additionalObjects(), nil
}
// GetAcceptParams gets the acceptParams property value. The acceptParams property
func (m *RequestModificationRequest_additionalObjects) GetAcceptParams()(RequestAcceptModificationRequestParametersable) {
    return m.acceptParams
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestModificationRequest_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *RequestModificationRequest_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestModificationRequest_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptParams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestAcceptModificationRequestParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptParams(val.(RequestAcceptModificationRequestParametersable))
        }
        return nil
    }
    res["audit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudit(val.(AuditInfoable))
        }
        return nil
    }
    res["resetStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestPasswordResetRequestStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetStatus(val.(RequestPasswordResetRequestStatusable))
        }
        return nil
    }
    return res
}
// GetResetStatus gets the resetStatus property value. The resetStatus property
func (m *RequestModificationRequest_additionalObjects) GetResetStatus()(RequestPasswordResetRequestStatusable) {
    return m.resetStatus
}
// Serialize serializes information the current object
func (m *RequestModificationRequest_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("acceptParams", m.GetAcceptParams())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resetStatus", m.GetResetStatus())
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
// SetAcceptParams sets the acceptParams property value. The acceptParams property
func (m *RequestModificationRequest_additionalObjects) SetAcceptParams(value RequestAcceptModificationRequestParametersable)() {
    m.acceptParams = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestModificationRequest_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *RequestModificationRequest_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetResetStatus sets the resetStatus property value. The resetStatus property
func (m *RequestModificationRequest_additionalObjects) SetResetStatus(value RequestPasswordResetRequestStatusable)() {
    m.resetStatus = value
}
// RequestModificationRequest_additionalObjectsable 
type RequestModificationRequest_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptParams()(RequestAcceptModificationRequestParametersable)
    GetAudit()(AuditInfoable)
    GetResetStatus()(RequestPasswordResetRequestStatusable)
    SetAcceptParams(value RequestAcceptModificationRequestParametersable)()
    SetAudit(value AuditInfoable)()
    SetResetStatus(value RequestPasswordResetRequestStatusable)()
}
