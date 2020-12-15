// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HarmonyBridge

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HarmonyBridgeABI is the input ABI used to generate the binding from.
const HarmonyBridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"LogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonySender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogNewUnlockClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"LogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"LogUnlockCompleted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"contractBridgeBank\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"completeUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonySender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"newUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unlockClaims\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"harmonySender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumHarmonyBridge.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HarmonyBridgeBin is the compiled bytecode used for deploying new contracts.
var HarmonyBridgeBin = "0x608060405234801561001057600080fd5b50604051611bdc380380611bdc8339818101604052604081101561003357600080fd5b8101908080519060200190929190805190602001909291905050506000600281905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160146101000a81548160ff0219169083151502179055506000600160156101000a81548160ff0219169083151502179055505050611abe8061011e6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637dc0d1d01161008c5780639e67206d116100665780639e67206d146104b8578063b1493ce0146104e6578063cf93b56c14610504578063fb7831f21461054a576100ea565b80637dc0d1d0146103e05780637f54af0c1461042a578063814c92c314610474576100ea565b806360cf6648116100c857806360cf6648146101c957806369294a4e146102575780637adbf973146102795780637c5d7590146102bd576100ea565b80630e41f373146100ef5780631a86e37714610139578063570ca7351461017f575b600080fd5b6100f761056c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101656004803603602081101561014f57600080fd5b8101908080359060200190929190505050610592565b604051808215151515815260200191505060405180910390f35b6101876105d8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610255600480360360808110156101df57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105fd565b005b61025f610d90565b604051808215151515815260200191505060405180910390f35b6102bb6004803603602081101561028f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610da3565b005b6102e9600480360360208110156102d357600080fd5b8101908080359060200190929190505050610fae565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260038111156103c757fe5b60ff168152602001965050505050505060405180910390f35b6103e8611077565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61043261109d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104b66004803603602081101561048a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110c3565b005b6104e4600480360360208110156104ce57600080fd5b81019080803590602001909291905050506112ce565b005b6104ee61146c565b6040518082815260200191505060405180910390f35b6105306004803603602081101561051a57600080fd5b8101908080359060200190929190505050611472565b604051808215151515815260200191505060405180910390f35b61055261158b565b604051808215151515815260200191505060405180910390f35b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160038111156105a157fe5b6005600084815260200190815260200160002060050160009054906101000a900460ff1660038111156105d057fe5b149050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60011515600160149054906101000a900460ff161515148015610633575060011515600160159054906101000a900460ff161515145b610688576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604681526020018061197c6046913960600191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561072757600080fd5b505afa15801561073b573d6000803e3d6000fd5b505050506040513d602081101561075157600080fd5b81019080805190602001909291905050506107d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b60011515600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633b6e750f846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561087757600080fd5b505afa15801561088b573d6000803e3d6000fd5b505050506040513d60208110156108a157600080fd5b8101908080519060200190929190505050151514610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f496e76616c696420546f6b656e0000000000000000000000000000000000000081525060200191505060405180910390fd5b80600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a7f01e4d846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156109c757600080fd5b505afa1580156109db573d6000803e3d6000fd5b505050506040513d60208110156109f157600080fd5b81019080805190602001909291905050501015610a59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611a29603a913960400191505060405180910390fd5b610a616118e2565b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160016003811115610aef57fe5b8152509050610b0a600160025461159e90919063ffffffff16565b6002819055508060056000600254815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff02191690836003811115610c7057fe5b02179055509050507f55d012d5114ba624a65f9c34c00bd550c726eb7fbc784eecd5da93bd8c5cd8ff6002548686338787604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405180910390a15050505050565b600160159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e65576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600160149054906101000a900460ff1615610ecb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260318152602001806119f86031913960400191505060405180910390fd5b60018060146101000a81548160ff02191690831515021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff16905086565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611185576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600160159054906101000a900460ff16156111eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806119c26036913960400191505060405180910390fd5b60018060156101000a81548160ff02191690831515021790555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b806112d881610592565b61134a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f556e6c6f636b20636c61696d206973206e6f742061637469766500000000000081525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113f0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611a636027913960400191505060405180910390fd5b60026005600084815260200190815260200160002060050160006101000a81548160ff0219169083600381111561142357fe5b021790555061143182611626565b7f2ffd08179cc0bd6247ca8f49a84122125e1f5961053a67a73757854f0459124d826040518082815260200191505060405180910390a15050565b60025481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c6005600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561154957600080fd5b505afa15801561155d573d6000803e3d6000fd5b505050506040513d602081101561157357600080fd5b81019080805190602001909291905050509050919050565b600160149054906101000a900460ff1681565b60008082840190508381101561161c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b61162e6118e2565b600560008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff1660038111156117cc57fe5b60038111156117d757fe5b815250509050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359508f8f8260200151836060015184608001516040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156118c657600080fd5b505af11580156118da573d6000803e3d6000fd5b505050505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000600381111561197557fe5b8152509056fe546865204f70657261746f72206d7573742073657420746865206f7261636c6520616e64206272696467652062616e6b20666f72206272696467652061637469766174696f6e546865204272696467652042616e6b2063616e6e6f742062652075706461746564206f6e636520697420686173206265656e20736574546865204f7261636c652063616e6e6f742062652075706461746564206f6e636520697420686173206265656e207365744e6f7420656e6f756768206c6f636b65642061737365747320746f20636f6d706c657465207468652070726f706f7365642070726f70686563794f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f70686563696573a265627a7a7231582031b520cd7b6c3718801d530b48e018559fff2b9bb843df201d65b128a9019e5764736f6c63430005110032"

