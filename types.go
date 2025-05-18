package main

import "fmt"


/* Go is a statically typed programming language. What that means is the compiler
 * always wants to know what the type is for every value in the program. When the
 * compiler knows the type information ahead of time, it can help to make sure that
 * the program is working with values in a safe way. This helps to reduce potential
 * memory corruption and bugs, and provides the compiler the opportunity to produce
 * more performant code.
 */

 // User-defined types

 /*
  *
  * Go allows you the ability to declare your own types. When you declare a new type, the
  * declaration is constructed to provide the compiler with size and representation information,
  * similar to how the built-in types work. There are two ways to declare a userdefined
  * type in Go. The most common way is to use the keyword struct, which allows
  * you to create a composite type.
  */


/*
 * Struct types are declared by composing a fixed set of unique fields. Each field in a
 * struct is declared with a known type, which could be a built-in type or another userdefined
 * type.
 */
  type user struct {
	name			string
	email			string
	age				int
	is_admin		bool
  }

/*
 *
 * A second way to declare a user-defined type is by taking an existing type and using
 * it as the type specification for the new type. These types are great when you need a
 * new type that can be represented by an existing type. The standard library uses this
 * type declaration to create high-level functionality from the built-in types.
 */

 type unitprice int


// Methods

/* Methods provide a way to add behavior to user-defined types. Methods are really functions
 * that contain an extra parameter that’s declared between the keyword func and
 * the function name.
 */

 func (u user) notify(){
	fmt.Println("New registered user")

	fmt.Printf("User's name is %s\n", u.name)
	fmt.Printf("User's email is %s\n", u.email)
	fmt.Printf("Is admin? %t\n", u.is_admin)
	fmt.Println()
 }

 func (u user) show_privilege(){
	fmt.Printf("%s's new admin privilige: %t\n", u.name, u.is_admin)
 }

 func (u *user) change_privilege(is_admin bool) {
	u.is_admin = is_admin
 }

 func main() {

	var moussa user
	moussa = user{
		name:		"Moussa",
		email:		"moussa@email.com",
		age:		27,
		is_admin:	true,
	}

	mohamed := user{"Mohamed","mohamed@email.com", 28, false}

	//fmt.Printf("User 1: %v\n", moussa)
	//fmt.Printf("User 2: %v\n", mohamed)

	moussa.notify()

	mohamed.notify()
	mohamed.change_privilege(true)
	mohamed.show_privilege()

	var price unitprice
	price = 80000

	fmt.Printf("Price is: %d\n", price)
 }