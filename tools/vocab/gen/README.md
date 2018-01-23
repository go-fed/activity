# gen

This library is used to generate the `vocab` library implementation. Changes to
this will fundamentally change the implementation of the `vocab` library. If
custom ActivityStream types are required, no changes to this are required.
Instead, the `tools/gen/vocab` library will need edits.

Creating a static type implementation from the specification is not
straightforward due to the following considerations that stem from its roots in
JSON-LD:

* Any property could be an IRI (`url.URL`)
* Most properties could be another ActivityStream type
  (`map[string]interface{}`) or a value (`time.Time`, etc)
* Most properties can have multiple values (`[]interface{}`)
* In a multiple value scenario, each value could once more be an IRI, another
  ActivityStream type, or a value

The current generation algorithm leaves ample opportunity for improvements and
optimizations.
