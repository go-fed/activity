//
package streams

import (
	"github.com/go-fed/activity/vocab"
	"net/url"
	"time"
)

// Indicates that the actor is blocking the object. Blocking is a stronger form of Ignore. The typical use is to support social systems that allow one user to block activities or content of other users. The target and origin typically have no defined meaning. This is a convenience wrapper of a type with the same name in the vocab package. Accessing it with the Raw function allows direct manipulaton of the object, and does not provide the same integrity guarantees as this package.
type Block struct {
	// The raw type from the vocab package
	raw *vocab.Block
}

// Raw returns the vocab type for manaual manipulation. Note that manipulating the underlying type to be in an inconsistent state may cause this convenience type's methods to later fail.
func (t *Block) Raw() (n *vocab.Block) {
	return t.raw

}

// Serialize turns this object into a map[string]interface{}.
func (t *Block) Serialize() (m map[string]interface{}, err error) {
	return t.raw.Serialize()

}

// IsPublic returns true if the 'to', 'bto', 'cc', or 'bcc' properties address the special Public ActivityPub collection
func (t *Block) IsPublic() (b bool) {
	return t.raw.IsPublic()

}

// LenActor returns the number of values this property contains. Each index be used with HasActor to determine if GetActor is safe to call or if raw handling would be needed.
func (t *Block) LenActor() (idx int) {
	return t.raw.ActorLen()

}

// GetActor attempts to get this 'actor' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetActor(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsActorIRI(idx) {
		k = t.raw.GetActorIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsActorObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsActorLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendActor appends the value for property 'actor'.
func (t *Block) AppendActor(k *url.URL) {
	t.raw.AppendActorIRI(k)

}

// PrependActor prepends the value for property 'actor'.
func (t *Block) PrependActor(k *url.URL) {
	t.raw.PrependActorIRI(k)

}

// RemoveActor deletes the value from the specified index for property 'actor'.
func (t *Block) RemoveActor(idx int) {
	t.raw.RemoveActorIRI(idx)

}

// HasActor returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasActor(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsActorIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsActorLink(idx) {
		p = RawPresence
	} else if t.raw.IsActorIRI(idx) {
		p = RawPresence
	}
	return

}

// LenObject returns the number of values this property contains. Each index be used with HasObject to determine if ResolveObject is safe to call or if raw handling would be needed.
func (t *Block) LenObject() (idx int) {
	return t.raw.ObjectLen()

}

// ResolveObject passes the actual concrete type to the resolver for handing property object. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveObject(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsObject(idx) {
		handled, err = r.dispatch(t.raw.GetObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsObjectIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// AppendObject appends the value for property 'object'.
func (t *Block) AppendObject(i vocab.ObjectType) {
	t.raw.AppendObject(i)

}

// PrependObject prepends the value for property 'object'.
func (t *Block) PrependObject(i vocab.ObjectType) {
	t.raw.PrependObject(i)

}

// RemoveObject deletes the value from the specified index for property 'object'.
func (t *Block) RemoveObject(idx int) {
	t.raw.RemoveObject(idx)

}

// HasObject returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasObject(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsObjectIRI(idx) {
		p = RawPresence
	}
	return

}

// LenTarget returns the number of values this property contains. Each index be used with HasTarget to determine if GetTarget is safe to call or if raw handling would be needed.
func (t *Block) LenTarget() (idx int) {
	return t.raw.TargetLen()

}

// GetTarget attempts to get this 'target' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetTarget(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsTargetIRI(idx) {
		k = t.raw.GetTargetIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsTargetObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsTargetLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendTarget appends the value for property 'target'.
func (t *Block) AppendTarget(k *url.URL) {
	t.raw.AppendTargetIRI(k)

}

// PrependTarget prepends the value for property 'target'.
func (t *Block) PrependTarget(k *url.URL) {
	t.raw.PrependTargetIRI(k)

}

// RemoveTarget deletes the value from the specified index for property 'target'.
func (t *Block) RemoveTarget(idx int) {
	t.raw.RemoveTargetIRI(idx)

}

// HasTarget returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasTarget(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsTargetIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsTargetLink(idx) {
		p = RawPresence
	} else if t.raw.IsTargetIRI(idx) {
		p = RawPresence
	}
	return

}

// LenResult returns the number of values this property contains. Each index be used with HasResult to determine if ResolveResult is safe to call or if raw handling would be needed.
func (t *Block) LenResult() (idx int) {
	return t.raw.ResultLen()

}

