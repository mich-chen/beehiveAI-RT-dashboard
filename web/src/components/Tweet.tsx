import React from 'react';

const Tweet: React.FC<{ name: string; text: string }> = ({ name, text }) => {
  return (
    <div className="tweet">
      <h4>{name}</h4>
      <p>{text}</p>
    </div>
  )
}

export default Tweet;