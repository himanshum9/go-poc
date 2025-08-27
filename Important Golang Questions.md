Important Golang Questions for interview:
Ques:Why Go?
Ans: 1. Highly Scalable.
    2. Statically typed Language.(which makes it fast as no conversions at run time)
    3. Easy to understand.
    4. Directly run on machine code without the need of any virtual machines.
    5. Robust library support.
    6. Open source, so it has a huge community.
    7. Automatic Garbage collection, which makes it memory efficient.
    8. concurreny is simpler.
    9. Cross platform-single binary multiple platform.

Ques:what are string literals?
Ans: not that imp, how strings are basically displayed in go like how it display in `` or "" or \n or \" etc.

Ques: What data types does Golang use?
Ans: Go uses:
Numeric:
    1. int
    2. float
    3. complex number
Boolean
String
Array
Struct
Reference data types:
    slice
    maps
    channels
7. interface type

Que: What are packages in Go?
Ans: Everything we do in go comes under some kind of package,Go is nothing without packages,it helps to collect the similar functionality togethar.Packages are used to use different functionalities into existing projects.
1. Collection of related source files.
2. better reusability,better maintainability.
3. many builtin packages for common functionality
4. every program starts with a main package,which is the entry point.
5. custom packages can be created.

Que: What form of type conversion does Go support?
Ans: Explicit type conversion


Que: Encasulation in Golang?
Ans:Encapsulation is about hiding the details of an object (its data and implementation) and exposing only what’s necessary through methods.
Achieved through by writing small starting letter so that not accessable from outside.
Access through get and set functions.

Que: Abstraction in Go?
Ans: Hiding the complex logic not like hiding data which is encapsulation.
2. Achieved through interfaces
3. fmt.Println(We dont know how internally it is using which code, we are just calling the function this is abstraction)

Que: What is Variadic Function?
Ans:1. Functions that have variable number of arguments passed in the function.
    2. type of all the variable should be same.
    3.Syntax: func myFunc(nums ...int) {
                // nums is a slice of ints
                }
    4. it must be the last variable of the function and will be treated as slice.

Que: What is anonymous function?
Ans:1. Function with no name,often invoked directly.
    2. Syntax: func() {
            fmt.Println("Hello, World!")
            }()
    3. Can be assigned to a variable.
    4. Can be passed as parameter to another function.(In which case we should follow this approach?)
    5. Closure satisfied.

Que: What is deffered function?
Ans: 1. Functions which executes after the surrounding function.
     2. Multiple deffered function execute in LIFO order.
     3. Used for resource cleanup like closing file and unlocking mutex.

Que: What are decorators and can we implement in Golang?
Ans: 1. No built-in functionality,possible through higher order function.
     2. Used as a wrapper to add further functionality.
     3. work well in functional programming.

Que: What is composition in Go?
Ans:1. Combining multiple component together and not only relying on inheritance.
    2. promotes code reuse by embedding types.
    3. Struct inside another struct.
    4. Can combine multiple interface in 1.
    5. Code reusability
    6. Decoupling
    7. Flexibility
    8. Clean Design

Que: What are pointers?
Ans:1. Variable that stores memory address of another variable.
    2. The zero value of pointer is nil.
    3. Pass by Reference

Que: Implement Linkelist in Go?
Ans:1. Check go-poc folder.

Que: Implement Tree in Go?
Ans:1. Check go-poc folder.

Que: How to handle error in Go?
Ans:1. we have error package in go in build standard library
    2. In Go we have built in Error type interface which is used to define error.
    3. syntax var err Error
    4. errors.New("message") for simple error msg
    5. fmt.Errorf("error occurred in %s", "function") for formatted error
    6. Implement error Interface by defining custom Error function.

Que: Panic and recover in Golang?
Ans:1. Also used as error handlening
    2. Should be used carefully as it can halt the application.
    3. use recover to handle panic.
    4. Syntax: defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Recovered from panic:", r)
                    }
                }()
                panic("unexpected error")
    5. If Panic occur defer function will run in LIFO order and will clear the stack and crashes if not recovered.

Que: What is Sentinel Error in Go?
Ans:1.predefined,constant error which are not changing.
    2. var ErrNotFound = errors.New("item not found")

Que: What is Custom error in Go?









Que: What is Error propagation in Go?
Ans:1. practice of returning error from function and handeling it effeciently.
    2. We can wrap error to propagate better.
    3. Help in better error understanding and cleaner code.