// DeployHarmonyBridge deploys a new Ethereum contract, binding an instance of HarmonyBridge to it.
func DeployHarmonyBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address) (common.Address, *types.Transaction, *HarmonyBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(HarmonyBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HarmonyBridgeBin), backend, _operator, _valset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HarmonyBridge{HarmonyBridgeCaller: HarmonyBridgeCaller{contract: contract}, HarmonyBridgeTransactor: HarmonyBridgeTransactor{contract: contract}, HarmonyBridgeFilterer: HarmonyBridgeFilterer{contract: contract}}, nil
}

// HarmonyBridge is an auto generated Go binding around an Ethereum contract.
type HarmonyBridge struct {
	HarmonyBridgeCaller     // Read-only binding to the contract
	HarmonyBridgeTransactor // Write-only binding to the contract
	HarmonyBridgeFilterer   // Log filterer for contract events
}

// HarmonyBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HarmonyBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HarmonyBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HarmonyBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HarmonyBridgeSession struct {
	Contract     *HarmonyBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HarmonyBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HarmonyBridgeCallerSession struct {
	Contract *HarmonyBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HarmonyBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HarmonyBridgeTransactorSession struct {
	Contract     *HarmonyBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HarmonyBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HarmonyBridgeRaw struct {
	Contract *HarmonyBridge // Generic contract binding to access the raw methods on
}

// HarmonyBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HarmonyBridgeCallerRaw struct {
	Contract *HarmonyBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// HarmonyBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HarmonyBridgeTransactorRaw struct {
	Contract *HarmonyBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHarmonyBridge creates a new instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridge(address common.Address, backend bind.ContractBackend) (*HarmonyBridge, error) {
	contract, err := bindHarmonyBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridge{HarmonyBridgeCaller: HarmonyBridgeCaller{contract: contract}, HarmonyBridgeTransactor: HarmonyBridgeTransactor{contract: contract}, HarmonyBridgeFilterer: HarmonyBridgeFilterer{contract: contract}}, nil
}

// NewHarmonyBridgeCaller creates a new read-only instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeCaller(address common.Address, caller bind.ContractCaller) (*HarmonyBridgeCaller, error) {
	contract, err := bindHarmonyBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeCaller{contract: contract}, nil
}

// NewHarmonyBridgeTransactor creates a new write-only instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*HarmonyBridgeTransactor, error) {
	contract, err := bindHarmonyBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeTransactor{contract: contract}, nil
}

// NewHarmonyBridgeFilterer creates a new log filterer instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*HarmonyBridgeFilterer, error) {
	contract, err := bindHarmonyBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeFilterer{contract: contract}, nil
}

