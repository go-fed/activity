package vocab

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
type ActivityStreamsLink interface {
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
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the Link type.
	// Note that this should not be used by app developers. It is only
	// used to help determine which implementation is LessThan the other.
	// Developers who are creating a different implementation of this
	// type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Link type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Link is lesser, with an arbitrary but stable
	// determination.
	LessThan(o ActivityStreamsLink) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsAttributedTo sets the "attributedTo" property.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsHeight sets the "height" property.
	SetActivityStreamsHeight(i ActivityStreamsHeightProperty)
	// SetActivityStreamsHref sets the "href" property.
	SetActivityStreamsHref(i ActivityStreamsHrefProperty)
	// SetActivityStreamsHreflang sets the "hreflang" property.
	SetActivityStreamsHreflang(i ActivityStreamsHreflangProperty)
	// SetActivityStreamsId sets the "id" property.
	SetActivityStreamsId(i ActivityStreamsIdProperty)
	// SetActivityStreamsMediaType sets the "mediaType" property.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName sets the "name" property.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsPreview sets the "preview" property.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsRel sets the "rel" property.
	SetActivityStreamsRel(i ActivityStreamsRelProperty)
	// SetActivityStreamsSummary sets the "summary" property.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsType sets the "type" property.
	SetActivityStreamsType(i ActivityStreamsTypeProperty)
	// SetActivityStreamsWidth sets the "width" property.
	SetActivityStreamsWidth(i ActivityStreamsWidthProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
