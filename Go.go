// Basic Syntax

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func main() { // Main function, entry point of the program
    fmt.Println("Hello, World!") // Prints "Hello, World!" to the console
}
// Variables and Constants

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func main() { // Main function, entry point of the program
    var a int = 10 // Declares an integer variable 'a' and initializes it with 10
    b := 20 // Declares an integer variable 'b' and initializes it with 20 (short variable declaration)
    const c = "constant" // Declares a constant 'c' with value "constant"

    fmt.Println("a:", a) // Prints the value of 'a' to the console
    fmt.Println("b:", b) // Prints the value of 'b' to the console
    fmt.Println("c:", c) // Prints the value of 'c' to the console
}
// Control Structures

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func main() { // Main function, entry point of the program
    // If-Else
    age := 18 // Declares an integer variable 'age' and initializes it with 18
    if age >= 18 { // Checks if 'age' is greater than or equal to 18
        fmt.Println("Adult") // Prints "Adult" if the condition is true
    } else { // Executes if the condition is false
        fmt.Println("Minor") // Prints "Minor" if the condition is false
    }

    // For Loop
    for i := 0; i < 5; i++ { // Loop from 0 to 4
        fmt.Println("Iteration", i) // Prints the current iteration number
    }

    // Switch Case
    day := 2 // Declares an integer variable 'day' and initializes it with 2
    switch day { // Switch statement on 'day'
    case 1: // Case when 'day' equals 1
        fmt.Println("Monday") // Prints "Monday"
    case 2: // Case when 'day' equals 2
        fmt.Println("Tuesday") // Prints "Tuesday"
    default: // Executes if no case matches
        fmt.Println("Other day") // Prints "Other day"
    }
}
// Arrays and Slices

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func main() { // Main function, entry point of the program
    // Array
    var arr [3]int = [3]int{1, 2, 3} // Declares an array of 3 integers and initializes it with 1, 2, 3
    fmt.Println("Array:", arr) // Prints the array

    // Slice
    slice := []int{4, 5, 6} // Declares a slice of integers and initializes it with 4, 5, 6
    slice = append(slice, 7) // Appends the integer 7 to the slice
    fmt.Println("Slice:", slice) // Prints the slice
}
// Maps

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func main() { // Main function, entry point of the program
    // Map
    m := make(map[string]int) // Creates a map with string keys and integer values
    m["one"] = 1 // Sets the value for the key "one" to 1
    m["two"] = 2 // Sets the value for the key "two" to 2

    for key, value := range m { // Iterates over the map
        fmt.Println(key, ":", value) // Prints each key-value pair
    }
}
// Functions

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Function that adds two integers
func add(a int, b int) int {
    return a + b // Returns the sum of a and b
}

// Function with named return value
func multiply(a int, b int) (result int) {
    result = a * b // Sets the named return value
    return // Returns result
}

func main() { // Main function, entry point of the program
    sum := add(5, 10) // Calls add function with arguments 5 and 10
    fmt.Println("Sum:", sum) // Prints the result of addition

    product := multiply(4, 6) // Calls multiply function with arguments 4 and 6
    fmt.Println("Product:", product) // Prints the result of multiplication
}
// Methods

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a struct type
type Person struct {
    Name string
    Age  int
}

// Method associated with Person type
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name // Returns a greeting message
}

func main() { // Main function, entry point of the program
    person := Person{Name: "Alice", Age: 30} // Creates a Person instance
    fmt.Println(person.Greet()) // Calls the Greet method and prints the result
}
// Packages

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "myapp/mathutils" // Imports the custom package 'mathutils'
)

func main() { // Main function, entry point of the program
    result := mathutils.Add(3, 7) // Calls Add function from the mathutils package
    fmt.Println("Result of addition:", result) // Prints the result
}
// Custom Package: mathutils/add.go

package mathutils // Declares the package name

// Function to add two integers
func Add(a int, b int) int {
    return a + b // Returns the sum of a and b
}
// Error Handling

package main // Declares the main package

import (
    "errors" // Imports the errors package for error handling
    "fmt" // Imports the fmt package for formatted I/O
)

// Function that returns an error if the divisor is zero
func divide(a int, b int) (int, error) {
    if b == 0 { // Checks if the divisor is zero
        return 0, errors.New("cannot divide by zero") // Returns an error
    }
    return a / b, nil // Returns the result of division and no error
}

func main() { // Main function, entry point of the program
    result, err := divide(10, 0) // Calls divide function with a divisor of zero
    if err != nil { // Checks if there was an error
        fmt.Println("Error:", err) // Prints the error
    } else {
        fmt.Println("Result:", result) // Prints the result if no error
    }
}
// Goroutines

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for sleeping
)

// Function that prints numbers
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i) // Prints the number
        time.Sleep(time.Millisecond * 500) // Sleeps for 500 milliseconds
    }
}

// Function that prints letters
func printLetters() {
    for _, letter := range "ABCDE" {
        fmt.Println(string(letter)) // Prints the letter
        time.Sleep(time.Millisecond * 500) // Sleeps for 500 milliseconds
    }
}

func main() { // Main function, entry point of the program
    go printNumbers() // Starts printNumbers as a goroutine
    go printLetters() // Starts printLetters as a goroutine

    time.Sleep(time.Second * 3) // Sleeps for 3 seconds to allow goroutines to finish
}
// Channels

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for sleeping
)

func sender(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Sends value to the channel
        time.Sleep(time.Millisecond * 500) // Sleeps for 500 milliseconds
    }
    close(ch) // Closes the channel
}

func receiver(ch <-chan int) {
    for value := range ch { // Receives values from the channel
        fmt.Println(value) // Prints the received value
    }
}

func main() { // Main function, entry point of the program
    ch := make(chan int) // Creates a new channel

    go sender(ch) // Starts sender function as a goroutine
    receiver(ch) // Starts receiver function

    // No need to explicitly wait for goroutines as receiver will handle it
}
// Select Statement

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for sleeping
)

func main() { // Main function, entry point of the program
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() { // Anonymous goroutine for sending to ch1
        time.Sleep(time.Second * 1)
        ch1 <- "Message from ch1"
    }()

    go func() { // Anonymous goroutine for sending to ch2
        time.Sleep(time.Second * 2)
        ch2 <- "Message from ch2"
    }()

    select {
    case msg1 := <-ch1: // Receives from ch1
        fmt.Println("Received", msg1)
    case msg2 := <-ch2: // Receives from ch2
        fmt.Println("Received", msg2)
    case <-time.After(time.Second * 3): // Times out after 3 seconds
        fmt.Println("Timeout")
    }
}
// Mutexes

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

var (
    count int
    mu    sync.Mutex // Creates a new mutex
)

func increment(wg *sync.WaitGroup) {
    mu.Lock() // Acquires the lock
    count++   // Increments the count
    mu.Unlock() // Releases the lock
    wg.Done() // Signals that this goroutine is done
}

func main() { // Main function, entry point of the program
    var wg sync.WaitGroup // Creates a new WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1) // Adds a count to the WaitGroup
        go increment(&wg) // Starts increment function as a goroutine
    }

    wg.Wait() // Waits for all goroutines to finish
    fmt.Println("Final count:", count) // Prints the final count
}
// Channels with Buffering

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

func main() { // Main function, entry point of the program
    ch := make(chan int, 2) // Creates a buffered channel with a capacity of 2

    ch <- 1 // Sends value 1 to the channel
    ch <- 2 // Sends value 2 to the channel

    // No blocking occurs here since the channel is buffered

    fmt.Println(<-ch) // Receives and prints the first value from the channel
    fmt.Println(<-ch) // Receives and prints the second value from the channel
}
// WaitGroup

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Signals that this goroutine is done
    fmt.Printf("Worker %d starting\n", id)
    // Simulate work
    // time.Sleep(time.Second * time.Duration(id)) // Uncomment to see delays
    fmt.Printf("Worker %d done\n", id)
}

func main() { // Main function, entry point of the program
    var wg sync.WaitGroup // Creates a new WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Adds a count to the WaitGroup
        go worker(i, &wg) // Starts worker function as a goroutine
    }

    wg.Wait() // Waits for all goroutines to finish
}
// Interfaces

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

// Define an interface
type Speaker interface {
    Speak() string
}

// Implement the interface with a type
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hi, my name is " + p.Name
}

func main() { // Main function, entry point of the program
    var speaker Speaker // Declares a variable of type Speaker
    speaker = Person{Name: "Bob"} // Assigns a Person instance to the speaker
    fmt.Println(speaker.Speak()) // Calls Speak method and prints the result
}
// Reflection

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "reflect" // Imports the reflect package for reflection
)

func main() { // Main function, entry point of the program
    x := 42 // An integer value
    v := reflect.ValueOf(x) // Gets the reflect.Value of x
    fmt.Println("Type:", v.Type()) // Prints the type of x
    fmt.Println("Kind:", v.Kind()) // Prints the kind of x
    fmt.Println("Value:", v.Interface()) // Prints the value of x
}
// Error Handling with Custom Errors

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

