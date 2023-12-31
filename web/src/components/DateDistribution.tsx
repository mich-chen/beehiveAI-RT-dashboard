import React from 'react';
import { BarChart } from '@mui/x-charts/BarChart';
import Widget from './Widget';

interface DistributionData {
  [key: string]: number 
}

const DateDistribution: React.FC<{data: DistributionData}> = ({ data }) => {
  const header = "Date Distribution"

  const sortedData = Object.keys(data).length ? Object.fromEntries(Object.entries(data).sort()) : null;

  const content = sortedData != null ? (<BarChart
  xAxis={[{ scaleType: 'band', data: Object.keys(sortedData) }]}
  series={[{ data: Object.values(sortedData) }]}
  height={250}
  width={475}
/>) : null

  return (
    <div className="metric">
      <Widget header={header} content={content} />
    </div>
  );
}

export default DateDistribution;