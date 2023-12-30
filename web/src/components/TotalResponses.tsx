import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardHeader from '@mui/material/CardHeader';
import Typography from '@mui/material/Typography';

const TotalResponses: React.FC<{ total: number }> = ({ total }) => {
  const header = "Total Tweets"
  return (
    <div className="metric">
      <Card variant="outlined" sx={{ backgroundColor: "#90caf9" }}>
        <CardHeader title={header} />
        <CardContent className="widget">
          <Typography variant="h3">
            {total}
          </Typography>
        </CardContent>
      </Card>
    </div>
  )
}

export default TotalResponses;