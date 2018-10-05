// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gomoku

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

// GomokuABI is the input ABI used to generate the binding from.
const GomokuABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"settle_deadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"settle_timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"isTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"queryResult\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stateproof\",\"type\":\"bytes\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"intendSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"board\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"queryState\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"turn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"game_phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"move_deadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint8\"},{\"name\":\"_y\",\"type\":\"uint8\"}],\"name\":\"placeStone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized_time\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"},{\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSettleTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"move_timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"players\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_player1\",\"type\":\"address\"},{\"name\":\"_player2\",\"type\":\"address\"},{\"name\":\"_move_timeout\",\"type\":\"uint256\"},{\"name\":\"_settle_timeout\",\"type\":\"uint256\"},{\"name\":\"_min_stone_offchain\",\"type\":\"uint8\"},{\"name\":\"_max_stone_onchain\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"IntendSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ConfirmSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"x\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint8\"}],\"name\":\"PlaceStone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"winner\",\"type\":\"uint256\"}],\"name\":\"GameOver\",\"type\":\"event\"}]"

// GomokuBin is the compiled bytecode used for deploying new contracts.
const GomokuBin = `0x608060405234801561001057600080fd5b5060405160c08061211a83398101604090815281516020830151918301516060840151608085015160a09095015160018054600160a060020a03958616600160a060020a03199182161790915560028054909116949095169390931790935560155560179190915560198054600060128190556013805461010061ffff1990911617905564ff000000001990911664010000000060ff958616021765ff000000000019166501000000000094909316939093029190911763ffffffff19169055805460ff191681556120329081906100e890396000f3006080604052600436106101065763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025b3e94811461010b5780632db91aaf1461013257806332c2695714610147578063338304b314610170578063754d7f281461018757806377961e17146101e05780638086a92a14610277578063873e9e19146102a85780638b299903146103325780638bae739814610347578063affed0e014610380578063b989ff2514610395578063c3657ff6146103aa578063c8282aae146103cb578063d2121c2b146103e0578063d71d04961461043b578063dfbf53ae14610450578063f49fec7b14610465578063f71d96cb1461047a575b600080fd5b34801561011757600080fd5b506101206104ae565b60408051918252519081900360200190f35b34801561013e57600080fd5b506101206104b4565b34801561015357600080fd5b5061015c6104ba565b604080519115158252519081900360200190f35b34801561017c57600080fd5b50610185610531565b005b34801561019357600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015c9436949293602493928401919081908401838280828437509497506105969650505050505050565b3480156101ec57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261018594369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506105db9650505050505050565b34801561028357600080fd5b506102926004356024356106eb565b6040805160ff9092168252519081900360200190f35b3480156102b457600080fd5b506102bd610721565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102f75781810151838201526020016102df565b50505050905090810190601f1680156103245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561033e57600080fd5b506102926108ae565b34801561035357600080fd5b5061035c6108bc565b6040518082600381111561036c57fe5b60ff16815260200191505060405180910390f35b34801561038c57600080fd5b506101206108c5565b3480156103a157600080fd5b506101206108cb565b3480156103b657600080fd5b5061018560ff600435811690602435166108d1565b3480156103d757600080fd5b50610120610d23565b3480156103ec57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015c9436949293602493928401919081908401838280828437509497505093359450610d299350505050565b34801561044757600080fd5b50610120610d54565b34801561045c57600080fd5b50610292610d5a565b34801561047157600080fd5b50610120610d63565b34801561048657600080fd5b50610492600435610d69565b60408051600160a060020a039092168252519081900360200190f35b60185481565b60175481565b6000600260005460ff1660038111156104cf57fe5b1480156104dd575060165443115b1561052a5760135460ff6101009091041660011415610505576105006002610d86565b610522565b601354610100900460ff1660021415610522576105226001610d86565b50600161052e565b5060005b90565b600160005460ff16600381111561054457fe5b148015610552575060185443115b15610594576000805460ff1916600217815560155443016016556040517f3fdef4cfa6ef64d0f8ecc3cd33edb3c0874a25c73d65b22a02894b32a8e401ea9190a15b565b80516000906001146105a757600080fd5b8160008151811015156105b657fe5b016020015160135460ff90811660f860020a9283900483029290920416149050919050565b6105e3611f73565b600360005460ff1660038111156105f657fe5b141561060157600080fd5b61060a82610e68565b90506106168382610e8d565b151561066c576040805160e560020a62461bcd02815260206004820152601260248201527f696e76616c6964207369676e6174757265730000000000000000000000000000604482015290519081900360640190fd5b600160005460ff16600381111561067f57fe5b14156106945760185443111561069457600080fd5b61069d8361113a565b6000805460ff1660038111156106af57fe5b14806106cb5750600260005460ff1660038111156106c957fe5b145b156106e6576000805460ff1916600117905560175443016018555b505050565b600382600f81106106f857fe5b0181600f811061070457fe5b602081049091015460ff601f9092166101000a9004169150829050565b6040805160e3808252610120820190925260609182916000918291829160208201611c60803883395050601354825192965060ff1660f860020a0291869150600090811061076b57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350601360019054906101000a900460ff1660f860020a028460018151811015156107c057fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060029250600091505b600f60ff831610156108a5575060005b600f60ff8216101561089a57600360ff8316600f811061082557fe5b0160ff8216600f811061083457fe5b602091828204019190069054906101000a900460ff1660f860020a02848481518110151561085e57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060019283019201610809565b6001909101906107f9565b50919392505050565b601354610100900460ff1681565b60005460ff1681565b60125481565b60165481565b600360005460ff1660038111156108e457fe5b141561093a576040805160e560020a62461bcd02815260206004820152600c60248201527f47616d65206973206f7665720000000000000000000000000000000000000000604482015290519081900360640190fd5b600160005460ff16600381111561094d57fe5b14156109a3576040805160e560020a62461bcd02815260206004820152601860248201527f536574746c696e67206f66662d636861696e2073746174650000000000000000604482015290519081900360640190fd5b60135460019060001960ff61010090920482160116600281106109c257fe5b0154600160a060020a03163314610a23576040805160e560020a62461bcd02815260206004820152600d60248201527f4e6f7420796f7572207475726e00000000000000000000000000000000000000604482015290519081900360640190fd5b610a2d828261149d565b1515610a83576040805160e560020a62461bcd02815260206004820152600f60248201527f4f7574206f6620626f756e646172790000000000000000000000000000000000604482015290519081900360640190fd5b600360ff8316600f8110610a9357fe5b0160ff8216600f8110610aa257fe5b602081049091015460ff601f9092166101000a90041615610b0d576040805160e560020a62461bcd02815260206004820152601960248201527f54686520736c6f7420697320616c72656164792074616b656e00000000000000604482015290519081900360640190fd5b60135460ff6101009091048116906003908416600f8110610b2a57fe5b0160ff8316600f8110610b3957fe5b6020808204929092018054601f9092166101000a60ff8181021990931694831602939093179092556012805460019081019091556019805463ffff00001961ffff19821661ffff928316850183161790811662010000918290048316909401909116029190911790556040805185841681529284169183019190915280517fa2292fc79d2ab39f4f11f31505ad91f6fb4b4aff88fc7cdc696fc25c02e1a36b9281900390910190a1610bf182826001600060016114bc565b80610c065750610c06828260006001806114bc565b80610c1b5750610c1b828260018060016114bc565b80610c325750610c328282600160001960016114bc565b15610c5057601354610c4b90610100900460ff16610d86565b610d1f565b60195461ffff1660e11480610c7c575060195465010000000000810460ff166201000090910461ffff16115b15610cd1576000805460ff191660031781556013805461ff001916905560408051918252517f3496ed15a97ad1d154265bf94e0068b3b45bed65b26cb72ecc4e1b8a44d9b3d7916020908290030190a1610d1f565b60135460ff6101009091041660011415610cf9576013805461ff001916610200179055610d09565b6013805461ff0019166101001790555b60155443016016556000805460ff191660021790555b5050565b60145481565b6000600360005460ff166003811115610d3e57fe5b148015610d4d57508160145411155b9392505050565b60185490565b60135460ff1681565b60155481565b60018160028110610d7657fe5b0154600160a060020a0316905081565b8060ff16600011158015610d9e575060028160ff1611155b1515610df4576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c69642077696e6e6572207374617465000000000000000000000000604482015290519081900360640190fd5b6013805460ff191660ff83811691909117918290551615610e65576000805460ff191660031790556013805461ff00198116909155436014556040805160ff9092168252517f3496ed15a97ad1d154265bf94e0068b3b45bed65b26cb72ecc4e1b8a44d9b3d7916020908290030190a15b50565b610e70611f73565b610e78611f73565b610e8560208485516115bd565b509392505050565b60008060008060008560000151516002141515610ea957600080fd5b866040518082805190602001908083835b60208310610ed95780518252601f199092019160209182019101610eba565b5181516000196020949094036101000a93909301928316921916919091179052604080519390910183900383207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008452601c8401819052905192839003603c019092208a5180519399509097506001945087935091600091508110610f5a57fe5b9060200190602002015188602001516000815181101515610f7757fe5b9060200190602002015189604001516000815181101515610f9457fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af1158015610ff5573d6000803e3d6000fd5b5050506020604051035191506001838760000151600181518110151561101757fe5b906020019060200201518860200151600181518110151561103457fe5b906020019060200201518960400151600181518110151561105157fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af11580156110b2573d6000803e3d6000fd5b5050604051601f19015191506001905060000154600160a060020a0383811691161480156110ed5750600254600160a060020a038281169116145b8061111d5750600254600160a060020a03838116911614801561111d5750600154600160a060020a038281169116145b1561112b5760019450611130565b600094505b5050505092915050565b611142611f95565b60008060006111508561177c565b60125481519195501061116257600080fd5b835160125560208401515160e31461117957600080fd5b6111a58460200151600081518110151561118f57fe5b016020015160f860020a90819004810204610d86565b60135460ff16151561146157602084015180516002945060019081106111c757fe5b01602001516013805460ff60f860020a93849004840293909304831661010090810261ff001990921691909117918290559004166000118015906112175750601354600261010090910460ff1611155b151561126d576040805160e560020a62461bcd02815260206004820152601260248201527f696e76616c6964207475726e2073746174650000000000000000000000000000604482015290519081900360640190fd5b6019805461ffff19169055600091505b600f60ff8316101561143a575060005b600f60ff8216101561142f5760208401518051600185019461ffff169081106112b257fe5b016020015160f860020a90819004810204600360ff8416600f81106112d357fe5b0160ff8316600f81106112e257fe5b602091828204019190066101000a81548160ff021916908360ff16021790555060038260ff16600f8110151561131457fe5b0160ff8216600f811061132357fe5b602081049091015460ff601f9092166101000a90041660001180159061137e57506002600360ff8416600f811061135657fe5b0160ff8316600f811061136557fe5b602081049091015460ff601f9092166101000a90041611155b15156113d4576040805160e560020a62461bcd02815260206004820152601360248201527f696e76616c696420626f61726420737461746500000000000000000000000000604482015290519081900360640190fd5b600360ff8316600f81106113e457fe5b0160ff8216600f81106113f357fe5b602081049091015460ff601f9092166101000a90041615611427576019805461ffff8082166001011661ffff199091161790555b60010161128d565b60019091019061127d565b60195460ff64010000000082041661ffff909116101561145957600080fd5b611461611799565b60125460408051918252517fce68db27527c6058059e8cbd1c6de0528ef1c417fe1c21c545919c7da3466d2a9181900360200190a15050505050565b6000600f60ff8416108015610d4d5750600f60ff831610905092915050565b60008086810b85900386820b859003826114d6838361149d565b15156114e457506001611558565b600360ff8b16600f81106114f457fe5b0160ff8a16600f811061150357fe5b602081049091015460ff601f9092166101000a90048116906003908516600f811061152a57fe5b0160ff8416600f811061153957fe5b602081049091015460ff601f9092166101000a90041614611558575060015b85806115615750805b15611575576115728a8a8a8a611872565b93505b85156115975760016115918b8b8b600019028b60001902611872565b03840193505b600560ff8516106115ab57600194506115b0565b600094505b5050505095945050505050565b6115c5611f73565b60006115cf611f73565b6115d7611fc8565b60008080895b8881018b101561165e576115f18b8b611944565b9c8d019c919550935091508360011415611621576116188b8b611612611981565b88611989565b8b019a50611659565b836002141561163d576116188b8b611637611981565b88611a0f565b8360031415610106576116188b8b611653611981565b88611a78565b6115dd565b9950898460016020020151604051908082528060200260200182016040528015611692578160200160208202803883390190505b5086526040808601518151818152602082810282010190925280156116c1578160200160208202803883390190505b50602087015284600360200201516040519080825280602002602001820160405280156116f8578160200160208202803883390190505b5060408701525b8881018b101561176b576117138b8b611944565b9c8d019c91955093509150836001141561173c576117338b8b8888611989565b8b019a50611766565b8360021415611751576117338b8b8888611a0f565b8360031415610106576117338b8b8888611a78565b6116ff565b509399969850959650505050505050565b611784611f95565b61178c611f95565b610e856020848551611ae1565b600080805b600f60ff841610156106e657600091505b600f60ff8316101561186757600360ff8416600f81106117cb57fe5b0160ff8316600f81106117da57fe5b602081049091015460ff601f9092166101000a9004169050801561185c57611807838360016000806114bc565b8061181d575061181d83836000600160006114bc565b806118325750611832838360018060006114bc565b8061184957506118498383600019600160006114bc565b1561185c5761185781610d86565b6106e6565b6001909101906117af565b60019092019161179e565b6000600181805b600560ff841611611939578260ff1686028860000b0191508260ff1685028760000b0190506118a8828261149d565b801561191c5750600360ff8916600f81106118bf57fe5b0160ff8816600f81106118ce57fe5b602081049091015460ff601f9092166101000a90048116906003908416600f81106118f557fe5b0160ff8316600f811061190457fe5b602081049091015460ff601f9092166101000a900416145b1561192c57600183019250611934565b829350611939565b611879565b505050949350505050565b60008060008060008060006119598989611c54565b9350935083600716600581111561196c57fe5b91506008840499919850919650945050505050565b61052a611f73565b60008060006119988787611c99565b915091506119a585611cb8565b156119bf57600184815b6020020180519091019052611a05565b84516020850151815184929181039081106119d657fe5b60ff9290921660209283029190910182015284015160001015611a0557600184815b6020020180519190910390525b9695505050505050565b6000806000611a1e8787611cbc565b91509150611a2b85611cb8565b15611a3a5760018460026119af565b6020850151604085015181518492918103908110611a5457fe5b60209081029091010152600084600260200201511115611a055760018460026119f8565b6000806000611a878787611cbc565b91509150611a9485611cb8565b15611aa35760018460036119af565b6040850151606085015181518492918103908110611abd57fe5b60209081029091010152600084600360200201511115611a055760018460036119f8565b611ae9611f95565b6000611af3611f95565b611afb611fe7565b60008080895b8881018b1015611b9757611b158b8b611944565b9c8d019c919550935091508360011415611b3e57611b358b8b8888611cd6565b8b019a50611b92565b8360021415611b5357611b358b8b8888611d19565b8360031415611b6857611b358b8b8888611d63565b8360041415611b7d57611b358b8b8888611dac565b836005141561010657611b358b8b8888611df5565b611b01565b809a505b8881018b101561176b57611baf8b8b611944565b9c8d019c919550935091508360011415611bdf57611bd68b8b611bd0611e3e565b88611cd6565b8b019a50611c4f565b8360021415611bfb57611bd68b8b611bf5611e3e565b88611d19565b8360031415611c1757611bd68b8b611c11611e3e565b88611d63565b8360041415611c3357611bd68b8b611c2d611e3e565b88611dac565b836005141561010657611bd68b8b611c49611e3e565b88611df5565b611b9b565b908101906000808080805b865160001a90508160070260020a81607f16028317925060018201915060018701965060808116608014611c5f5750909590945092505050565b600080600080611cab60018787611e46565b9097909650945050505050565b1590565b600080611ccb60208585611e90565b915091509250929050565b6000806000611ce58787611ed5565b91509150611cf285611cb8565b15611d0057600184816119af565b818552602084015160001015611a0557600184816119f8565b600060606000611d298787611ee7565b91509150611d3685611cb8565b15611d455760018460026119af565b60208501829052604084015160001015611a055760018460026119f8565b6000806000611d728787611cbc565b91509150611d7f85611cb8565b15611d8e5760018460036119af565b60408501829052606084015160001015611a055760018460036119f8565b6000806000611dbb8787611ed5565b91509150611dc885611cb8565b15611dd75760018460046119af565b60608501829052608084015160001015611a055760018460046119f8565b6000806000611e048787611ed5565b91509150611e1185611cb8565b15611e205760018460056119af565b6080850182905260a084015160001015611a055760018460056119f8565b61052a611f95565b6000806000806000611e588787611c54565b91509150600387019650858701519250600282036020036101000a83811515611e7d57fe5b0498600190920197509095505050505050565b6000806000806000611ea28787611c54565b909350915082820160ff60038a011614611ebb57600080fd5b505050600393909201830151949390920160ff1692915050565b600080600080611cab60208787611e46565b60606000611ccb8484606060008060006060611f038787611c54565b92509250826040519080825280601f01601f191660200182016040528015611f35578160200160208202803883390190505b5096860182019690506020810160005b848114611f6357885160001a82536001988901989182019101611f45565b5090979190920195509350505050565b6060604051908101604052806060815260200160608152602001606081525090565b60a06040519081016040528060008152602001606081526020016000801916815260200160008152602001600081525090565b6080604051908101604052806004906020820280388339509192915050565b60c06040519081016040528060069060208202803883395091929150505600a165627a7a723058205c85d3051ed2afff86802cff8ea7335ecffcc400399559ef7e61d8035b8a4a600029`

// DeployGomoku deploys a new Ethereum contract, binding an instance of Gomoku to it.
func DeployGomoku(auth *bind.TransactOpts, backend bind.ContractBackend, _player1 common.Address, _player2 common.Address, _move_timeout *big.Int, _settle_timeout *big.Int, _min_stone_offchain uint8, _max_stone_onchain uint8) (common.Address, *types.Transaction, *Gomoku, error) {
	parsed, err := abi.JSON(strings.NewReader(GomokuABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GomokuBin), backend, _player1, _player2, _move_timeout, _settle_timeout, _min_stone_offchain, _max_stone_onchain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gomoku{GomokuCaller: GomokuCaller{contract: contract}, GomokuTransactor: GomokuTransactor{contract: contract}, GomokuFilterer: GomokuFilterer{contract: contract}}, nil
}

// Gomoku is an auto generated Go binding around an Ethereum contract.
type Gomoku struct {
	GomokuCaller     // Read-only binding to the contract
	GomokuTransactor // Write-only binding to the contract
	GomokuFilterer   // Log filterer for contract events
}

// GomokuCaller is an auto generated read-only Go binding around an Ethereum contract.
type GomokuCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GomokuTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GomokuTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GomokuFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GomokuFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GomokuSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GomokuSession struct {
	Contract     *Gomoku           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GomokuCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GomokuCallerSession struct {
	Contract *GomokuCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GomokuTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GomokuTransactorSession struct {
	Contract     *GomokuTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GomokuRaw is an auto generated low-level Go binding around an Ethereum contract.
type GomokuRaw struct {
	Contract *Gomoku // Generic contract binding to access the raw methods on
}

// GomokuCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GomokuCallerRaw struct {
	Contract *GomokuCaller // Generic read-only contract binding to access the raw methods on
}

// GomokuTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GomokuTransactorRaw struct {
	Contract *GomokuTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGomoku creates a new instance of Gomoku, bound to a specific deployed contract.
func NewGomoku(address common.Address, backend bind.ContractBackend) (*Gomoku, error) {
	contract, err := bindGomoku(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gomoku{GomokuCaller: GomokuCaller{contract: contract}, GomokuTransactor: GomokuTransactor{contract: contract}, GomokuFilterer: GomokuFilterer{contract: contract}}, nil
}

// NewGomokuCaller creates a new read-only instance of Gomoku, bound to a specific deployed contract.
func NewGomokuCaller(address common.Address, caller bind.ContractCaller) (*GomokuCaller, error) {
	contract, err := bindGomoku(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GomokuCaller{contract: contract}, nil
}

// NewGomokuTransactor creates a new write-only instance of Gomoku, bound to a specific deployed contract.
func NewGomokuTransactor(address common.Address, transactor bind.ContractTransactor) (*GomokuTransactor, error) {
	contract, err := bindGomoku(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GomokuTransactor{contract: contract}, nil
}

// NewGomokuFilterer creates a new log filterer instance of Gomoku, bound to a specific deployed contract.
func NewGomokuFilterer(address common.Address, filterer bind.ContractFilterer) (*GomokuFilterer, error) {
	contract, err := bindGomoku(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GomokuFilterer{contract: contract}, nil
}

// bindGomoku binds a generic wrapper to an already deployed contract.
func bindGomoku(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GomokuABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gomoku *GomokuRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gomoku.Contract.GomokuCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gomoku *GomokuRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gomoku.Contract.GomokuTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gomoku *GomokuRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gomoku.Contract.GomokuTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gomoku *GomokuCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gomoku.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gomoku *GomokuTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gomoku.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gomoku *GomokuTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gomoku.Contract.contract.Transact(opts, method, params...)
}

// Board is a free data retrieval call binding the contract method 0x8086a92a.
//
// Solidity: function board( uint256,  uint256) constant returns(uint8)
func (_Gomoku *GomokuCaller) Board(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "board", arg0, arg1)
	return *ret0, err
}

// Board is a free data retrieval call binding the contract method 0x8086a92a.
//
// Solidity: function board( uint256,  uint256) constant returns(uint8)
func (_Gomoku *GomokuSession) Board(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Gomoku.Contract.Board(&_Gomoku.CallOpts, arg0, arg1)
}

// Board is a free data retrieval call binding the contract method 0x8086a92a.
//
// Solidity: function board( uint256,  uint256) constant returns(uint8)
func (_Gomoku *GomokuCallerSession) Board(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Gomoku.Contract.Board(&_Gomoku.CallOpts, arg0, arg1)
}

// FinalizedTime is a free data retrieval call binding the contract method 0xc8282aae.
//
// Solidity: function finalized_time() constant returns(uint256)
func (_Gomoku *GomokuCaller) FinalizedTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "finalized_time")
	return *ret0, err
}

// FinalizedTime is a free data retrieval call binding the contract method 0xc8282aae.
//
// Solidity: function finalized_time() constant returns(uint256)
func (_Gomoku *GomokuSession) FinalizedTime() (*big.Int, error) {
	return _Gomoku.Contract.FinalizedTime(&_Gomoku.CallOpts)
}

// FinalizedTime is a free data retrieval call binding the contract method 0xc8282aae.
//
// Solidity: function finalized_time() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) FinalizedTime() (*big.Int, error) {
	return _Gomoku.Contract.FinalizedTime(&_Gomoku.CallOpts)
}

// GamePhase is a free data retrieval call binding the contract method 0x8bae7398.
//
// Solidity: function game_phase() constant returns(uint8)
func (_Gomoku *GomokuCaller) GamePhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "game_phase")
	return *ret0, err
}

// GamePhase is a free data retrieval call binding the contract method 0x8bae7398.
//
// Solidity: function game_phase() constant returns(uint8)
func (_Gomoku *GomokuSession) GamePhase() (uint8, error) {
	return _Gomoku.Contract.GamePhase(&_Gomoku.CallOpts)
}

// GamePhase is a free data retrieval call binding the contract method 0x8bae7398.
//
// Solidity: function game_phase() constant returns(uint8)
func (_Gomoku *GomokuCallerSession) GamePhase() (uint8, error) {
	return _Gomoku.Contract.GamePhase(&_Gomoku.CallOpts)
}

// GetSettleTime is a free data retrieval call binding the contract method 0xd71d0496.
//
// Solidity: function getSettleTime() constant returns(uint256)
func (_Gomoku *GomokuCaller) GetSettleTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "getSettleTime")
	return *ret0, err
}

