package streams

import (
	propertyaccuracy "github.com/go-fed/activity/streams/impl/activitystreams/property_accuracy"
	propertyactor "github.com/go-fed/activity/streams/impl/activitystreams/property_actor"
	propertyaltitude "github.com/go-fed/activity/streams/impl/activitystreams/property_altitude"
	propertyanyof "github.com/go-fed/activity/streams/impl/activitystreams/property_anyof"
	propertyattachment "github.com/go-fed/activity/streams/impl/activitystreams/property_attachment"
	propertyattributedto "github.com/go-fed/activity/streams/impl/activitystreams/property_attributedto"
	propertyaudience "github.com/go-fed/activity/streams/impl/activitystreams/property_audience"
	propertybcc "github.com/go-fed/activity/streams/impl/activitystreams/property_bcc"
	propertybto "github.com/go-fed/activity/streams/impl/activitystreams/property_bto"
	propertycc "github.com/go-fed/activity/streams/impl/activitystreams/property_cc"
	propertyclosed "github.com/go-fed/activity/streams/impl/activitystreams/property_closed"
	propertycontent "github.com/go-fed/activity/streams/impl/activitystreams/property_content"
	propertycontext "github.com/go-fed/activity/streams/impl/activitystreams/property_context"
	propertycurrent "github.com/go-fed/activity/streams/impl/activitystreams/property_current"
	propertydeleted "github.com/go-fed/activity/streams/impl/activitystreams/property_deleted"
	propertydescribes "github.com/go-fed/activity/streams/impl/activitystreams/property_describes"
	propertyduration "github.com/go-fed/activity/streams/impl/activitystreams/property_duration"
	propertyendtime "github.com/go-fed/activity/streams/impl/activitystreams/property_endtime"
	propertyfirst "github.com/go-fed/activity/streams/impl/activitystreams/property_first"
	propertyformertype "github.com/go-fed/activity/streams/impl/activitystreams/property_formertype"
	propertygenerator "github.com/go-fed/activity/streams/impl/activitystreams/property_generator"
	propertyheight "github.com/go-fed/activity/streams/impl/activitystreams/property_height"
	propertyhref "github.com/go-fed/activity/streams/impl/activitystreams/property_href"
	propertyhreflang "github.com/go-fed/activity/streams/impl/activitystreams/property_hreflang"
	propertyicon "github.com/go-fed/activity/streams/impl/activitystreams/property_icon"
	propertyid "github.com/go-fed/activity/streams/impl/activitystreams/property_id"
	propertyimage "github.com/go-fed/activity/streams/impl/activitystreams/property_image"
	propertyinreplyto "github.com/go-fed/activity/streams/impl/activitystreams/property_inreplyto"
	propertyinstrument "github.com/go-fed/activity/streams/impl/activitystreams/property_instrument"
	propertyitems "github.com/go-fed/activity/streams/impl/activitystreams/property_items"
	propertylast "github.com/go-fed/activity/streams/impl/activitystreams/property_last"
	propertylatitude "github.com/go-fed/activity/streams/impl/activitystreams/property_latitude"
	propertylocation "github.com/go-fed/activity/streams/impl/activitystreams/property_location"
	propertylongitude "github.com/go-fed/activity/streams/impl/activitystreams/property_longitude"
	propertymediatype "github.com/go-fed/activity/streams/impl/activitystreams/property_mediatype"
	propertyname "github.com/go-fed/activity/streams/impl/activitystreams/property_name"
	propertynext "github.com/go-fed/activity/streams/impl/activitystreams/property_next"
	propertyobject "github.com/go-fed/activity/streams/impl/activitystreams/property_object"
	propertyoneof "github.com/go-fed/activity/streams/impl/activitystreams/property_oneof"
	propertyorigin "github.com/go-fed/activity/streams/impl/activitystreams/property_origin"
	propertypartof "github.com/go-fed/activity/streams/impl/activitystreams/property_partof"
	propertyprev "github.com/go-fed/activity/streams/impl/activitystreams/property_prev"
	propertypreview "github.com/go-fed/activity/streams/impl/activitystreams/property_preview"
	propertypublished "github.com/go-fed/activity/streams/impl/activitystreams/property_published"
	propertyradius "github.com/go-fed/activity/streams/impl/activitystreams/property_radius"
	propertyrel "github.com/go-fed/activity/streams/impl/activitystreams/property_rel"
	propertyrelationship "github.com/go-fed/activity/streams/impl/activitystreams/property_relationship"
	propertyreplies "github.com/go-fed/activity/streams/impl/activitystreams/property_replies"
	propertyresult "github.com/go-fed/activity/streams/impl/activitystreams/property_result"
	propertystartindex "github.com/go-fed/activity/streams/impl/activitystreams/property_startindex"
	propertystarttime "github.com/go-fed/activity/streams/impl/activitystreams/property_starttime"
	propertysubject "github.com/go-fed/activity/streams/impl/activitystreams/property_subject"
	propertysummary "github.com/go-fed/activity/streams/impl/activitystreams/property_summary"
	propertytag "github.com/go-fed/activity/streams/impl/activitystreams/property_tag"
	propertytarget "github.com/go-fed/activity/streams/impl/activitystreams/property_target"
	propertyto "github.com/go-fed/activity/streams/impl/activitystreams/property_to"
	propertytotalitems "github.com/go-fed/activity/streams/impl/activitystreams/property_totalitems"
	propertytype "github.com/go-fed/activity/streams/impl/activitystreams/property_type"
	propertyunits "github.com/go-fed/activity/streams/impl/activitystreams/property_units"
	propertyupdated "github.com/go-fed/activity/streams/impl/activitystreams/property_updated"
	propertyurl "github.com/go-fed/activity/streams/impl/activitystreams/property_url"
	propertywidth "github.com/go-fed/activity/streams/impl/activitystreams/property_width"
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// Manager manages interface types and deserializations for use by generated code.
// Application code implicitly uses this manager at run-time to create
// concrete implementations of the interfaces.
type Manager struct {
}

// DeserializeAcceptActivityStreams returns the deserialization method for the
// "AcceptInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AcceptInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AcceptInterface, error) {
		i, err := typeaccept.DeserializeAccept(m, aliasMap)
		return i, err
	}
}

