// Code generated by protoc-gen-sol. DO NOT EDIT.
// source: contracts/lib/proto/app.proto
pragma solidity ^0.5.0;
import "./Pb.sol";

library PbApp {
    using Pb for Pb.Buffer;  // so we can call Pb funcs on Buffer obj

    struct AppState {
        uint nonce;   // tag: 1
        uint seqNum;   // tag: 2
        bytes state;   // tag: 3
        uint timeout;   // tag: 4
    } // end struct AppState

    function decAppState(bytes memory raw) internal pure returns (AppState memory m) {
        Pb.Buffer memory buf = Pb.fromBytes(raw);

        uint tag;
        Pb.WireType wire;
        while (buf.hasMore()) {
            (tag, wire) = buf.decKey();
            if (false) {} // solidity has no switch/case
            else if (tag == 1) {
                m.nonce = uint(buf.decVarint());
            }
            else if (tag == 2) {
                m.seqNum = uint(buf.decVarint());
            }
            else if (tag == 3) {
                m.state = bytes(buf.decBytes());
            }
            else if (tag == 4) {
                m.timeout = uint(buf.decVarint());
            }
            else { buf.skipValue(wire); } // skip value of unknown tag
        }
    } // end decoder AppState

    struct StateProof {
        bytes appState;   // tag: 1
        bytes[] sigs;   // tag: 2
    } // end struct StateProof

    function decStateProof(bytes memory raw) internal pure returns (StateProof memory m) {
        Pb.Buffer memory buf = Pb.fromBytes(raw);

        uint[] memory cnts = buf.cntTags(2);
        m.sigs = new bytes[](cnts[2]);
        cnts[2] = 0;  // reset counter for later use
        
        uint tag;
        Pb.WireType wire;
        while (buf.hasMore()) {
            (tag, wire) = buf.decKey();
            if (false) {} // solidity has no switch/case
            else if (tag == 1) {
                m.appState = bytes(buf.decBytes());
            }
            else if (tag == 2) {
                m.sigs[cnts[2]] = bytes(buf.decBytes());
                cnts[2]++;
            }
            else { buf.skipValue(wire); } // skip value of unknown tag
        }
    } // end decoder StateProof

    struct SessionQuery {
        bytes32 session;   // tag: 1
        bytes query;   // tag: 2
    } // end struct SessionQuery

    function decSessionQuery(bytes memory raw) internal pure returns (SessionQuery memory m) {
        Pb.Buffer memory buf = Pb.fromBytes(raw);

        uint tag;
        Pb.WireType wire;
        while (buf.hasMore()) {
            (tag, wire) = buf.decKey();
            if (false) {} // solidity has no switch/case
            else if (tag == 1) {
                m.session = Pb._bytes32(buf.decBytes());
            }
            else if (tag == 2) {
                m.query = bytes(buf.decBytes());
            }
            else { buf.skipValue(wire); } // skip value of unknown tag
        }
    } // end decoder SessionQuery

}
