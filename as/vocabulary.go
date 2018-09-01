// Package extension outlines the interfaces required for an ActivityStreams
// extension.
package extension

import ()

// TODO: Figure out the new "Resolver" and "Property-Resolver" algorithms for
// the plugin architecture. Two kinds: for deserialization, and for dispatching.

// Plugin
type Plugin interface {
	NewTypeResolver()
}

// Resolver
type Resolver interface {
	Deserialize(map[string]interface{}) error
}

type PropertyResolver interface {
}

/*
EXISTING USAGE:

vocab.IsActivityType()
vocab.HasTypeTombstone()
vocab.HasTypeCreate()
vocab.Serializer
vocab.Typer
vocab.ObjectType
vocab.LinkType
vocab.ActivityType
vocab.IntransitiveActivityType
vocab.CollectionType
vocab.CollectionPageType
vocab.OrderedCollectionType
vocab.OrderedCollectionPageType
vocab.FollowType
vocab.TombstoneType
vocab.Object
vocab.Create
vocab.Accept
vocab.Reject
vocab.OrderedCollection
vocab.OrderedCollectionPage
vocab.Collection
vocab.CollectionPage
vocab.Tombstone

streams.Resolver
streams.Create
streams.Update
streams.Delete
streams.Follow
streams.Accept
streams.Reject
streams.Add
streams.Remove
streams.Like
streams.Undo
streams.Block
streams.Announce
streams.Dislike
streams.Flag
streams.Ignore
streams.Invite
streams.Join
streams.Leave
streams.Listen
streams.Move
streams.Offer
streams.Question
streams.Read
streams.TentativeAccept
streams.TentativeReject
streams.Travel
streams.View
streams.Collection
streams.OrderedCollection
streams.CollectionPage
streams.OrderedCollectionPage
*/
