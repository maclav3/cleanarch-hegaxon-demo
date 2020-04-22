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

The book already has a `Loan(*reader.Reader) error` method, we need an analogous one to return the book, 
with the exception that we don't need the reader context for the `Return() error` method.

Be sure to cover the loaning/returning scenarios thoroughly, remembering about:
1. Inactive readers cannot loan books
1. Only not-loaned books may be loaned; we expect `ErrAlreadyLoaned` otherwise.
1. Only loaned books may be returned; we expect a new error, say `ErrNotLoaned` otherwise.

As you can see, it's a pretty simple domain invented for the demonstration of a concept, but some non-trivial 
domain rules have to be obeyed.

#### The application layer

We are adding just one new use case: a command to return the book.
The command takes just the book ID.

In a real-life system, we would probably need some kind of authentication to ensure that evil Readers don't mark books
as returned without physically returning them to the library. Either a person on the desk, or some kind of automated
system using RFC labels could trigger the command.

Take inspiration from how the `Loan` command handler is implemented:
1. Validate command parameters.
1. Grab the `Book` aggregate from the repository.
1. Call the necessary domain methods.
1. Persist the aggregate.

Don't forget to add the command handler to the `Application.Commands` structure.

A test for the command handler would be nice, but may be superfluous if the remaining layers are well covered.
 
#### The ports layer

The CLI port is somewhat convoluted, but adding a new port should be rather trivial.
Copy the existing `book loan` port, and make sure that you perform the following:
1. Parse the necessary arguments from the command line.
1. Create an application-layer command structure.
1. Call the application command.
1. Return an error or format the OK response.

Don't forget to register the new command within the CLI router.

#### Wiring it all up

In this project, we have a very naïve approach to DI, coding all dependency injection by hand.
However, as the project is still very simple, all that is required is make sure that `Service.App.Commands` has the
`Return Book` command handler instantiated through the constructor that the app layer should be exposing.

### The solution

See solution.patch for a `git` patch that demonstrates a sample solution.
