# Chapter 18: Building for Reliability

## Think About the Future

* **What could happen, and how will we be ready?
  * What could break?
  * How can the system keep working if something breaks?
  * What could fail without realizing it? How can we make sure we'll be alerted?
  * What might we be asked to (or want to) change in the future? Will it be easy?
  * What would I want to know if something breaks?
  * What behavior would I want to be able to control or change if something goes wrong?
  * What could users do that we don't intend?
  * What could be inconvenient, slow, or error-prone when we have customer problem?
  * What might trip up a future developer of the system?
  * What precondition or required invariant isn't documented or tested?

## Design for Failure

* Look beyond teh goal of not crashing:
  * achieve unaltered or gently degraded user experiences even when things go wrong
  * can be recovered when they experience disaster, support disaster recovery.
* Failure resilience techniques:
  * **Redundancy**
  * **Graceful degradation**

## Plan for Data Loss

* The only safe routes:
  * **Offline copies of your data**, not just replicas.
  * **System state restorable from source control.**

## Build for Debugging (or Observability Is Everything)

* Build observability into a systems lets us ensure that
  * we can detect if it breaks;
  * when that happens, we have the tooling to figure out why.

## Build for Opeartor Intervention

* Build endpoints, dynamic configuration, UI controls to rapidly change a system's behavior without building and deploying new code.

## Build for Rollback

* New features are best deployed behind configuration control that allows them to be rolled back.
* Database schema changes should be backward-compatible.

## Don't Order Dependencies; Wait

* Better is to have each component calmly wait for its dependencies, retrying with some backoff; then you never need to maintain an explicit graph, and your system will tend to heal when the unexpected happens.

## Plan for Overload

* We may experience temporary request spikes beyond a system's capacity.
* The most common reason for that prolonged downtime is excessive queueing.
* Do
  * Bound your queue sizes; consider LIFO rather than FIFO!
  * Use a circuit-breaking library any time you use a service client.
  * Use an inbound loadshedder if you can!
  * Ensure you have some way to rate-limit inbound traffic.

## Data Layer Isolation (In Which Database Migrations Are Our Bigest Problem)

* Migrating from one system to another.
* MVCS pattern helps keep data access independently.
