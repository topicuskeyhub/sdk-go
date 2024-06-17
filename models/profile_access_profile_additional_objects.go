package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfile_additionalObjects struct {
    // The accounts property
    accounts ProfileAccessProfileAccountLinkableWrapperable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
}
// NewProfileAccessProfile_additionalObjects instantiates a new ProfileAccessProfile_additionalObjects and sets the default values.
func NewProfileAccessProfile_additionalObjects()(*ProfileAccessProfile_additionalObjects) {
    m := &ProfileAccessProfile_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProfileAccessProfile_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfile_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfile_additionalObjects(), nil
}
// GetAccounts gets the accounts property value. The accounts property
// returns a ProfileAccessProfileAccountLinkableWrapperable when successful
func (m *ProfileAccessProfile_additionalObjects) GetAccounts()(ProfileAccessProfileAccountLinkableWrapperable) {
    return m.accounts
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProfileAccessProfile_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *ProfileAccessProfile_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfile_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfileAccountLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccounts(val.(ProfileAccessProfileAccountLinkableWrapperable))
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
    return res
}
// Serialize serializes information the current object
func (m *ProfileAccessProfile_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accounts", m.GetAccounts())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *ProfileAccessProfile_additionalObjects) SetAccounts(value ProfileAccessProfileAccountLinkableWrapperable)() {
    m.accounts = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProfileAccessProfile_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ProfileAccessProfile_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
type ProfileAccessProfile_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()(ProfileAccessProfileAccountLinkableWrapperable)
    GetAudit()(AuditInfoable)
    SetAccounts(value ProfileAccessProfileAccountLinkableWrapperable)()
    SetAudit(value AuditInfoable)()
}