// Custom error type
type DivisionError struct {
    Divisor int
}

func (e *DivisionError) Error() string {
    return fmt.Sprintf("cannot divide by zero, divisor: %d", e.Divisor)
}

func divide(a int, b int) (int, error) {
    if b == 0 { // Checks if the divisor is zero
        return 0, &DivisionError{Divisor: b} // Returns a custom error
    }
    return a / b, nil // Returns the result of division and no error
}

func main() { // Main function, entry point of the program
    result, err := divide(10, 0) // Calls divide function with a divisor of zero
    if err != nil { // Checks if there was an error
        fmt.Println("Error:", err) // Prints the error
    } else {
        fmt.Println("Result:", result) // Prints the result if no error
    }
}
// Structs and Methods

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a struct type
type Rectangle struct {
    Width  int
    Height int
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() int {
    return r.Width * r.Height // Returns the area
}

// Method to calculate the perimeter of the rectangle
func (r Rectangle) Perimeter() int {
    return 2 * (r.Width + r.Height) // Returns the perimeter
}

func main() { // Main function, entry point of the program
    rect := Rectangle{Width: 10, Height: 5} // Creates a Rectangle instance

    fmt.Println("Area:", rect.Area()) // Prints the area of the rectangle
    fmt.Println("Perimeter:", rect.Perimeter()) // Prints the perimeter of the rectangle
}
// Embedded Structs

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a base struct type
type Animal struct {
    Name string
}

// Method for Animal
func (a Animal) Speak() string {
    return "Animal sound" // Returns a generic animal sound
}

// Define a derived struct type that embeds Animal
type Dog struct {
    Animal // Embedded field
    Breed  string
}

// Method for Dog that overrides Animal's method
func (d Dog) Speak() string {
    return "Woof" // Returns a dog-specific sound
}

func main() { // Main function, entry point of the program
    dog := Dog{Animal: Animal{Name: "Rex"}, Breed: "Golden Retriever"} // Creates a Dog instance

    fmt.Println("Name:", dog.Name) // Prints the name from the embedded Animal struct
    fmt.Println("Breed:", dog.Breed) // Prints the breed of the dog
    fmt.Println("Speak:", dog.Speak()) // Calls the Speak method of Dog
}
// Pointers

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func increment(value *int) { // Function takes a pointer to an integer
    *value++ // Dereferences the pointer to increment the value
}

func main() { // Main function, entry point of the program
    number := 5 // Declares an integer variable
    increment(&number) // Calls increment with the address of number
    fmt.Println("Incremented number:", number) // Prints the incremented value
}
// Variadic Functions

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Variadic function that sums integers
func sum(numbers ...int) int {
    total := 0 // Initialize total
    for _, number := range numbers { // Iterates over each number
        total += number // Adds the number to the total
    }
    return total // Returns the total sum
}

func main() { // Main function, entry point of the program
    result := sum(1, 2, 3, 4, 5) // Calls sum with multiple arguments
    fmt.Println("Sum:", result) // Prints the result of the sum
}
// Defer Statement

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

func example() { // Function that demonstrates the use of defer
    defer fmt.Println("Deferred") // This statement will execute after the surrounding function returns
    fmt.Println("Immediate") // This statement will execute first
}

func main() { // Main function, entry point of the program
    example() // Calls the example function
}
// JSON Encoding and Decoding

package main // Declares the main package

import (
    "encoding/json" // Imports the encoding/json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
)

type Person struct { // Define a struct type
    Name string `json:"name"` // JSON field name is "name"
    Age  int    `json:"age"`  // JSON field name is "age"
}

func main() { // Main function, entry point of the program
    // Create a new Person instance
    person := Person{Name: "Alice", Age: 30}

    // Encode the Person instance to JSON
    jsonData, _ := json.Marshal(person)
    fmt.Println("JSON:", string(jsonData)) // Prints the JSON representation

    // Decode JSON back to Person instance
    var decodedPerson Person
    json.Unmarshal(jsonData, &decodedPerson)
    fmt.Println("Decoded Person:", decodedPerson) // Prints the decoded Person instance
}
// File I/O

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "os" // Imports the os package for file handling
)

func main() { // Main function, entry point of the program
    // Create a new file
    file, err := os.Create("example.txt")
    if err != nil { // Checks for file creation error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }
    defer file.Close() // Ensures the file is closed after the function returns

    // Write data to the file
    _, err = file.WriteString("Hello, file I/O!")
    if err != nil { // Checks for write error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("File written successfully") // Confirms successful file write

    // Open the file for reading
    file, err = os.Open("example.txt")
    if err != nil { // Checks for file open error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }
    defer file.Close() // Ensures the file is closed after the function returns

    // Read data from the file
    buf := make([]byte, 100) // Creates a buffer to hold file contents
    _, err = file.Read(buf)
    if err != nil { // Checks for read error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("File content:", string(buf)) // Prints the content read from the file
}
// File I/O

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "os" // Imports the os package for file handling
)

func main() { // Main function, entry point of the program
    // Create a new file
    file, err := os.Create("example.txt")
    if err != nil { // Checks for file creation error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }
    defer file.Close() // Ensures the file is closed after the function returns

    // Write data to the file
    _, err = file.WriteString("Hello, file I/O!")
    if err != nil { // Checks for write error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("File written successfully") // Confirms successful file write

    // Open the file for reading
    file, err = os.Open("example.txt")
    if err != nil { // Checks for file open error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }
    defer file.Close() // Ensures the file is closed after the function returns

    // Read data from the file
    buf := make([]byte, 100) // Creates a buffer to hold file contents
    _, err = file.Read(buf)
    if err != nil { // Checks for read error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("File content:", string(buf)) // Prints the content read from the file
}
// HTTP Server

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "net/http" // Imports the net/http package for HTTP handling
)

// Handler function for HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!") // Writes response to the client
}

func main() { // Main function, entry point of the program
    http.HandleFunc("/", handler) // Registers handler function for root URL
    fmt.Println("Server starting on port 8080") // Prints a message indicating server start
    http.ListenAndServe(":8080", nil) // Starts the HTTP server on port 8080
}
// HTTP Client

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "io/ioutil" // Imports the ioutil package for reading response body
    "net/http" // Imports the net/http package for HTTP handling
)

func main() { // Main function, entry point of the program
    response, err := http.Get("http://example.com") // Sends a GET request to example.com
    if err != nil { // Checks for request error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }
    defer response.Body.Close() // Ensures the response body is closed after use

    body, err := ioutil.ReadAll(response.Body) // Reads the response body
    if err != nil { // Checks for reading error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("Response body:", string(body)) // Prints the response body
}
// Struct Tags

package main // Declares the main package

import (
    "encoding/json" // Imports the encoding/json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
)

// Define a struct type with tags
type Person struct {
    FirstName string `json:"first_name"` // JSON field name is "first_name"
    LastName  string `json:"last_name"`  // JSON field name is "last_name"
}

func main() { // Main function, entry point of the program
    person := Person{FirstName: "John", LastName: "Doe"} // Creates a Person instance

    jsonData, _ := json.Marshal(person) // Encodes the Person instance to JSON
    fmt.Println("JSON:", string(jsonData)) // Prints the JSON representation
}
// Time Handling

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func main() { // Main function, entry point of the program
    currentTime := time.Now() // Gets the current time
    fmt.Println("Current time:", currentTime) // Prints the current time

    formattedTime := currentTime.Format("2006-01-02 15:04:05") // Formats the time
    fmt.Println("Formatted time:", formattedTime) // Prints the formatted time

    duration := time.Duration(5 * time.Second) // Creates a duration of 5 seconds
    time.Sleep(duration) // Sleeps for the duration
    fmt.Println("Woke up after 5 seconds") // Prints a message after sleeping
}
// Command Line Arguments

package main // Declares the main package

import (
    "flag" // Imports the flag package for command-line flags
    "fmt" // Imports the fmt package for formatted I/O
)

func main() { // Main function, entry point of the program
    name := flag.String("name", "World", "a name to say hello to") // Defines a string flag
    flag.Parse() // Parses the command-line flags

    fmt.Println("Hello,", *name) // Prints the greeting with the provided name
}
// JSON to Struct

package main // Declares the main package

import (
    "encoding/json" // Imports the encoding/json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
)

type Person struct { // Define a struct type
    Name string `json:"name"` // JSON field name is "name"
    Age  int    `json:"age"`  // JSON field name is "age"
}

func main() { // Main function, entry point of the program
    jsonStr := `{"name":"Alice","age":30}` // JSON string
    var person Person // Declares a variable of type Person

    err := json.Unmarshal([]byte(jsonStr), &person) // Decodes JSON string into Person
    if err != nil { // Checks for decoding error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("Person:", person) // Prints the decoded Person instance
}
// HTML Template

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "html/template" // Imports the html/template package for HTML templates
    "net/http" // Imports the net/http package for HTTP handling
)

// HTML template string
const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Heading}}</h1>
    <p>{{.Body}}</p>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title   string
        Heading string
        Body    string
    }{
        Title:   "Template Example",
        Heading: "Hello, World!",
        Body:    "This is a simple HTML template example.",
    }

    t := template.Must(template.New("webpage").Parse(tpl)) // Parses the template
    t.Execute(w, data) // Executes the template with data
}

func main() { // Main function, entry point of the program
    http.HandleFunc("/", handler) // Registers handler function for root URL
    fmt.Println("Server starting on port 8080") // Prints a message indicating server start
    http.ListenAndServe(":8080", nil) // Starts the HTTP server on port 8080
}
// Regular Expressions

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "regexp" // Imports the regexp package for regular expressions
)

func main() { // Main function, entry point of the program
    pattern := `\d+` // Regular expression pattern to match digits
    re := regexp.MustCompile(pattern) // Compiles the regular expression

    text := "There are 123 apples and 456 oranges." // Sample text
    matches := re.FindAllString(text, -1) // Finds all matches in the text

    fmt.Println("Matches:", matches) // Prints the matches
}
// JSON Encoding with Struct Tags

package main // Declares the main package

import (
    "encoding/json" // Imports the encoding/json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
)

// Define a struct type with JSON tags
type Person struct {
    FirstName string `json:"first_name"` // JSON field name is "first_name"
    LastName  string `json:"last_name"`  // JSON field name is "last_name"
}

func main() { // Main function, entry point of the program
    person := Person{FirstName: "Jane", LastName: "Doe"} // Creates a Person instance

    jsonData, _ := json.Marshal(person) // Encodes the Person instance to JSON
    fmt.Println("JSON:", string(jsonData)) // Prints the JSON representation
}
// Goroutine Synchronization

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Signals that this goroutine is done
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // Simulates work with sleep
    fmt.Printf("Worker %d done\n", id)
}

