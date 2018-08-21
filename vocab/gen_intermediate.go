//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// repliesIntermediateType will only have one of its values set at most
type repliesIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *repliesIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// startTimeIntermediateType will only have one of its values set at most
type startTimeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *startTimeIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inboxIntermediateType will only have one of its values set at most
type inboxIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inboxIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// originIntermediateType will only have one of its values set at most
type originIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for origin property
	Object ObjectType
	// Stores possible LinkType type for origin property
	Link LinkType
	// Stores possible *url.URL type for origin property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *originIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *originIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// firstIntermediateType will only have one of its values set at most
type firstIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for first property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for first property
	Link LinkType
	// Stores possible *url.URL type for first property
	IRI *url.URL
	// Stores possible OrderedCollectionPageType type for first property
	OrderedCollectionPage OrderedCollectionPageType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *firstIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *firstIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// actorIntermediateType will only have one of its values set at most
type actorIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for actor property
	Object ObjectType
	// Stores possible LinkType type for actor property
	Link LinkType
	// Stores possible *url.URL type for actor property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *actorIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *actorIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// currentIntermediateType will only have one of its values set at most
type currentIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for current property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for current property
	Link LinkType
	// Stores possible *url.URL type for current property
	IRI *url.URL
	// Stores possible OrderedCollectionPageType type for current property
	OrderedCollectionPage OrderedCollectionPageType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *currentIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *currentIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// latitudeIntermediateType will only have one of its values set at most
type latitudeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for latitude property
	float *float64
	// Stores possible *url.URL type for latitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *latitudeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *latitudeIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// iconIntermediateType will only have one of its values set at most
type iconIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for icon property
	Image ImageType
	// Stores possible LinkType type for icon property
	Link LinkType
	// Stores possible *url.URL type for icon property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *iconIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *iconIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// btoIntermediateType will only have one of its values set at most
type btoIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bto property
	Object ObjectType
	// Stores possible LinkType type for bto property
	Link LinkType
	// Stores possible *url.URL type for bto property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *btoIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *btoIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// mediaTypeIntermediateType will only have one of its values set at most
type mediaTypeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.mimeMediaTypeValue, err = mimeMediaTypeValueDeserialize(i)
			if err != nil {
				t.mimeMediaTypeValue = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *mediaTypeIntermediateType) Serialize() (i interface{}, err error) {
	if t.mimeMediaTypeValue != nil {
		i = mimeMediaTypeValueSerialize(*t.mimeMediaTypeValue)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// resultIntermediateType will only have one of its values set at most
type resultIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for result property
	Object ObjectType
	// Stores possible LinkType type for result property
	Link LinkType
	// Stores possible *url.URL type for result property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *resultIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *resultIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// nameIntermediateType will only have one of its values set at most
type nameIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for name property
	stringName *string
	// Stores possible *string type for name property
	langString *string
	// Stores possible *url.URL type for name property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nameIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *nameIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// itemsIntermediateType will only have one of its values set at most
type itemsIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for items property
	Object ObjectType
	// Stores possible LinkType type for items property
	Link LinkType
	// Stores possible *url.URL type for items property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *itemsIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *itemsIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// partOfIntermediateType will only have one of its values set at most
type partOfIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible LinkType type for partOf property
	Link LinkType
	// Stores possible CollectionType type for partOf property
	Collection CollectionType
	// Stores possible *url.URL type for partOf property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *partOfIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *partOfIntermediateType) Serialize() (i interface{}, err error) {
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// updatedIntermediateType will only have one of its values set at most
type updatedIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *updatedIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// accuracyIntermediateType will only have one of its values set at most
type accuracyIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for accuracy property
	float *float64
	// Stores possible *url.URL type for accuracy property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *accuracyIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *accuracyIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// longitudeIntermediateType will only have one of its values set at most
type longitudeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for longitude property
	float *float64
	// Stores possible *url.URL type for longitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *longitudeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *longitudeIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// generatorIntermediateType will only have one of its values set at most
type generatorIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for generator property
	Object ObjectType
	// Stores possible LinkType type for generator property
	Link LinkType
	// Stores possible *url.URL type for generator property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *generatorIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *generatorIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// sharesIntermediateType will only have one of its values set at most
type sharesIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for shares property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for shares property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for shares property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sharesIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *sharesIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// hreflangIntermediateType will only have one of its values set at most
type hreflangIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for hreflang property
	bcp47LanguageTag *string
	// Stores possible *url.URL type for hreflang property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *hreflangIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.bcp47LanguageTag, err = bcp47LanguageTagDeserialize(i)
			if err != nil {
				t.bcp47LanguageTag = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *hreflangIntermediateType) Serialize() (i interface{}, err error) {
	if t.bcp47LanguageTag != nil {
		i = bcp47LanguageTagSerialize(*t.bcp47LanguageTag)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// startIndexIntermediateType will only have one of its values set at most
type startIndexIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for startIndex property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for startIndex property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startIndexIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *startIndexIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// anyOfIntermediateType will only have one of its values set at most
type anyOfIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for anyOf property
	Object ObjectType
	// Stores possible LinkType type for anyOf property
	Link LinkType
	// Stores possible *url.URL type for anyOf property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *anyOfIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *anyOfIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// describesIntermediateType will only have one of its values set at most
type describesIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for describes property
	Object ObjectType
	// Stores possible *url.URL type for describes property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *describesIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *describesIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// outboxIntermediateType will only have one of its values set at most
type outboxIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *outboxIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// unitsIntermediateType will only have one of its values set at most
type unitsIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for units property
	unitsValue *string
	// Stores possible *url.URL type for units property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *unitsIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.unitsValue, err = unitsValueDeserialize(i)
			if err != nil {
				t.unitsValue = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *unitsIntermediateType) Serialize() (i interface{}, err error) {
	if t.unitsValue != nil {
		i = unitsValueSerialize(*t.unitsValue)
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endTimeIntermediateType will only have one of its values set at most
type endTimeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endTimeIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// publishedIntermediateType will only have one of its values set at most
type publishedIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *publishedIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followersIntermediateType will only have one of its values set at most
type followersIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for followers property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for followers property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for followers property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followersIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followersIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// closedIntermediateType will only have one of its values set at most
type closedIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for closed property
	dateTime *time.Time
	// Stores possible *bool type for closed property
	boolean *bool
	// Stores possible ObjectType type for closed property
	Object ObjectType
	// Stores possible LinkType type for closed property
	Link LinkType
	// Stores possible *url.URL type for closed property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *closedIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.boolean, err = booleanDeserialize(i)
			if err != nil {
				t.boolean = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *closedIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.boolean != nil {
		i = booleanSerialize(*t.boolean)
		return
	}
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contentIntermediateType will only have one of its values set at most
type contentIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for content property
	stringName *string
	// Stores possible *string type for content property
	langString *string
	// Stores possible *url.URL type for content property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contentIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contentIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contextIntermediateType will only have one of its values set at most
type contextIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for context property
	Object ObjectType
	// Stores possible LinkType type for context property
	Link LinkType
	// Stores possible *url.URL type for context property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contextIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contextIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// oneOfIntermediateType will only have one of its values set at most
type oneOfIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for oneOf property
	Object ObjectType
	// Stores possible LinkType type for oneOf property
	Link LinkType
	// Stores possible *url.URL type for oneOf property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *oneOfIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *oneOfIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// radiusIntermediateType will only have one of its values set at most
type radiusIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for radius property
	float *float64
	// Stores possible *url.URL type for radius property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *radiusIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *radiusIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attachmentIntermediateType will only have one of its values set at most
type attachmentIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attachment property
	Object ObjectType
	// Stores possible LinkType type for attachment property
	Link LinkType
	// Stores possible *url.URL type for attachment property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attachmentIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attachmentIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// widthIntermediateType will only have one of its values set at most
type widthIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for width property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for width property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *widthIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *widthIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// orderedItemsIntermediateType will only have one of its values set at most
type orderedItemsIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for orderedItems property
	Object ObjectType
	// Stores possible LinkType type for orderedItems property
	Link LinkType
	// Stores possible *url.URL type for orderedItems property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *orderedItemsIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *orderedItemsIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// prevIntermediateType will only have one of its values set at most
type prevIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for prev property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for prev property
	Link LinkType
	// Stores possible *url.URL type for prev property
	IRI *url.URL
	// Stores possible OrderedCollectionPageType type for prev property
	OrderedCollectionPage OrderedCollectionPageType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *prevIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *prevIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// subjectIntermediateType will only have one of its values set at most
type subjectIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for subject property
	Object ObjectType
	// Stores possible LinkType type for subject property
	Link LinkType
	// Stores possible *url.URL type for subject property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *subjectIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *subjectIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// previewIntermediateType will only have one of its values set at most
type previewIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for preview property
	Object ObjectType
	// Stores possible LinkType type for preview property
	Link LinkType
	// Stores possible *url.URL type for preview property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *previewIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *previewIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// urlIntermediateType will only have one of its values set at most
type urlIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *urlIntermediateType) Serialize() (i interface{}, err error) {
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// toIntermediateType will only have one of its values set at most
type toIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for to property
	Object ObjectType
	// Stores possible LinkType type for to property
	Link LinkType
	// Stores possible *url.URL type for to property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *toIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *toIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// bccIntermediateType will only have one of its values set at most
type bccIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bcc property
	Object ObjectType
	// Stores possible LinkType type for bcc property
	Link LinkType
	// Stores possible *url.URL type for bcc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *bccIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *bccIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likesIntermediateType will only have one of its values set at most
type likesIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for likes property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for likes property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for likes property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likesIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likesIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// targetIntermediateType will only have one of its values set at most
type targetIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for target property
	Object ObjectType
	// Stores possible LinkType type for target property
	Link LinkType
	// Stores possible *url.URL type for target property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *targetIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *targetIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// altitudeIntermediateType will only have one of its values set at most
type altitudeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *altitudeIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likedIntermediateType will only have one of its values set at most
type likedIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for liked property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for liked property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for liked property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likedIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likedIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// totalItemsIntermediateType will only have one of its values set at most
type totalItemsIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for totalItems property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for totalItems property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *totalItemsIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *totalItemsIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// relationshipIntermediateType will only have one of its values set at most
type relationshipIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for relationship property
	Object ObjectType
	// Stores possible *url.URL type for relationship property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *relationshipIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *relationshipIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// tagIntermediateType will only have one of its values set at most
type tagIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for tag property
	Object ObjectType
	// Stores possible LinkType type for tag property
	Link LinkType
	// Stores possible *url.URL type for tag property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *tagIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *tagIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// relIntermediateType will only have one of its values set at most
type relIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for rel property
	linkRelation *string
	// Stores possible *url.URL type for rel property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *relIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.linkRelation, err = linkRelationDeserialize(i)
			if err != nil {
				t.linkRelation = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *relIntermediateType) Serialize() (i interface{}, err error) {
	if t.linkRelation != nil {
		i = linkRelationSerialize(*t.linkRelation)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// objectIntermediateType will only have one of its values set at most
type objectIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for object property
	Object ObjectType
	// Stores possible *url.URL type for object property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *objectIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *objectIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// instrumentIntermediateType will only have one of its values set at most
type instrumentIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for instrument property
	Object ObjectType
	// Stores possible LinkType type for instrument property
	Link LinkType
	// Stores possible *url.URL type for instrument property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *instrumentIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *instrumentIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inReplyToIntermediateType will only have one of its values set at most
type inReplyToIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for inReplyTo property
	Object ObjectType
	// Stores possible LinkType type for inReplyTo property
	Link LinkType
	// Stores possible *url.URL type for inReplyTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inReplyToIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inReplyToIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// locationIntermediateType will only have one of its values set at most
type locationIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for location property
	Object ObjectType
	// Stores possible LinkType type for location property
	Link LinkType
	// Stores possible *url.URL type for location property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *locationIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *locationIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// ccIntermediateType will only have one of its values set at most
type ccIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for cc property
	Object ObjectType
	// Stores possible LinkType type for cc property
	Link LinkType
	// Stores possible *url.URL type for cc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *ccIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *ccIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// nextIntermediateType will only have one of its values set at most
type nextIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for next property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for next property
	Link LinkType
	// Stores possible *url.URL type for next property
	IRI *url.URL
	// Stores possible OrderedCollectionPageType type for next property
	OrderedCollectionPage OrderedCollectionPageType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nextIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *nextIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// formerTypeIntermediateType will only have one of its values set at most
type formerTypeIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for formerType property
	stringName *string
	// Stores possible ObjectType type for formerType property
	Object ObjectType
	// Stores possible *url.URL type for formerType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *formerTypeIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *formerTypeIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// imageIntermediateType will only have one of its values set at most
type imageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for image property
	Image ImageType
	// Stores possible LinkType type for image property
	Link LinkType
	// Stores possible *url.URL type for image property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *imageIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *imageIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// durationIntermediateType will only have one of its values set at most
type durationIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.duration, err = durationDeserialize(i)
			if err != nil {
				t.duration = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *durationIntermediateType) Serialize() (i interface{}, err error) {
	if t.duration != nil {
		i = durationSerialize(*t.duration)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followingIntermediateType will only have one of its values set at most
type followingIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for following property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for following property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for following property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followingIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followingIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// preferredUsernameIntermediateType will only have one of its values set at most
type preferredUsernameIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *preferredUsernameIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// heightIntermediateType will only have one of its values set at most
type heightIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for height property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for height property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *heightIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *heightIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attributedToIntermediateType will only have one of its values set at most
type attributedToIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attributedTo property
	Object ObjectType
	// Stores possible LinkType type for attributedTo property
	Link LinkType
	// Stores possible *url.URL type for attributedTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attributedToIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attributedToIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// audienceIntermediateType will only have one of its values set at most
type audienceIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for audience property
	Object ObjectType
	// Stores possible LinkType type for audience property
	Link LinkType
	// Stores possible *url.URL type for audience property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *audienceIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *audienceIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// summaryIntermediateType will only have one of its values set at most
type summaryIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for summary property
	stringName *string
	// Stores possible *string type for summary property
	langString *string
	// Stores possible *url.URL type for summary property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *summaryIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *summaryIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// sourceIntermediateType will only have one of its values set at most
type sourceIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *sourceIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endpointsIntermediateType will only have one of its values set at most
type endpointsIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endpointsIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// lastIntermediateType will only have one of its values set at most
type lastIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for last property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for last property
	Link LinkType
	// Stores possible *url.URL type for last property
	IRI *url.URL
	// Stores possible OrderedCollectionPageType type for last property
	OrderedCollectionPage OrderedCollectionPageType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *lastIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *lastIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// deletedIntermediateType will only have one of its values set at most
type deletedIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for deleted property
	dateTime *time.Time
	// Stores possible *url.URL type for deleted property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *deletedIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *deletedIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// deserializerepliesIntermediateType will accept a map to create a repliesIntermediateType
func deserializeRepliesIntermediateType(in interface{}) (t *repliesIntermediateType, err error) {
	tmp := &repliesIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesIntermediateType will accept a slice to create a slice of repliesIntermediateType
func deserializeSliceRepliesIntermediateType(in []interface{}) (t []*repliesIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesIntermediateType will accept a repliesIntermediateType to create a map
func serializeRepliesIntermediateType(t *repliesIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesIntermediateType will accept a slice of repliesIntermediateType to create a slice result
func serializeSliceRepliesIntermediateType(s []*repliesIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeIntermediateType will accept a map to create a startTimeIntermediateType
func deserializeStartTimeIntermediateType(in interface{}) (t *startTimeIntermediateType, err error) {
	tmp := &startTimeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeIntermediateType will accept a slice to create a slice of startTimeIntermediateType
func deserializeSliceStartTimeIntermediateType(in []interface{}) (t []*startTimeIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeIntermediateType will accept a startTimeIntermediateType to create a map
func serializeStartTimeIntermediateType(t *startTimeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeIntermediateType will accept a slice of startTimeIntermediateType to create a slice result
func serializeSliceStartTimeIntermediateType(s []*startTimeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxIntermediateType will accept a map to create a inboxIntermediateType
func deserializeInboxIntermediateType(in interface{}) (t *inboxIntermediateType, err error) {
	tmp := &inboxIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxIntermediateType will accept a slice to create a slice of inboxIntermediateType
func deserializeSliceInboxIntermediateType(in []interface{}) (t []*inboxIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxIntermediateType will accept a inboxIntermediateType to create a map
func serializeInboxIntermediateType(t *inboxIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxIntermediateType will accept a slice of inboxIntermediateType to create a slice result
func serializeSliceInboxIntermediateType(s []*inboxIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoriginIntermediateType will accept a map to create a originIntermediateType
func deserializeOriginIntermediateType(in interface{}) (t *originIntermediateType, err error) {
	tmp := &originIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice originIntermediateType will accept a slice to create a slice of originIntermediateType
func deserializeSliceOriginIntermediateType(in []interface{}) (t []*originIntermediateType, err error) {
	for _, i := range in {
		tmp := &originIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoriginIntermediateType will accept a originIntermediateType to create a map
func serializeOriginIntermediateType(t *originIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoriginIntermediateType will accept a slice of originIntermediateType to create a slice result
func serializeSliceOriginIntermediateType(s []*originIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefirstIntermediateType will accept a map to create a firstIntermediateType
func deserializeFirstIntermediateType(in interface{}) (t *firstIntermediateType, err error) {
	tmp := &firstIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice firstIntermediateType will accept a slice to create a slice of firstIntermediateType
func deserializeSliceFirstIntermediateType(in []interface{}) (t []*firstIntermediateType, err error) {
	for _, i := range in {
		tmp := &firstIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefirstIntermediateType will accept a firstIntermediateType to create a map
func serializeFirstIntermediateType(t *firstIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefirstIntermediateType will accept a slice of firstIntermediateType to create a slice result
func serializeSliceFirstIntermediateType(s []*firstIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeactorIntermediateType will accept a map to create a actorIntermediateType
func deserializeActorIntermediateType(in interface{}) (t *actorIntermediateType, err error) {
	tmp := &actorIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice actorIntermediateType will accept a slice to create a slice of actorIntermediateType
func deserializeSliceActorIntermediateType(in []interface{}) (t []*actorIntermediateType, err error) {
	for _, i := range in {
		tmp := &actorIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeactorIntermediateType will accept a actorIntermediateType to create a map
func serializeActorIntermediateType(t *actorIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceactorIntermediateType will accept a slice of actorIntermediateType to create a slice result
func serializeSliceActorIntermediateType(s []*actorIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecurrentIntermediateType will accept a map to create a currentIntermediateType
func deserializeCurrentIntermediateType(in interface{}) (t *currentIntermediateType, err error) {
	tmp := &currentIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice currentIntermediateType will accept a slice to create a slice of currentIntermediateType
func deserializeSliceCurrentIntermediateType(in []interface{}) (t []*currentIntermediateType, err error) {
	for _, i := range in {
		tmp := &currentIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecurrentIntermediateType will accept a currentIntermediateType to create a map
func serializeCurrentIntermediateType(t *currentIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecurrentIntermediateType will accept a slice of currentIntermediateType to create a slice result
func serializeSliceCurrentIntermediateType(s []*currentIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelatitudeIntermediateType will accept a map to create a latitudeIntermediateType
func deserializeLatitudeIntermediateType(in interface{}) (t *latitudeIntermediateType, err error) {
	tmp := &latitudeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice latitudeIntermediateType will accept a slice to create a slice of latitudeIntermediateType
func deserializeSliceLatitudeIntermediateType(in []interface{}) (t []*latitudeIntermediateType, err error) {
	for _, i := range in {
		tmp := &latitudeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelatitudeIntermediateType will accept a latitudeIntermediateType to create a map
func serializeLatitudeIntermediateType(t *latitudeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelatitudeIntermediateType will accept a slice of latitudeIntermediateType to create a slice result
func serializeSliceLatitudeIntermediateType(s []*latitudeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconIntermediateType will accept a map to create a iconIntermediateType
func deserializeIconIntermediateType(in interface{}) (t *iconIntermediateType, err error) {
	tmp := &iconIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconIntermediateType will accept a slice to create a slice of iconIntermediateType
func deserializeSliceIconIntermediateType(in []interface{}) (t []*iconIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconIntermediateType will accept a iconIntermediateType to create a map
func serializeIconIntermediateType(t *iconIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconIntermediateType will accept a slice of iconIntermediateType to create a slice result
func serializeSliceIconIntermediateType(s []*iconIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoIntermediateType will accept a map to create a btoIntermediateType
func deserializeBtoIntermediateType(in interface{}) (t *btoIntermediateType, err error) {
	tmp := &btoIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoIntermediateType will accept a slice to create a slice of btoIntermediateType
func deserializeSliceBtoIntermediateType(in []interface{}) (t []*btoIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoIntermediateType will accept a btoIntermediateType to create a map
func serializeBtoIntermediateType(t *btoIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoIntermediateType will accept a slice of btoIntermediateType to create a slice result
func serializeSliceBtoIntermediateType(s []*btoIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeIntermediateType will accept a map to create a mediaTypeIntermediateType
func deserializeMediaTypeIntermediateType(in interface{}) (t *mediaTypeIntermediateType, err error) {
	tmp := &mediaTypeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeIntermediateType will accept a slice to create a slice of mediaTypeIntermediateType
func deserializeSliceMediaTypeIntermediateType(in []interface{}) (t []*mediaTypeIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeIntermediateType will accept a mediaTypeIntermediateType to create a map
func serializeMediaTypeIntermediateType(t *mediaTypeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeIntermediateType will accept a slice of mediaTypeIntermediateType to create a slice result
func serializeSliceMediaTypeIntermediateType(s []*mediaTypeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeresultIntermediateType will accept a map to create a resultIntermediateType
func deserializeResultIntermediateType(in interface{}) (t *resultIntermediateType, err error) {
	tmp := &resultIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice resultIntermediateType will accept a slice to create a slice of resultIntermediateType
func deserializeSliceResultIntermediateType(in []interface{}) (t []*resultIntermediateType, err error) {
	for _, i := range in {
		tmp := &resultIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeresultIntermediateType will accept a resultIntermediateType to create a map
func serializeResultIntermediateType(t *resultIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceresultIntermediateType will accept a slice of resultIntermediateType to create a slice result
func serializeSliceResultIntermediateType(s []*resultIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameIntermediateType will accept a map to create a nameIntermediateType
func deserializeNameIntermediateType(in interface{}) (t *nameIntermediateType, err error) {
	tmp := &nameIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameIntermediateType will accept a slice to create a slice of nameIntermediateType
func deserializeSliceNameIntermediateType(in []interface{}) (t []*nameIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameIntermediateType will accept a nameIntermediateType to create a map
func serializeNameIntermediateType(t *nameIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameIntermediateType will accept a slice of nameIntermediateType to create a slice result
func serializeSliceNameIntermediateType(s []*nameIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeitemsIntermediateType will accept a map to create a itemsIntermediateType
func deserializeItemsIntermediateType(in interface{}) (t *itemsIntermediateType, err error) {
	tmp := &itemsIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice itemsIntermediateType will accept a slice to create a slice of itemsIntermediateType
func deserializeSliceItemsIntermediateType(in []interface{}) (t []*itemsIntermediateType, err error) {
	for _, i := range in {
		tmp := &itemsIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeitemsIntermediateType will accept a itemsIntermediateType to create a map
func serializeItemsIntermediateType(t *itemsIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceitemsIntermediateType will accept a slice of itemsIntermediateType to create a slice result
func serializeSliceItemsIntermediateType(s []*itemsIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepartOfIntermediateType will accept a map to create a partOfIntermediateType
func deserializePartOfIntermediateType(in interface{}) (t *partOfIntermediateType, err error) {
	tmp := &partOfIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice partOfIntermediateType will accept a slice to create a slice of partOfIntermediateType
func deserializeSlicePartOfIntermediateType(in []interface{}) (t []*partOfIntermediateType, err error) {
	for _, i := range in {
		tmp := &partOfIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepartOfIntermediateType will accept a partOfIntermediateType to create a map
func serializePartOfIntermediateType(t *partOfIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepartOfIntermediateType will accept a slice of partOfIntermediateType to create a slice result
func serializeSlicePartOfIntermediateType(s []*partOfIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedIntermediateType will accept a map to create a updatedIntermediateType
func deserializeUpdatedIntermediateType(in interface{}) (t *updatedIntermediateType, err error) {
	tmp := &updatedIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedIntermediateType will accept a slice to create a slice of updatedIntermediateType
func deserializeSliceUpdatedIntermediateType(in []interface{}) (t []*updatedIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedIntermediateType will accept a updatedIntermediateType to create a map
func serializeUpdatedIntermediateType(t *updatedIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedIntermediateType will accept a slice of updatedIntermediateType to create a slice result
func serializeSliceUpdatedIntermediateType(s []*updatedIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaccuracyIntermediateType will accept a map to create a accuracyIntermediateType
func deserializeAccuracyIntermediateType(in interface{}) (t *accuracyIntermediateType, err error) {
	tmp := &accuracyIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice accuracyIntermediateType will accept a slice to create a slice of accuracyIntermediateType
func deserializeSliceAccuracyIntermediateType(in []interface{}) (t []*accuracyIntermediateType, err error) {
	for _, i := range in {
		tmp := &accuracyIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaccuracyIntermediateType will accept a accuracyIntermediateType to create a map
func serializeAccuracyIntermediateType(t *accuracyIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaccuracyIntermediateType will accept a slice of accuracyIntermediateType to create a slice result
func serializeSliceAccuracyIntermediateType(s []*accuracyIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelongitudeIntermediateType will accept a map to create a longitudeIntermediateType
func deserializeLongitudeIntermediateType(in interface{}) (t *longitudeIntermediateType, err error) {
	tmp := &longitudeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice longitudeIntermediateType will accept a slice to create a slice of longitudeIntermediateType
func deserializeSliceLongitudeIntermediateType(in []interface{}) (t []*longitudeIntermediateType, err error) {
	for _, i := range in {
		tmp := &longitudeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelongitudeIntermediateType will accept a longitudeIntermediateType to create a map
func serializeLongitudeIntermediateType(t *longitudeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelongitudeIntermediateType will accept a slice of longitudeIntermediateType to create a slice result
func serializeSliceLongitudeIntermediateType(s []*longitudeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorIntermediateType will accept a map to create a generatorIntermediateType
func deserializeGeneratorIntermediateType(in interface{}) (t *generatorIntermediateType, err error) {
	tmp := &generatorIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorIntermediateType will accept a slice to create a slice of generatorIntermediateType
func deserializeSliceGeneratorIntermediateType(in []interface{}) (t []*generatorIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorIntermediateType will accept a generatorIntermediateType to create a map
func serializeGeneratorIntermediateType(t *generatorIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorIntermediateType will accept a slice of generatorIntermediateType to create a slice result
func serializeSliceGeneratorIntermediateType(s []*generatorIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesharesIntermediateType will accept a map to create a sharesIntermediateType
func deserializeSharesIntermediateType(in interface{}) (t *sharesIntermediateType, err error) {
	tmp := &sharesIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sharesIntermediateType will accept a slice to create a slice of sharesIntermediateType
func deserializeSliceSharesIntermediateType(in []interface{}) (t []*sharesIntermediateType, err error) {
	for _, i := range in {
		tmp := &sharesIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesharesIntermediateType will accept a sharesIntermediateType to create a map
func serializeSharesIntermediateType(t *sharesIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesharesIntermediateType will accept a slice of sharesIntermediateType to create a slice result
func serializeSliceSharesIntermediateType(s []*sharesIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializehreflangIntermediateType will accept a map to create a hreflangIntermediateType
func deserializeHreflangIntermediateType(in interface{}) (t *hreflangIntermediateType, err error) {
	tmp := &hreflangIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice hreflangIntermediateType will accept a slice to create a slice of hreflangIntermediateType
func deserializeSliceHreflangIntermediateType(in []interface{}) (t []*hreflangIntermediateType, err error) {
	for _, i := range in {
		tmp := &hreflangIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializehreflangIntermediateType will accept a hreflangIntermediateType to create a map
func serializeHreflangIntermediateType(t *hreflangIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicehreflangIntermediateType will accept a slice of hreflangIntermediateType to create a slice result
func serializeSliceHreflangIntermediateType(s []*hreflangIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartIndexIntermediateType will accept a map to create a startIndexIntermediateType
func deserializeStartIndexIntermediateType(in interface{}) (t *startIndexIntermediateType, err error) {
	tmp := &startIndexIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startIndexIntermediateType will accept a slice to create a slice of startIndexIntermediateType
func deserializeSliceStartIndexIntermediateType(in []interface{}) (t []*startIndexIntermediateType, err error) {
	for _, i := range in {
		tmp := &startIndexIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartIndexIntermediateType will accept a startIndexIntermediateType to create a map
func serializeStartIndexIntermediateType(t *startIndexIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartIndexIntermediateType will accept a slice of startIndexIntermediateType to create a slice result
func serializeSliceStartIndexIntermediateType(s []*startIndexIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeanyOfIntermediateType will accept a map to create a anyOfIntermediateType
func deserializeAnyOfIntermediateType(in interface{}) (t *anyOfIntermediateType, err error) {
	tmp := &anyOfIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice anyOfIntermediateType will accept a slice to create a slice of anyOfIntermediateType
func deserializeSliceAnyOfIntermediateType(in []interface{}) (t []*anyOfIntermediateType, err error) {
	for _, i := range in {
		tmp := &anyOfIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeanyOfIntermediateType will accept a anyOfIntermediateType to create a map
func serializeAnyOfIntermediateType(t *anyOfIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceanyOfIntermediateType will accept a slice of anyOfIntermediateType to create a slice result
func serializeSliceAnyOfIntermediateType(s []*anyOfIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedescribesIntermediateType will accept a map to create a describesIntermediateType
func deserializeDescribesIntermediateType(in interface{}) (t *describesIntermediateType, err error) {
	tmp := &describesIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice describesIntermediateType will accept a slice to create a slice of describesIntermediateType
func deserializeSliceDescribesIntermediateType(in []interface{}) (t []*describesIntermediateType, err error) {
	for _, i := range in {
		tmp := &describesIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedescribesIntermediateType will accept a describesIntermediateType to create a map
func serializeDescribesIntermediateType(t *describesIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedescribesIntermediateType will accept a slice of describesIntermediateType to create a slice result
func serializeSliceDescribesIntermediateType(s []*describesIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxIntermediateType will accept a map to create a outboxIntermediateType
func deserializeOutboxIntermediateType(in interface{}) (t *outboxIntermediateType, err error) {
	tmp := &outboxIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxIntermediateType will accept a slice to create a slice of outboxIntermediateType
func deserializeSliceOutboxIntermediateType(in []interface{}) (t []*outboxIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxIntermediateType will accept a outboxIntermediateType to create a map
func serializeOutboxIntermediateType(t *outboxIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxIntermediateType will accept a slice of outboxIntermediateType to create a slice result
func serializeSliceOutboxIntermediateType(s []*outboxIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeunitsIntermediateType will accept a map to create a unitsIntermediateType
func deserializeUnitsIntermediateType(in interface{}) (t *unitsIntermediateType, err error) {
	tmp := &unitsIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice unitsIntermediateType will accept a slice to create a slice of unitsIntermediateType
func deserializeSliceUnitsIntermediateType(in []interface{}) (t []*unitsIntermediateType, err error) {
	for _, i := range in {
		tmp := &unitsIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeunitsIntermediateType will accept a unitsIntermediateType to create a map
func serializeUnitsIntermediateType(t *unitsIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceunitsIntermediateType will accept a slice of unitsIntermediateType to create a slice result
func serializeSliceUnitsIntermediateType(s []*unitsIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeIntermediateType will accept a map to create a endTimeIntermediateType
func deserializeEndTimeIntermediateType(in interface{}) (t *endTimeIntermediateType, err error) {
	tmp := &endTimeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeIntermediateType will accept a slice to create a slice of endTimeIntermediateType
func deserializeSliceEndTimeIntermediateType(in []interface{}) (t []*endTimeIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeIntermediateType will accept a endTimeIntermediateType to create a map
func serializeEndTimeIntermediateType(t *endTimeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeIntermediateType will accept a slice of endTimeIntermediateType to create a slice result
func serializeSliceEndTimeIntermediateType(s []*endTimeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedIntermediateType will accept a map to create a publishedIntermediateType
func deserializePublishedIntermediateType(in interface{}) (t *publishedIntermediateType, err error) {
	tmp := &publishedIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedIntermediateType will accept a slice to create a slice of publishedIntermediateType
func deserializeSlicePublishedIntermediateType(in []interface{}) (t []*publishedIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedIntermediateType will accept a publishedIntermediateType to create a map
func serializePublishedIntermediateType(t *publishedIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedIntermediateType will accept a slice of publishedIntermediateType to create a slice result
func serializeSlicePublishedIntermediateType(s []*publishedIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersIntermediateType will accept a map to create a followersIntermediateType
func deserializeFollowersIntermediateType(in interface{}) (t *followersIntermediateType, err error) {
	tmp := &followersIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersIntermediateType will accept a slice to create a slice of followersIntermediateType
func deserializeSliceFollowersIntermediateType(in []interface{}) (t []*followersIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersIntermediateType will accept a followersIntermediateType to create a map
func serializeFollowersIntermediateType(t *followersIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersIntermediateType will accept a slice of followersIntermediateType to create a slice result
func serializeSliceFollowersIntermediateType(s []*followersIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeclosedIntermediateType will accept a map to create a closedIntermediateType
func deserializeClosedIntermediateType(in interface{}) (t *closedIntermediateType, err error) {
	tmp := &closedIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice closedIntermediateType will accept a slice to create a slice of closedIntermediateType
func deserializeSliceClosedIntermediateType(in []interface{}) (t []*closedIntermediateType, err error) {
	for _, i := range in {
		tmp := &closedIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeclosedIntermediateType will accept a closedIntermediateType to create a map
func serializeClosedIntermediateType(t *closedIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceclosedIntermediateType will accept a slice of closedIntermediateType to create a slice result
func serializeSliceClosedIntermediateType(s []*closedIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentIntermediateType will accept a map to create a contentIntermediateType
func deserializeContentIntermediateType(in interface{}) (t *contentIntermediateType, err error) {
	tmp := &contentIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentIntermediateType will accept a slice to create a slice of contentIntermediateType
func deserializeSliceContentIntermediateType(in []interface{}) (t []*contentIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentIntermediateType will accept a contentIntermediateType to create a map
func serializeContentIntermediateType(t *contentIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentIntermediateType will accept a slice of contentIntermediateType to create a slice result
func serializeSliceContentIntermediateType(s []*contentIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextIntermediateType will accept a map to create a contextIntermediateType
func deserializeContextIntermediateType(in interface{}) (t *contextIntermediateType, err error) {
	tmp := &contextIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextIntermediateType will accept a slice to create a slice of contextIntermediateType
func deserializeSliceContextIntermediateType(in []interface{}) (t []*contextIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextIntermediateType will accept a contextIntermediateType to create a map
func serializeContextIntermediateType(t *contextIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextIntermediateType will accept a slice of contextIntermediateType to create a slice result
func serializeSliceContextIntermediateType(s []*contextIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoneOfIntermediateType will accept a map to create a oneOfIntermediateType
func deserializeOneOfIntermediateType(in interface{}) (t *oneOfIntermediateType, err error) {
	tmp := &oneOfIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice oneOfIntermediateType will accept a slice to create a slice of oneOfIntermediateType
func deserializeSliceOneOfIntermediateType(in []interface{}) (t []*oneOfIntermediateType, err error) {
	for _, i := range in {
		tmp := &oneOfIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoneOfIntermediateType will accept a oneOfIntermediateType to create a map
func serializeOneOfIntermediateType(t *oneOfIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoneOfIntermediateType will accept a slice of oneOfIntermediateType to create a slice result
func serializeSliceOneOfIntermediateType(s []*oneOfIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeradiusIntermediateType will accept a map to create a radiusIntermediateType
func deserializeRadiusIntermediateType(in interface{}) (t *radiusIntermediateType, err error) {
	tmp := &radiusIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice radiusIntermediateType will accept a slice to create a slice of radiusIntermediateType
func deserializeSliceRadiusIntermediateType(in []interface{}) (t []*radiusIntermediateType, err error) {
	for _, i := range in {
		tmp := &radiusIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeradiusIntermediateType will accept a radiusIntermediateType to create a map
func serializeRadiusIntermediateType(t *radiusIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceradiusIntermediateType will accept a slice of radiusIntermediateType to create a slice result
func serializeSliceRadiusIntermediateType(s []*radiusIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentIntermediateType will accept a map to create a attachmentIntermediateType
func deserializeAttachmentIntermediateType(in interface{}) (t *attachmentIntermediateType, err error) {
	tmp := &attachmentIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentIntermediateType will accept a slice to create a slice of attachmentIntermediateType
func deserializeSliceAttachmentIntermediateType(in []interface{}) (t []*attachmentIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentIntermediateType will accept a attachmentIntermediateType to create a map
func serializeAttachmentIntermediateType(t *attachmentIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentIntermediateType will accept a slice of attachmentIntermediateType to create a slice result
func serializeSliceAttachmentIntermediateType(s []*attachmentIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializewidthIntermediateType will accept a map to create a widthIntermediateType
func deserializeWidthIntermediateType(in interface{}) (t *widthIntermediateType, err error) {
	tmp := &widthIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice widthIntermediateType will accept a slice to create a slice of widthIntermediateType
func deserializeSliceWidthIntermediateType(in []interface{}) (t []*widthIntermediateType, err error) {
	for _, i := range in {
		tmp := &widthIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializewidthIntermediateType will accept a widthIntermediateType to create a map
func serializeWidthIntermediateType(t *widthIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicewidthIntermediateType will accept a slice of widthIntermediateType to create a slice result
func serializeSliceWidthIntermediateType(s []*widthIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeorderedItemsIntermediateType will accept a map to create a orderedItemsIntermediateType
func deserializeOrderedItemsIntermediateType(in interface{}) (t *orderedItemsIntermediateType, err error) {
	tmp := &orderedItemsIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice orderedItemsIntermediateType will accept a slice to create a slice of orderedItemsIntermediateType
func deserializeSliceOrderedItemsIntermediateType(in []interface{}) (t []*orderedItemsIntermediateType, err error) {
	for _, i := range in {
		tmp := &orderedItemsIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeorderedItemsIntermediateType will accept a orderedItemsIntermediateType to create a map
func serializeOrderedItemsIntermediateType(t *orderedItemsIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceorderedItemsIntermediateType will accept a slice of orderedItemsIntermediateType to create a slice result
func serializeSliceOrderedItemsIntermediateType(s []*orderedItemsIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeprevIntermediateType will accept a map to create a prevIntermediateType
func deserializePrevIntermediateType(in interface{}) (t *prevIntermediateType, err error) {
	tmp := &prevIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice prevIntermediateType will accept a slice to create a slice of prevIntermediateType
func deserializeSlicePrevIntermediateType(in []interface{}) (t []*prevIntermediateType, err error) {
	for _, i := range in {
		tmp := &prevIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeprevIntermediateType will accept a prevIntermediateType to create a map
func serializePrevIntermediateType(t *prevIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceprevIntermediateType will accept a slice of prevIntermediateType to create a slice result
func serializeSlicePrevIntermediateType(s []*prevIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesubjectIntermediateType will accept a map to create a subjectIntermediateType
func deserializeSubjectIntermediateType(in interface{}) (t *subjectIntermediateType, err error) {
	tmp := &subjectIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice subjectIntermediateType will accept a slice to create a slice of subjectIntermediateType
func deserializeSliceSubjectIntermediateType(in []interface{}) (t []*subjectIntermediateType, err error) {
	for _, i := range in {
		tmp := &subjectIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesubjectIntermediateType will accept a subjectIntermediateType to create a map
func serializeSubjectIntermediateType(t *subjectIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesubjectIntermediateType will accept a slice of subjectIntermediateType to create a slice result
func serializeSliceSubjectIntermediateType(s []*subjectIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewIntermediateType will accept a map to create a previewIntermediateType
func deserializePreviewIntermediateType(in interface{}) (t *previewIntermediateType, err error) {
	tmp := &previewIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewIntermediateType will accept a slice to create a slice of previewIntermediateType
func deserializeSlicePreviewIntermediateType(in []interface{}) (t []*previewIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewIntermediateType will accept a previewIntermediateType to create a map
func serializePreviewIntermediateType(t *previewIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewIntermediateType will accept a slice of previewIntermediateType to create a slice result
func serializeSlicePreviewIntermediateType(s []*previewIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlIntermediateType will accept a map to create a urlIntermediateType
func deserializeUrlIntermediateType(in interface{}) (t *urlIntermediateType, err error) {
	tmp := &urlIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlIntermediateType will accept a slice to create a slice of urlIntermediateType
func deserializeSliceUrlIntermediateType(in []interface{}) (t []*urlIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlIntermediateType will accept a urlIntermediateType to create a map
func serializeUrlIntermediateType(t *urlIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlIntermediateType will accept a slice of urlIntermediateType to create a slice result
func serializeSliceUrlIntermediateType(s []*urlIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoIntermediateType will accept a map to create a toIntermediateType
func deserializeToIntermediateType(in interface{}) (t *toIntermediateType, err error) {
	tmp := &toIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toIntermediateType will accept a slice to create a slice of toIntermediateType
func deserializeSliceToIntermediateType(in []interface{}) (t []*toIntermediateType, err error) {
	for _, i := range in {
		tmp := &toIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoIntermediateType will accept a toIntermediateType to create a map
func serializeToIntermediateType(t *toIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoIntermediateType will accept a slice of toIntermediateType to create a slice result
func serializeSliceToIntermediateType(s []*toIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccIntermediateType will accept a map to create a bccIntermediateType
func deserializeBccIntermediateType(in interface{}) (t *bccIntermediateType, err error) {
	tmp := &bccIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccIntermediateType will accept a slice to create a slice of bccIntermediateType
func deserializeSliceBccIntermediateType(in []interface{}) (t []*bccIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccIntermediateType will accept a bccIntermediateType to create a map
func serializeBccIntermediateType(t *bccIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccIntermediateType will accept a slice of bccIntermediateType to create a slice result
func serializeSliceBccIntermediateType(s []*bccIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesIntermediateType will accept a map to create a likesIntermediateType
func deserializeLikesIntermediateType(in interface{}) (t *likesIntermediateType, err error) {
	tmp := &likesIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesIntermediateType will accept a slice to create a slice of likesIntermediateType
func deserializeSliceLikesIntermediateType(in []interface{}) (t []*likesIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesIntermediateType will accept a likesIntermediateType to create a map
func serializeLikesIntermediateType(t *likesIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesIntermediateType will accept a slice of likesIntermediateType to create a slice result
func serializeSliceLikesIntermediateType(s []*likesIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetargetIntermediateType will accept a map to create a targetIntermediateType
func deserializeTargetIntermediateType(in interface{}) (t *targetIntermediateType, err error) {
	tmp := &targetIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice targetIntermediateType will accept a slice to create a slice of targetIntermediateType
func deserializeSliceTargetIntermediateType(in []interface{}) (t []*targetIntermediateType, err error) {
	for _, i := range in {
		tmp := &targetIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetargetIntermediateType will accept a targetIntermediateType to create a map
func serializeTargetIntermediateType(t *targetIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetargetIntermediateType will accept a slice of targetIntermediateType to create a slice result
func serializeSliceTargetIntermediateType(s []*targetIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeIntermediateType will accept a map to create a altitudeIntermediateType
func deserializeAltitudeIntermediateType(in interface{}) (t *altitudeIntermediateType, err error) {
	tmp := &altitudeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeIntermediateType will accept a slice to create a slice of altitudeIntermediateType
func deserializeSliceAltitudeIntermediateType(in []interface{}) (t []*altitudeIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeIntermediateType will accept a altitudeIntermediateType to create a map
func serializeAltitudeIntermediateType(t *altitudeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeIntermediateType will accept a slice of altitudeIntermediateType to create a slice result
func serializeSliceAltitudeIntermediateType(s []*altitudeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedIntermediateType will accept a map to create a likedIntermediateType
func deserializeLikedIntermediateType(in interface{}) (t *likedIntermediateType, err error) {
	tmp := &likedIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedIntermediateType will accept a slice to create a slice of likedIntermediateType
func deserializeSliceLikedIntermediateType(in []interface{}) (t []*likedIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedIntermediateType will accept a likedIntermediateType to create a map
func serializeLikedIntermediateType(t *likedIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedIntermediateType will accept a slice of likedIntermediateType to create a slice result
func serializeSliceLikedIntermediateType(s []*likedIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetotalItemsIntermediateType will accept a map to create a totalItemsIntermediateType
func deserializeTotalItemsIntermediateType(in interface{}) (t *totalItemsIntermediateType, err error) {
	tmp := &totalItemsIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice totalItemsIntermediateType will accept a slice to create a slice of totalItemsIntermediateType
func deserializeSliceTotalItemsIntermediateType(in []interface{}) (t []*totalItemsIntermediateType, err error) {
	for _, i := range in {
		tmp := &totalItemsIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetotalItemsIntermediateType will accept a totalItemsIntermediateType to create a map
func serializeTotalItemsIntermediateType(t *totalItemsIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetotalItemsIntermediateType will accept a slice of totalItemsIntermediateType to create a slice result
func serializeSliceTotalItemsIntermediateType(s []*totalItemsIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerelationshipIntermediateType will accept a map to create a relationshipIntermediateType
func deserializeRelationshipIntermediateType(in interface{}) (t *relationshipIntermediateType, err error) {
	tmp := &relationshipIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice relationshipIntermediateType will accept a slice to create a slice of relationshipIntermediateType
func deserializeSliceRelationshipIntermediateType(in []interface{}) (t []*relationshipIntermediateType, err error) {
	for _, i := range in {
		tmp := &relationshipIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerelationshipIntermediateType will accept a relationshipIntermediateType to create a map
func serializeRelationshipIntermediateType(t *relationshipIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerelationshipIntermediateType will accept a slice of relationshipIntermediateType to create a slice result
func serializeSliceRelationshipIntermediateType(s []*relationshipIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagIntermediateType will accept a map to create a tagIntermediateType
func deserializeTagIntermediateType(in interface{}) (t *tagIntermediateType, err error) {
	tmp := &tagIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagIntermediateType will accept a slice to create a slice of tagIntermediateType
func deserializeSliceTagIntermediateType(in []interface{}) (t []*tagIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagIntermediateType will accept a tagIntermediateType to create a map
func serializeTagIntermediateType(t *tagIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagIntermediateType will accept a slice of tagIntermediateType to create a slice result
func serializeSliceTagIntermediateType(s []*tagIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerelIntermediateType will accept a map to create a relIntermediateType
func deserializeRelIntermediateType(in interface{}) (t *relIntermediateType, err error) {
	tmp := &relIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice relIntermediateType will accept a slice to create a slice of relIntermediateType
func deserializeSliceRelIntermediateType(in []interface{}) (t []*relIntermediateType, err error) {
	for _, i := range in {
		tmp := &relIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerelIntermediateType will accept a relIntermediateType to create a map
func serializeRelIntermediateType(t *relIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerelIntermediateType will accept a slice of relIntermediateType to create a slice result
func serializeSliceRelIntermediateType(s []*relIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeobjectIntermediateType will accept a map to create a objectIntermediateType
func deserializeObjectIntermediateType(in interface{}) (t *objectIntermediateType, err error) {
	tmp := &objectIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice objectIntermediateType will accept a slice to create a slice of objectIntermediateType
func deserializeSliceObjectIntermediateType(in []interface{}) (t []*objectIntermediateType, err error) {
	for _, i := range in {
		tmp := &objectIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeobjectIntermediateType will accept a objectIntermediateType to create a map
func serializeObjectIntermediateType(t *objectIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceobjectIntermediateType will accept a slice of objectIntermediateType to create a slice result
func serializeSliceObjectIntermediateType(s []*objectIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinstrumentIntermediateType will accept a map to create a instrumentIntermediateType
func deserializeInstrumentIntermediateType(in interface{}) (t *instrumentIntermediateType, err error) {
	tmp := &instrumentIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice instrumentIntermediateType will accept a slice to create a slice of instrumentIntermediateType
func deserializeSliceInstrumentIntermediateType(in []interface{}) (t []*instrumentIntermediateType, err error) {
	for _, i := range in {
		tmp := &instrumentIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinstrumentIntermediateType will accept a instrumentIntermediateType to create a map
func serializeInstrumentIntermediateType(t *instrumentIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinstrumentIntermediateType will accept a slice of instrumentIntermediateType to create a slice result
func serializeSliceInstrumentIntermediateType(s []*instrumentIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToIntermediateType will accept a map to create a inReplyToIntermediateType
func deserializeInReplyToIntermediateType(in interface{}) (t *inReplyToIntermediateType, err error) {
	tmp := &inReplyToIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToIntermediateType will accept a slice to create a slice of inReplyToIntermediateType
func deserializeSliceInReplyToIntermediateType(in []interface{}) (t []*inReplyToIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToIntermediateType will accept a inReplyToIntermediateType to create a map
func serializeInReplyToIntermediateType(t *inReplyToIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToIntermediateType will accept a slice of inReplyToIntermediateType to create a slice result
func serializeSliceInReplyToIntermediateType(s []*inReplyToIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationIntermediateType will accept a map to create a locationIntermediateType
func deserializeLocationIntermediateType(in interface{}) (t *locationIntermediateType, err error) {
	tmp := &locationIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationIntermediateType will accept a slice to create a slice of locationIntermediateType
func deserializeSliceLocationIntermediateType(in []interface{}) (t []*locationIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationIntermediateType will accept a locationIntermediateType to create a map
func serializeLocationIntermediateType(t *locationIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationIntermediateType will accept a slice of locationIntermediateType to create a slice result
func serializeSliceLocationIntermediateType(s []*locationIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccIntermediateType will accept a map to create a ccIntermediateType
func deserializeCcIntermediateType(in interface{}) (t *ccIntermediateType, err error) {
	tmp := &ccIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccIntermediateType will accept a slice to create a slice of ccIntermediateType
func deserializeSliceCcIntermediateType(in []interface{}) (t []*ccIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccIntermediateType will accept a ccIntermediateType to create a map
func serializeCcIntermediateType(t *ccIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccIntermediateType will accept a slice of ccIntermediateType to create a slice result
func serializeSliceCcIntermediateType(s []*ccIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenextIntermediateType will accept a map to create a nextIntermediateType
func deserializeNextIntermediateType(in interface{}) (t *nextIntermediateType, err error) {
	tmp := &nextIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nextIntermediateType will accept a slice to create a slice of nextIntermediateType
func deserializeSliceNextIntermediateType(in []interface{}) (t []*nextIntermediateType, err error) {
	for _, i := range in {
		tmp := &nextIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenextIntermediateType will accept a nextIntermediateType to create a map
func serializeNextIntermediateType(t *nextIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenextIntermediateType will accept a slice of nextIntermediateType to create a slice result
func serializeSliceNextIntermediateType(s []*nextIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeformerTypeIntermediateType will accept a map to create a formerTypeIntermediateType
func deserializeFormerTypeIntermediateType(in interface{}) (t *formerTypeIntermediateType, err error) {
	tmp := &formerTypeIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice formerTypeIntermediateType will accept a slice to create a slice of formerTypeIntermediateType
func deserializeSliceFormerTypeIntermediateType(in []interface{}) (t []*formerTypeIntermediateType, err error) {
	for _, i := range in {
		tmp := &formerTypeIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeformerTypeIntermediateType will accept a formerTypeIntermediateType to create a map
func serializeFormerTypeIntermediateType(t *formerTypeIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceformerTypeIntermediateType will accept a slice of formerTypeIntermediateType to create a slice result
func serializeSliceFormerTypeIntermediateType(s []*formerTypeIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageIntermediateType will accept a map to create a imageIntermediateType
func deserializeImageIntermediateType(in interface{}) (t *imageIntermediateType, err error) {
	tmp := &imageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageIntermediateType will accept a slice to create a slice of imageIntermediateType
func deserializeSliceImageIntermediateType(in []interface{}) (t []*imageIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageIntermediateType will accept a imageIntermediateType to create a map
func serializeImageIntermediateType(t *imageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageIntermediateType will accept a slice of imageIntermediateType to create a slice result
func serializeSliceImageIntermediateType(s []*imageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationIntermediateType will accept a map to create a durationIntermediateType
func deserializeDurationIntermediateType(in interface{}) (t *durationIntermediateType, err error) {
	tmp := &durationIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationIntermediateType will accept a slice to create a slice of durationIntermediateType
func deserializeSliceDurationIntermediateType(in []interface{}) (t []*durationIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationIntermediateType will accept a durationIntermediateType to create a map
func serializeDurationIntermediateType(t *durationIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationIntermediateType will accept a slice of durationIntermediateType to create a slice result
func serializeSliceDurationIntermediateType(s []*durationIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingIntermediateType will accept a map to create a followingIntermediateType
func deserializeFollowingIntermediateType(in interface{}) (t *followingIntermediateType, err error) {
	tmp := &followingIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingIntermediateType will accept a slice to create a slice of followingIntermediateType
func deserializeSliceFollowingIntermediateType(in []interface{}) (t []*followingIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingIntermediateType will accept a followingIntermediateType to create a map
func serializeFollowingIntermediateType(t *followingIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingIntermediateType will accept a slice of followingIntermediateType to create a slice result
func serializeSliceFollowingIntermediateType(s []*followingIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameIntermediateType will accept a map to create a preferredUsernameIntermediateType
func deserializePreferredUsernameIntermediateType(in interface{}) (t *preferredUsernameIntermediateType, err error) {
	tmp := &preferredUsernameIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameIntermediateType will accept a slice to create a slice of preferredUsernameIntermediateType
func deserializeSlicePreferredUsernameIntermediateType(in []interface{}) (t []*preferredUsernameIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameIntermediateType will accept a preferredUsernameIntermediateType to create a map
func serializePreferredUsernameIntermediateType(t *preferredUsernameIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameIntermediateType will accept a slice of preferredUsernameIntermediateType to create a slice result
func serializeSlicePreferredUsernameIntermediateType(s []*preferredUsernameIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeheightIntermediateType will accept a map to create a heightIntermediateType
func deserializeHeightIntermediateType(in interface{}) (t *heightIntermediateType, err error) {
	tmp := &heightIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice heightIntermediateType will accept a slice to create a slice of heightIntermediateType
func deserializeSliceHeightIntermediateType(in []interface{}) (t []*heightIntermediateType, err error) {
	for _, i := range in {
		tmp := &heightIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeheightIntermediateType will accept a heightIntermediateType to create a map
func serializeHeightIntermediateType(t *heightIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceheightIntermediateType will accept a slice of heightIntermediateType to create a slice result
func serializeSliceHeightIntermediateType(s []*heightIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToIntermediateType will accept a map to create a attributedToIntermediateType
func deserializeAttributedToIntermediateType(in interface{}) (t *attributedToIntermediateType, err error) {
	tmp := &attributedToIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToIntermediateType will accept a slice to create a slice of attributedToIntermediateType
func deserializeSliceAttributedToIntermediateType(in []interface{}) (t []*attributedToIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToIntermediateType will accept a attributedToIntermediateType to create a map
func serializeAttributedToIntermediateType(t *attributedToIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToIntermediateType will accept a slice of attributedToIntermediateType to create a slice result
func serializeSliceAttributedToIntermediateType(s []*attributedToIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceIntermediateType will accept a map to create a audienceIntermediateType
func deserializeAudienceIntermediateType(in interface{}) (t *audienceIntermediateType, err error) {
	tmp := &audienceIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceIntermediateType will accept a slice to create a slice of audienceIntermediateType
func deserializeSliceAudienceIntermediateType(in []interface{}) (t []*audienceIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceIntermediateType will accept a audienceIntermediateType to create a map
func serializeAudienceIntermediateType(t *audienceIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceIntermediateType will accept a slice of audienceIntermediateType to create a slice result
func serializeSliceAudienceIntermediateType(s []*audienceIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryIntermediateType will accept a map to create a summaryIntermediateType
func deserializeSummaryIntermediateType(in interface{}) (t *summaryIntermediateType, err error) {
	tmp := &summaryIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryIntermediateType will accept a slice to create a slice of summaryIntermediateType
func deserializeSliceSummaryIntermediateType(in []interface{}) (t []*summaryIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryIntermediateType will accept a summaryIntermediateType to create a map
func serializeSummaryIntermediateType(t *summaryIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryIntermediateType will accept a slice of summaryIntermediateType to create a slice result
func serializeSliceSummaryIntermediateType(s []*summaryIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceIntermediateType will accept a map to create a sourceIntermediateType
func deserializeSourceIntermediateType(in interface{}) (t *sourceIntermediateType, err error) {
	tmp := &sourceIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceIntermediateType will accept a slice to create a slice of sourceIntermediateType
func deserializeSliceSourceIntermediateType(in []interface{}) (t []*sourceIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceIntermediateType will accept a sourceIntermediateType to create a map
func serializeSourceIntermediateType(t *sourceIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceIntermediateType will accept a slice of sourceIntermediateType to create a slice result
func serializeSliceSourceIntermediateType(s []*sourceIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsIntermediateType will accept a map to create a endpointsIntermediateType
func deserializeEndpointsIntermediateType(in interface{}) (t *endpointsIntermediateType, err error) {
	tmp := &endpointsIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsIntermediateType will accept a slice to create a slice of endpointsIntermediateType
func deserializeSliceEndpointsIntermediateType(in []interface{}) (t []*endpointsIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsIntermediateType will accept a endpointsIntermediateType to create a map
func serializeEndpointsIntermediateType(t *endpointsIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsIntermediateType will accept a slice of endpointsIntermediateType to create a slice result
func serializeSliceEndpointsIntermediateType(s []*endpointsIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelastIntermediateType will accept a map to create a lastIntermediateType
func deserializeLastIntermediateType(in interface{}) (t *lastIntermediateType, err error) {
	tmp := &lastIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice lastIntermediateType will accept a slice to create a slice of lastIntermediateType
func deserializeSliceLastIntermediateType(in []interface{}) (t []*lastIntermediateType, err error) {
	for _, i := range in {
		tmp := &lastIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelastIntermediateType will accept a lastIntermediateType to create a map
func serializeLastIntermediateType(t *lastIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelastIntermediateType will accept a slice of lastIntermediateType to create a slice result
func serializeSliceLastIntermediateType(s []*lastIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedeletedIntermediateType will accept a map to create a deletedIntermediateType
func deserializeDeletedIntermediateType(in interface{}) (t *deletedIntermediateType, err error) {
	tmp := &deletedIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice deletedIntermediateType will accept a slice to create a slice of deletedIntermediateType
func deserializeSliceDeletedIntermediateType(in []interface{}) (t []*deletedIntermediateType, err error) {
	for _, i := range in {
		tmp := &deletedIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedeletedIntermediateType will accept a deletedIntermediateType to create a map
func serializeDeletedIntermediateType(t *deletedIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedeletedIntermediateType will accept a slice of deletedIntermediateType to create a slice result
func serializeSliceDeletedIntermediateType(s []*deletedIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}
