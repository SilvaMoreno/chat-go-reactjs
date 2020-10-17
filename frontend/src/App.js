import React, { useEffect, useState } from "react";
import logo from "./logo.svg";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from "./components/Header";
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput";

function App() {
    const [message, setMessage] = useState(null);
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        if (message) {
            console.log("NEW message", message);
            setMessages([...messages, message]);
            console.log("LIST message", messages);
        }
    }, [message]);

    useEffect(() => {
        connect((msg) => {
            setMessage(msg);
        });
    });

    const send = (event) => {
        if (event.keyCode === 13) {
            sendMsg(event.target.value);
            event.target.value = "";
        }
    };
    return (
        <div className="App">
            <Header />
            <ChatHistory chatHistory={messages} />
            <ChatInput send={send} />
        </div>
    );
}

export default App;
