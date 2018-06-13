//
package streams

import (
	"github.com/go-fed/activity/vocab"
	"net/url"
)

// A Link is an indirect, qualified reference to a resource identified by a URL. The fundamental model for links is established by [ RFC5988]. Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link. When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the resource identified by the href. Properties of the Link are properties of the reference as opposed to properties of the resource. This is a convenience wrapper of a type with the same name in the vocab package. Accessing it with the Raw function allows direct manipulaton of the object, and does not provide the same integrity guarantees as this package.
type Link struct {
	// The raw type from the vocab package
	raw *vocab.Link
}

// Raw returns the vocab type for manaual manipulation. Note that manipulating the underlying type to be in an inconsistent state may cause this convenience type's methods to later fail.
func (t *Link) Raw() (n *vocab.Link) {
	return t.raw

}

// Serialize turns this object into a map[string]interface{}.
func (t *Link) Serialize() (m map[string]interface{}, err error) {
	return t.raw.Serialize()

}

// LenAttributedTo returns the number of values this property contains. Each index be used with HasAttributedTo to determine if GetAttributedTo is safe to call or if raw handling would be needed.
func (t *Link) LenAttributedTo() (idx int) {
	return t.raw.AttributedToLen()

}

// GetAttributedTo attempts to get this 'attributedTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetAttributedTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsAttributedToIRI(idx) {
		k = t.raw.GetAttributedToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAttributedToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsAttributedToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendAttributedTo appends the value for property 'attributedTo'.
func (t *Link) AppendAttributedTo(k *url.URL) {
	t.raw.AppendAttributedToIRI(k)

}

// PrependAttributedTo prepends the value for property 'attributedTo'.
func (t *Link) PrependAttributedTo(k *url.URL) {
	t.raw.PrependAttributedToIRI(k)

}

// RemoveAttributedTo deletes the value from the specified index for property 'attributedTo'.
func (t *Link) RemoveAttributedTo(idx int) {
	t.raw.RemoveAttributedToIRI(idx)

}

// HasAttributedTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasAttributedTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAttributedToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttributedToLink(idx) {
		p = RawPresence
	} else if t.raw.IsAttributedToIRI(idx) {
		p = RawPresence
	}
	return

}

// GetHref attempts to get this 'href' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetHref() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasHref() {
		k = t.raw.GetHref()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasHref returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasHref() (p Presence) {
	p = NoPresence
	if t.raw.HasHref() {
		p = ConvenientPresence
	}
	return

}

// SetHref sets the value for property 'href'.
func (t *Link) SetHref(k *url.URL) {
	t.raw.SetHref(k)

}

// GetId attempts to get this 'id' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetId() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasId() {
		k = t.raw.GetId()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasId returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasId() (p Presence) {
	p = NoPresence
	if t.raw.HasId() {
		p = ConvenientPresence
	}
	return

}

// SetId sets the value for property 'id'.
func (t *Link) SetId(k *url.URL) {
	t.raw.SetId(k)

}

// LenRel returns the number of values this property contains. Each index be used with HasRel to determine if GetRel is safe to call or if raw handling would be needed.
func (t *Link) LenRel() (idx int) {
	return t.raw.RelLen()

}

// GetRel attempts to get this 'rel' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetRel(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsRel(idx) {
		k = t.raw.GetRel(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsRelIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendRel appends the value for property 'rel'.
func (t *Link) AppendRel(k string) {
	t.raw.AppendRel(k)

}

// PrependRel prepends the value for property 'rel'.
func (t *Link) PrependRel(k string) {
	t.raw.PrependRel(k)

}

// RemoveRel deletes the value from the specified index for property 'rel'.
func (t *Link) RemoveRel(idx int) {
	t.raw.RemoveRel(idx)

}

// HasRel returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasRel(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsRel(idx) {
		p = ConvenientPresence
	} else if t.raw.IsRelIRI(idx) {
		p = RawPresence
	}
	return

}

// LenType returns the number of values this property contains. Each index be used with HasType to determine if GetType is safe to call or if raw handling would be needed.
func (t *Link) LenType() (idx int) {
	return t.raw.TypeLen()

}

// GetType attempts to get this 'type' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetType(idx int) (r Resolution, s string) {
	r = Unresolved
	if tmp := t.raw.GetType(idx); tmp != nil {
		ok := false
		if s, ok = tmp.(string); ok {
			r = Resolved
		} else {
			r = RawResolutionNeeded
		}
	}
	return

}

// AppendType appends the value for property 'type'.
func (t *Link) AppendType(i interface{}) {
	t.raw.AppendType(i)

}

// PrependType prepends the value for property 'type'.
func (t *Link) PrependType(i interface{}) {
	t.raw.PrependType(i)

}

// RemoveType deletes the value from the specified index for property 'type'.
func (t *Link) RemoveType(idx int) {
	t.raw.RemoveType(idx)

}

// GetMediaType attempts to get this 'mediaType' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetMediaType() (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsMediaType() {
		k = t.raw.GetMediaType()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsMediaTypeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasMediaType returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasMediaType() (p Presence) {
	p = NoPresence
	if t.raw.IsMediaType() {
		p = ConvenientPresence
	} else if t.raw.IsMediaTypeIRI() {
		p = RawPresence
	}
	return

}

// SetMediaType sets the value for property 'mediaType'.
func (t *Link) SetMediaType(k string) {
	t.raw.SetMediaType(k)

}

// LenName returns the number of values this property contains. Each index be used with HasName to determine if GetName is safe to call or if raw handling would be needed.
func (t *Link) LenName() (idx int) {
	return t.raw.NameLen()

}

// GetName attempts to get this 'name' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetName(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsNameString(idx) {
		k = t.raw.GetNameString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsNameLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsNameIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendName appends the value for property 'name'.