// DeserializeAccuracyPropertyActivityStreams returns the deserialization method
// for the "AccuracyPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAccuracyPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AccuracyPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AccuracyPropertyInterface, error) {
		i, err := propertyaccuracy.DeserializeAccuracyProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeActivityActivityStreams returns the deserialization method for the
// "ActivityInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityInterface, error) {
		i, err := typeactivity.DeserializeActivity(m, aliasMap)
		return i, err
	}
}

// DeserializeActorPropertyActivityStreams returns the deserialization method for
// the "ActorPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeActorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActorPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActorPropertyInterface, error) {
		i, err := propertyactor.DeserializeActorProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeAddActivityStreams returns the deserialization method for the
// "AddInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeAddActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AddInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AddInterface, error) {
		i, err := typeadd.DeserializeAdd(m, aliasMap)
		return i, err
	}
}

// DeserializeAltitudePropertyActivityStreams returns the deserialization method
// for the "AltitudePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAltitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AltitudePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AltitudePropertyInterface, error) {
		i, err := propertyaltitude.DeserializeAltitudeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeAnnounceActivityStreams returns the deserialization method for the
// "AnnounceInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAnnounceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AnnounceInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AnnounceInterface, error) {
		i, err := typeannounce.DeserializeAnnounce(m, aliasMap)
		return i, err
	}
}

// DeserializeAnyOfPropertyActivityStreams returns the deserialization method for
// the "AnyOfPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAnyOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AnyOfPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AnyOfPropertyInterface, error) {
		i, err := propertyanyof.DeserializeAnyOfProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeApplicationActivityStreams returns the deserialization method for
// the "ApplicationInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeApplicationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ApplicationInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ApplicationInterface, error) {
		i, err := typeapplication.DeserializeApplication(m, aliasMap)
		return i, err
	}
}

// DeserializeArriveActivityStreams returns the deserialization method for the
// "ArriveInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeArriveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ArriveInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ArriveInterface, error) {
		i, err := typearrive.DeserializeArrive(m, aliasMap)
		return i, err
	}
}

