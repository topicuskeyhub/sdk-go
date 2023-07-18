package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupClient 
type GroupGroupClient struct {
    Linkable
    // The activationRequired property
    activationRequired *bool
    // The additionalObjects property
    additionalObjects GroupGroupClient_additionalObjectsable
    // The client property
    client ClientClientApplicationPrimerable
    // The group property
    group GroupGroupPrimerable
    // The owner property
    owner GroupGroupPrimerable
    // The technicalAdministrator property
    technicalAdministrator GroupGroupPrimerable
}
// NewGroupGroupClient instantiates a new groupGroupClient and sets the default values.
func NewGroupGroupClient()(*GroupGroupClient) {
    m := &GroupGroupClient{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupClient"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupClient(), nil
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
func (m *GroupGroupClient) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupClient) GetAdditionalObjects()(GroupGroupClient_additionalObjectsable) {
    return m.additionalObjects
}
// GetClient gets the client property value. The client property
func (m *GroupGroupClient) GetClient()(ClientClientApplicationPrimerable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["activationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationRequired(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClient_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroupClient_additionalObjectsable))
        }
        return nil
    }
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(ClientClientApplicationPrimerable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["technicalAdministrator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnicalAdministrator(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *GroupGroupClient) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetOwner gets the owner property value. The owner property
func (m *GroupGroupClient) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetTechnicalAdministrator gets the technicalAdministrator property value. The technicalAdministrator property
func (m *GroupGroupClient) GetTechnicalAdministrator()(GroupGroupPrimerable) {
    return m.technicalAdministrator
}
// Serialize serializes information the current object
func (m *GroupGroupClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activationRequired", m.GetActivationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("technicalAdministrator", m.GetTechnicalAdministrator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationRequired sets the activationRequired property value. The activationRequired property
func (m *GroupGroupClient) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupClient) SetAdditionalObjects(value GroupGroupClient_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetClient sets the client property value. The client property
func (m *GroupGroupClient) SetClient(value ClientClientApplicationPrimerable)() {
    m.client = value
}
// SetGroup sets the group property value. The group property
func (m *GroupGroupClient) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetOwner sets the owner property value. The owner property
func (m *GroupGroupClient) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetTechnicalAdministrator sets the technicalAdministrator property value. The technicalAdministrator property
func (m *GroupGroupClient) SetTechnicalAdministrator(value GroupGroupPrimerable)() {
    m.technicalAdministrator = value
}
// GroupGroupClientable 
type GroupGroupClientable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationRequired()(*bool)
    GetAdditionalObjects()(GroupGroupClient_additionalObjectsable)
    GetClient()(ClientClientApplicationPrimerable)
    GetGroup()(GroupGroupPrimerable)
    GetOwner()(GroupGroupPrimerable)
    GetTechnicalAdministrator()(GroupGroupPrimerable)
    SetActivationRequired(value *bool)()
    SetAdditionalObjects(value GroupGroupClient_additionalObjectsable)()
    SetClient(value ClientClientApplicationPrimerable)()
    SetGroup(value GroupGroupPrimerable)()
    SetOwner(value GroupGroupPrimerable)()
    SetTechnicalAdministrator(value GroupGroupPrimerable)()
}
