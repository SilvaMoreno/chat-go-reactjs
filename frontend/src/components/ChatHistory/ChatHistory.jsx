import React from 'react'
import Message from '../Message'
import "./ChatHistory.scss"

const ChatHistory = ({chatHistory = []}) => {

    const messages = chatHistory.map((msg, index) => <Message key={index} message={msg.data} />)

    return (
        <div className="chatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    )
}

export default ChatHistory