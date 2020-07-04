"use strict";

import React from "react";

import ReactDOM from "react-dom";

const styles = document.createElement("style");
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

const App = React.createElement('div',{ id: 'app-container' }, '7g-tooling baybay');
const root = document.getElementById('app');

ReactDOM.render(App, root);