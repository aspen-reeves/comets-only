# api documentation
This directory contains the API documentation for the project.

## API Functions
- {host}/ - returns a landing page

- {host}/auth - call with a POST request to authenticate a user
    - request body: `{"username": string, "password": string}`
    - response body: `{"token": int}`
    - if password is wrong, response body is `{"wrong password"}`
    - if there is no user, it creates a new user and returns a token, `{"token": int}`, as well as a `{false}` boolean

- {host}/profile/{token} - call with a GET request to get a user's profile at the token
    - response body: 
    ```
    {
        id: int,
        name: string,
        age: int,
        lang: string,
        os: string,
        editor: string,
        lastShower: string,
        code: string
    }
    ```
    - if there is no user, response body is `"Profile not found"`

- {host}/signup - call with a POST request to create a new user
    - request body: 
    ```
    {
        "name": string,
        "age": int,
        "lang": string,
        "os": string,
        "editor": string,
        "lastShower": string,
        "code": string
    }
    ```
    - response body: 
    ```
    {
        "id": int,
        "name": string,
        "age": int,
        "lang": string,
        "os": string,
        "editor": string,
        "lastShower": string,
        "code": string
    }
    ```

- {host}/checkmatches - call with a POST request to check for matches
    - request body: `{int}`
    - response body: `{id1: int, id2: int, isMatch: boolean}`

- {host}/match - call with a POST request to create/edit a match, 
    - request body: `{id1: int, id2: int}`
    - no response body

- {host}/getbitches - call with a GET request to get a random user
    - response body: ```{id: int, name: string, age: int, lang: string, os: string, editor: string, lastShower: string, code: string}```
