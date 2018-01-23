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
	example1Type.AddType("Object")
	example1Type.SetId(MustParseURL("http://www.test.example/object/1"))
	example1Type.AddNameString("A Simple, non-specific object")
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
	example2Type.AddType("Link")
	example2Type.SetHref(href)
	example2Type.SetHreflang("en")
	example2Type.SetMediaType("text/html")
	example2Type.AddNameString("An example link")
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Note{}
	object.AddType("Note")
	object.AddNameString("A Note")
	example3Type.AddType("Activity")
	example3Type.AddSummaryString("Sally did something to a note")
	example3Type.AddActorObject(actor)
	example3Type.AddObject(object)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	place := &Place{}
	place.AddType("Place")
	place.AddNameString("Work")
	example4Type.AddType("Travel")
	example4Type.AddSummaryString("Sally went to work")
	example4Type.AddActorObject(person)
	example4Type.AddTargetObject(place)
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
	note1.AddType("Note")
	note1.AddNameString("A Simple Note")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Another Simple Note")
	example5Type.AddType("Collection")
	example5Type.AddSummaryString("Sally's notes")
	example5Type.SetTotalItems(2)
	example5Type.AddItemsObject(note1)
	example5Type.AddItemsObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("A Simple Note")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Another Simple Note")
	example6Type.AddType("OrderedCollection")
	example6Type.AddSummaryString("Sally's notes")
	example6Type.SetTotalItems(2)
	example6Type.AddOrderedItemsObject(note1)
	example6Type.AddOrderedItemsObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("A Simple Note")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Another Simple Note")
	link := MustParseURL("http://example.org/foo")
	example7Type.AddType("CollectionPage")
	example7Type.AddSummaryString("Page 1 of Sally's notes")
	example7Type.SetId(MustParseURL("http://example.org/foo?page=1"))
	example7Type.SetPartOfIRI(link)
	example7Type.AddItemsObject(note1)
	example7Type.AddItemsObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("A Simple Note")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Another Simple Note")
	link := MustParseURL("http://example.org/foo")
	example8Type.AddType("OrderedCollectionPage")
	example8Type.AddSummaryString("Page 1 of Sally's notes")
	example8Type.SetId(MustParseURL("http://example.org/foo?page=1"))
	example8Type.SetPartOfIRI(link)
	example8Type.AddOrderedItemsObject(note1)
	example8Type.AddOrderedItemsObject(note2)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	event := &Event{}
	event.AddType("Event")
	event.AddNameString("Going-Away Party for Jim")
	link := MustParseURL("http://john.example.org")
	invite := &Invite{}
	invite.AddType("Invite")
	invite.AddActorIRI(link)
	invite.AddObject(event)
	example9Type.AddType("Accept")
	example9Type.AddSummaryString("Sally accepted an invitation to a party")
	example9Type.AddActorObject(person)
	example9Type.AddObject(invite)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	object := &Person{}
	object.AddType("Person")
	object.AddNameString("Joe")
	target := &Group{}
	target.AddType("Group")
	target.AddNameString("The Club")
	example10Type.AddType("Accept")
	example10Type.AddSummaryString("Sally accepted Joe into the club")
	example10Type.AddActorObject(person)
	example10Type.AddObject(object)
	example10Type.AddTargetObject(target)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	event := &Event{}
	event.AddType("Event")
	event.AddNameString("Going-Away Party for Jim")
	link := MustParseURL("http://john.example.org")
	invite := &Invite{}
	invite.AddType("Invite")
	invite.AddActorIRI(link)
	invite.AddObject(event)
	example11Type.AddType("TentativeAccept")
	example11Type.AddSummaryString("Sally tentatively accepted an invitation to a party")
	example11Type.AddActorObject(person)
	example11Type.AddObject(invite)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	link := MustParseURL("http://example.org/abc")
	example12Type.AddType("Add")
	example12Type.AddSummaryString("Sally added an object")
	example12Type.AddActorObject(person)
	example12Type.AddObjectIRI(link)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	link := MustParseURL("http://example.org/img/cat.png")
	object := &Image{}
	object.AddType("Image")
	object.AddNameString("A picture of my cat")
	object.AddUrlAnyURI(link)
	origin := &Collection{}
	origin.AddType("Collection")
	origin.AddNameString("Camera Roll")
	target := &Collection{}
	target.AddType("Collection")
	target.AddNameString("My Cat Pictures")
	example13Type.AddType("Add")
	example13Type.AddSummaryString("Sally added a picture of her cat to her cat picture collection")
	example13Type.AddActorObject(person)
	example13Type.AddObject(object)
	example13Type.AddOriginObject(origin)
	example13Type.AddTargetObject(target)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	location := &Place{}
	location.AddType("Place")
	location.AddNameString("Work")
	origin := &Place{}
	origin.AddType("Place")
	origin.AddNameString("Home")
	example14Type.AddType("Arrive")
	example14Type.AddSummaryString("Sally arrived at work")
	example14Type.AddActorObject(person)
	example14Type.AddLocationObject(location)
	example14Type.AddOriginObject(origin)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Note{}
	object.AddType("Note")
	object.AddNameString("A Simple Note")
	object.AddContentString("This is a simple note")
	example15Type.AddType("Create")
	example15Type.AddSummaryString("Sally created a note")
	example15Type.AddActorObject(actor)
	example15Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	origin := &Collection{}
	origin.AddType("Collection")
	origin.AddNameString("Sally's Notes")
	link := MustParseURL("http://example.org/notes/1")
	example16Type.AddType("Delete")
	example16Type.AddSummaryString("Sally deleted a note")
	example16Type.AddActorObject(actor)
	example16Type.AddObjectIRI(link)
	example16Type.AddOriginObject(origin)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Person{}
	object.AddType("Person")
	object.AddNameString("John")
	example17Type.AddType("Follow")
	example17Type.AddSummaryString("Sally followed John")
	example17Type.AddActorObject(actor)
	example17Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example18Type.AddType("Ignore")
	example18Type.AddSummaryString("Sally ignored a note")
	example18Type.AddActorObject(actor)
	example18Type.AddObjectIRI(link)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Group{}
	object.AddType("Group")
	object.AddNameString("A Simple Group")
	example19Type.AddType("Join")
	example19Type.AddSummaryString("Sally joined a group")
	example19Type.AddActorObject(actor)
	example19Type.AddObject(object)

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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Place{}
	object.AddType("Place")
	object.AddNameString("Work")
	example20Type.AddType("Leave")
	example20Type.AddSummaryString("Sally left work")
	example20Type.AddActorObject(actor)
	example20Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Group{}
	object.AddType("Group")
	object.AddNameString("A Simple Group")
	example21Type.AddType("Leave")
	example21Type.AddSummaryString("Sally left a group")
	example21Type.AddActorObject(actor)
	example21Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example22Type.AddType("Like")
	example22Type.AddSummaryString("Sally liked a note")
	example22Type.AddActorObject(actor)
	example22Type.AddObjectIRI(link)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := make(map[string]interface{})
	object["type"] = "http://www.types.example/ProductOffer"
	object["name"] = "50% Off!"
	target := &Person{}
	target.AddType("Person")
	target.AddNameString("Lewis")
	example23Type.AddType("Offer")
	example23Type.AddSummaryString("Sally offered 50% off to Lewis")
	example23Type.AddActorObject(actor)
	example23Type.SetUnknownObject(object)
	example23Type.AddTargetObject(target)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Event{}
	object.AddType("Event")
	object.AddNameString("A Party")
	target1 := &Person{}
	target1.AddType("Person")
	target1.AddNameString("John")
	target2 := &Person{}
	target2.AddType("Person")
	target2.AddNameString("Lisa")
	example24Type.AddType("Invite")
	example24Type.AddSummaryString("Sally invited John and Lisa to a party")
	example24Type.AddActorObject(actor)
	example24Type.AddObject(object)
	example24Type.AddTargetObject(target1)
	example24Type.AddTargetObject(target2)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://john.example.org")
	inviteObject := &Event{}
	inviteObject.AddType("Event")
	inviteObject.AddNameString("Going-Away Party for Jim")
	object := &Invite{}
	object.AddType("Invite")
	object.AddActorIRI(link)
	object.AddObject(inviteObject)
	example25Type.AddType("Reject")
	example25Type.AddSummaryString("Sally rejected an invitation to a party")
	example25Type.AddActorObject(actor)
	example25Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://john.example.org")
	inviteObject := &Event{}
	inviteObject.AddType("Event")
	inviteObject.AddNameString("Going-Away Party for Jim")
	object := &Invite{}
	object.AddType("Invite")
	object.AddActorIRI(link)
	object.AddObject(inviteObject)
	example26Type.AddType("TentativeReject")
	example26Type.AddSummaryString("Sally tentatively rejected an invitation to a party")
	example26Type.AddActorObject(actor)
	example26Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	target := &Collection{}
	target.AddType("Collection")
	target.AddNameString("Notes Folder")
	example27Type.AddType("Remove")
	example27Type.AddSummaryString("Sally removed a note from her notes folder")
	example27Type.AddActorObject(actor)
	example27Type.AddObjectIRI(link)
	example27Type.AddTargetObject(target)
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
	object.AddType("Person")
	object.AddNameString("Sally")
	origin := &Group{}
	origin.AddType("Group")
	origin.AddNameString("A Simple Group")
	example28Type.AddType("Remove")
	example28Type.AddSummaryString("The moderator removed Sally from a group")
	example28Type.SetUnknownActor(actor)
	example28Type.AddObject(object)
	example28Type.AddOriginObject(origin)
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
	object.AddType("Offer")
	object.AddActorIRI(link)
	object.AddObjectIRI(objectLink)
	object.AddTargetIRI(targetLink)
	example29Type.AddType("Undo")
	example29Type.AddSummaryString("Sally retracted her offer to John")
	example29Type.AddActorIRI(link)
	example29Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/notes/1")
	example30Type.AddType("Update")
	example30Type.AddSummaryString("Sally updated her note")
	example30Type.AddActorObject(actor)
	example30Type.AddObjectIRI(link)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	object := &Article{}
	object.AddType("Article")
	object.AddNameString("What You Should Know About Activity Streams")
	example31Type.AddType("View")
	example31Type.AddSummaryString("Sally read an article")
	example31Type.AddActorObject(actor)
	example31Type.AddObject(object)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/music.mp3")
	example32Type.AddType("Listen")
	example32Type.AddSummaryString("Sally listened to a piece of music")
	example32Type.AddActorObject(actor)
	example32Type.AddObjectIRI(link)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/posts/1")
	example33Type.AddType("Read")
	example33Type.AddSummaryString("Sally read a blog post")
	example33Type.AddActorObject(actor)
	example33Type.AddObjectIRI(link)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	link := MustParseURL("http://example.org/posts/1")
	target := &Collection{}
	target.AddType("Collection")
	target.AddNameString("List B")
	origin := &Collection{}
	origin.AddType("Collection")
	origin.AddNameString("List A")
	example34Type.AddType("Move")
	example34Type.AddSummaryString("Sally moved a post from List A to List B")
	example34Type.AddActorObject(actor)
	example34Type.AddObjectIRI(link)
	example34Type.AddTargetObject(target)
	example34Type.AddOriginObject(origin)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	target := &Place{}
	target.AddType("Place")
	target.AddNameString("Home")
	origin := &Place{}
	origin.AddType("Place")
	origin.AddNameString("Work")
	example35Type.AddType("Travel")
	example35Type.AddSummaryString("Sally went home from work")
	example35Type.AddActorObject(actor)
	example35Type.AddTargetObject(target)
	example35Type.AddOriginObject(origin)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	actor.SetId(link)
	loc := &Place{}
	loc.AddType("Place")
	loc.AddNameString("Work")
	object := &Arrive{}
	object.AddType("Arrive")
	object.AddActorIRI(link)
	object.AddLocationObject(loc)
	example36Type.AddType("Announce")
	example36Type.AddSummaryString("Sally announced that she had arrived at work")
	example36Type.AddActorObject(actor)
	example36Type.AddObject(object)
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
	example37Type.AddType("Block")
	example37Type.AddSummaryString("Sally blocked Joe")
	example37Type.AddActorIRI(link)
	example37Type.AddObjectIRI(objLink)
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
	object.AddType("Note")
	object.AddContentString("An inappropriate note")
	example38Type.AddType("Flag")
	example38Type.AddSummaryString("Sally flagged an inappropriate note")
	example38Type.AddActorIRI(link)
	example38Type.AddObject(object)
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
	example39Type.AddType("Dislike")
	example39Type.AddSummaryString("Sally disliked a post")
	example39Type.AddActorIRI(link)
	example39Type.AddObjectIRI(objLink)
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
	note1.AddType("Note")
	note1.AddNameString("Option A")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Option B")
	example40Type.AddType("Question")
	example40Type.AddNameString("What is the answer?")
	example40Type.AddOneOfObject(note1)
	example40Type.AddOneOfObject(note2)
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
	example41Type.AddType("Question")
	example41Type.AddNameString("What is the answer?")
	example41Type.AddClosedDateTime(t)
}

