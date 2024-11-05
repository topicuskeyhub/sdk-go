package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReportErrorDetails struct {
    NonLinkable
    // The error property
    error *RequestModificationRequestReportError
    // The names property
    names []string
}
// NewRequestModificationRequestReportErrorDetails instantiates a new RequestModificationRequestReportErrorDetails and sets the default values.
func NewRequestModificationRequestReportErrorDetails()(*RequestModificationRequestReportErrorDetails) {
    m := &RequestModificationRequestReportErrorDetails{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReportErrorDetails"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReportErrorDetails(), nil
}
// GetError gets the error property value. The error property
// returns a *RequestModificationRequestReportError when successful
func (m *RequestModificationRequestReportErrorDetails) GetError()(*RequestModificationRequestReportError) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReportErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestReportError)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*RequestModificationRequestReportError))
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNames(res)
        }
        return nil
    }
    return res
}
// GetNames gets the names property value. The names property
// returns a []string when successful
func (m *RequestModificationRequestReportErrorDetails) GetNames()([]string) {
    return m.names
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReportErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetError() != nil {
        cast := (*m.GetError()).String()
        err = writer.WriteStringValue("error", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        err = writer.WriteCollectionOfStringValues("names", m.GetNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error property
func (m *RequestModificationRequestReportErrorDetails) SetError(value *RequestModificationRequestReportError)() {
    m.error = value
}
// SetNames sets the names property value. The names property
func (m *RequestModificationRequestReportErrorDetails) SetNames(value []string)() {
    m.names = value
}
type RequestModificationRequestReportErrorDetailsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*RequestModificationRequestReportError)
    GetNames()([]string)
    SetError(value *RequestModificationRequestReportError)()
    SetNames(value []string)()
}
