package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedNamespace struct {
    ProvisioningProvisionedSystem
    // The baseSystem property
    baseSystem ProvisioningProvisionedSystemPrimerable
    // The groupDN property
    groupDN *string
    // The serviceAccountDN property
    serviceAccountDN *string
}
// NewProvisioningProvisionedNamespace instantiates a new ProvisioningProvisionedNamespace and sets the default values.
func NewProvisioningProvisionedNamespace()(*ProvisioningProvisionedNamespace) {
    m := &ProvisioningProvisionedNamespace{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedNamespace"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedNamespaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedNamespaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedNamespace(), nil
}
// GetBaseSystem gets the baseSystem property value. The baseSystem property
// returns a ProvisioningProvisionedSystemPrimerable when successful
func (m *ProvisioningProvisionedNamespace) GetBaseSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.baseSystem
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedNamespace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
    res["baseSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    res["groupDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDN(val)
        }
        return nil
    }
    res["serviceAccountDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccountDN(val)
        }
        return nil
    }
    return res
}
// GetGroupDN gets the groupDN property value. The groupDN property
// returns a *string when successful
func (m *ProvisioningProvisionedNamespace) GetGroupDN()(*string) {
    return m.groupDN
}
// GetServiceAccountDN gets the serviceAccountDN property value. The serviceAccountDN property
// returns a *string when successful
func (m *ProvisioningProvisionedNamespace) GetServiceAccountDN()(*string) {
    return m.serviceAccountDN
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedNamespace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("baseSystem", m.GetBaseSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupDN", m.GetGroupDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceAccountDN", m.GetServiceAccountDN())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseSystem sets the baseSystem property value. The baseSystem property
func (m *ProvisioningProvisionedNamespace) SetBaseSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.baseSystem = value
}
// SetGroupDN sets the groupDN property value. The groupDN property
func (m *ProvisioningProvisionedNamespace) SetGroupDN(value *string)() {
    m.groupDN = value
}
// SetServiceAccountDN sets the serviceAccountDN property value. The serviceAccountDN property
func (m *ProvisioningProvisionedNamespace) SetServiceAccountDN(value *string)() {
    m.serviceAccountDN = value
}
type ProvisioningProvisionedNamespaceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetBaseSystem()(ProvisioningProvisionedSystemPrimerable)
    GetGroupDN()(*string)
    GetServiceAccountDN()(*string)
    SetBaseSystem(value ProvisioningProvisionedSystemPrimerable)()
    SetGroupDN(value *string)()
    SetServiceAccountDN(value *string)()
}
