package gorange

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/utils"
)

const PADDING = 20

func (n *Node) preAlloc(c Config) error {
	pv, err := n.WsProvider(context.Background())
	if err != nil {
		return err
	}
	defer pv.Close()

	faucet := pv.Accounts[0]

	_, tx, pa, err := DeployPreAlloc(faucet, pv)
	if err != nil {
		return err
	}
	if _, err = pv.WaitDeployedWithTimeout(tx, time.Minute); err != nil {
		return err
	}

	acl := len(c.Accounts)
	for i := 0; i < acl; i += PADDING {
		l := i + PADDING
		a := PADDING
		if acl < i+PADDING {
			l = acl
			a = l - i
		}

		tx, err = pa.Distribute(&bind.TransactOpts{
			From:    faucet.From,
			Context: pv.Context,
			Signer:  faucet.Signer,
			Value: new(big.Int).Mul(
				big.NewInt(int64(a)),
				utils.Ether(c.Balances),
			),
		}, c.Accounts[i:l], utils.Ether(c.Balances))
		if err != nil {
			return err
		}
		if receipt, err := pv.WaitMinedWithTimeout(tx, time.Minute); err != nil {
			return err
		} else {
			provider.PrintTxResult(tx, receipt)
		}
	}
	return nil
}
