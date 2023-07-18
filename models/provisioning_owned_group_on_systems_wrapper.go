package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningOwnedGroupOnSystemsWrapper 
type ProvisioningOwnedGroupOnSystemsWrapper struct {
    ProvisioningGroupOnSystemLinkableWrapper
    // The unlinkedCount property
    unlinkedCount *int64
}
// NewProvisioningOwnedGroupOnSystemsWrapper instantiates a new provisioningOwnedGroupOnSystemsWrapper and sets the default values.
func NewProvisioningOwnedGroupOnSystemsWrapper()(*ProvisioningOwnedGroupOnSystemsWrapper) {
    m := &ProvisioningOwnedGroupOnSystemsWrapper{
        ProvisioningGroupOnSystemLinkableWrapper: *NewProvisioningGroupOnSystemLinkableWrapper(),
    }
    return m
}
// CreateProvisioningOwnedGroupOnSystemsWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningOwnedGroupOnSystemsWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningOwnedGroupOnSystemsWrapper(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningOwnedGroupOnSystemsWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningGroupOnSystemLinkableWrapper.GetFieldDeserializers()
    res["unlinkedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlinkedCount(val)
        }
        return nil
    }
    return res
}
// GetUnlinkedCount gets the unlinkedCount property value. The unlinkedCount property
func (m *ProvisioningOwnedGroupOnSystemsWrapper) GetUnlinkedCount()(*int64) {
    return m.unlinkedCount
}
// Serialize serializes information the current object
func (m *ProvisioningOwnedGroupOnSystemsWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningGroupOnSystemLinkableWrapper.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("unlinkedCount", m.GetUnlinkedCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUnlinkedCount sets the unlinkedCount property value. The unlinkedCount property
func (m *ProvisioningOwnedGroupOnSystemsWrapper) SetUnlinkedCount(value *int64)() {
    m.unlinkedCount = value
}
// ProvisioningOwnedGroupOnSystemsWrapperable 
type ProvisioningOwnedGroupOnSystemsWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningGroupOnSystemLinkableWrapperable
    GetUnlinkedCount()(*int64)
    SetUnlinkedCount(value *int64)()
}