// DeserializeArticleActivityStreams returns the deserialization method for the
// "ArticleInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeArticleActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ArticleInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ArticleInterface, error) {
		i, err := typearticle.DeserializeArticle(m, aliasMap)
		return i, err
	}
}

// DeserializeAttachmentPropertyActivityStreams returns the deserialization method
// for the "AttachmentPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAttachmentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AttachmentPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AttachmentPropertyInterface, error) {
		i, err := propertyattachment.DeserializeAttachmentProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeAttributedToPropertyActivityStreams returns the deserialization
// method for the "AttributedToPropertyInterface" non-functional property in
// the vocabulary "ActivityStreams"
func (this Manager) DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AttributedToPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AttributedToPropertyInterface, error) {
		i, err := propertyattributedto.DeserializeAttributedToProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeAudiencePropertyActivityStreams returns the deserialization method
// for the "AudiencePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAudiencePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AudiencePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AudiencePropertyInterface, error) {
		i, err := propertyaudience.DeserializeAudienceProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeAudioActivityStreams returns the deserialization method for the
// "AudioInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeAudioActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AudioInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.AudioInterface, error) {
		i, err := typeaudio.DeserializeAudio(m, aliasMap)
		return i, err
	}
}

// DeserializeBccPropertyActivityStreams returns the deserialization method for
// the "BccPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeBccPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BccPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.BccPropertyInterface, error) {
		i, err := propertybcc.DeserializeBccProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeBlockActivityStreams returns the deserialization method for the
// "BlockInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeBlockActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BlockInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.BlockInterface, error) {
		i, err := typeblock.DeserializeBlock(m, aliasMap)
		return i, err
	}
}

// DeserializeBtoPropertyActivityStreams returns the deserialization method for
// the "BtoPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeBtoPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BtoPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.BtoPropertyInterface, error) {
		i, err := propertybto.DeserializeBtoProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeCcPropertyActivityStreams returns the deserialization method for the
// "CcPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCcPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CcPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.CcPropertyInterface, error) {
		i, err := propertycc.DeserializeCcProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeClosedPropertyActivityStreams returns the deserialization method for
// the "ClosedPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeClosedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ClosedPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ClosedPropertyInterface, error) {
		i, err := propertyclosed.DeserializeClosedProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeCollectionActivityStreams returns the deserialization method for the
// "CollectionInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.CollectionInterface, error) {
		i, err := typecollection.DeserializeCollection(m, aliasMap)
		return i, err
	}
}

// DeserializeCollectionPageActivityStreams returns the deserialization method for
// the "CollectionPageInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionPageInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.CollectionPageInterface, error) {
		i, err := typecollectionpage.DeserializeCollectionPage(m, aliasMap)
		return i, err
	}
}

// DeserializeContentPropertyActivityStreams returns the deserialization method
// for the "ContentPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeContentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ContentPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ContentPropertyInterface, error) {
		i, err := propertycontent.DeserializeContentProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeContextPropertyActivityStreams returns the deserialization method
// for the "ContextPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeContextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ContextPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ContextPropertyInterface, error) {
		i, err := propertycontext.DeserializeContextProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeCreateActivityStreams returns the deserialization method for the
// "CreateInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCreateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CreateInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.CreateInterface, error) {
		i, err := typecreate.DeserializeCreate(m, aliasMap)
		return i, err
	}
}

// DeserializeCurrentPropertyActivityStreams returns the deserialization method
// for the "CurrentPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeCurrentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CurrentPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.CurrentPropertyInterface, error) {
		i, err := propertycurrent.DeserializeCurrentProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeDeleteActivityStreams returns the deserialization method for the
// "DeleteInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDeleteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DeleteInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DeleteInterface, error) {
		i, err := typedelete.DeserializeDelete(m, aliasMap)
		return i, err
	}
}

// DeserializeDeletedPropertyActivityStreams returns the deserialization method
// for the "DeletedPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDeletedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DeletedPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DeletedPropertyInterface, error) {
		i, err := propertydeleted.DeserializeDeletedProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeDescribesPropertyActivityStreams returns the deserialization method
