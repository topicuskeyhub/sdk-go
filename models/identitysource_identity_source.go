package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentitysourceIdentitySource struct {
    IdentitysourceIdentitySourcePrimer
    // The additionalObjects property
    additionalObjects IdentitysourceIdentitySource_additionalObjectsable
}
// NewIdentitysourceIdentitySource instantiates a new IdentitysourceIdentitySource and sets the default values.
func NewIdentitysourceIdentitySource()(*IdentitysourceIdentitySource) {
    m := &IdentitysourceIdentitySource{
        IdentitysourceIdentitySourcePrimer: *NewIdentitysourceIdentitySourcePrimer(),
    }
    typeEscapedValue := "identitysource.IdentitySource"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentitysourceIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentitysourceIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "identitysource.AFASIdentitySource":
                        return NewIdentitysourceAFASIdentitySource(), nil
                }
            }
        }
    }
    return NewIdentitysourceIdentitySource(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a IdentitysourceIdentitySource_additionalObjectsable when successful
func (m *IdentitysourceIdentitySource) GetAdditionalObjects()(IdentitysourceIdentitySource_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentitysourceIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitysourceIdentitySourcePrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitysourceIdentitySource_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(IdentitysourceIdentitySource_additionalObjectsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IdentitysourceIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitysourceIdentitySourcePrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *IdentitysourceIdentitySource) SetAdditionalObjects(value IdentitysourceIdentitySource_additionalObjectsable)() {
    m.additionalObjects = value
}
type IdentitysourceIdentitySourceable interface {
    IdentitysourceIdentitySourcePrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(IdentitysourceIdentitySource_additionalObjectsable)
    SetAdditionalObjects(value IdentitysourceIdentitySource_additionalObjectsable)()
}
