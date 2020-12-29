import React, { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { Row, Col, Input, Select } from 'antd';
import { tokens } from 'constant/support';
import WalletsETH from 'components/WalletsETH';
import WalletsHmy from 'components/WalletsHmy';
import {
  RightOutlined,
  ArrowRightOutlined,
  ArrowDownOutlined,
  SwapOutlined
} from '@ant-design/icons';

import './style.scss';

const { Option } = Select;

const listWallets = {
  ETH: <WalletsETH />,
  ONE: <WalletsHmy />
};

function SwapPage() {
  const addressETH = useSelector(state => state.eth.address);
  const addressHmy = useSelector(state => state.hmy.address);

  const listAddressDest = useMemo(() => {
    return { 1: addressETH, 0: addressHmy };
  }, [addressETH, addressHmy]);

  const [indexRoadSwap, setIndexRoadSwap] = useState(0);
  const [enanbleBtnSwap, setEnanbleBtnSwap] = useState(false);
  const [amountSource, setAmountSource] = useState();
  const [tokenSource, setTokenSource] = useState();
  const [tokenDest, setTokenDest] = useState();
  const [walletSource, setWalletSource] = useState('ETH');
  const [walletDest, setWalletDest] = useState('ONE');
  const [addressDest, setAddressDest] = useState();
  const [toAddress, setToAddress] = useState();

  useEffect(() => {
    setAddressDest(listAddressDest[indexRoadSwap]);
    setToAddress(listAddressDest[indexRoadSwap]);
  }, [listAddressDest, indexRoadSwap, addressDest]);

  function reverseDirectionToken() {
    setTokenSource(tokenDest);
    setTokenDest(tokenSource);
  }

  function onChangeTokenSource(value) {
    setTokenSource(value);
    if (value === tokenDest) {
      setTokenDest();
    }
  }

  function onChangeTokenDest(value) {
    setTokenDest(value);
    if (value === tokenSource) {
      setTokenSource();
    }
  }

  const onChangeAddressTo = e => {
    const { value } = e.target;
    setToAddress(value);
  };

  const onBlurAddressTo = e => {
    const { value } = e.target;
    setToAddress(value);
  };

  const onChangeFormatNumber = e => {
    const { value } = e.target;
    const reg = /^-?\d*(\.\d*)?$/;
    if ((!isNaN(value) && reg.test(value)) || value === '') {
      setAmountSource(value);
    }
  };

  const chooseRoad = (from, to, index) => {
    setIndexRoadSwap(index);
    setWalletSource(from);
    setWalletDest(to);
    setAddressDest(listAddressDest[index]);
    setToAddress(listAddressDest[index]);
  };

  const onBlurFormatNumber = e => {
    const { value } = e.target;
    let valueTemp = value;
    if (value.charAt(value.length - 1) === '.') {
      valueTemp = value.slice(0, -1);
    }
    setAmountSource(valueTemp.replace(/0*(\d+)/, '$1'));
  };

  function setMyAdress() {
    setToAddress(addressDest);
  }

  return (
    <div className='swap-page'>
      <Row className='container' justify='space-between'>
        <Col xs={{ order: 2, span: 24 }} md={{ order: 1, span: 10 }}>
          <div className='content-swap'>
            <div className='input-and-select-token'>
              <div className='token-source-dest'>
                <div className='label-input-token'>
                  <div className='sc-hSdWYo euiRCS css-1rhdhic'>From</div>
                </div>
                <div className='box-input-token'>
                  <Input
                    size='large'
                    placeholder='0.0'
                    className='input-token'
                    onChange={e => onChangeFormatNumber(e)}
                    onBlur={e => onBlurFormatNumber(e)}
                    value={amountSource}
                  />
                  <Select
                    value={tokenSource}
                    style={{ width: 150 }}
                    showSearch
                    placeholder='Select a token'
                    optionFilterProp='children'
                    onChange={onChangeTokenSource}
                    filterOption={(input, option) =>
                      option.children[1].toLowerCase().indexOf(input.toLowerCase()) >= 0
                    }
                    className='button-select-token'
                  >
                    {tokens.map((token, i) => {
                      return (
                        <Option value={token.symbol} key={i}>
                          <img
                            alt='icon-token'
                            src={token.icon}
                            className='img-icon-token-select-option'
                          />
                          {token.symbol}
                        </Option>
                      );
                    })}
                  </Select>
                </div>
              </div>
              <div className='symbol-arrow-down'>
                <ArrowDownOutlined
                  className='icon'
                  onClick={() => {
                    reverseDirectionToken();
                  }}
                />
              </div>
              <div className='token-source-dest'>
                <div className='label-input-token'>
                  <div className='sc-hSdWYo euiRCS css-1rhdhic'>To</div>
                </div>
                <div className='box-input-token'>
                  <Input size='large' placeholder='0.0' className='input-token' disabled />
                  <Select
                    value={tokenDest}
                    style={{ width: 150 }}
                    showSearch
                    placeholder='Select a token'
                    optionFilterProp='children'
                    onChange={onChangeTokenDest}
                    filterOption={(input, option) =>
                      option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
                    }
                    className='button-select-token'
                  >
                    {tokens.map((token, i) => {
                      return (
                        <Option value={token.symbol} key={i}>
                          <img
                            alt='icon-token'
                            src={token.icon}
                            className='img-icon-token-select-option'
                          />
                          {token.symbol}
                        </Option>
                      );
                    })}
                  </Select>
                </div>
              </div>
              <div>
                <div className='label-input-token'>
                  <div className='sc-hSdWYo euiRCS css-1rhdhic'>To Address</div>
                </div>
                <div className='box-input-token'>
                  <Input
                    size='large'
                    placeholder='address...'
                    className='input-token'
                    value={toAddress}
                    onChange={e => onChangeAddressTo(e)}
                    onBlur={e => onBlurAddressTo(e)}
                  />
                </div>
                {addressDest ? (
                  <div className='button-use-my-address' onClick={() => setMyAdress()}>
                    Use my address
                  </div>
                ) : null}
              </div>
              <div className='button-swap'>
                <button disabled={enanbleBtnSwap} id='swap-button' className='swap-button'>
                  <div className='css-10ob8xa'>
                    <SwapOutlined /> Swap Anyway
                  </div>
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
                  'eth-to-harmony button-chain ' +
                  (indexRoadSwap === 0 ? 'enable-road' : 'disable-road')
                }
                onClick={() => {
                  chooseRoad('ETH', 'ONE', 0);
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
                  'harmony-to-eth button-chain ' +
                  (indexRoadSwap === 1 ? 'enable-road' : 'disable-road')
                }
                onClick={() => {
                  chooseRoad('ONE', 'ETH', 1);
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
                      {listWallets[walletSource]}
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
                      {listWallets[walletDest]}
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
