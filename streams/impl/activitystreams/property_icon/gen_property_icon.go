package propertyicon

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// IconPropertyIterator is an iterator for a property. It is permitted to be one
// of multiple value types. At most, one type of value can be present, or none
// at all. Setting a value will clear the other types of values so that only
// one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type IconPropertyIterator struct {
	ImageMember   vocab.ImageInterface
	LinkMember    vocab.LinkInterface
	MentionMember vocab.MentionInterface
	unknown       interface{}
	iri           *url.URL
	alias         string
	myIdx         int
	parent        vocab.IconPropertyInterface
}

// NewIconPropertyIterator creates a new icon property.
func NewIconPropertyIterator() *IconPropertyIterator {
	return &IconPropertyIterator{alias: ""}
}

// deserializeIconPropertyIterator creates an iterator from an element that has
// been unmarshalled from a text or binary format.
func deserializeIconPropertyIterator(i interface{}, aliasMap map[string]string) (*IconPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &IconPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap); err == nil {
			this := &IconPropertyIterator{
				ImageMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
			this := &IconPropertyIterator{
				LinkMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
			this := &IconPropertyIterator{
				MentionMember: v,
				alias:         alias,
			}
			return this, nil
		}
	}
	this := &IconPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
	return nil, fmt.Errorf("could not deserialize %q property", "icon")
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this IconPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetImage returns the value of this property. When IsImage returns false,
// GetImage will return an arbitrary value.
func (this IconPropertyIterator) GetImage() vocab.ImageInterface {
	return this.ImageMember
}

// GetLink returns the value of this property. When IsLink returns false, GetLink
// will return an arbitrary value.
func (this IconPropertyIterator) GetLink() vocab.LinkInterface {
	return this.LinkMember
}

// GetMention returns the value of this property. When IsMention returns false,
// GetMention will return an arbitrary value.
func (this IconPropertyIterator) GetMention() vocab.MentionInterface {
	return this.MentionMember
}

// HasAny returns true if any of the different values is set.
func (this IconPropertyIterator) HasAny() bool {
	return this.IsImage() ||
		this.IsLink() ||
		this.IsMention() ||
		this.iri != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this IconPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsImage returns true if this property has a type of "Image". When true, use the
// GetImage and SetImage methods to access and set this property.
func (this IconPropertyIterator) IsImage() bool {
	return this.ImageMember != nil
}

// IsLink returns true if this property has a type of "Link". When true, use the
// GetLink and SetLink methods to access and set this property.
func (this IconPropertyIterator) IsLink() bool {
	return this.LinkMember != nil
}

// IsMention returns true if this property has a type of "Mention". When true, use
// the GetMention and SetMention methods to access and set this property.
func (this IconPropertyIterator) IsMention() bool {
	return this.MentionMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this IconPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsImage() {
		child = this.GetImage().JSONLDContext()
	} else if this.IsLink() {
		child = this.GetLink().JSONLDContext()
	} else if this.IsMention() {
		child = this.GetMention().JSONLDContext()
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
func (this IconPropertyIterator) KindIndex() int {
	if this.IsImage() {
		return 0
	}
	if this.IsLink() {
		return 1
	}
	if this.IsMention() {
		return 2
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
func (this IconPropertyIterator) LessThan(o vocab.IconPropertyIteratorInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsImage() {
		return this.GetImage().LessThan(o.GetImage())
	} else if this.IsLink() {
		return this.GetLink().LessThan(o.GetLink())
	} else if this.IsMention() {
		return this.GetMention().LessThan(o.GetMention())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "icon".
func (this IconPropertyIterator) Name() string {
	return "icon"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this IconPropertyIterator) Next() vocab.IconPropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this IconPropertyIterator) Prev() vocab.IconPropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *IconPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetImage sets the value of this property. Calling IsImage afterwards returns
// true.
func (this *IconPropertyIterator) SetImage(v vocab.ImageInterface) {
	this.clear()
	this.ImageMember = v
}

// SetLink sets the value of this property. Calling IsLink afterwards returns true.
func (this *IconPropertyIterator) SetLink(v vocab.LinkInterface) {
	this.clear()
	this.LinkMember = v
}

// SetMention sets the value of this property. Calling IsMention afterwards
// returns true.
func (this *IconPropertyIterator) SetMention(v vocab.MentionInterface) {
	this.clear()
	this.MentionMember = v
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *IconPropertyIterator) clear() {
	this.ImageMember = nil
	this.LinkMember = nil
	this.MentionMember = nil
	this.unknown = nil
	this.iri = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this IconPropertyIterator) serialize() (interface{}, error) {
	if this.IsImage() {
		return this.GetImage().Serialize()
	} else if this.IsLink() {
		return this.GetLink().Serialize()
	} else if this.IsMention() {
		return this.GetMention().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// IconProperty is the non-functional property "icon". It is permitted to have one
// or more values, and of different value types.
type IconProperty struct {
	properties []*IconPropertyIterator
	alias      string
}

// DeserializeIconProperty creates a "icon" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeIconProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.IconPropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "icon"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "icon")
	}
	if i, ok := m[propName]; ok {
		this := &IconProperty{
			alias:      alias,
			properties: []*IconPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeIconPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeIconPropertyIterator(i, aliasMap); err != nil {
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

// NewIconProperty creates a new icon property.
func NewIconProperty() *IconProperty {
	return &IconProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "icon"
func (this *IconProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &IconPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendImage appends a Image value to the back of a list of the property "icon".
// Invalidates iterators that are traversing using Prev.
func (this *IconProperty) AppendImage(v vocab.ImageInterface) {
	this.properties = append(this.properties, &IconPropertyIterator{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendLink appends a Link value to the back of a list of the property "icon".
// Invalidates iterators that are traversing using Prev.
func (this *IconProperty) AppendLink(v vocab.LinkInterface) {
	this.properties = append(this.properties, &IconPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendMention appends a Mention value to the back of a list of the property
// "icon". Invalidates iterators that are traversing using Prev.
func (this *IconProperty) AppendMention(v vocab.MentionInterface) {
	this.properties = append(this.properties, &IconPropertyIterator{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this IconProperty) At(index int) vocab.IconPropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this IconProperty) Begin() vocab.IconPropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this IconProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this IconProperty) End() vocab.IconPropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this IconProperty) JSONLDContext() map[string]string {
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
func (this IconProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "icon" property.
func (this IconProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this IconProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetImage()
			rhs := this.properties[j].GetImage()
			return lhs.LessThan(rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetLink()
			rhs := this.properties[j].GetLink()
			return lhs.LessThan(rhs)
		} else if idx1 == 2 {
			lhs := this.properties[i].GetMention()
			rhs := this.properties[j].GetMention()
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
func (this IconProperty) LessThan(o vocab.IconPropertyInterface) bool {
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

// Name returns the name of this property: "icon".
func (this IconProperty) Name() string {
	return "icon"
}

// PrependIRI prepends an IRI value to the front of a list of the property "icon".
func (this *IconProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*IconPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependImage prepends a Image value to the front of a list of the property
// "icon". Invalidates all iterators.
func (this *IconProperty) PrependImage(v vocab.ImageInterface) {
	this.properties = append([]*IconPropertyIterator{{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLink prepends a Link value to the front of a list of the property
// "icon". Invalidates all iterators.
func (this *IconProperty) PrependLink(v vocab.LinkInterface) {
	this.properties = append([]*IconPropertyIterator{{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependMention prepends a Mention value to the front of a list of the property
// "icon". Invalidates all iterators.
func (this *IconProperty) PrependMention(v vocab.MentionInterface) {
	this.properties = append([]*IconPropertyIterator{{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "icon", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *IconProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &IconPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this IconProperty) Serialize() (interface{}, error) {
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

// SetIRI sets an IRI value to be at the specified index for the property "icon".
// Panics if the index is out of bounds.
func (this *IconProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &IconPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetImage sets a Image value to be at the specified index for the property
// "icon". Panics if the index is out of bounds. Invalidates all iterators.
func (this *IconProperty) SetImage(idx int, v vocab.ImageInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &IconPropertyIterator{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetLink sets a Link value to be at the specified index for the property "icon".
// Panics if the index is out of bounds. Invalidates all iterators.
func (this *IconProperty) SetLink(idx int, v vocab.LinkInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &IconPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetMention sets a Mention value to be at the specified index for the property
// "icon". Panics if the index is out of bounds. Invalidates all iterators.
func (this *IconProperty) SetMention(idx int, v vocab.MentionInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &IconPropertyIterator{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// Swap swaps the location of values at two indices for the "icon" property.
func (this IconProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
