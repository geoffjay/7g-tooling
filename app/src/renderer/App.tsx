import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Network from './pages/network';

import './app.scss';


const App: React.FC = () : JSX.Element=> {
    return (
        <div className="main-content">
            <Router>
                <Switch>
                    <Route path="/" component={Network} />
                </Switch>
            </Router>
        </div>
    );
};

// TODO: idk
// throws overload error if not component not called
export default <App />;
