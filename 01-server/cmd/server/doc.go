package main

import (
	_ "01-server/docs"
)

// @title           API Server
// @version         1.0
// @description     API with JWT auth
// @host            localhost:8081
// @BasePath        /

// @securityDefinitions.apikey  BearerAuth
// @in                           header
// @name                         Authorization
//
// @tag.name      Sellers
// @tag.description CRUD operations for Sellers
// @tag.name      Auth
// @tag.description Registration and login
//
// @tag.name        Products
// @tag.description Operations for products
// @tag.name        Customers
// @tag.description Operations for customers
// @tag.name        Orders
// @tag.description Operations for orders
