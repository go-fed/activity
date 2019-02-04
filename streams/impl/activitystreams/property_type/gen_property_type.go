package propertytype

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// TypePropertyIterator is an iterator for a property. It is permitted to be a
// single nilable value type.
type TypePropertyIterator struct {
	anyURIMember *url.URL
	unknown      []byte
	alias        string
	myIdx        int
	parent       vocab.TypePropertyInterface
}

// NewTypePropertyIterator creates a new type property.
func NewTypePropertyIterator() *TypePropertyIterator {
	return &TypePropertyIterator{alias: ""}
}

// deserializeTypePropertyIterator creates an iterator from an element that has
// been unmarshalled from a text or binary format.
func deserializeTypePropertyIterator(i interface{}, aliasMap map[string]string) (*TypePropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if v, err := anyuri.DeserializeAnyURI(i); err == nil {
		this := &TypePropertyIterator{
			alias:        alias,
			anyURIMember: v,
		}
		return this, nil
	} else if str, ok := i.(string); ok {
		this := &TypePropertyIterator{
			alias:   alias,
			unknown: []byte(str),
		}
		return this, nil
	}
	return nil, fmt.Errorf("could not deserialize %q property", "type")
}

// Get returns the value of this property. When IsAnyURI returns false, Get will
// return any arbitrary value.
func (this TypePropertyIterator) Get() *url.URL {
	return this.anyURIMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this TypePropertyIterator) GetIRI() *url.URL {
	return this.anyURIMember
}

// HasAny returns true if the value or IRI is set.
func (this TypePropertyIterator) HasAny() bool {
	return this.IsAnyURI()
}

// IsAnyURI returns true if this property is set and not an IRI.
func (this TypePropertyIterator) IsAnyURI() bool {
	return this.anyURIMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this TypePropertyIterator) IsIRI() bool {
	return this.anyURIMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this TypePropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
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
func (this TypePropertyIterator) KindIndex() int {
	if this.IsAnyURI() {
		return 0
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
func (this TypePropertyIterator) LessThan(o vocab.TypePropertyIteratorInterface) bool {
	if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsAnyURI() && !o.IsAnyURI() {
		// Both are unknowns.
		return false
	} else if this.IsAnyURI() && !o.IsAnyURI() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsAnyURI() && o.IsAnyURI() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return anyuri.LessAnyURI(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "type".
func (this TypePropertyIterator) Name() string {
	return "type"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this TypePropertyIterator) Next() vocab.TypePropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this TypePropertyIterator) Prev() vocab.TypePropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsAnyURI afterwards will return
// true.
func (this *TypePropertyIterator) Set(v *url.URL) {
	this.clear()
	this.anyURIMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *TypePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.Set(v)
}

// clear ensures no value of this property is set. Calling IsAnyURI afterwards
// will return false.
func (this *TypePropertyIterator) clear() {
	this.unknown = nil
	this.anyURIMember = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this TypePropertyIterator) serialize() (interface{}, error) {
	if this.IsAnyURI() {
		return anyuri.SerializeAnyURI(this.Get())
	}
	return string(this.unknown), nil
}

// TypeProperty is the non-functional property "type". It is permitted to have one
// or more values, and of different value types.
type TypeProperty struct {
	properties []*TypePropertyIterator
	alias      string
}

// DeserializeTypeProperty creates a "type" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeTypeProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.TypePropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "type"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "type")
	}
	if i, ok := m[propName]; ok {
		this := &TypeProperty{
			alias:      alias,
			properties: []*TypePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeTypePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeTypePropertyIterator(i, aliasMap); err != nil {
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

// NewTypeProperty creates a new type property.
func NewTypeProperty() *TypeProperty {
	return &TypeProperty{alias: ""}
}

// AppendAnyURI appends a anyURI value to the back of a list of the property
// "type". Invalidates iterators that are traversing using Prev.
func (this *TypeProperty) AppendAnyURI(v *url.URL) {
	this.properties = append(this.properties, &TypePropertyIterator{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendIRI appends an IRI value to the back of a list of the property "type"
func (this *TypeProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &TypePropertyIterator{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this TypeProperty) At(index int) vocab.TypePropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this TypeProperty) Begin() vocab.TypePropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this TypeProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this TypeProperty) End() vocab.TypePropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this TypeProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
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
func (this TypeProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "type" property.
func (this TypeProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this TypeProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return anyuri.LessAnyURI(lhs, rhs)
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
func (this TypeProperty) LessThan(o vocab.TypePropertyInterface) bool {
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

// Name returns the name of this property: "type".
func (this TypeProperty) Name() string {
	return "type"
}

// PrependAnyURI prepends a anyURI value to the front of a list of the property
// "type". Invalidates all iterators.
func (this *TypeProperty) PrependAnyURI(v *url.URL) {
	this.properties = append([]*TypePropertyIterator{{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property "type".
func (this *TypeProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*TypePropertyIterator{{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "type", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *TypeProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &TypePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this TypeProperty) Serialize() (interface{}, error) {
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

// Set sets a anyURI value to be at the specified index for the property "type".
// Panics if the index is out of bounds. Invalidates all iterators.
func (this *TypeProperty) Set(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &TypePropertyIterator{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        idx,
		parent:       this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property "type".
// Panics if the index is out of bounds.
func (this *TypeProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &TypePropertyIterator{
		alias:        this.alias,
		anyURIMember: v,
		myIdx:        idx,
		parent:       this,
	}
}

// Swap swaps the location of values at two indices for the "type" property.
func (this TypeProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
