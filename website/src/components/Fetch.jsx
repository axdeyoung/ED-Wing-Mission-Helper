import { useState, useEffect } from 'react';

const Fetch = () => {

    const [fetchedText, setFetchedText] = useState("[No server response]");

    useEffect(() => {
        fetch('http://127.0.0.1:31173/')
            .then((response) => {
                return response.text();
            })
            .then((text) => {
                console.log(text);
                setFetchedText(text);
            });
    }, []);

    return (
        <div>
            {fetchedText}
        </div>
    );
};

export default Fetch;