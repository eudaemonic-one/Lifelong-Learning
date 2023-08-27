# Chapter 2. Data Models and Query Languages

> Each layer hides the complexity of the layers below it by providing a clean data model.

## Relational Model Versus Document Model

SQL based on relational model: data is organized into *relations* (called *tables* in SQL), where each relation is an unordered collection of *tuples* (*rows* in SQL).

NoSQL: greater scalability, specialized query operations, dynamic and expressive data model.

### The Object-Relational Mismatch

If data is stored in relational tables, an awkward translation layer is required between the objects in the application code and the database model of tables.

Object-relational mapping (ORM) could reduce boilerplate.

JSON representation makes one-to-many relationships or say tree structure explicit.

### Many-to-One and Many-to-Many Relationships

Normalization vs. De-normalization

Normalizing data requires many-to-one relationships. The relationship indicates the need of joins. In document databases, the work of emulating joins is shifted from database to application code; the support for joins is often weak.

> Data has a tendency of becoming more interconnected as features are added to applications.

## Relational Versus Document Databases Today

Document data model are schema flexible, better performance due to locality, and closer to data structures used by the application. Relational model provides better support for joins, and many-to-one and many-to-many relationships.

For highly interconnected data, the document model is awkward, the relational model is acceptable, and graph models are the most natural.

Document databases are *schema-on-read* while the traditional relational databases are *schema-on-write*.

The locality advantage only applies if you need large parts of the document at the same time.

## Query Languages for Data

Declarative Query Language vs. Imperative Query Code/APIs

SQL is a *declarative query language*, which is up to the database system's query optimizer to decide what indexes and which join methods to use, and in which order to execute various parts of the query.

MapReduce is a low-level programming model for distributed execution on a cluster of machines, neither a declarative query language nor a fully imperative query API.

## Graph-Like Data Models

Graph == *vertices* + *edges*

All vertices in a graph may represent the same kind of thing. However, graphs are not limited to such *homogeneous* data.

For example, Facebook maintains a single graph with many different types of vertices and edges: vertices represent people, locations, events, comments made by users; edges indicate which people are friends with each other other, who commented on which post, who attended which event.

### Property Graph

Graph store == two relational tables (vertices, edges)

Any vertex can have an edge connecting it with any other vertex.

Given any vertex, you can efficiently find both its incoming and outgoing edges, and thus traverse the graph.

By using different labels for different kinds of relationships, you can store several different kinds of information iin a single graph, while still maintaining a clean data model.

### The Cypher Query Language

Cypher is a declarative query language for property graphs, created for the Neo4j graph database.

```text
CREATE
  (NAmerica:Location {name:'North America', type:'continent'}),
  (USA:Location      {name:'United States', type:'country'  }),
  (Idaho:Location    {name:'Idaho',         type:'state'    }),
  (Lucy:Person       {name:'Lucy' }),
  (Idaho) -[:WITHIN]->  (USA)  -[:WITHIN]-> (NAmerica),
  (Lucy)  -[:BORN_IN]-> (Idaho)

MATCH
  (person) -[:BORN_IN]->  () -[:WITHIN*0..]-> (us:Location {name:'United States'}),
  (person) -[:LIVES_IN]-> () -[:WITHIN*0..]-> (eu:Location {name:'Europe'})
RETURN person.name
```

### Graph Queries in SQL

In a relational database, you usually know in advance which joins you need in your query. In a graph query, the number of joins is not fixed in advance.

### Triple-Stores and SPARQL

In a triple-store, all information is stored in the form of: (subject, predicate, object). The predicate and object of the triple are equivalent to the key and value of a property on the subject vertex. The predicate is an edge in the graph.

```text
@prefix : <urn:example:>.
_:lucy     a       :Person.
_:lucy     :name   "Lucy".
_:lucy     :bornIn _:idaho.
_:idaho    a       :Location.
_:idaho    :name   "Idaho".
_:idaho    :type   "state".
_:idaho    :within _:usa.
_:usa      a       :Location.
_:usa      :name   "United States".
_:usa      :type   "country".
_:usa      :within _:namerica.
_:namerica a       :Location.
_:namerica :name   "North America".
_:namerica :type   "continent".
```

SPARQL is a query language for triple-stores using RDF data model.

```text
PREFIX : <urn:example:>

SELECT ?personName WHERE {
  ?person :name ?personName.
  ?person :bornIn  / :within* / :name "United States".
  ?person :livesIn / :within* / :name "Europe".
}
```

## Summary

big tree (hierarchical model) -> many-to-many relationships (relational model) -> one-to-many (document databases) -> anything-potentially-related-to-everything (graph databases)

document and graph databases don't enforce a schema for the data they store
