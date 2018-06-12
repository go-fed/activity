package defs

import (
	"bytes"
)

// Type definitions
const (
	baseURI = "https://www.w3.org/ns/activitystreams#"
)

var (
	objectType = &Type{
		Name:  "Object",
		URI:   baseURI + "Object",
		Notes: "Describes an object of any kind. The Object type serves as the base type for most of the other kinds of objects defined in the Activity Vocabulary, including other Core types such as Activity, IntransitiveActivity, Collection and OrderedCollection.",
		// Done within init: DisjointWith = []*Type{linkType}
	}
	linkType = &Type{
		Name:         "Link",
		URI:          baseURI + "Link",
		Notes:        "A Link is an indirect, qualified reference to a resource identified by a URL. The fundamental model for links is established by [ RFC5988]. Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link. When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the resource identified by the href. Properties of the Link are properties of the reference as opposed to properties of the resource.",
		DisjointWith: []*Type{objectType},
	}
	activityType = &Type{
		Name:    "Activity",
		URI:     baseURI + "Activity",
		Notes:   "An Activity is a subtype of Object that describes some form of action that may happen, is currently happening, or has already happened. The Activity type itself serves as an abstract base type for all types of activities. It is important to note that the Activity type itself does not carry any specific semantics about the kind of action being taken.",
		Extends: []*Type{objectType},
	}
	intransitiveActivityType = &Type{
		Name:              "IntransitiveActivity",
		URI:               baseURI + "IntransitiveActivity",
		Notes:             "Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions. The object property is therefore inappropriate for these activities.",
		Extends:           []*Type{activityType},
		WithoutProperties: []*PropertyType{objectPropertyType},
	}
	collectionType = &Type{
		Name:    "Collection",
		URI:     baseURI + "Collection",
		Notes:   "A Collection is a subtype of Object that represents ordered or unordered sets of Object or Link instances. Refer to the Activity Streams 2.0 Core specification for a complete description of the Collection type.",
		Extends: []*Type{objectType},
	}
	orderedCollectionType = &Type{
		Name:    "OrderedCollection",
		URI:     baseURI + "OrderedCollection",
		Notes:   "A subtype of Collection in which members of the logical collection are assumed to always be strictly ordered.",
		Extends: []*Type{collectionType},
		WithoutProperties: []*PropertyType{
			itemsPropertyType,
			currentPropertyType,
			firstPropertyType,
			lastPropertyType,
		},
	}
	collectionPageType = &Type{
		Name:    "CollectionPage",
		URI:     baseURI + "CollectionPage",
		Notes:   "Used to represent distinct subsets of items from a Collection. Refer to the Activity Streams 2.0 Core for a complete description of the CollectionPage object.",
		Extends: []*Type{collectionType},
	}
	orderedCollectionPageType = &Type{
		Name:    "OrderedCollectionPage",
		URI:     baseURI + "OrderedCollectionPage",
		Notes:   "Used to represent ordered subsets of items from an OrderedCollection. Refer to the Activity Streams 2.0 Core for a complete description of the OrderedCollectionPage object.",
		Extends: []*Type{orderedCollectionType, collectionPageType},
		WithoutProperties: []*PropertyType{
			itemsPropertyType,
			currentPropertyType,
			firstPropertyType,
			lastPropertyType,
			nextPropertyType,
			prevPropertyType,
		},
	}

	AllCoreTypes = []*Type{
		objectType,
		linkType,
		activityType,
		intransitiveActivityType,
		collectionType,
		orderedCollectionType,
		collectionPageType,
		orderedCollectionPageType,
	}
)

// Extended Type definitions
const (
	extendedBaseURI = "https://www.w3.org/ns/activitystreams#"
)

