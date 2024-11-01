# LostNFound
LostNFound is a web service designed to help people find lost items. Users can report found items, allowing the original owners to retrieve information about their lost belongings.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Screenshots](#screenshots)

## Features

- Add  lost items information and photo
- Simple and user-friendly interface for interaction

## Technologies

This project is built using the following technologies:

- **Backend**: Go (Golang) with the following libraries:
  - `gorilla/mux` for routing
  - `viper` for configuration management
  - `zap` for logging
  - `gorm` for ORM
- **Database**: MySQL
- **Object Storage**: MinIO
- **Containerization**: Docker Compose
<!-- - **Build Tool**: Makefile -->

## Getting Started

To get a local copy up and running, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/subliker/lost-n-found-service
   cd lost-n-found-service
   ```

2. **Set up your environment**:

   Create a `.env` file in the root directory and define the necessary environment variables. Refer to `.env.example` for a sample configuration.

   Additionally, you can configure the server using `config.toml` located in the `server/` directory.  A sample configuration is provided in `config_example.toml`.

3. **Build and run the application using Docker Compose**:

    ```bash
    docker-compose up --build 
    ```

<!-- 4. **Alternatively, use the Makefile**:

   The Makefile provides convenient commands for managing the project. You can run:

    ```bash
    make up      # to build and start the application 
    make down    # to stop and remove the containers  
    make logs    # to view logs  
    ``` -->

## Usage

Once the application is running, you can access the web interface at `http://localhost:80`.

1. **Adding a Lost Item**: Add a lost item by clicking the corresponding button on the main page.
2. **Searching for a Lost Item**: In the presented list of lost items, try to find yours.

## Logging

Server logs are stored in the `server/logs` directory. You can monitor this directory to keep track of application activities and debug issues. 

### Example Logs

Here are some examples of log entries you might see:
```log
{"level":"info","ts":"2024-11-01T06:38:55.107+0800","msg":"Item store connected"}
{"level":"info","ts":"2024-11-01T06:38:55.149+0800","msg":"Item store migrated"} 
{"level":"info","ts":"2024-11-01T06:38:55.159+0800","msg":"Server routes was initialized"} 
{"level":"info","ts":"2024-11-01T06:38:55.159+0800","msg":"Server instance created"} 
{"level":"info","ts":"2024-11-01T06:38:55.159+0800","msg":"App instance created"} 
{"level":"info","ts":"2024-11-01T06:38:55.159+0800","msg":"App running..."}
```

<!-- ## Screenshots

![Screenshot 1: Home Page](link-to-screenshot-1)  
*Description: Show the main interface where users can navigate the service.*

![Screenshot 2: Report Found Item](link-to-screenshot-2)  
*Description: Display the form for reporting a found item.*

![Screenshot 3: Search Results](link-to-screenshot-3)  
*Description: Illustrate the search results for lost items.* -->