# Chapter 2. Case Study: Designing a Document Editor

* “This chapter presents a case study in the design of a “What-You-See-Is-What-You-Get” (or “WYSIWYG”) document editor called **Lexi**.”

## 2.1 Design Problems

* Document structure: internal representation for the document.
* Formatting: arrange text and graphics into lines and columns.
* Embellishing the user interface: such as scroll bars, borders, and drop shadows.
* Supporting multiple look-and-feel standards: adapt easily to different look-and-feel standards.
* Support multiple window systems: independent of window system.
* User operations: provide a uniform mechanism both for accessing scattered functionality and for undoing its effects.
* Spelling checking and hyphenation: analytical operations.

## 2.2 Document Structure

* Goals of the internal representation:
  * “Maintaining the document’s physical structure, that is, the arrangement of text and graphics into lines, columns, tables, etc.”
  * “Generating and presenting the document visually.”
  * “Mapping positions on the display to elements in the internal representation. This lets Lexi determine what the user is referring to when he points to something in the visual representation.”
* Constraints:
  * Treat text and graphics uniformly.
  * Our implementation shouldn't have to distinguish between single elements and groups of elements in the internal representation.
  * The need to analyze the text.
* Recursive Composition
  * “A common way to represent hierarchically structured information is through a technique called **recursive composition**, which entails building increasingly complex elements out of simpler ones.”
  * “By using an object for each character and graphical element in the document, we promote flexibility at the finest levels of Lexi’s design.”
* Glyphs
  * “We’ll define a **Glyph** abstract class for all objects that can appear in a document structure. Its subclasses define both primitive graphical elements (like characters and images) and structural elements (like rows and columns).”
  * “Glyphs have three basic responsibilities. They know (1) how to draw themselves, (2) what space they occupy, and (3) their children and parent.”

| Responsibility    | Operations                              |
| ----------------- | --------------------------------------- |
| **appearance**    | `virtual void Draw(Window*)`            |
|                   | `virtual void Bounds(Rect&)`            |
| **hit detection** | `virtual bool Intersects(const Point&)` |
| **structure**     | `virtual void Insert(Glyph*, int)`      |
|                   | `virtual void Remove(Glyph*)`           |
|                   | `virtual Glyph* Child(int)`             |
|                   | `virtual Glyph* Parent()`               |

* Composite Pattern
  * “The Composite pattern captures the essence of recursive composition in object-oriented terms.”


## 2.3 Formatting

* **Encapsulating the Formatting Algorithm**
  * “Because Lexi is a WYSIWYG editor, an important trade-off to consider is the balance between formatting quality and formatting speed.”
  * “Because formatting algorithms tend to be complex, it’s also desirable to keep them well-contained or—better yet—completely independent of the document structure.”
  * “We should design Lexi so that it’s easy to change the formatting algorithm at least at compile-time, if not at run-time as well.”
    * “More specifically, we’ll define a separate class hierarchy for objects that encapsulate formatting algorithms.”
* **Compositor and Composition**
  * “We’ll define a **Compositor** class for objects that can encapsulate a formatting algorithm. The interface (Table 2.2) lets the compositor know *what* glyphs to format and *when* to do the formatting.”
    * “The glyphs it formats are the children of a special Glyph subclass called Composition.”
    * “A composition gets an instance of a Compositor subclass (specialized for a particular linebreaking algorithm) when it is created, and it tells the compositor to `Compose` its glyphs when necessary.”
    * “An unformatted Composition object contains only the visible glyphs that make up the document’s basic content. It doesn’t contain glyphs that determine the document’s physical structure, such as Row and Column.”
    * “When the composition needs formatting, it calls its compositor’s `Compose` operation. The compositor in turn iterates through the composition’s children and inserts new Row and Column glyphs according to its linebreaking algorithm.”


| Responsibility | Operations                          |
| -------------- | ----------------------------------- |
| what to format | `void SetComposition(Composition*)` |
| when to format | `virtual void Compose()`            |

* **Strategy Pattern**
  * “Encapsulating an algorithm in an object is the intent of the Strategy (315) pattern.”
    * “Compositors are strategies; they encapsulate different formatting algorithms.”


## 2.4 Embellishing the User Interface

* **Transparent Enclosure**
  * Inheritance extension -> no rearranging embellishments at run-time -> explosion of subclasses.
  * Object composition -> more workable and flexible extension.
  * Have the border contain the glyph makes sense because no modification is required to the corresponding Glyph subclass.
  * Transparant enclosure: single-child composition, compatible interfaces.
    * The enclosure delegates all its operations to its component.
