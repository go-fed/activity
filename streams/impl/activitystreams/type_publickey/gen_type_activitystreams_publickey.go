package typepublickey

import (
	vocab "github.com/go-fed/activity/streams/vocab"
)

// A public key represents a public cryptographical key for a user
type ActivityStreamsPublicKey struct {
	ActivityStreamsId           vocab.ActivityStreamsIdProperty
	ActivityStreamsOwner        vocab.ActivityStreamsOwnerProperty
	ActivityStreamsPublicKeyPem vocab.ActivityStreamsPublicKeyPemProperty
	ActivityStreamsType         vocab.ActivityStreamsTypeProperty
	alias                       string
	unknown                     map[string]interface{}
}

// ActivityStreamsPublicKeyExtends returns true if the PublicKey type extends from
// the other type.
func ActivityStreamsPublicKeyExtends(other vocab.Type) bool {
	// Shortcut implementation: this does not extend anything.
	return false
}

// DeserializePublicKey creates a PublicKey from a map representation that has
// been unmarshalled from a text or binary format.
func DeserializePublicKey(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsPublicKey, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	this := &ActivityStreamsPublicKey{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}

	// HACK: IGNORE TYPE
	// TODO: ENCODE THIS HACK IN THE CODE GENERATION

	// Begin: Known property deserialization
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsId = p
	}
	if p, err := mgr.DeserializeOwnerPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsOwner = p
	}
	if p, err := mgr.DeserializePublicKeyPemPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPublicKeyPem = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsType = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "id" {
			continue
		} else if k == "owner" {
			continue
		} else if k == "publicKeyPem" {
			continue
		} else if k == "type" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// IsOrExtendsPublicKey returns true if the other provided type is the PublicKey
// type or extends from the PublicKey type.
func IsOrExtendsPublicKey(other vocab.Type) bool {
	if other.GetTypeName() == "PublicKey" {
		return true
	}
	return PublicKeyIsExtendedBy(other)
}

// NewActivityStreamsPublicKey creates a new PublicKey type
func NewActivityStreamsPublicKey() *ActivityStreamsPublicKey {
	typeProp := typePropertyConstructor()
	typeProp.AppendXMLSchemaString("PublicKey")
	return &ActivityStreamsPublicKey{
		ActivityStreamsType: typeProp,
		alias:               "",
		unknown:             make(map[string]interface{}, 0),
	}
}

// PublicKeyIsDisjointWith returns true if the other provided type is disjoint
// with the PublicKey type.
func PublicKeyIsDisjointWith(other vocab.Type) bool {
	// Shortcut implementation: is not disjoint with anything.
	return false
}

// PublicKeyIsExtendedBy returns true if the other provided type extends from the
// PublicKey type. Note that it returns false if the types are the same; see
// the "IsOrExtendsPublicKey" variant instead.
func PublicKeyIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// GetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this ActivityStreamsPublicKey) GetActivityStreamsId() vocab.ActivityStreamsIdProperty {
	return this.ActivityStreamsId
}

// GetActivityStreamsOwner returns the "owner" property if it exists, and nil
// otherwise.
func (this ActivityStreamsPublicKey) GetActivityStreamsOwner() vocab.ActivityStreamsOwnerProperty {
	return this.ActivityStreamsOwner
}

// GetActivityStreamsPublicKeyPem returns the "publicKeyPem" property if it
// exists, and nil otherwise.
func (this ActivityStreamsPublicKey) GetActivityStreamsPublicKeyPem() vocab.ActivityStreamsPublicKeyPemProperty {
	return this.ActivityStreamsPublicKeyPem
}

// GetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this ActivityStreamsPublicKey) GetActivityStreamsType() vocab.ActivityStreamsTypeProperty {
	return this.ActivityStreamsType
}

// GetTypeName returns the name of this type.
func (this ActivityStreamsPublicKey) GetTypeName() string {
	return "PublicKey"
}

// GetUnknownProperties returns the unknown properties for the PublicKey type.
// Note that this should not be used by app developers. It is only used to
// help determine which implementation is LessThan the other. Developers who
// are creating a different implementation of this type's interface can use
// this method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this ActivityStreamsPublicKey) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the PublicKey type extends from the other type.
func (this ActivityStreamsPublicKey) IsExtending(other vocab.Type) bool {
	return ActivityStreamsPublicKeyExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this ActivityStreamsPublicKey) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	m = this.helperJSONLDContext(this.ActivityStreamsId, m)
	m = this.helperJSONLDContext(this.ActivityStreamsOwner, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPublicKeyPem, m)
	m = this.helperJSONLDContext(this.ActivityStreamsType, m)

	return m
}

// LessThan computes if this PublicKey is lesser, with an arbitrary but stable
// determination.
func (this ActivityStreamsPublicKey) LessThan(o vocab.ActivityStreamsPublicKey) bool {
	// Begin: Compare known properties
	// Compare property "id"
	if lhs, rhs := this.ActivityStreamsId, o.GetActivityStreamsId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "owner"
	if lhs, rhs := this.ActivityStreamsOwner, o.GetActivityStreamsOwner(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "publicKeyPem"
	if lhs, rhs := this.ActivityStreamsPublicKeyPem, o.GetActivityStreamsPublicKeyPem(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "type"
	if lhs, rhs := this.ActivityStreamsType, o.GetActivityStreamsType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this ActivityStreamsPublicKey) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "PublicKey"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "PublicKey"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "id"
	if this.ActivityStreamsId != nil {
		if i, err := this.ActivityStreamsId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsId.Name()] = i
		}
	}
	// Maybe serialize property "owner"
	if this.ActivityStreamsOwner != nil {
		if i, err := this.ActivityStreamsOwner.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsOwner.Name()] = i
		}
	}
	// Maybe serialize property "publicKeyPem"
	if this.ActivityStreamsPublicKeyPem != nil {
		if i, err := this.ActivityStreamsPublicKeyPem.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPublicKeyPem.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.ActivityStreamsType != nil {
		if i, err := this.ActivityStreamsType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsType.Name()] = i
		}
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetActivityStreamsId sets the "id" property.
func (this *ActivityStreamsPublicKey) SetActivityStreamsId(i vocab.ActivityStreamsIdProperty) {
	this.ActivityStreamsId = i
}

// SetActivityStreamsOwner sets the "owner" property.
func (this *ActivityStreamsPublicKey) SetActivityStreamsOwner(i vocab.ActivityStreamsOwnerProperty) {
	this.ActivityStreamsOwner = i
}

// SetActivityStreamsPublicKeyPem sets the "publicKeyPem" property.
func (this *ActivityStreamsPublicKey) SetActivityStreamsPublicKeyPem(i vocab.ActivityStreamsPublicKeyPemProperty) {
	this.ActivityStreamsPublicKeyPem = i
}

// SetActivityStreamsType sets the "type" property.
func (this *ActivityStreamsPublicKey) SetActivityStreamsType(i vocab.ActivityStreamsTypeProperty) {
	this.ActivityStreamsType = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this ActivityStreamsPublicKey) VocabularyURI() string {
	return "https://www.w3.org/ns/activitystreams"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this ActivityStreamsPublicKey) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}
