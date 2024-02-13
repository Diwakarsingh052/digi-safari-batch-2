package main

/*
Create a package library which has the following structures

1. `Author` - represents the author of a book. It should have the following fields: `Name` and `Biography`, both of which are of type string.
2. `Book` - represents a book in the library. It should have the following fields:
   - `ID` - a string that represents unique identifier of the book.
   - `Title` - a string that represents the title of the book.
   - `PageCount` - an integer representing the number of pages in the book.
   - `Author` - an `Author` struct representing the author of the book.
   - `IsBorrowed` - a boolean indicating if the book has been borrowed by a library customer.
   - `BorrowerName` - a string representing name of the person who has borrowed the book.

Having these structures in mind, implement the following functionalities:

1. A method for the `Book` struct called `GetDetails`. This will print out all the details of the book including the title, page count, author's name, whether or not it has been borrowed,
   and the name of the person who borrowed it (if applicable).
2. A method for the `Book` struct called `Borrow`. This method will accept a string parameter `borrowerName`.
   If the book is already borrowed, it should print a message indicating that it's already borrowed.
   If it's not borrowed, it should change the `IsBorrowed` status to `true`,
   set the `BorrowerName` to the provided name, and print a message indicating that the book has been successfully borrowed.
3. A method for the `Book` struct called `ReturnBook`. This method will reset the `IsBorrowed` status and clear the `BorrowerName`,
   while printing a message indicating that the book has been returned.
4. A method for the `Book` struct called `GetAuthorBio`. This method will print out the name and biography of the author of the book.

Finally, demonstrate usage of the functionalities in a `main` function by creating an `Author`,
a `Book`, borrowing the book, printing the author's biography,
returning the book, and checking the book's details at various stages.
*/
