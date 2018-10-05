
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:9545'));

const protoChainLoader = require('./helper/protoChainLoader');
const solidityLoader = require('./helper/solidityLoader');

var Gomoku = artifacts.require('Gomoku')

const getStateProofBytes = async (nonce, state) => {
  const protoChain = await protoChainLoader();
  const solidity = await solidityLoader();
  const stateProof = {
    nonce: solidity.uint256.create({ data: [nonce] }),
    state: state
  };
  const stateProofProto = protoChain.StateProof.create(stateProof);
  sp = protoChain.StateProof.encode(stateProofProto).finish().toJSON().data;
  return web3.utils.bytesToHex(sp);
};

const signMessage = async (msg, addr) => {
  const signature = await web3.eth.sign(msg, addr);
  const r = web3.utils.hexToBytes(signature.slice(0, 66));
  const s = web3.utils.hexToBytes('0x' + signature.slice(66, 130));
  const v = web3.utils.toDecimal(signature.slice(130, 132)) + 27
  return { r, s, v };
};

const getSignatures = async (msg, addr1, addr2) => {
  const protoChain = await protoChainLoader();
  const solidity = await solidityLoader();
  const messageHash = web3.utils.keccak256(msg);
  const sig1 = await signMessage(messageHash, addr1)
  const sig2 = await signMessage(messageHash, addr2)
  const multiSignature = {
    v: [
      solidity.uint256.create({ data: [sig1.v] }),
      solidity.uint256.create({ data: [sig2.v] })
    ],
    r: [
      solidity.bytes32.create({ data: sig1.r }),
      solidity.bytes32.create({ data: sig2.r })
    ],
    s: [
      solidity.bytes32.create({ data: sig1.s }),
      solidity.bytes32.create({ data: sig2.s })
    ]
  };
  const multiSignatureProto = protoChain.MultiSignature.create(multiSignature);
  sigs = protoChain.MultiSignature.encode(multiSignatureProto).finish().toJSON().data;
  return web3.utils.bytesToHex(sigs);
};

contract('Gomoku', async(accounts) => {
  console.log(web3.version)

  const player1 = accounts[0]
  const player2 = accounts[1]
  let instance
  let turn
  let res

  it('start a new game, player 2 submits state proof and wins', async() => {
    instance = await Gomoku.new(player1, player2, 3, 3, 5, 5)
    var state = new Array(227)
    state[0] = 2  // winner
    state[1] = 0  // turn
    const stateproof = await getStateProofBytes(12, state);
    const sigs = await getSignatures(stateproof, player1, player2);
    res = await instance.intendSettle(stateproof, sigs)
    res = await instance.isFinalized(web3.utils.bytesToHex([2]), 1000)
    assert.equal(res, true)
    res = await instance.queryResult(web3.utils.bytesToHex([2]))
    assert.equal(res, true)
  });

  it('start a new game, player 1 submits state proof', async() => {
    instance = await Gomoku.new(player1, player2, 3, 0, 5, 5)
    var state = new Array(227)
    state[0] = 0   // winner, not specified, but the board state will indicate player 1 wins.
    state[1] = 2   // turn
    state[2] = 1   // (0, 0)
    state[3] = 1   // (0, 1)
    state[4] = 1   // (0, 2)
    state[5] = 1   // (0, 3)
    state[101] = 2
    state[102] = 2
    state[103] = 2
    const stateproof = await getStateProofBytes(8, state);
    const sigs = await getSignatures(stateproof, player1, player2);
    // submit state proof
    res = await instance.intendSettle(stateproof, sigs)
    // read on-chain state
    res = await instance.queryState();
    state = web3.utils.hexToBytes(res)
    assert.equal(state[0], 0)
    assert.equal(state[1], 2)
    assert.equal(state[2], 1)
    assert.equal(state[3], 1)
    assert.equal(state[4], 1)
    assert.equal(state[5], 1)
    assert.equal(state[101], 2)
    assert.equal(state[102], 2)
    assert.equal(state[103], 2)
  });

  it('player 2 places a stone at (3, 12) and palyer 1 takes the turn', async() => {
    res = await instance.confirmSettle()
    res = await instance.placeStone(3, 12, {from: player2})
    turn = await instance.turn()
    assert.equal(turn.valueOf(), 1)
  });

  it('player 2 tries to place another stone and should fail', async() => {
    try {
      res = await instance.placeStone(4, 12, {from: player2})
      assert.fail('expected revert not received')
    } catch (e) {
      turn = await instance.turn()
      assert.equal(turn.valueOf(), 1)
    }
  });

  it('player 1 tries to place a stone at occupied slot (3, 12) and shoud fail', async() => {
    try {
      res = await instance.placeStone(3, 12, {from: player1})
      assert.fail('expected revert not received')
    } catch (e) {
      turn = await instance.turn()
      assert.equal(turn.valueOf(), 1)
    }
  });

  it('player 1 places a stone at (0, 4) and wins the game', async() => {
    res = await instance.placeStone(0, 4, {from: player1})
    turn = await instance.turn()
    assert.equal(turn.valueOf(), 0)
    res = await instance.isFinalized(web3.utils.bytesToHex([1]), 1000)
    assert.equal(res, true)
    res = await instance.queryResult(web3.utils.bytesToHex([1]))
    assert.equal(res, true)
  });
});