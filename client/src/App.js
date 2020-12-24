import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import './App.css';
// import { useEffect } from 'react';
// import { connectOneWallet } from './utils/connectOneWallet';
// import { connectMetamask } from './utils/connectMetamask';

import { Layout } from 'antd';

import Head from 'components/Head';
import SwapPage from 'pages/SwapPage';

function App() {
  // useEffect(() => {
  //   // connectOneWallet();
  //   connectMetamask();
  // });

  return (
    <Router>
      <Layout className='App'>
        <Head />
        <Switch>
          <Route exact path='/' component={SwapPage} />
        </Switch>
        {/* <FooterPage /> */}
      </Layout>
    </Router>
  );
}

export default App;
