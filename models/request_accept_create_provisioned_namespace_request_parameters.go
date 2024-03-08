package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAcceptCreateProvisionedNamespaceRequestParameters struct {
    RequestAcceptModificationRequestParameters
    // The groupDN property
    groupDN *string
    // The serviceAccountDN property
    serviceAccountDN *string
}
// NewRequestAcceptCreateProvisionedNamespaceRequestParameters instantiates a new RequestAcceptCreateProvisionedNamespaceRequestParameters and sets the default values.
func NewRequestAcceptCreateProvisionedNamespaceRequestParameters()(*RequestAcceptCreateProvisionedNamespaceRequestParameters) {
    m := &RequestAcceptCreateProvisionedNamespaceRequestParameters{
        RequestAcceptModificationRequestParameters: *NewRequestAcceptModificationRequestParameters(),
    }
    typeEscapedValue := "request.AcceptCreateProvisionedNamespaceRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptCreateProvisionedNamespaceRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAcceptCreateProvisionedNamespaceRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptCreateProvisionedNamespaceRequestParameters(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAcceptModificationRequestParameters.GetFieldDeserializers()
    res["groupDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDN(val)
        }
        return nil
    }
    res["serviceAccountDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccountDN(val)
        }
        return nil
    }
    return res
}
// GetGroupDN gets the groupDN property value. The groupDN property
// returns a *string when successful
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) GetGroupDN()(*string) {
    return m.groupDN
}
// GetServiceAccountDN gets the serviceAccountDN property value. The serviceAccountDN property
// returns a *string when successful
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) GetServiceAccountDN()(*string) {
    return m.serviceAccountDN
}
// Serialize serializes information the current object
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAcceptModificationRequestParameters.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupDN", m.GetGroupDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceAccountDN", m.GetServiceAccountDN())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupDN sets the groupDN property value. The groupDN property
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) SetGroupDN(value *string)() {
    m.groupDN = value
}
// SetServiceAccountDN sets the serviceAccountDN property value. The serviceAccountDN property
func (m *RequestAcceptCreateProvisionedNamespaceRequestParameters) SetServiceAccountDN(value *string)() {
    m.serviceAccountDN = value
}
type RequestAcceptCreateProvisionedNamespaceRequestParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAcceptModificationRequestParametersable
    GetGroupDN()(*string)
    GetServiceAccountDN()(*string)
    SetGroupDN(value *string)()
    SetServiceAccountDN(value *string)()
}
