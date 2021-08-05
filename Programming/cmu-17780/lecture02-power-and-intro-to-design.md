# Lecture 2 The Power of APIs and Introduction to API Design

## The Power of APIs

* Exponential growth in the power of APIs
* '00s - Data structures, higher-level abstractions, Web APIs: social media, cloud infrastructure
* Enabled code reuse on a grand scale
* Increased the level of abstraction dramatically
* Most of the functionality came from libraries
* APIs confer superpowers

## Characteristics of a good API

* Easy to learn
* Easy to use, even if you take away the documentation
* Hard to misuse
* Easy to read and maintain code that uses it
* Sufficiently powerful to satisfy requirements
* Handles edge and corner cases gracefully
* Easy to evolve
* Appropriate to audience

## A bit of APIs

* The way to learn API design is to design APIs
* Problem:
  * Can you write me an API for a thermometer?
  * It has to support Fahrenheit and Celsius because we'll have users in the US and Eurepe
* Design Choices
  * Store temperature
  * Temperature unit
    * Conversion
  * Changes/Updates
  * Representation
    * Temperature precision / range

```java
// My Thermometer Design
public enum TemperatureUnit {
  Celsius,
  Fahrenheit;
}

public class Thermometer {
  // Constructor.
  Thermometer(double temperature, TemperatureUnit unit);
  
  // Convert current temperature into Fahrenheit scale. No operation would be applied if the current scale is Fahrenheit.
  public void convertToFahrenheit();
  // Convert current temperature into Celsius scale. No operation would be applied if the current scale is Celsius.
  public void convertToCelsius();
  
  // Returns current temperature in given unit.
  public double getTemperature(TemperatureUnit unit);
  // Returns current temperature in default unit.
  public double getTemperature();
  // Returns current temperature unit.
  public TemperatureUnit getTemperatureUnit();
  
  // Sets the new temperature according to given temperature and unit (scale).
  public boolean changeTemperature(double newTemperature, TemperatureUnit unit);
}

// Thermometer Design 1
public class Thermometer {
  public static void setScaleFahrenheit(boolean fahrenheit);
  public static double temperature();
}

// Thermometer Design 1c
public class Thermometer {
  public enum Scale { FAHRENHEIT, CELSIUS }
  public static void setScale(Scale scale);
  public static double temperature();
}

// Thermometer Design 1d
public class Thermometer {
  public enum Scale { FAHRENHEIT, CELSIUS }
  public static double temperature(Scale scale);
}
// it's a pseudo-singleton
// can't test without sensor hardware to mock

// Thermometer Design 2
public class Thermometer {
  public enum Scale { FAHRENHEIT, CELSIUS }
  public Thermometer();
  public double temperature(Scale scale);
}

// Thermometer Design 3
public interface Thermometer {
  enum Scale { FAHRENHEIT, CELSIUS }
  double temperature(Scale scale);
}

public class Thermometer {
  // Returns "standard thermometer." Doc must say what that is!
  public Thermometer getThermometer(); // Static factory
}

// Thermometer Design 3b
public interface Thermometer {
  enum Scale { FAHRENHEIT, CELSIUS }
  double temperature();
  double temperature(Scale scale);
  Scale defaultScale();
}

// To invert control, use the Observer pattern
public interface Thermometer {
  enum Scale { FAHRENHEIT, CELSIUS }
  interface TemperatureChangeListener {
    void temperatureChange(double newTemp, Scale scale);
  }
  void addListener(TemperatureChangeListener listener);
  // Returns true if the listener was represent.
  boolean removeListener(TemperatureChangeListener listener);
  double temperature();
  double temperature(Scale scale);
  Scale defaultScale();
}
```

* Combinatorial explosion of possible designs
  * Shape: Utility class vs. true singleton vs. final concrete class vs. non-final concrete class vs. abstract class vs. interface
  * Scale specification: boolean vs. method clevage. vs. enum
  * setScale method vs. passing scale in to temperature
  * Whether to support a default scale
  * Traditional API vs. callbacks
* **Moral: even the simplest API designs are complex**