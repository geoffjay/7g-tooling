import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Network from './pages/network';

import './app.scss';


const App: React.FC = () : JSX.Element=> {
    return (
        <Router>
            <Switch>
                <Route path="/" component={Network} />
            </Switch>
        </Router>
    );
};

// throws overload error if not component not called
export default <App />;
