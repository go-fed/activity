// Package vocab provides an implementation of serializing and deserializing activity streams into native golang structs without relying on reflection. This package is code-generated from the vocabulary specification available at https://www.w3.org/TR/activitystreams-vocabulary and by design forgoes full resolution of raw JSON-LD data. However, custom extensions of the vocabulary are supported by modifying the data definitions in the generation tool and rerunning it. Do not modify this package directly.
package vocab

import (
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// Serializer implementations can serialize themselves to a generic map form.
type Serializer interface {
	Serialize() (m map[string]interface{}, e error)
}

// Deserializer implementations can deserialize themselves from a generic map form.
type Deserializer interface {
	Deserialize(m map[string]interface{}) (e error)
}

// Typer supports common functions for determining an ActivityStream type
type Typer interface {
	TypeLen() (l int)
	GetType(index int) (v interface{})
}

// Unknown is an entry whose root type is unknown.
type Unknown struct {
	// Raw unknown, untyped values
	u map[string]interface{}
}

// Serialize turns this object into a map[string]interface{}. Note that for the Unknown type, the "type" property is NOT populated with anything special during this process.
func (t *Unknown) Serialize() (m map[string]interface{}, err error) {
	m = t.u
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *Unknown) Deserialize(m map[string]interface{}) (err error) {
	t.u = m
	return

}

// HasField determines whether the call to GetField is safe with the specified field
func (t *Unknown) HasField(f string) (ok bool) {
	return t.u != nil && t.u[f] != nil

}

// GetField returns the unknown field value
func (t *Unknown) GetField(f string) (v interface{}) {
	return t.u[f]
}

// SetField sets the unknown field value
func (t *Unknown) SetField(f string, i interface{}) (this *Unknown) {
	if t.u == nil {
		t.u = make(map[string]interface{})
	}
	t.u[f] = i
	return t

}

// dateTimeDeserialize turns a string into a time.
func dateTimeDeserialize(v interface{}) (t *time.Time, err error) {
	var tmp time.Time
	if s, ok := v.(string); ok {
		tmp, err = time.Parse(time.RFC3339, s)
		if err != nil {
			tmp, err = time.Parse("2006-01-02T15:04Z07:00", s)
			if err != nil {
				err = fmt.Errorf("%s cannot be interpreted as xsd:dateTime", s)
			} else {
				t = &tmp
			}
		} else {
			t = &tmp
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for xsd:dateTime", v)
	}
	return

}

// dateTimeSerialize turns a time into a string
func dateTimeSerialize(t time.Time) (s string) {
	s = t.Format(time.RFC3339)
	return

}

// booleanDeserialize turns a interface{} into a bool.
func booleanDeserialize(v interface{}) (b *bool, err error) {
	if bv, ok := v.(bool); ok {
		b = &bv
	} else if bv, ok := v.(float64); ok {
		if bv == 0 {
			bvb := false
			b = &bvb
		} else if bv == 1 {
			bvb := true
			b = &bvb
		} else {
			err = fmt.Errorf("%d cannot be interpreted as a bool float64 for xsd:boolean", v)
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a bool for xsd:boolean", v)
	}
	return

}

// booleanSerialize simply returns the bool value
func booleanSerialize(v bool) (r bool) {
	r = v
	return

}

// anyURIDeserialize turns a string into a URI.
func anyURIDeserialize(v interface{}) (u *url.URL, err error) {
	if s, ok := v.(string); ok {
		u, err = url.Parse(s)
		if err != nil {
			err = fmt.Errorf("%s cannot be interpreted as xsd:anyURI", s)
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for xsd:anyURI", v)
	}
	return

}

// anyURISerialize turns a URI into a string
func anyURISerialize(u *url.URL) (s string) {
	s = u.String()
	return

}

// floatDeserialize turns a interface{} into a float64.
func floatDeserialize(v interface{}) (f *float64, err error) {
	if fv, ok := v.(float64); ok {
		f = &fv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a float for xsd:float", v)
	}
	return

}

// floatSerialize simply returns the float value
func floatSerialize(f float64) (r float64) {
	r = f
	return

}

// stringDeserialize turns a interface{} into a string.
func stringDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for xsd:string", v)
	}
	return

}

// stringSerialize simply returns the string value
func stringSerialize(s string) (r string) {
	r = s
	return

}

// langStringDeserialize turns a RDF interface{} into a string.
func langStringDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for rdf:langString", v)
	}
	return

}

