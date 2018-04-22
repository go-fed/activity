package pub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	postContentTypeHeader     = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	responseContentTypeHeader = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	getAcceptHeader           = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	contentTypeHeader         = "Content-Type"
	acceptHeader              = "Accept"
	publicActivityPub         = "https://www.w3.org/ns/activitystreams#Public"
	publicJsonLD              = "Public"
	publicJsonLDAS            = "as:Public"
	jsonLDContext             = "@context"
	activityPubContext        = "https://www.w3.org/ns/activitystreams"
)

var alternatives = []string{"application/activity+json"}

func trimAll(s []string) []string {
	var r []string
	for _, e := range s {
		r = append(r, strings.Trim(e, " "))
	}
	return r
}

func headerEqualsOneOf(header string, acceptable []string) bool {
	sanitizedHeader := strings.Join(trimAll(strings.Split(header, ";")), ";")
	for _, v := range acceptable {
		// Remove any number of whitespace after ;'s
		sanitizedV := strings.Join(trimAll(strings.Split(v, ";")), ";")
		if sanitizedHeader == sanitizedV {
			return true
		}
	}
	return false
}

func isActivityPubPost(r *http.Request) bool {
	return r.Method == "POST" && headerEqualsOneOf(r.Header.Get(contentTypeHeader), []string{postContentTypeHeader, contentTypeHeader})
}

func isActivityPubGet(r *http.Request) bool {
	return r.Method == "GET" && headerEqualsOneOf(r.Header.Get(acceptHeader), []string{getAcceptHeader, contentTypeHeader})
}

// isPublic determines if a target is the Public collection as defined in the
// spec, including JSON-LD compliant collections.
func isPublic(s string) bool {
	return s == publicActivityPub || s == publicJsonLD || s == publicJsonLDAS
}

func addJSONLDContext(m map[string]interface{}) {
	m[jsonLDContext] = activityPubContext
}

