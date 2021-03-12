# Lecture 11 IP 3

## Network Address Translation (NAT)

* Within organization: assign every host a private IP address
  * Route within organization by IP protocol, can do subnetting
* NAT translates between public and private IP addresses
* Firewall acts as proxy for client and relabels destination to local address
* Use port mapping to make servers available
* NAT has to be consistent during a session
* NAT only works for certain applications
  * Some applications (e.g., ftp) pass IP information in payload
  * Need application level gateways to do a matching translation
  * Peer-peer, multi-player games have problems - who is server?
* NATs help with security
  * Often combined with firewalls
  * Most internal hosts are not accessible from the public Internet

## Tunnels

### Tunneling

* Motivation
  * e.g., IP multicast
  * e.g., incremental deployment of IPv6
  * e.g., must have some address to use services
* Force a packet to go to a specific point in the network
* Achieved by adding an extra IP header to the packet with a new destination address
* Used increasingly to deal with special routing requirements or new features
  * e.g., Mobile IP, Multicast, IPv6

### IP-in-IP Tunneling

* IP source and destination address identify tunnel endpoints
* Several fields are copies of the inner-IP header
* Inner header is not modified, except for decrementing TTL

### Tunneling Applications

* Virtual private networks
  * Connect subnets of a corporation using IP tunnels
  * Often combined with IP Sec
  * Overlays private network on top of regular Internet
* Support for new or unusual protocols
  * Routers that support the protocols use tunnels to bypass routers that do not support it
  * e.g., multicast, IPv6
* Force packets to follow non-standard routes
  * Routing is based on outer-header
  * e.g., mobile IP

## IPv6

### IPv6 Addressing

* Most urgent issue: increasing address space
  * 128-bit addresses
* Simplified header for faster processing
  * No checksum
  * No fragmentation
* 128-bit addresses provide space for structure

### IPv6 Autoconfiguration

* Stateless and no manual config at all
* Link-local address
  * Uniqueness test ("anyone using this address?")
* DHCP takes some of the wind out of this

### Fast Path vs. Slow Path

* Common case: switched in silicon (fast path)
* Weird cases: handed to CPU (slow path)
  * Fragmentation
  * TTL expiration (traceroute)
  * IP option handling

### IPv6 Header Cleanup

* 32 IPv4 options -> variable length header
  * Rarely used and many routers do not support
  * Processed in slow path
* IPv6 options: "Next header" pointer
  * Combines "protocol" and "options" handling
  * Extension header
  * Makes it easy to implement host-based options
* No checksum
  * Motivation was efficiency
  * Useful when corruption frequent, bandwidth is expensive
* No fragmentation
  * Router discard packets, send ICMP "Packet Too Big"
  * Reduced packet processing and network complexity
  * Increased MTU a boon to application writers
  * Hosts can still fragment - using fragmentation header, but routers don't deal with it any more

### Migration from IPv4 to IPv6

* Interoperability with IPv4 is necessary for incremental deployment
* Combination of mechanisms:
  * Dual stack operation: IPv6 nodes support both address types
  * Tunnel IPv6 packets trhough IPv4 clouds
  * IPv4-IPv6 translation at edge of network
    * NAT should also translate between IPv4 and IPv6 protocols