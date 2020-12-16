import * as React from 'react';

import { Layout } from 'antd';

import './style.scss';

const { Header } = Layout;

export default function Head() {
  return (
    <Header className='header'>
      <div className='container'>
        <h1>Lending Swap</h1>
      </div>
    </Header>
  );
}