// for the "DescribesPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDescribesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DescribesPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DescribesPropertyInterface, error) {
		i, err := propertydescribes.DeserializeDescribesProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeDislikeActivityStreams returns the deserialization method for the
// "DislikeInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDislikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DislikeInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DislikeInterface, error) {
		i, err := typedislike.DeserializeDislike(m, aliasMap)
		return i, err
	}
}

// DeserializeDocumentActivityStreams returns the deserialization method for the
// "DocumentInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDocumentActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DocumentInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DocumentInterface, error) {
		i, err := typedocument.DeserializeDocument(m, aliasMap)
		return i, err
	}
}

// DeserializeDurationPropertyActivityStreams returns the deserialization method
// for the "DurationPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDurationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DurationPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.DurationPropertyInterface, error) {
		i, err := propertyduration.DeserializeDurationProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeEndTimePropertyActivityStreams returns the deserialization method
// for the "EndTimePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeEndTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.EndTimePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.EndTimePropertyInterface, error) {
		i, err := propertyendtime.DeserializeEndTimeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeEventActivityStreams returns the deserialization method for the
// "EventInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeEventActivityStreams() func(map[string]interface{}, map[string]string) (vocab.EventInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.EventInterface, error) {
		i, err := typeevent.DeserializeEvent(m, aliasMap)
		return i, err
	}
}

// DeserializeFirstPropertyActivityStreams returns the deserialization method for
// the "FirstPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeFirstPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FirstPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.FirstPropertyInterface, error) {
		i, err := propertyfirst.DeserializeFirstProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeFlagActivityStreams returns the deserialization method for the
// "FlagInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeFlagActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FlagInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.FlagInterface, error) {
		i, err := typeflag.DeserializeFlag(m, aliasMap)
		return i, err
	}
}

// DeserializeFollowActivityStreams returns the deserialization method for the
// "FollowInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeFollowActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FollowInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.FollowInterface, error) {
		i, err := typefollow.DeserializeFollow(m, aliasMap)
		return i, err
	}
}

// DeserializeFormerTypePropertyActivityStreams returns the deserialization method
// for the "FormerTypePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeFormerTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FormerTypePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.FormerTypePropertyInterface, error) {
		i, err := propertyformertype.DeserializeFormerTypeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeGeneratorPropertyActivityStreams returns the deserialization method
// for the "GeneratorPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeGeneratorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.GeneratorPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.GeneratorPropertyInterface, error) {
		i, err := propertygenerator.DeserializeGeneratorProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeGroupActivityStreams returns the deserialization method for the
// "GroupInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeGroupActivityStreams() func(map[string]interface{}, map[string]string) (vocab.GroupInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.GroupInterface, error) {
		i, err := typegroup.DeserializeGroup(m, aliasMap)
		return i, err
	}
}

// DeserializeHeightPropertyActivityStreams returns the deserialization method for
// the "HeightPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeHeightPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HeightPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.HeightPropertyInterface, error) {
		i, err := propertyheight.DeserializeHeightProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeHrefPropertyActivityStreams returns the deserialization method for
// the "HrefPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeHrefPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HrefPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.HrefPropertyInterface, error) {
		i, err := propertyhref.DeserializeHrefProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeHreflangPropertyActivityStreams returns the deserialization method
// for the "HreflangPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeHreflangPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HreflangPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.HreflangPropertyInterface, error) {
		i, err := propertyhreflang.DeserializeHreflangProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeIconPropertyActivityStreams returns the deserialization method for
// the "IconPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeIconPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IconPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.IconPropertyInterface, error) {
		i, err := propertyicon.DeserializeIconProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeIdPropertyActivityStreams returns the deserialization method for the
// "IdPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeIdPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IdPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.IdPropertyInterface, error) {
		i, err := propertyid.DeserializeIdProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeIgnoreActivityStreams returns the deserialization method for the
// "IgnoreInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeIgnoreActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IgnoreInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.IgnoreInterface, error) {
		i, err := typeignore.DeserializeIgnore(m, aliasMap)
		return i, err
	}
}

