# Projectgo

5 backend tasks of instagram API is implemented here 
In Projectgo , create user, get user, create a post
Get a post and get all posts of a user using GET and POST 
Methods. The packages are executed using go and tested using
Postman , by sending multiple GET and POST requests by
Changing the route in the localhost URL.

Mongodb is used to store user and post data in seperate
Databases. 

This project does not use any library other than golang's
Standard library. 

net/http is used for Routing.

Since http is used for Routing ,the Id of user and post 
For a GET request are parsed from the url path given after
/ ,which is Later used for finding the corresponding details. 

"Sync" module along with functions lock.Lock and 
lock.Unlock are used with a time interval of 1 second
So as to maintain the concurrency and safety of the 
Execution thread.

The password of the user is hashed using 'sha1' which
Cannot be reverse engineered by anyone.

