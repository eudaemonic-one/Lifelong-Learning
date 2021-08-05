# Lecture 23 Building a Next-Generation Serverless Platform for Converged AI/ML Workflows

## Hybrid Cloud

* **Cloud is Changing Discovery**
* Massive-scale, resource pooling rivals traditional large-scale systems
* New programming models dramatically simlify complex, large-scale workloads on Cloud
* Elastic, scalable data storage is changing data gravity landscape
* A paradigm shift fundamentally transforming large-scale workloads
  * Scale on demand
  * Simple user experience, infrastructure abstraction
  * Agility, standardization, open-source innovation
  * Portability, choice, access to competitive offerings

## Cloud AI

* Modeling + AI + Data Analytics
  * From siloed, model based simulation, to re-imagined, AI-enhanced hybrid workflows
* Emerging HPC (AI Everywhere)
  * Machine Learning
  * Simulation
  * Data Analytics
* Common reference patterns for AI-enabled workflows
  * Tightly Coupled
    * ML-in the loop
    * In-situ analytics
    * Co-located
    * Constrained response time
    * e.g., Massively Parallel Infrastructure for Adaptive Multiscale Simulation
  * Loosely coupled
    * Multi-domain
    * Physically distributed (on-prem+cloud)
    * Relaxed response time constraints
  * Recasting
    * Replacing existing methods with purely data-driven approaches
    * Often interactive
* Data-driven workflows
  * e.g., Reperforming a Nobel Prize Discovery on Kubernetes (CERN): elastic, on-demand platform for a data-driven approach

## Next Generation Workflow

* High-throughput, parallel analytics
* Modeling Simulation
* AL, neural networks
* Users are trying to weave capabilities together into more complex workflows
* This requires a unifying programming layer that seamlessly combines different software stacks and ecosystems
* With parts of the workflow running in different environments, users need a unified hybrid cloud platform with hybrid data orchestration
* Users can then leverage specialized hardware and infrastructure to optimize workload performance

## Future of Serverless

* Going beyond Function as a Service
  * Code on-demand automatically scaled
  * Pay for the time the code is running only
  * Allocation and deallocation of resources automatically driven by the workload
  * No infrastructure management - getting IT out of the way
  * No capacity management (an infinite pool of resources)
* Intent-driven adaptive software
  * Action and dependencies description transparently mapped into resources
  * Code execution on-demand automatically scaled
  * Allocation and deallocation of resources automatically driven by the workload
  * Runtime optimized for problems size/cost/performance
  * Pay for the time the code is running

### Serverless 1.0

* Functions-as-a-Service (FaaS)
* Data-shipping Architecture
  * Short-lived and non-addressable functions
  * Ephemeral, stateless
* No direct network addressability
  * No point-to-point communication
  * Data passed over some slow and expensive storage medium with high latency
* No Specialized Hardware
  * Timeslice of a CPU hyper thread and some amount of RAM

### Serverless 2.0

* Large-footprint, Long-running, Addressable agents
* Functional model with high-level, rich semantics
  * Stateful long-lived virtual addressable components
  * Dataflow and distributed states management at scale
  * Co-location and locality awareness support
* Direct communication
  * Point-to-point communication
  * Rich set of execution and communicating primitives
  * Unified data and storage support (caching)
* Heterogeneous Hardware Support
  * Automated allocation of code to specialized hardware

### Building End-to-End Distributed Applications

* Serverless (Easy to develop)
  * e.g., lambda
* Ad-hoc
  * e.g., MPI, Spark
* Re-factor (General)
  * e.g., gRPC, Docker
* AI, neural networks
* HPC
* Microservices

### Building the Serverless HPC Platform

* More complex workflows require advanced execution patterns, and low-latency data exchange support

## Emerging AI and Big Data

* Emerging workflows are complex pipelines requiring integration and coordination of different tools and runtimes
* Re-imaging middleware for discovery
  * Single unified runtime
  * Rapid, automated data sharing
  * Automated deployment, resource management and scaling across cloud environments

### Ray - An emerging distributed computing framework

* A single, dynamic graph execution (DAG) for both stateless (task) and stateful (actor) computing patterns
* Decoupling of compute specification and execution
* In-memory object store
* Transparent resource management
* Proven low-latency runtime (millions of tasks/second)
* High-performance compute/communication primitives that on-par performance with MPI

### Middleware for Converged Workflows

* Integration of multiple runtimes
* Efficient data transfer
* Simplified pipeline definition and scaling
* Simplified user experience for end-to-end workflows