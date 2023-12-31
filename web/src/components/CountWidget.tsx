import React from 'react';
import Typography from '@mui/material/Typography';
import Widget from './Widget';

const CountWidget: React.FC<{ header: string; count: number }> = ({ header, count }) => {
  const content = (
    <Typography variant="h3" data-testid="count">
      {count}
    </Typography>
  )

  return (
    <div className="metric">
      <Widget header={header} content={content} sx={{ textAlign: "center" }} />
    </div>
  )
}

export default CountWidget;