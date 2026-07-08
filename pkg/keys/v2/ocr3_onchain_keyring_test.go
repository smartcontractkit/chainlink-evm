package keys

import (
	"context"
	"encoding/hex"
	"errors"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/keystore"
)

var _ keystore.Storage = &TestMemoryStorage{}

/*
TestMemoryStorage implements a keystore storage prepopulated with the test ECDSA keys (used in test vectors):

	"0000000000000000000000000000000000000000000000000000000000000001",
	"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
	"4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
	"c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
	"f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
	"e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
	"69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",

The keys are accessible by their name using the first 4 characters ("0000", "ffff", "4c08", "c875", "f8b8", "e916", "69ec").
The storage is encrypted using the empty password and DefaultScryptParams.
*/
type TestMemoryStorage struct {
}

func (t TestMemoryStorage) GetEncryptedKeystore(_ context.Context) ([]byte, error) {
	const KS = "7b22636970686572223a226165732d3132382d637472222c2263697068657274657874223a22643233623331366430356531333339393632313961393564396634323437353164356435326561623365383339393363373535383835396662323964653539343936373063666334316431663136653638353364396436663665626233333566363661663035366536333138663533663039356532303566656435316162393935663932633830393664633838333563383361626263623137613030376137366430343536366138363639353137653065633031373437313365643562303362653166336330306537663166323266333562616339336534663434326135373730643037343533633666363265646337333335616232333539346264356537373137643163613765323364643766386138346331343363616435323164373631316637326536343362323632376632326263666266353164656362643663643434386264643639396361366638343836336539346530373430383965666332313364393034356239616236303334663438383833386236323436653366306637333563363232303131353433393932613964623534303939333634303931383066616331376264303764353930336537333234613632646438666532306564366662323637313466653639333932646463353039336563383131646566373430393961653837336361373539363530653362346535306433353433636164396263643036386433613063643963643466313432626233336334343262316331656363636431316533656634323939376362393535366661383036346534353431333331646661373330353166383761303865393661303938316337653362653965366535396566366433393337313863626533346562643832386130623965363834353532343165616165633963366239613764303764373438383962313233613665366437353662376131346364396435653638386537353661383535333536646163653964363165663730343034346566376535373234653637326132383338366334653866316635376465303736663263663564393163373263356261333534653133343561636464336464653738336666336564346433643337323362623261633866373165343138383434363832323065656436393834643232306366626537346366222c22636970686572706172616d73223a7b226976223a223230383936323338633466613965666261366665633435343337396164323963227d2c226b6466223a22736372797074222c226b6466706172616d73223a7b22646b6c656e223a33322c226e223a3236323134342c2270223a312c2272223a382c2273616c74223a2232346135633936323235656335626161613666323032393366376465613864393131333630333536343561336432396264366533353635343637343864646134227d2c226d6163223a2261666433333165626565626562656166373465373765303866663437636461663161653962653934623836633137386164316236663733353534303636313639227d"
	return hex.DecodeString(KS)
}

func (t TestMemoryStorage) PutEncryptedKeystore(_ context.Context, _ []byte) error {
	return errors.New("unimplemented")
}

