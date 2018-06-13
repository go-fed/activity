//
package vocab

import (
	"fmt"
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

// A Link is an indirect, qualified reference to a resource identified by a URL. The fundamental model for links is established by [ RFC5988]. Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link. When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the resource identified by the href. Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToLinkIntermediateType
	// The functional 'href' value holds a single type and a single value
	href *url.URL
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'rel' value could have multiple types and values
	rel []*relLinkIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeLinkIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameLinkIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The 'summary' value could have multiple types and values
	summary []*summaryLinkIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The functional 'hreflang' value could have multiple types, but only a single value
	hreflang *hreflangLinkIntermediateType
	// The functional 'height' value could have multiple types, but only a single value
	height *heightLinkIntermediateType
	// The functional 'width' value could have multiple types, but only a single value
	width *widthLinkIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewLinkIntermediateType
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
	t.attributedTo = append(t.attributedTo, &attributedToLinkIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Link) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToLinkIntermediateType{&attributedToLinkIntermediateType{Object: v}}, t.attributedTo...)

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
	t.attributedTo = append(t.attributedTo, &attributedToLinkIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Link) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToLinkIntermediateType{&attributedToLinkIntermediateType{Link: v}}, t.attributedTo...)

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
	t.attributedTo = append(t.attributedTo, &attributedToLinkIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Link) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToLinkIntermediateType{&attributedToLinkIntermediateType{IRI: v}}, t.attributedTo...)

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
	tmp := &attributedToLinkIntermediateType{}
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
	t.rel = append(t.rel, &relLinkIntermediateType{linkRelation: &v})

}

// PrependRel adds to the front of rel a string type
func (t *Link) PrependRel(v string) {
	t.rel = append([]*relLinkIntermediateType{&relLinkIntermediateType{linkRelation: &v}}, t.rel...)

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
	t.rel = append(t.rel, &relLinkIntermediateType{IRI: v})

}

