# Building Go Web Applications and Microservices Using Gin

## What is Gin

Gin is a high-performance micro-framework that can be used to build web applications and microservices. It makes it simple to build a request handling pipeline from modular, reusable pieces. It does this by allowing you to write middleware that can be plugged into one or more request handlers or groups of request handlers.

## Why Gin

* built-in net/http library that allows you to create an HTTP server effortlessly.
* a set of commonly used functionalities, e.g. routing, middleware support, rendering, that reduce boilerplate code and make writing web applications simpler.

## Designing the Application

Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response