func main() { // Main function, entry point of the program
    var wg sync.WaitGroup // Creates a new WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) // Adds a count to the WaitGroup
        go worker(i, &wg) // Starts worker function as a goroutine
    }

    wg.Wait() // Waits for all goroutines to finish
}
// Context Package

package main // Declares the main package

import (
    "context" // Imports the context package for managing context
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func work(ctx context.Context) {
    select {
    case <-time.After(time.Second * 2): // Simulate work
        fmt.Println("Work completed")
    case <-ctx.Done(): // Check if context is canceled
        fmt.Println("Work canceled:", ctx.Err())
    }
}

func main() { // Main function, entry point of the program
    ctx, cancel := context.WithTimeout(context.Background(), time.Second) // Creates a context with a timeout
    defer cancel() // Ensures cancel function is called to release resources

    go work(ctx) // Starts work function as a goroutine
    time.Sleep(time.Millisecond * 500) // Simulates main work
    fmt.Println("Main work done") // Prints a message from the main function
}
// Time Parsing

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func main() { // Main function, entry point of the program
    layout := "2006-01-02 15:04:05" // Time format layout
    str := "2024-09-07 12:34:56" // Time string to parse
    t, err := time.Parse(layout, str) // Parses the time string
    if err != nil { // Checks for parsing error
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("Parsed time:", t) // Prints the parsed time
}
// Command Line Flag Parsing

package main // Declares the main package

import (
    "flag" // Imports the flag package for command-line flags
    "fmt" // Imports the fmt package for formatted I/O
)

func main() { // Main function, entry point of the program
    var port int // Declares an integer flag for port
    var env string // Declares a string flag for environment

    flag.IntVar(&port, "port", 8080, "Port to listen on") // Defines an integer flag for port
    flag.StringVar(&env, "env", "development", "Environment") // Defines a string flag for environment
    flag.Parse() // Parses the command-line flags

    fmt.Println("Port:", port) // Prints the port value
    fmt.Println("Environment:", env) // Prints the environment value
}
// Function Types

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a function type
type Operation func(int, int) int

// Define functions that match the function type
func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func main() { // Main function, entry point of the program
    var op Operation // Declare a variable of type Operation
    op = add // Assign the add function to op
    fmt.Println("Addition:", op(3, 2)) // Prints the result of add function

    op = subtract // Assign the subtract function to op
    fmt.Println("Subtraction:", op(3, 2)) // Prints the result of subtract function
}
// Goroutine Communication

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func sendData(ch chan<- string) { // Function to send data to a channel
    time.Sleep(time.Second) // Simulates work with sleep
    ch <- "data" // Sends data to the channel
}

func receiveData(ch <-chan string) { // Function to receive data from a channel
    data := <-ch // Receives data from the channel
    fmt.Println("Received:", data) // Prints the received data
}

func main() { // Main function, entry point of the program
    ch := make(chan string) // Creates a new channel

    go sendData(ch) // Starts sendData function as a goroutine
    receiveData(ch) // Calls receiveData function
}
// String Manipulation

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "strings" // Imports the strings package for string manipulation
)

func main() { // Main function, entry point of the program
    str := "hello, world" // A sample string
    upperStr := strings.ToUpper(str) // Converts string to uppercase
    fmt.Println("Uppercase:", upperStr) // Prints the uppercase string

    replacedStr := strings.Replace(str, "world", "Go", -1) // Replaces "world" with "Go"
    fmt.Println("Replaced:", replacedStr) // Prints the replaced string

    splitStr := strings.Split(str, ", ") // Splits the string by comma
    fmt.Println("Split:", splitStr) // Prints the split string
}
// Interfaces and Polymorphism

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define an interface
type Shape interface {
    Area() float64
}

// Define a struct type
type Circle struct {
    Radius float64
}

// Implement the interface for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius // Calculates the area of the circle
}

// Define another struct type
type Rectangle struct {
    Width, Height float64
}

// Implement the interface for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height // Calculates the area of the rectangle
}

func printArea(s Shape) { // Function that accepts a Shape interface
    fmt.Println("Area:", s.Area()) // Prints the area of the shape
}

func main() { // Main function, entry point of the program
    circle := Circle{Radius: 5} // Creates a Circle instance
    rectangle := Rectangle{Width: 4, Height: 6} // Creates a Rectangle instance

    printArea(circle) // Calls printArea with Circle
    printArea(rectangle) // Calls printArea with Rectangle
}
// Channels and Select

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func main() { // Main function, entry point of the program
    ch1 := make(chan string) // Creates a channel for string
    ch2 := make(chan string) // Creates another channel for string

    go func() { // Anonymous goroutine for sending to ch1
        time.Sleep(time.Second) // Simulates work with sleep
        ch1 <- "Message from ch1" // Sends message to ch1
    }()

    go func() { // Anonymous goroutine for sending to ch2
        time.Sleep(2 * time.Second) // Simulates work with sleep
        ch2 <- "Message from ch2" // Sends message to ch2
    }()

    select {
    case msg1 := <-ch1: // Receives message from ch1
        fmt.Println("Received:", msg1) // Prints the received message
    case msg2 := <-ch2: // Receives message from ch2
        fmt.Println("Received:", msg2) // Prints the received message
    case <-time.After(time.Second * 3): // Times out after 3 seconds
        fmt.Println("Timeout") // Prints timeout message
    }
}
// Advanced Goroutine Patterns

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

// Worker function that performs a task
func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done() // Signals that this goroutine is done
    for task := range ch { // Iterates over tasks from the channel
        fmt.Printf("Worker %d processing task %d\n", id, task)
    }
}

func main() { // Main function, entry point of the program
    const numWorkers = 3 // Number of workers
    tasks := make(chan int, 10) // Creates a buffered channel for tasks
    var wg sync.WaitGroup // Creates a new WaitGroup

    // Start worker goroutines
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1) // Adds a count to the WaitGroup
        go worker(i, tasks, &wg) // Starts worker function as a goroutine
    }

    // Send tasks to the channel
    for i := 1; i <= 10; i++ {
        tasks <- i // Sends task to the channel
    }
    close(tasks) // Closes the channel after sending all tasks

    wg.Wait() // Waits for all goroutines to finish
}
// Mutex and Concurrent Map Access

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

// Define a struct to hold shared data
type SafeMap struct {
    data map[string]int
    mu   sync.Mutex // Mutex for synchronizing access
}

// Method to safely add a key-value pair to the map
func (sm *SafeMap) Add(key string, value int) {
    sm.mu.Lock() // Lock the mutex
    sm.data[key] = value // Modify the map
    sm.mu.Unlock() // Unlock the mutex
}

// Method to safely get a value from the map
func (sm *SafeMap) Get(key string) int {
    sm.mu.Lock() // Lock the mutex
    defer sm.mu.Unlock() // Unlock the mutex after the function returns
    return sm.data[key] // Return the value
}