// dereference makes an HTTP GET request to an IRI in order to obtain the
// ActivityStream representation.
func dereference(c HttpClient, u url.URL, agent string) ([]byte, error) {
	// TODO: (section 7.1) The server MUST dereference the collection (with the user's credentials)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(acceptHeader, getAcceptHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request to %s failed (%d): %s", u.String(), resp.StatusCode, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

// postToOutbox will attempt to send a POST request to the given URL with the
// body set to the provided bytes.
func postToOutbox(c HttpClient, b []byte, to url.URL, agent string) error {
	byteCopy := make([]byte, 0, len(b))
	copy(b, byteCopy)
	buf := bytes.NewBuffer(byteCopy)
	req, err := http.NewRequest("POST", to.String(), buf)
	if err != nil {
		return err
	}
	req.Header.Add(contentTypeHeader, postContentTypeHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request to %s failed (%d): %s", to.String(), resp.StatusCode, resp.Status)
	}
	return nil
}

// wrapInCreate will automatically wrap the provided object in a Create
// activity. This will copy over the 'to', 'bto', 'cc', 'bcc', and 'audience'
// properties. It will also copy over the published time if present.
func (f *federator) wrapInCreate(o vocab.ObjectType, actor url.URL) *vocab.Create {
	c := &vocab.Create{}
	c.AddObject(o)
	c.AddActorIRI(actor)
	if o.IsPublished() {
		c.SetPublished(o.GetPublished())
	}
	for i := 0; i < o.ToLen(); i++ {
		if o.IsToObject(i) {
			c.AddToObject(o.GetToObject(i))
		} else if o.IsToLink(i) {
			c.AddToLink(o.GetToLink(i))
		} else if o.IsToIRI(i) {
			c.AddToIRI(o.GetToIRI(i))
		}
	}
	for i := 0; i < o.BtoLen(); i++ {
		if o.IsBtoObject(i) {
			c.AddBtoObject(o.GetBtoObject(i))
		} else if o.IsBtoLink(i) {
			c.AddBtoLink(o.GetBtoLink(i))
		} else if o.IsBtoIRI(i) {
			c.AddBtoIRI(o.GetBtoIRI(i))
		}
	}
	for i := 0; i < o.CcLen(); i++ {
		if o.IsCcObject(i) {
			c.AddCcObject(o.GetCcObject(i))
		} else if o.IsCcLink(i) {
			c.AddCcLink(o.GetCcLink(i))
		} else if o.IsCcIRI(i) {
			c.AddCcIRI(o.GetCcIRI(i))
		}
	}
	for i := 0; i < o.BccLen(); i++ {
		if o.IsBccObject(i) {
			c.AddBccObject(o.GetBccObject(i))
		} else if o.IsBccLink(i) {
			c.AddBccLink(o.GetBccLink(i))
		} else if o.IsBccIRI(i) {
			c.AddBccIRI(o.GetBccIRI(i))
		}
	}
	for i := 0; i < o.AudienceLen(); i++ {
		if o.IsAudienceObject(i) {
			c.AddAudienceObject(o.GetAudienceObject(i))
		} else if o.IsAudienceLink(i) {
			c.AddAudienceLink(o.GetAudienceLink(i))
		} else if o.IsAudienceIRI(i) {
			c.AddAudienceIRI(o.GetAudienceIRI(i))
		}
	}
	return c
}

// Ensures the activity and object have the same 'to', 'bto', 'cc', 'bcc', and
// 'audience' properties. Copy the Activity's recipients to objects, and the
// objects to the activity, but does NOT copy objects' recipients to each other.
//
// If there is any disagreement between the activity and an object, we default
// to a no-op.
func (f *federator) sameRecipients(a vocab.ActivityType) error {
	// First, map recipients for each object and the activity.
	to := make([]map[string]interface{}, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		to[i] = make(map[string]interface{})
		if !a.IsObject(i) {
			return fmt.Errorf("sameRecipients does not support 'to' object IRIs on Activities")
		}
		o := a.GetObject(i)
		for j := 0; j < o.ToLen(); j++ {
			if o.IsToObject(j) {
				id := o.GetToObject(j).GetId()
				to[i][(&id).String()] = o.GetToObject(j)
			} else if o.IsToLink(j) {
				id := o.GetToLink(j).GetHref()
				to[i][(&id).String()] = o.GetToLink(j)
			} else if o.IsToIRI(j) {
				id := o.GetToIRI(j)
				to[i][(&id).String()] = id
			}
		}
	}
	toActivity := make(map[string]interface{})
	for i := 0; i < a.ToLen(); i++ {
		if a.IsToObject(i) {
			id := a.GetToObject(i).GetId()
			toActivity[(&id).String()] = a.GetToObject(i)
		} else if a.IsToLink(i) {
			id := a.GetToLink(i).GetHref()
			toActivity[(&id).String()] = a.GetToLink(i)
		} else if a.IsToIRI(i) {
			id := a.GetToIRI(i)
			toActivity[(&id).String()] = id
		}
	}
	bto := make([]map[string]interface{}, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		bto[i] = make(map[string]interface{})
		if !a.IsObject(i) {
			return fmt.Errorf("sameRecipients does not support 'bto' object IRIs on Activities")
		}
		o := a.GetObject(i)
		for j := 0; j < o.BtoLen(); j++ {
			if o.IsBtoObject(j) {
				id := o.GetBtoObject(j).GetId()
				bto[i][(&id).String()] = o.GetBtoObject(j)
			} else if o.IsBtoLink(j) {
				id := o.GetBtoLink(j).GetHref()
				bto[i][(&id).String()] = o.GetBtoLink(j)
			} else if o.IsBtoIRI(j) {
				id := o.GetBtoIRI(j)
				bto[i][(&id).String()] = id
			}
		}
	}
	btoActivity := make(map[string]interface{})
	for i := 0; i < a.BtoLen(); i++ {
		if a.IsBtoObject(i) {
			id := a.GetBtoObject(i).GetId()
			btoActivity[(&id).String()] = a.GetBtoObject(i)
		} else if a.IsBtoLink(i) {
			id := a.GetBtoLink(i).GetHref()
			btoActivity[(&id).String()] = a.GetBtoLink(i)
		} else if a.IsBtoIRI(i) {
			id := a.GetBtoIRI(i)
			btoActivity[(&id).String()] = id
		}
	}
	cc := make([]map[string]interface{}, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		cc[i] = make(map[string]interface{})
		if !a.IsObject(i) {
			return fmt.Errorf("sameRecipients does not support 'cc' object IRIs on Activities")
		}
		o := a.GetObject(i)
		for j := 0; j < o.CcLen(); j++ {
			if o.IsCcObject(j) {
				id := o.GetCcObject(j).GetId()
				cc[i][(&id).String()] = o.GetCcObject(j)
			} else if o.IsCcLink(j) {
				id := o.GetCcLink(j).GetHref()
				cc[i][(&id).String()] = o.GetCcLink(j)
			} else if o.IsCcIRI(j) {
				id := o.GetCcIRI(j)
				cc[i][(&id).String()] = id
			}
		}
	}
	ccActivity := make(map[string]interface{})
	for i := 0; i < a.CcLen(); i++ {
		if a.IsCcObject(i) {
			id := a.GetCcObject(i).GetId()
			ccActivity[(&id).String()] = a.GetCcObject(i)
		} else if a.IsCcLink(i) {
			id := a.GetCcLink(i).GetHref()
			ccActivity[(&id).String()] = a.GetCcLink(i)
		} else if a.IsCcIRI(i) {
			id := a.GetCcIRI(i)
			ccActivity[(&id).String()] = id
		}
	}
	bcc := make([]map[string]interface{}, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		bcc[i] = make(map[string]interface{})
		if !a.IsObject(i) {
			return fmt.Errorf("sameRecipients does not support 'bcc' object IRIs on Activities")
		}
		o := a.GetObject(i)
		for j := 0; j < o.BccLen(); j++ {
			if o.IsBccObject(j) {
				id := o.GetBccObject(j).GetId()
				bcc[i][(&id).String()] = o.GetBccObject(j)
			} else if o.IsBccLink(j) {
				id := o.GetBccLink(j).GetHref()
				bcc[i][(&id).String()] = o.GetBccLink(j)
			} else if o.IsBccIRI(j) {
				id := o.GetBccIRI(j)
				bcc[i][(&id).String()] = id
			}
		}
	}
	bccActivity := make(map[string]interface{})
	for i := 0; i < a.BccLen(); i++ {
		if a.IsBccObject(i) {
			id := a.GetBccObject(i).GetId()
			bccActivity[(&id).String()] = a.GetBccObject(i)
		} else if a.IsBccLink(i) {
			id := a.GetBccLink(i).GetHref()
			bccActivity[(&id).String()] = a.GetBccLink(i)
		} else if a.IsBccIRI(i) {
			id := a.GetBccIRI(i)
			bccActivity[(&id).String()] = id
		}
	}
	audience := make([]map[string]interface{}, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		audience[i] = make(map[string]interface{})
		if !a.IsObject(i) {
			return fmt.Errorf("sameRecipients does not support 'audience' object IRIs on Activities")
		}
		o := a.GetObject(i)
		for j := 0; j < o.AudienceLen(); j++ {
			if o.IsAudienceObject(j) {
				id := o.GetAudienceObject(j).GetId()
				audience[i][(&id).String()] = o.GetAudienceObject(j)
			} else if o.IsAudienceLink(j) {
				id := o.GetAudienceLink(j).GetHref()
				audience[i][(&id).String()] = o.GetAudienceLink(j)
			} else if o.IsAudienceIRI(j) {
				id := o.GetAudienceIRI(j)
				audience[i][(&id).String()] = id
			}
		}
	}
	audienceActivity := make(map[string]interface{})
	for i := 0; i < a.AudienceLen(); i++ {
		if a.IsAudienceObject(i) {
			id := a.GetAudienceObject(i).GetId()
			audienceActivity[(&id).String()] = a.GetAudienceObject(i)
		} else if a.IsAudienceLink(i) {
			id := a.GetAudienceLink(i).GetHref()
			audienceActivity[(&id).String()] = a.GetAudienceLink(i)
		} else if a.IsAudienceIRI(i) {
			id := a.GetAudienceIRI(i)
			audienceActivity[(&id).String()] = id
		}
	}
	// Next, add activity recipients to all objects if not already present
	for k, v := range toActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := to[i][k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.GetObject(i).AddToObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.GetObject(i).AddToLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.GetObject(i).AddToIRI(vIRI)
				}
			}
		}
	}
	for k, v := range btoActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := bto[i][k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.GetObject(i).AddBtoObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.GetObject(i).AddBtoLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.GetObject(i).AddBtoIRI(vIRI)
				}
			}
		}
	}
	for k, v := range ccActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := cc[i][k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.GetObject(i).AddCcObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.GetObject(i).AddCcLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.GetObject(i).AddCcIRI(vIRI)
				}
			}
		}
	}
	for k, v := range bccActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := bcc[i][k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.GetObject(i).AddBccObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.GetObject(i).AddBccLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.GetObject(i).AddBccIRI(vIRI)
				}
			}
		}
	}
	for k, v := range audienceActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := audience[i][k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.GetObject(i).AddAudienceObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.GetObject(i).AddAudienceLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.GetObject(i).AddAudienceIRI(vIRI)
				}
			}
		}
	}
	// Finally, add all the objects' recipients to the activity if not
	// already present.
	for i := 0; i < a.ObjectLen(); i++ {
		for k, v := range to[i] {
			if _, ok := toActivity[k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.AddToObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.AddToLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.AddToIRI(vIRI)
				}
			}
		}
		for k, v := range bto[i] {
			if _, ok := btoActivity[k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.AddBtoObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.AddBtoLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.AddBtoIRI(vIRI)
				}
			}
		}
		for k, v := range cc[i] {
			if _, ok := ccActivity[k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.AddCcObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.AddCcLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.AddCcIRI(vIRI)
				}
			}
		}
		for k, v := range bcc[i] {
			if _, ok := bccActivity[k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.AddBccObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.AddBccLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.AddBccIRI(vIRI)
				}
			}
		}
		for k, v := range audience[i] {
			if _, ok := audienceActivity[k]; !ok {
				if vObj, ok := v.(vocab.ObjectType); ok {
					a.AddAudienceObject(vObj)
				} else if vLink, ok := v.(vocab.LinkType); ok {
					a.AddAudienceLink(vLink)
				} else if vIRI, ok := v.(url.URL); ok {
					a.AddAudienceIRI(vIRI)
				}
			}
		}
	}
	return nil
}