Que; How file handleing happen in Go?
Ans:1. We use os built in package to handle file.
    2. os.open to open a file for reading.
    3. os.Create or os.OpenFile for writing in file.
    4. always use defer file.Close after proccess is completed.
    5. ioutil package also used to read the file os.File also read the file.
    6. Use os.File.Write, os.Create, or os.OpenFile to write data to a file.

Que: What if we do not write file.Close at the end?
Ans:1. Resource Leaks:
        file open->os resource assign -> limited os resource->can lead to resource execution.
    2. Data Corruption chances are high.->some packages do not commit to the file till its closing.
    3. File lock issue,no other process would be able to access the file.

Que: How concurrency happens in Golang?
Ans: Multiple tasks making progress independently.

Que: Difference between Concurrency and parallelism?
Ans: Concurrency is managing multiple tasks at once; parallelism is executing multiple tasks simultaneously.


Que: Goroutine in Golang?
Ans:1. A light weight concurrent ,execution managed by go runtime.
    2. Uses M:N scheduling,meaning M goroutine schduled on N os threads.
    3. Syntax: go myFunction()
    Goroutine Life Cycle:
        1.New->Created but not schdulled.
        2.Runnable -> Ready to execute,waiting for the schduler.
        3.Running -> Avtively executing
        4.Blocked -> Waiting for I/O,channel or sync.
        5.Terminated ->Finished Execution.

    Anonymous function as a Go routine:
    Syntax: go func() {
                fmt.Println("Hello from an anonymous goroutine!")
            }()
    4. Main function terminates all the running go routine.
    5. time.Sleep,sync.Waitgroup and channel used for wait in go routine.
    6. Check all the poc created by me.
    7. runtime.NumGoroutine() number of goroutine running.
    8. sync.Mutex and sync/Atomic for race condition handleing
    Go Schduller:
        1. brain for managing go routine.
        2. manages execution of all routine by distributing them across available cpu thread.
        3. uses work steal to balance the load,ideal threads steal work from busy threads.
        4. it has the power to stop the running goroutine and assign the cpu to different process.
        5. GOMAXPROCS:
            1. control the no. of thread available as runtime.
            2. runtime.GOMAXPROCS(4) -> to set number of threads.

Que: WHat role does escapse analysis have compare to Garbage Collector?
Ans: Escape Analysis: Done at compile time to decide if a variable should be allocated on the stack or heap
     Mark and Sweep (Garbage Collector): At runtime, it marks reachable objects and sweeps (frees) unreachable ones from the heap.

Que: Context package in Golang?
Ans:1.Context package used in Concurrent programming,timeout,cancellation and request scoped values.
    2. Controlling Timeouts
    3. Cancelling go routines
    4. passing matadata accross application
    5. Helps you to communicate between go routines also, like if 1 goroutine fails stop other also.

Que: What is Logrus?
Ans:1.A structured logging library for Go.
    2.Syntax: logrus.WithFields(logrus.Fields{
                    "user_id": 123,
                    "status":  "success",
                }).Info("User logged in")

Que: print odd,even,prime number for multithreading?

Que: Consume an Api, How will i fetch the data and display in json?

Que: Underlying data structure in Youtube Playlist?
Ans: Linklist

Que: TDD and other type of Development?
Ans:
TDD (Test-Driven Development): Write tests before writing code. Cycle: Red → Green → Refactor.
BDD (Behavior-Driven Development): Focus on user behavior; tests are written in natural language (often with tools like Ginkgo).
DDD (Domain-Driven Design): Code structure mirrors business domain, useful for complex systems.
FDD (Feature-Driven Development): Build software by planning and implementing features, suitable for larger teams.

Que: SRP?

Que: Ways to connect 2 microservices?
Ans:
    Ways to connect two microservices:

    1. HTTP REST APIs – Most common (e.g., GET /user/1).
    2. gRPC – Fast, efficient binary protocol (Protobuf-based).
    3. Message Brokers – Async via Kafka, RabbitMQ, or NATS.
    4. GraphQL – If aggregation from multiple services is needed.
    5. WebSockets – For real-time bidirectional communication.



Que: Create me a test case?

Que: 


name := "abc", Dont you think this line will introduce mutability problem in go?


one thread in its stack store information related to api request then dont you think if the waiting thread is blocked then again for new thread we need to store information in the new thread stack

16 -> 16 -> blocking 

