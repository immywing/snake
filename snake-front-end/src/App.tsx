import React from 'react';
import logo from './logo.svg';
import './App.css';
import Button from '@mui/material/Button';
import { useState } from 'react';
import { KeyboardEvent } from 'react';

function newConnection() {
  console.log("fetching websocket conn");
  return new WebSocket("http://localhost:8080/snake");
}

interface snake {
  Grid: number[][];
  Score: number
}

const emptyGrid = {
  Grid: [
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
  ],
  Score: 0
};

const inputSocket = new WebSocket("http://localhost:8080/input")
function handleKeyDown(event: globalThis.KeyboardEvent) {
  console.log(inputSocket.readyState)
  switch (event.key) {
    case 'w':
      inputSocket.send('0')
      break;
    case 'a':
      inputSocket.send('2')
      break;
    case 's':
      inputSocket.send('1')
      break;
    case 'd':
      inputSocket.send('3')
      break;
  }
};
document.addEventListener('keydown', handleKeyDown)

function App() {
  const [grid, setGrid] = useState<snake>(emptyGrid);
  const [socket, setSocket] = useState<WebSocket | null>();
  const onClick = () => {
    const newSocket = newConnection();
    setSocket(newSocket); 
    if (newSocket) {
      newSocket.onmessage = (event: MessageEvent<any>) => {
        const data: snake = JSON.parse(event.data);
        setGrid(data);
      };
    } else {
      console.error("Failed to create WebSocket connection");
    }
  };
  console.log(grid)
  return (
    <div className="App">
      <div className="button-container">
        <Button variant="contained" onClick={onClick}>
          New Game
        </Button>
      </div>
      <br/>
      <div className='Score'>Score: {grid.Score}</div>
      <div className="grid-container">
      {grid.Grid.map((row, rowIndex) => (
        <div className='grid-row' key={rowIndex}>
          {row.map((item, colIndex) => (
            <div key={rowIndex.toString() + colIndex.toString()}>
              {item === -1 ?
              (<div className="Grid-GameOver"></div>) : 
              item === 0 ?
              (<div className="Grid"></div>) : 
              item === 2 ? (<div className="Grid">*</div>) :
              (<div className="Snake"></div>)}
            </div>
          ))}
        </div>
      ))}
      </div>
    </div>
  );
}

export default App;
