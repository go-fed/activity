package gen

import (
	"fmt"
	"github.com/go-fed/activity/tools/defs"
	"strings"
)

func validateDomains(pt []*defs.PropertyType) error {
	var invalid []string
	for _, p := range pt {
		for _, d := range p.Domain {
			if d.T == nil {
				continue
				found := false
				for _, inv := range d.T.GetProperties() {
					if inv.Name == p.Name {
						found = true
						break
					}
				}
				if !found {
					invalid = append(invalid, fmt.Sprintf("Property %s has domain %s, but %s does not have property %s", p.Name, d.T.Name, d.T.Name, p.Name))
				}
			}
		}
	}
	if len(invalid) == 0 {
		return nil
	}
	return fmt.Errorf("ValidateDomains failed:\n\t%s", strings.Join(invalid, "\n\t"))
}

func validateProperties(ct []*defs.Type) error {
	var invalid []string
	for _, t := range ct {
		names := t.AllExtendsNames()
		for _, p := range t.GetProperties() {
			found := false
			for _, d := range p.Domain {
				if d.T == nil {
					continue
				}
				for _, n := range names {
					if d.T.Name == n {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				invalid = append(invalid, fmt.Sprintf("Type %s has property %s, but property %s does not have domain %s", t.Name, p.Name, p.Name, t.Name))
			}
		}
	}
	if len(invalid) == 0 {
		return nil
	}
	return fmt.Errorf("ValidateProperties failed:\n\t%s", strings.Join(invalid, "\n\t"))
}

func validateValues(vt []*defs.ValueType, pt []*defs.PropertyType) error {
	var invalid []string
	for _, v := range vt {
		if defs.IsIRIValueType(v) {
			continue
		}
		found := false
		for _, p := range pt {
			for _, r := range p.Range {
				if r.V != nil {
					if r.V.Name == v.Name {
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		if !found {
			invalid = append(invalid, fmt.Sprintf("Value %s is not referred to by any property", v.Name))
		}
	}
	if len(invalid) == 0 {
		return nil
	}
	return fmt.Errorf("ValidateValues failed:\n\t%s", strings.Join(invalid, "\n\t"))
}