* **Monoglyph**
  * MonoGlyph extends Glyph to serve as an abstract class for embellishment glyphs.
    * “MonoGlyph stores a reference to a component and forwards all requests to it that makes MonoGlyph totally transparent to clients by default. ”
    * MonoGlyph subclasses reimplement at least one of these forwarding operations.
* **Decorator Pattern**
  * “The Decorator pattern captures class and object relationships that support embellishment by transparent enclosure.”
    * “In the Decorator pattern, embellishment refers to anything that adds responsibilities to an object.”


## 2.5 Supporting Multiple Look-and-Feel Standards

* Diversity -> enforce uniformity between applications.
* Standards -> guidelines for how applications appear and react to the user.
* Design goals -> to make Lexi conform to multiple standards and easy to support newly introduced standards.
* **Abstracting Object Creation**
  * Two sets of widget glyph classes:
    * “A set of abstract Glyph subclasses for each category of widget glyph.”
      * “For example, an abstract class ScrollBar will augment the basic glyph interface to add general scrolling operations; Button is an abstract class that adds button-oriented operations; and so on.”
    * “A set of concrete subclasses for each abstract subclass that implement different look-and-feel standards.”
      * “For example, ScrollBar might have Motif ScrollBar and PMScrollBar subclasses that implement Motif and Presentation Manager-style scroll bars, respectively.”
  * Avoid making explicit calls to create widgets.
* **Factories and Product Classes**
  * `ScrollBar* sb = guiFactory->CreateScrollBar();`
  * GUIFactory is an abstract class that defines a general interface for creating widget glyphs.
  * “Regardless of how and when we decide to initialize `guiFactory`, we know that once we do, the application can create the appropriate look and feel without modification.”
* **Abstract Factory Pattern**
  * Abstract Factory Pattern creates families of related product objects without instantiating classes directly, which is appropriate when the number and general kinds of product objects stay constant, and there are differences in specific product families.
    * Choose among families by instantiating a particular concrete factory.

## 2.6 Supporting Multiple Window Systems

* **Encapsulating Implementation Dependencies**
  * Encapsulated Window class:
    * “They provide operations for drawing basic geometric shapes.”
    * “They can iconify and de-iconify themselves.”
    * “They can resize themselves.”
    * “They can (re)draw their contents on demand”
  * “The Window class must span the functionality of windows from different window systems.”

| Responsibility        | Operations                      |
| --------------------- | ------------------------------- |
| **window management** | `virtual void Redraw()`         |
|                       | `virtual void Raise()`          |
|                       | `virtual void Lower()`          |
|                       | `virtual void Iconify()`        |
|                       | `virtual void Deiconify()`      |
|                       | ...                             |
| **graphics**          | `virtual void DrawLine(...)`    |
|                       | `virtual void DrawRect(...)`    |
|                       | `virtual void DrawPolygon(...)` |
|                       | `virtual void DrawText(...)`    |
|                       | ...                             |

* **Window and WindowImp**
  * Define a separate **WindowImp** class hierarchy -> hide different window system implementations.
  * WindowImp: abstract class -> encapsulate window system dependent code.
  * Each window object with an instance of a WindowImp subclass for that system.
* **Bridge Pattern**
  * The intent behind Bridge is to allow separate class hierarchies to work together even as they evolve independently.
    * One for logical notion, another capturing different implementations.
    * Enhance logical abstractions without touching environment-dependent code.

## 2.7 User Operations

* Goal: multiple user interfaces to the same operations + change the interface in the future + implenentors want lightweight dependencies + support limited number of undo and redo of most *but not all* its functionality.
* **Encapsulating a Request**
  * MenuItem: subclass of Glyph + respond to a request from a client.
  * Define a subclass of MenuItem for every user operation -> proliferation of subclass implementations work + coupling teh request to a particular user interface.
  * Parameterize menu items by the request they should fulfill within an *object* -> use inheritance to extend and reuse the request's implementation + a place to store state and implement undo/redo.
* **Command Class and Subclasses**
  * **Command**: an abstract class to provide an interface for issuing a request.
    * “The basic interface consists of a single abstract operation called “Execute.”
    * “Subclasses of Command implement Execute in different ways to fulfill different requests. ”
  * Give each menu item an instance of Command subclass suitable for that menu item.
* **Undoability**
  * Add Unexecute operation to Command's interface -> reverse effects of a preceding Execute operation using whatever undo information Execute stored.
  * No net effect if execute a command -> no need for a undo request.
    * To determine if a command is undoable -> add an abstract Reversible operation to the Command interface.