// langStringSerialize returns a formatted RDF value.
func langStringSerialize(s string) (r string) {
	r = s
	return

}

// durationDeserialize turns a interface{} into a time.Duration.
func durationDeserialize(v interface{}) (d *time.Duration, err error) {
	if sv, ok := v.(string); ok {
		isNeg := false
		if sv[0] == '-' {
			isNeg = true
			sv = sv[1:]
		}
		if sv[0] != 'P' {
			err = fmt.Errorf("'%s' malformed: missing 'P' for xsd:duration", sv)
			return
		}
		re := regexp.MustCompile("P(\\d*Y)?(\\d*M)?(\\d*D)?(T(\\d*H)?(\\d*M)?(\\d*S)?)?")
		res := re.FindStringSubmatch(sv)
		var dur time.Duration
		nYear := res[1]
		if len(nYear) > 0 {
			nYear = nYear[:len(nYear)-1]
			vYear, err := strconv.ParseInt(nYear, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vYear) * time.Hour * 8760 // Hours per 365 days -- no way to count leap years here.
		}
		nMonth := res[2]
		if len(nMonth) > 0 {
			nMonth = nMonth[:len(nMonth)-1]
			vMonth, err := strconv.ParseInt(nMonth, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vMonth) * time.Hour * 720 // Hours per 30 days -- no way to tell if these months span 31 days, 28, or 29 each.
		}
		nDay := res[3]
		if len(nDay) > 0 {
			nDay = nDay[:len(nDay)-1]
			vDay, err := strconv.ParseInt(nDay, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vDay) * time.Hour * 24
		}
		nHour := res[5]
		if len(nHour) > 0 {
			nHour = nHour[:len(nHour)-1]
			vHour, err := strconv.ParseInt(nHour, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vHour) * time.Hour
		}
		nMinute := res[6]
		if len(nMinute) > 0 {
			nMinute = nMinute[:len(nMinute)-1]
			vMinute, err := strconv.ParseInt(nMinute, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vMinute) * time.Minute
		}
		nSecond := res[7]
		if len(nSecond) > 0 {
			nSecond = nSecond[:len(nSecond)-1]
			vSecond, err := strconv.ParseInt(nSecond, 10, 64)
			if err != nil {
				return nil, err
			}
			dur += time.Duration(vSecond) * time.Second
		}
		if isNeg {
			dur *= -1
		}
		d = &dur
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for xsd:duration", v)
	}
	return

}

// durationSerialize returns the duration as a string.
func durationSerialize(d time.Duration) (s string) {
	s = "P"
	if d < 0 {
		s = "-P"
		d = -1 * d
	}
	var tally time.Duration
	if years := d.Hours() / 8760.0; years >= 1 {
		nYears := int64(math.Floor(years))
		tally += time.Duration(nYears) * 8760 * time.Hour
		s = fmt.Sprintf("%s%dY", s, nYears)
	}
	if months := (d.Hours() - tally.Hours()) / 720.0; months >= 1 {
		nMonths := int64(math.Floor(months))
		tally += time.Duration(nMonths) * 720 * time.Hour
		s = fmt.Sprintf("%s%dM", s, nMonths)
	}
	if days := (d.Hours() - tally.Hours()) / 24.0; days >= 1 {
		nDays := int64(math.Floor(days))
		tally += time.Duration(nDays) * 24 * time.Hour
		s = fmt.Sprintf("%s%dD", s, nDays)
	}
	if tally < d {
		s = fmt.Sprintf("%sT", s)
		if hours := d.Hours() - tally.Hours(); hours >= 1 {
			nHours := int64(math.Floor(hours))
			tally += time.Duration(nHours) * time.Hour
			s = fmt.Sprintf("%s%dH", s, nHours)
		}
		if minutes := d.Minutes() - tally.Minutes(); minutes >= 1 {
			nMinutes := int64(math.Floor(minutes))
			tally += time.Duration(nMinutes) * time.Minute
			s = fmt.Sprintf("%s%dM", s, nMinutes)
		}
		if seconds := d.Seconds() - tally.Seconds(); seconds >= 1 {
			nSeconds := int64(math.Floor(seconds))
			tally += time.Duration(nSeconds) * time.Second
			s = fmt.Sprintf("%s%dS", s, nSeconds)
		}
	}
	return

}