func main() { // Main function, entry point of the program
    sm := SafeMap{data: make(map[string]int)} // Creates a SafeMap instance

    var wg sync.WaitGroup // Creates a new WaitGroup

    // Start multiple goroutines to modify the map
    for i := 0; i < 10; i++ {
        wg.Add(1) // Adds a count to the WaitGroup
        go func(i int) {
            defer wg.Done() // Signals that this goroutine is done
            sm.Add(fmt.Sprintf("key%d", i), i) // Adds data to the map
        }(i)
    }

    wg.Wait() // Waits for all goroutines to finish

    // Prints the map data
    for i := 0; i < 10; i++ {
        fmt.Printf("key%d: %d\n", i, sm.Get(fmt.Sprintf("key%d", i)))
    }
}
// Reflection in Go

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "reflect" // Imports the reflect package for reflection
)

// Define a struct with different types of fields
type Person struct {
    Name string
    Age  int
}

func main() { // Main function, entry point of the program
    p := Person{Name: "Alice", Age: 30} // Creates a Person instance

    v := reflect.ValueOf(p) // Gets the reflection value of p
    t := reflect.TypeOf(p) // Gets the reflection type of p

    fmt.Println("Type:", t) // Prints the type of the Person struct

    for i := 0; i < v.NumField(); i++ { // Iterates over fields
        field := v.Field(i) // Gets the field value
        fmt.Printf("Field %d: %s = %v\n", i, t.Field(i).Name, field) // Prints field name and value
    }
}
// Channel Buffering and Select with Timeout

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

func producer(ch chan<- int) { // Function to send data to a channel
    for i := 0; i < 5; i++ {
        ch <- i // Sends data to the channel
        time.Sleep(time.Millisecond * 500) // Simulates delay
    }
    close(ch) // Closes the channel after sending all data
}

func main() { // Main function, entry point of the program
    ch := make(chan int, 3) // Creates a buffered channel with capacity 3

    go producer(ch) // Starts producer function as a goroutine

    for {
        select {
        case msg, ok := <-ch: // Receives data from the channel
            if !ok { // Checks if the channel is closed
                fmt.Println("Channel closed")
                return // Exits the loop if channel is closed
            }
            fmt.Println("Received:", msg) // Prints the received data
        case <-time.After(time.Second): // Times out after 1 second
            fmt.Println("Timeout")
        }
    }
}
// Generics (Go 1.18+)

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a generic function with type parameter T
func Print[T any](value T) {
    fmt.Println(value) // Prints the value of type T
}

func main() { // Main function, entry point of the program
    Print("Hello, Generics!") // Calls Print with a string
    Print(123) // Calls Print with an integer
}
// Error Handling with Custom Errors

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "errors" // Imports the errors package for error handling
)

// Define a custom error type
type MyError struct {
    Code    int
    Message string
}

// Implement the Error() method for MyError
func (e MyError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func mightFail(flag bool) error { // Function that returns an error based on a flag
    if flag {
        return MyError{Code: 123, Message: "Something went wrong"} // Returns a custom error
    }
    return nil // Returns nil if no error
}

func main() { // Main function, entry point of the program
    err := mightFail(true) // Calls mightFail function
    if err != nil { // Checks for an error
        fmt.Println("Error:", err) // Prints the error
    } else {
        fmt.Println("No error occurred")
    }
}
// Middleware in HTTP Servers

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "net/http" // Imports the net/http package for HTTP handling
)

// Middleware function that logs requests
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request received:", r.Method, r.URL.Path) // Logs the request
        next.ServeHTTP(w, r) // Calls the next handler
    })
}

// Handler function for HTTP requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Middleware!") // Writes response to the client
}

func main() { // Main function, entry point of the program
    mux := http.NewServeMux() // Creates a new ServeMux
    mux.Handle("/", helloHandler) // Registers the handler function

    http.ListenAndServe(":8080", loggingMiddleware(mux)) // Starts the HTTP server with middleware
}
// Using `sync/atomic` for Atomic Operations

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync/atomic" // Imports the sync/atomic package for atomic operations
)

func main() { // Main function, entry point of the program
    var counter int64 // Defines an int64 variable for the counter

    // Increment the counter atomically
    atomic.AddInt64(&counter, 1)
    fmt.Println("Counter after increment:", atomic.LoadInt64(&counter)) // Prints the counter value

    // Decrement the counter atomically
    atomic.AddInt64(&counter, -1)
    fmt.Println("Counter after decrement:", atomic.LoadInt64(&counter)) // Prints the counter value
}
// Concurrent Data Structures

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

// Define a concurrent-safe stack using mutex
type Stack struct {
    mu    sync.Mutex
    items []int
}

// Method to push an item onto the stack
func (s *Stack) Push(item int) {
    s.mu.Lock() // Lock the mutex
    s.items = append(s.items, item) // Append item to the stack
    s.mu.Unlock() // Unlock the mutex
}

// Method to pop an item from the stack
func (s *Stack) Pop() (int, bool) {
    s.mu.Lock() // Lock the mutex
    defer s.mu.Unlock() // Unlock the mutex after the function returns
    if len(s.items) == 0 {
        return 0, false // Return false if stack is empty
    }
    item := s.items[len(s.items)-1] // Get the last item
    s.items = s.items[:len(s.items)-1] // Remove the last item
    return item, true // Return the item and true
}

func main() { // Main function, entry point of the program
    stack := Stack{} // Creates a new Stack instance

    stack.Push(1) // Pushes an item onto the stack
    stack.Push(2) // Pushes another item

    item, ok := stack.Pop() // Pops an item from the stack
    if ok {
        fmt.Println("Popped:", item) // Prints the popped item
    } else {
        fmt.Println("Stack is empty")
    }
}
// Type Assertion and Type Switch

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

// Define an interface with multiple implementations
type Animal interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
    return "Woof"
}

func (c Cat) Speak() string {
    return "Meow"
}

func printAnimalSound(a Animal) {
    switch v := a.(type) { // Type switch to determine the concrete type
    case Dog:
        fmt.Println("Dog says:", v.Speak()) // Prints sound of a Dog
    case Cat:
        fmt.Println("Cat says:", v.Speak()) // Prints sound of a Cat
    default:
        fmt.Println("Unknown animal")
    }
}

func main() { // Main function, entry point of the program
    var a Animal = Dog{} // Create a Dog as an Animal
    printAnimalSound(a) // Prints the sound of the Dog

    a = Cat{} // Change to a Cat
    printAnimalSound(a) // Prints the sound of the Cat
}
// Advanced Error Handling with Wrapping

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "errors" // Imports the errors package for error handling
)

// Define a custom error type with wrapping
type CustomError struct {
    Code    int
    Message string
    Err     error // Embedded error for wrapping
}

// Implement the Error() method for CustomError
func (e CustomError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("Code %d: %s - %v", e.Code, e.Message, e.Err)
    }
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
    err := errors.New("underlying error") // Create an underlying error
    return CustomError{Code: 404, Message: "Not Found", Err: err} // Wrap the underlying error
}

func main() { // Main function, entry point of the program
    err := doSomething() // Call a function that returns an error
    if err != nil { // Check if there is an error
        fmt.Println("Error:", err) // Prints the wrapped error
    }
}
// Custom JSON Unmarshalling

package main // Declares the main package

import (
    "encoding/json" // Imports the encoding/json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
)

// Define a struct type with custom unmarshalling
type Person struct {
    Name string
    Age  int
}

// Implement custom UnmarshalJSON method for Person
func (p *Person) UnmarshalJSON(data []byte) error {
    var raw map[string]interface{}
    if err := json.Unmarshal(data, &raw); err != nil {
        return err
    }
    if name, ok := raw["name"].(string); ok {
        p.Name = name
    }
    if age, ok := raw["age"].(float64); ok { // JSON numbers are float64 by default
        p.Age = int(age)
    }
    return nil
}

func main() { // Main function, entry point of the program
    jsonStr := `{"name":"Alice","age":30}` // JSON string
    var person Person // Declares a variable of type Person

    if err := json.Unmarshal([]byte(jsonStr), &person); err != nil { // Decodes JSON string
        fmt.Println("Error:", err) // Prints the error if any
        return // Exits the function
    }

    fmt.Println("Person:", person) // Prints the decoded Person instance
}
// Pipeline Pattern

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

// Function to generate a sequence of numbers
func generateNumbers(n int, ch chan<- int) {
    for i := 0; i < n; i++ {
        ch <- i // Sends numbers to the channel
    }
    close(ch) // Closes the channel after sending all numbers
}

// Function to process numbers from a channel
func processNumbers(in <-chan int, out chan<- int) {
    for num := range in { // Iterates over numbers from the channel
        out <- num * 2 // Sends processed numbers to the output channel
    }
    close(out) // Closes the output channel after processing all numbers
}

