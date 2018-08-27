package sssa 
import {
         "bytes"
	"fmt"
	"io"
        "github.com/tendermint/tendermint/crypto"
	amino "github.com/tendermint/go-amino"

        }
  
//-------------@Jhon stark-------


const  (
	SssaPrivKeyAminoRoute = "tendermint/PrivKeySssa"
	SssaPubKeyAminoRoute = "tendermint/PubKeySssa"
)
	var cdc =amino.NewCodec()

func init() {
            RegisterAmino(cdc)

}
func RegisterAmino(cdc *amino.cdc) {

        cdc.RegsiterInterface(*crypto.PubKey(nil),nil)
	cdc.RegisterConcrete(PubKeySssa{},
		SssaPubKeyAminoRoute,nil)

	cdc.RegisterInterface(*crypto.PrivKey(nil),nil)
	cdc.RegisterConcrete(PrivKeySssa{},
		SssaPrivKeyAminoRoutr,nil)
 


}
