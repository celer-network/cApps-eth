pragma solidity ^0.5.0;

library LibSignature {

    /**
     * @dev Recover signers from message and signatures
     * @param state Signed state message
     * @param sigs Signature array of the state
     * @return Siger address array
     */

    function recoverSigners(bytes memory state, bytes[] memory sigs, bool ordered)
        internal
        pure
        returns (address[] memory)
    {
        address[] memory signers = new address[](sigs.length);
        bytes32 digest = prefixedDigest(keccak256(abi.encodePacked(state)));

        address prev = address(0);
        for (uint i = 0; i < sigs.length; i++) {
            signers[i] = recoverSigner(digest, sigs[i]);
            if (ordered) {
                require(signers[i] > prev, "signers not in ascending order");
                prev = signers[i];
            }
        }
        return signers;
    }

    /**
     * @dev Recover signer address from message and signatures
     * @param digest Signed message digest
     * @param sig Signature of the message
     * @return Siger address
     */
    function recoverSigner(bytes32 digest, bytes memory sig)
        internal
        pure
        returns (address)
    {
        if (sig.length != 65) {
            return (address(0));
        }

        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
        if (v < 27) {
            v += 27;
        }

        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            return ecrecover(digest, v, r, s);
        }
    }

    /**
     * @dev ETH signed prefixed message digest
     */
    function prefixedDigest(bytes32 message) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", message));
    }
}