func TestSignVerifyTestVectorsOnchain(t *testing.T) {
	ks, err := keystore.LoadKeystore(t.Context(), TestMemoryStorage{}, "")
	require.NoError(t, err)
	for _, tv := range slices.Concat(ecdsaTestVectors, ecdsaNegativeTestVectors) {
		resp, err := ks.GetKeys(t.Context(), keystore.GetKeysRequest{KeyNames: []string{tv.signingKey[:4]}})
		require.NoError(t, err)
		require.Len(t, resp.Keys, 1)
		publicKey, err := crypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
		require.NoError(t, err)
		addr := crypto.PubkeyToAddress(*publicKey)
		kr := evmOnchainKeyring2{
			ks:      ks,
			keyPath: keystore.KeyPath{tv.signingKey[:4]},
			addr:    addr,
		}
		configDigestBytes, err := hex.DecodeString(tv.configDigest)
		require.NoError(t, err)
		configDigest := ocrtypes.ConfigDigest(configDigestBytes)
		reportBytes, err := hex.DecodeString(tv.report)
		require.NoError(t, err)
		report := ocr3types.ReportWithInfo[struct{}]{
			Report: reportBytes,
		}
		sig, err := kr.Sign(configDigest, tv.seqNr, report.Report)
		require.NoError(t, err)
		addressBytes, err := hex.DecodeString(tv.address)
		require.NoError(t, err)
		require.True(t, kr.Has(addressBytes))
		sigBytes, err := hex.DecodeString(tv.sig)
		require.NoError(t, err)
		if tv.sigValid {
			require.Equal(t, sig, sigBytes)
			require.True(t, kr.Verify(addressBytes, configDigest, tv.seqNr, report.Report, sigBytes))
		} else {
			require.False(t, kr.Verify(addressBytes, configDigest, tv.seqNr, report.Report, sigBytes))
		}
	}
}

type ecdsaTestVector struct {
	desc         string
	configDigest string
	seqNr        uint64
	report       string
	reportHash   string
	signingKey   string
	address      string
	r            string
	s            string
	v            byte
	sig          string
	sigValid     bool
}

