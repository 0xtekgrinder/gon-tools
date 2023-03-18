CLASS="tekgrinderNFTs"
NFT="token60"

ark transfer ics721 --from $WALLET_MINTER --recipient stars183e7ccwsnngj2q8lfxnmekunspnfxs6q8nzqcf --target-chain stargaze --source-channel $CHANNEL_1_TO_STARGAZE --collection $ARK_GON_COLLECTION --token arkalpha002
gon transfer --self-relay  --from gon --src gon-irishub-1 --dst elgafar-1 --class-id $CLASS --nft-id $NFT --channel-id nft-transfer/channel-22 -y
gon transfer --self-relay  --from gon --src elgafar-1 --dst uni-6 --class-id wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/$CLASS --nft-id $NFT --channel-id wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-230 -y
gon transfer --self-relay  --from gon --src uni-6 --dst uptick_7000-2 --class-id wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/$CLASS --nft-id $NFT --channel-id wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-86 -y
gon transfer --self-relay  --from gon --src uptick_7000-2 --dst gon-flixnet-1 --class-id nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/$CLASS --nft-id $NFT --channel-id nft-transfer/channel-9 -y
gon transfer --self-relay  --from gon --src gon-flixnet-1 --dst gon-irishub-1 --class-id nft-transfer/channel-42/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/$CLASS --nft-id $NFT --channel-id nft-transfer/channel-24 -y

IBC=$(iris query nft-transfer class-hash nft-transfer/channel-0/nft-transfer/channel-42/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/$CLASS --node http://34.80.93.133:26657/ --chain-id gon-irishub-1 | awk '{print $2}')
iris tx nft transfer iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz ibc/$IBC $NFT --from gon --node http://34.80.93.133:26657 --chain-id gon-irishub-1 --fees 2000uiris -y