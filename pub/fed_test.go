package pub

import (
	"context"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	testAgent = "test agent string"
)

func Must(l *time.Location, e error) *time.Location {
	if e != nil {
		panic(e)
	}
	return l
}

var (
	now = time.Date(2000, 2, 3, 4, 5, 6, 7, Must(time.LoadLocation("America/New_York")))
)

var _ Clock = &MockClock{}

type MockClock struct {
	now time.Time
}

func (m *MockClock) Now() time.Time {
	return m.now
}

var _ Application = &MockApplication{}

type MockApplication struct {
	t                    *testing.T
	get                  func(c context.Context, id url.URL) (PubObject, error)
	set                  func(c context.Context, o PubObject) error
	getInbox             func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	getOutbox            func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	postOutboxAuthorized func(c context.Context, r *http.Request) (bool, error)
	newId                func(c context.Context, t Typer) url.URL
	actorIRI             func(c context.Context, r *http.Request) (url.URL, error)
}

func (m *MockApplication) Get(c context.Context, id url.URL) (PubObject, error) {
	if m.get == nil {
		m.t.Fatal("unexpected call to MockApplication Get")
	}
	return m.get(c, id)
}

func (m *MockApplication) Set(c context.Context, o PubObject) error {
	if m.set == nil {
		m.t.Fatal("unexpected call to MockApplication Set")
	}
	return m.set(c, o)
}

func (m *MockApplication) GetInbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
	if m.getInbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetInbox")
	}
	return m.getInbox(c, r)
}

func (m *MockApplication) GetOutbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
	if m.getOutbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetOutbox")
	}
	return m.getOutbox(c, r)
}

func (m *MockApplication) PostOutboxAuthorized(c context.Context, r *http.Request) (bool, error) {
	if m.postOutboxAuthorized == nil {
		m.t.Fatal("unexpected call to MockApplication PostOutboxAuthorized")
	}
	return m.postOutboxAuthorized(c, r)
}

func (m *MockApplication) NewId(c context.Context, t Typer) url.URL {
	if m.newId == nil {
		m.t.Fatal("unexpected call to MockApplication NewId")
	}
	return m.newId(c, t)
}

func (m *MockApplication) ActorIRI(c context.Context, r *http.Request) (url.URL, error) {
	if m.actorIRI == nil {
		m.t.Fatal("unexpected call to MockApplication ActorIRI")
	}
	return m.actorIRI(c, r)
}

var _ SocialApp = &MockSocialApp{}

type MockSocialApp struct {
	t         *testing.T
	owns      func(c context.Context, id url.URL) bool
	canAdd    func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
	canRemove func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
}

func (m *MockSocialApp) Owns(c context.Context, id url.URL) bool {
	if m.owns == nil {
		m.t.Fatal("unexpected call to MockSocialApp Owns")
	}
	return m.owns(c, id)
}

func (m *MockSocialApp) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canAdd == nil {
		m.t.Fatal("unexpected call to MockSocialApp CanAdd")
	}
	return m.canAdd(c, o, t)
}

func (m *MockSocialApp) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canRemove == nil {
		m.t.Fatal("unexpected call to MockSocialApp CanRemove")
	}
	return m.canRemove(c, o, t)
}

var _ Callbacker = &MockCallbacker{}

type MockCallbacker struct {
	t      *testing.T
	create func(c context.Context, s *streams.Create) error
	update func(c context.Context, s *streams.Update) error
	delete func(c context.Context, s *streams.Delete) error
	add    func(c context.Context, s *streams.Add) error
	remove func(c context.Context, s *streams.Remove) error
	like   func(c context.Context, s *streams.Like) error
	block  func(c context.Context, s *streams.Block) error
	follow func(c context.Context, s *streams.Follow) error
	undo   func(c context.Context, s *streams.Undo) error
	accept func(c context.Context, s *streams.Accept) error
	reject func(c context.Context, s *streams.Reject) error
}

func (m *MockCallbacker) Create(c context.Context, s *streams.Create) error {
	if m.create == nil {
		m.t.Logf("unimplemented MockCallbacker Create called: %v %v", c, s)
		return nil
	} else {
		return m.create(c, s)
	}
}

func (m *MockCallbacker) Update(c context.Context, s *streams.Update) error {
	if m.update == nil {
		m.t.Logf("unimplemented MockCallbacker Update called: %v %v", c, s)
		return nil
	} else {
		return m.update(c, s)
	}
}

func (m *MockCallbacker) Delete(c context.Context, s *streams.Delete) error {
	if m.delete == nil {
		m.t.Logf("unimplemented MockCallbacker Delete called: %v %v", c, s)
		return nil
	} else {
		return m.delete(c, s)
	}
}

func (m *MockCallbacker) Add(c context.Context, s *streams.Add) error {
	if m.add == nil {
		m.t.Logf("unimplemented MockCallbacker Add called: %v %v", c, s)
		return nil
	} else {
		return m.add(c, s)
	}
}

