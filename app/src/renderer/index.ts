'use strict';

import ReactDOM from 'react-dom';

import App from './App';


const styles = document.createElement('style');
styles.innerText = `
    body {
        margin: 0;
    }

    #app {
        height: 100vh;
        width: 100vw;
    }

    #app-container {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }
`;

document.head.appendChild(styles);

const root = document.getElementById('app');

ReactDOM.render(App, root);
