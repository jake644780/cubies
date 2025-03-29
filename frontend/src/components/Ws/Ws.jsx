import styles from "./style.module.scss";
import React, { useState, useEffect } from "react";

export default function WebSocketChat(){
  const [socket, setSocket] = useState(null);
  const [username, setUsername] = useState("");
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);
  const [usernameSet, setUsernameSet] = useState(false);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("Connected to WebSocket server");
    };

    ws.onmessage = (event) => {
      console.log("Message from server:", event.data);
      setMessages((prevMessages) => [...prevMessages, event.data]);

      if (event.data.includes("Username set to")) {
        setUsernameSet(true);
      }
    };

    ws.onclose = () => {
        ws.send({
            type:""
        })
      console.log("Disconnected from WebSocket server");
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  const handleSetUsername = () => {
    if (!username.trim()) {
      alert("Please enter a username.");
      return;
    }

    const message = JSON.stringify({ setUsername: username });
    socket.send(message);
  };

  const handleSendMessage = () => {
    if (!message.trim()) return;

    socket.send(message);
    setMessages((prevMessages) => [...prevMessages, `You: ${message}`]);
    setMessage("");
  };

  return (
    <div>
      <h2>WebSocket Chat</h2>

      {!usernameSet ? (
        <div>
          <input
            type="text"
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <button onClick={handleSetUsername}>Set Username</button>
        </div>
      ) : (
        <div>
          <input
            type="text"
            placeholder="Type a message..."
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />
          <button onClick={handleSendMessage}>Send</button>
          <ul>
            {messages.map((msg, index) => (
              <li key={index}>{msg}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

