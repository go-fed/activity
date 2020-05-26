// +build generate
//go:generate go run ./astool -spec astool/activitystreams.jsonld -spec astool/security-v1.jsonld -path github.com/go-fed/activity ./streams

package activity
