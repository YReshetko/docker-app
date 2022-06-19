import 'bootstrap/dist/css/bootstrap.min.css';
import React, {useMemo, useState} from 'react'
import Services from "./components/Services";

function App() {
    console.log("Starting app")
    const [services, setServices] = useState([]);
    let fetchState = 0;

    useMemo(() => {
        console.log("Fetching services")
        fetch("/api/v1/services", {
            "method": "GET"
        })
            .then(response => response.json())
            .then(response => {
                setServices(Services(response));
            })
            .catch(err => {
                console.log(err);
            });
    }, [fetchState]);
    fetchState = 1;

    return (
        <div>
            {services}
        </div>
    );
}

export default App;
