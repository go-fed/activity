package streams

import (
	"github.com/go-fed/activity/streams/vocab"
	"net/url"
	"time"
)

func MustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

const example1 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Object",
  "id": "http://www.test.example/object/1",
  "name": "A Simple, non-specific object"
}`

func example1Type() vocab.ObjectInterface {
	example1Type := NewActivityStreamsObject()
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://www.test.example/object/1"))
	example1Type.SetId(id)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A Simple, non-specific object")
	example1Type.SetName(name)
	return example1Type
}

const example2 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link"
}`

func example2Type() vocab.LinkInterface {
	example2Type := NewActivityStreamsLink()
	hrefUrl := MustParseURL("http://example.org/abc")
	href := NewActivityStreamsHrefProperty()
	href.Set(hrefUrl)
	example2Type.SetHref(href)
	hrefLang := NewActivityStreamsHreflangProperty()
	hrefLang.Set("en")
	example2Type.SetHreflang(hrefLang)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/html")
	example2Type.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("An example link")
	example2Type.SetName(name)
	return example2Type
}

const example3 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Activity",
  "summary": "Sally did something to a note",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Note",
    "name": "A Note"
  }
}`

func example3Type() vocab.ActivityInterface {
	example3Type := NewActivityStreamsActivity()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	note := NewActivityStreamsNote()
	aNote := NewActivityStreamsNameProperty()
	aNote.AppendString("Sally")
	note.SetName(aNote)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally did something to a note")
	example3Type.SetSummary(summary)
	actor := NewActivityStreamsActorProperty()
	actor.AppendPerson(person)
	example3Type.SetActor(actor)
	object := NewActivityStreamsObjectProperty()
	object.AppendNote(note)
	example3Type.SetObject(object)
	return example3Type
}

const example4 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Travel",
  "summary": "Sally went to work",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "target": {
    "type": "Place",
    "name": "Work"
  }
}`

func example4Type() vocab.TravelInterface {
	example4Type := NewActivityStreamsTravel()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	place := NewActivityStreamsPlace()
	work := NewActivityStreamsNameProperty()
	work.AppendString("Work")
	place.SetName(work)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally went to work")
	example4Type.SetSummary(summary)
	actor := NewActivityStreamsActorProperty()
	actor.AppendPerson(person)
	example4Type.SetActor(actor)
	target := NewActivityStreamsTargetProperty()
	target.AppendPlace(place)
	example4Type.SetTarget(target)
	return example4Type
}

const example5 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's notes",
  "type": "Collection",
  "totalItems": 2,
  "items": [
    {
      "type": "Note",
      "name": "A Simple Note"
    },
    {
      "type": "Note",
      "name": "Another Simple Note"
    }
  ]
}`

func example5Type() vocab.CollectionInterface {
	example5Type := NewActivityStreamsCollection()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("A Simple Note")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Another Simple Note")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's notes")
	example5Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(2)
	example5Type.SetTotalItems(totalItems)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note1)
	items.AppendNote(note2)
	example5Type.SetItems(items)
	return example5Type
}

const example6 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's notes",
  "type": "OrderedCollection",
  "totalItems": 2,
  "orderedItems": [
    {
      "type": "Note",
      "name": "A Simple Note"
    },
    {
      "type": "Note",
      "name": "Another Simple Note"
    }
  ]
}`

func example6Type() vocab.OrderedCollectionInterface {
	example6Type := NewActivityStreamsOrderedCollection()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("A Simple Note")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Another Simple Note")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's notes")
	example6Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(2)
	example6Type.SetTotalItems(totalItems)
	orderedItems := NewActivityStreamsOrderedItemsProperty()
	orderedItems.AppendNote(note1)
	orderedItems.AppendNote(note2)
	example6Type.SetOrderedItems(orderedItems)
	return example6Type
}

const example7 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's notes",
  "type": "CollectionPage",
  "id": "http://example.org/foo?page=1",
  "partOf": "http://example.org/foo",
  "items": [
    {
      "type": "Note",
      "name": "A Simple Note"
    },
    {
      "type": "Note",
      "name": "Another Simple Note"
    }
  ]
}`

func example7Type() vocab.CollectionPageInterface {
	example7Type := NewActivityStreamsCollectionPage()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("A Simple Note")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Another Simple Note")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's notes")
	example7Type.SetSummary(summary)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://example.org/foo?page=1"))
	example7Type.SetId(id)
	partOf := NewActivityStreamsPartOfProperty()
	partOf.SetIRI(MustParseURL("http://example.org/foo"))
	example7Type.SetPartOf(partOf)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note1)
	items.AppendNote(note2)
	example7Type.SetItems(items)
	return example7Type
}

const example8 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's notes",
  "type": "OrderedCollectionPage",
  "id": "http://example.org/foo?page=1",
  "partOf": "http://example.org/foo",
  "orderedItems": [
    {
      "type": "Note",
      "name": "A Simple Note"
    },
    {
      "type": "Note",
      "name": "Another Simple Note"
    }
  ]
}`

func example8Type() vocab.OrderedCollectionPageInterface {
	example8Type := NewActivityStreamsOrderedCollectionPage()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("A Simple Note")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Another Simple Note")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's notes")
	example8Type.SetSummary(summary)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://example.org/foo?page=1"))
	example8Type.SetId(id)
	partOf := NewActivityStreamsPartOfProperty()
	partOf.SetIRI(MustParseURL("http://example.org/foo"))
	example8Type.SetPartOf(partOf)
	orderedItems := NewActivityStreamsOrderedItemsProperty()
	orderedItems.AppendNote(note1)
	orderedItems.AppendNote(note2)
	example8Type.SetOrderedItems(orderedItems)
	return example8Type
}

const example9 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally accepted an invitation to a party",
  "type": "Accept",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Invite",
    "actor": "http://john.example.org",
    "object": {
      "type": "Event",
      "name": "Going-Away Party for Jim"
    }
  }
}`

func example9Type() vocab.AcceptInterface {
	example9Type := NewActivityStreamsAccept()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	event := NewActivityStreamsEvent()
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	event.SetName(goingAway)
	invite := NewActivityStreamsInvite()
	actor := NewActivityStreamsActorProperty()
	actor.AppendIRI(MustParseURL("http://john.example.org"))
	invite.SetActor(actor)
	object := NewActivityStreamsObjectProperty()
	object.AppendEvent(event)
	invite.SetObject(object)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally accepted an invitation to a party")
	example9Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example9Type.SetActor(rootActor)
	inviteObject := NewActivityStreamsObjectProperty()
	inviteObject.AppendInvite(invite)
	example9Type.SetObject(inviteObject)
	return example9Type
}

const example10 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally accepted Joe into the club",
  "type": "Accept",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Person",
    "name": "Joe"
  },
  "target": {
    "type": "Group",
    "name": "The Club"
  }
}`

func example10Type() vocab.AcceptInterface {
	example10Type := NewActivityStreamsAccept()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	object := NewActivityStreamsPerson()
	joe := NewActivityStreamsNameProperty()
	joe.AppendString("Joe")
	object.SetName(joe)
	target := NewActivityStreamsGroup()
	club := NewActivityStreamsNameProperty()
	club.AppendString("The Club")
	target.SetName(club)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally accepted Joe into the club")
	example10Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example10Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example10Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendGroup(target)
	example10Type.SetTarget(tobj)
	return example10Type
}

const example11 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally tentatively accepted an invitation to a party",
  "type": "TentativeAccept",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Invite",
    "actor": "http://john.example.org",
    "object": {
      "type": "Event",
      "name": "Going-Away Party for Jim"
    }
  }
}`

func example11Type() vocab.TentativeAcceptInterface {
	example11Type := NewActivityStreamsTentativeAccept()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	event := NewActivityStreamsEvent()
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	event.SetName(goingAway)
	invite := NewActivityStreamsInvite()
	inviteActor := NewActivityStreamsActorProperty()
	inviteActor.AppendIRI(MustParseURL("http://john.example.org"))
	invite.SetActor(inviteActor)
	objInvite := NewActivityStreamsObjectProperty()
	objInvite.AppendEvent(event)
	invite.SetObject(objInvite)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally tentatively accepted an invitation to a party")
	example11Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example11Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendInvite(invite)
	example11Type.SetObject(obj)
	return example11Type
}

const example12 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally added an object",
  "type": "Add",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/abc"
}`

func example12Type() vocab.AddInterface {
	example12Type := NewActivityStreamsAdd()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	link := MustParseURL("http://example.org/abc")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally added an object")
	example12Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example12Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example12Type.SetObject(obj)
	return example12Type
}

const example13 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally added a picture of her cat to her cat picture collection",
  "type": "Add",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Image",
    "name": "A picture of my cat",
    "url": "http://example.org/img/cat.png"
  },
  "origin": {
    "type": "Collection",
    "name": "Camera Roll"
  },
  "target": {
    "type": "Collection",
    "name": "My Cat Pictures"
  }
}`

func example13Type() vocab.AddInterface {
	example13Type := NewActivityStreamsAdd()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	link := MustParseURL("http://example.org/img/cat.png")
	object := NewActivityStreamsImage()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("A picture of my cat")
	object.SetName(objectName)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(link)
	object.SetUrl(urlProp)
	origin := NewActivityStreamsCollection()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("Camera Roll")
	origin.SetName(originName)
	target := NewActivityStreamsCollection()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("My Cat Pictures")
	target.SetName(targetName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally added a picture of her cat to her cat picture collection")
	example13Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example13Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendImage(object)
	example13Type.SetObject(obj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendCollection(origin)
	example13Type.SetOrigin(originProp)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendCollection(target)
	example13Type.SetTarget(tobj)
	return example13Type
}

const example14 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally arrived at work",
  "type": "Arrive",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "location": {
    "type": "Place",
    "name": "Work"
  },
  "origin": {
    "type": "Place",
    "name": "Home"
  }
}`

func example14Type() vocab.ArriveInterface {
	example14Type := NewActivityStreamsArrive()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	location := NewActivityStreamsPlace()
	locationName := NewActivityStreamsNameProperty()
	locationName.AppendString("Work")
	location.SetName(locationName)
	origin := NewActivityStreamsPlace()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("Home")
	origin.SetName(originName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally arrived at work")
	example14Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(person)
	example14Type.SetActor(rootActor)
	loc := NewActivityStreamsLocationProperty()
	loc.AppendPlace(location)
	example14Type.SetLocation(loc)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendPlace(origin)
	example14Type.SetOrigin(originProp)
	return example14Type
}

const example15 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally created a note",
  "type": "Create",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Note",
    "name": "A Simple Note",
    "content": "This is a simple note"
  }
}`

func example15Type() vocab.CreateInterface {
	example15Type := NewActivityStreamsCreate()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsNote()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("A Simple Note")
	object.SetName(objectName)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is a simple note")
	object.SetContent(content)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally created a note")
	example15Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example15Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendNote(object)
	example15Type.SetObject(obj)
	return example15Type
}

const example16 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally deleted a note",
  "type": "Delete",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/notes/1",
  "origin": {
    "type": "Collection",
    "name": "Sally's Notes"
  }
}`

func example16Type() vocab.DeleteInterface {
	example16Type := NewActivityStreamsDelete()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	origin := NewActivityStreamsCollection()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("Sally's Notes")
	origin.SetName(originName)
	link := MustParseURL("http://example.org/notes/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally deleted a note")
	example16Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example16Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example16Type.SetObject(obj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendCollection(origin)
	example16Type.SetOrigin(originProp)
	return example16Type
}

const example17 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally followed John",
  "type": "Follow",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Person",
    "name": "John"
  }
}`

func example17Type() vocab.FollowInterface {
	example17Type := NewActivityStreamsFollow()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsPerson()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("John")
	object.SetName(objectName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally followed John")
	example17Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example17Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example17Type.SetObject(obj)
	return example17Type
}

const example18 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally ignored a note",
  "type": "Ignore",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/notes/1"
}`

func example18Type() vocab.IgnoreInterface {
	example18Type := NewActivityStreamsIgnore()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/notes/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally ignored a note")
	example18Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example18Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example18Type.SetObject(obj)
	return example18Type
}

const example19 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally joined a group",
  "type": "Join",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Group",
    "name": "A Simple Group"
  }
}`

func example19Type() vocab.JoinInterface {
	example19Type := NewActivityStreamsJoin()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsGroup()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("A Simple Group")
	object.SetName(objectName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally joined a group")
	example19Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example19Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendGroup(object)
	example19Type.SetObject(obj)
	return example19Type
}

const example20 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally left work",
  "type": "Leave",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Place",
    "name": "Work"
  }
}`

func example20Type() vocab.LeaveInterface {
	example20Type := NewActivityStreamsLeave()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsPlace()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("Work")
	object.SetName(objectName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally left work")
	example20Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example20Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPlace(object)
	example20Type.SetObject(obj)
	return example20Type
}

const example21 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally left a group",
  "type": "Leave",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Group",
    "name": "A Simple Group"
  }
}`

func example21Type() vocab.LeaveInterface {
	example21Type := NewActivityStreamsLeave()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsGroup()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("A Simple Group")
	object.SetName(objectName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally left a group")
	example21Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example21Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendGroup(object)
	example21Type.SetObject(obj)
	return example21Type
}

const example22 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally liked a note",
  "type": "Like",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/notes/1"
}`

func example22Type() vocab.LikeInterface {
	example22Type := NewActivityStreamsLike()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/notes/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally liked a note")
	example22Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example22Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example22Type.SetObject(obj)
	return example22Type
}

const example23 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered 50% off to Lewis",
  "type": "Offer",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "http://www.types.example/ProductOffer",
    "name": "50% Off!"
  },
  "target": {
    "type": "Person",
    "name": "Lewis"
  }
}`

var example23Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["object"] = map[string]interface{}{
		"type": "http://www.types.example/ProductOffer",
		"name": "50% Off!",
	}
	return m
}

func example23Type() vocab.OfferInterface {
	example23Type := NewActivityStreamsOffer()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	target := NewActivityStreamsPerson()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("Lewis")
	target.SetName(targetName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered 50% off to Lewis")
	example23Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example23Type.SetActor(rootActor)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendPerson(target)
	example23Type.SetTarget(tobj)
	return example23Type
}

const example24 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally invited John and Lisa to a party",
  "type": "Invite",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Event",
    "name": "A Party"
  },
  "target": [
    {
      "type": "Person",
      "name": "John"
    },
    {
      "type": "Person",
      "name": "Lisa"
    }
  ]
}`

func example24Type() vocab.InviteInterface {
	example24Type := NewActivityStreamsInvite()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsEvent()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("A Party")
	object.SetName(objectName)
	target1 := NewActivityStreamsPerson()
	target1Name := NewActivityStreamsNameProperty()
	target1Name.AppendString("John")
	target1.SetName(target1Name)
	target2 := NewActivityStreamsPerson()
	target2Name := NewActivityStreamsNameProperty()
	target2Name.AppendString("Lisa")
	target2.SetName(target2Name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally invited John and Lisa to a party")
	example24Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example24Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendEvent(object)
	example24Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendPerson(target1)
	tobj.AppendPerson(target2)
	example24Type.SetTarget(tobj)
	return example24Type
}

const example25 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally rejected an invitation to a party",
  "type": "Reject",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Invite",
    "actor": "http://john.example.org",
    "object": {
      "type": "Event",
      "name": "Going-Away Party for Jim"
    }
  }
}`

func example25Type() vocab.RejectInterface {
	example25Type := NewActivityStreamsReject()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	inviteObject := NewActivityStreamsEvent()
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	inviteObject.SetName(goingAway)
	object := NewActivityStreamsInvite()
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(MustParseURL("http://john.example.org"))
	object.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendEvent(inviteObject)
	object.SetObject(obj)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally rejected an invitation to a party")
	example25Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example25Type.SetActor(rootActor)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendInvite(object)
	example25Type.SetObject(objRoot)
	return example25Type
}

const example26 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally tentatively rejected an invitation to a party",
  "type": "TentativeReject",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Invite",
    "actor": "http://john.example.org",
    "object": {
      "type": "Event",
      "name": "Going-Away Party for Jim"
    }
  }
}`

