package vocab

// A specialized Link that represents an @mention.
//
// Example 58 (https://www.w3.org/TR/activitystreams-vocabulary/#ex181-jsonld):
//   {
//     "name": "Joe",
//     "summary": "Mention of Joe by Carrie in her note",
//     "type": "Mention",
//     "url": "http://example.org/joe"
//   }
type MentionInterface interface {
	// GetAttributedTo returns the "attributedTo" property if it exists, and
	// nil otherwise.
	GetAttributedTo() AttributedToPropertyInterface
	// GetHeight returns the "height" property if it exists, and nil otherwise.
	GetHeight() HeightPropertyInterface
	// GetHref returns the "href" property if it exists, and nil otherwise.
	GetHref() HrefPropertyInterface
	// GetHreflang returns the "hreflang" property if it exists, and nil
	// otherwise.
	GetHreflang() HreflangPropertyInterface
	// GetId returns the "id" property if it exists, and nil otherwise.
	GetId() IdPropertyInterface
	// GetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	GetMediaType() MediaTypePropertyInterface
	// GetName returns the "name" property if it exists, and nil otherwise.
	GetName() NamePropertyInterface
	// GetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	GetPreview() PreviewPropertyInterface
	// GetRel returns the "rel" property if it exists, and nil otherwise.
	GetRel() RelPropertyInterface
	// GetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	GetSummary() SummaryPropertyInterface
	// GetType returns the "type" property if it exists, and nil otherwise.
	GetType() TypePropertyInterface
	// GetUnknownProperties returns the unknown properties for the Mention
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// GetWidth returns the "width" property if it exists, and nil otherwise.
	GetWidth() WidthPropertyInterface
	// IsExtending returns true if the Mention type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Mention is lesser, with an arbitrary but
	// stable determination.
	LessThan(o MentionInterface) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetAttributedTo returns the "attributedTo" property if it exists, and
	// nil otherwise.
	SetAttributedTo(i AttributedToPropertyInterface)
	// SetHeight returns the "height" property if it exists, and nil otherwise.
	SetHeight(i HeightPropertyInterface)
	// SetHref returns the "href" property if it exists, and nil otherwise.
	SetHref(i HrefPropertyInterface)
	// SetHreflang returns the "hreflang" property if it exists, and nil
	// otherwise.
	SetHreflang(i HreflangPropertyInterface)
	// SetId returns the "id" property if it exists, and nil otherwise.
	SetId(i IdPropertyInterface)
	// SetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	SetMediaType(i MediaTypePropertyInterface)
	// SetName returns the "name" property if it exists, and nil otherwise.
	SetName(i NamePropertyInterface)
	// SetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	SetPreview(i PreviewPropertyInterface)
	// SetRel returns the "rel" property if it exists, and nil otherwise.
	SetRel(i RelPropertyInterface)
	// SetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	SetSummary(i SummaryPropertyInterface)
	// SetType returns the "type" property if it exists, and nil otherwise.
	SetType(i TypePropertyInterface)
	// SetWidth returns the "width" property if it exists, and nil otherwise.
	SetWidth(i WidthPropertyInterface)
}