// PrependRelIRI adds to the front of rel a *url.URL type
func (t *Link) PrependRelIRI(v *url.URL) {
	t.rel = append([]*relLinkIntermediateType{&relLinkIntermediateType{IRI: v}}, t.rel...)

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
	tmp := &relLinkIntermediateType{}
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
	t.mediaType = &mediaTypeLinkIntermediateType{mimeMediaTypeValue: &v}

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
	t.mediaType = &mediaTypeLinkIntermediateType{IRI: v}

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
	tmp := &mediaTypeLinkIntermediateType{}
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
	t.name = append(t.name, &nameLinkIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Link) PrependNameString(v string) {
	t.name = append([]*nameLinkIntermediateType{&nameLinkIntermediateType{stringName: &v}}, t.name...)

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
	t.name = append(t.name, &nameLinkIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Link) PrependNameLangString(v string) {
	t.name = append([]*nameLinkIntermediateType{&nameLinkIntermediateType{langString: &v}}, t.name...)

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
	t.name = append(t.name, &nameLinkIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Link) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameLinkIntermediateType{&nameLinkIntermediateType{IRI: v}}, t.name...)

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
	tmp := &nameLinkIntermediateType{}
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
	t.summary = append(t.summary, &summaryLinkIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Link) PrependSummaryString(v string) {
	t.summary = append([]*summaryLinkIntermediateType{&summaryLinkIntermediateType{stringName: &v}}, t.summary...)

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
	t.summary = append(t.summary, &summaryLinkIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Link) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryLinkIntermediateType{&summaryLinkIntermediateType{langString: &v}}, t.summary...)

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
	t.summary = append(t.summary, &summaryLinkIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Link) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryLinkIntermediateType{&summaryLinkIntermediateType{IRI: v}}, t.summary...)

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
	tmp := &summaryLinkIntermediateType{}
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
	t.hreflang = &hreflangLinkIntermediateType{bcp47LanguageTag: &v}

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
	t.hreflang = &hreflangLinkIntermediateType{IRI: v}

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
	tmp := &hreflangLinkIntermediateType{}
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
	t.height = &heightLinkIntermediateType{nonNegativeInteger: &v}

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
	t.height = &heightLinkIntermediateType{IRI: v}

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
	tmp := &heightLinkIntermediateType{}
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
	t.width = &widthLinkIntermediateType{nonNegativeInteger: &v}

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
	t.width = &widthLinkIntermediateType{IRI: v}

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
	tmp := &widthLinkIntermediateType{}
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
	t.preview = append(t.preview, &previewLinkIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Link) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewLinkIntermediateType{&previewLinkIntermediateType{Object: v}}, t.preview...)

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
	t.preview = append(t.preview, &previewLinkIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Link) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewLinkIntermediateType{&previewLinkIntermediateType{Link: v}}, t.preview...)

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
	t.preview = append(t.preview, &previewLinkIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Link) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewLinkIntermediateType{&previewLinkIntermediateType{IRI: v}}, t.preview...)

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
	tmp := &previewLinkIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *Link) AddUnknown(k string, i interface{}) (this *Link) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
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
	if v, err := serializeSliceAttributedToLinkIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceRelLinkIntermediateType(t.rel); err == nil && v != nil {
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
		if v, err := serializeMediaTypeLinkIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceNameLinkIntermediateType(t.name); err == nil && v != nil {
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
	if v, err := serializeSliceSummaryLinkIntermediateType(t.summary); err == nil && v != nil {
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
		if v, err := serializeHreflangLinkIntermediateType(t.hreflang); err == nil {
			m["hreflang"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.height != nil {
		if v, err := serializeHeightLinkIntermediateType(t.height); err == nil {
			m["height"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.width != nil {
		if v, err := serializeWidthLinkIntermediateType(t.width); err == nil {
			m["width"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSlicePreviewLinkIntermediateType(t.preview); err == nil && v != nil {
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
					tmp, err := deserializeAttributedToLinkIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToLinkIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToLinkIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToLinkIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToLinkIntermediateType{tmp}
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
					tmp, err := deserializeRelLinkIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.rel = []*relLinkIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.rel, err = deserializeSliceRelLinkIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &relLinkIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.rel = []*relLinkIntermediateType{tmp}
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
				t.mediaType, err = deserializeMediaTypeLinkIntermediateType(v)
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
					tmp, err := deserializeNameLinkIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameLinkIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameLinkIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameLinkIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameLinkIntermediateType{tmp}
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
					tmp, err := deserializeSummaryLinkIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryLinkIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryLinkIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryLinkIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryLinkIntermediateType{tmp}
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
				t.hreflang, err = deserializeHreflangLinkIntermediateType(v)
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
				t.height, err = deserializeHeightLinkIntermediateType(v)
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
				t.width, err = deserializeWidthLinkIntermediateType(v)
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
					tmp, err := deserializePreviewLinkIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewLinkIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewLinkIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewLinkIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewLinkIntermediateType{tmp}
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

// attributedToLinkIntermediateType will only have one of its values set at most
type attributedToLinkIntermediateType struct {
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
func (t *attributedToLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attributedToLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// relLinkIntermediateType will only have one of its values set at most
type relLinkIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for rel property
	linkRelation *string
	// Stores possible *url.URL type for rel property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *relLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *relLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// mediaTypeLinkIntermediateType will only have one of its values set at most
type mediaTypeLinkIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *mediaTypeLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// nameLinkIntermediateType will only have one of its values set at most
type nameLinkIntermediateType struct {
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
func (t *nameLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *nameLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// summaryLinkIntermediateType will only have one of its values set at most
type summaryLinkIntermediateType struct {
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
func (t *summaryLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *summaryLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// hreflangLinkIntermediateType will only have one of its values set at most
type hreflangLinkIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for hreflang property
	bcp47LanguageTag *string
	// Stores possible *url.URL type for hreflang property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *hreflangLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *hreflangLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// heightLinkIntermediateType will only have one of its values set at most
type heightLinkIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for height property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for height property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *heightLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *heightLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// widthLinkIntermediateType will only have one of its values set at most
type widthLinkIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for width property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for width property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *widthLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *widthLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// previewLinkIntermediateType will only have one of its values set at most
type previewLinkIntermediateType struct {
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
func (t *previewLinkIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *previewLinkIntermediateType) Serialize() (i interface{}, err error) {
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

// deserializeattributedToLinkIntermediateType will accept a map to create a attributedToLinkIntermediateType
func deserializeAttributedToLinkIntermediateType(in interface{}) (t *attributedToLinkIntermediateType, err error) {
	tmp := &attributedToLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToLinkIntermediateType will accept a slice to create a slice of attributedToLinkIntermediateType
func deserializeSliceAttributedToLinkIntermediateType(in []interface{}) (t []*attributedToLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToLinkIntermediateType will accept a attributedToLinkIntermediateType to create a map
func serializeAttributedToLinkIntermediateType(t *attributedToLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToLinkIntermediateType will accept a slice of attributedToLinkIntermediateType to create a slice result
func serializeSliceAttributedToLinkIntermediateType(s []*attributedToLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerelLinkIntermediateType will accept a map to create a relLinkIntermediateType
func deserializeRelLinkIntermediateType(in interface{}) (t *relLinkIntermediateType, err error) {
	tmp := &relLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice relLinkIntermediateType will accept a slice to create a slice of relLinkIntermediateType
func deserializeSliceRelLinkIntermediateType(in []interface{}) (t []*relLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &relLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerelLinkIntermediateType will accept a relLinkIntermediateType to create a map
func serializeRelLinkIntermediateType(t *relLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerelLinkIntermediateType will accept a slice of relLinkIntermediateType to create a slice result
func serializeSliceRelLinkIntermediateType(s []*relLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeLinkIntermediateType will accept a map to create a mediaTypeLinkIntermediateType
func deserializeMediaTypeLinkIntermediateType(in interface{}) (t *mediaTypeLinkIntermediateType, err error) {
	tmp := &mediaTypeLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeLinkIntermediateType will accept a slice to create a slice of mediaTypeLinkIntermediateType
func deserializeSliceMediaTypeLinkIntermediateType(in []interface{}) (t []*mediaTypeLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeLinkIntermediateType will accept a mediaTypeLinkIntermediateType to create a map
func serializeMediaTypeLinkIntermediateType(t *mediaTypeLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeLinkIntermediateType will accept a slice of mediaTypeLinkIntermediateType to create a slice result
func serializeSliceMediaTypeLinkIntermediateType(s []*mediaTypeLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameLinkIntermediateType will accept a map to create a nameLinkIntermediateType
func deserializeNameLinkIntermediateType(in interface{}) (t *nameLinkIntermediateType, err error) {
	tmp := &nameLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameLinkIntermediateType will accept a slice to create a slice of nameLinkIntermediateType
func deserializeSliceNameLinkIntermediateType(in []interface{}) (t []*nameLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameLinkIntermediateType will accept a nameLinkIntermediateType to create a map
func serializeNameLinkIntermediateType(t *nameLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameLinkIntermediateType will accept a slice of nameLinkIntermediateType to create a slice result
func serializeSliceNameLinkIntermediateType(s []*nameLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryLinkIntermediateType will accept a map to create a summaryLinkIntermediateType
func deserializeSummaryLinkIntermediateType(in interface{}) (t *summaryLinkIntermediateType, err error) {
	tmp := &summaryLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryLinkIntermediateType will accept a slice to create a slice of summaryLinkIntermediateType
func deserializeSliceSummaryLinkIntermediateType(in []interface{}) (t []*summaryLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryLinkIntermediateType will accept a summaryLinkIntermediateType to create a map
func serializeSummaryLinkIntermediateType(t *summaryLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryLinkIntermediateType will accept a slice of summaryLinkIntermediateType to create a slice result
func serializeSliceSummaryLinkIntermediateType(s []*summaryLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializehreflangLinkIntermediateType will accept a map to create a hreflangLinkIntermediateType
func deserializeHreflangLinkIntermediateType(in interface{}) (t *hreflangLinkIntermediateType, err error) {
	tmp := &hreflangLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice hreflangLinkIntermediateType will accept a slice to create a slice of hreflangLinkIntermediateType
func deserializeSliceHreflangLinkIntermediateType(in []interface{}) (t []*hreflangLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &hreflangLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializehreflangLinkIntermediateType will accept a hreflangLinkIntermediateType to create a map
func serializeHreflangLinkIntermediateType(t *hreflangLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicehreflangLinkIntermediateType will accept a slice of hreflangLinkIntermediateType to create a slice result
func serializeSliceHreflangLinkIntermediateType(s []*hreflangLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeheightLinkIntermediateType will accept a map to create a heightLinkIntermediateType
func deserializeHeightLinkIntermediateType(in interface{}) (t *heightLinkIntermediateType, err error) {
	tmp := &heightLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice heightLinkIntermediateType will accept a slice to create a slice of heightLinkIntermediateType
func deserializeSliceHeightLinkIntermediateType(in []interface{}) (t []*heightLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &heightLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeheightLinkIntermediateType will accept a heightLinkIntermediateType to create a map
func serializeHeightLinkIntermediateType(t *heightLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceheightLinkIntermediateType will accept a slice of heightLinkIntermediateType to create a slice result
func serializeSliceHeightLinkIntermediateType(s []*heightLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializewidthLinkIntermediateType will accept a map to create a widthLinkIntermediateType
func deserializeWidthLinkIntermediateType(in interface{}) (t *widthLinkIntermediateType, err error) {
	tmp := &widthLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice widthLinkIntermediateType will accept a slice to create a slice of widthLinkIntermediateType
func deserializeSliceWidthLinkIntermediateType(in []interface{}) (t []*widthLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &widthLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializewidthLinkIntermediateType will accept a widthLinkIntermediateType to create a map
func serializeWidthLinkIntermediateType(t *widthLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicewidthLinkIntermediateType will accept a slice of widthLinkIntermediateType to create a slice result
func serializeSliceWidthLinkIntermediateType(s []*widthLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewLinkIntermediateType will accept a map to create a previewLinkIntermediateType
func deserializePreviewLinkIntermediateType(in interface{}) (t *previewLinkIntermediateType, err error) {
	tmp := &previewLinkIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewLinkIntermediateType will accept a slice to create a slice of previewLinkIntermediateType
func deserializeSlicePreviewLinkIntermediateType(in []interface{}) (t []*previewLinkIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewLinkIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewLinkIntermediateType will accept a previewLinkIntermediateType to create a map
func serializePreviewLinkIntermediateType(t *previewLinkIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewLinkIntermediateType will accept a slice of previewLinkIntermediateType to create a slice result
func serializeSlicePreviewLinkIntermediateType(s []*previewLinkIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}
