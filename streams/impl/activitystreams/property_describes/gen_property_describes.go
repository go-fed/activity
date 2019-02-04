package propertydescribes

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// DescribesProperty is the functional property "describes". It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type DescribesProperty struct {
	ObjectMember                vocab.ObjectInterface
	AcceptMember                vocab.AcceptInterface
	ActivityMember              vocab.ActivityInterface
	AddMember                   vocab.AddInterface
	AnnounceMember              vocab.AnnounceInterface
	ApplicationMember           vocab.ApplicationInterface
	ArriveMember                vocab.ArriveInterface
	ArticleMember               vocab.ArticleInterface
	AudioMember                 vocab.AudioInterface
	BlockMember                 vocab.BlockInterface
	CollectionMember            vocab.CollectionInterface
	CollectionPageMember        vocab.CollectionPageInterface
	CreateMember                vocab.CreateInterface
	DeleteMember                vocab.DeleteInterface
	DislikeMember               vocab.DislikeInterface
	DocumentMember              vocab.DocumentInterface
	EventMember                 vocab.EventInterface
	FlagMember                  vocab.FlagInterface
	FollowMember                vocab.FollowInterface
	GroupMember                 vocab.GroupInterface
	IgnoreMember                vocab.IgnoreInterface
	ImageMember                 vocab.ImageInterface
	IntransitiveActivityMember  vocab.IntransitiveActivityInterface
	InviteMember                vocab.InviteInterface
	JoinMember                  vocab.JoinInterface
	LeaveMember                 vocab.LeaveInterface
	LikeMember                  vocab.LikeInterface
	ListenMember                vocab.ListenInterface
	MoveMember                  vocab.MoveInterface
	NoteMember                  vocab.NoteInterface
	OfferMember                 vocab.OfferInterface
	OrderedCollectionMember     vocab.OrderedCollectionInterface
	OrderedCollectionPageMember vocab.OrderedCollectionPageInterface
	OrganizationMember          vocab.OrganizationInterface
	PageMember                  vocab.PageInterface
	PersonMember                vocab.PersonInterface
	PlaceMember                 vocab.PlaceInterface
	ProfileMember               vocab.ProfileInterface
	QuestionMember              vocab.QuestionInterface
	ReadMember                  vocab.ReadInterface
	RejectMember                vocab.RejectInterface
	RelationshipMember          vocab.RelationshipInterface
	RemoveMember                vocab.RemoveInterface
	ServiceMember               vocab.ServiceInterface
	TentativeAcceptMember       vocab.TentativeAcceptInterface
	TentativeRejectMember       vocab.TentativeRejectInterface
	TombstoneMember             vocab.TombstoneInterface
	TravelMember                vocab.TravelInterface
	UndoMember                  vocab.UndoInterface
	UpdateMember                vocab.UpdateInterface
	VideoMember                 vocab.VideoInterface
	ViewMember                  vocab.ViewInterface
	unknown                     []byte
	iri                         *url.URL
	alias                       string
}

// DeserializeDescribesProperty creates a "describes" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeDescribesProperty(m map[string]interface{}, aliasMap map[string]string) (*DescribesProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "describes"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "describes")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &DescribesProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ObjectMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeAcceptActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					AcceptMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeActivityActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ActivityMember: v,
					alias:          alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeAddActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					AddMember: v,
					alias:     alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeAnnounceActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					AnnounceMember: v,
					alias:          alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeApplicationActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ApplicationMember: v,
					alias:             alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeArriveActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ArriveMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeArticleActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ArticleMember: v,
					alias:         alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeAudioActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					AudioMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeBlockActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					BlockMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					CollectionMember: v,
					alias:            alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					CollectionPageMember: v,
					alias:                alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCreateActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					CreateMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeDeleteActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					DeleteMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeDislikeActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					DislikeMember: v,
					alias:         alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeDocumentActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					DocumentMember: v,
					alias:          alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeEventActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					EventMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeFlagActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					FlagMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeFollowActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					FollowMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeGroupActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					GroupMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeIgnoreActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					IgnoreMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ImageMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeIntransitiveActivityActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					IntransitiveActivityMember: v,
					alias:                      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeInviteActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					InviteMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeJoinActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					JoinMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeLeaveActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					LeaveMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeLikeActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					LikeMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeListenActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ListenMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeMoveActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					MoveMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeNoteActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					NoteMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOfferActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					OfferMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					OrderedCollectionMember: v,
					alias:                   alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					OrderedCollectionPageMember: v,
					alias:                       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrganizationActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					OrganizationMember: v,
					alias:              alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializePageActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					PageMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializePersonActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					PersonMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializePlaceActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					PlaceMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeProfileActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ProfileMember: v,
					alias:         alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeQuestionActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					QuestionMember: v,
					alias:          alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeReadActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ReadMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeRejectActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					RejectMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeRelationshipActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					RelationshipMember: v,
					alias:              alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeRemoveActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					RemoveMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeServiceActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ServiceMember: v,
					alias:         alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeTentativeAcceptActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					TentativeAcceptMember: v,
					alias:                 alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeTentativeRejectActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					TentativeRejectMember: v,
					alias:                 alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeTombstoneActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					TombstoneMember: v,
					alias:           alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeTravelActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					TravelMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeUndoActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					UndoMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeUpdateActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					UpdateMember: v,
					alias:        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeVideoActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					VideoMember: v,
					alias:       alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeViewActivityStreams()(m, aliasMap); err == nil {
				this := &DescribesProperty{
					ViewMember: v,
					alias:      alias,
				}
				return this, nil
			}
		} else if str, ok := i.(string); ok {
			this := &DescribesProperty{
				alias:   alias,
				unknown: []byte(str),
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "describes")
		}
	}
	return nil, nil
}

