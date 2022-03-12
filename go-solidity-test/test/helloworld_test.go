package api

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"
	_ "time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"

	hwdApi "go-solc-test/m/go-deploy/helloworld"
)

type HelloworldTestSuite struct {
	suite.Suite
	auth       *bind.TransactOpts
	address    common.Address
	gAlloc     core.GenesisAlloc
	sim        *backends.SimulatedBackend
	helloworld *hwdApi.Helloworld
}

func TestRunHelloworldSuite(t *testing.T) {
	suite.Run(t, new(HelloworldTestSuite))
}

func (s *HelloworldTestSuite) SetupTest() {
	key, _ := crypto.GenerateKey()
	s.auth = bind.NewKeyedTransactor(key)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	s.address = s.auth.From
	fmt.Println("s.address:", s.address)

	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: balance},
	}

	blockGasLimit := uint64(4712388)
	s.sim = backends.NewSimulatedBackend(s.gAlloc, blockGasLimit)

	gasPrice, err := s.sim.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasprice:", gasPrice)

	s.auth.Value = big.NewInt(0)
	s.auth.GasLimit = uint64(557533)
	s.auth.GasPrice = gasPrice

	_, _, hw, e := hwdApi.DeployHelloworld(s.auth, s.sim)
	s.helloworld = hw
	s.Nil(e)
	s.sim.Commit()
}

func (s *HelloworldTestSuite) TestSay() {
	str, err := s.helloworld.Say(nil)
	s.Equal("hello etherworld", str)
	s.Nil(err)
}

func (s *HelloworldTestSuite) TestDeposit() {
	tx, err := s.helloworld.Deposit(s.auth, big.NewInt(10000))
	//s.Equal("hello etherworld", str)

	fmt.Println("deposit tx:", tx.Hash().Hex())
	fmt.Println("s.address:", s.address)
	fmt.Println("s.auth.From:", s.auth.From)
	s.Nil(err)
}

// func (s *HelloworldTestSuite) TestCheckBalance() {

// 	amount, err := s.helloworld.CheckBalance(&bind.CallOpts{}, s.auth.From)
// 	s.Equal(amount, big.NewInt(10000))
// 	s.Nil(err)
// }
