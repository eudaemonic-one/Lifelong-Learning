# Object-Oriented Design

## How to Approach

* **Step 1: Handle Ambiguity**
  * Object-oriented design (OOD) questions are often intentionally vague in order to test whether you'll make assumptions or if you'll ask clarifying questions
  * You should inquire who is going to use it and how they are going to use it
* **Step 2: Define the Core Objects**
  * For example, suppose we are asked to do the object-oriented design for a restaurant
  * Our core objects might be things like Table, Guest, Party, Order, Meal, Employee, Server, and Host
* **Step 3: Analyze Relationships**
  * Which objects are members of which other objects?
  * Do any objects inherit from any others?
  * Are relationships many-to-many or one-to-many?
* **Step 4: Investigate Actions**
  * What remains is to consider the key actions that the objects will take and how they relate to each other
  * You may find that you have forgotten some objects, and you will need to update your design

## Design Patterns

* The Singleton and Factory Method design patterns are widely used in interviews

### Singleton Class

* The Singleton pattern ensures that a class has only one instance and ensures access to the instance through the application
* It can be useful in cases where you have a "global" object with exactly one instance

```java
public class Restaurant {
  private static Restaurant _instance = null;
  protected Restaurant() { ... }
  public static Restaurant getInstance() {
    if (_instance == null) {
      _instance = new Restaurant();
    }
    return _instance;
  }
}
```

### Factory Method

* The Factory Method offers an interface for creating an instance of a class, with its subclasses deciding which class to instantiate
* You might want to implement this with the creator class being abstract and not providing an implementation for the Factory method
* Or, you could have the Creator class be a concrete class that provides an implementation for the Factory method

```java
public class CardGame {
  public static CardGame createCardGame(GameType type) {
    if (type == GameType.Poker) {
      return new PokerGame();
    } else if (type == GameType.BlackJack) {
      return new BlackJackGame();
    }
    return null;
  }
}
```

## Interview Questions

* **7.1 Deck of Cards:**
  * Design the data structures for a generic deck of cards. Explain how you would subclass the data structures to implement blackjack.
* **7.2 Call Center:**
  * Imagine you have a call center with three levels of employees: respondent, manager, and director. An incoming telephone call must be first allocated to a respondent who is free. If the respondent can't handle the call, he or she must escalate the call to a manager. If the manager is not free or not able to handle it, then the call should be escalated to a director. Design the classes and data structures for this problem. Implement a method dispatchCall() which assigns a call to the first available employee.
* **7.3 Jukebox:**
  * Design a musical jukebox using object-oriented principles.
* **7.4 Parking Lot:**
  * Design a parking lot using object-oriented principles.
* **7.5 Online Book Reader:**
  * Design the data structures for an online book reader system.
* **7.6 Jigsaw:**
  * Implement an NxN jigsaw puzzle. Design the data structures and explain an algorithm to solve the puzzle. You can assume that you have a fitsWith method which, when passed two puzzle edges, returns true if the two edges belong together.
* **7.7 Chat Server:**
  * Explain how you would design a chat server. In particular, provide details about the various backend components, classes, and methods. What would be the hardest problems to solve?
* **7.8 Othello:**
  * Othello is played as follows: Each Othello piece is white on one side and black on the other. When a piece is surrounded by its opponents on both the left and right sides, or both the top and bottom, it is said to be captured and its color is flipped. On your turn, you must capture at least one of your opponent's pieces. The game ends when either user has no more valid moves. The win is assigned to the person with the most pieces. Implement the object-oriented design for Othello.
* **7.9 Circular Array:**
  * Implement a CircularArray class that supports an array-like data structure which can be efficiently rotated. If possible, the class should use a generic type (also called a template), and should support iteration via the standard for (Obj o : circularArray) notation.