// NewDescribesProperty creates a new describes property.
func NewDescribesProperty() *DescribesProperty {
	return &DescribesProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *DescribesProperty) Clear() {
	this.ObjectMember = nil
	this.AcceptMember = nil
	this.ActivityMember = nil
	this.AddMember = nil
	this.AnnounceMember = nil
	this.ApplicationMember = nil
	this.ArriveMember = nil
	this.ArticleMember = nil
	this.AudioMember = nil
	this.BlockMember = nil
	this.CollectionMember = nil
	this.CollectionPageMember = nil
	this.CreateMember = nil
	this.DeleteMember = nil
	this.DislikeMember = nil
	this.DocumentMember = nil
	this.EventMember = nil
	this.FlagMember = nil
	this.FollowMember = nil
	this.GroupMember = nil
	this.IgnoreMember = nil
	this.ImageMember = nil
	this.IntransitiveActivityMember = nil
	this.InviteMember = nil
	this.JoinMember = nil
	this.LeaveMember = nil
	this.LikeMember = nil
	this.ListenMember = nil
	this.MoveMember = nil
	this.NoteMember = nil
	this.OfferMember = nil
	this.OrderedCollectionMember = nil
	this.OrderedCollectionPageMember = nil
	this.OrganizationMember = nil
	this.PageMember = nil
	this.PersonMember = nil
	this.PlaceMember = nil
	this.ProfileMember = nil
	this.QuestionMember = nil
	this.ReadMember = nil
	this.RejectMember = nil
	this.RelationshipMember = nil
	this.RemoveMember = nil
	this.ServiceMember = nil
	this.TentativeAcceptMember = nil
	this.TentativeRejectMember = nil
	this.TombstoneMember = nil
	this.TravelMember = nil
	this.UndoMember = nil
	this.UpdateMember = nil
	this.VideoMember = nil
	this.ViewMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetAccept returns the value of this property. When IsAccept returns false,
// GetAccept will return an arbitrary value.
func (this DescribesProperty) GetAccept() vocab.AcceptInterface {
	return this.AcceptMember
}

// GetActivity returns the value of this property. When IsActivity returns false,
// GetActivity will return an arbitrary value.
func (this DescribesProperty) GetActivity() vocab.ActivityInterface {
	return this.ActivityMember
}

// GetAdd returns the value of this property. When IsAdd returns false, GetAdd
// will return an arbitrary value.
func (this DescribesProperty) GetAdd() vocab.AddInterface {
	return this.AddMember
}

// GetAnnounce returns the value of this property. When IsAnnounce returns false,
// GetAnnounce will return an arbitrary value.
func (this DescribesProperty) GetAnnounce() vocab.AnnounceInterface {
	return this.AnnounceMember
}

// GetApplication returns the value of this property. When IsApplication returns
// false, GetApplication will return an arbitrary value.
func (this DescribesProperty) GetApplication() vocab.ApplicationInterface {
	return this.ApplicationMember
}

// GetArrive returns the value of this property. When IsArrive returns false,
// GetArrive will return an arbitrary value.
func (this DescribesProperty) GetArrive() vocab.ArriveInterface {
	return this.ArriveMember
}

// GetArticle returns the value of this property. When IsArticle returns false,
// GetArticle will return an arbitrary value.
func (this DescribesProperty) GetArticle() vocab.ArticleInterface {
	return this.ArticleMember
}

// GetAudio returns the value of this property. When IsAudio returns false,
// GetAudio will return an arbitrary value.
func (this DescribesProperty) GetAudio() vocab.AudioInterface {
	return this.AudioMember
}

// GetBlock returns the value of this property. When IsBlock returns false,
// GetBlock will return an arbitrary value.
func (this DescribesProperty) GetBlock() vocab.BlockInterface {
	return this.BlockMember
}

// GetCollection returns the value of this property. When IsCollection returns
// false, GetCollection will return an arbitrary value.
func (this DescribesProperty) GetCollection() vocab.CollectionInterface {
	return this.CollectionMember
}

// GetCollectionPage returns the value of this property. When IsCollectionPage
// returns false, GetCollectionPage will return an arbitrary value.
func (this DescribesProperty) GetCollectionPage() vocab.CollectionPageInterface {
	return this.CollectionPageMember
}

// GetCreate returns the value of this property. When IsCreate returns false,
// GetCreate will return an arbitrary value.
func (this DescribesProperty) GetCreate() vocab.CreateInterface {
	return this.CreateMember
}

// GetDelete returns the value of this property. When IsDelete returns false,
// GetDelete will return an arbitrary value.
func (this DescribesProperty) GetDelete() vocab.DeleteInterface {
	return this.DeleteMember
}

// GetDislike returns the value of this property. When IsDislike returns false,
// GetDislike will return an arbitrary value.
func (this DescribesProperty) GetDislike() vocab.DislikeInterface {
	return this.DislikeMember
}

// GetDocument returns the value of this property. When IsDocument returns false,
// GetDocument will return an arbitrary value.
func (this DescribesProperty) GetDocument() vocab.DocumentInterface {
	return this.DocumentMember
}

// GetEvent returns the value of this property. When IsEvent returns false,
// GetEvent will return an arbitrary value.
func (this DescribesProperty) GetEvent() vocab.EventInterface {
	return this.EventMember
}

// GetFlag returns the value of this property. When IsFlag returns false, GetFlag
// will return an arbitrary value.
func (this DescribesProperty) GetFlag() vocab.FlagInterface {
	return this.FlagMember
}

// GetFollow returns the value of this property. When IsFollow returns false,
// GetFollow will return an arbitrary value.
func (this DescribesProperty) GetFollow() vocab.FollowInterface {
	return this.FollowMember
}

// GetGroup returns the value of this property. When IsGroup returns false,
// GetGroup will return an arbitrary value.
func (this DescribesProperty) GetGroup() vocab.GroupInterface {
	return this.GroupMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this DescribesProperty) GetIRI() *url.URL {
	return this.iri
}

// GetIgnore returns the value of this property. When IsIgnore returns false,
// GetIgnore will return an arbitrary value.
func (this DescribesProperty) GetIgnore() vocab.IgnoreInterface {
	return this.IgnoreMember
}

// GetImage returns the value of this property. When IsImage returns false,
// GetImage will return an arbitrary value.
func (this DescribesProperty) GetImage() vocab.ImageInterface {
	return this.ImageMember
}

// GetIntransitiveActivity returns the value of this property. When
// IsIntransitiveActivity returns false, GetIntransitiveActivity will return
// an arbitrary value.
func (this DescribesProperty) GetIntransitiveActivity() vocab.IntransitiveActivityInterface {
	return this.IntransitiveActivityMember
}

// GetInvite returns the value of this property. When IsInvite returns false,
// GetInvite will return an arbitrary value.
func (this DescribesProperty) GetInvite() vocab.InviteInterface {
	return this.InviteMember
}

// GetJoin returns the value of this property. When IsJoin returns false, GetJoin
// will return an arbitrary value.
func (this DescribesProperty) GetJoin() vocab.JoinInterface {
	return this.JoinMember
}

// GetLeave returns the value of this property. When IsLeave returns false,
// GetLeave will return an arbitrary value.
func (this DescribesProperty) GetLeave() vocab.LeaveInterface {
	return this.LeaveMember
}

// GetLike returns the value of this property. When IsLike returns false, GetLike
// will return an arbitrary value.
func (this DescribesProperty) GetLike() vocab.LikeInterface {
	return this.LikeMember
}

// GetListen returns the value of this property. When IsListen returns false,
// GetListen will return an arbitrary value.
func (this DescribesProperty) GetListen() vocab.ListenInterface {
	return this.ListenMember
}

// GetMove returns the value of this property. When IsMove returns false, GetMove
// will return an arbitrary value.
func (this DescribesProperty) GetMove() vocab.MoveInterface {
	return this.MoveMember
}

// GetNote returns the value of this property. When IsNote returns false, GetNote
// will return an arbitrary value.
func (this DescribesProperty) GetNote() vocab.NoteInterface {
	return this.NoteMember
}

// GetObject returns the value of this property. When IsObject returns false,
// GetObject will return an arbitrary value.
func (this DescribesProperty) GetObject() vocab.ObjectInterface {
	return this.ObjectMember
}

// GetOffer returns the value of this property. When IsOffer returns false,
// GetOffer will return an arbitrary value.
func (this DescribesProperty) GetOffer() vocab.OfferInterface {
	return this.OfferMember
}

// GetOrderedCollection returns the value of this property. When
// IsOrderedCollection returns false, GetOrderedCollection will return an
// arbitrary value.
func (this DescribesProperty) GetOrderedCollection() vocab.OrderedCollectionInterface {
	return this.OrderedCollectionMember
}

// GetOrderedCollectionPage returns the value of this property. When
// IsOrderedCollectionPage returns false, GetOrderedCollectionPage will return
// an arbitrary value.
func (this DescribesProperty) GetOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return this.OrderedCollectionPageMember
}

