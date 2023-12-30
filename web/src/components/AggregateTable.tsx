import React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { AggregatedSentiments } from '../App';

interface AggregatedData {
  [key: string]: AggregatedSentiments
}

// FUTURE FOLLOWUP: 
// -- Can improve the aggregate table to be a data grid from @mui/x-data-grid
// -- dataGrid allows for additional functionalities such as filtering, sorting, etc

const AggregateTable: React.FC<({ data: AggregatedData })> = ({ data }) => {
  return (
    <div className="metric">
      <h2>Airline Aggregate</h2>
      {data ? <Table size="small" className="sentiments-table">
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
      : null}

    </div>
  )
}

export default AggregateTable;