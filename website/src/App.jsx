import React, { useState } from 'react';
import "./App.css";
import Fetch from './components/Fetch'

function App() {
    const [count, setCount] = useState(0);
    const localServerUrl = "http://127.0.0.1:31173/"
    
    var responseText = "[none]";



    return (
        <div className="App">
            <h1>Simple React App</h1>
            <p>Click the button to increase the count:</p>
            <button onClick={() => setCount(count + 1)}>
                Count: {count}
            </button>
            <p>Server status: <Fetch /></p>
        </div>
    );
}

export default App;