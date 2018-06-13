//
package vocab

import (
	"fmt"
	"net/url"
)

// MentionType is an interface for accepting types that extend from 'Mention'.
type MentionType interface {
	Serializer
	Deserializer
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v ObjectType)
	AppendAttributedToObject(v ObjectType)
	PrependAttributedToObject(v ObjectType)
	RemoveAttributedToObject(index int)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v LinkType)
	AppendAttributedToLink(v LinkType)
	PrependAttributedToLink(v LinkType)
	RemoveAttributedToLink(index int)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v *url.URL)
	AppendAttributedToIRI(v *url.URL)
	PrependAttributedToIRI(v *url.URL)
	RemoveAttributedToIRI(index int)
	HasHref() (ok bool)
	GetHref() (v *url.URL)
	SetHref(v *url.URL)
	HasUnknownHref() (ok bool)
	GetUnknownHref() (v interface{})
	SetUnknownHref(i interface{})
	HasId() (ok bool)
	GetId() (v *url.URL)
	SetId(v *url.URL)
	HasUnknownId() (ok bool)
	GetUnknownId() (v interface{})
	SetUnknownId(i interface{})
	RelLen() (l int)
	IsRel(index int) (ok bool)
	GetRel(index int) (v string)
	AppendRel(v string)
	PrependRel(v string)
	RemoveRel(index int)
	IsRelIRI(index int) (ok bool)
	GetRelIRI(index int) (v *url.URL)
	AppendRelIRI(v *url.URL)
	PrependRelIRI(v *url.URL)
	RemoveRelIRI(index int)
	TypeLen() (l int)
	GetType(index int) (v interface{})
	AppendType(v interface{})
	PrependType(v interface{})
	RemoveType(index int)
	IsMediaType() (ok bool)
	GetMediaType() (v string)
	SetMediaType(v string)
	IsMediaTypeIRI() (ok bool)
	GetMediaTypeIRI() (v *url.URL)
	SetMediaTypeIRI(v *url.URL)
	NameLen() (l int)
	IsNameString(index int) (ok bool)
	GetNameString(index int) (v string)
	AppendNameString(v string)
	PrependNameString(v string)
	RemoveNameString(index int)
	IsNameLangString(index int) (ok bool)
	GetNameLangString(index int) (v string)
	AppendNameLangString(v string)
	PrependNameLangString(v string)
	RemoveNameLangString(index int)
	IsNameIRI(index int) (ok bool)
	GetNameIRI(index int) (v *url.URL)
	AppendNameIRI(v *url.URL)
	PrependNameIRI(v *url.URL)
	RemoveNameIRI(index int)
	NameMapLanguages() (l []string)
	GetNameMap(l string) (v string)
	SetNameMap(l string, v string)
	SummaryLen() (l int)
	IsSummaryString(index int) (ok bool)
	GetSummaryString(index int) (v string)
	AppendSummaryString(v string)
	PrependSummaryString(v string)
	RemoveSummaryString(index int)
	IsSummaryLangString(index int) (ok bool)
	GetSummaryLangString(index int) (v string)
	AppendSummaryLangString(v string)
	PrependSummaryLangString(v string)
	RemoveSummaryLangString(index int)
	IsSummaryIRI(index int) (ok bool)
	GetSummaryIRI(index int) (v *url.URL)
	AppendSummaryIRI(v *url.URL)
	PrependSummaryIRI(v *url.URL)
	RemoveSummaryIRI(index int)
	SummaryMapLanguages() (l []string)
	GetSummaryMap(l string) (v string)
	SetSummaryMap(l string, v string)
	IsHreflang() (ok bool)
	GetHreflang() (v string)
	SetHreflang(v string)
	IsHreflangIRI() (ok bool)
	GetHreflangIRI() (v *url.URL)
	SetHreflangIRI(v *url.URL)
	IsHeight() (ok bool)
	GetHeight() (v int64)
	SetHeight(v int64)
	IsHeightIRI() (ok bool)
	GetHeightIRI() (v *url.URL)
	SetHeightIRI(v *url.URL)
	IsWidth() (ok bool)
	GetWidth() (v int64)
	SetWidth(v int64)
	IsWidthIRI() (ok bool)
	GetWidthIRI() (v *url.URL)
	SetWidthIRI(v *url.URL)
	PreviewLen() (l int)
	IsPreviewObject(index int) (ok bool)
	GetPreviewObject(index int) (v ObjectType)
	AppendPreviewObject(v ObjectType)
	PrependPreviewObject(v ObjectType)
	RemovePreviewObject(index int)
	IsPreviewLink(index int) (ok bool)
	GetPreviewLink(index int) (v LinkType)
	AppendPreviewLink(v LinkType)
	PrependPreviewLink(v LinkType)
	RemovePreviewLink(index int)
	IsPreviewIRI(index int) (ok bool)
	GetPreviewIRI(index int) (v *url.URL)
	AppendPreviewIRI(v *url.URL)
	PrependPreviewIRI(v *url.URL)
	RemovePreviewIRI(index int)
}

