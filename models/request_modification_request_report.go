package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReport struct {
    NonLinkable
    // The conclusion property
    conclusion *RequestModificationRequestReportConclusion
    // The errorDetails property
    errorDetails []RequestModificationRequestReportErrorDetailsable
    // The groupChangeDetails property
    groupChangeDetails []RequestModificationRequestReportGroupChangeDetailsable
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
    res["groupChangeDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequestModificationRequestReportGroupChangeDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequestModificationRequestReportGroupChangeDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RequestModificationRequestReportGroupChangeDetailsable)
                }
            }
            m.SetGroupChangeDetails(res)
        }
        return nil
    }
    return res
}
// GetGroupChangeDetails gets the groupChangeDetails property value. The groupChangeDetails property
// returns a []RequestModificationRequestReportGroupChangeDetailsable when successful
func (m *RequestModificationRequestReport) GetGroupChangeDetails()([]RequestModificationRequestReportGroupChangeDetailsable) {
    return m.groupChangeDetails
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetGroupChangeDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupChangeDetails()))
        for i, v := range m.GetGroupChangeDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupChangeDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConclusion sets the conclusion property value. The conclusion property
func (m *RequestModificationRequestReport) SetConclusion(value *RequestModificationRequestReportConclusion)() {
    m.conclusion = value
}
// SetErrorDetails sets the errorDetails property value. The errorDetails property
func (m *RequestModificationRequestReport) SetErrorDetails(value []RequestModificationRequestReportErrorDetailsable)() {
    m.errorDetails = value
}
// SetGroupChangeDetails sets the groupChangeDetails property value. The groupChangeDetails property
func (m *RequestModificationRequestReport) SetGroupChangeDetails(value []RequestModificationRequestReportGroupChangeDetailsable)() {
    m.groupChangeDetails = value
}
type RequestModificationRequestReportable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConclusion()(*RequestModificationRequestReportConclusion)
    GetErrorDetails()([]RequestModificationRequestReportErrorDetailsable)
    GetGroupChangeDetails()([]RequestModificationRequestReportGroupChangeDetailsable)
    SetConclusion(value *RequestModificationRequestReportConclusion)()
    SetErrorDetails(value []RequestModificationRequestReportErrorDetailsable)()
    SetGroupChangeDetails(value []RequestModificationRequestReportGroupChangeDetailsable)()
}