func (m *MockCallbacker) Remove(c context.Context, s *streams.Remove) error {
	if m.Remove == nil {
		m.t.Logf("unimplemented MockCallbacker Remove called: %v %v", c, s)
		return nil
	} else {
		return m.remove(c, s)
	}
}

func (m *MockCallbacker) Like(c context.Context, s *streams.Like) error {
	if m.like == nil {
		m.t.Logf("unimplemented MockCallbacker Like called: %v %v", c, s)
		return nil
	} else {
		return m.like(c, s)
	}
}

func (m *MockCallbacker) Block(c context.Context, s *streams.Block) error {
	if m.block == nil {
		m.t.Logf("unimplemented MockCallbacker Block called: %v %v", c, s)
		return nil
	} else {
		return m.block(c, s)
	}
}

func (m *MockCallbacker) Follow(c context.Context, s *streams.Follow) error {
	if m.follow == nil {
		m.t.Logf("unimplemented MockCallbacker Follow called: %v %v", c, s)
		return nil
	} else {
		return m.follow(c, s)
	}
}

func (m *MockCallbacker) Undo(c context.Context, s *streams.Undo) error {
	if m.undo == nil {
		m.t.Logf("unimplemented MockCallbacker Undo called: %v %v", c, s)
		return nil
	} else {
		return m.undo(c, s)
	}
}

func (m *MockCallbacker) Accept(c context.Context, s *streams.Accept) error {
	if m.accept == nil {
		m.t.Logf("unimplemented MockCallbacker Accept called: %v %v", c, s)
		return nil
	} else {
		return m.accept(c, s)
	}
}

func (m *MockCallbacker) Reject(c context.Context, s *streams.Reject) error {
	if m.reject == nil {
		m.t.Logf("unimplemented MockCallbacker Reject called: %v %v", c, s)
		return nil
	} else {
		return m.reject(c, s)
	}
}

var _ FederateApp = &MockFederateApp{}