// GetOrganization returns the value of this property. When IsOrganization returns
// false, GetOrganization will return an arbitrary value.
func (this DescribesProperty) GetOrganization() vocab.OrganizationInterface {
	return this.OrganizationMember
}

// GetPage returns the value of this property. When IsPage returns false, GetPage
// will return an arbitrary value.
func (this DescribesProperty) GetPage() vocab.PageInterface {
	return this.PageMember
}

// GetPerson returns the value of this property. When IsPerson returns false,
// GetPerson will return an arbitrary value.
func (this DescribesProperty) GetPerson() vocab.PersonInterface {
	return this.PersonMember
}

// GetPlace returns the value of this property. When IsPlace returns false,
// GetPlace will return an arbitrary value.
func (this DescribesProperty) GetPlace() vocab.PlaceInterface {
	return this.PlaceMember
}

// GetProfile returns the value of this property. When IsProfile returns false,
// GetProfile will return an arbitrary value.
func (this DescribesProperty) GetProfile() vocab.ProfileInterface {
	return this.ProfileMember
}

// GetQuestion returns the value of this property. When IsQuestion returns false,
// GetQuestion will return an arbitrary value.
func (this DescribesProperty) GetQuestion() vocab.QuestionInterface {
	return this.QuestionMember
}

// GetRead returns the value of this property. When IsRead returns false, GetRead
// will return an arbitrary value.
func (this DescribesProperty) GetRead() vocab.ReadInterface {
	return this.ReadMember
}

// GetReject returns the value of this property. When IsReject returns false,
// GetReject will return an arbitrary value.
func (this DescribesProperty) GetReject() vocab.RejectInterface {
	return this.RejectMember
}

// GetRelationship returns the value of this property. When IsRelationship returns
// false, GetRelationship will return an arbitrary value.
func (this DescribesProperty) GetRelationship() vocab.RelationshipInterface {
	return this.RelationshipMember
}

// GetRemove returns the value of this property. When IsRemove returns false,
// GetRemove will return an arbitrary value.
func (this DescribesProperty) GetRemove() vocab.RemoveInterface {
	return this.RemoveMember
}

// GetService returns the value of this property. When IsService returns false,
// GetService will return an arbitrary value.
func (this DescribesProperty) GetService() vocab.ServiceInterface {
	return this.ServiceMember
}

// GetTentativeAccept returns the value of this property. When IsTentativeAccept
// returns false, GetTentativeAccept will return an arbitrary value.
func (this DescribesProperty) GetTentativeAccept() vocab.TentativeAcceptInterface {
	return this.TentativeAcceptMember
}

// GetTentativeReject returns the value of this property. When IsTentativeReject
// returns false, GetTentativeReject will return an arbitrary value.
func (this DescribesProperty) GetTentativeReject() vocab.TentativeRejectInterface {
	return this.TentativeRejectMember
}

// GetTombstone returns the value of this property. When IsTombstone returns
// false, GetTombstone will return an arbitrary value.
func (this DescribesProperty) GetTombstone() vocab.TombstoneInterface {
	return this.TombstoneMember
}

// GetTravel returns the value of this property. When IsTravel returns false,
// GetTravel will return an arbitrary value.
func (this DescribesProperty) GetTravel() vocab.TravelInterface {
	return this.TravelMember
}

// GetUndo returns the value of this property. When IsUndo returns false, GetUndo
// will return an arbitrary value.
func (this DescribesProperty) GetUndo() vocab.UndoInterface {
	return this.UndoMember
}

// GetUpdate returns the value of this property. When IsUpdate returns false,
// GetUpdate will return an arbitrary value.
func (this DescribesProperty) GetUpdate() vocab.UpdateInterface {
	return this.UpdateMember
}

// GetVideo returns the value of this property. When IsVideo returns false,
// GetVideo will return an arbitrary value.
func (this DescribesProperty) GetVideo() vocab.VideoInterface {
	return this.VideoMember
}

// GetView returns the value of this property. When IsView returns false, GetView
// will return an arbitrary value.
func (this DescribesProperty) GetView() vocab.ViewInterface {
	return this.ViewMember
}

// HasAny returns true if any of the different values is set.
func (this DescribesProperty) HasAny() bool {
	return this.IsObject() ||
		this.IsAccept() ||
		this.IsActivity() ||
		this.IsAdd() ||
		this.IsAnnounce() ||
		this.IsApplication() ||
		this.IsArrive() ||
		this.IsArticle() ||
		this.IsAudio() ||
		this.IsBlock() ||
		this.IsCollection() ||
		this.IsCollectionPage() ||
		this.IsCreate() ||
		this.IsDelete() ||
		this.IsDislike() ||
		this.IsDocument() ||
		this.IsEvent() ||
		this.IsFlag() ||
		this.IsFollow() ||
		this.IsGroup() ||
		this.IsIgnore() ||
		this.IsImage() ||
		this.IsIntransitiveActivity() ||
		this.IsInvite() ||
		this.IsJoin() ||
		this.IsLeave() ||
		this.IsLike() ||
		this.IsListen() ||
		this.IsMove() ||
		this.IsNote() ||
		this.IsOffer() ||
		this.IsOrderedCollection() ||
		this.IsOrderedCollectionPage() ||
		this.IsOrganization() ||
		this.IsPage() ||
		this.IsPerson() ||
		this.IsPlace() ||
		this.IsProfile() ||
		this.IsQuestion() ||
		this.IsRead() ||
		this.IsReject() ||
		this.IsRelationship() ||
		this.IsRemove() ||
		this.IsService() ||
		this.IsTentativeAccept() ||
		this.IsTentativeReject() ||
		this.IsTombstone() ||
		this.IsTravel() ||
		this.IsUndo() ||
		this.IsUpdate() ||
		this.IsVideo() ||
		this.IsView() ||
		this.iri != nil
}

// IsAccept returns true if this property has a type of "Accept". When true, use
// the GetAccept and SetAccept methods to access and set this property.
func (this DescribesProperty) IsAccept() bool {
	return this.AcceptMember != nil
}

// IsActivity returns true if this property has a type of "Activity". When true,
// use the GetActivity and SetActivity methods to access and set this property.
func (this DescribesProperty) IsActivity() bool {
	return this.ActivityMember != nil
}

// IsAdd returns true if this property has a type of "Add". When true, use the
// GetAdd and SetAdd methods to access and set this property.
func (this DescribesProperty) IsAdd() bool {
	return this.AddMember != nil
}

// IsAnnounce returns true if this property has a type of "Announce". When true,
// use the GetAnnounce and SetAnnounce methods to access and set this property.
func (this DescribesProperty) IsAnnounce() bool {
	return this.AnnounceMember != nil
}

// IsApplication returns true if this property has a type of "Application". When
// true, use the GetApplication and SetApplication methods to access and set
// this property.
func (this DescribesProperty) IsApplication() bool {
	return this.ApplicationMember != nil
}

// IsArrive returns true if this property has a type of "Arrive". When true, use
// the GetArrive and SetArrive methods to access and set this property.
func (this DescribesProperty) IsArrive() bool {
	return this.ArriveMember != nil
}

// IsArticle returns true if this property has a type of "Article". When true, use
// the GetArticle and SetArticle methods to access and set this property.
func (this DescribesProperty) IsArticle() bool {
	return this.ArticleMember != nil
}

// IsAudio returns true if this property has a type of "Audio". When true, use the
// GetAudio and SetAudio methods to access and set this property.
func (this DescribesProperty) IsAudio() bool {
	return this.AudioMember != nil
}