// DeserializeImageActivityStreams returns the deserialization method for the
// "ImageInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeImageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ImageInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ImageInterface, error) {
		i, err := typeimage.DeserializeImage(m, aliasMap)
		return i, err
	}
}

// DeserializeImagePropertyActivityStreams returns the deserialization method for
// the "ImagePropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeImagePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ImagePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ImagePropertyInterface, error) {
		i, err := propertyimage.DeserializeImageProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeInReplyToPropertyActivityStreams returns the deserialization method
// for the "InReplyToPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeInReplyToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.InReplyToPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.InReplyToPropertyInterface, error) {
		i, err := propertyinreplyto.DeserializeInReplyToProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeInstrumentPropertyActivityStreams returns the deserialization method
// for the "InstrumentPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeInstrumentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.InstrumentPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.InstrumentPropertyInterface, error) {
		i, err := propertyinstrument.DeserializeInstrumentProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeIntransitiveActivityActivityStreams returns the deserialization
// method for the "IntransitiveActivityInterface" non-functional property in
// the vocabulary "ActivityStreams"
func (this Manager) DeserializeIntransitiveActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IntransitiveActivityInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.IntransitiveActivityInterface, error) {
		i, err := typeintransitiveactivity.DeserializeIntransitiveActivity(m, aliasMap)
		return i, err
	}
}

// DeserializeInviteActivityStreams returns the deserialization method for the
// "InviteInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeInviteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.InviteInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.InviteInterface, error) {
		i, err := typeinvite.DeserializeInvite(m, aliasMap)
		return i, err
	}
}

// DeserializeItemsPropertyActivityStreams returns the deserialization method for
// the "ItemsPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ItemsPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ItemsPropertyInterface, error) {
		i, err := propertyitems.DeserializeItemsProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeJoinActivityStreams returns the deserialization method for the
// "JoinInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeJoinActivityStreams() func(map[string]interface{}, map[string]string) (vocab.JoinInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.JoinInterface, error) {
		i, err := typejoin.DeserializeJoin(m, aliasMap)
		return i, err
	}
}

// DeserializeLastPropertyActivityStreams returns the deserialization method for
// the "LastPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeLastPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LastPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LastPropertyInterface, error) {
		i, err := propertylast.DeserializeLastProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeLatitudePropertyActivityStreams returns the deserialization method
// for the "LatitudePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLatitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LatitudePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LatitudePropertyInterface, error) {
		i, err := propertylatitude.DeserializeLatitudeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeLeaveActivityStreams returns the deserialization method for the
// "LeaveInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeLeaveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LeaveInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LeaveInterface, error) {
		i, err := typeleave.DeserializeLeave(m, aliasMap)
		return i, err
	}
}

// DeserializeLikeActivityStreams returns the deserialization method for the
// "LikeInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeLikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LikeInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LikeInterface, error) {
		i, err := typelike.DeserializeLike(m, aliasMap)
		return i, err
	}
}

// DeserializeLinkActivityStreams returns the deserialization method for the
// "LinkInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LinkInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LinkInterface, error) {
		i, err := typelink.DeserializeLink(m, aliasMap)
		return i, err
	}
}

// DeserializeListenActivityStreams returns the deserialization method for the
// "ListenInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeListenActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ListenInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ListenInterface, error) {
		i, err := typelisten.DeserializeListen(m, aliasMap)
		return i, err
	}
}

// DeserializeLocationPropertyActivityStreams returns the deserialization method
// for the "LocationPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLocationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LocationPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LocationPropertyInterface, error) {
		i, err := propertylocation.DeserializeLocationProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeLongitudePropertyActivityStreams returns the deserialization method
// for the "LongitudePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLongitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LongitudePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.LongitudePropertyInterface, error) {
		i, err := propertylongitude.DeserializeLongitudeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeMediaTypePropertyActivityStreams returns the deserialization method
// for the "MediaTypePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MediaTypePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.MediaTypePropertyInterface, error) {
		i, err := propertymediatype.DeserializeMediaTypeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeMentionActivityStreams returns the deserialization method for the
// "MentionInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MentionInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.MentionInterface, error) {
		i, err := typemention.DeserializeMention(m, aliasMap)
		return i, err
	}
}