const example42 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Application",
  "name": "Exampletron 3000"
}`

var example42Type = &Application{}

func init() {
	example42Type.AddType("Application")
	example42Type.AddNameString("Exampletron 3000")
}

const example43 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Group",
  "name": "Big Beards of Austin"
}`

var example43Type = &Group{}

func init() {
	example43Type.AddType("Group")
	example43Type.AddNameString("Big Beards of Austin")
}

const example44 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Organization",
  "name": "Example Co."
}`

var example44Type = &Organization{}

func init() {
	example44Type.AddType("Organization")
	example44Type.AddNameString("Example Co.")
}

const example45 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally Smith"
}`

var example45Type = &Person{}

func init() {
	example45Type.AddType("Person")
	example45Type.AddNameString("Sally Smith")
}

const example46 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Service",
  "name": "Acme Web Service"
}`

var example46Type = &Service{}

func init() {
	example46Type.AddType("Service")
	example46Type.AddNameString("Acme Web Service")
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
	subject.AddType("Person")
	subject.AddNameString("Sally")
	object := &Person{}
	object.AddType("Person")
	object.AddNameString("John")
	rel := MustParseURL("http://purl.org/vocab/relationship/acquaintanceOf")
	example47Type.AddType("Relationship")
	example47Type.AddSummaryString("Sally is an acquaintance of John")
	example47Type.SetSubjectObject(subject)
	example47Type.AddObject(object)
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
	example48Type.AddType("Article")
	example48Type.AddNameString("What a Crazy Day I Had")
	example48Type.AddAttributedToIRI(att)
	example48Type.AddContentString("<div>... you will never believe ...</div>")
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
	example49Type.AddType("Document")
	example49Type.AddNameString("4Q Sales Forecast")
	example49Type.AddUrlAnyURI(l)
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
	link.AddType("Link")
	link.SetHref(l)
	link.SetMediaType("audio/mp3")
	example50Type.AddType("Audio")
	example50Type.AddNameString("Interview With A Famous Technologist")
	example50Type.AddUrlLink(link)
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
	link1.AddType("Link")
	link1.SetHref(l1)
	link1.SetMediaType("image/jpeg")
	link2 := &Link{}
	link2.AddType("Link")
	link2.SetHref(l2)
	link2.SetMediaType("image/png")
	example51Type.AddType("Image")
	example51Type.AddNameString("Cat Jumping on Wagon")
	example51Type.AddUrlLink(link1)
	example51Type.AddUrlLink(link2)
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
	example52Type.AddType("Video")
	example52Type.AddNameString("Puppy Plays With Ball")
	example52Type.AddUrlAnyURI(l)
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
	example53Type.AddType("Note")
	example53Type.AddNameString("A Word of Warning")
	example53Type.AddContentString("Looks like it is going to rain today. Bring an umbrella!")
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
	example54Type.AddType("Page")
	example54Type.AddNameString("Omaha Weather Report")
	example54Type.AddUrlAnyURI(l)
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
	example55Type.AddType("Event")
	example55Type.AddNameString("Going-Away Party for Jim")
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
	example56Type.AddType("Place")
	example56Type.AddNameString("Work")
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
	example57Type.AddType("Place")
	example57Type.AddNameString("Fresno Area")
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
	example58Type.AddType("Mention")
	example58Type.AddSummaryString("Mention of Joe by Carrie in her note")
	example58Type.SetHref(l)
	example58Type.AddNameString("Joe")
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
	person.AddType("Person")
	person.AddNameString("Sally Smith")
	example59Type.AddType("Profile")
	example59Type.AddSummaryString("Sally's Profile")
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
	image1.AddType("Image")
	image1.SetId(MustParseURL("http://image.example/1"))
	tombstone := &Tombstone{}
	tombstone.AddType("Tombstone")
	tombstone.AddFormerTypeString("Image")
	tombstone.SetId(MustParseURL("http://image.example/2"))
	tombstone.SetDeleted(t)
	image2 := &Image{}
	image2.AddType("Image")
	image2.SetId(MustParseURL("http://image.example/3"))
	example60Type.AddType("OrderedCollection")
	example60Type.SetTotalItems(3)
	example60Type.AddNameString("Vacation photos 2016")
	example60Type.AddOrderedItemsObject(image1)
	example60Type.AddOrderedItemsObject(tombstone)
	example60Type.AddOrderedItemsObject(image2)
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
	example63Type.AddType("Offer")
	example63Type.AddSummaryString("Sally offered the Foo object")
	example63Type.AddActorIRI(l)
	example63Type.AddObjectIRI(o)
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
	actor.AddType("Person")
	actor.SetId(MustParseURL("http://sally.example.org"))
	actor.AddSummaryString("Sally")
	o := MustParseURL("http://example.org/foo")
	example64Type.AddType("Offer")
	example64Type.AddSummaryString("Sally offered the Foo object")
	example64Type.AddActorObject(actor)
	example64Type.AddObjectIRI(o)
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
	actor.AddType("Person")
	actor.SetId(MustParseURL("http://sally.example.org"))
	actor.AddNameString("Sally")
	o := MustParseURL("http://example.org/foo")
	l := MustParseURL("http://joe.example.org")
	example65Type.AddType("Offer")
	example65Type.AddSummaryString("Sally and Joe offered the Foo object")
	example65Type.AddObjectIRI(o)
	example65Type.AddActorIRI(l)
	example65Type.AddActorObject(actor)
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
	image.AddType("Image")
	image.AddContentString("This is what he looks like.")
	image.AddUrlAnyURI(l)
	example66Type.AddType("Note")
	example66Type.AddNameString("Have you seen my cat?")
	example66Type.AddAttachmentObject(image)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	example67Type.AddType("Image")
	example67Type.AddNameString("My cat taking a nap")
	example67Type.AddUrlAnyURI(l)
	example67Type.AddAttributedToObject(person)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	example68Type.AddType("Image")
	example68Type.AddNameString("My cat taking a nap")
	example68Type.AddUrlAnyURI(l)
	example68Type.AddAttributedToIRI(a)
	example68Type.AddAttributedToObject(person)
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
	example69Type.AddType("Note")
	example69Type.AddNameString("Holiday announcement")
	example69Type.AddContentString("Thursday will be a company-wide holiday. Enjoy your day off!")
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
	example70Type.AddType("Offer")
	example70Type.AddSummaryString("Sally offered a post to John")
	example70Type.AddActorIRI(a)
	example70Type.AddObjectIRI(o)
	example70Type.AddTargetIRI(t)
	example70Type.AddBccIRI(b)
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
	example71Type.AddType("Offer")
	example71Type.AddSummaryString("Sally offered a post to John")
	example71Type.AddActorIRI(a)
	example71Type.AddObjectIRI(o)
	example71Type.AddTargetIRI(t)
	example71Type.AddBtoIRI(b)
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
	example72Type.AddType("Offer")
	example72Type.AddSummaryString("Sally offered a post to John")
	example72Type.AddActorIRI(a)
	example72Type.AddObjectIRI(o)
	example72Type.AddTargetIRI(t)
	example72Type.AddCcIRI(b)
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
	offer.AddType("Offer")
	offer.AddActorIRI(oa)
	offer.AddObjectIRI(oo)
	offer.AddTargetIRI(ot)
	offer.AddContextIRI(oc)
	la := MustParseURL("http://joe.example.org")
	lo := MustParseURL("http://example.org/posts/2")
	lc := MustParseURL("http://example.org/contexts/1")
	like := &Like{}
	like.AddType("Like")
	like.AddActorIRI(la)
	like.AddObjectIRI(lo)
	like.AddContextIRI(lc)
	example73Type.AddType("Collection")
	example73Type.AddSummaryString("Activities in context 1")
	example73Type.AddItemsObject(offer)
	example73Type.AddItemsObject(like)
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
	example74Type.AddType("Collection")
	example74Type.SetTotalItems(3)
	example74Type.AddSummaryString("Sally's blog posts")
	example74Type.SetCurrentIRI(c)
	example74Type.AddItemsIRI(i1)
	example74Type.AddItemsIRI(i2)
	example74Type.AddItemsIRI(i3)
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
	link.AddType("Link")
	link.AddSummaryString("Most Recent Items")
	link.SetHref(href)
	example75Type.AddType("Collection")
	example75Type.SetTotalItems(3)
	example75Type.AddSummaryString("Sally's blog posts")
	example75Type.SetCurrentLink(link)
	example75Type.AddItemsIRI(i1)
	example75Type.AddItemsIRI(i2)
	example75Type.AddItemsIRI(i3)
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
	example76Type.AddType("Collection")
	example76Type.SetTotalItems(3)
	example76Type.AddSummaryString("Sally's blog posts")
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
	link.AddType("Link")
	link.AddSummaryString("First Page")
	link.SetHref(href)
	example77Type.AddType("Collection")
	example77Type.SetTotalItems(3)
	example77Type.AddSummaryString("Sally's blog posts")
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
	app.AddType("Application")
	app.AddNameString("Exampletron 3000")
	example78Type.AddType("Note")
	example78Type.AddSummaryString("A simple note")
	example78Type.AddContentString("This is all there is.")
	example78Type.AddGeneratorObject(app)
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
	image.AddType("Image")
	image.AddNameString("Note icon")
	image.AddUrlAnyURI(u)
	image.SetWidth(16)
	image.SetHeight(16)
	example79Type.AddType("Note")
	example79Type.AddSummaryString("A simple note")
	example79Type.AddContentString("This is all there is.")
	example79Type.AddIconImage(image)
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
	image1.AddType("Image")
	image1.AddSummaryString("Note (16x16)")
	image1.AddUrlAnyURI(u1)
	image1.SetWidth(16)
	image1.SetHeight(16)
	image2 := &Image{}
	image2.AddType("Image")
	image2.AddSummaryString("Note (32x32)")
	image2.AddUrlAnyURI(u2)
	image2.SetWidth(32)
	image2.SetHeight(32)
	example80Type.AddType("Note")
	example80Type.AddSummaryString("A simple note")
	example80Type.AddContentString("A simple note")
	example80Type.AddIconImage(image1)
	example80Type.AddIconImage(image2)
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
	image.AddType("Image")
	image.AddNameString("A Cat")
	image.AddUrlAnyURI(u)
	example81Type.AddType("Note")
	example81Type.AddNameString("A simple note")
	example81Type.AddContentString("This is all there is.")
	example81Type.AddImageImage(image)
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
	image1.AddType("Image")
	image1.AddNameString("Cat 1")
	image1.AddUrlAnyURI(u1)
	image2 := &Image{}
	image2.AddType("Image")
	image2.AddNameString("Cat 2")
	image2.AddUrlAnyURI(u2)
	example82Type.AddType("Note")
	example82Type.AddNameString("A simple note")
	example82Type.AddContentString("This is all there is.")
	example82Type.AddImageImage(image1)
	example82Type.AddImageImage(image2)
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
	note.AddType("Note")
	note.AddSummaryString("Previous note")
	note.AddContentString("What else is there?")
	example83Type.AddType("Note")
	example83Type.AddSummaryString("A simple note")
	example83Type.AddContentString("This is all there is.")
	example83Type.AddInReplyToObject(note)
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
	example84Type.AddType("Note")
	example84Type.AddSummaryString("A simple note")
	example84Type.AddContentString("This is all there is.")
	example84Type.AddInReplyToIRI(u)
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
	actor.AddType("Person")
	actor.AddNameString("Sally")
	service := &Service{}
	service.AddType("Service")
	service.AddNameString("Acme Music Service")
	example85Type.AddType("Listen")
	example85Type.AddSummaryString("Sally listened to a piece of music on the Acme Music Service")
	example85Type.AddActorObject(actor)
	example85Type.AddObjectIRI(u)
	example85Type.AddInstrumentObject(service)
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
	example86Type.AddType("Collection")
	example86Type.SetTotalItems(3)
	example86Type.AddSummaryString("A collection")
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
	link.AddType("Link")
	link.AddSummaryString("Last Page")
	link.SetHref(u)
	example87Type.AddType("Collection")
	example87Type.AddSummaryString("A collection")
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
	place.AddType("Place")
	place.AddNameString("Over the Arabian Sea, east of Socotra Island Nature Sanctuary")
	place.SetLongitude(12.34)
	place.SetLatitude(56.78)
	place.SetAltitude(90)
	place.SetUnitsUnitsValue("m")
	example88Type.AddType("Person")
	example88Type.AddNameString("Sally")
	example88Type.AddLocationObject(place)
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
	note1.AddType("Note")
	note1.AddNameString("Reminder for Going-Away Party")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Meeting 2016-11-17")
	example89Type.AddType("Collection")
	example89Type.AddSummaryString("Sally's notes")
	example89Type.SetTotalItems(2)
	example89Type.AddItemsObject(note1)
	example89Type.AddItemsObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("Meeting 2016-11-17")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Reminder for Going-Away Party")
	example90Type.AddType("OrderedCollection")
	example90Type.AddSummaryString("Sally's notes")
	example90Type.SetTotalItems(2)
	example90Type.AddOrderedItemsObject(note1)
	example90Type.AddOrderedItemsObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("Option A")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Option B")
	example91Type.AddType("Question")
	example91Type.AddNameString("What is the answer?")
	example91Type.AddOneOfObject(note1)
	example91Type.AddOneOfObject(note2)
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
	note1.AddType("Note")
	note1.AddNameString("Option A")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Option B")
	example92Type.AddType("Question")
	example92Type.AddNameString("What is the answer?")
	example92Type.AddAnyOfObject(note1)
	example92Type.AddAnyOfObject(note2)
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
	example93Type.AddType("Question")
	example93Type.AddNameString("What is the answer?")
	example93Type.AddClosedDateTime(t)
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
	target.AddType("Collection")
	target.AddNameString("List B")
	origin := &Collection{}
	origin.AddType("Collection")
	origin.AddNameString("List A")
	example94Type.AddType("Move")
	example94Type.AddSummaryString("Sally moved a post from List A to List B")
	example94Type.AddActorIRI(a)
	example94Type.AddObjectIRI(o)
	example94Type.AddTargetObject(target)
	example94Type.AddOriginObject(origin)
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
	example95Type.AddType("CollectionPage")
	example95Type.AddSummaryString("Page 2 of Sally's blog posts")
	example95Type.SetNextIRI(i)
	example95Type.AddItemsIRI(u1)
	example95Type.AddItemsIRI(u2)
	example95Type.AddItemsIRI(u3)
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
	link.AddType("Link")
	link.AddNameString("Next Page")
	link.SetHref(href)
	example96Type.AddType("CollectionPage")
	example96Type.AddSummaryString("Page 2 of Sally's blog posts")
	example96Type.SetNextLink(link)
	example96Type.AddItemsIRI(u1)
	example96Type.AddItemsIRI(u2)
	example96Type.AddItemsIRI(u3)
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
	example97Type.AddType("Like")
	example97Type.AddSummaryString("Sally liked a post")
	example97Type.AddActorIRI(a)
	example97Type.AddObjectIRI(o)
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
	note.AddType("Note")
	note.AddContentString("A simple note")
	example98Type.AddType("Like")
	example98Type.AddActorIRI(a)
	example98Type.AddObject(note)
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
	note.AddType("Note")
	note.AddSummaryString("A simple note")
	note.AddContentString("That is a tree.")
	example99Type.AddType("Like")
	example99Type.AddSummaryString("Sally liked a note")
	example99Type.AddActorIRI(a)
	example99Type.AddObjectIRI(o)
	example99Type.AddObject(note)
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
	example100Type.AddType("CollectionPage")
	example100Type.AddSummaryString("Page 1 of Sally's blog posts")
	example100Type.SetPrevIRI(p)
	example100Type.AddItemsIRI(u1)
	example100Type.AddItemsIRI(u2)
	example100Type.AddItemsIRI(u3)
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
	link.AddType("Link")
	link.AddNameString("Previous Page")
	link.SetHref(p)
	example101Type.AddType("CollectionPage")
	example101Type.AddSummaryString("Page 1 of Sally's blog posts")
	example101Type.SetPrevLink(link)
	example101Type.AddItemsIRI(u1)
	example101Type.AddItemsIRI(u2)
	example101Type.AddItemsIRI(u3)
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
	link.AddType("Link")
	link.SetMediaType("video/mkv")
	link.SetHref(u)
	video := &Video{}
	video.AddType("Video")
	video.AddNameString("Trailer")
	video.SetDuration(time.Minute)
	video.AddUrlLink(link)
	example102Type.AddType("Video")
	example102Type.AddNameString("Cool New Movie")
	example102Type.SetDuration(time.Hour*2 + time.Minute*30)
	example102Type.AddPreviewObject(video)
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
	example103Type.AddType("Activity")
	example103Type.AddType("http://www.verbs.example/Check")
	example103Type.AddSummaryString("Sally checked that her flight was on time")
	example103Type.AddActorIRI(a)
	example103Type.AddObjectIRI(o)
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
	note.AddType("Note")
	note.AddSummaryString("A response to the note")
	note.AddContentString("I am glad to hear it.")
	note.AddInReplyToIRI(i)
	replies := &Collection{}
	replies.AddType("Collection")
	replies.SetTotalItems(1)
	replies.AddItemsObject(note)
	example104Type.AddType("Note")
	example104Type.AddSummaryString("A simple note")
	example104Type.SetId(MustParseURL("http://www.test.example/notes/1"))
	example104Type.AddContentString("I am fine.")
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
	person.AddType("Person")
	person.SetId(MustParseURL("http://sally.example.org"))
	person.AddNameString("Sally")
	example105Type.AddType("Image")
	example105Type.AddSummaryString("Picture of Sally")
	example105Type.AddUrlAnyURI(u)
	example105Type.AddTagObject(person)
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
	example106Type.AddType("Offer")
	example106Type.AddSummaryString("Sally offered the post to John")
	example106Type.AddActorIRI(a)
	example106Type.AddObjectIRI(o)
	example106Type.AddTargetIRI(t)
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
	person.AddType("Person")
	person.AddNameString("John")
	example107Type.AddType("Offer")
	example107Type.AddSummaryString("Sally offered the post to John")
	example107Type.AddActorIRI(a)
	example107Type.AddObjectIRI(o)
	example107Type.AddTargetObject(person)
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
	example108Type.AddType("Offer")
	example108Type.AddSummaryString("Sally offered the post to John")
	example108Type.AddActorIRI(a)
	example108Type.AddObjectIRI(o)
	example108Type.AddTargetIRI(t)
	example108Type.AddToIRI(z)
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
	example109Type.AddType("Document")
	example109Type.AddNameString("4Q Sales Forecast")
	example109Type.AddUrlAnyURI(u)
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
	link.AddType("Link")
	link.SetHref(u)
	example110Type.AddType("Document")
	example110Type.AddNameString("4Q Sales Forecast")
	example110Type.AddUrlLink(link)
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
	link1.AddType("Link")
	link1.SetHref(u1)
	link1.SetMediaType("application/pdf")
	link2 := &Link{}
	link2.AddType("Link")
	link2.SetHref(u2)
	link2.SetMediaType("text/html")
	example111Type.AddType("Document")
	example111Type.AddNameString("4Q Sales Forecast")
	example111Type.AddUrlLink(link1)
	example111Type.AddUrlLink(link2)
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
	example112Type.AddType("Place")
	example112Type.AddNameString("Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China")
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
	example113Type.AddType("Place")
	example113Type.AddNameString("Fresno Area")
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
	example114Type.AddType("Note")
	example114Type.AddSummaryString("A simple note")
	example114Type.AddContentString("A <em>simple</em> note")
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
	example115Type.AddType("Note")
	example115Type.AddSummaryString("A simple note")
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
	example116Type.AddType("Note")
	example116Type.AddSummaryString("A simple note")
	example116Type.SetMediaType("text/markdown")
	example116Type.AddContentString("## A simple note\nA simple markdown `note`")
}

const example117 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A simple note"
}`

