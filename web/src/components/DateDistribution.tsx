import React from 'react';
import { BarChart } from '@mui/x-charts/BarChart';

interface DistributionData {
  [key: string]: number 
}

const DateDistribution: React.FC<{data: DistributionData}> = ({ data }) => {
  
  const sortedData = Object.keys(data).length ? Object.fromEntries(Object.entries(data).sort()) : null;

  return (
    <div className="metric">
      <h2>Date Distribution</h2>
      {sortedData != null ? <BarChart
        xAxis={[{ scaleType: 'band', data: Object.keys(sortedData) }]}
        series={[{ data: Object.values(sortedData) }]}
        width={500}
        height={300}
      /> : null}
      
    </div>
  );
}

export default DateDistribution;