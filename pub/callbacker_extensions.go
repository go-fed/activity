package pub

import (
	"context"
	"github.com/go-fed/activity/streams"
)

type callbackerAnnounce interface {
	Announce(c context.Context, s *streams.Announce) error
}

type callbackerArrive interface {
	Arrive(c context.Context, s *streams.Arrive) error
}

type callbackerDislike interface {
	Dislike(c context.Context, s *streams.Dislike) error
}

type callbackerFlag interface {
	Flag(c context.Context, s *streams.Flag) error
}

type callbackerIgnore interface {
	Ignore(c context.Context, s *streams.Ignore) error
}

type callbackerInvite interface {
	Invite(c context.Context, s *streams.Invite) error
}

type callbackerJoin interface {
	Join(c context.Context, s *streams.Join) error
}

type callbackerLeave interface {
	Leave(c context.Context, s *streams.Leave) error
}

type callbackerListen interface {
	Listen(c context.Context, s *streams.Listen) error
}

type callbackerMove interface {
	Move(c context.Context, s *streams.Move) error
}

type callbackerOffer interface {
	Offer(c context.Context, s *streams.Offer) error
}

type callbackerQuestion interface {
	Question(c context.Context, s *streams.Question) error
}

type callbackerRead interface {
	Read(c context.Context, s *streams.Read) error
}

type callbackerTentativeAccept interface {
	TentativeAccept(c context.Context, s *streams.TentativeAccept) error
}

type callbackerTentativeReject interface {
	TentativeReject(c context.Context, s *streams.TentativeReject) error
}

type callbackerTravel interface {
	Travel(c context.Context, s *streams.Travel) error
}

type callbackerView interface {
	View(c context.Context, s *streams.View) error
}
