# Building Go Web Applications and Microservices Using Gin

## What is Gin

Gin is a high-performance micro-framework that can be used to build web applications and microservices. It makes it simple to build a request handling pipeline from modular, reusable pieces. It does this by allowing you to write middleware that can be plugged into one or more request handlers or groups of request handlers.

## Why Gin

* built-in net/http library that allows you to create an HTTP server effortlessly.
* a set of commonly used functionalities, e.g. routing, middleware support, rendering, that reduce boilerplate code and make writing web applications simpler.

## Designing the Application

Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response

## Application Functionality

The application we’ll build is a simple article manager. This application should:

* Let users register with a username and a password (non-logged in users only),
* Let users log in with a username and a password (non-logged in users only),
* Let users log out (logged in users only),
* Let users create new articles (logged in users only),
* Display the list of all articles on the home page (for all users), and
* Display a single article on its own page (for all users).

In addition to this functionality, the list of articles and a single article should be accessible in the HTML, JSON and XML formats.

To achieve this, we will make use of the following functionalities offered by Gin:

* Routing — to handle various URLs,
* Custom rendering — to handle the response format, and
* Middleware — to implement authentication.

We’ll also write tests to validate that all the features work as intended.
