package keysv2

import (
	"context"
	"crypto/ed25519"
	"errors"

	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

var _ ragetypes.PeerKeyring = (*PeerKeyring)(nil)

const (
	PeerKeyringPrefix = "ragep2p_peer"
)

type PeerKeyring struct {
	ks      commonks.Keystore
	keyPath KeyPath
	pubKey  ragetypes.PeerPublicKey
}

func (k *PeerKeyring) KeyPath() KeyPath {
	return k.keyPath
}

func (k *PeerKeyring) PublicKey() ragetypes.PeerPublicKey {
	return k.pubKey
}

func (k *PeerKeyring) PeerID() (string, error) {
	peerID, err := ragetypes.PeerIDFromPublicKey(ed25519.PublicKey(k.pubKey[:]))
	if err != nil {
		return "", err
	}
	return peerID.String(), nil
}

func (k *PeerKeyring) MustPeerID() string {
	peerID, err := k.PeerID()
	if err != nil {
		panic(err)
	}
	return peerID
}

func (k *PeerKeyring) Sign(msg []byte) ([]byte, error) {
	resp, err := k.ks.Sign(context.Background(), commonks.SignRequest{
		KeyName: k.keyPath.String(),
		Data:    msg,
	})
	if err != nil {
		return nil, err
	}
	return resp.Signature, nil
}

func CreatePeerKeyring(ctx context.Context, ks commonks.Keystore, name string) (*PeerKeyring, error) {
	keyPath := NewKeyPath(PeerKeyringPrefix, name)
	createReq := commonks.CreateKeysRequest{
		Keys: []commonks.CreateKeyRequest{
			{KeyName: keyPath.String(), KeyType: commonks.Ed25519},
		},
	}
	resp, err := ks.CreateKeys(ctx, createReq)
	if err != nil {
		return nil, err
	}
	if len(resp.Keys) != 1 {
		return nil, errors.New("expected 1 key")
	}
	var peerPubKey ragetypes.PeerPublicKey
	copy(peerPubKey[:], resp.Keys[0].KeyInfo.PublicKey)
	return &PeerKeyring{ks: ks, keyPath: keyPath, pubKey: peerPubKey}, nil
}
