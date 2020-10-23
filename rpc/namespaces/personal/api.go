package personal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/ethermint/crypto/ethsecp256k1"
	"github.com/cosmos/ethermint/crypto/hd"
	"github.com/cosmos/ethermint/rpc/namespaces/eth"
	rpctypes "github.com/cosmos/ethermint/rpc/types"
)

// PrivateAccountAPI is the personal_ prefixed set of APIs in the Web3 JSON-RPC spec.
type PrivateAccountAPI struct {
	ethAPI           *eth.PublicEthereumAPI
	logger           log.Logger
	unlockedAccounts map[common.Address]bool
}

// NewAPI creates an instance of the public Personal Eth API.
func NewAPI(ethAPI *eth.PublicEthereumAPI) *PrivateAccountAPI {
	api := &PrivateAccountAPI{
		ethAPI:           ethAPI,
		logger:           log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "json-rpc", "namespace", "personal"),
		unlockedAccounts: make(map[common.Address]bool),
	}

	_ = api.ethAPI.GetKeyringInfo()
	return api
}

// ImportRawKey armors and encrypts a given raw hex encoded ECDSA key and stores it into the key directory.
// The name of the key will have the format "personal_<length-keys>", where <length-keys> is the total number of
// keys stored on the keyring.
// NOTE: The key will be both armored and encrypted using the same passphrase. By default the newly created account is
// locked.
func (api *PrivateAccountAPI) ImportRawKey(privkey, password string) (common.Address, error) {
	api.logger.Debug("personal_importRawKey")
	priv, err := crypto.HexToECDSA(privkey)
	if err != nil {
		return common.Address{}, err
	}

	privKey := ethsecp256k1.PrivKey(crypto.FromECDSA(priv))

	armor := mintkey.EncryptArmorPrivKey(privKey, password, ethsecp256k1.KeyType)

	// ignore error as we only care about the length of the list
	list, _ := api.ethAPI.ClientCtx().Keybase.List()
	privKeyName := fmt.Sprintf("personal_%d", len(list))

	if err := api.ethAPI.ClientCtx().Keybase.ImportPrivKey(privKeyName, armor, password); err != nil {
		return common.Address{}, err
	}

	addr := common.BytesToAddress(privKey.PubKey().Address().Bytes())
	api.unlockedAccounts[addr] = false

	api.logger.Info("key successfully imported", "name", privKeyName, "address", addr.String())
	return addr, nil
}

// ListAccounts will return a list of addresses for accounts this node manages.
func (api *PrivateAccountAPI) ListAccounts() ([]common.Address, error) {
	api.logger.Debug("personal_listAccounts")
	addrs := []common.Address{}

	infos, err := api.ethAPI.ClientCtx().Keybase.List()
	if err != nil {
		return addrs, err
	}

	for _, info := range infos {
		addressBytes := info.GetPubKey().Address().Bytes()
		addrs = append(addrs, common.BytesToAddress(addressBytes))
	}

	return addrs, nil
}

// LockAccount will lock the account associated with the given address when it's unlocked.
// It removes the key corresponding to the given address from the API's local keys.
func (api *PrivateAccountAPI) LockAccount(address common.Address) bool {
	api.logger.Debug("personal_lockAccount", "address", address.String())

	// return true if the account is already locked
	if !api.unlockedAccounts[address] {
		return true
	}

	addr := sdk.AccAddress(address.Bytes())

	_, err := api.ethAPI.ClientCtx().Keybase.GetByAddress(addr)
	if err != nil {
		return false
	}

	api.unlockedAccounts[address] = false

	api.logger.Debug("account unlocked", "address", address.String())
	return true
}

// NewAccount will create a new account and returns the address for the new account.
// By default the newly created account is locked.
func (api *PrivateAccountAPI) NewAccount(password string) (common.Address, error) {
	api.logger.Debug("personal_newAccount")

	name := "key_" + time.Now().UTC().Format(time.RFC3339)
	info, _, err := api.ethAPI.ClientCtx().Keybase.CreateMnemonic(name, keys.English, password, hd.EthSecp256k1)
	if err != nil {
		return common.Address{}, err
	}

	addr := common.BytesToAddress(info.GetPubKey().Address().Bytes())

	api.unlockedAccounts[addr] = false

	api.logger.Info("Your new key was generated", "address", addr.String())
	api.logger.Info("Please backup your key file!", "path", os.Getenv("HOME")+"/.ethermintd/"+name)
	api.logger.Info("Please remember your password!")
	return addr, nil
}

// UnlockAccount will unlock the account associated with the given address with
// the given password for duration seconds. If duration is nil it will use a
// default of 300 seconds. It returns an indication if the account was unlocked.
// It exports the private key corresponding to the given address from the keyring and stores it in the API's local keys.
func (api *PrivateAccountAPI) UnlockAccount(_ context.Context, addr common.Address, password string, _ *uint64) (bool, error) { // nolint: interfacer
	api.logger.Debug("personal_unlockAccount", "address", addr.String())
	// TODO: use duration

	// return true if the account is already unlocked
	if api.unlockedAccounts[addr] {
		return true, nil
	}

	info, err := api.ethAPI.ClientCtx().Keybase.GetByAddress(addr.Bytes())
	if err != nil {
		return false, err
	}

	name := info.GetName()

	armor, err := api.ethAPI.ClientCtx().Keybase.ExportPrivKey(name, password, password)
	if err != nil {
		return false, err
	}

	if err := api.ethAPI.ClientCtx().Keybase.Delete(name, password, true); err != nil {
		return false, err
	}

	if err := api.ethAPI.ClientCtx().Keybase.ImportPrivKey(name, armor, password); err != nil {
		return false, err
	}

	api.unlockedAccounts[addr] = true

	api.logger.Debug("account unlocked", "address", addr.String())
	return true, nil
}

// SendTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.To. If the given password isn't
// able to decrypt the key it fails.
func (api *PrivateAccountAPI) SendTransaction(_ context.Context, args rpctypes.SendTxArgs, _ string) (common.Hash, error) {
	if !api.unlockedAccounts[args.From] {
		return common.Hash{}, fmt.Errorf("account %s is locked", args.From.String())
	}

	return api.ethAPI.SendTransaction(args)
}

// Sign calculates an Ethereum ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The key used to calculate the signature is decrypted with the given password.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_sign
func (api *PrivateAccountAPI) Sign(_ context.Context, data hexutil.Bytes, addr common.Address, password string) (hexutil.Bytes, error) {
	api.logger.Debug("personal_sign", "data", data, "address", addr.String())

	info, err := api.ethAPI.ClientCtx().Keybase.GetByAddress(sdk.AccAddress(addr.Bytes()))
	if err != nil {
		return nil, err
	}

	if !api.unlockedAccounts[addr] {
		return nil, fmt.Errorf("account %s is locked", addr.String())
	}

	sig, _, err := api.ethAPI.ClientCtx().Keybase.Sign(info.GetName(), password, accounts.TextHash(data))
	if err != nil {
		return nil, err
	}

	sig[crypto.RecoveryIDOffset] += 27 // transform V from 0/1 to 27/28
	return sig, nil
}

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be 27 or 28 for legacy reasons.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecove
func (api *PrivateAccountAPI) EcRecover(_ context.Context, data, sig hexutil.Bytes) (common.Address, error) {
	api.logger.Debug("personal_ecRecover", "data", data, "sig", sig)

	if len(sig) != crypto.SignatureLength {
		return common.Address{}, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	pubkey, err := crypto.SigToPub(accounts.TextHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}
