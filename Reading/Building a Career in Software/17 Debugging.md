# Chapter 17: Debugging

## The Philosophy of Debugging

* Debug in a loop of hypothesizing and experimenting.
  * Mentally model system; hypothesize what's wrong if you can.
  * Imagine what data would most precisely confirm, disprove, or refine our theory.
  * Get that data.
  * Confirm, reject, or refine the model.

## Just Do Something, or "Debugging is Hard, Though"

* In the most orderly debugging processes, we steadily refine our hypothesis by collecting data of intuitive value.

## Debugging Menu

* **Logs**
  * e.g., syslog, systemed journal, Splunk, Kibana, /var/log, docker logs, database logs, kubectl logs.
  * try filtering by process name, service name, hostname, screen, username, entity ID, time.
* **Metrics**
  * e.g., counts, latencies, rates like on-disk bytes, response time, error rate.
  * => describe a system's aggregate behavior, and capture general systemic degradation.
* **Chrome dev tools**
  * e.g., UI misssing behavior, failing requests to the backend, or weird data.
* **Just grep the code**
* **Add new instrumentation**
  * => add logs on your laptop, in staging, or in a production environment.
* **Local system monitors**
  * e.g., `top`, `htop`, `iotop`, `iostat`.
* **git bisect**
  * => check out one commit at a time, and test it until you find the earliest broken one.
* **Debugger**
  * printing variables, stepping through functions, and scripting data structure traversals.
* **CPU profiler**
  * provides a view of where a program spends its CPU cycles over time => understand performance problems and crashes.
* **Heap introspection**
  * => for memory leaks.
* **tcpdump/ngrep/wireshark**
  * TCP interactions, encoding issues, or network issues.
* **Tracing frameworks**
  * e.g., Linux's strace, Linux perf, DTrace.

## Running Experiments

* The steps to run an experiment:
  * installing builds, restarting systems, changing configuration, hooking up external instruments, monitoring during the run, and finally collecting the right data at the end.
* Two methods:
  * **Making a detailed runbook** (checklist) ensure that all dashboarding, data collection docs, are prepared in advance.
  * **Recording *everything* done during the experiment.**
    * with screenshots, copy-pastes, and links in a single Google Doc; time-stamping the completion of stages of the checklist.
