import React, { useState, useEffect, useRef } from 'react';
import './App.css';
import Tweet from './components/Tweet';
import AggregateTable from './components/AggregateTable';
import CountWidget from './components/CountWidget';
import DateDistribution from './components/DateDistribution';
import Stack from '@mui/system/Stack';

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
    setTweets(data.messages);

    // Server sends updated object representing server's storage
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

  return (
    <div className="App">
      <div className="tweets-container" data-testid="tweets-container">
        <h2>Tweets</h2>
        <ul>
          {tweets.map((tweet) => {
            return (
              <Tweet name={tweet.name} text={tweet.text} key={tweet.tweetId}/>
            )
          })}
        </ul>
      </div>
      <div className="metrics-container" data-testid="metrics-container">
        <h2>Tweet Message Dashboard Metrics</h2>
        <div className="metrics-widgets">
        <Stack spacing={{ xs: 2, sm: 2, md: 2 }}>
          <Stack spacing={{ xs: 2 }} direction="row">
            <CountWidget header={"Tweets Count"} count={tweets.length} />
            <CountWidget header={"Airlines Count"} count={Object.keys(aggregateSentiments).length} />
          </Stack>

          <Stack spacing={{ xs: 2, sm: 2, md: 2 }} direction="row" useFlexGap flexWrap="wrap">
            <AggregateTable data={aggregateSentiments} />
            <DateDistribution data={dateDistribution} />
          </Stack>
        </Stack>
        </div>
      </div>
    </div>
  );
}

export default App;
