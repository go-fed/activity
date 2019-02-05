package typecollection

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeAltitudePropertyActivityStreams returns the deserialization
	// method for the "AltitudePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeAltitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AltitudePropertyInterface, error)
	// DeserializeAttachmentPropertyActivityStreams returns the
	// deserialization method for the "AttachmentPropertyInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeAttachmentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AttachmentPropertyInterface, error)
	// DeserializeAttributedToPropertyActivityStreams returns the
	// deserialization method for the "AttributedToPropertyInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AttributedToPropertyInterface, error)
	// DeserializeAudiencePropertyActivityStreams returns the deserialization
	// method for the "AudiencePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeAudiencePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AudiencePropertyInterface, error)
	// DeserializeBccPropertyActivityStreams returns the deserialization
	// method for the "BccPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeBccPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BccPropertyInterface, error)
	// DeserializeBtoPropertyActivityStreams returns the deserialization
	// method for the "BtoPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeBtoPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BtoPropertyInterface, error)
	// DeserializeCcPropertyActivityStreams returns the deserialization method
	// for the "CcPropertyInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCcPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CcPropertyInterface, error)
	// DeserializeContentPropertyActivityStreams returns the deserialization
	// method for the "ContentPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeContentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ContentPropertyInterface, error)
	// DeserializeContextPropertyActivityStreams returns the deserialization
	// method for the "ContextPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeContextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ContextPropertyInterface, error)
	// DeserializeCurrentPropertyActivityStreams returns the deserialization
	// method for the "CurrentPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeCurrentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CurrentPropertyInterface, error)
	// DeserializeDurationPropertyActivityStreams returns the deserialization
	// method for the "DurationPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeDurationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DurationPropertyInterface, error)
	// DeserializeEndTimePropertyActivityStreams returns the deserialization
	// method for the "EndTimePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeEndTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.EndTimePropertyInterface, error)
	// DeserializeFirstPropertyActivityStreams returns the deserialization
	// method for the "FirstPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeFirstPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FirstPropertyInterface, error)
	// DeserializeGeneratorPropertyActivityStreams returns the deserialization
	// method for the "GeneratorPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeGeneratorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.GeneratorPropertyInterface, error)
	// DeserializeIconPropertyActivityStreams returns the deserialization
	// method for the "IconPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeIconPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IconPropertyInterface, error)
	// DeserializeIdPropertyActivityStreams returns the deserialization method
	// for the "IdPropertyInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeIdPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IdPropertyInterface, error)
	// DeserializeImagePropertyActivityStreams returns the deserialization
	// method for the "ImagePropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeImagePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ImagePropertyInterface, error)
	// DeserializeInReplyToPropertyActivityStreams returns the deserialization
	// method for the "InReplyToPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeInReplyToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.InReplyToPropertyInterface, error)
	// DeserializeItemsPropertyActivityStreams returns the deserialization
	// method for the "ItemsPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ItemsPropertyInterface, error)
	// DeserializeLastPropertyActivityStreams returns the deserialization
	// method for the "LastPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeLastPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LastPropertyInterface, error)
	// DeserializeLikesPropertyActivityStreams returns the deserialization
	// method for the "LikesPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeLikesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LikesPropertyInterface, error)
	// DeserializeLocationPropertyActivityStreams returns the deserialization
	// method for the "LocationPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeLocationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LocationPropertyInterface, error)
	// DeserializeMediaTypePropertyActivityStreams returns the deserialization
	// method for the "MediaTypePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MediaTypePropertyInterface, error)
	// DeserializeNamePropertyActivityStreams returns the deserialization
	// method for the "NamePropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NamePropertyInterface, error)
	// DeserializeObjectPropertyActivityStreams returns the deserialization
	// method for the "ObjectPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeObjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ObjectPropertyInterface, error)
	// DeserializePreviewPropertyActivityStreams returns the deserialization
	// method for the "PreviewPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PreviewPropertyInterface, error)
	// DeserializePublishedPropertyActivityStreams returns the deserialization
	// method for the "PublishedPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializePublishedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PublishedPropertyInterface, error)
	// DeserializeRepliesPropertyActivityStreams returns the deserialization
	// method for the "RepliesPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeRepliesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RepliesPropertyInterface, error)
	// DeserializeSharesPropertyActivityStreams returns the deserialization
	// method for the "SharesPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeSharesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.SharesPropertyInterface, error)
	// DeserializeStartTimePropertyActivityStreams returns the deserialization
	// method for the "StartTimePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeStartTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.StartTimePropertyInterface, error)
	// DeserializeSummaryPropertyActivityStreams returns the deserialization
	// method for the "SummaryPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.SummaryPropertyInterface, error)
	// DeserializeTagPropertyActivityStreams returns the deserialization
	// method for the "TagPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeTagPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TagPropertyInterface, error)
	// DeserializeToPropertyActivityStreams returns the deserialization method
	// for the "ToPropertyInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ToPropertyInterface, error)
	// DeserializeTotalItemsPropertyActivityStreams returns the
	// deserialization method for the "TotalItemsPropertyInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeTotalItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TotalItemsPropertyInterface, error)
	// DeserializeTypePropertyActivityStreams returns the deserialization
	// method for the "TypePropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TypePropertyInterface, error)
	// DeserializeUpdatedPropertyActivityStreams returns the deserialization
	// method for the "UpdatedPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeUpdatedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UpdatedPropertyInterface, error)
	// DeserializeUrlPropertyActivityStreams returns the deserialization
	// method for the "UrlPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeUrlPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UrlPropertyInterface, error)
}

// jsonldContexter is a private interface to determine the JSON-LD contexts and
// aliases needed for functional and non-functional properties. It is a helper
// interface for this implementation.
type jsonldContexter interface {
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
