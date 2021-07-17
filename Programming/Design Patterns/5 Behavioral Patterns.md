# Chapter 5. Behavioral Patterns

* Behavioral patterns: algorithms and the assignment of responsibilities between objects + communication between them.
* Behavioral class patterns: use inheritance to distribute behavioral between classes.
* Behavioral object patterns: use object composition rather than inheritance.
  * How do peer objects know who to cooperate to perform some tasks with?

## Object Behavioral: Chain of Responsibility

* **Intent**
  * Avoid coupling the sender of a request to its receiver by giving more than one object change to handle the request.
  * Chain the receiving objects and pass the request along the chain until an object handles it.
* **Motivation**
  * Decouple the object that initiates the request from the objects that might handle the request.
  * The request has an implicit receiver.
  * Each object on the chain shares a common interface for handling requests and for accessing its **successor** on the chain.
* **Applicability**
  * Use then
    * more than one object may handle a request, and the handler isn't known *a priori*.
    * you want to issue a request to one of several objects without specifying the receiver explicitly.
    * the set of objects that can handle a request should be specified dynamically.
* **Structure**

![pg225fig01](images/5 Behavioral Patterns/pg225fig01.jpg)

![pg225fig02](images/5 Behavioral Patterns/pg225fig02.jpg)

* **Participants**
  * **Handler**
    * defines an interface for handling requests.
    * (optional) implements the successor link.
  * **ConcreteHandler**
    * handles requests it is responsible for.
    * can access its successor.
    * if the ConcreteHandler can handle the request, it does so; otherwise it forwards the request to its successor.
  * **Client**
    * initiates the request to a ConcreteHandler object on the chain.
* **Collaborations**
  * When a client issues a request, the request propagates along the chain until a ConcreteHandler object takes responsibility for handling it.
* **Consequences**
  * Reduced coupling.
    * Keep a single reference to their successor -> simplify object inter-connections.
  * Added flexibility in assigning responsibilities to objects.
    * Add or change responsibilities by adding or changing the chain at run-time.
  * Receipt isn't guaranteed.
    * The chain should be configured properly.
* **Implementation**
  * Implementing the successor chain.
    * Define new links.
    * Use existing links.
  * Connecting successors.
    * The Handler might provide a default implementation to forward the request to the successor -> ConcreteHandler doesn't have to override the operation if not interested in.
  * Representing requests.
    * Hardcode operation invocation -> convenient, safe.
    * A single handler function that takes a request code -> open-ended set of requests.
    * Separate request objects that bundle request parameters -> safer parameter-passing.
      * Define request kinds and parameters by subclassing.
* **Related Patterns**
  * Often applied in conjunction with Composite.

## Object Behavioral: Command

* **Intent**
  * Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue, or log requests, and support undoable operations.
* **Also Known As**
  * Action, Transaction
* **Motivation**
  * It's necessary to issue requests to objects without knowing anything about the operation being requested or the receiver of the request.
  * Command pattern: turn request into an object -> can be stored and passed around.
  * Command declares an interface for executing operations (Execute), while the receiver has the knowledge to carry out the request.
  * MacroCommand: a concrete Command subclass executing a sequence of Commands.
* **Applicability**
  * Use when
    * parameterize objects by an action to perform.
      * **callback**: register a function to be called at a later point.
    * specify, queue, and execute requests at different times.
      * let you transfer the request to a different process and fulfill the request there.
    * support undo.
      * Execute: store state for reversing its effects.
      * Executed commands are stored in a history list.
    * support logging changes so that they can be reapplied in case of a system crash.
      * support load and store operations -> persistent log of changes -> reloading logged commands from disk and reexecuting them with Execute.
    * structure a system aroung high-level operations built on primitives operations.
      * **transaction**: encapsulate a set of operations + common interface + easy to extend.
* **Structure**

![pg236fig01](images/5 Behavioral Patterns/pg236fig01.jpg)

* **Participants**
  * **Command**
    * declares an interface for executing an operation.
  * **ConcreteCommand**
    * defines a binding between a Receiver object and an action.
    * implements Execute by invoking the corresponding operation(s) on Receiver.
  * **Client**
    * creates a ConcreteCommand object and sets its receiver.
  * **Invoker**
    * asks the command to carry out the request.
  * **Receiver**
    * knows how to perform the operations associated with carrying out a request.
* **Collaborations**
  * The client creates a ConcreteCommand object and specifies its receiver.
  * An Invoker object stores the ConcreteCommand object.
  * The invoker issues a request by calling Execute on the command.
  * The ConcreteCommand object invokes operations on its receiver to carry out the request.

![pg237fig01](images/5 Behavioral Patterns/pg237fig01.jpg)