https://api.coindesk.com/v1/bpi/currentprice.json


Que: What is the difference between interface and generics any?
Ans: any was implemented for new developer as it is easy readable but actually it is an aliasis of interface{}
type any = interface {}
interface{} (empty interface)
    Accepts any type, but you lose type safety.
    Requires type assertion to get the real type.interface{} (empty interface)
    Accepts any type, but you lose type safety.
    Requires type assertion to get the real type.
Generics (any is an alias for interface{} but used in generics)
    Compile-time type safety.
    Avoids type assertions, supports reusable typed functions.

Que: How does Go’s scheduler manage goroutines, and how would you debug a goroutine leak?
Ans: Go scheduler is the heart of go runtime which makes concurrency/parallilsm so easy for the languagee.It is based on M:n schdulling where m go routine are bind to n os threads.there is a processor P which is bind with the OS and has a local queue. If there is enough space in the local queue the go routine sits in that and if it is full it sits in the global queue.The goroutine itself initialized by the pool.The go schduller manages load of each thread,assiging different processor goroutine to different to manage the load.To prevent the goroutine leaks we can check the number of running goroutine before entring the function and after exiting the function.we need to take care of how the goroutine is sending data and how we are recieving it, also pprof is used to catch the goroutine leak

Que: How do we deploy things in production?
Que: How do we manage the deployment?
Que: Reason for migrating from GCP to AWS?
ans: The company is migrating from GCP to AWS primarily for cost optimization, better integration with our existing tools, and to leverage AWS’s broader service offerings.
    
Que: What is package in Go and how can we create our own Package and use it?
Ans: used to place similar functionality together.for custom package create the package folder and include it in main to use

Que: What is scoping in go?
ans: how variables are accessible:
1. package scope -> declared outside function.-> visible to entire package
2. function scope -> Declared inside a function -> accessible only there.
3. Block scope -> Inside if,for etc -> limited to that block.

Que: Factory pattern implement?
Ans: Need to hide the creational logic then we use factory.
Que: Pointer implement?
Que: Goroutine what is it?

Que: If one goroutine fail how will you stop all the other go routine.
Ans: Use context with cancelation to notify other goroutine to stop.

Que: if api is taking so much time how will you handle this situation,how will you design the functionality.
Ans: Will use the context package with timeout.

Que: Defination of Factory Design Pattern?
Ans: it is a creation design pattern,it provides inteface to create object without exposing the instantiation logic.
    It is used when object creation is complex and depend on input.
    Definition: Encapsulates object creation logic in a separate method or struct.
    Purpose: Hides instantiation details and simplifies object creation.
    Use Case: When object creation is complex or varies based on input.
    Benefit: Improves maintainability and reduces code duplication.

Que: Defination of Builder Design Pattern?
Ans: 1. Creational design pattern.
     2. Used when object creation is complex.
     3. Object creation is step by step and creation logic is separate from main object.
     4. Object has multiple configurable fields.


Que: Reflection in Golang?
Ans: Reflection lets you inspect types and values at runtime using the reflect package.
        reflect.TypeOf(x)

Que: why we need _ packages at the first place?
Ans: Because if i use this it will initiate the init function which is mendatory for some package like db package.

Que: How golang handle cyclic package requirements?
Ans: Golang doesn't support cyclic package imports. To solve this, we can extract common logic into a shared package or use interfaces to decouple dependencies.

Que: Why multi threading in go better than java or c# etc.


Que: What are generics and have you used it in your project and how?
Ans: 1. helps to achieve type independent code using type parameter.
    code: func Print[T any](value T) {
                fmt.Println(value)
            }


Que: what are interface in golang?
Ans: Interface in Golang are set to specific method signature that a type can implement.think of it as a contract.
    1. No explicit implement keyword
    2. flexible design
    3. Easy to implement

Que: What is Polymorphism and how it can be achieved?
Ans: poly = many, so a function that can use different types using interface can be called as polymorphism concept.

Que:What are the Best Practices: When to Use An Interface type?
Ans: 1. When multiple type share common behavior for ex: Vehicle interface and car,bike,truck as struct.
     2. When we want to write type independent code ex: Instead of func Save(db *SQLDatabase), use func Save(db Database).
     3. writing mockable code for testing purpose to not invoke the dependency.
     4. When we want to achieve polymorphism.



