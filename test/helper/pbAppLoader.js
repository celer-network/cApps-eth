const protobuf = require('protobufjs');
protobuf.common('google/protobuf/descriptor.proto', {});

module.exports = async () => {
  pbApp = await protobuf.load(
    `${__dirname}/../../contracts/lib/proto/app.proto`
  );
  pbSession = await protobuf.load(
    `${__dirname}/../../contracts/lib/proto/session.proto`
  );

  const StateProof = pbApp.lookupType('app.StateProof');
  const AppState = pbApp.lookupType('app.AppState');
  const SessionQuery = pbSession.lookupType('session.SessionQuery');

  return {
    StateProof,
    AppState,
    SessionQuery,
  };
};
