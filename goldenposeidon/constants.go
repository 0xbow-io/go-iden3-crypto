package poseidon

import "github.com/n8wb/go-iden3-crypto/ffg"

const (
	// NROUNDSF const from original paper
	NROUNDSF = 8
	// NROUNDSP const from original paper
	NROUNDSP = 22
	// CAPLEN const
	CAPLEN = 4
	// mLen const
	mLen = 12
)

var (
	mcirc = []uint64{17, 15, 41, 16, 2, 28, 13, 13, 39, 18, 34, 20}
	mdiag = []uint64{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	c = []uint64{
		0xb585f766f2144405,
		0x7746a55f43921ad7,
		0xb2fb0d31cee799b4,
		0xf6760a4803427d7,
		0xe10d666650f4e012,
		0x8cae14cb07d09bf1,
		0xd438539c95f63e9f,
		0xef781c7ce35b4c3d,
		0xcdc4a239b0c44426,
		0x277fa208bf337bff,
		0xe17653a29da578a1,
		0xc54302f225db2c76,
		0xac6c9c2b4418dd61,
		0xe0888eb1e8a01286,
		0x813dbe952b98904e,
		0xcc3033609c9cf175,
		0x72cebc82a59c0f82,
		0x8150d8525753e741,
		0xb1122c74b268d66e,
		0x7c6ddd482375aa2,
		0xa4dd6f1ef49fb6af,
		0xd33b0d5b4f7ccfe5,
		0xc523112247209124,
		0x464804200134c32d,
		0xcd09dea180de4f2c,
		0xadb069225c93e4e6,
		0xbf01209b8a7c8534,
		0xb1eb37d319913823,
		0xdadf943b8d3e5a0d,
		0x6d15f3cb7a3520ba,
		0xf07af62b134ef181,
		0x568355076c6b0de6,
		0x31ca4bf93cab68b8,
		0xfbad37a125735ba,
		0x9d3a9caaf1ac9e0a,
		0x4f265810f020c095,
		0x6a84c9524e81a8bc,
		0x68ba410537925c79,
		0x422604631b34b07a,
		0x28e3a001f62f8290,
		0x3adfdccb8f734d41,
		0x73503e539baec66a,
		0xe8c1fd0142d9849c,
		0xe204ac13660546c5,
		0x8e2bb3ea97a40c53,
		0xac2800d1bf56548c,
		0x9494dca005d180d0,
		0xf36e1d066383ef53,
		0x8aa35b97a0e03c04,
		0xcf42a59addbd1f0c,
		0xa43ace89f8fdbd79,
		0x37585d8c243870c,
		0x4ab94ee3e26596fe,
		0xcee3abbb50d57b23,
		0xac91a7101a5ec55b,
		0x9173aa8462280d2d,
		0xaec1ca46ccb95105,
		0x57b2f2845db61e4a,
		0x95704158500c90c6,
		0x66e023b0e6c9df5f,
		0x315f63f4fec360ba,
		0xf3009795713abcf1,
		0xf4decc3fb00765ee,
		0x32620ac918682d50,
		0x49717d63a5fc742e,
		0x153516f22014ea2d,
		0xcc316380a2761fe4,
		0x2e49b3f7076d203d,
		0x44ac3e9bf0a2dc89,
		0x49d1e388d8e35c,
		0x53ec867cb39989fa,
		0xd2c9bcc8d65f5a62,
		0xc0cc930ee8540455,
		0x40651e0872505e8,
		0x168973b2ebafbe6c,
		0x9c7eecb3b40581c2,
		0x389473bcdfca97a2,
		0xb1cb0b3abe9753ad,
		0x41afceccffdb18e6,
		0x7bf841e237ccd6c9,
		0x6082a3f101fb888,
		0x8c1a39196f4163cc,
		0xb56664760c1c9476,
		0x2a02ac020d1eb5a3,
		0x6a9d48e8aa83605d,
		0x8a0d2f5c4c9c51b2,
		0x75fc65575b284ad4,
		0xadaedf7d1ce2a8dd,
		0x235bc889cc83968e,
		0xa8c30cf1781738f5,
		0x546b2a846753bcf8,
		0x9b68e8c06c04bd25,
		0x3fdf80794ebb443b,
		0x92ca132a9bec5a45,
		0x76133eecfd9bd1ff,
		0x3fb0fd5381054812,
		0xf15925978dbd52ff,
		0x2ee289ac37f0e879,
		0xd8af8654e9a2e659,
		0x8595bbd7f34c5e8a,
		0x206ddbf781e47b2,
		0xe101a767854a2f97,
		0xf4d4f0a01072c996,
		0x197aec2894aab642,
		0x8d0c3911220db49b,
		0xa62a8bad609227ca,
		0x1e4813a7e7b9cbce,
		0x6b547528731244eb,
		0xd08e48512bfea84e,
		0xb2920c88d3885857,
		0x1f0cd5d7a309fcc2,
		0x99a0ea0842fdb4fb,
		0xc227210554b6c53d,
		0x70e5269708f6f3a9,
		0xbe8f71c8c98bb3bd,
		0xf96fb39adc4baaf6,
		0x7f9a7555c60fc6c7,
		0xccaa5446d71fe6a5,
	}

	p = [][]uint64{
		{
			0x19,
			0x78566230aa7cc5d0,
			0xdbf23e50005e7f24,
			0xb4a02c5c826d523e,
			0x466d8f66a8f9fed5,
			0x68da2264f65ec3e,
			0xb59f9ff0ac6d5d78,
			0xcfb03c902d447551,
			0x2044ce14eaf8f5d9,
			0xfb9373c8481e0f0d,
			0x72af70cdcb99214f,
			0xe3ef40eacc6ff78d,
		},
		{
			0xf,
			0x817bd8a7869ed1b5,
			0x819f2c14a8366b1f,
			0x7a5cf5b7b922e946,
			0x727eca45c8d7bb71,
			0x605a82c52b5ad2f1,
			0x59ccc4d5184bc93a,
			0x66c8bab2096cfd38,
			0xeb4c0ce280c3e935,
			0x17f9202c16676b2f,
			0x9b6e5164ed35d878,
			0x6fadc9347faeee81,
		},
		{
			0x29,
			0xd267254bea1097f4,
			0x2dc10fce3233f443,
			0xfa9db0de2d852e7a,
			0xde2a0516f8c9d943,
			0xe6fdf23648931b99,
			0x3743057c07a5dbfa,
			0xa6fdb8ebccc51667,
			0x2c4916605e3dea58,
			0xe95c10ae32e05085,
			0x97f9b7d2cfc2ade5,
			0x9b6e5164ed35d878,
		},
		{
			0x10,
			0x60c33ebd1e023f0a,
			0xdb6945a20d277091,
			0x383dd77e07998487,
			0xe04ea1957ad8305c,
			0xd499fcbf63fbd266,
			0x462269e4b04620a5,
			0x63c9679d8572a867,
			0x81c44e9699915693,
			0x62ecbe05e02433fc,
			0xe95c10ae32e05085,
			0x17f9202c16676b2f,
		},
		{
			0x2,
			0xa89ef32ae1462322,
			0x77c1a153e73659e8,
			0x2aec981be4b62ed5,
			0xb70fb5f2b4f1f85f,
			0x7c66d474cd2087cb,
			0x39302966be7df654,
			0xb827c807875511c0,
			0xa4daffb3ffd0e78f,
			0x81c44e9699915693,
			0x2c4916605e3dea58,
			0xeb4c0ce280c3e935,
		},
		{
			0x1c,
			0x6250f5f176d483e7,
			0xaad1255d46e78f07,
			0x8a00c7c83c762584,
			0xc734f3829ed30b0c,
			0xb1a0132288b1619b,
			0x88685b4f0798dfd1,
			0xfc02e869e21b72f8,
			0xb827c807875511c0,
			0x63c9679d8572a867,
			0xa6fdb8ebccc51667,
			0x66c8bab2096cfd38,
		},
		{
			0xd,
			0xe16a6c1dee3ba347,
			0x13d316e45539aef4,
			0x577e0472764f061d,
			0x226a4dcf5db3316d,
			0x3373035a3ca3dac6,
			0x441f3a3747b5adb7,
			0x88685b4f0798dfd1,
			0x39302966be7df654,
			0x462269e4b04620a5,
			0x3743057c07a5dbfa,
			0x59ccc4d5184bc93a,
		},
		{
			0xd,
			0xec9730136b7c2c05,
			0xe1ecc5c21eec0646,
			0x956d3c8b5528e064,
			0x6df1d31fa84398f4,
			0xf4898a1a3554ee49,
			0x3373035a3ca3dac6,
			0xb1a0132288b1619b,
			0x7c66d474cd2087cb,
			0xd499fcbf63fbd266,
			0xe6fdf23648931b99,
			0x605a82c52b5ad2f1,
		},
		{
			0x27,
			0x3cf7c3a39d94c236,
			0x9e62c7d7b000cb0b,
			0xe202be7ad7265af6,
			0x82178371fa5fff69,
			0x6df1d31fa84398f4,
			0x226a4dcf5db3316d,
			0xc734f3829ed30b0c,
			0xb70fb5f2b4f1f85f,
			0xe04ea1957ad8305c,
			0xde2a0516f8c9d943,
			0x727eca45c8d7bb71,
		},
		{
			0x12,
			0xb4707207455f57e3,
			0x8e1de42b665c6706,
			0xee7b04568203481,
			0xe202be7ad7265af6,
			0x956d3c8b5528e064,
			0x577e0472764f061d,
			0x8a00c7c83c762584,
			0x2aec981be4b62ed5,
			0x383dd77e07998487,
			0xfa9db0de2d852e7a,
			0x7a5cf5b7b922e946,
		},
		{
			0x22,
			0xaadb39e83e76a9e0,
			0xcd9bf0bd292c5fda,
			0x8e1de42b665c6706,
			0x9e62c7d7b000cb0b,
			0xe1ecc5c21eec0646,
			0x13d316e45539aef4,
			0xaad1255d46e78f07,
			0x77c1a153e73659e8,
			0xdb6945a20d277091,
			0x2dc10fce3233f443,
			0x819f2c14a8366b1f,
		},
		{
			0x14,
			0x32f8ae916e567d39,
			0xaadb39e83e76a9e0,
			0xb4707207455f57e3,
			0x3cf7c3a39d94c236,
			0xec9730136b7c2c05,
			0xe16a6c1dee3ba347,
			0x6250f5f176d483e7,
			0xa89ef32ae1462322,
			0x60c33ebd1e023f0a,
			0xd267254bea1097f4,
			0x817bd8a7869ed1b5,
		},
	}

	s = []uint64{
		0x19,
		0x3d999c961b7c63b0,
		0x814e82efcd172529,
		0x2421e5d236704588,
		0x887af7d4dd482328,
		0xa5e9c291f6119b27,
		0xbdc52b2676a4b4aa,
		0x64832009d29bcf57,
		0x9c4155174a552cc,
		0x463f9ee03d290810,
		0xc810936e64982542,
		0x43b1c289f7bc3ac,
		0x94877900674181c3,
		0xc6c67cc37a2a2bbd,
		0xd667c2055387940f,
		0xba63a63e94b5ff0,
		0x99460cc41b8f079f,
		0x7ff02375ed524bb3,
		0xea0870b47a8caf0e,
		0xabcad82633b7bc9d,
		0x3b8d135261052241,
		0xfb4515f5e5b0d539,
		0x3ee8011c2b37f77c,
		0x19,
		0x673655aae8be5a8b,
		0xd510fe714f39fa10,
		0x2c68a099b51c9e73,
		0xa667bfa9aa96999d,
		0x4d67e72f063e2108,
		0xf84dde3e6acda179,
		0x40f9cc8c08f80981,
		0x5ead032050097142,
		0x6591b02092d671bb,
		0xe18c71963dd1b7,
		0x8a21bcd24a14218a,
		0xadef3740e71c726,
		0xa37bf67c6f986559,
		0xc6b16f7ed4fa1b00,
		0x6a065da88d8bfc3c,
		0x4cabc0916844b46f,
		0x407faac0f02e78d1,
		0x7a786d9cf0852cf,
		0x42433fb6949a629a,
		0x891682a147ce43b0,
		0x26cfd58e7b003b55,
		0x2bbf0ed7b657acb3,
		0x19,
		0x202800f4addbdc87,
		0xe4b5bdb1cc3504ff,
		0xbe32b32a825596e7,
		0x8e0f68c5dc223b9a,
		0x58022d9e1c256ce3,
		0x584d29227aa073ac,
		0x8b9352ad04bef9e7,
		0xaead42a3f445ecbf,
		0x3c667a1d833a3cca,
		0xda6f61838efa1ffe,
		0xe8f749470bd7c446,
		0x481ac7746b159c67,
		0xe367de32f108e278,
		0x73f260087ad28bec,
		0x5cfc82216bc1bdca,
		0xcaccc870a2663a0e,
		0xdb69cd7b4298c45d,
		0x7bc9e0c57243e62d,
		0x3cc51c5d368693ae,
		0x366b4e8cc068895b,
		0x2bd18715cdabbca4,
		0xa752061c4f33b8cf,
		0x19,
		0xc5b85bab9e5b3869,
		0x45245258aec51cf7,
		0x16e6b8e68b931830,
		0xe2ae0f051418112c,
		0x470e26a0093a65b,
		0x6bef71973a8146ed,
		0x119265be51812daf,
		0xb0be7356254bea2e,
		0x8584defff7589bd7,
		0x3c5fe4aeb1fb52ba,
		0x9e7cd88acf543a5e,
		0xb22d2432b72d5098,
		0x9e18a487f44d2fe4,
		0x4b39e14ce22abd3c,
		0x9e77fde2eb315e0d,
		0xca5e0385fe67014d,
		0xc2cb99bf1b6bddb,
		0x99ec1cd2a4460bfe,
		0x8577a815a2ff843f,
		0x7d80a6b4fd6518a5,
		0xeb6c67123eab62cb,
		0x8f7851650eca21a5,
		0x19,
		0x179be4bba87f0a8c,
		0xacf63d95d8887355,
		0x6696670196b0074f,
		0xd99ddf1fe75085f9,
		0xc2597881fef0283b,
		0xcf48395ee6c54f14,
		0x15226a8e4cd8d3b6,
		0xc053297389af5d3b,
		0x2c08893f0d1580e2,
		0xed3cbcff6fcc5ba,
		0xc82f510ecf81f6d0,
		0x11ba9a1b81718c2a,
		0x9f7d798a3323410c,
		0xa821855c8c1cf5e5,
		0x535e8d6fac0031b2,
		0x404e7c751b634320,
		0xa729353f6e55d354,
		0x4db97d92e58bb831,
		0xb53926c27897bf7d,
		0x965040d52fe115c5,
		0x9565fa41ebd31fd7,
		0xaae4438c877ea8f4,
		0x19,
		0x94b06183acb715cc,
		0x500392ed0d431137,
		0x861cc95ad5c86323,
		0x5830a443f86c4ac,
		0x3b68225874a20a7c,
		0x10b3309838e236fb,
		0x9b77fc8bcd559e2c,
		0xbdecf5e0cb9cb213,
		0x30276f1221ace5fa,
		0x7935dd342764a144,
		0xeac6db520bb03708,
		0x37f4e36af6073c6e,
		0x4edc0918210800e9,
		0xc44998e99eae4188,
		0x9f4310d05d068338,
		0x9ec7fe4350680f29,
		0xc5b2c1fdc0b50874,
		0xa01920c5ef8b2ebe,
		0x59fa6f8bd91d58ba,
		0x8bfc9eb89b515a82,
		0xbe86a7a2555ae775,
		0xcbb8bbaa3810babf,
		0x19,
		0x7186a80551025f8f,
		0x622247557e9b5371,
		0xc4cbe326d1ad9742,
		0x55f1523ac6a23ea2,
		0xa13dfe77a3d52f53,
		0xe30750b6301c0452,
		0x8bd488070a3a32b,
		0xcd800caef5b72ae3,
		0x83329c90f04233ce,
		0xb5b99e6664a0a3ee,
		0x6b0731849e200a7f,
		0x577f9a9e7ee3f9c2,
		0x88c522b949ace7b1,
		0x82f07007c8b72106,
		0x8283d37c6675b50e,
		0x98b074d9bbac1123,
		0x75c56fb7758317c1,
		0xfed24e206052bc72,
		0x26d7c3d1bc07dae5,
		0xf88c5e441e28dbb4,
		0x4fe27f9f96615270,
		0x514d4ba49c2b14fe,
		0x19,
		0xec3fabc192b01799,
		0x382b38cee8ee5375,
		0x3bfb6c3f0e616572,
		0x514abd0cf6c7bc86,
		0x47521b1361dcc546,
		0x178093843f863d14,
		0xad1003c5d28918e7,
		0x738450e42495bc81,
		0xaf947c59af5e4047,
		0x4653fb0685084ef2,
		0x57fde2062ae35bf,
		0xf02a3ac068ee110b,
		0xa3630dafb8ae2d7,
		0xce0dc874eaf9b55c,
		0x9a95f6cff5b55c7e,
		0x626d76abfed00c7b,
		0xa0c1cf1251c204ad,
		0xdaebd3006321052c,
		0x3d4bd48b625a8065,
		0x7f1e584e071f6ed2,
		0x720574f0501caed3,
		0xe3260ba93d23540a,
		0x19,
		0xe376678d843ce55e,
		0x66f3860d7514e7fc,
		0x7817f3dfff8b4ffa,
		0x3929624a9def725b,
		0x126ca37f215a80a,
		0xfce2f5d02762a303,
		0x1bc927375febbad7,
		0x85b481e5243f60bf,
		0x2d3c5f42a39c91a0,
		0x811719919351ae8,
		0xf669de0add993131,
		0xab1cbd41d8c1e335,
		0x9322ed4c0bc2df01,
		0x51c3c0983d4284e5,
		0x94178e291145c231,
		0xfd0f1a973d6b2085,
		0xd427ad96e2b39719,
		0x8a52437fecaac06b,
		0xdc20ee4b8c4c9a80,
		0xa2c98e9549da2100,
		0x1603fe12613db5b6,
		0xe174929433c5505,
		0x19,
		0x7de38bae084da92d,
		0x5b848442237e8a9b,
		0xf6c705da84d57310,
		0x31e6a4bdb6a49017,
		0x889489706e5c5c0f,
		0xe4a205459692a1b,
		0xbac3fa75ee26f299,
		0x5f5894f4057d755e,
		0xb0dc3ecd724bb076,
		0x5e34d8554a6452ba,
		0x4f78fd8c1fdcc5f,
		0x3d4eab2b8ef5f796,
		0xcfff421583896e22,
		0x4143cb32d39ac3d9,
		0x22365051b78a5b65,
		0x6f7fd010d027c9b6,
		0xd9dd36fba77522ab,
		0xa44cf1cb33e37165,
		0x3fc83d3038c86417,
		0xc4588d418e88d270,
		0xce1320f10ab80fe2,
		0xdb5eadbbec18de5d,
		0x19,
		0x4dd19c38779512ea,
		0xdb79ba02704620e9,
		0x92a29a3675a5d2be,
		0xd5177029fe495166,
		0xd32b3298a13330c1,
		0x251c4a3eb2c5f8fd,
		0xe1c48b26e0d98825,
		0x3301d3362a4ffccb,
		0x9bb6c88de8cd178,
		0xdc05b676564f538a,
		0x60192d883e473fee,
		0x1183dfce7c454afd,
		0x21cea4aa3d3ed949,
		0xfce6f70303f2304,
		0x19557d34b55551be,
		0x4c56f689afc5bbc9,
		0xa1e920844334f944,
		0xbad66d423d2ec861,
		0xf318c785dc9e0479,
		0x99e2032e765ddd81,
		0x400ccc9906d66f45,
		0xe1197454db2e0dd9,
		0x19,
		0x16b9774801ac44a0,
		0x3cb8411e786d3c8e,
		0xa86e9cf505072491,
		0x178928152e109ae,
		0x5317b905a6e1ab7b,
		0xda20b3be7f53d59f,
		0xcb97dedecebee9ad,
		0x4bd545218c59f58d,
		0x77dc8d856c05a44a,
		0x87948589e4f243fd,
		0x7e5217af969952c2,
		0x84d1ecc4d53d2ff1,
		0xd8af8b9ceb4e11b6,
		0x335856bb527b52f4,
		0xc756f17fb59be595,
		0xc0654e4ea5553a78,
		0x9e9a46b61f2ea942,
		0x14fc8b5b3b809127,
		0xd7009f0f103be413,
		0x3e0ee7b7a9fb4601,
		0xa74e888922085ed7,
		0xe80a7cde3d4ac526,
		0x19,
		0xbc58987d06a84e4d,
		0xb5d420244c9cae3,
		0xa3c4711b938c02c0,
		0x3aace640a3e03990,
		0x865a0f3249aacd8a,
		0x8d00b2a7dbed06c7,
		0x6eacb905beb7e2f8,
		0x45322b216ec3ec7,
		0xeb9de00d594828e6,
		0x88c5f20df9e5c26,
		0xf555f4112b19781f,
		0x238aa6daa612186d,
		0x9137a5c630bad4b4,
		0xc7db3817870c5eda,
		0x217e4f04e5718dc9,
		0xcae814e2817bd99d,
		0xe3292e7ab770a8ba,
		0x7bb36ef70b6b9482,
		0x3c7835fb85bca2d3,
		0xfe2cdf8ee3c25e86,
		0x61b3915ad7274b20,
		0xeab75ca7c918e4ef,
		0x19,
		0xa8cedbff1813d3a7,
		0x50dcaee0fd27d164,
		0xf1cb02417e23bd82,
		0xfaf322786e2abe8b,
		0x937a4315beb5d9b6,
		0x1b18992921a11d85,
		0x7d66c4368b3c497b,
		0xe7946317a6b4e99,
		0xbe4430134182978b,
		0x3771e82493ab262d,
		0xa671690d8095ce82,
		0xd6e15ffc055e154e,
		0xec67881f381a32bf,
		0xfbb1196092bf409c,
		0xdc9d2e07830ba226,
		0x698ef3245ff7988,
		0x194fae2974f8b576,
		0x7a5d9bea6ca4910e,
		0x7aebfea95ccdd1c9,
		0xf9bd38a67d5f0e86,
		0xfa65539de65492d8,
		0xf0dfcbe7653ff787,
		0x19,
		0xb035585f6e929d9d,
		0xba1579c7e219b954,
		0xcb201cf846db4ba3,
		0x287bf9177372cf45,
		0xa350e4f61147d0a6,
		0xd5d0ecfb50bcff99,
		0x2e166aa6c776ed21,
		0xe1e66c991990e282,
		0x662b329b01e7bb38,
		0x8aa674b36144d9a9,
		0xcbabf78f97f95e65,
		0xbd87ad390420258,
		0xad8617bca9e33c8,
		0xc00ad377a1e2666,
		0xac6fc58b3f0518f,
		0xc0cc8a892cc4173,
		0xc210accb117bc21,
		0xb73630dbb46ca18,
		0xc8be4920cbd4a54,
		0xbfe877a21be1690,
		0xae790559b0ded81,
		0xbf50db2f8d6ce31,
		0x19,
		0xeec24b15a06b53fe,
		0xc8a7aa07c5633533,
		0xefe9c6fa4311ad51,
		0xb9173f13977109a1,
		0x69ce43c9cc94aedc,
		0xecf623c9cd118815,
		0x28625def198c33c7,
		0xccfc5f7de5c3636a,
		0xf5e6c40f1621c299,
		0xcec0e58c34cb64b1,
		0xa868ea113387939f,
		0xcf29427ff7c58,
		0xbd9b3cf49eec8,
		0xd1dc8aa81fb26,
		0xbc792d5c394ef,
		0xd2ae0b2266453,
		0xd413f12c496c1,
		0xc84128cfed618,
		0xdb5ebd48fc0d4,
		0xd1b77326dcb90,
		0xbeb0ccc145421,
		0xd10e5b22b11d1,
		0x19,
		0xd8dddbdc5ce4ef45,
		0xacfc51de8131458c,
		0x146bb3c0fe499ac0,
		0x9e65309f15943903,
		0x80d0ad980773aa70,
		0xf97817d4ddbf0607,
		0xe4626620a75ba276,
		0xdfdc7fd6fc74f66,
		0xf464864ad6f2bb93,
		0x2d55e52a5d44414,
		0xdd8de62487c40925,
		0xe24c99adad8,
		0xcf389ed4bc8,
		0xe580cbf6966,
		0xcde5fd7e04f,
		0xe63628041b3,
		0xe7e81a87361,
		0xdabe78f6d98,
		0xefb14cac554,
		0xe5574743b10,
		0xd05709f42c1,
		0xe4690c96af1,
		0x19,
		0xc15acf44759545a3,
		0xcbfdcf39869719d4,
		0x33f62042e2f80225,
		0x2599c5ead81d8fa3,
		0xb306cb6c1d7c8d0,
		0x658c80d3df3729b1,
		0xe8d1b2b21b41429c,
		0xa1b67f09d4b3ccb8,
		0xe1adf8b84437180,
		0xd593a5e584af47b,
		0xa023d94c56e151c7,
		0xf7157bc98,
		0xe3006d948,
		0xfa65811e6,
		0xe0d127e2f,
		0xfc18bfe53,
		0xfd002d901,
		0xeed6461d8,
		0x1068562754,
		0xfa0236f50,
		0xe3af13ee1,
		0xfa460f6d1,
		0x19,
		0x49026cc3a4afc5a6,
		0xe06dff00ab25b91b,
		0xab38c561e8850ff,
		0x92c3c8275e105eeb,
		0xb65256e546889bd0,
		0x3c0468236ea142f6,
		0xee61766b889e18f2,
		0xa206f41b12c30415,
		0x2fe9d756c9f12d1,
		0xe9633210630cbf12,
		0x1ffea9fe85a0b0b1,
		0x11131738,
		0xf56d588,
		0x11050f86,
		0xf848f4f,
		0x111527d3,
		0x114369a1,
		0x106f2f38,
		0x11e2ca94,
		0x110a29f0,
		0xfa9f5c1,
		0x10f625d1,
		0x19,
		0x81d1ae8cc50240f3,
		0xf4c77a079a4607d7,
		0xed446b2315e3efc1,
		0xb0a6b70915178c3,
		0xb11ff3e089f15d9a,
		0x1d4dba0b7ae9cc18,
		0x65d74e2f43b48d05,
		0xa2df8c6b8ae0804a,
		0xa4e6f0a8c33348a6,
		0xc0a26efc7be5669b,
		0xa6b6582c547d0d60,
		0x11f718,
		0x10b6c8,
		0x134a96,
		0x10cf7f,
		0x124d03,
		0x13f8a1,
		0x117c58,
		0x132c94,
		0x134fc0,
		0x10a091,
		0x128961,
		0x19,
		0x84afc741f1c13213,
		0x2f8f43734fc906f3,
		0xde682d72da0a02d9,
		0xbb005236adb9ef2,
		0x5bdf35c10a8b5624,
		0x739a8a343950010,
		0x52f515f44785cfbc,
		0xcbaf4e5d82856c60,
		0xac9ea09074e3e150,
		0x8f0fa011a2035fb0,
		0x1a37905d8450904a,
		0x1300,
		0x1750,
		0x114e,
		0x131f,
		0x167b,
		0x1371,
		0x1230,
		0x182c,
		0x1368,
		0xf31,
		0x15c9,
		0x19,
		0x3abeb80def61cc85,
		0x9d19c9dd4eac4133,
		0x75a652d9641a985,
		0x9daf69ae1b67e667,
		0x364f71da77920a18,
		0x50bd769f745c95b1,
		0xf223d1180dbbf3fc,
		0x2f885e584e04aa99,
		0xb69a0fa70aea684a,
		0x9584acaa6e062a0,
		0xbc051640145b19b,
		0x14,
		0x22,
		0x12,
		0x27,
		0xd,
		0xd,
		0x1c,
		0x2,
		0x10,
		0x29,
		0xf,
	}

	// C is a constant array of element
	C []*ffg.Element
	// M is a matrix
	M [][]*ffg.Element
	// P is a matrix
	P [][]*ffg.Element
	// S is a array of element
	S []*ffg.Element
)

func init() {
	for i := 0; i < len(c); i++ {
		C = append(C, ffg.NewElementFromUint64(c[i]))
	}

	for i := 0; i < len(s); i++ {
		S = append(S, ffg.NewElementFromUint64(s[i]))
	}

	for i := 0; i < mLen; i++ {
		var (
			mRow []*ffg.Element
			pRow []*ffg.Element
		)
		for j := 0; j < mLen; j++ {
			ele := ffg.NewElementFromUint64(mcirc[(i-j+mLen)%mLen])
			if i == j {
				ele = ffg.NewElementFromUint64(mcirc[0] + mdiag[i])
			}
			mRow = append(mRow, ele)
			pRow = append(pRow, ffg.NewElementFromUint64(p[i][j]))
		}
		M = append(M, mRow)
		P = append(P, pRow)
	}
}