func example26Type() vocab.TentativeRejectInterface {
	example26Type := NewActivityStreamsTentativeReject()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	inviteObject := NewActivityStreamsEvent()
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	inviteObject.SetName(goingAway)
	object := NewActivityStreamsInvite()
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(MustParseURL("http://john.example.org"))
	object.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendEvent(inviteObject)
	object.SetObject(obj)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally tentatively rejected an invitation to a party")
	example26Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example26Type.SetActor(rootActor)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendInvite(object)
	example26Type.SetObject(objRoot)
	return example26Type
}

const example27 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally removed a note from her notes folder",
  "type": "Remove",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/notes/1",
  "target": {
    "type": "Collection",
    "name": "Notes Folder"
  }
}`

func example27Type() vocab.RemoveInterface {
	example27Type := NewActivityStreamsRemove()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/notes/1")
	target := NewActivityStreamsCollection()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("Notes Folder")
	target.SetName(targetName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally removed a note from her notes folder")
	example27Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example27Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example27Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendCollection(target)
	example27Type.SetTarget(tobj)
	return example27Type
}

const example28 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "The moderator removed Sally from a group",
  "type": "Remove",
  "actor": {
    "type": "http://example.org/Role",
    "name": "The Moderator"
  },
  "object": {
    "type": "Person",
    "name": "Sally"
  },
  "origin": {
    "type": "Group",
    "name": "A Simple Group"
  }
}`

var example28Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["actor"] = map[string]interface{}{
		"type": "http://example.org/Role",
		"name": "The Moderator",
	}
	return m
}

func example28Type() vocab.RemoveInterface {
	example28Type := NewActivityStreamsRemove()
	object := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	object.SetName(sally)
	origin := NewActivityStreamsGroup()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("A Simple Group")
	origin.SetName(originName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("The moderator removed Sally from a group")
	example28Type.SetSummary(summary)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example28Type.SetObject(obj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendGroup(origin)
	example28Type.SetOrigin(originProp)
	return example28Type
}

const example29 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally retracted her offer to John",
  "type": "Undo",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Offer",
    "actor": "http://sally.example.org",
    "object": "http://example.org/posts/1",
    "target": "http://john.example.org"
  }
}`

func example29Type() vocab.UndoInterface {
	example29Type := NewActivityStreamsUndo()
	link := MustParseURL("http://sally.example.org")
	objectLink := MustParseURL("http://example.org/posts/1")
	targetLink := MustParseURL("http://john.example.org")
	object := NewActivityStreamsOffer()
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(link)
	object.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(objectLink)
	object.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(targetLink)
	object.SetTarget(target)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally retracted her offer to John")
	example29Type.SetSummary(summary)
	actor := NewActivityStreamsActorProperty()
	actor.AppendIRI(link)
	example29Type.SetActor(actor)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendOffer(object)
	example29Type.SetObject(objRoot)
	return example29Type
}

const example30 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally updated her note",
  "type": "Update",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/notes/1"
}`

func example30Type() vocab.UpdateInterface {
	example30Type := NewActivityStreamsUpdate()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/notes/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally updated her note")
	example30Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example30Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example30Type.SetObject(obj)
	return example30Type
}

const example31 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally read an article",
  "type": "View",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Article",
    "name": "What You Should Know About Activity Streams"
  }
}`

func example31Type() vocab.ViewInterface {
	example31Type := NewActivityStreamsView()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	object := NewActivityStreamsArticle()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("What You Should Know About Activity Streams")
	object.SetName(objectName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally read an article")
	example31Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example31Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendArticle(object)
	example31Type.SetObject(obj)
	return example31Type
}

const example32 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally listened to a piece of music",
  "type": "Listen",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/music.mp3"
}`

func example32Type() vocab.ListenInterface {
	example32Type := NewActivityStreamsListen()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/music.mp3")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally listened to a piece of music")
	example32Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example32Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example32Type.SetObject(obj)
	return example32Type
}

const example33 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally read a blog post",
  "type": "Read",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/posts/1"
}`

func example33Type() vocab.ReadInterface {
	example33Type := NewActivityStreamsRead()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/posts/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally read a blog post")
	example33Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example33Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example33Type.SetObject(obj)
	return example33Type
}

const example34 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally moved a post from List A to List B",
  "type": "Move",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/posts/1",
  "target": {
    "type": "Collection",
    "name": "List B"
  },
  "origin": {
    "type": "Collection",
    "name": "List A"
  }
}`

func example34Type() vocab.MoveInterface {
	example34Type := NewActivityStreamsMove()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	link := MustParseURL("http://example.org/posts/1")
	target := NewActivityStreamsCollection()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("List B")
	target.SetName(targetName)
	origin := NewActivityStreamsCollection()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("List A")
	origin.SetName(originName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally moved a post from List A to List B")
	example34Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example34Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(link)
	example34Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendCollection(target)
	example34Type.SetTarget(tobj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendCollection(origin)
	example34Type.SetOrigin(originProp)
	return example34Type
}

const example35 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally went home from work",
  "type": "Travel",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "target": {
    "type": "Place",
    "name": "Home"
  },
  "origin": {
    "type": "Place",
    "name": "Work"
  }
}`

func example35Type() vocab.TravelInterface {
	example35Type := NewActivityStreamsTravel()
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	target := NewActivityStreamsPlace()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("Home")
	target.SetName(targetName)
	origin := NewActivityStreamsPlace()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("Work")
	origin.SetName(originName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally went home from work")
	example35Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example35Type.SetActor(rootActor)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendPlace(target)
	example35Type.SetTarget(tobj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendPlace(origin)
	example35Type.SetOrigin(originProp)
	return example35Type
}

const example36 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally announced that she had arrived at work",
  "type": "Announce",
  "actor": {
    "type": "Person",
    "id": "http://sally.example.org",
    "name": "Sally"
  },
  "object": {
    "type": "Arrive",
    "actor": "http://sally.example.org",
    "location": {
      "type": "Place",
      "name": "Work"
    }
  }
}`

func example36Type() vocab.AnnounceInterface {
	example36Type := NewActivityStreamsAnnounce()
	link := MustParseURL("http://sally.example.org")
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	id := NewActivityStreamsIdProperty()
	id.Set(link)
	actor.SetId(id)
	loc := NewActivityStreamsPlace()
	locName := NewActivityStreamsNameProperty()
	locName.AppendString("Work")
	loc.SetName(locName)
	object := NewActivityStreamsArrive()
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(link)
	object.SetActor(objectActor)
	location := NewActivityStreamsLocationProperty()
	location.AppendPlace(loc)
	object.SetLocation(location)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally announced that she had arrived at work")
	example36Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example36Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendArrive(object)
	example36Type.SetObject(obj)
	return example36Type
}

const example37 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally blocked Joe",
  "type": "Block",
  "actor": "http://sally.example.org",
  "object": "http://joe.example.org"
}`

func example37Type() vocab.BlockInterface {
	example37Type := NewActivityStreamsBlock()
	link := MustParseURL("http://sally.example.org")
	objLink := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally blocked Joe")
	example37Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(link)
	example37Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(objLink)
	example37Type.SetObject(obj)
	return example37Type
}

const example38 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally flagged an inappropriate note",
  "type": "Flag",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Note",
    "content": "An inappropriate note"
  }
}`

func example38Type() vocab.FlagInterface {
	example38Type := NewActivityStreamsFlag()
	link := MustParseURL("http://sally.example.org")
	object := NewActivityStreamsNote()
	content := NewActivityStreamsContentProperty()
	content.AppendString("An inappropriate note")
	object.SetContent(content)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally flagged an inappropriate note")
	example38Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(link)
	example38Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendNote(object)
	example38Type.SetObject(obj)
	return example38Type
}

const example39 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally disliked a post",
  "type": "Dislike",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`

func example39Type() vocab.DislikeInterface {
	example39Type := NewActivityStreamsDislike()
	link := MustParseURL("http://sally.example.org")
	objLink := MustParseURL("http://example.org/posts/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally disliked a post")
	example39Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(link)
	example39Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(objLink)
	example39Type.SetObject(obj)
	return example39Type
}

const example40 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "oneOf": [
    {
      "type": "Note",
      "name": "Option A"
    },
    {
      "type": "Note",
      "name": "Option B"
    }
  ]
}`

func example40Type() vocab.QuestionInterface {
	example40Type := NewActivityStreamsQuestion()
	note1 := NewActivityStreamsNote()
	note1Name := NewActivityStreamsNameProperty()
	note1Name.AppendString("Option A")
	note1.SetName(note1Name)
	note2 := NewActivityStreamsNote()
	note2Name := NewActivityStreamsNameProperty()
	note2Name.AppendString("Option B")
	note2.SetName(note2Name)
	name := NewActivityStreamsNameProperty()
	name.AppendString("What is the answer?")
	example40Type.SetName(name)
	oneOf1 := NewActivityStreamsOneOfProperty()
	oneOf1.AppendNote(note1)
	example40Type.SetOneOf(oneOf1)
	oneOf2 := NewActivityStreamsOneOfProperty()
	oneOf2.AppendNote(note2)
	example40Type.SetOneOf(oneOf2)
	return example40Type
}

const example41 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "closed": "2016-05-10T00:00:00Z"
}`

func example41Type() vocab.QuestionInterface {
	example41Type := NewActivityStreamsQuestion()
	t, err := time.Parse(time.RFC3339, "2016-05-10T00:00:00Z")
	if err != nil {
		panic(err)
	}
	name := NewActivityStreamsNameProperty()
	name.AppendString("What is the answer?")
	example41Type.SetName(name)
	closed := NewActivityStreamsClosedProperty()
	closed.AppendDateTime(t)
	example41Type.SetClosed(closed)
	return example41Type
}

const example42 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Application",
  "name": "Exampletron 3000"
}`

func example42Type() vocab.ApplicationInterface {
	example42Type := NewActivityStreamsApplication()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Exampletron 3000")
	example42Type.SetName(name)
	return example42Type
}

const example43 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Group",
  "name": "Big Beards of Austin"
}`

func example43Type() vocab.GroupInterface {
	example43Type := NewActivityStreamsGroup()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Big Beards of Austin")
	example43Type.SetName(name)
	return example43Type
}

const example44 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Organization",
  "name": "Example Co."
}`

func example44Type() vocab.OrganizationInterface {
	example44Type := NewActivityStreamsOrganization()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Example Co.")
	example44Type.SetName(name)
	return example44Type
}

const example45 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally Smith"
}`

func example45Type() vocab.PersonInterface {
	example45Type := NewActivityStreamsPerson()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Sally Smith")
	example45Type.SetName(name)
	return example45Type
}

const example46 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Service",
  "name": "Acme Web Service"
}`

func example46Type() vocab.ServiceInterface {
	example46Type := NewActivityStreamsService()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Acme Web Service")
	example46Type.SetName(name)
	return example46Type
}

const example47 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally is an acquaintance of John",
  "type": "Relationship",
  "subject": {
    "type": "Person",
    "name": "Sally"
  },
  "relationship": "http://purl.org/vocab/relationship/acquaintanceOf",
  "object": {
    "type": "Person",
    "name": "John"
  }
}`

func example47Type() vocab.RelationshipInterface {
	example47Type := NewActivityStreamsRelationship()
	subject := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	subject.SetName(sally)
	object := NewActivityStreamsPerson()
	objectName := NewActivityStreamsNameProperty()
	objectName.AppendString("John")
	object.SetName(objectName)
	rel := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally is an acquaintance of John")
	example47Type.SetSummary(summary)
	subj := NewActivityStreamsSubjectProperty()
	subj.SetPerson(subject)
	example47Type.SetSubject(subj)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example47Type.SetObject(obj)
	relationship := NewActivityStreamsRelationshipProperty()
	relationship.AppendIRI(rel)
	example47Type.SetRelationship(relationship)
	return example47Type
}

const example48 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Article",
  "name": "What a Crazy Day I Had",
  "content": "<div>... you will never believe ...</div>",
  "attributedTo": "http://sally.example.org"
}`

func example48Type() vocab.ArticleInterface {
	example48Type := NewActivityStreamsArticle()
	att := MustParseURL("http://sally.example.org")
	name := NewActivityStreamsNameProperty()
	name.AppendString("What a Crazy Day I Had")
	example48Type.SetName(name)
	attr := NewActivityStreamsAttributedToProperty()
	attr.AppendIRI(att)
	example48Type.SetAttributedTo(attr)
	content := NewActivityStreamsContentProperty()
	content.AppendString("<div>... you will never believe ...</div>")
	example48Type.SetContent(content)
	return example48Type
}

const example49 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": "http://example.org/4q-sales-forecast.pdf"
}`

func example49Type() vocab.DocumentInterface {
	example49Type := NewActivityStreamsDocument()
	l := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	name := NewActivityStreamsNameProperty()
	name.AppendString("4Q Sales Forecast")
	example49Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(l)
	example49Type.SetUrl(urlProp)
	return example49Type
}

const example50 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Audio",
  "name": "Interview With A Famous Technologist",
  "url": {
    "type": "Link",
    "href": "http://example.org/podcast.mp3",
    "mediaType": "audio/mp3"
  }
}`

func example50Type() vocab.AudioInterface {
	example50Type := NewActivityStreamsAudio()
	l := MustParseURL("http://example.org/podcast.mp3")
	link := NewActivityStreamsLink()
	href := NewActivityStreamsHrefProperty()
	href.Set(l)
	link.SetHref(href)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("audio/mp3")
	link.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Interview With A Famous Technologist")
	example50Type.SetName(name)
	urlProperty := NewActivityStreamsUrlProperty()
	urlProperty.AppendLink(link)
	example50Type.SetUrl(urlProperty)
	return example50Type
}

const example51 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "Cat Jumping on Wagon",
  "url": [
    {
      "type": "Link",
      "href": "http://example.org/image.jpeg",
      "mediaType": "image/jpeg"
    },
    {
      "type": "Link",
      "href": "http://example.org/image.png",
      "mediaType": "image/png"
    }
  ]
}`

func example51Type() vocab.ImageInterface {
	example51Type := NewActivityStreamsImage()
	l1 := MustParseURL("http://example.org/image.jpeg")
	l2 := MustParseURL("http://example.org/image.png")
	link1 := NewActivityStreamsLink()
	href1 := NewActivityStreamsHrefProperty()
	href1.Set(l1)
	link1.SetHref(href1)
	mediaType1 := NewActivityStreamsMediaTypeProperty()
	mediaType1.Set("image/jpeg")
	link1.SetMediaType(mediaType1)
	link2 := NewActivityStreamsLink()
	href2 := NewActivityStreamsHrefProperty()
	href2.Set(l2)
	link2.SetHref(href2)
	mediaType2 := NewActivityStreamsMediaTypeProperty()
	mediaType2.Set("image/png")
	link2.SetMediaType(mediaType2)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Cat Jumping on Wagon")
	example51Type.SetName(name)
	urlProperty1 := NewActivityStreamsUrlProperty()
	urlProperty1.AppendLink(link1)
	example51Type.SetUrl(urlProperty1)
	urlProperty2 := NewActivityStreamsUrlProperty()
	urlProperty2.AppendLink(link2)
	example51Type.SetUrl(urlProperty2)
	return example51Type
}

const example52 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Puppy Plays With Ball",
  "url": "http://example.org/video.mkv",
  "duration": "PT2H"
}`

func example52Type() vocab.VideoInterface {
	example52Type := NewActivityStreamsVideo()
	l := MustParseURL("http://example.org/video.mkv")
	name := NewActivityStreamsNameProperty()
	name.AppendString("Puppy Plays With Ball")
	example52Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(l)
	example52Type.SetUrl(urlProp)
	dur := NewActivityStreamsDurationProperty()
	dur.Set(time.Hour * 2)
	example52Type.SetDuration(dur)
	return example52Type
}

const example53 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A Word of Warning",
  "content": "Looks like it is going to rain today. Bring an umbrella!"
}`

func example53Type() vocab.NoteInterface {
	example53Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendString("A Word of Warning")
	example53Type.SetName(name)
	content := NewActivityStreamsContentProperty()
	content.AppendString("Looks like it is going to rain today. Bring an umbrella!")
	example53Type.SetContent(content)
	return example53Type
}

const example54 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Page",
  "name": "Omaha Weather Report",
  "url": "http://example.org/weather-in-omaha.html"
}`