var example117Type = &Note{}

func init() {
	example117Type.AddType("Note")
	example117Type.AddNameString("A simple note")
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
	example118Type.AddType("Note")
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
	example119Type.AddType("Video")
	example119Type.AddNameString("Birds Flying")
	example119Type.AddUrlAnyURI(u)
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
	example120Type.AddType("Link")
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
	example121Type.AddType("Link")
	example121Type.SetHref(u)
	example121Type.SetMediaType("text/html")
	example121Type.AddNameString("Previous")
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
	example122Type.AddType("Link")
	example122Type.SetHref(u)
	example122Type.SetMediaType("text/html")
	example122Type.AddNameString("Previous")
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
	note1.AddType("Note")
	note1.AddNameString("Pizza Toppings to Try")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Thought about California")
	example123Type.AddType("CollectionPage")
	example123Type.AddSummaryString("Page 1 of Sally's notes")
	example123Type.SetId(MustParseURL("http://example.org/collection?page=1"))
	example123Type.SetPartOfIRI(u)
	example123Type.AddItemsObject(note1)
	example123Type.AddItemsObject(note2)
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
	example124Type.AddType("Place")
	example124Type.AddNameString("Fresno Area")
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
	example125Type.AddType("Place")
	example125Type.AddNameString("Fresno Area")
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
	example126Type.AddType("Link")
	example126Type.SetHref(u)
	example126Type.SetHreflang("en")
	example126Type.SetMediaType("text/html")
	example126Type.AddNameString("Next")
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
	example127Type.AddType("Event")
	example127Type.AddNameString("Going-Away Party for Jim")
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
	example128Type.AddType("Note")
	example128Type.AddSummaryString("A simple note")
	example128Type.AddContentString("Fish swim.")
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
	example129Type.AddType("Event")
	example129Type.AddNameString("Going-Away Party for Jim")
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
	example130Type.AddType("Place")
	example130Type.AddNameString("Fresno Area")
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
	example131Type.AddType("Link")
	example131Type.SetHref(u)
	example131Type.SetHreflang("en")
	example131Type.SetMediaType("text/html")
	example131Type.AddNameString("Preview")
	example131Type.AddRel("canonical")
	example131Type.AddRel("preview")
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
	note1.AddType("Note")
	note1.AddNameString("Density of Water")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Air Mattress Idea")
	example132Type.AddType("OrderedCollectionPage")
	example132Type.AddSummaryString("Page 1 of Sally's notes")
	example132Type.SetStartIndex(0)
	example132Type.AddOrderedItemsObject(note1)
	example132Type.AddOrderedItemsObject(note2)
}

