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
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewActivityStreamsHreflangProperty creates a new HreflangPropertyInterface
func NewActivityStreamsHreflangProperty() vocab.HreflangPropertyInterface {
	return propertyhreflang.NewHreflangProperty()
}

// NewActivityStreamsIdProperty creates a new IdPropertyInterface
func NewActivityStreamsIdProperty() vocab.IdPropertyInterface {
	return propertyid.NewIdProperty()
}

// NewActivityStreamsEndTimeProperty creates a new EndTimePropertyInterface
func NewActivityStreamsEndTimeProperty() vocab.EndTimePropertyInterface {
	return propertyendtime.NewEndTimeProperty()
}

// NewActivityStreamsHrefProperty creates a new HrefPropertyInterface
func NewActivityStreamsHrefProperty() vocab.HrefPropertyInterface {
	return propertyhref.NewHrefProperty()
}

// NewActivityStreamsUpdatedProperty creates a new UpdatedPropertyInterface
func NewActivityStreamsUpdatedProperty() vocab.UpdatedPropertyInterface {
	return propertyupdated.NewUpdatedProperty()
}

// NewActivityStreamsNextProperty creates a new NextPropertyInterface
func NewActivityStreamsNextProperty() vocab.NextPropertyInterface {
	return propertynext.NewNextProperty()
}

// NewActivityStreamsSubjectProperty creates a new SubjectPropertyInterface
func NewActivityStreamsSubjectProperty() vocab.SubjectPropertyInterface {
	return propertysubject.NewSubjectProperty()
}

// NewActivityStreamsDurationProperty creates a new DurationPropertyInterface
func NewActivityStreamsDurationProperty() vocab.DurationPropertyInterface {
	return propertyduration.NewDurationProperty()
}

// NewActivityStreamsFirstProperty creates a new FirstPropertyInterface
func NewActivityStreamsFirstProperty() vocab.FirstPropertyInterface {
	return propertyfirst.NewFirstProperty()
}

// NewActivityStreamsDeletedProperty creates a new DeletedPropertyInterface
func NewActivityStreamsDeletedProperty() vocab.DeletedPropertyInterface {
	return propertydeleted.NewDeletedProperty()
}

// NewActivityStreamsUnitsProperty creates a new UnitsPropertyInterface
func NewActivityStreamsUnitsProperty() vocab.UnitsPropertyInterface {
	return propertyunits.NewUnitsProperty()
}

// NewActivityStreamsTotalItemsProperty creates a new TotalItemsPropertyInterface
func NewActivityStreamsTotalItemsProperty() vocab.TotalItemsPropertyInterface {
	return propertytotalitems.NewTotalItemsProperty()
}

// NewActivityStreamsLatitudeProperty creates a new LatitudePropertyInterface
func NewActivityStreamsLatitudeProperty() vocab.LatitudePropertyInterface {
	return propertylatitude.NewLatitudeProperty()
}

// NewActivityStreamsRepliesProperty creates a new RepliesPropertyInterface
func NewActivityStreamsRepliesProperty() vocab.RepliesPropertyInterface {
	return propertyreplies.NewRepliesProperty()
}

// NewActivityStreamsLongitudeProperty creates a new LongitudePropertyInterface
func NewActivityStreamsLongitudeProperty() vocab.LongitudePropertyInterface {
	return propertylongitude.NewLongitudeProperty()
}

// NewActivityStreamsMediaTypeProperty creates a new MediaTypePropertyInterface
func NewActivityStreamsMediaTypeProperty() vocab.MediaTypePropertyInterface {
	return propertymediatype.NewMediaTypeProperty()
}

// NewActivityStreamsAccuracyProperty creates a new AccuracyPropertyInterface
func NewActivityStreamsAccuracyProperty() vocab.AccuracyPropertyInterface {
	return propertyaccuracy.NewAccuracyProperty()
}

// NewActivityStreamsWidthProperty creates a new WidthPropertyInterface
func NewActivityStreamsWidthProperty() vocab.WidthPropertyInterface {
	return propertywidth.NewWidthProperty()
}

// NewActivityStreamsDescribesProperty creates a new DescribesPropertyInterface
func NewActivityStreamsDescribesProperty() vocab.DescribesPropertyInterface {
	return propertydescribes.NewDescribesProperty()
}

// NewActivityStreamsPartOfProperty creates a new PartOfPropertyInterface
func NewActivityStreamsPartOfProperty() vocab.PartOfPropertyInterface {
	return propertypartof.NewPartOfProperty()
}

