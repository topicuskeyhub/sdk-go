package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestJoinVaultRequest 
type RequestJoinVaultRequest struct {
    RequestModificationRequest
}
// NewRequestJoinVaultRequest instantiates a new requestJoinVaultRequest and sets the default values.
func NewRequestJoinVaultRequest()(*RequestJoinVaultRequest) {
    m := &RequestJoinVaultRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.JoinVaultRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestJoinVaultRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestJoinVaultRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestJoinVaultRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestJoinVaultRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestJoinVaultRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestJoinVaultRequestable 
type RequestJoinVaultRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
}
