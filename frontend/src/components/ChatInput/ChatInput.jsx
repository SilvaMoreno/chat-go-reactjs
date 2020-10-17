import React from 'react'
import './ChatInput.scss'

const ChatInput = ({send}) => {

    return (
        <div className="ChatInput">
            <input type="text" onKeyDown={send}/>
        </div>
    )
}

export default ChatInput