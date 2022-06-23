# Local Node

Hardhat or Anvil(Foundry) local nodes


```go
ctx := context.Background()
logger := logging.NewLogger()
nodeURL := os.GetEnv("NODE_URL")
blockNumber := "14675022"

ln, err := localnode.New(logger, localnode.Anvil, nodeURL, "13858047")
if err != nil {
    log.Fatal(err)
}
defer ln.Stop()

client, err := ethclient.DialContext(ctx, ln.GetNodeURL())
if err != nil {
    log.Fatal(err)
}

err = ln.ReplaceContract(
    ctx,
    "../testing/contracts/source/Booster.sol",
    "Booster",
    common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"),
)
if err != nil {
    log.Fatal(err)
}

```