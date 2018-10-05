const protobuf = require("protobufjs");

module.exports = async () =>
  protobuf.load(`${__dirname}/../../contracts/lib/proto/chain.proto`)
    .then(function (root) {
      const StateProof = root.lookupType("rpc.StateProof");
      const MultiSignature = root.lookupType("rpc.MultiSignature");

      return {
        StateProof,
        MultiSignature
      }
    });