When to Use What?
Keyword	            Works With	                        Returns                             Example Usage
new	                Pointers (struct, int, etc.)	    Pointer	                            p := new(int)
make	            slice, map, channel	                Value	                            m := make(map[string]int)
var	                All types	                        Zero value	                        var x int (0)
:=	                All types (inside functions)	    Initialized value	                x := 10

Best Practice: Use make for slices/maps/channels, new rarely, and prefer := for regular variables.
new is less flexible.
new gives pointer to a zero value.
&T{} create pointer to a literal struct initialized with zero values


What happens when you close a channel that's already closed? What about sending/receiving on a closed channel?
Ans: 1. Closing a already closed channel will give panic.
Sending data to a closed channel will also panic.But recieving data from a closed channel will not panic but return 0 value.



Que: How does go handle method overloading?
Ans: Go does not support the conventional way of overloading like we have in java,c++.
but in go we can use iterface, variadic function that can help you achieve the similar behavior.

Que: What is type embedding in Go and when would you use it?
Ans: It is a way to achieve the composition in go as go does  not support inheritance.
    2. Struct inside struct,interface inside a struct

Que: Can we have multiple init function in a file?
Ans: Yes, there can be multiple init function in a file.

Que: What type of authentication does your company implemented?
Ans:I haven’t worked on implementing the authentication and authorization myself, but I’ve interacted with it closely. In our Go backend, we use JWTs for authentication. When users log in, a token is issued and passed in the Authorization header as a Bearer token—something I regularly work with while testing APIs in Postman or integrating frontend requests.

Que: What is JWT token?
Ans: json web token, used for authentication and authorization.It has 3 parts:
1. Header: token type and signing algorithm
2. Payload: payload consist of claims and claims tells about entity like user,employee etc
3. Signature: created with the use of encoded header,payload and secret algo mentioned in header.

The backend validates this token on each request. If the token is valid and not expired, the request is allowed; otherwise, it's rejected. For authorization, we follow a Role-Based Access Control model, so different user roles determine access to certain API endpoints or frontend views. On the frontend, built in Vue.js, we use navigation guards to control access to routes based on user roles or login status.

So, while I haven't developed the system myself, I understand how it works and how to work with it effectively in both backend and frontend contexts.


Que: VAlue types in which a copy is passed in a function:
    1. int
    2. float
    3. bool
    4. string
    5. struct
    6. Array (not slices)
    Referance type:
    1. Slice
    2. Map
    3. Channel
    4. function
    5. interface
    6. Pointer
    
    
    
Concurrency in Go vs other languages?
Parallelism in Go?
Binary in Go?
How Go executable binaries run
Types in Golang?
What is new in Go?When to use new?
Interface in Golang?
Where to use interfaces in Go?
Special use cases for interfaces in Go?
Generics in Go?
Inheritance in Go - code example?
Struct embedding in Go?
Two goroutines: publisher generates 1 to 10, subscriber prints those numbers (did not allow other packages than fmt)
Dockerfile example?
Where is the Docker image built?
Multi-stage Dockerfile?
Mocking in Go and why to mock
Fan in fan out workerpool 1-1000 print krvaana hai go routine mutux locking. Race condition. Mutual exclusion
2 go routine 2
Singles pass in channels using struct. 
Routing patterns?
Worker pool
Microservice go
Dbr v2
What are go.mod and go.sum files?
Defer, Panic and Recover?
Goroutines & channels?
Implement a custom middleware for a Go HTTP server that performs advanced request logging and modifies responses based on certain conditions/ The middleware should be configurable and support different logging levels (eg. info, debug, error, etc.).
 Provide a test case demonstrating its functionality.
Program to remove zeros from the input array.
Explain go memory management.
Explain error wrapping in go?
How do you compare two structures (Hint: reflect pkg)
Write a golang package that stores KeyValueStore and write methods like Get, Set and Delete using mutex?
Arr: [1, 3, 5, 7, 10, 2, 4] Product: 70. Give the indexes of the two integers that equals the product?
Develop rest apis, where you forward the request towards dummy open source endpoints to fetch json response?
Design and develop an application for payments gateway where you should use Payment interface with Pay method and provide different methods like UPI, credit card and give discount as per payment method.

Que: What is Semophore in golang, have you heard about it?



Que: What is closure in go?
Ans: A closure in go is function that can capture variable from its outer function and can store itself even when the outer function is already executed.

strconv.Itoa() - > integer to string
strconv.Atoi()  -> string to integer
strconv.ParseFloat().  -> str to float

