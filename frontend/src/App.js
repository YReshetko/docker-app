import 'bootstrap/dist/css/bootstrap.min.css';
import React, { useState, useEffect } from 'react'
import ReactDOM from 'react-dom'
import MyCard from "./components/MyCard";
import {Stack} from "react-bootstrap";

function App() {
    let tempCards = [];
    const [cards, setCards] = useState([]);
    useEffect(() => {
        fetch("/containers", {
            "method": "GET"
        })
            .then(response => response.json())
            .then(response => {
                for (let i = 0; i < response.length; i++) {
                    tempCards.push(MyCard(response[i]));
                    setCards([...tempCards]);
                }
            })
            .catch(err => {
                console.log(err);
            });
    }, [])
  return (
    <div>
        {cards}
    </div>
  );
}

export default App;
