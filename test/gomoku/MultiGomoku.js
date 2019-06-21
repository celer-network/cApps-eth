const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

const utils = require('../helper/utils');
const pbApp = require('../helper/pbAppFactory');
const MultiGomoku = artifacts.require('MultiGomoku');

const fs = require('fs');
const GAS_USED_LOG = 'gas/MultiGomoku.txt';

contract('MultiGomoku', async accounts => {
  const players1 = [accounts[0], accounts[1]];
  const players2 = [accounts[1], accounts[2]];
  players1.sort();
  players2.sort();
  const nonce1 = 1;
  const nonce2 = 2;
  const session1 = utils.getSessionID(nonce1, players1[0], players1[1]);
  const session2 = utils.getSessionID(nonce2, players2[0], players2[1]);
  const timeout = 2;
  let instance;
  let res;
  let gasUsed;

  before(async () => {
    fs.writeFileSync(GAS_USED_LOG, '******* Gas Used in SingleGomoku Tests ********\n\n');
  });

  it('new app deployed', async () => {
    instance = await MultiGomoku.new(5, 5);
    gasUsed = await utils.getDeployGasUsed(instance);
    fs.appendFileSync(GAS_USED_LOG, 'contract deploy: ' + gasUsed + '\n');
  });

  it('player 1 submits state proof shoudl success', async () => {
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
      nonce1,
      seq,
      state,
      timeout,
      players1
    );
    // submit state proof
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'first intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.getState(session1, 2);
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
      nonce1,
      seq,
      state,
      timeout,
      players1
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
      nonce1,
      seq,
      state,
      timeout,
      players1
    );
    // submit state proof
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'second intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.getState(session1, 2);
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

  it('player 2 places a stone at (3, 12) before settle finalized time should fail', async () => {
    try {
      await instance.applyAction(session1, [3, 12], {
        from: players1[1],
      });
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('player 2 places a stone at (3, 12) after settle finalized time should success', async () => {
    const settleFinalizedTime = await instance.getSettleFinalizedTime(session1);
    await utils.waitBlock(settleFinalizedTime, accounts[0]);
    res = await instance.applyAction(session1, [3, 12], { from: players1[1] });
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const turn = await instance.getState(session1, 0);
    assert.equal(turn.valueOf(), 1);
  });

  it('player 2 tries to place another stone and should fail', async () => {
    try {
      await instance.applyAction(session1, [4, 12], { from: players1[1] });
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
      await instance.applyAction(session1, [3, 12], { from: players1[0] });
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
    res = await instance.applyAction(session1, [0, 4], { from: players1[0] });
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const turn = await instance.getState(session1, 0);
    assert.equal(turn.valueOf(), 0);
    res = await instance.isFinalized(session1);
    assert.equal(res, true);
    const arg = await pbApp.encodeSessionQuery(session1, [1]);
    res = await instance.getOutcome(arg);
    assert.equal(res, true);
  });

  it('intend settle with a different nonce for player 2 and 3 should success', async () => {
    const seq = 1;
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
      nonce2,
      seq,
      state,
      timeout,
      players2
    );
    // submit state proof
    await instance.intendSettle(stateProof);
    res = await instance.getState(session2, 2);
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

  it('player 1 places a stone for session 2 should fail', async () => {
    const settleFinalizedTime = await instance.getSettleFinalizedTime(session2);
    await utils.waitBlock(settleFinalizedTime, accounts[0]);
    try {
      await instance.applyAction(session2, [3, 12], {
        from: players1[0],
      });
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('player 3 places a stone at (3, 12) for session 2 should success', async () => {
    res = await instance.applyAction(session2, [3, 12], { from: players2[1] });
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const turn = await instance.getState(session2, 0);
    assert.equal(turn.valueOf(), 1);
  });

  it('finalizeOnTimeout before action deadline should fail', async () => {
    try {
      await instance.finalizeOnActionTimeout(session2);
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('finalizeOnTimeout after action deadline should success', async () => {
    const deadline = await instance.getActionDeadline(session2);
    await utils.waitBlock(deadline, accounts[0]);
    res = await instance.finalizeOnActionTimeout(session2);
    fs.appendFileSync(GAS_USED_LOG, 'finalizeOnActionTimeout: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.isFinalized(session2);
    assert.equal(res, true);
    const arg = await pbApp.encodeSessionQuery(session2, [2]);
    res = await instance.getOutcome(arg);
    assert.equal(res, true);
  });
});
