# Chapter 9. GUI Applications

* Swing framework
  * certain tasks must run in the Swing event thread => safety.
  * you cannot execute long-running tasks in the event thread => lest the UI become unresponsive.
  * Swing data strutcures are not thread-safe => confine them to the event thread.
* Nearly all GUI toolkits are implemented as *single-threaded subsystems* in which all GUI activity is confined to a single thread.