const example133 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "Cane Sugar Processing",
  "type": "Note",
  "summary": "A simple <em>note</em>"
}`

var example133Type = &Note{}

func init() {
	example133Type.AddType("Note")
	example133Type.AddNameString("Cane Sugar Processing")
	example133Type.AddSummaryString("A simple <em>note</em>")
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
	example134Type.AddType("Note")
	example134Type.AddNameString("Cane Sugar Processing")
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
	note1.AddType("Note")
	note1.AddNameString("Which Staircase Should I Use")
	note2 := &Note{}
	note2.AddType("Note")
	note2.AddNameString("Something to Remember")
	example135Type.AddType("Collection")
	example135Type.AddSummaryString("Sally's notes")
	example135Type.SetTotalItems(2)
	example135Type.AddItemsObject(note1)
	example135Type.AddItemsObject(note2)
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
	example136Type.AddType("Place")
	example136Type.AddNameString("Fresno Area")
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
	example137Type.AddType("Note")
	example137Type.AddNameString("Cranberry Sauce Idea")
	example137Type.AddContentString("Mush it up so it does not have the same shape as the can.")
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
	example138Type.AddType("Link")
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
	subject.AddType("Person")
	subject.AddNameString("Sally")
	object := &Person{}
	object.AddType("Person")
	object.AddNameString("John")
	example139Type.AddType("Relationship")
	example139Type.AddSummaryString("Sally is an acquaintance of John's")
	example139Type.SetSubjectObject(subject)
	example139Type.AddObject(object)
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
	subject.AddType("Person")
	subject.AddNameString("Sally")
	object := &Person{}
	object.AddType("Person")
	object.AddNameString("John")
	example140Type.AddType("Relationship")
	example140Type.AddSummaryString("Sally is an acquaintance of John's")
	example140Type.SetSubjectObject(subject)
	example140Type.AddObject(object)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	example141Type.AddType("Profile")
	example141Type.AddSummaryString("Sally's profile")
	example141Type.AddUrlAnyURI(u)
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
	example142Type.AddType("Tombstone")
	example142Type.AddSummaryString("This image has been deleted")
	example142Type.AddFormerTypeString("Image")
	example142Type.AddUrlAnyURI(u)
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
	example143Type.AddType("Tombstone")
	example143Type.AddSummaryString("This image has been deleted")
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
	audience.AddType("Group")
	audience.AddNameString("Project XYZ Working Group")
	note := &Note{}
	note.AddType("Note")
	note.AddSummaryString("A note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AddContentString("A note")
	create := &Create{}
	create.AddType("Create")
	create.AddSummaryString("Sally created a note")
	create.SetId(MustParseURL("http://activities.example.com/1"))
	create.AddActorIRI(sally)
	create.AddObject(note)
	create.SetUnknownContext(context)
	create.AddAudienceObject(audience)
	create.AddToIRI(john)
	like := &Like{}
	like.AddType("Like")
	like.AddSummaryString("John liked Sally's note")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AddActorIRI(john)
	like.AddObjectIRI(o)
	like.SetUnknownContext(context)
	like.AddAudienceObject(audience)
	like.AddToIRI(sally)
	example144Type.AddType("Collection")
	example144Type.AddSummaryString("Activities in Project XYZ")
	example144Type.AddItemsObject(create)
	example144Type.AddItemsObject(like)
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
	sally.AddType("Person")
	sally.AddNameString("Sally")
	jane := &Person{}
	jane.AddType("Person")
	jane.AddNameString("Jane")
	joe := &Person{}
	joe.AddType("Person")
	joe.AddNameString("Joe")
	joeRel := &Relationship{}
	joeRel.AddType("Relationship")
	joeRel.AddSummaryString("Sally is influenced by Joe")
	joeRel.SetSubjectObject(sally)
	joeRel.AddObject(joe)
	joeRel.SetRelationshipIRI(influenced)
	janeRel := &Relationship{}
	janeRel.AddType("Relationship")
	janeRel.AddSummaryString("Sally is a friend of Jane")
	janeRel.SetSubjectObject(sally)
	janeRel.AddObject(jane)
	janeRel.SetRelationshipIRI(friend)
	example145Type.AddType("Collection")
	example145Type.AddSummaryString("Sally's friends list")
	example145Type.AddItemsObject(joeRel)
	example145Type.AddItemsObject(janeRel)
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
	relationship.AddType("Relationship")
	relationship.SetSubjectIRI(s)
	relationship.SetRelationshipIRI(friend)
	relationship.AddObjectIRI(m)
	relationship.SetStartTime(t)
	example146Type.AddType("Create")
	example146Type.AddSummaryString("Sally became a friend of Matt")
	example146Type.AddActorIRI(s)
	example146Type.AddObject(relationship)
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
	rel.AddType("Relationship")
	rel.AddSummaryString("Sally and John's friendship")
	rel.SetId(MustParseURL("http://example.org/connections/123"))
	rel.SetSubjectIRI(s)
	rel.AddObjectIRI(t)
	rel.SetRelationshipIRI(friend)
	example147Type.AddType("Offer")
	example147Type.AddSummaryString("Sally requested to be a friend of John")
	example147Type.SetId(MustParseURL("http://example.org/connection-requests/123"))
	example147Type.AddActorIRI(s)
	example147Type.AddObject(rel)
	example147Type.AddTargetIRI(t)
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
	jc.AddType("Collection")
	jc.AddSummaryString("John's Connections")
	sc := &Collection{}
	sc.AddType("Collection")
	sc.AddSummaryString("Sally's Connections")
	o1 := &Accept{}
	o1.AddType("Accept")
	o1.SetId(MustParseURL("http://example.org/activities/122"))
	o1.AddSummaryString("John accepted Sally's friend request")
	o1.AddObjectIRI(req123)
	o1.AddInReplyToIRI(req123)
	o1.AddContextIRI(conn123)
	o1.AddResultIRI(a123)
	o1.AddResultIRI(a124)
	o1.AddResultIRI(a125)
	o1.AddResultIRI(a126)
	o1.AddActorIRI(john)
	o2 := &Follow{}
	o2.AddType("Follow")
	o2.SetId(MustParseURL("http://example.org/activities/123"))
	o2.AddActorIRI(john)
	o2.AddObjectIRI(sally)
	o2.AddContextIRI(conn123)
	o2.AddSummaryString("John followed Sally")
	o3 := &Follow{}
	o3.AddType("Follow")
	o3.SetId(MustParseURL("http://example.org/activities/124"))
	o3.AddActorIRI(sally)
	o3.AddObjectIRI(john)
	o3.AddContextIRI(conn123)
	o3.AddSummaryString("Sally followed John")
	o4 := &Add{}
	o4.AddType("Add")
	o4.SetId(MustParseURL("http://example.org/activities/125"))
	o4.AddSummaryString("John added Sally to his friends list")
	o4.AddActorIRI(john)
	o4.AddObjectIRI(conn123)
	o4.AddContextIRI(conn123)
	o4.AddTargetObject(jc)
	o5 := &Add{}
	o5.AddType("Add")
	o5.SetId(MustParseURL("http://example.org/activities/126"))
	o5.AddSummaryString("Sally added John to her friends list")
	o5.AddActorIRI(sally)
	o5.AddObjectIRI(conn123)
	o5.AddContextIRI(conn123)
	o5.AddTargetObject(sc)
	example148Type.AddType("Collection")
	example148Type.AddSummaryString("Sally and John's relationship history")
	example148Type.AddItemsObject(o1)
	example148Type.AddItemsObject(o2)
	example148Type.AddItemsObject(o3)
	example148Type.AddItemsObject(o4)
	example148Type.AddItemsObject(o5)
}

const example149 = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "San Francisco, CA"
}`

