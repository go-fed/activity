package propertytype

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsTypePropertyIterator is an iterator for a property. It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsTypePropertyIterator struct {
	xmlschemaAnyURIMember *url.URL
	xmlschemaStringMember string
	hasStringMember       bool
	unknown               interface{}
	alias                 string
	myIdx                 int
	parent                vocab.ActivityStreamsTypeProperty
}

// NewActivityStreamsTypePropertyIterator creates a new ActivityStreamsType
// property.
func NewActivityStreamsTypePropertyIterator() *ActivityStreamsTypePropertyIterator {
	return &ActivityStreamsTypePropertyIterator{alias: ""}
}

// deserializeActivityStreamsTypePropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsTypePropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsTypePropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	if v, err := anyuri.DeserializeAnyURI(i); err == nil {
		this := &ActivityStreamsTypePropertyIterator{
			alias:                 alias,
			xmlschemaAnyURIMember: v,
		}
		return this, nil
	} else if v, err := string1.DeserializeString(i); err == nil {
		this := &ActivityStreamsTypePropertyIterator{
			alias:                 alias,
			hasStringMember:       true,
			xmlschemaStringMember: v,
		}
		return this, nil
	}
	this := &ActivityStreamsTypePropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsTypePropertyIterator) GetIRI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaAnyURI returns the value of this property. When IsXMLSchemaAnyURI
// returns false, GetXMLSchemaAnyURI will return an arbitrary value.
func (this ActivityStreamsTypePropertyIterator) GetXMLSchemaAnyURI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this ActivityStreamsTypePropertyIterator) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsTypePropertyIterator) HasAny() bool {
	return this.IsXMLSchemaAnyURI() ||
		this.IsXMLSchemaString()
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsTypePropertyIterator) IsIRI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI". When
// true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI methods to access
// and set this property.
func (this ActivityStreamsTypePropertyIterator) IsXMLSchemaAnyURI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.
func (this ActivityStreamsTypePropertyIterator) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsTypePropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsTypePropertyIterator) KindIndex() int {
	if this.IsXMLSchemaAnyURI() {
		return 0
	}
	if this.IsXMLSchemaString() {
		return 1
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsTypePropertyIterator) LessThan(o vocab.ActivityStreamsTypePropertyIterator) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaAnyURI() {
		return anyuri.LessAnyURI(this.GetXMLSchemaAnyURI(), o.GetXMLSchemaAnyURI())
	} else if this.IsXMLSchemaString() {
		return string1.LessString(this.GetXMLSchemaString(), o.GetXMLSchemaString())
	}
	return false
}

// Name returns the name of this property: "ActivityStreamsType".
func (this ActivityStreamsTypePropertyIterator) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "ActivityStreamsType"
	} else {
		return "ActivityStreamsType"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsTypePropertyIterator) Next() vocab.ActivityStreamsTypePropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsTypePropertyIterator) Prev() vocab.ActivityStreamsTypePropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsTypePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.SetXMLSchemaAnyURI(v)
}

// SetXMLSchemaAnyURI sets the value of this property. Calling IsXMLSchemaAnyURI
// afterwards returns true.
func (this *ActivityStreamsTypePropertyIterator) SetXMLSchemaAnyURI(v *url.URL) {
	this.clear()
	this.xmlschemaAnyURIMember = v
}

// SetXMLSchemaString sets the value of this property. Calling IsXMLSchemaString
// afterwards returns true.
func (this *ActivityStreamsTypePropertyIterator) SetXMLSchemaString(v string) {
	this.clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsTypePropertyIterator) clear() {
	this.xmlschemaAnyURIMember = nil
	this.hasStringMember = false
	this.unknown = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsTypePropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaAnyURI() {
		return anyuri.SerializeAnyURI(this.GetXMLSchemaAnyURI())
	} else if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	}
	return this.unknown, nil
}

// ActivityStreamsTypeProperty is the non-functional property "type". It is
// permitted to have one or more values, and of different value types.
type ActivityStreamsTypeProperty struct {
	properties []*ActivityStreamsTypePropertyIterator
	alias      string
}

// DeserializeTypeProperty creates a "type" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeTypeProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTypeProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "type"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "type")
	}
	i, ok := m[propName]

	if ok {
		this := &ActivityStreamsTypeProperty{
			alias:      alias,
			properties: []*ActivityStreamsTypePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsTypePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsTypePropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsTypeProperty creates a new type property.
func NewActivityStreamsTypeProperty() *ActivityStreamsTypeProperty {
	return &ActivityStreamsTypeProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "type"
func (this *ActivityStreamsTypeProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// AppendXMLSchemaAnyURI appends a anyURI value to the back of a list of the
// property "type". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsTypeProperty) AppendXMLSchemaAnyURI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// AppendXMLSchemaString appends a string value to the back of a list of the
// property "type". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsTypeProperty) AppendXMLSchemaString(v string) {
	this.properties = append(this.properties, &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaStringMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsTypeProperty) At(index int) vocab.ActivityStreamsTypePropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsTypeProperty) Begin() vocab.ActivityStreamsTypePropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsTypeProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsTypeProperty) End() vocab.ActivityStreamsTypePropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "type".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaAnyURI inserts a anyURI value at the specified index for a
// property "type". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) InsertXMLSchemaAnyURI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaString inserts a string value at the specified index for a
// property "type". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) InsertXMLSchemaString(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsTypeProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this ActivityStreamsTypeProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "type" property.
func (this ActivityStreamsTypeProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsTypeProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetXMLSchemaAnyURI()
			rhs := this.properties[j].GetXMLSchemaAnyURI()
			return anyuri.LessAnyURI(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetXMLSchemaString()
			rhs := this.properties[j].GetXMLSchemaString()
			return string1.LessString(lhs, rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsTypeProperty) LessThan(o vocab.ActivityStreamsTypeProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property ("type") with any alias.
func (this ActivityStreamsTypeProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "type"
	} else {
		return "type"
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property "type".
func (this *ActivityStreamsTypeProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsTypePropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaAnyURI prepends a anyURI value to the front of a list of the
// property "type". Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) PrependXMLSchemaAnyURI(v *url.URL) {
	this.properties = append([]*ActivityStreamsTypePropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaString prepends a string value to the front of a list of the
// property "type". Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) PrependXMLSchemaString(v string) {
	this.properties = append([]*ActivityStreamsTypePropertyIterator{{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 0,
		parent:                this,
		xmlschemaStringMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "type", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsTypeProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsTypePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsTypeProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetIRI sets an IRI value to be at the specified index for the property "type".
// Panics if the index is out of bounds.
func (this *ActivityStreamsTypeProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// SetXMLSchemaAnyURI sets a anyURI value to be at the specified index for the
// property "type". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsTypeProperty) SetXMLSchemaAnyURI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// SetXMLSchemaString sets a string value to be at the specified index for the
// property "type". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsTypeProperty) SetXMLSchemaString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
}

// Swap swaps the location of values at two indices for the "type" property.
func (this ActivityStreamsTypeProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