// nonNegativeIntegerDeserialize turns a interface{} into a positive int64 value.
func nonNegativeIntegerDeserialize(v interface{}) (i *int64, err error) {
	if fv, ok := v.(float64); ok {
		iv := int64(fv)
		if iv >= 0 {
			i = &iv
		} else {
			err = fmt.Errorf("%d is a negative integer for xsd:nonNegativeInteger", iv)
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a float for xsd:nonNegativeInteger", v)
	}
	return

}

// nonNegativeIntegerSerialize simply returns the int64 value.
func nonNegativeIntegerSerialize(i int64) (r int64) {
	r = i
	return

}

// bcp47LanguageTagDeserialize turns a interface{} into a string.
func bcp47LanguageTagDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for BCP 47 Language Tag", v)
	}
	return

}

// bcp47LanguageTagSerialize simply returns the string value
func bcp47LanguageTagSerialize(s string) (r string) {
	r = s
	return

}

// mimeMediaTypeValueDeserialize turns a interface{} into a string.
func mimeMediaTypeValueDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for MIME media type value", v)
	}
	return

}

// mimeMediaTypeValueSerialize simply returns the string value
func mimeMediaTypeValueSerialize(s string) (r string) {
	r = s
	return

}

// linkRelationDeserialize turns a interface{} into a string.
func linkRelationDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for link relation", v)
	}
	return

}

// linkRelationSerialize simply returns the string value
func linkRelationSerialize(s string) (r string) {
	r = s
	return

}

// unitsValueDeserialize turns a interface{} into a string.
func unitsValueDeserialize(v interface{}) (s *string, err error) {
	if sv, ok := v.(string); ok {
		s = &sv
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for units value", v)
	}
	return

}

// unitsValueSerialize simply returns the string value
func unitsValueSerialize(s string) (r string) {
	r = s
	return

}

// IRIDeserialize turns a string into a URI.
func IRIDeserialize(v interface{}) (u *url.URL, err error) {
	if s, ok := v.(string); ok {
		u, err = url.Parse(s)
		if err != nil {
			err = fmt.Errorf("%s cannot be interpreted as IRI", s)
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for IRI", v)
	}
	return

}

// IRISerialize turns an IRI into a string
func IRISerialize(u *url.URL) (s string) {
	s = u.String()
	return

}

// HasTypeObject returns true if the Typer has a type of Object.
func HasTypeObject(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Object" {
				return true
			}
		}
	}
	return false

}

// HasTypeLink returns true if the Typer has a type of Link.
func HasTypeLink(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Link" {
				return true
			}
		}
	}
	return false

}

// HasTypeActivity returns true if the Typer has a type of Activity.
func HasTypeActivity(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Activity" {
				return true
			}
		}
	}
	return false

}

// HasTypeIntransitiveActivity returns true if the Typer has a type of IntransitiveActivity.
func HasTypeIntransitiveActivity(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "IntransitiveActivity" {
				return true
			}
		}
	}
	return false

}

// HasTypeCollection returns true if the Typer has a type of Collection.
func HasTypeCollection(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Collection" {
				return true
			}
		}
	}
	return false

}

// HasTypeOrderedCollection returns true if the Typer has a type of OrderedCollection.
func HasTypeOrderedCollection(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "OrderedCollection" {
				return true
			}
		}
	}
	return false

}