func main() { // Main function, entry point of the program
    numChan := make(chan int) // Creates a channel for numbers
    processedChan := make(chan int) // Creates a channel for processed numbers

    var wg sync.WaitGroup // Creates a new WaitGroup

    // Start number generation
    wg.Add(1)
    go func() {
        defer wg.Done() // Signals that this goroutine is done
        generateNumbers(10, numChan) // Generates numbers
    }()

    // Start number processing
    wg.Add(1)
    go func() {
        defer wg.Done() // Signals that this goroutine is done
        processNumbers(numChan, processedChan) // Processes numbers
    }()

    // Print processed numbers
    go func() {
        for num := range processedChan { // Iterates over processed numbers
            fmt.Println(num) // Prints each processed number
        }
    }()

    wg.Wait() // Waits for all goroutines to finish
}
// Custom HTTP Middleware with Context

package main // Declares the main package

import (
    "context" // Imports the context package for managing context
    "fmt" // Imports the fmt package for formatted I/O
    "net/http" // Imports the net/http package for HTTP handling
)

// Middleware function to add values to context
func contextMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "user", "John Doe") // Add value to context
        next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler with updated context
    })
}

// Handler function to read values from context
func handler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user") // Retrieve value from context
    fmt.Fprintf(w, "Hello, %s!", user) // Writes response to the client
}

func main() { // Main function, entry point of the program
    mux := http.NewServeMux() // Creates a new ServeMux
    mux.Handle("/", handler) // Registers the handler function

    http.ListenAndServe(":8080", contextMiddleware(mux)) // Starts the HTTP server with middleware
}
// Implementing a Custom Iterator

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

// Define an iterator for a custom collection
type Iterator struct {
    data  []int
    index int
}

// Create a new iterator
func NewIterator(data []int) *Iterator {
    return &Iterator{data: data}
}

// Check if there are more items to iterate
func (it *Iterator) HasNext() bool {
    return it.index < len(it.data)
}

// Get the next item in the iteration
func (it *Iterator) Next() (int, bool) {
    if it.HasNext() {
        item := it.data[it.index]
        it.index++
        return item, true
    }
    return 0, false
}

func main() { // Main function, entry point of the program
    data := []int{1, 2, 3, 4, 5} // Sample data
    it := NewIterator(data) // Creates a new iterator

    // Iterate over the data using the iterator
    for it.HasNext() {
        item, _ := it.Next()
        fmt.Println(item) // Prints each item
    }
}
// Advanced Usage of Channels

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
)

// Function to produce a stream of numbers
func numberStream(start, end int, out chan<- int) {
    for i := start; i <= end; i++ {
        out <- i // Sends numbers to the channel
    }
    close(out) // Closes the channel after sending all numbers
}

// Function to consume numbers from a channel and process them
func processNumbers(in <-chan int, wg *sync.WaitGroup) {
    defer wg.Done() // Signals that this goroutine is done
    for num := range in { // Iterates over numbers from the channel
        fmt.Println("Processed:", num) // Prints processed numbers
    }
}

func main() { // Main function, entry point of the program
    const numWorkers = 3
    var wg sync.WaitGroup

    // Create a channel to stream numbers
    numChan := make(chan int, 10)

    // Start number stream
    go numberStream(1, 15, numChan)

    // Start multiple processing workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go processNumbers(numChan, &wg)
    }

    wg.Wait() // Waits for all processing goroutines to finish
}
// Advanced Concurrency with Worker Pools

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Define a worker function that processes tasks
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done() // Signals that this goroutine is done
    for task := range tasks { // Processes tasks from the channel
        fmt.Printf("Worker %d processing task %d\n", id, task)
        time.Sleep(time.Millisecond * 100) // Simulate work
        results <- task * task // Send the result back to the results channel
    }
}

// Define a function to generate tasks
func generateTasks(n int, tasks chan<- int) {
    for i := 1; i <= n; i++ {
        tasks <- i // Send tasks to the channel
    }
    close(tasks) // Close the channel after sending all tasks
}

func main() { // Main function, entry point of the program
    const numWorkers = 4 // Number of worker goroutines
    const numTasks = 10 // Number of tasks to process

    tasks := make(chan int, numTasks) // Creates a buffered channel for tasks
    results := make(chan int, numTasks) // Creates a buffered channel for results
    var wg sync.WaitGroup // Creates a new WaitGroup

    // Start worker goroutines
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, results, &wg)
    }

    // Generate tasks
    go generateTasks(numTasks, tasks)

    // Wait for all workers to finish
    go func() {
        wg.Wait()
        close(results)
    }()

    // Print results
    for result := range results {
        fmt.Println("Result:", result)
    }
}
// Network Programming with TCP Server

package main // Declares the main package

import (
    "bufio" // Imports the bufio package for buffered I/O
    "fmt" // Imports the fmt package for formatted I/O
    "net" // Imports the net package for network programming
    "os" // Imports the os package for OS operations
)

// Define a function to handle client connections
func handleConnection(conn net.Conn) {
    defer conn.Close() // Ensure the connection is closed when done
    reader := bufio.NewReader(conn) // Create a buffered reader for the connection

    for {
        message, err := reader.ReadString('\n') // Read a line from the connection
        if err != nil {
            fmt.Println("Error reading from connection:", err)
            return
        }
        fmt.Print("Received:", message) // Print the received message
        _, err = conn.Write([]byte("Message received\n")) // Send a response to the client
        if err != nil {
            fmt.Println("Error writing to connection:", err)
            return
        }
    }
}

func main() { // Main function, entry point of the program
    // Start a TCP server
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }
    defer listener.Close() // Ensure the listener is closed when done

    fmt.Println("Server is listening on port 8080")

    for {
        conn, err := listener.Accept() // Accept a new connection
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn) // Handle the connection in a new goroutine
    }
}
// Using the `context` Package for Cancellation and Timeouts

package main // Declares the main package

import (
    "context" // Imports the context package for managing context
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

// Define a function that simulates a long-running task
func longRunningTask(ctx context.Context) {
    select {
    case <-time.After(5 * time.Second): // Simulate work with a delay
        fmt.Println("Task completed")
    case <-ctx.Done(): // Handle context cancellation
        fmt.Println("Task canceled:", ctx.Err())
    }
}

func main() { // Main function, entry point of the program
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel() // Ensure the context is canceled when done

    longRunningTask(ctx) // Start a long-running task

    // Wait to demonstrate the timeout
    time.Sleep(4 * time.Second)
}
// Implementing a Custom HTTP Server with Custom Logging

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "log" // Imports the log package for logging
    "net/http" // Imports the net/http package for HTTP handling
    "time" // Imports the time package for time handling
)

// Custom logger middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now() // Record start time
        next.ServeHTTP(w, r) // Call the next handler
        log.Printf("Request %s %s took %v", r.Method, r.URL.Path, time.Since(start)) // Log the request
    })
}

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Custom Logging!") // Writes response to the client
}

func main() { // Main function, entry point of the program
    mux := http.NewServeMux() // Creates a new ServeMux
    mux.Handle("/", helloHandler) // Registers the handler function

    // Create a custom HTTP server with logging middleware
    server := &http.Server{
        Addr:    ":8080",
        Handler: loggingMiddleware(mux),
    }

    log.Println("Starting server on port 8080")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal("Server failed:", err) // Log and exit if server fails
    }
}
// Understanding Go's Memory Model and Race Conditions

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Shared data structure with a race condition
type SharedData struct {
    counter int
}

func (sd *SharedData) Increment() {
    sd.counter++ // Increment counter
}

func main() { // Main function, entry point of the program
    sd := &SharedData{} // Create a new SharedData instance
    var wg sync.WaitGroup // Create a new WaitGroup

    // Start multiple goroutines that access shared data
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done() // Signal when the goroutine is done
            sd.Increment() // Increment shared counter
        }()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("Final counter value:", sd.counter) // Print the final counter value

    // Uncomment the following line to run the race detector:
    // go test -race
}
// Using Go Generics for a Generic Queue Implementation

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a generic queue data structure
type Queue[T any] struct {
    items []T
}

// Method to enqueue an item
func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item) // Add item to the queue
}

// Method to dequeue an item
func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false // Return zero value if queue is empty
    }
    item := q.items[0] // Get the first item
    q.items = q.items[1:] // Remove the first item
    return item, true // Return the item
}

func main() { // Main function, entry point of the program
    q := Queue[int]{} // Create a new queue for integers

    // Enqueue some items
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    // Dequeue and print items
    for {
        item, ok := q.Dequeue()
        if !ok {
            break
        }
        fmt.Println("Dequeued:", item) // Print each dequeued item
    }
}
// Advanced Use of `sync/atomic` for Concurrent Counters

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync/atomic" // Imports the sync/atomic package for atomic operations
)

// Define an atomic counter with int64
type AtomicCounter struct {
    value int64
}

// Method to increment the counter atomically
func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1) // Atomically increment the counter
}

// Method to get the counter value
func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.value) // Atomically load the counter value
}

