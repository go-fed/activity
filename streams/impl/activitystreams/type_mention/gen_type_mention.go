package typemention

import vocab "github.com/go-fed/activity/streams/vocab"

// A specialized Link that represents an @mention.
//
// Example 58 (https://www.w3.org/TR/activitystreams-vocabulary/#ex181-jsonld):
//   {
//     "name": "Joe",
//     "summary": "Mention of Joe by Carrie in her note",
//     "type": "Mention",
//     "url": "http://example.org/joe"
//   }
type Mention struct {
	AttributedTo vocab.AttributedToPropertyInterface
	Height       vocab.HeightPropertyInterface
	Href         vocab.HrefPropertyInterface
	Hreflang     vocab.HreflangPropertyInterface
	Id           vocab.IdPropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Rel          vocab.RelPropertyInterface
	Type         vocab.TypePropertyInterface
	Width        vocab.WidthPropertyInterface
	alias        string
	unknown      map[string]interface{}
}

// DeserializeMention creates a Mention from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeMention(m map[string]interface{}, aliasMap map[string]string) (*Mention, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	this := &Mention{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.AttributedTo = p
	}
	if p, err := mgr.DeserializeHeightPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Height = p
	}
	if p, err := mgr.DeserializeHrefPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Href = p
	}
	if p, err := mgr.DeserializeHreflangPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Hreflang = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Id = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.MediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Name = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Preview = p
	}
	if p, err := mgr.DeserializeRelPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Rel = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Type = p
	}
	if p, err := mgr.DeserializeWidthPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Width = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "attributedTo" {
			continue
		} else if k == "height" {
			continue
		} else if k == "href" {
			continue
		} else if k == "hreflang" {
			continue
		} else if k == "id" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "rel" {
			continue
		} else if k == "type" {
			continue
		} else if k == "width" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// MentionExtends returns true if the Mention type extends from the other type.
func MentionExtends(other vocab.Type) bool {
	extensions := []string{"Link"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// MentionIsDisjointWith returns true if the other provided type is disjoint with
// the Mention type.
func MentionIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Object", "Note", "Tombstone", "Relationship", "Collection", "CollectionPage", "OrderedCollectionPage", "OrderedCollection", "Person", "Application", "Activity", "Reject", "TentativeReject", "Read", "View", "Flag", "Update", "Accept", "TentativeAccept", "Undo", "Offer", "Invite", "Follow", "Create", "Move", "Listen", "Remove", "Add", "Join", "Delete", "Dislike", "Ignore", "Block", "IntransitiveActivity", "Travel", "Arrive", "Question", "Leave", "Announce", "Like", "Group", "Event", "Service", "Profile", "Article", "Organization", "Document", "Audio", "Page", "Image", "Video", "Place"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// MentionIsExtendedBy returns true if the other provided type extends from the
// Mention type.
func MentionIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// NewMention creates a new Mention type
func NewMention() *Mention {
	return &Mention{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Mention) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetHeight returns the "height" property if it exists, and nil otherwise.
func (this Mention) GetHeight() vocab.HeightPropertyInterface {
	return this.Height
}

// GetHref returns the "href" property if it exists, and nil otherwise.
func (this Mention) GetHref() vocab.HrefPropertyInterface {
	return this.Href
}

// GetHreflang returns the "hreflang" property if it exists, and nil otherwise.
func (this Mention) GetHreflang() vocab.HreflangPropertyInterface {
	return this.Hreflang
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this Mention) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Mention) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this Mention) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Mention) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetRel returns the "rel" property if it exists, and nil otherwise.
func (this Mention) GetRel() vocab.RelPropertyInterface {
	return this.Rel
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this Mention) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the Mention type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this Mention) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetWidth returns the "width" property if it exists, and nil otherwise.
func (this Mention) GetWidth() vocab.WidthPropertyInterface {
	return this.Width
}

// IsExtending returns true if the Mention type extends from the other type.
func (this Mention) IsExtending(other vocab.Type) bool {
	return MentionExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this Mention) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.AttributedTo, m)
	m = this.helperJSONLDContext(this.Height, m)
	m = this.helperJSONLDContext(this.Href, m)
	m = this.helperJSONLDContext(this.Hreflang, m)
	m = this.helperJSONLDContext(this.Id, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Rel, m)
	m = this.helperJSONLDContext(this.Type, m)
	m = this.helperJSONLDContext(this.Width, m)

	return m
}