// HasTypeCollectionPage returns true if the Typer has a type of CollectionPage.
func HasTypeCollectionPage(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "CollectionPage" {
				return true
			}
		}
	}
	return false

}

// HasTypeOrderedCollectionPage returns true if the Typer has a type of OrderedCollectionPage.
func HasTypeOrderedCollectionPage(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "OrderedCollectionPage" {
				return true
			}
		}
	}
	return false

}

// HasTypeAccept returns true if the Typer has a type of Accept.
func HasTypeAccept(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Accept" {
				return true
			}
		}
	}
	return false

}

// HasTypeTentativeAccept returns true if the Typer has a type of TentativeAccept.
func HasTypeTentativeAccept(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "TentativeAccept" {
				return true
			}
		}
	}
	return false

}

// HasTypeAdd returns true if the Typer has a type of Add.
func HasTypeAdd(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Add" {
				return true
			}
		}
	}
	return false

}

// HasTypeArrive returns true if the Typer has a type of Arrive.
func HasTypeArrive(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Arrive" {
				return true
			}
		}
	}
	return false

}

// HasTypeCreate returns true if the Typer has a type of Create.
func HasTypeCreate(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Create" {
				return true
			}
		}
	}
	return false

}

// HasTypeDelete returns true if the Typer has a type of Delete.
func HasTypeDelete(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Delete" {
				return true
			}
		}
	}
	return false

}

// HasTypeFollow returns true if the Typer has a type of Follow.
func HasTypeFollow(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Follow" {
				return true
			}
		}
	}
	return false

}

// HasTypeIgnore returns true if the Typer has a type of Ignore.
func HasTypeIgnore(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Ignore" {
				return true
			}
		}
	}
	return false

}

// HasTypeJoin returns true if the Typer has a type of Join.
func HasTypeJoin(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Join" {
				return true
			}
		}
	}
	return false

}

// HasTypeLeave returns true if the Typer has a type of Leave.
func HasTypeLeave(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Leave" {
				return true
			}
		}
	}
	return false

}

// HasTypeLike returns true if the Typer has a type of Like.
func HasTypeLike(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Like" {
				return true
			}
		}
	}
	return false

}

// HasTypeOffer returns true if the Typer has a type of Offer.
func HasTypeOffer(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Offer" {
				return true
			}
		}
	}
	return false

}

// HasTypeInvite returns true if the Typer has a type of Invite.
func HasTypeInvite(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Invite" {
				return true
			}
		}
	}
	return false

}

// HasTypeReject returns true if the Typer has a type of Reject.
func HasTypeReject(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Reject" {
				return true
			}
		}
	}
	return false

}

// HasTypeTentativeReject returns true if the Typer has a type of TentativeReject.
func HasTypeTentativeReject(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "TentativeReject" {
				return true
			}
		}
	}
	return false

}

// HasTypeRemove returns true if the Typer has a type of Remove.
func HasTypeRemove(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Remove" {
				return true
			}
		}
	}
	return false

}

// HasTypeUndo returns true if the Typer has a type of Undo.
func HasTypeUndo(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Undo" {
				return true
			}
		}
	}
	return false

}

// HasTypeUpdate returns true if the Typer has a type of Update.
func HasTypeUpdate(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Update" {
				return true
			}
		}
	}
	return false

}

// HasTypeView returns true if the Typer has a type of View.
func HasTypeView(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "View" {
				return true
			}
		}
	}
	return false

}

// HasTypeListen returns true if the Typer has a type of Listen.
func HasTypeListen(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Listen" {
				return true
			}
		}
	}
	return false

}

// HasTypeRead returns true if the Typer has a type of Read.
func HasTypeRead(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Read" {
				return true
			}
		}
	}
	return false

}

// HasTypeMove returns true if the Typer has a type of Move.
func HasTypeMove(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Move" {
				return true
			}
		}
	}
	return false

}

