# Data Analysis Microservice

## Description

This microservice focuses on providing real-time and historical stock data analysis. Built using Golang and the Echo web framework, it aims to serve as a robust and scalable service for fetching and analyzing stock data.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Routes](#routes)
- [Testing](#testing)
- [Future Enhancements](#future-enhancements)
- [License](#license)

## Installation

To get started, clone the repository to your local machine.

```bash
git clone https://github.com/yourusername/data_analysis_microservice.git
```

Navigate into the project directory and install the required packages.

```bash
cd data_analysis_microservice
go mod download
```

## Usage

Run the service using:

```bash
go run main.go
```

This will start the server at `localhost:8080`.

## Routes

Below are the initial MVP routes for this microservice:

- **Fetch Stock Data by Symbol**
    - **Route**: `GET /stocks/:symbol`
    - **Purpose**: To get the latest stock data for a given symbol.

- **Fetch Stock Data for a Date Range**
    - **Route**: `GET /stocks/:symbol/:startDate/:endDate`
    - **Purpose**: To get historical stock data for a given symbol between specific dates.

- **Analyze Stock Trends**
    - **Route**: `GET /stocks/:symbol/analysis`
    - **Purpose**: To provide basic stock trend analysis for a given symbol.

- **Stock Recommendations**
    - **Route**: `GET /stocks/recommendations`
    - **Purpose**: To provide stock recommendations based on pre-defined logic.

- **Bulk Fetch**
    - **Route**: `POST /stocks/bulk-fetch`
    - **Purpose**: To fetch data for multiple stock symbols in a single request.

- **Health Check**
    - **Route**: `GET /health`
    - **Purpose**: Basic health check to ensure that the microservice is up and running.

- **List Supported Stock Symbols**
    - **Route**: `GET /stocks/supported`
    - **Purpose**: To list all stock symbols that the service can provide data for.

## Testing

Unit tests have been written for the main functionalities. Run the tests using:

```bash
go test ./...
```

## Future Enhancements

- Add user authentication
- Extend data analysis capabilities
- Add caching mechanisms for improved performance

## License

This project is under the MIT License.
