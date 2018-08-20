package pub

import (
	"bytes"
	"context"
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
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
	dateHeader                = "Date"
	digestHeader              = "Digest"
	acceptHeader              = "Accept"
	publicActivityPub         = "https://www.w3.org/ns/activitystreams#Public"
	publicJsonLD              = "Public"
	publicJsonLDAS            = "as:Public"
	jsonLDContext             = "@context"
	activityPubContext        = "https://www.w3.org/ns/activitystreams"
	sha256Digest              = "SHA-256"
	digestDelimiter           = "="
)

var mediaTypes []string

func init() {
	mediaTypes = []string{
		"application/activity+json",
	}
	jsonLdType := "application/ld+json"
	for _, semi := range []string{";", " ;", " ; ", "; "} {
		for _, profile := range []string{"profile=https://www.w3.org/ns/activitystreams", "profile=\"https://www.w3.org/ns/activitystreams\""} {
			mediaTypes = append(mediaTypes, fmt.Sprintf("%s%s%s", jsonLdType, semi, profile))
		}
	}
}

func trimAll(s []string) []string {
	var r []string
	for _, e := range s {
		r = append(r, strings.Trim(e, " "))
	}
	return r
}

func headerIsActivityPubMediaType(header string) bool {
	for _, mediaType := range mediaTypes {
		if strings.Contains(header, mediaType) {
			return true
		}
	}
	return false
}

func isActivityPubPost(r *http.Request) bool {
	return r.Method == "POST" && headerIsActivityPubMediaType(r.Header.Get(contentTypeHeader))
}

func isActivityPubGet(r *http.Request) bool {
	return r.Method == "GET" && headerIsActivityPubMediaType(r.Header.Get(acceptHeader))
}

// isPublic determines if a target is the Public collection as defined in the
// spec, including JSON-LD compliant collections.
func isPublic(s string) bool {
	return s == publicActivityPub || s == publicJsonLD || s == publicJsonLDAS
}

func addJSONLDContext(m map[string]interface{}) {
	m[jsonLDContext] = activityPubContext
}

