# Exercise 02 – implement the functionality for querying the readers list

### Introduction

The demo has a command to loan a book, but there is no functionality to return it.

Business tells us that we should be able to return books, with the following domain rules:
1. We don't care who returned the book, and we don't register the fact.
1. Only a loaned book may be returned.

### Goal of the exercise

Implement the functionality of returning loaned books, end-to-end.

### How to approach it (don't read for a bigger challenge)

#### The domain layer

It is a good idea to start implementation of new domain logic from the domain layer.
We can quickly verify our grasp of the idea, cover the business cases with unit tests.