-------Strings package function
str := "    apple,  orange,banana,orange.  "


parts := strings.Split(str,",").  -> convert string to array
strings.Count(str,"orange"). - > counting occurance in strings
strings.TrimSpace(str).  -> trim the whitespace before and after text
strings.Join([]string{str1,str2}, " ").  -> concat 2 strings

-------- time package function:
curr := time.Now()   -> will give the current time (with extra data like utc etc)
curr.Format("2006-01-02 15:04:05") - > formated time in string
val := curr.Format("2006-01-02 3:04 PM"). -> for 12 hr time in string
time.Parse("2006-01-02",val) - > convert it to time.Time
time.Sleep(4* time.Second)  -> add sleep time


---------- files operation package os and io/ioutil
file,err := os.Create("example.txt") -> it will create new file or truncate existing
out,err := io.WriteString(file, "This is the strind") - >it will add string in file
file.Close() - > always when done working with file
buff := make([]byte, 1024)
for {
n,err := file.Read(buff). - > read the file content into buffer
if err == io.EOF{
    break;
}
fmt.Println(string(buff[:n]))
}

####if we do not want to buffer read we can completly read the data using ioutil.ReadFile function

content,err := ioutil.ReadFile("example.txt")-> read the file and store byte slice []byte

fmt.Println(string(content))


net/url Package -> For everything related to url's

myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

parsedUrl,err:= url.Parse(myUrl)
parsedUrl.Scheme -> will give me http/https
parsedUrl.Host - > will give me hostname
parsedUrl.Path - > will give me PATH of the url




Que: WHat is Rest Api?
Ans: Mechanism for communication between services, it is based on http protocols and uses mainly 4 methods(get,post,put,delete),it uses url to identify resource and it is stateless in nature(each request independent of future and past)


Que. Explain the deployment process that your company follows?
Ans: All steps Regarding deployment process:
1.    A feature branch is created → code is developed and tested there.
2.    A release (rel) branch is created → the final code is merged into it.
3.    A Git tag (e.g., v1.54.0) is created and linked to the rel branch.
4.    The GitHub tag is published → this triggers CircleCI via webhook integration with github.
5.    CircleCI builds Go binaries for all backend services.
6.    CircleCI builds a Docker image tagged with the Git SHA.
7.    The image is pushed to quay.io → this SHA is the "commit hash".
8.    CircleCI then clones the terraform repo.
9.    CircleCI writes that commit hash to the backend_image_hash_prod.auto.tfvars file.
10.    This commit triggers Terraform Cloud (TFC).
11.    TFC runs terraform plan and apply.
12.    This updates Kubernetes deployments (e.g., worker pods etc) in production.


Que: how go garbage collection works?
Ans: Used Mark and Sweep algorithm,Mark all the unreachable variables that are not needed and sweep them ,it works concurrently with the program, so very minimum to nill down time is faced,it also automatically keeps the longer vairable into heap from stack so that it does not have to check everytime.


Que: Hoe to decode in json?
Ans: json.NewDecoder(r.Body).Decode(&myStruct)
    json.UnMarshal(io.ReadAll(r.Body),&mystruct)





// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	str := "Hello, 世界"
	leng := len(str)
	out := []rune{}
	for i := leng - 1; i >= 0; i-- {
		//fmt.println(str[i])
		out = append(out, rune(str[i]))
	}
	fmt.Println(string(out))
}

Que: Project Folder structure?
Ans:
| Folder         |Purpose                                                                   |
| -------------- | ------------------------------------------------------------------------ |
| `cmd/`         | App entry points (`main.go`) — one per service/microservice              |
| `internal/`    | Core app logic (handlers, services, repos, etc.) — **private to module** |
| `pkg/`         | Reusable code shared across projects                                     |
| `configs/`     | Config files and config loaders                                          |
| `db/`          | Migrations, seeders, and SQL helpers                                     |
| `deployments/` | Docker/K8s configs                                                       |
| `build/`       | CI/CD scripts, build tools                                               |
| `docs/`        | Swagger/OpenAPI or architecture docs                                     |
| `.env`         | App environment variables                                                |
| `Makefile`     | CLI for build/test/run convenience                                       |


Que: What is Software Development Life Cycle and how do you estimate the time for completion of some task?
Ans: A story point estimates the effort, complexity, and risk of a task to help plan work without tying it directly to time.
-> SDLC a structured process for planning, creating, testing, and maintaining software.
The SDLC process typically includes: Requirement Gathering → Design → Development → Testing → Deployment → Maintenance.