// bindHarmonyBridge binds a generic wrapper to an already deployed contract.
func bindHarmonyBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HarmonyBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarmonyBridge *HarmonyBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarmonyBridge.Contract.HarmonyBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarmonyBridge *HarmonyBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.HarmonyBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarmonyBridge *HarmonyBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.HarmonyBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarmonyBridge *HarmonyBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarmonyBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarmonyBridge *HarmonyBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarmonyBridge *HarmonyBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) BridgeBank() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeBank(&_HarmonyBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) BridgeBank() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeBank(&_HarmonyBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "hasBridgeBank")
	return *ret0, err
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) HasBridgeBank() (bool, error) {
	return _HarmonyBridge.Contract.HasBridgeBank(&_HarmonyBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) HasBridgeBank() (bool, error) {
	return _HarmonyBridge.Contract.HasBridgeBank(&_HarmonyBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) HasOracle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "hasOracle")
	return *ret0, err
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) HasOracle() (bool, error) {
	return _HarmonyBridge.Contract.HasOracle(&_HarmonyBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) HasOracle() (bool, error) {
	return _HarmonyBridge.Contract.HasOracle(&_HarmonyBridge.CallOpts)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) IsUnlockClaimActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "isUnlockClaimActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) IsUnlockClaimValidatorActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "isUnlockClaimValidatorActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimValidatorActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimValidatorActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Operator() (common.Address, error) {
	return _HarmonyBridge.Contract.Operator(&_HarmonyBridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Operator() (common.Address, error) {
	return _HarmonyBridge.Contract.Operator(&_HarmonyBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Oracle() (common.Address, error) {
	return _HarmonyBridge.Contract.Oracle(&_HarmonyBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Oracle() (common.Address, error) {
	return _HarmonyBridge.Contract.Oracle(&_HarmonyBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCaller) UnlockClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "unlockClaimCount")
	return *ret0, err
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeSession) UnlockClaimCount() (*big.Int, error) {
	return _HarmonyBridge.Contract.UnlockClaimCount(&_HarmonyBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCallerSession) UnlockClaimCount() (*big.Int, error) {
	return _HarmonyBridge.Contract.UnlockClaimCount(&_HarmonyBridge.CallOpts)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeCaller) UnlockClaims(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		HarmonySender     common.Address
		EthereumReceiver  common.Address
		OriginalValidator common.Address
		Token             common.Address
		Amount            *big.Int
		Status            uint8
	})
	out := ret
	err := _HarmonyBridge.contract.Call(opts, out, "unlockClaims", arg0)
	return *ret, err
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeSession) UnlockClaims(arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _HarmonyBridge.Contract.UnlockClaims(&_HarmonyBridge.CallOpts, arg0)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeCallerSession) UnlockClaims(arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _HarmonyBridge.Contract.UnlockClaims(&_HarmonyBridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Valset() (common.Address, error) {
	return _HarmonyBridge.Contract.Valset(&_HarmonyBridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Valset() (common.Address, error) {
	return _HarmonyBridge.Contract.Valset(&_HarmonyBridge.CallOpts)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) CompleteUnlockClaim(opts *bind.TransactOpts, _unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "completeUnlockClaim", _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.CompleteUnlockClaim(&_HarmonyBridge.TransactOpts, _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.CompleteUnlockClaim(&_HarmonyBridge.TransactOpts, _unlockID)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) NewUnlockClaim(opts *bind.TransactOpts, _harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "newUnlockClaim", _harmonySender, _ethereumReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeSession) NewUnlockClaim(_harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.NewUnlockClaim(&_HarmonyBridge.TransactOpts, _harmonySender, _ethereumReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) NewUnlockClaim(_harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.NewUnlockClaim(&_HarmonyBridge.TransactOpts, _harmonySender, _ethereumReceiver, _token, _amount)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_HarmonyBridge *HarmonyBridgeSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.SetBridgeBank(&_HarmonyBridge.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.SetBridgeBank(&_HarmonyBridge.TransactOpts, _bridgeBank)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_HarmonyBridge *HarmonyBridgeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.SetOracle(&_HarmonyBridge.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.SetOracle(&_HarmonyBridge.TransactOpts, _oracle)
}

// HarmonyBridgeLogBridgeBankSetIterator is returned from FilterLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for LogBridgeBankSet events raised by the HarmonyBridge contract.
type HarmonyBridgeLogBridgeBankSetIterator struct {
	Event *HarmonyBridgeLogBridgeBankSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HarmonyBridgeLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeLogBridgeBankSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HarmonyBridgeLogBridgeBankSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HarmonyBridgeLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeLogBridgeBankSet represents a LogBridgeBankSet event raised by the HarmonyBridge contract.
type HarmonyBridgeLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeBankSet is a free log retrieval operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterLogBridgeBankSet(opts *bind.FilterOpts) (*HarmonyBridgeLogBridgeBankSetIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeLogBridgeBankSetIterator{contract: _HarmonyBridge.contract, event: "LogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchLogBridgeBankSet is a free log subscription operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeLogBridgeBankSet)
				if err := _HarmonyBridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBridgeBankSet is a log parse operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseLogBridgeBankSet(log types.Log) (*HarmonyBridgeLogBridgeBankSet, error) {
	event := new(HarmonyBridgeLogBridgeBankSet)
	if err := _HarmonyBridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarmonyBridgeLogNewUnlockClaimIterator is returned from FilterLogNewUnlockClaim and is used to iterate over the raw logs and unpacked data for LogNewUnlockClaim events raised by the HarmonyBridge contract.
type HarmonyBridgeLogNewUnlockClaimIterator struct {
	Event *HarmonyBridgeLogNewUnlockClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HarmonyBridgeLogNewUnlockClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeLogNewUnlockClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HarmonyBridgeLogNewUnlockClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HarmonyBridgeLogNewUnlockClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeLogNewUnlockClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeLogNewUnlockClaim represents a LogNewUnlockClaim event raised by the HarmonyBridge contract.
type HarmonyBridgeLogNewUnlockClaim struct {
	UnlockID         *big.Int
	HarmonySender    common.Address
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewUnlockClaim is a free log retrieval operation binding the contract event 0x55d012d5114ba624a65f9c34c00bd550c726eb7fbc784eecd5da93bd8c5cd8ff.
//
// Solidity: event LogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterLogNewUnlockClaim(opts *bind.FilterOpts) (*HarmonyBridgeLogNewUnlockClaimIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "LogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeLogNewUnlockClaimIterator{contract: _HarmonyBridge.contract, event: "LogNewUnlockClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewUnlockClaim is a free log subscription operation binding the contract event 0x55d012d5114ba624a65f9c34c00bd550c726eb7fbc784eecd5da93bd8c5cd8ff.
//
// Solidity: event LogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchLogNewUnlockClaim(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeLogNewUnlockClaim) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "LogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeLogNewUnlockClaim)
				if err := _HarmonyBridge.contract.UnpackLog(event, "LogNewUnlockClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewUnlockClaim is a log parse operation binding the contract event 0x55d012d5114ba624a65f9c34c00bd550c726eb7fbc784eecd5da93bd8c5cd8ff.
//
// Solidity: event LogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseLogNewUnlockClaim(log types.Log) (*HarmonyBridgeLogNewUnlockClaim, error) {
	event := new(HarmonyBridgeLogNewUnlockClaim)
	if err := _HarmonyBridge.contract.UnpackLog(event, "LogNewUnlockClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarmonyBridgeLogOracleSetIterator is returned from FilterLogOracleSet and is used to iterate over the raw logs and unpacked data for LogOracleSet events raised by the HarmonyBridge contract.
type HarmonyBridgeLogOracleSetIterator struct {
	Event *HarmonyBridgeLogOracleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HarmonyBridgeLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeLogOracleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HarmonyBridgeLogOracleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HarmonyBridgeLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeLogOracleSet represents a LogOracleSet event raised by the HarmonyBridge contract.
type HarmonyBridgeLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogOracleSet is a free log retrieval operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterLogOracleSet(opts *bind.FilterOpts) (*HarmonyBridgeLogOracleSetIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeLogOracleSetIterator{contract: _HarmonyBridge.contract, event: "LogOracleSet", logs: logs, sub: sub}, nil
}

// WatchLogOracleSet is a free log subscription operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchLogOracleSet(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeLogOracleSet)
				if err := _HarmonyBridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogOracleSet is a log parse operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseLogOracleSet(log types.Log) (*HarmonyBridgeLogOracleSet, error) {
	event := new(HarmonyBridgeLogOracleSet)
	if err := _HarmonyBridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarmonyBridgeLogUnlockCompletedIterator is returned from FilterLogUnlockCompleted and is used to iterate over the raw logs and unpacked data for LogUnlockCompleted events raised by the HarmonyBridge contract.
type HarmonyBridgeLogUnlockCompletedIterator struct {
	Event *HarmonyBridgeLogUnlockCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HarmonyBridgeLogUnlockCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeLogUnlockCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HarmonyBridgeLogUnlockCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HarmonyBridgeLogUnlockCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeLogUnlockCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeLogUnlockCompleted represents a LogUnlockCompleted event raised by the HarmonyBridge contract.
type HarmonyBridgeLogUnlockCompleted struct {
	UnlockID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUnlockCompleted is a free log retrieval operation binding the contract event 0x2ffd08179cc0bd6247ca8f49a84122125e1f5961053a67a73757854f0459124d.
//
// Solidity: event LogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterLogUnlockCompleted(opts *bind.FilterOpts) (*HarmonyBridgeLogUnlockCompletedIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "LogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeLogUnlockCompletedIterator{contract: _HarmonyBridge.contract, event: "LogUnlockCompleted", logs: logs, sub: sub}, nil
}

// WatchLogUnlockCompleted is a free log subscription operation binding the contract event 0x2ffd08179cc0bd6247ca8f49a84122125e1f5961053a67a73757854f0459124d.
//
// Solidity: event LogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchLogUnlockCompleted(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeLogUnlockCompleted) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "LogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeLogUnlockCompleted)
				if err := _HarmonyBridge.contract.UnpackLog(event, "LogUnlockCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUnlockCompleted is a log parse operation binding the contract event 0x2ffd08179cc0bd6247ca8f49a84122125e1f5961053a67a73757854f0459124d.
//
// Solidity: event LogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseLogUnlockCompleted(log types.Log) (*HarmonyBridgeLogUnlockCompleted, error) {
	event := new(HarmonyBridgeLogUnlockCompleted)
	if err := _HarmonyBridge.contract.UnpackLog(event, "LogUnlockCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
