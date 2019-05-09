// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisessionapp

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleMultiSessionAppABI is the input ABI used to generate the binding from.
const SimpleMultiSessionAppABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"}],\"name\":\"getSettleFinalizedTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stateProof\",\"type\":\"bytes\"}],\"name\":\"intendSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"}],\"name\":\"getSeqNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"getSessionID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"}],\"name\":\"finalizeOnActionTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"}],\"name\":\"getActionDeadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"},{\"name\":\"_action\",\"type\":\"bytes\"}],\"name\":\"applyAction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_timeout\",\"type\":\"uint256\"},{\"name\":\"_playerNum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"session\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"seq\",\"type\":\"uint256\"}],\"name\":\"IntendSettle\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_session\",\"type\":\"bytes32\"},{\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimpleMultiSessionAppBin is the compiled bytecode used for deploying new contracts.
const SimpleMultiSessionAppBin = `0x608060405234801561001057600080fd5b506040516040806120938339810180604052604081101561003057600080fd5b810190808051906020019092919080519060200190929190505050818181818160008001819055508060006001018190555050505050505061201c806100776000396000f3fe6080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806309b65d86146100b4578063130d33fe1461010357806329dd2f8e146101895780633b6de66f146102475780634d8bedec146102965780635de28ae01461033a578063b89fa28b14610397578063bcdbda94146103d2578063cab9244614610470578063ef4592fb146104bf578063f3c771921461055d575b600080fd5b3480156100c057600080fd5b506100ed600480360360208110156100d757600080fd5b81019080803590602001909291905050506105ed565b6040518082815260200191505060405180910390f35b34801561010f57600080fd5b506101876004803603602081101561012657600080fd5b810190808035906020019064010000000081111561014357600080fd5b82018360208201111561015557600080fd5b8035906020019184600183028401116401000000008311171561017757600080fd5b9091929391929390505050610659565b005b34801561019557600080fd5b506101cc600480360360408110156101ac57600080fd5b810190808035906020019092919080359060200190929190505050610a63565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561020c5780820151818401526020810190506101f1565b50505050905090810190601f1680156102395780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025357600080fd5b506102806004803603602081101561026a57600080fd5b8101908080359060200190929190505050610ad6565b6040518082815260200191505060405180910390f35b3480156102a257600080fd5b50610324600480360360408110156102b957600080fd5b8101908080359060200190929190803590602001906401000000008111156102e057600080fd5b8201836020820111156102f257600080fd5b8035906020019184602083028401116401000000008311171561031457600080fd5b9091929391929390505050610af6565b6040518082815260200191505060405180910390f35b34801561034657600080fd5b506103736004803603602081101561035d57600080fd5b8101908080359060200190929190505050610b5f565b6040518082600381111561038357fe5b60ff16815260200191505060405180910390f35b3480156103a357600080fd5b506103d0600480360360208110156103ba57600080fd5b8101908080359060200190929190505050610b8c565b005b3480156103de57600080fd5b50610456600480360360208110156103f557600080fd5b810190808035906020019064010000000081111561041257600080fd5b82018360208201111561042457600080fd5b8035906020019184600183028401116401000000008311171561044657600080fd5b9091929391929390505050610c50565b604051808215151515815260200191505060405180910390f35b34801561047c57600080fd5b506104a96004803603602081101561049357600080fd5b8101908080359060200190929190505050610ce7565b6040518082815260200191505060405180910390f35b3480156104cb57600080fd5b50610543600480360360208110156104e257600080fd5b81019080803590602001906401000000008111156104ff57600080fd5b82018360208201111561051157600080fd5b8035906020019184600183028401116401000000008311171561053357600080fd5b9091929391929390505050610db9565b604051808215151515815260200191505060405180910390f35b34801561056957600080fd5b506105eb6004803603604081101561058057600080fd5b8101908080359060200190929190803590602001906401000000008111156105a757600080fd5b8201836020820111156105b957600080fd5b803590602001918460018302840111640100000000831117156105db57600080fd5b9091929391929390505050610e2d565b005b6000600160038111156105fc57fe5b6002600084815260200190815260200160002060030160009054906101000a900460ff16600381111561062b57fe5b141561064f5760026000838152602001908152602001600020600201549050610654565b600090505b919050565b610661611eb0565b6106ae83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061106a565b905060606106c68260000151836020015160016111f5565b90506000600101548151141515610745576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f696e76616c6964206e756d626572206f6620706c61796572730000000000000081525060200191505060405180910390fd5b61074d611eca565b61075a8360000151611420565b905060008160000151836040516020018083815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156107b0578082015181840152602081019050610795565b5050505090500193505050506040516020818303038152906040528051906020012090506000600260008381526020019081526020016000209050600060038111156107f857fe5b8160030160009054906101000a900460ff16600381111561081557fe5b14156108355783816000019080519060200190610833929190611eec565b505b60038081111561084157fe5b8160030160009054906101000a900460ff16600381111561085e57fe5b141515156108d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f73746174652069732066696e616c697a6564000000000000000000000000000081525060200191505060405180910390fd5b82602001518160010154101515610953576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c69642073657175656e6365206e756d62657200000000000000000081525060200191505060405180910390fd5b8260200151816001018190555060018160030160006101000a81548160ff0219169083600381111561098157fe5b02179055506000800154430181600201819055506109a38284604001516114ed565b1515610a17576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f737461746520757064617465206661696c65640000000000000000000000000081525060200191505060405180910390fd5b7f82c4eeba939ff9358877334330e22a5cdb0472113cd14f90625ea634b60d2e5b828260010154604051808381526020018281526020019250505060405180910390a150505050505050565b60608060206040519080825280601f01601f191660200182016040528015610a9a5781602001600182028038833980820191505090505b50905060006003600086815260200190815260200160002060000160009054906101000a900460ff169050806020830152819250505092915050565b600060026000838152602001908152602001600020600101549050919050565b600083838360405160200180848152602001806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509450505050506040516020818303038152906040528051906020012090509392505050565b60006002600083815260200190815260200160002060030160009054906101000a900460ff169050919050565b60006002600083815260200190815260200160002060030160009054906101000a900460ff16905060006002600084815260200190815260200160002060020154905060026003811115610bdc57fe5b826003811115610be857fe5b1415610c01578043111515610bfc57600080fd5b610c41565b60016003811115610c0e57fe5b826003811115610c1a57fe5b1415610c39576000800154810143111515610c3457600080fd5b610c40565b5050610c4d565b5b610c4a83611698565b50505b50565b600080610ca084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506116d3565b9050600380811115610cae57fe5b6002600083815260200190815260200160002060030160009054906101000a900460ff166003811115610cdd57fe5b1491505092915050565b600060026003811115610cf657fe5b6002600084815260200190815260200160002060030160009054906101000a900460ff166003811115610d2557fe5b1415610d495760026000838152602001908152602001600020600201549050610db4565b60016003811115610d5657fe5b6002600084815260200190815260200160002060030160009054906101000a900460ff166003811115610d8557fe5b1415610daf5760008001546002600084815260200190815260200160002060020154019050610db4565b600090505b919050565b6000610dc3611f76565b610e1084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506116f1565b9050610e24816000015182602001516117a5565b91505092915050565b6000600260008581526020019081526020016000209050600380811115610e5057fe5b8160030160009054906101000a900460ff166003811115610e6d57fe5b14151515610ee3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6170702073746174652069732066696e616c697a65640000000000000000000081525060200191505060405180910390fd5b60016003811115610ef057fe5b8160030160009054906101000a900460ff166003811115610f0d57fe5b148015610f1d5750438160020154105b15610f495760028160030160006101000a81548160ff02191690836003811115610f4357fe5b02179055505b60026003811115610f5657fe5b8160030160009054906101000a900460ff166003811115610f7357fe5b141515610fe8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f617070206e6f7420696e20616374696f6e206d6f64650000000000000000000081525060200191505060405180910390fd5b80600101600081548092919060010191905055506000800154430181600201819055506110598484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506118d0565b151561106457600080fd5b50505050565b611072611eb0565b61107a611f93565b61108383611a7b565b9050606061109b600283611aac90919063ffffffff16565b90508060028151811015156110ac57fe5b906020019060200201516040519080825280602002602001820160405280156110e957816020015b60608152602001906001900390816110d45790505b508360200181905250600081600281518110151561110357fe5b90602001906020020181815250506000805b61111e84611b55565b156111ec5761112c84611b6a565b8092508193505050600015611140576111e7565b600182141561115f5761115284611ba0565b85600001819052506111e6565b60028214156111d15761117184611ba0565b856020015184600281518110151561118557fe5b9060200190602002015181518110151561119b57fe5b906020019060200201819052508260028151811015156111b757fe5b9060200190602002018051809190600101815250506111e5565b6111e48185611c5b90919063ffffffff16565b5b5b5b611115565b50505050919050565b60608083516040519080825280602002602001820160405280156112285781602001602082028038833980820191505090505b50905060006112a9866040516020018082805190602001908083835b6020831015156112695780518252602082019150602081019050602083039250611244565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120611ced565b9050600080905060008090505b8651811015611412576112e08388838151811015156112d157fe5b90602001906020020151611d45565b84828151811015156112ee57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508515611405578173ffffffffffffffffffffffffffffffffffffffff16848281518110151561135557fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff161115156113ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f7369676e657273206e6f7420696e20617363656e64696e67206f72646572000081525060200191505060405180910390fd5b83818151811015156113f857fe5b9060200190602002015191505b80806001019150506112b6565b508293505050509392505050565b611428611eca565b611430611f93565b61143983611a7b565b90506000805b61144883611b55565b156114e55761145683611b6a565b809250819350505060001561146a576114e0565b600182141561148a5761147c83611e27565b8460000181815250506114df565b60028214156114aa5761149c83611e27565b8460200181815250506114de565b60038214156114c9576114bc83611ba0565b84604001819052506114dd565b6114dc8184611c5b90919063ffffffff16565b5b5b5b5b61143f565b505050919050565b600060018251141515611568576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c6964207374617465206c656e67746800000000000000000000000081525060200191505060405180910390fd5b600060036000858152602001908152602001600020905082600081518110151561158e57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f010000000000000000000000000000000000000000000000000000000000000090048160000160006101000a81548160ff021916908360ff16021790555060018160000160009054906101000a900460ff1660ff16148061164f575060028160000160009054906101000a900460ff1660ff16145b1561168d5760036002600086815260200190815260200160002060030160006101000a81548160ff0219169083600381111561168757fe5b02179055505b600191505092915050565b60036002600083815260200190815260200160002060030160006101000a81548160ff021916908360038111156116cb57fe5b021790555050565b6000602082511415156116e557600080fd5b60208201519050919050565b6116f9611f76565b611701611f93565b61170a83611a7b565b90506000805b61171983611b55565b1561179d5761172783611b6a565b809250819350505060001561173b57611798565b60018214156117635761175561175084611ba0565b6116d3565b846000018181525050611797565b60028214156117825761177583611ba0565b8460200181905250611796565b6117958184611c5b90919063ffffffff16565b5b5b5b611710565b505050919050565b600060018251141515611820576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c6964207175657279206c656e67746800000000000000000000000081525060200191505060405180910390fd5b81600081518110151561182f57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900460ff166003600085815260200190815260200160002060000160009054906101000a900460ff1660ff1614905092915050565b60006001825114151561194b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c696420616374696f6e206c656e677468000000000000000000000081525060200191505060405180910390fd5b600060036000858152602001908152602001600020905082600081518110151561197157fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f010000000000000000000000000000000000000000000000000000000000000090048160000160006101000a81548160ff021916908360ff16021790555060018160000160009054906101000a900460ff1660ff161480611a32575060028160000160009054906101000a900460ff1660ff16145b15611a705760036002600086815260200190815260200160002060030160006101000a81548160ff02191690836003811115611a6a57fe5b02179055505b600191505092915050565b611a83611f93565b60018251111515611a9357600080fd5b8181602001819052506000816000018181525050919050565b606060008360000151905060018301604051908082528060200260200182016040528015611ae95781602001602082028038833980820191505090505b5091506000805b611af986611b55565b15611b4257611b0786611b6a565b809250819350505060018483815181101515611b1f57fe5b9060200190602002018181510191508181525050611b3d8682611c5b565b611af0565b8286600001818152505050505092915050565b60008160200151518260000151109050919050565b6000806000611b7884611e27565b9050600881811515611b8657fe5b049250600781166005811115611b9857fe5b915050915091565b60606000611bad83611e27565b905060008184600001510190508360200151518111151515611bce57600080fd5b816040519080825280601f01601f191660200182016040528015611c015781602001600182028038833980820191505090505b50925060608460200151905060008086600001519050602086019150806020840101905060008090505b85811015611c46578082015181840152602081019050611c2b565b50838760000181815250505050505050919050565b60006005811115611c6857fe5b816005811115611c7457fe5b1415611c8957611c8382611e27565b50611ce9565b60026005811115611c9657fe5b816005811115611ca257fe5b1415611ce3576000611cb383611e27565b905080836000018181510191508181525050826020015151836000015111151515611cdd57600080fd5b50611ce8565b600080fd5b5b5050565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c01828152602001915050604051602081830303815290604052805190602001209050919050565b600060418251141515611d5b5760009050611e21565b60008060006020850151925060408501519150606085015160001a9050601b8160ff161015611d8b57601b810190505b601b8160ff1614158015611da35750601c8160ff1614155b15611db45760009350505050611e21565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611e11573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60008060608360200151905083600001519250826020820101519150600080935060008090505b600a811015611ea55783811a915060078102607f83169060020a02851794506000608083161415611e98576001810186600001818151019150818152505084945050505050611eab565b8080600101915050611e4e565b50600080fd5b919050565b604080519081016040528060608152602001606081525090565b6060604051908101604052806000815260200160008152602001606081525090565b828054828255906000526020600020908101928215611f65579160200282015b82811115611f645782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611f0c565b5b509050611f729190611fad565b5090565b604080519081016040528060008019168152602001606081525090565b604080519081016040528060008152602001606081525090565b611fed91905b80821115611fe957600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611fb3565b5090565b9056fea165627a7a723058201bc37429ad50a84293c4bff12124b837f58f4a3828e7cef614e40af7fed7f6fe0029`

// DeploySimpleMultiSessionApp deploys a new Ethereum contract, binding an instance of SimpleMultiSessionApp to it.
func DeploySimpleMultiSessionApp(auth *bind.TransactOpts, backend bind.ContractBackend, _timeout *big.Int, _playerNum *big.Int) (common.Address, *types.Transaction, *SimpleMultiSessionApp, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMultiSessionAppABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleMultiSessionAppBin), backend, _timeout, _playerNum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleMultiSessionApp{SimpleMultiSessionAppCaller: SimpleMultiSessionAppCaller{contract: contract}, SimpleMultiSessionAppTransactor: SimpleMultiSessionAppTransactor{contract: contract}, SimpleMultiSessionAppFilterer: SimpleMultiSessionAppFilterer{contract: contract}}, nil
}