func addResponseHeaders(h http.Header, c Clock, responseContent []byte) {
	h.Set(contentTypeHeader, responseContentTypeHeader)
	// RFC 7231 §7.1.1.2
	h.Set(dateHeader, c.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	// RFC 3230 and RFC 5843
	var b bytes.Buffer
	b.WriteString(sha256Digest)
	b.WriteString(digestDelimiter)
	hashed := sha256.Sum256(responseContent)
	b.WriteString(base64.StdEncoding.EncodeToString(hashed[:]))
	h.Set(digestHeader, b.String())
}

// dereference makes an HTTP GET request to an IRI in order to obtain the
// ActivityStream representation.
//
// creds is allowed to be nil.
func dereference(c HttpClient, u *url.URL, agent string, creds *creds, clock Clock) ([]byte, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(acceptHeader, getAcceptHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", clock.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	if creds != nil {
		err := creds.signer.SignRequest(creds.privKey, creds.pubKeyId, req)
		creds.privKey = nil
		if err != nil {
			return nil, err
		}
	}
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

type creds struct {
	signer   httpsig.Signer
	privKey  crypto.PrivateKey
	pubKeyId string
}

// dereferenceAsUser is meant to be used by the activity handlers that need to
// handle IRI use cases where objects are expected. Returns an error if not
// federating.
func (f *federator) dereferenceAsUser(boxIRI, fetchIRI *url.URL) (obj vocab.ObjectType, err error) {
	if !f.EnableServer {
		err = fmt.Errorf("cannot dereference iri as user if not federating: %q", fetchIRI)
		return
	}
	creds := &creds{}
	creds.signer, err = f.FederateAPI.NewSigner()
	if err != nil {
		return
	}
	creds.privKey, creds.pubKeyId, err = f.FederateAPI.PrivateKey(boxIRI)
	if err != nil {
		return
	}
	resp, err := dereference(f.Client, fetchIRI, f.Agent, creds, f.Clock)
	if err != nil {
		return
	}
	var m map[string]interface{}
	if err = json.Unmarshal(resp, &m); err != nil {
		return
	}
	return toAnyObject(m)
}

// postToOutbox will attempt to send a POST request to the given URL with the
// body set to the provided bytes.
//
// creds is able to be nil.
func postToOutbox(c HttpClient, b []byte, to *url.URL, agent string, creds *creds, clock Clock) error {
	byteCopy := make([]byte, len(b))
	copy(byteCopy, b)
	buf := bytes.NewBuffer(byteCopy)
	req, err := http.NewRequest("POST", to.String(), buf)
	if err != nil {
		return err
	}
	req.Header.Add(contentTypeHeader, postContentTypeHeader)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", clock.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s (go-fed ActivityPub)", agent))
	if creds != nil {
		err := creds.signer.SignRequest(creds.privKey, creds.pubKeyId, req)
		creds.privKey = nil
		if err != nil {
			return err
		}
	}
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

// addNewIds will add new IDs not just for an activity, but all objects
// contained within the activity if it is a Create activity.
func (f *federator) addNewIds(c context.Context, a vocab.ActivityType) {
	newId := f.App.NewId(c, a)
	a.SetId(newId)
	if vocab.HasTypeCreate(a) {
		for i := 0; i < a.ObjectLen(); i++ {
			if a.IsObject(i) {
				obj := a.GetObject(i)
				obj.SetId(f.App.NewId(c, obj))
			}
		}
	}
}

// addNewIdsIntransitive will add new IDs for an intransitive activity.
func (f *federator) addNewIdsIntransitive(c context.Context, a vocab.IntransitiveActivityType) {
	newId := f.App.NewId(c, a)
	a.SetId(newId)
}

// wrapInCreate will automatically wrap the provided object in a Create
// activity. This will copy over the 'to', 'bto', 'cc', 'bcc', and 'audience'
// properties. It will also copy over the published time if present.
func (f *federator) wrapInCreate(o vocab.ObjectType, actor *url.URL) (c *vocab.Create, err error) {
	c = &vocab.Create{}
	c.AppendType("Create")
	c.AppendObject(o)
	c.AppendActorIRI(actor)
	if o.IsPublished() {
		c.SetPublished(o.GetPublished())
	}
	for i := 0; i < o.ToLen(); i++ {
		var to *url.URL
		if o.IsToObject(i) {
			obj := o.GetToObject(i)
			if !obj.HasId() {
				err = fmt.Errorf("to object missing id")
				return
			}
			to = obj.GetId()
		} else if o.IsToLink(i) {
			href := o.GetToLink(i)
			if !href.HasHref() {
				err = fmt.Errorf("to link missing href")
				return
			}
			to = href.GetHref()
		} else if o.IsToIRI(i) {
			to = o.GetToIRI(i)
		}
		c.AppendToIRI(to)
	}
	for i := 0; i < o.BtoLen(); i++ {
		var bto *url.URL
		if o.IsBtoObject(i) {
			obj := o.GetBtoObject(i)
			if !obj.HasId() {
				err = fmt.Errorf("bto object missing id")
				return
			}
			bto = obj.GetId()
		} else if o.IsBtoLink(i) {
			href := o.GetBtoLink(i)
			if !href.HasHref() {
				err = fmt.Errorf("bto link missing href")
				return
			}
			bto = href.GetHref()
		} else if o.IsBtoIRI(i) {
			bto = o.GetBtoIRI(i)
		}
		c.AppendBtoIRI(bto)
	}
	for i := 0; i < o.CcLen(); i++ {
		var cc *url.URL
		if o.IsCcObject(i) {
			obj := o.GetCcObject(i)
			if !obj.HasId() {
				err = fmt.Errorf("cc object missing id")
				return
			}
			cc = obj.GetId()
		} else if o.IsCcLink(i) {
			href := o.GetCcLink(i)
			if !href.HasHref() {
				err = fmt.Errorf("cc link missing href")
				return
			}
			cc = href.GetHref()
		} else if o.IsCcIRI(i) {
			cc = o.GetCcIRI(i)
		}
		c.AppendCcIRI(cc)
	}
	for i := 0; i < o.BccLen(); i++ {
		var bcc *url.URL
		if o.IsBccObject(i) {
			obj := o.GetBccObject(i)
			if !obj.HasId() {
				err = fmt.Errorf("bcc object missing id")
				return
			}
			bcc = obj.GetId()
		} else if o.IsBccLink(i) {
			href := o.GetBccLink(i)
			if !href.HasHref() {
				err = fmt.Errorf("bcc link missing href")
				return
			}
			bcc = href.GetHref()
		} else if o.IsBccIRI(i) {
			bcc = o.GetBccIRI(i)
		}
		c.AppendBccIRI(bcc)
	}
	for i := 0; i < o.AudienceLen(); i++ {
		var audience *url.URL
		if o.IsAudienceObject(i) {
			obj := o.GetAudienceObject(i)
			if !obj.HasId() {
				err = fmt.Errorf("audience object missing id")
				return
			}
			audience = obj.GetId()
		} else if o.IsAudienceLink(i) {
			href := o.GetAudienceLink(i)
			if !href.HasHref() {
				err = fmt.Errorf("audience link missing href")
				return
			}
			audience = href.GetHref()
		} else if o.IsAudienceIRI(i) {
			audience = o.GetAudienceIRI(i)
		}
		c.AppendAudienceIRI(audience)
	}
	return
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
				to[i][id.String()] = o.GetToObject(j)
			} else if o.IsToLink(j) {
				id := o.GetToLink(j).GetHref()
				to[i][id.String()] = o.GetToLink(j)
			} else if o.IsToIRI(j) {
				id := o.GetToIRI(j)
				to[i][id.String()] = id
			}
		}
	}
	toActivity := make(map[string]interface{})
	for i := 0; i < a.ToLen(); i++ {
		if a.IsToObject(i) {
			id := a.GetToObject(i).GetId()
			toActivity[id.String()] = a.GetToObject(i)
		} else if a.IsToLink(i) {
			id := a.GetToLink(i).GetHref()
			toActivity[id.String()] = a.GetToLink(i)
		} else if a.IsToIRI(i) {
			id := a.GetToIRI(i)
			toActivity[id.String()] = id
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
				bto[i][id.String()] = o.GetBtoObject(j)
			} else if o.IsBtoLink(j) {
				id := o.GetBtoLink(j).GetHref()
				bto[i][id.String()] = o.GetBtoLink(j)
			} else if o.IsBtoIRI(j) {
				id := o.GetBtoIRI(j)
				bto[i][id.String()] = id
			}
		}
	}
	btoActivity := make(map[string]interface{})
	for i := 0; i < a.BtoLen(); i++ {
		if a.IsBtoObject(i) {
			id := a.GetBtoObject(i).GetId()
			btoActivity[id.String()] = a.GetBtoObject(i)
		} else if a.IsBtoLink(i) {
			id := a.GetBtoLink(i).GetHref()
			btoActivity[id.String()] = a.GetBtoLink(i)
		} else if a.IsBtoIRI(i) {
			id := a.GetBtoIRI(i)
			btoActivity[id.String()] = id
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
				cc[i][id.String()] = o.GetCcObject(j)
			} else if o.IsCcLink(j) {
				id := o.GetCcLink(j).GetHref()
				cc[i][id.String()] = o.GetCcLink(j)
			} else if o.IsCcIRI(j) {
				id := o.GetCcIRI(j)
				cc[i][id.String()] = id
			}
		}
	}
	ccActivity := make(map[string]interface{})
	for i := 0; i < a.CcLen(); i++ {
		if a.IsCcObject(i) {
			id := a.GetCcObject(i).GetId()
			ccActivity[id.String()] = a.GetCcObject(i)
		} else if a.IsCcLink(i) {
			id := a.GetCcLink(i).GetHref()
			ccActivity[id.String()] = a.GetCcLink(i)
		} else if a.IsCcIRI(i) {
			id := a.GetCcIRI(i)
			ccActivity[id.String()] = id
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
				bcc[i][id.String()] = o.GetBccObject(j)
			} else if o.IsBccLink(j) {
				id := o.GetBccLink(j).GetHref()
				bcc[i][id.String()] = o.GetBccLink(j)
			} else if o.IsBccIRI(j) {
				id := o.GetBccIRI(j)
				bcc[i][id.String()] = id
			}
		}
	}
	bccActivity := make(map[string]interface{})
	for i := 0; i < a.BccLen(); i++ {
		if a.IsBccObject(i) {
			id := a.GetBccObject(i).GetId()
			bccActivity[id.String()] = a.GetBccObject(i)
		} else if a.IsBccLink(i) {
			id := a.GetBccLink(i).GetHref()
			bccActivity[id.String()] = a.GetBccLink(i)
		} else if a.IsBccIRI(i) {
			id := a.GetBccIRI(i)
			bccActivity[id.String()] = id
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
				audience[i][id.String()] = o.GetAudienceObject(j)
			} else if o.IsAudienceLink(j) {
				id := o.GetAudienceLink(j).GetHref()
				audience[i][id.String()] = o.GetAudienceLink(j)
			} else if o.IsAudienceIRI(j) {
				id := o.GetAudienceIRI(j)
				audience[i][id.String()] = id
			}
		}
	}
	audienceActivity := make(map[string]interface{})
	for i := 0; i < a.AudienceLen(); i++ {
		if a.IsAudienceObject(i) {
			id := a.GetAudienceObject(i).GetId()
			audienceActivity[id.String()] = a.GetAudienceObject(i)
		} else if a.IsAudienceLink(i) {
			id := a.GetAudienceLink(i).GetHref()
			audienceActivity[id.String()] = a.GetAudienceLink(i)
		} else if a.IsAudienceIRI(i) {
			id := a.GetAudienceIRI(i)
			audienceActivity[id.String()] = id
		}
	}
	// Next, add activity recipients to all objects if not already present
	for k, v := range toActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := to[i][k]; !ok {
				var to *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("to object missing id")
					}
					to = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("to link missing href")
					}
					to = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					to = vIRI
				}
				a.GetObject(i).AppendToIRI(to)
			}
		}
	}
	for k, v := range btoActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := bto[i][k]; !ok {
				var bto *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("bto object missing id")
					}
					bto = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("bto link missing href")
					}
					bto = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					bto = vIRI
				}
				a.GetObject(i).AppendBtoIRI(bto)
			}
		}
	}
	for k, v := range ccActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := cc[i][k]; !ok {
				var cc *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("cc object missing id")
					}
					cc = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("cc link missing href")
					}
					cc = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					cc = vIRI
				}
				a.GetObject(i).AppendCcIRI(cc)
			}
		}
	}
	for k, v := range bccActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := bcc[i][k]; !ok {
				var bcc *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("bcc object missing id")
					}
					bcc = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("bcc link missing href")
					}
					bcc = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					bcc = vIRI
				}
				a.GetObject(i).AppendBccIRI(bcc)
			}
		}
	}
	for k, v := range audienceActivity {
		for i := 0; i < a.ObjectLen(); i++ {
			if _, ok := audience[i][k]; !ok {
				var activity *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("activity object missing id")
					}
					activity = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("activity link missing href")
					}
					activity = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					activity = vIRI
				}
				a.GetObject(i).AppendAudienceIRI(activity)
			}
		}
	}
	// Finally, add all the objects' recipients to the activity if not
	// already present.
	for i := 0; i < a.ObjectLen(); i++ {
		for k, v := range to[i] {
			if _, ok := toActivity[k]; !ok {
				var to *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("to object missing id")
					}
					to = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("to link missing href")
					}
					to = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					to = vIRI
				}
				a.AppendToIRI(to)
			}
		}
		for k, v := range bto[i] {
			if _, ok := btoActivity[k]; !ok {
				var bto *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("bto object missing id")
					}
					bto = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("bto link missing href")
					}
					bto = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					bto = vIRI
				}
				a.AppendBtoIRI(bto)
			}
		}
		for k, v := range cc[i] {
			if _, ok := ccActivity[k]; !ok {
				var cc *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("cc object missing id")
					}
					cc = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("cc link missing href")
					}
					cc = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					cc = vIRI
				}
				a.AppendCcIRI(cc)
			}
		}
		for k, v := range bcc[i] {
			if _, ok := bccActivity[k]; !ok {
				var bcc *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("bcc object missing id")
					}
					bcc = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("bcc link missing href")
					}
					bcc = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					bcc = vIRI
				}
				a.AppendBccIRI(bcc)
			}
		}
		for k, v := range audience[i] {
			if _, ok := audienceActivity[k]; !ok {
				var audience *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("audience object missing id")
					}
					audience = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("audience link missing href")
					}
					audience = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					audience = vIRI
				}
				a.AppendAudienceIRI(audience)
			}
		}
	}
	return nil
}

