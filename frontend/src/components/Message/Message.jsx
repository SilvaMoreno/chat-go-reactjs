import React from 'react'
import './Message.scss'

const Message = ({message}) => {
    const tmp = JSON.parse(message)

    return <div className="Message">{tmp.body}</div>
}

export default Message