// A specialized Link that represents an @mention.
type Mention struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToMentionIntermediateType
	// The functional 'href' value holds a single type and a single value
	href *url.URL
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'rel' value could have multiple types and values
	rel []*relMentionIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeMentionIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameMentionIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The 'summary' value could have multiple types and values
	summary []*summaryMentionIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The functional 'hreflang' value could have multiple types, but only a single value
	hreflang *hreflangMentionIntermediateType
	// The functional 'height' value could have multiple types, but only a single value
	height *heightMentionIntermediateType
	// The functional 'width' value could have multiple types, but only a single value
	width *widthMentionIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewMentionIntermediateType
}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *Mention) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *Mention) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *Mention) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *Mention) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToMentionIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Mention) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToMentionIntermediateType{&attributedToMentionIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *Mention) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *Mention) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *Mention) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *Mention) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToMentionIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Mention) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToMentionIntermediateType{&attributedToMentionIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *Mention) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *Mention) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *Mention) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *Mention) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToMentionIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Mention) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToMentionIntermediateType{&attributedToMentionIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *Mention) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *Mention) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *Mention) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *Mention) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToMentionIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// HasHref determines whether the call to GetHref is safe
func (t *Mention) HasHref() (ok bool) {
	return t.href != nil

}

// GetHref returns the value for href
func (t *Mention) GetHref() (v *url.URL) {
	return t.href

}

// SetHref sets the value of href
func (t *Mention) SetHref(v *url.URL) {
	t.href = v

}

// HasUnknownHref determines whether the call to GetUnknownHref is safe
func (t *Mention) HasUnknownHref() (ok bool) {
	return t.unknown_ != nil && t.unknown_["href"] != nil

}

// GetUnknownHref returns the unknown value for href
func (t *Mention) GetUnknownHref() (v interface{}) {
	return t.unknown_["href"]

}

// SetUnknownHref sets the unknown value of href
func (t *Mention) SetUnknownHref(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["href"] = i

}

// HasId determines whether the call to GetId is safe
func (t *Mention) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *Mention) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *Mention) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *Mention) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *Mention) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *Mention) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// RelLen determines the number of elements able to be used for the IsRel, GetRel, and RemoveRel functions
func (t *Mention) RelLen() (l int) {
	return len(t.rel)

}

// IsRel determines whether the call to GetRel is safe for the specified index
func (t *Mention) IsRel(index int) (ok bool) {
	return t.rel[index].linkRelation != nil

}

// GetRel returns the value safely if IsRel returned true for the specified index
func (t *Mention) GetRel(index int) (v string) {
	return *t.rel[index].linkRelation

}

// AppendRel adds to the back of rel a string type
func (t *Mention) AppendRel(v string) {
	t.rel = append(t.rel, &relMentionIntermediateType{linkRelation: &v})

}

// PrependRel adds to the front of rel a string type
func (t *Mention) PrependRel(v string) {
	t.rel = append([]*relMentionIntermediateType{&relMentionIntermediateType{linkRelation: &v}}, t.rel...)

}

// RemoveRel deletes the value from the specified index
func (t *Mention) RemoveRel(index int) {
	copy(t.rel[index:], t.rel[index+1:])
	t.rel[len(t.rel)-1] = nil
	t.rel = t.rel[:len(t.rel)-1]

}

// IsRelIRI determines whether the call to GetRelIRI is safe for the specified index
func (t *Mention) IsRelIRI(index int) (ok bool) {
	return t.rel[index].IRI != nil

}

// GetRelIRI returns the value safely if IsRelIRI returned true for the specified index
func (t *Mention) GetRelIRI(index int) (v *url.URL) {
	return t.rel[index].IRI

}