func example54Type() vocab.PageInterface {
	example54Type := NewActivityStreamsPage()
	l := MustParseURL("http://example.org/weather-in-omaha.html")
	name := NewActivityStreamsNameProperty()
	name.AppendString("Omaha Weather Report")
	example54Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(l)
	example54Type.SetUrl(urlProp)
	return example54Type
}

const example55 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

func example55Type() vocab.EventInterface {
	example55Type := NewActivityStreamsEvent()
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	example55Type.SetName(goingAway)
	startTime := NewActivityStreamsStartTimeProperty()
	startTime.Set(t1)
	example55Type.SetStartTime(startTime)
	endTime := NewActivityStreamsEndTimeProperty()
	endTime.Set(t2)
	example55Type.SetEndTime(endTime)
	return example55Type
}

const example56 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Work"
}`

func example56Type() vocab.PlaceInterface {
	example56Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Work")
	example56Type.SetName(name)
	return example56Type
}

const example57 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`

func example57Type() vocab.PlaceInterface {
	example57Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example57Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example57Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example57Type.SetLongitude(lon)
	rad := NewActivityStreamsRadiusProperty()
	rad.Set(15)
	example57Type.SetRadius(rad)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example57Type.SetUnits(units)
	return example57Type
}

const example58 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Mention",
  "summary": "Mention of Joe by Carrie in her note",
  "href": "http://example.org/joe",
  "name": "Joe"
}`

func example58Type() vocab.MentionInterface {
	example58Type := NewActivityStreamsMention()
	l := MustParseURL("http://example.org/joe")
	href := NewActivityStreamsHrefProperty()
	href.Set(l)
	example58Type.SetHref(href)
	joe := NewActivityStreamsNameProperty()
	joe.AppendString("Joe")
	example58Type.SetName(joe)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Mention of Joe by Carrie in her note")
	example58Type.SetSummary(summary)
	return example58Type
}

const example59 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Profile",
  "summary": "Sally's Profile",
  "describes": {
    "type": "Person",
    "name": "Sally Smith"
  }
}`

func example59Type() vocab.ProfileInterface {
	example59Type := NewActivityStreamsProfile()
	person := NewActivityStreamsPerson()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Sally Smith")
	person.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's Profile")
	example59Type.SetSummary(summary)
	describes := NewActivityStreamsDescribesProperty()
	describes.SetPerson(person)
	example59Type.SetDescribes(describes)
	return example59Type
}

// Note that the @context is missing from the spec!
const example60 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "OrderedCollection",
  "totalItems": 3,
  "name": "Vacation photos 2016",
  "orderedItems": [
    {
      "type": "Image",
      "id": "http://image.example/1"
    },
    {
      "type": "Tombstone",
      "formerType": "Image",
      "id": "http://image.example/2",
      "deleted": "2016-03-17T00:00:00Z"
    },
    {
      "type": "Image",
      "id": "http://image.example/3"
    }
  ]
}`

func example60Type() vocab.OrderedCollectionInterface {
	example60Type := NewActivityStreamsOrderedCollection()
	t, err := time.Parse(time.RFC3339, "2016-03-17T00:00:00Z")
	if err != nil {
		panic(err)
	}
	image1 := NewActivityStreamsImage()
	imgId1 := NewActivityStreamsIdProperty()
	imgId1.Set(MustParseURL("http://image.example/1"))
	image1.SetId(imgId1)
	tombstone := NewActivityStreamsTombstone()
	ft := NewActivityStreamsFormerTypeProperty()
	ft.AppendString("Image")
	tombstone.SetFormerType(ft)
	tombId := NewActivityStreamsIdProperty()
	tombId.Set(MustParseURL("http://image.example/2"))
	tombstone.SetId(tombId)
	deleted := NewActivityStreamsDeletedProperty()
	deleted.Set(t)
	tombstone.SetDeleted(deleted)
	image2 := NewActivityStreamsImage()
	imgId2 := NewActivityStreamsIdProperty()
	imgId2.Set(MustParseURL("http://image.example/3"))
	image2.SetId(imgId2)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example60Type.SetTotalItems(totalItems)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Vacation photos 2016")
	example60Type.SetName(name)
	orderedItems := NewActivityStreamsOrderedItemsProperty()
	orderedItems.AppendImage(image1)
	orderedItems.AppendTombstone(tombstone)
	orderedItems.AppendImage(image2)
	example60Type.SetOrderedItems(orderedItems)
	return example60Type
}

const example61 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Foo",
  "id": "http://example.org/foo"
}`

func example61Type() interface{} {
	return nil
}

var example61Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["id"] = "http://example.org/foo"
	m["name"] = "Foo"
	return m
}

const example62 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A foo",
  "type": "http://example.org/Foo"
}`

func example62Type() interface{} {
	return nil
}

var example62Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["type"] = "http://example.org/Foo"
	m["summary"] = "A foo"
	return m
}

const example63 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the Foo object",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/foo"
}`

func example63Type() vocab.OfferInterface {
	example63Type := NewActivityStreamsOffer()
	l := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/foo")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered the Foo object")
	example63Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(l)
	example63Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example63Type.SetObject(obj)
	return example63Type
}

const example64 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the Foo object",
  "type": "Offer",
  "actor": {
    "type": "Person",
    "id": "http://sally.example.org",
    "summary": "Sally"
  },
  "object": "http://example.org/foo"
}`

func example64Type() vocab.OfferInterface {
	example64Type := NewActivityStreamsOffer()
	actor := NewActivityStreamsPerson()
	actorId := NewActivityStreamsIdProperty()
	actorId.Set(MustParseURL("http://sally.example.org"))
	actor.SetId(actorId)
	summaryActor := NewActivityStreamsSummaryProperty()
	summaryActor.AppendString("Sally")
	actor.SetSummary(summaryActor)
	o := MustParseURL("http://example.org/foo")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered the Foo object")
	example64Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example64Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example64Type.SetObject(obj)
	return example64Type
}

const example65 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally and Joe offered the Foo object",
  "type": "Offer",
  "actor": [
    "http://joe.example.org",
    {
      "type": "Person",
      "id": "http://sally.example.org",
      "name": "Sally"
    }
  ],
  "object": "http://example.org/foo"
}`

func example65Type() vocab.OfferInterface {
	example65Type := NewActivityStreamsOffer()
	actor := NewActivityStreamsPerson()
	actorId := NewActivityStreamsIdProperty()
	actorId.Set(MustParseURL("http://sally.example.org"))
	actor.SetId(actorId)
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	o := MustParseURL("http://example.org/foo")
	l := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally and Joe offered the Foo object")
	example65Type.SetSummary(summary)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example65Type.SetObject(obj)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(l)
	example65Type.SetActor(objectActor)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example65Type.SetActor(rootActor)
	return example65Type
}

// NOTE: Changed to not be an array value for "attachment" to keep in line with other examples in spec!
const example66 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "Have you seen my cat?",
  "attachment": {
    "type": "Image",
    "content": "This is what he looks like.",
    "url": "http://example.org/cat.jpeg"
  }
}`

func example66Type() vocab.NoteInterface {
	example66Type := NewActivityStreamsNote()
	l := MustParseURL("http://example.org/cat.jpeg")
	image := NewActivityStreamsImage()
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is what he looks like.")
	image.SetContent(content)
	imgProp := NewActivityStreamsUrlProperty()
	imgProp.AppendIRI(l)
	image.SetUrl(imgProp)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Have you seen my cat?")
	example66Type.SetName(name)
	attachment := NewActivityStreamsAttachmentProperty()
	attachment.AppendImage(image)
	example66Type.SetAttachment(attachment)
	return example66Type
}

// NOTE: Changed to not be an array value for "attributedTo" to keep in line with other examples in spec!
const example67 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "My cat taking a nap",
  "url": "http://example.org/cat.jpeg",
  "attributedTo": {
    "type": "Person",
    "name": "Sally"
  }
}`

func example67Type() vocab.ImageInterface {
	example67Type := NewActivityStreamsImage()
	l := MustParseURL("http://example.org/cat.jpeg")
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	name := NewActivityStreamsNameProperty()
	name.AppendString("My cat taking a nap")
	example67Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(l)
	example67Type.SetUrl(urlProp)
	attr := NewActivityStreamsAttributedToProperty()
	attr.AppendPerson(person)
	example67Type.SetAttributedTo(attr)
	return example67Type
}

const example68 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "My cat taking a nap",
  "url": "http://example.org/cat.jpeg",
  "attributedTo": [
    "http://joe.example.org",
    {
      "type": "Person",
      "name": "Sally"
    }
  ]
}`

func example68Type() vocab.ImageInterface {
	example68Type := NewActivityStreamsImage()
	l := MustParseURL("http://example.org/cat.jpeg")
	a := MustParseURL("http://joe.example.org")
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	name := NewActivityStreamsNameProperty()
	name.AppendString("My cat taking a nap")
	example68Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(l)
	example68Type.SetUrl(urlProp)
	attr := NewActivityStreamsAttributedToProperty()
	attr.AppendIRI(a)
	example68Type.SetAttributedTo(attr)
	attrPerson := NewActivityStreamsAttributedToProperty()
	attrPerson.AppendPerson(person)
	example68Type.SetAttributedTo(attrPerson)
	return example68Type
}

const example69 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Holiday announcement",
  "type": "Note",
  "content": "Thursday will be a company-wide holiday. Enjoy your day off!",
  "audience": {
    "type": "http://example.org/Organization",
    "name": "ExampleCo LLC"
  }
}`

var example69Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["audience"] = map[string]interface{}{
		"type": "http://example.org/Organization",
		"name": "ExampleCo LLC",
	}
	return m
}

func example69Type() vocab.NoteInterface {
	example69Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Holiday announcement")
	example69Type.SetName(name)
	content := NewActivityStreamsContentProperty()
	content.AppendString("Thursday will be a company-wide holiday. Enjoy your day off!")
	example69Type.SetContent(content)
	return example69Type
}

// NOTE: Changed to not be an array value for "bcc" to keep in line with other examples in spec!
const example70 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered a post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "bcc": "http://joe.example.org"
}`

func example70Type() vocab.OfferInterface {
	example70Type := NewActivityStreamsOffer()
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered a post to John")
	example70Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example70Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example70Type.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example70Type.SetTarget(target)
	bcc := NewActivityStreamsBccProperty()
	bcc.AppendIRI(b)
	example70Type.SetBcc(bcc)
	return example70Type
}

// NOTE: Changed to not be an array value for "bto" to keep in line with other examples in spec!
const example71 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered a post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "bto": "http://joe.example.org"
}`

func example71Type() vocab.OfferInterface {
	example71Type := NewActivityStreamsOffer()
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered a post to John")
	example71Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example71Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example71Type.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example71Type.SetTarget(target)
	bto := NewActivityStreamsBtoProperty()
	bto.AppendIRI(b)
	example71Type.SetBto(bto)
	return example71Type
}

// NOTE: Changed to not be an array value for "cc" to keep in line with other examples in spec!
const example72 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered a post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "cc": "http://joe.example.org"
}`

func example72Type() vocab.OfferInterface {
	example72Type := NewActivityStreamsOffer()
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered a post to John")
	example72Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example72Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example72Type.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example72Type.SetTarget(target)
	cc := NewActivityStreamsCcProperty()
	cc.AppendIRI(b)
	example72Type.SetCc(cc)
	return example72Type
}

const example73 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Activities in context 1",
  "type": "Collection",
  "items": [
    {
      "type": "Offer",
      "actor": "http://sally.example.org",
      "object": "http://example.org/posts/1",
      "target": "http://john.example.org",
      "context": "http://example.org/contexts/1"
    },
    {
      "type": "Like",
      "actor": "http://joe.example.org",
      "object": "http://example.org/posts/2",
      "context": "http://example.org/contexts/1"
    }
  ]
}`

func example73Type() vocab.CollectionInterface {
	example73Type := NewActivityStreamsCollection()
	oa := MustParseURL("http://sally.example.org")
	oo := MustParseURL("http://example.org/posts/1")
	ot := MustParseURL("http://john.example.org")
	oc := MustParseURL("http://example.org/contexts/1")
	offer := NewActivityStreamsOffer()
	offerActor := NewActivityStreamsActorProperty()
	offerActor.AppendIRI(oa)
	offer.SetActor(offerActor)
	objOffer := NewActivityStreamsObjectProperty()
	objOffer.AppendIRI(oo)
	offer.SetObject(objOffer)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(ot)
	offer.SetTarget(target)
	ctx := NewActivityStreamsContextProperty()
	ctx.AppendIRI(oc)
	offer.SetContext(ctx)
	la := MustParseURL("http://joe.example.org")
	lo := MustParseURL("http://example.org/posts/2")
	lc := MustParseURL("http://example.org/contexts/1")
	like := NewActivityStreamsLike()
	likeActor := NewActivityStreamsActorProperty()
	likeActor.AppendIRI(la)
	like.SetActor(likeActor)
	objLike := NewActivityStreamsObjectProperty()
	objLike.AppendIRI(lo)
	like.SetObject(objLike)
	ctxLike := NewActivityStreamsContextProperty()
	ctxLike.AppendIRI(lc)
	like.SetContext(ctxLike)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Activities in context 1")
	example73Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendOffer(offer)
	items.AppendLike(like)
	example73Type.SetItems(items)
	return example73Type
}

const example74 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "current": "http://example.org/collection",
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example74Type() vocab.CollectionInterface {
	example74Type := NewActivityStreamsCollection()
	c := MustParseURL("http://example.org/collection")
	i1 := MustParseURL("http://example.org/posts/1")
	i2 := MustParseURL("http://example.org/posts/2")
	i3 := MustParseURL("http://example.org/posts/3")
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example74Type.SetTotalItems(totalItems)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's blog posts")
	example74Type.SetSummary(summary)
	current := NewActivityStreamsCurrentProperty()
	current.SetIRI(c)
	example74Type.SetCurrent(current)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(i1)
	items.AppendIRI(i2)
	items.AppendIRI(i3)
	example74Type.SetItems(items)
	return example74Type
}

const example75 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "current": {
    "type": "Link",
    "summary": "Most Recent Items",
    "href": "http://example.org/collection"
  },
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example75Type() vocab.CollectionInterface {
	example75Type := NewActivityStreamsCollection()
	i1 := MustParseURL("http://example.org/posts/1")
	i2 := MustParseURL("http://example.org/posts/2")
	i3 := MustParseURL("http://example.org/posts/3")
	href := MustParseURL("http://example.org/collection")
	link := NewActivityStreamsLink()
	summaryLink := NewActivityStreamsSummaryProperty()
	summaryLink.AppendString("Most Recent Items")
	link.SetSummary(summaryLink)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(href)
	link.SetHref(hrefLink)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example75Type.SetTotalItems(totalItems)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's blog posts")
	example75Type.SetSummary(summary)
	current := NewActivityStreamsCurrentProperty()
	current.SetLink(link)
	example75Type.SetCurrent(current)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(i1)
	items.AppendIRI(i2)
	items.AppendIRI(i3)
	example75Type.SetItems(items)
	return example75Type
}

const example76 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "first": "http://example.org/collection?page=0"
}`

func example76Type() vocab.CollectionInterface {
	example76Type := NewActivityStreamsCollection()
	f := MustParseURL("http://example.org/collection?page=0")
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example76Type.SetTotalItems(totalItems)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's blog posts")
	example76Type.SetSummary(summary)
	first := NewActivityStreamsFirstProperty()
	first.SetIRI(f)
	example76Type.SetFirst(first)
	return example76Type
}

const example77 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "first": {
    "type": "Link",
    "summary": "First Page",
    "href": "http://example.org/collection?page=0"
  }
}`

