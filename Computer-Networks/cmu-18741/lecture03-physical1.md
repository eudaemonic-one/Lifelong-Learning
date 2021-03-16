# Lecture 3 Physical Layer I

## End-to-end Argument

* End-to-end deals with where to place functionality
  * Inside the network (in switching elements)
  * At the edges
* Reliable file transfer
  * Solution 1: make each step reliable, and then concatenate them
  * Solution 2: end-to-end check and retry

## Digital Networking

* Analog information takes on continuous values
* Digital information takes on discrete values

### Block vs. Stream

* Block
  * e.g., text message, data file, JPEG, MPEG
* Stream
  * e.g., real-time voice, streaming video

### Why Digital?

* Economically advantageous
* Multimedia applications want to mix different types of data
* Digital transmission can recover from errors (e.g., noise, distortion)
* Problems with analog long-distance communication
  * Each repeater attempts to restore analog signal to its original form
  * Restoration is imperfect
    * Distortion is not completely eliminated
    * Noise & interference is only partially removed
  * Signal quality decreases with # of repeaters
* Digital long-distance communication
  * Regenerator recovers original data sequence and restransmits on next segment
  * Can design so error probability is very small