// DeserializeMoveActivityStreams returns the deserialization method for the
// "MoveInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeMoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MoveInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.MoveInterface, error) {
		i, err := typemove.DeserializeMove(m, aliasMap)
		return i, err
	}
}

// DeserializeNamePropertyActivityStreams returns the deserialization method for
// the "NamePropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NamePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.NamePropertyInterface, error) {
		i, err := propertyname.DeserializeNameProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeNextPropertyActivityStreams returns the deserialization method for
// the "NextPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeNextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NextPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.NextPropertyInterface, error) {
		i, err := propertynext.DeserializeNextProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeNoteActivityStreams returns the deserialization method for the
// "NoteInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeNoteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NoteInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.NoteInterface, error) {
		i, err := typenote.DeserializeNote(m, aliasMap)
		return i, err
	}
}

// DeserializeObjectActivityStreams returns the deserialization method for the
// "ObjectInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeObjectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ObjectInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ObjectInterface, error) {
		i, err := typeobject.DeserializeObject(m, aliasMap)
		return i, err
	}
}

// DeserializeObjectPropertyActivityStreams returns the deserialization method for
// the "ObjectPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeObjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ObjectPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ObjectPropertyInterface, error) {
		i, err := propertyobject.DeserializeObjectProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeOfferActivityStreams returns the deserialization method for the
// "OfferInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeOfferActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OfferInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OfferInterface, error) {
		i, err := typeoffer.DeserializeOffer(m, aliasMap)
		return i, err
	}
}

// DeserializeOneOfPropertyActivityStreams returns the deserialization method for
// the "OneOfPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeOneOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OneOfPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OneOfPropertyInterface, error) {
		i, err := propertyoneof.DeserializeOneOfProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeOrderedCollectionActivityStreams returns the deserialization method
// for the "OrderedCollectionInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OrderedCollectionInterface, error) {
		i, err := typeorderedcollection.DeserializeOrderedCollection(m, aliasMap)
		return i, err
	}
}

// DeserializeOrderedCollectionPageActivityStreams returns the deserialization
// method for the "OrderedCollectionPageInterface" non-functional property in
// the vocabulary "ActivityStreams"
func (this Manager) DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionPageInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OrderedCollectionPageInterface, error) {
		i, err := typeorderedcollectionpage.DeserializeOrderedCollectionPage(m, aliasMap)
		return i, err
	}
}

// DeserializeOrganizationActivityStreams returns the deserialization method for
// the "OrganizationInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeOrganizationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrganizationInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OrganizationInterface, error) {
		i, err := typeorganization.DeserializeOrganization(m, aliasMap)
		return i, err
	}
}

// DeserializeOriginPropertyActivityStreams returns the deserialization method for
// the "OriginPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeOriginPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OriginPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.OriginPropertyInterface, error) {
		i, err := propertyorigin.DeserializeOriginProperty(m, aliasMap)
		return i, err
	}
}

// DeserializePageActivityStreams returns the deserialization method for the
// "PageInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializePageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PageInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PageInterface, error) {
		i, err := typepage.DeserializePage(m, aliasMap)
		return i, err
	}
}

// DeserializePartOfPropertyActivityStreams returns the deserialization method for
// the "PartOfPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePartOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PartOfPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PartOfPropertyInterface, error) {
		i, err := propertypartof.DeserializePartOfProperty(m, aliasMap)
		return i, err
	}
}

// DeserializePersonActivityStreams returns the deserialization method for the
// "PersonInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePersonActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PersonInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PersonInterface, error) {
		i, err := typeperson.DeserializePerson(m, aliasMap)
		return i, err
	}
}

// DeserializePlaceActivityStreams returns the deserialization method for the
// "PlaceInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializePlaceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PlaceInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PlaceInterface, error) {
		i, err := typeplace.DeserializePlace(m, aliasMap)
		return i, err
	}
}

// DeserializePrevPropertyActivityStreams returns the deserialization method for
// the "PrevPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePrevPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PrevPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PrevPropertyInterface, error) {
		i, err := propertyprev.DeserializePrevProperty(m, aliasMap)
		return i, err
	}
}

