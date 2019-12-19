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

## Setting Up the Route

In the previous section, we created the route and the route definition in the main.go file itself. As the application grows, it would make sense to move the routes definitions in its own file. We’ll create a function initializeRoutes() in the routes.go file and call this function from the main() function to set up all the routes. Instead of defining the route handler inline, we’ll define them as separate functions.

```go
// routes.go

package main

func initializeRoutes() {
  // Handle the index route
  router.GET("/", showIndexPage)
}
```

## Designing the Article Model

We will keep the article structure simple with just three fields – Id, Title and Content. This can be represented with a struct as follows:

```go
type article struct {
  ID      int    json:"id"
  Title   string json:"title"
  Content string json:"content"
}
```

Most applications will use a database to persist the data. To keep things simple, we will keep the list of articles in memory and will initialize the list with two hard-coded articles as follows:

```go
var articleList = []article{
  article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
  article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}
```

We will place the above code in a new file named models.article.go. At this stage, we need a function that will return the list of all articles. We will name this function getAllArticles() and place it in the same file. We will also write a test for it. This test will be named TestGetAllArticles and will be placed in the models.article_test.go file.

Let’s start by creating the unit test (TestGetAllArticles) for the getAllArticles() function. After creating this unit test, the models.article_test.go file should contain the following code:

```go
// models.article_test.go

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
  alist := getAllArticles()

  // Check that the length of the list of articles returned is the
  // same as the length of the global variable holding the list
  if len(alist) != len(articleList) {
    t.Fail()
  }

  // Check that each member is identical
  for i, v := range alist {
    if v.Content != articleList[i].Content ||
      v.ID != articleList[i].ID ||
      v.Title != articleList[i].Title {

      t.Fail()
      break
    }
  }
}
```

Once we have written the test, we can proceed to write the actual code. The models.article.go file should contain the following code:

```go
// models.article.go

package main

type article struct {
  ID      int    json:"id"
  Title   string json:"title"
  Content string json:"content"
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
  article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
  article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
  return articleList
}
```

## Creating the View Template

We assume that the list of articles will be passed to the template in a variable named payload. With this assumption, the following snippet should show the list of all articles:

```html
  {{range .payload }}
    <!--Create the link for the article based on its ID-->
    <a href="/article/view/{{.ID}}">
      <!--Display the title of the article -->
      <h2>{{.Title}}</h2>
    </a>
    <!--Display the content of the article-->
    <p>{{.Content}}</p>
  {{end}}
```

The above snippet will also link to each article. However, since we have not yet defined route handlers for displaying individual articles, these links won’t work as expected.

## Specifying the Requirement for the Route Handler With a Unit Test

Before we create the handler for the index route, we will create a test to define the expected behavior of this route handler. This test will check for the following conditions:

1. The handler responds with an HTTP status code of 200,
2. The returned HTML contains a title tag containing the text Home Page.

The code for the test will be placed in the TestShowIndexPageUnauthenticated function in the handlers.article_test.go file. We will place helper functions used by this function in the common_test.go file.

```go
// handlers.article_test.go

package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
  r := getRouter(true)

  r.GET("/", showIndexPage)

  // Create a request to send to the above route
  req, _ := http.NewRequest("GET", "/", nil)

  testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
    // Test that the http status code is 200
    statusOK := w.Code == http.StatusOK

    // Test that the page title is "Home Page"
    // You can carry out a lot more detailed tests using libraries that can
    // parse and process HTML pages
    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}
```

```go
The content of common_test.go is as follows:

package main

import (
  "net/http"
  "net/http/httptest"
  "os"
  "testing"

  "github.com/gin-gonic/gin"
)

var tmpArticleList []article

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
  //Set Gin to Test Mode
  gin.SetMode(gin.TestMode)

  // Run the other tests
  os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
  r := gin.Default()
  if withTemplates {
    r.LoadHTMLGlob("templates/*")
  }
  return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

  // Create a response recorder
  w := httptest.NewRecorder()

  // Create the service and process the above request.
  r.ServeHTTP(w, req)

  if !f(w) {
    t.Fail()
  }
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
  tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
  articleList = tmpArticleList
}
```