// SimpleMultiSessionApp is an auto generated Go binding around an Ethereum contract.
type SimpleMultiSessionApp struct {
	SimpleMultiSessionAppCaller     // Read-only binding to the contract
	SimpleMultiSessionAppTransactor // Write-only binding to the contract
	SimpleMultiSessionAppFilterer   // Log filterer for contract events
}

// SimpleMultiSessionAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleMultiSessionAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSessionAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleMultiSessionAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSessionAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleMultiSessionAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSessionAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleMultiSessionAppSession struct {
	Contract     *SimpleMultiSessionApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleMultiSessionAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleMultiSessionAppCallerSession struct {
	Contract *SimpleMultiSessionAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SimpleMultiSessionAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleMultiSessionAppTransactorSession struct {
	Contract     *SimpleMultiSessionAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SimpleMultiSessionAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleMultiSessionAppRaw struct {
	Contract *SimpleMultiSessionApp // Generic contract binding to access the raw methods on
}

// SimpleMultiSessionAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleMultiSessionAppCallerRaw struct {
	Contract *SimpleMultiSessionAppCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleMultiSessionAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleMultiSessionAppTransactorRaw struct {
	Contract *SimpleMultiSessionAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleMultiSessionApp creates a new instance of SimpleMultiSessionApp, bound to a specific deployed contract.
func NewSimpleMultiSessionApp(address common.Address, backend bind.ContractBackend) (*SimpleMultiSessionApp, error) {
	contract, err := bindSimpleMultiSessionApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSessionApp{SimpleMultiSessionAppCaller: SimpleMultiSessionAppCaller{contract: contract}, SimpleMultiSessionAppTransactor: SimpleMultiSessionAppTransactor{contract: contract}, SimpleMultiSessionAppFilterer: SimpleMultiSessionAppFilterer{contract: contract}}, nil
}

// NewSimpleMultiSessionAppCaller creates a new read-only instance of SimpleMultiSessionApp, bound to a specific deployed contract.
func NewSimpleMultiSessionAppCaller(address common.Address, caller bind.ContractCaller) (*SimpleMultiSessionAppCaller, error) {
	contract, err := bindSimpleMultiSessionApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSessionAppCaller{contract: contract}, nil
}

// NewSimpleMultiSessionAppTransactor creates a new write-only instance of SimpleMultiSessionApp, bound to a specific deployed contract.
func NewSimpleMultiSessionAppTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleMultiSessionAppTransactor, error) {
	contract, err := bindSimpleMultiSessionApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSessionAppTransactor{contract: contract}, nil
}

// NewSimpleMultiSessionAppFilterer creates a new log filterer instance of SimpleMultiSessionApp, bound to a specific deployed contract.
func NewSimpleMultiSessionAppFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleMultiSessionAppFilterer, error) {
	contract, err := bindSimpleMultiSessionApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSessionAppFilterer{contract: contract}, nil
}

// bindSimpleMultiSessionApp binds a generic wrapper to an already deployed contract.
func bindSimpleMultiSessionApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMultiSessionAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleMultiSessionApp.Contract.SimpleMultiSessionAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.SimpleMultiSessionAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.SimpleMultiSessionAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleMultiSessionApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.contract.Transact(opts, method, params...)
}

