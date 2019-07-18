package chain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/storage/contract"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipweb-group/go-ipws-config"
)

const StorageDepositContractAddress = "0x0000000000000000000000000000000000000010"

const RetryTime = 30

func getStorageAccount(client *ethclient.Client, cfg *config.Config, fileAddress common.Address) (*contract.StorageAccount, error) {
	storageDepositContractAddr := common.HexToAddress(StorageDepositContractAddress)

	storageDeposit, err := contract.NewStorageDeposit(storageDepositContractAddr, client)
	if err != nil {
		return nil, err
	}

	storageAccountContractAddr, err := storageDeposit.GetStorageAccount(nil, fileAddress)
	if err != nil {
		return nil, err
	}

	return contract.NewStorageAccount(storageAccountContractAddr, client)
}

func newKeyedTransactor(cfg *config.Config) *bind.TransactOpts {
	privateKey, err := cfg.Chain.WalletKey()
	if err != nil {
		return nil
	}
	walletKey, _ := crypto.HexToECDSA(privateKey)
	return bind.NewKeyedTransactor(walletKey)
}

func CommitBlockInfo(cfg *config.Config, fHash string, bIdx uint32, bHash string, proof []byte) error {
	fileAddress := common.BytesToAddress(crypto.Keccak256([]byte(fHash)))

	client, err := ethclient.Dial(cfg.Chain.URL)
	if err != nil {
		return err
	}

	storageAccount, err := getStorageAccount(client, cfg, fileAddress)
	if err != nil {
		return err
	}

	retry := 0
	var tx *types.Transaction
	var e error

	for {
		if retry > RetryTime {
			break
		}
		e = nil

		auth := newKeyedTransactor(cfg)

		tx, e = storageAccount.CommitBlockInfo(auth, fileAddress, big.NewInt(int64(bIdx)), []byte(bHash), []byte(cfg.Identity.PeerID), proof)
		if e == nil {
			break
		}
		retry = retry + 1
		time.Sleep(time.Second)
	}
	if e == nil {
		wait(client, tx.Hash())
	}

	return e
}

func DownloadBlock(cfg *config.Config, fHash string, bIdx uint32, bHash string, proof []byte) error {
	fileAddress := common.BytesToAddress(crypto.Keccak256([]byte(fHash)))

	client, err := ethclient.Dial(cfg.Chain.URL)
	if err != nil {
		return err
	}

	storageAccount, err := getStorageAccount(client, cfg, fileAddress)
	if err != nil {
		return err
	}

	retry := 0
	var tx *types.Transaction
	var e error

	for {
		if retry > RetryTime {
			break
		}
		e = nil

		auth := newKeyedTransactor(cfg)

		tx, e = storageAccount.DownloadBlock(auth, fileAddress, big.NewInt(int64(bIdx)), []byte(bHash), []byte(cfg.Identity.PeerID), proof)
		if e == nil {
			break
		}
		retry = retry + 1
		time.Sleep(time.Second)
	}
	if e == nil {
		wait(client, tx.Hash())
	}

	return e
}

func wait(b *ethclient.Client, tx common.Hash) error {
	ctx := context.Background()
	for {
		receipt, err := b.TransactionReceipt(ctx, tx)
		if err != nil {
			if err != ethereum.NotFound {
				return err
			}
			time.Sleep(time.Second)
			continue
		}
		// TODO:update ethereum version
		if receipt.Status != 1 {
			return fmt.Errorf("tx %s status is failed", tx.String())
		}
		return nil
	}
	return nil
}

func Proof(data []byte, privKey string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(data)
	return crypto.Sign(hash, privateKey)
}
