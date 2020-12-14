import logo from './logo.svg';
import './App.css';
import { useEffect } from 'react';
// import { connectOneWallet } from './utils/connectOneWallet';
// import { connectMetamask } from './utils/connectMetamask';

function App() {
  useEffect(() => {
    // connectOneWallet();
    // connectMetamask();
  });

  return (
    <div className='App'>
      <header className='App-header'>
        <img src={logo} className='App-logo' alt='logo' />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className='App-link'
          href='https://reactjs.org'
          target='_blank'
          rel='noopener noreferrer'
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