// GetActionDeadline is a free data retrieval call binding the contract method 0xcab92446.
//
// Solidity: function getActionDeadline(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetActionDeadline(opts *bind.CallOpts, _session [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getActionDeadline", _session)
	return *ret0, err
}

// GetActionDeadline is a free data retrieval call binding the contract method 0xcab92446.
//
// Solidity: function getActionDeadline(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetActionDeadline(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetActionDeadline(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetActionDeadline is a free data retrieval call binding the contract method 0xcab92446.
//
// Solidity: function getActionDeadline(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetActionDeadline(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetActionDeadline(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetResult(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getResult", _query)
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetResult(_query []byte) (bool, error) {
	return _SimpleMultiSessionApp.Contract.GetResult(&_SimpleMultiSessionApp.CallOpts, _query)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetResult(_query []byte) (bool, error) {
	return _SimpleMultiSessionApp.Contract.GetResult(&_SimpleMultiSessionApp.CallOpts, _query)
}

// GetSeqNum is a free data retrieval call binding the contract method 0x3b6de66f.
//
// Solidity: function getSeqNum(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetSeqNum(opts *bind.CallOpts, _session [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getSeqNum", _session)
	return *ret0, err
}

// GetSeqNum is a free data retrieval call binding the contract method 0x3b6de66f.
//
// Solidity: function getSeqNum(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetSeqNum(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetSeqNum(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetSeqNum is a free data retrieval call binding the contract method 0x3b6de66f.
//
// Solidity: function getSeqNum(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetSeqNum(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetSeqNum(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetSessionID is a free data retrieval call binding the contract method 0x4d8bedec.
//
// Solidity: function getSessionID(uint256 _nonce, address[] _signers) constant returns(bytes32)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetSessionID(opts *bind.CallOpts, _nonce *big.Int, _signers []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getSessionID", _nonce, _signers)
	return *ret0, err
}