func main() { // Main function, entry point of the program
    counter := &AtomicCounter{} // Create a new AtomicCounter
    var wg sync.WaitGroup // Create a new WaitGroup

    // Start multiple goroutines to increment the counter
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done() // Signal when the goroutine is done
            counter.Increment() // Increment the counter
        }()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("Final counter value:", counter.Value()) // Print the final counter value
}
// Implementing a Custom Mutex with Spinlock

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync/atomic" // Imports the sync/atomic package for atomic operations
    "time" // Imports the time package for time handling
)

// Define a custom spinlock
type Spinlock struct {
    locked int32
}

// Lock the spinlock
func (s *Spinlock) Lock() {
    for !atomic.CompareAndSwapInt32(&s.locked, 0, 1) {
        time.Sleep(time.Nanosecond) // Busy-wait (spin) if the lock is held
    }
}

// Unlock the spinlock
func (s *Spinlock) Unlock() {
    atomic.StoreInt32(&s.locked, 0) // Release the lock
}

func main() { // Main function, entry point of the program
    var lock Spinlock // Create a new Spinlock
    var wg sync.WaitGroup // Create a new WaitGroup

    // Start multiple goroutines that use the spinlock
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done() // Signal when the goroutine is done
            lock.Lock() // Acquire the lock
            fmt.Printf("Goroutine %d has the lock\n", id)
            time.Sleep(time.Second) // Simulate work
            lock.Unlock() // Release the lock
        }(i)
    }

    wg.Wait() // Wait for all goroutines to finish
}
// Using Reflection to Inspect and Manipulate Data

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "reflect" // Imports the reflect package for reflection
)

// Define a struct for demonstration
type Person struct {
    Name string
    Age  int
}

func printStructFields(v interface{}) {
    val := reflect.ValueOf(v) // Get the reflect.Value of the input
    typ := reflect.TypeOf(v) // Get the reflect.Type of the input

    if val.Kind() != reflect.Struct {
        fmt.Println("Expected a struct")
        return
    }

    // Iterate over the fields of the struct
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i) // Get the field value
        fieldType := typ.Field(i) // Get the field type
        fmt.Printf("Field Name: %s, Field Value: %v\n", fieldType.Name, field.Interface())
    }
}

func main() { // Main function, entry point of the program
    person := Person{Name: "Alice", Age: 30} // Create a Person instance
    printStructFields(person) // Print the fields of the Person struct
}
// Implementing a Bounded Buffer with Channels

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
)

// Define a bounded buffer using channels
type BoundedBuffer struct {
    buffer chan int
}

// Create a new BoundedBuffer with a given size
func NewBoundedBuffer(size int) *BoundedBuffer {
    return &BoundedBuffer{buffer: make(chan int, size)}
}

// Add an item to the buffer
func (b *BoundedBuffer) Put(item int) {
    b.buffer <- item // Send the item to the buffer
}

// Remove an item from the buffer
func (b *BoundedBuffer) Get() int {
    return <-b.buffer // Receive an item from the buffer
}

func main() { // Main function, entry point of the program
    buf := NewBoundedBuffer(3) // Create a buffer with size 3

    // Producer
    go func() {
        for i := 0; i < 10; i++ {
            buf.Put(i) // Add items to the buffer
            fmt.Println("Produced:", i)
        }
    }()

    // Consumer
    go func() {
        for i := 0; i < 10; i++ {
            item := buf.Get() // Remove items from the buffer
            fmt.Println("Consumed:", item)
        }
    }()

    // Wait to allow goroutines to finish
    time.Sleep(time.Second * 2)
}
// Advanced Data Structures: Implementing a Binary Search Tree

package main // Declares the main package

import "fmt" // Imports the fmt package for formatted I/O

// Define a node in the binary search tree
type Node struct {
    value int
    left  *Node
    right *Node
}

// Insert a value into the binary search tree
func (n *Node) Insert(value int) {
    if value < n.value {
        if n.left == nil {
            n.left = &Node{value: value}
        } else {
            n.left.Insert(value)
        }
    } else {
        if n.right == nil {
            n.right = &Node{value: value}
        } else {
            n.right.Insert(value)
        }
    }
}

// In-order traversal of the binary search tree
func (n *Node) InOrderTraversal() {
    if n.left != nil {
        n.left.InOrderTraversal()
    }
    fmt.Print(n.value, " ") // Print the node value
    if n.right != nil {
        n.right.InOrderTraversal()
    }
}

func main() { // Main function, entry point of the program
    root := &Node{value: 10} // Create the root node
    root.Insert(5) // Insert values
    root.Insert(15)
    root.Insert(3)
    root.Insert(7)
    root.Insert(12)
    root.Insert(18)

    fmt.Print("In-order Traversal: ")
    root.InOrderTraversal() // Print the values in sorted order
    fmt.Println()
}
// Advanced Usage of Channels: Buffered Channel Patterns

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "time" // Imports the time package for time handling
)

// Define a function to generate numbers
func generateNumbers(n int, ch chan<- int) {
    for i := 1; i <= n; i++ {
        ch <- i // Send numbers to the channel
        time.Sleep(time.Millisecond * 100) // Simulate delay
    }
    close(ch) // Close the channel after sending all numbers
}

// Define a function to consume numbers from a buffered channel
func consumeNumbers(ch <-chan int) {
    for num := range ch { // Receive numbers from the channel
        fmt.Println("Consumed:", num) // Print each number
    }
}

func main() { // Main function, entry point of the program
    ch := make(chan int, 5) // Create a buffered channel with capacity 5

    go generateNumbers(20, ch) // Start generating numbers

    consumeNumbers(ch) // Start consuming numbers
}
// Creating a Custom Command-Line Tool with Flag Parsing

package main // Declares the main package

import (
    "flag" // Imports the flag package for command-line flag parsing
    "fmt" // Imports the fmt package for formatted I/O
)

// Define flags for command-line arguments
var (
    name   = flag.String("name", "World", "Name to greet")
    repeat = flag.Int("repeat", 1, "Number of times to repeat the greeting")
)

func main() { // Main function, entry point of the program
    flag.Parse() // Parse the command-line flags

    for i := 0; i < *repeat; i++ {
        fmt.Printf("Hello, %s!\n", *name) // Print the greeting
    }
}
// Implementing a Custom Logger with Levels

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "os" // Imports the os package for OS operations
)

// Define log levels
const (
    INFO  = iota
    WARN
    ERROR
)

// Custom logger struct
type Logger struct {
    level int
}

// Create a new logger with a specified level
func NewLogger(level int) *Logger {
    return &Logger{level: level}
}

// Log a message with a specific level
func (l *Logger) Log(level int, msg string) {
    if level >= l.level {
        var prefix string
        switch level {
        case INFO:
            prefix = "INFO: "
        case WARN:
            prefix = "WARN: "
        case ERROR:
            prefix = "ERROR: "
        }
        fmt.Fprintf(os.Stderr, "%s%s\n", prefix, msg)
    }
}

func main() { // Main function, entry point of the program
    logger := NewLogger(INFO) // Create a new logger with INFO level

    logger.Log(INFO, "This is an info message")
    logger.Log(WARN, "This is a warning message")
    logger.Log(ERROR, "This is an error message")
}
// Advanced Web Server with Database Integration and Middleware

package main // Declares the main package

import (
    "database/sql" // Imports the sql package for database operations
    "encoding/json" // Imports the json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
    "log" // Imports the log package for logging
    "net/http" // Imports the net/http package for HTTP handling
    "os" // Imports the os package for OS operations
    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/gorilla/mux" // Imports the Gorilla Mux router
)

// Define a struct for holding user data
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Database connection configuration
const (
    dbUser     = "yourusername"
    dbPassword = "yourpassword"
    dbName     = "yourdbname"
)

// Create a global variable for the database connection
var db *sql.DB

