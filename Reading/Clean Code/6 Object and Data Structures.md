# Chapter 6: Objects and Data Structures

## Data Abstraction

* Hiding implementation is about abstractions.
  * A class exposes abstract interfaces that allow its users to manipulate the *essence* of the data, without having to know its implementation.

```java
public class Point {
  public double x;
  public double y;
}
```

```java
public interface Point {
  double getX();
  double getY();
  void setCartesian(double x, double y);
  double getR();
  double getTheta();
  void setPolar(double r, double theta);
}
```

```java
public interface Vehicle {
  double getFuelTankCapacityInGallons();
  double getGallonsOfGasoline();
}
```

```java
public interface Vehicle {
  double getPercentFuelRemaining();
}
```

## Data/Object Anti-Symmetry

* Objects hide their data behind abstractions and expose functions that operate on that data. Data structure expose their data and have no meaningful functions.
* Procedural code (code using data structures) makes it easy to add new functions without changing the existing data structures. OO code, on the other hand, makes it easy to add new classes without changing existing functions.
* Procedural code makes it hard to add new data structures because all functions must change. OO code makes it hard to add new functions because all the classes must change.

```java
public class Square {
  public Point topLeft;
  public double side;
}

public class Rectangle {
  public Point topLeft;
  public double height;
  public double width;
}

public class Circle {
  public Point center;
  public double radius;
}

public class Geometry {
  public final double PI = 3.141592653589793;

  public double area(Object shape) throws NoSuchShapeException {
    if (shape instanceof Square) {
      Square s = (Square)shape;
      return s.side * s.side;
    } else if (shape instanceof Rectangle) {
      Rectangle r = (Rectangle)shape;
      return r.height * r.width;
    } else if (shape instanceof Circle) {
      Circle c = (Circle)shape;
      return PI * c.radius * c.radius;
    }
    throw new NoSuchShapeException();
  }
}
```

```java
public class Square implements Shape {
  private Point topLeft;
  private double side;

  public double area() {
    return side*side;
  }
}

public class Rectangle implements Shape {
  private Point topLeft;
  private double height;
  private double width;
 
  public double area() {
    return height * width;
  }
}

public class Circle implements Shape {
  private Point center;
  private double radius;
  public final double PI = 3.141592653589793;

  public double area() {
    return PI * radius * radius;
  }
}
```


## The Law of Demeter

* A module should not know about the innards of the *objects* it manipulates.

### Train Wreeks

* Chains of calls like *train wreck* are generally considered to be sloppy style and should be avoided.

```java
Options opts = ctxt.getOptions();
File scratchDir = opts.getScratchDir();
final String outputDir = scratchDir.getAbsolutePath();
```

* The use of accessor functions confuses the issue.

```java
final String outputDir = ctxt.options.scratchDir.absolutePath;
```

### Hybrids

* Hybrid structures have functions that do significant things, and they also have either public variables or public accessors and mutators that, for all intents and purposes, make the private variables public, tempting other external functions to use those variables the way a procedural program would use a data structure.

### Hiding Structure

* Becasue objects are supposed to hide their internal structure, we should not be able to navigate through them.
  * e.g., `ctxt.getAbsolutePathOfScratchDirectoryOption();` vs. `ctx.getScratchDirectoryOption().getAbsolutePath()`
  * The first option could lead to an explosion of methods in the `ctxt` object. The second presumes that `getScratchDirectoryOption()` returns a data structure, not an object. Neither option feels good.
  * We see that the intent of getting the absolute path of the scratch directory was to create a scratch file of given name.
    * e.g., `BufferedOutputStream bos = ctxt.createScratchFileStream(classFileName);`
    * This allows `ctxt` to hide its internals and prevents the current function from having to violate the Law of Demeter by navigating through objects it shouldn't know about.

## Data Transfer Objects

* A data transfer object is a class with public variables and no functions.
  * They often become the first in a series of translation stages that convert raw data in a database into objects in the application code.

### Active Record

* Active Records are special forms of DTOs. They are data structures with public variables; but they typically have navigational methods like `save` and `find`.
* We can treat the Active Record as a data sturcture and to create separate objects that contain the business rules and that hide their internal data.

## Conclusion

* “In any given system we will sometimes want the flexibility to add new data types, and so we prefer objects for that part of the system. Other times we will want the flexibility to add new behaviors, and so in that part of the system we prefer data types and procedures. Good software developers understand these issues without prejudice and choose the approach that is best for the job at hand.”
