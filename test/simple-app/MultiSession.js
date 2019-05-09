const utils = require('../helper/utils');
const pbApp = require('../helper/pbAppFactory');
const SimpleMultiSessionApp = artifacts.require('SimpleMultiSessionApp');

const fs = require('fs');
const GAS_USED_LOG = 'gas/SimpleMultiSessionApp.txt';

contract('SimpleMultiSessionApp', async accounts => {
  const players = [accounts[0], accounts[1]];
  players.sort();
  const nonce1 = 1;
  const nonce2 = 2;
  const session1 = utils.getSessionID(nonce1, players[0], players[1]);
  const session2 = utils.getSessionID(nonce2, players[0], players[1]);
  const timeout = 2;
  let instance;
  let res;
  let gasUsed;

  before(async () => {
    fs.writeFileSync(GAS_USED_LOG, '******* Gas Used in SimpleMultiSessionApp Tests ********\n\n');
  });

  it('new app deployed and first intend settle should success', async () => {
    instance = await SimpleMultiSessionApp.new(timeout, 2);
    gasUsed = await utils.getDeployGasUsed(instance);
    fs.appendFileSync(GAS_USED_LOG, 'contract deploy: ' + gasUsed + '\n');
    const seq = 2;
    const state = [5];
    const stateProof = await pbApp.encodeStateProof(
      nonce1,
      seq,
      state,
      timeout,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'first intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    const { event, args } = res.logs[0];
    assert.equal(event, 'IntendSettle');
    assert.equal(args.session, session1);

    const arg = await pbApp.encodeSessionQuery(session1, [5]);
    res = await instance.getResult(arg);
    assert.equal(res, true);
    res = await instance.isFinalized(session1);
    assert.equal(res, false);
  });

  it('apply action before settle finalized time should fail', async () => {
    try {
      await instance.applyAction(session1, [1]);
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
    const settleFinalizedTime = await instance.getSettleFinalizedTime(session1);
    await utils.waitBlock(settleFinalizedTime, accounts[0]);

    res = await instance.applyAction(session1, [3]);
    fs.appendFileSync(GAS_USED_LOG, 'applyAction: ' + utils.getCallGasUsed(res) + '\n');
    const query = await pbApp.encodeSessionQuery(session1, [3]);
    res = await instance.getResult(query);
    assert.equal(res, true);
  });

  it('apply action on a different session should fail', async () => {
    try {
      await instance.applyAction(session2, [1]);
    } catch (e) {
      assert.isAbove(
        e.message.search('VM Exception while processing transaction'),
        -1
      );
      return;
    }
    assert.fail('should have thrown before');
  });

  it('intend settle with invalid sequence number should fail', async () => {
    const seq = 1;
    const state = [2];
    const stateProof = await pbApp.encodeStateProof(
      nonce1,
      seq,
      state,
      timeout,
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

  it('intend settle with different player sigs should fail', async () => {
    const seq = 4;
    const state = [2];
    const stateProof = await pbApp.encodeStateProof(
      nonce1,
      seq,
      state,
      timeout,
      [players[0], accounts[2]]
    );
    await instance.intendSettle(stateProof);
    const arg = await pbApp.encodeSessionQuery(session1, [3]);
    const res = await instance.getResult(arg);
    assert.equal(res, true);
  });

  it('intend settle with a different nonce should fail', async () => {
    const seq = 4;
    const state = [2];
    const stateProof = await pbApp.encodeStateProof(
      nonce2,
      seq,
      state,
      timeout,
      players
    );
    await instance.intendSettle(stateProof);
    const arg = await pbApp.encodeSessionQuery(session1, [3]);
    const res = await instance.getResult(arg);
    assert.equal(res, true);
  });

  it('intend settle with a valid seq, sig, nonce should success', async () => {
    const seq = 4;
    const state = [2];
    const stateProof = await pbApp.encodeStateProof(
      nonce1,
      seq,
      state,
      timeout,
      players
    );
    res = await instance.intendSettle(stateProof);
    fs.appendFileSync(GAS_USED_LOG, 'second intendSettle: ' + utils.getCallGasUsed(res) + '\n');
    const arg = await pbApp.encodeSessionQuery(session1, [2]);
    res = await instance.getResult(arg);
    assert.equal(res, true);
    res = await instance.isFinalized(session1);
    assert.equal(res, true);
  });

  it('apply action after finalized should fail', async () => {
    try {
      await instance.applyAction(session1, [1]);
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
    const seq = 5;
    const state = [1];
    const stateProof = await pbApp.encodeStateProof(
      nonce1,
      seq,
      state,
      timeout,
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
});