// IsBlock returns true if this property has a type of "Block". When true, use the
// GetBlock and SetBlock methods to access and set this property.
func (this DescribesProperty) IsBlock() bool {
	return this.BlockMember != nil
}

// IsCollection returns true if this property has a type of "Collection". When
// true, use the GetCollection and SetCollection methods to access and set
// this property.
func (this DescribesProperty) IsCollection() bool {
	return this.CollectionMember != nil
}

// IsCollectionPage returns true if this property has a type of "CollectionPage".
// When true, use the GetCollectionPage and SetCollectionPage methods to
// access and set this property.
func (this DescribesProperty) IsCollectionPage() bool {
	return this.CollectionPageMember != nil
}

// IsCreate returns true if this property has a type of "Create". When true, use
// the GetCreate and SetCreate methods to access and set this property.
func (this DescribesProperty) IsCreate() bool {
	return this.CreateMember != nil
}

// IsDelete returns true if this property has a type of "Delete". When true, use
// the GetDelete and SetDelete methods to access and set this property.
func (this DescribesProperty) IsDelete() bool {
	return this.DeleteMember != nil
}

// IsDislike returns true if this property has a type of "Dislike". When true, use
// the GetDislike and SetDislike methods to access and set this property.
func (this DescribesProperty) IsDislike() bool {
	return this.DislikeMember != nil
}

// IsDocument returns true if this property has a type of "Document". When true,
// use the GetDocument and SetDocument methods to access and set this property.
func (this DescribesProperty) IsDocument() bool {
	return this.DocumentMember != nil
}

// IsEvent returns true if this property has a type of "Event". When true, use the
// GetEvent and SetEvent methods to access and set this property.
func (this DescribesProperty) IsEvent() bool {
	return this.EventMember != nil
}

// IsFlag returns true if this property has a type of "Flag". When true, use the
// GetFlag and SetFlag methods to access and set this property.
func (this DescribesProperty) IsFlag() bool {
	return this.FlagMember != nil
}

// IsFollow returns true if this property has a type of "Follow". When true, use
// the GetFollow and SetFollow methods to access and set this property.
func (this DescribesProperty) IsFollow() bool {
	return this.FollowMember != nil
}

// IsGroup returns true if this property has a type of "Group". When true, use the
// GetGroup and SetGroup methods to access and set this property.
func (this DescribesProperty) IsGroup() bool {
	return this.GroupMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this DescribesProperty) IsIRI() bool {
	return this.iri != nil
}

// IsIgnore returns true if this property has a type of "Ignore". When true, use
// the GetIgnore and SetIgnore methods to access and set this property.
func (this DescribesProperty) IsIgnore() bool {
	return this.IgnoreMember != nil
}

// IsImage returns true if this property has a type of "Image". When true, use the
// GetImage and SetImage methods to access and set this property.
func (this DescribesProperty) IsImage() bool {
	return this.ImageMember != nil
}

// IsIntransitiveActivity returns true if this property has a type of
// "IntransitiveActivity". When true, use the GetIntransitiveActivity and
// SetIntransitiveActivity methods to access and set this property.
func (this DescribesProperty) IsIntransitiveActivity() bool {
	return this.IntransitiveActivityMember != nil
}

// IsInvite returns true if this property has a type of "Invite". When true, use
// the GetInvite and SetInvite methods to access and set this property.
func (this DescribesProperty) IsInvite() bool {
	return this.InviteMember != nil
}

// IsJoin returns true if this property has a type of "Join". When true, use the
// GetJoin and SetJoin methods to access and set this property.
func (this DescribesProperty) IsJoin() bool {
	return this.JoinMember != nil
}

// IsLeave returns true if this property has a type of "Leave". When true, use the
// GetLeave and SetLeave methods to access and set this property.
func (this DescribesProperty) IsLeave() bool {
	return this.LeaveMember != nil
}

// IsLike returns true if this property has a type of "Like". When true, use the
// GetLike and SetLike methods to access and set this property.
func (this DescribesProperty) IsLike() bool {
	return this.LikeMember != nil
}

// IsListen returns true if this property has a type of "Listen". When true, use
// the GetListen and SetListen methods to access and set this property.
func (this DescribesProperty) IsListen() bool {
	return this.ListenMember != nil
}

// IsMove returns true if this property has a type of "Move". When true, use the
// GetMove and SetMove methods to access and set this property.
func (this DescribesProperty) IsMove() bool {
	return this.MoveMember != nil
}

// IsNote returns true if this property has a type of "Note". When true, use the
// GetNote and SetNote methods to access and set this property.
func (this DescribesProperty) IsNote() bool {
	return this.NoteMember != nil
}

// IsObject returns true if this property has a type of "Object". When true, use
// the GetObject and SetObject methods to access and set this property.
func (this DescribesProperty) IsObject() bool {
	return this.ObjectMember != nil
}

// IsOffer returns true if this property has a type of "Offer". When true, use the
// GetOffer and SetOffer methods to access and set this property.
func (this DescribesProperty) IsOffer() bool {
	return this.OfferMember != nil
}

// IsOrderedCollection returns true if this property has a type of
// "OrderedCollection". When true, use the GetOrderedCollection and
// SetOrderedCollection methods to access and set this property.
func (this DescribesProperty) IsOrderedCollection() bool {
	return this.OrderedCollectionMember != nil
}

// IsOrderedCollectionPage returns true if this property has a type of
// "OrderedCollectionPage". When true, use the GetOrderedCollectionPage and
// SetOrderedCollectionPage methods to access and set this property.
func (this DescribesProperty) IsOrderedCollectionPage() bool {
	return this.OrderedCollectionPageMember != nil
}

// IsOrganization returns true if this property has a type of "Organization". When
// true, use the GetOrganization and SetOrganization methods to access and set
// this property.
func (this DescribesProperty) IsOrganization() bool {
	return this.OrganizationMember != nil
}

// IsPage returns true if this property has a type of "Page". When true, use the
// GetPage and SetPage methods to access and set this property.
func (this DescribesProperty) IsPage() bool {
	return this.PageMember != nil
}

// IsPerson returns true if this property has a type of "Person". When true, use
// the GetPerson and SetPerson methods to access and set this property.
func (this DescribesProperty) IsPerson() bool {
	return this.PersonMember != nil
}

// IsPlace returns true if this property has a type of "Place". When true, use the
// GetPlace and SetPlace methods to access and set this property.
func (this DescribesProperty) IsPlace() bool {
	return this.PlaceMember != nil
}

// IsProfile returns true if this property has a type of "Profile". When true, use
// the GetProfile and SetProfile methods to access and set this property.
func (this DescribesProperty) IsProfile() bool {
	return this.ProfileMember != nil
}

// IsQuestion returns true if this property has a type of "Question". When true,
// use the GetQuestion and SetQuestion methods to access and set this property.
func (this DescribesProperty) IsQuestion() bool {
	return this.QuestionMember != nil
}

// IsRead returns true if this property has a type of "Read". When true, use the
// GetRead and SetRead methods to access and set this property.
func (this DescribesProperty) IsRead() bool {
	return this.ReadMember != nil
}

// IsReject returns true if this property has a type of "Reject". When true, use
// the GetReject and SetReject methods to access and set this property.
func (this DescribesProperty) IsReject() bool {
	return this.RejectMember != nil
}

// IsRelationship returns true if this property has a type of "Relationship". When
// true, use the GetRelationship and SetRelationship methods to access and set
// this property.
func (this DescribesProperty) IsRelationship() bool {
	return this.RelationshipMember != nil
}

