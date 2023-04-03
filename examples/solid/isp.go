package main

import (
	"context"
	"errors"
)

//Keep interfaces small so that users don’t end up depending on things they don’t need.

type Product struct {
}

type Money struct {
}

// type User interface {
// 	AddToShoppingCart(product Product)
// 	IsLoggedIn() bool
// 	Pay(money Money) error
// 	HasPremium() bool
// 	HasDiscountFor(product Product) bool
// }

type ShoppingCart struct {
}

func (s ShoppingCart) Add(product Product) {}

// type Guest struct {
// 	cart ShoppingCart
// 	//
// 	// some additional fields
// 	//
// }

// func (g *Guest) AddToShoppingCart(product Product) {
// 	g.cart.Add(product)
// }

// func (g *Guest) IsLoggedIn() bool {
// 	return false
// }

// func (g *Guest) Pay(Money) error {
// 	return errors.New("user is not logged in")
// }

// func (g *Guest) HasPremium() bool {
// 	return false
// }

// func (g *Guest) HasDiscountFor(Product) bool {
// 	return false
// }

type Wallet struct{}

func (w *Wallet) Deduct(money Money) error {
	return nil
}

// type NormalCustomer struct {
// 	cart   ShoppingCart
// 	wallet Wallet
// 	//
// 	// some additional fields
// 	//
// }

// func (c *NormalCustomer) AddToShoppingCart(product Product) {
// 	c.cart.Add(product)
// }

// func (c *NormalCustomer) IsLoggedIn() bool {
// 	return true
// }

// func (c *NormalCustomer) Pay(money Money) error {
// 	return c.wallet.Deduct(money)
// }

// func (c *NormalCustomer) HasPremium() bool {
// 	return false
// }

// func (c *NormalCustomer) HasDiscountFor(Product) bool {
// 	return false
// }

type DiscountPolicy struct{}

// type PremiumCustomer struct {
// 	cart     ShoppingCart
// 	wallet   Wallet
// 	policies []DiscountPolicy
// 	//
// 	// some additional fields
// 	//
// }

// func (c *PremiumCustomer) AddToShoppingCart(product Product) {
// 	c.cart.Add(product)
// }

// func (c *PremiumCustomer) IsLoggedIn() bool {
// 	return true
// }

// func (c *PremiumCustomer) Pay(money Money) error {
// 	return c.wallet.Deduct(money)
// }

// func (c *PremiumCustomer) HasPremium() bool {
// 	return true
// }

// func (c *PremiumCustomer) HasDiscountFor(product Product) bool {
// 	// for _, p := range c.policies {
// 	// 	if true {
// 	// 		return true
// 	// 	}
// 	// }

// 	return false
// }

// type UserService struct {
// 	//
// 	// some fields
// 	//
// }

// func (u *UserService) Checkout(ctx context.Context, user User, product Product) error {
// 	if !user.IsLoggedIn() {
// 		return errors.New("user is not logged in")
// 	}

// 	var money Money
// 	//
// 	// some calculation
// 	//
// 	if user.HasDiscountFor(product) {
// 		//
// 		// apply discount
// 		//
// 	}

// 	return user.Pay(money)
// }

/*
So, to have better generalization, we got:

Many structs with unused methods.
Methods that we need somehow to mark so that others do not use them.
Much additional code for unit testing.
Unnatural polymorphism.
…
*/
// Refactoring
type User interface {
	AddToShoppingCart(product Product)
}

type LoggedInUser interface {
	User
	Pay(money Money) error
}

type PremiumUser interface {
	LoggedInUser
	HasDiscountFor(product Product) bool
}

type Guest struct {
	cart ShoppingCart
	//
	// some additional fields
	//
}

func (g *Guest) AddToShoppingCart(product Product) {
	g.cart.Add(product)
}

type NormalCustomer struct {
	cart   ShoppingCart
	wallet Wallet
}

func (c *NormalCustomer) AddToShoppingCart(product Product) {
	c.cart.Add(product)
}

func (c *NormalCustomer) Pay(money Money) error {
	return c.wallet.Deduct(money)
}

type PremiumCustomer struct {
	cart     ShoppingCart
	wallet   Wallet
	policies []DiscountPolicy
	//
	// some additional fields
	//
}

func (c *PremiumCustomer) AddToShoppingCart(product Product) {
	c.cart.Add(product)
}

func (c *PremiumCustomer) Pay(money Money) error {
	return c.wallet.Deduct(money)
}

func (c *PremiumCustomer) HasDiscountFor(product Product) bool {
	// for _, p := range c.policies {
	// 	if p.IsApplicableFor(c, product) {
	// 		return true
	// 	}
	// }

	return false
}

type UserService struct {
	//
	// some fields
	//
}

func (u *UserService) Checkout(ctx context.Context, user User, product Product) error {
	loggedIn, ok := user.(LoggedInUser)
	if !ok {
		return errors.New("user is not logged in")
	}

	var money Money
	// some calculation
	//
	if premium, ok := loggedIn.(PremiumUser); ok && premium.HasDiscountFor(product) {
		//
		// apply discount
		//
	}

	return loggedIn.Pay(money)
}
