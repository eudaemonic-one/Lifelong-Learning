# Lecture 1 Intro

## What is a Network

* An infrastructure that allows distributed users to communicate with each other
* Building Block: node -> link -> node
* Scaling the network
* Message Switching Architecture
  * Store-and-Forward operation
    * Messages were decoded
    * Next-hop in route determined by destination address of a message
    * Each message was carried by hand to next line
* Circuit Switching
  * Source first establishes a connection (circuit) to destination
  * Source sends the data over the circuit
  * Predictable performance (Fast and simple data transfer)
  * Circuit will be idle for significant periods of time
* Packet Switching
  * Source sends information as self-contained messages that have an address
  * Each packet transfers independently to destination

## What is the Internet

* Inter-net: network of networks
* Enable communication between diverse applications on diverse devices
* Over very diverse infrastructures: WIFI, cellular, data center networks, corporate networks
* Application needs/demands
  * Traffic data rate and loss sensitivity
  * Traffic pattern
  * Traffic target
* Multiplexing
  * Need to share network resources

## Internet Design

* In order to inter-operate, all participating networks must follow a common set of rules
* Service Model: the commitment made to applications
* Standards allow players to share risk and benefits of a new market
  *e.g., 802.11 LAN, UP, HTTP/SMTP
