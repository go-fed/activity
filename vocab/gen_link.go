//
package vocab

import (
	"net/url"
)

// LinkType is an interface for accepting types that extend from 'Link'.
type LinkType interface {
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
	HasUnknownAttributedTo() (ok bool)
	GetUnknownAttributedTo() (v interface{})
	SetUnknownAttributedTo(i interface{})
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
	HasUnknownRel() (ok bool)
	GetUnknownRel() (v interface{})
	SetUnknownRel(i interface{})
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
	HasUnknownMediaType() (ok bool)
	GetUnknownMediaType() (v interface{})
	SetUnknownMediaType(i interface{})
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
	HasUnknownName() (ok bool)
	GetUnknownName() (v interface{})
	SetUnknownName(i interface{})
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
	HasUnknownSummary() (ok bool)
	GetUnknownSummary() (v interface{})
	SetUnknownSummary(i interface{})
	SummaryMapLanguages() (l []string)
	GetSummaryMap(l string) (v string)
	SetSummaryMap(l string, v string)
	IsHreflang() (ok bool)
	GetHreflang() (v string)
	SetHreflang(v string)
	IsHreflangIRI() (ok bool)
	GetHreflangIRI() (v *url.URL)
	SetHreflangIRI(v *url.URL)
	HasUnknownHreflang() (ok bool)
	GetUnknownHreflang() (v interface{})
	SetUnknownHreflang(i interface{})
	IsHeight() (ok bool)
	GetHeight() (v int64)
	SetHeight(v int64)
	IsHeightIRI() (ok bool)
	GetHeightIRI() (v *url.URL)
	SetHeightIRI(v *url.URL)
	HasUnknownHeight() (ok bool)
	GetUnknownHeight() (v interface{})
	SetUnknownHeight(i interface{})
	IsWidth() (ok bool)
	GetWidth() (v int64)
	SetWidth(v int64)
	IsWidthIRI() (ok bool)
	GetWidthIRI() (v *url.URL)
	SetWidthIRI(v *url.URL)
	HasUnknownWidth() (ok bool)
	GetUnknownWidth() (v interface{})
	SetUnknownWidth(i interface{})
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
	HasUnknownPreview() (ok bool)
	GetUnknownPreview() (v interface{})
	SetUnknownPreview(i interface{})
}

// A Link is an indirect, qualified reference to a resource identified by a URL. The fundamental model for links is established by [ RFC5988]. Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link. When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the resource identified by the href. Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToIntermediateType
	// The functional 'href' value holds a single type and a single value
	href *url.URL
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'rel' value could have multiple types and values
	rel []*relIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The 'summary' value could have multiple types and values
	summary []*summaryIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The functional 'hreflang' value could have multiple types, but only a single value
	hreflang *hreflangIntermediateType
	// The functional 'height' value could have multiple types, but only a single value
	height *heightIntermediateType
	// The functional 'width' value could have multiple types, but only a single value
	width *widthIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewIntermediateType
}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *Link) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *Link) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *Link) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *Link) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Link) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *Link) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *Link) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *Link) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *Link) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Link) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *Link) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *Link) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *Link) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *Link) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Link) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *Link) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *Link) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *Link) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *Link) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// HasHref determines whether the call to GetHref is safe
func (t *Link) HasHref() (ok bool) {
	return t.href != nil

}

// GetHref returns the value for href
func (t *Link) GetHref() (v *url.URL) {
	return t.href

}

// SetHref sets the value of href
func (t *Link) SetHref(v *url.URL) {
	t.href = v

}

// HasUnknownHref determines whether the call to GetUnknownHref is safe
func (t *Link) HasUnknownHref() (ok bool) {
	return t.unknown_ != nil && t.unknown_["href"] != nil

}

// GetUnknownHref returns the unknown value for href
func (t *Link) GetUnknownHref() (v interface{}) {
	return t.unknown_["href"]

}

// SetUnknownHref sets the unknown value of href
func (t *Link) SetUnknownHref(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["href"] = i

}

// HasId determines whether the call to GetId is safe
func (t *Link) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *Link) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *Link) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *Link) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *Link) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *Link) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// RelLen determines the number of elements able to be used for the IsRel, GetRel, and RemoveRel functions
func (t *Link) RelLen() (l int) {
	return len(t.rel)

}

// IsRel determines whether the call to GetRel is safe for the specified index
func (t *Link) IsRel(index int) (ok bool) {
	return t.rel[index].linkRelation != nil

}

// GetRel returns the value safely if IsRel returned true for the specified index
func (t *Link) GetRel(index int) (v string) {
	return *t.rel[index].linkRelation

}

// AppendRel adds to the back of rel a string type
func (t *Link) AppendRel(v string) {
	t.rel = append(t.rel, &relIntermediateType{linkRelation: &v})

}

// PrependRel adds to the front of rel a string type
func (t *Link) PrependRel(v string) {
	t.rel = append([]*relIntermediateType{&relIntermediateType{linkRelation: &v}}, t.rel...)

}