// TODO: (Section 7) HTTP caching mechanisms [RFC7234] SHOULD be respected when appropriate, both when receiving responses from other servers as well as sending responses to other servers.

// deliver will complete the peer-to-peer sending of a federated message to
// another server.
func (f *federator) deliver(obj vocab.ActivityType, boxIRI *url.URL) error {
	recipients, err := f.prepare(boxIRI, obj)
	if err != nil {
		return err
	}
	creds := &creds{}
	creds.signer, err = f.FederateAPI.NewSigner()
	if err != nil {
		return err
	}
	creds.privKey, creds.pubKeyId, err = f.FederateAPI.PrivateKey(boxIRI)
	if err != nil {
		return err
	}
	return f.deliverToRecipients(obj, recipients, creds)
}

// deliverToRecipients will take a prepared Activity and send it to specific
// recipients without examining the activity.
func (f *federator) deliverToRecipients(obj vocab.ActivityType, recipients []*url.URL, creds *creds) error {
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
		f.deliverer.Do(b, to, func(b []byte, u *url.URL) error {
			return postToOutbox(f.Client, b, u, f.Agent, creds, f.Clock)
		})
	}
	return nil
}

// prepare takes a deliverableObject and returns a list of the proper recipient
// target URIs. Additionally, the deliverableObject will have any hidden
// hidden recipients ("bto" and "bcc") stripped from it.
func (c *federator) prepare(boxIRI *url.URL, o deliverableObject) ([]*url.URL, error) {
	// Get inboxes of recipients
	var r []*url.URL
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
	receiverActors, err := c.resolveInboxes(boxIRI, r, 0, c.MaxDeliveryDepth)
	if err != nil {
		return nil, err
	}
	targets, err := getInboxes(receiverActors)
	if err != nil {
		return nil, err
	}
	// Get inboxes of sender(s)
	senderActors, err := c.resolveInboxes(boxIRI, getActorsAttributedToURI(o), 0, c.MaxDeliveryDepth)
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
func (c *federator) resolveInboxes(boxIRI *url.URL, r []*url.URL, depth int, max int) ([]actor, error) {
	if depth >= max {
		return nil, nil
	}
	a := make([]actor, 0, len(r))
	for _, u := range r {
		// Do not retry here -- if a dereference fails, then fail the
		// entire delivery.
		actor, co, oc, cp, ocp, cr, err := c.dereferenceForResolvingInboxes(u, boxIRI, nil)
		if err != nil {
			return nil, err
		}
		var uris []*url.URL
		if co != nil {
			uris := getURIsInItemer(co.Raw())
			actors, err := c.resolveInboxes(boxIRI, uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if oc != nil {
			uris := getURIsInOrderedItemer(oc.Raw())
			actors, err := c.resolveInboxes(boxIRI, uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if cp != nil {
			cb := func(c vocab.CollectionPageType) error {
				uris = getURIsInItemer(c)
				return nil
			}
			err := doForCollectionPage(c.Client, c.Agent, cb, cp.Raw(), cr, c.Clock)
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(boxIRI, uris, depth+1, max)
			if err != nil {
				return nil, err
			}
			a = append(a, actors...)
		} else if ocp != nil {
			cb := func(c vocab.OrderedCollectionPageType) error {
				uris = getURIsInOrderedItemer(c)
				return nil
			}
			err := doForOrderedCollectionPage(c.Client, c.Agent, cb, ocp.Raw(), cr, c.Clock)
			if err != nil {
				return nil, err
			}
			actors, err := c.resolveInboxes(boxIRI, uris, depth+1, max)
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

func (c *federator) dereferenceForResolvingInboxes(u, boxIRI *url.URL, cr *creds) (actor actor, co *streams.Collection, oc *streams.OrderedCollection, cp *streams.CollectionPage, ocp *streams.OrderedCollectionPage, cred *creds, err error) {
	// To pass back to calling function, since may be set recursively:
	cred = cr
	var resp []byte
	resp, err = dereference(c.Client, u, c.Agent, cr, c.Clock)
	if err != nil {
		return
	}
	var m map[string]interface{}
	if err = json.Unmarshal(resp, &m); err != nil {
		return
	}
	// Set AT MOST one of: co, oc, cp, ocp
	// If none of these are set, attempt to use actor
	if err = toActorCollectionResolver(&actor, &co, &oc, &cp, &ocp).Deserialize(m); err != nil {
		return
	}
	// If a recipient is a Collection or OrderedCollection, then the
	// server MUST dereference the collection, WITH the user's
	// credentials.
	//
	// Note that this also applies to CollectionPage and
	// OrderedCollectionPage.
	//
	// This jumps to the label above ONLY if we have not yet set the
	// creds -- which happens at most once.
	if (co != nil || oc != nil || cp != nil || ocp != nil) && cr == nil {
		cr = &creds{}
		cr.signer, err = c.FederateAPI.NewSigner()
		if err != nil {
			return
		}
		cr.privKey, cr.pubKeyId, err = c.FederateAPI.PrivateKey(boxIRI)
		if err != nil {
			return
		}
		return c.dereferenceForResolvingInboxes(u, boxIRI, cr)
	}
	return
}

// getInboxes extracts the 'inbox' IRIs from actors.
func getInboxes(a []actor) ([]*url.URL, error) {
	var u []*url.URL
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
func getActorsAttributedToURI(a actorObject) []*url.URL {
	var u []*url.URL
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
		} else if o.IsBccIRI(0) {
			o.RemoveBccIRI(0)
		}
	}
}

// dedupeIRIs will deduplicate final inbox IRIs. The ignore list is applied to
// the final list
func dedupeIRIs(recipients, ignored []*url.URL) (out []*url.URL) {
	ignoredMap := make(map[string]bool, len(ignored))
	for _, elem := range ignored {
		ignoredMap[elem.String()] = true
	}
	outMap := make(map[string]bool, len(recipients))
	for _, k := range recipients {
		kStr := k.String()
		if !ignoredMap[kStr] && !outMap[kStr] {
			out = append(out, k)
			outMap[kStr] = true
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
			id = iri.String()
		} else if oc.IsOrderedItemsLink(i) {
			removeFn = oc.RemoveOrderedItemsLink
			iri := oc.GetOrderedItemsLink(i).GetId()
			id = iri.String()
		} else if oc.IsOrderedItemsIRI(i) {
			removeFn = oc.RemoveOrderedItemsIRI
			id = oc.GetOrderedItemsIRI(i).String()
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
func filterURLs(u []*url.URL, fn func(s string) bool) []*url.URL {
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

func getToIRIs(o deliverableObject) []*url.URL {
	var r []*url.URL
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

func getBToIRIs(o deliverableObject) []*url.URL {
	var r []*url.URL
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

func getCcIRIs(o deliverableObject) []*url.URL {
	var r []*url.URL
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

func getBccIRIs(o deliverableObject) []*url.URL {
	var r []*url.URL
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

func getAudienceIRIs(o deliverableObject) []*url.URL {
	var r []*url.URL
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
func doForCollectionPage(h HttpClient, agent string, cb func(c vocab.CollectionPageType) error, c vocab.CollectionPageType, creds *creds, clock Clock) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForCollectionPage(h, agent, cb, c.GetNextCollectionPage(), creds, clock)
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent, creds, clock)
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
				return doForCollectionPage(h, agent, cb, next.Raw(), creds, clock)
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent, creds, clock)
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
			return doForCollectionPage(h, agent, cb, next.Raw(), creds, clock)
		}
	}
	return nil
}

// doForOrderedCollectionPage applies a function over a collection and its
// subsequent pages recursively. It returns the first non-nil error it
// encounters.
func doForOrderedCollectionPage(h HttpClient, agent string, cb func(c vocab.OrderedCollectionPageType) error, c vocab.OrderedCollectionPageType, creds *creds, clock Clock) error {
	err := cb(c)
	if err != nil {
		return err
	}
	if c.IsNextOrderedCollectionPage() {
		// Handle this one weird trick that other peers HATE federating
		// with.
		return doForOrderedCollectionPage(h, agent, cb, c.GetNextOrderedCollectionPage(), creds, clock)
	} else if c.IsNextLink() {
		l := c.GetNextLink()
		if l.HasHref() {
			u := l.GetHref()
			resp, err := dereference(h, u, agent, creds, clock)
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
				return doForOrderedCollectionPage(h, agent, cb, next.Raw(), creds, clock)
			}
		}
	} else if c.IsNextIRI() {
		u := c.GetNextIRI()
		resp, err := dereference(h, u, agent, creds, clock)
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
			return doForOrderedCollectionPage(h, agent, cb, next.Raw(), creds, clock)
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
	GetItemsIRI(index int) (v *url.URL)
}

var _ itemer = &vocab.Collection{}
var _ itemer = &vocab.CollectionPage{}

// getURIsInItemer will extract 'items' from the provided Collection or
// CollectionPage.
func getURIsInItemer(i itemer) []*url.URL {
	var u []*url.URL
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
	GetOrderedItemsIRI(index int) (v *url.URL)
}

var _ orderedItemer = &vocab.OrderedCollection{}
var _ orderedItemer = &vocab.OrderedCollectionPage{}

// getURIsInOrderedItemer will extract 'items' from the provided
// OrderedCollection or OrderedCollectionPage.
func getURIsInOrderedItemer(i orderedItemer) []*url.URL {
	var u []*url.URL
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
	GetObjectIRI(index int) (v *url.URL)
}

func getObjectIds(a activityWithObject) (u []*url.URL, e error) {
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
	GetTargetIRI(index int) (v *url.URL)
}

func getTargetIds(a activityWithTarget) (u []*url.URL, e error) {
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

func removeCollectionItemWithId(c vocab.CollectionType, iri *url.URL) {
	for i := 0; i < c.ItemsLen(); i++ {
		if c.IsItemsObject(i) {
			o := c.GetItemsObject(i)
			if !o.HasId() {
				continue
			}
			if *o.GetId() == *iri {
				c.RemoveItemsObject(i)
				return
			}
		} else if c.IsItemsLink(i) {
			l := c.GetItemsLink(i)
			if !l.HasHref() {
				continue
			}
			if *l.GetHref() == *iri {
				c.RemoveItemsLink(i)
				return
			}
		} else if c.IsItemsIRI(i) {
			if *c.GetItemsIRI(i) == *iri {
				c.RemoveItemsIRI(i)
				return
			}
		}
	}
}

func removeOrderedCollectionItemWithId(c vocab.OrderedCollectionType, iri *url.URL) {
	for i := 0; i < c.OrderedItemsLen(); i++ {
		if c.IsOrderedItemsObject(i) {
			o := c.GetOrderedItemsObject(i)
			if !o.HasId() {
				continue
			}
			if *o.GetId() == *iri {
				c.RemoveOrderedItemsObject(i)
				return
			}
		} else if c.IsOrderedItemsLink(i) {
			l := c.GetOrderedItemsLink(i)
			if !l.HasHref() {
				continue
			}
			if *l.GetHref() == *iri {
				c.RemoveOrderedItemsLink(i)
				return
			}
		} else if c.IsOrderedItemsIRI(i) {
			if *c.GetOrderedItemsIRI(i) == *iri {
				c.RemoveOrderedItemsIRI(i)
				return
			}
		}
	}
}

// toTombstone creates a Tombstone for the given object.
func toTombstone(obj vocab.ObjectType, id *url.URL, now time.Time) vocab.TombstoneType {
	tomb := &vocab.Tombstone{}
	tomb.SetId(id)
	for i := 0; i < obj.TypeLen(); i++ {
		if s, ok := obj.GetType(i).(string); ok {
			tomb.AppendFormerTypeString(s)
		} else if fObj, ok := obj.GetType(i).(vocab.ObjectType); ok {
			tomb.AppendFormerTypeObject(fObj)
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

type getActorCollectionFn func(actor vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (isIRI bool, e error)

func (f *federator) addAllObjectsToActorCollection(ctx context.Context, getter getActorCollectionFn, c vocab.ActivityType, prepend bool) error {
	for i := 0; i < c.ActorLen(); i++ {
		var iri *url.URL
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
			// TODO: Fetch or just store
			continue
		}
		var actor vocab.ObjectType
		pObj, err := f.App.Get(ctx, iri, ReadWrite)
		if err != nil {
			return err
		}
		ok := false
		actor, ok = pObj.(vocab.ObjectType)
		if !ok {
			return fmt.Errorf("actor is not vocab.ObjectType")
		}
		// Obtain ordered/unordered collection
		var lc vocab.CollectionType
		var loc vocab.OrderedCollectionType
		isIRI := false
		if isIRI, err = getter(actor, &lc, &loc); err != nil {
			return err
		}
		// Duplication detection
		var iriSet map[string]bool
		if lc != nil {
			iriSet, err = getIRISetFromItems(lc)
		} else if loc != nil {
			iriSet, err = getIRISetFromOrderedItems(loc)
		}
		if err != nil {
			return err
		}
		// Add object to collection if not a duplicate
		for i := 0; i < c.ObjectLen(); i++ {
			var iri *url.URL
			if c.IsObjectIRI(i) {
				iri = c.GetObjectIRI(i)
			} else if c.IsObject(i) {
				obj := c.GetObject(i)
				if !obj.HasId() {
					return fmt.Errorf("object at index %d has no id", i)
				}
				iri = obj.GetId()
			}
			if iriSet[iri.String()] {
				continue
			}
			if lc != nil {
				if prepend {
					lc.PrependItemsIRI(iri)
				} else {
					lc.AppendItemsIRI(iri)
				}
			} else if loc != nil {
				if prepend {
					loc.PrependOrderedItemsIRI(iri)
				} else {
					loc.AppendOrderedItemsIRI(iri)
				}
			}
		}
		if isIRI {
			if lc != nil {
				err = f.App.Set(ctx, lc)
			} else if loc != nil {
				err = f.App.Set(ctx, loc)
			}
			if err != nil {
				return err
			}
		} else if err := f.App.Set(ctx, actor); err != nil {
			return err
		}
	}
	return nil
}

type getObjectCollectionFn func(object vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (isIRI bool, e error)

func (f *federator) addAllActorsToObjectCollection(ctx context.Context, getter getObjectCollectionFn, c vocab.ActivityType, prepend bool) (bool, error) {
	ownsAny := false
	for i := 0; i < c.ObjectLen(); i++ {
		var iri *url.URL
		if c.IsObject(i) {
			obj := c.GetObject(i)
			if !obj.HasId() {
				return ownsAny, fmt.Errorf("object does not have id")
			}
			iri = obj.GetId()
		} else if c.IsObjectIRI(i) {
			iri = c.GetObjectIRI(i)
		}
		if !f.App.Owns(ctx, iri) {
			// TODO: Fetch or just store
			continue
		}
		ownsAny = true
		var object vocab.ObjectType
		pObj, err := f.App.Get(ctx, iri, ReadWrite)
		if err != nil {
			return ownsAny, err
		}
		ok := false
		object, ok = pObj.(vocab.ObjectType)
		if !ok {
			return ownsAny, fmt.Errorf("object is not vocab.ObjectType")
		}
		// Obtain ordered/unordered collection
		var lc vocab.CollectionType
		var loc vocab.OrderedCollectionType
		isIRI := false
		if isIRI, err = getter(object, &lc, &loc); err != nil {
			return ownsAny, err
		}
		// Duplication detection
		var iriSet map[string]bool
		if lc != nil {
			iriSet, err = getIRISetFromItems(lc)
		} else if loc != nil {
			iriSet, err = getIRISetFromOrderedItems(loc)
		}
		if err != nil {
			return ownsAny, err
		}
		// Add actor to collection if not a duplicate
		for i := 0; i < c.ActorLen(); i++ {
			var iri *url.URL
			if c.IsActorIRI(i) {
				iri = c.GetActorIRI(i)
			} else if c.IsActorObject(i) {
				obj := c.GetActorObject(i)
				if !obj.HasId() {
					return ownsAny, fmt.Errorf("actor object at index %d has no id", i)
				}
				iri = obj.GetId()
			} else if c.IsActorLink(i) {
				l := c.GetActorLink(i)
				if !l.HasHref() {
					return ownsAny, fmt.Errorf("actor link at index %d has no id", i)
				}
				iri = l.GetHref()
			}
			if iriSet[iri.String()] {
				continue
			}
			if lc != nil {
				if prepend {
					lc.AppendItemsIRI(iri)
				} else {
					lc.PrependItemsIRI(iri)
				}
			} else if loc != nil {
				if prepend {
					loc.PrependOrderedItemsIRI(iri)
				} else {
					loc.AppendOrderedItemsIRI(iri)
				}
			}
		}
		if isIRI {
			if lc != nil {
				err = f.App.Set(ctx, lc)
			} else if loc != nil {
				err = f.App.Set(ctx, loc)
			}
			if err != nil {
				return ownsAny, err
			}
		} else if err := f.App.Set(ctx, object); err != nil {
			return ownsAny, err
		}
	}
	return ownsAny, nil
}

func (f *federator) addActivityToObjectCollection(ctx context.Context, getter getObjectCollectionFn, c vocab.ActivityType, prepend bool) (bool, error) {
	ownsAny := false
	for i := 0; i < c.ObjectLen(); i++ {
		var objIri *url.URL
		if c.IsObject(i) {
			obj := c.GetObject(i)
			if !obj.HasId() {
				return ownsAny, fmt.Errorf("object does not have id")
			}
			objIri = obj.GetId()
		} else if c.IsObjectIRI(i) {
			objIri = c.GetObjectIRI(i)
		}
		if !f.App.Owns(ctx, objIri) {
			// TODO: Fetch or just store
			continue
		}
		ownsAny = true
		var object vocab.ObjectType
		pObj, err := f.App.Get(ctx, objIri, ReadWrite)
		if err != nil {
			return ownsAny, err
		}
		ok := false
		object, ok = pObj.(vocab.ObjectType)
		if !ok {
			return ownsAny, fmt.Errorf("object is not vocab.ObjectType")
		}
		// Obtain ordered/unordered collection
		var lc vocab.CollectionType
		var loc vocab.OrderedCollectionType
		isIRI := false
		if isIRI, err = getter(object, &lc, &loc); err != nil {
			return ownsAny, err
		}
		// Duplication detection
		var iriSet map[string]bool
		if lc != nil {
			iriSet, err = getIRISetFromItems(lc)
		} else if loc != nil {
			iriSet, err = getIRISetFromOrderedItems(loc)
		}
		if err != nil {
			return ownsAny, err
		}
		// Add activity to collection if not a duplicate
		if !c.HasId() {
			return ownsAny, fmt.Errorf("activity has no id")
		}
		iri := c.GetId()
		if iriSet[iri.String()] {
			continue
		}
		if lc != nil {
			if prepend {
				lc.AppendItemsIRI(iri)
			} else {
				lc.PrependItemsIRI(iri)
			}
		} else if loc != nil {
			if prepend {
				loc.PrependOrderedItemsIRI(iri)
			} else {
				loc.AppendOrderedItemsIRI(iri)
			}
		}
		if isIRI {
			if lc != nil {
				err = f.App.Set(ctx, lc)
			} else if loc != nil {
				err = f.App.Set(ctx, loc)
			}
			if err != nil {
				return ownsAny, err
			}
		} else if err := f.App.Set(ctx, object); err != nil {
			return ownsAny, err
		}
	}
	return ownsAny, nil
}

func (f *federator) ownsAnyObjects(c context.Context, a vocab.ActivityType) (bool, error) {
	var iris []*url.URL
	for i := 0; i < a.ObjectLen(); i++ {
		if a.IsObject(i) {
			obj := a.GetObject(i)
			if !obj.HasId() {
				return false, fmt.Errorf("object missing id")
			}
			iris = append(iris, obj.GetId())
		} else if a.IsObjectIRI(i) {
			iris = append(iris, a.GetObjectIRI(i))
		}
	}
	return f.ownsAnyIRIs(c, iris), nil
}

func (f *federator) addToOutbox(c context.Context, r *http.Request, m map[string]interface{}) error {
	outbox, err := f.App.GetOutbox(c, r, ReadWrite)
	if err != nil {
		return err
	}
	activity, err := toAnyActivity(m)
	if err != nil {
		return err
	}
	if err := f.App.Set(c, activity); err != nil {
		return err
	}
	if !activity.HasId() {
		return fmt.Errorf("activity missing id")
	}
	outbox.PrependOrderedItemsIRI(activity.GetId())
	return f.App.Set(c, outbox)
}

func (f *federator) addToInboxIfNew(c context.Context, r *http.Request, m map[string]interface{}, callback func() error) error {
	inbox, err := f.App.GetInbox(c, r, ReadWrite)
	if err != nil {
		return err
	}
	activity, err := toAnyActivity(m)
	if err != nil {
		return err
	}
	iriSet, err := getIRISetFromOrderedItems(inbox)
	if err != nil {
		return err
	}
	if !activity.HasId() {
		return fmt.Errorf("activity missing id")
	}
	if !iriSet[activity.GetId().String()] {
		if err := callback(); err != nil {
			return err
		}
		inbox.PrependOrderedItemsIRI(activity.GetId())
		return f.App.Set(c, inbox)
	}
	return nil
}

// Note: This is a mechanism for causing other victim servers to DDOS
// or forward spam on a malicious user's behalf. The trick is a simple
// one: Reply to a user, and CC a ton of 'follower' collections owned
// by the victim server. Bonus points for listing more 'follower'
// collections from other popular instances as well. Leveraging the
// Inbox Forwarding mechanism, a storm of messages will ensue.
//
// I don't want users of this library to be vulnerable to this kind of
// spam/DDOS storm. So here we allow the client application to filter
// out recipient collections.
func (f *federator) inboxForwarding(c context.Context, m map[string]interface{}) error {
	a, err := toAnyActivity(m)
	if err != nil {
		return err
	}
	// 1. Must be first time we have seen this Activity.
	if ok, err := f.App.Has(c, a.GetId()); err != nil {
		return err
	} else if ok {
		return nil
	}
	// 2. The values of 'to', 'cc', or 'audience' are Collections owned by
	//    this server.
	var r []*url.URL
	r = append(r, getToIRIs(a)...)
	r = append(r, getCcIRIs(a)...)
	r = append(r, getAudienceIRIs(a)...)
	var myIRIs []*url.URL
	col := make(map[string]vocab.CollectionType, 0)
	oCol := make(map[string]vocab.OrderedCollectionType, 0)
	for _, iri := range r {
		if ok, err := f.App.Has(c, iri); err != nil {
			return err
		} else if !ok {
			continue
		}
		obj, err := f.App.Get(c, iri, Read)
		if err != nil {
			return err
		}
		if c, ok := obj.(vocab.CollectionType); ok {
			col[iri.String()] = c
			myIRIs = append(myIRIs, iri)
		} else if oc, ok := obj.(vocab.OrderedCollectionType); ok {
			oCol[iri.String()] = oc
			myIRIs = append(myIRIs, iri)
		}
	}
	if len(myIRIs) == 0 {
		return nil
	}
	// 3. The values of 'inReplyTo', 'object', 'target', or 'tag' are owned
	//    by this server.
	ownsValue := false
	objs, l, iris := getInboxForwardingValues(a)
	for _, obj := range objs {
		if f.hasInboxForwardingValues(c, 0, f.MaxInboxForwardingDepth, obj) {
			ownsValue = true
			break
		}
	}
	if !ownsValue && f.ownsAnyLinks(c, l) {
		ownsValue = true
	}
	if !ownsValue && f.ownsAnyIRIs(c, iris) {
		ownsValue = true
	}
	if !ownsValue {
		return nil
	}
	// Do the inbox forwarding since the above conditions hold true. Support
	// the behavior of letting the application filter out the resulting
	// collections to be targeted.
	toSend, err := f.FederateAPI.FilterForwarding(c, a, myIRIs)
	if err != nil {
		return err
	}
	recipients := make([]*url.URL, 0, len(toSend))
	for _, iri := range toSend {
		if c, ok := col[iri.String()]; ok {
			for i := 0; i < c.ItemsLen(); i++ {
				if c.IsItemsObject(i) {
					obj := c.GetItemsObject(i)
					if obj.HasId() {
						recipients = append(recipients, obj.GetId())
					}
				} else if c.IsItemsLink(i) {
					l := c.GetItemsLink(i)
					if l.HasHref() {
						recipients = append(recipients, l.GetHref())
					}
				} else if c.IsItemsIRI(i) {
					recipients = append(recipients, c.GetItemsIRI(i))
				}
			}
		} else if oc, ok := oCol[iri.String()]; ok {
			for i := 0; i < oc.OrderedItemsLen(); i++ {
				if oc.IsOrderedItemsObject(i) {
					obj := oc.GetOrderedItemsObject(i)
					if obj.HasId() {
						recipients = append(recipients, obj.GetId())
					}
				} else if oc.IsOrderedItemsLink(i) {
					l := oc.GetItemsLink(i)
					if l.HasHref() {
						recipients = append(recipients, l.GetHref())
					}
				} else if oc.IsOrderedItemsIRI(i) {
					recipients = append(recipients, oc.GetOrderedItemsIRI(i))
				}
			}
		}
	}
	return f.deliverToRecipients(a, recipients, nil)
}

// Given an 'inReplyTo', 'object', 'target', or 'tag' object, recursively
// examines those same values to determine if the app owns any, up to a maximum
// depth.
func (f *federator) hasInboxForwardingValues(c context.Context, depth, maxDepth int, o vocab.ObjectType) bool {
	if depth == maxDepth {
		return false
	}
	if f.App.Owns(c, o.GetId()) {
		return true
	}
	objs, l, iris := getInboxForwardingValues(o)
	for _, obj := range objs {
		if f.hasInboxForwardingValues(c, depth+1, maxDepth, obj) {
			return true
		}
	}
	if f.ownsAnyLinks(c, l) {
		return true
	}
	return f.ownsAnyIRIs(c, iris)
}

func (f *federator) ownsAnyIRIs(c context.Context, iris []*url.URL) bool {
	for _, iri := range iris {
		if f.App.Owns(c, iri) {
			return true
		}
		// TODO: Dereference the IRI
	}
	return false
}

func (f *federator) ownsAnyLinks(c context.Context, links []vocab.LinkType) bool {
	for _, link := range links {
		if !link.HasHref() {
			continue
		}
		href := link.GetHref()
		if f.App.Owns(c, href) {
			return true
		}
		// TODO: Dereference the IRI
	}
	return false
}

func (f *federator) ensureActivityOriginMatchesObjects(a vocab.ActivityType) error {
	if !a.HasId() {
		return fmt.Errorf("activity has no iri")
	}
	originIRI := a.GetId()
	originHost := originIRI.Host
	for i := 0; i < a.ObjectLen(); i++ {
		if a.IsObject(i) {
			obj := a.GetObject(i)
			if !obj.HasId() {
				return fmt.Errorf("object at index %d has no id", i)
			}
			iri := obj.GetId()
			if originHost != iri.Host {
				return fmt.Errorf("object %q: not in activity origin", iri)
			}
		} else if a.IsObjectIRI(i) {
			iri := a.GetObjectIRI(i)
			if originHost != iri.Host {
				return fmt.Errorf("object %q: not in activity origin", iri)
			}
		}
	}
	return nil
}

func (f *federator) ensureActivityActorsMatchObjectActors(a vocab.ActivityType) error {
	actorSet := make(map[string]bool, a.ActorLen())
	for i := 0; i < a.ActorLen(); i++ {
		if a.IsActorObject(i) {
			obj := a.GetActorObject(i)
			if !obj.HasId() {
				return fmt.Errorf("actor object at index %d has no id", i)
			}
			actorSet[obj.GetId().String()] = true
		} else if a.IsActorLink(i) {
			l := a.GetActorLink(i)
			if !l.HasHref() {
				return fmt.Errorf("actor link at index %d has no href", i)
			}
			actorSet[l.GetHref().String()] = true
		} else if a.IsActorIRI(i) {
			actorSet[a.GetActorIRI(i).String()] = true
		}
	}
	objectActors := make(map[string]bool, a.ObjectLen())
	for i := 0; i < a.ObjectLen(); i++ {
		if a.IsObject(i) {
			obj := a.GetObject(i)
			if !obj.HasId() {
				return fmt.Errorf("object at index %d has no id", i)
			}
			objectActivity, ok := obj.(vocab.ActivityType)
			if !ok {
				return fmt.Errorf("object at index %d is not an activity", i)
			}
			for j := 0; j < objectActivity.ActorLen(); j++ {
				if objectActivity.IsActorObject(j) {
					obj := objectActivity.GetActorObject(j)
					if !obj.HasId() {
						return fmt.Errorf("actor object at index (%d,%d) has no id", i, j)
					}
					objectActors[obj.GetId().String()] = true
				} else if objectActivity.IsActorLink(j) {
					l := objectActivity.GetActorLink(j)
					if !l.HasHref() {
						return fmt.Errorf("actor link at index (%d,%d) has no href", i, j)
					}
					objectActors[l.GetHref().String()] = true
				} else if objectActivity.IsActorIRI(j) {
					objectActors[objectActivity.GetActorIRI(j).String()] = true
				}
			}
		} else if a.IsObjectIRI(i) {
			// TODO: Dereference IRI
			iri := a.GetObjectIRI(i)
			return fmt.Errorf("unimplemented: fetching IRI for UNDO verification of ownership of %q", iri)
		}
	}
	for k := range objectActors {
		if !actorSet[k] {
			return fmt.Errorf("at least 1 activity actors missing: %q", k)
		}
	}
	return nil
}

func getInboxForwardingValues(o vocab.ObjectType) (objs []vocab.ObjectType, l []vocab.LinkType, iri []*url.URL) {
	// 'inReplyTo'
	for i := 0; i < o.InReplyToLen(); i++ {
		if o.IsInReplyToObject(i) {
			objs = append(objs, o.GetInReplyToObject(i))
		} else if o.IsInReplyToLink(i) {
			l = append(l, o.GetInReplyToLink(i))
		} else if o.IsInReplyToIRI(i) {
			iri = append(iri, o.GetInReplyToIRI(i))
		}
	}
	// 'tag'
	for i := 0; i < o.TagLen(); i++ {
		if o.IsTagObject(i) {
			objs = append(objs, o.GetTagObject(i))
		} else if o.IsTagLink(i) {
			l = append(l, o.GetTagLink(i))
		} else if o.IsTagIRI(i) {
			iri = append(iri, o.GetTagIRI(i))
		}
	}
	if a, ok := o.(vocab.ActivityType); ok {
		// 'object'
		for i := 0; i < a.ObjectLen(); i++ {
			if a.IsObject(i) {
				objs = append(objs, a.GetObject(i))
			} else if a.IsObjectIRI(i) {
				iri = append(iri, a.GetObjectIRI(i))
			}
		}
		// 'target'
		for i := 0; i < a.TargetLen(); i++ {
			if a.IsTargetObject(i) {
				objs = append(objs, a.GetTargetObject(i))
			} else if a.IsTargetLink(i) {
				l = append(l, a.GetTargetLink(i))
			} else if a.IsTargetIRI(i) {
				iri = append(iri, a.GetTargetIRI(i))
			}
		}
	}
	return
}

// Fetches an "object" on a raw JSON map of an Activity with the matching 'id'
// field. If there is no object matching the IRI, or the object just is an IRI,
// or the object wth the matching id is not in the array of objects, then a nil
// map is returned.
func getRawObject(m map[string]interface{}, id string) map[string]interface{} {
	for k, v := range m {
		if k == "object" {
			switch val := v.(type) {
			case map[string]interface{}:
				if r, ok := val["id"]; ok {
					if rId, ok := r.(string); ok && rId == id {
						return val
					}
				}
			case []interface{}:
				for _, elem := range val {
					if elemVal, ok := elem.(map[string]interface{}); ok {
						if r, ok := elemVal["id"]; ok {
							if rId, ok := r.(string); ok && rId == id {
								return elemVal
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// recursivelyApplyDeletes takes input map 'm' and deletes entries in its maps
// where the 'hasNils' map has nils. If the interface{} of the maps are
// themselves map[string]interface{} types, then this function recurs.
func recursivelyApplyDeletes(m, hasNils map[string]interface{}) {
	for k, v := range hasNils {
		if _, ok := m[k]; v == nil && ok {
			delete(m, k)
		} else if nilsSubMap, ok := v.(map[string]interface{}); ok {
			if mSub, ok := m[k]; ok {
				if mSubMap, ok := mSub.(map[string]interface{}); ok {
					recursivelyApplyDeletes(mSubMap, nilsSubMap)
				}
			}
		}
	}
}

func getIRISetFromItems(c vocab.CollectionType) (map[string]bool, error) {
	r := make(map[string]bool, c.ItemsLen())
	for i := 0; i < c.ItemsLen(); i++ {
		if c.IsItemsObject(i) {
			obj := c.GetItemsObject(i)
			if !obj.HasId() {
				return r, fmt.Errorf("items object at index %d has no id", i)
			}
			r[obj.GetId().String()] = true
		} else if c.IsItemsLink(i) {
			l := c.GetItemsLink(i)
			if !l.HasHref() {
				return r, fmt.Errorf("items link at index %d has no href", i)
			}
			r[l.GetHref().String()] = true
		} else if c.IsItemsIRI(i) {
			r[c.GetItemsIRI(i).String()] = true
		}
	}
	return r, nil
}

func getIRISetFromOrderedItems(c vocab.OrderedCollectionType) (map[string]bool, error) {
	r := make(map[string]bool, c.OrderedItemsLen())
	for i := 0; i < c.OrderedItemsLen(); i++ {
		if c.IsOrderedItemsObject(i) {
			obj := c.GetOrderedItemsObject(i)
			if !obj.HasId() {
				return r, fmt.Errorf("items object at index %d has no id", i)
			}
			r[obj.GetId().String()] = true
		} else if c.IsOrderedItemsLink(i) {
			l := c.GetOrderedItemsLink(i)
			if !l.HasHref() {
				return r, fmt.Errorf("items link at index %d has no href", i)
			}
			r[l.GetHref().String()] = true
		} else if c.IsOrderedItemsIRI(i) {
			r[c.GetOrderedItemsIRI(i).String()] = true
		}
	}
	return r, nil
}

func clearSensitiveFields(obj vocab.ObjectType) {
	for i := 0; i < obj.BtoLen(); i++ {
		if obj.IsBtoObject(0) {
			obj.RemoveBtoObject(0)
		} else if obj.IsBtoLink(0) {
			obj.RemoveBtoLink(0)
		} else if obj.IsBtoIRI(0) {
			obj.RemoveBtoIRI(0)
		}
	}
	for i := 0; i < obj.BccLen(); i++ {
		if obj.IsBccObject(0) {
			obj.RemoveBccObject(0)
		} else if obj.IsBccLink(0) {
			obj.RemoveBccLink(0)
		} else if obj.IsBccIRI(0) {
			obj.RemoveBccIRI(0)
		}
	}
}
