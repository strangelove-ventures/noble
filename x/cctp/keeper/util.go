package keeper

import (
	"bytes"
	sdkerrors "cosmossdk.io/errors"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func parseIntoMessage(msg []byte) types.Message {
	message := types.Message{
		Version:           binary.BigEndian.Uint32(msg[0:4]),
		SourceDomainBytes: msg[4:8],
		SourceDomain:      binary.BigEndian.Uint32(msg[4:8]),
		DestinationDomain: binary.BigEndian.Uint32(msg[8:12]),
		NonceBytes:        msg[12:20],
		Nonce:             binary.BigEndian.Uint64(msg[12:20]),
		Sender:            msg[20:52],
		Recipient:         msg[52:84],
		DestinationCaller: msg[84:116],
		MessageBody:       msg[116:],
	}

	return message
}

func parseIntoBurnMessage(msg []byte) types.BurnMessage {
	message := types.BurnMessage{
		Version:       binary.BigEndian.Uint32(msg[0:4]),
		BurnToken:     msg[4:36],
		MintRecipient: msg[36:68],
		Amount:        binary.BigEndian.Uint64(msg[68:100]),
		MessageSender: msg[100:132],
	}

	return message
}

func parseBurnMessageIntoBytes(msg types.BurnMessage) []byte {
	result := make([]byte, burnMessageLength)

	versionBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(versionBytes, msg.Version)

	amountBytes := make([]byte, 32)
	binary.LittleEndian.PutUint64(amountBytes, msg.Amount)

	copyBytes(0, 4, versionBytes, &result)
	copyBytes(4, 36, msg.BurnToken, &result)
	copyBytes(36, 68, msg.MintRecipient, &result)
	copyBytes(68, 100, amountBytes, &result)

	return result
}

func copyBytes(start int, end int, copyFrom []byte, copyInto *[]byte) {
	for i := start; i < end; i++ {
		(*copyInto)[i] = copyFrom[i]
	}
}

/*
* Rules for valid attestation:
* 1. length of `_attestation` == 65 (signature length) * signatureThreshold
* 2. addresses recovered from attestation must be in increasing order.
* 	For example, if signature A is signed by address 0x1..., and signature B
* 		is signed by address 0x2..., attestation must be passed as AB.
* 3. no duplicate signers
* 4. all signers must be enabled attesters
 */
func verifyAttestationSignatures(
	message []byte,
	attestation []byte,
	publicKeys []types.PublicKeys,
	signatureThreshold uint32) (bool, error) {

	if uint32(len(attestation)) != signatureLength*signatureThreshold {
		return false, sdkerrors.Wrap(types.ErrSignatureVerification, "invalid attestation length")
	}

	if signatureThreshold == 0 {
		return false, sdkerrors.Wrap(types.ErrSignatureVerification, "signature verification threshold cannot be 0")
	}

	// public keys cannot be empty, so the recovered key should be bigger than latestPublicKey
	var latestPublicKey []byte

	digest := crypto.Keccak256Hash(message).Bytes()

	for i := uint32(0); i < signatureThreshold; i++ {
		signature := attestation[i*signatureLength : (i*signatureLength)+signatureLength]

		recoveredKey, err := crypto.Ecrecover(digest, signature)
		if err != nil {
			return false, sdkerrors.Wrap(types.ErrSignatureVerification, "failed to recover public key")
		}

		// Signatures must be in increasing order of address, and may not duplicate signatures from same address
		if bytes.Compare(latestPublicKey, recoveredKey) > -1 {
			return false, sdkerrors.Wrap(types.ErrSignatureVerification, "Invalid signature order or dupe")
		}

		// check that recovered key is a valid
		contains := false
		for _, key := range publicKeys {
			if key.Key == string(recoveredKey) {
				contains = true
			}
		}

		if !contains {
			return false, sdkerrors.Wrap(types.ErrSignatureVerification, "Invalid signature: not attester")
		}

		latestPublicKey = recoveredKey

	}
	return true, nil
}
