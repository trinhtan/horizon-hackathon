import React, { useState } from 'react';

import './style.scss';

import { Row, Col, Input, Select } from 'antd';
import { RightOutlined, ArrowRightOutlined, ArrowDownOutlined } from '@ant-design/icons';

const { Option } = Select;

function SwapPage() {
  const [roadSwap, setRoadSwap] = useState(0);

  function onChange(value) {
    console.log(`selected ${value}`);
  }

  function onBlur() {
    console.log('blur');
  }

  function onFocus() {
    console.log('focus');
  }

  function onSearch(val) {
    console.log('search:', val);
  }

  return (
    <div className='swap-page'>
      <Row className='container' justify='space-between'>
        <Col xs={{ order: 2, span: 24 }} md={{ order: 1, span: 10 }}>
          <div className='content-swap'>
            <div className='input-and-select-token'>
              <div className='token-source-dest'>
                <div className='label-input-token'>
                  <div class='sc-hSdWYo euiRCS css-1rhdhic'>From</div>
                </div>
                <div className='box-input-token'>
                  <Input size='large' placeholder='0.0' className='input-token' />
                  <Select
                    showSearch
                    placeholder='Select a token'
                    optionFilterProp='children'
                    onChange={onChange}
                    onFocus={onFocus}
                    onBlur={onBlur}
                    onSearch={onSearch}
                    filterOption={(input, option) =>
                      option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
                    }
                    className='button-select-token'
                  >
                    <Option value='jack'>Jack</Option>
                    <Option value='lucy'>Lucy</Option>
                    <Option value='tom'>Tom</Option>
                  </Select>
                </div>
              </div>
              <div className='symbol-arrow-down'>
                <ArrowDownOutlined />
              </div>
              <div className='token-source-dest'>
                <div className='label-input-token'>
                  <div class='sc-hSdWYo euiRCS css-1rhdhic'>To</div>
                </div>
                <div className='box-input-token'>
                  <Input size='large' placeholder='0.0' className='input-token' />
                  <Select
                    showSearch
                    placeholder='Select a token'
                    optionFilterProp='children'
                    onChange={onChange}
                    onFocus={onFocus}
                    onBlur={onBlur}
                    onSearch={onSearch}
                    filterOption={(input, option) =>
                      option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
                    }
                    className='button-select-token'
                  >
                    <Option value='jack'>Jack</Option>
                    <Option value='lucy'>Lucy</Option>
                    <Option value='tom'>Tom</Option>
                  </Select>
                </div>
              </div>
              <div className='button-swap'>
                <button id='swap-button' class='swap-button'>
                  <div class='css-10ob8xa'>Swap Anyway</div>
                </button>
              </div>
            </div>
          </div>
        </Col>
        <Col xs={{ order: 1, span: 24 }} md={{ order: 2, span: 12 }}>
          <div className='area-chain-to-chain'>
            <Row justify='space-between'>
              <Col
                span={11}
                className={
                  'eth-to-harmony button-chain ' + (roadSwap === 0 ? 'enable-road' : 'disable-road')
                }
                onClick={() => {
                  setRoadSwap(0);
                }}
              >
                <Row>
                  <Col span={10}>
                    <div className='token-source'>
                      <h3>
                        <img alt='icon' src='/eth.svg' className='icon-token' />
                        ETH
                      </h3>
                    </div>
                  </Col>
                  <Col span={4}>
                    <div className='symbol-arrow'>
                      <RightOutlined />
                    </div>
                  </Col>
                  <Col span={10}>
                    <div className='token-dest'>
                      <h3>
                        <img alt='icon' src='/one.svg' className='icon-token' /> ONE
                      </h3>
                    </div>
                  </Col>
                </Row>
              </Col>
              <Col
                span={11}
                className={
                  'harmony-to-eth button-chain ' + (roadSwap === 1 ? 'enable-road' : 'disable-road')
                }
                onClick={() => {
                  setRoadSwap(1);
                }}
              >
                <Row>
                  <Col span={10}>
                    <div className='token-source'>
                      <h3>
                        <img alt='icon' src='/one.svg' className='icon-token' /> ONE
                      </h3>
                    </div>
                  </Col>
                  <Col span={4}>
                    <div className='symbol-arrow'>
                      <RightOutlined />
                    </div>
                  </Col>
                  <Col span={10}>
                    <div className='token-dest'>
                      <h3>
                        <img alt='icon' src='/eth.svg' className='icon-token' />
                        ETH
                      </h3>
                    </div>
                  </Col>
                </Row>
              </Col>
              <Col span={24} className='wallets-area'>
                <Row>
                  <Col span={11}>
                    <div className='wallet-source'>
                      <div className='text-source-dest'>
                        <h2>Wallet Source</h2>
                      </div>
                      <Row justify='center' className='list-wallets'>
                        <Col>
                          <div className='info-wallet'>
                            <img alt='icon' className='icon-wallet' src='/metamask.png' />
                            <h4>Metamask</h4>
                          </div>
                        </Col>
                        <Col>
                          <div className='info-wallet'>
                            <img alt='icon' className='icon-wallet' src='/trust.png' />
                            <h4>Trust</h4>
                          </div>
                        </Col>
                      </Row>
                    </div>
                  </Col>
                  <Col span={2} className='symbol-arrow-wallet'>
                    <ArrowRightOutlined />
                  </Col>
                  <Col span={11}>
                    <div className='wallet-dest'>
                      <div className='text-source-dest'>
                        <h2>Wallet Dest</h2>
                      </div>
                      <Row justify='center' className='list-wallets'>
                        <Col>
                          <div className='info-wallet'>
                            <img alt='icon' className='icon-wallet' src='/trust.png' />
                            <h4>Trust</h4>
                          </div>
                        </Col>
                      </Row>
                    </div>
                  </Col>
                </Row>
              </Col>
            </Row>
          </div>
        </Col>
      </Row>
    </div>
  );
}

export default SwapPage;
