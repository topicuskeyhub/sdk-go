package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedAccount 
type ProvisioningProvisionedAccount struct {
    AuthAccountPrimer
    // The additionalObjects property
    additionalObjects ProvisioningProvisionedAccount_additionalObjectsable
    // The uid property
    uid *int64
}
// NewProvisioningProvisionedAccount instantiates a new provisioningProvisionedAccount and sets the default values.
func NewProvisioningProvisionedAccount()(*ProvisioningProvisionedAccount) {
    m := &ProvisioningProvisionedAccount{
        AuthAccountPrimer: *NewAuthAccountPrimer(),
    }
    typeEscapedValue := "provisioning.ProvisionedAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedAccount(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionedAccount) GetAdditionalObjects()(ProvisioningProvisionedAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccountPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProvisioningProvisionedAccount_additionalObjectsable))
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    return res
}
// GetUid gets the uid property value. The uid property
func (m *ProvisioningProvisionedAccount) GetUid()(*int64) {
    return m.uid
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthAccountPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionedAccount) SetAdditionalObjects(value ProvisioningProvisionedAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetUid sets the uid property value. The uid property
func (m *ProvisioningProvisionedAccount) SetUid(value *int64)() {
    m.uid = value
}
// ProvisioningProvisionedAccountable 
type ProvisioningProvisionedAccountable interface {
    AuthAccountPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(ProvisioningProvisionedAccount_additionalObjectsable)
    GetUid()(*int64)
    SetAdditionalObjects(value ProvisioningProvisionedAccount_additionalObjectsable)()
    SetUid(value *int64)()
}
