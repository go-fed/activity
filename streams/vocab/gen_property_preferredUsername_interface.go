package vocab

import "net/url"

// A short username which may be used to refer to the actor, with no uniqueness
// guarantees
type PreferredUsernamePropertyInterface interface {
	// Clear ensures no value and no language map for this property is set.
	// Calling HasAny or any of the 'Is' methods afterwards will return
	// false.
	Clear()
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLangString returns the value of this property. When IsLangString
	// returns false, GetLangString will return an arbitrary value.
	GetLangString() map[string]string
	// GetLanguage returns the value for the specified BCP47 language code, or
	// an empty string if it is either not a language map or no value is
	// present.
	GetLanguage(bcp47 string) string
	// GetString returns the value of this property. When IsString returns
	// false, GetString will return an arbitrary value.
	GetString() string
	// HasAny returns true if any of the values are set, except for the
	// natural language map. When true, the specific has, getter, and
	// setter methods may be used to determine what kind of value there is
	// to access and set this property. To determine if the property was
	// set as a natural language map, use the IsLangString method instead.
	HasAny() bool
	// HasLanguage returns true if the natural language map has an entry for
	// the specified BCP47 language code.
	HasLanguage(bcp47 string) bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsLangString returns true if this property has a type of "langString".
	// When true, use the GetLangString and SetLangString methods to
	// access and set this property.. To determine if the property was set
	// as a natural language map, use the IsLangString method instead.
	IsLangString() bool
	// IsString returns true if this property has a type of "string". When
	// true, use the GetString and SetString methods to access and set
	// this property.. To determine if the property was set as a natural
	// language map, use the IsLangString method instead.
	IsString() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o PreferredUsernamePropertyInterface) bool
	// Name returns the name of this property: "preferredUsername".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLangString sets the value of this property and clears the natural
	// language map. Calling IsLangString afterwards will return true.
	// Calling IsLangString afterwards returns false.
	SetLangString(v map[string]string)
	// SetLanguage sets the value for the specified BCP47 language code.
	SetLanguage(bcp47, value string)
	// SetString sets the value of this property and clears the natural
	// language map. Calling IsString afterwards will return true. Calling
	// IsLangString afterwards returns false.
	SetString(v string)
}
