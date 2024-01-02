# Real-Time Webhook Dashboard

This is a full-stack webhook and websocket dashboard to that receives qualitative data via webhook in real-time, with a Go server with `/webhook` and `/websocket` endpoints and a React client. The data source is [Twitter US Airline Sentiment](https://www.kaggle.com/datasets/crowdflower/twitter-airline-sentiment). The React dashboard displays new received Tweet messages in real-time and updates basic dashboard metrics in real-time.

Current server handles creating new websocket connection, adding it to a connections map to store for broadcasting and a webhook endpoint to receive JSON data, parse and aggregate, and then broadcast to websocket clients.

## Contents
* [Tech Stack](#tech-stack)
* [Installation and Deployment](#installation-and-deployment)
* [Testing](#testing)
* [Follow Up](#fast-follows-and-tradoffs)
* [Author](#author)

## Tech Stack

- Server
  - Go backend server
  - `net/http` for HTTP server
  - `gorilla/websocket` for websocket connections
  - `go-playground/validator` for JSON validation
  - `sync` for mutex
  - `net/httptest` for testing

- Client
  - React and Typescript
  - `@mui` materials UI for diagrams, charts, and layout components
  - React Testing Library for testing
  - `jest-websocket-mock` for websocket mock and testing

## Installation and Deployment

### Server

1. git clone the repo either by SSH or https
2. `cd beehiveAI-RT-dashboard` to be in root directory run the following commands to start the server on locally port `:8080`

```
$ go install
$ go build -race      // enable race detector
$ ./beehiveAI-RT-dashboard
```

3. Alternatively, you can run the following commands and not the executable

```
$ go install
$ go run -race .      // enable race detector
```

4. If using Postman, you can follow the Postman [docs](https://learning.postman.com/docs/sending-requests/websocket/create-a-websocket-request/) to create new websocket request. Set the websocket endpoint to `ws://localhost:8080/websocket`

5. Add new HTTP request tab and set the endpoint to `http://localhost:8080/webhook` with `POST` method and begin sending each row as `JSON` data from the above data source.
  - Expected behavior is sending single row from data source to simulate a feed of tweets
```
// example POST request body
{
    "tweetId": 5323041546422352922,
    "airlineSentiment": "positive",
    "airlineSentimentConfidence": 1,
    "negativereason": null,
    "negativereasonConfidence": null,
    "airline": "Delta",
    "airlineSentimentGold": null,
    "name": "Test",
    "negativereasonGold": null,
    "retweetCount": 0,
    "text": "@delta expecting some delays due to weather conditions",
    "tweetCord": "[40.74804263, -73.99295302]",
    "tweetCreated": "2024-01-01 11:35:52 -0800",
    "tweetLocation": "California",
    "userTimezone": "Pacific Time (US & Canada)"
}
```

6. Switch back to websocket tab in Postman and validate message received

### Web

1. Open new terminal window and change directory to `/web` and run following commands to start React client

```
$ cd web/
$ yarn install
$ yarn start
```

2. Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

## Testing

Run the following commands to run tests on server and client

#### Server Side Tests

1. Change to repo root directory

```
$ go clean -testcache   // clean up cache if any
$ go test -race ./...    // runs main_test.go and subdirectory tests with race detector flag
```

- **Please note:** The comment in `main_test.go` line 60 with follow-up fix and tradeoff in test set up

```
// CAVEAT TO FIX TEST SETUP IN FOLLOUP:
	// - test fails when accsessing var server when running $ go test -race .
	// 		- race detected even though when t.Log(server) you will see connection successfully added to server.conns map
	// - uncomment below and update line 52 with server variable and rerun tests with $ go test . (removing -race flag) and tests will pass showing connection successfully added to server.conns map
```

#### Client Side Tests

1. Have 2 separate terminals open, one terminal running the Server `$ go run .`
2. Second terminal to run Client, change to `/web/src` directory
3. **KNOWN ISSUE:** `@recharts` is using `d3` library that has a [known issue](https://github.com/recharts/recharts/issues/2991) causing tests with **react-testing-library** to fail due to an export issue in their library. Due to scope of this project, our workaround will be to comment out the `<DateDistribution />` in **line 112** in `App.tsx` before running FE tests. Issue not present when testing manually running client locally if you would like to test manually.

- Server terminal
```
$ go run .
```

- Client terminal
```
$ cd web/src/
$ yarn test
```

## Fast Follows and Tradoffs

#### Server Side 

- Update server to use Go channels for handling concurrencies. My tradeoff was using `mutex` instead to lock variables, but in fast follow I'd like to improve my knowledge on Go channels and add that to this server.
- Server websocket endpoint to handle reconnections and adding Ping/Pong handler where client disappears / lost connections according to `gorilla/websocket` 's [Control Messages](https://pkg.go.dev/github.com/gorilla/websocket#hdr-Control_Messages). Unfortunately, I did not have a chance to implement reconnection but would be great opportunity to do so as a fast follow.
- Handling authentication and update broadcasting logic to appropriate clients

#### Client Side

- `<Loading />` state and component until websocket connection returns successful
- Add animation on tweet feed showing newest tweet fading in or some animation/indication it is new (currently shows newest tweet as top message in feed)
- Implement a feature to filter, categorize, or search the data on the dashboard
- Allow drag / drop / reordering of dashboard widgets for customization
- Add basic authentication for dashboard

## Author

**Michelle Chen** \
[LinkedIn](https://www.linkedin.com/in/mich-chen) \
Contact: mich.chen.94@gmail.com \
Please feel free to contact me for any issues with deployment and running the application!