func example77Type() vocab.CollectionInterface {
	example77Type := NewActivityStreamsCollection()
	href := MustParseURL("http://example.org/collection?page=0")
	link := NewActivityStreamsLink()
	summaryLink := NewActivityStreamsSummaryProperty()
	summaryLink.AppendString("First Page")
	link.SetSummary(summaryLink)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(href)
	link.SetHref(hrefLink)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example77Type.SetTotalItems(totalItems)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's blog posts")
	example77Type.SetSummary(summary)
	first := NewActivityStreamsFirstProperty()
	first.SetLink(link)
	example77Type.SetFirst(first)
	return example77Type
}

const example78 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "generator": {
    "type": "Application",
    "name": "Exampletron 3000"
  }
}`

func example78Type() vocab.NoteInterface {
	example78Type := NewActivityStreamsNote()
	app := NewActivityStreamsApplication()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Exampletron 3000")
	app.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example78Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example78Type.SetContent(content)
	gen := NewActivityStreamsGeneratorProperty()
	gen.AppendApplication(app)
	example78Type.SetGenerator(gen)
	return example78Type
}

const example79 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "icon": {
    "type": "Image",
    "name": "Note icon",
    "url": "http://example.org/note.png",
    "width": 16,
    "height": 16
  }
}`

func example79Type() vocab.NoteInterface {
	example79Type := NewActivityStreamsNote()
	u := MustParseURL("http://example.org/note.png")
	image := NewActivityStreamsImage()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Note icon")
	image.SetName(name)
	imgProp := NewActivityStreamsUrlProperty()
	imgProp.AppendIRI(u)
	image.SetUrl(imgProp)
	width := NewActivityStreamsWidthProperty()
	width.Set(16)
	image.SetWidth(width)
	height := NewActivityStreamsHeightProperty()
	height.Set(16)
	image.SetHeight(height)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example79Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example79Type.SetContent(content)
	icon := NewActivityStreamsIconProperty()
	icon.AppendImage(image)
	example79Type.SetIcon(icon)
	return example79Type
}

const example80 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "icon": [
    {
      "type": "Image",
      "summary": "Note (16x16)",
      "url": "http://example.org/note1.png",
      "width": 16,
      "height": 16
    },
    {
      "type": "Image",
      "summary": "Note (32x32)",
      "url": "http://example.org/note2.png",
      "width": 32,
      "height": 32
    }
  ]
}`

func example80Type() vocab.NoteInterface {
	example80Type := NewActivityStreamsNote()
	u1 := MustParseURL("http://example.org/note1.png")
	u2 := MustParseURL("http://example.org/note2.png")
	image1 := NewActivityStreamsImage()
	summaryImg1 := NewActivityStreamsSummaryProperty()
	summaryImg1.AppendString("Note (16x16)")
	image1.SetSummary(summaryImg1)
	imgProp1 := NewActivityStreamsUrlProperty()
	imgProp1.AppendIRI(u1)
	image1.SetUrl(imgProp1)
	width1 := NewActivityStreamsWidthProperty()
	width1.Set(16)
	image1.SetWidth(width1)
	height1 := NewActivityStreamsHeightProperty()
	height1.Set(16)
	image1.SetHeight(height1)
	image2 := NewActivityStreamsImage()
	summaryImg2 := NewActivityStreamsSummaryProperty()
	summaryImg2.AppendString("Note (32x32)")
	image2.SetSummary(summaryImg2)
	imgProp2 := NewActivityStreamsUrlProperty()
	imgProp2.AppendIRI(u2)
	image2.SetUrl(imgProp2)
	width2 := NewActivityStreamsWidthProperty()
	width2.Set(32)
	image2.SetWidth(width2)
	height2 := NewActivityStreamsHeightProperty()
	height2.Set(32)
	image2.SetHeight(height2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example80Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("A simple note")
	example80Type.SetContent(content)
	icon := NewActivityStreamsIconProperty()
	icon.AppendImage(image1)
	icon.AppendImage(image2)
	example80Type.SetIcon(icon)
	return example80Type
}

const example81 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "image": {
    "type": "Image",
    "name": "A Cat",
    "url": "http://example.org/cat.png"
  }
}`

func example81Type() vocab.NoteInterface {
	example81Type := NewActivityStreamsNote()
	u := MustParseURL("http://example.org/cat.png")
	image := NewActivityStreamsImage()
	imageName := NewActivityStreamsNameProperty()
	imageName.AppendString("A Cat")
	image.SetName(imageName)
	imgProp := NewActivityStreamsUrlProperty()
	imgProp.AppendIRI(u)
	image.SetUrl(imgProp)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A simple note")
	example81Type.SetName(name)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example81Type.SetContent(content)
	imageProp := NewActivityStreamsImageProperty()
	imageProp.AppendImage(image)
	example81Type.SetImage(imageProp)
	return example81Type
}

const example82 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "image": [
    {
      "type": "Image",
      "name": "Cat 1",
      "url": "http://example.org/cat1.png"
    },
    {
      "type": "Image",
      "name": "Cat 2",
      "url": "http://example.org/cat2.png"
    }
  ]
}`

func example82Type() vocab.NoteInterface {
	example82Type := NewActivityStreamsNote()
	u1 := MustParseURL("http://example.org/cat1.png")
	u2 := MustParseURL("http://example.org/cat2.png")
	image1 := NewActivityStreamsImage()
	image1Name := NewActivityStreamsNameProperty()
	image1Name.AppendString("Cat 1")
	image1.SetName(image1Name)
	imgProp1 := NewActivityStreamsUrlProperty()
	imgProp1.AppendIRI(u1)
	image1.SetUrl(imgProp1)
	image2 := NewActivityStreamsImage()
	image2Name := NewActivityStreamsNameProperty()
	image2Name.AppendString("Cat 2")
	image2.SetName(image2Name)
	imgProp2 := NewActivityStreamsUrlProperty()
	imgProp2.AppendIRI(u2)
	image2.SetUrl(imgProp2)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A simple note")
	example82Type.SetName(name)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example82Type.SetContent(content)
	imageProp := NewActivityStreamsImageProperty()
	imageProp.AppendImage(image1)
	imageProp.AppendImage(image2)
	example82Type.SetImage(imageProp)
	return example82Type
}

const example83 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "inReplyTo": {
    "summary": "Previous note",
    "type": "Note",
    "content": "What else is there?"
  }
}`

func example83Type() vocab.NoteInterface {
	example83Type := NewActivityStreamsNote()
	note := NewActivityStreamsNote()
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("Previous note")
	note.SetSummary(summaryNote)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("What else is there?")
	note.SetContent(contentNote)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example83Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example83Type.SetContent(content)
	inReplyTo := NewActivityStreamsInReplyToProperty()
	inReplyTo.AppendNote(note)
	example83Type.SetInReplyTo(inReplyTo)
	return example83Type
}

const example84 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "inReplyTo": "http://example.org/posts/1"
}`

func example84Type() vocab.NoteInterface {
	example84Type := NewActivityStreamsNote()
	u := MustParseURL("http://example.org/posts/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example84Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("This is all there is.")
	example84Type.SetContent(content)
	inReplyTo := NewActivityStreamsInReplyToProperty()
	inReplyTo.AppendIRI(u)
	example84Type.SetInReplyTo(inReplyTo)
	return example84Type
}

const example85 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally listened to a piece of music on the Acme Music Service",
  "type": "Listen",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/foo.mp3",
  "instrument": {
    "type": "Service",
    "name": "Acme Music Service"
  }
}`

func example85Type() vocab.ListenInterface {
	example85Type := NewActivityStreamsListen()
	u := MustParseURL("http://example.org/foo.mp3")
	actor := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	actor.SetName(sally)
	service := NewActivityStreamsService()
	serviceName := NewActivityStreamsNameProperty()
	serviceName.AppendString("Acme Music Service")
	service.SetName(serviceName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally listened to a piece of music on the Acme Music Service")
	example85Type.SetSummary(summary)
	rootActor := NewActivityStreamsActorProperty()
	rootActor.AppendPerson(actor)
	example85Type.SetActor(rootActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(u)
	example85Type.SetObject(obj)
	inst := NewActivityStreamsInstrumentProperty()
	inst.AppendService(service)
	example85Type.SetInstrument(inst)
	return example85Type
}

const example86 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A collection",
  "type": "Collection",
  "totalItems": 3,
  "last": "http://example.org/collection?page=1"
}`

func example86Type() vocab.CollectionInterface {
	example86Type := NewActivityStreamsCollection()
	u := MustParseURL("http://example.org/collection?page=1")
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	example86Type.SetTotalItems(totalItems)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A collection")
	example86Type.SetSummary(summary)
	last := NewActivityStreamsLastProperty()
	last.SetIRI(u)
	example86Type.SetLast(last)
	return example86Type
}

const example87 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A collection",
  "type": "Collection",
  "totalItems": 5,
  "last": {
    "type": "Link",
    "summary": "Last Page",
    "href": "http://example.org/collection?page=1"
  }
}`

func example87Type() vocab.CollectionInterface {
	example87Type := NewActivityStreamsCollection()
	u := MustParseURL("http://example.org/collection?page=1")
	link := NewActivityStreamsLink()
	summaryLink := NewActivityStreamsSummaryProperty()
	summaryLink.AppendString("Last Page")
	link.SetSummary(summaryLink)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	link.SetHref(hrefLink)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A collection")
	example87Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(5)
	example87Type.SetTotalItems(totalItems)
	last := NewActivityStreamsLastProperty()
	last.SetLink(link)
	example87Type.SetLast(last)
	return example87Type
}

const example88 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally",
  "location": {
    "name": "Over the Arabian Sea, east of Socotra Island Nature Sanctuary",
    "type": "Place",
    "longitude": 12.34,
    "latitude": 56.78,
    "altitude": 90,
    "units": "m"
  }
}`

func example88Type() vocab.PersonInterface {
	example88Type := NewActivityStreamsPerson()
	place := NewActivityStreamsPlace()
	placeName := NewActivityStreamsNameProperty()
	placeName.AppendString("Over the Arabian Sea, east of Socotra Island Nature Sanctuary")
	place.SetName(placeName)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(12.34)
	place.SetLongitude(lon)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(56.78)
	place.SetLatitude(lat)
	alt := NewActivityStreamsAltitudeProperty()
	alt.Set(90)
	place.SetAltitude(alt)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("m")
	place.SetUnits(units)
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	example88Type.SetName(sally)
	location := NewActivityStreamsLocationProperty()
	location.AppendPlace(place)
	example88Type.SetLocation(location)
	return example88Type
}

const example89 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's notes",
  "type": "Collection",
  "totalItems": 2,
  "items": [
    {
      "type": "Note",
      "name": "Reminder for Going-Away Party"
    },
    {
      "type": "Note",
      "name": "Meeting 2016-11-17"
    }
  ]
}`

func example89Type() vocab.CollectionInterface {
	example89Type := NewActivityStreamsCollection()
	note1 := NewActivityStreamsNote()
	note1Name := NewActivityStreamsNameProperty()
	note1Name.AppendString("Reminder for Going-Away Party")
	note1.SetName(note1Name)
	note2 := NewActivityStreamsNote()
	note2Name := NewActivityStreamsNameProperty()
	note2Name.AppendString("Meeting 2016-11-17")
	note2.SetName(note2Name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's notes")
	example89Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(2)
	example89Type.SetTotalItems(totalItems)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note1)
	items.AppendNote(note2)
	example89Type.SetItems(items)
	return example89Type
}

const example90 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's notes",
  "type": "OrderedCollection",
  "totalItems": 2,
  "orderedItems": [
    {
      "type": "Note",
      "name": "Meeting 2016-11-17"
    },
    {
      "type": "Note",
      "name": "Reminder for Going-Away Party"
    }
  ]
}`

func example90Type() vocab.OrderedCollectionInterface {
	example90Type := NewActivityStreamsOrderedCollection()
	note1 := NewActivityStreamsNote()
	note1Name := NewActivityStreamsNameProperty()
	note1Name.AppendString("Meeting 2016-11-17")
	note1.SetName(note1Name)
	note2 := NewActivityStreamsNote()
	note2Name := NewActivityStreamsNameProperty()
	note2Name.AppendString("Reminder for Going-Away Party")
	note2.SetName(note2Name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's notes")
	example90Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(2)
	example90Type.SetTotalItems(totalItems)
	orderedItems := NewActivityStreamsOrderedItemsProperty()
	orderedItems.AppendNote(note1)
	orderedItems.AppendNote(note2)
	example90Type.SetOrderedItems(orderedItems)
	return example90Type
}

const example91 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "oneOf": [
    {
      "type": "Note",
      "name": "Option A"
    },
    {
      "type": "Note",
      "name": "Option B"
    }
  ]
}`

func example91Type() vocab.QuestionInterface {
	example91Type := NewActivityStreamsQuestion()
	note1 := NewActivityStreamsNote()
	note1Name := NewActivityStreamsNameProperty()
	note1Name.AppendString("Option A")
	note1.SetName(note1Name)
	note2 := NewActivityStreamsNote()
	note2Name := NewActivityStreamsNameProperty()
	note2Name.AppendString("Option B")
	note2.SetName(note2Name)
	name := NewActivityStreamsNameProperty()
	name.AppendString("What is the answer?")
	example91Type.SetName(name)
	oneOf1 := NewActivityStreamsOneOfProperty()
	oneOf1.AppendNote(note1)
	example91Type.SetOneOf(oneOf1)
	oneOf2 := NewActivityStreamsOneOfProperty()
	oneOf2.AppendNote(note2)
	example91Type.SetOneOf(oneOf2)
	return example91Type
}

const example92 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "anyOf": [
    {
      "type": "Note",
      "name": "Option A"
    },
    {
      "type": "Note",
      "name": "Option B"
    }
  ]
}`

func example92Type() vocab.QuestionInterface {
	example92Type := NewActivityStreamsQuestion()
	note1 := NewActivityStreamsNote()
	note1Name := NewActivityStreamsNameProperty()
	note1Name.AppendString("Option A")
	note1.SetName(note1Name)
	note2 := NewActivityStreamsNote()
	note2Name := NewActivityStreamsNameProperty()
	note2Name.AppendString("Option B")
	note2.SetName(note2Name)
	name := NewActivityStreamsNameProperty()
	name.AppendString("What is the answer?")
	example92Type.SetName(name)
	anyOf := NewActivityStreamsAnyOfProperty()
	anyOf.AppendNote(note1)
	anyOf.AppendNote(note2)
	example92Type.SetAnyOf(anyOf)
	return example92Type
}

const example93 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "closed": "2016-05-10T00:00:00Z"
}`

func example93Type() vocab.QuestionInterface {
	example93Type := NewActivityStreamsQuestion()
	t, err := time.Parse(time.RFC3339, "2016-05-10T00:00:00Z")
	if err != nil {
		panic(err)
	}
	name := NewActivityStreamsNameProperty()
	name.AppendString("What is the answer?")
	example93Type.SetName(name)
	closed := NewActivityStreamsClosedProperty()
	closed.AppendDateTime(t)
	example93Type.SetClosed(closed)
	return example93Type
}

const example94 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally moved a post from List A to List B",
  "type": "Move",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": {
    "type": "Collection",
    "name": "List B"
  },
  "origin": {
    "type": "Collection",
    "name": "List A"
  }
}`

func example94Type() vocab.MoveInterface {
	example94Type := NewActivityStreamsMove()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	target := NewActivityStreamsCollection()
	targetName := NewActivityStreamsNameProperty()
	targetName.AppendString("List B")
	target.SetName(targetName)
	origin := NewActivityStreamsCollection()
	originName := NewActivityStreamsNameProperty()
	originName.AppendString("List A")
	origin.SetName(originName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally moved a post from List A to List B")
	example94Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example94Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example94Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendCollection(target)
	example94Type.SetTarget(tobj)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendCollection(origin)
	example94Type.SetOrigin(originProp)
	return example94Type
}

const example95 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 2 of Sally's blog posts",
  "type": "CollectionPage",
  "next": "http://example.org/collection?page=2",
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example95Type() vocab.CollectionPageInterface {
	example95Type := NewActivityStreamsCollectionPage()
	i := MustParseURL("http://example.org/collection?page=2")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 2 of Sally's blog posts")
	example95Type.SetSummary(summary)
	next := NewActivityStreamsNextProperty()
	next.SetIRI(i)
	example95Type.SetNext(next)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(u1)
	items.AppendIRI(u2)
	items.AppendIRI(u3)
	example95Type.SetItems(items)
	return example95Type
}

const example96 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 2 of Sally's blog posts",
  "type": "CollectionPage",
  "next": {
    "type": "Link",
    "name": "Next Page",
    "href": "http://example.org/collection?page=2"
  },
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example96Type() vocab.CollectionPageInterface {
	example96Type := NewActivityStreamsCollectionPage()
	href := MustParseURL("http://example.org/collection?page=2")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	link := NewActivityStreamsLink()
	linkName := NewActivityStreamsNameProperty()
	linkName.AppendString("Next Page")
	link.SetName(linkName)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(href)
	link.SetHref(hrefLink)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 2 of Sally's blog posts")
	example96Type.SetSummary(summary)
	next := NewActivityStreamsNextProperty()
	next.SetLink(link)
	example96Type.SetNext(next)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(u1)
	items.AppendIRI(u2)
	items.AppendIRI(u3)
	example96Type.SetItems(items)
	return example96Type
}

const example97 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally liked a post",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`

