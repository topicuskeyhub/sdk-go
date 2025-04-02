package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentitysourceAFASIdentitySource struct {
    IdentitysourceIdentitySource
    // The token property
    token *string
    // The url property
    url *string
}
// NewIdentitysourceAFASIdentitySource instantiates a new IdentitysourceAFASIdentitySource and sets the default values.
func NewIdentitysourceAFASIdentitySource()(*IdentitysourceAFASIdentitySource) {
    m := &IdentitysourceAFASIdentitySource{
        IdentitysourceIdentitySource: *NewIdentitysourceIdentitySource(),
    }
    typeEscapedValue := "identitysource.AFASIdentitySource"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentitysourceAFASIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentitysourceAFASIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentitysourceAFASIdentitySource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentitysourceAFASIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitysourceIdentitySource.GetFieldDeserializers()
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *IdentitysourceAFASIdentitySource) GetToken()(*string) {
    return m.token
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *IdentitysourceAFASIdentitySource) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *IdentitysourceAFASIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitysourceIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("token", m.GetToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetToken sets the token property value. The token property
func (m *IdentitysourceAFASIdentitySource) SetToken(value *string)() {
    m.token = value
}
// SetUrl sets the url property value. The url property
func (m *IdentitysourceAFASIdentitySource) SetUrl(value *string)() {
    m.url = value
}
type IdentitysourceAFASIdentitySourceable interface {
    IdentitysourceIdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetToken()(*string)
    GetUrl()(*string)
    SetToken(value *string)()
    SetUrl(value *string)()
}