type MockFederateApp struct {
	t            *testing.T
	owns         func(c context.Context, id url.URL) bool
	canAdd       func(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool
	canRemove    func(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool
	onFollow     func(c context.Context, s *streams.Follow) FollowResponse
	unblocked    func(c context.Context, actorIRIs []url.URL) error
	getFollowing func(c context.Context, actor url.URL) (vocab.CollectionType, error)
}

func (m *MockFederateApp) Owns(c context.Context, id url.URL) bool {
	if m.owns == nil {
		m.t.Fatal("unexpected call to MockFederateApp Owns")
	}
	return m.owns(c, id)
}

func (m *MockFederateApp) CanAdd(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool {
	if m.canAdd == nil {
		m.t.Fatal("unexpected call to MockFederateApp CanAdd")
	}
	return m.canAdd(c, obj, target)
}

func (m *MockFederateApp) CanRemove(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool {
	if m.canRemove == nil {
		m.t.Fatal("unexpected call to MockFederateApp CanRemove")
	}
	return m.canRemove(c, obj, target)
}

func (m *MockFederateApp) OnFollow(c context.Context, s *streams.Follow) FollowResponse {
	if m.onFollow == nil {
		m.t.Fatal("unexpected call to MockFederateApp OnFollow")
	}
	return m.onFollow(c, s)
}

func (m *MockFederateApp) Unblocked(c context.Context, actorIRIs []url.URL) error {
	if m.unblocked == nil {
		m.t.Fatal("unexpected call to MockFederateApp Unblocked")
	}
	return m.unblocked(c, actorIRIs)
}

func (m *MockFederateApp) GetFollowing(c context.Context, actor url.URL) (vocab.CollectionType, error) {
	if m.getFollowing == nil {
		m.t.Fatal("unexpected call to MockFederateApp GetFollowing")
	}
	return m.getFollowing(c, actor)
}

var _ DeliveryPersister = &MockPersister{}

type MockPersister struct {
	t             *testing.T
	sending       func(b []byte, to url.URL) string
	cancel        func(id string)
	successful    func(id string)
	retrying      func(id string)
	undeliverable func(id string)
}

func (m *MockPersister) Sending(b []byte, to url.URL) string {
	if m.sending == nil {
		m.t.Fatal("unexpected call to MockPersister Sending")
	}
	return m.sending(b, to)
}

func (m *MockPersister) Cancel(id string) {
	if m.cancel == nil {
		m.t.Fatal("unexpected call to MockPersister Cancel")
	}
	m.cancel(id)
}

func (m *MockPersister) Successful(id string) {
	if m.successful == nil {
		m.t.Fatal("unexpected call to MockPersister Successful")
	}
	m.successful(id)
}

func (m *MockPersister) Retrying(id string) {
	if m.retrying == nil {
		m.t.Fatal("unexpected call to MockPersister Retrying")
	}
	m.retrying(id)
}

func (m *MockPersister) Undeliverable(id string) {
	if m.undeliverable == nil {
		m.t.Fatal("unexpected call to MockPersister Undeliverable")
	}
	m.undeliverable(id)
}

func NewSocialPubberTest(t *testing.T) (app *MockApplication, socialApp *MockSocialApp, cb *MockCallbacker, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	socialApp = &MockSocialApp{t: t}
	cb = &MockCallbacker{t: t}
	p = NewSocialPubber(clock, app, socialApp, cb)
	return
}

func NewFederatingPubberTest(t *testing.T) (app *MockApplication, fedApp *MockFederateApp, cb *MockCallbacker, per *MockPersister, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	fedApp = &MockFederateApp{t: t}
	cb = &MockCallbacker{t: t}
	per = &MockPersister{t: t}
	d := DeliveryOptions{
		Client:           &http.Client{},
		Agent:            testAgent,
		MaxDeliveryDepth: 1,
		InitialRetryTime: time.Second,
		MaximumRetryTime: time.Second,
		BackoffFactor:    1,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1, 1),
		Persister:        per,
	}
	p = NewFederatingPubber(clock, app, fedApp, cb, d)
	return
}

func NewPubberTest(t *testing.T) (app *MockApplication, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, per *MockPersister, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	socialApp = &MockSocialApp{t: t}
	fedApp = &MockFederateApp{t: t}
	socialCb = &MockCallbacker{t: t}
	fedCb = &MockCallbacker{t: t}
	per = &MockPersister{t: t}
	d := DeliveryOptions{
		Client:           &http.Client{},
		Agent:            testAgent,
		MaxDeliveryDepth: 1,
		InitialRetryTime: time.Second,
		MaximumRetryTime: time.Second,
		BackoffFactor:    1,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1, 1),
		Persister:        per,
	}
	p = NewPubber(clock, app, socialApp, fedApp, socialCb, fedCb, d)
	return
}

func TestSocialPubber_Stop(t *testing.T) {
	_, _, _, p := NewSocialPubberTest(t)
	p.Stop()
}

func TestSocialPubber_RejectPostInbox(t *testing.T) {
	// TODO: Implement
}

func TestSocialPubber_GetInbox(t *testing.T) {
	// TODO: Implement
}

func TestSocialPubber_PostOutbox(t *testing.T) {
	// TODO: Implement
}

func TestSocialPubber_GetOutbox(t *testing.T) {
	// TODO: Implement
}

func TestFederatingPubber_Stop(t *testing.T) {
	_, _, _, _, p := NewFederatingPubberTest(t)
	p.Stop()
}

func TestFederatingPubber_PostInbox(t *testing.T) {
	// TODO: Implement
}

func TestFederatingPubber_GetInbox(t *testing.T) {
	// TODO: Implement
}

func TestFederatingPubber_RejectPostOutbox(t *testing.T) {
	// TODO: Implement
}

func TestFederatingPubber_GetOutbox(t *testing.T) {
	// TODO: Implement
}

func TestPubber_Stop(t *testing.T) {
	_, _, _, _, _, _, p := NewPubberTest(t)
	p.Stop()
}

func TestPubber_PostInbox(t *testing.T) {
	// TODO: Implement
}

func TestPubber_GetInbox(t *testing.T) {
	// TODO: Implement
}

func TestPubber_PostOutbox(t *testing.T) {
	// TODO: Implement
}

func TestPubber_GetOutbox(t *testing.T) {
	// TODO: Implement
}

func PostInbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func PostInbox_HandlesBlocked(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Create_SetsObject(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Create_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Update_SetsObject(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Update_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Delete_SetsTombstone(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Delete_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Follow_DoNothing(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Follow_AutoReject(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Follow_AutoAccept(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Follow_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Accept_DoesNothingIfNotAcceptingFollow(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Accept_AcceptFollowAddsToFollowers(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Accept_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Reject_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_RequireType(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_DoesNotAddIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_AddIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Add_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_RequireType(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_DoesNotAddIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_AddIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Remove_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Like_AddsToLikeCollection(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Like_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostInbox_Undo_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func GetInbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func GetInbox_SetsContentTypeHeader(t *testing.T) {
	// TODO: Implement
}

func GetInbox_DeduplicateInboxItems(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_RejectUnauthorized(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_WrapInCreateActivity(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_CopyToAttributedTo(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_CopyRecipientsToObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_SetCreatedObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Create_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Update_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Update_OverwriteUpdatedFields(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Update_SetUpdatedObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Update_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Update_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Delete_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Delete_SetsTombstone(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Delete_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Delete_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Follow_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Follow_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Follow_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Accept_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Accept_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Reject_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Reject_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_RequireType(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_DoesNotAddIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_AddsIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Add_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_RequireType(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_DoesNotRemoveIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_RemoveIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_DoesNotRemoveIfAppCannotRemove(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_RemoveIfAppCanRemove(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Remove_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_RequireType(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_AddsToLikedCollection(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_AddsToLikedOrderedCollection(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Like_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Undo_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Undo_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Undo_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Block_RequireObject(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Block_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_Block_IsNotDelivered(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_DoesNotDeliverNondeliverable(t *testing.T) {
	// TODO: Implement
}

func PostOutbox_SetsLocationHeader(t *testing.T) {
	// TODO: Implement
}

func GetOutbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func GetOutbox_SetsContentTypeHeader(t *testing.T) {
	// TODO: Implement
}