var (
	// Activity extensions
	acceptExtendedType = &Type{
		Name:    "Accept",
		URI:     extendedBaseURI + "Accept",
		Notes:   "Indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate the context into which the object has been accepted.",
		Extends: []*Type{activityType},
	}
	tentativeAcceptExtendedType = &Type{
		Name:    "TentativeAccept",
		URI:     extendedBaseURI + "TentativeAccept",
		Notes:   "A specialization of Accept indicating that the acceptance is tentative.",
		Extends: []*Type{acceptExtendedType},
	}
	addExtendedType = &Type{
		Name:    "Add",
		URI:     extendedBaseURI + "Add",
		Notes:   "Indicates that the actor has added the object to the target. If the target property is not explicitly specified, the target would need to be determined implicitly by context. The origin can be used to identify the context from which the object originated.",
		Extends: []*Type{activityType},
	}
	arriveExtendedType = &Type{
		Name:    "Arrive",
		URI:     extendedBaseURI + "Arrive",
		Notes:   "An IntransitiveActivity that indicates that the actor has arrived at the location. The origin can be used to identify the context from which the actor originated. The target typically has no defined meaning.",
		Extends: []*Type{intransitiveActivityType},
	}
	createExtendedType = &Type{
		Name:    "Create",
		URI:     extendedBaseURI + "Create",
		Notes:   "Indicates that the actor has created the object.",
		Extends: []*Type{activityType},
	}
	deleteExtendedType = &Type{
		Name:    "Delete",
		URI:     extendedBaseURI + "Delete",
		Notes:   "Indicates that the actor has deleted the object. If specified, the origin indicates the context from which the object was deleted.",
		Extends: []*Type{activityType},
	}
	followExtendedType = &Type{
		Name:    "Follow",
		URI:     extendedBaseURI + "Follow",
		Notes:   "Indicates that the actor is \"following\" the object. Following is defined in the sense typically used within Social systems in which the actor is interested in any activity performed by or on the object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	ignoreExtendedType = &Type{
		Name:    "Ignore",
		URI:     extendedBaseURI + "Ignore",
		Notes:   "Indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	joinExtendedType = &Type{
		Name:    "Join",
		URI:     extendedBaseURI + "Join",
		Notes:   "Indicates that the actor has joined the object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	leaveExtendedType = &Type{
		Name:    "Leave",
		URI:     extendedBaseURI + "Leave",
		Notes:   "Indicates that the actor has left the object. The target and origin typically have no meaning.",
		Extends: []*Type{activityType},
	}
	likeExtendedType = &Type{
		Name:    "Like",
		URI:     extendedBaseURI + "Like",
		Notes:   "Indicates that the actor likes, recommends or endorses the object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	offerExtendedType = &Type{
		Name:    "Offer",
		URI:     extendedBaseURI + "Offer",
		Notes:   "Indicates that the actor is offering the object. If specified, the target indicates the entity to which the object is being offered.",
		Extends: []*Type{activityType},
	}
	inviteExtendedType = &Type{
		Name:    "Invite",
		URI:     extendedBaseURI + "Invite",
		Notes:   "A specialization of Offer in which the actor is extending an invitation for the object to the target.",
		Extends: []*Type{offerExtendedType},
	}
	rejectExtendedType = &Type{
		Name:    "Reject",
		URI:     extendedBaseURI + "Reject",
		Notes:   "Indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	tentativeRejectExtendedType = &Type{
		Name:    "TentativeReject",
		URI:     extendedBaseURI + "TentativeReject",
		Notes:   "A specialization of Reject in which the rejection is considered tentative.",
		Extends: []*Type{rejectExtendedType},
	}
	removeExtendedType = &Type{
		Name:    "Remove",
		URI:     extendedBaseURI + "Remove",
		Notes:   "Indicates that the actor is removing the object. If specified, the origin indicates the context from which the object is being removed.",
		Extends: []*Type{activityType},
	}
	undoExtendedType = &Type{
		Name:    "Undo",
		URI:     extendedBaseURI + "Undo",
		Notes:   "Indicates that the actor is undoing the object. In most cases, the object will be an Activity describing some previously performed action (for instance, a person may have previously \"liked\" an article but, for whatever reason, might choose to undo that like at some later point in time). The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	updateExtendedType = &Type{
		Name:    "Update",
		URI:     extendedBaseURI + "Update",
		Notes:   "Indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object. The target and origin typically have no defined meaning.",
		Extends: []*Type{activityType},
	}
	viewExtendedType = &Type{
		Name:    "View",
		URI:     extendedBaseURI + "View",
		Notes:   "Indicates that the actor has viewed the object.",
		Extends: []*Type{activityType},
	}
	listenExtendedType = &Type{
		Name:    "Listen",
		URI:     extendedBaseURI + "Listen",
		Notes:   "Indicates that the actor has listened to the object.",
		Extends: []*Type{activityType},
	}
	readExtendedType = &Type{
		Name:    "Read",
		URI:     extendedBaseURI + "Read",
		Notes:   "Indicates that the actor has read the object.",
		Extends: []*Type{activityType},
	}
	moveExtendedType = &Type{
		Name:    "Move",
		URI:     extendedBaseURI + "Move",
		Notes:   "Indicates that the actor has moved object from origin to target. If the origin or target are not specified, either can be determined by context.",
		Extends: []*Type{activityType},
	}
	travelExtendedType = &Type{
		Name:    "Travel",
		URI:     extendedBaseURI + "Travel",
		Notes:   "Indicates that the actor is traveling to target from origin. Travel is an IntransitiveObject whose actor specifies the direct object. If the target or origin are not specified, either can be determined by context.",
		Extends: []*Type{intransitiveActivityType},
	}
	announceExtendedType = &Type{
		Name:    "Announce",
		URI:     extendedBaseURI + "Announce",
		Notes:   "Indicates that the actor is calling the target's attention the object. The origin typically has no defined meaning.",
		Extends: []*Type{activityType},
	}
	blockExtendedType = &Type{
		Name:    "Block",
		URI:     extendedBaseURI + "Block",
		Notes:   "Indicates that the actor is blocking the object. Blocking is a stronger form of Ignore. The typical use is to support social systems that allow one user to block activities or content of other users. The target and origin typically have no defined meaning.",
		Extends: []*Type{ignoreExtendedType},
	}
	flagExtendedType = &Type{
		Name:    "Flag",
		URI:     extendedBaseURI + "Flag",
		Notes:   "Indicates that the actor is \"flagging\" the object. Flagging is defined in the sense common to many social platforms as reporting content as being inappropriate for any number of reasons.",
		Extends: []*Type{activityType},
	}
	dislikeExtendedType = &Type{
		Name:    "Dislike",
		URI:     extendedBaseURI + "Dislike",
		Notes:   "Indicates that the actor dislikes the object.",
		Extends: []*Type{activityType},
	}
	questionExtendedType = &Type{
		Name:    "Question",
		URI:     extendedBaseURI + "Question",
		Notes:   "Represents a question being asked. Question objects are an extension of IntransitiveActivity. That is, the Question object is an Activity, but the direct object is the question itself and therefore it would not contain an object property. Either of the anyOf and oneOf properties may be used to express possible answers, but a Question object must not have both properties.",
		Extends: []*Type{intransitiveActivityType},
	}

	// Actor extensions
	applicationExtendedType = &Type{
		Name:    "Application",
		URI:     extendedBaseURI + "Application",
		Notes:   "Describes a software application.",
		Extends: []*Type{objectType},
	}
	groupExtendedType = &Type{
		Name:    "Group",
		URI:     extendedBaseURI + "Group",
		Notes:   "Represents a formal or informal collective of Actors.",
		Extends: []*Type{objectType},
	}
	organizationExtendedType = &Type{
		Name:    "Organization",
		URI:     extendedBaseURI + "Organization",
		Notes:   "Represents an organization.",
		Extends: []*Type{objectType},
	}
	personExtendedType = &Type{
		Name:    "Person",
		URI:     extendedBaseURI + "Person",
		Notes:   "Represents an individual person.",
		Extends: []*Type{objectType},
	}
	serviceExtendedType = &Type{
		Name:    "Service",
		URI:     extendedBaseURI + "Service",
		Notes:   "Represents a service of any kind.",
		Extends: []*Type{objectType},
	}

	// Object extensions
	relationshipExtendedType = &Type{
		Name:    "Relationship",
		URI:     extendedBaseURI + "Relationship",
		Notes:   "Describes a relationship between two individuals. The subject and object properties are used to identify the connected individuals. See 5.2 Representing Relationships Between Entities for additional information.",
		Extends: []*Type{objectType},
	}
	articleExtendedType = &Type{
		Name:    "Article",
		URI:     extendedBaseURI + "Article",
		Notes:   "Represents any kind of multi-paragraph written work.",
		Extends: []*Type{objectType},
	}
	documentExtendedType = &Type{
		Name:    "Document",
		URI:     extendedBaseURI + "Document",
		Notes:   "Represents a document of any kind.",
		Extends: []*Type{objectType},
	}
	audioExtendedType = &Type{
		Name:    "Audio",
		URI:     extendedBaseURI + "Audio",
		Notes:   "Represents an audio document of any kind.",
		Extends: []*Type{documentExtendedType},
	}
	imageExtendedType = &Type{
		Name:    "Image",
		URI:     extendedBaseURI + "Image",
		Notes:   "An image document of any kind",
		Extends: []*Type{documentExtendedType},
	}
	videoExtendedType = &Type{
		Name:    "Video",
		URI:     extendedBaseURI + "Video",
		Notes:   "Represents a video document of any kind.",
		Extends: []*Type{documentExtendedType},
	}
	noteExtendedType = &Type{
		Name:    "Note",
		URI:     extendedBaseURI + "Note",
		Notes:   "Represents a short written work typically less than a single paragraph in length.",
		Extends: []*Type{objectType},
	}
	pageExtendedType = &Type{
		Name:    "Page",
		URI:     extendedBaseURI + "Page",
		Notes:   "Represents a Web Page.",
		Extends: []*Type{documentExtendedType},
	}
	eventExtendedType = &Type{
		Name:    "Event",
		URI:     extendedBaseURI + "Event",
		Notes:   "Represents any kind of event.",
		Extends: []*Type{objectType},
	}
	placeExtendedType = &Type{
		Name:    "Place",
		URI:     extendedBaseURI + "Place",
		Notes:   "Represents a logical or physical location. See 5.3 Representing Places for additional information.",
		Extends: []*Type{objectType},
	}
	profileExtendedType = &Type{
		Name:    "Profile",
		URI:     extendedBaseURI + "Profile",
		Notes:   "A Profile is a content object that describes another Object, typically used to describe Actor Type objects. The describes property is used to reference the object being described by the profile.",
		Extends: []*Type{objectType},
	}
	tombstoneExtendedType = &Type{
		Name:    "Tombstone",
		URI:     extendedBaseURI + "Tombstone",
		Notes:   "A Tombstone represents a content object that has been deleted. It can be used in Collections to signify that there used to be an object at this position, but it has been deleted.",
		Extends: []*Type{objectType},
	}

	// Link extensions
	mentionExtendedType = &Type{
		Name:    "Mention",
		URI:     extendedBaseURI + "Mention",
		Notes:   "A specialized Link that represents an @mention.",
		Extends: []*Type{linkType},
	}

	AllExtendedTypes = []*Type{
		acceptExtendedType,
		tentativeAcceptExtendedType,
		addExtendedType,
		arriveExtendedType,
		createExtendedType,
		deleteExtendedType,
		followExtendedType,
		ignoreExtendedType,
		joinExtendedType,
		leaveExtendedType,
		likeExtendedType,
		offerExtendedType,
		inviteExtendedType,
		rejectExtendedType,
		tentativeRejectExtendedType,
		removeExtendedType,
		undoExtendedType,
		updateExtendedType,
		viewExtendedType,
		listenExtendedType,
		readExtendedType,
		moveExtendedType,
		travelExtendedType,
		announceExtendedType,
		blockExtendedType,
		flagExtendedType,
		dislikeExtendedType,
		questionExtendedType,
		applicationExtendedType,
		groupExtendedType,
		organizationExtendedType,
		personExtendedType,
		serviceExtendedType,
		relationshipExtendedType,
		articleExtendedType,
		documentExtendedType,
		audioExtendedType,
		imageExtendedType,
		videoExtendedType,
		noteExtendedType,
		pageExtendedType,
		eventExtendedType,
		placeExtendedType,
		profileExtendedType,
		tombstoneExtendedType,
		mentionExtendedType,
	}
)

// PropertyType definitions
const (
	propertyBaseURI = "https://www.w3.org/ns/activitystreams#"
)

