# Lecture 24 Inside Azure Datacenter Architecture

## Inside Datacenters

* Azure from Cloud to Edge
  * Azure Sphere / Azure IoT devices, Azure Stack Hub and HCI / Azure Private Edge Zones / Azure Edge Zones / Azure Regions
  * Consistent security, identity, management, and AI
* 65+ Azure regions
* Azure region architecture
  * Data Residency Boundary -> Region -> Availability Zone -> Server Clusters
* Modular Datacenters
  * High availability module for uninterrupted power
  * Azure Energy sustainability

## Inside Intelligent Infrastructure

### AIOps

* Fast and actionable anomaly detection
* Auto-communication
* Automatic impacted service identification
* Impact assement and incident aggregation
* Root cause service identification
* Efficient outage management

### Brain

* Aggregator
  * Customer experience
  * Azure Services
  * Infrastructure devices
  * Critical Environment and Mechanical

### Resource Central

* Model training platform
* Services and features -> Inference -> Model training
* Live data and validation -> Model training
* Usage metadata -> Model development

#### Project Narya

* Predictive and adaptive failure prevention
* ML failure prediction + Prediction rule -> A/B testing + Multi-armed bandit decision -> Composite actions (Block allocation, Live migrate, Service heal, Soft kernel reboot) -> Customer impact + Diagnostics -> Feedback loop

### Azure Digital Twins

* Input: IoT Hub / Workflow integration / Business systems / services integration via REST APIs
* Azure Digital Twins -> Client appas / External compute
* Output: Workflow integration / Cold Storage / Time Series Insight / Azure Synapse Analytics

## Inside Networking

### Rack Hardware

* SmartNIC/FPGA
* Software Defined Appliances

### DC Hardware

* Azure Firewall
* Azure DDos detection
* Load balancing

### Intra-region

* Regional network gateways

### Microsoft WAN

* Dedicated dark fiber WAN
* Optical networking (DWDM)
* Network Watcher
* Network Performance Monitoring

### Last Mile and Enterprise Connectivity

* ExpressRoute
* CDN

### Azure Orbital

* Ground Station as a Service
  * Earth observation
  * Global communication

## Inside Servers

* Memory optimized servers
  * Godzilla (2014)
  * Beast (2017)
  * Beast V2 (2019)
  * Mega-Godzilla-Beast (2020)
* Massively scalable AI supercomputer
  * NDv4
* Liquid cooling
* Quantum controller

## Inside Azure Resource Manager

* Azure portal / CLI and PowerShell / SDKs / REST API
* ARM (Auth and RBAC / Activity logs and telemetry / Resource metadata)
  * Traffic Manager -> Load Balancers
* Resource Providers (Compute / Network / Storage / SQL)
* Azure infrastructure (Hardware manager / Edge infrastructure)

### VM Applications

* Azure DevOps + Github -> CI/CD
* Application Package
* Azure Artifacts Gallery (Private/Public) -> Install/Update
* Virtual Machines / Virtual Machine Scale Sets

## Inside Compute

### Dapr

* Programming languages and frameworks
* Building blocks
* Cloud infrastructures

## Inside Storage and Data

### Data protection

* Existing encryption -> Data at rest / Data in transit
* Confidential computing -> Data in use

### Azure Storage Portfolio

* Disk storage
* Object storage
* File storage
* Backup
* Data transport
* Hybrid storage
* Future technology