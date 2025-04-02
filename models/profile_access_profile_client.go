package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileClient struct {
    Linkable
    // The accessProfile property
    accessProfile ProfileAccessProfilePrimerable
    // The additionalObjects property
    additionalObjects ProfileAccessProfileClient_additionalObjectsable
    // The client property
    client ClientClientApplicationPrimerable
}
// NewProfileAccessProfileClient instantiates a new ProfileAccessProfileClient and sets the default values.
func NewProfileAccessProfileClient()(*ProfileAccessProfileClient) {
    m := &ProfileAccessProfileClient{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "profile.AccessProfileClient"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileClient(), nil
}
// GetAccessProfile gets the accessProfile property value. The accessProfile property
// returns a ProfileAccessProfilePrimerable when successful
func (m *ProfileAccessProfileClient) GetAccessProfile()(ProfileAccessProfilePrimerable) {
    return m.accessProfile
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProfileAccessProfileClient_additionalObjectsable when successful
func (m *ProfileAccessProfileClient) GetAdditionalObjects()(ProfileAccessProfileClient_additionalObjectsable) {
    return m.additionalObjects
}
// GetClient gets the client property value. The client property
// returns a ClientClientApplicationPrimerable when successful
func (m *ProfileAccessProfileClient) GetClient()(ClientClientApplicationPrimerable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accessProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfilePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfile(val.(ProfileAccessProfilePrimerable))
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfileClient_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProfileAccessProfileClient_additionalObjectsable))
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
    return res
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessProfile", m.GetAccessProfile())
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
    return nil
}
// SetAccessProfile sets the accessProfile property value. The accessProfile property
func (m *ProfileAccessProfileClient) SetAccessProfile(value ProfileAccessProfilePrimerable)() {
    m.accessProfile = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfileClient) SetAdditionalObjects(value ProfileAccessProfileClient_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetClient sets the client property value. The client property
func (m *ProfileAccessProfileClient) SetClient(value ClientClientApplicationPrimerable)() {
    m.client = value
}
type ProfileAccessProfileClientable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessProfile()(ProfileAccessProfilePrimerable)
    GetAdditionalObjects()(ProfileAccessProfileClient_additionalObjectsable)
    GetClient()(ClientClientApplicationPrimerable)
    SetAccessProfile(value ProfileAccessProfilePrimerable)()
    SetAdditionalObjects(value ProfileAccessProfileClient_additionalObjectsable)()
    SetClient(value ClientClientApplicationPrimerable)()
}
