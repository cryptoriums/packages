package hardhat

import (
	"context"
	"os"
	"path"
	"strings"

	contraget "github.com/cryptoriums/contraget/pkg/cli"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

func ReplaceContract(ctx context.Context, nodeURL string, contractPath string, contractName string, contractAddrToReplace common.Address) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}

	cfg := &contraget.Cli{
		Path:       contractPath,
		Name:       "contract",
		ObjectsDst: "tmp",
	}
	defer os.RemoveAll(cfg.DownloadDst)

	if err := contraget.Run(cfg); err != nil {
		return errors.Wrap(err, "generating the contract bin file")
	}

	_bin, err := os.ReadFile(path.Join(cfg.ObjectsDst, contractName+".bin"))
	if err != nil {
		return errors.Wrap(err, "reading the bin file")
	}
	bin := string(_bin)
	indexDeployBin := strings.LastIndex(bin, "60806040523480156")

	err = rpcClient.CallContext(ctx, nil, "hardhat_setCode", contractAddrToReplace, "0x"+bin[indexDeployBin:])
	if err != nil {
		return errors.Wrap(err, "hardhat_setCode call")
	}

	return nil
}
