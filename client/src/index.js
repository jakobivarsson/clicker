import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import Counter from './components/counter/Counter';
import About from './components/about/About';
import { Router, Route, browserHistory } from 'react-router';
import './index.css';

ReactDOM.render(
  <Router history={browserHistory}>
    <Route path="/" component={App} />
    <Route path="/counter" component={Counter} />
    <Route path="/about" component={About} />
  </Router>,
  document.getElementById('root')
);
