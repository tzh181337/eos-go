package contract_call

import (
	"context"
	"fmt"
	"github.com/armoniax/eos-go"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

const (
	api = "https://test-chain.ambt.art"
)

type TokenStats struct {
	Supply struct {
		Amount int `json:"amount"`
		Symbol struct {
			Id       int64 `json:"id"`
			ParentId int   `json:"parent_id"`
		} `json:"symbol"`
	} `json:"supply"`
	MaxSupply struct {
		Amount int `json:"amount"`
		Symbol struct {
			Id       int64 `json:"id"`
			ParentId int   `json:"parent_id"`
		} `json:"symbol"`
	} `json:"max_supply"`
	TokenUri    string `json:"token_uri"`
	Ipowner     string `json:"ipowner"`
	Notary      string `json:"notary"`
	Issuer      string `json:"issuer"`
	IssuedAt    string `json:"issued_at"`
	NotarizedAt string `json:"notarized_at"`
	Paused      int    `json:"paused"`
}

// 卖家出价表
type SellOrders struct {
	ID    uint64 `json:"id"`
	Sn    uint64 `json:"sn"`
	Price struct {
		Value  string `json:"value"`
		Symbol struct {
			Id       int64 `json:"id"`
			ParentId int   `json:"parent_id"`
		} `json:"symbol"`
	} `json:"price"`
	Frozen    int64  `json:"frozen"`
	Maker     string `json:"maker"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 买家议价表
type BuyerBids struct {
	ID          uint64 `json:"id"`
	SellOrderId uint64 `json:"sell_order_id"`
	Price       struct {
		Value  string `json:"value"`
		Symbol struct {
			Id       int64 `json:"id"`
			ParentId int   `json:"parent_id"`
		} `json:"symbol"`
	} `json:"price"`
	Frozen    string `json:"frozen"`
	Buyer     string `json:"buyer"`
	CreatedAt string `json:"created_at"`
}

type MartGlobal struct {
	Rows []struct {
		Admin            string `json:"admin"`
		DevFeeCollector  string `json:"dev_fee_collector"`
		DevFeeRate       string `json:"dev_fee_rate"`
		CreatorFeeRate   string `json:"creator_fee_rate"`
		IpownerFeeRate   string `json:"ipowner_fee_rate"`
		NotaryFeeRate    string `json:"notary_fee_rate"`
		OrderExpiryHours int    `json:"order_expiry_hours"`
	} `json:"rows"`
	More    bool   `json:"more"`
	NextKey string `json:"next_key"`
}

// https://github.com/armoniax/nftone.contracts/blob/main/contracts/nftone.mart/include/nftone.mart/nftone.mart.db.hpp
func TestQueryBuyerBids(t *testing.T) {
	client := eos.New(api)
	ctx := context.Background()

	buyer := "amaxpro11111"
	orderId := uint64(16)

	// 如果是组合索引查询，要计算
	lowerBound, upperBound, err := getQuery(buyer, orderId)
	if err != nil {
		fmt.Printf("get table %s rows error: %s", "buyerbids", err.Error())
		return
	}

	// 根据sell_order_id + buyer +created_at 组合索引查询 buyerbids 表
	params := eos.GetTableRowsRequest{
		Code:  "nftone.mart",
		Scope: "nftone.mart",
		Table: "buyerbids",
		//LowerBound: "0x000000000000000392ae869a79e8000000000000000000000000000000000000",
		//UpperBound: "0x000000000000000392ae869a79e80000ffffffffffffffffffffffffffffffff",
		LowerBound: lowerBound,
		UpperBound: upperBound,
		Limit:      10,
		KeyType:    "sha256",
		Index:      "4",
		EncodeType: "i256",
		Reverse:    true,
		JSON:       true,
	}

	var tableRows *eos.GetTableRowsResp
	// test
	for i := 0; i < 4; i++ {
		fmt.Printf("query loop i: %d\n", i)
		tableRows, err = client.GetTableRows(ctx, params)
		if err != nil {
			fmt.Printf("get table %s rows error: %s", "buyerbids", err.Error())
		}

		if err == nil && len(tableRows.Rows) > 0 {
			break
		}
	}

	var buyerBids []BuyerBids
	err = tableRows.JSONToStructs(&buyerBids)
	if err != nil {
		fmt.Printf("json table rows err: %s\n", err.Error())
		return
	}

	fmt.Printf("buyerBids: %+v\n", buyerBids)
}

func getQuery(name string, orderId uint64) (string, string, error) {
	actual, err := eos.ExtendedStringToName(name)
	if err != nil {
		fmt.Printf("string to name error: %s", err.Error())
		return "", "", err
	}

	binName := strconv.FormatUint(actual, 2)

	var nameBuilder strings.Builder
	if len(binName) < 64 {
		for i := 0; i < 64-len(binName); i++ {
			nameBuilder.WriteString("0")
		}
	}

	binNames := fmt.Sprintf("%s%s", nameBuilder.String(), binName)

	var idBuilder strings.Builder
	ids := strconv.FormatUint(orderId, 2)
	if len(ids) < 64 {
		for i := 0; i < 64-len(ids); i++ {
			idBuilder.WriteString("0")
		}
	}

	//binOrderId := idBuilder.String()
	binOrderId := fmt.Sprintf("%s%s", idBuilder.String(), ids)

	var sumBuilder strings.Builder
	sumBuilder.WriteString(binOrderId)
	sumBuilder.WriteString(binNames)

	sumBin := sumBuilder.String()

	sumBig, _ := new(big.Int).SetString(sumBin, 2)
	hexSum := toHexInt(sumBig)

	//
	var prefixSum strings.Builder
	if len(hexSum) < 32 {
		for i := 0; i < 32-len(hexSum); i++ {
			prefixSum.WriteString("0")
		}
	}

	var suffixZero strings.Builder
	var suffixf strings.Builder
	for i := 0; i < 32; i++ {
		suffixZero.WriteString("0")
	}

	for i := 0; i < 32; i++ {
		suffixf.WriteString("f")
	}

	var lower strings.Builder
	var upper strings.Builder

	//lower.WriteString("0x")
	lower.WriteString(prefixSum.String())
	lower.WriteString(hexSum)
	lower.WriteString(suffixZero.String())

	//upper.WriteString("0x")
	upper.WriteString(prefixSum.String())
	upper.WriteString(hexSum)
	upper.WriteString(suffixf.String())

	lowerBound := lower.String()
	upperBound := upper.String()
	return lowerBound, upperBound, nil
}

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func TestQuerySellOrders(t *testing.T) {
	client := eos.New(api)
	ctx := context.Background()

	// uint128_t by_maker_created_at()const { return (uint128_t) maker.value << 64 | (uint128_t) created_at.sec_since_epoch(); }
	// 根据maker+created_at 组合索引查询 sellOrder订单
	//amaxamaxmara

	maker := "yanjingxun33"
	actual, err := eos.ExtendedStringToName(maker)
	if err != nil {
		fmt.Printf("string to name error: %s", err.Error())
		return
	}

	binName := strconv.FormatUint(actual, 2)

	var nameBuilder strings.Builder
	if len(binName) < 16 {
		for i := 0; i < 16-len(binName); i++ {
			nameBuilder.WriteString("0")
		}
	}

	binNames := fmt.Sprintf("%s%s", nameBuilder.String(), binName)

	nameBig, _ := new(big.Int).SetString(binNames, 2)
	hexName := toHexInt(nameBig)

	fmt.Printf("hexName:===== %s\n", hexName)
	fmt.Printf("len name:===== %d\n", len(hexName)) // 16

	var suffixZero strings.Builder
	var suffixf strings.Builder
	for i := 0; i < 16; i++ {
		suffixZero.WriteString("0")
	}

	for i := 0; i < 16; i++ {
		suffixf.WriteString("f")
	}

	var lower strings.Builder
	var upper strings.Builder

	lower.WriteString("0x")
	lower.WriteString(hexName)
	lower.WriteString(suffixZero.String())

	upper.WriteString("0x")
	upper.WriteString(hexName)
	upper.WriteString(suffixf.String())

	lowerBound := lower.String()
	upperBound := upper.String()
	//lowerBound := "0xcab19d61570800000000000000000000"
	//upperBound := "0xcab19d6157080000ffffffffffffffff"

	params := eos.GetTableRowsRequest{
		//Code:       "martz", //测试
		Code:       "nftone.mart", //线上
		Scope:      "2852126739",  // token_id
		Table:      "sellorders",
		LowerBound: lowerBound,
		UpperBound: upperBound,
		Limit:      100,
		KeyType:    "i128",
		Index:      "3",
		EncodeType: "i128",
		Reverse:    true,
		JSON:       true,
	}

	var tableRows *eos.GetTableRowsResp
	for i := 0; i < 4; i++ {
		fmt.Printf("query loop i: %d\n", i)
		tableRows, err := client.GetTableRows(ctx, params)
		if err != nil {
			fmt.Printf("get table %s rows error: %s", "sellorders", err.Error())
		}

		if err == nil && len(tableRows.Rows) > 2 {
			break
		}
	}

	var sellOrders []SellOrders

	if tableRows != nil {
		err := tableRows.JSONToStructs(&sellOrders)
		if err != nil {
			fmt.Printf("json table rows err: %s\n", err.Error())
			return
		}

		fmt.Printf("orders: %+v\n", sellOrders)
	}
}

// 根据
func TestQueryAllOrders(t *testing.T) {
	client := eos.New(api)
	ctx := context.Background()

	params := eos.GetTableRowsRequest{
		Code:  "martf",
		Scope: "989856054",
		Table: "sellorders",
		//LowerBound: lowerBound,
		//UpperBound: upperBound,
		Limit: 1000,
		//KeyType:    "i128",
		//Index:      "3",
		//EncodeType: "i128",
		Reverse: false,
		JSON:    true,
	}

	var tableRows *eos.GetTableRowsResp
	var err error
	tableRows, err = retryQuery(ctx, 4, client, params)
	if err != nil {
		fmt.Printf("retury get table %s rows error: %s", params.Table, err.Error())
		return
	}

	var sellOrders []SellOrders

	if tableRows != nil && len(tableRows.Rows) > 2 {
		err := tableRows.JSONToStructs(&sellOrders)
		if err != nil {
			fmt.Printf("json table rows err: %s\n", err.Error())
			return
		}

		fmt.Printf("orders: %+v\n", sellOrders)
	}
}

func TestQueryToken(t *testing.T) {
	client := eos.New(api)
	ctx := context.Background()

	params := eos.GetTableRowsRequest{
		Code:       "versontoken2",
		Scope:      "versontoken2", // token_id
		Table:      "tokenstats",
		LowerBound: "6600014",
		//UpperBound: upperBound,
		Limit: 1,
		//KeyType:    "i128",
		//Index:      "3",
		//EncodeType: "i128",
		Reverse: false,
		JSON:    true,
	}

	var tableRows *eos.GetTableRowsResp
	var err error
	tableRows, err = retryQuery(ctx, 4, client, params)
	if err != nil {
		fmt.Printf("retury get table %s rows error: %s", params.Table, err.Error())
		return
	}

	if tableRows != nil {
		var tokenstats []TokenStats
		err = tableRows.JSONToStructs(&tokenstats)
		if err != nil {
			fmt.Printf("json table rows err: %s\n", err.Error())
			return
		}

		fmt.Printf("tokenstats: %+v\n", tokenstats)
	}
}

func TestQueryAccount(t *testing.T) {
	client := eos.New(api)
	ctx := context.Background()

	params := eos.GetTableRowsRequest{
		Code:  "amax.ntoken",
		Scope: "merchantx", // token_id
		Table: "accounts",
		//LowerBound: "603980154",
		//UpperBound: upperBound,
		Limit: 100,
		//KeyType:    "i128",
		//Index:      "3",
		//EncodeType: "i128",
		Reverse: true,
		JSON:    true,
	}

	var tableRows *eos.GetTableRowsResp
	var err error
	tableRows, err = retryQuery(ctx, 4, client, params)
	if err != nil {
		fmt.Printf("retury get table %s rows error: %s", params.Table, err.Error())
		return
	}

	if tableRows != nil {
		var tokenstats []TokenStats
		err = tableRows.JSONToStructs(&tokenstats)
		if err != nil {
			fmt.Printf("json table rows err: %s\n", err.Error())
			return
		}

		fmt.Printf("tokenstats: %+v\n", tokenstats)
	}
}

func TestQueryTSellOrder(t *testing.T) {
	api := "http://hk-t1.nchain.me:18888"
	client := eos.New(api)
	ctx := context.Background()

	var order SellOrders

	// 根据maker+created_at 组合索引查询 sellOrder订单
	params := eos.GetTableRowsRequest{
		Code:       "martf",
		Scope:      "989856054", // token_id
		Table:      "sellorders",
		LowerBound: "132",
		Index:      "1",
		Limit:      1,
		Reverse:    false,
		JSON:       true,
	}

	// 这里先加上重试4次，后面把同步order写到定时任务里
	var tableRows *eos.GetTableRowsResp
	var err error
	tableRows, err = retryQuery(ctx, 4, client, params)
	if err != nil {
		fmt.Printf("retury get table %s rows error: %s", params.Table, err.Error())
		return
	}

	if len(tableRows.Rows) == 0 {
		fmt.Printf("get table %s rows error, res is nil\n", params.Table)
		return
	}

	var orders []SellOrders
	err = tableRows.JSONToStructs(&orders)
	if err != nil {
		fmt.Printf("json table rows err: %s\n", err.Error())
		return
	}

	if len(orders) >= 1 {
		order = orders[0]
	}

	fmt.Printf("one sellorders: %+v\n", order)
}

func retryQuery(ctx context.Context, num int, client *eos.API, params eos.GetTableRowsRequest) (*eos.GetTableRowsResp, error) {
	var tableRows *eos.GetTableRowsResp
	var err error
	for i := 0; i < num; i++ {
		fmt.Printf("query table: %s, loop i: %d\n", params.Table, i)
		tableRows, err = client.GetTableRows(ctx, params)
		if err != nil {
			fmt.Printf("get table %s rows error: %s", params.Table, err.Error())
		}

		if err == nil && len(tableRows.Rows) > 2 {
			break
		}
	}

	return tableRows, nil
}