Que: how many containers can I run inside of a pod?
Ans: You can run multiple containers inside a pod, but typically 2–3 is ideal; technically there's no strict limit, but it's bounded by node resources (CPU, memory).

Que: in docker file I have keywords like cmd and I have keyword run. Basically both of them are executing the command. What's the difference?
Ans: RUN creates a new layer in the Docker image during build, while CMD defines the default command that runs when the container starts.

Que: imagine that you want to build an image and push it to docker registry, right? And what stages should your build job have?
Ans: Code Checkout → linting-> Build Image → Run Tests → Tag Image → Push to Registry → Deploy

Que: What is SRP(Single Responsiblity Principle)?
Ans: SRP (Single Responsibility Principle) means a class should have only one reason to change—i.e., it should do only one job.

Que: What is OCP(Open Closed Principle)?
Ans: You should be able to add new features (extend) without altering existing, tested code (closed to modification).
Often done by interfaces and composition in golang.

Que: What is LSP (Liskov Substitution Principle)?
Ans: When you replace the parent object with the child one, nothing should break.
    Subclasses must behave so that anywhere you use a base class, the subclass can stand in without breaking things.

Que: What is Interface Segregation Principle?
Ans: means don’t force a class to implement interfaces it doesn’t use—split large interfaces into smaller, specific ones.

Que: What is DIP (Dependency Inversion Principle)?
Ans: High-level modules and low-level modules both depend on abstractions (interfaces), not on each other’s concrete details.

Que: what is the major idea behind microservices?
Ans: The major idea behind microservices is to build an application as a collection of small, independent services—each focused on one business function and able to be developed, deployed, and scaled separately.


Que: How to sort value of map using decorator function?

Implementing SOLID Principles in Golang | by Anshul Prakash | FAUN — Developer Community 🐾

https://go.dev/tour/welcome/1

https://learn.innoskrit.in/blog/print-odd-even-numbers-synchronously-using-goroutines-and-channels-in-go/

https://leetcode.com/problems/number-of-islands/description/

https://dev.to/ankitmalikg/create-custom-errors-in-golang-3mdl

https://gobyexample.com/interfaces

https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6



Que: How to design a good Rest Api?
Ans: It should be consistant, easy to consume and maintain, proper http methods and protocols should be used,error should be gracefully handled,Pagination should also be handled,If required version should be done.


Que: What is the Zero Value in Go?
Ans: In Go, the zero value is the default value assigned to a variable when no explicit value is given.
It ensures that every declared variable is initialized, even without a constructor or assignment.


Que: From Linux, let's say you have to deploy a binary to a Windows system. Can you create a binary from your Ubuntu system that can run in the Windows system?
Ans: GOOS=windows GOARCH=amd64 go build -o app.exe main.go
We use env variable like the above to define the target and architecture.

Que: Can you tell me what a story point measures?
Ans: Complexity, Amount of work and uncertainity.

Que: What protections and preventions do we need to put in our code against OWASP Top 10 attacks, specifically the 30% that are code-level?
Ans: 1. Injections (use prepared statement,avoid string concatination)
    2. Broken Access Control(auth check at business logic,use middleware,decorator)
    3. Cross site scripting(sanitize user input while rendering,can use templating engine)
    4. Insecure Deserialisation(use strict schema during unmarshalling)
    5. Security misconfiguration.(keep internal endoint safe,set secure headers,close connection when fail)
    

........................Top 10 GraphQl interview Questions....................


Que1. What is Graphql and how it is different?
Ans: It is a query language for Api's that let client request only the data they need.
    But in Rest, it uses a single endpoint to fetch the data.

Que2: Explain the structure of a basic GraphQL query.
Ans: A basic query includes a root field and the exact fields you want:
                {
                    user(id: "1") {
                        name
                        email
                    }
                }

Que3: What are resolvers in Graphql and how do they work?
Ans: Resolvers are nothing but just a function that fetch data for each field mentioned in the query.

Que4: What are the key component in Graphql?
Ans: Types(like users)-> Queries(read operations) -> Mutations(write operations) -> Resolvers(connect schema to source)

