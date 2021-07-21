# Chapter 7. Cancellation and Shutdown

* Java does not provide any mechanism for safely forcing a thread to stop.
* Instead, it provides *interruption*, a cooperative mechanism that lets one thread ask another to stop.
* We rarely want a task, thread, or service to stop *immediately*, since that could leave shared data structures in an inconsistent state.
* Instead, tasks and services can be coded so that, when requested, they clean up any work currently in progress and *then* terminate.
