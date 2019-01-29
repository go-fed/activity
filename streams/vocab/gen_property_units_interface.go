package vocab

import "net/url"

// Specifies the measurement units for the radius and altitude properties on a
// Place object. If not specified, the default is assumed to be "m" for
// "meters".
//
// Example 136 (https://www.w3.org/TR/activitystreams-vocabulary/#ex157-jsonld):
//   {
//     "latitude": 36.75,
//     "longitude": 119.7667,
//     "name": "Fresno Area",
//     "radius": 15,
//     "type": "Place",
//     "units": "miles"
//   }
type UnitsPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
	// GetAnyURI returns the value of this property. When IsAnyURI returns
	// false, GetAnyURI will return an arbitrary value.
	GetAnyURI() *url.URL
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetString returns the value of this property. When IsString returns
	// false, GetString will return an arbitrary value.
	GetString() string
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsAnyURI returns true if this property has a type of "anyURI". When
	// true, use the GetAnyURI and SetAnyURI methods to access and set
	// this property.
	IsAnyURI() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsString returns true if this property has a type of "string". When
	// true, use the GetString and SetString methods to access and set
	// this property.
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
	LessThan(o UnitsPropertyInterface) bool
	// Name returns the name of this property: "units".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetAnyURI sets the value of this property. Calling IsAnyURI afterwards
	// returns true.
	SetAnyURI(v *url.URL)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetString sets the value of this property. Calling IsString afterwards
	// returns true.
	SetString(v string)
}