// ResolveResult passes the actual concrete type to the resolver for handing property result. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveResult(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsResultObject(idx) {
		handled, err = r.dispatch(t.raw.GetResultObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsResultLink(idx) {
		handled, err = r.dispatch(t.raw.GetResultLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsResultIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasResult returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasResult(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsResultObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsResultLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsResultIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendResult appends an 'Object' typed value.
func (t *Block) AppendResult(i vocab.ObjectType) {
	t.raw.AppendResultObject(i)

}

// PrependResult prepends an 'Object' typed value.
func (t *Block) PrependResult(i vocab.ObjectType) {
	t.raw.PrependResultObject(i)

}

// AppendResultLink appends a 'Link' typed value.
func (t *Block) AppendResultLink(i vocab.LinkType) {
	t.raw.AppendResultLink(i)

}

// PrependResultLink prepends a 'Link' typed value.
func (t *Block) PrependResultLink(i vocab.LinkType) {
	t.raw.PrependResultLink(i)

}

// LenOrigin returns the number of values this property contains. Each index be used with HasOrigin to determine if ResolveOrigin is safe to call or if raw handling would be needed.
func (t *Block) LenOrigin() (idx int) {
	return t.raw.OriginLen()

}

// ResolveOrigin passes the actual concrete type to the resolver for handing property origin. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveOrigin(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsOriginObject(idx) {
		handled, err = r.dispatch(t.raw.GetOriginObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsOriginLink(idx) {
		handled, err = r.dispatch(t.raw.GetOriginLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsOriginIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasOrigin returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasOrigin(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsOriginObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsOriginLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsOriginIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendOrigin appends an 'Object' typed value.
func (t *Block) AppendOrigin(i vocab.ObjectType) {
	t.raw.AppendOriginObject(i)

}

// PrependOrigin prepends an 'Object' typed value.
func (t *Block) PrependOrigin(i vocab.ObjectType) {
	t.raw.PrependOriginObject(i)

}

// AppendOriginLink appends a 'Link' typed value.
func (t *Block) AppendOriginLink(i vocab.LinkType) {
	t.raw.AppendOriginLink(i)

}

// PrependOriginLink prepends a 'Link' typed value.
func (t *Block) PrependOriginLink(i vocab.LinkType) {
	t.raw.PrependOriginLink(i)

}

// LenInstrument returns the number of values this property contains. Each index be used with HasInstrument to determine if ResolveInstrument is safe to call or if raw handling would be needed.
func (t *Block) LenInstrument() (idx int) {
	return t.raw.InstrumentLen()

}

// ResolveInstrument passes the actual concrete type to the resolver for handing property instrument. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveInstrument(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsInstrumentObject(idx) {
		handled, err = r.dispatch(t.raw.GetInstrumentObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsInstrumentLink(idx) {
		handled, err = r.dispatch(t.raw.GetInstrumentLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsInstrumentIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasInstrument returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasInstrument(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsInstrumentObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsInstrumentLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsInstrumentIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendInstrument appends an 'Object' typed value.
func (t *Block) AppendInstrument(i vocab.ObjectType) {
	t.raw.AppendInstrumentObject(i)

}

// PrependInstrument prepends an 'Object' typed value.
func (t *Block) PrependInstrument(i vocab.ObjectType) {
	t.raw.PrependInstrumentObject(i)

}

// AppendInstrumentLink appends a 'Link' typed value.
func (t *Block) AppendInstrumentLink(i vocab.LinkType) {
	t.raw.AppendInstrumentLink(i)

}

// PrependInstrumentLink prepends a 'Link' typed value.
func (t *Block) PrependInstrumentLink(i vocab.LinkType) {
	t.raw.PrependInstrumentLink(i)

}

// GetAltitude attempts to get this 'altitude' property as a float64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetAltitude() (r Resolution, k float64) {
	r = Unresolved
	handled := false
	if t.raw.IsAltitude() {
		k = t.raw.GetAltitude()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAltitudeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasAltitude returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasAltitude() (p Presence) {
	p = NoPresence
	if t.raw.IsAltitude() {
		p = ConvenientPresence
	} else if t.raw.IsAltitudeIRI() {
		p = RawPresence
	}
	return

}

// SetAltitude sets the value for property 'altitude'.
func (t *Block) SetAltitude(k float64) {
	t.raw.SetAltitude(k)

}

// LenAttachment returns the number of values this property contains. Each index be used with HasAttachment to determine if ResolveAttachment is safe to call or if raw handling would be needed.
func (t *Block) LenAttachment() (idx int) {
	return t.raw.AttachmentLen()

}

// ResolveAttachment passes the actual concrete type to the resolver for handing property attachment. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveAttachment(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsAttachmentObject(idx) {
		handled, err = r.dispatch(t.raw.GetAttachmentObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsAttachmentLink(idx) {
		handled, err = r.dispatch(t.raw.GetAttachmentLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsAttachmentIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasAttachment returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasAttachment(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAttachmentObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttachmentLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttachmentIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendAttachment appends an 'Object' typed value.
func (t *Block) AppendAttachment(i vocab.ObjectType) {
	t.raw.AppendAttachmentObject(i)

}

// PrependAttachment prepends an 'Object' typed value.
func (t *Block) PrependAttachment(i vocab.ObjectType) {
	t.raw.PrependAttachmentObject(i)

}

// AppendAttachmentLink appends a 'Link' typed value.
func (t *Block) AppendAttachmentLink(i vocab.LinkType) {
	t.raw.AppendAttachmentLink(i)

}

// PrependAttachmentLink prepends a 'Link' typed value.
func (t *Block) PrependAttachmentLink(i vocab.LinkType) {
	t.raw.PrependAttachmentLink(i)

}

// LenAttributedTo returns the number of values this property contains. Each index be used with HasAttributedTo to determine if GetAttributedTo is safe to call or if raw handling would be needed.
func (t *Block) LenAttributedTo() (idx int) {
	return t.raw.AttributedToLen()

}

// GetAttributedTo attempts to get this 'attributedTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetAttributedTo(idx int) (r Resolution, k *url.URL) {
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
func (t *Block) AppendAttributedTo(k *url.URL) {
	t.raw.AppendAttributedToIRI(k)

}

// PrependAttributedTo prepends the value for property 'attributedTo'.
func (t *Block) PrependAttributedTo(k *url.URL) {
	t.raw.PrependAttributedToIRI(k)

}

// RemoveAttributedTo deletes the value from the specified index for property 'attributedTo'.
func (t *Block) RemoveAttributedTo(idx int) {
	t.raw.RemoveAttributedToIRI(idx)

}

// HasAttributedTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasAttributedTo(idx int) (p Presence) {
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

// LenAudience returns the number of values this property contains. Each index be used with HasAudience to determine if GetAudience is safe to call or if raw handling would be needed.
func (t *Block) LenAudience() (idx int) {
	return t.raw.AudienceLen()

}

// GetAudience attempts to get this 'audience' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetAudience(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsAudienceIRI(idx) {
		k = t.raw.GetAudienceIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAudienceObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsAudienceLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendAudience appends the value for property 'audience'.
func (t *Block) AppendAudience(k *url.URL) {
	t.raw.AppendAudienceIRI(k)

}

// PrependAudience prepends the value for property 'audience'.
func (t *Block) PrependAudience(k *url.URL) {
	t.raw.PrependAudienceIRI(k)

}

// RemoveAudience deletes the value from the specified index for property 'audience'.
func (t *Block) RemoveAudience(idx int) {
	t.raw.RemoveAudienceIRI(idx)

}

// HasAudience returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasAudience(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAudienceIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAudienceLink(idx) {
		p = RawPresence
	} else if t.raw.IsAudienceIRI(idx) {
		p = RawPresence
	}
	return

}

// LenContent returns the number of values this property contains. Each index be used with HasContent to determine if GetContent is safe to call or if raw handling would be needed.
func (t *Block) LenContent() (idx int) {
	return t.raw.ContentLen()

}

// GetContent attempts to get this 'content' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetContent(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsContentString(idx) {
		k = t.raw.GetContentString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsContentLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsContentIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendContent appends the value for property 'content'.
func (t *Block) AppendContent(k string) {
	t.raw.AppendContentString(k)

}

// PrependContent prepends the value for property 'content'.
func (t *Block) PrependContent(k string) {
	t.raw.PrependContentString(k)

}

// RemoveContent deletes the value from the specified index for property 'content'.
func (t *Block) RemoveContent(idx int) {
	t.raw.RemoveContentString(idx)

}

// HasContent returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasContent(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsContentString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContentLangString(idx) {
		p = RawPresence
	} else if t.raw.IsContentIRI(idx) {
		p = RawPresence
	}
	return

}

// ContentLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Block) ContentLanguages() (l []string) {
	return t.raw.ContentMapLanguages()

}

// GetContentMap retrieves the value of 'content' for the specified language, or an empty string if it does not exist
func (t *Block) GetContentForLanguage(l string) (v string) {
	return t.raw.GetContentMap(l)

}

// SetContentForLanguage sets the value of 'content' for the specified language
func (t *Block) SetContentForLanguage(l string, v string) {
	t.raw.SetContentMap(l, v)

}

// LenContext returns the number of values this property contains. Each index be used with HasContext to determine if ResolveContext is safe to call or if raw handling would be needed.
func (t *Block) LenContext() (idx int) {
	return t.raw.ContextLen()

}

// ResolveContext passes the actual concrete type to the resolver for handing property context. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveContext(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsContextObject(idx) {
		handled, err = r.dispatch(t.raw.GetContextObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsContextLink(idx) {
		handled, err = r.dispatch(t.raw.GetContextLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsContextIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasContext returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasContext(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsContextObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContextLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContextIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendContext appends an 'Object' typed value.
func (t *Block) AppendContext(i vocab.ObjectType) {
	t.raw.AppendContextObject(i)

}

// PrependContext prepends an 'Object' typed value.
func (t *Block) PrependContext(i vocab.ObjectType) {
	t.raw.PrependContextObject(i)

}

// AppendContextLink appends a 'Link' typed value.
func (t *Block) AppendContextLink(i vocab.LinkType) {
	t.raw.AppendContextLink(i)

}

// PrependContextLink prepends a 'Link' typed value.
func (t *Block) PrependContextLink(i vocab.LinkType) {
	t.raw.PrependContextLink(i)

}

// LenName returns the number of values this property contains. Each index be used with HasName to determine if GetName is safe to call or if raw handling would be needed.
func (t *Block) LenName() (idx int) {
	return t.raw.NameLen()

}

// GetName attempts to get this 'name' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetName(idx int) (r Resolution, k string) {
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
func (t *Block) AppendName(k string) {
	t.raw.AppendNameString(k)

}

// PrependName prepends the value for property 'name'.
func (t *Block) PrependName(k string) {
	t.raw.PrependNameString(k)

}

// RemoveName deletes the value from the specified index for property 'name'.
func (t *Block) RemoveName(idx int) {
	t.raw.RemoveNameString(idx)

}

// HasName returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasName(idx int) (p Presence) {
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
func (t *Block) NameLanguages() (l []string) {
	return t.raw.NameMapLanguages()

}

// GetNameMap retrieves the value of 'name' for the specified language, or an empty string if it does not exist
func (t *Block) GetNameForLanguage(l string) (v string) {
	return t.raw.GetNameMap(l)

}

// SetNameForLanguage sets the value of 'name' for the specified language
func (t *Block) SetNameForLanguage(l string, v string) {
	t.raw.SetNameMap(l, v)

}

// GetEndTime attempts to get this 'endTime' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetEndTime() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsEndTime() {
		k = t.raw.GetEndTime()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsEndTimeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasEndTime returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasEndTime() (p Presence) {
	p = NoPresence
	if t.raw.IsEndTime() {
		p = ConvenientPresence
	} else if t.raw.IsEndTimeIRI() {
		p = RawPresence
	}
	return

}

// SetEndTime sets the value for property 'endTime'.
func (t *Block) SetEndTime(k time.Time) {
	t.raw.SetEndTime(k)

}

// LenGenerator returns the number of values this property contains. Each index be used with HasGenerator to determine if ResolveGenerator is safe to call or if raw handling would be needed.
func (t *Block) LenGenerator() (idx int) {
	return t.raw.GeneratorLen()

}

// ResolveGenerator passes the actual concrete type to the resolver for handing property generator. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveGenerator(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsGeneratorObject(idx) {
		handled, err = r.dispatch(t.raw.GetGeneratorObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsGeneratorLink(idx) {
		handled, err = r.dispatch(t.raw.GetGeneratorLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsGeneratorIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasGenerator returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasGenerator(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsGeneratorObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsGeneratorLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsGeneratorIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendGenerator appends an 'Object' typed value.
func (t *Block) AppendGenerator(i vocab.ObjectType) {
	t.raw.AppendGeneratorObject(i)

}

// PrependGenerator prepends an 'Object' typed value.
func (t *Block) PrependGenerator(i vocab.ObjectType) {
	t.raw.PrependGeneratorObject(i)

}

// AppendGeneratorLink appends a 'Link' typed value.
func (t *Block) AppendGeneratorLink(i vocab.LinkType) {
	t.raw.AppendGeneratorLink(i)

}

// PrependGeneratorLink prepends a 'Link' typed value.
func (t *Block) PrependGeneratorLink(i vocab.LinkType) {
	t.raw.PrependGeneratorLink(i)

}

// LenIcon returns the number of values this property contains. Each index be used with HasIcon to determine if ResolveIcon is safe to call or if raw handling would be needed.
func (t *Block) LenIcon() (idx int) {
	return t.raw.IconLen()

}

// ResolveIcon passes the actual concrete type to the resolver for handing property icon. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveIcon(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsIconImage(idx) {
		handled, err = r.dispatch(t.raw.GetIconImage(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsIconLink(idx) {
		s = RawResolutionNeeded
	} else if t.raw.IsIconIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// AppendIcon appends the value for property 'icon'.
func (t *Block) AppendIcon(i vocab.ImageType) {
	t.raw.AppendIconImage(i)

}

// PrependIcon prepends the value for property 'icon'.
func (t *Block) PrependIcon(i vocab.ImageType) {
	t.raw.PrependIconImage(i)

}

// RemoveIcon deletes the value from the specified index for property 'icon'.
func (t *Block) RemoveIcon(idx int) {
	t.raw.RemoveIconImage(idx)

}

// HasIcon returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasIcon(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsIconImage(idx) {
		p = ConvenientPresence
	} else if t.raw.IsIconLink(idx) {
		p = RawPresence
	} else if t.raw.IsIconIRI(idx) {
		p = RawPresence
	}
	return

}

// GetId attempts to get this 'id' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetId() (r Resolution, k *url.URL) {
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
func (t *Block) HasId() (p Presence) {
	p = NoPresence
	if t.raw.HasId() {
		p = ConvenientPresence
	}
	return

}

// SetId sets the value for property 'id'.
func (t *Block) SetId(k *url.URL) {
	t.raw.SetId(k)

}

// LenImage returns the number of values this property contains. Each index be used with HasImage to determine if ResolveImage is safe to call or if raw handling would be needed.
func (t *Block) LenImage() (idx int) {
	return t.raw.ImageLen()

}

// ResolveImage passes the actual concrete type to the resolver for handing property image. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveImage(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsImageImage(idx) {
		handled, err = r.dispatch(t.raw.GetImageImage(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsImageLink(idx) {
		s = RawResolutionNeeded
	} else if t.raw.IsImageIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// AppendImage appends the value for property 'image'.
func (t *Block) AppendImage(i vocab.ImageType) {
	t.raw.AppendImageImage(i)

}

// PrependImage prepends the value for property 'image'.
func (t *Block) PrependImage(i vocab.ImageType) {
	t.raw.PrependImageImage(i)

}

// RemoveImage deletes the value from the specified index for property 'image'.
func (t *Block) RemoveImage(idx int) {
	t.raw.RemoveImageImage(idx)

}

// HasImage returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasImage(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsImageImage(idx) {
		p = ConvenientPresence
	} else if t.raw.IsImageLink(idx) {
		p = RawPresence
	} else if t.raw.IsImageIRI(idx) {
		p = RawPresence
	}
	return

}

// LenInReplyTo returns the number of values this property contains. Each index be used with HasInReplyTo to determine if GetInReplyTo is safe to call or if raw handling would be needed.
func (t *Block) LenInReplyTo() (idx int) {
	return t.raw.InReplyToLen()

}

// GetInReplyTo attempts to get this 'inReplyTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetInReplyTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsInReplyToIRI(idx) {
		k = t.raw.GetInReplyToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsInReplyToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsInReplyToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendInReplyTo appends the value for property 'inReplyTo'.
func (t *Block) AppendInReplyTo(k *url.URL) {
	t.raw.AppendInReplyToIRI(k)

}

// PrependInReplyTo prepends the value for property 'inReplyTo'.
func (t *Block) PrependInReplyTo(k *url.URL) {
	t.raw.PrependInReplyToIRI(k)

}

// RemoveInReplyTo deletes the value from the specified index for property 'inReplyTo'.
func (t *Block) RemoveInReplyTo(idx int) {
	t.raw.RemoveInReplyToIRI(idx)

}

// HasInReplyTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasInReplyTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsInReplyToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsInReplyToLink(idx) {
		p = RawPresence
	} else if t.raw.IsInReplyToIRI(idx) {
		p = RawPresence
	}
	return

}

// LenLocation returns the number of values this property contains. Each index be used with HasLocation to determine if ResolveLocation is safe to call or if raw handling would be needed.
func (t *Block) LenLocation() (idx int) {
	return t.raw.LocationLen()

}

// ResolveLocation passes the actual concrete type to the resolver for handing property location. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveLocation(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLocationObject(idx) {
		handled, err = r.dispatch(t.raw.GetLocationObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLocationLink(idx) {
		handled, err = r.dispatch(t.raw.GetLocationLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLocationIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasLocation returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasLocation(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsLocationObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsLocationLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsLocationIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendLocation appends an 'Object' typed value.
func (t *Block) AppendLocation(i vocab.ObjectType) {
	t.raw.AppendLocationObject(i)

}

// PrependLocation prepends an 'Object' typed value.
func (t *Block) PrependLocation(i vocab.ObjectType) {
	t.raw.PrependLocationObject(i)

}

// AppendLocationLink appends a 'Link' typed value.
func (t *Block) AppendLocationLink(i vocab.LinkType) {
	t.raw.AppendLocationLink(i)

}

// PrependLocationLink prepends a 'Link' typed value.
func (t *Block) PrependLocationLink(i vocab.LinkType) {
	t.raw.PrependLocationLink(i)

}

// LenPreview returns the number of values this property contains. Each index be used with HasPreview to determine if ResolvePreview is safe to call or if raw handling would be needed.
func (t *Block) LenPreview() (idx int) {
	return t.raw.PreviewLen()

}

// ResolvePreview passes the actual concrete type to the resolver for handing property preview. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolvePreview(r *Resolver, idx int) (s Resolution, err error) {
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
func (t *Block) HasPreview(idx int) (p Presence) {
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
func (t *Block) AppendPreview(i vocab.ObjectType) {
	t.raw.AppendPreviewObject(i)

}

// PrependPreview prepends an 'Object' typed value.
func (t *Block) PrependPreview(i vocab.ObjectType) {
	t.raw.PrependPreviewObject(i)

}

// AppendPreviewLink appends a 'Link' typed value.
func (t *Block) AppendPreviewLink(i vocab.LinkType) {
	t.raw.AppendPreviewLink(i)

}

// PrependPreviewLink prepends a 'Link' typed value.
func (t *Block) PrependPreviewLink(i vocab.LinkType) {
	t.raw.PrependPreviewLink(i)

}

// GetPublished attempts to get this 'published' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetPublished() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsPublished() {
		k = t.raw.GetPublished()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsPublishedIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasPublished returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasPublished() (p Presence) {
	p = NoPresence
	if t.raw.IsPublished() {
		p = ConvenientPresence
	} else if t.raw.IsPublishedIRI() {
		p = RawPresence
	}
	return

}

// SetPublished sets the value for property 'published'.
func (t *Block) SetPublished(k time.Time) {
	t.raw.SetPublished(k)

}

// ResolveReplies passes the actual concrete type to the resolver for handing property replies. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveReplies(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsReplies() {
		handled, err = r.dispatch(t.raw.GetReplies())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsRepliesIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasReplies returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasReplies() (p Presence) {
	p = NoPresence
	if t.raw.IsReplies() {
		p = ConvenientPresence
	} else if t.raw.IsRepliesIRI() {
		p = RawPresence
	}
	return

}

// SetReplies sets this value to be a 'Collection' type.
func (t *Block) SetReplies(i vocab.CollectionType) {
	t.raw.SetReplies(i)

}

// GetStartTime attempts to get this 'startTime' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetStartTime() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsStartTime() {
		k = t.raw.GetStartTime()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsStartTimeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasStartTime returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasStartTime() (p Presence) {
	p = NoPresence
	if t.raw.IsStartTime() {
		p = ConvenientPresence
	} else if t.raw.IsStartTimeIRI() {
		p = RawPresence
	}
	return

}

// SetStartTime sets the value for property 'startTime'.
func (t *Block) SetStartTime(k time.Time) {
	t.raw.SetStartTime(k)

}

// LenSummary returns the number of values this property contains. Each index be used with HasSummary to determine if GetSummary is safe to call or if raw handling would be needed.
func (t *Block) LenSummary() (idx int) {
	return t.raw.SummaryLen()

}

// GetSummary attempts to get this 'summary' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetSummary(idx int) (r Resolution, k string) {
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
func (t *Block) AppendSummary(k string) {
	t.raw.AppendSummaryString(k)

}

// PrependSummary prepends the value for property 'summary'.
func (t *Block) PrependSummary(k string) {
	t.raw.PrependSummaryString(k)

}

// RemoveSummary deletes the value from the specified index for property 'summary'.
func (t *Block) RemoveSummary(idx int) {
	t.raw.RemoveSummaryString(idx)

}

// HasSummary returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasSummary(idx int) (p Presence) {
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
func (t *Block) SummaryLanguages() (l []string) {
	return t.raw.SummaryMapLanguages()

}

// GetSummaryMap retrieves the value of 'summary' for the specified language, or an empty string if it does not exist
func (t *Block) GetSummaryForLanguage(l string) (v string) {
	return t.raw.GetSummaryMap(l)

}

// SetSummaryForLanguage sets the value of 'summary' for the specified language
func (t *Block) SetSummaryForLanguage(l string, v string) {
	t.raw.SetSummaryMap(l, v)

}

// LenTag returns the number of values this property contains. Each index be used with HasTag to determine if ResolveTag is safe to call or if raw handling would be needed.
func (t *Block) LenTag() (idx int) {
	return t.raw.TagLen()

}

// ResolveTag passes the actual concrete type to the resolver for handing property tag. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveTag(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsTagObject(idx) {
		handled, err = r.dispatch(t.raw.GetTagObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsTagLink(idx) {
		handled, err = r.dispatch(t.raw.GetTagLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsTagIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasTag returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasTag(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsTagObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsTagLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsTagIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendTag appends an 'Object' typed value.
func (t *Block) AppendTag(i vocab.ObjectType) {
	t.raw.AppendTagObject(i)

}

// PrependTag prepends an 'Object' typed value.
func (t *Block) PrependTag(i vocab.ObjectType) {
	t.raw.PrependTagObject(i)

}

// AppendTagLink appends a 'Link' typed value.
func (t *Block) AppendTagLink(i vocab.LinkType) {
	t.raw.AppendTagLink(i)

}

// PrependTagLink prepends a 'Link' typed value.
func (t *Block) PrependTagLink(i vocab.LinkType) {
	t.raw.PrependTagLink(i)

}

// LenType returns the number of values this property contains. Each index be used with HasType to determine if GetType is safe to call or if raw handling would be needed.
func (t *Block) LenType() (idx int) {
	return t.raw.TypeLen()

}

// GetType attempts to get this 'type' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetType(idx int) (r Resolution, s string) {
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
func (t *Block) AppendType(i interface{}) {
	t.raw.AppendType(i)

}

// PrependType prepends the value for property 'type'.
func (t *Block) PrependType(i interface{}) {
	t.raw.PrependType(i)

}

// RemoveType deletes the value from the specified index for property 'type'.
func (t *Block) RemoveType(idx int) {
	t.raw.RemoveType(idx)

}

// GetUpdated attempts to get this 'updated' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetUpdated() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsUpdated() {
		k = t.raw.GetUpdated()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsUpdatedIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasUpdated returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasUpdated() (p Presence) {
	p = NoPresence
	if t.raw.IsUpdated() {
		p = ConvenientPresence
	} else if t.raw.IsUpdatedIRI() {
		p = RawPresence
	}
	return

}

// SetUpdated sets the value for property 'updated'.
func (t *Block) SetUpdated(k time.Time) {
	t.raw.SetUpdated(k)

}

// LenUrl returns the number of values this property contains. Each index be used with HasUrl to determine if GetUrl is safe to call or if raw handling would be needed.
func (t *Block) LenUrl() (idx int) {
	return t.raw.UrlLen()

}

// GetUrl attempts to get this 'url' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetUrl(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsUrlAnyURI(idx) {
		k = t.raw.GetUrlAnyURI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsUrlLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendUrl appends the value for property 'url'.
func (t *Block) AppendUrl(k *url.URL) {
	t.raw.AppendUrlAnyURI(k)

}

// PrependUrl prepends the value for property 'url'.
func (t *Block) PrependUrl(k *url.URL) {
	t.raw.PrependUrlAnyURI(k)

}

// RemoveUrl deletes the value from the specified index for property 'url'.
func (t *Block) RemoveUrl(idx int) {
	t.raw.RemoveUrlAnyURI(idx)

}

// HasUrl returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasUrl(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsUrlAnyURI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsUrlLink(idx) {
		p = RawPresence
	}
	return

}

// LenTo returns the number of values this property contains. Each index be used with HasTo to determine if GetTo is safe to call or if raw handling would be needed.
func (t *Block) LenTo() (idx int) {
	return t.raw.ToLen()

}

// GetTo attempts to get this 'to' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsToIRI(idx) {
		k = t.raw.GetToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendTo appends the value for property 'to'.
func (t *Block) AppendTo(k *url.URL) {
	t.raw.AppendToIRI(k)

}

// PrependTo prepends the value for property 'to'.
func (t *Block) PrependTo(k *url.URL) {
	t.raw.PrependToIRI(k)

}

// RemoveTo deletes the value from the specified index for property 'to'.
func (t *Block) RemoveTo(idx int) {
	t.raw.RemoveToIRI(idx)

}

// HasTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsToLink(idx) {
		p = RawPresence
	} else if t.raw.IsToIRI(idx) {
		p = RawPresence
	}
	return

}

// LenBto returns the number of values this property contains. Each index be used with HasBto to determine if GetBto is safe to call or if raw handling would be needed.
func (t *Block) LenBto() (idx int) {
	return t.raw.BtoLen()

}

// GetBto attempts to get this 'bto' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetBto(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsBtoIRI(idx) {
		k = t.raw.GetBtoIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsBtoObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsBtoLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendBto appends the value for property 'bto'.
func (t *Block) AppendBto(k *url.URL) {
	t.raw.AppendBtoIRI(k)

}

// PrependBto prepends the value for property 'bto'.
func (t *Block) PrependBto(k *url.URL) {
	t.raw.PrependBtoIRI(k)

}

// RemoveBto deletes the value from the specified index for property 'bto'.
func (t *Block) RemoveBto(idx int) {
	t.raw.RemoveBtoIRI(idx)

}

// HasBto returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasBto(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsBtoIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsBtoLink(idx) {
		p = RawPresence
	} else if t.raw.IsBtoIRI(idx) {
		p = RawPresence
	}
	return

}

// LenCc returns the number of values this property contains. Each index be used with HasCc to determine if GetCc is safe to call or if raw handling would be needed.
func (t *Block) LenCc() (idx int) {
	return t.raw.CcLen()

}

// GetCc attempts to get this 'cc' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetCc(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsCcIRI(idx) {
		k = t.raw.GetCcIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsCcObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsCcLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendCc appends the value for property 'cc'.
func (t *Block) AppendCc(k *url.URL) {
	t.raw.AppendCcIRI(k)

}

// PrependCc prepends the value for property 'cc'.
func (t *Block) PrependCc(k *url.URL) {
	t.raw.PrependCcIRI(k)

}

// RemoveCc deletes the value from the specified index for property 'cc'.
func (t *Block) RemoveCc(idx int) {
	t.raw.RemoveCcIRI(idx)

}

// HasCc returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasCc(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsCcIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsCcLink(idx) {
		p = RawPresence
	} else if t.raw.IsCcIRI(idx) {
		p = RawPresence
	}
	return

}

// LenBcc returns the number of values this property contains. Each index be used with HasBcc to determine if GetBcc is safe to call or if raw handling would be needed.
func (t *Block) LenBcc() (idx int) {
	return t.raw.BccLen()

}

// GetBcc attempts to get this 'bcc' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetBcc(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsBccIRI(idx) {
		k = t.raw.GetBccIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsBccObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsBccLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendBcc appends the value for property 'bcc'.
func (t *Block) AppendBcc(k *url.URL) {
	t.raw.AppendBccIRI(k)

}

// PrependBcc prepends the value for property 'bcc'.
func (t *Block) PrependBcc(k *url.URL) {
	t.raw.PrependBccIRI(k)

}

// RemoveBcc deletes the value from the specified index for property 'bcc'.
func (t *Block) RemoveBcc(idx int) {
	t.raw.RemoveBccIRI(idx)

}

// HasBcc returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasBcc(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsBccIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsBccLink(idx) {
		p = RawPresence
	} else if t.raw.IsBccIRI(idx) {
		p = RawPresence
	}
	return

}

// GetMediaType attempts to get this 'mediaType' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetMediaType() (r Resolution, k string) {
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
func (t *Block) HasMediaType() (p Presence) {
	p = NoPresence
	if t.raw.IsMediaType() {
		p = ConvenientPresence
	} else if t.raw.IsMediaTypeIRI() {
		p = RawPresence
	}
	return

}

// SetMediaType sets the value for property 'mediaType'.
func (t *Block) SetMediaType(k string) {
	t.raw.SetMediaType(k)

}

// GetDuration attempts to get this 'duration' property as a time.Duration. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetDuration() (r Resolution, k time.Duration) {
	r = Unresolved
	handled := false
	if t.raw.IsDuration() {
		k = t.raw.GetDuration()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsDurationIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasDuration returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasDuration() (p Presence) {
	p = NoPresence
	if t.raw.IsDuration() {
		p = ConvenientPresence
	} else if t.raw.IsDurationIRI() {
		p = RawPresence
	}
	return

}

// SetDuration sets the value for property 'duration'.
func (t *Block) SetDuration(k time.Duration) {
	t.raw.SetDuration(k)

}

// ResolveSource passes the actual concrete type to the resolver for handing property source. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveSource(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsSource() {
		handled, err = r.dispatch(t.raw.GetSource())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsSourceIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasSource returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasSource() (p Presence) {
	p = NoPresence
	if t.raw.IsSource() {
		p = ConvenientPresence
	} else if t.raw.IsSourceIRI() {
		p = RawPresence
	}
	return

}

// SetSource sets this value to be a 'Object' type.
func (t *Block) SetSource(i vocab.ObjectType) {
	t.raw.SetSource(i)

}

// ResolveInbox passes the actual concrete type to the resolver for handing property inbox. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveInbox(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsInboxOrderedCollection() {
		handled, err = r.dispatch(t.raw.GetInboxOrderedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsInboxAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasInbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasInbox() (p Presence) {
	p = NoPresence
	if t.raw.IsInboxOrderedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsInboxAnyURI() {
		p = RawPresence
	}
	return

}

// SetInbox sets this value to be a 'OrderedCollection' type.
func (t *Block) SetInbox(i vocab.OrderedCollectionType) {
	t.raw.SetInboxOrderedCollection(i)

}

// ResolveOutbox passes the actual concrete type to the resolver for handing property outbox. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveOutbox(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsOutboxOrderedCollection() {
		handled, err = r.dispatch(t.raw.GetOutboxOrderedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsOutboxAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasOutbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasOutbox() (p Presence) {
	p = NoPresence
	if t.raw.IsOutboxOrderedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsOutboxAnyURI() {
		p = RawPresence
	}
	return

}

// SetOutbox sets this value to be a 'OrderedCollection' type.
func (t *Block) SetOutbox(i vocab.OrderedCollectionType) {
	t.raw.SetOutboxOrderedCollection(i)

}

// ResolveFollowing passes the actual concrete type to the resolver for handing property following. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveFollowing(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsFollowingCollection() {
		handled, err = r.dispatch(t.raw.GetFollowingCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsFollowingOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsFollowingAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasFollowing returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasFollowing() (p Presence) {
	p = NoPresence
	if t.raw.IsFollowingCollection() {
		p = ConvenientPresence
	} else if t.raw.IsFollowingOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsFollowingAnyURI() {
		p = RawPresence
	}
	return

}

// SetFollowing sets this value to be a 'Collection' type.
func (t *Block) SetFollowing(i vocab.CollectionType) {
	t.raw.SetFollowingCollection(i)

}

// ResolveFollowers passes the actual concrete type to the resolver for handing property followers. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveFollowers(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsFollowersCollection() {
		handled, err = r.dispatch(t.raw.GetFollowersCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsFollowersOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsFollowersAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasFollowers returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasFollowers() (p Presence) {
	p = NoPresence
	if t.raw.IsFollowersCollection() {
		p = ConvenientPresence
	} else if t.raw.IsFollowersOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsFollowersAnyURI() {
		p = RawPresence
	}
	return

}

// SetFollowers sets this value to be a 'Collection' type.
func (t *Block) SetFollowers(i vocab.CollectionType) {
	t.raw.SetFollowersCollection(i)

}

// ResolveLiked passes the actual concrete type to the resolver for handing property liked. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveLiked(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLikedCollection() {
		handled, err = r.dispatch(t.raw.GetLikedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLikedOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsLikedAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasLiked returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasLiked() (p Presence) {
	p = NoPresence
	if t.raw.IsLikedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsLikedOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsLikedAnyURI() {
		p = RawPresence
	}
	return

}

// SetLiked sets this value to be a 'Collection' type.
func (t *Block) SetLiked(i vocab.CollectionType) {
	t.raw.SetLikedCollection(i)

}

// ResolveLikes passes the actual concrete type to the resolver for handing property likes. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveLikes(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLikesCollection() {
		handled, err = r.dispatch(t.raw.GetLikesCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLikesOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsLikesAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasLikes returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasLikes() (p Presence) {
	p = NoPresence
	if t.raw.IsLikesCollection() {
		p = ConvenientPresence
	} else if t.raw.IsLikesOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsLikesAnyURI() {
		p = RawPresence
	}
	return

}

// SetLikes sets this value to be a 'Collection' type.
func (t *Block) SetLikes(i vocab.CollectionType) {
	t.raw.SetLikesCollection(i)

}

// LenStreams returns the number of values this property contains. Each index be used with HasStreams to determine if GetStreams is safe to call or if raw handling would be needed.
func (t *Block) LenStreams() (idx int) {
	return t.raw.StreamsLen()

}

// GetStreams attempts to get this 'streams' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetStreams(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if /*t.raw.HasStreams(idx)*/ true {
		k = t.raw.GetStreams(idx)
		if handled {
			r = Resolved
		}
	}
	return

}

// AppendStreams appends the value for property 'streams'.
func (t *Block) AppendStreams(k *url.URL) {
	t.raw.AppendStreams(k)

}

// PrependStreams prepends the value for property 'streams'.
func (t *Block) PrependStreams(k *url.URL) {
	t.raw.PrependStreams(k)

}

// RemoveStreams deletes the value from the specified index for property 'streams'.
func (t *Block) RemoveStreams(idx int) {
	t.raw.RemoveStreams(idx)

}

// GetPreferredUsername attempts to get this 'preferredUsername' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetPreferredUsername() (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsPreferredUsername() {
		k = t.raw.GetPreferredUsername()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsPreferredUsernameIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasPreferredUsername returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasPreferredUsername() (p Presence) {
	p = NoPresence
	if t.raw.IsPreferredUsername() {
		p = ConvenientPresence
	} else if t.raw.IsPreferredUsernameIRI() {
		p = RawPresence
	}
	return

}

// SetPreferredUsername sets the value for property 'preferredUsername'.
func (t *Block) SetPreferredUsername(k string) {
	t.raw.SetPreferredUsername(k)

}

// PreferredUsernameLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Block) PreferredUsernameLanguages() (l []string) {
	return t.raw.PreferredUsernameMapLanguages()

}

// GetPreferredUsernameMap retrieves the value of 'preferredUsername' for the specified language, or an empty string if it does not exist
func (t *Block) GetPreferredUsernameForLanguage(l string) (v string) {
	return t.raw.GetPreferredUsernameMap(l)

}

// SetPreferredUsernameForLanguage sets the value of 'preferredUsername' for the specified language
func (t *Block) SetPreferredUsernameForLanguage(l string, v string) {
	t.raw.SetPreferredUsernameMap(l, v)

}

// ResolveEndpoints passes the actual concrete type to the resolver for handing property endpoints. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveEndpoints(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsEndpoints() {
		handled, err = r.dispatch(t.raw.GetEndpoints())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsEndpointsIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasEndpoints returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasEndpoints() (p Presence) {
	p = NoPresence
	if t.raw.IsEndpoints() {
		p = ConvenientPresence
	} else if t.raw.IsEndpointsIRI() {
		p = RawPresence
	}
	return

}

// SetEndpoints sets this value to be a 'Object' type.
func (t *Block) SetEndpoints(i vocab.ObjectType) {
	t.raw.SetEndpoints(i)

}

// GetProxyUrl attempts to get this 'proxyUrl' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetProxyUrl() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasProxyUrl() {
		k = t.raw.GetProxyUrl()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasProxyUrl returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasProxyUrl() (p Presence) {
	p = NoPresence
	if t.raw.HasProxyUrl() {
		p = ConvenientPresence
	}
	return

}

// SetProxyUrl sets the value for property 'proxyUrl'.
func (t *Block) SetProxyUrl(k *url.URL) {
	t.raw.SetProxyUrl(k)

}

// GetOauthAuthorizationEndpoint attempts to get this 'oauthAuthorizationEndpoint' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetOauthAuthorizationEndpoint() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasOauthAuthorizationEndpoint() {
		k = t.raw.GetOauthAuthorizationEndpoint()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasOauthAuthorizationEndpoint returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasOauthAuthorizationEndpoint() (p Presence) {
	p = NoPresence
	if t.raw.HasOauthAuthorizationEndpoint() {
		p = ConvenientPresence
	}
	return

}

// SetOauthAuthorizationEndpoint sets the value for property 'oauthAuthorizationEndpoint'.
func (t *Block) SetOauthAuthorizationEndpoint(k *url.URL) {
	t.raw.SetOauthAuthorizationEndpoint(k)

}

// GetOauthTokenEndpoint attempts to get this 'oauthTokenEndpoint' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetOauthTokenEndpoint() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasOauthTokenEndpoint() {
		k = t.raw.GetOauthTokenEndpoint()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasOauthTokenEndpoint returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasOauthTokenEndpoint() (p Presence) {
	p = NoPresence
	if t.raw.HasOauthTokenEndpoint() {
		p = ConvenientPresence
	}
	return

}

// SetOauthTokenEndpoint sets the value for property 'oauthTokenEndpoint'.
func (t *Block) SetOauthTokenEndpoint(k *url.URL) {
	t.raw.SetOauthTokenEndpoint(k)

}

// GetProvideClientKey attempts to get this 'provideClientKey' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetProvideClientKey() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasProvideClientKey() {
		k = t.raw.GetProvideClientKey()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasProvideClientKey returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasProvideClientKey() (p Presence) {
	p = NoPresence
	if t.raw.HasProvideClientKey() {
		p = ConvenientPresence
	}
	return

}

// SetProvideClientKey sets the value for property 'provideClientKey'.
func (t *Block) SetProvideClientKey(k *url.URL) {
	t.raw.SetProvideClientKey(k)

}

// GetSignClientKey attempts to get this 'signClientKey' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetSignClientKey() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasSignClientKey() {
		k = t.raw.GetSignClientKey()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasSignClientKey returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasSignClientKey() (p Presence) {
	p = NoPresence
	if t.raw.HasSignClientKey() {
		p = ConvenientPresence
	}
	return

}

// SetSignClientKey sets the value for property 'signClientKey'.
func (t *Block) SetSignClientKey(k *url.URL) {
	t.raw.SetSignClientKey(k)

}

// GetSharedInbox attempts to get this 'sharedInbox' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Block) GetSharedInbox() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasSharedInbox() {
		k = t.raw.GetSharedInbox()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasSharedInbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasSharedInbox() (p Presence) {
	p = NoPresence
	if t.raw.HasSharedInbox() {
		p = ConvenientPresence
	}
	return

}

// SetSharedInbox sets the value for property 'sharedInbox'.
func (t *Block) SetSharedInbox(k *url.URL) {
	t.raw.SetSharedInbox(k)

}

// ResolveShares passes the actual concrete type to the resolver for handing property shares. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) ResolveShares(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsSharesCollection() {
		handled, err = r.dispatch(t.raw.GetSharesCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsSharesOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsSharesAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasShares returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Block) HasShares() (p Presence) {
	p = NoPresence
	if t.raw.IsSharesCollection() {
		p = ConvenientPresence
	} else if t.raw.IsSharesOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsSharesAnyURI() {
		p = RawPresence
	}
	return

}

// SetShares sets this value to be a 'Collection' type.
func (t *Block) SetShares(i vocab.CollectionType) {
	t.raw.SetSharesCollection(i)

}
