package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeValue_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The previous property
    previous IdentityAccountAttributeValueable
    // The selection property
    selection IdentityAccountAttributeValueSelectionable
}
// NewIdentityAccountAttributeValue_additionalObjects instantiates a new IdentityAccountAttributeValue_additionalObjects and sets the default values.
func NewIdentityAccountAttributeValue_additionalObjects()(*IdentityAccountAttributeValue_additionalObjects) {
    m := &IdentityAccountAttributeValue_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIdentityAccountAttributeValue_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeValue_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeValue_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IdentityAccountAttributeValue_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *IdentityAccountAttributeValue_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeValue_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["previous"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrevious(val.(IdentityAccountAttributeValueable))
        }
        return nil
    }
    res["selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeValueSelectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelection(val.(IdentityAccountAttributeValueSelectionable))
        }
        return nil
    }
    return res
}
// GetPrevious gets the previous property value. The previous property
// returns a IdentityAccountAttributeValueable when successful
func (m *IdentityAccountAttributeValue_additionalObjects) GetPrevious()(IdentityAccountAttributeValueable) {
    return m.previous
}
// GetSelection gets the selection property value. The selection property
// returns a IdentityAccountAttributeValueSelectionable when successful
func (m *IdentityAccountAttributeValue_additionalObjects) GetSelection()(IdentityAccountAttributeValueSelectionable) {
    return m.selection
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeValue_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previous", m.GetPrevious())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("selection", m.GetSelection())
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
func (m *IdentityAccountAttributeValue_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *IdentityAccountAttributeValue_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetPrevious sets the previous property value. The previous property
func (m *IdentityAccountAttributeValue_additionalObjects) SetPrevious(value IdentityAccountAttributeValueable)() {
    m.previous = value
}
// SetSelection sets the selection property value. The selection property
func (m *IdentityAccountAttributeValue_additionalObjects) SetSelection(value IdentityAccountAttributeValueSelectionable)() {
    m.selection = value
}
type IdentityAccountAttributeValue_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetPrevious()(IdentityAccountAttributeValueable)
    GetSelection()(IdentityAccountAttributeValueSelectionable)
    SetAudit(value AuditInfoable)()
    SetPrevious(value IdentityAccountAttributeValueable)()
    SetSelection(value IdentityAccountAttributeValueSelectionable)()
}