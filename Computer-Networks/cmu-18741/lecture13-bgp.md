# Lecture 13 Border Gateway Protocol

## Routing Hierarchies

* Flat routing doesn't scale
* Solution: hierarchy
  * Inter and intra domain routing in the Internet
  * "Areas" inside OSPF
  * Switched LAN technologies

## Border Gateway Protocol (BGP)

### Autonomous System

* Autonomous System
  * A set of routers under a single technical administration, using an interior gateway protocol (IGP) and common metrics to route packets within the AS and using an exterior gateway protocol (EGP) to route packets to other AS's
  * Assigned unique ID (16-bit)
* Valley-free routing
  * Number linnks as (+1, 0, -1) for provider, peer, and customer
* Result: BGP introduced as routing protocol
  * Link state or distance vector?
* Solution: distance vector with path
  * Each routing update carries the entire path
  * Loops are detected as follows:
    * When AS get route, check if AS already in path

### Policy-based Routing

* Inter-connecting BGP peers
  * BGP uses TCP to connect peers
* Hop-by-hop Model
  * BGP advertises to neighbors only those routes that it uses
  * BGP enforces policies by
    * Choosing paths from multiple alternatives
    * Controlling advertisement to other AS's
* Examples of BGP Policies
  * A multi-homed AS refuses to act as transit
    * Limit path advertisement
  * A multi-homed AS can become transit for some AS's
    * Only advertise paths to some ASs
  * An AS can favor or disfavor certain AS's for traffic transit from itself
    * By choosing those paths among the options

### BGP Messages

* Open
  * Announces AS ID
  * Determines hold timer - interval between keep_alive or update messages, zero interval implies no keep_alive
* Keep-Alive
  * Sent periodically to peers to ensure connectivity
  * Send in place of an UPDATE message
* Notification
  * Used for error notification
  * TCP connection is closed immediately after notification
* UPDATE
  * List of withdrawn routes
  * Network layer reachability information (List of reachable prefixes)
  * Path attributes (Origin, Path, Metrics)
  * All prefixes advertised in message have same path attributed
* LOCAL-PREF
  * Local (within an AS) mechanism to provide relative priority among BGP routers
  * Prefer to use peering connection rather than transit
  * In general, customer > peer > provider
* AS-PATH
  * List of traversed AS's
* Multi-Exit Discriminator (MED)
  * Hint to external neighbors about the preferred path into an AS
  * Used when two AS's connect to each other in more than one place
  * MED is typically used in provider/subscriber scenarios

### Path Selection Criteria

* Attribute + external (policy) information
* Rough ordering for path selection
  * Highest LOCAL-PREF
  * Shortest AS-PATH
  * Lowest origin type
  * Lowest MED (if routes learned from same neighbor)
  * eBGP over iBGP-learned
  * Lowest internal routing cost to border router
  * Tie breaker, e.g., lowest router ID

### ISPs Peer

* Public peering: use network to connect large number of ISPs in Inernet eXchange Point (IXP)
* Private peering: directly connect ISP border router
  * Set up as private connection

### BGP Summary

* Wide area Internet structure and routing driven by economic considerations
  * Customer, provider, and peer
* BGP designed to
  * Provide hierarchy that allows scalability
  * Allow enforcement of policies related to structure
* Mechanisms
  * Path vector - scalable, hides structure from neighbors, detects lookp quickly
  * eBGP versus iBGP