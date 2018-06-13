//
package streams

import (
	"github.com/go-fed/activity/vocab"
	"net/url"
)

// A specialized Link that represents an @mention. This is a convenience wrapper of a type with the same name in the vocab package. Accessing it with the Raw function allows direct manipulaton of the object, and does not provide the same integrity guarantees as this package.
type Mention struct {
	// The raw type from the vocab package
	raw *vocab.Mention
}

// Raw returns the vocab type for manaual manipulation. Note that manipulating the underlying type to be in an inconsistent state may cause this convenience type's methods to later fail.
func (t *Mention) Raw() (n *vocab.Mention) {
	return t.raw

}

// Serialize turns this object into a map[string]interface{}.
func (t *Mention) Serialize() (m map[string]interface{}, err error) {
	return t.raw.Serialize()

}

// LenAttributedTo returns the number of values this property contains. Each index be used with HasAttributedTo to determine if GetAttributedTo is safe to call or if raw handling would be needed.
func (t *Mention) LenAttributedTo() (idx int) {
	return t.raw.AttributedToLen()

}

// GetAttributedTo attempts to get this 'attributedTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetAttributedTo(idx int) (r Resolution, k *url.URL) {
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
func (t *Mention) AppendAttributedTo(k *url.URL) {
	t.raw.AppendAttributedToIRI(k)

}

// PrependAttributedTo prepends the value for property 'attributedTo'.
func (t *Mention) PrependAttributedTo(k *url.URL) {
	t.raw.PrependAttributedToIRI(k)

}

// RemoveAttributedTo deletes the value from the specified index for property 'attributedTo'.
func (t *Mention) RemoveAttributedTo(idx int) {
	t.raw.RemoveAttributedToIRI(idx)

}

// HasAttributedTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Mention) HasAttributedTo(idx int) (p Presence) {
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
func (t *Mention) GetHref() (r Resolution, k *url.URL) {
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
func (t *Mention) HasHref() (p Presence) {
	p = NoPresence
	if t.raw.HasHref() {
		p = ConvenientPresence
	}
	return

}

// SetHref sets the value for property 'href'.
func (t *Mention) SetHref(k *url.URL) {
	t.raw.SetHref(k)

}

// GetId attempts to get this 'id' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetId() (r Resolution, k *url.URL) {
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
func (t *Mention) HasId() (p Presence) {
	p = NoPresence
	if t.raw.HasId() {
		p = ConvenientPresence
	}
	return

}

// SetId sets the value for property 'id'.
func (t *Mention) SetId(k *url.URL) {
	t.raw.SetId(k)

}

// LenRel returns the number of values this property contains. Each index be used with HasRel to determine if GetRel is safe to call or if raw handling would be needed.
func (t *Mention) LenRel() (idx int) {
	return t.raw.RelLen()

}

// GetRel attempts to get this 'rel' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetRel(idx int) (r Resolution, k string) {
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
func (t *Mention) AppendRel(k string) {
	t.raw.AppendRel(k)

}

// PrependRel prepends the value for property 'rel'.
func (t *Mention) PrependRel(k string) {
	t.raw.PrependRel(k)

}

// RemoveRel deletes the value from the specified index for property 'rel'.
func (t *Mention) RemoveRel(idx int) {
	t.raw.RemoveRel(idx)

}

// HasRel returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Mention) HasRel(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsRel(idx) {
		p = ConvenientPresence
	} else if t.raw.IsRelIRI(idx) {
		p = RawPresence
	}
	return

}

// LenType returns the number of values this property contains. Each index be used with HasType to determine if GetType is safe to call or if raw handling would be needed.
func (t *Mention) LenType() (idx int) {
	return t.raw.TypeLen()

}

// GetType attempts to get this 'type' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetType(idx int) (r Resolution, s string) {
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
func (t *Mention) AppendType(i interface{}) {
	t.raw.AppendType(i)

}

// PrependType prepends the value for property 'type'.
func (t *Mention) PrependType(i interface{}) {
	t.raw.PrependType(i)

}

// RemoveType deletes the value from the specified index for property 'type'.
func (t *Mention) RemoveType(idx int) {
	t.raw.RemoveType(idx)

}

// GetMediaType attempts to get this 'mediaType' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetMediaType() (r Resolution, k string) {
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
func (t *Mention) HasMediaType() (p Presence) {
	p = NoPresence
	if t.raw.IsMediaType() {
		p = ConvenientPresence
	} else if t.raw.IsMediaTypeIRI() {
		p = RawPresence
	}
	return

}

// SetMediaType sets the value for property 'mediaType'.
func (t *Mention) SetMediaType(k string) {
	t.raw.SetMediaType(k)

}

// LenName returns the number of values this property contains. Each index be used with HasName to determine if GetName is safe to call or if raw handling would be needed.
func (t *Mention) LenName() (idx int) {
	return t.raw.NameLen()

}

