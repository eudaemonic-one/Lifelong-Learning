# Math and Logical Puzzles

## Prime Numbers

### Divisibility

* The prime number law stated above means that, in order for a number x to divide a number y (written $x\\y$, or $mod (y, x) = 0$), all primes in x's prime factorization must be in y's prime factorization

### Checking for Primality

* Iterate only up through the square root of n, checking for divisibility on each iteration

```java
boolean prime(int n) {
	if (n < 2) {
    return false;
  }
  int sqrt = (int) Math.sqrt(n);
  for (int i = 2; i <= sqrt; i++) {
    if (n % i == 0) {
      return false;
    }
  }
  return true;
}
```

### Generating a List of Primes: The Size of Eratosthenes

* The Sieve of Eratosthenes works by recognizing that all non-prime numbers are divisible by a prime number

```java
boolean[] sieveOfEratosthenes(int max) {
  boolean[] flags = new boolean[max + 1];
  int count = 0;
  init(flags); // Set all flags to true other than 0 and 1
  int prime = 2;
  while (prime <= Math.sqrt(max)) {
    crossOff(flags, prime);
    prime = getNextPrime(flags, prime);
  }
  return flags;
}

void crossOff(boolean[] flags, int prime) {
  for (int i = prime * prime; i < flags.length; i += prime) {
    flags[i] = false;
  }
}

int getNextPrime(boolean[] flags, int prime) {
  int next = prime + 1;
  while (next < flags.length && !flags[next]) {
    next++;
  }
  return next;
}
```

## Probability

### Probability of A and B

* $P(A and B) = P(B given A) P(A)$
* $P(A given B) = P(B given A) P(A) / P(B)$

### Probability of A or B

* $P(A or B) = P(A) + P(B) - P(A and B)$

### Independence

* $P(A and B) = P(A) P(B)$

### Mutual Exclusivity

* $P(A or B) = P(A) + P(B)$
* $P(A and B) = 0$

## Develop Rules and Patterns

* In many cases, you will find it useful to write down "rules" or patterns that you discover while solving the problem

### Brainteaser Example

```text
You have two ropes, and each takes exactly one hour to burn. How would you use them to time exactly 15 minutes? Note that the ropes are of uneven densities, so half the rope length-wise does not necessarily take half an hour to burn.
```

* Rule 1: Given a rope that takes x minutes to burn and another that takes y minutes, we can time x+y minutes
* Rule 2: Given a rope that takes x minutes to burn, we can time x/2 minutes
* Rule 3: If rope 1 takes x minutes to burn and rope 2 takes y minutes, we can turn rope 2 into a rope that takes (y - x) minutes or (y - x/2) minutes

## Worst Case Shifting

* Many brainteasers are worst-case minimization problems, worded either in terms of minimizing an action or in doing something at most a specific number of times
* A useful technique is to try to "balance" the worst case
  * That is, if an early decision results in a skewing of the worst case, we can sometimes change the deci­sion to balance out the worst case

```text
The "nine balls" question is a classic interview question. You have nine balls. Eight are of the same weight, and one is heavier. You are given a balance which tells you only whether the left side or the right side is heavier. Find the heavy ball in just two uses of the scale.
```

## Interview Questions

* **6.1 The Heavy Pill:**
  * You have 20 bottles of pills. 19 bottles have 1.0 gram pills, but one has pills of weight 1.1 grams. Given a scale that provides an exact measurement, how would you find the heavy bottle? You can only use the scale once.
* **6.2 Basketball:**
  * You have a basketball hoop and someone says that you can play one of two games.
  * Game 1: You get one shot to make the hoop.
  * Game 2: You get three shots and you have to make two of three shots.
  * If p is the probability of making a particular shot, for which values of p should you pick one game or the other?
* **6.3 Dominos:**
  * There is an 8x8 chessboard in which two diagonally opposite corners have been cut off. You are given 31 dominos, and a single domino can cover exactly two squares. Can you use the 31 dominos to cover the entire board? Prove your answer (by providing an example or showing why it's impossible).
* **6.4 Ants on a Triangle:**
  * There are three ants on different vertices of a triangle. What is the probability of collision (between any two or all of them) if they start walking on the sides of the triangle? Assume that each ant randomly picks a direction, with either direction being equally likely to be chosen, and that they walk at the same speed.
* **6.5 Jugs of Water:**
  * You have a five-quart jug, a three-quart jug, and an unlimited supply of water (but no measuring cups). How would you come up with exactly four quarts of water? Note that the jugs are oddly shaped, such that filling up exactly "half" of the jug would be impossible.
* **6.6 Blue-Eyed Island:**
  * A bunch of people are living on an island, when a visitor comes with a strange order: all blue-eyed people must leave the island as soon as possible. There will be a flight out at 8:00 pm every evening. Each person can see everyone else's eye color, but they do not know their own (nor is anyone allowed to tell them). Additionally, they do not know how many people have blue eyes, although they do know that at least one person does. How many days will it take the blue-eyed people to leave?
* **6.7 The Apocalypse:**
  * In the new post-apocalyptic world, the world queen is desperately concerned about the birth rate. Therefore, she decrees that all families should ensure that they have one girl or else they face massive fines. If all families abide by this policy-that is, they have continue to have children until they have one girl, at which point they immediately stop-what will the gender ratio of the new generation be? (Assume that the odds of someone having a boy or a girl on any given pregnancy is equal.) Solve this out logically and then write a computer simulation of it.
* **6.8 The Egg Drop Problem:**
  * There is a building of 100 floors. If an egg drops from the Nth floor or above, it will break. If it's dropped from any floor below, it will not break. You're given two eggs. Find N, while minimizing the number of drops for the worst case.
* **6.9 100 Lockers:**
  * There are 100 closed lockers in a hallway. A man begins by opening all 100 lockers. Next, he closes every second locker. Then, on his third pass, he toggles every third locker (closes it if it is open or opens it if it is closed). This process continues for 100 passes, such that on each pass i, the man toggles every ith locker. After his 100th pass in the hallway, in which he toggles only locker #100, how many lockers are open?
* **6.10 Poison:**
  * You have 1000 bottles of soda, and exactly one is poisoned. You have 10 test strips which can be used to detect poison. A single drop of poison will turn the test strip positive permanently. You can put any number of drops on a test strip at once and you can reuse a test strip as many times as you'd like (as long as the results are negative). However, you can only run tests once per day and it takes seven days to return a result. How would you figure out the poisoned bottle in as few days as possible?
