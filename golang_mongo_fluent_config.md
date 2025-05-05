# 🚀 Golang + MongoDB-style Fluent Configuration Example

This Go program mimics how MongoDB drivers (like `mongo-go-driver`) allow configuration using a **fluent-style interface**.

Let’s break it down step by step:

---

## 📦 `package main`

The entry point for the Go program — everything begins from here.

---

## 📥 Imports

```go
import "fmt"
```

Brings in the fmt package, used here for printing to the console.

---

## 🧱 Struct Definition

```go
type ClientOptions struct {
	URI string
}
```

Defines a struct called ClientOptions, with one field: URI.

Think of this like a configuration object (just like MongoDB client options).

---

## 🔧 Function: NewClientOptions

```go
func NewClientOptions() *ClientOptions {
	return &ClientOptions{}
}
```

This function creates a new ClientOptions struct and returns a pointer to it.

It’s like a constructor: options.Client() in the MongoDB Go driver.

---

## 🔁 Method: ApplyURI

```go
func (co *ClientOptions) ApplyURI(uri string) *ClientOptions {
	co.URI = uri
	return co
}
```

A method with a pointer receiver (*ClientOptions), so it modifies the actual object.

It sets the URI field.

Then it returns the same pointer (co), allowing method chaining (fluent interface).

---

## 🧪 In main()

```go
clientOptions := NewClientOptions().ApplyURI("mongodb://localhost:27017")
```

Calls NewClientOptions() → returns a *ClientOptions

Calls .ApplyURI(...) on the returned object → sets the URI

The result is stored in clientOptions, which now has URI = "mongodb://localhost:27017"

---

## 🖨️ Output

```go
fmt.Println("Mongo URI:", clientOptions.URI)
```

Outputs:

```
Mongo URI: mongodb://localhost:27017
```

---

## ✅ Summary

This code:

- Creates a struct (ClientOptions)
- Adds a constructor-like function (NewClientOptions)
- Adds a method (ApplyURI) to set values in a chainable way
- Demonstrates how method chaining makes config building clean and readable

---

## ❓ Why Use NewClientOptions() Instead of &ClientOptions{}?

You're right — you could just write:

```go
options := &ClientOptions{}
```

But using a constructor-like function gives you many benefits:

---

### ✅ 1. Encapsulation / Abstraction

You can hide internal logic, set defaults, or log actions:

```go
func NewClientOptions() *ClientOptions {
	fmt.Println("Creating new ClientOptions")
	return &ClientOptions{URI: "mongodb://localhost:27017"}  // Set defaults
}
```

---

### ✅ 2. Consistency with APIs

- Matches the style of popular libraries like the MongoDB Go driver
- Makes your config code fluent and uniform

---

### ✅ 3. Future-proofing

You can easily add logic like reading from environment variables:

```go
func NewClientOptions() *ClientOptions {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	return &ClientOptions{URI: uri}
}
```

Update the logic once, use it everywhere.

---

### ✅ 4. Improves Readability / Intent

This:

```go
options := NewClientOptions()
```

Is more expressive than:

```go
options := &ClientOptions{}
```

---

## 🧠 Summary Comparison

| Approach            | Pros                            | Cons                                |
|---------------------|----------------------------------|-------------------------------------|
| `&ClientOptions{}`  | Simple, direct                   | Harder to maintain if struct grows  |
| `NewClientOptions()`| Encapsulated, extensible, idiomatic | Slightly more verbose at first     |
