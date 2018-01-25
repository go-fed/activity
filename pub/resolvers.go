package pub

import (
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"net/url"
)

func toActorResolver(a *ActorObject) *streams.Resolver {
	return &streams.Resolver{
		AnyObjectCallback: func(i vocab.ObjectType) error {
			if o, ok := i.(ActorObject); ok {
				*a = o
			}
			return nil
		},
	}
}

func toActorCollectionResolver(a *ActorObject, c **streams.Collection, oc **streams.OrderedCollection, cp **streams.CollectionPage, ocp **streams.OrderedCollectionPage) *streams.Resolver {
	r := toActorResolver(a)
	r.CollectionCallback = func(i *streams.Collection) error {
		*c = i
		return nil
	}
	r.OrderedCollectionCallback = func(i *streams.OrderedCollection) error {
		*oc = i
		return nil
	}
	r.CollectionPageCallback = func(i *streams.CollectionPage) error {
		*cp = i
		return nil
	}
	r.OrderedCollectionPageCallback = func(i *streams.OrderedCollectionPage) error {
		*ocp = i
		return nil
	}
	return r
}

func toIdResolver(ok *bool, u *url.URL) *streams.Resolver {
	return &streams.Resolver{
		AnyObjectCallback: func(i vocab.ObjectType) error {
			*ok = i.HasId()
			if *ok {
				*u = i.GetId()
			}
			return nil
		},
	}
}

func toCollectionPage(m map[string]interface{}) (c *streams.CollectionPage, err error) {
	r := &streams.Resolver{
		CollectionPageCallback: func(i *streams.CollectionPage) error {
			c = i
			return nil
		},
	}
	err = r.Deserialize(m)
	return
}

func toOrderedCollectionPage(m map[string]interface{}) (c *streams.OrderedCollectionPage, err error) {
	r := &streams.Resolver{
		OrderedCollectionPageCallback: func(i *streams.OrderedCollectionPage) error {
			c = i
			return nil
		},
	}
	err = r.Deserialize(m)
	return
}
