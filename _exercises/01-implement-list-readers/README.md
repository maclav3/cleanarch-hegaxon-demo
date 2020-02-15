# Exercise 01 – implement the functionality for querying the readers list

### Introduction

The demo has implemented commands for adding and activating/deactivating readers.

We've been told that we need a way to list all current users, and optionally filter them by their activated/deactivated state.

### Goal of the exercise

Implement the functionality of listing existing readers, end-to-end.

### How to approach it (don't read for a bigger challenge)

The implementation of this feature will span almost all layers of our application.

In practice, often it is the easiest to begin with the application layer, because it defines the use cases, and the implementation
of particular adapters or ports may come later.

In this exercise, there will probably be no major changes to the domain logic.
Listing all readers is not part of the interface of the domain repository `domain/reader/repository.go`

#### The application layer

There is a stub query handler in `app/query/reader` that handles a `ListQuery`.

`ListQuery` is empty for now, but we need a way to filter activated/deactivated users.

Change `ListQuery` to reflect that and implement calling the injected repository, and handling the response/errors.

Don't worry about implementing the `ListReaders` method for now, we can take care of that in the next step.

#### The adapters layer

The `ListQuery` query handler in the application layer already accepts a repository implementation that has a `ListReaders` method.

We need to implement the method in `adapters/reader:MemoryRepository`, taking into account that the new parameter in `ListQuery`
must be handled correctly.

We may write a test for the new functionality, but we could also test the QueryHandler and see that it works as expected for any
implementation of the repository. Currently, we only have the memory repo, but if we decide to replace it in the future,
for example, with a MySQL repo, the test will still apply, regardless of the repository implementation.

#### The ports layer

Currently, the only I/O that we have is the CLI router.
Add a cobra command that prepares `ListQuery` from cli args, and renders a list of readers to the output buffer.

You can base your implementation on the command listing books.

The advantage of clean architecture in this case is that the I/O code is totally separated from the concerns of logging
or persistence. You could develop an HTTP port, for example, allowing I/O via a web browser, and most other code would stay untouched.

#### A test

It is always a good idea to test a user story-sized unit of work with some kind of an end-to-end test.
Of course, in addition to unit tests for chunks of code that are important, and well unit-testable.

Check `app/query/book/list_test.go` for an example of how to test an application query.

In-memory repositories are nice in that they enable testing application commands/queries without the need to run any external
dependencies, for example in a docker container.

However, sometimes it is more relevant to use an actual production repository, and maintaining corresponding in-memory
and production repositories rarely works out well in the long run.

### The solution

See solution.patch for a `git` patch that demonstrates a sample solution.