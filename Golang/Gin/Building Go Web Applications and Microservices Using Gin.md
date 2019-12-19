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

## Routing

In our application, we will:

* Serve the index page at route / (HTTP GET request),
Group user-related routes under the /u route,
* Serve the login page at /u/login (HTTP GET request),
  * Process the login credentials at /u/login (HTTP POST request),
  * Log out at /u/logout (HTTP GET request),
  * Serve the registration page at /u/register (HTTP GET request),
  * Process the registration information at /u/register (HTTP POST request) ,
* Group article related routes under the /article route,
  * Serve the article creation page at /article/create (HTTP GET request),
  * Process the submitted article at /article/create (HTTP POST request), and
  * Serve the article page at /article/view/:article_id (HTTP GET request). Take note of the :article_id part in this route. The : at the beginning indicates that this is a dynamic route. This means that :article_id can contain any value and Gin will make this value available in the route handler.

## Rendering

A web application can render a response in various formats like HTML, text, JSON, XML or other formats. API endpoints and microservices typically respond with data, commonly in JSON format but also in any other desired format.

## Middleware

In the context of a Go web application, middleware is a piece of code that can be executed at any stage while handling an HTTP request. It is typically used to encapsulate common functionality that you want to apply to multiple routes. We can use middleware before and/or after an HTTP request is handled. Some common uses of middleware include authorization, validation, etc.