// HasTypeTravel returns true if the Typer has a type of Travel.
func HasTypeTravel(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Travel" {
				return true
			}
		}
	}
	return false

}

// HasTypeAnnounce returns true if the Typer has a type of Announce.
func HasTypeAnnounce(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Announce" {
				return true
			}
		}
	}
	return false

}

// HasTypeBlock returns true if the Typer has a type of Block.
func HasTypeBlock(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Block" {
				return true
			}
		}
	}
	return false

}

// HasTypeFlag returns true if the Typer has a type of Flag.
func HasTypeFlag(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Flag" {
				return true
			}
		}
	}
	return false

}

// HasTypeDislike returns true if the Typer has a type of Dislike.
func HasTypeDislike(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Dislike" {
				return true
			}
		}
	}
	return false

}

// HasTypeQuestion returns true if the Typer has a type of Question.
func HasTypeQuestion(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Question" {
				return true
			}
		}
	}
	return false

}

// HasTypeApplication returns true if the Typer has a type of Application.
func HasTypeApplication(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Application" {
				return true
			}
		}
	}
	return false

}

// HasTypeGroup returns true if the Typer has a type of Group.
func HasTypeGroup(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Group" {
				return true
			}
		}
	}
	return false

}

// HasTypeOrganization returns true if the Typer has a type of Organization.
func HasTypeOrganization(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Organization" {
				return true
			}
		}
	}
	return false

}

// HasTypePerson returns true if the Typer has a type of Person.
func HasTypePerson(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Person" {
				return true
			}
		}
	}
	return false

}

// HasTypeService returns true if the Typer has a type of Service.
func HasTypeService(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Service" {
				return true
			}
		}
	}
	return false

}

// HasTypeRelationship returns true if the Typer has a type of Relationship.
func HasTypeRelationship(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Relationship" {
				return true
			}
		}
	}
	return false

}

// HasTypeArticle returns true if the Typer has a type of Article.
func HasTypeArticle(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Article" {
				return true
			}
		}
	}
	return false

}

// HasTypeDocument returns true if the Typer has a type of Document.
func HasTypeDocument(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Document" {
				return true
			}
		}
	}
	return false

}

// HasTypeAudio returns true if the Typer has a type of Audio.
func HasTypeAudio(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Audio" {
				return true
			}
		}
	}
	return false

}

// HasTypeImage returns true if the Typer has a type of Image.
func HasTypeImage(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Image" {
				return true
			}
		}
	}
	return false

}

// HasTypeVideo returns true if the Typer has a type of Video.
func HasTypeVideo(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Video" {
				return true
			}
		}
	}
	return false

}

// HasTypeNote returns true if the Typer has a type of Note.
func HasTypeNote(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Note" {
				return true
			}
		}
	}
	return false

}

// HasTypePage returns true if the Typer has a type of Page.
func HasTypePage(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Page" {
				return true
			}
		}
	}
	return false

}

// HasTypeEvent returns true if the Typer has a type of Event.
func HasTypeEvent(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Event" {
				return true
			}
		}
	}
	return false

}

// HasTypePlace returns true if the Typer has a type of Place.
func HasTypePlace(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Place" {
				return true
			}
		}
	}
	return false

}

// HasTypeProfile returns true if the Typer has a type of Profile.
func HasTypeProfile(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Profile" {
				return true
			}
		}
	}
	return false

}

// HasTypeTombstone returns true if the Typer has a type of Tombstone.
func HasTypeTombstone(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Tombstone" {
				return true
			}
		}
	}
	return false

}

// HasTypeMention returns true if the Typer has a type of Mention.
func HasTypeMention(t Typer) (b bool) {
	for i := 0; i < t.TypeLen(); i++ {
		v := t.GetType(i)
		if s, ok := v.(string); ok {
			if s == "Mention" {
				return true
			}
		}
	}
	return false

}