func example97Type() vocab.LikeInterface {
	example97Type := NewActivityStreamsLike()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally liked a post")
	example97Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example97Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example97Type.SetObject(obj)
	return example97Type
}

const example98 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Note",
    "content": "A simple note"
  }
}`

func example98Type() vocab.LikeInterface {
	example98Type := NewActivityStreamsLike()
	a := MustParseURL("http://sally.example.org")
	note := NewActivityStreamsNote()
	content := NewActivityStreamsContentProperty()
	content.AppendString("A simple note")
	note.SetContent(content)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example98Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendNote(note)
	example98Type.SetObject(obj)
	return example98Type
}

const example99 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally liked a note",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": [
    "http://example.org/posts/1",
    {
      "type": "Note",
      "summary": "A simple note",
      "content": "That is a tree."
    }
  ]
}`

func example99Type() vocab.LikeInterface {
	example99Type := NewActivityStreamsLike()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	note := NewActivityStreamsNote()
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("A simple note")
	note.SetSummary(summaryNote)
	content := NewActivityStreamsContentProperty()
	content.AppendString("That is a tree.")
	note.SetContent(content)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally liked a note")
	example99Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example99Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example99Type.SetObject(obj)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendNote(note)
	example99Type.SetObject(objRoot)
	return example99Type
}

const example100 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's blog posts",
  "type": "CollectionPage",
  "prev": "http://example.org/collection?page=1",
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example100Type() vocab.CollectionPageInterface {
	example100Type := NewActivityStreamsCollectionPage()
	p := MustParseURL("http://example.org/collection?page=1")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's blog posts")
	example100Type.SetSummary(summary)
	prev := NewActivityStreamsPrevProperty()
	prev.SetIRI(p)
	example100Type.SetPrev(prev)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(u1)
	items.AppendIRI(u2)
	items.AppendIRI(u3)
	example100Type.SetItems(items)
	return example100Type
}

const example101 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's blog posts",
  "type": "CollectionPage",
  "prev": {
    "type": "Link",
    "name": "Previous Page",
    "href": "http://example.org/collection?page=1"
  },
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`

func example101Type() vocab.CollectionPageInterface {
	example101Type := NewActivityStreamsCollectionPage()
	p := MustParseURL("http://example.org/collection?page=1")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	link := NewActivityStreamsLink()
	linkName := NewActivityStreamsNameProperty()
	linkName.AppendString("Previous Page")
	link.SetName(linkName)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(p)
	link.SetHref(hrefLink)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's blog posts")
	example101Type.SetSummary(summary)
	prev := NewActivityStreamsPrevProperty()
	prev.SetLink(link)
	example101Type.SetPrev(prev)
	items := NewActivityStreamsItemsProperty()
	items.AppendIRI(u1)
	items.AppendIRI(u2)
	items.AppendIRI(u3)
	example101Type.SetItems(items)
	return example101Type
}

// NOTE: The 'url' field has added the 'type' property
const example102 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Cool New Movie",
  "duration": "PT2H30M",
  "preview": {
    "type": "Video",
    "name": "Trailer",
    "duration": "PT1M",
    "url": {
      "type": "Link",
      "href": "http://example.org/trailer.mkv",
      "mediaType": "video/mkv"
    }
  }
}`

func example102Type() vocab.VideoInterface {
	example102Type := NewActivityStreamsVideo()
	u := MustParseURL("http://example.org/trailer.mkv")
	link := NewActivityStreamsLink()
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("video/mkv")
	link.SetMediaType(mediaType)
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	link.SetHref(hrefLink)
	video := NewActivityStreamsVideo()
	videoName := NewActivityStreamsNameProperty()
	videoName.AppendString("Trailer")
	video.SetName(videoName)
	durVideo := NewActivityStreamsDurationProperty()
	durVideo.Set(time.Minute)
	video.SetDuration(durVideo)
	urlProperty := NewActivityStreamsUrlProperty()
	urlProperty.AppendLink(link)
	video.SetUrl(urlProperty)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Cool New Movie")
	example102Type.SetName(name)
	dur := NewActivityStreamsDurationProperty()
	dur.Set(time.Hour*2 + time.Minute*30)
	example102Type.SetDuration(dur)
	preview := NewActivityStreamsPreviewProperty()
	preview.AppendVideo(video)
	example102Type.SetPreview(preview)
	return example102Type
}

const example103 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally checked that her flight was on time",
  "type": ["Activity", "http://www.verbs.example/Check"],
  "actor": "http://sally.example.org",
  "object": "http://example.org/flights/1",
  "result": {
    "type": "http://www.types.example/flightstatus",
    "name": "On Time"
  }
}`

var example103Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["result"] = map[string]interface{}{
		"type": "http://www.types.example/flightstatus",
		"name": "On Time",
	}
	return m
}

func example103Type() vocab.ActivityInterface {
	example103Type := NewActivityStreamsActivity()
	o := MustParseURL("http://example.org/flights/1")
	a := MustParseURL("http://sally.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally checked that her flight was on time")
	example103Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example103Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example103Type.SetObject(obj)
	return example103Type
}

// NOTE: Changed to not be an array value for "items" to keep in line with other examples in spec!
const example104 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "id": "http://www.test.example/notes/1",
  "content": "I am fine.",
  "replies": {
    "type": "Collection",
    "totalItems": 1,
    "items": {
      "summary": "A response to the note",
      "type": "Note",
      "content": "I am glad to hear it.",
      "inReplyTo": "http://www.test.example/notes/1"
    }
  }
}`

func example104Type() vocab.NoteInterface {
	example104Type := NewActivityStreamsNote()
	i := MustParseURL("http://www.test.example/notes/1")
	note := NewActivityStreamsNote()
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("A response to the note")
	note.SetSummary(summaryNote)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("I am glad to hear it.")
	note.SetContent(contentNote)
	inReplyTo := NewActivityStreamsInReplyToProperty()
	inReplyTo.AppendIRI(i)
	note.SetInReplyTo(inReplyTo)
	replies := NewActivityStreamsCollection()
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(1)
	replies.SetTotalItems(totalItems)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note)
	replies.SetItems(items)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example104Type.SetSummary(summary)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://www.test.example/notes/1"))
	example104Type.SetId(id)
	content := NewActivityStreamsContentProperty()
	content.AppendString("I am fine.")
	example104Type.SetContent(content)
	reply := NewActivityStreamsRepliesProperty()
	reply.SetCollection(replies)
	example104Type.SetReplies(reply)
	return example104Type
}

// NOTE: Changed to not be an array value for "tag" to keep in line with other examples in spec!
const example105 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Image",
  "summary": "Picture of Sally",
  "url": "http://example.org/sally.jpg",
  "tag": {
    "type": "Person",
    "id": "http://sally.example.org",
    "name": "Sally"
  }
}`

func example105Type() vocab.ImageInterface {
	example105Type := NewActivityStreamsImage()
	u := MustParseURL("http://example.org/sally.jpg")
	person := NewActivityStreamsPerson()
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://sally.example.org"))
	person.SetId(id)
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Picture of Sally")
	example105Type.SetSummary(summary)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(u)
	example105Type.SetUrl(urlProp)
	tag := NewActivityStreamsTagProperty()
	tag.AppendPerson(person)
	example105Type.SetTag(tag)
	return example105Type
}

const example106 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org"
}`

func example106Type() vocab.OfferInterface {
	example106Type := NewActivityStreamsOffer()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	t := MustParseURL("http://john.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered the post to John")
	example106Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example106Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example106Type.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example106Type.SetTarget(target)
	return example106Type
}

const example107 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": {
    "type": "Person",
    "name": "John"
  }
}`

func example107Type() vocab.OfferInterface {
	example107Type := NewActivityStreamsOffer()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	person := NewActivityStreamsPerson()
	personName := NewActivityStreamsNameProperty()
	personName.AppendString("John")
	person.SetName(personName)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered the post to John")
	example107Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example107Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example107Type.SetObject(obj)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendPerson(person)
	example107Type.SetTarget(tobj)
	return example107Type
}

// NOTE: Changed to not be an array value for "to" to keep in line with other examples in spec!
const example108 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "to": "http://joe.example.org"
}`

func example108Type() vocab.OfferInterface {
	example108Type := NewActivityStreamsOffer()
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	t := MustParseURL("http://john.example.org")
	z := MustParseURL("http://joe.example.org")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally offered the post to John")
	example108Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(a)
	example108Type.SetActor(objectActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(o)
	example108Type.SetObject(obj)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example108Type.SetTarget(target)
	to := NewActivityStreamsToProperty()
	to.AppendIRI(z)
	example108Type.SetTo(to)
	return example108Type
}

const example109 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": "http://example.org/4q-sales-forecast.pdf"
}`

func example109Type() vocab.DocumentInterface {
	example109Type := NewActivityStreamsDocument()
	u := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	name := NewActivityStreamsNameProperty()
	name.AppendString("4Q Sales Forecast")
	example109Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(u)
	example109Type.SetUrl(urlProp)
	return example109Type
}

const example110 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": {
    "type": "Link",
    "href": "http://example.org/4q-sales-forecast.pdf"
  }
}`

func example110Type() vocab.DocumentInterface {
	example110Type := NewActivityStreamsDocument()
	u := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	link := NewActivityStreamsLink()
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	link.SetHref(hrefLink)
	name := NewActivityStreamsNameProperty()
	name.AppendString("4Q Sales Forecast")
	example110Type.SetName(name)
	urlProperty := NewActivityStreamsUrlProperty()
	urlProperty.AppendLink(link)
	example110Type.SetUrl(urlProperty)
	return example110Type
}

const example111 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": [
    {
      "type": "Link",
      "href": "http://example.org/4q-sales-forecast.pdf",
      "mediaType": "application/pdf"
    },
    {
      "type": "Link",
      "href": "http://example.org/4q-sales-forecast.html",
      "mediaType": "text/html"
    }
  ]
}`

func example111Type() vocab.DocumentInterface {
	example111Type := NewActivityStreamsDocument()
	u1 := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	u2 := MustParseURL("http://example.org/4q-sales-forecast.html")
	link1 := NewActivityStreamsLink()
	hrefLink1 := NewActivityStreamsHrefProperty()
	hrefLink1.Set(u1)
	link1.SetHref(hrefLink1)
	mediaType1 := NewActivityStreamsMediaTypeProperty()
	mediaType1.Set("application/pdf")
	link1.SetMediaType(mediaType1)
	link2 := NewActivityStreamsLink()
	hrefLink2 := NewActivityStreamsHrefProperty()
	hrefLink2.Set(u2)
	link2.SetHref(hrefLink2)
	mediaType2 := NewActivityStreamsMediaTypeProperty()
	mediaType2.Set("text/html")
	link2.SetMediaType(mediaType2)
	name := NewActivityStreamsNameProperty()
	name.AppendString("4Q Sales Forecast")
	example111Type.SetName(name)
	urlProperty1 := NewActivityStreamsUrlProperty()
	urlProperty1.AppendLink(link1)
	example111Type.SetUrl(urlProperty1)
	urlProperty2 := NewActivityStreamsUrlProperty()
	urlProperty2.AppendLink(link2)
	example111Type.SetUrl(urlProperty2)
	return example111Type
}

const example112 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China",
  "type": "Place",
  "latitude": 36.75,
  "longitude": 119.7667,
  "accuracy": 94.5
}`

func example112Type() vocab.PlaceInterface {
	example112Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China")
	example112Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example112Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example112Type.SetLongitude(lon)
	acc := NewActivityStreamsAccuracyProperty()
	acc.Set(94.5)
	example112Type.SetAccuracy(acc)
	return example112Type
}

const example113 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "altitude": 15.0,
  "latitude": 36.75,
  "longitude": 119.7667,
  "units": "miles"
}`

func example113Type() vocab.PlaceInterface {
	example113Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example113Type.SetName(name)
	alt := NewActivityStreamsAltitudeProperty()
	alt.Set(15.0)
	example113Type.SetAltitude(alt)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example113Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example113Type.SetLongitude(lon)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example113Type.SetUnits(units)
	return example113Type
}

const example114 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "A <em>simple</em> note"
}`

func example114Type() vocab.NoteInterface {
	example114Type := NewActivityStreamsNote()
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example114Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("A <em>simple</em> note")
	example114Type.SetContent(content)
	return example114Type
}

const example115 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "contentMap": {
    "en": "A <em>simple</em> note",
    "es": "Una nota <em>sencilla</em>",
    "zh-Hans": "一段<em>简单的</em>笔记"
  }
}`

func example115Type() vocab.NoteInterface {
	example115Type := NewActivityStreamsNote()
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example115Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendLangString(map[string]string{
		"en":      "A <em>simple</em> note",
		"es":      "Una nota <em>sencilla</em>",
		"zh-Hans": "一段<em>简单的</em>笔记",
	})
	example115Type.SetContent(content)
	return example115Type
}

const example116 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "mediaType": "text/markdown",
  "content": "## A simple note\nA simple markdown ` + "`note`" + `"
}`

func example116Type() vocab.NoteInterface {
	example116Type := NewActivityStreamsNote()
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example116Type.SetSummary(summary)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/markdown")
	example116Type.SetMediaType(mediaType)
	content := NewActivityStreamsContentProperty()
	content.AppendString("## A simple note\nA simple markdown `note`")
	example116Type.SetContent(content)
	return example116Type
}

const example117 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A simple note"
}`

func example117Type() vocab.NoteInterface {
	example117Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendString("A simple note")
	example117Type.SetName(name)
	return example117Type
}

const example118 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "nameMap": {
    "en": "A simple note",
    "es": "Una nota sencilla",
    "zh-Hans": "一段简单的笔记"
  }
}`

func example118Type() vocab.NoteInterface {
	example118Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendLangString(map[string]string{
		"en":      "A simple note",
		"es":      "Una nota sencilla",
		"zh-Hans": "一段简单的笔记",
	})
	example118Type.SetName(name)
	return example118Type
}

const example119 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Birds Flying",
  "url": "http://example.org/video.mkv",
  "duration": "PT2H"
}`

func example119Type() vocab.VideoInterface {
	example119Type := NewActivityStreamsVideo()
	u := MustParseURL("http://example.org/video.mkv")
	name := NewActivityStreamsNameProperty()
	name.AppendString("Birds Flying")
	example119Type.SetName(name)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(u)
	example119Type.SetUrl(urlProp)
	dur := NewActivityStreamsDurationProperty()
	dur.Set(time.Hour * 2)
	example119Type.SetDuration(dur)
	return example119Type
}

const example120 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/image.png",
  "height": 100,
  "width": 100
}`

func example120Type() vocab.LinkInterface {
	example120Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/image.png")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example120Type.SetHref(hrefLink)
	width := NewActivityStreamsWidthProperty()
	width.Set(100)
	example120Type.SetWidth(width)
	height := NewActivityStreamsHeightProperty()
	height.Set(100)
	example120Type.SetHeight(height)
	return example120Type
}

