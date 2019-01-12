package gen

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
)

func GenRootPackageComment(pkgName string) string {
	return codegen.FormatPackageDocumentation(fmt.Sprintf("Package %s "+
		"contains constructors and functions necessary for "+
		"applications to serialize, deserialize, and use "+
		"ActivityStreams types in Go. This package is code-generated "+
		"and subject to the same license as the go-fed tool used to "+
		"generate it.\n\n"+
		"This package is useful to three classes of developers: "+
		"end-user-application developers, specification writers "+
		"creating an ActivityStream Extension, and ActivityPub "+
		"implementors wanting to create an alternate ActivityStreams "+
		"implementation that still satisfies the interfaces generated "+
		"by the go-fed tool.\n\n"+
		"Application developers should limit their use to the "+
		"Resolver type, the constructors beginning with \"New\", the "+
		"\"Extends\" functions, the \"DisjointWith\" functions, the "+
		"\"ExtendedBy\" functions, and any interfaces returned in "+
		"those functions in this package. This lets applications use "+
		"Resolvers to Deserialize or Dispatch specific types. The "+
		"types themselves can Serialize as needed. The \"Extends\", "+
		"\"DisjointWith\", and \"ExtendedBy\" functions help navigate "+
		"the ActivityStreams hierarchy since it is not equivalent to "+
		"object-oriented inheritance.\n\n"+
		"When creating an ActivityStreams extension, developers will "+
		"want to ensure that the generated code builds correctly and "+
		"check that the properties, types, extensions, and "+
		"disjointedness is set up correctly. Writing unit tests with "+
		"concrete types is then the next step. If the tool has an "+
		"error generating this code, a fix is needed in the tool as "+
		"it is likely there is a new RDF type being used in the "+
		"extension that the tool does not know how to resolve. Thus, "+
		"most development will focus on the go-fed tool itself."+
		"\n\n"+
		"Finally, ActivityStreams implementors that want drop-in "+
		"replacement while still using the generated interfaces are "+
		"highly encouraged to examine the Manager type in this "+
		"package (in addition to the constructors) as these are the "+
		"locations where concrete types are instantiated. When "+
		"supplying a different type in these two locations, the "+
		"other generated code will propagate it throughout the "+
		"rest of an application.\n\n"+
		"Subdirectories of this package include implementation "+
		"files and functions that are not intended to be directly "+
		"linked to applications, but are used by this particular "+
		"package. It is strongly recommended to only use the "+
		"property interfaces and type interfaces in subdirectories "+
		"and limiting concrete types to those in this package. The "+
		"go-fed tool is likely to contain a pruning feature in the "+
		"future which will analyze an application and eliminate "+
		"code that would be dead if it were to be generated which "+
		"reduces the compilation time, compilation resources, and "+
		"binary size of an application. Such a feature will not be "+
		"compatible with applications that use the concrete "+
		"implementation types.",
		pkgName))
}
