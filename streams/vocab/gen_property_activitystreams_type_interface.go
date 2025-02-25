package vocab

import "net/url"

// ActivityStreamsTypePropertyIterator represents a single value for the "type"
// property.
type ActivityStreamsTypePropertyIterator interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetXMLSchemaAnyURI returns the value of this property. When
	// IsXMLSchemaAnyURI returns false, GetXMLSchemaAnyURI will return an
	// arbitrary value.
	GetXMLSchemaAnyURI() *url.URL
	// GetXMLSchemaString returns the value of this property. When
	// IsXMLSchemaString returns false, GetXMLSchemaString will return an
	// arbitrary value.
	GetXMLSchemaString() string
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI".
	// When true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI
	// methods to access and set this property.
	IsXMLSchemaAnyURI() bool
	// IsXMLSchemaString returns true if this property has a type of "string".
	// When true, use the GetXMLSchemaString and SetXMLSchemaString
	// methods to access and set this property.
	IsXMLSchemaString() bool
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
	LessThan(o ActivityStreamsTypePropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsType".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsTypePropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsTypePropertyIterator
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetXMLSchemaAnyURI sets the value of this property. Calling
	// IsXMLSchemaAnyURI afterwards returns true.
	SetXMLSchemaAnyURI(v *url.URL)
	// SetXMLSchemaString sets the value of this property. Calling
	// IsXMLSchemaString afterwards returns true.
	SetXMLSchemaString(v string)
}

// Identifies the Object or Link type. Multiple values may be specified.
//
// Example 62 (https://www.w3.org/TR/activitystreams-vocabulary/#extype-jsonld):
//   {
//     "summary": "A foo",
//     "type": "http://example.org/Foo"
//   }
type ActivityStreamsTypeProperty interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "type"
	AppendIRI(v *url.URL)
	// AppendXMLSchemaAnyURI appends a anyURI value to the back of a list of
	// the property "type". Invalidates iterators that are traversing
	// using Prev.
	AppendXMLSchemaAnyURI(v *url.URL)
	// AppendXMLSchemaString appends a string value to the back of a list of
	// the property "type". Invalidates iterators that are traversing
	// using Prev.
	AppendXMLSchemaString(v string)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsTypePropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsTypePropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsTypePropertyIterator
	// Insert inserts an IRI value at the specified index for a property
	// "type". Existing elements at that index and higher are shifted back
	// once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// InsertXMLSchemaAnyURI inserts a anyURI value at the specified index for
	// a property "type". Existing elements at that index and higher are
	// shifted back once. Invalidates all iterators.
	InsertXMLSchemaAnyURI(idx int, v *url.URL)
	// InsertXMLSchemaString inserts a string value at the specified index for
	// a property "type". Existing elements at that index and higher are
	// shifted back once. Invalidates all iterators.
	InsertXMLSchemaString(idx int, v string)
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
	// Len returns the number of values that exist for the "type" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsTypeProperty) bool
	// Name returns the name of this property: "type".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "type".
	PrependIRI(v *url.URL)
	// PrependXMLSchemaAnyURI prepends a anyURI value to the front of a list
	// of the property "type". Invalidates all iterators.
	PrependXMLSchemaAnyURI(v *url.URL)
	// PrependXMLSchemaString prepends a string value to the front of a list
	// of the property "type". Invalidates all iterators.
	PrependXMLSchemaString(v string)
	// Remove deletes an element at the specified index from a list of the
	// property "type", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "type". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetXMLSchemaAnyURI sets a anyURI value to be at the specified index for
	// the property "type". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetXMLSchemaAnyURI(idx int, v *url.URL)
	// SetXMLSchemaString sets a string value to be at the specified index for
	// the property "type". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetXMLSchemaString(idx int, v string)
	// Swap swaps the location of values at two indices for the "type"
	// property.
	Swap(i, j int)
}
