# RSS Aggregator

RSS Aggregator is a backend-only tool for aggregating and managing RSS feeds from various sources. It provides functionality for fetching, parsing, and storing RSS feed data.

## Features

- **Subscription Management**: Users can add, edit, and delete RSS feed subscriptions programmatically.
- **Feed Parsing**: Automatically fetches and parses RSS feeds from subscribed sources.
- **Data Storage**: Stores feed data in a database for easy retrieval and management.
- **Scheduled Tasks**: Supports scheduled tasks for updating and refreshing feed data at regular intervals.

## Installation

To install and run the RSS aggregator backend, follow these steps:

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/deepak04112002/RSS_aggregator.git
    ```

2. Navigate to the project directory:

    ```bash
    cd RSS_aggregator
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Build the project:

    ```bash
    go build
    ```

5. Run the executable:

    ```bash
    ./rssagg
    ```

6. The backend server will start running on `http://localhost:8080`.

## Usage

- **API Endpoints**: The backend exposes API endpoints for managing RSS feed subscriptions and fetching feed data.
- **Scheduled Tasks**: Use cron jobs or scheduler libraries to periodically trigger feed update tasks.
- **Database Management**: Configure and manage the database according to your requirements.

## Technologies Used

- **Go**: Backend programming language.
- **SQL**: Database management.
- **HTTP Server**: Built-in HTTP server for handling API requests.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Hat tip to anyone whose code was used
- Inspiration
- etc.
