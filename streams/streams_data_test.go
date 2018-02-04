package streams

var repoExample1 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Object",
  "id": "http://www.test.example/object/1",
  "name": "A Simple, non-specific object"
}`
var repoExample3 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Activity",
  "name": "Sally did something to a note",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Note",
    "name": "A Note"
  }
}`
var repoExample5 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally's notes",
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
var repoExample7 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
      "name": "A Party!"
    }
  }
}`
var repoExample9 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally added an object",
  "type": "Add",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/abc"
}`
var repoExample11 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally arrived at work",
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
var repoExample13 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally deleted a note",
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
var repoExample15 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally followed John",
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
var repoExample17 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally joined a group",
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
var repoExample19 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally left a group",
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
var repoExample21 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered 50% off to Lewis",
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
var repoExample24 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally invited John and Lisa to a party",
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
var repoExample26 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally rejected an invitation to a party",
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
      "name": "A Party!"
    }
  }
}`
var repoExample28 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally removed a note from her notes folder",
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
var repoExample32 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally retracted her offer to John",
  "type": "Undo",
  "actor": "http://sally.example.org",
  "object": {
    "type": "Offer",
    "actor": "http://sally.example.org",
    "object": "http://example.org/posts/1",
    "target": "http://john.example.org"
  }
}`
var repoExample34 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Application",
  "name": "My Software Application."
}`
var repoExample37 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Group",
  "name": "A Simple Group."
}`
var repoExample39 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally Smith."
}`
var repoExample42 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Service",
  "name": "Acme Web Service"
}`
var repoExample48 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": "http://example.org/4q-sales-forecast.pdf"
}`
var repoExample50 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "A Simple Image",
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
var repoExample52 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A Short Note",
  "content": "This is a short note"
}`
var repoExample55 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
var repoExample57 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Work"
}`
var repoExample59 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered the Foo object",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/foo"
}`
var repoExample61 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally and Joe offered the Foo object",
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
var repoExample64 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Note",
  "name": "A Simple Note",
  "attachment": {
    "type": "Image",
    "content": "A simple Image",
    "url": "http://example.org/cat.jpeg"
  }
}`
var repoExample66 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "A Simple Image",
  "url": "http://example.org/cat.jpeg",
  "attributedTo": [
    "http://joe.example.org",
    {
      "type": "Person",
      "name": "Sally"
    }
  ]
}`
var repoExample68 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered a post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "bcc": "http://joe.example.org"
}`
var repoExample70 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered a post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "cc": "http://joe.example.org"
}`
var repoExample72 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "current": {
    "type": "Link",
    "name": "Most Recent Items",
    "href": "http://example.org/collection"
  },
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`
var repoExample74 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally's blog posts",
  "type": "Collection",
  "totalItems": 3,
  "first": {
    "type": "Link",
    "name": "First Page",
    "href": "http://example.org/collection?page=0"
  }
}`
var repoExample77 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "icon": {
    "type": "Image",
    "name": "Note",
    "url": "http://example.org/note.png",
    "width": 16,
    "height": 16
  }
}`
var repoExample80 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "image": {
    "type": "Image",
    "name": "A Cat",
    "url": "http://example.org/cat.png"
  }
}`
var repoExample83 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "inReplyTo": {
    "name": "Another note",
    "type": "Note",
    "content": "Another note"
  }
}`
var repoExample87 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A collection",
  "type": "Collection",
  "totalItems": 3,
  "last": "http://example.org/collection?page=1"
}`
var repoExample89 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
var repoExample91 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally's notes",
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
var repoExample93 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
var repoExample96 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Page 2 of Sally's blog posts",
  "type": "CollectionPage",
  "next": "http://example.org/collection?page=2",
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`
var repoExample98 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally liked a post",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`
var repoExample100 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally liked a note",
  "type": "Like",
  "actor": "http://sally.example.org",
  "object": [
    "http://example.org/posts/1",
    {
      "type": "Note",
      "name": "A simple note",
      "content": "A simple note"
    }
  ]
}`
var repoExample104 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Page 1 of Sally's blog posts",
  "type": "CollectionPage",
  "prev": "http://example.org/collection?page=1",
  "items": [
    "http://example.org/posts/1",
    "http://example.org/posts/2",
    "http://example.org/posts/3"
  ]
}`
var repoExample106 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Video",
  "name": "Cool New Movie",
  "duration": "PT2H30M",
  "preview": {
    "type": "Video",
    "name": "Trailer",
    "duration": "PT1M",
    "url": {
      "href": "http://example.org/trailer.mkv",
      "mediaType": "video/mkv"
    }
  }
}`
var repoExample108 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally checked that her flight was on time",
  "type": ["Activity", "http://www.verbs.example/Check"],
  "actor": "http://sally.example.org",
  "object": "http://example.org/flights/1",
  "result": {
    "type": "http://www.types.example/flightstatus",
    "name": "On Time"
  }
}`
var repoExample112 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "id": "http://www.test.example/notes/1",
  "content": "A simple note",
  "replies": {
    "type": "Collection",
    "totalItems": 1,
    "items": {
      "name": "A response to the note",
      "type": "Note",
      "content": "A response to the note",
      "inReplyTo": "http://www.test.example/notes/1"
    }
  }
}`
var repoExample118 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Image",
  "name": "Picture of Sally",
  "url": "http://example.org/sally.jpg",
  "tag": {
    "type": "Person",
    "id": "http://sally.example.org",
    "name": "Sally"
  }
}`
var repoExample120 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org"
}`
var repoExample123 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": "http://john.example.org",
  "to": "http://joe.example.org"
}`
var repoExample125 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Document",
  "name": "4Q Sales Forecast",
  "url": {
    "type": "Link",
    "href": "http://example.org/4q-sales-forecast.pdf"
  }
}`
var repoExample127 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China",
  "type": "Place",
  "latitude": 36.75,
  "longitude": 119.7667,
  "accuracy": 94.5
}`
var repoExample129 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "altitude": 15.0,
  "latitude": 36.75,
  "longitude": 119.7667,
  "units": "miles"
}`
var repoExample131 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "contentMap": {
    "en": "A &lt;i&gt;simple&lt;/i&gt; note",
    "sp": "Una &lt;i&gt;simple&lt;/i&gt; nota"
  }
}`
var repoExample133 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Note",
  "nameMap": {
    "en": "A simple note",
    "sp": "Una simple nota"
  }
}`
var repoExample136 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/image.png",
  "height": 100,
  "width": 100
}`
var repoExample138 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link"
}`
var repoExample140 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Place",
  "name": "Fresno Area",
  "latitude": 36.75,
  "longitude": 119.7667,
  "radius": 15,
  "units": "miles"
}`
var repoExample142 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link"
}`
var repoExample144 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "A Party!",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`
var repoExample146 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Event",
  "name": "A Party!",
  "startTime": "2014-12-31T23:00:00-08:00",
  "endTime": "2015-01-01T06:00:00-08:00"
}`
var repoExample149 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link",
  "rel": ["canonical", "preview"]
}`
var repoExample152 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "summary": "A simple &lt;i&gt;note&lt;/i&gt;"
}`
var repoExample156 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally's notes",
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
var repoExample158 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "updated": "2014-12-12T12:12:12Z"
}`
var repoExample161 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally read an article about Activity Streams",
  "type": "View",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Article",
    "name": "An article about Activity Streams"
  }
}`
var repoExample163 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally listened to a piece of music",
  "type": "Listen",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": "http://example.org/music.mp3"
}`
var repoExample166 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally moved a post from List A to List B",
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
var repoExample168 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally moved a post from List A to List B",
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
var repoExample170 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally announced that she had arrived at work",
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
var repoExample173 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally blocked Joe",
  "type": "Block",
  "actor": "http://sally.example.org",
  "object": "http://joe.example.org"
}`
var repoExample175 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "Sally disliked a post",
  "type": "Dislike",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1"
}`
var repoExample180 = `{
"@context": "http://www.w3.org/ns/activitystreams",
"name": "Sally's friends list",
"type": "Collection",
"items": [
{
 "name": "Sally is following Joe",
 "type": "Relationship",
 "subject": {
   "type": "Person",
   "name": "Sally"
 },
 "relationship": "IsFollowing",
 "object": {
   "type": "Person",
   "name": "Joe"
 }
},
{
 "name": "Sally is a contact of Jane",
 "type": "Relationship",
 "subject": {
   "type": "Person",
   "name": "Sally"
 },
 "relationship": "IsContact",
 "object": {
   "type": "Person",
   "name": "Jane"
 }
}
]
}`
var repoExample182 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Travel",
  "name": "Sally went to work",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "target": {
    "type": "Place",
    "name": "Work"
  }
}`
var repoExample184 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
var repoExample186 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "type": "Organization",
  "name": "Example Co."
}`
var repoExample188 = `{
"@context": "http://www.w3.org/ns/activitystreams",
"name": "Sally and John's relationship history",
"type": "Collection",
"items": [
{
 "name": "John accepted Sally's friend request",
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
 "name": "John followed Sally",
 "id": "http://example.org/activities/123",
 "type": "Follow",
 "actor": "acct:john@example.org",
 "object": "acct:sally@example.org",
 "context": "http://example.org/connections/123"
},
{
 "name": "Sally followed John",
 "id": "http://example.org/activities/124",
 "type": "Follow",
 "actor": "acct:sally@example.org",
 "object": "acct:john@example.org",
 "context": "http://example.org/connections/123"
},
{
 "name": "John added Sally to his friends list",
 "id": "http://example.org/activities/125",
 "type": "Add",
 "actor": "acct:john@example.org",
 "object": "http://example.org/connections/123",
 "target": {
   "type": "Collection",
   "name": "John's Connections"
 },
 "context": "http://example.org/connections/123"
},
{
 "name": "Sally added John to her friends list",
 "id": "http://example.org/activities/126",
 "type": "Add",
 "actor": "acct:sally@example.org",
 "object": "http://example.org/connections/123",
 "target": {
   "type": "Collection",
   "name": "Sally's Connections"
 },
 "context": "http://example.org/connections/123"
}
]
}`
var repoExample190 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "id": "http://polls.example.org/question/1",
  "name": "A question about robots",
  "type": "Question",
   "content": "I'd like to build a robot to feed my cat. Which platform is best?",
   "oneOf": [
     {"name": "arduino"},
     {"name": "raspberry pi"}
   ]
 }`
var repoExample192 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
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
     "content": "Users are favoriting &amp;quot;arduino&amp;quot; by a 33% margin."
   }
 }`
