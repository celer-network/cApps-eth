function getSessionID (nonce, addr1, addr2) {
  let session;
  if (addr1 < addr2) {
    session = web3.eth.abi.encodeParameters(['uint', 'address[]'], [nonce, [addr1, addr2]]);
  } else {
    session = web3.eth.abi.encodeParameters(['uint', 'address[]'], [nonce, [addr2, addr1]]);
  }
  return web3.utils.keccak256(session);
}

function padLeft (data) {
  const pad = (32 - (data.length - 2) % 32) % 32;
  let x = data.substr(2, data.length);
  for (let i = 0; i < pad; i++) {
    x = 0 + x;
  }
  return '0x' + x;
}

async function waitBlock (block, addr) {
  let blocknumber = await web3.eth.getBlockNumber();
  while (blocknumber <= block) {
    await web3.eth.sendTransaction({ from: addr });
    blocknumber = await web3.eth.getBlockNumber();
  }
}

async function getDeployGasUsed(instance) {
  let receipt = await web3.eth.getTransactionReceipt(instance.transactionHash);
  return receipt.gasUsed;
}

function getCallGasUsed(tx) {
  return tx.receipt.gasUsed;
}

module.exports.getSessionID = getSessionID;
module.exports.padLeft = padLeft;
module.exports.waitBlock = waitBlock;
module.exports.getDeployGasUsed = getDeployGasUsed;
module.exports.getCallGasUsed = getCallGasUsed;
