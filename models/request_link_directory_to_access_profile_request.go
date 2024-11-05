package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestLinkDirectoryToAccessProfileRequest struct {
    RequestAbstractAccessProfileModificationRequest
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
}
// NewRequestLinkDirectoryToAccessProfileRequest instantiates a new RequestLinkDirectoryToAccessProfileRequest and sets the default values.
func NewRequestLinkDirectoryToAccessProfileRequest()(*RequestLinkDirectoryToAccessProfileRequest) {
    m := &RequestLinkDirectoryToAccessProfileRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.LinkDirectoryToAccessProfileRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestLinkDirectoryToAccessProfileRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestLinkDirectoryToAccessProfileRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestLinkDirectoryToAccessProfileRequest(), nil
}
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *RequestLinkDirectoryToAccessProfileRequest) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestLinkDirectoryToAccessProfileRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectoryPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(DirectoryAccountDirectoryPrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestLinkDirectoryToAccessProfileRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectory sets the directory property value. The directory property
func (m *RequestLinkDirectoryToAccessProfileRequest) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
type RequestLinkDirectoryToAccessProfileRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
}