const example121 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "mediaType": "text/html",
  "name": "Previous"
}`

func example121Type() vocab.LinkInterface {
	example121Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/abc")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example121Type.SetHref(hrefLink)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/html")
	example121Type.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Previous")
	example121Type.SetName(name)
	return example121Type
}

const example122 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "Previous"
}`

func example122Type() vocab.LinkInterface {
	example122Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/abc")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example122Type.SetHref(hrefLink)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/html")
	example122Type.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Previous")
	example122Type.SetName(name)
	hreflang := NewActivityStreamsHreflangProperty()
	hreflang.Set("en")
	example122Type.SetHreflang(hreflang)
	return example122Type
}

const example123 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's notes",
  "type": "CollectionPage",
  "id": "http://example.org/collection?page=1",
  "partOf": "http://example.org/collection",
  "items": [
    {
      "type": "Note",
      "name": "Pizza Toppings to Try"
    },
    {
      "type": "Note",
      "name": "Thought about California"
    }
  ]
}`

func example123Type() vocab.CollectionPageInterface {
	example123Type := NewActivityStreamsCollectionPage()
	u := MustParseURL("http://example.org/collection")
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("Pizza Toppings to Try")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Thought about California")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's notes")
	example123Type.SetSummary(summary)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://example.org/collection?page=1"))
	example123Type.SetId(id)
	partOf := NewActivityStreamsPartOfProperty()
	partOf.SetIRI(u)
	example123Type.SetPartOf(partOf)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note1)
	items.AppendNote(note2)
	example123Type.SetItems(items)
	return example123Type
}

const example124 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`

func example124Type() vocab.PlaceInterface {
	example124Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example124Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example124Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example124Type.SetLongitude(lon)
	rad := NewActivityStreamsRadiusProperty()
	rad.Set(15)
	example124Type.SetRadius(rad)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example124Type.SetUnits(units)
	return example124Type
}

const example125 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`

func example125Type() vocab.PlaceInterface {
	example125Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example125Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example125Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example125Type.SetLongitude(lon)
	rad := NewActivityStreamsRadiusProperty()
	rad.Set(15)
	example125Type.SetRadius(rad)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example125Type.SetUnits(units)
	return example125Type
}

const example126 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "Next"
}`

func example126Type() vocab.LinkInterface {
	example126Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/abc")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example126Type.SetHref(hrefLink)
	hreflang := NewActivityStreamsHreflangProperty()
	hreflang.Set("en")
	example126Type.SetHreflang(hreflang)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/html")
	example126Type.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Next")
	example126Type.SetName(name)
	return example126Type
}

const example127 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

func example127Type() vocab.EventInterface {
	example127Type := NewActivityStreamsEvent()
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	example127Type.SetName(goingAway)
	startTime := NewActivityStreamsStartTimeProperty()
	startTime.Set(t1)
	example127Type.SetStartTime(startTime)
	endTime := NewActivityStreamsEndTimeProperty()
	endTime.Set(t2)
	example127Type.SetEndTime(endTime)
	return example127Type
}

const example128 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "Fish swim.",
  "published": "2014-12-12T12:12:12Z"
}`

func example128Type() vocab.NoteInterface {
	example128Type := NewActivityStreamsNote()
	t, err := time.Parse(time.RFC3339, "2014-12-12T12:12:12Z")
	if err != nil {
		panic(err)
	}
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple note")
	example128Type.SetSummary(summary)
	content := NewActivityStreamsContentProperty()
	content.AppendString("Fish swim.")
	example128Type.SetContent(content)
	published := NewActivityStreamsPublishedProperty()
	published.Set(t)
	example128Type.SetPublished(published)
	return example128Type
}

const example129 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

func example129Type() vocab.EventInterface {
	example129Type := NewActivityStreamsEvent()
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	goingAway := NewActivityStreamsNameProperty()
	goingAway.AppendString("Going-Away Party for Jim")
	example129Type.SetName(goingAway)
	startTime := NewActivityStreamsStartTimeProperty()
	startTime.Set(t1)
	example129Type.SetStartTime(startTime)
	endTime := NewActivityStreamsEndTimeProperty()
	endTime.Set(t2)
	example129Type.SetEndTime(endTime)
	return example129Type
}

const example130 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`

func example130Type() vocab.PlaceInterface {
	example130Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example130Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example130Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example130Type.SetLongitude(lon)
	rad := NewActivityStreamsRadiusProperty()
	rad.Set(15)
	example130Type.SetRadius(rad)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example130Type.SetUnits(units)
	return example130Type
}

const example131 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "Preview",
  "rel": ["canonical", "preview"]
}`

func example131Type() vocab.LinkInterface {
	example131Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/abc")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example131Type.SetHref(hrefLink)
	hreflang := NewActivityStreamsHreflangProperty()
	hreflang.Set("en")
	example131Type.SetHreflang(hreflang)
	mediaType := NewActivityStreamsMediaTypeProperty()
	mediaType.Set("text/html")
	example131Type.SetMediaType(mediaType)
	name := NewActivityStreamsNameProperty()
	name.AppendString("Preview")
	example131Type.SetName(name)
	rel := NewActivityStreamsRelProperty()
	rel.AppendRfc5988("canonical")
	rel.AppendRfc5988("preview")
	example131Type.SetRel(rel)
	return example131Type
}

const example132 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Page 1 of Sally's notes",
  "type": "OrderedCollectionPage",
  "startIndex": 0,
  "orderedItems": [
    {
      "type": "Note",
      "name": "Density of Water"
    },
    {
      "type": "Note",
      "name": "Air Mattress Idea"
    }
  ]
}`

func example132Type() vocab.OrderedCollectionPageInterface {
	example132Type := NewActivityStreamsOrderedCollectionPage()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("Density of Water")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Air Mattress Idea")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Page 1 of Sally's notes")
	example132Type.SetSummary(summary)
	start := NewActivityStreamsStartIndexProperty()
	start.Set(0)
	example132Type.SetStartIndex(start)
	orderedItems := NewActivityStreamsOrderedItemsProperty()
	orderedItems.AppendNote(note1)
	orderedItems.AppendNote(note2)
	example132Type.SetOrderedItems(orderedItems)
	return example132Type
}

const example133 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cane Sugar Processing",
  "type": "Note",
  "summary": "A simple <em>note</em>"
}`

func example133Type() vocab.NoteInterface {
	example133Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Cane Sugar Processing")
	example133Type.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("A simple <em>note</em>")
	example133Type.SetSummary(summary)
	return example133Type
}

const example134 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cane Sugar Processing",
  "type": "Note",
  "summaryMap": {
    "en": "A simple <em>note</em>",
    "es": "Una <em>nota</em> sencilla",
    "zh-Hans": "一段<em>简单的</em>笔记"
  }
}`

func example134Type() vocab.NoteInterface {
	example134Type := NewActivityStreamsNote()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Cane Sugar Processing")
	example134Type.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendLangString(map[string]string{
		"en":      "A simple <em>note</em>",
		"es":      "Una <em>nota</em> sencilla",
		"zh-Hans": "一段<em>简单的</em>笔记",
	})
	example134Type.SetSummary(summary)
	return example134Type
}

const example135 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's notes",
  "type": "Collection",
  "totalItems": 2,
  "items": [
    {
      "type": "Note",
      "name": "Which Staircase Should I Use"
    },
    {
      "type": "Note",
      "name": "Something to Remember"
    }
  ]
}`

func example135Type() vocab.CollectionInterface {
	example135Type := NewActivityStreamsCollection()
	note1 := NewActivityStreamsNote()
	name1 := NewActivityStreamsNameProperty()
	name1.AppendString("Which Staircase Should I Use")
	note1.SetName(name1)
	note2 := NewActivityStreamsNote()
	name2 := NewActivityStreamsNameProperty()
	name2.AppendString("Something to Remember")
	note2.SetName(name2)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's notes")
	example135Type.SetSummary(summary)
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(2)
	example135Type.SetTotalItems(totalItems)
	items := NewActivityStreamsItemsProperty()
	items.AppendNote(note1)
	items.AppendNote(note2)
	example135Type.SetItems(items)
	return example135Type
}

const example136 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`

func example136Type() vocab.PlaceInterface {
	example136Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Fresno Area")
	example136Type.SetName(name)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(36.75)
	example136Type.SetLatitude(lat)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(119.7667)
	example136Type.SetLongitude(lon)
	rad := NewActivityStreamsRadiusProperty()
	rad.Set(15)
	example136Type.SetRadius(rad)
	units := NewActivityStreamsUnitsProperty()
	units.SetString("miles")
	example136Type.SetUnits(units)
	return example136Type
}

const example137 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cranberry Sauce Idea",
  "type": "Note",
  "content": "Mush it up so it does not have the same shape as the can.",
  "updated": "2014-12-12T12:12:12Z"
}`

func example137Type() vocab.NoteInterface {
	example137Type := NewActivityStreamsNote()
	t, err := time.Parse(time.RFC3339, "2014-12-12T12:12:12Z")
	if err != nil {
		panic(err)
	}
	name := NewActivityStreamsNameProperty()
	name.AppendString("Cranberry Sauce Idea")
	example137Type.SetName(name)
	content := NewActivityStreamsContentProperty()
	content.AppendString("Mush it up so it does not have the same shape as the can.")
	example137Type.SetContent(content)
	updated := NewActivityStreamsUpdatedProperty()
	updated.Set(t)
	example137Type.SetUpdated(updated)
	return example137Type
}

const example138 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/image.png",
  "height": 100,
  "width": 100
}`

func example138Type() vocab.LinkInterface {
	example138Type := NewActivityStreamsLink()
	u := MustParseURL("http://example.org/image.png")
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	example138Type.SetHref(hrefLink)
	width := NewActivityStreamsWidthProperty()
	width.Set(100)
	example138Type.SetWidth(width)
	height := NewActivityStreamsHeightProperty()
	height.Set(100)
	example138Type.SetHeight(height)
	return example138Type
}

const example139 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally is an acquaintance of John's",
  "type": "Relationship",
  "subject": {
    "type": "Person",
    "name": "Sally"
  },
  "relationship": "http://purl.org/vocab/relationship/acquaintanceOf",
  "object": {
    "type": "Person",
    "name": "John"
  }
}`

func example139Type() vocab.RelationshipInterface {
	example139Type := NewActivityStreamsRelationship()
	u := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	subject := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	subject.SetName(sally)
	object := NewActivityStreamsPerson()
	name := NewActivityStreamsNameProperty()
	name.AppendString("John")
	object.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally is an acquaintance of John's")
	example139Type.SetSummary(summary)
	subj := NewActivityStreamsSubjectProperty()
	subj.SetPerson(subject)
	example139Type.SetSubject(subj)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example139Type.SetObject(obj)
	relationship := NewActivityStreamsRelationshipProperty()
	relationship.AppendIRI(u)
	example139Type.SetRelationship(relationship)
	return example139Type
}

const example140 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally is an acquaintance of John's",
  "type": "Relationship",
  "subject": {
    "type": "Person",
    "name": "Sally"
  },
  "relationship": "http://purl.org/vocab/relationship/acquaintanceOf",
  "object": {
    "type": "Person",
    "name": "John"
  }
}`

func example140Type() vocab.RelationshipInterface {
	example140Type := NewActivityStreamsRelationship()
	u := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	subject := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	subject.SetName(sally)
	object := NewActivityStreamsPerson()
	name := NewActivityStreamsNameProperty()
	name.AppendString("John")
	object.SetName(name)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally is an acquaintance of John's")
	example140Type.SetSummary(summary)
	subj := NewActivityStreamsSubjectProperty()
	subj.SetPerson(subject)
	example140Type.SetSubject(subj)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(object)
	example140Type.SetObject(obj)
	relationship := NewActivityStreamsRelationshipProperty()
	relationship.AppendIRI(u)
	example140Type.SetRelationship(relationship)
	return example140Type
}

const example141 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's profile",
  "type": "Profile",
  "describes": {
    "type": "Person",
    "name": "Sally"
  },
  "url": "http://sally.example.org"
}`

func example141Type() vocab.ProfileInterface {
	example141Type := NewActivityStreamsProfile()
	u := MustParseURL("http://sally.example.org")
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's profile")
	example141Type.SetSummary(summary)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(u)
	example141Type.SetUrl(urlProp)
	describes := NewActivityStreamsDescribesProperty()
	describes.SetPerson(person)
	example141Type.SetDescribes(describes)
	return example141Type
}

const example142 = `{
"@context": "https://www.w3.org/ns/activitystreams",
"summary": "This image has been deleted",
"type": "Tombstone",
"formerType": "Image",
"url": "http://example.org/image/2"
}`

func example142Type() vocab.TombstoneInterface {
	example142Type := NewActivityStreamsTombstone()
	u := MustParseURL("http://example.org/image/2")
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("This image has been deleted")
	example142Type.SetSummary(summary)
	former := NewActivityStreamsFormerTypeProperty()
	former.AppendString("Image")
	example142Type.SetFormerType(former)
	urlProp := NewActivityStreamsUrlProperty()
	urlProp.AppendIRI(u)
	example142Type.SetUrl(urlProp)
	return example142Type
}

const example143 = `{
"@context": "https://www.w3.org/ns/activitystreams",
"summary": "This image has been deleted",
"type": "Tombstone",
"deleted": "2016-05-03T00:00:00Z"
}`

func example143Type() vocab.TombstoneInterface {
	example143Type := NewActivityStreamsTombstone()
	t, err := time.Parse(time.RFC3339, "2016-05-03T00:00:00Z")
	if err != nil {
		panic(err)
	}
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("This image has been deleted")
	example143Type.SetSummary(summary)
	deleted := NewActivityStreamsDeletedProperty()
	deleted.Set(t)
	example143Type.SetDeleted(deleted)
	return example143Type
}

const example144 = `{
 "@context": "https://www.w3.org/ns/activitystreams",
 "summary": "Activities in Project XYZ",
 "type": "Collection",
 "items": [
   {
     "summary": "Sally created a note",
     "type": "Create",
     "id": "http://activities.example.com/1",
     "actor": "http://sally.example.org",
     "object": {
       "summary": "A note",
       "type": "Note",
       "id": "http://notes.example.com/1",
       "content": "A note"
     },
     "context": {
       "type": "http://example.org/Project",
       "name": "Project XYZ"
     },
     "audience": {
       "type": "Group",
       "name": "Project XYZ Working Group"
     },
     "to": "http://john.example.org"
   },
   {
     "summary": "John liked Sally's note",
     "type": "Like",
     "id": "http://activities.example.com/1",
     "actor": "http://john.example.org",
     "object": "http://notes.example.com/1",
     "context": {
       "type": "http://example.org/Project",
       "name": "Project XYZ"
     },
     "audience": {
       "type": "Group",
       "name": "Project XYZ Working Group"
     },
     "to": "http://sally.example.org"
   }
 ]
}`

var example144Unknown = func(m map[string]interface{}) map[string]interface{} {
	items := m["items"].([]interface{})
	create := items[0].(map[string]interface{})
	like := items[1].(map[string]interface{})
	create["context"] = map[string]interface{}{
		"type": "http://example.org/Project",
		"name": "Project XYZ",
	}
	like["context"] = map[string]interface{}{
		"type": "http://example.org/Project",
		"name": "Project XYZ",
	}
	m["items"] = []interface{}{create, like}
	return m
}

func example144Type() vocab.CollectionInterface {
	example144Type := NewActivityStreamsCollection()
	sally := MustParseURL("http://sally.example.org")
	john := MustParseURL("http://john.example.org")
	o := MustParseURL("http://notes.example.com/1")
	audience := NewActivityStreamsGroup()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Project XYZ Working Group")
	audience.SetName(name)
	note := NewActivityStreamsNote()
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("A note")
	note.SetSummary(summaryNote)
	noteId := NewActivityStreamsIdProperty()
	noteId.Set(MustParseURL("http://notes.example.com/1"))
	note.SetId(noteId)
	content := NewActivityStreamsContentProperty()
	content.AppendString("A note")
	note.SetContent(content)
	create := NewActivityStreamsCreate()
	summaryCreate := NewActivityStreamsSummaryProperty()
	summaryCreate.AppendString("Sally created a note")
	create.SetSummary(summaryCreate)
	createId := NewActivityStreamsIdProperty()
	createId.Set(MustParseURL("http://activities.example.com/1"))
	create.SetId(createId)
	createActor := NewActivityStreamsActorProperty()
	createActor.AppendIRI(sally)
	create.SetActor(createActor)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendNote(note)
	create.SetObject(obj)
	aud1 := NewActivityStreamsAudienceProperty()
	aud1.AppendGroup(audience)
	create.SetAudience(aud1)
	toCreate := NewActivityStreamsToProperty()
	toCreate.AppendIRI(john)
	create.SetTo(toCreate)
	like := NewActivityStreamsLike()
	summaryLike := NewActivityStreamsSummaryProperty()
	summaryLike.AppendString("John liked Sally's note")
	like.SetSummary(summaryLike)
	likeId := NewActivityStreamsIdProperty()
	likeId.Set(MustParseURL("http://activities.example.com/1"))
	like.SetId(likeId)
	likeActor := NewActivityStreamsActorProperty()
	likeActor.AppendIRI(john)
	like.SetActor(likeActor)
	objLike := NewActivityStreamsObjectProperty()
	objLike.AppendIRI(o)
	like.SetObject(objLike)
	aud2 := NewActivityStreamsAudienceProperty()
	aud2.AppendGroup(audience)
	like.SetAudience(aud2)
	toLike := NewActivityStreamsToProperty()
	toLike.AppendIRI(sally)
	like.SetTo(toLike)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Activities in Project XYZ")
	example144Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendCreate(create)
	items.AppendLike(like)
	example144Type.SetItems(items)
	return example144Type
}

const example145 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's friends list",
  "type": "Collection",
  "items": [
    {
      "summary": "Sally is influenced by Joe",
      "type": "Relationship",
      "subject": {
        "type": "Person",
        "name": "Sally"
      },
      "relationship": "http://purl.org/vocab/relationship/influencedBy",
      "object": {
        "type": "Person",
        "name": "Joe"
      }
    },
    {
      "summary": "Sally is a friend of Jane",
      "type": "Relationship",
      "subject": {
        "type": "Person",
        "name": "Sally"
      },
      "relationship": "http://purl.org/vocab/relationship/friendOf",
      "object": {
        "type": "Person",
        "name": "Jane"
      }
    }
  ]
}`

