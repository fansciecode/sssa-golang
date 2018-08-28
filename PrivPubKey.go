package sssa 
import {
	"crypto/sha256"
        "bytes"
	"fmt"
	"io"
        "github.com/tendermint/tendermint/crypto"
	amino "github.com/tendermint/go-amino"
	"github/tendermint/tendermint/crypto/secp256k1"
	"flag"
	"io/ioutil"
	"os"
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

/*
func (privKey PrivKeySssa) Sign(msg  []byte )([]byte ,error) {

	priv,_ := sssa.Create(2,3,privKey)
	sig,err :=priv.
	if err != nil {
		return nil ,err
	}

         return  sig.Serialize(),nil
}
*/

var _crypto.PrivKey = PrivKeySssa{}

type PrivKeySssa  [32]byte  

func (privKey PrivKeySssa) Bytes() []byte {
	return cdc.MustMarshalBinaryBare(privKey)
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
func GenPrivKeySssa()(secret []byte) PrivKeySssa {

	privKey32 := sha256.sum256(secret)
	return PrivKeySssa(privKey32)

      }
func splitSecret(secret []byte) []string {
		var minShares int
		var totalShares int

		fmt.Print("Insert Number of Slices: ")
		fmt.Scan(&totalShares)
		fmt.Print("Insert Minumim number of Slices to decrypt the secret: ")
		fmt.Scan(&minShares)
		if (minShares > totalShares) || (minShares <= 0) {
			fmt.Println("The minimum nuber of Slices must be lower than the total number of Slices")
			os.Exit(1)
								}

			fmt.Println("[+] Generating Slices")
		shares, err := sssa.Create(minShares, totalShares, string(secret))
		if err != nil {
		fmt.Println("Unable to create Shamir's Shares")
		os.Exit(1)
        } 																								return shares
}
func decrypt (){

	var n int
		var shares []string
		fmt.Print("Number of slices: ")
		fmt.Scan(&n)
		if n < 1 {
		fmt.Println("You can't inster a negative integer")
		os.Exit(1)
		}
		for i := 0; i < n; i++ {
		var share string
		fmt.Print("Insert share :")
		fmt.Scanln(&share)
		shares = append(shares, share)
				}																				secret, err := sssa.Combine(shares)
			if err != nil {																						fmt.Println("Error during Decryption")																			os.Exit(1)
	}
	fmt.Println(secret)
}