// GetSettleTime is a free data retrieval call binding the contract method 0xd71d0496.
//
// Solidity: function getSettleTime() constant returns(uint256)
func (_Gomoku *GomokuSession) GetSettleTime() (*big.Int, error) {
	return _Gomoku.Contract.GetSettleTime(&_Gomoku.CallOpts)
}

// GetSettleTime is a free data retrieval call binding the contract method 0xd71d0496.
//
// Solidity: function getSettleTime() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) GetSettleTime() (*big.Int, error) {
	return _Gomoku.Contract.GetSettleTime(&_Gomoku.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(_query bytes, _timeout uint256) constant returns(bool)
func (_Gomoku *GomokuCaller) IsFinalized(opts *bind.CallOpts, _query []byte, _timeout *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "isFinalized", _query, _timeout)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(_query bytes, _timeout uint256) constant returns(bool)
func (_Gomoku *GomokuSession) IsFinalized(_query []byte, _timeout *big.Int) (bool, error) {
	return _Gomoku.Contract.IsFinalized(&_Gomoku.CallOpts, _query, _timeout)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(_query bytes, _timeout uint256) constant returns(bool)
func (_Gomoku *GomokuCallerSession) IsFinalized(_query []byte, _timeout *big.Int) (bool, error) {
	return _Gomoku.Contract.IsFinalized(&_Gomoku.CallOpts, _query, _timeout)
}

// MoveDeadline is a free data retrieval call binding the contract method 0xb989ff25.
//
// Solidity: function move_deadline() constant returns(uint256)
func (_Gomoku *GomokuCaller) MoveDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "move_deadline")
	return *ret0, err
}