* **Command History**
  * Define a **command history**: a list of commands that have been executed.
  * To redo a command that's just been undone, we do the same thing in reverse.
* **Command Pattern**
  * The Command pattern -> prescribe a uniform interface for issuing requests + shield clients from the request's implementation.
  * A command may delegate all, part, none of the request's implementation to other objects.

## 2.8 Spelling Checking and Hyphenation

* Design problem: textual analysis, specifically checking for misspellings and introducing hyphenation points.
* Design goal: a diverse set of algorithms for space/time/quality trade-offs + avoid coupling analysis with document structure.
* Puzzles:
  * accessing the information to be analyzed
  * doing the analysis
* **Accessing Scattered Information**
  * Access mechanism must accommodate differing data structures, and we must support different kinds of traversals, such as preorder, postorder, and inorder.
* **Encapsulating Access and Traversal**
  * The approach of adding abstract operations to Glyph's interface still has problems:
    * Can't support new traversals.
    * Have to change existing declarations.
    * Hard to modify or extend without changing lots of classes.
    * Difficult to reuse the mechanism to traverse other kinds of object structures.
    * Can't have more than one traversal in progress on a structure.

```cpp
void First(Traversal kind)
void Next()
bool IsDone()
Glyph* GetCurrent()
void Insert(Glyph*)
  
Glyph* g;
for (g->First(PREORDER); !g->IsDone(); g->Next()) {
  Glyph* current = g->GetCurrent();
  // do some analysis
}
```

* **Iterator Class and Subclasses**
  * **Iterator**: A general interface for access and traversal.
    * Each Iterator subclass has a reference to the structure it traverses.
  * “A glyph subclass that has children will override CreateIterator to return an instance of a different Iterator subclass.”


```cpp
Glyph* g;
Iterator<Glyph*>* i = g->CreateIterator();
for (i->First(); !i->IsDone(); i->Next()) {
  Glyph* child = i->CurrentItem();
  // do something with current child
}
```

* **Iterator Pattern**
  * The Iterator pattern supports access and traversal over object structures.
    * Applicable to composite structures and collections.
    * Abstract the traversal algorithm and shield clients from the internal structure.
* **Traversal versus Traversal Actions**
  * We want to accumulate information during the traversal.
    * We can actually reuse the same set of iterators for different analyses.
  * Analysis must be able to distinguish different kinds of glyphs.
  * Put the analytical capability into the glyph classes -> interface expands with every new analytical capability -> obscure the basic Glyph interface.
* **Encapsulating the Analysis**
  * A separate object -> use in conjunction with an iterator -> perform analysis at each point in the traversal.
  * But how the analysis object distinguishes different kinds of glyphs without resorting to type casts or downcasts?
    * We can add `void CheckMe(SpellingChecker&)` to the Glyph class where, in turn, the `SpellingChecker` class includes an operation like `CheckGlyphSubclass` for every Glyph subclass.
    * It seems we have to add an operation like `CheckMe(SpellingChecker&)` to Glyph and its subclasses whenever we add a new kind of analysis.
    * But we can actually give all analysis classes the same interface -> an analysis-independent operation on the Glyph subclass.
* **Vistor Class and Subclasses**
  * Concrete subclasses of `Visitor` perform different analyses.
  * `CheckMe` for Glyph -> `Accept` -> takes a `Visitor&` as argument -> reflect the fact that it can accept any visitor.
  * New analysis requires a new subclass of `Visitor`.

```cpp
class Visitor {
public:
  virtual void VisitCharacter(Character*) { }
  virtual void VisitRow(Row*) { }
  virtual void VisitImage(Image*) { }
  // ... and so forth
}
```

* **Visitor Pattern**
  * The Visitor pattern support an open-ended number of analyses of the object structure without having to change the object classes themselves.
    * Applicable to composites or any object structure.
  * Visitors can work across class hierarchies.
  * Suitable when performing different operations on a stable class structure.
    * But whenever you add a subclass to the structure -> have to update all visitor interfaces to include a `Visit...` operation for that subclass.

## 2.9 Summary

* Composite to represent the document's physical structure.
* Strategy to allow different formatting algorithms.
* Decorator for embellishing the user interface.
* Abstract Factory for supporting multiple look-and-feel standards.
* Bridge to allow multiple window platforms.
* Command for undoable user operations.
* Iterator for accessing and traversing object structures.
* Visitor for allowing an open-ended number of analytical capabilities without complicating the document structure's implementation.