var example149Type = &Place{}

func init() {
	example149Type.AddType("Place")
	example149Type.AddNameString("San Francisco, CA")
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
	example150Type.AddType("Place")
	example150Type.AddNameString("San Francisco, CA")
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
	example151Type.AddType("Question")
	example151Type.AddNameString("A question about robots")
	example151Type.SetId(MustParseURL("http://help.example.org/question/1"))
	example151Type.AddContentString("I'd like to build a robot to feed my cat. Should I use Arduino or Raspberry Pi?")
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
	example152Type.AddType("Question")
	example152Type.AddNameString("A question about robots")
	example152Type.SetId(MustParseURL("http://polls.example.org/question/1"))
	example152Type.AddContentString("I'd like to build a robot to feed my cat. Which platform is best?")
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
	replies.AddType("Collection")
	replies.SetTotalItems(3)
	replies.SetUnknownItems(items)
	note := &Note{}
	note.AddType("Note")
	note.AddContentString("Users are favoriting &quot;arduino&quot; by a 33% margin.")
	example154Type.AddType("Question")
	example154Type.AddNameString("A question about robots")
	example154Type.SetId(MustParseURL("http://polls.example.org/question/1"))
	example154Type.AddContentString("I'd like to build a robot to feed my cat. Which platform is best?")
	example154Type.SetUnknownOneOf(oneOf)
	example154Type.SetReplies(replies)
	example154Type.AddResultObject(note)
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
	note.AddType("Note")
	note.AddSummaryString("John's note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AddContentString("My note")
	note.AddAttributedToIRI(john)
	like := &Like{}
	like.AddType("Like")
	like.AddSummaryString("Sally liked John's note")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AddActorIRI(sally)
	like.SetPublished(t1)
	like.AddObject(note)
	dislike := &Dislike{}
	dislike.AddType("Dislike")
	dislike.AddSummaryString("Sally disliked John's note")
	dislike.SetId(MustParseURL("http://activities.example.com/2"))
	dislike.AddActorIRI(sally)
	dislike.SetPublished(t2)
	dislike.AddObject(note)
	example155Type.AddType("Collection")
	example155Type.AddSummaryString("History of John's note")
	example155Type.AddItemsObject(like)
	example155Type.AddItemsObject(dislike)
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
	note.AddType("Note")
	note.SetId(MustParseURL("http://notes.example.com/1"))
	note.AddSummaryString("John's note")
	note.AddAttributedToIRI(john)
	note.AddContentString("My note")
	like := &Like{}
	like.AddType("Like")
	like.SetId(MustParseURL("http://activities.example.com/1"))
	like.AddSummaryString("Sally liked John's note")
	like.AddActorIRI(sally)
	like.SetPublished(t1)
	like.AddObject(note)
	undo := &Undo{}
	undo.AddType("Undo")
	undo.SetId(MustParseURL("http://activities.example.com/2"))
	undo.AddSummaryString("Sally no longer likes John's note")
	undo.AddActorIRI(sally)
	undo.SetPublished(t2)
	undo.AddObjectIRI(a)
	example156Type.AddType("Collection")
	example156Type.AddSummaryString("History of John's note")
	example156Type.AddItemsObject(like)
	example156Type.AddItemsObject(undo)
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
	person.AddType("Person")
	person.AddNameString("Sally")
	person.SetId(MustParseURL("http://sally.example.org"))
	example157Type.AddType("Note")
	example157Type.AddNameString("A thank-you note")
	example157Type.AddContentString("Thank you <a href='http://sally.example.org'>@sally</a> for all your hard work! <a href='http://example.org/tags/givingthanks'>#givingthanks</a>")
	example157Type.AddToObject(person)
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
	mention.AddType("Mention")
	mention.SetHref(u)
	mention.AddNameString("@sally")
	example158Type.AddType("Note")
	example158Type.AddNameString("A thank-you note")
	example158Type.AddContentString("Thank you @sally for all your hard work! #givingthanks")
	example158Type.AddTagLink(mention)
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
	obj.AddType("Document")
	obj.AddNameString("sales figures")
	origin := &Collection{}
	origin.AddType("Collection")
	origin.AddNameString("Folder A")
	target := &Collection{}
	target.AddType("Collection")
	target.AddNameString("Folder B")
	example159Type.AddType("Move")
	example159Type.AddSummaryString("Sally moved the sales figures from Folder A to Folder B")
	example159Type.AddActorIRI(sally)
	example159Type.AddObject(obj)
	example159Type.AddOriginObject(origin)
	example159Type.AddTargetObject(target)
}
