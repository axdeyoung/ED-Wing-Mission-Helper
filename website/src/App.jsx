import React, { useRef } from 'react';
import "./App.css";
import Fetch from './components/Fetch'

function App() {
    const localServerUrl = "http://127.0.0.1:31173/"
    const fetchRef = useRef()

    
    return (
        <div className="App">
            <h1>Simple React App</h1>
            <p>Click the button to update the server status:</p>
            <button onClick={() => fetchRef.current.refresh()}>
                Refresh
            </button>
            <p>Server status: <Fetch ref={fetchRef} url={localServerUrl} defaultText="Loading..." /></p>
        </div>
    );
}

export default App;