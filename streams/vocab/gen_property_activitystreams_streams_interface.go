package vocab

import "net/url"

// ActivityStreamsStreamsPropertyIterator represents a single value for the
// "streams" property.
type ActivityStreamsStreamsPropertyIterator interface {
	// GetActivityStreamsCollection returns the value of this property. When
	// IsActivityStreamsCollection returns false,
	// GetActivityStreamsCollection will return an arbitrary value.
	GetActivityStreamsCollection() ActivityStreamsCollection
	// GetActivityStreamsCollectionPage returns the value of this property.
	// When IsActivityStreamsCollectionPage returns false,
	// GetActivityStreamsCollectionPage will return an arbitrary value.
	GetActivityStreamsCollectionPage() ActivityStreamsCollectionPage
	// GetActivityStreamsOrderedCollection returns the value of this property.
	// When IsActivityStreamsOrderedCollection returns false,
	// GetActivityStreamsOrderedCollection will return an arbitrary value.
	GetActivityStreamsOrderedCollection() ActivityStreamsOrderedCollection
	// GetActivityStreamsOrderedCollectionPage returns the value of this
	// property. When IsActivityStreamsOrderedCollectionPage returns
	// false, GetActivityStreamsOrderedCollectionPage will return an
	// arbitrary value.
	GetActivityStreamsOrderedCollectionPage() ActivityStreamsOrderedCollectionPage
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsCollection returns true if this property has a type of
	// "Collection". When true, use the GetActivityStreamsCollection and
	// SetActivityStreamsCollection methods to access and set this
	// property.
	IsActivityStreamsCollection() bool
	// IsActivityStreamsCollectionPage returns true if this property has a
	// type of "CollectionPage". When true, use the
	// GetActivityStreamsCollectionPage and
	// SetActivityStreamsCollectionPage methods to access and set this
	// property.
	IsActivityStreamsCollectionPage() bool
	// IsActivityStreamsOrderedCollection returns true if this property has a
	// type of "OrderedCollection". When true, use the
	// GetActivityStreamsOrderedCollection and
	// SetActivityStreamsOrderedCollection methods to access and set this
	// property.
	IsActivityStreamsOrderedCollection() bool
	// IsActivityStreamsOrderedCollectionPage returns true if this property
	// has a type of "OrderedCollectionPage". When true, use the
	// GetActivityStreamsOrderedCollectionPage and
	// SetActivityStreamsOrderedCollectionPage methods to access and set
	// this property.
	IsActivityStreamsOrderedCollectionPage() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsStreamsPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsStreams".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsStreamsPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsStreamsPropertyIterator
	// SetActivityStreamsCollection sets the value of this property. Calling
	// IsActivityStreamsCollection afterwards returns true.
	SetActivityStreamsCollection(v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets the value of this property.
	// Calling IsActivityStreamsCollectionPage afterwards returns true.
	SetActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// SetActivityStreamsOrderedCollection sets the value of this property.
	// Calling IsActivityStreamsOrderedCollection afterwards returns true.
	SetActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets the value of this
	// property. Calling IsActivityStreamsOrderedCollectionPage afterwards
	// returns true.
	SetActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}

// A list of supplementary Collections which may be of interest
type ActivityStreamsStreamsProperty interface {
	// AppendActivityStreamsCollection appends a Collection value to the back
	// of a list of the property "streams". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsCollection(v ActivityStreamsCollection)
	// AppendActivityStreamsCollectionPage appends a CollectionPage value to
	// the back of a list of the property "streams". Invalidates iterators
	// that are traversing using Prev.
	AppendActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// AppendActivityStreamsOrderedCollection appends a OrderedCollection
	// value to the back of a list of the property "streams". Invalidates
	// iterators that are traversing using Prev.
	AppendActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// AppendActivityStreamsOrderedCollectionPage appends a
	// OrderedCollectionPage value to the back of a list of the property
	// "streams". Invalidates iterators that are traversing using Prev.
	AppendActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "streams"
	AppendIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "streams". Invalidates iterators that are traversing
	// using Prev. Returns an error if the type is not a valid one to set
	// for this property.
	AppendType(t Type) error
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsStreamsPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsStreamsPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsStreamsPropertyIterator
	// InsertActivityStreamsCollection inserts a Collection value at the
	// specified index for a property "streams". Existing elements at that
	// index and higher are shifted back once. Invalidates all iterators.
	InsertActivityStreamsCollection(idx int, v ActivityStreamsCollection)
	// InsertActivityStreamsCollectionPage inserts a CollectionPage value at
	// the specified index for a property "streams". Existing elements at
	// that index and higher are shifted back once. Invalidates all
	// iterators.
	InsertActivityStreamsCollectionPage(idx int, v ActivityStreamsCollectionPage)
	// InsertActivityStreamsOrderedCollection inserts a OrderedCollection
	// value at the specified index for a property "streams". Existing
	// elements at that index and higher are shifted back once.
	// Invalidates all iterators.
	InsertActivityStreamsOrderedCollection(idx int, v ActivityStreamsOrderedCollection)
	// InsertActivityStreamsOrderedCollectionPage inserts a
	// OrderedCollectionPage value at the specified index for a property
	// "streams". Existing elements at that index and higher are shifted
	// back once. Invalidates all iterators.
	InsertActivityStreamsOrderedCollectionPage(idx int, v ActivityStreamsOrderedCollectionPage)
	// Insert inserts an IRI value at the specified index for a property
	// "streams". Existing elements at that index and higher are shifted
	// back once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "streams". Invalidates all iterators. Returns an error
	// if the type is not a valid one to set for this property.
	InsertType(idx int, t Type) error
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "streams" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsStreamsProperty) bool
	// Name returns the name of this property: "streams".
	Name() string
	// PrependActivityStreamsCollection prepends a Collection value to the
	// front of a list of the property "streams". Invalidates all
	// iterators.
	PrependActivityStreamsCollection(v ActivityStreamsCollection)
	// PrependActivityStreamsCollectionPage prepends a CollectionPage value to
	// the front of a list of the property "streams". Invalidates all
	// iterators.
	PrependActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// PrependActivityStreamsOrderedCollection prepends a OrderedCollection
	// value to the front of a list of the property "streams". Invalidates
	// all iterators.
	PrependActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// PrependActivityStreamsOrderedCollectionPage prepends a
	// OrderedCollectionPage value to the front of a list of the property
	// "streams". Invalidates all iterators.
	PrependActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "streams".
	PrependIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "streams". Invalidates all iterators. Returns an error
	// if the type is not a valid one to set for this property.
	PrependType(t Type) error
	// Remove deletes an element at the specified index from a list of the
	// property "streams", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsCollection sets a Collection value to be at the
	// specified index for the property "streams". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsCollection(idx int, v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets a CollectionPage value to be at
	// the specified index for the property "streams". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetActivityStreamsCollectionPage(idx int, v ActivityStreamsCollectionPage)
	// SetActivityStreamsOrderedCollection sets a OrderedCollection value to
	// be at the specified index for the property "streams". Panics if the
	// index is out of bounds. Invalidates all iterators.
	SetActivityStreamsOrderedCollection(idx int, v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets a OrderedCollectionPage
	// value to be at the specified index for the property "streams".
	// Panics if the index is out of bounds. Invalidates all iterators.
	SetActivityStreamsOrderedCollectionPage(idx int, v ActivityStreamsOrderedCollectionPage)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "streams". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetType sets an arbitrary type value to the specified index of the
	// property "streams". Invalidates all iterators. Returns an error if
	// the type is not a valid one to set for this property. Panics if the
	// index is out of bounds.
	SetType(idx int, t Type) error
	// Swap swaps the location of values at two indices for the "streams"
	// property.
	Swap(i, j int)
}