func (t *Link) AppendName(k string) {
	t.raw.AppendNameString(k)

}

// PrependName prepends the value for property 'name'.
func (t *Link) PrependName(k string) {
	t.raw.PrependNameString(k)

}

// RemoveName deletes the value from the specified index for property 'name'.
func (t *Link) RemoveName(idx int) {
	t.raw.RemoveNameString(idx)

}

// HasName returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasName(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsNameString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsNameLangString(idx) {
		p = RawPresence
	} else if t.raw.IsNameIRI(idx) {
		p = RawPresence
	}
	return

}

// NameLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Link) NameLanguages() (l []string) {
	return t.raw.NameMapLanguages()

}

// GetNameMap retrieves the value of 'name' for the specified language, or an empty string if it does not exist
func (t *Link) GetNameForLanguage(l string) (v string) {
	return t.raw.GetNameMap(l)

}

// SetNameForLanguage sets the value of 'name' for the specified language
func (t *Link) SetNameForLanguage(l string, v string) {
	t.raw.SetNameMap(l, v)

}

// LenSummary returns the number of values this property contains. Each index be used with HasSummary to determine if GetSummary is safe to call or if raw handling would be needed.
func (t *Link) LenSummary() (idx int) {
	return t.raw.SummaryLen()

}

// GetSummary attempts to get this 'summary' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetSummary(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsSummaryString(idx) {
		k = t.raw.GetSummaryString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsSummaryLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsSummaryIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendSummary appends the value for property 'summary'.
func (t *Link) AppendSummary(k string) {
	t.raw.AppendSummaryString(k)

}

// PrependSummary prepends the value for property 'summary'.
func (t *Link) PrependSummary(k string) {
	t.raw.PrependSummaryString(k)

}

// RemoveSummary deletes the value from the specified index for property 'summary'.
func (t *Link) RemoveSummary(idx int) {
	t.raw.RemoveSummaryString(idx)

}

// HasSummary returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasSummary(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsSummaryString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsSummaryLangString(idx) {
		p = RawPresence
	} else if t.raw.IsSummaryIRI(idx) {
		p = RawPresence
	}
	return

}

// SummaryLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Link) SummaryLanguages() (l []string) {
	return t.raw.SummaryMapLanguages()

}

// GetSummaryMap retrieves the value of 'summary' for the specified language, or an empty string if it does not exist
func (t *Link) GetSummaryForLanguage(l string) (v string) {
	return t.raw.GetSummaryMap(l)

}

// SetSummaryForLanguage sets the value of 'summary' for the specified language
func (t *Link) SetSummaryForLanguage(l string, v string) {
	t.raw.SetSummaryMap(l, v)

}

// GetHreflang attempts to get this 'hreflang' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetHreflang() (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsHreflang() {
		k = t.raw.GetHreflang()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsHreflangIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasHreflang returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasHreflang() (p Presence) {
	p = NoPresence
	if t.raw.IsHreflang() {
		p = ConvenientPresence
	} else if t.raw.IsHreflangIRI() {
		p = RawPresence
	}
	return

}

// SetHreflang sets the value for property 'hreflang'.
func (t *Link) SetHreflang(k string) {
	t.raw.SetHreflang(k)

}

// GetHeight attempts to get this 'height' property as a int64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetHeight() (r Resolution, k int64) {
	r = Unresolved
	handled := false
	if t.raw.IsHeight() {
		k = t.raw.GetHeight()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsHeightIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasHeight returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasHeight() (p Presence) {
	p = NoPresence
	if t.raw.IsHeight() {
		p = ConvenientPresence
	} else if t.raw.IsHeightIRI() {
		p = RawPresence
	}
	return

}

// SetHeight sets the value for property 'height'.
func (t *Link) SetHeight(k int64) {
	t.raw.SetHeight(k)

}

// GetWidth attempts to get this 'width' property as a int64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Link) GetWidth() (r Resolution, k int64) {
	r = Unresolved
	handled := false
	if t.raw.IsWidth() {
		k = t.raw.GetWidth()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsWidthIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasWidth returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasWidth() (p Presence) {
	p = NoPresence
	if t.raw.IsWidth() {
		p = ConvenientPresence
	} else if t.raw.IsWidthIRI() {
		p = RawPresence
	}
	return

}

// SetWidth sets the value for property 'width'.
func (t *Link) SetWidth(k int64) {
	t.raw.SetWidth(k)

}

// LenPreview returns the number of values this property contains. Each index be used with HasPreview to determine if ResolvePreview is safe to call or if raw handling would be needed.
func (t *Link) LenPreview() (idx int) {
	return t.raw.PreviewLen()

}

// ResolvePreview passes the actual concrete type to the resolver for handing property preview. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) ResolvePreview(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsPreviewObject(idx) {
		handled, err = r.dispatch(t.raw.GetPreviewObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsPreviewLink(idx) {
		handled, err = r.dispatch(t.raw.GetPreviewLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsPreviewIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasPreview returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Link) HasPreview(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsPreviewObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsPreviewLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsPreviewIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendPreview appends an 'Object' typed value.
func (t *Link) AppendPreview(i vocab.ObjectType) {
	t.raw.AppendPreviewObject(i)

}

// PrependPreview prepends an 'Object' typed value.
func (t *Link) PrependPreview(i vocab.ObjectType) {
	t.raw.PrependPreviewObject(i)

}

// AppendPreviewLink appends a 'Link' typed value.
func (t *Link) AppendPreviewLink(i vocab.LinkType) {
	t.raw.AppendPreviewLink(i)

}

// PrependPreviewLink prepends a 'Link' typed value.
func (t *Link) PrependPreviewLink(i vocab.LinkType) {
	t.raw.PrependPreviewLink(i)

}
