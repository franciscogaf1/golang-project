# golang-project

This project is the final exercise of the golang bootcamp

## Description

### Golang Development Program - Level 7

**Statement**

You are to design the backend side of a system for the following business idea.

We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference. We will only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer. In essence, each question can only have 0 or 1 answer.

The backend should support the following operations:

-   Get one question by its ID
    
-   Get a list of all questions
    
-   Get all the questions created by a given user
    
-   Create a new question
    
-   Update an existing question (the statement and/or the answer)
    
-   Delete an existing question
    

No user tracking or security needed for this version.

Database design is up to you.

We would like to receive code that runs, so remember to focus on the MVP functionality. You can document what’s missing that you wish you had more time for? Please think about the different problems you might encounter if the business idea is successful. This would include considerations such as increased load, increased data, and an upvoting feature.

## Prerequisites

### Mandatory
- Docker
- Golang
- Postman (or any Restful API testing tool)

### Optional
- DBeaver (or any database tool that has postgresql driver)

## How to run

1. Clone this Project
2. Create `golang-app` docker image
	- navigate to `path/to/project/golang-project`
	- execute `docker build -t golang-app .`
3. Run the containers using *docker-compose.yml*
	- `docker-compose up -d`
4. Verify if the containers are running
	- `docker ps`
5. Container address
	- golang-app: `localhost:8080`
	- postgres: `localhost:5432`

## Endpoints

### users

- `localhost:8080/users`
	- Accepts: GET, POST
	- GET: Request all Users
		- Payload: `[{id: userId, name: "username"}]`
	- POST: Create a new User
		- Payload: `{name:"username"}`
- `localhost:8080/users/{id}`
	- Accepts: GET, PUT, PATCH, DELETE
	- GET: Request particular User 
		- Payload: `{id: userId, name: "username"}`
	- PUT / PATCH: Update User
		- Payload: `{name:"username"}`
	- DELETE: Delete User
- `localhost:8080/users/{id}/questions`
	- Accepts: GET
	- GET: Request particular User 
		- Payload: `[{id: questionId, question: "question", answer:"answer", "userId": userId}]`

### questions

- `localhost:8080/questions`
	- Accepts: GET, POST
	- GET: Request all Questions
		- Payload: `[{id: questionId, question: "question", answer:"answer", "userId": userId}]`
	- POST: Create a new Question
		- Payload: `{question: "question", answer:"answer", "userId": userId}`
- `localhost:8080/questions/{id}`
	- Accepts: GET, PUT, PATCH, DELETE
	- GET: Request particular Question 
		- Payload: `{question: "question", answer:"answer", "userId": userId}`
	- PUT / PATCH: Update Question
		- Payload: `{question: "question", answer:"answer", "userId": userId}`
	- DELETE: Delete Question

## Improvements
Since the purpose of this exercise is to provide a MVP, the project needs some improvements regarding the code and execution. These improvements may be done in the future and don't add any new functionalities to the project.

1. Clean and organize the code properly: There are a lot of repeating lines and duplicated code.
2. There is a bug that might occur: When a request throws an error handled by the code, the next one might fail and crash the app. Since the restart policy of the container is set to `always`, the container restarts and the application continues to run normally.
3. Develop unit and maybe integration tests.