package assets

import (
	"github.com/ethereum/go-ethereum/common"
)

var PolyDataFeeds = map[common.Address]common.Address{
	// AAVE/ETH
	common.HexToAddress("0x12528b3026cd252E8b4435E22f5615f7f3B8dA77"): common.HexToAddress("0xD6DF932A45C0f255f85145f286eA0b292B21C90B"),
	// BAL/ETH
	common.HexToAddress("0x7C66609Db07C8983f324098DDc53F03af716aba7"): common.HexToAddress("0x9a71012B13CA4d3D0Cdc72A177DF3ef03b0E76A3"),
	// CRV/ETH
	common.HexToAddress("0xC1670343d479ddEA6e90A108741b8Acc23aBe847"): common.HexToAddress("0x172370d5Cd63279eFa6d502DAB29171933a610AF"),
	// DAI/ETH
	common.HexToAddress("0x5777Ca61f29cac50250a3b136b52328D05DBa8C9"): common.HexToAddress("0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"),
	// DPI/ETH
	common.HexToAddress("0x02228Ab6bA5FfeE5ba15cd1477987FF4D4BfEce4"): common.HexToAddress("0x85955046DF4668e1DD369D2DE9f3AEB98DD2A369"),
	// GHST/ETH
	common.HexToAddress("0xec35e6f084ce365a819e99bcd1f89319e519fdf3"): common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
	// LINK/ETH
	common.HexToAddress("0xd2aa3fc2585999ef9ca66a1b6be18123b5774be7"): common.HexToAddress("0x53E0bca35eC356BD5ddDFebbD1Fc0fD03FaBad39"),
	// MATIC/ETH
	common.HexToAddress("0xc6d82423c6f8b0c406c1c34aee8e988b14d5f685"): common.HexToAddress("0x0000000000000000000000000000000000001010"),
	// MKR/ETH
	common.HexToAddress("0xa5043abeb607a370a16cdbd885fb7da6485a4e2c"): common.HexToAddress("0x6f7C932e7684666C9fd1d44527765433e01fF61d"),
	// QUICK/ETH
	common.HexToAddress("0x836faa493e68fac2dd6b9250ace9666fd48c4f09"): common.HexToAddress("0xf28164A485B0B2C90639E47b0f377b4a438a16B1"),
	// SUSHI/ETH
	common.HexToAddress("0x5826bdde4e50b2dc78f62103e921b3dcd14d4fd7"): common.HexToAddress("0x0b3F868E0BE5597D5DB7fEB59E1CADBb0fdDa50a"),
	// UNI/ETH
	common.HexToAddress("0x811ec390710bf071a585e32a465d06890a420937"): common.HexToAddress("0xb33EaAd8d922B1083446DC23f610c2567fB5180f"),
	// USDC/ETH
	common.HexToAddress("0xec5dc23f4fa6aac7fcbbcc2849571b04fcacd75f"): common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
	// USDT/ETH
	common.HexToAddress("0x40c9e6e3ba0324e58c0c88c78c8e733a93ac1b61"): common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
	// WBTC/ETH
	common.HexToAddress("0x921c2121af8d68f397cb2d2a3d24ffc180bcdcb7"): common.HexToAddress("0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6"),
	// YFI/ETH
	common.HexToAddress("0xf7c72effe3975c25252539685574a5f5bde19423"): common.HexToAddress("0xDA537104D6A5edd53c6fBba9A898708E465260b6"),

	// TEST UNI CONTRACT, SHOULD GET ALOT OF HITS.
	// common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"): common.HexToAddress("0x0000000000000000000000000000000000006969"),
	// common.HexToAddress("0x1618674520b125dbcD03B6dF25A20b47Aedbe3b1"): common.HexToAddress("0x0000000000000000000000000000000000006969"),
	// common.HexToAddress("0x226004a7C4CCe2f8649EfB936b32782a6B9665Af"): common.HexToAddress("0x0000000000000000000000000000000000006969"),
}
