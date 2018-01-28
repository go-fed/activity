package pub

import (
	"github.com/go-fed/activity/vocab"
	"net/url"
)

// ActorObject is an object that has "actor" or "attributedTo" properties,
// representing the author or originator of the object.
type ActorObject interface {
	HasInbox() (ok bool)
	GetInbox() (v url.URL)
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v vocab.ObjectType)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v vocab.LinkType)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v url.URL)
	ActorLen() (l int)
	IsActorObject(index int) (ok bool)
	GetActorObject(index int) (v vocab.ObjectType)
	IsActorLink(index int) (ok bool)
	GetActorLink(index int) (v vocab.LinkType)
	IsActorIRI(index int) (ok bool)
	GetActorIRI(index int) (v url.URL)
}

// DeliverableObject is an object that is able to be sent to recipients via the
// "to", "bto", "cc", "bcc", and "audience" objects and/or links and/or IRIs.
type DeliverableObject interface {
	ActorObject
	ToLen() (l int)
	IsToObject(index int) (ok bool)
	GetToObject(index int) (v vocab.ObjectType)
	IsToLink(index int) (ok bool)
	GetToLink(index int) (v vocab.LinkType)
	IsToIRI(index int) (ok bool)
	GetToIRI(index int) (v url.URL)
	BtoLen() (l int)
	IsBtoObject(index int) (ok bool)
	GetBtoObject(index int) (v vocab.ObjectType)
	RemoveBtoObject(index int)
	IsBtoLink(index int) (ok bool)
	GetBtoLink(index int) (v vocab.LinkType)
	RemoveBtoLink(index int)
	IsBtoIRI(index int) (ok bool)
	GetBtoIRI(index int) (v url.URL)
	RemoveBtoIRI(index int)
	CcLen() (l int)
	IsCcObject(index int) (ok bool)
	GetCcObject(index int) (v vocab.ObjectType)
	IsCcLink(index int) (ok bool)
	GetCcLink(index int) (v vocab.LinkType)
	IsCcIRI(index int) (ok bool)
	GetCcIRI(index int) (v url.URL)
	BccLen() (l int)
	IsBccObject(index int) (ok bool)
	GetBccObject(index int) (v vocab.ObjectType)
	RemoveBccObject(index int)
	IsBccLink(index int) (ok bool)
	GetBccLink(index int) (v vocab.LinkType)
	RemoveBccLink(index int)
	IsBccIRI(index int) (ok bool)
	GetBccIRI(index int) (v url.URL)
	RemoveBccIRI(index int)
	AudienceLen() (l int)
	IsAudienceObject(index int) (ok bool)
	GetAudienceObject(index int) (v vocab.ObjectType)
	IsAudienceLink(index int) (ok bool)
	GetAudienceLink(index int) (v vocab.LinkType)
	IsAudienceIRI(index int) (ok bool)
	GetAudienceIRI(index int) (v url.URL)
}
