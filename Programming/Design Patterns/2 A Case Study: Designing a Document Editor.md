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
