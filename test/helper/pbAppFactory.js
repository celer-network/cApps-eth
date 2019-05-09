const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:7545'));

const pbAppLoader = require('./pbAppLoader');

async function encodeStateProof (nonce, seqNum, state, timeout, players) {
  const pbApp = await pbAppLoader();
  const appState = {
    nonce,
    seqNum,
    state,
    timeout,
  };
  const appStateMsg = pbApp.AppState.create(appState);
  const appStateBuffer = pbApp.AppState.encode(appStateMsg).finish();
  const appStateHash = web3.utils.keccak256(appStateBuffer);

  const sigs = await Promise.all(
    players.map(async player => {
      const sig = await web3.eth.sign(appStateHash, player);
      return web3.utils.hexToBytes(sig);
    })
  );

  const stateProof = {
    sigs,
    appState: appStateBuffer,
  };
  const stateProofMsg = pbApp.StateProof.create(stateProof);
  const stateProofBuffer = pbApp.StateProof.encode(stateProofMsg).finish();

  return web3.utils.bytesToHex(stateProofBuffer);
}

async function encodeSessionQuery (session, query) {
  const pbApp = await pbAppLoader();
  const request = {
    query,
    session: web3.utils.hexToBytes(session),
  };
  const requestMsg = pbApp.SessionQuery.create(request);
  const requestBuffer = pbApp.SessionQuery.encode(requestMsg).finish();
  return web3.utils.bytesToHex(requestBuffer);
}

module.exports.encodeStateProof = encodeStateProof;
module.exports.encodeSessionQuery = encodeSessionQuery;
