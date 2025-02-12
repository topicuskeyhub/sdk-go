package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientClientConsent struct {
    Linkable
    // The allowedScopes property
    allowedScopes []string
    // The client property
    client ClientClientApplicationPrimerable
    // The identifier property
    identifier *string
}
// NewClientClientConsent instantiates a new ClientClientConsent and sets the default values.
func NewClientClientConsent()(*ClientClientConsent) {
    m := &ClientClientConsent{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "client.ClientConsent"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientClientConsentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientClientConsentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientClientConsent(), nil
}
// GetAllowedScopes gets the allowedScopes property value. The allowedScopes property
// returns a []string when successful
func (m *ClientClientConsent) GetAllowedScopes()([]string) {
    return m.allowedScopes
}
// GetClient gets the client property value. The client property
// returns a ClientClientApplicationPrimerable when successful
func (m *ClientClientConsent) GetClient()(ClientClientApplicationPrimerable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientClientConsent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["allowedScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedScopes(res)
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
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *string when successful
func (m *ClientClientConsent) GetIdentifier()(*string) {
    return m.identifier
}
// Serialize serializes information the current object
func (m *ClientClientConsent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedScopes() != nil {
        err = writer.WriteCollectionOfStringValues("allowedScopes", m.GetAllowedScopes())
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
// SetAllowedScopes sets the allowedScopes property value. The allowedScopes property
func (m *ClientClientConsent) SetAllowedScopes(value []string)() {
    m.allowedScopes = value
}
// SetClient sets the client property value. The client property
func (m *ClientClientConsent) SetClient(value ClientClientApplicationPrimerable)() {
    m.client = value
}
// SetIdentifier sets the identifier property value. The identifier property
func (m *ClientClientConsent) SetIdentifier(value *string)() {
    m.identifier = value
}
type ClientClientConsentable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedScopes()([]string)
    GetClient()(ClientClientApplicationPrimerable)
    GetIdentifier()(*string)
    SetAllowedScopes(value []string)()
    SetClient(value ClientClientApplicationPrimerable)()
    SetIdentifier(value *string)()
}
