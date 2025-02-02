# create user with error (missing required fields like firstname, lastname)
POST http://localhost:3005/api/v1/users
{
    "email": "john@example.com",
    "password": "password1"
}
HTTP 400
[Asserts]
jsonpath "$.error" == "database validation failed"



# Create first user
POST http://localhost:3005/api/v1/users
{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password1"
}
HTTP 201
[Captures]
user1_id: jsonpath "$.id"

# Create second user
POST http://localhost:3005/api/v1/users
{
    "first_name": "Jane",
    "last_name": "Smith",
    "email": "jane@example.com",
    "password": "password2"
}
HTTP 201
[Captures]
user2_id: jsonpath "$.id"

# List all users (authenticated)
GET http://localhost:3005/api/v1/users
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.data" count == 2
jsonpath "$.count" == 2
jsonpath "$.page" == 1
jsonpath "$.per_page" == 10
jsonpath "$.total" == 2
jsonpath "$.data[*].first_name" includes "John"
jsonpath "$.data[*].first_name" includes "Jane"

# list all users with filter
GET http://localhost:3005/api/v1/users?first_name=John
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.data" count == 1
jsonpath "$.count" == 1
jsonpath "$.page" == 1
jsonpath "$.per_page" == 10
jsonpath "$.total" == 1
jsonpath "$.data[0].first_name" == "John"


# Get specific user (authenticated)
GET http://localhost:3005/api/v1/users/{{user1_id}}
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.first_name" == "John"
jsonpath "$.last_name" == "Doe"
jsonpath "$.email" == "john@example.com"


# Update user (authenticated)
PATCH http://localhost:3005/api/v1/users/{{user1_id}}
[BasicAuth]
john@example.com: password1
{
    "first_name": "Johnny"
}
HTTP 200
[Asserts]
jsonpath "$.first_name" == "Johnny"
jsonpath "$.id" == "{{user1_id}}"

# Verify update with GET request
GET http://localhost:3005/api/v1/users/{{user1_id}}
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.first_name" == "Johnny"

# create a topic linked to user john
POST http://localhost:3005/api/v1/topics
[BasicAuth]
john@example.com: password1
{
    "name": "javascript",
    "description": "a topic about your fav programming language",
    "user_id": "{{user1_id}}"
}
HTTP 201
[Captures]
topic1_id: jsonpath "$.id"
[Asserts]
jsonpath "$.name" == "javascript"
jsonpath "$.description" == "a topic about your fav programming language"
jsonpath "$.user_id" == "{{user1_id}}"

# get the topic
GET http://localhost:3005/api/v1/topic/{{topic1_id}}
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.name" == "John"
jsonpath "$.description" == "Doe"
jsonpath "$.user_id" == "{{user1_id}}"
jsonpath "$.id" == "{{topic1_id}}"

# update the topic
PATCH http://localhost:3005/api/v1/topic/{{topic_id}}
[BasicAuth]
john@example.com: password1
{
    "name": "javascript",
    "description": "a topic about your favorite programming language",
    "user_id": "{{user1_id}}"
}
HTTP 201
[Asserts]
jsonpath "$.name" == "javascript"
jsonpath "$.description" == "a topic about your favorite programming language"
jsonpath "$.user_id" == "{{user1_id}}"

# get all topics
GET http://localhost:3005/api/v1/topics
[BasicAuth]
john@example.com: password1
HTTP 200
[Asserts]
jsonpath "$.data" count == 1
jsonpath "$.count" == 1
jsonpath "$.page" == 1
jsonpath "$.per_page" == 10
jsonpath "$.total" == 1
jsonpath "$.data[0].name" == "javascript"
jsonpath "$.data[0].description" == "a topic about your favorite programming language"
jsonpath "$.data[0].user_id" == "{{user1_id}}"

# delete the topic
DELETE http://localhost:3005/api/v1/topics/{{topic1_id}}
[BasicAuth]
john@example.com: password1
HTTP 204

# Delete all users (authenticated)
DELETE http://localhost:3005/api/v1/users/{{user2_id}}
[BasicAuth]
john@example.com: password1
HTTP 204

DELETE http://localhost:3005/api/v1/users/{{user1_id}}
[BasicAuth]
john@example.com: password1
HTTP 204

