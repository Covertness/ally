package ethclient

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	client "github.com/ethereum/go-ethereum/ethclient"

	"github.com/Covertness/ally/contracts"
	"github.com/Covertness/ally/pkg/hdwallet"
)

var ins *Client

// Client the instance content
type Client struct {
	*client.Client
	myAddressFactory *contracts.AddressFactory
	myERC20Token     *contracts.ERC20
}

// Init create the eth client with infura url
func Init(infuraURL string) (*Client, error) {
	my, err := New(infuraURL)
	ins = &Client{
		Client: my,
	}
	return ins, err
}

// GetInstance get the instance
func GetInstance() *Client {
	return ins
}

// New create a new internal instance
func New(infuraURL string) (*client.Client, error) {
	return client.Dial(infuraURL)
}

// GenNonce generate the transaction nonce
func (c *Client) GenNonce() (*big.Int, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	fromAddress := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	mNonce := big.NewInt(int64(nonce))
	return mNonce, nil
}

// GetSigner get the transaction signer
func (c *Client) GetSigner(nonce *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nonce
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5500000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func getPrivateKey() (*ecdsa.PrivateKey, error) {
	root, err := hdwallet.RootAddress()
	if err != nil {
		return nil, err
	}

	return hdwallet.AddressPrivateKey(root)
}