var (
	idPropertyType = &PropertyType{
		Name:       "id",
		URI:        "@id",
		Notes:      "Provides the globally unique identifier for an Object or Link.",
		Domain:     []DomainReference{{T: objectType}, {T: linkType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true,
	}
	typePropertyType = &PropertyType{
		Name:   "type",
		URI:    "@type",
		Notes:  "Identifies the Object or Link type. Multiple values may be specified. Note that when serializing, the appropriate Activitystream type is added if it is not already added by the caller.",
		Domain: []DomainReference{{T: objectType}, {T: linkType}},
		Range:  []RangeReference{{Any: true}},
	}
	actorPropertyType = &PropertyType{
		Name:                 "actor",
		URI:                  propertyBaseURI + "actor",
		Notes:                "Describes one or more entities that either performed or are expected to perform the activity. Any single activity can have multiple actors. The actor may be specified using an indirect Link.",
		Domain:               []DomainReference{{T: activityType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		SubpropertyOf:        attributedToPropertyType,
		PreferIRIConvenience: true,
	}
	attachmentPropertyType = &PropertyType{
		Name:   "attachment",
		URI:    propertyBaseURI + "attachment",
		Notes:  "Identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	attributedToPropertyType = &PropertyType{
		Name:                 "attributedTo",
		URI:                  propertyBaseURI + "attributedTo",
		Notes:                "Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.",
		Domain:               []DomainReference{{T: objectType}, {T: linkType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	audiencePropertyType = &PropertyType{
		Name:                 "audience",
		URI:                  propertyBaseURI + "audience",
		Notes:                "Identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	bccPropertyType = &PropertyType{
		Name:                 "bcc",
		URI:                  propertyBaseURI + "bcc",
		Notes:                "Identifies one or more Objects that are part of the private secondary audience of this Object.",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	btoPropertyType = &PropertyType{
		Name:                 "bto",
		URI:                  propertyBaseURI + "bto",
		Notes:                "Identifies an Object that is part of the private primary audience of this Object.",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	ccPropertyType = &PropertyType{
		Name:                 "cc",
		URI:                  propertyBaseURI + "cc",
		Notes:                "Identifies an Object that is part of the public secondary audience of this Object.",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	contextPropertyType = &PropertyType{
		Name:   "context",
		URI:    propertyBaseURI + "context",
		Notes:  "Identifies the context within which the object exists or an activity was performed. The notion of \"context\" used is intentionally vague. The intended function is to serve as a means of grouping objects and activities that share a common originating context or purpose. An example could be all activities relating to a common project or event.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	currentPropertyType = &PropertyType{
		Name:                 "current",
		URI:                  propertyBaseURI + "current",
		Notes:                "In a paged Collection, indicates the page that contains the most recently updated member items.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: collectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	currentOrderedPropertyType = &PropertyType{
		Name:                 "current",
		URI:                  propertyBaseURI + "current",
		Notes:                "In a paged OrderedCollection, indicates the page that contains the most recently updated member items.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: orderedCollectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	firstPropertyType = &PropertyType{
		Name:                 "first",
		URI:                  propertyBaseURI + "first",
		Notes:                "In a paged Collection, indicates the furthest preceeding page of items in the collection.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: collectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	firstOrderedPropertyType = &PropertyType{
		Name:                 "first",
		URI:                  propertyBaseURI + "first",
		Notes:                "In a paged OrderedCollection, indicates the furthest preceeding page of items in the collection.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: orderedCollectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	generatorPropertyType = &PropertyType{
		Name:   "generator",
		URI:    propertyBaseURI + "generator",
		Notes:  "Identifies the entity (e.g. an application) that generated the object.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	iconPropertyType = &PropertyType{
		Name:   "icon",
		URI:    propertyBaseURI + "icon",
		Notes:  "Indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: imageExtendedType}, {T: linkType}},
	}
	imagePropertyType = &PropertyType{
		Name:   "image",
		URI:    propertyBaseURI + "image",
		Notes:  "Indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: imageExtendedType}, {T: linkType}},
	}
	inReplyToPropertyType = &PropertyType{
		Name:                 "inReplyTo",
		URI:                  propertyBaseURI + "inReplyTo",
		Notes:                "Indicates one or more entities for which this object is considered a response.",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	instrumentPropertyType = &PropertyType{
		Name:   "instrument",
		URI:    propertyBaseURI + "instrument",
		Notes:  "Identifies one or more objects used (or to be used) in the completion of an Activity.",
		Domain: []DomainReference{{T: activityType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	lastPropertyType = &PropertyType{
		Name:                 "last",
		URI:                  propertyBaseURI + "last",
		Notes:                "In a paged Collection, indicates the furthest proceeding page of the collection.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: collectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	lastOrderedPropertyType = &PropertyType{
		Name:                 "last",
		URI:                  propertyBaseURI + "last",
		Notes:                "In a paged OrderedCollection, indicates the furthest proceeding page of the collection.",
		Domain:               []DomainReference{{T: collectionType}},
		Range:                []RangeReference{{T: orderedCollectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	locationPropertyType = &PropertyType{
		Name:   "location",
		URI:    propertyBaseURI + "location",
		Notes:  "Indicates one or more physical or logical locations associated with the object.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	itemsPropertyType = &PropertyType{
		Name:   "items",
		URI:    propertyBaseURI + "items",
		Notes:  "Identifies the items contained in a collection. The items might be ordered or unordered.",
		Domain: []DomainReference{{T: collectionType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	orderedItemsPropertyType = &PropertyType{ // Implicit within spec!
		Name:   "orderedItems",
		URI:    propertyBaseURI + "items",
		Notes:  "Identifies the items contained in a collection. The items might be ordered or unordered.",
		Domain: []DomainReference{{T: orderedCollectionType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	oneOfPropertyType = &PropertyType{
		Name:   "oneOf",
		URI:    propertyBaseURI + "oneOf",
		Notes:  "Identifies an exclusive option for a Question. Use of oneOf implies that the Question can have only a single answer. To indicate that a Question can have multiple answers, use anyOf.",
		Domain: []DomainReference{{T: questionExtendedType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	anyOfPropertyType = &PropertyType{
		Name:   "anyOf",
		URI:    propertyBaseURI + "anyOf",
		Notes:  "Identifies an inclusive option for a Question. Use of anyOf implies that the Question can have multiple answers. To indicate that a Question can have only one answer, use oneOf.",
		Domain: []DomainReference{{T: questionExtendedType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	closedPropertyType = &PropertyType{
		Name:   "closed",
		URI:    propertyBaseURI + "closed",
		Notes:  "Indicates that a question has been closed, and answers are no longer accepted.",
		Domain: []DomainReference{{T: questionExtendedType}},
		Range:  []RangeReference{{V: xsdDateTimeValueType}, {V: xsdBooleanValueType}, {T: objectType}, {T: linkType}},
	}
	originPropertyType = &PropertyType{
		Name:   "origin",
		URI:    propertyBaseURI + "origin",
		Notes:  "Describes an indirect object of the activity from which the activity is directed. The precise meaning of the origin is the object of the English preposition \"from\". For instance, in the activity \"John moved an item to List B from List A\", the origin of the activity is \"List A\".",
		Domain: []DomainReference{{T: activityType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	nextPropertyType = &PropertyType{
		Name:                 "next",
		URI:                  propertyBaseURI + "next",
		Notes:                "In a paged Collection, indicates the next page of items.",
		Domain:               []DomainReference{{T: collectionPageType}},
		Range:                []RangeReference{{T: collectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	nextOrderedPropertyType = &PropertyType{
		Name:                 "next",
		URI:                  propertyBaseURI + "next",
		Notes:                "In a paged OrderedCollection, indicates the next page of items.",
		Domain:               []DomainReference{{T: orderedCollectionPageType}},
		Range:                []RangeReference{{T: orderedCollectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	objectPropertyType = &PropertyType{
		Name:   "object",
		URI:    propertyBaseURI + "object",
		Notes:  "When used within an Activity, describes the direct object of the activity. For instance, in the activity \"John added a movie to his wishlist\", the object of the activity is the movie added. When used within a Relationship describes the entity to which the subject is related.",
		Domain: []DomainReference{{T: activityType}, {T: relationshipExtendedType}},
		Range:  []RangeReference{{T: objectType}},
	}
	prevPropertyType = &PropertyType{
		Name:                 "prev",
		URI:                  propertyBaseURI + "prev",
		Notes:                "In a paged Collection, identifies the previous page of items.",
		Domain:               []DomainReference{{T: collectionPageType}},
		Range:                []RangeReference{{T: collectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	prevOrderedPropertyType = &PropertyType{
		Name:                 "prev",
		URI:                  propertyBaseURI + "prev",
		Notes:                "In a paged OrderedCollection, identifies the previous page of items.",
		Domain:               []DomainReference{{T: orderedCollectionPageType}},
		Range:                []RangeReference{{T: orderedCollectionPageType}, {T: linkType}},
		Functional:           true,
		PreferIRIConvenience: true,
	}
	previewPropertyType = &PropertyType{
		Name:   "preview",
		URI:    propertyBaseURI + "preview",
		Notes:  "Identifies an entity that provides a preview of this object.",
		Domain: []DomainReference{{T: objectType}, {T: linkType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	resultPropertyType = &PropertyType{
		Name:   "result",
		URI:    propertyBaseURI + "result",
		Notes:  "Describes the result of the activity. For instance, if a particular action results in the creation of a new resource, the result property can be used to describe that new resource.",
		Domain: []DomainReference{{T: activityType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	repliesPropertyType = &PropertyType{
		Name:       "replies",
		URI:        propertyBaseURI + "replies",
		Notes:      "Identifies a Collection containing objects considered to be responses to this object.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: collectionType}},
		Functional: true,
	}
	tagPropertyType = &PropertyType{
		Name:   "tag",
		URI:    propertyBaseURI + "tag",
		Notes:  "One or more \"tags\" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{T: objectType}, {T: linkType}},
	}
	targetPropertyType = &PropertyType{
		Name:                 "target",
		URI:                  propertyBaseURI + "target",
		Notes:                "Describes the indirect object, or target, of the activity. The precise meaning of the target is largely dependent on the type of action being described but will often be the object of the English preposition \"to\". For instance, in the activity \"John added a movie to his wishlist\", the target of the activity is John's wishlist. An activity can have more than one target.",
		Domain:               []DomainReference{{T: activityType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	toPropertyType = &PropertyType{
		Name:                 "to",
		URI:                  propertyBaseURI + "to",
		Notes:                "Identifies an entity considered to be part of the public primary audience of an Object",
		Domain:               []DomainReference{{T: objectType}},
		Range:                []RangeReference{{T: objectType}, {T: linkType}},
		PreferIRIConvenience: true,
	}
	urlPropertyType = &PropertyType{
		Name:   "url",
		URI:    propertyBaseURI + "url",
		Notes:  "Identifies one or more links to representations of the object",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{V: xsdAnyURIValueType}, {T: linkType}},
	}
	accuracyPropertyType = &PropertyType{
		Name:       "accuracy",
		URI:        propertyBaseURI + "accuracy",
		Notes:      "Indicates the accuracy of position coordinates on a Place objects. Expressed in properties of percentage. e.g. \"94.0\" means \"94.0%% accurate\".",
		Domain:     []DomainReference{{T: placeExtendedType}},
		Range:      []RangeReference{{V: xsdFloatValueType}},
		Functional: true,
	}
	altitudePropertyType = &PropertyType{
		Name:       "altitude",
		URI:        propertyBaseURI + "altitude",
		Notes:      "ndicates the altitude of a place. The measurement units is indicated using the units property. If units is not specified, the default is assumed to be \"m\" indicating meters.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdFloatValueType}},
		Functional: true,
	}
	contentPropertyType = &PropertyType{
		Name:               "content",
		URI:                propertyBaseURI + "content",
		Notes:              "The content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type. The content may be expressed using multiple language-tagged values.",
		NaturalLanguageMap: true,
		Domain:             []DomainReference{{T: objectType}},
		Range:              []RangeReference{{V: xsdStringValueType}, {V: rdfLangStringValueType}},
	}
	namePropertyType = &PropertyType{
		Name:               "name",
		URI:                propertyBaseURI + "name",
		Notes:              "A simple, human-readable, plain-text name for the object. HTML markup must not be included. The name may be expressed using multiple language-tagged values.",
		NaturalLanguageMap: true,
		Domain:             []DomainReference{{T: objectType}, {T: linkType}},
		Range:              []RangeReference{{V: xsdStringValueType}, {V: rdfLangStringValueType}},
	}
	durationPropertyType = &PropertyType{
		Name:       "duration",
		URI:        propertyBaseURI + "duration",
		Notes:      "When the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value must be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as \"PT5S\").",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdDurationValueType}},
		Functional: true,
	}
	heightPropertyType = &PropertyType{
		Name:       "height",
		URI:        propertyBaseURI + "height",
		Notes:      "On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.",
		Domain:     []DomainReference{{T: linkType}, {T: imageExtendedType}},
		Range:      []RangeReference{{V: xsdNonNegativeIntegerValueType}},
		Functional: true,
	}
	hrefPropertyType = &PropertyType{
		Name:       "href",
		URI:        propertyBaseURI + "href",
		Notes:      "The target resource pointed to by a Link.",
		Domain:     []DomainReference{{T: linkType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true,
	}
	hrefLangPropertyType = &PropertyType{
		Name:       "hreflang",
		URI:        propertyBaseURI + "hreflang",
		Notes:      "Hints as to the language used by the target resource. Value must be a [BCP47] Language-Tag.",
		Domain:     []DomainReference{{T: linkType}},
		Range:      []RangeReference{{V: bcp47LangTag}},
		Functional: true,
	}
	partOfPropertyType = &PropertyType{
		Name:       "partOf",
		URI:        propertyBaseURI + "partOf",
		Notes:      "Identifies the Collection to which a CollectionPage objects items belong.",
		Domain:     []DomainReference{{T: collectionPageType}},
		Range:      []RangeReference{{T: linkType}, {T: collectionType}},
		Functional: true,
	}
	latitudePropertyType = &PropertyType{
		Name:       "latitude",
		URI:        propertyBaseURI + "latitude",
		Notes:      "The latitude of a place",
		Domain:     []DomainReference{{T: placeExtendedType}},
		Range:      []RangeReference{{V: xsdFloatValueType}},
		Functional: true,
	}
	longitudePropertyType = &PropertyType{
		Name:       "longitude",
		URI:        propertyBaseURI + "longitude",
		Notes:      "The longitude of a place",
		Domain:     []DomainReference{{T: placeExtendedType}},
		Range:      []RangeReference{{V: xsdFloatValueType}},
		Functional: true,
	}
	mediaTypePropertyType = &PropertyType{
		Name:       "mediaType",
		URI:        propertyBaseURI + "mediaType",
		Notes:      "When used on a Link, identifies the MIME media type of the referenced resource. When used on an Object, identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.",
		Domain:     []DomainReference{{T: linkType}, {T: objectType}},
		Range:      []RangeReference{{V: mimeMediaValueType}},
		Functional: true,
	}
	endTimePropertyType = &PropertyType{
		Name:       "endTime",
		URI:        propertyBaseURI + "endTime",
		Notes:      "The date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdDateTimeValueType}},
		Functional: true,
	}
	publishedPropertyType = &PropertyType{
		Name:       "published",
		URI:        propertyBaseURI + "published",
		Notes:      "The date and time at which the object was published",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdDateTimeValueType}},
		Functional: true,
	}
	startTimePropertyType = &PropertyType{
		Name:       "startTime",
		URI:        propertyBaseURI + "startTime",
		Notes:      "The date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdDateTimeValueType}},
		Functional: true,
	}
	radiusPropertyType = &PropertyType{
		Name:       "radius",
		URI:        propertyBaseURI + "property",
		Notes:      "The radius from the given latitude and longitude for a Place. The units is expressed by the units property. If units is not specified, the default is assumed to be \"m\" indicating \"meters\".",
		Domain:     []DomainReference{{T: placeExtendedType}},
		Range:      []RangeReference{{V: xsdFloatValueType}},
		Functional: true,
	}
	relPropertyType = &PropertyType{
		Name:   "rel",
		URI:    propertyBaseURI + "rel",
		Notes:  "A link relation associated with a Link. The value must conform to both the [HTML5] and [RFC5988] \"link relation\" definitions. In the [HTML5], any string not containing the \"space\" U+0020, \"tab\" (U+0009), \"LF\" (U+000A), \"FF\" (U+000C), \"CR\" (U+000D) or \",\" (U+002C) characters can be used as a valid link relation.",
		Domain: []DomainReference{{T: linkType}},
		Range:  []RangeReference{{V: linkRelationValueType}},
	}
	startIndexPropertyType = &PropertyType{
		Name:       "startIndex",
		URI:        propertyBaseURI + "startIndex",
		Notes:      "A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.",
		Domain:     []DomainReference{{T: orderedCollectionPageType}},
		Range:      []RangeReference{{V: xsdNonNegativeIntegerValueType}},
		Functional: true,
	}
	summaryPropertyType = &PropertyType{
		Name:               "summary",
		URI:                propertyBaseURI + "summary",
		Notes:              "A natural language summarization of the object encoded as HTML. Multiple language tagged summaries may be provided.",
		NaturalLanguageMap: true,
		Domain:             []DomainReference{{T: objectType}, {T: linkType}}, // Link conflicts with spec!
		Range:              []RangeReference{{V: xsdStringValueType}, {V: rdfLangStringValueType}},
	}
	totalItemsPropertyType = &PropertyType{
		Name:       "totalItems",
		URI:        propertyBaseURI + "totalItems",
		Notes:      "A non-negative integer specifying the total number of objects contained by the logical view of the collection. This number might not reflect the actual number of items serialized within the Collection object instance.",
		Domain:     []DomainReference{{T: collectionType}},
		Range:      []RangeReference{{V: xsdNonNegativeIntegerValueType}},
		Functional: true,
	}
	unitsPropertyType = &PropertyType{
		Name:       "units",
		URI:        propertyBaseURI + "units",
		Notes:      "Specifies the measurement units for the radius and altitude properties on a Place object. If not specified, the default is assumed to be \"m\" for \"meters\".",
		Domain:     []DomainReference{{T: placeExtendedType}},
		Range:      []RangeReference{{V: unitsValueType}, {V: xsdAnyURIValueType}},
		Functional: true,
	}
	updatedPropertyType = &PropertyType{
		Name:       "updated",
		URI:        propertyBaseURI + "updated",
		Notes:      "The date and time at which the object was updated",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdDateTimeValueType}},
		Functional: true,
	}
	widthPropertyType = &PropertyType{
		Name:       "width",
		URI:        propertyBaseURI + "width",
		Notes:      "On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.",
		Domain:     []DomainReference{{T: linkType}, {T: imageExtendedType}},
		Range:      []RangeReference{{V: xsdNonNegativeIntegerValueType}},
		Functional: true,
	}
	subjectPropertyType = &PropertyType{
		Name:       "subject",
		URI:        propertyBaseURI + "subject",
		Notes:      "On a Relationship object, the subject property identifies one of the connected individuals. For instance, for a Relationship object describing \"John is related to Sally\", subject would refer to John.",
		Domain:     []DomainReference{{T: relationshipExtendedType}},
		Range:      []RangeReference{{T: objectType}, {T: linkType}},
		Functional: true,
	}
	relationshipPropertyType = &PropertyType{
		Name:       "relationship",
		URI:        propertyBaseURI + "relationship",
		Notes:      "On a Relationship object, the relationship property identifies the kind of relationship that exists between subject and object.",
		Domain:     []DomainReference{{T: relationshipExtendedType}},
		Range:      []RangeReference{{T: objectType}},
		Functional: true,
	}
	describesPropertyType = &PropertyType{
		Name:       "describes",
		URI:        propertyBaseURI + "describes",
		Notes:      "On a Profile object, the describes property identifies the object described by the Profile.",
		Domain:     []DomainReference{{T: profileExtendedType}},
		Range:      []RangeReference{{T: objectType}},
		Functional: true,
	}
	formerTypePropertyType = &PropertyType{
		Name:   "formerType",
		URI:    propertyBaseURI + "formerType",
		Notes:  "On a Tombstone object, the formerType property identifies the type of the object that was deleted.",
		Domain: []DomainReference{{T: tombstoneExtendedType}},
		Range:  []RangeReference{{V: xsdStringValueType}, {T: objectType}}, // xsd:String not in spec!
	}
	deletedPropertyType = &PropertyType{
		Name:       "deleted",
		URI:        propertyBaseURI + "deleted",
		Notes:      "On a Tombstone object, the deleted property is a timestamp for when the object was deleted.",
		Domain:     []DomainReference{{T: tombstoneExtendedType}},
		Range:      []RangeReference{{V: xsdDateTimeValueType}},
		Functional: true,
	}

	// Specific to ActivityPub specification
	sourcePropertyType = &PropertyType{
		Name:       "source",
		URI:        propertyBaseURI + "source", // Missing from spec!
		Notes:      "The source property is intended to convey some sort of source from which the content markup was derived, as a form of provenance, or to support future editing by clients.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: objectType}},
		Functional: true, // Missing from spec!
	}
	inboxPropertyType = &PropertyType{
		Name:       "inbox",
		URI:        propertyBaseURI + "inbox", // Missing from spec!
		Notes:      "A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	outboxPropertyType = &PropertyType{
		Name:       "outbox",
		URI:        propertyBaseURI + "outbox", // Missing from spec!
		Notes:      "An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	followingPropertyType = &PropertyType{
		Name:       "following",
		URI:        propertyBaseURI + "following",
		Notes:      "A link to an [ActivityStreams] collection of the actors that this actor is following",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: collectionType}, {T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	followersPropertyType = &PropertyType{
		Name:       "followers",
		URI:        propertyBaseURI + "followers",
		Notes:      "A link to an [ActivityStreams] collection of the actors that follow this actor",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: collectionType}, {T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	likedPropertyType = &PropertyType{
		Name:       "liked",
		URI:        propertyBaseURI + "liked",
		Notes:      "A link to an [ActivityStreams] collection of objects this actor has liked",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: collectionType}, {T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	likesPropertyType = &PropertyType{
		Name:       "likes",
		URI:        propertyBaseURI + "likes",
		Notes:      "A link to an [ActivityStreams] collection of objects this actor has liked",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: collectionType}, {T: orderedCollectionType}, {V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	streamsPropertyType = &PropertyType{
		Name:   "streams",
		URI:    propertyBaseURI + "streams",
		Notes:  "A list of supplementary Collections which may be of interest.",
		Domain: []DomainReference{{T: objectType}},
		Range:  []RangeReference{{V: xsdAnyURIValueType}},
	}
	preferredUsernamePropertyType = &PropertyType{
		Name:               "preferredUsername",
		URI:                propertyBaseURI + "preferredUsername",
		Notes:              "A short username which may be used to refer to the actor, with no uniqueness guarantees.",
		NaturalLanguageMap: true,
		Domain:             []DomainReference{{T: objectType}},
		Range:              []RangeReference{{V: xsdStringValueType}},
		Functional:         true, // Missing from spec!
	}
	endpointsPropertyType = &PropertyType{
		Name:       "endpoints",
		URI:        propertyBaseURI + "endpoints",
		Notes:      "A json object which maps additional (typically server/domain-wide) endpoints which may be useful either for this actor or someone referencing this actor. This mapping may be nested inside the actor document as the value or may be a link to a JSON-LD document with these properties.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{T: objectType}},
		Functional: true, // Missing from spec!
	}
	proxyUrlPropertyType = &PropertyType{
		Name:       "proxyUrl",
		URI:        propertyBaseURI + "proxyUrl",
		Notes:      "Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being the id of the requested ActivityStreams object.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	oauthAuthorizationEndpointPropertyType = &PropertyType{
		Name:       "oauthAuthorizationEndpoint",
		URI:        propertyBaseURI + "oauthAuthorizationEndpoint",
		Notes:      "If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions, this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	oauthTokenEndpointPropertyType = &PropertyType{
		Name:       "oauthTokenEndpoint",
		URI:        propertyBaseURI + "oauthTokenEndpoint",
		Notes:      "If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions, this endpoint specifies a URI at which a client may acquire an access token.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	provideClientKeyPropertyType = &PropertyType{
		Name:       "provideClientKey",
		URI:        propertyBaseURI + "provideClientKey",
		Notes:      "If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization, this endpoint specifies a URI at which browser-authenticated users may authorize a client's public key for client to server interactions.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	signClientKeyPropertyType = &PropertyType{
		Name:       "signClientKey",
		URI:        propertyBaseURI + "signClientKey",
		Notes:      "If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization, this endpoint specifies a URI at which a client key may be signed by the actor's key for a time window to act on behalf of the actor in interacting with foreign servers.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}
	sharedInboxPropertyType = &PropertyType{
		Name:       "sharedInbox",
		URI:        propertyBaseURI + "sharedInbox",
		Notes:      "An optional endpoint used for wide delivery of publicly addressed activities and activities sent to followers. sharedInbox endpoints SHOULD also be publicly readable OrderedCollection objects containing objects addressed to the Public special collection. Reading from the sharedInbox endpoint MUST NOT present objects which are not addressed to the Public endpoint.",
		Domain:     []DomainReference{{T: objectType}},
		Range:      []RangeReference{{V: xsdAnyURIValueType}},
		Functional: true, // Missing from spec!
	}

	AllPropertyTypes = []*PropertyType{
		idPropertyType,
		typePropertyType,
		actorPropertyType,
		attachmentPropertyType,
		attributedToPropertyType,
		audiencePropertyType,
		bccPropertyType,
		btoPropertyType,
		ccPropertyType,
		contextPropertyType,
		currentPropertyType,
		currentOrderedPropertyType,
		firstPropertyType,
		firstOrderedPropertyType,
		generatorPropertyType,
		iconPropertyType,
		imagePropertyType,
		inReplyToPropertyType,
		instrumentPropertyType,
		lastPropertyType,
		lastOrderedPropertyType,
		locationPropertyType,
		itemsPropertyType,
		orderedItemsPropertyType,
		oneOfPropertyType,
		anyOfPropertyType,
		closedPropertyType,
		originPropertyType,
		nextPropertyType,
		nextOrderedPropertyType,
		objectPropertyType,
		prevPropertyType,
		prevOrderedPropertyType,
		previewPropertyType,
		resultPropertyType,
		repliesPropertyType,
		tagPropertyType,
		targetPropertyType,
		toPropertyType,
		urlPropertyType,
		accuracyPropertyType,
		altitudePropertyType,
		contentPropertyType,
		namePropertyType,
		durationPropertyType,
		heightPropertyType,
		hrefPropertyType,
		hrefLangPropertyType,
		partOfPropertyType,
		latitudePropertyType,
		longitudePropertyType,
		mediaTypePropertyType,
		endTimePropertyType,
		publishedPropertyType,
		startTimePropertyType,
		radiusPropertyType,
		relPropertyType,
		startIndexPropertyType,
		summaryPropertyType,
		totalItemsPropertyType,
		unitsPropertyType,
		updatedPropertyType,
		widthPropertyType,
		subjectPropertyType,
		relationshipPropertyType,
		describesPropertyType,
		formerTypePropertyType,
		deletedPropertyType,
		sourcePropertyType,
		inboxPropertyType,
		outboxPropertyType,
		followingPropertyType,
		followersPropertyType,
		likedPropertyType,
		likesPropertyType,
		streamsPropertyType,
		preferredUsernamePropertyType,
		endpointsPropertyType,
		proxyUrlPropertyType,
		oauthAuthorizationEndpointPropertyType,
		oauthTokenEndpointPropertyType,
		provideClientKeyPropertyType,
		signClientKeyPropertyType,
		sharedInboxPropertyType,
	}
)

// ValueType definitions
const (
	xsdBaseURI = "http://www.w3.org/2001/XMLSchema#"
	rdfBaseURI = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
)

var (
	xsdDateTimeValueType = &ValueType{
		Name:           "dateTime",
		URI:            xsdBaseURI + "dateTime",
		DefinitionType: "*time.Time",
		Zero:           "time.Time{}",
		DeserializeFn: &FunctionDef{
			Name:    "dateTimeDeserialize",
			Comment: "dateTimeDeserialize turns a string into a time.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"t", "*time.Time"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("var tmp time.Time\n")
				b.WriteString("if s, ok := v.(string); ok {\n")
				b.WriteString("tmp, err = time.Parse(time.RFC3339, s)\n")
				b.WriteString("if err != nil {\n")
				b.WriteString("tmp, err = time.Parse(\"2006-01-02T15:04Z07:00\", s)\n")
				b.WriteString("if err != nil {\n")
				b.WriteString("err = fmt.Errorf(\"%s cannot be interpreted as xsd:dateTime\", s)\n")
				b.WriteString("} else {\n")
				b.WriteString("t = &tmp\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("t = &tmp\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for xsd:dateTime\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "dateTimeSerialize",
			Comment: "dateTimeSerialize turns a time into a string",
			Args:    []*FunctionVarDef{{"t", "time.Time"}},
			Return:  []*FunctionVarDef{{"s", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = t.Format(time.RFC3339)\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdBooleanValueType = &ValueType{
		Name:           "boolean",
		URI:            xsdBaseURI + "boolean",
		DefinitionType: "*bool",
		Zero:           "false",
		DeserializeFn: &FunctionDef{
			Name:    "booleanDeserialize",
			Comment: "booleanDeserialize turns a interface{} into a bool.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"b", "*bool"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if bv, ok := v.(bool); ok {\n")
				b.WriteString("b = &bv\n")
				b.WriteString("} else if bv, ok := v.(float64); ok {\n")
				b.WriteString("if bv == 0 {\n")
				b.WriteString("bvb := false\n")
				b.WriteString("b = &bvb\n")
				b.WriteString("} else if bv == 1 {\n")
				b.WriteString("bvb := true\n")
				b.WriteString("b = &bvb\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%d cannot be interpreted as a bool float64 for xsd:boolean\", v)\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a bool for xsd:boolean\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "booleanSerialize",
			Comment: "booleanSerialize simply returns the bool value",
			Args:    []*FunctionVarDef{{"v", "bool"}},
			Return:  []*FunctionVarDef{{"r", "bool"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = v\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdAnyURIValueType = &ValueType{
		Name:           "anyURI",
		URI:            xsdBaseURI + "anyURI",
		DefinitionType: "*url.URL",
		Zero:           "\"\"",
		DeserializeFn: &FunctionDef{
			Name:    "anyURIDeserialize",
			Comment: "anyURIDeserialize turns a string into a URI.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"u", "*url.URL"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if s, ok := v.(string); ok {\n")
				b.WriteString("u, err = url.Parse(s)\n")
				b.WriteString("if err != nil {\n")
				b.WriteString("err = fmt.Errorf(\"%s cannot be interpreted as xsd:anyURI\", s)\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for xsd:anyURI\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "anyURISerialize",
			Comment: "anyURISerialize turns a URI into a string",
			Args:    []*FunctionVarDef{{"u", "*url.URL"}},
			Return:  []*FunctionVarDef{{"s", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = u.String()\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdFloatValueType = &ValueType{
		Name:           "float",
		URI:            xsdBaseURI + "float",
		DefinitionType: "*float64",
		Zero:           "0",
		DeserializeFn: &FunctionDef{
			Name:    "floatDeserialize",
			Comment: "floatDeserialize turns a interface{} into a float64.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"f", "*float64"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if fv, ok := v.(float64); ok {\n")
				b.WriteString("f = &fv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a float for xsd:float\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "floatSerialize",
			Comment: "floatSerialize simply returns the float value",
			Args:    []*FunctionVarDef{{"f", "float64"}},
			Return:  []*FunctionVarDef{{"r", "float64"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = f\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdStringValueType = &ValueType{
		Name:           "string",
		URI:            xsdBaseURI + "string",
		DefinitionType: "*string",
		Zero:           "0",
		DeserializeFn: &FunctionDef{
			Name:    "stringDeserialize",
			Comment: "stringDeserialize turns a interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for xsd:string\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "stringSerialize",
			Comment: "stringSerialize simply returns the string value",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	rdfLangStringValueType = &ValueType{
		Name:           "langString",
		URI:            rdfBaseURI + "langString",
		DefinitionType: "*string",
		Zero:           "0",
		DeserializeFn: &FunctionDef{
			Name:    "langStringDeserialize",
			Comment: "langStringDeserialize turns a RDF interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for rdf:langString\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "langStringSerialize",
			Comment: "langStringSerialize returns a formatted RDF value.",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdDurationValueType = &ValueType{
		Name:           "duration",
		URI:            xsdBaseURI + "duration",
		DefinitionType: "*time.Duration",
		Zero:           "0",
		DeserializeFn: &FunctionDef{
			Name:    "durationDeserialize",
			Comment: "durationDeserialize turns a interface{} into a time.Duration.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"d", "*time.Duration"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("isNeg := false\n")
				b.WriteString("if sv[0] == '-' {\n")
				b.WriteString("isNeg = true\n")
				b.WriteString("sv = sv[1:]\n")
				b.WriteString("}\n")
				b.WriteString("if sv[0] != 'P' {\n")
				b.WriteString("err = fmt.Errorf(\"'%s' malformed: missing 'P' for xsd:duration\", sv)\n")
				b.WriteString("return\n")
				b.WriteString("}\n")
				b.WriteString(`re := regexp.MustCompile("P(\\d*Y)?(\\d*M)?(\\d*D)?(T(\\d*H)?(\\d*M)?(\\d*S)?)?")`)
				b.WriteString("\n")
				b.WriteString("res := re.FindStringSubmatch(sv)\n")
				b.WriteString("var dur time.Duration\n")
				b.WriteString("nYear := res[1]\n")
				b.WriteString("if len(nYear) > 0 {\n")
				b.WriteString("nYear = nYear[:len(nYear)-1]\n")
				b.WriteString("vYear, err := strconv.ParseInt(nYear, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vYear) * time.Hour * 8760 // Hours per 365 days -- no way to count leap years here.\n")
				b.WriteString("}\n")
				b.WriteString("nMonth := res[2]\n")
				b.WriteString("if len(nMonth) > 0 {\n")
				b.WriteString("nMonth = nMonth[:len(nMonth)-1]\n")
				b.WriteString("vMonth, err := strconv.ParseInt(nMonth, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vMonth) * time.Hour * 720 // Hours per 30 days -- no way to tell if these months span 31 days, 28, or 29 each.\n")
				b.WriteString("}\n")
				b.WriteString("nDay := res[3]\n")
				b.WriteString("if len(nDay) > 0 {\n")
				b.WriteString("nDay = nDay[:len(nDay)-1]\n")
				b.WriteString("vDay, err := strconv.ParseInt(nDay, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vDay) * time.Hour * 24\n")
				b.WriteString("}\n")
				b.WriteString("nHour := res[5]\n")
				b.WriteString("if len(nHour) > 0 {\n")
				b.WriteString("nHour = nHour[:len(nHour)-1]\n")
				b.WriteString("vHour, err := strconv.ParseInt(nHour, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vHour) * time.Hour\n")
				b.WriteString("}\n")
				b.WriteString("nMinute := res[6]\n")
				b.WriteString("if len(nMinute) > 0 {\n")
				b.WriteString("nMinute = nMinute[:len(nMinute)-1]\n")
				b.WriteString("vMinute, err := strconv.ParseInt(nMinute, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vMinute) * time.Minute\n")
				b.WriteString("}\n")
				b.WriteString("nSecond := res[7]\n")
				b.WriteString("if len(nSecond) > 0 {\n")
				b.WriteString("nSecond = nSecond[:len(nSecond)-1]\n")
				b.WriteString("vSecond, err := strconv.ParseInt(nSecond, 10, 64)\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("dur += time.Duration(vSecond) * time.Second\n")
				b.WriteString("}\n")
				b.WriteString("if isNeg {\n")
				b.WriteString("dur *= -1\n")
				b.WriteString("}\n")
				b.WriteString("d = &dur\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for xsd:duration\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "durationSerialize",
			Comment: "durationSerialize returns the duration as a string.",
			Args:    []*FunctionVarDef{{"d", "time.Duration"}},
			Return:  []*FunctionVarDef{{"s", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = \"P\"\n")
				b.WriteString("if d < 0 {\n")
				b.WriteString("s = \"-P\"\n")
				b.WriteString("d = -1 * d\n")
				b.WriteString("}\n")
				b.WriteString("var tally time.Duration\n")
				b.WriteString("if years := d.Hours() / 8760.0; years >= 1 {\n")
				b.WriteString("nYears := int64(math.Floor(years))\n")
				b.WriteString("tally += time.Duration(nYears) * 8760 * time.Hour\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dY\", s, nYears)\n")
				b.WriteString("}\n")
				b.WriteString("if months := (d.Hours() - tally.Hours()) / 720.0; months >= 1 {\n")
				b.WriteString("nMonths := int64(math.Floor(months))\n")
				b.WriteString("tally += time.Duration(nMonths) * 720 * time.Hour\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dM\", s, nMonths)\n")
				b.WriteString("}\n")
				b.WriteString("if days := (d.Hours() - tally.Hours()) / 24.0; days >= 1 {\n")
				b.WriteString("nDays := int64(math.Floor(days))\n")
				b.WriteString("tally += time.Duration(nDays) * 24 * time.Hour\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dD\", s, nDays)\n")
				b.WriteString("}\n")
				b.WriteString("if tally < d {\n")
				b.WriteString("s = fmt.Sprintf(\"%sT\", s)\n")
				b.WriteString("if hours := d.Hours() - tally.Hours(); hours >= 1 {\n")
				b.WriteString("nHours := int64(math.Floor(hours))\n")
				b.WriteString("tally += time.Duration(nHours) * time.Hour\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dH\", s, nHours)\n")
				b.WriteString("}\n")
				b.WriteString("if minutes := d.Minutes() - tally.Minutes(); minutes >= 1 {\n")
				b.WriteString("nMinutes := int64(math.Floor(minutes))\n")
				b.WriteString("tally += time.Duration(nMinutes) * time.Minute\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dM\", s, nMinutes)\n")
				b.WriteString("}\n")
				b.WriteString("if seconds := d.Seconds() - tally.Seconds(); seconds >= 1 {\n")
				b.WriteString("nSeconds := int64(math.Floor(seconds))\n")
				b.WriteString("tally += time.Duration(nSeconds) * time.Second\n")
				b.WriteString("s = fmt.Sprintf(\"%s%dS\", s, nSeconds)\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	xsdNonNegativeIntegerValueType = &ValueType{
		Name:           "nonNegativeInteger",
		URI:            xsdBaseURI + "nonNegativeInteger",
		DefinitionType: "*int64",
		Zero:           "0",
		DeserializeFn: &FunctionDef{
			Name:    "nonNegativeIntegerDeserialize",
			Comment: "nonNegativeIntegerDeserialize turns a interface{} into a positive int64 value.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"i", "*int64"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if fv, ok := v.(float64); ok {\n")
				b.WriteString("iv := int64(fv)\n")
				b.WriteString("if iv >= 0 {\n")
				b.WriteString("i = &iv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%d is a negative integer for xsd:nonNegativeInteger\", iv)\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a float for xsd:nonNegativeInteger\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "nonNegativeIntegerSerialize",
			Comment: "nonNegativeIntegerSerialize simply returns the int64 value.",
			Args:    []*FunctionVarDef{{"i", "int64"}},
			Return:  []*FunctionVarDef{{"r", "int64"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = i\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	bcp47LangTag = &ValueType{
		Name:           "bcp47LanguageTag",
		DefinitionType: "*string",
		Zero:           "\"\"",
		DeserializeFn: &FunctionDef{
			Name:    "bcp47LanguageTagDeserialize",
			Comment: "bcp47LanguageTagDeserialize turns a interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for BCP 47 Language Tag\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "bcp47LanguageTagSerialize",
			Comment: "bcp47LanguageTagSerialize simply returns the string value",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	mimeMediaValueType = &ValueType{
		Name:           "mimeMediaTypeValue",
		DefinitionType: "*string",
		Zero:           "\"\"",
		DeserializeFn: &FunctionDef{
			Name:    "mimeMediaTypeValueDeserialize",
			Comment: "mimeMediaTypeValueDeserialize turns a interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for MIME media type value\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "mimeMediaTypeValueSerialize",
			Comment: "mimeMediaTypeValueSerialize simply returns the string value",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	linkRelationValueType = &ValueType{
		Name:           "linkRelation",
		DefinitionType: "*string",
		Zero:           "\"\"",
		DeserializeFn: &FunctionDef{
			Name:    "linkRelationDeserialize",
			Comment: "linkRelationDeserialize turns a interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for link relation\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "linkRelationSerialize",
			Comment: "linkRelationSerialize simply returns the string value",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
	unitsValueType = &ValueType{
		Name:           "unitsValue",
		DefinitionType: "*string",
		Zero:           "\"\"",
		DeserializeFn: &FunctionDef{
			Name:    "unitsValueDeserialize",
			Comment: "unitsValueDeserialize turns a interface{} into a string.",
			Args:    []*FunctionVarDef{{"v", "interface{}"}},
			Return:  []*FunctionVarDef{{"s", "*string"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if sv, ok := v.(string); ok {\n")
				b.WriteString("s = &sv\n")
				b.WriteString("} else {\n")
				b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for units value\", v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &FunctionDef{
			Name:    "unitsValueSerialize",
			Comment: "unitsValueSerialize simply returns the string value",
			Args:    []*FunctionVarDef{{"s", "string"}},
			Return:  []*FunctionVarDef{{"r", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = s\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}

	AllValueTypes = []*ValueType{
		xsdDateTimeValueType,
		xsdBooleanValueType,
		xsdAnyURIValueType,
		xsdFloatValueType,
		xsdStringValueType,
		rdfLangStringValueType,
		xsdDurationValueType,
		xsdNonNegativeIntegerValueType,
		bcp47LangTag,
		mimeMediaValueType,
		linkRelationValueType,
		unitsValueType,
	}
)

func init() {
	objectType.DisjointWith = []*Type{linkType}

	// Type properties
	objectType.Properties = []*PropertyType{
		altitudePropertyType, // Missing from spec!
		attachmentPropertyType,
		attributedToPropertyType,
		audiencePropertyType,
		contentPropertyType,
		contextPropertyType,
		namePropertyType,
		endTimePropertyType,
		generatorPropertyType,
		iconPropertyType,
		idPropertyType, // Missing from spec!
		imagePropertyType,
		inReplyToPropertyType,
		locationPropertyType,
		previewPropertyType,
		publishedPropertyType,
		repliesPropertyType,
		startTimePropertyType,
		summaryPropertyType,
		tagPropertyType,
		typePropertyType, // Missing from spec!
		updatedPropertyType,
		urlPropertyType,
		toPropertyType,
		btoPropertyType,
		ccPropertyType,
		bccPropertyType,
		mediaTypePropertyType,
		durationPropertyType,
		// ActivityPub specific properties
		sourcePropertyType,
		inboxPropertyType,
		outboxPropertyType,
		followingPropertyType,
		followersPropertyType,
		likedPropertyType,
		likesPropertyType,
		streamsPropertyType,
		preferredUsernamePropertyType,
		endpointsPropertyType,
		proxyUrlPropertyType,
		oauthAuthorizationEndpointPropertyType,
		oauthTokenEndpointPropertyType,
		provideClientKeyPropertyType,
		signClientKeyPropertyType,
		sharedInboxPropertyType,
	}
	linkType.Properties = []*PropertyType{
		attributedToPropertyType, // Missing from spec!
		hrefPropertyType,
		idPropertyType, // Missing from spec!
		relPropertyType,
		typePropertyType, // Missing from spec!
		mediaTypePropertyType,
		namePropertyType,
		summaryPropertyType, // Only in example 58, missing from spec!
		hrefLangPropertyType,
		heightPropertyType,
		widthPropertyType,
		previewPropertyType,
	}
	activityType.Properties = []*PropertyType{
		actorPropertyType,
		objectPropertyType,
		targetPropertyType,
		resultPropertyType,
		originPropertyType,
		instrumentPropertyType,
	}
	collectionType.Properties = []*PropertyType{
		totalItemsPropertyType,
		currentPropertyType,
		firstPropertyType,
		lastPropertyType,
		itemsPropertyType,
	}
	orderedCollectionType.Properties = []*PropertyType{
		orderedItemsPropertyType, // Missing from spec!
		currentOrderedPropertyType,
		firstOrderedPropertyType,
		lastOrderedPropertyType,
	}
	collectionPageType.Properties = []*PropertyType{
		partOfPropertyType,
		nextPropertyType,
		prevPropertyType,
	}
	orderedCollectionPageType.Properties = []*PropertyType{
		startIndexPropertyType,
		nextOrderedPropertyType,
		prevOrderedPropertyType,
	}

	// ExtendedType properties
	imageExtendedType.Properties = []*PropertyType{
		heightPropertyType, // Missing from spec!
		widthPropertyType,  // Missing from spec!
	}
	questionExtendedType.Properties = []*PropertyType{
		oneOfPropertyType,
		anyOfPropertyType,
		closedPropertyType,
	}
	relationshipExtendedType.Properties = []*PropertyType{
		subjectPropertyType,
		objectPropertyType,
		relationshipPropertyType,
	}
	placeExtendedType.Properties = []*PropertyType{
		accuracyPropertyType,
		altitudePropertyType,
		latitudePropertyType,
		longitudePropertyType,
		radiusPropertyType,
		unitsPropertyType,
	}
	profileExtendedType.Properties = []*PropertyType{
		describesPropertyType,
	}
	tombstoneExtendedType.Properties = []*PropertyType{
		formerTypePropertyType,
		deletedPropertyType,
	}
}

var IriValueType = &ValueType{
	Name:           "IRI",
	URI:            "IRI",
	DefinitionType: "*url.URL",
	Zero:           "nil",
	DeserializeFn: &FunctionDef{
		Name:    "IRIDeserialize",
		Comment: "IRIDeserialize turns a string into a URI.",
		Args:    []*FunctionVarDef{{"v", "interface{}"}},
		Return:  []*FunctionVarDef{{"u", "*url.URL"}, {"err", "error"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("if s, ok := v.(string); ok {\n")
			b.WriteString("u, err = url.Parse(s)\n")
			b.WriteString("if err != nil {\n")
			b.WriteString("err = fmt.Errorf(\"%s cannot be interpreted as IRI\", s)\n")
			b.WriteString("}\n")
			b.WriteString("} else {\n")
			b.WriteString("err = fmt.Errorf(\"%v cannot be interpreted as a string for IRI\", v)\n")
			b.WriteString("}\n")
			b.WriteString("return\n")
			return b.String()
		},
	},
	SerializeFn: &FunctionDef{
		Name:    "IRISerialize",
		Comment: "IRISerialize turns an IRI into a string",
		Args:    []*FunctionVarDef{{"u", "*url.URL"}},
		Return:  []*FunctionVarDef{{"s", "string"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("s = u.String()\n")
			b.WriteString("return\n")
			return b.String()
		},
	},
}

func init() {
	for _, p := range AllPropertyTypes {
		if !HasAnyURI(p.Range) {
			p.Range = append(p.Range, RangeReference{V: IriValueType})
		}
	}
}

func HasAnyURI(r []RangeReference) bool {
	for _, v := range r {
		if v.V != nil && v.V == xsdAnyURIValueType {
			return true
		}
	}
	return false
}

func IsIRIValueTypeString(v *ValueType) bool {
	return v.DefinitionType == "*url.URL"
}

func IsIRIValueType(v *ValueType) bool {
	return v == IriValueType
}

func IsAnyURIValueType(v *ValueType) bool {
	return v == xsdAnyURIValueType
}

func AnyURIValueTypeName() string {
	return xsdAnyURIValueType.Name
}

func IRIFuncs() []*FunctionDef {
	return []*FunctionDef{IriValueType.DeserializeFn, IriValueType.SerializeFn}
}

func IsActivity(t *Type) bool {
	var recur func(t *Type) bool
	recur = func(t *Type) bool {
		for _, e := range t.Extends {
			if e == activityType {
				return true
			} else if recur(e) {
				return true
			}
		}
		return false
	}
	return recur(t)
}
