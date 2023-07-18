package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedAzureTenant 
type ProvisioningProvisionedAzureTenant struct {
    ProvisioningProvisionedSystem
    // The clientId property
    clientId *string
    // The clientSecret property
    clientSecret *string
    // The idpDomain property
    idpDomain *string
    // The tenant property
    tenant *string
}
// NewProvisioningProvisionedAzureTenant instantiates a new provisioningProvisionedAzureTenant and sets the default values.
func NewProvisioningProvisionedAzureTenant()(*ProvisioningProvisionedAzureTenant) {
    m := &ProvisioningProvisionedAzureTenant{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedAzureTenant"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedAzureTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedAzureTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedAzureTenant(), nil
}
// GetClientId gets the clientId property value. The clientId property
func (m *ProvisioningProvisionedAzureTenant) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. The clientSecret property
func (m *ProvisioningProvisionedAzureTenant) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedAzureTenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
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
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["idpDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdpDomain(val)
        }
        return nil
    }
    res["tenant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenant(val)
        }
        return nil
    }
    return res
}
// GetIdpDomain gets the idpDomain property value. The idpDomain property
func (m *ProvisioningProvisionedAzureTenant) GetIdpDomain()(*string) {
    return m.idpDomain
}
// GetTenant gets the tenant property value. The tenant property
func (m *ProvisioningProvisionedAzureTenant) GetTenant()(*string) {
    return m.tenant
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedAzureTenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("idpDomain", m.GetIdpDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenant", m.GetTenant())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The clientId property
func (m *ProvisioningProvisionedAzureTenant) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. The clientSecret property
func (m *ProvisioningProvisionedAzureTenant) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetIdpDomain sets the idpDomain property value. The idpDomain property
func (m *ProvisioningProvisionedAzureTenant) SetIdpDomain(value *string)() {
    m.idpDomain = value
}
// SetTenant sets the tenant property value. The tenant property
func (m *ProvisioningProvisionedAzureTenant) SetTenant(value *string)() {
    m.tenant = value
}
// ProvisioningProvisionedAzureTenantable 
type ProvisioningProvisionedAzureTenantable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetIdpDomain()(*string)
    GetTenant()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetIdpDomain(value *string)()
    SetTenant(value *string)()
}
