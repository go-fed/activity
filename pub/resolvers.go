package pub

import (
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"net/url"
)

func toActorResolver(a *actorObject) *streams.Resolver {
	return &streams.Resolver{
		AnyObjectCallback: func(i vocab.ObjectType) error {
			if o, ok := i.(actorObject); ok {
				*a = o
			}
			return nil
		},
	}
}

func toActorCollectionResolver(a *actorObject, c **streams.Collection, oc **streams.OrderedCollection, cp **streams.CollectionPage, ocp **streams.OrderedCollectionPage) *streams.Resolver {
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

func toTypeIder(m map[string]interface{}) (tid typeIder, err error) {
	var t []typeIder
	r := &streams.Resolver{
		AnyObjectCallback: func(i vocab.ObjectType) error {
			t = append(t, i)
			return nil
		},
		AnyLinkCallback: func(i vocab.LinkType) error {
			t = append(t, i)
			return nil
		},
	}
	err = r.Deserialize(m)
	if err != nil {
		return
	}
	// TODO: Support more than one, which will enable creating multiple
	// objects simultaneously.
	if len(t) != 1 {
		err = fmt.Errorf("too many object/links: %d", len(t))
		return
	}
	tid = t[0]
	return
}

func toAnyActivity(m map[string]interface{}) (o vocab.ActivityType, err error) {
	r := &streams.Resolver{
		AnyActivityCallback: func(i vocab.ActivityType) error {
			o = i
			return nil
		},
	}
	err = r.Deserialize(m)
	return
}