// MoveDeadline is a free data retrieval call binding the contract method 0xb989ff25.
//
// Solidity: function move_deadline() constant returns(uint256)
func (_Gomoku *GomokuSession) MoveDeadline() (*big.Int, error) {
	return _Gomoku.Contract.MoveDeadline(&_Gomoku.CallOpts)
}

// MoveDeadline is a free data retrieval call binding the contract method 0xb989ff25.
//
// Solidity: function move_deadline() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) MoveDeadline() (*big.Int, error) {
	return _Gomoku.Contract.MoveDeadline(&_Gomoku.CallOpts)
}

// MoveTimeout is a free data retrieval call binding the contract method 0xf49fec7b.
//
// Solidity: function move_timeout() constant returns(uint256)
func (_Gomoku *GomokuCaller) MoveTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "move_timeout")
	return *ret0, err
}

// MoveTimeout is a free data retrieval call binding the contract method 0xf49fec7b.
//
// Solidity: function move_timeout() constant returns(uint256)
func (_Gomoku *GomokuSession) MoveTimeout() (*big.Int, error) {
	return _Gomoku.Contract.MoveTimeout(&_Gomoku.CallOpts)
}

// MoveTimeout is a free data retrieval call binding the contract method 0xf49fec7b.
//
// Solidity: function move_timeout() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) MoveTimeout() (*big.Int, error) {
	return _Gomoku.Contract.MoveTimeout(&_Gomoku.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Gomoku *GomokuCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Gomoku *GomokuSession) Nonce() (*big.Int, error) {
	return _Gomoku.Contract.Nonce(&_Gomoku.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) Nonce() (*big.Int, error) {
	return _Gomoku.Contract.Nonce(&_Gomoku.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players( uint256) constant returns(address)
func (_Gomoku *GomokuCaller) Players(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "players", arg0)
	return *ret0, err
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players( uint256) constant returns(address)
func (_Gomoku *GomokuSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Gomoku.Contract.Players(&_Gomoku.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players( uint256) constant returns(address)
func (_Gomoku *GomokuCallerSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Gomoku.Contract.Players(&_Gomoku.CallOpts, arg0)
}

// QueryResult is a free data retrieval call binding the contract method 0x754d7f28.
//
// Solidity: function queryResult(_query bytes) constant returns(bool)
func (_Gomoku *GomokuCaller) QueryResult(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "queryResult", _query)
	return *ret0, err
}

// QueryResult is a free data retrieval call binding the contract method 0x754d7f28.
//
// Solidity: function queryResult(_query bytes) constant returns(bool)
func (_Gomoku *GomokuSession) QueryResult(_query []byte) (bool, error) {
	return _Gomoku.Contract.QueryResult(&_Gomoku.CallOpts, _query)
}

// QueryResult is a free data retrieval call binding the contract method 0x754d7f28.
//
// Solidity: function queryResult(_query bytes) constant returns(bool)
func (_Gomoku *GomokuCallerSession) QueryResult(_query []byte) (bool, error) {
	return _Gomoku.Contract.QueryResult(&_Gomoku.CallOpts, _query)
}

// QueryState is a free data retrieval call binding the contract method 0x873e9e19.
//
// Solidity: function queryState() constant returns(bytes)
func (_Gomoku *GomokuCaller) QueryState(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "queryState")
	return *ret0, err
}

// QueryState is a free data retrieval call binding the contract method 0x873e9e19.
//
// Solidity: function queryState() constant returns(bytes)
func (_Gomoku *GomokuSession) QueryState() ([]byte, error) {
	return _Gomoku.Contract.QueryState(&_Gomoku.CallOpts)
}

// QueryState is a free data retrieval call binding the contract method 0x873e9e19.
//
// Solidity: function queryState() constant returns(bytes)
func (_Gomoku *GomokuCallerSession) QueryState() ([]byte, error) {
	return _Gomoku.Contract.QueryState(&_Gomoku.CallOpts)
}

// SettleDeadline is a free data retrieval call binding the contract method 0x025b3e94.
//
// Solidity: function settle_deadline() constant returns(uint256)
func (_Gomoku *GomokuCaller) SettleDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "settle_deadline")
	return *ret0, err
}

// SettleDeadline is a free data retrieval call binding the contract method 0x025b3e94.
//
// Solidity: function settle_deadline() constant returns(uint256)
func (_Gomoku *GomokuSession) SettleDeadline() (*big.Int, error) {
	return _Gomoku.Contract.SettleDeadline(&_Gomoku.CallOpts)
}

// SettleDeadline is a free data retrieval call binding the contract method 0x025b3e94.
//
// Solidity: function settle_deadline() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) SettleDeadline() (*big.Int, error) {
	return _Gomoku.Contract.SettleDeadline(&_Gomoku.CallOpts)
}