// Test vectors have been generated in offchain-reporting
// https://github.com/smartcontractkit/offchain-reporting/blob/c72a6e434f1397b6ede9c31fd5c540a782df3476/lib/offchainreporting2plus/internal/ocr3/keyring/onchain_keyring_ecdsa_testvectors_test.go
var ecdsaTestVectors = []ecdsaTestVector{
	{
		desc:         "key=00000000...0001 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ec",
		s:            "20936516ffe5cd25d9353afb44016c3986b97209f7d8d541f7bfa0fa2df4dda0",
		v:            1,
		sig:          "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ecdf6c9ae9001a32da26cac504bbfe93c533f56adcb76fcaf9c812bd92a24163a1",
		sigValid:     true,
	},
	{
		desc:         "key=00000000...0001 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "1e22f945d3e0aa61bf6b5a1b1dd831421c8df3f30024060b9110f4054f0b813e",
		s:            "6052f1af0594e0986af54e1302a25a80394d1eddce2067edab8a7690b2a33f7d",
		v:            0,
		sig:          "1e22f945d3e0aa61bf6b5a1b1dd831421c8df3f30024060b9110f4054f0b813e6052f1af0594e0986af54e1302a25a80394d1eddce2067edab8a7690b2a33f7d",
		sigValid:     true,
	},
	{
		desc:         "key=00000000...0001 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "dbee1e1fe06c459b06f6379c520e87e672390f74a54fee9749c42a5ba370bf71",
		s:            "298a455436e51191dfed39795c1c8ed0330c85c3903dcec10add404d7a66a7cd",
		v:            0,
		sig:          "dbee1e1fe06c459b06f6379c520e87e672390f74a54fee9749c42a5ba370bf71298a455436e51191dfed39795c1c8ed0330c85c3903dcec10add404d7a66a7cd",
		sigValid:     true,
	},
	{
		desc:         "key=00000000...0001 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "24b5feb9879f034070de5974785768ab9ef30883f4b1323ce00ca32a829ef47c",
		s:            "338618290274465e2763a73112d18f1620cdc68ceadaf011fe38cd85fe16457d",
		v:            0,
		sig:          "24b5feb9879f034070de5974785768ab9ef30883f4b1323ce00ca32a829ef47c338618290274465e2763a73112d18f1620cdc68ceadaf011fe38cd85fe16457d",
		sigValid:     true,
	},
	{
		desc:         "key=00000000...0001 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "e19ff678871fbe78eb55085b155c3f39afe8db2e1ebf932cacbfa9ef6d66b421",
		s:            "556dffb2245945e24577fee6075f0e099815bdf0bcaf2e34024f947c1ac8be34",
		v:            0,
		sig:          "e19ff678871fbe78eb55085b155c3f39afe8db2e1ebf932cacbfa9ef6d66b421556dffb2245945e24577fee6075f0e099815bdf0bcaf2e34024f947c1ac8be34",
		sigValid:     true,
	},
	{
		desc:         "key=ffffffff...4140 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
		address:      "80c0dbf239224071c59dd8970ab9d542e3414ab2",
		r:            "8e5fd9590aba664e5fb6fb9ad20e14017095ce7b2224e6b010938c9d09d3c466",
		s:            "1a338ed6ec9f598fc9f6509897e9de05711e963f7869973bee978647ee37fe8b",
		v:            0,
		sig:          "8e5fd9590aba664e5fb6fb9ad20e14017095ce7b2224e6b010938c9d09d3c4661a338ed6ec9f598fc9f6509897e9de05711e963f7869973bee978647ee37fe8b",
		sigValid:     true,
	},
	{
		desc:         "key=ffffffff...4140 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
		address:      "80c0dbf239224071c59dd8970ab9d542e3414ab2",
		r:            "b647f922f44bd6e58b519ff65bd4b835ad32cafe8abd97e571ef32e42f58e366",
		s:            "5a6dd19682006a3b284e97b9b245ab402527e36a7a570a0a2cc85c1155ce5cac",
		v:            0,
		sig:          "b647f922f44bd6e58b519ff65bd4b835ad32cafe8abd97e571ef32e42f58e3665a6dd19682006a3b284e97b9b245ab402527e36a7a570a0a2cc85c1155ce5cac",
		sigValid:     true,
	},
	{
		desc:         "key=ffffffff...4140 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
		address:      "80c0dbf239224071c59dd8970ab9d542e3414ab2",
		r:            "49db88554ed0bd94cbcfea6aaf62a1885a795242246efe07d296b9aa3d6096d2",
		s:            "03169c934d0d102770b114a9f24307c31c74c29627cb3cfade24e30964272c9c",
		v:            1,
		sig:          "49db88554ed0bd94cbcfea6aaf62a1885a795242246efe07d296b9aa3d6096d2fce9636cb2f2efd88f4eeb560dbcf83b9e3a1a50877d6340e1ad7b836c0f14a5",
		sigValid:     true,
	},
	{
		desc:         "key=ffffffff...4140 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
		address:      "80c0dbf239224071c59dd8970ab9d542e3414ab2",
		r:            "e4854608a40360cb17dcb81e209077323c0f2568b7bc323b6e5951776a17478b",
		s:            "12fb5b73dab25aa987a794236f476a5e974041fd15938656c35e1750384396aa",
		v:            1,
		sig:          "e4854608a40360cb17dcb81e209077323c0f2568b7bc323b6e5951776a17478bed04a48c254da55678586bdc90b895a0236e9ae999b519e4fc74473c97f2aa97",
		sigValid:     true,
	},
	{
		desc:         "key=ffffffff...4140 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
		address:      "80c0dbf239224071c59dd8970ab9d542e3414ab2",
		r:            "2500e000fa35f24302b921b11b01c2348ec4948002ab2dee446a1a039dcd5713",
		s:            "38f316967b09f825a280b0f4ffd46bd0a8364238212d2078b005e2fed6e01ccd",
		v:            1,
		sig:          "2500e000fa35f24302b921b11b01c2348ec4948002ab2dee446a1a039dcd5713c70ce96984f607da5d7f4f0b002b942e12789aae8e1b7fc30fcc7b8df9562474",
		sigValid:     true,
	},
	{
		desc:         "key=4c0883a6...0fb0 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
		address:      "ef9124c50ac6abdeb9717cf0a55ee3e5a757a7a7",
		r:            "8be03fe5bfc15fb8bed8cd29e8918e5b62ccdd56438804efeb446104178c7343",
		s:            "16ccf3e15773f613df5d8d800da27fce738ca1538fed88d077c086a25a3189a8",
		v:            0,
		sig:          "8be03fe5bfc15fb8bed8cd29e8918e5b62ccdd56438804efeb446104178c734316ccf3e15773f613df5d8d800da27fce738ca1538fed88d077c086a25a3189a8",
		sigValid:     true,
	},
	{
		desc:         "key=4c0883a6...0fb0 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
		address:      "ef9124c50ac6abdeb9717cf0a55ee3e5a757a7a7",
		r:            "4efe2c34c38a645f41acb70bd6f5a653370e6e18c1f4fcab9ec48da12285017b",
		s:            "62102f237e2df6a602c9e80c5206412c18de1b89ddb224909aedeee9a9be1e39",
		v:            0,
		sig:          "4efe2c34c38a645f41acb70bd6f5a653370e6e18c1f4fcab9ec48da12285017b62102f237e2df6a602c9e80c5206412c18de1b89ddb224909aedeee9a9be1e39",
		sigValid:     true,
	},
	{
		desc:         "key=4c0883a6...0fb0 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
		address:      "ef9124c50ac6abdeb9717cf0a55ee3e5a757a7a7",
		r:            "3d915ad6af7d24f0899a380dc067f673c102060c8328bbc3897fd555cfebcc44",
		s:            "1179fe1871d3ce181af4e1780a5c8d093d1e33f1379d2a4b36c62e8a1226fd4a",
		v:            0,
		sig:          "3d915ad6af7d24f0899a380dc067f673c102060c8328bbc3897fd555cfebcc441179fe1871d3ce181af4e1780a5c8d093d1e33f1379d2a4b36c62e8a1226fd4a",
		sigValid:     true,
	},
	{
		desc:         "key=4c0883a6...0fb0 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
		address:      "ef9124c50ac6abdeb9717cf0a55ee3e5a757a7a7",
		r:            "bf06bd4d785c5c459a17cea72e51b015a1ea5a4d5578a2506e633de297c9888d",
		s:            "19536c85a2f3ea64d4a5e0cd887be3e51cfee028b7f910b5b38217114768b77e",
		v:            0,
		sig:          "bf06bd4d785c5c459a17cea72e51b015a1ea5a4d5578a2506e633de297c9888d19536c85a2f3ea64d4a5e0cd887be3e51cfee028b7f910b5b38217114768b77e",
		sigValid:     true,
	},
	{
		desc:         "key=4c0883a6...0fb0 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "4c0883a69102937d6231471b5dbb6204fe512961708279f1d7b1b86a2f130fb0",
		address:      "ef9124c50ac6abdeb9717cf0a55ee3e5a757a7a7",
		r:            "002235dbbb03712e5609c268862d75988847c158c4ef181ac34a3692d010785b",
		s:            "4fc2789cc68b7f883013056df66acb5167d77d166f3054ec7770ca8b32423e19",
		v:            1,
		sig:          "002235dbbb03712e5609c268862d75988847c158c4ef181ac34a3692d010785bb03d876339748077cfecfa92099534ad52d75fd040184b4f486194019df40328",
		sigValid:     true,
	},
	{
		desc:         "key=c87509a1...c0d3 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
		address:      "627306090abab3a6e1400e9345bc60c78a8bef57",
		r:            "4427683b75323c793236b9a198cf0ef9537e40dfd27ea3eb46c3b935a515b511",
		s:            "175984bd0adacee6ca014c983cdd19503cfd591f5f6afcbe34a8cd5ca5176e73",
		v:            1,
		sig:          "4427683b75323c793236b9a198cf0ef9537e40dfd27ea3eb46c3b935a515b511e8a67b42f525311935feb367c322e6ae7db183c74fdda37d8b2991302b1ed2ce",
		sigValid:     true,
	},
	{
		desc:         "key=c87509a1...c0d3 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
		address:      "627306090abab3a6e1400e9345bc60c78a8bef57",
		r:            "14d1a04922e2a9408543c44d3757fafc080aef21b82b76dc42f983c56c39fa4c",
		s:            "6d62da300c0fd29970874a2e0c07c390a14b650d20ee4cf976efc5c300b4c68b",
		v:            0,
		sig:          "14d1a04922e2a9408543c44d3757fafc080aef21b82b76dc42f983c56c39fa4c6d62da300c0fd29970874a2e0c07c390a14b650d20ee4cf976efc5c300b4c68b",
		sigValid:     true,
	},
	{
		desc:         "key=c87509a1...c0d3 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
		address:      "627306090abab3a6e1400e9345bc60c78a8bef57",
		r:            "049be5f3938a514467e37132e31eb7f8b513ea8093e43b2ec4abbbf06b883763",
		s:            "37a9328edf6786328ac3cd0459868fffd131e4395cf171a971ce3d3d003c11a2",
		v:            1,
		sig:          "049be5f3938a514467e37132e31eb7f8b513ea8093e43b2ec4abbbf06b883763c856cd71209879cd753c32fba6796ffee97cf8ad52572e924e04214fcffa2f9f",
		sigValid:     true,
	},
	{
		desc:         "key=c87509a1...c0d3 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
		address:      "627306090abab3a6e1400e9345bc60c78a8bef57",
		r:            "22e10a8486749dcc68c0f02a0a05b16807731ba956d442b9dcfaf801e49ea346",
		s:            "3c2c7cc5cf78f768881537a77a130140fcf360bb4d8f13c7b1b5fb751f81595f",
		v:            1,
		sig:          "22e10a8486749dcc68c0f02a0a05b16807731ba956d442b9dcfaf801e49ea346c3d3833a3087089777eac85885ecfebdbdbb7c2b61b98c740e1c6317b0b4e7e2",
		sigValid:     true,
	},
	{
		desc:         "key=c87509a1...c0d3 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3",
		address:      "627306090abab3a6e1400e9345bc60c78a8bef57",
		r:            "48397612187f6c6131cc5092cc4332cb1443037a3bd8d1b2b4a3659bcbb7ed62",
		s:            "05e11c1e9d55ac37136bc4bcf50d80123c59c86928d6d7d7a546a5ca182dbed3",
		v:            1,
		sig:          "48397612187f6c6131cc5092cc4332cb1443037a3bd8d1b2b4a3659bcbb7ed62fa1ee3e162aa53c8ec943b430af27fec7e55147d8671c8641a8bb8c2b808826e",
		sigValid:     true,
	},
	{
		desc:         "key=f8b8af8c...4181 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
		address:      "3b7afc0192e752307a611e53ad177d55482bce2a",
		r:            "7f6f2c1e35077382964856c346a71f4312dff20daa256f1fcc79dfb07ad15e2f",
		s:            "212abf396d60b4168e45c7590df331e249c34b5b119012f0c6e114441640d962",
		v:            0,
		sig:          "7f6f2c1e35077382964856c346a71f4312dff20daa256f1fcc79dfb07ad15e2f212abf396d60b4168e45c7590df331e249c34b5b119012f0c6e114441640d962",
		sigValid:     true,
	},
	{
		desc:         "key=f8b8af8c...4181 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
		address:      "3b7afc0192e752307a611e53ad177d55482bce2a",
		r:            "08bd9d8114e38c2a99d7aeaafa7ba7045b76352f35ea9b85c721358b947a60b2",
		s:            "786ceb931e3127bc3ef6fd5b068bc4be64dcc7f8b5cfe73b4568002ec526023a",
		v:            1,
		sig:          "08bd9d8114e38c2a99d7aeaafa7ba7045b76352f35ea9b85c721358b947a60b28793146ce1ced843c10902a4f9743b4055d214edf978b9007a6a5e5e0b103f07",
		sigValid:     true,
	},
	{
		desc:         "key=f8b8af8c...4181 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
		address:      "3b7afc0192e752307a611e53ad177d55482bce2a",
		r:            "2628d9918f7b19b23920b5ef29100a90f0fd77e7872ae19209bdf7608cb3ff5d",
		s:            "321adf17499d5d28491766b873a16d8898d23d63f1100057c3c3ce1a8ce584db",
		v:            0,
		sig:          "2628d9918f7b19b23920b5ef29100a90f0fd77e7872ae19209bdf7608cb3ff5d321adf17499d5d28491766b873a16d8898d23d63f1100057c3c3ce1a8ce584db",
		sigValid:     true,
	},
	{
		desc:         "key=f8b8af8c...4181 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
		address:      "3b7afc0192e752307a611e53ad177d55482bce2a",
		r:            "b8bd492006a4e8e83ea1c24d138b23d47055816d8a93ca9bd2868d254b8db112",
		s:            "2e9ed31d0981db11d3b5d023621e169cc07479e9dbb1c9b8c05de8849d28280f",
		v:            0,
		sig:          "b8bd492006a4e8e83ea1c24d138b23d47055816d8a93ca9bd2868d254b8db1122e9ed31d0981db11d3b5d023621e169cc07479e9dbb1c9b8c05de8849d28280f",
		sigValid:     true,
	},
	{
		desc:         "key=f8b8af8c...4181 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
		address:      "3b7afc0192e752307a611e53ad177d55482bce2a",
		r:            "8d186bec27d486e2adf26fc6212d535bfef2f9ee70433634467de5964445b534",
		s:            "2cfccb8e033398d85af2ff1e3955cffd16a408fa14599121ae65e1e9ae87c5c3",
		v:            0,
		sig:          "8d186bec27d486e2adf26fc6212d535bfef2f9ee70433634467de5964445b5342cfccb8e033398d85af2ff1e3955cffd16a408fa14599121ae65e1e9ae87c5c3",
		sigValid:     true,
	},
	{
		desc:         "key=e91671c4...55c2 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
		address:      "1b4a166527da727855095608bf2a0809011d3842",
		r:            "9daca4a62f8e2c7673fb90df7226f106df7b04e152017fcac77cac52f8d4c945",
		s:            "243ae5c9e02c0f455ff6bb0611f54932471b36c8e1e3930a354abb76085d288f",
		v:            0,
		sig:          "9daca4a62f8e2c7673fb90df7226f106df7b04e152017fcac77cac52f8d4c945243ae5c9e02c0f455ff6bb0611f54932471b36c8e1e3930a354abb76085d288f",
		sigValid:     true,
	},
	{
		desc:         "key=e91671c4...55c2 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
		address:      "1b4a166527da727855095608bf2a0809011d3842",
		r:            "8d615180fe87218c880f9ffb5767a53b4adc0e36957dd816520f56a3aebd01eb",
		s:            "57446ef8c7ba3945c994955820c0a0ed9b5374f0a10f2c2a00aab9342e095082",
		v:            1,
		sig:          "8d615180fe87218c880f9ffb5767a53b4adc0e36957dd816520f56a3aebd01eba8bb91073845c6ba366b6aa7df3f5f111f5b67f60e397411bf27a558a22cf0bf",
		sigValid:     true,
	},
	{
		desc:         "key=e91671c4...55c2 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
		address:      "1b4a166527da727855095608bf2a0809011d3842",
		r:            "eeb534134ffe7e4c41da15a684c828a487d7f696a56ebcb4a761db8f269e4c61",
		s:            "3d85d8a2d5422b0b1fc4739682513c9820c75f749ab5f20eeee86d9d649277f6",
		v:            1,
		sig:          "eeb534134ffe7e4c41da15a684c828a487d7f696a56ebcb4a761db8f269e4c61c27a275d2abdd4f4e03b8c697daec36699e77d721492ae2cd0e9f0ef6ba3c94b",
		sigValid:     true,
	},
	{
		desc:         "key=e91671c4...55c2 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
		address:      "1b4a166527da727855095608bf2a0809011d3842",
		r:            "46f0ea74d7b90488c0e2290ff6e0ee3e542b40c9ee53a77fad33575f400657b0",
		s:            "52570a99fd8ddd2dce434fb47c30a11fe40a3479c2f52e2b269d8566dd539957",
		v:            0,
		sig:          "46f0ea74d7b90488c0e2290ff6e0ee3e542b40c9ee53a77fad33575f400657b052570a99fd8ddd2dce434fb47c30a11fe40a3479c2f52e2b269d8566dd539957",
		sigValid:     true,
	},
	{
		desc:         "key=e91671c4...55c2 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
		address:      "1b4a166527da727855095608bf2a0809011d3842",
		r:            "664edaaf3d2d27670a106707640ea22336fbc469460dbbab6b8ddea474b315eb",
		s:            "3e11a5a08059a1ff8337ae815e1468eb09ca1665183dac4cfe55bed960c00865",
		v:            0,
		sig:          "664edaaf3d2d27670a106707640ea22336fbc469460dbbab6b8ddea474b315eb3e11a5a08059a1ff8337ae815e1468eb09ca1665183dac4cfe55bed960c00865",
		sigValid:     true,
	},
	{
		desc:         "key=69ec59ea...ba64 | zero digest, seqNr=0, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",
		address:      "607c2cbf9fe450ffcfc8d4609c595e7271782721",
		r:            "dee70153867f15874ff8c2222bdc766e87e9f5406b8c56ff9ed30c9ae22135a7",
		s:            "2834a5223aafe201171d0cef86b2c98b36f12aa131042e51a4e3230c2b191eae",
		v:            1,
		sig:          "dee70153867f15874ff8c2222bdc766e87e9f5406b8c56ff9ed30c9ae22135a7d7cb5addc5501dfee8e2f310794d367383bdb2457e4471ea1aef3b80a51d2293",
		sigValid:     true,
	},
	{
		desc:         "key=69ec59ea...ba64 | zero digest, seqNr=1, one-byte report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        1,
		report:       "42",
		reportHash:   "ed2d6d299a4cd568b2732306b21a75302ccc42a6249256b5b58c93fe5ebbfb32",
		signingKey:   "69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",
		address:      "607c2cbf9fe450ffcfc8d4609c595e7271782721",
		r:            "d79f6e0a636522933f2021a49762c3a605c00528bfb0e11e332bafb94b4f8ada",
		s:            "128160327fefaed34bad71aed9761b6e5a3617b760fd28ef381c19ab5b44fdb2",
		v:            0,
		sig:          "d79f6e0a636522933f2021a49762c3a605c00528bfb0e11e332bafb94b4f8ada128160327fefaed34bad71aed9761b6e5a3617b760fd28ef381c19ab5b44fdb2",
		sigValid:     true,
	},
	{
		desc:         "key=69ec59ea...ba64 | zero digest, seqNr=max, empty report",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        18446744073709551615,
		report:       "",
		reportHash:   "6e0d3574a8cda294d3ee359b0a6e39ab066d5fb0fa013c8415fbf124f511ec17",
		signingKey:   "69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",
		address:      "607c2cbf9fe450ffcfc8d4609c595e7271782721",
		r:            "af3ff5c266763c6ec6df423f62d4076bd51d3889c99e2da29b5651d7b3111522",
		s:            "1182f4faf6e4d11aedff0970fe99978801adab4579a28a9e977830af8b6fddcc",
		v:            0,
		sig:          "af3ff5c266763c6ec6df423f62d4076bd51d3889c99e2da29b5651d7b31115221182f4faf6e4d11aedff0970fe99978801adab4579a28a9e977830af8b6fddcc",
		sigValid:     true,
	},
	{
		desc:         "key=69ec59ea...ba64 | realistic digest, seqNr=1000, short report",
		configDigest: "0001abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345",
		seqNr:        1000,
		report:       "68656c6c6f20776f726c64",
		reportHash:   "384945bf326e6c56341a041a57e7aea3f2e0373909e839e01cce882a658ab427",
		signingKey:   "69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",
		address:      "607c2cbf9fe450ffcfc8d4609c595e7271782721",
		r:            "92de01c3df67b3671a114b7ff8875a039c772a8ead14988f1769a69908cd423a",
		s:            "27a0a0bc082ebdc17a75bd979594b5e395dc35a7c74e83e01b7a9cb3a0dcb090",
		v:            0,
		sig:          "92de01c3df67b3671a114b7ff8875a039c772a8ead14988f1769a69908cd423a27a0a0bc082ebdc17a75bd979594b5e395dc35a7c74e83e01b7a9cb3a0dcb090",
		sigValid:     true,
	},
	{
		desc:         "key=69ec59ea...ba64 | all-ones digest, seqNr=2^32, medium report",
		configDigest: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		seqNr:        4294967296,
		report:       "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		reportHash:   "005538eb461897c9eb8b57527ed070be20008b916e41d9949840f7574911915f",
		signingKey:   "69ec59eaa1f4f2e36b639716b7c30ca86d9a5375c7b38d8918bd9c0ebc80ba64",
		address:      "607c2cbf9fe450ffcfc8d4609c595e7271782721",
		r:            "b77378921c491e668ca9f4acb94798d1c6dbd3565caf5deef5629226a5d7846a",
		s:            "726a25173ecf01d4eaa91b1ee5b5d884b5889876d66ab5e7a2899e2d8b391ba6",
		v:            1,
		sig:          "b77378921c491e668ca9f4acb94798d1c6dbd3565caf5deef5629226a5d7846a8d95dae8c130fe2b1556e4e11a4a277a0526446fd8ddea541d48c05f44fd259b",
		sigValid:     true,
	},
}

