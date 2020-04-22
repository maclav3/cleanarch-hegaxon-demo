# Exercise 03 – emit an event if book was returned after due date

*WARNING* – it is required to finish Exercise 02 before attempting this one.

### Introduction

The demo has a loan period `from` and `to` some time, but nothing happens if the book was returned after due date.

Business tells us that we should communicate to the fines department that this happened, they will take care of fining the reader.

### Goal of the exercise

Emit an event `BookReturnedLate` if book was returned after the due date.

After negotiations with the stakeholders, you agreed to the following schema:

```json
{
  "book_id": "string",
  "reader_id": "string",
  "loaned_to": "2009-11-10T23:00:00Z",
  "returned_at": "2009-12-10T23:00:00Z"
}
```

Don't bother using an actual pub/sub implementation. You can use a mock that performs a noop instead.

### Notes

In the real life, we will run into a typical challenge that occurs when persisting an aggregate and emitting 
events.
We need to ensure that the save and sending the event are performed transactionally.

We don't want to save the aggreagate and not send the event, or send the event that does not correspond
to the actual persisted state.

The goal of this exercise is to highlight how dependencies may be injected into the domain layer, so we'll
ignore the difficulties with making the persistence/event store transactionally bound.

Just be aware that this is something that you will need to tackle in real-life applications.


### How to approach it (don't read for a bigger challenge)

#### The domain layer

It's always a good idea to begin with the domain layer, to get a good idea of what we want to achieve.

In this case, we'll need to inject a dependency into the domain layer in the form of a _domain service_.

Define the interface of the event publisher in the domain layer, because we don't want to inject lower layers here.
The implementation of the event publisher will lay in the adapters layer.

Sometimes, when dealing with time-dependent domain logic, it may be a good idea to
mock the time provider to make tests easier. The proposed solution skips that for simplicity.

The easiest way to inject a dependency into a domain method is to just accept is as a parameter, like:
```go
func (b *Book) Return(publisher BookReturnedLatePublisher) error
```

Another way would be to have, for example, a domain service for loaning/returning books, something like:
```go
type loanReturnService struct {
    bookReturnedLatePublisher BookReturnedLatePublisher
}

func (s loanReturnService) LoanBook(b Book, forReader *reader.Reader) error
func (s loanReturnService) Return(b Book) error
```

Take your pick, but the solution presented here uses the first method.

Make sure that you have the domain functionality well-covered with tests.
Use a mock that registers that the event publishing was triggered for the given event, something like

```go
type bookReturnedLatePublisherMock struct {
    capturedEvents []BookPublishedLate
}
```

#### The application layer and wiring it all up

There was a lot of work in the domain layer, but now it's easy to wrap it up.

Write a noop implementation of the publisher in `adapters`.
Make the application command handler for `Return` take the publisher in its constructor, and pass it along 
to the domain method.

Make sure that the dependency injection code injects the proper implementation of the publisher.

### Solution
 
`0001-base-exercise-02.patch` contains the implementation of Exercise 02 that is needed as base for this exercise.
`0002-solution-exercise-03.patch` contains the solution proper.