// AppendRelIRI adds to the back of rel a *url.URL type
func (t *Mention) AppendRelIRI(v *url.URL) {
	t.rel = append(t.rel, &relMentionIntermediateType{IRI: v})

}

// PrependRelIRI adds to the front of rel a *url.URL type
func (t *Mention) PrependRelIRI(v *url.URL) {
	t.rel = append([]*relMentionIntermediateType{&relMentionIntermediateType{IRI: v}}, t.rel...)

}

// RemoveRelIRI deletes the value from the specified index
func (t *Mention) RemoveRelIRI(index int) {
	copy(t.rel[index:], t.rel[index+1:])
	t.rel[len(t.rel)-1] = nil
	t.rel = t.rel[:len(t.rel)-1]

}

// HasUnknownRel determines whether the call to GetUnknownRel is safe
func (t *Mention) HasUnknownRel() (ok bool) {
	return t.rel != nil && t.rel[0].unknown_ != nil

}

// GetUnknownRel returns the unknown value for rel
func (t *Mention) GetUnknownRel() (v interface{}) {
	return t.rel[0].unknown_

}

// SetUnknownRel sets the unknown value of rel
func (t *Mention) SetUnknownRel(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &relMentionIntermediateType{}
	tmp.unknown_ = i
	t.rel = append(t.rel, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *Mention) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *Mention) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *Mention) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *Mention) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *Mention) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *Mention) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *Mention) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *Mention) SetMediaType(v string) {
	t.mediaType = &mediaTypeMentionIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *Mention) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *Mention) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *Mention) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeMentionIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *Mention) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *Mention) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *Mention) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeMentionIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *Mention) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *Mention) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *Mention) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *Mention) AppendNameString(v string) {
	t.name = append(t.name, &nameMentionIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Mention) PrependNameString(v string) {
	t.name = append([]*nameMentionIntermediateType{&nameMentionIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *Mention) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *Mention) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *Mention) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *Mention) AppendNameLangString(v string) {
	t.name = append(t.name, &nameMentionIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Mention) PrependNameLangString(v string) {
	t.name = append([]*nameMentionIntermediateType{&nameMentionIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *Mention) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *Mention) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *Mention) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *Mention) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameMentionIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Mention) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameMentionIntermediateType{&nameMentionIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *Mention) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *Mention) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *Mention) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *Mention) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameMentionIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Mention) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Mention) GetNameMap(l string) (v string) {
	if t.nameMap == nil {
		return ""
	}
	ok := false
	v, ok = t.nameMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetNameMap sets the value of the property for the specified language
func (t *Mention) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *Mention) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *Mention) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *Mention) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *Mention) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryMentionIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Mention) PrependSummaryString(v string) {
	t.summary = append([]*summaryMentionIntermediateType{&summaryMentionIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *Mention) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *Mention) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *Mention) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *Mention) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryMentionIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Mention) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryMentionIntermediateType{&summaryMentionIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *Mention) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *Mention) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *Mention) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *Mention) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryMentionIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Mention) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryMentionIntermediateType{&summaryMentionIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *Mention) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *Mention) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *Mention) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *Mention) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryMentionIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Mention) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Mention) GetSummaryMap(l string) (v string) {
	if t.summaryMap == nil {
		return ""
	}
	ok := false
	v, ok = t.summaryMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetSummaryMap sets the value of the property for the specified language
func (t *Mention) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// IsHreflang determines whether the call to GetHreflang is safe
func (t *Mention) IsHreflang() (ok bool) {
	return t.hreflang != nil && t.hreflang.bcp47LanguageTag != nil

}

// GetHreflang returns the value safely if IsHreflang returned true
func (t *Mention) GetHreflang() (v string) {
	return *t.hreflang.bcp47LanguageTag

}

// SetHreflang sets the value of hreflang to be of string type
func (t *Mention) SetHreflang(v string) {
	t.hreflang = &hreflangMentionIntermediateType{bcp47LanguageTag: &v}

}

// IsHreflangIRI determines whether the call to GetHreflangIRI is safe
func (t *Mention) IsHreflangIRI() (ok bool) {
	return t.hreflang != nil && t.hreflang.IRI != nil

}

// GetHreflangIRI returns the value safely if IsHreflangIRI returned true
func (t *Mention) GetHreflangIRI() (v *url.URL) {
	return t.hreflang.IRI

}

// SetHreflangIRI sets the value of hreflang to be of *url.URL type
func (t *Mention) SetHreflangIRI(v *url.URL) {
	t.hreflang = &hreflangMentionIntermediateType{IRI: v}

}

// HasUnknownHreflang determines whether the call to GetUnknownHreflang is safe
func (t *Mention) HasUnknownHreflang() (ok bool) {
	return t.hreflang != nil && t.hreflang.unknown_ != nil

}

// GetUnknownHreflang returns the unknown value for hreflang
func (t *Mention) GetUnknownHreflang() (v interface{}) {
	return t.hreflang.unknown_

}

// SetUnknownHreflang sets the unknown value of hreflang
func (t *Mention) SetUnknownHreflang(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &hreflangMentionIntermediateType{}
	tmp.unknown_ = i
	t.hreflang = tmp

}

// IsHeight determines whether the call to GetHeight is safe
func (t *Mention) IsHeight() (ok bool) {
	return t.height != nil && t.height.nonNegativeInteger != nil

}

// GetHeight returns the value safely if IsHeight returned true
func (t *Mention) GetHeight() (v int64) {
	return *t.height.nonNegativeInteger

}

// SetHeight sets the value of height to be of int64 type
func (t *Mention) SetHeight(v int64) {
	t.height = &heightMentionIntermediateType{nonNegativeInteger: &v}

}

// IsHeightIRI determines whether the call to GetHeightIRI is safe
func (t *Mention) IsHeightIRI() (ok bool) {
	return t.height != nil && t.height.IRI != nil

}

// GetHeightIRI returns the value safely if IsHeightIRI returned true
func (t *Mention) GetHeightIRI() (v *url.URL) {
	return t.height.IRI

}

// SetHeightIRI sets the value of height to be of *url.URL type
func (t *Mention) SetHeightIRI(v *url.URL) {
	t.height = &heightMentionIntermediateType{IRI: v}

}

// HasUnknownHeight determines whether the call to GetUnknownHeight is safe
func (t *Mention) HasUnknownHeight() (ok bool) {
	return t.height != nil && t.height.unknown_ != nil

}

// GetUnknownHeight returns the unknown value for height
func (t *Mention) GetUnknownHeight() (v interface{}) {
	return t.height.unknown_

}

// SetUnknownHeight sets the unknown value of height
func (t *Mention) SetUnknownHeight(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &heightMentionIntermediateType{}
	tmp.unknown_ = i
	t.height = tmp

}

// IsWidth determines whether the call to GetWidth is safe
func (t *Mention) IsWidth() (ok bool) {
	return t.width != nil && t.width.nonNegativeInteger != nil

}

// GetWidth returns the value safely if IsWidth returned true
func (t *Mention) GetWidth() (v int64) {
	return *t.width.nonNegativeInteger

}

// SetWidth sets the value of width to be of int64 type
func (t *Mention) SetWidth(v int64) {
	t.width = &widthMentionIntermediateType{nonNegativeInteger: &v}

}

// IsWidthIRI determines whether the call to GetWidthIRI is safe
func (t *Mention) IsWidthIRI() (ok bool) {
	return t.width != nil && t.width.IRI != nil

}

// GetWidthIRI returns the value safely if IsWidthIRI returned true
func (t *Mention) GetWidthIRI() (v *url.URL) {
	return t.width.IRI

}

// SetWidthIRI sets the value of width to be of *url.URL type
func (t *Mention) SetWidthIRI(v *url.URL) {
	t.width = &widthMentionIntermediateType{IRI: v}

}

// HasUnknownWidth determines whether the call to GetUnknownWidth is safe
func (t *Mention) HasUnknownWidth() (ok bool) {
	return t.width != nil && t.width.unknown_ != nil

}

// GetUnknownWidth returns the unknown value for width
func (t *Mention) GetUnknownWidth() (v interface{}) {
	return t.width.unknown_

}

// SetUnknownWidth sets the unknown value of width
func (t *Mention) SetUnknownWidth(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &widthMentionIntermediateType{}
	tmp.unknown_ = i
	t.width = tmp

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *Mention) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *Mention) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *Mention) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *Mention) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewMentionIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Mention) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewMentionIntermediateType{&previewMentionIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *Mention) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *Mention) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *Mention) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *Mention) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewMentionIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Mention) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewMentionIntermediateType{&previewMentionIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *Mention) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *Mention) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *Mention) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *Mention) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewMentionIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Mention) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewMentionIntermediateType{&previewMentionIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *Mention) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *Mention) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *Mention) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *Mention) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewMentionIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *Mention) AddUnknown(k string, i interface{}) (this *Mention) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *Mention) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *Mention) RemoveUnknown(k string) (this *Mention) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "Mention" if not manually set by the caller
func (t *Mention) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "Mention" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "Mention")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttributedToMentionIntermediateType(t.attributedTo); err == nil && v != nil {
		if len(v) == 1 {
			m["attributedTo"] = v[0]
		} else {
			m["attributedTo"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	if t.href != nil {
		hrefSerializeFunc := func() (interface{}, error) {
			v := t.href
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		hrefResult, err := hrefSerializeFunc()
		if err == nil {
			m["href"] = hrefResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.id != nil {
		idSerializeFunc := func() (interface{}, error) {
			v := t.id
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		idResult, err := idSerializeFunc()
		if err == nil {
			m["id"] = idResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceRelMentionIntermediateType(t.rel); err == nil && v != nil {
		if len(v) == 1 {
			m["rel"] = v[0]
		} else {
			m["rel"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalAnyDefinition
	if t.typeName != nil {
		if len(t.typeName) == 1 {
			m["type"] = t.typeName[0]
		} else {
			m["type"] = t.typeName
		}
	}
	// End generation by generateNonFunctionalAnyDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.mediaType != nil {
		if v, err := serializeMediaTypeMentionIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceNameMentionIntermediateType(t.name); err == nil && v != nil {
		if len(v) == 1 {
			m["name"] = v[0]
		} else {
			m["name"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.nameMap != nil && len(t.nameMap) >= 0 {
		m["nameMap"] = t.nameMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryMentionIntermediateType(t.summary); err == nil && v != nil {
		if len(v) == 1 {
			m["summary"] = v[0]
		} else {
			m["summary"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.summaryMap != nil && len(t.summaryMap) >= 0 {
		m["summaryMap"] = t.summaryMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.hreflang != nil {
		if v, err := serializeHreflangMentionIntermediateType(t.hreflang); err == nil {
			m["hreflang"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.height != nil {
		if v, err := serializeHeightMentionIntermediateType(t.height); err == nil {
			m["height"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.width != nil {
		if v, err := serializeWidthMentionIntermediateType(t.width); err == nil {
			m["width"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSlicePreviewMentionIntermediateType(t.preview); err == nil && v != nil {
		if len(v) == 1 {
			m["preview"] = v[0]
		} else {
			m["preview"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *Mention) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToMentionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToMentionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToMentionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToMentionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToMentionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "href" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.href = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "id" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.id = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "rel" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeRelMentionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.rel = []*relMentionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.rel, err = deserializeSliceRelMentionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &relMentionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.rel = []*relMentionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalAnyDefinition
			if k == "type" {
				if tmpTypeSlice, ok := v.([]interface{}); ok {
					t.typeName = tmpTypeSlice
					handled = true
				} else {
					t.typeName = []interface{}{v}
					handled = true
				}
			}
			// End generation by generateNonFunctionalAnyDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeMentionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameMentionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameMentionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameMentionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameMentionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameMentionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "nameMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.nameMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "summary" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeSummaryMentionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryMentionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryMentionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryMentionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryMentionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "summaryMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.summaryMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "hreflang" {
				t.hreflang, err = deserializeHreflangMentionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "height" {
				t.height, err = deserializeHeightMentionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "width" {
				t.width, err = deserializeWidthMentionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewMentionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewMentionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewMentionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewMentionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewMentionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled && k != "@context" {
			if t.unknown_ == nil {
				t.unknown_ = make(map[string]interface{})
			}
			t.unknown_[k] = unknownValueDeserialize(v)
		}
	}
	return

}

// attributedToMentionIntermediateType will only have one of its values set at most
type attributedToMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attributedTo property
	Object ObjectType
	// Stores possible LinkType type for attributedTo property
	Link LinkType
	// Stores possible *url.URL type for attributedTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attributedToMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attributedToMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// relMentionIntermediateType will only have one of its values set at most
type relMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for rel property
	linkRelation *string
	// Stores possible *url.URL type for rel property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *relMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.linkRelation, err = linkRelationDeserialize(i)
			if err != nil {
				t.linkRelation = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *relMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.linkRelation != nil {
		i = linkRelationSerialize(*t.linkRelation)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// mediaTypeMentionIntermediateType will only have one of its values set at most
type mediaTypeMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.mimeMediaTypeValue, err = mimeMediaTypeValueDeserialize(i)
			if err != nil {
				t.mimeMediaTypeValue = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *mediaTypeMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.mimeMediaTypeValue != nil {
		i = mimeMediaTypeValueSerialize(*t.mimeMediaTypeValue)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// nameMentionIntermediateType will only have one of its values set at most
type nameMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for name property
	stringName *string
	// Stores possible *string type for name property
	langString *string
	// Stores possible *url.URL type for name property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nameMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *nameMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// summaryMentionIntermediateType will only have one of its values set at most
type summaryMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for summary property
	stringName *string
	// Stores possible *string type for summary property
	langString *string
	// Stores possible *url.URL type for summary property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *summaryMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *summaryMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// hreflangMentionIntermediateType will only have one of its values set at most
type hreflangMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for hreflang property
	bcp47LanguageTag *string
	// Stores possible *url.URL type for hreflang property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *hreflangMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.bcp47LanguageTag, err = bcp47LanguageTagDeserialize(i)
			if err != nil {
				t.bcp47LanguageTag = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *hreflangMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.bcp47LanguageTag != nil {
		i = bcp47LanguageTagSerialize(*t.bcp47LanguageTag)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// heightMentionIntermediateType will only have one of its values set at most
type heightMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for height property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for height property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *heightMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *heightMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// widthMentionIntermediateType will only have one of its values set at most
type widthMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for width property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for width property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *widthMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *widthMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// previewMentionIntermediateType will only have one of its values set at most
type previewMentionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for preview property
	Object ObjectType
	// Stores possible LinkType type for preview property
	Link LinkType
	// Stores possible *url.URL type for preview property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *previewMentionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *previewMentionIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// deserializeattributedToMentionIntermediateType will accept a map to create a attributedToMentionIntermediateType
func deserializeAttributedToMentionIntermediateType(in interface{}) (t *attributedToMentionIntermediateType, err error) {
	tmp := &attributedToMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToMentionIntermediateType will accept a slice to create a slice of attributedToMentionIntermediateType
func deserializeSliceAttributedToMentionIntermediateType(in []interface{}) (t []*attributedToMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToMentionIntermediateType will accept a attributedToMentionIntermediateType to create a map
func serializeAttributedToMentionIntermediateType(t *attributedToMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToMentionIntermediateType will accept a slice of attributedToMentionIntermediateType to create a slice result
func serializeSliceAttributedToMentionIntermediateType(s []*attributedToMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerelMentionIntermediateType will accept a map to create a relMentionIntermediateType
func deserializeRelMentionIntermediateType(in interface{}) (t *relMentionIntermediateType, err error) {
	tmp := &relMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice relMentionIntermediateType will accept a slice to create a slice of relMentionIntermediateType
func deserializeSliceRelMentionIntermediateType(in []interface{}) (t []*relMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &relMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerelMentionIntermediateType will accept a relMentionIntermediateType to create a map
func serializeRelMentionIntermediateType(t *relMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerelMentionIntermediateType will accept a slice of relMentionIntermediateType to create a slice result
func serializeSliceRelMentionIntermediateType(s []*relMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeMentionIntermediateType will accept a map to create a mediaTypeMentionIntermediateType
func deserializeMediaTypeMentionIntermediateType(in interface{}) (t *mediaTypeMentionIntermediateType, err error) {
	tmp := &mediaTypeMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeMentionIntermediateType will accept a slice to create a slice of mediaTypeMentionIntermediateType
func deserializeSliceMediaTypeMentionIntermediateType(in []interface{}) (t []*mediaTypeMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeMentionIntermediateType will accept a mediaTypeMentionIntermediateType to create a map
func serializeMediaTypeMentionIntermediateType(t *mediaTypeMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeMentionIntermediateType will accept a slice of mediaTypeMentionIntermediateType to create a slice result
func serializeSliceMediaTypeMentionIntermediateType(s []*mediaTypeMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameMentionIntermediateType will accept a map to create a nameMentionIntermediateType
func deserializeNameMentionIntermediateType(in interface{}) (t *nameMentionIntermediateType, err error) {
	tmp := &nameMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameMentionIntermediateType will accept a slice to create a slice of nameMentionIntermediateType
func deserializeSliceNameMentionIntermediateType(in []interface{}) (t []*nameMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameMentionIntermediateType will accept a nameMentionIntermediateType to create a map
func serializeNameMentionIntermediateType(t *nameMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameMentionIntermediateType will accept a slice of nameMentionIntermediateType to create a slice result
func serializeSliceNameMentionIntermediateType(s []*nameMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryMentionIntermediateType will accept a map to create a summaryMentionIntermediateType
func deserializeSummaryMentionIntermediateType(in interface{}) (t *summaryMentionIntermediateType, err error) {
	tmp := &summaryMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryMentionIntermediateType will accept a slice to create a slice of summaryMentionIntermediateType
func deserializeSliceSummaryMentionIntermediateType(in []interface{}) (t []*summaryMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryMentionIntermediateType will accept a summaryMentionIntermediateType to create a map
func serializeSummaryMentionIntermediateType(t *summaryMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryMentionIntermediateType will accept a slice of summaryMentionIntermediateType to create a slice result
func serializeSliceSummaryMentionIntermediateType(s []*summaryMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializehreflangMentionIntermediateType will accept a map to create a hreflangMentionIntermediateType
func deserializeHreflangMentionIntermediateType(in interface{}) (t *hreflangMentionIntermediateType, err error) {
	tmp := &hreflangMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice hreflangMentionIntermediateType will accept a slice to create a slice of hreflangMentionIntermediateType
func deserializeSliceHreflangMentionIntermediateType(in []interface{}) (t []*hreflangMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &hreflangMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializehreflangMentionIntermediateType will accept a hreflangMentionIntermediateType to create a map
func serializeHreflangMentionIntermediateType(t *hreflangMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicehreflangMentionIntermediateType will accept a slice of hreflangMentionIntermediateType to create a slice result
func serializeSliceHreflangMentionIntermediateType(s []*hreflangMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeheightMentionIntermediateType will accept a map to create a heightMentionIntermediateType
func deserializeHeightMentionIntermediateType(in interface{}) (t *heightMentionIntermediateType, err error) {
	tmp := &heightMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice heightMentionIntermediateType will accept a slice to create a slice of heightMentionIntermediateType
func deserializeSliceHeightMentionIntermediateType(in []interface{}) (t []*heightMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &heightMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeheightMentionIntermediateType will accept a heightMentionIntermediateType to create a map
func serializeHeightMentionIntermediateType(t *heightMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceheightMentionIntermediateType will accept a slice of heightMentionIntermediateType to create a slice result
func serializeSliceHeightMentionIntermediateType(s []*heightMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializewidthMentionIntermediateType will accept a map to create a widthMentionIntermediateType
func deserializeWidthMentionIntermediateType(in interface{}) (t *widthMentionIntermediateType, err error) {
	tmp := &widthMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice widthMentionIntermediateType will accept a slice to create a slice of widthMentionIntermediateType
func deserializeSliceWidthMentionIntermediateType(in []interface{}) (t []*widthMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &widthMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializewidthMentionIntermediateType will accept a widthMentionIntermediateType to create a map
func serializeWidthMentionIntermediateType(t *widthMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicewidthMentionIntermediateType will accept a slice of widthMentionIntermediateType to create a slice result
func serializeSliceWidthMentionIntermediateType(s []*widthMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewMentionIntermediateType will accept a map to create a previewMentionIntermediateType
func deserializePreviewMentionIntermediateType(in interface{}) (t *previewMentionIntermediateType, err error) {
	tmp := &previewMentionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewMentionIntermediateType will accept a slice to create a slice of previewMentionIntermediateType
func deserializeSlicePreviewMentionIntermediateType(in []interface{}) (t []*previewMentionIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewMentionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewMentionIntermediateType will accept a previewMentionIntermediateType to create a map
func serializePreviewMentionIntermediateType(t *previewMentionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewMentionIntermediateType will accept a slice of previewMentionIntermediateType to create a slice result
func serializeSlicePreviewMentionIntermediateType(s []*previewMentionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}
