# Clean/Hexagonal Architecture demo app

## Introduction

This is a very simple app that demonstrates an architectural style inspired by Clean Architecture and Hexagonal Architecture concepts.

The naming of particular packages may differ from the names established in literature; they reflect what has worked to date in the team I was working in and do not adhere to the original sources religiously.

## The business domain

The app's purpose is management of a library inventory.

*READERS* are clients of the library, and they *loan* *BOOKS*.
A *BOOK* may be *loaned* only for a specific, configurable period of time - _load period_.

After a *READER* is done reading the book, they must *return* the *BOOK* to the library.

_To be continued. Exercises that familiarize with the project structure will appear here._
