package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReportChangeDetails struct {
    NonLinkable
    // The change property
    change *RequestModificationRequestReportChange
    // The objectName property
    objectName *string
    // The subjectName property
    subjectName *string
}
// NewRequestModificationRequestReportChangeDetails instantiates a new RequestModificationRequestReportChangeDetails and sets the default values.
func NewRequestModificationRequestReportChangeDetails()(*RequestModificationRequestReportChangeDetails) {
    m := &RequestModificationRequestReportChangeDetails{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReportChangeDetails"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportChangeDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportChangeDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReportChangeDetails(), nil
}
// GetChange gets the change property value. The change property
// returns a *RequestModificationRequestReportChange when successful
func (m *RequestModificationRequestReportChangeDetails) GetChange()(*RequestModificationRequestReportChange) {
    return m.change
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReportChangeDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["change"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestReportChange)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChange(val.(*RequestModificationRequestReportChange))
        }
        return nil
    }
    res["objectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectName(val)
        }
        return nil
    }
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    return res
}
// GetObjectName gets the objectName property value. The objectName property
// returns a *string when successful
func (m *RequestModificationRequestReportChangeDetails) GetObjectName()(*string) {
    return m.objectName
}
// GetSubjectName gets the subjectName property value. The subjectName property
// returns a *string when successful
func (m *RequestModificationRequestReportChangeDetails) GetSubjectName()(*string) {
    return m.subjectName
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReportChangeDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChange() != nil {
        cast := (*m.GetChange()).String()
        err = writer.WriteStringValue("change", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectName", m.GetObjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChange sets the change property value. The change property
func (m *RequestModificationRequestReportChangeDetails) SetChange(value *RequestModificationRequestReportChange)() {
    m.change = value
}
// SetObjectName sets the objectName property value. The objectName property
func (m *RequestModificationRequestReportChangeDetails) SetObjectName(value *string)() {
    m.objectName = value
}
// SetSubjectName sets the subjectName property value. The subjectName property
func (m *RequestModificationRequestReportChangeDetails) SetSubjectName(value *string)() {
    m.subjectName = value
}
type RequestModificationRequestReportChangeDetailsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChange()(*RequestModificationRequestReportChange)
    GetObjectName()(*string)
    GetSubjectName()(*string)
    SetChange(value *RequestModificationRequestReportChange)()
    SetObjectName(value *string)()
    SetSubjectName(value *string)()
}
