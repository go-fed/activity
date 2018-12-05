package rdf

import (
	"fmt"
)

const (
	typeSpec                = "@type"
	typeActivityStreamsSpec = "type"
	IdSpec                  = "@id"
	IdActivityStreamsSpec   = "id"
	ContainerSpec           = "@container"
	IndexSpec               = "@index"
)

// jsonLDNodes contains the well-known set of nodes as defined by the JSON-LD
// specification.
func jsonLDNodes(r *RDFRegistry) []RDFNode {
	// Order matters -- we want to be able to distinguish the types of
	// things without other nodes hijacking the flow.
	return []RDFNode{
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     typeSpec,
			Delegate: &typeLD{r: r},
		},
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     typeActivityStreamsSpec,
			Delegate: &typeLD{r: r},
		},
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     IdSpec,
			Delegate: &idLD{},
		},
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     IdActivityStreamsSpec,
			Delegate: &idLD{},
		},
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     ContainerSpec,
			Delegate: &ContainerLD{},
		},
		&AliasedDelegate{
			Spec:     "",
			Alias:    "",
			Name:     IndexSpec,
			Delegate: &IndexLD{},
		},
	}
}

var _ RDFNode = &idLD{}

type idLD struct{}

func (i *idLD) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("id cannot be entered")
}

func (i *idLD) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("id cannot be exited")
}

func (i *idLD) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	if ctx.Current == nil {
		return true, nil
	} else if ider, ok := ctx.Current.(URISetter); !ok {
		return true, fmt.Errorf("id apply called with non-URISetter")
	} else if str, ok := value.(string); !ok {
		return true, fmt.Errorf("id apply called with non-string value")
	} else {
		return true, ider.SetURI(str)
	}
}

var _ RDFNode = &typeLD{}

type typeLD struct {
	r *RDFRegistry
}

func (t *typeLD) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, nil
}

func (t *typeLD) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, nil
}

func (t *typeLD) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	fmt.Printf("@type key=%s, value=%v\n", key, value)
	vs, ok := value.(string)
	if !ok {
		return true, fmt.Errorf("@type is not string")
	}
	n, e := t.r.getNode(vs)
	if e != nil {
		return true, e
	}
	return n.Apply(vs, nil, ctx)
}

var _ RDFNode = &ContainerLD{}

type ContainerLD struct {
	ContainsNode RDFNode
}

func (c *ContainerLD) Enter(key string, ctx *ParsingContext) (bool, error) {
	fmt.Println("===container enter")
	if ctx.OnlyApplyThisNodeNextLevel != nil {
		return true, fmt.Errorf("@container parsing context exit already has non-nil node")
	}
	ctx.SetOnlyApplyThisNodeNextLevel(c.ContainsNode)
	return true, nil
}

func (c *ContainerLD) Exit(key string, ctx *ParsingContext) (bool, error) {
	fmt.Println("===container exit")
	if ctx.OnlyApplyThisNodeNextLevel == nil {
		return true, fmt.Errorf("@container parsing context exit already has nil node")
	}
	ctx.ResetOnlyAppliedThisNodeNextLevel()
	return true, nil
}

func (c *ContainerLD) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	return true, nil
}

var _ RDFNode = &IndexLD{}

type IndexLD struct{}

func (i *IndexLD) Enter(key string, ctx *ParsingContext) (bool, error) {
	fmt.Println(">>> enter index")
	return true, nil
}

func (i *IndexLD) Exit(key string, ctx *ParsingContext) (bool, error) {
	fmt.Println(">>> exit index")
	return true, nil
}

func (i *IndexLD) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	return true, nil
}
