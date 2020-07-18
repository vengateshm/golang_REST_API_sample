# RestApiSampleGo
A simple REST api using Go lang which demonstrates CRUD operations for a Book Repository.

This application uses array data structure in place of a database to hold the book data.  It uses the following http multiplexer https://github.com/gorilla/mux to handle the incoming requests.
It runs on the localhost in port number 8000.

The following are the end points

Get all books - http://localhost:8000/api/v1/books - GET\
Add book - http://localhost:8000/api/v1/books - POST   
Update book - http://localhost:8000/api/v1/books/{id} - POST\
Delete book - http://localhost:8000/api/v1/books/{id} - DELETE 

References

https://youtu.be/_c1b6VFuSTk  
https://golang.org/