// GetSessionID is a free data retrieval call binding the contract method 0x4d8bedec.
//
// Solidity: function getSessionID(uint256 _nonce, address[] _signers) constant returns(bytes32)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetSessionID(_nonce *big.Int, _signers []common.Address) ([32]byte, error) {
	return _SimpleMultiSessionApp.Contract.GetSessionID(&_SimpleMultiSessionApp.CallOpts, _nonce, _signers)
}

// GetSessionID is a free data retrieval call binding the contract method 0x4d8bedec.
//
// Solidity: function getSessionID(uint256 _nonce, address[] _signers) constant returns(bytes32)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetSessionID(_nonce *big.Int, _signers []common.Address) ([32]byte, error) {
	return _SimpleMultiSessionApp.Contract.GetSessionID(&_SimpleMultiSessionApp.CallOpts, _nonce, _signers)
}

// GetSettleFinalizedTime is a free data retrieval call binding the contract method 0x09b65d86.
//
// Solidity: function getSettleFinalizedTime(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetSettleFinalizedTime(opts *bind.CallOpts, _session [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getSettleFinalizedTime", _session)
	return *ret0, err
}

// GetSettleFinalizedTime is a free data retrieval call binding the contract method 0x09b65d86.
//
// Solidity: function getSettleFinalizedTime(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetSettleFinalizedTime(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetSettleFinalizedTime(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetSettleFinalizedTime is a free data retrieval call binding the contract method 0x09b65d86.
//
// Solidity: function getSettleFinalizedTime(bytes32 _session) constant returns(uint256)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetSettleFinalizedTime(_session [32]byte) (*big.Int, error) {
	return _SimpleMultiSessionApp.Contract.GetSettleFinalizedTime(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetState is a free data retrieval call binding the contract method 0x29dd2f8e.
//
// Solidity: function getState(bytes32 _session, uint256 _key) constant returns(bytes)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetState(opts *bind.CallOpts, _session [32]byte, _key *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getState", _session, _key)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x29dd2f8e.
//
// Solidity: function getState(bytes32 _session, uint256 _key) constant returns(bytes)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetState(_session [32]byte, _key *big.Int) ([]byte, error) {
	return _SimpleMultiSessionApp.Contract.GetState(&_SimpleMultiSessionApp.CallOpts, _session, _key)
}