Que5: How does graphql handle versioning to Rest?
Ans: In Rest, to manage versioning we create new endpoint and include versioning in it. but in Graphql versioning is avoided because data is fetched based on the query and without versioning this problem can be resolved and also Deprecated field can remain without breaking existing clients.Old field which are not needed can be marked as deprecated so that client can have some time for migrating to new one.

Que6: What are mutations and how they are different from queries?
Ans: Mutations change data like create or update but queries are used for read operations, Mutations can have side effects.
    mutations {
        AddUser(name:"John") {
            id
        }
    }

Que 7. What are subscription in Graphql and when would you use them?
Ans: Subscriptions are used for real time update using websockets.Use them to update chats or update status at real time.

Que 8: How do you implement authentication and authorization in Graphql?
Ans: Use Middleware or context to pass userInfr from JWT tokens.Resolvers will then check permission before returning data.

Que 9: How do you handle errors in graphql?
Ans: Graphql separates data and error in the response. You can throw custom error in resolvers also and catch them later.

Que9: What are some performance pitfalls of graphql and how do you mitigate them?
Ans:  performance pitfalls can be improved using batched proccessing.
        Complex query limit.
        Query abuse(over feetching) -> only allow pre approved queries
















........................Top 10 GRPC interview Questions....................

Que1: What are GRPC and how does it work?
Ans1: grpc is high performance open source rpc framework by Google that uses http/2 and protocol buffers for fast and efficient communication between services.

Que2: How grpc different from rest?
Ans:  Grpc uses Http2.0 (support multiplexing/streaming)
      uses binary proto buffs for communication instead of json.
      strongly types contracts via .proto files

Que3: What are protocol Buffer and how they are used in GRPC?
Ans: it is a language neutral binary serialization format. grpc uses .proto files to define services and messages then generate client/server code in multiple language.

Que4: Explain the types of grpc communication methods?
ans: Unary (single request-> single response)
    Server Streaming (single request -> Streamed response)
    client Streaming (streamed request -> single response)
    bidirectional streaming (both side streams)

Que 5: How does grpc handle service defination and code generation?
Ans:  Define service and messages in .proto files, then use protoc compiler to generate server and client stubs in your desired language.

Ues 6:  What are the advantages and disadvantages of gRPC?
        ✅ Fast (binary), efficient, supports streaming, strongly typed, cross-language
        ❌ Harder to debug (binary), not well-suited for browser clients directly (needs proxy)

Que 7: How do you implement authentication and security in gRPC?
Ans:    Use SSL/TLS for transport encryption.
        Use metadata headers for passing tokens (e.g., JWT).
        Integrate with OAuth2 or mTLS for identity verification.

Que 8: What is deadline and cancellation in gRPC?
Ans:    Deadlines define a timeout for requests.
        Clients or servers can cancel requests using context (e.g., in Go: ctx, cancel := context.WithTimeout(...)).

Que 9: How do you handle error handling and status codes in gRPC?
Ans: gRPC uses status codes (codes.*) like codes.NotFound, codes.Internal, etc., with structured error messages. Return them using language-specific gRPC helpers.

Que 10: What are some common use cases or limitations of gRPC?
Ans:    ✅ Microservices, internal service-to-service communication, real-time streaming
        ❌ Not ideal for browser-based apps without a proxy (e.g., Envoy), limited support for human-readable debugging.




........................Top 10 Docker interview Questions....................


Que: What is docker and how does it differ from Virtual machine?
Ans: Docker is a containerisation platform that packages apps with dependencies in isolated containers, is shares the underline os host and are lightweight and faster.

Que2: What is a Dockerfile and what are its key instructions?
Ans: A dockerfile is a script to build docker image.
        FROM -> base image
        RUN -> Execute commands
        COPY/ADD -> Adding files
        CMD/ENTRYPOINT - > default command
        EXPOSE -> For exposing ports

Que3: What is the difference between docker image and docker container?
Ans: Image -> Blueprint of all the code,libraries and configs
     Container - > Running Instance of the image(or where images are run)

Que4: How does docker networking work?
Ans: Docker creates a default bridge network where containers can communicate.
-> via Bridge
-> using user defined networks(auto dns)
-> or using host

Que 5: What is docker compose and when would you use it?
Ans: Docker compose is used for managing multi container app via docker-compose.yml files, Ideal for managing services like app + db + cache together.

Que 6: How do you optimize docker  image size?
    -> Use minimal base image
    -> Combine run Commands(each run adds new layer so combining them reduces layer)
    -> Avoid Unneccessary Files
    -> clean cache

