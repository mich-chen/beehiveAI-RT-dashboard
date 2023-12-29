#Instructions

1. Server-Side Component (Go): 
● Create a WebSocket server using the Go programming language. The server should handle incoming webhook data, with the data source being Twitter US Airline Sentiment. We will post each row from the dataset to the webhook endpoint. To download the dataset, please follow this link 
(https://www.kaggle.com/datasets/crowdflower/twitter-airline-sentiment). 
● Implement a simple storage mechanism to store and manage the received responses. ● Ensure that the server is capable of handling multiple WebSocket connections concurrently. 
● Implement appropriate error handling for robustness. 
2. Webhook Handling: 
● Design an endpoint to receive webhook responses. The server should parse incoming data and broadcast it to all connected WebSocket clients in real-time.
3. Client-Side Component (React - TypeScript): 
● Develop a React-based dashboard that connects to the WebSocket server. ● Display incoming tweet text in real-time on the dashboard. 
● Include basic metrics on the dashboard, such as 
○ The total responses received 
○ Aggregate sentiment (airline_sentiment): 
■ Negative 
■ neutral 
■ Positive 
○ The date distribution of tweet_created 
● Ensure the user interface is user-friendly and free of UX issues. 
4. WebSocket Communication: 
● Use WebSockets for real-time communication between the server and the client. ● Implement appropriate error handling and reconnection mechanisms in the client to handle potential disruptions in the WebSocket connection. 
5. Code Structure and Style: 
● Follow best practices for code organization and structure in both the Go server and React client. 
● Write clean and well-documented code. 
● Use consistent coding styles for readability. 
3. User Experience: 
● When a webhook receives data, it should be immediately displayed in the dashboard in real-time. 
● All metrics should be updated in real-time. 
● Provide a clean and user-friendly interface. 
● In case of errors, the user should receive the message. The error should also be logged on the server side. 
7. Bonus Features (Optional): 
● Implement a feature to filter, categorize, or search the data on the dashboard. ● Add basic authentication for the dashboard. 
● Include visualizations for response metrics. 
Submission Guidelines: 
1. Repository Structure:
● Create a clear and organized repository structure. 
● Include a README file with instructions on how to run the application. 2. Code Quality: 
● Write clean, modular, and well-documented code. 
● Use meaningful variable and function names. 
● Implement error handling where necessary. 
3. Testing: 
● Include unit tests for critical components. 
● Ensure that the application works correctly in different scenarios. 4. Deployment: 
● Provide instructions for deploying the application locally. 
5. Submission: 
● Submit a link to the GitHub repository containing the code