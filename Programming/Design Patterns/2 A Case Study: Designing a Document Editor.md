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
