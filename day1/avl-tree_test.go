package main

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func TestLeftLeft(t *testing.T) {
	tr := new(AVL)
	tr.insert(1)
	tr.insert(2)
	tr.insert(3)
	tr.insert(4)
	tr.insert(5)
	l := tr.toList()
	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Printf("%d ", v.Value.(int))
	}

	if tr.root.v != 2 {
		t.Fatal("Invalid value for root")
	}

	if tr.root.left.v != 1 {
		t.Fatal("Invalid value for left node")
	}

	if tr.root.right.v != 4 {
		t.Fatal("Invalid value for right node")
	}
	if tr.root.right.left.v != 3 {
		t.Fatal("Invalid value for right node")
	}
	if tr.root.right.v != 4 {
		t.Fatal("Invalid value for right node")
	}
}

func TestRightRight(t *testing.T) {
	tr := new(AVL)
	tr.insert(5)
	tr.insert(4)
	tr.insert(3)
	tr.insert(2)
	tr.insert(1)
	l := tr.toList()
	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Printf("%d ", v.Value.(int))
	}

	if tr.root.v != 4 {
		t.Fatal("Invalid value for root")
	}

	if tr.root.left.v != 2 {
		t.Fatal("Invalid value for left node")
	}
	if tr.root.left.left.v != 1 {
		t.Fatal("Invalid value for left node")
	}
	if tr.root.left.right.v != 3 {
		t.Fatal("Invalid value for left node")
	}

	if tr.root.right.v != 5 {
		t.Fatal("Invalid value for right node")
	}
}

func TestSorting(t *testing.T) {
	v := manyValues
	avl := new(AVL)
	for _, a := range v {
		avl.insert(a)
	}

	slices.Sort(v)
	l := avl.toList()
	idx := 0
	for x := l.Front(); x != nil; x = x.Next() {
		if x.Value.(int) != v[idx] {
			t.Fail()
			log.Printf("%d != %d", x.Value.(int), v[idx])
		}
		idx = idx + 1
	}
}

func TestManyValues(t *testing.T) {

	avl := new(AVL)
	for _, a := range manyValues {
		avl.insert(a)
	}

	for val := range avl.traverse() {
		log.Printf("%d ", val)
	}

	if avl.count != 1001 {
		t.Fatalf("Expected 1001 items in the avl, had %d", avl.count)
	}
}

