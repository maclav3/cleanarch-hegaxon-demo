# Clean/Hexagonal Architecture demo app

## Introduction

This is a very simple app that demonstrates an architectural style inspired by Clean Architecture and Hexagonal Architecture concepts.

The naming of particular packages may differ from the names established in literature; they reflect what has worked to date in the team I was working in and do not adhere to the original sources religiously.

## The business domain

The app's purpose is management of a library inventory.

*READERS* are clients of the library, and they *loan* *BOOKS*.
A *BOOK* may be *loaned* only for a specific, configurable period of time - _loan period_.

After a *READER* is done reading the book, they must *return* the *BOOK* to the library.

## Exercises

### Exercise 1 – implement the functionality for querying the readers list

[Exercise 1](https://github.com/maclav3/cleanarch-hegaxon-demo/tree/master/_exercises/01-implement-list-readers) covers adding a new functionality for querying the registry of READERs. Changes in the application, adapters and ports are required.


[Exercise 2](more to come)
