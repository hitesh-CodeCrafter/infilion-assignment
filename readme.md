# Assignment

## Overview
This is a web application built with the Go programming language using the Gin framework. This application includes basic routing setup and a health check endpoint.
Prerequisites

    Go (1.16 or higher)
    Gin framework

## Installation

    Clone the repository:

    bash

git clone https://github.com/hitesh-CodeCrafter/infilion-assignment.git

Navigate to the project directory:

bash

cd myapp

Install dependencies:

bash

    go mod tidy

Running the Application

    Start the application:

    bash

    go run main.go

    Access the application:
        Health check endpoint: http://localhost:8080/
        Get person info: GET /person/:person_id/info
        Create person: POST /person/create

## Endpoints
Health Check

    URL: /

    Method: GET

    Description: Returns a JSON response indicating the health status of the application.

    Response:

    json

    {
      "error": false,
      "message": "check health of the app"
    }

Get Person Info

    URL: /person/:person_id/info
    Method: GET
    Description: Retrieves information for a specific person by person_id.
    Parameters:
        person_id (path parameter): ID of the person to retrieve information for.

Create Person

    URL: /person/create
    Method: POST
    Description: Creates a new person record.
    Request Body: JSON object with the details of the person to create.