// DeserializePreviewPropertyActivityStreams returns the deserialization method
// for the "PreviewPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PreviewPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PreviewPropertyInterface, error) {
		i, err := propertypreview.DeserializePreviewProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeProfileActivityStreams returns the deserialization method for the
// "ProfileInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeProfileActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ProfileInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ProfileInterface, error) {
		i, err := typeprofile.DeserializeProfile(m, aliasMap)
		return i, err
	}
}

// DeserializePublishedPropertyActivityStreams returns the deserialization method
// for the "PublishedPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializePublishedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PublishedPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.PublishedPropertyInterface, error) {
		i, err := propertypublished.DeserializePublishedProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeQuestionActivityStreams returns the deserialization method for the
// "QuestionInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeQuestionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.QuestionInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.QuestionInterface, error) {
		i, err := typequestion.DeserializeQuestion(m, aliasMap)
		return i, err
	}
}

// DeserializeRadiusPropertyActivityStreams returns the deserialization method for
// the "RadiusPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRadiusPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RadiusPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RadiusPropertyInterface, error) {
		i, err := propertyradius.DeserializeRadiusProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeReadActivityStreams returns the deserialization method for the
// "ReadInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeReadActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ReadInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ReadInterface, error) {
		i, err := typeread.DeserializeRead(m, aliasMap)
		return i, err
	}
}

// DeserializeRejectActivityStreams returns the deserialization method for the
// "RejectInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RejectInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RejectInterface, error) {
		i, err := typereject.DeserializeReject(m, aliasMap)
		return i, err
	}
}

// DeserializeRelPropertyActivityStreams returns the deserialization method for
// the "RelPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRelPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RelPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RelPropertyInterface, error) {
		i, err := propertyrel.DeserializeRelProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeRelationshipActivityStreams returns the deserialization method for
// the "RelationshipInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRelationshipActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RelationshipInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RelationshipInterface, error) {
		i, err := typerelationship.DeserializeRelationship(m, aliasMap)
		return i, err
	}
}

// DeserializeRelationshipPropertyActivityStreams returns the deserialization
// method for the "RelationshipPropertyInterface" non-functional property in
// the vocabulary "ActivityStreams"
func (this Manager) DeserializeRelationshipPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RelationshipPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RelationshipPropertyInterface, error) {
		i, err := propertyrelationship.DeserializeRelationshipProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeRemoveActivityStreams returns the deserialization method for the
// "RemoveInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRemoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RemoveInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RemoveInterface, error) {
		i, err := typeremove.DeserializeRemove(m, aliasMap)
		return i, err
	}
}

// DeserializeRepliesPropertyActivityStreams returns the deserialization method
// for the "RepliesPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeRepliesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RepliesPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.RepliesPropertyInterface, error) {
		i, err := propertyreplies.DeserializeRepliesProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeResultPropertyActivityStreams returns the deserialization method for
// the "ResultPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeResultPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ResultPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ResultPropertyInterface, error) {
		i, err := propertyresult.DeserializeResultProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeServiceActivityStreams returns the deserialization method for the
// "ServiceInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeServiceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ServiceInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ServiceInterface, error) {
		i, err := typeservice.DeserializeService(m, aliasMap)
		return i, err
	}
}

// DeserializeStartIndexPropertyActivityStreams returns the deserialization method
// for the "StartIndexPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeStartIndexPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.StartIndexPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.StartIndexPropertyInterface, error) {
		i, err := propertystartindex.DeserializeStartIndexProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeStartTimePropertyActivityStreams returns the deserialization method
// for the "StartTimePropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeStartTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.StartTimePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.StartTimePropertyInterface, error) {
		i, err := propertystarttime.DeserializeStartTimeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeSubjectPropertyActivityStreams returns the deserialization method
// for the "SubjectPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeSubjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.SubjectPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.SubjectPropertyInterface, error) {
		i, err := propertysubject.DeserializeSubjectProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeSummaryPropertyActivityStreams returns the deserialization method