func example145Type() vocab.CollectionInterface {
	example145Type := NewActivityStreamsCollection()
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	influenced := MustParseURL("http://purl.org/vocab/relationship/influencedBy")
	sally := NewActivityStreamsPerson()
	sallyName := NewActivityStreamsNameProperty()
	sallyName.AppendString("Sally")
	sally.SetName(sallyName)
	jane := NewActivityStreamsPerson()
	name := NewActivityStreamsNameProperty()
	name.AppendString("Jane")
	jane.SetName(name)
	joe := NewActivityStreamsPerson()
	joeName := NewActivityStreamsNameProperty()
	joeName.AppendString("Joe")
	joe.SetName(joeName)
	joeRel := NewActivityStreamsRelationship()
	summaryJoe := NewActivityStreamsSummaryProperty()
	summaryJoe.AppendString("Sally is influenced by Joe")
	joeRel.SetSummary(summaryJoe)
	subj1 := NewActivityStreamsSubjectProperty()
	subj1.SetPerson(sally)
	joeRel.SetSubject(subj1)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendPerson(joe)
	joeRel.SetObject(obj)
	relationship := NewActivityStreamsRelationshipProperty()
	relationship.AppendIRI(influenced)
	joeRel.SetRelationship(relationship)
	janeRel := NewActivityStreamsRelationship()
	summaryJane := NewActivityStreamsSummaryProperty()
	summaryJane.AppendString("Sally is a friend of Jane")
	janeRel.SetSummary(summaryJane)
	subj2 := NewActivityStreamsSubjectProperty()
	subj2.SetPerson(sally)
	janeRel.SetSubject(subj2)
	objJane := NewActivityStreamsObjectProperty()
	objJane.AppendPerson(jane)
	janeRel.SetObject(objJane)
	relationshipJane := NewActivityStreamsRelationshipProperty()
	relationshipJane.AppendIRI(friend)
	janeRel.SetRelationship(relationshipJane)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally's friends list")
	example145Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendRelationship(joeRel)
	items.AppendRelationship(janeRel)
	example145Type.SetItems(items)
	return example145Type
}

// NOTE: Added `Z` to `startTime` to make align to spec!
const example146 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally became a friend of Matt",
  "type": "Create",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Relationship",
    "subject": "http://sally.example.org",
    "relationship": "http://purl.org/vocab/relationship/friendOf",
    "object": "http://matt.example.org",
    "startTime": "2015-04-21T12:34:56Z"
  }
}`

func example146Type() vocab.CreateInterface {
	example146Type := NewActivityStreamsCreate()
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	m := MustParseURL("http://matt.example.org")
	s := MustParseURL("http://sally.example.org")
	t, err := time.Parse(time.RFC3339, "2015-04-21T12:34:56Z")
	if err != nil {
		panic(err)
	}
	relationship := NewActivityStreamsRelationship()
	subj := NewActivityStreamsSubjectProperty()
	subj.SetIRI(s)
	relationship.SetSubject(subj)
	friendRel := NewActivityStreamsRelationshipProperty()
	friendRel.AppendIRI(friend)
	relationship.SetRelationship(friendRel)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(m)
	relationship.SetObject(obj)
	startTime := NewActivityStreamsStartTimeProperty()
	startTime.Set(t)
	relationship.SetStartTime(startTime)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally became a friend of Matt")
	example146Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(s)
	example146Type.SetActor(objectActor)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendRelationship(relationship)
	example146Type.SetObject(objRoot)
	return example146Type
}

const example147 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "http://example.org/connection-requests/123",
  "summary": "Sally requested to be a friend of John",
  "type": "Offer",
  "actor": "acct:sally@example.org",
  "object": {
    "summary": "Sally and John's friendship",
    "id": "http://example.org/connections/123",
    "type": "Relationship",
    "subject": "acct:sally@example.org",
    "relationship": "http://purl.org/vocab/relationship/friendOf",
    "object": "acct:john@example.org"
  },
  "target": "acct:john@example.org"
}`

func example147Type() vocab.OfferInterface {
	example147Type := NewActivityStreamsOffer()
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	s := MustParseURL("acct:sally@example.org")
	t := MustParseURL("acct:john@example.org")
	rel := NewActivityStreamsRelationship()
	summaryRel := NewActivityStreamsSummaryProperty()
	summaryRel.AppendString("Sally and John's friendship")
	rel.SetSummary(summaryRel)
	relId := NewActivityStreamsIdProperty()
	relId.Set(MustParseURL("http://example.org/connections/123"))
	rel.SetId(relId)
	subj := NewActivityStreamsSubjectProperty()
	subj.SetIRI(s)
	rel.SetSubject(subj)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(t)
	rel.SetObject(obj)
	relationship := NewActivityStreamsRelationshipProperty()
	relationship.AppendIRI(friend)
	rel.SetRelationship(relationship)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally requested to be a friend of John")
	example147Type.SetSummary(summary)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://example.org/connection-requests/123"))
	example147Type.SetId(id)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(s)
	example147Type.SetActor(objectActor)
	objRel := NewActivityStreamsObjectProperty()
	objRel.AppendRelationship(rel)
	example147Type.SetObject(objRel)
	objRoot := NewActivityStreamsObjectProperty()
	objRoot.AppendRelationship(rel)
	example147Type.SetObject(objRoot)
	target := NewActivityStreamsTargetProperty()
	target.AppendIRI(t)
	example147Type.SetTarget(target)
	return example147Type
}

const example148 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally and John's relationship history",
  "type": "Collection",
  "items": [
    {
      "summary": "John accepted Sally's friend request",
      "id": "http://example.org/activities/122",
      "type": "Accept",
      "actor": "acct:john@example.org",
      "object": "http://example.org/connection-requests/123",
      "inReplyTo": "http://example.org/connection-requests/123",
      "context": "http://example.org/connections/123",
      "result": [
        "http://example.org/activities/123",
        "http://example.org/activities/124",
        "http://example.org/activities/125",
        "http://example.org/activities/126"
      ]
    },
    {
      "summary": "John followed Sally",
      "id": "http://example.org/activities/123",
      "type": "Follow",
      "actor": "acct:john@example.org",
      "object": "acct:sally@example.org",
      "context": "http://example.org/connections/123"
    },
    {
      "summary": "Sally followed John",
      "id": "http://example.org/activities/124",
      "type": "Follow",
      "actor": "acct:sally@example.org",
      "object": "acct:john@example.org",
      "context": "http://example.org/connections/123"
    },
    {
      "summary": "John added Sally to his friends list",
      "id": "http://example.org/activities/125",
      "type": "Add",
      "actor": "acct:john@example.org",
      "object": "http://example.org/connections/123",
      "target": {
        "type": "Collection",
        "summary": "John's Connections"
      },
      "context": "http://example.org/connections/123"
    },
    {
      "summary": "Sally added John to her friends list",
      "id": "http://example.org/activities/126",
      "type": "Add",
      "actor": "acct:sally@example.org",
      "object": "http://example.org/connections/123",
      "target": {
        "type": "Collection",
        "summary": "Sally's Connections"
      },
      "context": "http://example.org/connections/123"
    }
  ]
}`

func example148Type() vocab.CollectionInterface {
	example148Type := NewActivityStreamsCollection()
	john := MustParseURL("acct:john@example.org")
	sally := MustParseURL("acct:sally@example.org")
	req123 := MustParseURL("http://example.org/connection-requests/123")
	conn123 := MustParseURL("http://example.org/connections/123")
	a123 := MustParseURL("http://example.org/activities/123")
	a124 := MustParseURL("http://example.org/activities/124")
	a125 := MustParseURL("http://example.org/activities/125")
	a126 := MustParseURL("http://example.org/activities/126")
	jc := NewActivityStreamsCollection()
	summaryJc := NewActivityStreamsSummaryProperty()
	summaryJc.AppendString("John's Connections")
	jc.SetSummary(summaryJc)
	sc := NewActivityStreamsCollection()
	summarySc := NewActivityStreamsSummaryProperty()
	summarySc.AppendString("Sally's Connections")
	sc.SetSummary(summarySc)
	o1 := NewActivityStreamsAccept()
	oId1 := NewActivityStreamsIdProperty()
	oId1.Set(MustParseURL("http://example.org/activities/122"))
	o1.SetId(oId1)
	summary1 := NewActivityStreamsSummaryProperty()
	summary1.AppendString("John accepted Sally's friend request")
	o1.SetSummary(summary1)
	obj1 := NewActivityStreamsObjectProperty()
	obj1.AppendIRI(req123)
	o1.SetObject(obj1)
	inReplyTo := NewActivityStreamsInReplyToProperty()
	inReplyTo.AppendIRI(req123)
	o1.SetInReplyTo(inReplyTo)
	ctx1 := NewActivityStreamsContextProperty()
	ctx1.AppendIRI(conn123)
	o1.SetContext(ctx1)
	result := NewActivityStreamsResultProperty()
	result.AppendIRI(a123)
	result.AppendIRI(a124)
	result.AppendIRI(a125)
	result.AppendIRI(a126)
	o1.SetResult(result)
	objectActor1 := NewActivityStreamsActorProperty()
	objectActor1.AppendIRI(john)
	o1.SetActor(objectActor1)
	o2 := NewActivityStreamsFollow()
	oId2 := NewActivityStreamsIdProperty()
	oId2.Set(MustParseURL("http://example.org/activities/123"))
	o2.SetId(oId2)
	objectActor2 := NewActivityStreamsActorProperty()
	objectActor2.AppendIRI(john)
	o2.SetActor(objectActor2)
	obj2 := NewActivityStreamsObjectProperty()
	obj2.AppendIRI(sally)
	o2.SetObject(obj2)
	ctx2 := NewActivityStreamsContextProperty()
	ctx2.AppendIRI(conn123)
	o2.SetContext(ctx2)
	summary2 := NewActivityStreamsSummaryProperty()
	summary2.AppendString("John followed Sally")
	o2.SetSummary(summary2)
	o3 := NewActivityStreamsFollow()
	oId3 := NewActivityStreamsIdProperty()
	oId3.Set(MustParseURL("http://example.org/activities/124"))
	o3.SetId(oId3)
	objectActor3 := NewActivityStreamsActorProperty()
	objectActor3.AppendIRI(sally)
	o3.SetActor(objectActor3)
	obj3 := NewActivityStreamsObjectProperty()
	obj3.AppendIRI(john)
	o3.SetObject(obj3)
	ctx3 := NewActivityStreamsContextProperty()
	ctx3.AppendIRI(conn123)
	o3.SetContext(ctx3)
	summary3 := NewActivityStreamsSummaryProperty()
	summary3.AppendString("Sally followed John")
	o3.SetSummary(summary3)
	o4 := NewActivityStreamsAdd()
	oId4 := NewActivityStreamsIdProperty()
	oId4.Set(MustParseURL("http://example.org/activities/125"))
	o4.SetId(oId4)
	summary4 := NewActivityStreamsSummaryProperty()
	summary4.AppendString("John added Sally to his friends list")
	o4.SetSummary(summary4)
	objectActor4 := NewActivityStreamsActorProperty()
	objectActor4.AppendIRI(john)
	o4.SetActor(objectActor4)
	obj4 := NewActivityStreamsObjectProperty()
	obj4.AppendIRI(conn123)
	o4.SetObject(obj4)
	ctx4 := NewActivityStreamsContextProperty()
	ctx4.AppendIRI(conn123)
	o4.SetContext(ctx4)
	tobj4 := NewActivityStreamsTargetProperty()
	tobj4.AppendCollection(jc)
	o4.SetTarget(tobj4)
	o5 := NewActivityStreamsAdd()
	oId5 := NewActivityStreamsIdProperty()
	oId5.Set(MustParseURL("http://example.org/activities/126"))
	o5.SetId(oId5)
	summary5 := NewActivityStreamsSummaryProperty()
	summary5.AppendString("Sally added John to her friends list")
	o5.SetSummary(summary5)
	objectActor5 := NewActivityStreamsActorProperty()
	objectActor5.AppendIRI(sally)
	o5.SetActor(objectActor5)
	obj5 := NewActivityStreamsObjectProperty()
	obj5.AppendIRI(conn123)
	o5.SetObject(obj5)
	ctx5 := NewActivityStreamsContextProperty()
	ctx5.AppendIRI(conn123)
	o5.SetContext(ctx5)
	tobj5 := NewActivityStreamsTargetProperty()
	tobj5.AppendCollection(sc)
	o5.SetTarget(tobj5)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally and John's relationship history")
	example148Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendAccept(o1)
	items.AppendFollow(o2)
	items.AppendFollow(o3)
	items.AppendAdd(o4)
	items.AppendAdd(o5)
	example148Type.SetItems(items)
	return example148Type
}

const example149 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "San Francisco, CA"
}`

func example149Type() vocab.PlaceInterface {
	example149Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("San Francisco, CA")
	example149Type.SetName(name)
	return example149Type
}

// NOTE: Un-stringified the longitude and latitude values.
const example150 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "San Francisco, CA",
  "longitude": 122.4167,
  "latitude": 37.7833
}`

func example150Type() vocab.PlaceInterface {
	example150Type := NewActivityStreamsPlace()
	name := NewActivityStreamsNameProperty()
	name.AppendString("San Francisco, CA")
	example150Type.SetName(name)
	lon := NewActivityStreamsLongitudeProperty()
	lon.Set(122.4167)
	example150Type.SetLongitude(lon)
	lat := NewActivityStreamsLatitudeProperty()
	lat.Set(37.7833)
	example150Type.SetLatitude(lat)
	return example150Type
}

const example151 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A question about robots",
  "id": "http://help.example.org/question/1",
  "type": "Question",
  "content": "I'd like to build a robot to feed my cat. Should I use Arduino or Raspberry Pi?"
}`

