package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yanue/starkex"
)

//var baseURL = "https://api.dydx.exchange"
var baseURL = "https://api.stage.dydx.exchange/"

/*


dYdX
https://trade.dydx.exchange
0x87016b...b1734c12
Message
action:
dYdX STARK Key
onlySignOn:
https://trade.dydx.exchange


dYdX
https://trade.dydx.exchange
0x87016b...b1734c12
Message
action:
dYdX Onboarding
onlySignOn:
https://trade.dydx.exchange

Request	Limit
GET v3/*	100 requests per 10 seconds.
PUT v3/emails/send-verification-email	2 requests for 10 minutes.
DELETE v3/orders	See Cancel-Order Rate Limits
POST v3/orders	See Place-Order Rate-Limits
POST v3/testnet/tokens	5 requests per 24 hours.
GET v3/active-orders	See Active-Order Rate-Limits
DELETE v3/active-orders	See Active-Order Rate-Limits
All other requests	10 requests per minute.
*/

var RequestTestToken = "v3/testnet/tokens"

var recipient = "0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E"
var stark_key = "7530be28e4ec85e0be0f4a864dbaa0b6078ad3de4f4531b3dc4e2d5294e8d4a"

func main() {
	//request()
	//post()

	// const MOCK_PRIVATE_KEY = "67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3"  //0x4F211267896C4D3f2388025263AC6BD67B0f2C54
	// const NETWORK_ID_ROPSTEN = 3
	orderSign()
	//transferSign()
	//	testToken()
}
func testToken() {
	url := baseURL + RequestTestToken

	postBody, _ := json.Marshal(map[string]string{
		"ethereum_address": "0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00",
	})
	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}

func transferSign() {
	MOCK_PRIVATE_KEY := stark_key
	param := starkex.TransferSignParam{
		NetworkId:          starkex.NETWORK_ID_ROPSTEN,
		CreditAmount:       "1",
		DebitAmount:        "2",
		SenderPositionId:   12345,
		ReceiverPositionId: 67890,
		ReceiverPublicKey:  "04a9ecd28a67407c3cff8937f329ca24fd631b1d9ca2b9f2df47c7ebf72bf0b0",
		ReceiverAddress:    recipient,
		Expiration:         "2022-09-17T04:15:55.028Z",
		ClientId:           "This is an ID that the client came up with to describe this transfer",
	}
	sign, err := starkex.TransferSign(MOCK_PRIVATE_KEY, param)
	// 0278aeb361938d4c377950487bb770fc9464bf5352e19117c03243efad4e10a302bb3983e05676c7952caa4acdc1a83426d5c8cb8c56d7f6c477cfdafd37718a
	fmt.Println("sign,err", sign, err)
}
func orderSign() {
	MOCK_PRIVATE_KEY := stark_key

	param := starkex.OrderSignParam{
		NetworkId:  starkex.NETWORK_ID_ROPSTEN,
		Market:     "ETH-USD",
		Side:       "BUY",
		PositionId: 12345,
		HumanSize:  "145.0005",
		HumanPrice: "350.00067",
		LimitFee:   "0.125",
		ClientId:   "This is an ID that the client came up with to describe this order",
		Expiration: "2022-09-17T04:15:55.028Z",
	}
	sign, err := starkex.OrderSign(MOCK_PRIVATE_KEY, param)
	// 00cecbe513ecdbf782cd02b2a5efb03e58d5f63d15f2b840e9bc0029af04e8dd0090b822b16f50b2120e4ea9852b340f7936ff6069d02acca02f2ed03029ace5
	fmt.Println("sign,err", sign, err)

}
func request() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func post() {

	url := "https://trade.dydx.exchange/v3/api-keys"
	//url := "https://trade.dydx.exchange/v3/testnet/tokens"

	postBody, _ := json.Marshal(map[string]string{
		"ethereum_address": "0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00",
	})
	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
