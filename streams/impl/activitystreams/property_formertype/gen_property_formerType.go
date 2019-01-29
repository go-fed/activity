package propertyformertype

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// FormerTypePropertyIterator is an iterator for a property. It is permitted to be
// a single nilable value type.
type FormerTypePropertyIterator struct {
	ObjectMember vocab.ObjectInterface
	unknown      []byte
	iri          *url.URL
	alias        string
	myIdx        int
	parent       vocab.FormerTypePropertyInterface
}

// NewFormerTypePropertyIterator creates a new formerType property.
func NewFormerTypePropertyIterator() *FormerTypePropertyIterator {
	return &FormerTypePropertyIterator{alias: ""}
}

// deserializeFormerTypePropertyIterator creates an iterator from an element that
// has been unmarshalled from a text or binary format.
func deserializeFormerTypePropertyIterator(i interface{}, aliasMap map[string]string) (*FormerTypePropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		if err == nil {
			this := &FormerTypePropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err != nil {
			this := &FormerTypePropertyIterator{
				ObjectMember: v,
				alias:        alias,
			}
			return this, nil
		}
	} else if v, ok := i.([]byte); ok {
		this := &FormerTypePropertyIterator{
			alias:   alias,
			unknown: v,
		}
		return this, nil
	}
	return nil, fmt.Errorf("could not deserialize %q property", "formerType")
}

// Get returns the value of this property. When IsObject returns false, Get will
// return any arbitrary value.
func (this FormerTypePropertyIterator) Get() vocab.ObjectInterface {
	return this.ObjectMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this FormerTypePropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this FormerTypePropertyIterator) HasAny() bool {
	return this.IsObject() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this FormerTypePropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsObject returns true if this property is set and not an IRI.
func (this FormerTypePropertyIterator) IsObject() bool {
	return this.ObjectMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this FormerTypePropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsObject() {
		child = this.Get().JSONLDContext()
	}
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
func (this FormerTypePropertyIterator) KindIndex() int {
	if this.IsObject() {
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
func (this FormerTypePropertyIterator) LessThan(o vocab.FormerTypePropertyIteratorInterface) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsObject() && !o.IsObject() {
		// Both are unknowns.
		return false
	} else if this.IsObject() && !o.IsObject() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsObject() && o.IsObject() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "formerType".
func (this FormerTypePropertyIterator) Name() string {
	return "formerType"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this FormerTypePropertyIterator) Next() vocab.FormerTypePropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this FormerTypePropertyIterator) Prev() vocab.FormerTypePropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsObject afterwards will return
// true.
func (this *FormerTypePropertyIterator) Set(v vocab.ObjectInterface) {
	this.clear()
	this.ObjectMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *FormerTypePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// clear ensures no value of this property is set. Calling IsObject afterwards
// will return false.
func (this *FormerTypePropertyIterator) clear() {
	this.unknown = nil
	this.iri = nil
	this.ObjectMember = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this FormerTypePropertyIterator) serialize() (interface{}, error) {
	if this.IsObject() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// FormerTypeProperty is the non-functional property "formerType". It is permitted
// to have one or more values, and of different value types.
type FormerTypeProperty struct {
	properties []*FormerTypePropertyIterator
	alias      string
}

// DeserializeFormerTypeProperty creates a "formerType" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeFormerTypeProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.FormerTypePropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	var this *FormerTypeProperty
	propName := "formerType"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "formerType")
	}
	if i, ok := m[propName]; ok {
		this := &FormerTypeProperty{
			alias:      alias,
			properties: []*FormerTypePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeFormerTypePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeFormerTypePropertyIterator(i, aliasMap); err != nil {
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
	}
	return this, nil
}

// NewFormerTypeProperty creates a new formerType property.
func NewFormerTypeProperty() *FormerTypeProperty {
	return &FormerTypeProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property
// "formerType"
func (this *FormerTypeProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &FormerTypePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendObject appends a Object value to the back of a list of the property
// "formerType". Invalidates iterators that are traversing using Prev.
func (this *FormerTypeProperty) AppendObject(v vocab.ObjectInterface) {
	this.properties = append(this.properties, &FormerTypePropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this FormerTypeProperty) At(index int) vocab.FormerTypePropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this FormerTypeProperty) Begin() vocab.FormerTypePropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this FormerTypeProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this FormerTypeProperty) End() vocab.FormerTypePropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this FormerTypeProperty) JSONLDContext() map[string]string {
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
func (this FormerTypeProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "formerType" property.
func (this FormerTypeProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this FormerTypeProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return lhs.LessThan(rhs)
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
func (this FormerTypeProperty) LessThan(o vocab.FormerTypePropertyInterface) bool {
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

// Name returns the name of this property: "formerType".
func (this FormerTypeProperty) Name() string {
	return "formerType"
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "formerType".
func (this *FormerTypeProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*FormerTypePropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependObject prepends a Object value to the front of a list of the property
// "formerType". Invalidates all iterators.
func (this *FormerTypeProperty) PrependObject(v vocab.ObjectInterface) {
	this.properties = append([]*FormerTypePropertyIterator{{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "formerType", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *FormerTypeProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &FormerTypePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this FormerTypeProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	return s, nil
}

// Set sets a Object value to be at the specified index for the property
// "formerType". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *FormerTypeProperty) Set(idx int, v vocab.ObjectInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &FormerTypePropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "formerType". Panics if the index is out of bounds.
func (this *FormerTypeProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &FormerTypePropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// Swap swaps the location of values at two indices for the "formerType" property.
func (this FormerTypeProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