func example151Type() vocab.QuestionInterface {
	example151Type := NewActivityStreamsQuestion()
	name := NewActivityStreamsNameProperty()
	name.AppendString("A question about robots")
	example151Type.SetName(name)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://help.example.org/question/1"))
	example151Type.SetId(id)
	content := NewActivityStreamsContentProperty()
	content.AppendString("I'd like to build a robot to feed my cat. Should I use Arduino or Raspberry Pi?")
	example151Type.SetContent(content)
	return example151Type
}

const example152 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "http://polls.example.org/question/1",
  "name": "A question about robots",
  "type": "Question",
  "content": "I'd like to build a robot to feed my cat. Which platform is best?",
  "oneOf": [
    {"name": "arduino"},
    {"name": "raspberry pi"}
  ]
}`

var example152Unknown = func(m map[string]interface{}) map[string]interface{} {
	ard := make(map[string]interface{})
	ard["name"] = "arduino"
	ras := make(map[string]interface{})
	ras["name"] = "raspberry pi"
	m["oneOf"] = []interface{}{ard, ras}
	return m
}

func example152Type() vocab.QuestionInterface {
	example152Type := NewActivityStreamsQuestion()
	name := NewActivityStreamsNameProperty()
	name.AppendString("A question about robots")
	example152Type.SetName(name)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://polls.example.org/question/1"))
	example152Type.SetId(id)
	content := NewActivityStreamsContentProperty()
	content.AppendString("I'd like to build a robot to feed my cat. Which platform is best?")
	example152Type.SetContent(content)
	return example152Type
}

const example153 = `{
 "@context": "https://www.w3.org/ns/activitystreams",
 "attributedTo": "http://sally.example.org",
 "inReplyTo": "http://polls.example.org/question/1",
 "name": "arduino"
}`

func example153Type() interface{} {
	return nil
}

var example153Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["attributedTo"] = "http://sally.example.org"
	m["inReplyTo"] = "http://polls.example.org/question/1"
	m["name"] = "arduino"
	return m
}

const example154 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A question about robots",
  "id": "http://polls.example.org/question/1",
  "type": "Question",
  "content": "I'd like to build a robot to feed my cat. Which platform is best?",
  "oneOf": [
    {"name": "arduino"},
    {"name": "raspberry pi"}
  ],
  "replies": {
    "type": "Collection",
    "totalItems": 3,
    "items": [
      {
        "attributedTo": "http://sally.example.org",
        "inReplyTo": "http://polls.example.org/question/1",
        "name": "arduino"
      },
      {
        "attributedTo": "http://joe.example.org",
        "inReplyTo": "http://polls.example.org/question/1",
        "name": "arduino"
      },
      {
        "attributedTo": "http://john.example.org",
        "inReplyTo": "http://polls.example.org/question/1",
        "name": "raspberry pi"
      }
    ]
  },
  "result": {
    "type": "Note",
    "content": "Users are favoriting &quot;arduino&quot; by a 33% margin."
  }
}`

var example154Unknown = func(m map[string]interface{}) map[string]interface{} {
	ard := make(map[string]interface{})
	ard["name"] = "arduino"
	ras := make(map[string]interface{})
	ras["name"] = "raspberry pi"
	m["oneOf"] = []interface{}{ard, ras}
	// replies
	one := make(map[string]interface{})
	one["attributedTo"] = "http://sally.example.org"
	one["inReplyTo"] = "http://polls.example.org/question/1"
	one["name"] = "arduino"
	two := make(map[string]interface{})
	two["attributedTo"] = "http://joe.example.org"
	two["inReplyTo"] = "http://polls.example.org/question/1"
	two["name"] = "arduino"
	three := make(map[string]interface{})
	three["attributedTo"] = "http://john.example.org"
	three["inReplyTo"] = "http://polls.example.org/question/1"
	three["name"] = "raspberry pi"
	items := []interface{}{one, two, three}
	replies := m["replies"].(map[string]interface{})
	replies["items"] = items
	return m
}

func example154Type() vocab.QuestionInterface {
	example154Type := NewActivityStreamsQuestion()
	replies := NewActivityStreamsCollection()
	totalItems := NewActivityStreamsTotalItemsProperty()
	totalItems.Set(3)
	replies.SetTotalItems(totalItems)
	note := NewActivityStreamsNote()
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("Users are favoriting &quot;arduino&quot; by a 33% margin.")
	note.SetContent(contentNote)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A question about robots")
	example154Type.SetName(name)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://polls.example.org/question/1"))
	example154Type.SetId(id)
	content := NewActivityStreamsContentProperty()
	content.AppendString("I'd like to build a robot to feed my cat. Which platform is best?")
	example154Type.SetContent(content)
	reply := NewActivityStreamsRepliesProperty()
	reply.SetCollection(replies)
	example154Type.SetReplies(reply)
	result := NewActivityStreamsResultProperty()
	result.AppendNote(note)
	example154Type.SetResult(result)
	return example154Type
}

const example155 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "History of John's note",
  "type": "Collection",
  "items": [
    {
      "summary": "Sally liked John's note",
      "type": "Like",
      "actor": "http://sally.example.org",
      "id": "http://activities.example.com/1",
      "published": "2015-11-12T12:34:56Z",
      "object": {
        "summary": "John's note",
        "type": "Note",
        "id": "http://notes.example.com/1",
        "attributedTo": "http://john.example.org",
        "content": "My note"
      }
    },
    {
      "summary": "Sally disliked John's note",
      "type": "Dislike",
      "actor": "http://sally.example.org",
      "id": "http://activities.example.com/2",
      "published": "2015-12-11T21:43:56Z",
      "object": {
        "summary": "John's note",
        "type": "Note",
        "id": "http://notes.example.com/1",
        "attributedTo": "http://john.example.org",
        "content": "My note"
      }
    }
  ]
}`

func example155Type() vocab.CollectionInterface {
	example155Type := NewActivityStreamsCollection()
	john := MustParseURL("http://john.example.org")
	sally := MustParseURL("http://sally.example.org")
	t1, err := time.Parse(time.RFC3339, "2015-11-12T12:34:56Z")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-12-11T21:43:56Z")
	if err != nil {
		panic(err)
	}
	note := NewActivityStreamsNote()
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("John's note")
	note.SetSummary(summaryNote)
	noteId := NewActivityStreamsIdProperty()
	noteId.Set(MustParseURL("http://notes.example.com/1"))
	note.SetId(noteId)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("My note")
	note.SetContent(contentNote)
	attr := NewActivityStreamsAttributedToProperty()
	attr.AppendIRI(john)
	note.SetAttributedTo(attr)
	like := NewActivityStreamsLike()
	summaryLike := NewActivityStreamsSummaryProperty()
	summaryLike.AppendString("Sally liked John's note")
	like.SetSummary(summaryLike)
	likeId := NewActivityStreamsIdProperty()
	likeId.Set(MustParseURL("http://activities.example.com/1"))
	like.SetId(likeId)
	likeActor := NewActivityStreamsActorProperty()
	likeActor.AppendIRI(sally)
	like.SetActor(likeActor)
	published1 := NewActivityStreamsPublishedProperty()
	published1.Set(t1)
	like.SetPublished(published1)
	objLike := NewActivityStreamsObjectProperty()
	objLike.AppendNote(note)
	like.SetObject(objLike)
	dislike := NewActivityStreamsDislike()
	summaryDislike := NewActivityStreamsSummaryProperty()
	summaryDislike.AppendString("Sally disliked John's note")
	dislike.SetSummary(summaryDislike)
	dislikeId := NewActivityStreamsIdProperty()
	dislikeId.Set(MustParseURL("http://activities.example.com/2"))
	dislike.SetId(dislikeId)
	dislikeActor := NewActivityStreamsActorProperty()
	dislikeActor.AppendIRI(sally)
	dislike.SetActor(dislikeActor)
	published2 := NewActivityStreamsPublishedProperty()
	published2.Set(t2)
	dislike.SetPublished(published2)
	objDislike := NewActivityStreamsObjectProperty()
	objDislike.AppendNote(note)
	dislike.SetObject(objDislike)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("History of John's note")
	example155Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendLike(like)
	items.AppendDislike(dislike)
	example155Type.SetItems(items)
	return example155Type
}

const example156 = `{
 "@context": "https://www.w3.org/ns/activitystreams",
 "summary": "History of John's note",
 "type": "Collection",
 "items": [
   {
     "summary": "Sally liked John's note",
     "type": "Like",
     "id": "http://activities.example.com/1",
     "actor": "http://sally.example.org",
     "published": "2015-11-12T12:34:56Z",
     "object": {
       "summary": "John's note",
       "type": "Note",
       "id": "http://notes.example.com/1",
       "attributedTo": "http://john.example.org",
       "content": "My note"
     }
   },
   {
     "summary": "Sally no longer likes John's note",
     "type": "Undo",
     "id": "http://activities.example.com/2",
     "actor": "http://sally.example.org",
     "published": "2015-12-11T21:43:56Z",
     "object": "http://activities.example.com/1"
   }
 ]
}`

func example156Type() vocab.CollectionInterface {
	example156Type := NewActivityStreamsCollection()
	john := MustParseURL("http://john.example.org")
	sally := MustParseURL("http://sally.example.org")
	a := MustParseURL("http://activities.example.com/1")
	t1, err := time.Parse(time.RFC3339, "2015-11-12T12:34:56Z")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-12-11T21:43:56Z")
	if err != nil {
		panic(err)
	}
	note := NewActivityStreamsNote()
	noteId := NewActivityStreamsIdProperty()
	noteId.Set(MustParseURL("http://notes.example.com/1"))
	note.SetId(noteId)
	summaryNote := NewActivityStreamsSummaryProperty()
	summaryNote.AppendString("John's note")
	note.SetSummary(summaryNote)
	attr := NewActivityStreamsAttributedToProperty()
	attr.AppendIRI(john)
	note.SetAttributedTo(attr)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("My note")
	note.SetContent(contentNote)
	like := NewActivityStreamsLike()
	likeId := NewActivityStreamsIdProperty()
	likeId.Set(MustParseURL("http://activities.example.com/1"))
	like.SetId(likeId)
	summaryLike := NewActivityStreamsSummaryProperty()
	summaryLike.AppendString("Sally liked John's note")
	like.SetSummary(summaryLike)
	likeActor := NewActivityStreamsActorProperty()
	likeActor.AppendIRI(sally)
	like.SetActor(likeActor)
	published1 := NewActivityStreamsPublishedProperty()
	published1.Set(t1)
	like.SetPublished(published1)
	objLike := NewActivityStreamsObjectProperty()
	objLike.AppendNote(note)
	like.SetObject(objLike)
	undo := NewActivityStreamsUndo()
	undoId := NewActivityStreamsIdProperty()
	undoId.Set(MustParseURL("http://activities.example.com/2"))
	undo.SetId(undoId)
	summaryUndo := NewActivityStreamsSummaryProperty()
	summaryUndo.AppendString("Sally no longer likes John's note")
	undo.SetSummary(summaryUndo)
	undoActor := NewActivityStreamsActorProperty()
	undoActor.AppendIRI(sally)
	undo.SetActor(undoActor)
	published2 := NewActivityStreamsPublishedProperty()
	published2.Set(t2)
	undo.SetPublished(published2)
	obj := NewActivityStreamsObjectProperty()
	obj.AppendIRI(a)
	undo.SetObject(obj)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("History of John's note")
	example156Type.SetSummary(summary)
	items := NewActivityStreamsItemsProperty()
	items.AppendLike(like)
	items.AppendUndo(undo)
	example156Type.SetItems(items)
	return example156Type
}

// NOTE: The `content` field has been inlined to keep within JSON spec.
const example157 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A thank-you note",
  "type": "Note",
  "content": "Thank you <a href='http://sally.example.org'>@sally</a> for all your hard work! <a href='http://example.org/tags/givingthanks'>#givingthanks</a>",
  "to": {
    "name": "Sally",
    "type": "Person",
    "id": "http://sally.example.org"
  },
  "tag": {
    "id": "http://example.org/tags/givingthanks",
    "name": "#givingthanks"
  }
}`

var example157Unknown = func(m map[string]interface{}) map[string]interface{} {
	m["tag"] = map[string]interface{}{
		"id":   "http://example.org/tags/givingthanks",
		"name": "#givingthanks",
	}
	return m
}

func example157Type() vocab.NoteInterface {
	example157Type := NewActivityStreamsNote()
	person := NewActivityStreamsPerson()
	sally := NewActivityStreamsNameProperty()
	sally.AppendString("Sally")
	person.SetName(sally)
	id := NewActivityStreamsIdProperty()
	id.Set(MustParseURL("http://sally.example.org"))
	person.SetId(id)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A thank-you note")
	example157Type.SetName(name)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("Thank you <a href='http://sally.example.org'>@sally</a> for all your hard work! <a href='http://example.org/tags/givingthanks'>#givingthanks</a>")
	example157Type.SetContent(contentNote)
	to := NewActivityStreamsToProperty()
	to.AppendPerson(person)
	example157Type.SetTo(to)
	return example157Type
}

const example158 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A thank-you note",
  "type": "Note",
  "content": "Thank you @sally for all your hard work! #givingthanks",
  "tag": [
    {
      "type": "Mention",
      "href": "http://example.org/people/sally",
      "name": "@sally"
    },
    {
      "id": "http://example.org/tags/givingthanks",
      "name": "#givingthanks"
    }
  ]
}`

var example158Unknown = func(m map[string]interface{}) map[string]interface{} {
	existing := m["tag"]
	next := map[string]interface{}{
		"id":   "http://example.org/tags/givingthanks",
		"name": "#givingthanks",
	}
	m["tag"] = []interface{}{existing, next}
	return m
}

func example158Type() vocab.NoteInterface {
	example158Type := NewActivityStreamsNote()
	u := MustParseURL("http://example.org/people/sally")
	mention := NewActivityStreamsMention()
	hrefLink := NewActivityStreamsHrefProperty()
	hrefLink.Set(u)
	mention.SetHref(hrefLink)
	mentionName := NewActivityStreamsNameProperty()
	mentionName.AppendString("@sally")
	mention.SetName(mentionName)
	name := NewActivityStreamsNameProperty()
	name.AppendString("A thank-you note")
	example158Type.SetName(name)
	contentNote := NewActivityStreamsContentProperty()
	contentNote.AppendString("Thank you @sally for all your hard work! #givingthanks")
	example158Type.SetContent(contentNote)
	tag := NewActivityStreamsTagProperty()
	tag.AppendMention(mention)
	example158Type.SetTag(tag)
	return example158Type
}

const example159 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally moved the sales figures from Folder A to Folder B",
  "type": "Move",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Document",
    "name": "sales figures"
  },
  "origin": {
    "type": "Collection",
    "name": "Folder A"
  },
  "target": {
    "type": "Collection",
    "name": "Folder B"
  }
}`

func example159Type() vocab.MoveInterface {
	example159Type := NewActivityStreamsMove()
	sally := MustParseURL("http://sally.example.org")
	obj := NewActivityStreamsDocument()
	nameObj := NewActivityStreamsNameProperty()
	nameObj.AppendString("sales figures")
	obj.SetName(nameObj)
	origin := NewActivityStreamsCollection()
	nameOrigin := NewActivityStreamsNameProperty()
	nameOrigin.AppendString("Folder A")
	origin.SetName(nameOrigin)
	target := NewActivityStreamsCollection()
	nameTarget := NewActivityStreamsNameProperty()
	nameTarget.AppendString("Folder B")
	target.SetName(nameTarget)
	summary := NewActivityStreamsSummaryProperty()
	summary.AppendString("Sally moved the sales figures from Folder A to Folder B")
	example159Type.SetSummary(summary)
	objectActor := NewActivityStreamsActorProperty()
	objectActor.AppendIRI(sally)
	example159Type.SetActor(objectActor)
	object := NewActivityStreamsObjectProperty()
	object.AppendDocument(obj)
	example159Type.SetObject(object)
	originProp := NewActivityStreamsOriginProperty()
	originProp.AppendCollection(origin)
	example159Type.SetOrigin(originProp)
	tobj := NewActivityStreamsTargetProperty()
	tobj.AppendCollection(target)
	example159Type.SetTarget(tobj)
	return example159Type
}