// for the "SummaryPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.SummaryPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.SummaryPropertyInterface, error) {
		i, err := propertysummary.DeserializeSummaryProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeTagPropertyActivityStreams returns the deserialization method for
// the "TagPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTagPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TagPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TagPropertyInterface, error) {
		i, err := propertytag.DeserializeTagProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeTargetPropertyActivityStreams returns the deserialization method for
// the "TargetPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTargetPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TargetPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TargetPropertyInterface, error) {
		i, err := propertytarget.DeserializeTargetProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeTentativeAcceptActivityStreams returns the deserialization method
// for the "TentativeAcceptInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTentativeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TentativeAcceptInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TentativeAcceptInterface, error) {
		i, err := typetentativeaccept.DeserializeTentativeAccept(m, aliasMap)
		return i, err
	}
}

// DeserializeTentativeRejectActivityStreams returns the deserialization method
// for the "TentativeRejectInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTentativeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TentativeRejectInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TentativeRejectInterface, error) {
		i, err := typetentativereject.DeserializeTentativeReject(m, aliasMap)
		return i, err
	}
}

// DeserializeToPropertyActivityStreams returns the deserialization method for the
// "ToPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ToPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ToPropertyInterface, error) {
		i, err := propertyto.DeserializeToProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeTombstoneActivityStreams returns the deserialization method for the
// "TombstoneInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTombstoneActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TombstoneInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TombstoneInterface, error) {
		i, err := typetombstone.DeserializeTombstone(m, aliasMap)
		return i, err
	}
}

// DeserializeTotalItemsPropertyActivityStreams returns the deserialization method
// for the "TotalItemsPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTotalItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TotalItemsPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TotalItemsPropertyInterface, error) {
		i, err := propertytotalitems.DeserializeTotalItemsProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeTravelActivityStreams returns the deserialization method for the
// "TravelInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTravelActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TravelInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TravelInterface, error) {
		i, err := typetravel.DeserializeTravel(m, aliasMap)
		return i, err
	}
}

// DeserializeTypePropertyActivityStreams returns the deserialization method for
// the "TypePropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TypePropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.TypePropertyInterface, error) {
		i, err := propertytype.DeserializeTypeProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeUndoActivityStreams returns the deserialization method for the
// "UndoInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeUndoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UndoInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.UndoInterface, error) {
		i, err := typeundo.DeserializeUndo(m, aliasMap)
		return i, err
	}
}

// DeserializeUnitsPropertyActivityStreams returns the deserialization method for
// the "UnitsPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUnitsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UnitsPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.UnitsPropertyInterface, error) {
		i, err := propertyunits.DeserializeUnitsProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeUpdateActivityStreams returns the deserialization method for the
// "UpdateInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUpdateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UpdateInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.UpdateInterface, error) {
		i, err := typeupdate.DeserializeUpdate(m, aliasMap)
		return i, err
	}
}

// DeserializeUpdatedPropertyActivityStreams returns the deserialization method
// for the "UpdatedPropertyInterface" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeUpdatedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UpdatedPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.UpdatedPropertyInterface, error) {
		i, err := propertyupdated.DeserializeUpdatedProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeUrlPropertyActivityStreams returns the deserialization method for
// the "UrlPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUrlPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UrlPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.UrlPropertyInterface, error) {
		i, err := propertyurl.DeserializeUrlProperty(m, aliasMap)
		return i, err
	}
}

// DeserializeVideoActivityStreams returns the deserialization method for the
// "VideoInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeVideoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.VideoInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.VideoInterface, error) {
		i, err := typevideo.DeserializeVideo(m, aliasMap)
		return i, err
	}
}

// DeserializeViewActivityStreams returns the deserialization method for the
// "ViewInterface" non-functional property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeViewActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ViewInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ViewInterface, error) {
		i, err := typeview.DeserializeView(m, aliasMap)
		return i, err
	}
}

// DeserializeWidthPropertyActivityStreams returns the deserialization method for
// the "WidthPropertyInterface" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeWidthPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.WidthPropertyInterface, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.WidthPropertyInterface, error) {
		i, err := propertywidth.DeserializeWidthProperty(m, aliasMap)
		return i, err
	}
}
