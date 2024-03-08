package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultRecord struct {
    VaultVaultRecordPrimer
    // The additionalObjects property
    additionalObjects VaultVaultRecord_additionalObjectsable
    // The derived property
    derived *bool
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The filename property
    filename *string
    // The types property
    types []VaultVaultSecretType
    // The url property
    url *string
    // The username property
    username *string
    // The warningPeriod property
    warningPeriod *VaultVaultRecordWarningPeriod
}
// NewVaultVaultRecord instantiates a new VaultVaultRecord and sets the default values.
func NewVaultVaultRecord()(*VaultVaultRecord) {
    m := &VaultVaultRecord{
        VaultVaultRecordPrimer: *NewVaultVaultRecordPrimer(),
    }
    typeEscapedValue := "vault.VaultRecord"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecord(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a VaultVaultRecord_additionalObjectsable when successful
func (m *VaultVaultRecord) GetAdditionalObjects()(VaultVaultRecord_additionalObjectsable) {
    return m.additionalObjects
}
// GetDerived gets the derived property value. The derived property
// returns a *bool when successful
func (m *VaultVaultRecord) GetDerived()(*bool) {
    return m.derived
}
// GetEndDate gets the endDate property value. The endDate property
// returns a *DateOnly when successful
func (m *VaultVaultRecord) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VaultVaultRecordPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecord_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(VaultVaultRecord_additionalObjectsable))
        }
        return nil
    }
    res["derived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerived(val)
        }
        return nil
    }
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilename(val)
        }
        return nil
    }
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseVaultVaultSecretType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VaultVaultSecretType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*VaultVaultSecretType))
                }
            }
            m.SetTypes(res)
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
    res["warningPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultVaultRecordWarningPeriod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarningPeriod(val.(*VaultVaultRecordWarningPeriod))
        }
        return nil
    }
    return res
}
// GetFilename gets the filename property value. The filename property
// returns a *string when successful
func (m *VaultVaultRecord) GetFilename()(*string) {
    return m.filename
}
// GetTypes gets the types property value. The types property
// returns a []VaultVaultSecretType when successful
func (m *VaultVaultRecord) GetTypes()([]VaultVaultSecretType) {
    return m.types
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *VaultVaultRecord) GetUrl()(*string) {
    return m.url
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *VaultVaultRecord) GetUsername()(*string) {
    return m.username
}
// GetWarningPeriod gets the warningPeriod property value. The warningPeriod property
// returns a *VaultVaultRecordWarningPeriod when successful
func (m *VaultVaultRecord) GetWarningPeriod()(*VaultVaultRecordWarningPeriod) {
    return m.warningPeriod
}
// Serialize serializes information the current object
func (m *VaultVaultRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VaultVaultRecordPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filename", m.GetFilename())
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
    {
        err = writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    if m.GetWarningPeriod() != nil {
        cast := (*m.GetWarningPeriod()).String()
        err = writer.WriteStringValue("warningPeriod", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *VaultVaultRecord) SetAdditionalObjects(value VaultVaultRecord_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDerived sets the derived property value. The derived property
func (m *VaultVaultRecord) SetDerived(value *bool)() {
    m.derived = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *VaultVaultRecord) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetFilename sets the filename property value. The filename property
func (m *VaultVaultRecord) SetFilename(value *string)() {
    m.filename = value
}
// SetTypes sets the types property value. The types property
func (m *VaultVaultRecord) SetTypes(value []VaultVaultSecretType)() {
    m.types = value
}
// SetUrl sets the url property value. The url property
func (m *VaultVaultRecord) SetUrl(value *string)() {
    m.url = value
}
// SetUsername sets the username property value. The username property
func (m *VaultVaultRecord) SetUsername(value *string)() {
    m.username = value
}
// SetWarningPeriod sets the warningPeriod property value. The warningPeriod property
func (m *VaultVaultRecord) SetWarningPeriod(value *VaultVaultRecordWarningPeriod)() {
    m.warningPeriod = value
}
type VaultVaultRecordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VaultVaultRecordPrimerable
    GetAdditionalObjects()(VaultVaultRecord_additionalObjectsable)
    GetDerived()(*bool)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetFilename()(*string)
    GetTypes()([]VaultVaultSecretType)
    GetUrl()(*string)
    GetUsername()(*string)
    GetWarningPeriod()(*VaultVaultRecordWarningPeriod)
    SetAdditionalObjects(value VaultVaultRecord_additionalObjectsable)()
    SetDerived(value *bool)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetFilename(value *string)()
    SetTypes(value []VaultVaultSecretType)()
    SetUrl(value *string)()
    SetUsername(value *string)()
    SetWarningPeriod(value *VaultVaultRecordWarningPeriod)()
}
