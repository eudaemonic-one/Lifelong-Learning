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

## Creating Reusable Templates

Our application will display a web page using its template. However, there will be several parts such as the header, menu, sidebar, and footer, which will be common across all pages. Go allows us to create reusable template snippets that can be imported in other templates.

The template for the index page makes use of the header and the footer and displays a simple Hello Gin message:

```html
<!--index.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

  <h1>Hello Gin!</h1>

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## Completing and Validating the Setup

Once you have created the templates, it’s time to create the entry file for your application. We’ll create the main.go file for this with the simplest possible web application that will use the index template. We can do this using Gin in four steps:

### Create the router

```go
    router := gin.Default()
```

This creates a router which can be used to define the build of the application.

### Load the templates

```go
router.LoadHTMLGlob("templates/*")
```

This loads all the template files located in the templates folder. Once loaded, these don’t have to be read again on every request making Gin web applications very fast.

### Define the route handler

At the heart of Gin is how you divide the application into various routes and define handlers for each route. We will create a route for the index page and an inline route handler.

```go
router.GET("/", func(c *gin.Context) {
  // Call the HTML method of the Context to render a template
  c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
          "title": "Home Page",
      },
  )
})
```

The router.GET method is used to define a route handler for a GET request. It takes in as parameters the route (/) and one or more route handlers which are just functions.

The route handler has a pointer to the context (gin.Context) as its parameter. This context contains all the information about the request that the handler might need to process it. For example, it includes information about the headers, cookies, etc.

The Context also has methods to render a response in HTML, text, JSON and XML formats. In this case, we use the context.HTML method to render an HTML template (index.html). The call to this method includes additional data in which the value of title is set to Home Page. This is a value that the HTML template can make use of. In this case, we use this value in the \<title\> tag in the header’s template.

### Start the application

```go
router.Run()
```

This starts the application on localhost and serves on the 8080 port by default.

To execute the application from the command line, go to your application directory and execute the following command:

```shell
go build -o app
./app
```

If all goes well, you should be able to access your application at [http://localhost:8080](http://localhost:8080).

The directory structure of your application at this stage should be as follows:

```text
├── main.go
└── templates
    ├── footer.html
    ├── header.html
    ├── index.html
    └── menu.html
```