// IsRemove returns true if this property has a type of "Remove". When true, use
// the GetRemove and SetRemove methods to access and set this property.
func (this DescribesProperty) IsRemove() bool {
	return this.RemoveMember != nil
}

// IsService returns true if this property has a type of "Service". When true, use
// the GetService and SetService methods to access and set this property.
func (this DescribesProperty) IsService() bool {
	return this.ServiceMember != nil
}

// IsTentativeAccept returns true if this property has a type of
// "TentativeAccept". When true, use the GetTentativeAccept and
// SetTentativeAccept methods to access and set this property.
func (this DescribesProperty) IsTentativeAccept() bool {
	return this.TentativeAcceptMember != nil
}

// IsTentativeReject returns true if this property has a type of
// "TentativeReject". When true, use the GetTentativeReject and
// SetTentativeReject methods to access and set this property.
func (this DescribesProperty) IsTentativeReject() bool {
	return this.TentativeRejectMember != nil
}

// IsTombstone returns true if this property has a type of "Tombstone". When true,
// use the GetTombstone and SetTombstone methods to access and set this
// property.
func (this DescribesProperty) IsTombstone() bool {
	return this.TombstoneMember != nil
}

// IsTravel returns true if this property has a type of "Travel". When true, use
// the GetTravel and SetTravel methods to access and set this property.
func (this DescribesProperty) IsTravel() bool {
	return this.TravelMember != nil
}

// IsUndo returns true if this property has a type of "Undo". When true, use the
// GetUndo and SetUndo methods to access and set this property.
func (this DescribesProperty) IsUndo() bool {
	return this.UndoMember != nil
}

// IsUpdate returns true if this property has a type of "Update". When true, use
// the GetUpdate and SetUpdate methods to access and set this property.
func (this DescribesProperty) IsUpdate() bool {
	return this.UpdateMember != nil
}

// IsVideo returns true if this property has a type of "Video". When true, use the
// GetVideo and SetVideo methods to access and set this property.
func (this DescribesProperty) IsVideo() bool {
	return this.VideoMember != nil
}

