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
)

var mgr *Manager

// init handles the 'magic' of creating a Manager and dependency-injecting it into
// every other code-generated package. This gives the implementations access
// to create any type needed to deserialize, without relying on the other
// specific concrete implementations. In order to replace a go-fed created
// type with your own, be sure to have the manager call your own
// implementation's deserialize functions instead of the built-in type.
// Finally, each implementation views the Manager as an interface with only a
// subset of funcitons available. This means this Manager implements the union
// of those interfaces.
func init() {
	mgr = &Manager{}
	propertylongitude.SetManager(mgr)
	typerelationship.SetManager(mgr)
	typeadd.SetManager(mgr)
	typecollectionpage.SetManager(mgr)
	typeignore.SetManager(mgr)
	typereject.SetManager(mgr)
	propertybto.SetManager(mgr)
	typeblock.SetManager(mgr)
	propertyname.SetManager(mgr)
	propertylast.SetManager(mgr)
	propertystartindex.SetManager(mgr)
	typeorderedcollection.SetManager(mgr)
	typeflag.SetManager(mgr)
	typeleave.SetManager(mgr)
	typemove.SetManager(mgr)
	typeservice.SetManager(mgr)
	typevideo.SetManager(mgr)
	propertycontext.SetManager(mgr)
	typeapplication.SetManager(mgr)
	typeorderedcollectionpage.SetManager(mgr)
	typetentativeaccept.SetManager(mgr)
	propertyorigin.SetManager(mgr)
	typeintransitiveactivity.SetManager(mgr)
	typeview.SetManager(mgr)
	propertyheight.SetManager(mgr)
	propertyprev.SetManager(mgr)
	typearticle.SetManager(mgr)
	typepage.SetManager(mgr)
	typetombstone.SetManager(mgr)
	typeundo.SetManager(mgr)
	propertyitems.SetManager(mgr)
	propertyurl.SetManager(mgr)
	propertywidth.SetManager(mgr)
	typejoin.SetManager(mgr)
	typeorganization.SetManager(mgr)
	typeactivity.SetManager(mgr)
	typeupdate.SetManager(mgr)
	propertyaltitude.SetManager(mgr)
	propertyattachment.SetManager(mgr)
	propertyinreplyto.SetManager(mgr)
	typeperson.SetManager(mgr)
	typeevent.SetManager(mgr)
	typeprofile.SetManager(mgr)
	propertylocation.SetManager(mgr)
	typeaudio.SetManager(mgr)
	propertyicon.SetManager(mgr)
	propertyreplies.SetManager(mgr)
	typedislike.SetManager(mgr)
	propertydescribes.SetManager(mgr)
	propertyduration.SetManager(mgr)
	propertylatitude.SetManager(mgr)
	propertypublished.SetManager(mgr)
	propertytag.SetManager(mgr)
	propertybcc.SetManager(mgr)
	typeplace.SetManager(mgr)
	propertyattributedto.SetManager(mgr)
	propertyhreflang.SetManager(mgr)
	propertyradius.SetManager(mgr)
	typeannounce.SetManager(mgr)
	typenote.SetManager(mgr)
	propertytype.SetManager(mgr)
	typefollow.SetManager(mgr)
	typeremove.SetManager(mgr)
	typetravel.SetManager(mgr)
	propertyformertype.SetManager(mgr)
	propertymediatype.SetManager(mgr)
	typelink.SetManager(mgr)
	propertycontent.SetManager(mgr)
	propertydeleted.SetManager(mgr)
	propertysummary.SetManager(mgr)
	typedocument.SetManager(mgr)
	propertycurrent.SetManager(mgr)
	propertycc.SetManager(mgr)
	typetentativereject.SetManager(mgr)
	propertyimage.SetManager(mgr)
	typelike.SetManager(mgr)
	typeinvite.SetManager(mgr)
	typeobject.SetManager(mgr)
	propertyhref.SetManager(mgr)
	typegroup.SetManager(mgr)
	propertyclosed.SetManager(mgr)
	propertyendtime.SetManager(mgr)
	propertyid.SetManager(mgr)
	propertyoneof.SetManager(mgr)
	typecreate.SetManager(mgr)
	typeoffer.SetManager(mgr)
	typearrive.SetManager(mgr)
	propertyfirst.SetManager(mgr)
	propertyinstrument.SetManager(mgr)
	propertynext.SetManager(mgr)
	propertyupdated.SetManager(mgr)
	propertyaccuracy.SetManager(mgr)
	typemention.SetManager(mgr)
	typequestion.SetManager(mgr)
	typecollection.SetManager(mgr)
	typeread.SetManager(mgr)
	propertystarttime.SetManager(mgr)
	propertysubject.SetManager(mgr)
	propertyunits.SetManager(mgr)
	typelisten.SetManager(mgr)
	propertyactor.SetManager(mgr)
	propertyanyof.SetManager(mgr)
	propertygenerator.SetManager(mgr)
	propertypartof.SetManager(mgr)
	typedelete.SetManager(mgr)
	propertytarget.SetManager(mgr)
	propertyto.SetManager(mgr)
	propertyobject.SetManager(mgr)
	propertyresult.SetManager(mgr)
	propertyaudience.SetManager(mgr)
	propertyrel.SetManager(mgr)
	propertyrelationship.SetManager(mgr)
	typeimage.SetManager(mgr)
	propertypreview.SetManager(mgr)
	propertytotalitems.SetManager(mgr)
	typeaccept.SetManager(mgr)
}
