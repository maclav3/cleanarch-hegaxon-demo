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

See `_exercises` for examples of exercises to implement on your own.
Proposed solutions are included, but you are encouraged to formulate your own solutions.

Be aware that a lot of real-world complexity is simplified in these exercises.
Their purpose is to provied an example of how to use the Clean-Architeture inspired coding style, not to 
deliver a robust and functional system.
Places where simplifications were made are marked in the code and the solutions.

