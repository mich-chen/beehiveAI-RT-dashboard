import React, { useState, useEffect, useRef } from 'react';
import './App.css';
import Tweet from './components/Tweet';
import AggregateTable from './components/AggregateTable';
import TotalResponses from './components/TotalResponses';
import DateDistribution from './components/DateDistribution';

interface Message {
  tweetId: number;
  airlineSentiment: string;
  airlineSentimentConfidence: number;
  negativereason: string;
  negativereasonConfidence: number;
  airline: string;
  airlineSentimentGold: string;
  name: string;
  negativereasonGold: string;
  retweetCount: number
  text: string;
  tweetCoord: string;
  tweetCreated: string;
  tweetLocation: string;
  userTimezone: string;
}

export interface AggregatedSentiments {
  total: number;
  positive: number;
  negative: number;
  neutral: number;
}

interface Metrics {
  aggregatedSentiments: { [key: string]: AggregatedSentiments };
  dateDistributions: { [key: string]: number };
}

interface WebsocketMessage {
  messages: Message[];
  metrics: Metrics;
}

function App() {
  const connection = useRef<WebSocket | null>(null);

  const [tweets, setTweets] = useState<Message[]>([]);
  const [aggregateSentiments, setAggregateSentiments] = useState<{ [key: string]: AggregatedSentiments }>({});
  const [dateDistribution, setDateDistribution] = useState<{ [key: string]: number }>({});

  const handleData = (data: WebsocketMessage) => {
    // to optimize, implement a caching 
    setTweets(data.messages);

    // Server sends updated object therefore do not need prevState
    setAggregateSentiments(data.metrics.aggregatedSentiments);
    setDateDistribution(data.metrics.dateDistributions);
  }

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/websocket");

    console.log('Connecting...')
  
    socket.onopen = () => {
      console.log('Successfully connected to websocket!');
    }

    // receives message
    socket.onmessage = async (event) => {
      console.log("received message from server: ");
      const data = await JSON.parse(event.data);
      handleData(data);
    }

    socket.onclose = (event) => {
      console.log('Closed websocket connection!', event);
    }
  
    socket.onerror = (err) => {
      console.log('Websocket err:', err)
    }

    connection.current = socket;

    // cleanup
    return () => connection.current?.close();
  },[])

  console.log("aggregated sentiments", aggregateSentiments);
  console.log("date distribution", dateDistribution);

  return (
    <div className="App">
      <div className="tweets-container">
        <ul>
          {tweets.map((tweet) => {
            return (
              <Tweet name={tweet.name} text={tweet.text} />
            )
          })}
        </ul>
      </div>
      <div className="metrics-container">
        <h2>Tweet Message Dashboard Metrics</h2>
        <div className="metrics-widgets">
          <TotalResponses total={tweets.length} />
          
          <AggregateTable data={aggregateSentiments} />
          
          <DateDistribution data={dateDistribution} />
        </div>
        
      </div>
    </div>
  );
}

export default App;