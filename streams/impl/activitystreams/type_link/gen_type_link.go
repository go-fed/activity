package typelink

import vocab "github.com/go-fed/activity/streams/vocab"

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988]. Many of the
// properties defined by the Activity Vocabulary allow values that are either
// instances of Object or Link. When a Link is used, it establishes a
// qualified relation connecting the subject (the containing object) to the
// resource identified by the href. Properties of the Link are properties of
// the reference as opposed to properties of the resource.
//
// Example 2 (https://www.w3.org/TR/activitystreams-vocabulary/#ex2-jsonld):
//   {
//     "hreflang": "en",
//     "mediaType": "text/html",
//     "name": "An example link",
//     "type": "owl:Class",
//     "url": "http://example.org/abc"
//   }
type Link struct {
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

// DeserializeLink creates a Link from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeLink(m map[string]interface{}, aliasMap map[string]string) (*Link, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	this := &Link{
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

// LinkExtends returns true if the Link type extends from the other type.
func LinkExtends(other vocab.Type) bool {
	// Shortcut implementation: this does not extend anything.
	return false
}

// LinkIsDisjointWith returns true if the other provided type is disjoint with the
// Link type.
func LinkIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Object", "Note", "Tombstone", "Relationship", "Collection", "CollectionPage", "OrderedCollectionPage", "OrderedCollection", "Person", "Application", "Activity", "Reject", "TentativeReject", "Read", "View", "Flag", "Update", "Accept", "TentativeAccept", "Undo", "Offer", "Invite", "Follow", "Create", "Move", "Listen", "Remove", "Add", "Join", "Delete", "Dislike", "Ignore", "Block", "IntransitiveActivity", "Travel", "Arrive", "Question", "Leave", "Announce", "Like", "Group", "Event", "Service", "Profile", "Article", "Organization", "Document", "Audio", "Page", "Image", "Video", "Place"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// LinkIsExtendedBy returns true if the other provided type extends from the Link
// type.
func LinkIsExtendedBy(other vocab.Type) bool {
	extensions := []string{"Mention"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// NewLink creates a new Link type
func NewLink() *Link {
	return &Link{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Link) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetHeight returns the "height" property if it exists, and nil otherwise.
func (this Link) GetHeight() vocab.HeightPropertyInterface {
	return this.Height
}

// GetHref returns the "href" property if it exists, and nil otherwise.
func (this Link) GetHref() vocab.HrefPropertyInterface {
	return this.Href
}

// GetHreflang returns the "hreflang" property if it exists, and nil otherwise.
func (this Link) GetHreflang() vocab.HreflangPropertyInterface {
	return this.Hreflang
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this Link) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Link) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this Link) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Link) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetRel returns the "rel" property if it exists, and nil otherwise.
func (this Link) GetRel() vocab.RelPropertyInterface {
	return this.Rel
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this Link) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the Link type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this Link) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetWidth returns the "width" property if it exists, and nil otherwise.
func (this Link) GetWidth() vocab.WidthPropertyInterface {
	return this.Width
}

// IsExtending returns true if the Link type extends from the other type.
func (this Link) IsExtending(other vocab.Type) bool {
	return LinkExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this Link) JSONLDContext() map[string]string {
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

// LessThan computes if this Link is lesser, with an arbitrary but stable
// determination.
func (this Link) LessThan(o vocab.LinkInterface) bool {
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
func (this Link) Serialize() (interface{}, error) {
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
func (this Link) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetHeight returns the "height" property if it exists, and nil otherwise.
func (this Link) SetHeight(i vocab.HeightPropertyInterface) {
	this.Height = i
}

// SetHref returns the "href" property if it exists, and nil otherwise.
func (this Link) SetHref(i vocab.HrefPropertyInterface) {
	this.Href = i
}

// SetHreflang returns the "hreflang" property if it exists, and nil otherwise.
func (this Link) SetHreflang(i vocab.HreflangPropertyInterface) {
	this.Hreflang = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this Link) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Link) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this Link) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Link) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetRel returns the "rel" property if it exists, and nil otherwise.
func (this Link) SetRel(i vocab.RelPropertyInterface) {
	this.Rel = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this Link) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetWidth returns the "width" property if it exists, and nil otherwise.
func (this Link) SetWidth(i vocab.WidthPropertyInterface) {
	this.Width = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this Link) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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
