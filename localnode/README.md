# Local Node

Hardhat or Anvil(Foundry) local nodes


```go
ctx := context.Background()
logger := logging.NewLogger()
nodeURL := os.GetEnv("NODE_URL")
blockNumber := "14675022"

cmd := localnode.Fork(logger, "npx", "hardhat", "node", "--fork", nodeURL, "--fork-block-number", blockNumber)
defer testutil.KillCmd(t, cmd)

client, err := client.NewClientCachedNetID(ctx, logger, localnode.DefaultUrl)
if err != nil {
    log.Fatal(err)
}

err = ReplaceContract(
    ctx,
    localnode.Hardhat,
    localnode.DefaultUrl,
    "../testing/contracts/source/Booster.sol",
    "Booster",
    common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"),
)

if err != nil {
    log.Fatal(err)
}

```