// TODO: (Section 7) HTTP caching mechanisms [RFC7234] SHOULD be respected when appropriate, both when receiving responses from other servers as well as sending responses to other servers.

// deliver will complete the peer-to-peer sending of a federated message to
// another server.
func (f *federator) deliver(obj vocab.ActivityType) error {
	recipients, err := f.prepare(obj)
	if err != nil {
		return err
	}
	m, err := obj.Serialize()
	if err != nil {
		return err
	}
	addJSONLDContext(m)
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	for _, to := range recipients {
		f.deliverer.Do(b, to, func(b []byte, u url.URL) error {
			return postToOutbox(f.Client, b, u, f.Agent)
		})
	}
	return nil
}

// prepare takes a deliverableObject and returns a list of the proper recipient
// target URIs. Additionally, the deliverableObject will have any hidden
// hidden recipients ("bto" and "bcc") stripped from it.
func (c *federator) prepare(o deliverableObject) ([]url.URL, error) {
	// Get inboxes of recipients
	var r []url.URL
	r = append(r, getToIRIs(o)...)
	r = append(r, getBToIRIs(o)...)
	r = append(r, getCcIRIs(o)...)
	r = append(r, getBccIRIs(o)...)
	r = append(r, getAudienceIRIs(o)...)
	// TODO: Support delivery to shared inbox
	// 1. When an object is being delivered to the originating actor's
	// followers, a server MAY reduce the number of receiving actors
	// delivered to by identifying all followers which share the same
	// sharedInbox who would otherwise be individual recipients and instead
	// deliver objects to said sharedInbox.
	// 2. If an object is addressed to the Public special collection, a
	// server MAY deliver that object to all known sharedInbox endpoints on
	// the network.
	r = filterURLs(r, isPublic)
	receiverActors, err := c.resolveInboxes(r, 0, c.MaxDeliveryDepth)
	if err != nil {
		return nil, err
	}
	targets, err := getInboxes(receiverActors)
	if err != nil {
		return nil, err
	}
	// Get inboxes of sender(s)
	senderActors, err := c.resolveInboxes(getActorsAttributedToURI(o), 0, c.MaxDeliveryDepth)
	if err != nil {
		return nil, err
	}
	ignore, err := getInboxes(senderActors)
	if err != nil {
		return nil, err
	}
	// Post-processing
	r = dedupeIRIs(targets, ignore)
	stripHiddenRecipients(o)
	return r, nil
}