var ecdsaNegativeTestVectors = []ecdsaTestVector{
	{
		desc:         "configure digest used for sig not matching",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ec",
		s:            "20936516ffe5cd25d9353afb44016c3986b97209f7d8d541f7bfa0fa2df4dda0",
		v:            1,
		sig:          "00cb1b6df5e3f977dc55ab1aac38bb17a4f23c5b7673d098a0afb40d9ecb46d5da2e772610be3c8316b56561310a8757f234aaa0f69317089fc25d34c491c931",
		sigValid:     false,
	},
	{
		desc:         "seqNr used for sig not matching",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ec",
		s:            "20936516ffe5cd25d9353afb44016c3986b97209f7d8d541f7bfa0fa2df4dda0",
		v:            1,
		sig:          "7dd3a9896c791e98d2560fe1999736ac10ff563f9a1f7cdc5456580950ed049c1caf74f71dd2dc1c756b93ebdc005b2dc9f3ae257791ac29f4b814b4c4c22595",
		sigValid:     false,
	},
	{
		desc:         "sig compression invalid",
		configDigest: "0000000000000000000000000000000000000000000000000000000000000000",
		seqNr:        0,
		report:       "",
		reportHash:   "fb8b8f91fdf6aec9b78c8c73bf7ad4ff942dc378659c50d97a6391ce8fa8ad23",
		signingKey:   "0000000000000000000000000000000000000000000000000000000000000001",
		address:      "7e5f4552091a69125d5dfcb7b8c2659029395bdf",
		r:            "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ec",
		s:            "20936516ffe5cd25d9353afb44016c3986b97209f7d8d541f7bfa0fa2df4dda0",
		v:            1,
		sig:          "2c757bd8ff2fcb408a1e26cf42e18ddecfda5656e573cd5255eee329705615ec20936516ffe5cd25d9353afb44016c3986b97209f7d8d541f7bfa0fa2df4dda0",
		sigValid:     false,
	},
}