// RemoveRel deletes the value from the specified index
func (t *Link) RemoveRel(index int) {
	copy(t.rel[index:], t.rel[index+1:])
	t.rel[len(t.rel)-1] = nil
	t.rel = t.rel[:len(t.rel)-1]

}

// IsRelIRI determines whether the call to GetRelIRI is safe for the specified index
func (t *Link) IsRelIRI(index int) (ok bool) {
	return t.rel[index].IRI != nil

}

// GetRelIRI returns the value safely if IsRelIRI returned true for the specified index
func (t *Link) GetRelIRI(index int) (v *url.URL) {
	return t.rel[index].IRI

}

// AppendRelIRI adds to the back of rel a *url.URL type
func (t *Link) AppendRelIRI(v *url.URL) {
	t.rel = append(t.rel, &relIntermediateType{IRI: v})

}

// PrependRelIRI adds to the front of rel a *url.URL type
func (t *Link) PrependRelIRI(v *url.URL) {
	t.rel = append([]*relIntermediateType{&relIntermediateType{IRI: v}}, t.rel...)

}

// RemoveRelIRI deletes the value from the specified index
func (t *Link) RemoveRelIRI(index int) {
	copy(t.rel[index:], t.rel[index+1:])
	t.rel[len(t.rel)-1] = nil
	t.rel = t.rel[:len(t.rel)-1]

}

// HasUnknownRel determines whether the call to GetUnknownRel is safe
func (t *Link) HasUnknownRel() (ok bool) {
	return t.rel != nil && t.rel[0].unknown_ != nil

}

// GetUnknownRel returns the unknown value for rel
func (t *Link) GetUnknownRel() (v interface{}) {
	return t.rel[0].unknown_

}

// SetUnknownRel sets the unknown value of rel
func (t *Link) SetUnknownRel(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &relIntermediateType{}
	tmp.unknown_ = i
	t.rel = append(t.rel, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *Link) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *Link) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *Link) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *Link) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *Link) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *Link) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *Link) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *Link) SetMediaType(v string) {
	t.mediaType = &mediaTypeIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *Link) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *Link) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *Link) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *Link) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *Link) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *Link) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *Link) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *Link) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *Link) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *Link) AppendNameString(v string) {
	t.name = append(t.name, &nameIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Link) PrependNameString(v string) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *Link) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *Link) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *Link) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *Link) AppendNameLangString(v string) {
	t.name = append(t.name, &nameIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Link) PrependNameLangString(v string) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *Link) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *Link) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *Link) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *Link) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Link) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *Link) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *Link) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *Link) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *Link) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Link) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Link) GetNameMap(l string) (v string) {
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
func (t *Link) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *Link) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *Link) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *Link) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *Link) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Link) PrependSummaryString(v string) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *Link) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *Link) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *Link) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *Link) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Link) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *Link) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *Link) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *Link) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *Link) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Link) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *Link) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *Link) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *Link) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *Link) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Link) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Link) GetSummaryMap(l string) (v string) {
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
func (t *Link) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// IsHreflang determines whether the call to GetHreflang is safe
func (t *Link) IsHreflang() (ok bool) {
	return t.hreflang != nil && t.hreflang.bcp47LanguageTag != nil

}

// GetHreflang returns the value safely if IsHreflang returned true
func (t *Link) GetHreflang() (v string) {
	return *t.hreflang.bcp47LanguageTag

}

// SetHreflang sets the value of hreflang to be of string type
func (t *Link) SetHreflang(v string) {
	t.hreflang = &hreflangIntermediateType{bcp47LanguageTag: &v}

}

// IsHreflangIRI determines whether the call to GetHreflangIRI is safe
func (t *Link) IsHreflangIRI() (ok bool) {
	return t.hreflang != nil && t.hreflang.IRI != nil

}

// GetHreflangIRI returns the value safely if IsHreflangIRI returned true
func (t *Link) GetHreflangIRI() (v *url.URL) {
	return t.hreflang.IRI

}

// SetHreflangIRI sets the value of hreflang to be of *url.URL type
func (t *Link) SetHreflangIRI(v *url.URL) {
	t.hreflang = &hreflangIntermediateType{IRI: v}

}

// HasUnknownHreflang determines whether the call to GetUnknownHreflang is safe
func (t *Link) HasUnknownHreflang() (ok bool) {
	return t.hreflang != nil && t.hreflang.unknown_ != nil

}

// GetUnknownHreflang returns the unknown value for hreflang
func (t *Link) GetUnknownHreflang() (v interface{}) {
	return t.hreflang.unknown_

}

// SetUnknownHreflang sets the unknown value of hreflang
func (t *Link) SetUnknownHreflang(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &hreflangIntermediateType{}
	tmp.unknown_ = i
	t.hreflang = tmp

}

// IsHeight determines whether the call to GetHeight is safe
func (t *Link) IsHeight() (ok bool) {
	return t.height != nil && t.height.nonNegativeInteger != nil

}

// GetHeight returns the value safely if IsHeight returned true
func (t *Link) GetHeight() (v int64) {
	return *t.height.nonNegativeInteger

}

// SetHeight sets the value of height to be of int64 type
func (t *Link) SetHeight(v int64) {
	t.height = &heightIntermediateType{nonNegativeInteger: &v}

}

// IsHeightIRI determines whether the call to GetHeightIRI is safe
func (t *Link) IsHeightIRI() (ok bool) {
	return t.height != nil && t.height.IRI != nil

}

// GetHeightIRI returns the value safely if IsHeightIRI returned true
func (t *Link) GetHeightIRI() (v *url.URL) {
	return t.height.IRI

}

// SetHeightIRI sets the value of height to be of *url.URL type
func (t *Link) SetHeightIRI(v *url.URL) {
	t.height = &heightIntermediateType{IRI: v}

}

// HasUnknownHeight determines whether the call to GetUnknownHeight is safe
func (t *Link) HasUnknownHeight() (ok bool) {
	return t.height != nil && t.height.unknown_ != nil

}

// GetUnknownHeight returns the unknown value for height
func (t *Link) GetUnknownHeight() (v interface{}) {
	return t.height.unknown_

}

// SetUnknownHeight sets the unknown value of height
func (t *Link) SetUnknownHeight(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &heightIntermediateType{}
	tmp.unknown_ = i
	t.height = tmp

}

// IsWidth determines whether the call to GetWidth is safe
func (t *Link) IsWidth() (ok bool) {
	return t.width != nil && t.width.nonNegativeInteger != nil

}

// GetWidth returns the value safely if IsWidth returned true
func (t *Link) GetWidth() (v int64) {
	return *t.width.nonNegativeInteger

}

// SetWidth sets the value of width to be of int64 type
func (t *Link) SetWidth(v int64) {
	t.width = &widthIntermediateType{nonNegativeInteger: &v}

}

// IsWidthIRI determines whether the call to GetWidthIRI is safe
func (t *Link) IsWidthIRI() (ok bool) {
	return t.width != nil && t.width.IRI != nil

}

// GetWidthIRI returns the value safely if IsWidthIRI returned true
func (t *Link) GetWidthIRI() (v *url.URL) {
	return t.width.IRI

}

// SetWidthIRI sets the value of width to be of *url.URL type
func (t *Link) SetWidthIRI(v *url.URL) {
	t.width = &widthIntermediateType{IRI: v}

}

// HasUnknownWidth determines whether the call to GetUnknownWidth is safe
func (t *Link) HasUnknownWidth() (ok bool) {
	return t.width != nil && t.width.unknown_ != nil

}

// GetUnknownWidth returns the unknown value for width
func (t *Link) GetUnknownWidth() (v interface{}) {
	return t.width.unknown_

}

// SetUnknownWidth sets the unknown value of width
func (t *Link) SetUnknownWidth(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &widthIntermediateType{}
	tmp.unknown_ = i
	t.width = tmp

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *Link) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *Link) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *Link) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *Link) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Link) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *Link) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *Link) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *Link) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *Link) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Link) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *Link) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *Link) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *Link) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *Link) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Link) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *Link) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *Link) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *Link) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *Link) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// AddUnknown adds an unknown property to this object with the specified key
func (t *Link) AddUnknown(k string, i interface{}) (this *Link) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown property with the specified key
func (t *Link) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *Link) RemoveUnknown(k string) (this *Link) {
	delete(t.unknown_, k)
	return t

}

