package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountPrimer 
type AuthAccountPrimer struct {
    Linkable
    // The displayName property
    displayName *string
    // The lastActive property
    lastActive *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The username property
    username *string
    // The uuid property
    uuid *string
    // The validity property
    validity *AuthAccountValidity
}
// NewAuthAccountPrimer instantiates a new authAccountPrimer and sets the default values.
func NewAuthAccountPrimer()(*AuthAccountPrimer) {
    m := &AuthAccountPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "auth.AccountPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "auth.Account":
                        return NewAuthAccount(), nil
                    case "auth.InternalAccount":
                        return NewAuthInternalAccount(), nil
                    case "group.GroupAccount":
                        return NewGroupGroupAccount(), nil
                    case "organization.OrganizationalUnitAccount":
                        return NewOrganizationOrganizationalUnitAccount(), nil
                    case "provisioning.ProvisionedAccount":
                        return NewProvisioningProvisionedAccount(), nil
                }
            }
        }
    }
    return NewAuthAccountPrimer(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AuthAccountPrimer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActive(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
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
    res["validity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthAccountValidity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidity(val.(*AuthAccountValidity))
        }
        return nil
    }
    return res
}
// GetLastActive gets the lastActive property value. The lastActive property
func (m *AuthAccountPrimer) GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActive
}
// GetUsername gets the username property value. The username property
func (m *AuthAccountPrimer) GetUsername()(*string) {
    return m.username
}
// GetUuid gets the uuid property value. The uuid property
func (m *AuthAccountPrimer) GetUuid()(*string) {
    return m.uuid
}
// GetValidity gets the validity property value. The validity property
func (m *AuthAccountPrimer) GetValidity()(*AuthAccountValidity) {
    return m.validity
}
// Serialize serializes information the current object
func (m *AuthAccountPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValidity() != nil {
        cast := (*m.GetValidity()).String()
        err = writer.WriteStringValue("validity", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuthAccountPrimer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastActive sets the lastActive property value. The lastActive property
func (m *AuthAccountPrimer) SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActive = value
}
// SetUsername sets the username property value. The username property
func (m *AuthAccountPrimer) SetUsername(value *string)() {
    m.username = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *AuthAccountPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// SetValidity sets the validity property value. The validity property
func (m *AuthAccountPrimer) SetValidity(value *AuthAccountValidity)() {
    m.validity = value
}
// AuthAccountPrimerable 
type AuthAccountPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUsername()(*string)
    GetUuid()(*string)
    GetValidity()(*AuthAccountValidity)
    SetDisplayName(value *string)()
    SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUsername(value *string)()
    SetUuid(value *string)()
    SetValidity(value *AuthAccountValidity)()
}
