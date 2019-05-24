# Github Analyser

Github Analyser was created as part of my study at the CODE University of Applied Sciences Berlin (@codeuniversity).
It is a REST-API, written in Golang that faciliates the GitHub GraphQL Api to get extensive stats about
GitHub users. Additionally it applies a statistical score metrics on the data to rank the users by for example
activity and popularity. This makes this data more compareable and understandable.

## Project Structure

Package | Description
------------ | -------------
main | Contains all configuration files of the project and the http-server initialization.
app | Is responsible for defining the http-server's own REST-API endpoints and handling their requests. It does not contain any endpoint logic but uses wrapper functions to handle the responses from the core functions of the project.
controller | Is responsible for the main functionality of the project and is depending on most other packages. Functions include the fetching and caching of user data, score calculation and updating.
db | The db package established and manages all database clients (mongodb, redis, elasticsearch) and there utility functions to work with the respective database.
errorutil | Is a helper package for creating custom errors that contain more information additional to the stack trace.
graphql | Is responsible for all requests to the GitHub GraphQL API. It uses a custom client that is inspired by [Machinebox GraphQL Client for Go](http://github.com/machinebox/graphql) and saves the queries as .gql-files within a sup-directory
httputil | Is a helper package for http-responses. It provides simple functions to create and write success and error responses.
logger | Provides additional functions for logging. It uses ASCI-colors and mails to provide overview about all events and issues on the server.
mailer | Is responsible for sending mails using SMTP. Currently mainly used by `logger`.
metrix | The metrix package contains the logic for fetching GitHub user data and calculating the score metrix of this populated data. It only runs with the `FLAG_DO_METRIX` flag being set to 1.
setup | The setup package contains the logic for scraping GitHub logins/usernames and saving them as a .txt-List in the input folder of the `metrix` package. Is only executed if the `FLAG_DO_SETUP`is set to 1.
util | Contains utility functions for all packages. All functions only have a single reponsibility and can be tested by unit tests.

## Deployment

The project uses Docker and Docker Compose to manage all it's services. This also enables a very easy setup.
Please make sure that you have installed the latest version of Docker and Docker Compose.

1. Clone this repository into a folder of your choice:
```shell
> git clone https://github.com/m-lukas/github_analyser.git
```
2. Go into the created folder:
```shell
> cd github-analyser/
```
3. Rename the file ".env.example" to ".env" and fill out the configuration. A reference to the configuration can be found in the following section: "Configuration". For further help, please contact @m-lukas.
4. Build the project (this will create a new Docker image for the deployment and cares about all dependencies):
```shell
> docker-compose build
```
5. Start the databases. App has to be started afterwards because it's depending on the databases being ready for connections:
```shell
> docker-compose up mongo redis elasticsearch
```
6. After all databases have been started properly, start the app:
```shell
> docker-compose up deployment
```
7. The app should be running 🎉

If you want to change the flags within the .env file to start the app in another state, you have to rebuild and start the app again. This is similar to the workflow of point 4 - 6.
For deployments on servers I recommend Google Cloud Computing Engine with the Containerized-OS: https://cloud.google.com/community/tutorials/docker-compose-on-container-optimized-os

## Configuration

Property | Description
------------ | -------------
