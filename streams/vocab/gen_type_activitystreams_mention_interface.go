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
type ActivityStreamsMention interface {
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() ActivityStreamsAttributedToProperty
	// GetActivityStreamsHeight returns the "height" property if it exists,
	// and nil otherwise.
	GetActivityStreamsHeight() ActivityStreamsHeightProperty
	// GetActivityStreamsHref returns the "href" property if it exists, and
	// nil otherwise.
	GetActivityStreamsHref() ActivityStreamsHrefProperty
	// GetActivityStreamsHreflang returns the "hreflang" property if it
	// exists, and nil otherwise.
	GetActivityStreamsHreflang() ActivityStreamsHreflangProperty
	// GetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	GetActivityStreamsId() ActivityStreamsIdProperty
	// GetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	GetActivityStreamsMediaType() ActivityStreamsMediaTypeProperty
	// GetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	GetActivityStreamsName() ActivityStreamsNameProperty
	// GetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPreview() ActivityStreamsPreviewProperty
	// GetActivityStreamsRel returns the "rel" property if it exists, and nil
	// otherwise.
	GetActivityStreamsRel() ActivityStreamsRelProperty
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsType returns the "type" property if it exists, and
	// nil otherwise.
	GetActivityStreamsType() ActivityStreamsTypeProperty
	// GetActivityStreamsWidth returns the "width" property if it exists, and
	// nil otherwise.
	GetActivityStreamsWidth() ActivityStreamsWidthProperty
	// GetName returns the name of this type.
	GetName() string
	// GetUnknownProperties returns the unknown properties for the Mention
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Mention type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Mention is lesser, with an arbitrary but
	// stable determination.
	LessThan(o ActivityStreamsMention) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsHeight returns the "height" property if it exists,
	// and nil otherwise.
	SetActivityStreamsHeight(i ActivityStreamsHeightProperty)
	// SetActivityStreamsHref returns the "href" property if it exists, and
	// nil otherwise.
	SetActivityStreamsHref(i ActivityStreamsHrefProperty)
	// SetActivityStreamsHreflang returns the "hreflang" property if it
	// exists, and nil otherwise.
	SetActivityStreamsHreflang(i ActivityStreamsHreflangProperty)
	// SetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	SetActivityStreamsId(i ActivityStreamsIdProperty)
	// SetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsRel returns the "rel" property if it exists, and nil
	// otherwise.
	SetActivityStreamsRel(i ActivityStreamsRelProperty)
	// SetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsType returns the "type" property if it exists, and
	// nil otherwise.
	SetActivityStreamsType(i ActivityStreamsTypeProperty)
	// SetActivityStreamsWidth returns the "width" property if it exists, and
	// nil otherwise.
	SetActivityStreamsWidth(i ActivityStreamsWidthProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
