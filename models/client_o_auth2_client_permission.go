package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientOAuth2ClientPermission 
type ClientOAuth2ClientPermission struct {
    Linkable
    // The additionalObjects property
    additionalObjects ClientOAuth2ClientPermission_additionalObjectsable
    // The forGroup property
    forGroup GroupGroupPrimerable
    // The forSystem property
    forSystem ProvisioningProvisionedSystemPrimerable
    // The value property
    value *ClientOAuth2ClientPermissionType
}
// NewClientOAuth2ClientPermission instantiates a new clientOAuth2ClientPermission and sets the default values.
func NewClientOAuth2ClientPermission()(*ClientOAuth2ClientPermission) {
    m := &ClientOAuth2ClientPermission{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "client.OAuth2ClientPermission"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientOAuth2ClientPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientOAuth2ClientPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "client.OAuth2ClientPermissionWithClient":
                        return NewClientOAuth2ClientPermissionWithClient(), nil
                }
            }
        }
    }
    return NewClientOAuth2ClientPermission(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *ClientOAuth2ClientPermission) GetAdditionalObjects()(ClientOAuth2ClientPermission_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientOAuth2ClientPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2ClientPermission_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ClientOAuth2ClientPermission_additionalObjectsable))
        }
        return nil
    }
    res["forGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["forSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientOAuth2ClientPermissionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*ClientOAuth2ClientPermissionType))
        }
        return nil
    }
    return res
}
// GetForGroup gets the forGroup property value. The forGroup property
func (m *ClientOAuth2ClientPermission) GetForGroup()(GroupGroupPrimerable) {
    return m.forGroup
}
// GetForSystem gets the forSystem property value. The forSystem property
func (m *ClientOAuth2ClientPermission) GetForSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.forSystem
}
// GetValue gets the value property value. The value property
func (m *ClientOAuth2ClientPermission) GetValue()(*ClientOAuth2ClientPermissionType) {
    return m.value
}
// Serialize serializes information the current object
func (m *ClientOAuth2ClientPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("forGroup", m.GetForGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("forSystem", m.GetForSystem())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
        err = writer.WriteStringValue("value", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ClientOAuth2ClientPermission) SetAdditionalObjects(value ClientOAuth2ClientPermission_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetForGroup sets the forGroup property value. The forGroup property
func (m *ClientOAuth2ClientPermission) SetForGroup(value GroupGroupPrimerable)() {
    m.forGroup = value
}
// SetForSystem sets the forSystem property value. The forSystem property
func (m *ClientOAuth2ClientPermission) SetForSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.forSystem = value
}
// SetValue sets the value property value. The value property
func (m *ClientOAuth2ClientPermission) SetValue(value *ClientOAuth2ClientPermissionType)() {
    m.value = value
}
// ClientOAuth2ClientPermissionable 
type ClientOAuth2ClientPermissionable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(ClientOAuth2ClientPermission_additionalObjectsable)
    GetForGroup()(GroupGroupPrimerable)
    GetForSystem()(ProvisioningProvisionedSystemPrimerable)
    GetValue()(*ClientOAuth2ClientPermissionType)
    SetAdditionalObjects(value ClientOAuth2ClientPermission_additionalObjectsable)()
    SetForGroup(value GroupGroupPrimerable)()
    SetForSystem(value ProvisioningProvisionedSystemPrimerable)()
    SetValue(value *ClientOAuth2ClientPermissionType)()
}