var repoExample194 = `{
 "@context": "http://www.w3.org/ns/activitystreams",
 "name": "Activities in Project XYZ",
 "type": "Collection",
 "items": [
   {
     "name": "Sally created a note",
     "type": "Create",
     "id": "http://activities.example.com/1",
     "actor": "http://sally.example.org",
     "object": {
      "name": "A note",
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
     "name": "John liked Sally's note",
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
var repoExample196 = `{
  "@context": "http://www.w3.org/ns/activitystreams",
  "name": "A thank-you note",
  "type": "Note",
  "content": "Thank you &lt;a href='http://sally.example.org'&gt;@sally&lt;/a&gt;\nfor all your hard work!\n&lt;a href='http://example.org/tags/givingthanks'&gt;#givingthanks&lt;/a&gt;",
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
var repoExample198 = `{
    "@context": "http://www.w3.org/ns/activitystreams",
    "name": "Sally moved the sales figures from Folder A to Folder B",
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
var allRepoExamples = map[string]string{
	"Example 1":   repoExample1,
	"Example 3":   repoExample3,
	"Example 5":   repoExample5,
	"Example 7":   repoExample7,
	"Example 9":   repoExample9,
	"Example 11":  repoExample11,
	"Example 13":  repoExample13,
	"Example 15":  repoExample15,
	"Example 17":  repoExample17,
	"Example 19":  repoExample19,
	"Example 21":  repoExample21,
	"Example 24":  repoExample24,
	"Example 26":  repoExample26,
	"Example 28":  repoExample28,
	"Example 32":  repoExample32,
	"Example 34":  repoExample34,
	"Example 37":  repoExample37,
	"Example 39":  repoExample39,
	"Example 42":  repoExample42,
	"Example 48":  repoExample48,
	"Example 50":  repoExample50,
	"Example 52":  repoExample52,
	"Example 55":  repoExample55,
	"Example 57":  repoExample57,
	"Example 59":  repoExample59,
	"Example 61":  repoExample61,
	"Example 64":  repoExample64,
	"Example 66":  repoExample66,
	"Example 68":  repoExample68,
	"Example 70":  repoExample70,
	"Example 72":  repoExample72,
	"Example 74":  repoExample74,
	"Example 77":  repoExample77,
	"Example 80":  repoExample80,
	"Example 83":  repoExample83,
	"Example 87":  repoExample87,
	"Example 89":  repoExample89,
	"Example 91":  repoExample91,
	"Example 93":  repoExample93,
	"Example 96":  repoExample96,
	"Example 98":  repoExample98,
	"Example 100": repoExample100,
	"Example 104": repoExample104,
	"Example 106": repoExample106,
	"Example 108": repoExample108,
	"Example 112": repoExample112,
	"Example 118": repoExample118,
	"Example 120": repoExample120,
	"Example 123": repoExample123,
	"Example 125": repoExample125,
	"Example 127": repoExample127,
	"Example 129": repoExample129,
	"Example 131": repoExample131,
	"Example 133": repoExample133,
	"Example 136": repoExample136,
	"Example 138": repoExample138,
	"Example 140": repoExample140,
	"Example 142": repoExample142,
	"Example 144": repoExample144,
	"Example 146": repoExample146,
	"Example 149": repoExample149,
	"Example 152": repoExample152,
	"Example 156": repoExample156,
	"Example 158": repoExample158,
	"Example 161": repoExample161,
	"Example 163": repoExample163,
	"Example 166": repoExample166,
	"Example 168": repoExample168,
	"Example 170": repoExample170,
	"Example 173": repoExample173,
	"Example 175": repoExample175,
	"Example 180": repoExample180,
	"Example 182": repoExample182,
	"Example 184": repoExample184,
	"Example 186": repoExample186,
	"Example 188": repoExample188,
	"Example 190": repoExample190,
	"Example 192": repoExample192,
	"Example 194": repoExample194,
	"Example 196": repoExample196,
	"Example 198": repoExample198,
}
