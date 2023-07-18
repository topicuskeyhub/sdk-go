package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestReviewAuditRequest 
type RequestReviewAuditRequest struct {
    RequestModificationRequest
}
// NewRequestReviewAuditRequest instantiates a new requestReviewAuditRequest and sets the default values.
func NewRequestReviewAuditRequest()(*RequestReviewAuditRequest) {
    m := &RequestReviewAuditRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.ReviewAuditRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestReviewAuditRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestReviewAuditRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestReviewAuditRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestReviewAuditRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestReviewAuditRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestReviewAuditRequestable 
type RequestReviewAuditRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
}
