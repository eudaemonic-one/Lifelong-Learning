# Chapter 2 Application Layer

## Principles of Network Applications

* At the core of network application development is writing programs that run on different end systems and communicate with each other over the network

### Network Application Architectures

* In a **client-server architecture**, there is an always-on host, called the *server*, which services requests from many other hosts, called *clients*
  * The server has a fixed, well-known address, called an IP address
  * A **data center**, housing a large number of hosts, is often used to create a powerful virtual server
* In a **P2P architecture**, there is minimal (or no) reliance on dedicated servers in data centers
  * The application exploits direct communication between pairs of intermittently connected hosts, called *peers*
  * **self-scalability**
  * Challenges
    * Most residential ISPs have been dimensioned for “asymmetrical” bandwidth usage, that is, for much more downstream than upstream traffic
    * Because of their highly distributed and open nature, P2P applications can be a challenge to secure
    * The success of future P2P applications also depends on convincing users to volunteer bandwidth, storage, and computation resources to the appli- cations

### Processes Communicating

* Processes on two different end systems communicate with each other by exchang- ing **messages** across the computer network
* A sending process creates and sends mes- sages into the network; a receiving process receives these messages and possibly responds by sending messages back
* **Client and Server Processes**
  * In the context of a communication session between a pair of processes, the process that initiates the communication (that is, initially contacts the other process at the beginning of the session) is labeled as the **client**
  * The process that waits to be contacted to begin the session is the **server**
* **The Interface Between the Process and the Computer Network**
  * A process sends messages into, and receives messages from, the network through a software interface called a **socket**
  * A socket is the interface (API) between the application layer and the transport layer within a host
  * The only control that the application developer has on the transport-layer side is
    * the choice of transport protocol
    * perhaps the ability to fix a few transport-layer parameters such as maximum buffer and maximum segment sizes
* **Addressing Processes**
  * The host is identified by its **IP address**
  * A destination **port number** serves the purpose of identifying the receiving process

### Transport Services Available to Applications

* The application at the sending side pushes messages through the socket
* At the other side of the socket, the transport-layer protocol has the responsibility of getting the messages to the socket of the receiving process
* **Reliable Data Transfer**
  * When a transport-layer protocol doesn’t provide reliable data transfer, some of the data sent by the sending process may never arrive at the receiving process
  * This may be acceptable for **loss-tolerant applications**
* **Throughput**
  * Applications that have throughput requirements are said to be **bandwidth-sensitive applications**
  * While bandwidth-sensitive applications have specific throughput requirements, **elastic applications** can make use of as much, or as little, throughput as happens to be available
* **Timing**
  * A transport-layer protocol can also provide timing guarantees
* **Security**
  * Finally, a transport protocol can provide an application with one or more security services

### Transport Services Provided by the Internet

* The Internet (and, more generally, TCP/IP networks) makes two transport protocols available to applications, UDP and TCP

| Application                           | Data Loss     | Throughput                                 | Time-Sensitive    |
| ------------------------------------- | ------------- | ------------------------------------------ | ----------------- |
| File transfer/download                | No loss       | Elastic                                    | No                |
| E-mail                                | No loss       | Elastic                                    | No                |
| Web documents                         | No loss       | Elastic                                    | No                |
| Internet telephony/Video conferencing | Loss-tolerant | Audio: few kbps-1Mbps; Video: 10kbps-5Mbps | Yes: 100s of msec |
| Streaming stored audio/vide           | Loss-tolerant | Same as above                              | Yes: few seconds  |
| Interactive games                     | Loss-tolerant | Few kbps-10kbps                            | Yes: 100s of msec |
| Instant messaging                     | No loss       | Elastic                                    | Yes and no        |

* **TCP Services**
  * Includes a connection-oriented service and a reliable data transfer service
  * After the handshaking phase, a **TCP connection** is said to exist between the sockets of the two processes
  * When one side of the application passes a stream of bytes into a socket, it can count on TCP to deliver the same stream of bytes to the receiving socket, with no missing or duplicate bytes
  * Also includes a congestion-control mechanism
* **UDP Services**
  * UDP is a no-frills, lightweight transport protocol, providing minimal services
  * UDP is connectionless, so there is no handshaking before the two processes start to communicate
  * UDP provides an unreliable data transfer service
  * Furthermore, messages that do arrive at the receiving process may arrive out of order
  * UDP does not include a congestion-control mechanism

### Application-Layer Protocols

* An **application-layer protocol** defines how an application’s processes, running on different end systems, pass messages to each other
  * The types of messages exchanged, for example, request messages and response messages
  * The syntax of the various message types, such as the fields in the message and how the fields are delineated
  * The semantics of the fields, that is, the meaning of the information in the fields
  * Rules for determining when and how a process sends messages and responds to messages