// Initialize the database connection
func initDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Handler function to get all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, "Error retrieving users", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            http.Error(w, "Error scanning user", http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Handler function to get a single user by ID
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
    var user User
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// Handler function to create a new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// Handler function to update a user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
    if err != nil {
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// Handler function to delete a user
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func main() { // Main function, entry point of the program
    var err error
    db, err = initDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }

    r := mux.NewRouter()
    r.Use(loggingMiddleware)

    r.HandleFunc("/users", getUsersHandler).Methods(http.MethodGet)
    r.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods(http.MethodGet)
    r.HandleFunc("/users", createUserHandler).Methods(http.MethodPost)
    r.HandleFunc("/users/{id:[0-9]+}", updateUserHandler).Methods(http.MethodPut)
    r.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods(http.MethodDelete)

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server
}
// Complex Concurrent Processing System

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "math/rand" // Imports the rand package for random number generation
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Define a struct for a task
type Task struct {
    ID    int
    Data  int
    Result int
}

// Function to process a task
func processTask(task Task, wg *sync.WaitGroup, results chan<- Task) {
    defer wg.Done() // Signal when the goroutine is done

    // Simulate processing time
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    task.Result = task.Data * 2 // Dummy processing: multiply by 2

    results <- task // Send the result to the results channel
}

// Function to distribute tasks to worker goroutines
func distributeTasks(tasks []Task, numWorkers int) []Task {
    var wg sync.WaitGroup
    results := make(chan Task, len(tasks))
    workers := make(chan struct{}, numWorkers)

    // Start worker goroutines
    for i := 0; i < numWorkers; i++ {
        workers <- struct{}{}
    }

    for _, task := range tasks {
        wg.Add(1)
        <-workers // Get a worker slot

        go func(t Task) {
            defer func() { workers <- struct{}{} }()
            processTask(t, &wg, results)
        }(task)
    }

    wg.Wait() // Wait for all tasks to complete
    close(results) // Close the results channel

    // Collect results
    var processedTasks []Task
    for result := range results {
        processedTasks = append(processedTasks, result)
    }

    return processedTasks
}

func main() { // Main function, entry point of the program
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator

    // Create tasks
    var tasks []Task
    for i := 0; i < 100; i++ {
        tasks = append(tasks, Task{ID: i, Data: rand.Intn(100)})
    }

    fmt.Println("Distributing tasks...")
    processedTasks := distributeTasks(tasks, 10) // Distribute tasks with 10 workers

    // Print results
    fmt.Println("Processed tasks:")
    for _, task := range processedTasks {
        fmt.Printf("Task ID: %d, Data: %d, Result: %d\n", task.ID, task.Data, task.Result)
    }
}
// Advanced Data Processing Pipeline

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Define a struct for a data item
type DataItem struct {
    ID   int
    Data int
}

// Function to generate data items
func generateData(numItems int, ch chan<- DataItem) {
    for i := 0; i < numItems; i++ {
        ch <- DataItem{ID: i, Data: i * 2}
    }
    close(ch) // Close the channel after generating all items
}

// Function to process data items
func processData(in <-chan DataItem, out chan<- DataItem, wg *sync.WaitGroup) {
    defer wg.Done() // Signal when the goroutine is done
    for item := range in {
        // Simulate processing time
        time.Sleep(time.Duration(item.Data%10) * time.Millisecond)
        processedItem := DataItem{ID: item.ID, Data: item.Data + 1}
        out <- processedItem
    }
    close(out) // Close the output channel after processing
}

// Function to collect and display processed data
func collectData(in <-chan DataItem) {
    for item := range in {
        fmt.Printf("Processed DataItem ID: %d, Data: %d\n", item.ID, item.Data)
    }
}

func main() { // Main function, entry point of the program
    numItems := 100
    numWorkers := 10

    // Create channels
    dataChan := make(chan DataItem, numItems)
    processedChan := make(chan DataItem, numItems)

    var wg sync.WaitGroup

    // Start data generation
    go generateData(numItems, dataChan)

    // Start worker goroutines for processing
    wg.Add(numWorkers)
    for i := 0; i < numWorkers; i++ {
        go processData(dataChan, processedChan, &wg)
    }

    // Collect and display processed data
    go collectData(processedChan)

    wg.Wait() // Wait for all workers to finish
}
// Custom Distributed Task Queue

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "math/rand" // Imports the rand package for random number generation
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Define a struct for a task
type Task struct {
    ID    int
    Data  int
}

// Define a struct for a worker
type Worker struct {
    ID int
}

// Function for worker to process tasks
func (w *Worker) processTasks(tasks <-chan Task, results chan<- Task, wg *sync.WaitGroup) {
    defer wg.Done() // Signal when the goroutine is done
    for task := range tasks {
        // Simulate processing time
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        task.Data += 1 // Dummy processing: increment data
        results <- task // Send result to results channel
    }
}

// Function to create and distribute tasks to workers
func createTaskQueue(numTasks, numWorkers int) []Task {
    tasks := make(chan Task, numTasks)
    results := make(chan Task, numTasks)
    var wg sync.WaitGroup

    // Start worker goroutines
    workers := make([]Worker, numWorkers)
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go workers[i].processTasks(tasks, results, &wg)
    }

    // Generate tasks
    go func() {
        for i := 0; i < numTasks; i++ {
            tasks <- Task{ID: i, Data: rand.Intn(100)}
        }
        close(tasks) // Close the tasks channel after generating all tasks
    }()

    wg.Wait() // Wait for all workers to finish
    close(results) // Close the results channel after processing

    // Collect results
    var processedTasks []Task
    for result := range results {
        processedTasks = append(processedTasks, result)
    }

    return processedTasks
}

func main() { // Main function, entry point of the program
    numTasks := 50
    numWorkers := 5

    fmt.Println("Creating task queue...")
    processedTasks := createTaskQueue(numTasks, numWorkers)

    // Print processed tasks
    fmt.Println("Processed tasks:")
    for _, task := range processedTasks {
        fmt.Printf("Task ID: %d, Data: %d\n", task.ID, task.Data)
    }
}
// Distributed Key-Value Store with Replication

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "sync" // Imports the sync package for synchronization
    "net/http" // Imports the net/http package for HTTP handling
    "encoding/json" // Imports the json package for JSON handling
    "io/ioutil" // Imports the ioutil package for reading files
)

// Define a struct for the key-value store
type KeyValueStore struct {
    data map[string]string
    mu   sync.RWMutex
}

// Create a new key-value store
func NewKeyValueStore() *KeyValueStore {
    return &KeyValueStore{data: make(map[string]string)}
}

// Handler function to get a value by key
func (store *KeyValueStore) getHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    store.mu.RLock()
    value, exists := store.data[key]
    store.mu.RUnlock()

    if !exists {
        http.Error(w, "Key not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"value": value})
}

// Handler function to set a value by key
func (store *KeyValueStore) setHandler(w http.ResponseWriter, r *http.Request) {
    var input map[string]string
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }
    if err := json.Unmarshal(body, &input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    key, exists := input["key"]
    value, exists := input["value"]
    if !exists {
        http.Error(w, "Missing key or value", http.StatusBadRequest)
        return
    }

    store.mu.Lock()
    store.data[key] = value
    store.mu.Unlock()

    w.WriteHeader(http.StatusNoContent)
}

// Replicate data to another store
func replicateData(store *KeyValueStore, targetURL string) {
    store.mu.RLock()
    data := make(map[string]string)
    for k, v := range store.data {
        data[k] = v
    }
    store.mu.RUnlock()

    body, _ := json.Marshal(data)
    _, err := http.Post(targetURL, "application/json", bytes.NewBuffer(body))
    if err != nil {
        fmt.Printf("Error replicating data: %v\n", err)
    }
}

func main() { // Main function, entry point of the program
    store := NewKeyValueStore()
    targetURL := "http://localhost:8081/replicate"

    go http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            store.getHandler(w, r)
        case http.MethodPost:
            store.setHandler(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))

    // Simulate replication every minute
    go func() {
        for {
            time.Sleep(time.Minute)
            replicateData(store, targetURL)
        }
    }()

    // Keep the server running
    select {}
}
// Event-Driven System with WebSocket and REST API

package main // Declares the main package

import (
    "fmt" // Imports the fmt package for formatted I/O
    "net/http" // Imports the net/http package for HTTP handling
    "github.com/gorilla/websocket" // Imports the Gorilla WebSocket package
    "sync" // Imports the sync package for synchronization
    "encoding/json" // Imports the json package for JSON handling
    "time" // Imports the time package for time handling
)

// Define a struct for events
type Event struct {
    Type string `json:"type"`
    Data string `json:"data"`
}

// Define a struct for a WebSocket connection
type Client struct {
    Conn *websocket.Conn
    Send chan Event
}

// Global list of connected clients
var clients = make(map[*Client]bool)
var clientsMutex sync.Mutex

// Upgrade HTTP connection to WebSocket
func upgradeConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := websocket.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }

    client := &Client{Conn: conn, Send: make(chan Event)}
    clientsMutex.Lock()
    clients[client] = true
    clientsMutex.Unlock()

    go handleClient(client)
}

// Handle incoming WebSocket messages
func handleClient(client *Client) {
    defer func() {
        clientsMutex.Lock()
        delete(clients, client)
        clientsMutex.Unlock()
        client.Conn.Close()
    }()

    for {
        var msg Event
        if err := client.Conn.ReadJSON(&msg); err != nil {
            fmt.Println("Error reading message:", err)
            return
        }
        // Broadcast event to all connected clients
        clientsMutex.Lock()
        for c := range clients {
            if c != client {
                c.Send <- msg
            }
        }
        clientsMutex.Unlock()
    }
}

// Handle REST API requests
func apiHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        var event Event
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        // Broadcast event to all connected clients
        clientsMutex.Lock()
        for client := range clients {
            client.Send <- event
        }
        clientsMutex.Unlock()

        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Handle WebSocket connection and API server
