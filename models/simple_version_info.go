package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SimpleVersionInfo struct {
    NonLinkable
    // The contractVersions property
    contractVersions []int32
    // The keyHubVersion property
    keyHubVersion *string
}
// NewSimpleVersionInfo instantiates a new SimpleVersionInfo and sets the default values.
func NewSimpleVersionInfo()(*SimpleVersionInfo) {
    m := &SimpleVersionInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "simple.VersionInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateSimpleVersionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimpleVersionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleVersionInfo(), nil
}
// GetContractVersions gets the contractVersions property value. The contractVersions property
// returns a []int32 when successful
func (m *SimpleVersionInfo) GetContractVersions()([]int32) {
    return m.contractVersions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimpleVersionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["contractVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetContractVersions(res)
        }
        return nil
    }
    res["keyHubVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyHubVersion(val)
        }
        return nil
    }
    return res
}
// GetKeyHubVersion gets the keyHubVersion property value. The keyHubVersion property
// returns a *string when successful
func (m *SimpleVersionInfo) GetKeyHubVersion()(*string) {
    return m.keyHubVersion
}
// Serialize serializes information the current object
func (m *SimpleVersionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContractVersions() != nil {
        err = writer.WriteCollectionOfInt32Values("contractVersions", m.GetContractVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyHubVersion", m.GetKeyHubVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContractVersions sets the contractVersions property value. The contractVersions property
func (m *SimpleVersionInfo) SetContractVersions(value []int32)() {
    m.contractVersions = value
}
// SetKeyHubVersion sets the keyHubVersion property value. The keyHubVersion property
func (m *SimpleVersionInfo) SetKeyHubVersion(value *string)() {
    m.keyHubVersion = value
}
type SimpleVersionInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContractVersions()([]int32)
    GetKeyHubVersion()(*string)
    SetContractVersions(value []int32)()
    SetKeyHubVersion(value *string)()
}
