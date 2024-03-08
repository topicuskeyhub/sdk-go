package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionNumberSequence struct {
    Linkable
    // The accountCount property
    accountCount *int32
    // The additionalObjects property
    additionalObjects ProvisioningProvisionNumberSequence_additionalObjectsable
    // The name property
    name *string
    // The nextUID property
    nextUID *int64
}
// NewProvisioningProvisionNumberSequence instantiates a new ProvisioningProvisionNumberSequence and sets the default values.
func NewProvisioningProvisionNumberSequence()(*ProvisioningProvisionNumberSequence) {
    m := &ProvisioningProvisionNumberSequence{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisionNumberSequence"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionNumberSequenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionNumberSequenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionNumberSequence(), nil
}
// GetAccountCount gets the accountCount property value. The accountCount property
// returns a *int32 when successful
func (m *ProvisioningProvisionNumberSequence) GetAccountCount()(*int32) {
    return m.accountCount
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProvisioningProvisionNumberSequence_additionalObjectsable when successful
func (m *ProvisioningProvisionNumberSequence) GetAdditionalObjects()(ProvisioningProvisionNumberSequence_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionNumberSequence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accountCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountCount(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionNumberSequence_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProvisioningProvisionNumberSequence_additionalObjectsable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["nextUID"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextUID(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ProvisioningProvisionNumberSequence) GetName()(*string) {
    return m.name
}
// GetNextUID gets the nextUID property value. The nextUID property
// returns a *int64 when successful
func (m *ProvisioningProvisionNumberSequence) GetNextUID()(*int64) {
    return m.nextUID
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionNumberSequence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("nextUID", m.GetNextUID())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountCount sets the accountCount property value. The accountCount property
func (m *ProvisioningProvisionNumberSequence) SetAccountCount(value *int32)() {
    m.accountCount = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionNumberSequence) SetAdditionalObjects(value ProvisioningProvisionNumberSequence_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetName sets the name property value. The name property
func (m *ProvisioningProvisionNumberSequence) SetName(value *string)() {
    m.name = value
}
// SetNextUID sets the nextUID property value. The nextUID property
func (m *ProvisioningProvisionNumberSequence) SetNextUID(value *int64)() {
    m.nextUID = value
}
type ProvisioningProvisionNumberSequenceable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountCount()(*int32)
    GetAdditionalObjects()(ProvisioningProvisionNumberSequence_additionalObjectsable)
    GetName()(*string)
    GetNextUID()(*int64)
    SetAccountCount(value *int32)()
    SetAdditionalObjects(value ProvisioningProvisionNumberSequence_additionalObjectsable)()
    SetName(value *string)()
    SetNextUID(value *int64)()
}
