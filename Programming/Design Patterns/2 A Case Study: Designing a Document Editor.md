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