// resolveInboxes takes a list of Actor id URIs and returns them as concrete
// instances of actorObject. It applies recursively when it encounters a target
// that is a Collection or OrderedCollection.
func (c *federator) resolveInboxes(r []url.URL, depth int, max int) ([]actor, error) {
	if depth >= max {
		return nil, nil
	}
	a := make([]actor, 0, len(r))
	for _, u := range r {
		// Do not retry here -- if a dereference fails, then fail the
		// entire delivery.
		resp, err := dereference(c.Client, u, c.Agent)
		if err != nil {
			return nil, err
		}
		var m map[string]interface{}
		if err = json.Unmarshal(resp, &m); err != nil {
			return nil, err
		}
		var actor actor
		var co *streams.Collection
		var oc *streams.OrderedCollection
		var cp *streams.CollectionPage
		var ocp *streams.OrderedCollectionPage
		// Set AT MOST one of: co, oc, cp, ocp
		// If none of these are set, attempt to use actor
		if err = toActorCollectionResolver(&actor, &co, &oc, &cp, &ocp).Deserialize(m); err != nil {
			return nil, err
		}
		// If a recipient is a Collection or OrderedCollection, then the
		// server MUST dereference the collection. Note that this also
		// applies to CollectionPage and OrderedCollectionPage.
		var uris []url.URL
		if co != nil {
			uris := getURIsInItemer(co.Raw())
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if oc != nil {
			uris := getURIsInOrderedItemer(oc.Raw())
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if cp != nil {
			cb := func(c vocab.CollectionPageType) error {
				uris = getURIsInItemer(c)
				return nil
			}
			err := doForCollectionPage(c.Client, c.Agent, cb, cp.Raw())
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if ocp != nil {
			cb := func(c vocab.OrderedCollectionPageType) error {
				uris = getURIsInOrderedItemer(c)
				return nil
			}
			err := doForOrderedCollectionPage(c.Client, c.Agent, cb, ocp.Raw())
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if actor != nil {
			a = append(a, actor)
		}
	}
	return a, nil
}

// getInboxes extracts the 'inbox' IRIs from actors.
func getInboxes(a []actor) ([]url.URL, error) {
	var u []url.URL
	for _, actor := range a {
		if actor.IsInboxAnyURI() {
			u = append(u, actor.GetInboxAnyURI())
		} else if actor.IsInboxOrderedCollection() {
			oc := actor.GetInboxOrderedCollection()
			if !oc.HasId() {
				return nil, fmt.Errorf("actor inbox OrderedCollection has no IRI")
			}
			u = append(u, oc.GetId())
		}
	}
	return u, nil
}

// getActorAttributedToURI attempts to find the URIs for the "actor" and
// "attributedTo" originators on the object.
func getActorsAttributedToURI(a actorObject) []url.URL {
	var u []url.URL
	for i := 0; i < a.AttributedToLen(); i++ {
		if a.IsAttributedToObject(i) {
			obj := a.GetAttributedToObject(i)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if a.IsAttributedToLink(i) {
			l := a.GetAttributedToLink(i)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if a.IsAttributedToIRI(i) {
			u = append(u, a.GetAttributedToIRI(i))
		}
	}
	for i := 0; i < a.ActorLen(); i++ {
		if a.IsActorObject(i) {
			obj := a.GetActorObject(i)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if a.IsActorLink(i) {
			l := a.GetActorLink(i)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if a.IsActorIRI(i) {
			u = append(u, a.GetActorIRI(i))
		}
	}
	return u
}

// stripHiddenRecipients removes "bto" and "bcc" from the deliverableObject.
// Note that this requirement of the specification is under "Section 6: Client
// to Server Interactions", the Social API, and not the Federative API.
func stripHiddenRecipients(o deliverableObject) {
	for o.BtoLen() > 0 {
		if o.IsBtoObject(0) {
			o.RemoveBtoObject(0)
		} else if o.IsBtoLink(0) {
			o.RemoveBtoLink(0)
		} else if o.IsBtoIRI(0) {
			o.RemoveBtoIRI(0)
		}
	}
	for o.BccLen() > 0 {
		if o.IsBccObject(0) {
			o.RemoveBccObject(0)
		} else if o.IsBccLink(0) {
			o.RemoveBccLink(0)
		} else if o.IsBtoIRI(0) {
			o.RemoveBccIRI(0)
		}
	}
}

// dedupeIRIs will deduplicate final inbox IRIs. The ignore list is applied to
// the final list
func dedupeIRIs(recipients, ignored []url.URL) (out []url.URL) {
	ignoredMap := make(map[url.URL]bool, len(ignored))
	for _, elem := range ignored {
		ignoredMap[elem] = true
	}
	outMap := make(map[url.URL]bool, len(recipients))
	for _, k := range recipients {
		if !ignoredMap[k] && !outMap[k] {
			out = append(out, k)
			outMap[k] = true
		}
	}
	return
}

// dedupeOrderedItems will deduplicate the 'orderedItems' within an ordered
// collection type. Deduplication happens by simply examining the 'id'.
func (f *federator) dedupeOrderedItems(oc vocab.OrderedCollectionType) (vocab.OrderedCollectionType, error) {
	i := 0
	seen := make(map[string]bool, oc.OrderedItemsLen())
	for i < oc.OrderedItemsLen() {
		var id string
		var removeFn func(int)
		if oc.IsOrderedItemsObject(i) {
			removeFn = oc.RemoveOrderedItemsObject
			iri := oc.GetOrderedItemsObject(i).GetId()
			pIri := &iri
			id = pIri.String()
		} else if oc.IsOrderedItemsLink(i) {
			removeFn = oc.RemoveOrderedItemsLink
			iri := oc.GetOrderedItemsLink(i).GetId()
			pIri := &iri
			id = pIri.String()
		} else if oc.IsOrderedItemsIRI(i) {
			removeFn = oc.RemoveOrderedItemsIRI
			b, err := dereference(f.Client, oc.GetOrderedItemsIRI(i), f.Agent)
			var m map[string]interface{}
			if err := json.Unmarshal(b, &m); err != nil {
				return oc, err
			}
			var iri url.URL
			var hasIri bool
			if err = toIdResolver(&hasIri, &iri).Deserialize(m); err != nil {
				return oc, err
			}
			pIri := &iri
			id = pIri.String()
		}
		if seen[id] {
			removeFn(i)
		} else {
			seen[id] = true
			i++
		}
	}
	return oc, nil
}

// filterURLs removes urls whose strings match the provided filter
func filterURLs(u []url.URL, fn func(s string) bool) []url.URL {
	i := 0
	for i < len(u) {
		if fn(u[i].String()) {
			u = append(u[:i], u[i+1:]...)
		} else {
			i++
		}
	}
	return u
}

func getToIRIs(o deliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.ToLen(); i++ {
		if o.IsToObject(i) {
			obj := o.GetToObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsToLink(i) {
			l := o.GetToLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsToIRI(i) {
			r = append(r, o.GetToIRI(i))
		}
	}
	return r
}

func getBToIRIs(o deliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.BtoLen(); i++ {
		if o.IsBtoObject(i) {
			obj := o.GetBtoObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsBtoLink(i) {
			l := o.GetBtoLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsBtoIRI(i) {
			r = append(r, o.GetBtoIRI(i))
		}
	}
	return r
}

func getCcIRIs(o deliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.CcLen(); i++ {
		if o.IsCcObject(i) {
			obj := o.GetCcObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsCcLink(i) {
			l := o.GetCcLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsCcIRI(i) {
			r = append(r, o.GetCcIRI(i))
		}
	}
	return r
}

func getBccIRIs(o deliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.BccLen(); i++ {
		if o.IsBccObject(i) {
			obj := o.GetBccObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsBccLink(i) {
			l := o.GetBccLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsBccIRI(i) {
			r = append(r, o.GetBccIRI(i))
		}
	}
	return r
}

func getAudienceIRIs(o deliverableObject) []url.URL {
	var r []url.URL
	for i := 0; i < o.AudienceLen(); i++ {
		if o.IsAudienceObject(i) {
			obj := o.GetAudienceObject(i)
			if obj.HasId() {
				r = append(r, obj.GetId())
			}
		} else if o.IsAudienceLink(i) {
			l := o.GetAudienceLink(i)
			if l.HasHref() {
				r = append(r, l.GetHref())
			}
		} else if o.IsAudienceIRI(i) {
			r = append(r, o.GetAudienceIRI(i))
		}
	}
	return r
}

// doForCollectionPage applies a function over a collection and its subsequent
// pages recursively. It returns the first non-nil error it encounters.
func doForCollectionPage(h HttpClient, agent string, cb func(c vocab.CollectionPageType) error, c vocab.CollectionPageType) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForCollectionPage(h, agent, cb, c.GetNextCollectionPage())
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent)
			if err != nil {
				return err
			}
			var m map[string]interface{}
			err = json.Unmarshal(resp, &m)
			if err != nil {
				return err
			}
			next, err := toCollectionPage(m)
			if err != nil {
				return err
			}
			if next != nil {
				return doForCollectionPage(h, agent, cb, next.Raw())
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent)
		if err != nil {
			return err
		}
		var m map[string]interface{}
		err = json.Unmarshal(resp, &m)
		if err != nil {
			return err
		}
		next, err := toCollectionPage(m)
		if err != nil {
			return err
		}
		if next != nil {
			return doForCollectionPage(h, agent, cb, next.Raw())
		}
	}
	return nil
}

// doForOrderedCollectionPage applies a function over a collection and its
// subsequent pages recursively. It returns the first non-nil error it
// encounters.
func doForOrderedCollectionPage(h HttpClient, agent string, cb func(c vocab.OrderedCollectionPageType) error, c vocab.OrderedCollectionPageType) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextOrderedCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForOrderedCollectionPage(h, agent, cb, c.GetNextOrderedCollectionPage())
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent)
			if err != nil {
				return err
			}
			var m map[string]interface{}
			err = json.Unmarshal(resp, &m)
			if err != nil {
				return err
			}
			next, err := toOrderedCollectionPage(m)
			if err != nil {
				return err
			}
			if next != nil {
				return doForOrderedCollectionPage(h, agent, cb, next.Raw())
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent)
		if err != nil {
			return err
		}
		var m map[string]interface{}
		err = json.Unmarshal(resp, &m)
		if err != nil {
			return err
		}
		next, err := toOrderedCollectionPage(m)
		if err != nil {
			return err
		}
		if next != nil {
			return doForOrderedCollectionPage(h, agent, cb, next.Raw())
		}
	}
	return nil
}

type itemer interface {
	ItemsLen() (l int)
	IsItemsObject(index int) (ok bool)
	GetItemsObject(index int) (v vocab.ObjectType)
	IsItemsLink(index int) (ok bool)
	GetItemsLink(index int) (v vocab.LinkType)
	IsItemsIRI(index int) (ok bool)
	GetItemsIRI(index int) (v url.URL)
}

// getURIsInItemer will extract 'items' from the provided Collection or
// CollectionPage.
func getURIsInItemer(i itemer) []url.URL {
	var u []url.URL
	for j := 0; j < i.ItemsLen(); j++ {
		if i.IsItemsObject(j) {
			obj := i.GetItemsObject(j)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if i.IsItemsLink(j) {
			l := i.GetItemsLink(j)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if i.IsItemsIRI(j) {
			u = append(u, i.GetItemsIRI(j))
		}
	}
	return u
}

type orderedItemer interface {
	OrderedItemsLen() (l int)
	IsOrderedItemsObject(index int) (ok bool)
	GetOrderedItemsObject(index int) (v vocab.ObjectType)
	IsOrderedItemsLink(index int) (ok bool)
	GetOrderedItemsLink(index int) (v vocab.LinkType)
	IsOrderedItemsIRI(index int) (ok bool)
	GetOrderedItemsIRI(index int) (v url.URL)
}

// getURIsInOrderedItemer will extract 'items' from the provided
// OrderedCollection or OrderedCollectionPage.
func getURIsInOrderedItemer(i orderedItemer) []url.URL {
	var u []url.URL
	for j := 0; j < i.OrderedItemsLen(); j++ {
		if i.IsOrderedItemsObject(j) {
			obj := i.GetOrderedItemsObject(j)
			if obj.HasId() {
				u = append(u, obj.GetId())
			}
		} else if i.IsOrderedItemsLink(j) {
			l := i.GetOrderedItemsLink(j)
			if l.HasHref() {
				u = append(u, l.GetHref())
			}
		} else if i.IsOrderedItemsIRI(j) {
			u = append(u, i.GetOrderedItemsIRI(j))
		}
	}
	return u
}

type activityWithObject interface {
	ObjectLen() (l int)
	IsObject(index int) (ok bool)
	GetObject(index int) (v vocab.ObjectType)
	IsObjectIRI(index int) (ok bool)
	GetObjectIRI(index int) (v url.URL)
}

func getObjectIds(a activityWithObject) (u []url.URL, e error) {
	for i := 0; i < a.ObjectLen(); i++ {
		if a.IsObject(i) {
			obj := a.GetObject(i)
			if !obj.HasId() {
				e = fmt.Errorf("object has no id: %v", obj)
				return
			}
			u = append(u, obj.GetId())
		} else if a.IsObjectIRI(i) {
			u = append(u, a.GetObjectIRI(i))
		}
	}
	return
}

type activityWithTarget interface {
	TargetLen() (l int)
	IsTargetObject(index int) (ok bool)
	GetTargetObject(index int) (v vocab.ObjectType)
	IsTargetIRI(index int) (ok bool)
	GetTargetIRI(index int) (v url.URL)
}

func getTargetIds(a activityWithTarget) (u []url.URL, e error) {
	for i := 0; i < a.TargetLen(); i++ {
		if a.IsTargetObject(i) {
			obj := a.GetTargetObject(i)
			if !obj.HasId() {
				e = fmt.Errorf("object has no id: %v", obj)
				return
			}
			u = append(u, obj.GetId())
		} else if a.IsTargetIRI(i) {
			u = append(u, a.GetTargetIRI(i))
		}
	}
	return
}

func removeCollectionItemWithId(c vocab.CollectionType, iri url.URL) {
	for i := 0; i < c.ItemsLen(); i++ {
		if c.IsItemsObject(i) {
			o := c.GetItemsObject(i)
			if !o.HasId() {
				continue
			}
			if o.GetId() == iri {
				c.RemoveItemsObject(i)
				return
			}
		} else if c.IsItemsLink(i) {
			l := c.GetItemsLink(i)
			if !l.HasHref() {
				continue
			}
			if l.GetHref() == iri {
				c.RemoveItemsLink(i)
				return
			}
		} else if c.IsItemsIRI(i) {
			if c.GetItemsIRI(i) == iri {
				c.RemoveItemsIRI(i)
				return
			}
		}
	}
}

func removeOrderedCollectionItemWithId(c vocab.OrderedCollectionType, iri url.URL) {
	for i := 0; i < c.OrderedItemsLen(); i++ {
		if c.IsOrderedItemsObject(i) {
			o := c.GetOrderedItemsObject(i)
			if !o.HasId() {
				continue
			}
			if o.GetId() == iri {
				c.RemoveOrderedItemsObject(i)
				return
			}
		} else if c.IsOrderedItemsLink(i) {
			l := c.GetOrderedItemsLink(i)
			if !l.HasHref() {
				continue
			}
			if l.GetHref() == iri {
				c.RemoveOrderedItemsLink(i)
				return
			}
		} else if c.IsOrderedItemsIRI(i) {
			if c.GetOrderedItemsIRI(i) == iri {
				c.RemoveOrderedItemsIRI(i)
				return
			}
		}
	}
}

// toTombstone creates a Tombstone for the given object.
func toTombstone(obj vocab.ObjectType, id url.URL, now time.Time) vocab.TombstoneType {
	tomb := &vocab.Tombstone{}
	tomb.SetId(id)
	for i := 0; i < obj.TypeLen(); i++ {
		if s, ok := obj.GetType(i).(string); ok {
			tomb.AddFormerTypeString(s)
		} else if fObj, ok := obj.GetType(i).(vocab.ObjectType); ok {
			tomb.AddFormerTypeObject(fObj)
		}
	}
	if obj.IsPublished() {
		tomb.SetPublished(obj.GetPublished())
	} else if obj.IsPublishedIRI() {
		tomb.SetPublishedIRI(obj.GetPublishedIRI())
	}
	if obj.IsUpdated() {
		tomb.SetUpdated(obj.GetUpdated())
	} else if obj.IsUpdatedIRI() {
		tomb.SetUpdatedIRI(obj.GetUpdatedIRI())
	}
	tomb.SetDeleted(now)
	return tomb
}

func (f *federator) addAllObjectsToActorCollection(ctx context.Context, getter func(actor vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) error, c vocab.ActivityType) error {
	for i := 0; i < c.ActorLen(); i++ {
		var iri url.URL
		if c.IsActorObject(i) {
			obj := c.GetActorObject(i)
			if !obj.HasId() {
				return fmt.Errorf("actor does not have id")
			}
			iri = obj.GetId()
		} else if c.IsActorLink(i) {
			l := c.GetActorLink(i)
			if !l.HasHref() {
				return fmt.Errorf("actor Link href required")
			}
			iri = l.GetHref()
		} else if c.IsActorIRI(i) {
			iri = c.GetActorIRI(i)
		}
		if !f.App.Owns(ctx, iri) {
			continue
		}
		var actor vocab.ObjectType
		pObj, err := f.App.Get(ctx, iri)
		if err != nil {
			return err
		}
		ok := false
		actor, ok = pObj.(vocab.ObjectType)
		if !ok {
			return fmt.Errorf("actor is not vocab.ObjectType")
		}
		var lc vocab.CollectionType
		var loc vocab.OrderedCollectionType
		if err := getter(actor, &lc, &loc); err != nil {
			return err
		}
		for i := 0; i < c.ObjectLen(); i++ {
			if c.IsObjectIRI(i) {
				if lc != nil {
					lc.AddItemsIRI(c.GetObjectIRI(i))
				} else if loc != nil {
					loc.AddOrderedItemsIRI(c.GetObjectIRI(i))
				}
			} else if c.IsObject(i) {
				if lc != nil {
					lc.AddItemsObject(c.GetObject(i))
				} else if loc != nil {
					loc.AddOrderedItemsObject(c.GetObject(i))
				}
			}
		}
		if lc != nil {
			if err := f.App.Set(ctx, lc); err != nil {
				return err
			}
		} else if loc != nil {
			if err := f.App.Set(ctx, loc); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *federator) addAllActorsToObjectCollection(ctx context.Context, getter func(object vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) error, c vocab.ActivityType) error {
	for i := 0; i < c.ObjectLen(); i++ {
		var iri url.URL
		if c.IsObject(i) {
			obj := c.GetObject(i)
			if !obj.HasId() {
				return fmt.Errorf("object does not have id")
			}
			iri = obj.GetId()
		} else if c.IsObjectIRI(i) {
			iri = c.GetObjectIRI(i)
		}
		if !f.App.Owns(ctx, iri) {
			continue
		}
		var object vocab.ObjectType
		pObj, err := f.App.Get(ctx, iri)
		if err != nil {
			return err
		}
		ok := false
		object, ok = pObj.(vocab.ObjectType)
		if !ok {
			return fmt.Errorf("object is not vocab.ObjectType")
		}
		var lc vocab.CollectionType
		var loc vocab.OrderedCollectionType
		if err := getter(object, &lc, &loc); err != nil {
			return err
		}
		for i := 0; i < c.ActorLen(); i++ {
			if c.IsActorIRI(i) {
				if lc != nil {
					lc.AddItemsIRI(c.GetActorIRI(i))
				} else if loc != nil {
					loc.AddOrderedItemsIRI(c.GetActorIRI(i))
				}
			} else if c.IsActorObject(i) {
				if lc != nil {
					lc.AddItemsObject(c.GetActorObject(i))
				} else if loc != nil {
					loc.AddOrderedItemsObject(c.GetActorObject(i))
				}
			} else if c.IsActorLink(i) {
				if lc != nil {
					lc.AddItemsLink(c.GetActorLink(i))
				} else if loc != nil {
					loc.AddOrderedItemsLink(c.GetActorLink(i))
				}
			}
		}
		if lc != nil {
			if err := f.App.Set(ctx, lc); err != nil {
				return err
			}
		} else if loc != nil {
			if err := f.App.Set(ctx, loc); err != nil {
				return err
			}
		}
	}
	return nil
}

// TODO: Move this to vocab package.
var activityTypes = []string{"Accept", "Add", "Announce", "Arrive", "Block", "Create", "Delete", "Dislike", "Flag", "Follow", "Ignore", "Invite", "Join", "Leave", "Like", "Listen", "Move", "Offer", "Question", "Reject", "Read", "Remove", "TentativeReject", "TentativeAccept", "Travel", "Undo", "Update", "View"}

func isActivityType(t Typer) bool {
	hasType := make(map[string]bool, 1)
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			hasType[s] = true
		}
	}
	for _, t := range activityTypes {
		if hasType[t] {
			return true
		}
	}
	return false
}
