package cmd

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gjermundgaraba/gon/chains"
	"github.com/spf13/cobra"
)

func transferNFTInteractive(cmd *cobra.Command) error {
	src, _ := cmd.Flags().GetString(flagSrcChainID)
	var sourceChain chains.Chain
	for _, chain := range chains.Chains {
		if string(chain.ChainID()) == src {
			sourceChain = chain
		}
	}

	setAddressPrefixes(sourceChain.Bech32Prefix())

	key := chooseOrCreateKey(cmd, sourceChain)
	if err := cmd.Flags().Set(flags.FlagFrom, key); err != nil {
		panic(err)
	}

	clientCtx := getClientTxContext(cmd, sourceChain)
	fromAddress := getAddressForChain(clientCtx, sourceChain, key)

	dst, _ := cmd.Flags().GetString(flagDstChainID)
	var destinationChain chains.Chain
	for _, chain := range chains.Chains {
		if string(chain.ChainID()) == dst {
			destinationChain = chain
		}
	}

	class, _ := cmd.Flags().GetString(flagNFTClassID)
	selectedClass := getUsersNfts(cmd.Context(), clientCtx, sourceChain, fromAddress, class)
	fmt.Println("Class ID: ", selectedClass)
	if len(selectedClass.NFTs) == 0 {
		fmt.Println("No NFT classes found")
		return nil
	}

	var selectedNFT chains.NFT
	nft, _ := cmd.Flags().GetString(flagNFTID)
	for _, cNft := range selectedClass.NFTs {
		if strings.ToLower(cNft.ID) == strings.ToLower(nft) {
			selectedNFT = cNft
		}
	}

	destinationAddress, _ := cmd.Flags().GetString(flagDestAddress)
	if destinationAddress == "" {
		destinationAddress = getAddressForChain(clientCtx, destinationChain, key)
		fmt.Println("Destination address:", destinationAddress)
	}

	channelId, _ := cmd.Flags().GetString(flagChannelID)
	var chosenChannel chains.NFTChannel
	var chosenConnection chains.NFTConnection
	connections := sourceChain.GetConnectionsTo(destinationChain)
	for _, connection := range connections {
		if connection.ChannelA.Label() == channelId {
			chosenChannel = connection.ChannelA
			chosenConnection = connection
		}
	}

	tryToForceTimeout, _ := cmd.Flags().GetBool(flagTryToForceTimeout)
	targetChainHeight, targetChainTimestamp := getCurrentChainStatus(cmd.Context(), getQueryClientContext(cmd, destinationChain))
	timeoutHeight, timeoutTimestamp := sourceChain.GetIBCTimeouts(clientCtx, chosenChannel.Port, chosenChannel.Channel, targetChainHeight, targetChainTimestamp, tryToForceTimeout)

	msg := sourceChain.CreateTransferNFTMsg(chosenChannel, selectedClass, selectedNFT, fromAddress, destinationAddress, timeoutHeight, timeoutTimestamp)
	if tryToForceTimeout {
		clientCtx = clientCtx.WithSkipConfirmation(true)
	}

	txResponse, err := sendTX(clientCtx, cmd.Flags(), msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("The destination ibc trace (full Class ID on destination chain) will be:")
	expectedDestinationClass, isRewind := calculateClassTrace(selectedClass.FullPathClassID, chosenConnection)
	if isRewind {
		fmt.Println("(This is a rewind transaction)")
	}
	fmt.Println(expectedDestinationClass)

	if len(strings.Split(expectedDestinationClass, "/")) > 2 && destinationChain.NFTImplementation() == chains.CosmosSDK {
		fmt.Println()
		fmt.Println("Class hash:")
		fmt.Println(calculateClassHash(expectedDestinationClass))
	}

	fmt.Println()
	selfRelay, _ := cmd.Flags().GetBool(flagSelfRelay)
	waitAndPrintIBCTrail(cmd, sourceChain, destinationChain, txResponse.TxHash, selfRelay)

	fmt.Println()
	fmt.Println("The destination ibc trace (full Class ID on destination chain):")
	if isRewind {
		fmt.Println("(This is a rewind transaction)")
	}
	fmt.Println(expectedDestinationClass)

	return nil
}