// Returns true if the provided Typer is an Activity.
func IsActivityType(t Typer) (b bool) {
	var activityTypes = []string{"IntransitiveActivity", "Accept", "TentativeAccept", "Add", "Arrive", "Create", "Delete", "Follow", "Ignore", "Join", "Leave", "Like", "Offer", "Invite", "Reject", "TentativeReject", "Remove", "Undo", "Update", "View", "Listen", "Read", "Move", "Travel", "Announce", "Block", "Flag", "Dislike", "Question"}
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

// resolveObject turns a string type that extends Object into a concrete type.
func resolveObject(s string) (i interface{}) {
	if s == "Object" {
		return &Object{}
	}
	if s == "Activity" {
		return &Activity{}
	}
	if s == "IntransitiveActivity" {
		return &IntransitiveActivity{}
	}
	if s == "Collection" {
		return &Collection{}
	}
	if s == "OrderedCollection" {
		return &OrderedCollection{}
	}
	if s == "CollectionPage" {
		return &CollectionPage{}
	}
	if s == "OrderedCollectionPage" {
		return &OrderedCollectionPage{}
	}
	if s == "Accept" {
		return &Accept{}
	}
	if s == "TentativeAccept" {
		return &TentativeAccept{}
	}
	if s == "Add" {
		return &Add{}
	}
	if s == "Arrive" {
		return &Arrive{}
	}
	if s == "Create" {
		return &Create{}
	}
	if s == "Delete" {
		return &Delete{}
	}
	if s == "Follow" {
		return &Follow{}
	}
	if s == "Ignore" {
		return &Ignore{}
	}
	if s == "Join" {
		return &Join{}
	}
	if s == "Leave" {
		return &Leave{}
	}
	if s == "Like" {
		return &Like{}
	}
	if s == "Offer" {
		return &Offer{}
	}
	if s == "Invite" {
		return &Invite{}
	}
	if s == "Reject" {
		return &Reject{}
	}
	if s == "TentativeReject" {
		return &TentativeReject{}
	}
	if s == "Remove" {
		return &Remove{}
	}
	if s == "Undo" {
		return &Undo{}
	}
	if s == "Update" {
		return &Update{}
	}
	if s == "View" {
		return &View{}
	}
	if s == "Listen" {
		return &Listen{}
	}
	if s == "Read" {
		return &Read{}
	}
	if s == "Move" {
		return &Move{}
	}
	if s == "Travel" {
		return &Travel{}
	}
	if s == "Announce" {
		return &Announce{}
	}
	if s == "Block" {
		return &Block{}
	}
	if s == "Flag" {
		return &Flag{}
	}
	if s == "Dislike" {
		return &Dislike{}
	}
	if s == "Question" {
		return &Question{}
	}
	if s == "Application" {
		return &Application{}
	}
	if s == "Group" {
		return &Group{}
	}
	if s == "Organization" {
		return &Organization{}
	}
	if s == "Person" {
		return &Person{}
	}
	if s == "Service" {
		return &Service{}
	}
	if s == "Relationship" {
		return &Relationship{}
	}
	if s == "Article" {
		return &Article{}
	}
	if s == "Document" {
		return &Document{}
	}
	if s == "Audio" {
		return &Audio{}
	}
	if s == "Image" {
		return &Image{}
	}
	if s == "Video" {
		return &Video{}
	}
	if s == "Note" {
		return &Note{}
	}
	if s == "Page" {
		return &Page{}
	}
	if s == "Event" {
		return &Event{}
	}
	if s == "Place" {
		return &Place{}
	}
	if s == "Profile" {
		return &Profile{}
	}
	if s == "Tombstone" {
		return &Tombstone{}
	}
	return nil

}

// resolveLink turns a string type that extends Link into a concrete type.
func resolveLink(s string) (i interface{}) {
	if s == "Link" {
		return &Link{}
	}
	if s == "Mention" {
		return &Mention{}
	}
	return nil

}

// unknownValueDeserialize transparently stores the object.
func unknownValueDeserialize(v interface{}) (o interface{}) {
	o = v
	return

}

// unknownValueSerialize transparently returns the object.
func unknownValueSerialize(v interface{}) (o interface{}) {
	o = v
	return

}