// NewActivityStreamsRadiusProperty creates a new RadiusPropertyInterface
func NewActivityStreamsRadiusProperty() vocab.RadiusPropertyInterface {
	return propertyradius.NewRadiusProperty()
}

// NewActivityStreamsStartIndexProperty creates a new StartIndexPropertyInterface
func NewActivityStreamsStartIndexProperty() vocab.StartIndexPropertyInterface {
	return propertystartindex.NewStartIndexProperty()
}

// NewActivityStreamsPrevProperty creates a new PrevPropertyInterface
func NewActivityStreamsPrevProperty() vocab.PrevPropertyInterface {
	return propertyprev.NewPrevProperty()
}

// NewActivityStreamsCurrentProperty creates a new CurrentPropertyInterface
func NewActivityStreamsCurrentProperty() vocab.CurrentPropertyInterface {
	return propertycurrent.NewCurrentProperty()
}

// NewActivityStreamsStartTimeProperty creates a new StartTimePropertyInterface
func NewActivityStreamsStartTimeProperty() vocab.StartTimePropertyInterface {
	return propertystarttime.NewStartTimeProperty()
}

// NewActivityStreamsLastProperty creates a new LastPropertyInterface
func NewActivityStreamsLastProperty() vocab.LastPropertyInterface {
	return propertylast.NewLastProperty()
}

// NewActivityStreamsAltitudeProperty creates a new AltitudePropertyInterface
func NewActivityStreamsAltitudeProperty() vocab.AltitudePropertyInterface {
	return propertyaltitude.NewAltitudeProperty()
}

// NewActivityStreamsHeightProperty creates a new HeightPropertyInterface
func NewActivityStreamsHeightProperty() vocab.HeightPropertyInterface {
	return propertyheight.NewHeightProperty()
}

// NewActivityStreamsPublishedProperty creates a new PublishedPropertyInterface
func NewActivityStreamsPublishedProperty() vocab.PublishedPropertyInterface {
	return propertypublished.NewPublishedProperty()
}

// NewActivityStreamsInstrumentProperty creates a new InstrumentPropertyInterface
func NewActivityStreamsInstrumentProperty() vocab.InstrumentPropertyInterface {
	return propertyinstrument.NewInstrumentProperty()
}

// NewActivityStreamsImageProperty creates a new ImagePropertyInterface
func NewActivityStreamsImageProperty() vocab.ImagePropertyInterface {
	return propertyimage.NewImageProperty()
}

// NewActivityStreamsObjectProperty creates a new ObjectPropertyInterface
func NewActivityStreamsObjectProperty() vocab.ObjectPropertyInterface {
	return propertyobject.NewObjectProperty()
}

// NewActivityStreamsContentProperty creates a new ContentPropertyInterface
func NewActivityStreamsContentProperty() vocab.ContentPropertyInterface {
	return propertycontent.NewContentProperty()
}

// NewActivityStreamsGeneratorProperty creates a new GeneratorPropertyInterface
func NewActivityStreamsGeneratorProperty() vocab.GeneratorPropertyInterface {
	return propertygenerator.NewGeneratorProperty()
}

// NewActivityStreamsBtoProperty creates a new BtoPropertyInterface
func NewActivityStreamsBtoProperty() vocab.BtoPropertyInterface {
	return propertybto.NewBtoProperty()
}

// NewActivityStreamsActorProperty creates a new ActorPropertyInterface
func NewActivityStreamsActorProperty() vocab.ActorPropertyInterface {
	return propertyactor.NewActorProperty()
}

// NewActivityStreamsLocationProperty creates a new LocationPropertyInterface
func NewActivityStreamsLocationProperty() vocab.LocationPropertyInterface {
	return propertylocation.NewLocationProperty()
}

// NewActivityStreamsAudienceProperty creates a new AudiencePropertyInterface
func NewActivityStreamsAudienceProperty() vocab.AudiencePropertyInterface {
	return propertyaudience.NewAudienceProperty()
}

// NewActivityStreamsResultProperty creates a new ResultPropertyInterface
func NewActivityStreamsResultProperty() vocab.ResultPropertyInterface {
	return propertyresult.NewResultProperty()
}

// NewActivityStreamsRelationshipProperty creates a new
// RelationshipPropertyInterface
func NewActivityStreamsRelationshipProperty() vocab.RelationshipPropertyInterface {
	return propertyrelationship.NewRelationshipProperty()
}

// NewActivityStreamsInReplyToProperty creates a new InReplyToPropertyInterface
func NewActivityStreamsInReplyToProperty() vocab.InReplyToPropertyInterface {
	return propertyinreplyto.NewInReplyToProperty()
}

