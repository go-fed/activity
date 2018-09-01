package extension

import (
	"fmt"
	"sync"
)

// The global mapping of names to specific vocabularies and the plugins
// providing their implementations.
var (
	pluginRegistry = make(map[string]Plugin)
	pluginRegistryMu sync.RWMutex
)

func RegisteredPlugins() []string {
	pluginRegistryMu.RLock()
	defer pluginRegistryMu.RUnlock()
	s := make([]string, 0, len(pluginRegistry))
	for name, _ := range pluginRegistry {
		s = append(s, name)
	}
	return s
}

/* BEGIN TEST SPECIFIC STUFF */

func unregisterAllExtensions() {
	pluginRegistryMu.Lock()
	defer pluginRegistryMu.Unlock()
	pluginRegistry = make(map[string]Plugin)
}

/* BEGIN STATIC SPECIFIC STUFF */

func RegisterExtension(name string, plugin Plugin) {
	pluginRegistryMu.Lock()
	defer pluginRegistryMu.Unlock()
	if plugin == nil {
		panic("nil plugin passed to RegisterExtension")
	}
	if _, ok := pluginRegistry[name]; ok {
		panic("duplicate RegisterExtension name: " + name)
	}
	pluginRegistry[name] = plugin
}

/* BEGIN DYNAMIC SPECIFIC STUFF */

var (
	pluginProvider PluginProvider = nil
	pluginProviderMu sync.RWMutex
)

type NamedPlugin struct {
	Name string
	Plugin Plugin
}

type PluginProvider interface {
	Add() <-chan NamedPlugin
	Remove() <-chan string
	// nil errors will be sent in response to successful calls to Add or Remove
	Err() chan<- error
	Done() <-chan struct{}
}

func RegisterPluginProvider(p PluginProvider) {
	pluginProviderMu.Lock()
	defer pluginProviderMu.Unlock()
	if pluginProvider != nil {
		panic("RegisterPluginProvider already called")
	}
	pluginProvider = p
	go func() {
		addCh := pluginProvider.Add()
		remCh := pluginProvider.Remove()
		errCh := pluginProvider.Err()
		doneCh := pluginProvider.Done()
		for {
			select {
			case namedPlugin := <-addCh:
				pluginRegistryMu.Lock() // LOCK
				if _, ok := pluginRegistry[namedPlugin.Name]; ok {
					errCh <- fmt.Errorf("cannot add: plugin already registered: %s", namedPlugin.Name)
				} else {
					pluginRegistry[namedPlugin.Name] = namedPlugin.Plugin
					errCh <- nil
				}
				pluginRegistryMu.Unlock() // UNLOCK
			case removeName := <-remCh:
				pluginRegistryMu.Lock() // LOCK
				if _, ok := pluginRegistry[removeName]; !ok {
					errCh <- fmt.Errorf("cannot remove: plugin not registered: %s", removeName)
				} else {
					delete(pluginRegistry, removeName)
					errCh <- nil
				}
				pluginRegistryMu.Unlock() // UNLOCK
			case <-doneCh:
				close(errCh)
				return
			}
		}
	}()
}
