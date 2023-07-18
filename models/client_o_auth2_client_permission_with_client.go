package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientOAuth2ClientPermissionWithClient 
type ClientOAuth2ClientPermissionWithClient struct {
    ClientOAuth2ClientPermission
    // The client property
    client ClientOAuth2Clientable
}
// NewClientOAuth2ClientPermissionWithClient instantiates a new clientOAuth2ClientPermissionWithClient and sets the default values.
func NewClientOAuth2ClientPermissionWithClient()(*ClientOAuth2ClientPermissionWithClient) {
    m := &ClientOAuth2ClientPermissionWithClient{
        ClientOAuth2ClientPermission: *NewClientOAuth2ClientPermission(),
    }
    typeEscapedValue := "client.OAuth2ClientPermissionWithClient"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientOAuth2ClientPermissionWithClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientOAuth2ClientPermissionWithClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientOAuth2ClientPermissionWithClient(), nil
}
// GetClient gets the client property value. The client property
func (m *ClientOAuth2ClientPermissionWithClient) GetClient()(ClientOAuth2Clientable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientOAuth2ClientPermissionWithClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientOAuth2ClientPermission.GetFieldDeserializers()
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2ClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(ClientOAuth2Clientable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ClientOAuth2ClientPermissionWithClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClientOAuth2ClientPermission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClient sets the client property value. The client property
func (m *ClientOAuth2ClientPermissionWithClient) SetClient(value ClientOAuth2Clientable)() {
    m.client = value
}
// ClientOAuth2ClientPermissionWithClientable 
type ClientOAuth2ClientPermissionWithClientable interface {
    ClientOAuth2ClientPermissionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClient()(ClientOAuth2Clientable)
    SetClient(value ClientOAuth2Clientable)()
}
