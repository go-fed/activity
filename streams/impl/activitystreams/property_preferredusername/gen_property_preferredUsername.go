package propertypreferredusername

import (
	"fmt"
	langstring "github.com/go-fed/activity/streams/values/langString"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// PreferredUsernameProperty is the functional property "preferredUsername". It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type PreferredUsernameProperty struct {
	stringMember     string
	hasStringMember  bool
	langStringMember map[string]string
	unknown          interface{}
	iri              *url.URL
	alias            string
	langMap          map[string]string
}

// DeserializePreferredUsernameProperty creates a "preferredUsername" property
// from an interface representation that has been unmarshalled from a text or
// binary format.
func DeserializePreferredUsernameProperty(m map[string]interface{}, aliasMap map[string]string) (*PreferredUsernameProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "preferredUsername"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "preferredUsername")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &PreferredUsernameProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := string1.DeserializeString(i); err == nil {
			this := &PreferredUsernameProperty{
				alias:           alias,
				hasStringMember: true,
				stringMember:    v,
			}
			return this, nil
		} else if v, err := langstring.DeserializeLangString(i); err == nil {
			this := &PreferredUsernameProperty{
				alias:            alias,
				langStringMember: v,
			}
			return this, nil
		}
		this := &PreferredUsernameProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewPreferredUsernameProperty creates a new preferredUsername property.
func NewPreferredUsernameProperty() *PreferredUsernameProperty {
	return &PreferredUsernameProperty{alias: ""}
}

// Clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *PreferredUsernameProperty) Clear() {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	this.langMap = nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this PreferredUsernameProperty) GetIRI() *url.URL {
	return this.iri
}

// GetLangString returns the value of this property. When IsLangString returns
// false, GetLangString will return an arbitrary value.
func (this PreferredUsernameProperty) GetLangString() map[string]string {
	return this.langStringMember
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this PreferredUsernameProperty) GetLanguage(bcp47 string) string {
	if this.langMap == nil {
		return ""
	} else if v, ok := this.langMap[bcp47]; ok {
		return v
	} else {
		return ""
	}
}

// GetString returns the value of this property. When IsString returns false,
// GetString will return an arbitrary value.
func (this PreferredUsernameProperty) GetString() string {
	return this.stringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsLanguageMap method instead.
func (this PreferredUsernameProperty) HasAny() bool {
	return this.IsString() ||
		this.IsLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this PreferredUsernameProperty) HasLanguage(bcp47 string) bool {
	if this.langMap == nil {
		return false
	} else {
		_, ok := this.langMap[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this PreferredUsernameProperty) IsIRI() bool {
	return this.iri != nil
}

// IsLangString returns true if this property has a type of "langString". When
// true, use the GetLangString and SetLangString methods to access and set
// this property.. To determine if the property was set as a natural language
// map, use the IsLanguageMap method instead.
func (this PreferredUsernameProperty) IsLangString() bool {
	return this.langStringMember != nil
}

// IsLanguageMap determines if this property is represented by a natural language
// map. When true, use HasLanguage, GetLanguage, and SetLanguage methods to
// access and mutate the natural language map. The Clear method can be used to
// clear the natural language map. Note that this method is only used for
// natural language representations, and does not determine the presence nor
// absence of other values for this property.
func (this PreferredUsernameProperty) IsLanguageMap() bool {
	return this.langMap != nil
}

// IsString returns true if this property has a type of "string". When true, use
// the GetString and SetString methods to access and set this property.. To
// determine if the property was set as a natural language map, use the
// IsLanguageMap method instead.
func (this PreferredUsernameProperty) IsString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this PreferredUsernameProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string

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
func (this PreferredUsernameProperty) KindIndex() int {
	if this.IsString() {
		return 0
	}
	if this.IsLangString() {
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
func (this PreferredUsernameProperty) LessThan(o vocab.PreferredUsernamePropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsString() {
		return string1.LessString(this.GetString(), o.GetString())
	} else if this.IsLangString() {
		return langstring.LessLangString(this.GetLangString(), o.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "preferredUsername".
func (this PreferredUsernameProperty) Name() string {
	return "preferredUsername"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this PreferredUsernameProperty) Serialize() (interface{}, error) {
	if this.IsString() {
		return string1.SerializeString(this.GetString())
	} else if this.IsLangString() {
		return langstring.SerializeLangString(this.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *PreferredUsernameProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetLangString sets the value of this property and clears the natural language
// map. Calling IsLangString afterwards will return true. Calling
// IsLanguageMap afterwards returns false.
func (this *PreferredUsernameProperty) SetLangString(v map[string]string) {
	this.Clear()
	this.langStringMember = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *PreferredUsernameProperty) SetLanguage(bcp47, value string) {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	if this.langMap == nil {
		this.langMap = make(map[string]string)
	}
	this.langMap[bcp47] = value
}

// SetString sets the value of this property and clears the natural language map.
// Calling IsString afterwards will return true. Calling IsLanguageMap
// afterwards returns false.
func (this *PreferredUsernameProperty) SetString(v string) {
	this.Clear()
	this.stringMember = v
	this.hasStringMember = true
}