* **Consequences**
  * Command decouples the object that invokes the operation from the one that knows how to perform it.
  * Commands are first-class object and thus can be manipulated and extended.
  * You can assemble commands into a composite command.
  * It's easy to add new Commands.
* **Implementation**
  * How intelligent should a command be?
    * Depends on its knowledge to find the receiver dynamically.
  * Supporting undo and redo.
    * Store additional state in ConcreteCommand.
    * Last command versus history list.
    * An undoable command might have to be copied before it can be placed on the history list.
  * Avoiding error accumulation in the undo process.
    * Errors can accumulate as commands are executed, unexecuted, and reexecuted repeatedly -> eventual state diverges from the original ones.
    * The Memento patterns: give the command access to the state information without exposing the internals of other objects.
  * Using C++ templates.
* **Related Patterns**
  * Composite: implement MacroCommands.
  * Memento: keep state the command requires to undo its effect.
  * A command that must be copied before being placed on the history list acts as a Prototype.

## Class Behavioral: Interpreter

* **Intent**
  * Given a language, define a representation for its grammer along with an interpreter that uses the representation to interpret sentences in the language.
* **Motivation**
  * Express problems as sentences in a simple language -> interpret these sentences with an interpreter.
  * The pattern describes how to define a grammar, represent a particular expression, and how to interpret the expression.

![248prog01](images/5 Behavioral Patterns/248prog01.jpg)

* **Applicability**
  * Use when there is a language to interpret, and you can represent statements in the language as abstract syntax trees. It works best when
    * the grammar is simple.
      * parser generator: a better alternative for complex grammer hierarchy.
    * efficiency is not a critical concern.
      * translating languages > interpreting parse trees.
* **Structure**

![pg245fig01](images/5 Behavioral Patterns/pg245fig01.jpg)

* **Participants**
  * **AbstractExpression**
    * declares an abstract Interpret operation that is common to all nodes in the abstract syntax tree.
  * **TerminalExpression**
    * implements an Interpret operation associated with terminal symbols in the grammar.
    * an instance is required for every terminal symbol in a sentence,
  * **NonterminalExpression**
    * one such class is required for every rule in the grammar.
    * maintains instance variables of type AbstractExpression for each of the symbols.
    * implements an Interpret operation for nonterminal symbols in the grammar.
  * **Context**
    * contains information that's global to the interpreter.
  * **Client**
    * builds an abstract syntax tree representing a particular sentence in the language that the grammar defines.
    * invokes the Interpret operation.
* **Collaborations**
  * The client builds the sentence as an abstract syntax tree of NonterminalExpression and TerminalExpression instances. Then the client initializes the context and invokes the Interpret operation.
  * Each NonterminalExpression node defines Interpret in terms of Interpret on each subexpression. The Interpret operation of each TerminalExpression defines the base case in the recursion.
  * The Interpret operations at each node use the context to store and access the state of the interpreter.
* **Consequences**
  * It's easy to change and extend the grammar.
    * Use inheritance to change or extend.
    * Modify existing expression incrementally.
    * Define new expressions as variations on old ones.
  * Implementing the grammar is easy, too.
  * Complex grammars are hard to maintain.
  * Adding new ways to interpret expressions.
    * e.g., pretty printing, type-checking.
* **Implementation**
  * Creating the abstract syntax tree.
    * By a table-driven parse, hand-crafted parser, or directly by the client.
  * Defining the Interpret operation.
    * Use a visitor to avoid defining operations on every grammar class.
  * Sharing terminal symbols with the Flyweight pattern.
    * Terminal nodes don't store information about their position, while parent nodes pass them the context during interpretation -> Flyweight pattern applies.
* **Related Patterns**
  * Composite: The abstract syntax tree is an instance of the Composite pattern.
  * Flyweight: share terminal symbols.
  * Iterator: traverse the structure.
  * Visitor: maintain the behavior in each node in one class.

## Object Behavioral: Iterator

* **Intent**
  * Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
* **Also Known As**
  * Cursor
* **Motivation**
  * Give a way to access elements without exposing internal structure.
  * Traverse an aggregate object in different ways.
  * **Iterator**: access + traverse out of an aggregate object.
* **Applicability**
  * Use when
    * to access an aggregate object's contents without exposing its internal representation.
    * to support multiple traversals of aggregate object.
    * to provide a uniform interface for traversing different aggregate structures (polymorphic iteration).
* **Structure**

![pg259fig01](images/5 Behavioral Patterns/pg259fig01.jpg)

* **Participants**
  * **Iterator**
    * defines an interface for accessing and traversing elements.
  * **ConcreteIterator**
    * implements the Iterator interface.
    * keeps track of the current position in the traversal of the aggregate.
  * **Aggregate**
    * defines an interface for creating an Iterator object.
  * **ConcreteAggregate**
    * implements the Iterator creation interface to return an instance of the proper ConcreteIterator.
