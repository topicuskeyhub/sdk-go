package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientClientApplicationPrimer struct {
    Linkable
    // The type property
    clientClientApplicationPrimerType *ClientClientApplicationType
    // The clientId property
    clientId *string
    // The name property
    name *string
    // The scopes property
    scopes []string
    // The ssoApplication property
    ssoApplication *bool
    // The uuid property
    uuid *string
}
// NewClientClientApplicationPrimer instantiates a new ClientClientApplicationPrimer and sets the default values.
func NewClientClientApplicationPrimer()(*ClientClientApplicationPrimer) {
    m := &ClientClientApplicationPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "client.ClientApplicationPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientClientApplicationPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientClientApplicationPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "client.ClientApplication":
                        return NewClientClientApplication(), nil
                    case "client.LdapClient":
                        return NewClientLdapClient(), nil
                    case "client.OAuth2Client":
                        return NewClientOAuth2Client(), nil
                    case "client.Saml2Client":
                        return NewClientSaml2Client(), nil
                    case "organization.OrganizationalUnitClientApplication":
                        return NewOrganizationOrganizationalUnitClientApplication(), nil
                }
            }
        }
    }
    return NewClientClientApplicationPrimer(), nil
}
// GetClientClientApplicationPrimerType gets the type property value. The type property
// returns a *ClientClientApplicationType when successful
func (m *ClientClientApplicationPrimer) GetClientClientApplicationPrimerType()(*ClientClientApplicationType) {
    return m.clientClientApplicationPrimerType
}
// GetClientId gets the clientId property value. The clientId property
// returns a *string when successful
func (m *ClientClientApplicationPrimer) GetClientId()(*string) {
    return m.clientId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientClientApplicationPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientClientApplicationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientClientApplicationPrimerType(val.(*ClientClientApplicationType))
        }
        return nil
    }
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetScopes(res)
        }
        return nil
    }
    res["ssoApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsoApplication(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ClientClientApplicationPrimer) GetName()(*string) {
    return m.name
}
// GetScopes gets the scopes property value. The scopes property
// returns a []string when successful
func (m *ClientClientApplicationPrimer) GetScopes()([]string) {
    return m.scopes
}
// GetSsoApplication gets the ssoApplication property value. The ssoApplication property
// returns a *bool when successful
func (m *ClientClientApplicationPrimer) GetSsoApplication()(*bool) {
    return m.ssoApplication
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *ClientClientApplicationPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *ClientClientApplicationPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientClientApplicationPrimerType() != nil {
        cast := (*m.GetClientClientApplicationPrimerType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        err = writer.WriteCollectionOfStringValues("scopes", m.GetScopes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientClientApplicationPrimerType sets the type property value. The type property
func (m *ClientClientApplicationPrimer) SetClientClientApplicationPrimerType(value *ClientClientApplicationType)() {
    m.clientClientApplicationPrimerType = value
}
// SetClientId sets the clientId property value. The clientId property
func (m *ClientClientApplicationPrimer) SetClientId(value *string)() {
    m.clientId = value
}
// SetName sets the name property value. The name property
func (m *ClientClientApplicationPrimer) SetName(value *string)() {
    m.name = value
}
// SetScopes sets the scopes property value. The scopes property
func (m *ClientClientApplicationPrimer) SetScopes(value []string)() {
    m.scopes = value
}
// SetSsoApplication sets the ssoApplication property value. The ssoApplication property
func (m *ClientClientApplicationPrimer) SetSsoApplication(value *bool)() {
    m.ssoApplication = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *ClientClientApplicationPrimer) SetUuid(value *string)() {
    m.uuid = value
}
type ClientClientApplicationPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientClientApplicationPrimerType()(*ClientClientApplicationType)
    GetClientId()(*string)
    GetName()(*string)
    GetScopes()([]string)
    GetSsoApplication()(*bool)
    GetUuid()(*string)
    SetClientClientApplicationPrimerType(value *ClientClientApplicationType)()
    SetClientId(value *string)()
    SetName(value *string)()
    SetScopes(value []string)()
    SetSsoApplication(value *bool)()
    SetUuid(value *string)()
}
