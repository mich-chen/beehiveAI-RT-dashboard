import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardHeader from '@mui/material/CardHeader';
import Typography from '@mui/material/Typography';

const Tweet: React.FC<{ name: string; text: string }> = ({ name, text }) => {
  const userName = '@'.concat(name);
  return (
    <Card variant="outlined" sx={{ margin: 2 }}>
      <CardHeader title={userName} sx={{ backgroundColor: "#42a5f5" }}></CardHeader>
      <CardContent className="tweet">
        <Typography variant="body1">
          {text}
        </Typography>
      </CardContent>
    </Card>
  )
}

export default Tweet;