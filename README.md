# go_api

# Intro
  This is a simple golang web api to do some basic operations such as authentication, get post put delete. 
  I believe it's a good point to start go language. As a Php developer this project showed me how powerful golang technology.
  Especially when you compare Laravel and golang there are some awesome results. Such as some endpoints returns around 5ms 
  which is something impossible to do with Laravel. 

# Auth Layer
  This is a middleware between routing and controller methods. I decided to use JWT to make it more challenging. You may check 
  `routes.go` and see this line of code `SetMiddlewareAuthentication` this is where to decide whether this routing gonna be in 
  auth layer or not.

# DB
  I decided to use mysql in order to make it more like a best practice. There is also migration and seeder which are always called 
  in main.go so whenever the application restarts or run it's gonna drop all tables and their contents then create and seed them 
  all again.
  
  to be continued... 