// GetState is a free data retrieval call binding the contract method 0x29dd2f8e.
//
// Solidity: function getState(bytes32 _session, uint256 _key) constant returns(bytes)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetState(_session [32]byte, _key *big.Int) ([]byte, error) {
	return _SimpleMultiSessionApp.Contract.GetState(&_SimpleMultiSessionApp.CallOpts, _session, _key)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _session) constant returns(uint8)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) GetStatus(opts *bind.CallOpts, _session [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "getStatus", _session)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _session) constant returns(uint8)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) GetStatus(_session [32]byte) (uint8, error) {
	return _SimpleMultiSessionApp.Contract.GetStatus(&_SimpleMultiSessionApp.CallOpts, _session)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _session) constant returns(uint8)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) GetStatus(_session [32]byte) (uint8, error) {
	return _SimpleMultiSessionApp.Contract.GetStatus(&_SimpleMultiSessionApp.CallOpts, _session)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCaller) IsFinalized(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimpleMultiSessionApp.contract.Call(opts, out, "isFinalized", _query)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) IsFinalized(_query []byte) (bool, error) {
	return _SimpleMultiSessionApp.Contract.IsFinalized(&_SimpleMultiSessionApp.CallOpts, _query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppCallerSession) IsFinalized(_query []byte) (bool, error) {
	return _SimpleMultiSessionApp.Contract.IsFinalized(&_SimpleMultiSessionApp.CallOpts, _query)
}

// ApplyAction is a paid mutator transaction binding the contract method 0xf3c77192.
//
// Solidity: function applyAction(bytes32 _session, bytes _action) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactor) ApplyAction(opts *bind.TransactOpts, _session [32]byte, _action []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.contract.Transact(opts, "applyAction", _session, _action)
}