// LessThan computes if this Mention is lesser, with an arbitrary but stable
// determination.
func (this Mention) LessThan(o vocab.MentionInterface) bool {
	// Begin: Compare known properties
	// Compare property "attributedTo"
	if this.AttributedTo.LessThan(o.GetAttributedTo()) {
		return true
	} else if o.GetAttributedTo().LessThan(this.AttributedTo) {
		return false
	}
	// Compare property "height"
	if this.Height.LessThan(o.GetHeight()) {
		return true
	} else if o.GetHeight().LessThan(this.Height) {
		return false
	}
	// Compare property "href"
	if this.Href.LessThan(o.GetHref()) {
		return true
	} else if o.GetHref().LessThan(this.Href) {
		return false
	}
	// Compare property "hreflang"
	if this.Hreflang.LessThan(o.GetHreflang()) {
		return true
	} else if o.GetHreflang().LessThan(this.Hreflang) {
		return false
	}
	// Compare property "id"
	if this.Id.LessThan(o.GetId()) {
		return true
	} else if o.GetId().LessThan(this.Id) {
		return false
	}
	// Compare property "mediaType"
	if this.MediaType.LessThan(o.GetMediaType()) {
		return true
	} else if o.GetMediaType().LessThan(this.MediaType) {
		return false
	}
	// Compare property "name"
	if this.Name.LessThan(o.GetName()) {
		return true
	} else if o.GetName().LessThan(this.Name) {
		return false
	}
	// Compare property "preview"
	if this.Preview.LessThan(o.GetPreview()) {
		return true
	} else if o.GetPreview().LessThan(this.Preview) {
		return false
	}
	// Compare property "rel"
	if this.Rel.LessThan(o.GetRel()) {
		return true
	} else if o.GetRel().LessThan(this.Rel) {
		return false
	}
	// Compare property "type"
	if this.Type.LessThan(o.GetType()) {
		return true
	} else if o.GetType().LessThan(this.Type) {
		return false
	}
	// Compare property "width"
	if this.Width.LessThan(o.GetWidth()) {
		return true
	} else if o.GetWidth().LessThan(this.Width) {
		return false
	}
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
func (this Mention) Serialize() (interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "attributedTo"
	if i, err := this.AttributedTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.AttributedTo.Name()] = i
	}
	// Maybe serialize property "height"
	if i, err := this.Height.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Height.Name()] = i
	}
	// Maybe serialize property "href"
	if i, err := this.Href.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Href.Name()] = i
	}
	// Maybe serialize property "hreflang"
	if i, err := this.Hreflang.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Hreflang.Name()] = i
	}
	// Maybe serialize property "id"
	if i, err := this.Id.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Id.Name()] = i
	}
	// Maybe serialize property "mediaType"
	if i, err := this.MediaType.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.MediaType.Name()] = i
	}
	// Maybe serialize property "name"
	if i, err := this.Name.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Name.Name()] = i
	}
	// Maybe serialize property "preview"
	if i, err := this.Preview.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Preview.Name()] = i
	}
	// Maybe serialize property "rel"
	if i, err := this.Rel.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Rel.Name()] = i
	}
	// Maybe serialize property "type"
	if i, err := this.Type.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Type.Name()] = i
	}
	// Maybe serialize property "width"
	if i, err := this.Width.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Width.Name()] = i
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

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Mention) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetHeight returns the "height" property if it exists, and nil otherwise.
func (this Mention) SetHeight(i vocab.HeightPropertyInterface) {
	this.Height = i
}

// SetHref returns the "href" property if it exists, and nil otherwise.
func (this Mention) SetHref(i vocab.HrefPropertyInterface) {
	this.Href = i
}

// SetHreflang returns the "hreflang" property if it exists, and nil otherwise.
func (this Mention) SetHreflang(i vocab.HreflangPropertyInterface) {
	this.Hreflang = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this Mention) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Mention) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this Mention) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Mention) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetRel returns the "rel" property if it exists, and nil otherwise.
func (this Mention) SetRel(i vocab.RelPropertyInterface) {
	this.Rel = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this Mention) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetWidth returns the "width" property if it exists, and nil otherwise.
func (this Mention) SetWidth(i vocab.WidthPropertyInterface) {
	this.Width = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this Mention) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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