// GetUnknown fetches an unknown property from this object with the specified key. Note that this will panic if HasUnknown would return false.
func (t *Link) GetUnknown(k string) (i interface{}) {
	return t.unknown_[k]

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "Link" if not manually set by the caller
func (t *Link) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "Link" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "Link")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttributedToIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceRelIntermediateType(t.rel); err == nil && v != nil {
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
		if v, err := serializeMediaTypeIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceNameIntermediateType(t.name); err == nil && v != nil {
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
	if v, err := serializeSliceSummaryIntermediateType(t.summary); err == nil && v != nil {
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
		if v, err := serializeHreflangIntermediateType(t.hreflang); err == nil {
			m["hreflang"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.height != nil {
		if v, err := serializeHeightIntermediateType(t.height); err == nil {
			m["height"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.width != nil {
		if v, err := serializeWidthIntermediateType(t.width); err == nil {
			m["width"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSlicePreviewIntermediateType(t.preview); err == nil && v != nil {
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
func (t *Link) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntermediateType{tmp}
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
					tmp, err := deserializeRelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.rel = []*relIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.rel, err = deserializeSliceRelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &relIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.rel = []*relIntermediateType{tmp}
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
				t.mediaType, err = deserializeMediaTypeIntermediateType(v)
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
					tmp, err := deserializeNameIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameIntermediateType{tmp}
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
					tmp, err := deserializeSummaryIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntermediateType{tmp}
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
				t.hreflang, err = deserializeHreflangIntermediateType(v)
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
				t.height, err = deserializeHeightIntermediateType(v)
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
				t.width, err = deserializeWidthIntermediateType(v)
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
					tmp, err := deserializePreviewIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewIntermediateType{tmp}
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