func main() { // Main function, entry point of the program
    http.HandleFunc("/ws", upgradeConnection)
    http.HandleFunc("/api/event", apiHandler)

    go func() {
        for {
            time.Sleep(time.Second)
            clientsMutex.Lock()
            for client := range clients {
                select {
                case event := <-client.Send:
                    client.Conn.WriteJSON(event)
                default:
                }
            }
            clientsMutex.Unlock()
        }
    }()

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server error:", err)
    }
}
// Microservice with gRPC and Protobuf

package main // Declares the main package

import (
    "context" // Imports the context package for context management
    "fmt" // Imports the fmt package for formatted I/O
    "log" // Imports the log package for logging
    "net" // Imports the net package for network operations
    "google.golang.org/grpc" // Imports the gRPC package
    pb "path/to/your/protobuf" // Imports the generated protobuf package
)

// Define a struct for the gRPC server
type server struct {
    pb.UnimplementedYourServiceServer
}

// Implement the gRPC method
func (s *server) YourMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
    // Process the request and prepare a response
    response := &pb.YourResponse{
        Message: fmt.Sprintf("Hello, %s!", req.Name),
    }
    return response, nil
}

func main() { // Main function, entry point of the program
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterYourServiceServer(s, &server{})

    fmt.Println("gRPC server listening on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
// Microservices Architecture with gRPC and REST

package main // Declares the main package

import (
    "context" // Imports the context package for context management
    "fmt" // Imports the fmt package for formatted I/O
    "log" // Imports the log package for logging
    "net" // Imports the net package for network operations
    "net/http" // Imports the net/http package for HTTP handling
    "google.golang.org/grpc" // Imports the gRPC package
    pb "path/to/your/protobuf" // Imports the generated protobuf package
    "encoding/json" // Imports the json package for JSON handling
)

// Define a struct for the gRPC server
type server struct {
    pb.UnimplementedYourServiceServer
}

// Implement the gRPC method
func (s *server) YourMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
    response := &pb.YourResponse{
        Message: fmt.Sprintf("Hello, %s!", req.Name),
    }
    return response, nil
}

// REST API server
func restAPIHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var request map[string]string
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        name, ok := request["name"]
        if !ok {
            http.Error(w, "Missing 'name' field", http.StatusBadRequest)
            return
        }

        response := map[string]string{
            "message": fmt.Sprintf("Hello, %s!", name),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() { // Main function, entry point of the program
    // Start gRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterYourServiceServer(grpcServer, &server{})

    go func() {
        fmt.Println("gRPC server listening on port 50051")
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("Failed to serve gRPC: %v", err)
        }
    }()

    // Start REST API server
    http.HandleFunc("/api/hello", restAPIHandler)

    fmt.Println("REST API server listening on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to serve HTTP: %v", err)
    }
}
// Event-Driven System with WebSocket and Kafka

package main // Declares the main package

import (
    "bytes" // Imports the bytes package for byte buffer operations
    "context" // Imports the context package for context management
    "encoding/json" // Imports the json package for JSON handling
    "fmt" // Imports the fmt package for formatted I/O
    "log" // Imports the log package for logging
    "net/http" // Imports the net/http package for HTTP handling
    "github.com/gorilla/websocket" // Imports the Gorilla WebSocket package
    "github.com/Shopify/sarama" // Imports the Sarama package for Kafka
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
)

// Define a struct for events
type Event struct {
    Type string `json:"type"`
    Data string `json:"data"`
}

// Define a struct for a WebSocket connection
type Client struct {
    Conn *websocket.Conn
    Send chan Event
}

// Global list of connected clients
var clients = make(map[*Client]bool)
var clientsMutex sync.Mutex

// Kafka producer configuration
var producer sarama.SyncProducer

func init() {
    var err error
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatalf("Failed to create Kafka producer: %v", err)
    }
}

// Upgrade HTTP connection to WebSocket
func upgradeConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := websocket.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }

    client := &Client{Conn: conn, Send: make(chan Event)}
    clientsMutex.Lock()
    clients[client] = true
    clientsMutex.Unlock()

    go handleClient(client)
}

// Handle incoming WebSocket messages
func handleClient(client *Client) {
    defer func() {
        clientsMutex.Lock()
        delete(clients, client)
        clientsMutex.Unlock()
        client.Conn.Close()
    }()

    for {
        var msg Event
        if err := client.Conn.ReadJSON(&msg); err != nil {
            fmt.Println("Error reading message:", err)
            return
        }
        // Send event to Kafka
        msgBytes, _ := json.Marshal(msg)
        _, _, err := producer.SendMessage(&sarama.ProducerMessage{
            Topic: "events",
            Value: sarama.StringEncoder(msgBytes),
        })
        if err != nil {
            fmt.Println("Error sending message to Kafka:", err)
        }

        // Broadcast event to all connected clients
        clientsMutex.Lock()
        for c := range clients {
            c.Send <- msg
        }
        clientsMutex.Unlock()
    }
}

// Handle REST API requests to produce events to Kafka
func apiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var event Event
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        // Send event to Kafka
        eventBytes, _ := json.Marshal(event)
        _, _, err := producer.SendMessage(&sarama.ProducerMessage{
            Topic: "events",
            Value: sarama.StringEncoder(eventBytes),
        })
        if err != nil {
            http.Error(w, "Error sending message to Kafka", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() { // Main function, entry point of the program
    http.HandleFunc("/ws", upgradeConnection)
    http.HandleFunc("/api/event", apiHandler)

    // Start WebSocket and REST API servers
    go func() {
        fmt.Println("WebSocket server listening on port 8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Failed to serve WebSocket: %v", err)
        }
    }()

    fmt.Println("REST API server listening on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Failed to serve REST API: %v", err)
    }
}
// Advanced CLI Tool with Config Parsing and Subcommands

package main // Declares the main package

import (
    "flag" // Imports the flag package for command-line flag parsing
    "fmt" // Imports the fmt package for formatted I/O
    "io/ioutil" // Imports the ioutil package for reading files
    "os" // Imports the os package for OS operations
    "encoding/json" // Imports the json package for JSON handling
)

// Define a struct for configuration
type Config struct {
    ServerAddress string `json:"server_address"`
    LogLevel      string `json:"log_level"`
}

// Load configuration from a file
func loadConfig(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err
    }

    return &config, nil
}

// Define a function for the 'start' subcommand
func startCommand(config *Config) {
    fmt.Printf("Starting server at %s with log level %s\n", config.ServerAddress, config.LogLevel)
    // Implementation for starting the server goes here
}

// Define a function for the 'stop' subcommand
func stopCommand() {
    fmt.Println("Stopping server...")
    // Implementation for stopping the server goes here
}

func main() { // Main function, entry point of the program
    if len(os.Args) < 2 {
        fmt.Println("Usage: cli-tool <command> [options]")
        os.Exit(1)
    }

    // Define subcommands
    startCmd := flag.NewFlagSet("start", flag.ExitOnError)
    stopCmd := flag.NewFlagSet("stop", flag.ExitOnError)
    configFile := startCmd.String("config", "config.json", "Path to the config file")

    switch os.Args[1] {
    case "start":
        startCmd.Parse(os.Args[2:])
        config, err := loadConfig(*configFile)
        if err != nil {
            fmt.Printf("Error loading config: %v\n", err)
            os.Exit(1)
        }
        startCommand(config)
    case "stop":
        stopCmd.Parse(os.Args[2:])
        stopCommand()
    default:
        fmt.Println("Unknown command:", os.Args[1])
        os.Exit(1)
    }
}
// High-Performance File Processor with Custom I/O and Profiling

package main // Declares the main package

import (
    "bufio" // Imports the bufio package for buffered I/O
    "fmt" // Imports the fmt package for formatted I/O
    "os" // Imports the os package for OS operations
    "sync" // Imports the sync package for synchronization
    "time" // Imports the time package for time handling
    "runtime/pprof" // Imports the pprof package for profiling
)

// Define a struct for file processing stats
type FileStats struct {
    LinesProcessed int
    BytesProcessed int
}

// Process file content
func processFile(filePath string, stats *FileStats, wg *sync.WaitGroup) {
    defer wg.Done() // Signal when the goroutine is done

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err.Error() != "EOF" {
                fmt.Printf("Error reading file: %v\n", err)
            }
            break
        }
        stats.LinesProcessed++
        stats.BytesProcessed += len(line)
    }
}

// Main function for file processing
func main() { // Main function, entry point of the program
    if len(os.Args) < 2 {
        fmt.Println("Usage: fileprocessor <file1> <file2> ...")
        os.Exit(1)
    }

    filePaths := os.Args[1:]

    // Set up profiling
    f, err := os.Create("cpu_profile.prof")
    if err != nil {
        fmt.Printf("Error creating profile file: %v\n", err)
        os.Exit(1)
    }
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    var wg sync.WaitGroup
    stats := &FileStats{}
    for _, filePath := range filePaths {
        wg.Add(1)
        go processFile(filePath, stats, &wg)
    }
    wg.Wait()

    fmt.Printf("Processed %d lines and %d bytes\n", stats.LinesProcessed, stats.BytesProcessed)
}