* **Collaborations**
  * A ConcreteIterator keeps track of the current object in the aggregate and can compute the succeeding object in the traversal.
* **Consequences**
  * It supports variations in the traversaal of an aggregate.
    * Replace the iterator instance to change the traversal algorithm.
  * Iterators simplify the Aggregate interface.
  * More than one traversal can be pending on an aggregate.
* **Implementation**
  * Who controls the iteration?
    * external iterator: clients advance the traversal and request the next element explicitly from the iterator.
    * internal iterator: the client hands an internal operation to perform, and the iterator applies operation to every element in the aggregate.
    * external iterators > internal iterators.
  * Who defines the traversal algorithm?
    * cursor: the aggregate defines the traversal algorithm and use the iterator to store the state of the iteration -> might violate the encapsulation of the aggregate.
  * How robust is the iterator?
    * modify an aggregate during traversal -> dangerous.
    * simple solution: copy the aggregate and traverse the copy -> expensive.
    * **robust iterator**: insertions, removals won't interfere with traversal.
  * Additional Iterator operations.
    * minimal interface: First, Next, IsDone, CurrentItem.
    * additional operations: Previous, SkipTo.
  * Using polymorphic iterators in C++.
    * polymorphic iterators: allocated dynamically by a factory method + have their cost.
    * The client is responsible for deleting the polymorphic iterators -> error-prone to forget to free heap-allocated iterator objects.
  * Iterators may have priviledged access.
    * The iterator and the aggregate are tightly coupled.
    * The Iterator class can include `protected` operations for accessing important but publicly unavailable members of the aggregate -> Iterator subclasses gain privileged access to the aggregate.
  * Iterators for composites.
    * Use an internal iterator -> record current position by calling itself recursively -> suitable for recursive aggregate structures.
    * cursor-based iterator: a better alternative if a Composite have an interface for moving from a node to its siblings, parents, and children.
    * Common traversal patterns: preorder, postorder, inorder, breadth-first.
  * Null iterators.
    * **NullIterator**: a degenerate iterator -> *always* done with traversal -> handling boundary conditions.
* **Related Patterns**
  * Composite: often applied to recursive structures such as Composite.
  * Factory Method: instantiate the appropriate Iterator subclass with factory methods -> polymorphic iterators.
  * Memento: in conjunction with the Iterator pattern -> capture the state of an iteration.

## Object Behavioral: Mediator

* **Intent**
  * Define an object that encapsulates how a set of objects interact -> keeping objects from referring to each other explicitly -> promote loose coupling -> vary their interaction independently.
* **Motivation**
  * Proliferating interconnections -> can not work without others' support -> reduce reusability.
  * **mediator**: a separate object encapsulating collective behavior.
    * Controlling and coordinating the interactions of a group of objects.
    * Keeps object in the group from referring to each other explicitly.
    * The objects only know the mediator -> reducing the number of interconnections.
* **Applicability**
  * Use when
    * a set of objects communicate in well-defined but complex ways.
    * reusing an object is difficult because it refers to and communicates with many other objects.
    * a behavior that's distributed between several classes should be customizable without a lot of subclassing.
* **Structure**

![pg276fig01](images/5 Behavioral Patterns/pg276fig01.jpg)

![pg276fig02](images/5 Behavioral Patterns/pg276fig02.jpg)

* **Participants**
  * **Mediator**
    * defines an interface for communicating with Colleague objects.
  * **ConcreteMediator**
    * implements cooperative behavior by coordinating Colleague objects.
    * knows and maintains its colleague.
  * **Colleague classes**
    * each Colleague class knows its Mediator object.
    * each colleague communicates with its mediator whenever it would have otherwise communicated with another colleague.
* **Collaborations**
  * Colleagues send and receive requests from a Mediator object. The mediator implements the cooperative behavior by routing requests between the appropriate colleague(s).
* **Consequences**
  * It limits subclassing.
    * A mediator localizes behavior -> subclassing Mediator only to change behavior.
  * It decouples colleagues.
  * It simplifies object protocols.
    * One-to-many: easier to understand, maintain, and extend.
  * It abstracts how objects cooperate.
  * It centralizes control.
    * mediator: monolith -> hard to maintain.
* **Implementation**
  * Omitting the abstract Mediator class.
    * There's no need to define an abstract Mediator class if colleagues work with only one mediator.
  * Colleague-Mediator communication.
    * Observer pattern: the colleague sends notifications to the mediator whenever they change state -> the mediator propagates the effects to other colleagues.
    * A specialized notification interface.
* **Related Patterns**
  * Colleagues can communicate with the mediator using the Observer pattern.