// NewActivityStreamsOriginProperty creates a new OriginPropertyInterface
func NewActivityStreamsOriginProperty() vocab.OriginPropertyInterface {
	return propertyorigin.NewOriginProperty()
}

// NewActivityStreamsAnyOfProperty creates a new AnyOfPropertyInterface
func NewActivityStreamsAnyOfProperty() vocab.AnyOfPropertyInterface {
	return propertyanyof.NewAnyOfProperty()
}

// NewActivityStreamsTagProperty creates a new TagPropertyInterface
func NewActivityStreamsTagProperty() vocab.TagPropertyInterface {
	return propertytag.NewTagProperty()
}

// NewActivityStreamsNameProperty creates a new NamePropertyInterface
func NewActivityStreamsNameProperty() vocab.NamePropertyInterface {
	return propertyname.NewNameProperty()
}

// NewActivityStreamsContextProperty creates a new ContextPropertyInterface
func NewActivityStreamsContextProperty() vocab.ContextPropertyInterface {
	return propertycontext.NewContextProperty()
}

// NewActivityStreamsSummaryProperty creates a new SummaryPropertyInterface
func NewActivityStreamsSummaryProperty() vocab.SummaryPropertyInterface {
	return propertysummary.NewSummaryProperty()
}

// NewActivityStreamsUrlProperty creates a new UrlPropertyInterface
func NewActivityStreamsUrlProperty() vocab.UrlPropertyInterface {
	return propertyurl.NewUrlProperty()
}

// NewActivityStreamsAttributedToProperty creates a new
// AttributedToPropertyInterface
func NewActivityStreamsAttributedToProperty() vocab.AttributedToPropertyInterface {
	return propertyattributedto.NewAttributedToProperty()
}

// NewActivityStreamsOneOfProperty creates a new OneOfPropertyInterface
func NewActivityStreamsOneOfProperty() vocab.OneOfPropertyInterface {
	return propertyoneof.NewOneOfProperty()
}

// NewActivityStreamsItemsProperty creates a new ItemsPropertyInterface
func NewActivityStreamsItemsProperty() vocab.ItemsPropertyInterface {
	return propertyitems.NewItemsProperty()
}

// NewActivityStreamsTypeProperty creates a new TypePropertyInterface
func NewActivityStreamsTypeProperty() vocab.TypePropertyInterface {
	return propertytype.NewTypeProperty()
}

// NewActivityStreamsAttachmentProperty creates a new AttachmentPropertyInterface
func NewActivityStreamsAttachmentProperty() vocab.AttachmentPropertyInterface {
	return propertyattachment.NewAttachmentProperty()
}

// NewActivityStreamsTargetProperty creates a new TargetPropertyInterface
func NewActivityStreamsTargetProperty() vocab.TargetPropertyInterface {
	return propertytarget.NewTargetProperty()
}

// NewActivityStreamsFormerTypeProperty creates a new FormerTypePropertyInterface
func NewActivityStreamsFormerTypeProperty() vocab.FormerTypePropertyInterface {
	return propertyformertype.NewFormerTypeProperty()
}

// NewActivityStreamsPreviewProperty creates a new PreviewPropertyInterface
func NewActivityStreamsPreviewProperty() vocab.PreviewPropertyInterface {
	return propertypreview.NewPreviewProperty()
}

// NewActivityStreamsClosedProperty creates a new ClosedPropertyInterface
func NewActivityStreamsClosedProperty() vocab.ClosedPropertyInterface {
	return propertyclosed.NewClosedProperty()
}

// NewActivityStreamsToProperty creates a new ToPropertyInterface
func NewActivityStreamsToProperty() vocab.ToPropertyInterface {
	return propertyto.NewToProperty()
}

// NewActivityStreamsIconProperty creates a new IconPropertyInterface
func NewActivityStreamsIconProperty() vocab.IconPropertyInterface {
	return propertyicon.NewIconProperty()
}

// NewActivityStreamsRelProperty creates a new RelPropertyInterface
func NewActivityStreamsRelProperty() vocab.RelPropertyInterface {
	return propertyrel.NewRelProperty()
}

// NewActivityStreamsBccProperty creates a new BccPropertyInterface
func NewActivityStreamsBccProperty() vocab.BccPropertyInterface {
	return propertybcc.NewBccProperty()
}

// NewActivityStreamsCcProperty creates a new CcPropertyInterface
func NewActivityStreamsCcProperty() vocab.CcPropertyInterface {
	return propertycc.NewCcProperty()
}
