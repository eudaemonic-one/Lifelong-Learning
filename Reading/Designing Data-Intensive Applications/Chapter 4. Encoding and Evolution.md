# Chapter 4. Encoding and Evolution

The old and new versions of the code, and old and new data formats, may potentially all coexist in the system at the same time.

Backward compatibility: newer code can read data that was written by older code.

Forward compatibility: Older code can read data that was written by newer code.

## Formats for Encoding Data

when writing data to a file or sending it over the network, data is encoded as some kind of self-contained sequence of bytes.

### Language-Specific Formats

Many programming languages come with built-in support for encoding in-memory objects into byte sequences.

* Tied to a particular language and hard to read in another language.
* The decoding process needs to be able to instantiate arbitrary classes.
* Efficiency is often an afterthought.

### JSON, XML, and Binary Variants

JSON, XML, CSV are textual formats and human-readable.

* Ambiguity around encoding of numbers.
* No support for binary strings, except for Base64-encoded hacky support.
* Depend on hardcoded encoding and decoding logic in application code.

For data that is used only internally within your organization, there is less pressure to use a lowest-common-denominator encoding format.

Binary encodings for JSON: MessagePack, BSON, BJSON, etc.

### Thrift and Protocol Buffers

Thrift: two binary encoding formats, i.e., *BinaryProtocol* and *CompactProtocol* (packing the field type and tag number into single byte, and by using variable-length integers).

```text
struct Person {
  1: required string       userName,
  2: optional i64          favoriteNumber,
  3: optional list<string> interests
}
```

Protocol Buffers: binary encoding library with bit packing. `required` and `optional` only enable a runtime check but makes no difference to how the field is encoded. It doesn't have a lst or array datatype, but instead has a `repeated` marker.

```text
message Person {
    required string user_name       = 1;
    optional int64  favorite_number = 2;
    repeated string interests       = 3;
}
```

*schema evolution*:

* You cannot change a field's tag, sine that would make all existing encoded data invalid.
* If old code tries to read data written by new ode, it can simply ignore the new field with a tag number it doesn't recognize.
* New code can always read old data, because the tag numbers still have the same meaning.
* New fields could either be optional or have a default value.
* Removing a field is just like adding a field. You can never remove a field that is required.

*datatypes and schema evolution* => there is a risk that values will lose precision or get truncated; maybe unsupported by encoding formats.

### Avro

Apache Avro: binary encoding format for Hadoop use case; no tag numbers; compact encoding; any mismatch in the schema would mean incorrectly decoded data.

```text
record Person {
    string               userName;
    union { null, long } favoriteNumber = null;
    array<string>        interests;
}
```

writer's schema vs. reader's schema => don't have to be the same, but need to be compatible.

With Arvo, to maintain compatibility, you may only add or remove a field that has a default value.

Avro supports dynamically generated schemas and can accommodate schema changes without reshuffling field tags.

### The Merits of Schemas

Binary encodings have nice properties:

* Much more compact.
* A value form of documentation.
* Keeping a database of schemas allow you to check forward and backward compatibility of schema changes, before anything is deployed.

## Modes of Dataflow

### Dataflow Through Databases

>> Storing something in the database is simply sending a message to your future self.

*data outlives code*: old data in original encoding will exist there unless you have explicitly rewritten it.

*archival storage*: data dump for backup purpose or for loading into a data warehouse is written in one go and is thereafter immutable.

### Dataflow Through Services: REST and RPC

REST is a design philosophy that emphasizes simple data formats, using URLs for identifying resources and using HTTP features for cache control, authentication, and content type negotiation.

SOAP is an XML-based protocol for making network API requests; it aims to be independent from HTTP; it comes with multitude of related standards (WS-*).

RPC (*remote procedure call*) makes a request to a remote network service look the same as calling a function or a method.

* A local function call is predictable and either succeeds or fails. A network request is unpredictable; the request or response may be lost or the remote machine may be slow or unavailable.
* A local function call either returns a result, or throws an exception, or never returns. A network request has another possible outcome: it may return without a result, due to a timeout.
* Retry a failed network request would require a robust mechanism for idempotency.
* A network request is much slower than a function call, and its latency is also wildly variable.
* When you call a local function, you can efficiently pass it references (pointers) to objects in local memory. When you make a network request, all those parameters need to be encoded into a sequence of bytes that can be sent over the network.
* The client and the service may be implemented in different programming languages, so the RPC framework must translate datatypes from one language into another.

Current directions for RPC:

* Use *futures* (*promises*) to encapsulate asynchronous actions that may fail.
* Support *streams*, where a call may consist of a series of requests and responses over time.
* Provide *service discovery* to allow clients to find out at which IP address and port number it can find a particular service.
* The main focus of RPC frameworks is on requests between services owned by the same organization, typically within the same datacenter.

Service compatibility is made harder by the fact that RPC is often used for communication across organizational boundaries, so the provider of a service often has no control over its clients and cannot force them to upgrade. Thus, compatibility needs to be maintained for a long time, perhaps indefinitely.

### Message-Passing Dataflow

*asynchronous message-passing system* (*message broker* or *message queue* or *message-oriented middleware*)

* It can act as a buffer if the recipient is unavailable or overloaded, and thus improve system reliability.
* It can automatically redeliver messages to a process that has crashed, and thus prevent messages from being lost.
* It avoids the sender needing to know the IP address and port number of the recipient.
* It allows one message to be sent to several recipients.
* It logically decouples the sender from the recipient.
* The message-passing communication is usually one-way; it's possible for a process to send a response, but this would be done on a separate channel.

*actor model*: encapsulate logic in actors; each actor represents one client or entity, and it communicates with other actors by sending and receiving asynchronous messages; message delivery is not guaranteed.

*distributed actor framework*: actors are allocated across multiple nodes; messages are transparently encoded into byte sequences, sent over the network, and decoded on the other side.

*location transparency* works better in the actor model than in RPC, because the actor model already assumes that messages may be lost.
