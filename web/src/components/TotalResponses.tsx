import React from 'react';

const TotalResponses: React.FC<{ total: number }> = ({ total }) => {
  const header = "Total Tweets Received"
  return (
    <div className="metric">
      <h4>{header}</h4>
      <div>{total}</div>
    </div>
  )
}

export default TotalResponses;