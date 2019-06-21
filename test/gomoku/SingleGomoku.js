const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

const utils = require('../helper/utils');
const pbApp = require('../helper/pbAppFactory');
const SingleGomoku = artifacts.require('SingleGomoku');

const fs = require('fs');
const GAS_USED_LOG = 'gas/SingleGomoku.txt';

contract('SingleGomoku', async accounts => {
  const players = [accounts[0], accounts[1]];
  const nonce = 0;
  const timeout = 2;
  let instance;
  let res;
  let gasUsed;

  before(async () => {
    fs.writeFileSync(GAS_USED_LOG, '******* Gas Used in SingleGomoku Tests ********\n\n');
  });

  it('start a new game and intend settle, player 2 should win', async () => {
    instance = await SingleGomoku.new(players, nonce, timeout, 5, 5);
    gasUsed = await utils.getDeployGasUsed(instance);
    fs.appendFileSync(GAS_USED_LOG, 'contract deploy: ' + gasUsed + '\n');
    const seq = 2;
    const state = new Array(227);
    state[0] = 2; // winner
    state[1] = 0; // turn
    const stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      null,
      players
    );
    await instance.intendSettle(stateProof);
    res = await instance.isFinalized([]);
    assert.equal(res, true);
    res = await instance.getOutcome(web3.utils.bytesToHex([2]));
    assert.equal(res, true);
  });

  it('start a new game and intend settle should success', async () => {
    instance = await SingleGomoku.new(players, nonce, timeout, 5, 5);
    const seq = 3;
    const state = new Array(227);
    state[0] = 0;
    state[1] = 1;
    state[2] = 2;
    state[3] = 2;
    state[4] = 1;
    state[5] = 1;
    state[6] = 2;
    state[7] = 2;
    state[8] = 1;
    const stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      null,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'first intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.getState(2);
    const onchainState = web3.utils.hexToBytes(res);
    assert.equal(onchainState[0], 0);
    assert.equal(onchainState[1], 1);
    assert.equal(onchainState[2], 2);
    assert.equal(onchainState[3], 2);
    assert.equal(onchainState[4], 1);
    assert.equal(onchainState[5], 1);
    assert.equal(onchainState[6], 2);
    assert.equal(onchainState[7], 2);
    assert.equal(onchainState[8], 1);
  });

  it('intend settle with invalid sequence number should fail', async () => {
    const seq = 2;
    const state = new Array(227);
    const stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      null,
      players
    );
    try {
      await instance.intendSettle(stateProof);
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('intend settle with higher sequence number should success', async () => {
    const seq = 4;
    const state = new Array(227);
    state[0] = 0; // winner
    state[1] = 2; // turn
    state[2] = 1; // (0, 0)
    state[3] = 1; // (0, 1)
    state[4] = 1; // (0, 2)
    state[5] = 1; // (0, 3)
    state[101] = 2;
    state[102] = 2;
    state[103] = 2;
    const stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      null,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'second intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.getState(2);
    const onchainState = web3.utils.hexToBytes(res);
    assert.equal(onchainState[0], 0);
    assert.equal(onchainState[1], 2);
    assert.equal(onchainState[2], 1);
    assert.equal(onchainState[3], 1);
    assert.equal(onchainState[4], 1);
    assert.equal(onchainState[5], 1);
    assert.equal(onchainState[6], 0);
    assert.equal(onchainState[7], 0);
    assert.equal(onchainState[8], 0);
    assert.equal(onchainState[101], 2);
    assert.equal(onchainState[102], 2);
    assert.equal(onchainState[103], 2);
  });

  it('player 2 places a stone at (3, 12) and player 1 takes the turn', async () => {
    const settleFinalizedTime = await instance.getSettleFinalizedTime();
    await utils.waitBlock(settleFinalizedTime, accounts[0]);
    res = await instance.applyAction([3, 12], { from: players[1] });
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const turn = await instance.getState(0);
    assert.equal(turn.valueOf(), 1);
  });

  it('player 2 tries to place another stone and should fail', async () => {
    try {
      await instance.applyAction([4, 12], { from: players[1] });
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('player 1 tries to place a stone at occupied slot (3, 12) and shoud fail', async () => {
    try {
      await instance.applyAction([3, 12], { from: players[0] });
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('player 1 places a stone at (0, 4) and wins the game', async () => {
    res = await instance.applyAction([0, 4], { from: players[0] });
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const turn = await instance.getState(0);
    assert.equal(turn.valueOf(), 0);
    res = await instance.isFinalized([]);
    assert.equal(res, true);
    res = await instance.getOutcome(web3.utils.bytesToHex([1]));
    assert.equal(res, true);
  });

  it('finalizeOnActionTimeout before action deadline should fail', async () => {
    instance = await SingleGomoku.new(players, nonce, timeout, 5, 5);
    const seq = 3;
    const state = new Array(227);
    state[0] = 0; // winner
    state[1] = 2; // turn
    state[2] = 1; // (0, 0)
    state[3] = 1; // (0, 1)
    state[4] = 1; // (0, 2)
    state[5] = 1; // (0, 3)
    state[101] = 2;
    state[102] = 2;
    state[103] = 2;
    const stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      null,
      players
    );
    await instance.intendSettle(stateProof);

    try {
      await instance.finalizeOnActionTimeout();
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('finalizeOnActionTimeout after action deadline should success', async () => {
    const deadline = await instance.getActionDeadline();
    await utils.waitBlock(deadline, accounts[0]);
    res = await instance.finalizeOnActionTimeout();
    fs.appendFileSync(GAS_USED_LOG, 'finalizeOnActionTimeout: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.isFinalized([]);
    assert.equal(res, true);
    res = await instance.getOutcome(web3.utils.bytesToHex([1]));
    assert.equal(res, true);
  });
});
