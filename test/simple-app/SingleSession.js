const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

const utils = require('../helper/utils');
const pbApp = require('../helper/pbAppFactory');
const SimpleSingleSessionApp = artifacts.require('SimpleSingleSessionApp');

const fs = require('fs');
const GAS_USED_LOG = 'gas/SimpleSingleSessionApp.txt';

contract('SimpleSingleSessionApp', async accounts => {
  const players = [accounts[0], accounts[1]];
  const nonce = 0;
  const timeout = 2;
  let instance;
  let seq;
  let res;
  let state;
  let stateProof;
  let gasUsed;

  before(async () => {
    fs.writeFileSync(GAS_USED_LOG, '******* Gas Used in SimpleSingleSessionApp Tests ********\n\n');
  });

  it('new app deployed and first apply action should fail', async () => {
    instance = await SimpleSingleSessionApp.new(players, nonce, timeout);
    gasUsed = await utils.getDeployGasUsed(instance);
    fs.appendFileSync(GAS_USED_LOG, 'contract deploy: ' + gasUsed + '\n');
    try {
      res = await instance.applyAction(web3.utils.bytesToHex([1]));
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('first intend settle should success', async () => {
    seq = 2;
    state = [5];
    stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      timeout,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'first intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    const { event, args } = res.logs[0];
    assert.equal(event, 'IntendSettle');
    assert.equal(args.seq.toString(10), seq);
    res = await instance.getResult(web3.utils.bytesToHex([5]));
    assert.equal(res, true);
    res = await instance.isFinalized([]);
    assert.equal(res, false);
  });

  it('apply action before settle finalized time should fail', async () => {
    try {
      res = await instance.applyAction(web3.utils.bytesToHex([1]));
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('apply action after settle finalized time should success', async () => {
    const settleFinalizedTime = await instance.getSettleFinalizedTime();
    await utils.waitBlock(settleFinalizedTime, accounts[0]);
    res = await instance.applyAction(web3.utils.bytesToHex([3]));
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    res = await instance.getResult(web3.utils.bytesToHex([3]));
    assert.equal(res, true);
  });

  it('intend settle with invalid sequence number should fail', async () => {
    seq = 1;
    state = [2];
    stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      timeout,
      players
    );
    try {
      res = await instance.intendSettle(stateProof);
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
    seq = 4;
    state = [2];
    stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      timeout,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'second intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    const { event, args } = res.logs[0];
    assert.equal(event, 'IntendSettle');
    assert.equal(args.seq.toString(10), seq);
    res = await instance.getResult(web3.utils.bytesToHex([2]));
    assert.equal(res, true);
    res = await instance.isFinalized([]);
    assert.equal(res, true);
  });

  it('apply action after finalized should fail', async () => {
    try {
      res = await instance.applyAction(web3.utils.bytesToHex([1]));
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('intend settle after finalized should fail', async () => {
    seq = 5;
    state = [1];
    stateProof = await pbApp.encodeStateProof(
      nonce,
      seq,
      state,
      timeout,
      players
    );
    try {
      res = await instance.intendSettle(stateProof);
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });
});