// SettleTimeout is a free data retrieval call binding the contract method 0x2db91aaf.
//
// Solidity: function settle_timeout() constant returns(uint256)
func (_Gomoku *GomokuCaller) SettleTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "settle_timeout")
	return *ret0, err
}

// SettleTimeout is a free data retrieval call binding the contract method 0x2db91aaf.
//
// Solidity: function settle_timeout() constant returns(uint256)
func (_Gomoku *GomokuSession) SettleTimeout() (*big.Int, error) {
	return _Gomoku.Contract.SettleTimeout(&_Gomoku.CallOpts)
}

// SettleTimeout is a free data retrieval call binding the contract method 0x2db91aaf.
//
// Solidity: function settle_timeout() constant returns(uint256)
func (_Gomoku *GomokuCallerSession) SettleTimeout() (*big.Int, error) {
	return _Gomoku.Contract.SettleTimeout(&_Gomoku.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() constant returns(uint8)
func (_Gomoku *GomokuCaller) Turn(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "turn")
	return *ret0, err
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() constant returns(uint8)
func (_Gomoku *GomokuSession) Turn() (uint8, error) {
	return _Gomoku.Contract.Turn(&_Gomoku.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() constant returns(uint8)
func (_Gomoku *GomokuCallerSession) Turn() (uint8, error) {
	return _Gomoku.Contract.Turn(&_Gomoku.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() constant returns(uint8)
func (_Gomoku *GomokuCaller) Winner(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gomoku.contract.Call(opts, out, "winner")
	return *ret0, err
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() constant returns(uint8)
func (_Gomoku *GomokuSession) Winner() (uint8, error) {
	return _Gomoku.Contract.Winner(&_Gomoku.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() constant returns(uint8)
func (_Gomoku *GomokuCallerSession) Winner() (uint8, error) {
	return _Gomoku.Contract.Winner(&_Gomoku.CallOpts)
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x338304b3.
//
// Solidity: function confirmSettle() returns()
func (_Gomoku *GomokuTransactor) ConfirmSettle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gomoku.contract.Transact(opts, "confirmSettle")
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x338304b3.
//
// Solidity: function confirmSettle() returns()
func (_Gomoku *GomokuSession) ConfirmSettle() (*types.Transaction, error) {
	return _Gomoku.Contract.ConfirmSettle(&_Gomoku.TransactOpts)
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x338304b3.
//
// Solidity: function confirmSettle() returns()
func (_Gomoku *GomokuTransactorSession) ConfirmSettle() (*types.Transaction, error) {
	return _Gomoku.Contract.ConfirmSettle(&_Gomoku.TransactOpts)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x77961e17.
//
// Solidity: function intendSettle(_stateproof bytes, _signatures bytes) returns()
func (_Gomoku *GomokuTransactor) IntendSettle(opts *bind.TransactOpts, _stateproof []byte, _signatures []byte) (*types.Transaction, error) {
	return _Gomoku.contract.Transact(opts, "intendSettle", _stateproof, _signatures)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x77961e17.
//
// Solidity: function intendSettle(_stateproof bytes, _signatures bytes) returns()
func (_Gomoku *GomokuSession) IntendSettle(_stateproof []byte, _signatures []byte) (*types.Transaction, error) {
	return _Gomoku.Contract.IntendSettle(&_Gomoku.TransactOpts, _stateproof, _signatures)
}

// IntendSettle is a paid mutator transaction binding the contract method 0x77961e17.
//
// Solidity: function intendSettle(_stateproof bytes, _signatures bytes) returns()
func (_Gomoku *GomokuTransactorSession) IntendSettle(_stateproof []byte, _signatures []byte) (*types.Transaction, error) {
	return _Gomoku.Contract.IntendSettle(&_Gomoku.TransactOpts, _stateproof, _signatures)
}

// IsTimeout is a paid mutator transaction binding the contract method 0x32c26957.
//
// Solidity: function isTimeout() returns(bool)
func (_Gomoku *GomokuTransactor) IsTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gomoku.contract.Transact(opts, "isTimeout")
}

// IsTimeout is a paid mutator transaction binding the contract method 0x32c26957.
//
// Solidity: function isTimeout() returns(bool)
func (_Gomoku *GomokuSession) IsTimeout() (*types.Transaction, error) {
	return _Gomoku.Contract.IsTimeout(&_Gomoku.TransactOpts)
}

// IsTimeout is a paid mutator transaction binding the contract method 0x32c26957.
//
// Solidity: function isTimeout() returns(bool)
func (_Gomoku *GomokuTransactorSession) IsTimeout() (*types.Transaction, error) {
	return _Gomoku.Contract.IsTimeout(&_Gomoku.TransactOpts)
}

// PlaceStone is a paid mutator transaction binding the contract method 0xc3657ff6.
//
// Solidity: function placeStone(_x uint8, _y uint8) returns()
func (_Gomoku *GomokuTransactor) PlaceStone(opts *bind.TransactOpts, _x uint8, _y uint8) (*types.Transaction, error) {
	return _Gomoku.contract.Transact(opts, "placeStone", _x, _y)
}

// PlaceStone is a paid mutator transaction binding the contract method 0xc3657ff6.
//
// Solidity: function placeStone(_x uint8, _y uint8) returns()
func (_Gomoku *GomokuSession) PlaceStone(_x uint8, _y uint8) (*types.Transaction, error) {
	return _Gomoku.Contract.PlaceStone(&_Gomoku.TransactOpts, _x, _y)
}

// PlaceStone is a paid mutator transaction binding the contract method 0xc3657ff6.
//
// Solidity: function placeStone(_x uint8, _y uint8) returns()
func (_Gomoku *GomokuTransactorSession) PlaceStone(_x uint8, _y uint8) (*types.Transaction, error) {
	return _Gomoku.Contract.PlaceStone(&_Gomoku.TransactOpts, _x, _y)
}

// GomokuConfirmSettleIterator is returned from FilterConfirmSettle and is used to iterate over the raw logs and unpacked data for ConfirmSettle events raised by the Gomoku contract.
type GomokuConfirmSettleIterator struct {
	Event *GomokuConfirmSettle // Event containing the contract specifics and raw log

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
func (it *GomokuConfirmSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GomokuConfirmSettle)
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
		it.Event = new(GomokuConfirmSettle)
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
func (it *GomokuConfirmSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GomokuConfirmSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GomokuConfirmSettle represents a ConfirmSettle event raised by the Gomoku contract.
type GomokuConfirmSettle struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterConfirmSettle is a free log retrieval operation binding the contract event 0x3fdef4cfa6ef64d0f8ecc3cd33edb3c0874a25c73d65b22a02894b32a8e401ea.
//
// Solidity: e ConfirmSettle()
func (_Gomoku *GomokuFilterer) FilterConfirmSettle(opts *bind.FilterOpts) (*GomokuConfirmSettleIterator, error) {

	logs, sub, err := _Gomoku.contract.FilterLogs(opts, "ConfirmSettle")
	if err != nil {
		return nil, err
	}
	return &GomokuConfirmSettleIterator{contract: _Gomoku.contract, event: "ConfirmSettle", logs: logs, sub: sub}, nil
}

// WatchConfirmSettle is a free log subscription operation binding the contract event 0x3fdef4cfa6ef64d0f8ecc3cd33edb3c0874a25c73d65b22a02894b32a8e401ea.
//
// Solidity: e ConfirmSettle()
func (_Gomoku *GomokuFilterer) WatchConfirmSettle(opts *bind.WatchOpts, sink chan<- *GomokuConfirmSettle) (event.Subscription, error) {

	logs, sub, err := _Gomoku.contract.WatchLogs(opts, "ConfirmSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GomokuConfirmSettle)
				if err := _Gomoku.contract.UnpackLog(event, "ConfirmSettle", log); err != nil {
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

// GomokuGameOverIterator is returned from FilterGameOver and is used to iterate over the raw logs and unpacked data for GameOver events raised by the Gomoku contract.
type GomokuGameOverIterator struct {
	Event *GomokuGameOver // Event containing the contract specifics and raw log

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
func (it *GomokuGameOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GomokuGameOver)
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
		it.Event = new(GomokuGameOver)
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
func (it *GomokuGameOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GomokuGameOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GomokuGameOver represents a GameOver event raised by the Gomoku contract.
type GomokuGameOver struct {
	Winner *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameOver is a free log retrieval operation binding the contract event 0x3496ed15a97ad1d154265bf94e0068b3b45bed65b26cb72ecc4e1b8a44d9b3d7.
//
// Solidity: e GameOver(winner uint256)
func (_Gomoku *GomokuFilterer) FilterGameOver(opts *bind.FilterOpts) (*GomokuGameOverIterator, error) {

	logs, sub, err := _Gomoku.contract.FilterLogs(opts, "GameOver")
	if err != nil {
		return nil, err
	}
	return &GomokuGameOverIterator{contract: _Gomoku.contract, event: "GameOver", logs: logs, sub: sub}, nil
}

// WatchGameOver is a free log subscription operation binding the contract event 0x3496ed15a97ad1d154265bf94e0068b3b45bed65b26cb72ecc4e1b8a44d9b3d7.
//
// Solidity: e GameOver(winner uint256)
func (_Gomoku *GomokuFilterer) WatchGameOver(opts *bind.WatchOpts, sink chan<- *GomokuGameOver) (event.Subscription, error) {

	logs, sub, err := _Gomoku.contract.WatchLogs(opts, "GameOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GomokuGameOver)
				if err := _Gomoku.contract.UnpackLog(event, "GameOver", log); err != nil {
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

// GomokuIntendSettleIterator is returned from FilterIntendSettle and is used to iterate over the raw logs and unpacked data for IntendSettle events raised by the Gomoku contract.
type GomokuIntendSettleIterator struct {
	Event *GomokuIntendSettle // Event containing the contract specifics and raw log

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
func (it *GomokuIntendSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GomokuIntendSettle)
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
		it.Event = new(GomokuIntendSettle)
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
func (it *GomokuIntendSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GomokuIntendSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GomokuIntendSettle represents a IntendSettle event raised by the Gomoku contract.
type GomokuIntendSettle struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIntendSettle is a free log retrieval operation binding the contract event 0xce68db27527c6058059e8cbd1c6de0528ef1c417fe1c21c545919c7da3466d2a.
//
// Solidity: e IntendSettle(nonce uint256)
func (_Gomoku *GomokuFilterer) FilterIntendSettle(opts *bind.FilterOpts) (*GomokuIntendSettleIterator, error) {

	logs, sub, err := _Gomoku.contract.FilterLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return &GomokuIntendSettleIterator{contract: _Gomoku.contract, event: "IntendSettle", logs: logs, sub: sub}, nil
}

// WatchIntendSettle is a free log subscription operation binding the contract event 0xce68db27527c6058059e8cbd1c6de0528ef1c417fe1c21c545919c7da3466d2a.
//
// Solidity: e IntendSettle(nonce uint256)
func (_Gomoku *GomokuFilterer) WatchIntendSettle(opts *bind.WatchOpts, sink chan<- *GomokuIntendSettle) (event.Subscription, error) {

	logs, sub, err := _Gomoku.contract.WatchLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GomokuIntendSettle)
				if err := _Gomoku.contract.UnpackLog(event, "IntendSettle", log); err != nil {
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

// GomokuPlaceStoneIterator is returned from FilterPlaceStone and is used to iterate over the raw logs and unpacked data for PlaceStone events raised by the Gomoku contract.
type GomokuPlaceStoneIterator struct {
	Event *GomokuPlaceStone // Event containing the contract specifics and raw log

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
func (it *GomokuPlaceStoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GomokuPlaceStone)
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
		it.Event = new(GomokuPlaceStone)
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
func (it *GomokuPlaceStoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GomokuPlaceStoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GomokuPlaceStone represents a PlaceStone event raised by the Gomoku contract.
type GomokuPlaceStone struct {
	X   uint8
	Y   uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPlaceStone is a free log retrieval operation binding the contract event 0xa2292fc79d2ab39f4f11f31505ad91f6fb4b4aff88fc7cdc696fc25c02e1a36b.
//
// Solidity: e PlaceStone(x uint8, y uint8)
func (_Gomoku *GomokuFilterer) FilterPlaceStone(opts *bind.FilterOpts) (*GomokuPlaceStoneIterator, error) {

	logs, sub, err := _Gomoku.contract.FilterLogs(opts, "PlaceStone")
	if err != nil {
		return nil, err
	}
	return &GomokuPlaceStoneIterator{contract: _Gomoku.contract, event: "PlaceStone", logs: logs, sub: sub}, nil
}

// WatchPlaceStone is a free log subscription operation binding the contract event 0xa2292fc79d2ab39f4f11f31505ad91f6fb4b4aff88fc7cdc696fc25c02e1a36b.
//
// Solidity: e PlaceStone(x uint8, y uint8)
func (_Gomoku *GomokuFilterer) WatchPlaceStone(opts *bind.WatchOpts, sink chan<- *GomokuPlaceStone) (event.Subscription, error) {

	logs, sub, err := _Gomoku.contract.WatchLogs(opts, "PlaceStone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GomokuPlaceStone)
				if err := _Gomoku.contract.UnpackLog(event, "PlaceStone", log); err != nil {
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

// PbABI is the input ABI used to generate the binding from.
const PbABI = "[]"

// PbBin is the compiled bytecode used for deploying new contracts.
const PbBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c4964bc817a33e4136a9ba06b0231f9b8a70ae5fd845060ab76647fdb444da8a0029`

// DeployPb deploys a new Ethereum contract, binding an instance of Pb to it.
func DeployPb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pb, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// Pb is an auto generated Go binding around an Ethereum contract.
type Pb struct {
	PbCaller     // Read-only binding to the contract
	PbTransactor // Write-only binding to the contract
	PbFilterer   // Log filterer for contract events
}

// PbCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSession struct {
	Contract     *Pb               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbCallerSession struct {
	Contract *PbCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbTransactorSession struct {
	Contract     *PbTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRaw struct {
	Contract *Pb // Generic contract binding to access the raw methods on
}

// PbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbCallerRaw struct {
	Contract *PbCaller // Generic read-only contract binding to access the raw methods on
}

// PbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbTransactorRaw struct {
	Contract *PbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPb creates a new instance of Pb, bound to a specific deployed contract.
func NewPb(address common.Address, backend bind.ContractBackend) (*Pb, error) {
	contract, err := bindPb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// NewPbCaller creates a new read-only instance of Pb, bound to a specific deployed contract.
func NewPbCaller(address common.Address, caller bind.ContractCaller) (*PbCaller, error) {
	contract, err := bindPb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbCaller{contract: contract}, nil
}

// NewPbTransactor creates a new write-only instance of Pb, bound to a specific deployed contract.
func NewPbTransactor(address common.Address, transactor bind.ContractTransactor) (*PbTransactor, error) {
	contract, err := bindPb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbTransactor{contract: contract}, nil
}

// NewPbFilterer creates a new log filterer instance of Pb, bound to a specific deployed contract.
func NewPbFilterer(address common.Address, filterer bind.ContractFilterer) (*PbFilterer, error) {
	contract, err := bindPb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbFilterer{contract: contract}, nil
}

// bindPb binds a generic wrapper to an already deployed contract.
func bindPb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.PbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transact(opts, method, params...)
}

// PbRpcMultiSignatureABI is the input ABI used to generate the binding from.
const PbRpcMultiSignatureABI = "[]"

// PbRpcMultiSignatureBin is the compiled bytecode used for deploying new contracts.
const PbRpcMultiSignatureBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204cf83b034e122108b7e2a212e70f8c1e8ee4f92702416093210919c5b9a057b70029`

// DeployPbRpcMultiSignature deploys a new Ethereum contract, binding an instance of PbRpcMultiSignature to it.
func DeployPbRpcMultiSignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcMultiSignature, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcMultiSignatureABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcMultiSignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcMultiSignature{PbRpcMultiSignatureCaller: PbRpcMultiSignatureCaller{contract: contract}, PbRpcMultiSignatureTransactor: PbRpcMultiSignatureTransactor{contract: contract}, PbRpcMultiSignatureFilterer: PbRpcMultiSignatureFilterer{contract: contract}}, nil
}

// PbRpcMultiSignature is an auto generated Go binding around an Ethereum contract.
type PbRpcMultiSignature struct {
	PbRpcMultiSignatureCaller     // Read-only binding to the contract
	PbRpcMultiSignatureTransactor // Write-only binding to the contract
	PbRpcMultiSignatureFilterer   // Log filterer for contract events
}

// PbRpcMultiSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcMultiSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcMultiSignatureSession struct {
	Contract     *PbRpcMultiSignature // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbRpcMultiSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcMultiSignatureCallerSession struct {
	Contract *PbRpcMultiSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PbRpcMultiSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcMultiSignatureTransactorSession struct {
	Contract     *PbRpcMultiSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PbRpcMultiSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcMultiSignatureRaw struct {
	Contract *PbRpcMultiSignature // Generic contract binding to access the raw methods on
}

// PbRpcMultiSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureCallerRaw struct {
	Contract *PbRpcMultiSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcMultiSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureTransactorRaw struct {
	Contract *PbRpcMultiSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcMultiSignature creates a new instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignature(address common.Address, backend bind.ContractBackend) (*PbRpcMultiSignature, error) {
	contract, err := bindPbRpcMultiSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignature{PbRpcMultiSignatureCaller: PbRpcMultiSignatureCaller{contract: contract}, PbRpcMultiSignatureTransactor: PbRpcMultiSignatureTransactor{contract: contract}, PbRpcMultiSignatureFilterer: PbRpcMultiSignatureFilterer{contract: contract}}, nil
}

// NewPbRpcMultiSignatureCaller creates a new read-only instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureCaller(address common.Address, caller bind.ContractCaller) (*PbRpcMultiSignatureCaller, error) {
	contract, err := bindPbRpcMultiSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureCaller{contract: contract}, nil
}

// NewPbRpcMultiSignatureTransactor creates a new write-only instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcMultiSignatureTransactor, error) {
	contract, err := bindPbRpcMultiSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureTransactor{contract: contract}, nil
}

// NewPbRpcMultiSignatureFilterer creates a new log filterer instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcMultiSignatureFilterer, error) {
	contract, err := bindPbRpcMultiSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureFilterer{contract: contract}, nil
}

// bindPbRpcMultiSignature binds a generic wrapper to an already deployed contract.
func bindPbRpcMultiSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcMultiSignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcMultiSignature *PbRpcMultiSignatureCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcMultiSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcMultiSignature *PbRpcMultiSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcMultiSignature *PbRpcMultiSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.contract.Transact(opts, method, params...)
}

// PbRpcStateProofABI is the input ABI used to generate the binding from.
const PbRpcStateProofABI = "[]"

// PbRpcStateProofBin is the compiled bytecode used for deploying new contracts.
const PbRpcStateProofBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582013541641ed169ab801166ea9d461894b54cfd61cf0310db596ab63244e4425a40029`

// DeployPbRpcStateProof deploys a new Ethereum contract, binding an instance of PbRpcStateProof to it.
func DeployPbRpcStateProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcStateProof, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcStateProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcStateProof{PbRpcStateProofCaller: PbRpcStateProofCaller{contract: contract}, PbRpcStateProofTransactor: PbRpcStateProofTransactor{contract: contract}, PbRpcStateProofFilterer: PbRpcStateProofFilterer{contract: contract}}, nil
}

// PbRpcStateProof is an auto generated Go binding around an Ethereum contract.
type PbRpcStateProof struct {
	PbRpcStateProofCaller     // Read-only binding to the contract
	PbRpcStateProofTransactor // Write-only binding to the contract
	PbRpcStateProofFilterer   // Log filterer for contract events
}

// PbRpcStateProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcStateProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcStateProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcStateProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcStateProofSession struct {
	Contract     *PbRpcStateProof  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRpcStateProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcStateProofCallerSession struct {
	Contract *PbRpcStateProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PbRpcStateProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcStateProofTransactorSession struct {
	Contract     *PbRpcStateProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PbRpcStateProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcStateProofRaw struct {
	Contract *PbRpcStateProof // Generic contract binding to access the raw methods on
}

// PbRpcStateProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcStateProofCallerRaw struct {
	Contract *PbRpcStateProofCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcStateProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcStateProofTransactorRaw struct {
	Contract *PbRpcStateProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcStateProof creates a new instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProof(address common.Address, backend bind.ContractBackend) (*PbRpcStateProof, error) {
	contract, err := bindPbRpcStateProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProof{PbRpcStateProofCaller: PbRpcStateProofCaller{contract: contract}, PbRpcStateProofTransactor: PbRpcStateProofTransactor{contract: contract}, PbRpcStateProofFilterer: PbRpcStateProofFilterer{contract: contract}}, nil
}

// NewPbRpcStateProofCaller creates a new read-only instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofCaller(address common.Address, caller bind.ContractCaller) (*PbRpcStateProofCaller, error) {
	contract, err := bindPbRpcStateProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofCaller{contract: contract}, nil
}

// NewPbRpcStateProofTransactor creates a new write-only instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcStateProofTransactor, error) {
	contract, err := bindPbRpcStateProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofTransactor{contract: contract}, nil
}

// NewPbRpcStateProofFilterer creates a new log filterer instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcStateProofFilterer, error) {
	contract, err := bindPbRpcStateProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofFilterer{contract: contract}, nil
}

// bindPbRpcStateProof binds a generic wrapper to an already deployed contract.
func bindPbRpcStateProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateProof *PbRpcStateProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateProof.Contract.PbRpcStateProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateProof *PbRpcStateProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.PbRpcStateProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateProof *PbRpcStateProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.PbRpcStateProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateProof *PbRpcStateProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateProof *PbRpcStateProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateProof *PbRpcStateProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.contract.Transact(opts, method, params...)
}
