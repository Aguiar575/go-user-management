# User Management System

This is a simple user management system written in Go. It provides an API to manage user permissions in different contexts. 

## Features

- Get user permissions
- Set new permissions for a user
- Remove existing permissions

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker (optional, for running a database container)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/user-management.git
cd user-management
```

2. Install dependencies:

```sh
go mod tidy
```

3. Create a `.env` file in the root of the project and set the following environment variables:

```
HOST=your_db_host
PORT=your_db_port
USERNAME=your_db_username
PASSWORD=your_db_password
```

4. (Optional) Run a database container using Docker:

```sh
docker run --name user_management_db -e POSTGRES_USER=your_db_username -e POSTGRES_PASSWORD=your_db_password -e POSTGRES_DB=user_management_db -p 5432:5432 -d postgres
```

5. Run the application:

```sh
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Get User Permissions

**GET** `/permissions/{userID}`

Fetches permissions for a user.

#### Request Parameters

- `userID` (path): The ID of the user whose permissions are being fetched.

#### Example

```sh
curl -X GET http://localhost:8080/permissions/1
```

#### Response

```json
[
    {
        "id": 1,
        "userId": 1,
        "contextId": 100,
        "read": true,
        "write": false
    },
    {
        "id": 2,
        "userId": 1,
        "contextId": 101,
        "read": true,
        "write": true
    }
]
```

### Set Permission

**POST** `/permissions`

Sets a new permission for a user.

#### Request Body

```json
{
    "userId": 1,
    "contextId": 100,
    "read": true,
    "write": false
}
```

#### Example

```sh
curl -X POST http://localhost:8080/permissions -H "Content-Type: application/json" -d '{"userId": 1, "contextId": 100, "read": true, "write": false}'
```

#### Response

- `201 Created`

### Remove Permission

**DELETE** `/permissions/{permissionID}`

Removes a permission by ID.

#### Request Parameters

- `permissionID` (path): The ID of the permission to be removed.

#### Example

```sh
curl -X DELETE http://localhost:8080/permissions/1
```

#### Response

- `204 No Content`
