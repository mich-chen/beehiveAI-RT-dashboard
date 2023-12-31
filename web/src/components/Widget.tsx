import React, { ReactNode } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardHeader from '@mui/material/CardHeader';

const Widget: React.FC<{ header: string; content: ReactNode; sx?: {} | undefined }> = ({ header, content, sx }) => {
  const headerSx = {
    backgroundColor: "#90caf9",
    height: 20,
  }

  const constentSx = {
    padding: 0,
    ...sx,
  }
  return (
    <Card>
      <CardHeader title={header} sx={headerSx} />
      <CardContent sx={constentSx}>
        {content}
      </CardContent>
    </Card>
  );
}

export default Widget;