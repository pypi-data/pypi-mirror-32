# Persistence

There are three kinds of objects that are considered on the system:


## Tree objects

Called Resources that implement `guillotina.interfaces.IResource`. This object
has a `__name__` and a `__parent__` field that indicate the id on the tree and
the link to the parent. By themselves they don't have access to their children,
they need to interact with the transaction object to get them.


## Nested

Objects that are linked at some attribute inside the Tree object, this object
are serialized with the main object and may lead to conflicts if there are lots
of this kind of objects.

It can belong to:

- A field that is an object


## Nested References

Base objects that belong to a specific object, it is big enough to have its own
entity and be saved in a different persistence object. Its not an element of the tree.

It can belong to:

- An annotation that is stored on a different object
- A BaseObject inherited field object
