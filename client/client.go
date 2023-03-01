package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"poginteractor/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	privateKey string

	cli      *ethclient.Client
	instance *contract.POGToken
}

func NewClient(infuraURL, rawPrivateKey, contractAddress string) (*Client, error) {
	if infuraURL == "" {
		return nil, errors.New("missed infra url")
	}
	if rawPrivateKey == "" {
		return nil, errors.New("missed private key")
	}

	cli, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, fmt.Errorf("dial failed: %w", err)
	}

	var instance *contract.POGToken
	if contractAddress != "" {
		instance, err = contract.NewPOGToken(common.HexToAddress(contractAddress), cli)
		if err != nil {
			return nil, fmt.Errorf("failed to create contract instance: %w", err)
		}
	}

	return &Client{
		privateKey: rawPrivateKey,
		cli:        cli,
		instance:   instance,
	}, nil
}

func (c *Client) prepareAuth() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(c.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey: %w", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// estimate gas price
	gasPrice, err := c.cli.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	opts := bind.NewKeyedTransactor(privateKey)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei
	opts.GasLimit = 0
	opts.GasPrice = gasPrice

	return opts, nil
}

func (c *Client) Deploy(supply *big.Int) (common.Address, *types.Transaction, error) {
	opts, err := c.prepareAuth()
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("auth failed: %w", err)
	}

	addr, tx, _, err := contract.DeployPOGToken(opts, c.cli, supply)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to deploy contract: %w", err)
	}

	return addr, tx, nil
}

func (c *Client) AddViewer(address string, username string) (*types.Transaction, error) {
	opts, err := c.prepareAuth()
	if err != nil {
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	tx, err := c.instance.AddViewer(opts, common.HexToAddress(address), username, true)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract method: %w", err)
	}

	return tx, nil
}

func (c *Client) RemoveViewer(address string) (*types.Transaction, error) {
	opts, err := c.prepareAuth()
	if err != nil {
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	tx, err := c.instance.RemoveViewer(opts, common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("failed to call contract method: %w", err)
	}

	return tx, nil
}

func (c *Client) PrintLogsAdd() error {
	tx, err := c.instance.FilterLogViewerAdded(&bind.FilterOpts{})
	if err != nil {
		return fmt.Errorf("failed to initialize filter: %w", err)
	}

	for tx.Next() {
		event := tx.Event

		fmt.Printf("[%s] New viewer added:\n - address:  %s\n - username:  %s\n\n", tx.Event.Raw.Address, event.Addr, event.Username)
	}

	return tx.Error()
}

func (c *Client) PrintLogsRemove() error {
	tx, err := c.instance.FilterLogViewerRemoved(&bind.FilterOpts{})
	if err != nil {
		return fmt.Errorf("failed to initialize filter: %w", err)
	}

	for tx.Next() {
		event := tx.Event

		fmt.Printf("[%s] A viewer has been removed:\n - address:  %s\n - username:  %s\n\n", tx.Event.Raw.Address, event.Addr, event.Username)
	}

	return tx.Error()
}
