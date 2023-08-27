# Chapter 1. Reliable, Scalable, and Maintainable Applications

## Overview

Standard building blocks: databases, caches, search indexes, stream processing, batch processing.

## Thinking About Data Systems

* Reliability: The system should continue to work correctly (performing the correct function at the desired level of performance) even in the face of adversity (hardware or software faults, and even human error).
* Scalability: As the system grows (in data volume, traffic volume, or complexity), there should be reasonable ways of dealing with that growth.
* Maintainability: Over time, many different people will work on the system (engineering and operations, both maintaining current behavior and adapting the systems to new use cases), and they should all be able to work on it productively.

## Reliability

faults != failure?

fault == system deviating from its spec

failure == system as a whole stops providing the required service

fault-tolerant vs. prevent faults?

### Hardware Faults

Hardware Faults: hard disks crash, faulty RAM, power grid blackout, unplug wrong network cable

add redundancy -> more fault-tolerant but not preventing hardware problems from causing failures

sometimes high availability is not absolutely essential; while platforms are designed to prioritize flexibility and elasticity over single-machine reliability

systems could tolerate the loss of individual machines

### Software Errors

Software Errors: software bug, runaway process due to CPU, memory, disk space, or network bandwidth draining, unresponsive dependent systems, cascading failures

### Human Errors

> Humans are known to be unreliable

approaches to tolerate human errors:

* well-designed abstractions, APIs, admin interfaces
* provide fully featured non-production sandbox environments where people can explore and experiment safely, using real data, without affecting real users
* test thoroughly at all levels
* allow quick and easy recovery from human errors, like fast to rollback but roll out new change gradually
* set up detailed and clear monitoring

## Scalability

### Describing Load

load parameters: requests per second, ratio of reads to writes, simultaneous active users, hit rate on a cache, etc.

daily average vs. bottleneck by a small number of extreme?

read-intensive vs. write-intensive?

### Describing Performance

Latency and/vs response time

* arithmetic mean vs. percentiles (p95, p99, p999)
* tail latencies == high percentiles of response times are affect ing user's experience directly
* queueing delay == head-of-line blocking == slow requests holding up the processing of subsequent requests
* artificial load generator -> ordering, shorter queue -> skews the measurements

service level objectives (*SLOs*) vs. service level agreements (*SLAs*)

### Approaches for Coping with Load

scaling up (*vertical scaling*) vs. scaling out (*horizontal scaling*)

## Maintainability

Maintenance: fix bugs, keep system operational, investigate failures, adapt to new platforms, modify for new use cases, repay technical debt, etc.

### Operability: Making Life Easy for Operations

a good operation team is responsible for:

* monitoring the health of the system and quickly restoring service if it goes into a bad state
* tracking down the cause of problems
* keeping software and platforms up to date
* establishing good practices and tools for deployment, configuration management
* defining processes that make operations predictable and help keep the production environment stable

data systems could do various things to make routine operation tasks easy:

* providing visibility into the runtime behavior and internals of the systems, with good monitoring
* providing good support for automation and integration with standard tools
* avoiding dependency on individual machines/systems
* providing good documentation and an easy-to-understand operational model
* providing good default behavior
* exhibiting predictable behavior, minimizing surprises

### Simplicity: Managing Complexity

Symptoms of complexity: explosion of the state space, tight coupling of modules, tangled dependencies, inconsistent naming and terminology, hacks aimed at solving performance problems, special-casing to work around issues, etc

> When complexity makes maintenance hard, budgets and schedules are often overrun.

A good abstraction can hide a great deal of implementation detail behind a clean, simple-to-understand facade.

### Evolvability: Making Change Easy

Agile vs. Test-Driven Development (TDD) and Refactoring

## Summary

functional requirements (what it should do) vs. non-functional requirements (security, reliability, compliance, scalability, compatibility, and maintainability)