var manyValues = []int{76569,
	38663,
	60350,
	60350, // duplicate
	35330,
	88681,
	30057,
	55455,
	48398,
	60451,
	23979,
	20498,
	18170,
	27928,
	45219,
	13289,
	24458,
	33053,
	30998,
	63320,
	18321,
	47393,
	71112,
	21392,
	60110,
	45686,
	43026,
	23080,
	81346,
	45701,
	67858,
	61483,
	52308,
	98286,
	99386,
	31064,
	90518,
	86055,
	77548,
	67600,
	13410,
	14441,
	73737,
	61602,
	88419,
	31757,
	51810,
	37775,
	42112,
	82029,
	84902,
	77147,
	55928,
	30168,
	35206,
	76009,
	51573,
	54931,
	23466,
	34344,
	56118,
	57495,
	20063,
	50850,
	94982,
	82710,
	54893,
	17029,
	31042,
	77274,
	36836,
	64993,
	36429,
	14228,
	93810,
	82201,
	11071,
	62119,
	74858,
	16332,
	41635,
	26392,
	30081,
	46852,
	61987,
	45439,
	16638,
	46884,
	78775,
	28266,
	50240,
	36216,
	63061,
	55186,
	33521,
	88738,
	17872,
	42966,
	74894,
	68870,
	32631,
	22114,
	54345,
	79241,
	75791,
	89775,
	36196,
	97340,
	70717,
	97532,
	17841,
	64433,
	57281,
	60158,
	38601,
	95768,
	74339,
	72283,
	66648,
	77312,
	24237,
	34725,
	81905,
	49368,
	70080,
	36941,
	25380,
	75150,
	95428,
	65063,
	95286,
	76555,
	69305,
	46692,
	60637,
	70348,
	63575,
	18822,
	10053,
	32355,
	33467,
	78154,
	21154,
	34321,
	63763,
	41970,
	95379,
	29009,
	57863,
	93891,
	55942,
	51505,
	57349,
	18516,
	47757,
	88631,
	82224,
	94668,
	37537,
	84583,
	62925,
	83652,
	48631,
	34328,
	17959,
	92497,
	74945,
	30957,
	20808,
	68043,
	92100,
	83140,
	50451,
	99877,
	17414,
	74459,
	71551,
	48028,
	45144,
	77667,
	55278,
	50857,
	71058,
	14884,
	26322,
	19517,
	61263,
	41865,
	46019,
	84622,
	72924,
	65922,
	57609,
	96221,
	58965,
	48467,
	38822,
	89271,
	50233,
	72132,
	50114,
	92854,
	68597,
	31426,
	26290,
	46519,
	99530,
	24195,
	24855,
	12474,
	79857,
	84581,
	14589,
	27460,
	95913,
	96860,
	68478,
	61539,
	54264,
	11712,
	67895,
	21806,
	94924,
	79279,
	71937,
	38121,
	99405,
	23777,
	15904,
	57158,
	90077,
	27713,
	94225,
	26741,
	76579,
	98566,
	96383,
	34920,
	30857,
	72631,
	97426,
	90023,
	42078,
	29177,
	15290,
	30020,
	48828,
	55399,
	92532,
	84224,
	56107,
	94015,
	74539,
	46023,
	65289,
	76828,
	31980,
	29563,
	37156,
	90827,
	90960,
	52081,
	37539,
	45714,
	33020,
	65172,
	48114,
	95266,
	14319,
	15016,
	44772,
	13318,
	66222,
	52139,
	89430,
	94339,
	51914,
	49546,
	77235,
	34977,
	32268,
	32335,
	87797,
	70688,
	34550,
	48762,
	36225,
	14156,
	54748,
	60824,
	23664,
	75308,
	85167,
	31875,
	80274,
	53914,
	65328,
	15842,
	64679,
	72913,
	26977,
	18805,
	50403,
	66822,
	25130,
	43179,
	28529,
	84485,
	88144,
	19444,
	25094,
	23999,
	76777,
	80395,
	87765,
	54941,
	35440,
	47296,
	82072,
	99634,
	66069,
	38850,
	30653,
	56745,
	21852,
	19323,
	83540,
	83398,
	68475,
	32324,
	75842,
	75024,
	30361,
	23965,
	89063,
	33958,
	52713,
	79512,
	57978,
	67269,
	52343,
	68862,
	23827,
	78795,
	51665,
	22205,
	44361,
	32727,
	51222,
	69372,
	87748,
	20468,
	64252,
	85477,
	74888,
	15886,
	50741,
	91573,
	64501,
	58070,
	48736,
	86830,
	42764,
	39435,
	57300,
	27997,
	94682,
	63535,
	61579,
	76805,
	80476,
	78044,
	20073,
	60476,
	67653,
	59888,
	51791,
	69612,
	92872,
	76912,
	39675,
	18997,
	39931,
	54580,
	80984,
	28979,
	90883,
	87336,
	15586,
	54359,
	61884,
	47021,
	44988,
	81841,
	10463,
	20507,
	42642,
	45174,
	98791,
	26063,
	53202,
	68734,
	11240,
	89762,
	26879,
	35769,
	86763,
	23079,
	62599,
	90177,
	89027,
	78085,
	23181,
	60101,
	70182,
	76082,
	23126,
	94557,
	39903,
	53351,
	27851,
	55945,
	94512,
	59442,
	46111,
	33188,
	69496,
	53706,
	84226,
	84530,
	15619,
	92909,
	82233,
	32228,
	22622,
	87096,
	45821,
	53028,
	87359,
	24280,
	33180,
	29942,
	62815,
	34378,
	19267,
	41823,
	19487,
	35166,
	53863,
	37077,
	21593,
	15250,
	97412,
	41281,
	50748,
	85109,
	93496,
	89991,
	21350,
	83836,
	99183,
	83195,
	57577,
	86197,
	75241,
	28959,
	19106,
	24181,
	58346,
	64260,
	63855,
	51041,
	12780,
	39423,
	15750,
	96117,
	48705,
	60266,
	56098,
	84332,
	59843,
	47823,
	92701,
	44807,
	87047,
	12269,
	32926,
	87543,
	37843,
	89968,
	81630,
	78677,
	97744,
	42608,
	80685,
	50122,
	13469,
	46215,
	19736,
	33883,
	57314,
	97638,
	53667,
	57195,
	25721,
	26068,
	25422,
	21805,
	30357,
	84783,
	51674,
	31130,
	66530,
	56565,
	45075,
	20799,
	62469,
	29429,
	28582,
	63901,
	82320,
	13257,
	60804,
	90339,
	42668,
	24336,
	26868,
	50073,
	84600,
	65452,
	95481,
	70538,
	99276,
	92621,
	91369,
	52612,
	36522,
	70248,
	88603,
	14573,
	67260,
	51191,
	44550,
	14867,
	40341,
	36004,
	81443,
	76574,
	87171,
	75979,
	74276,
	35093,
	83347,
	61220,
	75436,
	91705,
	44016,
	77374,
	56171,
	12623,
	67912,
	62519,
	50408,
	92685,
	69052,
	90063,
	19924,
	24684,
	94603,
	93869,
	36287,
	19610,
	41304,
	80834,
	73110,
	34006,
	49159,
	94986,
	23162,
	87756,
	78614,
	48810,
	30686,
	93411,
	92796,
	33708,
	69230,
	80346,
	31632,
	49454,
	87503,
	12993,
	24790,
	27278,
	31393,
	13890,
	45822,
	76509,
	73570,
	68202,
	11950,
	86295,
	59431,
	10073,
	46959,
	38024,
	59467,
	97947,
	52066,
	78026,
	67778,
	36416,
	28572,
	43037,
	13183,
	95497,
	59588,
	87551,
	47751,
	64227,
	25037,
	69189,
	78690,
	97053,
	86924,
	21239,
	70294,
	64633,
	63033,
	20972,
	57094,
	66334,
	84175,
	69423,
	59629,
	26492,
	24233,
	42703,
	43484,
	25830,
	27148,
	70173,
	37962,
	53875,
	62770,
	84383,
	38809,
	14689,
	58731,
	69616,
	12184,
	19032,
	21184,
	81673,
	92102,
	50429,
	41374,
	62668,
	38465,
	48412,
	85695,
	85327,
	40141,
	78149,
	19512,
	90774,
	30680,
	33876,
	48744,
	41261,
	94345,
	17561,
	82921,
	22353,
	66005,
	31257,
	47358,
	89407,
	16602,
	77071,
	33095,
	92791,
	90001,
	76457,
	45373,
	26903,
	94655,
	87065,
	59434,
	31424,
	98326,
	67336,
	63174,
	91623,
	30997,
	16487,
	36955,
	41035,
	40808,
	70906,
	83126,
	27778,
	29362,
	96931,
	49573,
	17169,
	62241,
	44935,
	56537,
	70264,
	53581,
	10468,
	72084,
	10700,
	95034,
	95618,
	44077,
	62523,
	73000,
	86080,
	70455,
	86887,
	36257,
	69290,
	84621,
	28070,
	67896,
	91077,
	28096,
	82719,
	13753,
	13359,
	73187,
	46578,
	51474,
	59914,
	32994,
	47213,
	27705,
	43458,
	55713,
	88478,
	76437,
	47552,
	61870,
	82030,
	66782,
	58762,
	73779,
	39491,
	56237,
	39046,
	83723,
	15550,
	45705,
	61209,
	34422,
	48150,
	54683,
	80717,
	73513,
	17391,
	93794,
	61475,
	53260,
	28755,
	61345,
	70447,
	17075,
	38645,
	34399,
	96669,
	53242,
	73388,
	43262,
	57113,
	99180,
	71205,
	23208,
	16013,
	48327,
	54890,
	35094,
	91186,
	57274,
	75941,
	57396,
	26112,
	54308,
	16053,
	92418,
	33397,
	68006,
	71144,
	41560,
	11541,
	12314,
	84666,
	45896,
	37274,
	25332,
	47691,
	19593,
	59745,
	16391,
	48465,
	27282,
	44459,
	56759,
	31978,
	49713,
	71736,
	48701,
	58335,
	85731,
	36730,
	24576,
	16255,
	86861,
	44160,
	42957,
	37675,
	46289,
	32863,
	44893,
	19026,
	14203,
	79022,
	90878,
	28194,
	99094,
	83283,
	97463,
	21838,
	33966,
	31560,
	51238,
	87950,
	97414,
	22335,
	43355,
	33567,
	77568,
	73302,
	20178,
	41008,
	61032,
	92382,
	28557,
	66152,
	55335,
	94368,
	96652,
	39401,
	33854,
	12726,
	58078,
	86375,
	93905,
	73642,
	10862,
	69937,
	16483,
	33726,
	91167,
	49496,
	40092,
	88996,
	87107,
	95555,
	10341,
	85542,
	67878,
	23412,
	50765,
	32794,
	96832,
	83174,
	64070,
	54569,
	96530,
	17512,
	71376,
	75803,
	85967,
	34569,
	96867,
	56321,
	31972,
	11367,
	16755,
	33148,
	27699,
	58246,
	81416,
	64682,
	14103,
	97226,
	76494,
	33262,
	83176,
	44601,
	19627,
	82257,
	33338,
	65644,
	40304,
	63980,
	35134,
	78787,
	31147,
	70378,
	96224,
	75530,
	43443,
	17552,
	15131,
	59594,
	74481,
	91569,
	67128,
	63154,
	51586,
	41213,
	66534,
	49095,
	25436,
	95020,
	93614,
	46555,
	35898,
	63537,
	71258,
	92852,
	97125,
	88632,
	57247,
	20277,
	10269,
	46681,
	31928,
	12487,
	76196,
	24895,
	46101,
	26438,
	86041,
	99574,
	14990,
	42622,
	76205,
	88565,
	88447,
	11332,
	47458,
	19133,
	58142,
	77904,
	45223,
	32352,
	88508,
	35832,
	79827,
	53932,
	16064,
	73085,
	17739,
	54478,
	28744,
	87515,
	86271,
	15579,
	51749,
	46430,
	24869,
	97219,
	13106,
	22072,
	67674,
	68925,
	33588,
	95625,
	31443,
	94542,
	21244,
	38829,
	68255,
	39081,
	25336,
	31219,
	87477,
	29146,
	37019,
	94223,
	20896,
	15884,
	76809,
	56403,
	23533,
	82416,
	24373}