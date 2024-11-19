import { useState, useEffect, forwardRef, useImperativeHandle } from 'react';
import PropTypes from 'prop-types';

const Fetch = forwardRef(({ className, url, defaultText }, ref) => {
    const [fetchedText, setFetchedText] = useState(defaultText);

    const fetchData = async () => {
        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const text = await response.text();
            setFetchedText(text);
        } catch (err) {
            console.error('Fetch error:', err);
            setFetchedText(defaultText);
        }
    };

    useImperativeHandle(ref, () => ({
        refresh: fetchData
    }));

    useEffect(() => {
        fetchData();
    }, [url]);

    return <span className={className}>{fetchedText}</span>;
});

// Add display name for debugging
Fetch.displayName = 'Fetch';

Fetch.propTypes = {
    className: PropTypes.string,
    url: PropTypes.string.isRequired,
    defaultText: PropTypes.string.isRequired,
};

Fetch.defaultProps = {
    className: '',
};

export default Fetch;