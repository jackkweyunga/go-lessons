Most times we will need to structure our project as a collection of modules which we create our selves.

Start by creating another module folder greetings. Within we store go files with greetings functionalities.

We will need to run code from our greetings module in another module.

Let us go on to creating another module folder namely hello. This module will use functions from the greetings module.

#

Since we are using an in-built module, we need to tell go to search for this module from this folder's path using the go edit command.

```cmd
go mod edit -replace example.com/greetings=../greetings 
go mod tidy
```
#
reference: https://go.dev/doc/tutorial/call-module-code