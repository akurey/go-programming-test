# Go Programming Test

## Getting started

* Make sure you have Go 1.15 or newer installed
* Download the project into your `$GOPATH`. The project expects to be in `$GOPATH/src/github.com/akurey/go-programming-test`
* Dependencies are managed with `go mod`(https://blog.golang.org/using-go-modules).
* To run the app and test it:
  * Build the service: `go build`
  * Run the service: `PORT={server-port} ./go-programming-test`
  * Upload an order: `curl -F "file=@example_order.csv" http://localhost:{server-port}/orders`
* To view the products
  * `/products`: Lists all products
  * `/products/:id`: Show a single product
* To view the orders
  * `/orders/:id`: Show a single order
* Note: The service is using an in-memory db. If you shut it down, you will lose all data!

## Tech Specs

* Go 1.15+

## Project Overview

This project is created to test your expertice with Go programming language. 
This includes the code writing and thinkering of the solution.

In this test, we will use as scenario a marketplace where a client can buy 
products, and an order is created.

### Vocabulary

Items refer to individual pieces of a product. Items are grouped by product.

An item's important data is:

* Status - available, sold, reserved
* Price sold

A product's important data is:

* Retail price - the original price (per item) at which we sell the product to clients
* Category
* Name

### Tasks

To sell products in the application, the user must upload a csv file with the items to buy. The items must be available or, if the buyer is a premium user, reserved.

You should be able to play around with the app by uploading the CSV file (`example_order.csv`).

The business has asked that we make the following changes:

* We're noticing a decrease of sells in Food products after lunch. As result, we want to sell Food products with a 30% discount if the current time is after 1PM UTC, but it returns to normal at 6AM, to motivate more clients to buy.
* We want to be able to hide the products with the status `sold` from listing all products.
* The products already show enough information for the clients, but as the inventory grows, it's harder for premium user to find the reserved products. We want to implement the following improvements:
  * The premium user can request to list only reserved products.
  * The premium user can sort the products to show the reserved products at top, and then sort by category.
* Right now when the system creates a new order, there is not way to know if the order should include reserved items or not. We want to add a flag to the request to know if the order should include or reject them.

### FAQ

For questions and notifications, please send a email to `aktest@akurey.com`.