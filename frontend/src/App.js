import 'bootstrap/dist/css/bootstrap.min.css';
import React, {useEffect, useState} from 'react'
import MyCard from "./components/MyCard";

function App() {
    let tempCards = [];
    const [cards, setCards] = useState([]);
    useEffect(() => {
        fetch("/api/v1/containers", {
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
