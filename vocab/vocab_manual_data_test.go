package vocab

import (
	"time"
)

const example1 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Object",
  "id": "http://www.test.example/object/1",
  "name": "A Simple, non-specific object"
}`

var example1Type = &Object{}

func init() {
	example1Type.AppendType("Object")
	example1Type.SetId(MustParseURL("http://www.test.example/object/1"))
	example1Type.AppendNameString("A Simple, non-specific object")
}

const example2 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link"
}`

var example2Type = &Link{}

func init() {
	href := MustParseURL("http://example.org/abc")
	example2Type.AppendType("Link")
	example2Type.SetHref(href)
	example2Type.SetHreflang("en")
	example2Type.SetMediaType("text/html")
	example2Type.AppendNameString("An example link")
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

var example3Type = &Activity{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Note{}
	object.AppendType("Note")
	object.AppendNameString("A Note")
	example3Type.AppendType("Activity")
	example3Type.AppendSummaryString("Sally did something to a note")
	example3Type.AppendActorObject(actor)
	example3Type.AppendObject(object)
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

var example4Type = &Travel{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	place := &Place{}
	place.AppendType("Place")
	place.AppendNameString("Work")
	example4Type.AppendType("Travel")
	example4Type.AppendSummaryString("Sally went to work")
	example4Type.AppendActorObject(person)
	example4Type.AppendTargetObject(place)
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

var example5Type = &Collection{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("A Simple Note")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Another Simple Note")
	example5Type.AppendType("Collection")
	example5Type.AppendSummaryString("Sally's notes")
	example5Type.SetTotalItems(2)
	example5Type.AppendItemsObject(note1)
	example5Type.AppendItemsObject(note2)
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

var example6Type = &OrderedCollection{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("A Simple Note")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Another Simple Note")
	example6Type.AppendType("OrderedCollection")
	example6Type.AppendSummaryString("Sally's notes")
	example6Type.SetTotalItems(2)
	example6Type.AppendOrderedItemsObject(note1)
	example6Type.AppendOrderedItemsObject(note2)
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

var example7Type = &CollectionPage{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("A Simple Note")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Another Simple Note")
	link := MustParseURL("http://example.org/foo")
	example7Type.AppendType("CollectionPage")
	example7Type.AppendSummaryString("Page 1 of Sally's notes")
	example7Type.SetId(MustParseURL("http://example.org/foo?page=1"))
	example7Type.SetPartOfIRI(link)
	example7Type.AppendItemsObject(note1)
	example7Type.AppendItemsObject(note2)
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

var example8Type = &OrderedCollectionPage{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("A Simple Note")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Another Simple Note")
	link := MustParseURL("http://example.org/foo")
	example8Type.AppendType("OrderedCollectionPage")
	example8Type.AppendSummaryString("Page 1 of Sally's notes")
	example8Type.SetId(MustParseURL("http://example.org/foo?page=1"))
	example8Type.SetPartOfIRI(link)
	example8Type.AppendOrderedItemsObject(note1)
	example8Type.AppendOrderedItemsObject(note2)
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

var example9Type = &Accept{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	event := &Event{}
	event.AppendType("Event")
	event.AppendNameString("Going-Away Party for Jim")
	link := MustParseURL("http://john.example.org")
	invite := &Invite{}
	invite.AppendType("Invite")
	invite.AppendActorIRI(link)
	invite.AppendObject(event)
	example9Type.AppendType("Accept")
	example9Type.AppendSummaryString("Sally accepted an invitation to a party")
	example9Type.AppendActorObject(person)
	example9Type.AppendObject(invite)
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

var example10Type = &Accept{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("Joe")
	target := &Group{}
	target.AppendType("Group")
	target.AppendNameString("The Club")
	example10Type.AppendType("Accept")
	example10Type.AppendSummaryString("Sally accepted Joe into the club")
	example10Type.AppendActorObject(person)
	example10Type.AppendObject(object)
	example10Type.AppendTargetObject(target)
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

var example11Type = &TentativeAccept{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	event := &Event{}
	event.AppendType("Event")
	event.AppendNameString("Going-Away Party for Jim")
	link := MustParseURL("http://john.example.org")
	invite := &Invite{}
	invite.AppendType("Invite")
	invite.AppendActorIRI(link)
	invite.AppendObject(event)
	example11Type.AppendType("TentativeAccept")
	example11Type.AppendSummaryString("Sally tentatively accepted an invitation to a party")
	example11Type.AppendActorObject(person)
	example11Type.AppendObject(invite)
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

var example12Type = &Add{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	link := MustParseURL("http://example.org/abc")
	example12Type.AppendType("Add")
	example12Type.AppendSummaryString("Sally added an object")
	example12Type.AppendActorObject(person)
	example12Type.AppendObjectIRI(link)
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

var example13Type = &Add{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	link := MustParseURL("http://example.org/img/cat.png")
	object := &Image{}
	object.AppendType("Image")
	object.AppendNameString("A picture of my cat")
	object.AppendUrlAnyURI(link)
	origin := &Collection{}
	origin.AppendType("Collection")
	origin.AppendNameString("Camera Roll")
	target := &Collection{}
	target.AppendType("Collection")
	target.AppendNameString("My Cat Pictures")
	example13Type.AppendType("Add")
	example13Type.AppendSummaryString("Sally added a picture of her cat to her cat picture collection")
	example13Type.AppendActorObject(person)
	example13Type.AppendObject(object)
	example13Type.AppendOriginObject(origin)
	example13Type.AppendTargetObject(target)
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

var example14Type = &Arrive{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	location := &Place{}
	location.AppendType("Place")
	location.AppendNameString("Work")
	origin := &Place{}
	origin.AppendType("Place")
	origin.AppendNameString("Home")
	example14Type.AppendType("Arrive")
	example14Type.AppendSummaryString("Sally arrived at work")
	example14Type.AppendActorObject(person)
	example14Type.AppendLocationObject(location)
	example14Type.AppendOriginObject(origin)
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

var example15Type = &Create{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Note{}
	object.AppendType("Note")
	object.AppendNameString("A Simple Note")
	object.AppendContentString("This is a simple note")
	example15Type.AppendType("Create")
	example15Type.AppendSummaryString("Sally created a note")
	example15Type.AppendActorObject(actor)
	example15Type.AppendObject(object)
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

var example16Type = &Delete{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	origin := &Collection{}
	origin.AppendType("Collection")
	origin.AppendNameString("Sally's Notes")
	link := MustParseURL("http://example.org/notes/1")
	example16Type.AppendType("Delete")
	example16Type.AppendSummaryString("Sally deleted a note")
	example16Type.AppendActorObject(actor)
	example16Type.AppendObjectIRI(link)
	example16Type.AppendOriginObject(origin)
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

var example17Type = &Follow{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("John")
	example17Type.AppendType("Follow")
	example17Type.AppendSummaryString("Sally followed John")
	example17Type.AppendActorObject(actor)
	example17Type.AppendObject(object)
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

var example18Type = &Ignore{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example18Type.AppendType("Ignore")
	example18Type.AppendSummaryString("Sally ignored a note")
	example18Type.AppendActorObject(actor)
	example18Type.AppendObjectIRI(link)
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

var example19Type = &Join{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Group{}
	object.AppendType("Group")
	object.AppendNameString("A Simple Group")
	example19Type.AppendType("Join")
	example19Type.AppendSummaryString("Sally joined a group")
	example19Type.AppendActorObject(actor)
	example19Type.AppendObject(object)

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

var example20Type = &Leave{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Place{}
	object.AppendType("Place")
	object.AppendNameString("Work")
	example20Type.AppendType("Leave")
	example20Type.AppendSummaryString("Sally left work")
	example20Type.AppendActorObject(actor)
	example20Type.AppendObject(object)
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

var example21Type = &Leave{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Group{}
	object.AppendType("Group")
	object.AppendNameString("A Simple Group")
	example21Type.AppendType("Leave")
	example21Type.AppendSummaryString("Sally left a group")
	example21Type.AppendActorObject(actor)
	example21Type.AppendObject(object)
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

var example22Type = &Like{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example22Type.AppendType("Like")
	example22Type.AppendSummaryString("Sally liked a note")
	example22Type.AppendActorObject(actor)
	example22Type.AppendObjectIRI(link)
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

var example23Type = &Offer{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := make(map[string]interface{})
	object["type"] = "http://www.types.example/ProductOffer"
	object["name"] = "50% Off!"
	target := &Person{}
	target.AppendType("Person")
	target.AppendNameString("Lewis")
	example23Type.AppendType("Offer")
	example23Type.AppendSummaryString("Sally offered 50% off to Lewis")
	example23Type.AppendActorObject(actor)
	example23Type.SetUnknownObject(object)
	example23Type.AppendTargetObject(target)
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

var example24Type = &Invite{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Event{}
	object.AppendType("Event")
	object.AppendNameString("A Party")
	target1 := &Person{}
	target1.AppendType("Person")
	target1.AppendNameString("John")
	target2 := &Person{}
	target2.AppendType("Person")
	target2.AppendNameString("Lisa")
	example24Type.AppendType("Invite")
	example24Type.AppendSummaryString("Sally invited John and Lisa to a party")
	example24Type.AppendActorObject(actor)
	example24Type.AppendObject(object)
	example24Type.AppendTargetObject(target1)
	example24Type.AppendTargetObject(target2)
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

var example25Type = &Reject{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://john.example.org")
	inviteObject := &Event{}
	inviteObject.AppendType("Event")
	inviteObject.AppendNameString("Going-Away Party for Jim")
	object := &Invite{}
	object.AppendType("Invite")
	object.AppendActorIRI(link)
	object.AppendObject(inviteObject)
	example25Type.AppendType("Reject")
	example25Type.AppendSummaryString("Sally rejected an invitation to a party")
	example25Type.AppendActorObject(actor)
	example25Type.AppendObject(object)
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

var example26Type = &TentativeReject{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://john.example.org")
	inviteObject := &Event{}
	inviteObject.AppendType("Event")
	inviteObject.AppendNameString("Going-Away Party for Jim")
	object := &Invite{}
	object.AppendType("Invite")
	object.AppendActorIRI(link)
	object.AppendObject(inviteObject)
	example26Type.AppendType("TentativeReject")
	example26Type.AppendSummaryString("Sally tentatively rejected an invitation to a party")
	example26Type.AppendActorObject(actor)
	example26Type.AppendObject(object)
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

var example27Type = &Remove{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	target := &Collection{}
	target.AppendType("Collection")
	target.AppendNameString("Notes Folder")
	example27Type.AppendType("Remove")
	example27Type.AppendSummaryString("Sally removed a note from her notes folder")
	example27Type.AppendActorObject(actor)
	example27Type.AppendObjectIRI(link)
	example27Type.AppendTargetObject(target)
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

var example28Type = &Remove{}

func init() {
	actor := make(map[string]interface{})
	actor["type"] = "http://example.org/Role"
	actor["name"] = "The Moderator"
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("Sally")
	origin := &Group{}
	origin.AppendType("Group")
	origin.AppendNameString("A Simple Group")
	example28Type.AppendType("Remove")
	example28Type.AppendSummaryString("The moderator removed Sally from a group")
	example28Type.SetUnknownActor(actor)
	example28Type.AppendObject(object)
	example28Type.AppendOriginObject(origin)
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

var example29Type = &Undo{}

func init() {
	link := MustParseURL("http://sally.example.org")
	objectLink := MustParseURL("http://example.org/posts/1")
	targetLink := MustParseURL("http://john.example.org")
	object := &Offer{}
	object.AppendType("Offer")
	object.AppendActorIRI(link)
	object.AppendObjectIRI(objectLink)
	object.AppendTargetIRI(targetLink)
	example29Type.AppendType("Undo")
	example29Type.AppendSummaryString("Sally retracted her offer to John")
	example29Type.AppendActorIRI(link)
	example29Type.AppendObject(object)
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

var example30Type = &Update{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example30Type.AppendType("Update")
	example30Type.AppendSummaryString("Sally updated her note")
	example30Type.AppendActorObject(actor)
	example30Type.AppendObjectIRI(link)
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

var example31Type = &View{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	object := &Article{}
	object.AppendType("Article")
	object.AppendNameString("What You Should Know About Activity Streams")
	example31Type.AppendType("View")
	example31Type.AppendSummaryString("Sally read an article")
	example31Type.AppendActorObject(actor)
	example31Type.AppendObject(object)
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

var example32Type = &Listen{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/music.mp3")
	example32Type.AppendType("Listen")
	example32Type.AppendSummaryString("Sally listened to a piece of music")
	example32Type.AppendActorObject(actor)
	example32Type.AppendObjectIRI(link)
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

var example33Type = &Read{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/posts/1")
	example33Type.AppendType("Read")
	example33Type.AppendSummaryString("Sally read a blog post")
	example33Type.AppendActorObject(actor)
	example33Type.AppendObjectIRI(link)
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

var example34Type = &Move{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	link := MustParseURL("http://example.org/posts/1")
	target := &Collection{}
	target.AppendType("Collection")
	target.AppendNameString("List B")
	origin := &Collection{}
	origin.AppendType("Collection")
	origin.AppendNameString("List A")
	example34Type.AppendType("Move")
	example34Type.AppendSummaryString("Sally moved a post from List A to List B")
	example34Type.AppendActorObject(actor)
	example34Type.AppendObjectIRI(link)
	example34Type.AppendTargetObject(target)
	example34Type.AppendOriginObject(origin)
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

var example35Type = &Travel{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	target := &Place{}
	target.AppendType("Place")
	target.AppendNameString("Home")
	origin := &Place{}
	origin.AppendType("Place")
	origin.AppendNameString("Work")
	example35Type.AppendType("Travel")
	example35Type.AppendSummaryString("Sally went home from work")
	example35Type.AppendActorObject(actor)
	example35Type.AppendTargetObject(target)
	example35Type.AppendOriginObject(origin)
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

var example36Type = &Announce{}

func init() {
	link := MustParseURL("http://sally.example.org")
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	actor.SetId(link)
	loc := &Place{}
	loc.AppendType("Place")
	loc.AppendNameString("Work")
	object := &Arrive{}
	object.AppendType("Arrive")
	object.AppendActorIRI(link)
	object.AppendLocationObject(loc)
	example36Type.AppendType("Announce")
	example36Type.AppendSummaryString("Sally announced that she had arrived at work")
	example36Type.AppendActorObject(actor)
	example36Type.AppendObject(object)
}

const example37 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally blocked Joe",
  "type": "Block",
  "actor": "http://sally.example.org",
  "object": "http://joe.example.org"
}`

var example37Type = &Block{}

func init() {
	link := MustParseURL("http://sally.example.org")
	objLink := MustParseURL("http://joe.example.org")
	example37Type.AppendType("Block")
	example37Type.AppendSummaryString("Sally blocked Joe")
	example37Type.AppendActorIRI(link)
	example37Type.AppendObjectIRI(objLink)
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

var example38Type = &Flag{}

func init() {
	link := MustParseURL("http://sally.example.org")
	object := &Note{}
	object.AppendType("Note")
	object.AppendContentString("An inappropriate note")
	example38Type.AppendType("Flag")
	example38Type.AppendSummaryString("Sally flagged an inappropriate note")
	example38Type.AppendActorIRI(link)
	example38Type.AppendObject(object)
}

const example39 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally disliked a post",
  "type": "Dislike",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`

var example39Type = &Dislike{}

func init() {
	link := MustParseURL("http://sally.example.org")
	objLink := MustParseURL("http://example.org/posts/1")
	example39Type.AppendType("Dislike")
	example39Type.AppendSummaryString("Sally disliked a post")
	example39Type.AppendActorIRI(link)
	example39Type.AppendObjectIRI(objLink)
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

var example40Type = &Question{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Option A")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Option B")
	example40Type.AppendType("Question")
	example40Type.AppendNameString("What is the answer?")
	example40Type.AppendOneOfObject(note1)
	example40Type.AppendOneOfObject(note2)
}

const example41 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "closed": "2016-05-10T00:00:00Z"
}`

var example41Type = &Question{}

func init() {
	t, err := time.Parse(time.RFC3339, "2016-05-10T00:00:00Z")
	if err != nil {
		panic(err)
	}
	example41Type.AppendType("Question")
	example41Type.AppendNameString("What is the answer?")
	example41Type.AppendClosedDateTime(t)
}

const example42 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Application",
  "name": "Exampletron 3000"
}`

var example42Type = &Application{}

func init() {
	example42Type.AppendType("Application")
	example42Type.AppendNameString("Exampletron 3000")
}

const example43 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Group",
  "name": "Big Beards of Austin"
}`

var example43Type = &Group{}

func init() {
	example43Type.AppendType("Group")
	example43Type.AppendNameString("Big Beards of Austin")
}

const example44 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Organization",
  "name": "Example Co."
}`

var example44Type = &Organization{}

func init() {
	example44Type.AppendType("Organization")
	example44Type.AppendNameString("Example Co.")
}

const example45 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally Smith"
}`

var example45Type = &Person{}

func init() {
	example45Type.AppendType("Person")
	example45Type.AppendNameString("Sally Smith")
}

const example46 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Service",
  "name": "Acme Web Service"
}`

var example46Type = &Service{}

func init() {
	example46Type.AppendType("Service")
	example46Type.AppendNameString("Acme Web Service")
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

var example47Type = &Relationship{}

func init() {
	subject := &Person{}
	subject.AppendType("Person")
	subject.AppendNameString("Sally")
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("John")
	rel := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	example47Type.AppendType("Relationship")
	example47Type.AppendSummaryString("Sally is an acquaintance of John")
	example47Type.SetSubjectObject(subject)
	example47Type.AppendObject(object)
	example47Type.SetRelationshipIRI(rel)
}

const example48 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Article",
  "name": "What a Crazy Day I Had",
  "content": "<div>... you will never believe ...</div>",
  "attributedTo": "http://sally.example.org"
}`

var example48Type = &Article{}

func init() {
	att := MustParseURL("http://sally.example.org")
	example48Type.AppendType("Article")
	example48Type.AppendNameString("What a Crazy Day I Had")
	example48Type.AppendAttributedToIRI(att)
	example48Type.AppendContentString("<div>... you will never believe ...</div>")
}

const example49 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": "http://example.org/4q-sales-forecast.pdf"
}`

var example49Type = &Document{}

func init() {
	l := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	example49Type.AppendType("Document")
	example49Type.AppendNameString("4Q Sales Forecast")
	example49Type.AppendUrlAnyURI(l)
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

var example50Type = &Audio{}

func init() {
	l := MustParseURL("http://example.org/podcast.mp3")
	link := &Link{}
	link.AppendType("Link")
	link.SetHref(l)
	link.SetMediaType("audio/mp3")
	example50Type.AppendType("Audio")
	example50Type.AppendNameString("Interview With A Famous Technologist")
	example50Type.AppendUrlLink(link)
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

var example51Type = &Image{}

func init() {
	l1 := MustParseURL("http://example.org/image.jpeg")
	l2 := MustParseURL("http://example.org/image.png")
	link1 := &Link{}
	link1.AppendType("Link")
	link1.SetHref(l1)
	link1.SetMediaType("image/jpeg")
	link2 := &Link{}
	link2.AppendType("Link")
	link2.SetHref(l2)
	link2.SetMediaType("image/png")
	example51Type.AppendType("Image")
	example51Type.AppendNameString("Cat Jumping on Wagon")
	example51Type.AppendUrlLink(link1)
	example51Type.AppendUrlLink(link2)
}

const example52 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Puppy Plays With Ball",
  "url": "http://example.org/video.mkv",
  "duration": "PT2H"
}`

var example52Type = &Video{}

func init() {
	l := MustParseURL("http://example.org/video.mkv")
	example52Type.AppendType("Video")
	example52Type.AppendNameString("Puppy Plays With Ball")
	example52Type.AppendUrlAnyURI(l)
	example52Type.SetDuration(time.Hour * 2)
}

const example53 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A Word of Warning",
  "content": "Looks like it is going to rain today. Bring an umbrella!"
}`

var example53Type = &Note{}

func init() {
	example53Type.AppendType("Note")
	example53Type.AppendNameString("A Word of Warning")
	example53Type.AppendContentString("Looks like it is going to rain today. Bring an umbrella!")
}

const example54 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Page",
  "name": "Omaha Weather Report",
  "url": "http://example.org/weather-in-omaha.html"
}`

var example54Type = &Page{}

func init() {
	l := MustParseURL("http://example.org/weather-in-omaha.html")
	example54Type.AppendType("Page")
	example54Type.AppendNameString("Omaha Weather Report")
	example54Type.AppendUrlAnyURI(l)
}

const example55 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

var example55Type = &Event{}

func init() {
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	example55Type.AppendType("Event")
	example55Type.AppendNameString("Going-Away Party for Jim")
	example55Type.SetStartTime(t1)
	example55Type.SetEndTime(t2)
}

const example56 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Work"
}`

var example56Type = &Place{}

func init() {
	example56Type.AppendType("Place")
	example56Type.AppendNameString("Work")
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

var example57Type = &Place{}

func init() {
	example57Type.AppendType("Place")
	example57Type.AppendNameString("Fresno Area")
	example57Type.SetLatitude(36.75)
	example57Type.SetLongitude(119.7667)
	example57Type.SetRadius(15)
	example57Type.SetUnitsUnitsValue("miles")
}

const example58 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Mention of Joe by Carrie in her note",
  "type": "Mention",
  "href": "http://example.org/joe",
  "name": "Joe"
}`

var example58Type = &Mention{}

func init() {
	l := MustParseURL("http://example.org/joe")
	example58Type.AppendType("Mention")
	example58Type.AppendSummaryString("Mention of Joe by Carrie in her note")
	example58Type.SetHref(l)
	example58Type.AppendNameString("Joe")
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

var example59Type = &Profile{}

func init() {
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally Smith")
	example59Type.AppendType("Profile")
	example59Type.AppendSummaryString("Sally's Profile")
	example59Type.SetDescribes(person)
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

var example60Type = &OrderedCollection{}

func init() {
	t, err := time.Parse(time.RFC3339, "2016-03-17T00:00:00Z")
	if err != nil {
		panic(err)
	}
	image1 := &Image{}
	image1.AppendType("Image")
	image1.SetId(MustParseURL("http://image.example/1"))
	tombstone := &Tombstone{}
	tombstone.AppendType("Tombstone")
	tombstone.AppendFormerTypeString("Image")
	tombstone.SetId(MustParseURL("http://image.example/2"))
	tombstone.SetDeleted(t)
	image2 := &Image{}
	image2.AppendType("Image")
	image2.SetId(MustParseURL("http://image.example/3"))
	example60Type.AppendType("OrderedCollection")
	example60Type.SetTotalItems(3)
	example60Type.AppendNameString("Vacation photos 2016")
	example60Type.AppendOrderedItemsObject(image1)
	example60Type.AppendOrderedItemsObject(tombstone)
	example60Type.AppendOrderedItemsObject(image2)
}

const example61 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Foo",
  "id": "http://example.org/foo"
}`

var example61Type = &Unknown{}

func init() {
	example61Type.SetField("id", "http://example.org/foo")
	example61Type.SetField("name", "Foo")
}

const example62 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A foo",
  "type": "http://example.org/Foo"
}`

var example62Type = &Unknown{}

func init() {
	example62Type.SetField("type", "http://example.org/Foo")
	example62Type.SetField("summary", "A foo")
}

const example63 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the Foo object",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/foo"
}`

var example63Type = &Offer{}

func init() {
	l := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/foo")
	example63Type.AppendType("Offer")
	example63Type.AppendSummaryString("Sally offered the Foo object")
	example63Type.AppendActorIRI(l)
	example63Type.AppendObjectIRI(o)
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

var example64Type = &Offer{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.SetId(MustParseURL("http://sally.example.org"))
	actor.AppendSummaryString("Sally")
	o := MustParseURL("http://example.org/foo")
	example64Type.AppendType("Offer")
	example64Type.AppendSummaryString("Sally offered the Foo object")
	example64Type.AppendActorObject(actor)
	example64Type.AppendObjectIRI(o)
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

var example65Type = &Offer{}

func init() {
	actor := &Person{}
	actor.AppendType("Person")
	actor.SetId(MustParseURL("http://sally.example.org"))
	actor.AppendNameString("Sally")
	o := MustParseURL("http://example.org/foo")
	l := MustParseURL("http://joe.example.org")
	example65Type.AppendType("Offer")
	example65Type.AppendSummaryString("Sally and Joe offered the Foo object")
	example65Type.AppendObjectIRI(o)
	example65Type.AppendActorIRI(l)
	example65Type.AppendActorObject(actor)
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

var example66Type = &Note{}

func init() {
	l := MustParseURL("http://example.org/cat.jpeg")
	image := &Image{}
	image.AppendType("Image")
	image.AppendContentString("This is what he looks like.")
	image.AppendUrlAnyURI(l)
	example66Type.AppendType("Note")
	example66Type.AppendNameString("Have you seen my cat?")
	example66Type.AppendAttachmentObject(image)
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

var example67Type = &Image{}

func init() {
	l := MustParseURL("http://example.org/cat.jpeg")
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	example67Type.AppendType("Image")
	example67Type.AppendNameString("My cat taking a nap")
	example67Type.AppendUrlAnyURI(l)
	example67Type.AppendAttributedToObject(person)
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

var example68Type = &Image{}

func init() {
	l := MustParseURL("http://example.org/cat.jpeg")
	a := MustParseURL("http://joe.example.org")
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	example68Type.AppendType("Image")
	example68Type.AppendNameString("My cat taking a nap")
	example68Type.AppendUrlAnyURI(l)
	example68Type.AppendAttributedToIRI(a)
	example68Type.AppendAttributedToObject(person)
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

var example69Type = &Note{}

func init() {
	audience := make(map[string]interface{})
	audience["type"] = "http://example.org/Organization"
	audience["name"] = "ExampleCo LLC"
	example69Type.AppendType("Note")
	example69Type.AppendNameString("Holiday announcement")
	example69Type.AppendContentString("Thursday will be a company-wide holiday. Enjoy your day off!")
	example69Type.AddUnknown("audience", audience)
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

var example70Type = &Offer{}

func init() {
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	example70Type.AppendType("Offer")
	example70Type.AppendSummaryString("Sally offered a post to John")
	example70Type.AppendActorIRI(a)
	example70Type.AppendObjectIRI(o)
	example70Type.AppendTargetIRI(t)
	example70Type.AppendBccIRI(b)
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

var example71Type = &Offer{}

func init() {
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	example71Type.AppendType("Offer")
	example71Type.AppendSummaryString("Sally offered a post to John")
	example71Type.AppendActorIRI(a)
	example71Type.AppendObjectIRI(o)
	example71Type.AppendTargetIRI(t)
	example71Type.AppendBtoIRI(b)
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

var example72Type = &Offer{}

func init() {
	o := MustParseURL("http://example.org/posts/1")
	a := MustParseURL("http://sally.example.org")
	t := MustParseURL("http://john.example.org")
	b := MustParseURL("http://joe.example.org")
	example72Type.AppendType("Offer")
	example72Type.AppendSummaryString("Sally offered a post to John")
	example72Type.AppendActorIRI(a)
	example72Type.AppendObjectIRI(o)
	example72Type.AppendTargetIRI(t)
	example72Type.AppendCcIRI(b)
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

var example73Type = &Collection{}

func init() {
	oa := MustParseURL("http://sally.example.org")
	oo := MustParseURL("http://example.org/posts/1")
	ot := MustParseURL("http://john.example.org")
	oc := MustParseURL("http://example.org/contexts/1")
	offer := &Offer{}
	offer.AppendType("Offer")
	offer.AppendActorIRI(oa)
	offer.AppendObjectIRI(oo)
	offer.AppendTargetIRI(ot)
	offer.AppendContextIRI(oc)
	la := MustParseURL("http://joe.example.org")
	lo := MustParseURL("http://example.org/posts/2")
	lc := MustParseURL("http://example.org/contexts/1")
	like := &Like{}
	like.AppendType("Like")
	like.AppendActorIRI(la)
	like.AppendObjectIRI(lo)
	like.AppendContextIRI(lc)
	example73Type.AppendType("Collection")
	example73Type.AppendSummaryString("Activities in context 1")
	example73Type.AppendItemsObject(offer)
	example73Type.AppendItemsObject(like)
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

var example74Type = &Collection{}

func init() {
	c := MustParseURL("http://example.org/collection")
	i1 := MustParseURL("http://example.org/posts/1")
	i2 := MustParseURL("http://example.org/posts/2")
	i3 := MustParseURL("http://example.org/posts/3")
	example74Type.AppendType("Collection")
	example74Type.SetTotalItems(3)
	example74Type.AppendSummaryString("Sally's blog posts")
	example74Type.SetCurrentIRI(c)
	example74Type.AppendItemsIRI(i1)
	example74Type.AppendItemsIRI(i2)
	example74Type.AppendItemsIRI(i3)
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

var example75Type = &Collection{}

func init() {
	i1 := MustParseURL("http://example.org/posts/1")
	i2 := MustParseURL("http://example.org/posts/2")
	i3 := MustParseURL("http://example.org/posts/3")
	href := MustParseURL("http://example.org/collection")
	link := &Link{}
	link.AppendType("Link")
	link.AppendSummaryString("Most Recent Items")
	link.SetHref(href)
	example75Type.AppendType("Collection")
	example75Type.SetTotalItems(3)
	example75Type.AppendSummaryString("Sally's blog posts")
	example75Type.SetCurrentLink(link)
	example75Type.AppendItemsIRI(i1)
	example75Type.AppendItemsIRI(i2)
	example75Type.AppendItemsIRI(i3)
}

const example76 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "first": "http://example.org/collection?page=0"
}`

var example76Type = &Collection{}

func init() {
	f := MustParseURL("http://example.org/collection?page=0")
	example76Type.AppendType("Collection")
	example76Type.SetTotalItems(3)
	example76Type.AppendSummaryString("Sally's blog posts")
	example76Type.SetFirstIRI(f)
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

var example77Type = &Collection{}

func init() {
	href := MustParseURL("http://example.org/collection?page=0")
	link := &Link{}
	link.AppendType("Link")
	link.AppendSummaryString("First Page")
	link.SetHref(href)
	example77Type.AppendType("Collection")
	example77Type.SetTotalItems(3)
	example77Type.AppendSummaryString("Sally's blog posts")
	example77Type.SetFirstLink(link)
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

var example78Type = &Note{}

func init() {
	app := &Application{}
	app.AppendType("Application")
	app.AppendNameString("Exampletron 3000")
	example78Type.AppendType("Note")
	example78Type.AppendSummaryString("A simple note")
	example78Type.AppendContentString("This is all there is.")
	example78Type.AppendGeneratorObject(app)
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

var example79Type = &Note{}

func init() {
	u := MustParseURL("http://example.org/note.png")
	image := &Image{}
	image.AppendType("Image")
	image.AppendNameString("Note icon")
	image.AppendUrlAnyURI(u)
	image.SetWidth(16)
	image.SetHeight(16)
	example79Type.AppendType("Note")
	example79Type.AppendSummaryString("A simple note")
	example79Type.AppendContentString("This is all there is.")
	example79Type.AppendIconImage(image)
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

var example80Type = &Note{}

func init() {
	u1 := MustParseURL("http://example.org/note1.png")
	u2 := MustParseURL("http://example.org/note2.png")
	image1 := &Image{}
	image1.AppendType("Image")
	image1.AppendSummaryString("Note (16x16)")
	image1.AppendUrlAnyURI(u1)
	image1.SetWidth(16)
	image1.SetHeight(16)
	image2 := &Image{}
	image2.AppendType("Image")
	image2.AppendSummaryString("Note (32x32)")
	image2.AppendUrlAnyURI(u2)
	image2.SetWidth(32)
	image2.SetHeight(32)
	example80Type.AppendType("Note")
	example80Type.AppendSummaryString("A simple note")
	example80Type.AppendContentString("A simple note")
	example80Type.AppendIconImage(image1)
	example80Type.AppendIconImage(image2)
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

var example81Type = &Note{}

func init() {
	u := MustParseURL("http://example.org/cat.png")
	image := &Image{}
	image.AppendType("Image")
	image.AppendNameString("A Cat")
	image.AppendUrlAnyURI(u)
	example81Type.AppendType("Note")
	example81Type.AppendNameString("A simple note")
	example81Type.AppendContentString("This is all there is.")
	example81Type.AppendImageImage(image)
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

var example82Type = &Note{}

func init() {
	u1 := MustParseURL("http://example.org/cat1.png")
	u2 := MustParseURL("http://example.org/cat2.png")
	image1 := &Image{}
	image1.AppendType("Image")
	image1.AppendNameString("Cat 1")
	image1.AppendUrlAnyURI(u1)
	image2 := &Image{}
	image2.AppendType("Image")
	image2.AppendNameString("Cat 2")
	image2.AppendUrlAnyURI(u2)
	example82Type.AppendType("Note")
	example82Type.AppendNameString("A simple note")
	example82Type.AppendContentString("This is all there is.")
	example82Type.AppendImageImage(image1)
	example82Type.AppendImageImage(image2)
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

var example83Type = &Note{}

func init() {
	note := &Note{}
	note.AppendType("Note")
	note.AppendSummaryString("Previous note")
	note.AppendContentString("What else is there?")
	example83Type.AppendType("Note")
	example83Type.AppendSummaryString("A simple note")
	example83Type.AppendContentString("This is all there is.")
	example83Type.AppendInReplyToObject(note)
}

const example84 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "inReplyTo": "http://example.org/posts/1"
}`

var example84Type = &Note{}

func init() {
	u := MustParseURL("http://example.org/posts/1")
	example84Type.AppendType("Note")
	example84Type.AppendSummaryString("A simple note")
	example84Type.AppendContentString("This is all there is.")
	example84Type.AppendInReplyToIRI(u)
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

var example85Type = &Listen{}

func init() {
	u := MustParseURL("http://example.org/foo.mp3")
	actor := &Person{}
	actor.AppendType("Person")
	actor.AppendNameString("Sally")
	service := &Service{}
	service.AppendType("Service")
	service.AppendNameString("Acme Music Service")
	example85Type.AppendType("Listen")
	example85Type.AppendSummaryString("Sally listened to a piece of music on the Acme Music Service")
	example85Type.AppendActorObject(actor)
	example85Type.AppendObjectIRI(u)
	example85Type.AppendInstrumentObject(service)
}

const example86 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A collection",
  "type": "Collection",
  "totalItems": 3,
  "last": "http://example.org/collection?page=1"
}`

var example86Type = &Collection{}

func init() {
	u := MustParseURL("http://example.org/collection?page=1")
	example86Type.AppendType("Collection")
	example86Type.SetTotalItems(3)
	example86Type.AppendSummaryString("A collection")
	example86Type.SetLastIRI(u)
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

var example87Type = &Collection{}

func init() {
	u := MustParseURL("http://example.org/collection?page=1")
	link := &Link{}
	link.AppendType("Link")
	link.AppendSummaryString("Last Page")
	link.SetHref(u)
	example87Type.AppendType("Collection")
	example87Type.AppendSummaryString("A collection")
	example87Type.SetTotalItems(5)
	example87Type.SetLastLink(link)
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

var example88Type = &Person{}

func init() {
	place := &Place{}
	place.AppendType("Place")
	place.AppendNameString("Over the Arabian Sea, east of Socotra Island Nature Sanctuary")
	place.SetLongitude(12.34)
	place.SetLatitude(56.78)
	place.SetAltitude(90)
	place.SetUnitsUnitsValue("m")
	example88Type.AppendType("Person")
	example88Type.AppendNameString("Sally")
	example88Type.AppendLocationObject(place)
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

var example89Type = &Collection{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Reminder for Going-Away Party")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Meeting 2016-11-17")
	example89Type.AppendType("Collection")
	example89Type.AppendSummaryString("Sally's notes")
	example89Type.SetTotalItems(2)
	example89Type.AppendItemsObject(note1)
	example89Type.AppendItemsObject(note2)
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

var example90Type = &OrderedCollection{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Meeting 2016-11-17")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Reminder for Going-Away Party")
	example90Type.AppendType("OrderedCollection")
	example90Type.AppendSummaryString("Sally's notes")
	example90Type.SetTotalItems(2)
	example90Type.AppendOrderedItemsObject(note1)
	example90Type.AppendOrderedItemsObject(note2)
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

var example91Type = &Question{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Option A")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Option B")
	example91Type.AppendType("Question")
	example91Type.AppendNameString("What is the answer?")
	example91Type.AppendOneOfObject(note1)
	example91Type.AppendOneOfObject(note2)
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

var example92Type = &Question{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Option A")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Option B")
	example92Type.AppendType("Question")
	example92Type.AppendNameString("What is the answer?")
	example92Type.AppendAnyOfObject(note1)
	example92Type.AppendAnyOfObject(note2)
}

const example93 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Question",
  "name": "What is the answer?",
  "closed": "2016-05-10T00:00:00Z"
}`

var example93Type = &Question{}

func init() {
	t, err := time.Parse(time.RFC3339, "2016-05-10T00:00:00Z")
	if err != nil {
		panic(err)
	}
	example93Type.AppendType("Question")
	example93Type.AppendNameString("What is the answer?")
	example93Type.AppendClosedDateTime(t)
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

var example94Type = &Move{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	target := &Collection{}
	target.AppendType("Collection")
	target.AppendNameString("List B")
	origin := &Collection{}
	origin.AppendType("Collection")
	origin.AppendNameString("List A")
	example94Type.AppendType("Move")
	example94Type.AppendSummaryString("Sally moved a post from List A to List B")
	example94Type.AppendActorIRI(a)
	example94Type.AppendObjectIRI(o)
	example94Type.AppendTargetObject(target)
	example94Type.AppendOriginObject(origin)
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

var example95Type = &CollectionPage{}

func init() {
	i := MustParseURL("http://example.org/collection?page=2")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	example95Type.AppendType("CollectionPage")
	example95Type.AppendSummaryString("Page 2 of Sally's blog posts")
	example95Type.SetNextIRI(i)
	example95Type.AppendItemsIRI(u1)
	example95Type.AppendItemsIRI(u2)
	example95Type.AppendItemsIRI(u3)
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

var example96Type = &CollectionPage{}

func init() {
	href := MustParseURL("http://example.org/collection?page=2")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	link := &Link{}
	link.AppendType("Link")
	link.AppendNameString("Next Page")
	link.SetHref(href)
	example96Type.AppendType("CollectionPage")
	example96Type.AppendSummaryString("Page 2 of Sally's blog posts")
	example96Type.SetNextLink(link)
	example96Type.AppendItemsIRI(u1)
	example96Type.AppendItemsIRI(u2)
	example96Type.AppendItemsIRI(u3)
}

const example97 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally liked a post",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`

var example97Type = &Like{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	example97Type.AppendType("Like")
	example97Type.AppendSummaryString("Sally liked a post")
	example97Type.AppendActorIRI(a)
	example97Type.AppendObjectIRI(o)
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

var example98Type = &Like{}

func init() {
	a := MustParseURL("http://sally.example.org")
	note := &Note{}
	note.AppendType("Note")
	note.AppendContentString("A simple note")
	example98Type.AppendType("Like")
	example98Type.AppendActorIRI(a)
	example98Type.AppendObject(note)
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

var example99Type = &Like{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	note := &Note{}
	note.AppendType("Note")
	note.AppendSummaryString("A simple note")
	note.AppendContentString("That is a tree.")
	example99Type.AppendType("Like")
	example99Type.AppendSummaryString("Sally liked a note")
	example99Type.AppendActorIRI(a)
	example99Type.AppendObjectIRI(o)
	example99Type.AppendObject(note)
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

var example100Type = &CollectionPage{}

func init() {
	p := MustParseURL("http://example.org/collection?page=1")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	example100Type.AppendType("CollectionPage")
	example100Type.AppendSummaryString("Page 1 of Sally's blog posts")
	example100Type.SetPrevIRI(p)
	example100Type.AppendItemsIRI(u1)
	example100Type.AppendItemsIRI(u2)
	example100Type.AppendItemsIRI(u3)
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

var example101Type = &CollectionPage{}

func init() {
	p := MustParseURL("http://example.org/collection?page=1")
	u1 := MustParseURL("http://example.org/posts/1")
	u2 := MustParseURL("http://example.org/posts/2")
	u3 := MustParseURL("http://example.org/posts/3")
	link := &Link{}
	link.AppendType("Link")
	link.AppendNameString("Previous Page")
	link.SetHref(p)
	example101Type.AppendType("CollectionPage")
	example101Type.AppendSummaryString("Page 1 of Sally's blog posts")
	example101Type.SetPrevLink(link)
	example101Type.AppendItemsIRI(u1)
	example101Type.AppendItemsIRI(u2)
	example101Type.AppendItemsIRI(u3)
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

var example102Type = &Video{}

func init() {
	u := MustParseURL("http://example.org/trailer.mkv")
	link := &Link{}
	link.AppendType("Link")
	link.SetMediaType("video/mkv")
	link.SetHref(u)
	video := &Video{}
	video.AppendType("Video")
	video.AppendNameString("Trailer")
	video.SetDuration(time.Minute)
	video.AppendUrlLink(link)
	example102Type.AppendType("Video")
	example102Type.AppendNameString("Cool New Movie")
	example102Type.SetDuration(time.Hour*2 + time.Minute*30)
	example102Type.AppendPreviewObject(video)
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

var example103Type = &Activity{}

func init() {
	o := MustParseURL("http://example.org/flights/1")
	a := MustParseURL("http://sally.example.org")
	status := make(map[string]interface{})
	status["type"] = "http://www.types.example/flightstatus"
	status["name"] = "On Time"
	example103Type.AppendType("Activity")
	example103Type.AppendType("http://www.verbs.example/Check")
	example103Type.AppendSummaryString("Sally checked that her flight was on time")
	example103Type.AppendActorIRI(a)
	example103Type.AppendObjectIRI(o)
	example103Type.AddUnknown("result", status)
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

var example104Type = &Note{}

func init() {
	i := MustParseURL("http://www.test.example/notes/1")
	note := &Note{}
	note.AppendType("Note")
	note.AppendSummaryString("A response to the note")
	note.AppendContentString("I am glad to hear it.")
	note.AppendInReplyToIRI(i)
	replies := &Collection{}
	replies.AppendType("Collection")
	replies.SetTotalItems(1)
	replies.AppendItemsObject(note)
	example104Type.AppendType("Note")
	example104Type.AppendSummaryString("A simple note")
	example104Type.SetId(MustParseURL("http://www.test.example/notes/1"))
	example104Type.AppendContentString("I am fine.")
	example104Type.SetReplies(replies)
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

var example105Type = &Image{}

func init() {
	u := MustParseURL("http://example.org/sally.jpg")
	person := &Person{}
	person.AppendType("Person")
	person.SetId(MustParseURL("http://sally.example.org"))
	person.AppendNameString("Sally")
	example105Type.AppendType("Image")
	example105Type.AppendSummaryString("Picture of Sally")
	example105Type.AppendUrlAnyURI(u)
	example105Type.AppendTagObject(person)
}

const example106 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org"
}`

var example106Type = &Offer{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	t := MustParseURL("http://john.example.org")
	example106Type.AppendType("Offer")
	example106Type.AppendSummaryString("Sally offered the post to John")
	example106Type.AppendActorIRI(a)
	example106Type.AppendObjectIRI(o)
	example106Type.AppendTargetIRI(t)
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

var example107Type = &Offer{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("John")
	example107Type.AppendType("Offer")
	example107Type.AppendSummaryString("Sally offered the post to John")
	example107Type.AppendActorIRI(a)
	example107Type.AppendObjectIRI(o)
	example107Type.AppendTargetObject(person)
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

var example108Type = &Offer{}

func init() {
	a := MustParseURL("http://sally.example.org")
	o := MustParseURL("http://example.org/posts/1")
	t := MustParseURL("http://john.example.org")
	z := MustParseURL("http://joe.example.org")
	example108Type.AppendType("Offer")
	example108Type.AppendSummaryString("Sally offered the post to John")
	example108Type.AppendActorIRI(a)
	example108Type.AppendObjectIRI(o)
	example108Type.AppendTargetIRI(t)
	example108Type.AppendToIRI(z)
}

const example109 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": "http://example.org/4q-sales-forecast.pdf"
}`

var example109Type = &Document{}

func init() {
	u := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	example109Type.AppendType("Document")
	example109Type.AppendNameString("4Q Sales Forecast")
	example109Type.AppendUrlAnyURI(u)
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

var example110Type = &Document{}

func init() {
	u := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	link := &Link{}
	link.AppendType("Link")
	link.SetHref(u)
	example110Type.AppendType("Document")
	example110Type.AppendNameString("4Q Sales Forecast")
	example110Type.AppendUrlLink(link)
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

var example111Type = &Document{}

func init() {
	u1 := MustParseURL("http://example.org/4q-sales-forecast.pdf")
	u2 := MustParseURL("http://example.org/4q-sales-forecast.html")
	link1 := &Link{}
	link1.AppendType("Link")
	link1.SetHref(u1)
	link1.SetMediaType("application/pdf")
	link2 := &Link{}
	link2.AppendType("Link")
	link2.SetHref(u2)
	link2.SetMediaType("text/html")
	example111Type.AppendType("Document")
	example111Type.AppendNameString("4Q Sales Forecast")
	example111Type.AppendUrlLink(link1)
	example111Type.AppendUrlLink(link2)
}

const example112 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China",
  "type": "Place",
  "latitude": 36.75,
  "longitude": 119.7667,
  "accuracy": 94.5
}`

var example112Type = &Place{}

func init() {
	example112Type.AppendType("Place")
	example112Type.AppendNameString("Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China")
	example112Type.SetLatitude(36.75)
	example112Type.SetLongitude(119.7667)
	example112Type.SetAccuracy(94.5)
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

var example113Type = &Place{}

func init() {
	example113Type.AppendType("Place")
	example113Type.AppendNameString("Fresno Area")
	example113Type.SetAltitude(15.0)
	example113Type.SetLatitude(36.75)
	example113Type.SetLongitude(119.7667)
	example113Type.SetUnitsUnitsValue("miles")
}

const example114 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "A <em>simple</em> note"
}`

var example114Type = &Note{}

func init() {
	example114Type.AppendType("Note")
	example114Type.AppendSummaryString("A simple note")
	example114Type.AppendContentString("A <em>simple</em> note")
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

var example115Type = &Note{}

func init() {
	example115Type.AppendType("Note")
	example115Type.AppendSummaryString("A simple note")
	example115Type.SetContentMap("en", "A <em>simple</em> note")
	example115Type.SetContentMap("es", "Una nota <em>sencilla</em>")
	example115Type.SetContentMap("zh-Hans", "一段<em>简单的</em>笔记")
}

const example116 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "mediaType": "text/markdown",
  "content": "## A simple note\nA simple markdown ` + "`note`" + `"
}`

var example116Type = &Note{}

func init() {
	example116Type.AppendType("Note")
	example116Type.AppendSummaryString("A simple note")
	example116Type.SetMediaType("text/markdown")
	example116Type.AppendContentString("## A simple note\nA simple markdown `note`")
}

const example117 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A simple note"
}`

var example117Type = &Note{}

func init() {
	example117Type.AppendType("Note")
	example117Type.AppendNameString("A simple note")
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

var example118Type = &Note{}

func init() {
	example118Type.AppendType("Note")
	example118Type.SetNameMap("en", "A simple note")
	example118Type.SetNameMap("es", "Una nota sencilla")
	example118Type.SetNameMap("zh-Hans", "一段简单的笔记")
}

const example119 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Birds Flying",
  "url": "http://example.org/video.mkv",
  "duration": "PT2H"
}`

var example119Type = &Video{}

func init() {
	u := MustParseURL("http://example.org/video.mkv")
	example119Type.AppendType("Video")
	example119Type.AppendNameString("Birds Flying")
	example119Type.AppendUrlAnyURI(u)
	example119Type.SetDuration(time.Hour * 2)
}

const example120 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/image.png",
  "height": 100,
  "width": 100
}`

var example120Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/image.png")
	example120Type.AppendType("Link")
	example120Type.SetHref(u)
	example120Type.SetHeight(100)
	example120Type.SetWidth(100)
}

const example121 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "mediaType": "text/html",
  "name": "Previous"
}`

var example121Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/abc")
	example121Type.AppendType("Link")
	example121Type.SetHref(u)
	example121Type.SetMediaType("text/html")
	example121Type.AppendNameString("Previous")
}

const example122 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "Previous"
}`

var example122Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/abc")
	example122Type.AppendType("Link")
	example122Type.SetHref(u)
	example122Type.SetMediaType("text/html")
	example122Type.AppendNameString("Previous")
	example122Type.SetHreflang("en")
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

var example123Type = &CollectionPage{}

func init() {
	u := MustParseURL("http://example.org/collection")
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Pizza Toppings to Try")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Thought about California")
	example123Type.AppendType("CollectionPage")
	example123Type.AppendSummaryString("Page 1 of Sally's notes")
	example123Type.SetId(MustParseURL("http://example.org/collection?page=1"))
	example123Type.SetPartOfIRI(u)
	example123Type.AppendItemsObject(note1)
	example123Type.AppendItemsObject(note2)
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

var example124Type = &Place{}

func init() {
	example124Type.AppendType("Place")
	example124Type.AppendNameString("Fresno Area")
	example124Type.SetLatitude(36.75)
	example124Type.SetLongitude(119.7667)
	example124Type.SetRadius(15)
	example124Type.SetUnitsUnitsValue("miles")
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

var example125Type = &Place{}

func init() {
	example125Type.AppendType("Place")
	example125Type.AppendNameString("Fresno Area")
	example125Type.SetLatitude(36.75)
	example125Type.SetLongitude(119.7667)
	example125Type.SetRadius(15)
	example125Type.SetUnitsUnitsValue("miles")
}

const example126 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "Next"
}`

var example126Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/abc")
	example126Type.AppendType("Link")
	example126Type.SetHref(u)
	example126Type.SetHreflang("en")
	example126Type.SetMediaType("text/html")
	example126Type.AppendNameString("Next")
}

const example127 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

var example127Type = &Event{}

func init() {
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	example127Type.AppendType("Event")
	example127Type.AppendNameString("Going-Away Party for Jim")
	example127Type.SetStartTime(t1)
	example127Type.SetEndTime(t2)
}

const example128 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "Fish swim.",
  "published": "2014-12-12T12:12:12Z"
}`

var example128Type = &Note{}

func init() {
	t, err := time.Parse(time.RFC3339, "2014-12-12T12:12:12Z")
	if err != nil {
		panic(err)
	}
	example128Type.AppendType("Note")
	example128Type.AppendSummaryString("A simple note")
	example128Type.AppendContentString("Fish swim.")
	example128Type.SetPublished(t)
}

const example129 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "Going-Away Party for Jim",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`

var example129Type = &Event{}

func init() {
	t1, err := time.Parse(time.RFC3339, "2014-12-31T23:00:00-08:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(time.RFC3339, "2015-01-01T06:00:00-08:00")
	if err != nil {
		panic(err)
	}
	example129Type.AppendType("Event")
	example129Type.AppendNameString("Going-Away Party for Jim")
	example129Type.SetStartTime(t1)
	example129Type.SetEndTime(t2)
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

var example130Type = &Place{}

func init() {
	example130Type.AppendType("Place")
	example130Type.AppendNameString("Fresno Area")
	example130Type.SetLatitude(36.75)
	example130Type.SetLongitude(119.7667)
	example130Type.SetRadius(15)
	example130Type.SetUnitsUnitsValue("miles")
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

var example131Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/abc")
	example131Type.AppendType("Link")
	example131Type.SetHref(u)
	example131Type.SetHreflang("en")
	example131Type.SetMediaType("text/html")
	example131Type.AppendNameString("Preview")
	example131Type.AppendRel("canonical")
	example131Type.AppendRel("preview")
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

var example132Type = &OrderedCollectionPage{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Density of Water")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Air Mattress Idea")
	example132Type.AppendType("OrderedCollectionPage")
	example132Type.AppendSummaryString("Page 1 of Sally's notes")
	example132Type.SetStartIndex(0)
	example132Type.AppendOrderedItemsObject(note1)
	example132Type.AppendOrderedItemsObject(note2)
}

const example133 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cane Sugar Processing",
  "type": "Note",
  "summary": "A simple <em>note</em>"
}`

var example133Type = &Note{}

func init() {
	example133Type.AppendType("Note")
	example133Type.AppendNameString("Cane Sugar Processing")
	example133Type.AppendSummaryString("A simple <em>note</em>")
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

var example134Type = &Note{}

func init() {
	example134Type.AppendType("Note")
	example134Type.AppendNameString("Cane Sugar Processing")
	example134Type.SetSummaryMap("en", "A simple <em>note</em>")
	example134Type.SetSummaryMap("es", "Una <em>nota</em> sencilla")
	example134Type.SetSummaryMap("zh-Hans", "一段<em>简单的</em>笔记")
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

var example135Type = &Collection{}

func init() {
	note1 := &Note{}
	note1.AppendType("Note")
	note1.AppendNameString("Which Staircase Should I Use")
	note2 := &Note{}
	note2.AppendType("Note")
	note2.AppendNameString("Something to Remember")
	example135Type.AppendType("Collection")
	example135Type.AppendSummaryString("Sally's notes")
	example135Type.SetTotalItems(2)
	example135Type.AppendItemsObject(note1)
	example135Type.AppendItemsObject(note2)
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

var example136Type = &Place{}

func init() {
	example136Type.AppendType("Place")
	example136Type.AppendNameString("Fresno Area")
	example136Type.SetLatitude(36.75)
	example136Type.SetLongitude(119.7667)
	example136Type.SetRadius(15)
	example136Type.SetUnitsUnitsValue("miles")
}

const example137 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cranberry Sauce Idea",
  "type": "Note",
  "content": "Mush it up so it does not have the same shape as the can.",
  "updated": "2014-12-12T12:12:12Z"
}`

var example137Type = &Note{}

func init() {
	t, err := time.Parse(time.RFC3339, "2014-12-12T12:12:12Z")
	if err != nil {
		panic(err)
	}
	example137Type.AppendType("Note")
	example137Type.AppendNameString("Cranberry Sauce Idea")
	example137Type.AppendContentString("Mush it up so it does not have the same shape as the can.")
	example137Type.SetUpdated(t)
}

const example138 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/image.png",
  "height": 100,
  "width": 100
}`

var example138Type = &Link{}

func init() {
	u := MustParseURL("http://example.org/image.png")
	example138Type.AppendType("Link")
	example138Type.SetHref(u)
	example138Type.SetHeight(100)
	example138Type.SetWidth(100)
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

var example139Type = &Relationship{}

func init() {
	u := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	subject := &Person{}
	subject.AppendType("Person")
	subject.AppendNameString("Sally")
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("John")
	example139Type.AppendType("Relationship")
	example139Type.AppendSummaryString("Sally is an acquaintance of John's")
	example139Type.SetSubjectObject(subject)
	example139Type.AppendObject(object)
	example139Type.SetRelationshipIRI(u)
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

var example140Type = &Relationship{}

func init() {
	u := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	subject := &Person{}
	subject.AppendType("Person")
	subject.AppendNameString("Sally")
	object := &Person{}
	object.AppendType("Person")
	object.AppendNameString("John")
	example140Type.AppendType("Relationship")
	example140Type.AppendSummaryString("Sally is an acquaintance of John's")
	example140Type.SetSubjectObject(subject)
	example140Type.AppendObject(object)
	example140Type.SetRelationshipIRI(u)
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

var example141Type = &Profile{}

func init() {
	u := MustParseURL("http://sally.example.org")
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	example141Type.AppendType("Profile")
	example141Type.AppendSummaryString("Sally's profile")
	example141Type.AppendUrlAnyURI(u)
	example141Type.SetDescribes(person)
}

const example142 = `{
"@context": "https://www.w3.org/ns/activitystreams",
"summary": "This image has been deleted",
"type": "Tombstone",
"formerType": "Image",
"url": "http://example.org/image/2"
}`

var example142Type = &Tombstone{}

func init() {
	u := MustParseURL("http://example.org/image/2")
	example142Type.AppendType("Tombstone")
	example142Type.AppendSummaryString("This image has been deleted")
	example142Type.AppendFormerTypeString("Image")
	example142Type.AppendUrlAnyURI(u)
}

const example143 = `{
"@context": "https://www.w3.org/ns/activitystreams",
"summary": "This image has been deleted",
"type": "Tombstone",
"deleted": "2016-05-03T00:00:00Z"
}`

var example143Type = &Tombstone{}

func init() {
	t, err := time.Parse(time.RFC3339, "2016-05-03T00:00:00Z")
	if err != nil {
		panic(err)
	}
	example143Type.AppendType("Tombstone")
	example143Type.AppendSummaryString("This image has been deleted")
	example143Type.SetDeleted(t)
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

var example144Type = &Collection{}

func init() {
	sally := MustParseURL("http://sally.example.org")
	john := MustParseURL("http://john.example.org")
	o := MustParseURL("http://notes.example.com/1")
	context := make(map[string]interface{})
	context["type"] = "http://example.org/Project"
	context["name"] = "Project XYZ"
	audience := &Group{}
	audience.AppendType("Group")
	audience.AppendNameString("Project XYZ Working Group")
	note := &Note{}
	note.AppendType("Note")
	note.AppendSummaryString("A note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AppendContentString("A note")
	create := &Create{}
	create.AppendType("Create")
	create.AppendSummaryString("Sally created a note")
	create.SetId(MustParseURL("http://activities.example.com/1"))
	create.AppendActorIRI(sally)
	create.AppendObject(note)
	create.SetUnknownContext(context)
	create.AppendAudienceObject(audience)
	create.AppendToIRI(john)
	like := &Like{}
	like.AppendType("Like")
	like.AppendSummaryString("John liked Sally's note")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AppendActorIRI(john)
	like.AppendObjectIRI(o)
	like.SetUnknownContext(context)
	like.AppendAudienceObject(audience)
	like.AppendToIRI(sally)
	example144Type.AppendType("Collection")
	example144Type.AppendSummaryString("Activities in Project XYZ")
	example144Type.AppendItemsObject(create)
	example144Type.AppendItemsObject(like)
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

var example145Type = &Collection{}

func init() {
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	influenced := MustParseURL("http://purl.org/vocab/relationship/influencedBy")
	sally := &Person{}
	sally.AppendType("Person")
	sally.AppendNameString("Sally")
	jane := &Person{}
	jane.AppendType("Person")
	jane.AppendNameString("Jane")
	joe := &Person{}
	joe.AppendType("Person")
	joe.AppendNameString("Joe")
	joeRel := &Relationship{}
	joeRel.AppendType("Relationship")
	joeRel.AppendSummaryString("Sally is influenced by Joe")
	joeRel.SetSubjectObject(sally)
	joeRel.AppendObject(joe)
	joeRel.SetRelationshipIRI(influenced)
	janeRel := &Relationship{}
	janeRel.AppendType("Relationship")
	janeRel.AppendSummaryString("Sally is a friend of Jane")
	janeRel.SetSubjectObject(sally)
	janeRel.AppendObject(jane)
	janeRel.SetRelationshipIRI(friend)
	example145Type.AppendType("Collection")
	example145Type.AppendSummaryString("Sally's friends list")
	example145Type.AppendItemsObject(joeRel)
	example145Type.AppendItemsObject(janeRel)
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

var example146Type = &Create{}

func init() {
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	m := MustParseURL("http://matt.example.org")
	s := MustParseURL("http://sally.example.org")
	t, err := time.Parse(time.RFC3339, "2015-04-21T12:34:56Z")
	if err != nil {
		panic(err)
	}
	relationship := &Relationship{}
	relationship.AppendType("Relationship")
	relationship.SetSubjectIRI(s)
	relationship.SetRelationshipIRI(friend)
	relationship.AppendObjectIRI(m)
	relationship.SetStartTime(t)
	example146Type.AppendType("Create")
	example146Type.AppendSummaryString("Sally became a friend of Matt")
	example146Type.AppendActorIRI(s)
	example146Type.AppendObject(relationship)
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

var example147Type = &Offer{}

func init() {
	friend := MustParseURL("http://purl.org/vocab/relationship/friendOf")
	s := MustParseURL("acct:sally@example.org")
	t := MustParseURL("acct:john@example.org")
	rel := &Relationship{}
	rel.AppendType("Relationship")
	rel.AppendSummaryString("Sally and John's friendship")
	rel.SetId(MustParseURL("http://example.org/connections/123"))
	rel.SetSubjectIRI(s)
	rel.AppendObjectIRI(t)
	rel.SetRelationshipIRI(friend)
	example147Type.AppendType("Offer")
	example147Type.AppendSummaryString("Sally requested to be a friend of John")
	example147Type.SetId(MustParseURL("http://example.org/connection-requests/123"))
	example147Type.AppendActorIRI(s)
	example147Type.AppendObject(rel)
	example147Type.AppendTargetIRI(t)
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

var example148Type = &Collection{}

func init() {
	john := MustParseURL("acct:john@example.org")
	sally := MustParseURL("acct:sally@example.org")
	req123 := MustParseURL("http://example.org/connection-requests/123")
	conn123 := MustParseURL("http://example.org/connections/123")
	a123 := MustParseURL("http://example.org/activities/123")
	a124 := MustParseURL("http://example.org/activities/124")
	a125 := MustParseURL("http://example.org/activities/125")
	a126 := MustParseURL("http://example.org/activities/126")
	jc := &Collection{}
	jc.AppendType("Collection")
	jc.AppendSummaryString("John's Connections")
	sc := &Collection{}
	sc.AppendType("Collection")
	sc.AppendSummaryString("Sally's Connections")
	o1 := &Accept{}
	o1.AppendType("Accept")
	o1.SetId(MustParseURL("http://example.org/activities/122"))
	o1.AppendSummaryString("John accepted Sally's friend request")
	o1.AppendObjectIRI(req123)
	o1.AppendInReplyToIRI(req123)
	o1.AppendContextIRI(conn123)
	o1.AppendResultIRI(a123)
	o1.AppendResultIRI(a124)
	o1.AppendResultIRI(a125)
	o1.AppendResultIRI(a126)
	o1.AppendActorIRI(john)
	o2 := &Follow{}
	o2.AppendType("Follow")
	o2.SetId(MustParseURL("http://example.org/activities/123"))
	o2.AppendActorIRI(john)
	o2.AppendObjectIRI(sally)
	o2.AppendContextIRI(conn123)
	o2.AppendSummaryString("John followed Sally")
	o3 := &Follow{}
	o3.AppendType("Follow")
	o3.SetId(MustParseURL("http://example.org/activities/124"))
	o3.AppendActorIRI(sally)
	o3.AppendObjectIRI(john)
	o3.AppendContextIRI(conn123)
	o3.AppendSummaryString("Sally followed John")
	o4 := &Add{}
	o4.AppendType("Add")
	o4.SetId(MustParseURL("http://example.org/activities/125"))
	o4.AppendSummaryString("John added Sally to his friends list")
	o4.AppendActorIRI(john)
	o4.AppendObjectIRI(conn123)
	o4.AppendContextIRI(conn123)
	o4.AppendTargetObject(jc)
	o5 := &Add{}
	o5.AppendType("Add")
	o5.SetId(MustParseURL("http://example.org/activities/126"))
	o5.AppendSummaryString("Sally added John to her friends list")
	o5.AppendActorIRI(sally)
	o5.AppendObjectIRI(conn123)
	o5.AppendContextIRI(conn123)
	o5.AppendTargetObject(sc)
	example148Type.AppendType("Collection")
	example148Type.AppendSummaryString("Sally and John's relationship history")
	example148Type.AppendItemsObject(o1)
	example148Type.AppendItemsObject(o2)
	example148Type.AppendItemsObject(o3)
	example148Type.AppendItemsObject(o4)
	example148Type.AppendItemsObject(o5)
}

const example149 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "San Francisco, CA"
}`

var example149Type = &Place{}

func init() {
	example149Type.AppendType("Place")
	example149Type.AppendNameString("San Francisco, CA")
}

// NOTE: Un-stringified the longitude and latitude values.
const example150 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "San Francisco, CA",
  "longitude": 122.4167,
  "latitude": 37.7833
}`

var example150Type = &Place{}

func init() {
	example150Type.AppendType("Place")
	example150Type.AppendNameString("San Francisco, CA")
	example150Type.SetLongitude(122.4167)
	example150Type.SetLatitude(37.7833)
}

const example151 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "A question about robots",
  "id": "http://help.example.org/question/1",
  "type": "Question",
  "content": "I'd like to build a robot to feed my cat. Should I use Arduino or Raspberry Pi?"
}`

var example151Type = &Question{}

func init() {
	example151Type.AppendType("Question")
	example151Type.AppendNameString("A question about robots")
	example151Type.SetId(MustParseURL("http://help.example.org/question/1"))
	example151Type.AppendContentString("I'd like to build a robot to feed my cat. Should I use Arduino or Raspberry Pi?")
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

var example152Type = &Question{}

func init() {
	ard := make(map[string]interface{})
	ard["name"] = "arduino"
	ras := make(map[string]interface{})
	ras["name"] = "raspberry pi"
	oneOf := []interface{}{ard, ras}
	example152Type.AppendType("Question")
	example152Type.AppendNameString("A question about robots")
	example152Type.SetId(MustParseURL("http://polls.example.org/question/1"))
	example152Type.AppendContentString("I'd like to build a robot to feed my cat. Which platform is best?")
	example152Type.SetUnknownOneOf(oneOf)
}

const example153 = `{
 "@context": "https://www.w3.org/ns/activitystreams",
 "attributedTo": "http://sally.example.org",
 "inReplyTo": "http://polls.example.org/question/1",
 "name": "arduino"
}`

var example153Type = &Unknown{}

func init() {
	example153Type.SetField("attributedTo", "http://sally.example.org")
	example153Type.SetField("inReplyTo", "http://polls.example.org/question/1")
	example153Type.SetField("name", "arduino")
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

var example154Type = &Question{}

func init() {
	ard := make(map[string]interface{})
	ard["name"] = "arduino"
	ras := make(map[string]interface{})
	ras["name"] = "raspberry pi"
	oneOf := []interface{}{ard, ras}
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
	replies := &Collection{}
	replies.AppendType("Collection")
	replies.SetTotalItems(3)
	replies.SetUnknownItems(items)
	note := &Note{}
	note.AppendType("Note")
	note.AppendContentString("Users are favoriting &quot;arduino&quot; by a 33% margin.")
	example154Type.AppendType("Question")
	example154Type.AppendNameString("A question about robots")
	example154Type.SetId(MustParseURL("http://polls.example.org/question/1"))
	example154Type.AppendContentString("I'd like to build a robot to feed my cat. Which platform is best?")
	example154Type.SetUnknownOneOf(oneOf)
	example154Type.SetReplies(replies)
	example154Type.AppendResultObject(note)
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

var example155Type = &Collection{}

func init() {
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
	note := &Note{}
	note.AppendType("Note")
	note.AppendSummaryString("John's note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AppendContentString("My note")
	note.AppendAttributedToIRI(john)
	like := &Like{}
	like.AppendType("Like")
	like.AppendSummaryString("Sally liked John's note")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AppendActorIRI(sally)
	like.SetPublished(t1)
	like.AppendObject(note)
	dislike := &Dislike{}
	dislike.AppendType("Dislike")
	dislike.AppendSummaryString("Sally disliked John's note")
	dislike.SetId(MustParseURL("http://activities.example.com/2"))
	dislike.AppendActorIRI(sally)
	dislike.SetPublished(t2)
	dislike.AppendObject(note)
	example155Type.AppendType("Collection")
	example155Type.AppendSummaryString("History of John's note")
	example155Type.AppendItemsObject(like)
	example155Type.AppendItemsObject(dislike)
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

var example156Type = &Collection{}

func init() {
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
	note := &Note{}
	note.AppendType("Note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AppendSummaryString("John's note")
	note.AppendAttributedToIRI(john)
	note.AppendContentString("My note")
	like := &Like{}
	like.AppendType("Like")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AppendSummaryString("Sally liked John's note")
	like.AppendActorIRI(sally)
	like.SetPublished(t1)
	like.AppendObject(note)
	undo := &Undo{}
	undo.AppendType("Undo")
	undo.SetId(MustParseURL("http://activities.example.com/2"))
	undo.AppendSummaryString("Sally no longer likes John's note")
	undo.AppendActorIRI(sally)
	undo.SetPublished(t2)
	undo.AppendObjectIRI(a)
	example156Type.AppendType("Collection")
	example156Type.AppendSummaryString("History of John's note")
	example156Type.AppendItemsObject(like)
	example156Type.AppendItemsObject(undo)
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

var example157Type = &Note{}

func init() {
	tag := make(map[string]interface{})
	tag["id"] = "http://example.org/tags/givingthanks"
	tag["name"] = "#givingthanks"
	person := &Person{}
	person.AppendType("Person")
	person.AppendNameString("Sally")
	person.SetId(MustParseURL("http://sally.example.org"))
	example157Type.AppendType("Note")
	example157Type.AppendNameString("A thank-you note")
	example157Type.AppendContentString("Thank you <a href='http://sally.example.org'>@sally</a> for all your hard work! <a href='http://example.org/tags/givingthanks'>#givingthanks</a>")
	example157Type.AppendToObject(person)
	example157Type.SetUnknownTag(tag)
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

var example158Type = &Note{}

func init() {
	u := MustParseURL("http://example.org/people/sally")
	tag := make(map[string]interface{})
	tag["id"] = "http://example.org/tags/givingthanks"
	tag["name"] = "#givingthanks"
	mention := &Mention{}
	mention.AppendType("Mention")
	mention.SetHref(u)
	mention.AppendNameString("@sally")
	example158Type.AppendType("Note")
	example158Type.AppendNameString("A thank-you note")
	example158Type.AppendContentString("Thank you @sally for all your hard work! #givingthanks")
	example158Type.AppendTagLink(mention)
	example158Type.SetUnknownTag(tag)
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

var example159Type = &Move{}

func init() {
	sally := MustParseURL("http://sally.example.org")
	obj := &Document{}
	obj.AppendType("Document")
	obj.AppendNameString("sales figures")
	origin := &Collection{}
	origin.AppendType("Collection")
	origin.AppendNameString("Folder A")
	target := &Collection{}
	target.AppendType("Collection")
	target.AppendNameString("Folder B")
	example159Type.AppendType("Move")
	example159Type.AppendSummaryString("Sally moved the sales figures from Folder A to Folder B")
	example159Type.AppendActorIRI(sally)
	example159Type.AppendObject(obj)
	example159Type.AppendOriginObject(origin)
	example159Type.AppendTargetObject(target)
}
