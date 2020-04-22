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

### How to approach it (don't read for a bigger challenge)

(TO BE CONTINUED)
