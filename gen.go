// +build generate
//go:generate go run ./astool -spec astool/activitystreams.jsonld -spec astool/security-v1.jsonld -spec astool/toot.jsonld -spec astool/forgefed.jsonld -path github.com/go-fed/activity ./streams

package activity
