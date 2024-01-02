import React from 'react';
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react';
import App from './App';
import WS from 'jest-websocket-mock';

const mockMessage = {
  "messages": [
      {
          "tweetId": 5323041546422352922,
          "airlineSentiment": "positive",
          "airlineSentimentConfidence": 1,
          "airline": "Delta",
          "name": "test",
          "text": "@testing weather conditions",
          "tweetCord": "[40.74804263, -73.99295302]",
          "tweetCreated": "2024-01-01 11:35:52 -0800",
          "tweetLocation": "California",
          "userTimezone": "Pacific Time (US \u0026 Canada)"
      }
  ],
  "metrics": {
      "aggregatedSentiments": {
          "Delta": {
              "positive": 1,
              "negative": 0,
              "neutral": 0
          }
      },
      "dateDistributions": {
          "2024-01-01": 1
      }
  }
}

describe("Test dashboard", () => {
  let server: WS;
  let client: WebSocket;
  beforeEach(async () => {
    server = new WS('ws://localhost:8080/websocket', { jsonProtocol: true });
    client = new WebSocket('ws://localhost:8080/websocket');
    await server.connected;
  })

  afterEach(() => {
    server.close();
    WS.clean();
  })

  test('renders initial content', async () => {
    render(<App />)

    const tweetsContainer = screen.getByTestId('tweets-container');
    const metricsContainer = screen.getByTestId('metrics-container');

    expect(tweetsContainer).toBeInTheDocument();
    expect(metricsContainer).toBeInTheDocument();

    // Check initial empty state
    expect(screen.getByText('Tweet Message Dashboard Metrics')).toBeInTheDocument();
    expect(screen.getByText('Tweets')).toBeInTheDocument();
    expect(screen.getByText('Tweets Count')).toBeInTheDocument();
    expect(screen.getByText('Airlines Count')).toBeInTheDocument();
    expect(screen.getByText('Airline Aggregate')).toBeInTheDocument();
    expect(screen.queryAllByText('0').length).toBe(2); // both count widget values are 0
  })

  test('client updates in real-time from received ws message', async () => {
    render(<App />)

    let message = null;
    client.onmessage = (e) => {
      message = e.data;
    }
    server.send(mockMessage)
    expect(message).toEqual(JSON.stringify(mockMessage));

    await waitFor(() => {
      expect(screen.getByText('Tweet Message Dashboard Metrics')).toBeInTheDocument();
      expect(screen.getByText('@test')).toBeInTheDocument();
      expect(screen.getByText('Delta')).toBeInTheDocument();
      expect(screen.getByText('@testing weather conditions')).toBeInTheDocument();
      expect(screen.getByText('Tweets Count')).toBeInTheDocument();
      expect(screen.getByText('Airlines Count')).toBeInTheDocument();
      expect(screen.getByText('Airline Aggregate')).toBeInTheDocument();

      const counts = screen.queryAllByText('1');
      // 3 counts include: tweets, airlines, and airline aggregate
      expect(counts.length).toBe(3);
    });
  });
}
)
