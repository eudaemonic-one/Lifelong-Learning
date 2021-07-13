# Chapter 12. Serialization

* “THIS chapter concerns *object serialization*, which is Java’s framework for encoding objects as byte streams (*serializing*) and reconstructing objects from their encodings (*deserializing*). Once an object has been serialized, its encoding can be sent from one VM to another or stored on disk for later deserialization. This chapter focuses on the dangers of serialization and how to minimize them.”


## Item 85: Prefer alternatives to Java serialization

* “A fundamental problem with serialization is that its *attack surface* is too big to protect, and constantly growing: Object graphs are deserialized by invoking the `readObject` method on an `ObjectInputStream`. This method is essentially a magic constructor that can be made to instantiate objects of almost any type on the class path, so long as the type implements the `Serializable` interface. In the process of deserializing a byte stream, this method can execute code from any of these types, so the code for all of these types is part of the attack surface.”
* “Without using any gadgets, you can easily mount a denial-of-service attack by causing the deserialization of a short stream that requires a long time to deserialize. Such streams are known as *deserialization bombs* [Svoboda16].”
* **“The best way to avoid serialization exploits is never to deserialize anything.”**
* **“There is no reason to use Java serialization in any new system you write.”**
  * “If you can’t avoid Java serialization entirely, perhaps because you’re working in the context of a legacy system that requires it, your next best alternative is to **never deserialize untrusted data**.”

* “The leading cross-platform structured data representations are JSON [JSON] and Protocol Buffers, also known as protobuf [Protobuf].”
  * “The most significant differences between JSON and protobuf are that JSON is text-based and human-readable, whereas protobuf is binary and substantially more efficient; and that JSON is exclusively a data representation, whereas protobuf offers *schemas* (types) to document and enforce appropriate usage. ”
* “If you can’t avoid serialization and you aren’t absolutely certain of the safety of the data you’re deserializing, use the object deserialization filtering added in Java 9 and backported to earlier releases (`java.io.ObjectInputFilter`). ”
  * “This facility lets you specify a filter that is applied to data streams before they’re deserialized. It operates at the class granularity, letting you accept or reject certain classes.”
  * “Accepting classes by default and rejecting a list of potentially dangerous ones is known as *blacklisting*; rejecting classes by default and accepting a list of those that are presumed safe is known as *whitelisting*.”
  * “**Prefer whitelisting to blacklisting**, as blacklisting only protects you against known threats.”
* **“In summary, serialization is dangerous and should be avoided. If you are designing a system from scratch, use a cross-platform structured-data representation such as JSON or protobuf instead. Do not deserialize untrusted data. If you must do so, use object deserialization filtering, but be aware that it is not guaranteed to thwart all attacks. Avoid writing serializable classes. If you must do so, exercise great caution.”**