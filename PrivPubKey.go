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
		SssaPrivKeyAminoRoute,nil)

}

func GenPrivKey () PriveKeySssa{

    
	return genPrivKey(crypto.CReader())
}
func genPrivKey(raw io.Reader) PrivKeySssa {

	 privKeyBytes  :=[32]byte {}
	 _,err := io.ReadFull (raw, privKeyBytes[:])
	 if err != nil {
		panic(err)
	 }

	 return PrivKeySssa(PrivKeyBytes)
}
