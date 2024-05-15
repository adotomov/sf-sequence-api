# sf-sequence-api

sf-sequence-api is a RESTful API service written in Go language that utilizes the Gorilla Mux router and PostgreSQL to manage email sequences. This API allows users to create, retrieve, update, and delete email sequences and their corresponding steps.

## Installation

To run the sf-sequence-api locally, you need to have Go and PostgreSQL installed on your system.

1. Clone the repository:

   ```
   git clone https://github.com/your-username/sf-sequence-api.git
   ```

2. Navigate to the project directory:

   ```
   cd sf-sequence-api
   ```

3. Install dependencies:

   ```
   go mod tidy
   ```

4. Set up your PostgreSQL database and configure the connection details in `config.go` file.

5. Run the tests

   ```
   make test
   ```

6. Run the application:
   ```
   make run
   ```

## Endpoints

### GET /sequences

Fetches all sequences from the database.

#### Response

- Status Code: 200 OK
- Response Body:
  ```
  [
      {
          "id": 1,
          "name": "Welcome Sequence",
          "openTrackingEnabled": true,
          "clickTrackingEnabled": true,
          "stepsInterval": 1
      },
      {
          "id": 2,
          "name": "Product Launch Sequence",
          "openTrackingEnabled": true,
          "clickTrackingEnabled": true,
          "stepsInterval": 2
      },
      ...
  ]
  ```

### GET /sequence?sequenceId={id}

Fetches a single sequence by its ID.

#### Parameters

- sequenceId: The ID of the sequence to fetch.

#### Response

- Status Code: 200 OK
- Response Body:
  ```
  {
      "id": 1,
      "name": "Welcome Sequence",
      "openTrackingEnabled": true,
      "clickTrackingEnabled": true,
      "stepsInterval": 1
  }
  ```

### POST /sequence

Creates a sequence.

#### Request Body

- Name (string): The name of the sequence.
- OpenTrackingEnabled (boolean): Indicates whether open tracking is enabled for the sequence.
- ClickTrackingEnabled (boolean): Indicates whether click tracking is enabled for the sequence.
- StepsInterval (integer): The interval between each step in the sequence.

#### Example Request Body

```
{
    "name": "Abandoned Cart Sequence",
    "openTrackingEnabled": true,
    "clickTrackingEnabled": true,
    "stepsInterval": 3
}
```

#### Response

- Status Code: 201 Created
- Response Body:
  ```
  {
      "id": 3,
      "name": "Abandoned Cart Sequence",
      "openTrackingEnabled": true,
      "clickTrackingEnabled": true,
      "stepsInterval": 3
  }
  ```

### POST /step

Creates a step.

#### Request Body

- sequenceId (string): The id of the sequence the step belongs to.
- emailSubject (string): The subject of the email
- emailContent (string): The content of the email

#### Example Request Body

```
{
	    "sequenceId": "d934f39e-6d40-406f-8b75-dcea99e4e5c1",
	    "emailSubject": "Test Subject One",
	    "emailContent": "Test Content One"
}
```

#### Response

- Status Code: 201 Created
- Response Body:
  ```
  {
      "id": 3
  }
  ```

### PUT /step

Updates a step.

#### Request Body

- id (string): The id of the step.
- emailSubject (string): The value to update with
- emailContent (string): The value to update with

#### Example Request Body

```
{
	    "id": "4f6e9fb9-af39-4199-91ec-f9606b4af790",
	    "emailSubject": "Updated Subject",
	    "emailContent": "Updated content"
}
```

#### Response

- Status Code: 200 OK
- Response Body:
  ```
  {
      "id": "4f6e9fb9-af39-4199-91ec-f9606b4af790"
  }
  ```

### DELETE /step?stepId={id}

Deletes a step by its ID

#### Parameters

- stepId: The ID of the step to delete.

#### Response

- Status Code: 200 OK

## Contributors

- Your Name <adrian.tomov@live.com>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```

```
