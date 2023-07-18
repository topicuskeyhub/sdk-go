package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceaccountServiceAccountPrimer 
type ServiceaccountServiceAccountPrimer struct {
    Linkable
    // The active property
    active *bool
    // The name property
    name *string
    // The system property
    system ProvisioningProvisionedSystemPrimerable
    // The username property
    username *string
    // The uuid property
    uuid *string
}
// NewServiceaccountServiceAccountPrimer instantiates a new serviceaccountServiceAccountPrimer and sets the default values.
func NewServiceaccountServiceAccountPrimer()(*ServiceaccountServiceAccountPrimer) {
    m := &ServiceaccountServiceAccountPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccountPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceaccountServiceAccountPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "serviceaccount.ServiceAccount":
                        return NewServiceaccountServiceAccount(), nil
                }
            }
        }
    }
    return NewServiceaccountServiceAccountPrimer(), nil
}
// GetActive gets the active property value. The active property
func (m *ServiceaccountServiceAccountPrimer) GetActive()(*bool) {
    return m.active
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceaccountServiceAccountPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
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
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *ServiceaccountServiceAccountPrimer) GetName()(*string) {
    return m.name
}
// GetSystem gets the system property value. The system property
func (m *ServiceaccountServiceAccountPrimer) GetSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.system
}
// GetUsername gets the username property value. The username property
func (m *ServiceaccountServiceAccountPrimer) GetUsername()(*string) {
    return m.username
}
// GetUuid gets the uuid property value. The uuid property
func (m *ServiceaccountServiceAccountPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("active", m.GetActive())
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
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. The active property
func (m *ServiceaccountServiceAccountPrimer) SetActive(value *bool)() {
    m.active = value
}
// SetName sets the name property value. The name property
func (m *ServiceaccountServiceAccountPrimer) SetName(value *string)() {
    m.name = value
}
// SetSystem sets the system property value. The system property
func (m *ServiceaccountServiceAccountPrimer) SetSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.system = value
}
// SetUsername sets the username property value. The username property
func (m *ServiceaccountServiceAccountPrimer) SetUsername(value *string)() {
    m.username = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *ServiceaccountServiceAccountPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// ServiceaccountServiceAccountPrimerable 
type ServiceaccountServiceAccountPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetName()(*string)
    GetSystem()(ProvisioningProvisionedSystemPrimerable)
    GetUsername()(*string)
    GetUuid()(*string)
    SetActive(value *bool)()
    SetName(value *string)()
    SetSystem(value ProvisioningProvisionedSystemPrimerable)()
    SetUsername(value *string)()
    SetUuid(value *string)()
}
