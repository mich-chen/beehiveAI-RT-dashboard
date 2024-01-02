import React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { AggregatedSentiments } from '../App';
import TableContainer from '@mui/material/TableContainer';
import Widget from './Widget';

// FUTURE FOLLOWUP: 
// -- Can improve the aggregate table to be a data grid from @mui/x-data-grid
// -- dataGrid allows for additional functionalities such as filtering, sorting, etc

interface AggregatedData {
  [key: string]: AggregatedSentiments
}

const AggregateTable: React.FC<({ data: AggregatedData })> = ({ data }) => {
  const header = "Airline Aggregate"
  const content = data ? (
  <TableContainer sx={{ maxHeight: 250, maxWidth: 450 }} >
    <Table size="small" stickyHeader data-testid="airline-aggregate-table">
    <TableHead>
      <TableRow>
      <TableCell>Airline</TableCell>
      <TableCell>Positive</TableCell>
      <TableCell>Negative</TableCell>
      <TableCell>Neutral</TableCell>
      </TableRow>
    </TableHead>

    <TableBody>
      {Object.entries(data).map(([airline, val]) => (
        <TableRow key={airline}>
          <TableCell>{airline}</TableCell>
          <TableCell>{val.positive}</TableCell>
          <TableCell>{val.negative}</TableCell>
          <TableCell>{val.neutral}</TableCell>
        </TableRow>
      ))}
    </TableBody>
  </Table>
</TableContainer>
  ) : null

  return (
    <div className="metric" >
      <Widget header={header} content={content} />
    </div>
  )
}

export default AggregateTable;