To implement this test, we have written some helper functions. These will also help us reduce boilerplate code when we write additional tests to test similar functionality.

To check the HTTP code and the returned HTML, we’ll do the following:

1. Create a new router,
2. Define a route to use the same handler that the main app uses (showIndexPage),
3. Create a new request to access this route,
4. Create a function that processes the response to test the HTTP code and HTML, and
5. Call testHTTPResponse() with this new function to complete the test.

## Creating the Route Handler

### Fetches the list of articles

```go
articles := getAllArticles()
```

### Renders the index.html template passing it the article list

```go
c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "index.html",
    // Pass the data that the page uses
    gin.H{
        "title":   "Home Page",
        "payload": articles,
    },
)
```

These are the new files added in this section:

```text
├── common_test.go
├── handlers.article.go
├── handlers.article_test.go
├── models.article.go
├── models.article_test.go
└── routes.go
```

## Setting Up the Route for a Single Article

We can set up a new route to handle requests for a single article in the same manner as in the previous route. However, we need to account for the fact that while the handler for all articles would be the same, the URL for each article would be different. Gin allows us to handle such conditions by defining route parameters as follows:

```go
router.GET("/article/view/:article_id", getArticle)
```

This route will match all requests matching the above path and will store the value of the last part of the route in the route parameter named article_id which we can access in the route handler. For this route, we will define the handler in a function named getArticle.

## Creating the Route Handler for a Single Article

### Extracts the ID of the article to display

```go
c.Param("article_id")
```

where c is the Gin Context which is a parameter to any route handler when using Gin.

### Fetches the article

```go
article, err := getArticleByID(articleID)
```

```go
func getArticleByID(id int) (*article, error) {
  for _, a := range articleList {
    if a.ID == id {
      return &a, nil
    }
  }
  return nil, errors.New("Article not found")
}
```

This function loops through the article list and returns the article whose ID matches the ID passed in. If no matching article is found it returns an error indicating the same.

### Renders the article.html template passing it the article

```go
c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the article.html template
    "article.html",
    // Pass the data that the page uses
    gin.H{
        "title":   article.Title,
        "payload": article,
    },
)
```

The new files added in this section are as follows:

```text
└── templates
    └── article.html
```

## Responding With JSON/XML

In this section, we will refactor the application a bit so that, depending on the request headers, our application can respond in HTML, JSON or XML format.

### Creating a Reusable Function

So far, we’ve been using the HTML method of Gin’s context to render directly from route handlers. This is fine when we always want to render HTML. However, if we want to change the format of the response based on the request, we should refactor this part out into a single function that takes care of rendering. By doing this, we can let the route handler focus on validation and data fetching.

A route handler has to do the same kind of validation, data fetching and data processing irrespective of the desired response format. Once this part is done, this data can be used to generate the response in the desired format. If we need an HTML response, we can pass this data to the HTML template and generate the page. If wee need a JSON response, we can convert this data to JSON and send it back. Likewise for XML.

We’ll create a render function in main.go that will be used by all the route handlers. This function will take care of rendering in the right format based on the request’s Accept header.

In Gin, the Context passed to a route handler contains a field named Request. This field contains the Header field which contains all the request headers. We can use the Get method on Header to extract the Accept header as follows:

```go
// c is the Gin Context
c.Request.Header.Get("Accept")
```

* If this is set to application/json, the function will render JSON,
* If this is set to application/xml, the function will render XML, and
* If this is set to anything else or is empty, the function will render HTML.

The complete render function is as follows:

```go
// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
  switch c.Request.Header.Get("Accept") {
  case "application/json":
    // Respond with JSON
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    // Respond with XML
    c.XML(http.StatusOK, data["payload"])
  default:
    // Respond with HTML
    c.HTML(http.StatusOK, templateName, data)
  }
}
```
