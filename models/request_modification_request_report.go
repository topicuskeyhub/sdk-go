package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReport struct {
    NonLinkable
    // The changeDetails property
    changeDetails []RequestModificationRequestReportObjectChangeDetailsable
    // The conclusion property
    conclusion *RequestModificationRequestReportConclusion
    // The errorDetails property
    errorDetails []RequestModificationRequestReportErrorDetailsable
}
// NewRequestModificationRequestReport instantiates a new RequestModificationRequestReport and sets the default values.
func NewRequestModificationRequestReport()(*RequestModificationRequestReport) {
    m := &RequestModificationRequestReport{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReport"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReport(), nil
}
// GetChangeDetails gets the changeDetails property value. The changeDetails property
// returns a []RequestModificationRequestReportObjectChangeDetailsable when successful
func (m *RequestModificationRequestReport) GetChangeDetails()([]RequestModificationRequestReportObjectChangeDetailsable) {
    return m.changeDetails
}
// GetConclusion gets the conclusion property value. The conclusion property
// returns a *RequestModificationRequestReportConclusion when successful
func (m *RequestModificationRequestReport) GetConclusion()(*RequestModificationRequestReportConclusion) {
    return m.conclusion
}
// GetErrorDetails gets the errorDetails property value. The errorDetails property
// returns a []RequestModificationRequestReportErrorDetailsable when successful
func (m *RequestModificationRequestReport) GetErrorDetails()([]RequestModificationRequestReportErrorDetailsable) {
    return m.errorDetails
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["changeDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequestModificationRequestReportObjectChangeDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequestModificationRequestReportObjectChangeDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RequestModificationRequestReportObjectChangeDetailsable)
                }
            }
            m.SetChangeDetails(res)
        }
        return nil
    }
    res["conclusion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestReportConclusion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val.(*RequestModificationRequestReportConclusion))
        }
        return nil
    }
    res["errorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequestModificationRequestReportErrorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequestModificationRequestReportErrorDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RequestModificationRequestReportErrorDetailsable)
                }
            }
            m.SetErrorDetails(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChangeDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChangeDetails()))
        for i, v := range m.GetChangeDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("changeDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConclusion() != nil {
        cast := (*m.GetConclusion()).String()
        err = writer.WriteStringValue("conclusion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetErrorDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrorDetails()))
        for i, v := range m.GetErrorDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("errorDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeDetails sets the changeDetails property value. The changeDetails property
func (m *RequestModificationRequestReport) SetChangeDetails(value []RequestModificationRequestReportObjectChangeDetailsable)() {
    m.changeDetails = value
}
// SetConclusion sets the conclusion property value. The conclusion property
func (m *RequestModificationRequestReport) SetConclusion(value *RequestModificationRequestReportConclusion)() {
    m.conclusion = value
}
// SetErrorDetails sets the errorDetails property value. The errorDetails property
func (m *RequestModificationRequestReport) SetErrorDetails(value []RequestModificationRequestReportErrorDetailsable)() {
    m.errorDetails = value
}
type RequestModificationRequestReportable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeDetails()([]RequestModificationRequestReportObjectChangeDetailsable)
    GetConclusion()(*RequestModificationRequestReportConclusion)
    GetErrorDetails()([]RequestModificationRequestReportErrorDetailsable)
    SetChangeDetails(value []RequestModificationRequestReportObjectChangeDetailsable)()
    SetConclusion(value *RequestModificationRequestReportConclusion)()
    SetErrorDetails(value []RequestModificationRequestReportErrorDetailsable)()
}