// IsView returns true if this property has a type of "View". When true, use the
// GetView and SetView methods to access and set this property.
func (this DescribesProperty) IsView() bool {
	return this.ViewMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this DescribesProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsObject() {
		child = this.GetObject().JSONLDContext()
	} else if this.IsAccept() {
		child = this.GetAccept().JSONLDContext()
	} else if this.IsActivity() {
		child = this.GetActivity().JSONLDContext()
	} else if this.IsAdd() {
		child = this.GetAdd().JSONLDContext()
	} else if this.IsAnnounce() {
		child = this.GetAnnounce().JSONLDContext()
	} else if this.IsApplication() {
		child = this.GetApplication().JSONLDContext()
	} else if this.IsArrive() {
		child = this.GetArrive().JSONLDContext()
	} else if this.IsArticle() {
		child = this.GetArticle().JSONLDContext()
	} else if this.IsAudio() {
		child = this.GetAudio().JSONLDContext()
	} else if this.IsBlock() {
		child = this.GetBlock().JSONLDContext()
	} else if this.IsCollection() {
		child = this.GetCollection().JSONLDContext()
	} else if this.IsCollectionPage() {
		child = this.GetCollectionPage().JSONLDContext()
	} else if this.IsCreate() {
		child = this.GetCreate().JSONLDContext()
	} else if this.IsDelete() {
		child = this.GetDelete().JSONLDContext()
	} else if this.IsDislike() {
		child = this.GetDislike().JSONLDContext()
	} else if this.IsDocument() {
		child = this.GetDocument().JSONLDContext()
	} else if this.IsEvent() {
		child = this.GetEvent().JSONLDContext()
	} else if this.IsFlag() {
		child = this.GetFlag().JSONLDContext()
	} else if this.IsFollow() {
		child = this.GetFollow().JSONLDContext()
	} else if this.IsGroup() {
		child = this.GetGroup().JSONLDContext()
	} else if this.IsIgnore() {
		child = this.GetIgnore().JSONLDContext()
	} else if this.IsImage() {
		child = this.GetImage().JSONLDContext()
	} else if this.IsIntransitiveActivity() {
		child = this.GetIntransitiveActivity().JSONLDContext()
	} else if this.IsInvite() {
		child = this.GetInvite().JSONLDContext()
	} else if this.IsJoin() {
		child = this.GetJoin().JSONLDContext()
	} else if this.IsLeave() {
		child = this.GetLeave().JSONLDContext()
	} else if this.IsLike() {
		child = this.GetLike().JSONLDContext()
	} else if this.IsListen() {
		child = this.GetListen().JSONLDContext()
	} else if this.IsMove() {
		child = this.GetMove().JSONLDContext()
	} else if this.IsNote() {
		child = this.GetNote().JSONLDContext()
	} else if this.IsOffer() {
		child = this.GetOffer().JSONLDContext()
	} else if this.IsOrderedCollection() {
		child = this.GetOrderedCollection().JSONLDContext()
	} else if this.IsOrderedCollectionPage() {
		child = this.GetOrderedCollectionPage().JSONLDContext()
	} else if this.IsOrganization() {
		child = this.GetOrganization().JSONLDContext()
	} else if this.IsPage() {
		child = this.GetPage().JSONLDContext()
	} else if this.IsPerson() {
		child = this.GetPerson().JSONLDContext()
	} else if this.IsPlace() {
		child = this.GetPlace().JSONLDContext()
	} else if this.IsProfile() {
		child = this.GetProfile().JSONLDContext()
	} else if this.IsQuestion() {
		child = this.GetQuestion().JSONLDContext()
	} else if this.IsRead() {
		child = this.GetRead().JSONLDContext()
	} else if this.IsReject() {
		child = this.GetReject().JSONLDContext()
	} else if this.IsRelationship() {
		child = this.GetRelationship().JSONLDContext()
	} else if this.IsRemove() {
		child = this.GetRemove().JSONLDContext()
	} else if this.IsService() {
		child = this.GetService().JSONLDContext()
	} else if this.IsTentativeAccept() {
		child = this.GetTentativeAccept().JSONLDContext()
	} else if this.IsTentativeReject() {
		child = this.GetTentativeReject().JSONLDContext()
	} else if this.IsTombstone() {
		child = this.GetTombstone().JSONLDContext()
	} else if this.IsTravel() {
		child = this.GetTravel().JSONLDContext()
	} else if this.IsUndo() {
		child = this.GetUndo().JSONLDContext()
	} else if this.IsUpdate() {
		child = this.GetUpdate().JSONLDContext()
	} else if this.IsVideo() {
		child = this.GetVideo().JSONLDContext()
	} else if this.IsView() {
		child = this.GetView().JSONLDContext()
	}
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
func (this DescribesProperty) KindIndex() int {
	if this.IsObject() {
		return 0
	}
	if this.IsAccept() {
		return 1
	}
	if this.IsActivity() {
		return 2
	}
	if this.IsAdd() {
		return 3
	}
	if this.IsAnnounce() {
		return 4
	}
	if this.IsApplication() {
		return 5
	}
	if this.IsArrive() {
		return 6
	}
	if this.IsArticle() {
		return 7
	}
	if this.IsAudio() {
		return 8
	}
	if this.IsBlock() {
		return 9
	}
	if this.IsCollection() {
		return 10
	}
	if this.IsCollectionPage() {
		return 11
	}
	if this.IsCreate() {
		return 12
	}
	if this.IsDelete() {
		return 13
	}
	if this.IsDislike() {
		return 14
	}
	if this.IsDocument() {
		return 15
	}
	if this.IsEvent() {
		return 16
	}
	if this.IsFlag() {
		return 17
	}
	if this.IsFollow() {
		return 18
	}
	if this.IsGroup() {
		return 19
	}
	if this.IsIgnore() {
		return 20
	}
	if this.IsImage() {
		return 21
	}
	if this.IsIntransitiveActivity() {
		return 22
	}
	if this.IsInvite() {
		return 23
	}
	if this.IsJoin() {
		return 24
	}
	if this.IsLeave() {
		return 25
	}
	if this.IsLike() {
		return 26
	}
	if this.IsListen() {
		return 27
	}
	if this.IsMove() {
		return 28
	}
	if this.IsNote() {
		return 29
	}
	if this.IsOffer() {
		return 30
	}
	if this.IsOrderedCollection() {
		return 31
	}
	if this.IsOrderedCollectionPage() {
		return 32
	}
	if this.IsOrganization() {
		return 33
	}
	if this.IsPage() {
		return 34
	}
	if this.IsPerson() {
		return 35
	}
	if this.IsPlace() {
		return 36
	}
	if this.IsProfile() {
		return 37
	}
	if this.IsQuestion() {
		return 38
	}
	if this.IsRead() {
		return 39
	}
	if this.IsReject() {
		return 40
	}
	if this.IsRelationship() {
		return 41
	}
	if this.IsRemove() {
		return 42
	}
	if this.IsService() {
		return 43
	}
	if this.IsTentativeAccept() {
		return 44
	}
	if this.IsTentativeReject() {
		return 45
	}
	if this.IsTombstone() {
		return 46
	}
	if this.IsTravel() {
		return 47
	}
	if this.IsUndo() {
		return 48
	}
	if this.IsUpdate() {
		return 49
	}
	if this.IsVideo() {
		return 50
	}
	if this.IsView() {
		return 51
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
func (this DescribesProperty) LessThan(o vocab.DescribesPropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsObject() {
		return this.GetObject().LessThan(o.GetObject())
	} else if this.IsAccept() {
		return this.GetAccept().LessThan(o.GetAccept())
	} else if this.IsActivity() {
		return this.GetActivity().LessThan(o.GetActivity())
	} else if this.IsAdd() {
		return this.GetAdd().LessThan(o.GetAdd())
	} else if this.IsAnnounce() {
		return this.GetAnnounce().LessThan(o.GetAnnounce())
	} else if this.IsApplication() {
		return this.GetApplication().LessThan(o.GetApplication())
	} else if this.IsArrive() {
		return this.GetArrive().LessThan(o.GetArrive())
	} else if this.IsArticle() {
		return this.GetArticle().LessThan(o.GetArticle())
	} else if this.IsAudio() {
		return this.GetAudio().LessThan(o.GetAudio())
	} else if this.IsBlock() {
		return this.GetBlock().LessThan(o.GetBlock())
	} else if this.IsCollection() {
		return this.GetCollection().LessThan(o.GetCollection())
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().LessThan(o.GetCollectionPage())
	} else if this.IsCreate() {
		return this.GetCreate().LessThan(o.GetCreate())
	} else if this.IsDelete() {
		return this.GetDelete().LessThan(o.GetDelete())
	} else if this.IsDislike() {
		return this.GetDislike().LessThan(o.GetDislike())
	} else if this.IsDocument() {
		return this.GetDocument().LessThan(o.GetDocument())
	} else if this.IsEvent() {
		return this.GetEvent().LessThan(o.GetEvent())
	} else if this.IsFlag() {
		return this.GetFlag().LessThan(o.GetFlag())
	} else if this.IsFollow() {
		return this.GetFollow().LessThan(o.GetFollow())
	} else if this.IsGroup() {
		return this.GetGroup().LessThan(o.GetGroup())
	} else if this.IsIgnore() {
		return this.GetIgnore().LessThan(o.GetIgnore())
	} else if this.IsImage() {
		return this.GetImage().LessThan(o.GetImage())
	} else if this.IsIntransitiveActivity() {
		return this.GetIntransitiveActivity().LessThan(o.GetIntransitiveActivity())
	} else if this.IsInvite() {
		return this.GetInvite().LessThan(o.GetInvite())
	} else if this.IsJoin() {
		return this.GetJoin().LessThan(o.GetJoin())
	} else if this.IsLeave() {
		return this.GetLeave().LessThan(o.GetLeave())
	} else if this.IsLike() {
		return this.GetLike().LessThan(o.GetLike())
	} else if this.IsListen() {
		return this.GetListen().LessThan(o.GetListen())
	} else if this.IsMove() {
		return this.GetMove().LessThan(o.GetMove())
	} else if this.IsNote() {
		return this.GetNote().LessThan(o.GetNote())
	} else if this.IsOffer() {
		return this.GetOffer().LessThan(o.GetOffer())
	} else if this.IsOrderedCollection() {
		return this.GetOrderedCollection().LessThan(o.GetOrderedCollection())
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().LessThan(o.GetOrderedCollectionPage())
	} else if this.IsOrganization() {
		return this.GetOrganization().LessThan(o.GetOrganization())
	} else if this.IsPage() {
		return this.GetPage().LessThan(o.GetPage())
	} else if this.IsPerson() {
		return this.GetPerson().LessThan(o.GetPerson())
	} else if this.IsPlace() {
		return this.GetPlace().LessThan(o.GetPlace())
	} else if this.IsProfile() {
		return this.GetProfile().LessThan(o.GetProfile())
	} else if this.IsQuestion() {
		return this.GetQuestion().LessThan(o.GetQuestion())
	} else if this.IsRead() {
		return this.GetRead().LessThan(o.GetRead())
	} else if this.IsReject() {
		return this.GetReject().LessThan(o.GetReject())
	} else if this.IsRelationship() {
		return this.GetRelationship().LessThan(o.GetRelationship())
	} else if this.IsRemove() {
		return this.GetRemove().LessThan(o.GetRemove())
	} else if this.IsService() {
		return this.GetService().LessThan(o.GetService())
	} else if this.IsTentativeAccept() {
		return this.GetTentativeAccept().LessThan(o.GetTentativeAccept())
	} else if this.IsTentativeReject() {
		return this.GetTentativeReject().LessThan(o.GetTentativeReject())
	} else if this.IsTombstone() {
		return this.GetTombstone().LessThan(o.GetTombstone())
	} else if this.IsTravel() {
		return this.GetTravel().LessThan(o.GetTravel())
	} else if this.IsUndo() {
		return this.GetUndo().LessThan(o.GetUndo())
	} else if this.IsUpdate() {
		return this.GetUpdate().LessThan(o.GetUpdate())
	} else if this.IsVideo() {
		return this.GetVideo().LessThan(o.GetVideo())
	} else if this.IsView() {
		return this.GetView().LessThan(o.GetView())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "describes".
func (this DescribesProperty) Name() string {
	return "describes"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this DescribesProperty) Serialize() (interface{}, error) {
	if this.IsObject() {
		return this.GetObject().Serialize()
	} else if this.IsAccept() {
		return this.GetAccept().Serialize()
	} else if this.IsActivity() {
		return this.GetActivity().Serialize()
	} else if this.IsAdd() {
		return this.GetAdd().Serialize()
	} else if this.IsAnnounce() {
		return this.GetAnnounce().Serialize()
	} else if this.IsApplication() {
		return this.GetApplication().Serialize()
	} else if this.IsArrive() {
		return this.GetArrive().Serialize()
	} else if this.IsArticle() {
		return this.GetArticle().Serialize()
	} else if this.IsAudio() {
		return this.GetAudio().Serialize()
	} else if this.IsBlock() {
		return this.GetBlock().Serialize()
	} else if this.IsCollection() {
		return this.GetCollection().Serialize()
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().Serialize()
	} else if this.IsCreate() {
		return this.GetCreate().Serialize()
	} else if this.IsDelete() {
		return this.GetDelete().Serialize()
	} else if this.IsDislike() {
		return this.GetDislike().Serialize()
	} else if this.IsDocument() {
		return this.GetDocument().Serialize()
	} else if this.IsEvent() {
		return this.GetEvent().Serialize()
	} else if this.IsFlag() {
		return this.GetFlag().Serialize()
	} else if this.IsFollow() {
		return this.GetFollow().Serialize()
	} else if this.IsGroup() {
		return this.GetGroup().Serialize()
	} else if this.IsIgnore() {
		return this.GetIgnore().Serialize()
	} else if this.IsImage() {
		return this.GetImage().Serialize()
	} else if this.IsIntransitiveActivity() {
		return this.GetIntransitiveActivity().Serialize()
	} else if this.IsInvite() {
		return this.GetInvite().Serialize()
	} else if this.IsJoin() {
		return this.GetJoin().Serialize()
	} else if this.IsLeave() {
		return this.GetLeave().Serialize()
	} else if this.IsLike() {
		return this.GetLike().Serialize()
	} else if this.IsListen() {
		return this.GetListen().Serialize()
	} else if this.IsMove() {
		return this.GetMove().Serialize()
	} else if this.IsNote() {
		return this.GetNote().Serialize()
	} else if this.IsOffer() {
		return this.GetOffer().Serialize()
	} else if this.IsOrderedCollection() {
		return this.GetOrderedCollection().Serialize()
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().Serialize()
	} else if this.IsOrganization() {
		return this.GetOrganization().Serialize()
	} else if this.IsPage() {
		return this.GetPage().Serialize()
	} else if this.IsPerson() {
		return this.GetPerson().Serialize()
	} else if this.IsPlace() {
		return this.GetPlace().Serialize()
	} else if this.IsProfile() {
		return this.GetProfile().Serialize()
	} else if this.IsQuestion() {
		return this.GetQuestion().Serialize()
	} else if this.IsRead() {
		return this.GetRead().Serialize()
	} else if this.IsReject() {
		return this.GetReject().Serialize()
	} else if this.IsRelationship() {
		return this.GetRelationship().Serialize()
	} else if this.IsRemove() {
		return this.GetRemove().Serialize()
	} else if this.IsService() {
		return this.GetService().Serialize()
	} else if this.IsTentativeAccept() {
		return this.GetTentativeAccept().Serialize()
	} else if this.IsTentativeReject() {
		return this.GetTentativeReject().Serialize()
	} else if this.IsTombstone() {
		return this.GetTombstone().Serialize()
	} else if this.IsTravel() {
		return this.GetTravel().Serialize()
	} else if this.IsUndo() {
		return this.GetUndo().Serialize()
	} else if this.IsUpdate() {
		return this.GetUpdate().Serialize()
	} else if this.IsVideo() {
		return this.GetVideo().Serialize()
	} else if this.IsView() {
		return this.GetView().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return string(this.unknown), nil
}

// SetAccept sets the value of this property. Calling IsAccept afterwards returns
// true.
func (this *DescribesProperty) SetAccept(v vocab.AcceptInterface) {
	this.Clear()
	this.AcceptMember = v
}

// SetActivity sets the value of this property. Calling IsActivity afterwards
// returns true.
func (this *DescribesProperty) SetActivity(v vocab.ActivityInterface) {
	this.Clear()
	this.ActivityMember = v
}

// SetAdd sets the value of this property. Calling IsAdd afterwards returns true.
func (this *DescribesProperty) SetAdd(v vocab.AddInterface) {
	this.Clear()
	this.AddMember = v
}

// SetAnnounce sets the value of this property. Calling IsAnnounce afterwards
// returns true.
func (this *DescribesProperty) SetAnnounce(v vocab.AnnounceInterface) {
	this.Clear()
	this.AnnounceMember = v
}

// SetApplication sets the value of this property. Calling IsApplication
// afterwards returns true.
func (this *DescribesProperty) SetApplication(v vocab.ApplicationInterface) {
	this.Clear()
	this.ApplicationMember = v
}

// SetArrive sets the value of this property. Calling IsArrive afterwards returns
// true.
func (this *DescribesProperty) SetArrive(v vocab.ArriveInterface) {
	this.Clear()
	this.ArriveMember = v
}

// SetArticle sets the value of this property. Calling IsArticle afterwards
// returns true.
func (this *DescribesProperty) SetArticle(v vocab.ArticleInterface) {
	this.Clear()
	this.ArticleMember = v
}

// SetAudio sets the value of this property. Calling IsAudio afterwards returns
// true.
func (this *DescribesProperty) SetAudio(v vocab.AudioInterface) {
	this.Clear()
	this.AudioMember = v
}

// SetBlock sets the value of this property. Calling IsBlock afterwards returns
// true.
func (this *DescribesProperty) SetBlock(v vocab.BlockInterface) {
	this.Clear()
	this.BlockMember = v
}

// SetCollection sets the value of this property. Calling IsCollection afterwards
// returns true.
func (this *DescribesProperty) SetCollection(v vocab.CollectionInterface) {
	this.Clear()
	this.CollectionMember = v
}

// SetCollectionPage sets the value of this property. Calling IsCollectionPage
// afterwards returns true.
func (this *DescribesProperty) SetCollectionPage(v vocab.CollectionPageInterface) {
	this.Clear()
	this.CollectionPageMember = v
}

// SetCreate sets the value of this property. Calling IsCreate afterwards returns
// true.
func (this *DescribesProperty) SetCreate(v vocab.CreateInterface) {
	this.Clear()
	this.CreateMember = v
}

// SetDelete sets the value of this property. Calling IsDelete afterwards returns
// true.
func (this *DescribesProperty) SetDelete(v vocab.DeleteInterface) {
	this.Clear()
	this.DeleteMember = v
}

// SetDislike sets the value of this property. Calling IsDislike afterwards
// returns true.
func (this *DescribesProperty) SetDislike(v vocab.DislikeInterface) {
	this.Clear()
	this.DislikeMember = v
}

// SetDocument sets the value of this property. Calling IsDocument afterwards
// returns true.
func (this *DescribesProperty) SetDocument(v vocab.DocumentInterface) {
	this.Clear()
	this.DocumentMember = v
}

// SetEvent sets the value of this property. Calling IsEvent afterwards returns
// true.
func (this *DescribesProperty) SetEvent(v vocab.EventInterface) {
	this.Clear()
	this.EventMember = v
}

// SetFlag sets the value of this property. Calling IsFlag afterwards returns true.
func (this *DescribesProperty) SetFlag(v vocab.FlagInterface) {
	this.Clear()
	this.FlagMember = v
}

// SetFollow sets the value of this property. Calling IsFollow afterwards returns
// true.
func (this *DescribesProperty) SetFollow(v vocab.FollowInterface) {
	this.Clear()
	this.FollowMember = v
}

// SetGroup sets the value of this property. Calling IsGroup afterwards returns
// true.
func (this *DescribesProperty) SetGroup(v vocab.GroupInterface) {
	this.Clear()
	this.GroupMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *DescribesProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetIgnore sets the value of this property. Calling IsIgnore afterwards returns
// true.
func (this *DescribesProperty) SetIgnore(v vocab.IgnoreInterface) {
	this.Clear()
	this.IgnoreMember = v
}

// SetImage sets the value of this property. Calling IsImage afterwards returns
// true.
func (this *DescribesProperty) SetImage(v vocab.ImageInterface) {
	this.Clear()
	this.ImageMember = v
}

// SetIntransitiveActivity sets the value of this property. Calling
// IsIntransitiveActivity afterwards returns true.
func (this *DescribesProperty) SetIntransitiveActivity(v vocab.IntransitiveActivityInterface) {
	this.Clear()
	this.IntransitiveActivityMember = v
}

// SetInvite sets the value of this property. Calling IsInvite afterwards returns
// true.
func (this *DescribesProperty) SetInvite(v vocab.InviteInterface) {
	this.Clear()
	this.InviteMember = v
}

// SetJoin sets the value of this property. Calling IsJoin afterwards returns true.
func (this *DescribesProperty) SetJoin(v vocab.JoinInterface) {
	this.Clear()
	this.JoinMember = v
}

// SetLeave sets the value of this property. Calling IsLeave afterwards returns
// true.
func (this *DescribesProperty) SetLeave(v vocab.LeaveInterface) {
	this.Clear()
	this.LeaveMember = v
}

// SetLike sets the value of this property. Calling IsLike afterwards returns true.
func (this *DescribesProperty) SetLike(v vocab.LikeInterface) {
	this.Clear()
	this.LikeMember = v
}

// SetListen sets the value of this property. Calling IsListen afterwards returns
// true.
func (this *DescribesProperty) SetListen(v vocab.ListenInterface) {
	this.Clear()
	this.ListenMember = v
}

// SetMove sets the value of this property. Calling IsMove afterwards returns true.
func (this *DescribesProperty) SetMove(v vocab.MoveInterface) {
	this.Clear()
	this.MoveMember = v
}

// SetNote sets the value of this property. Calling IsNote afterwards returns true.
func (this *DescribesProperty) SetNote(v vocab.NoteInterface) {
	this.Clear()
	this.NoteMember = v
}

// SetObject sets the value of this property. Calling IsObject afterwards returns
// true.
func (this *DescribesProperty) SetObject(v vocab.ObjectInterface) {
	this.Clear()
	this.ObjectMember = v
}

// SetOffer sets the value of this property. Calling IsOffer afterwards returns
// true.
func (this *DescribesProperty) SetOffer(v vocab.OfferInterface) {
	this.Clear()
	this.OfferMember = v
}

// SetOrderedCollection sets the value of this property. Calling
// IsOrderedCollection afterwards returns true.
func (this *DescribesProperty) SetOrderedCollection(v vocab.OrderedCollectionInterface) {
	this.Clear()
	this.OrderedCollectionMember = v
}

// SetOrderedCollectionPage sets the value of this property. Calling
// IsOrderedCollectionPage afterwards returns true.
func (this *DescribesProperty) SetOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.Clear()
	this.OrderedCollectionPageMember = v
}

// SetOrganization sets the value of this property. Calling IsOrganization
// afterwards returns true.
func (this *DescribesProperty) SetOrganization(v vocab.OrganizationInterface) {
	this.Clear()
	this.OrganizationMember = v
}

// SetPage sets the value of this property. Calling IsPage afterwards returns true.
func (this *DescribesProperty) SetPage(v vocab.PageInterface) {
	this.Clear()
	this.PageMember = v
}

// SetPerson sets the value of this property. Calling IsPerson afterwards returns
// true.
func (this *DescribesProperty) SetPerson(v vocab.PersonInterface) {
	this.Clear()
	this.PersonMember = v
}

// SetPlace sets the value of this property. Calling IsPlace afterwards returns
// true.
func (this *DescribesProperty) SetPlace(v vocab.PlaceInterface) {
	this.Clear()
	this.PlaceMember = v
}

// SetProfile sets the value of this property. Calling IsProfile afterwards
// returns true.
func (this *DescribesProperty) SetProfile(v vocab.ProfileInterface) {
	this.Clear()
	this.ProfileMember = v
}

// SetQuestion sets the value of this property. Calling IsQuestion afterwards
// returns true.
func (this *DescribesProperty) SetQuestion(v vocab.QuestionInterface) {
	this.Clear()
	this.QuestionMember = v
}

// SetRead sets the value of this property. Calling IsRead afterwards returns true.
func (this *DescribesProperty) SetRead(v vocab.ReadInterface) {
	this.Clear()
	this.ReadMember = v
}

// SetReject sets the value of this property. Calling IsReject afterwards returns
// true.
func (this *DescribesProperty) SetReject(v vocab.RejectInterface) {
	this.Clear()
	this.RejectMember = v
}

// SetRelationship sets the value of this property. Calling IsRelationship
// afterwards returns true.
func (this *DescribesProperty) SetRelationship(v vocab.RelationshipInterface) {
	this.Clear()
	this.RelationshipMember = v
}

// SetRemove sets the value of this property. Calling IsRemove afterwards returns
// true.
func (this *DescribesProperty) SetRemove(v vocab.RemoveInterface) {
	this.Clear()
	this.RemoveMember = v
}

// SetService sets the value of this property. Calling IsService afterwards
// returns true.
func (this *DescribesProperty) SetService(v vocab.ServiceInterface) {
	this.Clear()
	this.ServiceMember = v
}

// SetTentativeAccept sets the value of this property. Calling IsTentativeAccept
// afterwards returns true.
func (this *DescribesProperty) SetTentativeAccept(v vocab.TentativeAcceptInterface) {
	this.Clear()
	this.TentativeAcceptMember = v
}

// SetTentativeReject sets the value of this property. Calling IsTentativeReject
// afterwards returns true.
func (this *DescribesProperty) SetTentativeReject(v vocab.TentativeRejectInterface) {
	this.Clear()
	this.TentativeRejectMember = v
}

// SetTombstone sets the value of this property. Calling IsTombstone afterwards
// returns true.
func (this *DescribesProperty) SetTombstone(v vocab.TombstoneInterface) {
	this.Clear()
	this.TombstoneMember = v
}

// SetTravel sets the value of this property. Calling IsTravel afterwards returns
// true.
func (this *DescribesProperty) SetTravel(v vocab.TravelInterface) {
	this.Clear()
	this.TravelMember = v
}

// SetUndo sets the value of this property. Calling IsUndo afterwards returns true.
func (this *DescribesProperty) SetUndo(v vocab.UndoInterface) {
	this.Clear()
	this.UndoMember = v
}

// SetUpdate sets the value of this property. Calling IsUpdate afterwards returns
// true.
func (this *DescribesProperty) SetUpdate(v vocab.UpdateInterface) {
	this.Clear()
	this.UpdateMember = v
}

// SetVideo sets the value of this property. Calling IsVideo afterwards returns
// true.
func (this *DescribesProperty) SetVideo(v vocab.VideoInterface) {
	this.Clear()
	this.VideoMember = v
}

// SetView sets the value of this property. Calling IsView afterwards returns true.
func (this *DescribesProperty) SetView(v vocab.ViewInterface) {
	this.Clear()
	this.ViewMember = v
}
