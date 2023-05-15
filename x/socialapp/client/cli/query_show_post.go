package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"socialapp/x/socialapp/types"
)

var _ = strconv.Itoa(0)

func CmdShowPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-post [profile-id] [id]",
		Short: "Query show-post",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqProfileId := args[0]
			reqId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowPostRequest{

				ProfileId: reqProfileId,
				Id:        reqId,
			}

			res, err := queryClient.ShowPost(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