// ApplyAction is a paid mutator transaction binding the contract method 0xf3c77192.
//
// Solidity: function applyAction(bytes32 _session, bytes _action) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) ApplyAction(_session [32]byte, _action []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.ApplyAction(&_SimpleMultiSessionApp.TransactOpts, _session, _action)
}

// ApplyAction is a paid mutator transaction binding the contract method 0xf3c77192.
//
// Solidity: function applyAction(bytes32 _session, bytes _action) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactorSession) ApplyAction(_session [32]byte, _action []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.ApplyAction(&_SimpleMultiSessionApp.TransactOpts, _session, _action)
}

// FinalizeOnActionTimeout is a paid mutator transaction binding the contract method 0xb89fa28b.
//
// Solidity: function finalizeOnActionTimeout(bytes32 _session) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactor) FinalizeOnActionTimeout(opts *bind.TransactOpts, _session [32]byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.contract.Transact(opts, "finalizeOnActionTimeout", _session)
}

// FinalizeOnActionTimeout is a paid mutator transaction binding the contract method 0xb89fa28b.
//
// Solidity: function finalizeOnActionTimeout(bytes32 _session) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) FinalizeOnActionTimeout(_session [32]byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.FinalizeOnActionTimeout(&_SimpleMultiSessionApp.TransactOpts, _session)
}