Que7: What are volumes in docker and why they are used?
Ans: Volumes are used to persist data outside the container,They are stored on the host and can survive restart, rebuilds or deletion.

Que 8: How do you handle env variable in docker?
Ans: 

Que9: WHat is the role of entrypoint and cmd in dockerfile?
Ans: CMD -> default argument can be overridden
     Entrypoint -> fixed executable that can not be changed

Que10: How do you secure docker containers in production?
Ans: Use minimal images
     Run as non root
     keep docker and dependecies up to date.
     enable resource limit



Que: Syntax of Db connect?
Que: Syntax of RabbitMq connect?
Que: Design a rate Limitor?
Que: Full detail understanding of authentication and authorization?
Que: Implement Redis?(LRU etc)
Ans: https://youtu.be/UpzaB6bX7tc?si=RvrAsqW7MhBH3Inx(all redis related problems)
que. Implement gin for rest api's?
Que. Implement Middleware?
Que. Purpose of circuit breaker and how to implement that?



Technical Interviewer
13:06
Describe how Go handles error handling with the use of multiple return values and the error interface. What are the benefits of this approach?
Technical Interviewer
13:09
In a concurrent scenario, how can you detect and handle deadlocks in Go? What strategies or techniques can be employed to avoid or mitigate deadlocks?
Technical Interviewer
13:10
How does Go manage dependencies with its module system, and what are some best practices for versioning and dependency management in Go projects?
Technical Interviewer
13:13
Can you explain the difference between synchronous and asynchronous communication in the context of distributed systems?
Technical Interviewer
13:15
How do you ensure fault tolerance in a microservices architecture?
Technical Interviewer
13:16
What are the advantages of using containerization (e.g., Docker) in microservices architecture?
Technical Interviewer
13:18
What’s the difference between EC2, ECS/Fargate, and Lambda?
Technical Interviewer
13:19
You deployed an RDS instance without enabling backups. How do you enable backups without downtime?
Explain the difference between SNS and SQS, and when to use each.
Technical Interviewer
13:21
You need to allow only a specific set of IP addresses to access an API Gateway endpoint. How would you do it?
Technical Interviewer
13:24
Write a Go HTTP server with one endpoint /greet?name=John that responds with:
Hello, John!
If the name query parameter is missing, return "Hello, Guest!".
https://www.programiz.com/golang/online-compiler/
vkj-dage-wbr





------------------------Kubernetes All Questions for interview---------------------------


Que: What is Kubernetes (k8s) and why?
ans: Container orchestrator.
    Manages-> deployments, scaling, auto healing, load balancing of containers etc.
    think of it as an os for containers.

Que: What happens when you run kubectl apply -f deployment.yaml?
Ans: yaml file goes to api server, stored in etcd clustered memory,-> controller notice the desired state-> schduler finds the node and kubelet start the pod.

Que: Pod vs ReplicaSet vs Deployment vs StatefulSet?
Ans: Pod - > Single running app container.(Eg -> Nginx pod)
     ReplicaSet -> Makes sure n copies of the pod should run.
     Deployment -> manages Replicaset and allow updates/rollback
     StatefulSet - > similar to deployment but manages storage and pod names and make them fixed.


gcloud container clusters get-credentials CLUSTER_NAME --region REGION
|^for authentication

kubectl config get-clusters 

Que: 











Employees(id, name, dept_id, salary, manager_id)
Departments(id, name)

second highest salary
SLEECT MAX(SALARY) FROM employees WHERE salary < (SELECT MAX(salary) from employyes)

Find employees who earn more than their manager.

SELECT * FROM employees e
JOIN employees m on e.manager_id = m.id
WHERE e.salary > m.salary

Find duplicate employee names.

SELECT name, count(*) FROM employees group by name HAVING count(*) > 1;

Get department with the highest average salary.

SELECT department,avg(salary) as avgSal from employee JOIN
department on department.id = employee.id
WHERE Group by employee.name order by avgSal DESC LIMIT 1;

Find nth highest salary (say 3rd highest).

SELECT salary From employee WHERE salary = (SELECT DISTINCT(salary) from employee ORDER by salary DESC LIMIT 1 OFFSET 2);

List employees who don’t belong to any department.

SELECT * FROM employee e 
left join department d on d.id = e.dept_id 
WHERE e.id IS NULL;

Show employees with the highest salary in each department.









Find departments with no employees.
