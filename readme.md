**Senior GoLang Developer Test**

Welcome to the technical assessment for the Senior GoLang Developer role at Shaffra. This test is designed to evaluate your proficiency in GoLang, problem-solving skills, debugging capabilities, and your ability to design scalable systems. Please ensure that you follow all instructions carefully.

You will be tested on three core aspects:
1. **GoLang Application Development**
2. **Debugging and Code Review**
3. **Systems Design and Architecture**

### General Instructions:
- You are expected to complete all parts of the test independently.
- Please submit your answers, including the code, test cases, and documentation in a GitHub repository.
- Include a README.md file that explains how to run your code and any additional comments.
- The total time allocated for the test is 8 hours, divided between the three tasks as outlined below. You can distribute the time according to your comfort level.

---

### **Part 1: GoLang Problem-Solving Task**
#### Time Limit: 3-4 hours

**Objective:**  
You will be developing a simple GoLang microservice that handles user management. Your task is to create a RESTful API that is scalable, efficient, and easy to maintain.

**Requirements:**
1. **API Endpoints:**
   - Implement the following user management operations:
     - `POST /users`: Create a new user.
     - `GET /users/{id}`: Retrieve user details by ID.
     - `PUT /users/{id}`: Update user details by ID.
     - `DELETE /users/{id}`: Delete a user by ID.
   - Assume users have a name, email, and age fields.
   
2. **Concurrency and Goroutines:**
   - The API should handle requests concurrently.
   - Use goroutines and channels to log the time each request takes from start to finish. Ensure proper synchronization using `sync.WaitGroup`.

3. **Database Interaction:**
   - Integrate either PostgreSQL or MongoDB (your choice) to persist user data.
   - The API should perform CRUD operations on the database (Create, Read, Update, Delete).
   
4. **Testing:**
   - Write unit tests for each API endpoint and database operation. Aim for high test coverage.

5. **Error Handling:**
   - Implement robust error handling for the API and database interactions. Return appropriate status codes and messages for different scenarios (e.g., user not found, invalid input).

**Deliverables:**
- A GoLang project with a well-structured directory.
- Unit tests with reasonable coverage.
- A README.md explaining how to run the application and tests.

**Evaluation Criteria:**
- Code readability and structure.
- Efficient use of GoLang constructs (goroutines, channels, error handling).
- Clean and well-documented code.
- Effective database interaction and schema design.
- Testing coverage and proper error handling.

---

### **Part 2: Debugging and Problem-Solving (Code Review Task)**
#### Time Limit: 2 hours

**Objective:**  
In this part, we want to test your debugging skills. You will be provided with a GoLang project containing several coding flaws and inefficiencies. Your job is to identify and fix these issues.

**Instructions:**
1. **The GoLang project is in the file:** _buggy_project.go_
2. **Identify Issues:** Review the code and identify at least **three major issues** that could affect performance, security, or scalability. Common examples include:
   - Incorrect use of goroutines, race conditions, or deadlocks.
   - Mismanagement of database connections (e.g., failure to close connections).
   - Poor error handling or lack of logging.

3. **Fix the Issues:** Correct the issues you’ve identified and write a brief explanation (in your README.md file) of:
   - The problem you found.
   - Why it could cause a failure in production.
   - How you fixed it.

**Deliverables:**
- Updated project code with the issues resolved.
- A README.md file that explains your solutions.

**Evaluation Criteria:**
- Ability to identify subtle bugs and inefficiencies.
- Quality of your fixes and the justifications you provide.
- Code quality after the fixes are implemented.


### **Part 3: Systems Design and Architecture**
#### Time Limit: 2 hours

**Objective:**  
You will be required to demonstrate your understanding of scalable system design by proposing a simple microservices architecture for a hypothetical scenario.

**Scenario:**
You have been tasked with designing the backend architecture for an **online marketplace platform**. This platform should allow users to browse products, make purchases, and manage their orders. Your job is to outline a microservices-based architecture that supports scalability and fault tolerance.

**Requirements:**
1. **Services:**
   - Design at least three microservices:
     - User Service (handles user accounts, authentication, etc.).
     - Product Service (manages the product catalog).
     - Order Service (handles order creation, updates, and tracking).
   
2. **Database:**
   - Propose the appropriate database(s) for each microservice. Explain whether you would use SQL or NoSQL databases for each, and justify your choice.

3. **Scaling Considerations:**
   - Discuss how each service can be scaled horizontally. Include suggestions for handling load balancing and service discovery.
   
4. **CI/CD Pipeline:**
   - Briefly explain how you would set up a CI/CD pipeline for these services using tools like Docker, Kubernetes, and any cloud platform (e.g., AWS, GCP, or Azure).
   
**Deliverables:**
- A detailed design document (1-2 pages) that outlines the architecture.
- An explanation of database and scaling choices.
- A brief description of the CI/CD pipeline.

**Evaluation Criteria:**
- Clarity and soundness of your architectural decisions.
- Understanding of microservices architecture.
- Knowledge of cloud platforms, databases, and scalability.
- Practicality of your CI/CD pipeline design.

---

### **Additional Notes:**
- The goal of this test is to assess your problem-solving ability and GoLang expertise. Please ensure that all code is your own and original.
- **We strongly discourage using large language models (LLMs) like ChatGPT to generate code for this test.** We will assess for this, and if detected, it may result in disqualification. Be prepared to discuss and explain all of your solutions in detail during a follow-up interview.

---

We look forward to seeing your solutions and discussing them further with you. Good luck!

--- 