// FinalizeOnActionTimeout is a paid mutator transaction binding the contract method 0xb89fa28b.
//
// Solidity: function finalizeOnActionTimeout(bytes32 _session) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactorSession) FinalizeOnActionTimeout(_session [32]byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.FinalizeOnActionTimeout(&_SimpleMultiSessionApp.TransactOpts, _session)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x130d33fe.
//
// Solidity: function intendSettle(bytes _stateProof) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactor) IntendSettle(opts *bind.TransactOpts, _stateProof []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.contract.Transact(opts, "intendSettle", _stateProof)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x130d33fe.
//
// Solidity: function intendSettle(bytes _stateProof) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppSession) IntendSettle(_stateProof []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.IntendSettle(&_SimpleMultiSessionApp.TransactOpts, _stateProof)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x130d33fe.
//
// Solidity: function intendSettle(bytes _stateProof) returns()
func (_SimpleMultiSessionApp *SimpleMultiSessionAppTransactorSession) IntendSettle(_stateProof []byte) (*types.Transaction, error) {
	return _SimpleMultiSessionApp.Contract.IntendSettle(&_SimpleMultiSessionApp.TransactOpts, _stateProof)
}

// SimpleMultiSessionAppIntendSettleIterator is returned from FilterIntendSettle and is used to iterate over the raw logs and unpacked data for IntendSettle events raised by the SimpleMultiSessionApp contract.
type SimpleMultiSessionAppIntendSettleIterator struct {
	Event *SimpleMultiSessionAppIntendSettle // Event containing the contract specifics and raw log

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
func (it *SimpleMultiSessionAppIntendSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleMultiSessionAppIntendSettle)
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
		it.Event = new(SimpleMultiSessionAppIntendSettle)
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
func (it *SimpleMultiSessionAppIntendSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleMultiSessionAppIntendSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleMultiSessionAppIntendSettle represents a IntendSettle event raised by the SimpleMultiSessionApp contract.
type SimpleMultiSessionAppIntendSettle struct {
	Session [32]byte
	Seq     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIntendSettle is a free log retrieval operation binding the contract event 0x82c4eeba939ff9358877334330e22a5cdb0472113cd14f90625ea634b60d2e5b.
//
// Solidity: event IntendSettle(bytes32 session, uint256 seq)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppFilterer) FilterIntendSettle(opts *bind.FilterOpts) (*SimpleMultiSessionAppIntendSettleIterator, error) {

	logs, sub, err := _SimpleMultiSessionApp.contract.FilterLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSessionAppIntendSettleIterator{contract: _SimpleMultiSessionApp.contract, event: "IntendSettle", logs: logs, sub: sub}, nil
}

// WatchIntendSettle is a free log subscription operation binding the contract event 0x82c4eeba939ff9358877334330e22a5cdb0472113cd14f90625ea634b60d2e5b.
//
// Solidity: event IntendSettle(bytes32 session, uint256 seq)
func (_SimpleMultiSessionApp *SimpleMultiSessionAppFilterer) WatchIntendSettle(opts *bind.WatchOpts, sink chan<- *SimpleMultiSessionAppIntendSettle) (event.Subscription, error) {

	logs, sub, err := _SimpleMultiSessionApp.contract.WatchLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleMultiSessionAppIntendSettle)
				if err := _SimpleMultiSessionApp.contract.UnpackLog(event, "IntendSettle", log); err != nil {
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
