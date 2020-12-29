import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import './App.css';

import { Layout } from 'antd';

import Head from 'components/Head';
import Footer from 'components/Footer';
import SwapPage from 'pages/SwapPage';

function App() {
  return (
    <Router>
      <Layout className='App'>
        <Head />
        <Switch>
          <Route exact path='/' component={SwapPage} />
        </Switch>
        <Footer />
      </Layout>
    </Router>
  );
}

export default App;
