package propertyinstrument

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// InstrumentPropertyIterator is an iterator for a property. It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type InstrumentPropertyIterator struct {
	ObjectMember vocab.ObjectInterface
	LinkMember   vocab.LinkInterface
	unknown      []byte
	iri          *url.URL
	alias        string
	myIdx        int
	parent       vocab.InstrumentPropertyInterface
}

// NewInstrumentPropertyIterator creates a new instrument property.
func NewInstrumentPropertyIterator() *InstrumentPropertyIterator {
	return &InstrumentPropertyIterator{alias: ""}
}

// deserializeInstrumentPropertyIterator creates an iterator from an element that
// has been unmarshalled from a text or binary format.
func deserializeInstrumentPropertyIterator(i interface{}, aliasMap map[string]string) (*InstrumentPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		if err == nil {
			this := &InstrumentPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err != nil {
			this := &InstrumentPropertyIterator{
				ObjectMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err != nil {
			this := &InstrumentPropertyIterator{
				LinkMember: v,
				alias:      alias,
			}
			return this, nil
		}
	} else if v, ok := i.([]byte); ok {
		this := &InstrumentPropertyIterator{
			alias:   alias,
			unknown: v,
		}
		return this, nil
	}
	return nil, fmt.Errorf("could not deserialize %q property", "instrument")
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this InstrumentPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetLink returns the value of this property. When IsLink returns false, GetLink
// will return an arbitrary value.
func (this InstrumentPropertyIterator) GetLink() vocab.LinkInterface {
	return this.LinkMember
}

// GetObject returns the value of this property. When IsObject returns false,
// GetObject will return an arbitrary value.
func (this InstrumentPropertyIterator) GetObject() vocab.ObjectInterface {
	return this.ObjectMember
}

// HasAny returns true if any of the different values is set.
func (this InstrumentPropertyIterator) HasAny() bool {
	return this.IsObject() ||
		this.IsLink() ||
		this.iri != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this InstrumentPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsLink returns true if this property has a type of "Link". When true, use the
// GetLink and SetLink methods to access and set this property.
func (this InstrumentPropertyIterator) IsLink() bool {
	return this.LinkMember != nil
}

// IsObject returns true if this property has a type of "Object". When true, use
// the GetObject and SetObject methods to access and set this property.
func (this InstrumentPropertyIterator) IsObject() bool {
	return this.ObjectMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this InstrumentPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsObject() {
		child = this.GetObject().JSONLDContext()
	} else if this.IsLink() {
		child = this.GetLink().JSONLDContext()
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
func (this InstrumentPropertyIterator) KindIndex() int {
	if this.IsObject() {
		return 0
	}
	if this.IsLink() {
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
func (this InstrumentPropertyIterator) LessThan(o vocab.InstrumentPropertyIteratorInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsObject() {
		return this.GetObject().LessThan(o.GetObject())
	} else if this.IsLink() {
		return this.GetLink().LessThan(o.GetLink())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "instrument".
func (this InstrumentPropertyIterator) Name() string {
	return "instrument"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this InstrumentPropertyIterator) Next() vocab.InstrumentPropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this InstrumentPropertyIterator) Prev() vocab.InstrumentPropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *InstrumentPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetLink sets the value of this property. Calling IsLink afterwards returns true.
func (this *InstrumentPropertyIterator) SetLink(v vocab.LinkInterface) {
	this.clear()
	this.LinkMember = v
}

// SetObject sets the value of this property. Calling IsObject afterwards returns
// true.
func (this *InstrumentPropertyIterator) SetObject(v vocab.ObjectInterface) {
	this.clear()
	this.ObjectMember = v
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *InstrumentPropertyIterator) clear() {
	this.ObjectMember = nil
	this.LinkMember = nil
	this.unknown = nil
	this.iri = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this InstrumentPropertyIterator) serialize() (interface{}, error) {
	if this.IsObject() {
		return this.GetObject().Serialize()
	} else if this.IsLink() {
		return this.GetLink().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// InstrumentProperty is the non-functional property "instrument". It is permitted
// to have one or more values, and of different value types.
type InstrumentProperty struct {
	properties []*InstrumentPropertyIterator
	alias      string
}

// DeserializeInstrumentProperty creates a "instrument" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeInstrumentProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.InstrumentPropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	var this *InstrumentProperty
	propName := "instrument"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "instrument")
	}
	if i, ok := m[propName]; ok {
		this := &InstrumentProperty{
			alias:      alias,
			properties: []*InstrumentPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeInstrumentPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeInstrumentPropertyIterator(i, aliasMap); err != nil {
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

// NewInstrumentProperty creates a new instrument property.
func NewInstrumentProperty() *InstrumentProperty {
	return &InstrumentProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property
// "instrument"
func (this *InstrumentProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &InstrumentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendLink appends a Link value to the back of a list of the property
// "instrument". Invalidates iterators that are traversing using Prev.
func (this *InstrumentProperty) AppendLink(v vocab.LinkInterface) {
	this.properties = append(this.properties, &InstrumentPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendObject appends a Object value to the back of a list of the property
// "instrument". Invalidates iterators that are traversing using Prev.
func (this *InstrumentProperty) AppendObject(v vocab.ObjectInterface) {
	this.properties = append(this.properties, &InstrumentPropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this InstrumentProperty) At(index int) vocab.InstrumentPropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this InstrumentProperty) Begin() vocab.InstrumentPropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this InstrumentProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this InstrumentProperty) End() vocab.InstrumentPropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this InstrumentProperty) JSONLDContext() map[string]string {
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
func (this InstrumentProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "instrument" property.
func (this InstrumentProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this InstrumentProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetObject()
			rhs := this.properties[j].GetObject()
			return lhs.LessThan(rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetLink()
			rhs := this.properties[j].GetLink()
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
func (this InstrumentProperty) LessThan(o vocab.InstrumentPropertyInterface) bool {
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

// Name returns the name of this property: "instrument".
func (this InstrumentProperty) Name() string {
	return "instrument"
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "instrument".
func (this *InstrumentProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*InstrumentPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLink prepends a Link value to the front of a list of the property
// "instrument". Invalidates all iterators.
func (this *InstrumentProperty) PrependLink(v vocab.LinkInterface) {
	this.properties = append([]*InstrumentPropertyIterator{{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependObject prepends a Object value to the front of a list of the property
// "instrument". Invalidates all iterators.
func (this *InstrumentProperty) PrependObject(v vocab.ObjectInterface) {
	this.properties = append([]*InstrumentPropertyIterator{{
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
// "instrument", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *InstrumentProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &InstrumentPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this InstrumentProperty) Serialize() (interface{}, error) {
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

// SetIRI sets an IRI value to be at the specified index for the property
// "instrument". Panics if the index is out of bounds.
func (this *InstrumentProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &InstrumentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetLink sets a Link value to be at the specified index for the property
// "instrument". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *InstrumentProperty) SetLink(idx int, v vocab.LinkInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &InstrumentPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetObject sets a Object value to be at the specified index for the property
// "instrument". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *InstrumentProperty) SetObject(idx int, v vocab.ObjectInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &InstrumentPropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// Swap swaps the location of values at two indices for the "instrument" property.
func (this InstrumentProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
