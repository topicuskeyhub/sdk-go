package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationModificationRequestNotification 
type NotificationModificationRequestNotification struct {
    NotificationNotification
    // The admins property
    admins []string
    // The groups property
    groups []string
    // The request property
    request RequestModificationRequestable
}
// NewNotificationModificationRequestNotification instantiates a new notificationModificationRequestNotification and sets the default values.
func NewNotificationModificationRequestNotification()(*NotificationModificationRequestNotification) {
    m := &NotificationModificationRequestNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.ModificationRequestNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationModificationRequestNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationModificationRequestNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationModificationRequestNotification(), nil
}
// GetAdmins gets the admins property value. The admins property
func (m *NotificationModificationRequestNotification) GetAdmins()([]string) {
    return m.admins
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationModificationRequestNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdmins(res)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestModificationRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(RequestModificationRequestable))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
func (m *NotificationModificationRequestNotification) GetGroups()([]string) {
    return m.groups
}
// GetRequest gets the request property value. The request property
func (m *NotificationModificationRequestNotification) GetRequest()(RequestModificationRequestable) {
    return m.request
}
// Serialize serializes information the current object
func (m *NotificationModificationRequestNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdmins() != nil {
        err = writer.WriteCollectionOfStringValues("admins", m.GetAdmins())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        err = writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("request", m.GetRequest())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdmins sets the admins property value. The admins property
func (m *NotificationModificationRequestNotification) SetAdmins(value []string)() {
    m.admins = value
}
// SetGroups sets the groups property value. The groups property
func (m *NotificationModificationRequestNotification) SetGroups(value []string)() {
    m.groups = value
}
// SetRequest sets the request property value. The request property
func (m *NotificationModificationRequestNotification) SetRequest(value RequestModificationRequestable)() {
    m.request = value
}
// NotificationModificationRequestNotificationable 
type NotificationModificationRequestNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmins()([]string)
    GetGroups()([]string)
    GetRequest()(RequestModificationRequestable)
    SetAdmins(value []string)()
    SetGroups(value []string)()
    SetRequest(value RequestModificationRequestable)()
}