// GetName attempts to get this 'name' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetName(idx int) (r Resolution, k string) {
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
func (t *Mention) AppendName(k string) {
	t.raw.AppendNameString(k)

}

// PrependName prepends the value for property 'name'.
func (t *Mention) PrependName(k string) {
	t.raw.PrependNameString(k)

}

// RemoveName deletes the value from the specified index for property 'name'.
func (t *Mention) RemoveName(idx int) {
	t.raw.RemoveNameString(idx)

}

// HasName returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Mention) HasName(idx int) (p Presence) {
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
func (t *Mention) NameLanguages() (l []string) {
	return t.raw.NameMapLanguages()

}

// GetNameMap retrieves the value of 'name' for the specified language, or an empty string if it does not exist
func (t *Mention) GetNameForLanguage(l string) (v string) {
	return t.raw.GetNameMap(l)

}

// SetNameForLanguage sets the value of 'name' for the specified language
func (t *Mention) SetNameForLanguage(l string, v string) {
	t.raw.SetNameMap(l, v)

}

// LenSummary returns the number of values this property contains. Each index be used with HasSummary to determine if GetSummary is safe to call or if raw handling would be needed.
func (t *Mention) LenSummary() (idx int) {
	return t.raw.SummaryLen()

}

// GetSummary attempts to get this 'summary' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetSummary(idx int) (r Resolution, k string) {
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
func (t *Mention) AppendSummary(k string) {
	t.raw.AppendSummaryString(k)

}

// PrependSummary prepends the value for property 'summary'.
func (t *Mention) PrependSummary(k string) {
	t.raw.PrependSummaryString(k)

}

// RemoveSummary deletes the value from the specified index for property 'summary'.
func (t *Mention) RemoveSummary(idx int) {
	t.raw.RemoveSummaryString(idx)

}

// HasSummary returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Mention) HasSummary(idx int) (p Presence) {
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
func (t *Mention) SummaryLanguages() (l []string) {
	return t.raw.SummaryMapLanguages()

}

// GetSummaryMap retrieves the value of 'summary' for the specified language, or an empty string if it does not exist
func (t *Mention) GetSummaryForLanguage(l string) (v string) {
	return t.raw.GetSummaryMap(l)

}

// SetSummaryForLanguage sets the value of 'summary' for the specified language
func (t *Mention) SetSummaryForLanguage(l string, v string) {
	t.raw.SetSummaryMap(l, v)

}

// GetHreflang attempts to get this 'hreflang' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetHreflang() (r Resolution, k string) {
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
func (t *Mention) HasHreflang() (p Presence) {
	p = NoPresence
	if t.raw.IsHreflang() {
		p = ConvenientPresence
	} else if t.raw.IsHreflangIRI() {
		p = RawPresence
	}
	return

}

// SetHreflang sets the value for property 'hreflang'.
func (t *Mention) SetHreflang(k string) {
	t.raw.SetHreflang(k)

}

// GetHeight attempts to get this 'height' property as a int64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetHeight() (r Resolution, k int64) {
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
func (t *Mention) HasHeight() (p Presence) {
	p = NoPresence
	if t.raw.IsHeight() {
		p = ConvenientPresence
	} else if t.raw.IsHeightIRI() {
		p = RawPresence
	}
	return

}

// SetHeight sets the value for property 'height'.
func (t *Mention) SetHeight(k int64) {
	t.raw.SetHeight(k)

}

// GetWidth attempts to get this 'width' property as a int64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Mention) GetWidth() (r Resolution, k int64) {
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
func (t *Mention) HasWidth() (p Presence) {
	p = NoPresence
	if t.raw.IsWidth() {
		p = ConvenientPresence
	} else if t.raw.IsWidthIRI() {
		p = RawPresence
	}
	return

}

// SetWidth sets the value for property 'width'.
func (t *Mention) SetWidth(k int64) {
	t.raw.SetWidth(k)

}

// LenPreview returns the number of values this property contains. Each index be used with HasPreview to determine if ResolvePreview is safe to call or if raw handling would be needed.
func (t *Mention) LenPreview() (idx int) {
	return t.raw.PreviewLen()

}

// ResolvePreview passes the actual concrete type to the resolver for handing property preview. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Mention) ResolvePreview(r *Resolver, idx int) (s Resolution, err error) {
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
func (t *Mention) HasPreview(idx int) (p Presence) {
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
func (t *Mention) AppendPreview(i vocab.ObjectType) {
	t.raw.AppendPreviewObject(i)

}

// PrependPreview prepends an 'Object' typed value.
func (t *Mention) PrependPreview(i vocab.ObjectType) {
	t.raw.PrependPreviewObject(i)

}

// AppendPreviewLink appends a 'Link' typed value.
func (t *Mention) AppendPreviewLink(i vocab.LinkType) {
	t.raw.AppendPreviewLink(i)

}

// PrependPreviewLink prepends a 'Link' typed value.
func (t *Mention) PrependPreviewLink(i vocab.LinkType) {
	t.raw.PrependPreviewLink(i)

}
