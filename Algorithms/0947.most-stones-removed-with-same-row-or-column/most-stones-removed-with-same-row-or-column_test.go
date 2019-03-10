package problem0947

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones [][]int
	ans    int
}{
	{
		[][]int{{0, 1}, {1, 0}, {1, 1}},
		2,
	},

	{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
		5,
	},

	{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
		3,
	},

	{[][]int{{0, 0}},
		0,
	},

	{
		[][]int{{279, 2814}, {2343, 975}, {2680, 1479}, {1922, 451}, {1538, 2188}, {2180, 1620}, {2452, 1425}, {1626, 1625}, {2697, 1281}, {1450, 339}, {1252, 219}, {2256, 1772}, {150, 1442}, {691, 2958}, {2477, 1565}, {1730, 1884}, {1077, 291}, {641, 1267}, {1196, 1419}, {2590, 2078}, {1285, 2006}, {2608, 2610}, {173, 438}, {191, 2537}, {1240, 1797}, {1681, 2174}, {1145, 1306}, {2623, 2705}, {421, 538}, {2114, 428}, {1270, 2004}, {2091, 2376}, {754, 2827}, {2816, 1698}, {2606, 1299}, {1881, 1299}, {1511, 420}, {2929, 44}, {1945, 258}, {2700, 2455}, {36, 1556}, {1049, 1545}, {1242, 1486}, {893, 2322}, {2117, 2955}, {269, 1299}, {497, 2145}, {2495, 1110}, {2851, 1273}, {1221, 864}, {16, 1491}, {403, 2549}, {2125, 1281}, {160, 2062}, {2887, 141}, {2583, 2216}, {2247, 754}, {504, 1141}, {2372, 2404}, {1711, 1047}, {751, 2211}, {1232, 244}, {2309, 961}, {2943, 488}, {822, 1485}, {1273, 0}, {2306, 2209}, {2223, 712}, {2596, 86}, {2245, 128}, {2567, 727}, {2526, 2479}, {44, 2623}, {968, 735}, {1256, 1831}, {234, 1302}, {2181, 586}, {2365, 1033}, {529, 2910}, {822, 777}, {131, 489}, {729, 2030}, {1621, 299}, {932, 1589}, {1589, 1649}, {498, 1813}, {1218, 366}, {2309, 2754}, {2709, 2553}, {1671, 795}, {2851, 700}, {1944, 1143}, {2430, 2714}, {2283, 1291}, {2730, 838}, {257, 146}, {2862, 1167}, {2093, 2561}, {1379, 2025}, {901, 1019}, {2261, 2414}, {1597, 2796}, {1221, 2543}, {2688, 2399}, {1927, 363}, {113, 720}, {241, 411}, {2507, 668}, {1984, 2278}, {409, 1460}, {2916, 2547}, {1576, 2495}, {737, 1125}, {2566, 1241}, {2040, 344}, {2600, 2363}, {2798, 1218}, {1424, 2763}, {2180, 1459}, {2532, 2072}, {2493, 2318}, {595, 1889}, {1860, 353}, {2880, 1402}, {1191, 2453}, {1488, 392}, {777, 1609}, {1879, 180}, {2189, 2377}, {458, 1878}, {2350, 1404}, {1936, 779}, {615, 512}, {1562, 2045}, {1052, 2307}, {1024, 520}, {894, 1624}, {1733, 773}, {1807, 222}, {914, 563}, {2719, 448}, {1059, 520}, {340, 742}, {1951, 2170}, {2066, 2825}, {104, 604}, {2836, 1969}, {1676, 2523}, {961, 1456}, {1390, 647}, {2703, 2911}, {955, 2698}, {849, 183}, {1376, 2044}, {740, 2163}, {2577, 2983}, {688, 1846}, {328, 753}, {1177, 2262}, {1149, 1682}, {1837, 90}, {313, 2584}, {880, 1181}, {351, 2241}, {298, 2527}, {2260, 1626}, {2185, 1390}, {1019, 1915}, {2213, 600}, {2840, 1166}, {22, 722}, {2435, 2438}, {1829, 1793}, {1190, 1733}, {986, 815}, {1467, 2135}, {2255, 2238}, {826, 1067}, {1807, 1723}, {1355, 345}, {728, 184}, {2908, 174}, {312, 73}, {2878, 1597}, {1881, 14}, {306, 2664}, {2944, 3}, {2063, 2607}, {2403, 1494}, {1797, 549}, {2691, 2615}, {2118, 1281}, {1214, 1329}, {1655, 885}, {2973, 2797}, {365, 968}, {777, 1458}, {2922, 1403}, {1444, 2366}, {1573, 2718}, {1979, 1655}, {1869, 1847}, {184, 1298}, {1357, 1123}, {1744, 2903}, {2892, 2386}, {2798, 193}, {2126, 706}, {718, 2582}, {215, 220}, {2428, 1921}, {321, 1705}, {852, 2383}, {293, 771}, {1294, 933}, {961, 1037}, {2892, 2951}, {1026, 875}, {1642, 155}, {349, 824}, {2910, 2343}, {780, 726}, {1861, 2576}, {1456, 227}, {1961, 2710}, {2035, 2853}, {2179, 902}, {2295, 1126}, {62, 1761}, {378, 1722}, {2254, 505}, {203, 2621}, {209, 2429}, {2226, 2085}, {1412, 827}, {790, 755}, {2375, 893}, {968, 1049}, {457, 1696}, {2999, 244}, {1164, 2494}, {2850, 132}, {1556, 304}, {2104, 1394}, {2276, 2504}, {1856, 1838}, {1677, 893}, {221, 2677}, {2784, 2818}, {779, 1153}, {2962, 1161}, {2489, 85}, {37, 279}, {1033, 870}, {1502, 2978}, {1495, 1391}, {2700, 2851}, {45, 46}, {1335, 2155}, {2859, 296}, {2380, 447}, {1957, 328}, {562, 2135}, {2070, 1078}, {233, 48}, {407, 2666}, {1981, 564}, {250, 485}, {857, 2496}, {2194, 73}, {765, 261}, {2045, 1768}, {773, 1612}, {2728, 2325}, {1830, 640}, {2618, 1253}, {1392, 1088}, {2258, 1000}, {2738, 478}, {1735, 2406}, {600, 1774}, {1682, 1222}, {161, 2918}, {1887, 599}, {1498, 2535}, {1213, 612}, {2915, 795}, {2845, 976}, {424, 260}, {1075, 639}, {264, 2945}, {855, 715}, {1130, 2745}, {2562, 1625}, {763, 360}, {1724, 1043}, {2813, 1930}, {384, 2363}, {197, 275}, {914, 2153}, {809, 2885}, {143, 2347}, {659, 2872}, {2855, 2995}, {2686, 2122}, {623, 1043}, {2263, 2172}, {2634, 2672}, {990, 1897}, {1628, 672}, {863, 1034}, {1216, 129}, {1268, 2810}, {1858, 2332}, {2061, 632}, {1851, 2069}, {696, 1381}, {659, 859}, {2299, 322}, {2109, 989}, {65, 763}, {1873, 515}, {2954, 1859}, {544, 887}, {691, 1224}, {927, 968}, {433, 1468}, {2781, 1977}, {1435, 2594}, {275, 2802}, {523, 324}, {1266, 2146}, {10, 1139}, {1708, 121}, {2367, 2481}, {2663, 862}, {1068, 2308}, {2861, 2962}, {637, 1667}, {632, 2781}, {2685, 2013}, {1268, 1591}, {678, 1511}, {2365, 2337}, {1903, 2032}, {913, 275}, {710, 1831}, {719, 338}, {2089, 655}, {2179, 2849}, {2292, 2253}, {765, 1246}, {961, 926}, {1309, 1459}, {1856, 2023}, {601, 1582}, {156, 2969}, {1504, 1473}, {244, 2074}, {1472, 1647}, {445, 1769}, {1815, 2070}, {2084, 1787}, {1205, 2895}, {151, 1647}, {340, 368}, {2664, 2296}, {280, 667}, {49, 1935}, {1578, 2696}, {526, 2172}, {487, 673}, {2054, 1884}, {2236, 1559}, {823, 2519}, {861, 1870}, {2032, 1947}, {390, 450}, {654, 803}, {1985, 1291}, {336, 368}, {1124, 2793}, {968, 117}, {943, 2655}, {606, 1483}, {1126, 130}, {2155, 1903}, {2614, 874}, {1493, 1090}, {1699, 907}, {359, 1374}, {1104, 1391}, {2701, 2259}, {2242, 1176}, {2238, 569}, {492, 1431}, {232, 110}, {14, 1962}, {403, 113}, {1451, 2821}, {912, 2753}, {1160, 1590}, {576, 2686}, {2497, 2859}, {1217, 1186}, {2321, 1629}, {2024, 169}, {651, 1429}, {935, 2049}, {2980, 2642}, {2802, 1412}, {1946, 1963}, {1305, 801}, {2952, 1858}, {457, 276}, {2806, 667}, {1011, 2407}, {2861, 1668}, {190, 2286}, {1267, 2597}, {923, 369}, {698, 691}, {1398, 84}, {1174, 90}, {1540, 2547}, {370, 1549}, {2983, 114}, {1160, 1962}, {772, 1517}, {1230, 1953}, {2337, 1546}, {2122, 2784}, {1477, 2012}, {1991, 2178}, {1370, 2680}, {2207, 1009}, {588, 1200}, {1700, 139}, {1884, 1720}, {75, 2688}, {2171, 327}, {1038, 929}, {584, 1958}, {2458, 2163}, {1468, 2194}, {2348, 1135}, {557, 1244}, {2608, 1495}, {1754, 2722}, {2776, 2969}, {2699, 1236}, {225, 777}, {821, 397}, {271, 1004}, {540, 555}, {1590, 1583}, {557, 2920}, {1217, 1053}, {755, 2122}, {2758, 47}, {2868, 1563}, {1986, 2891}, {833, 369}, {2779, 1709}, {1132, 685}, {318, 2527}, {1461, 1803}, {1088, 1481}, {1694, 7}, {2805, 2829}, {2620, 2654}, {1355, 1701}, {1109, 414}, {766, 221}, {2617, 916}, {2155, 693}, {2918, 18}, {2830, 2616}, {1266, 346}, {2774, 2537}, {937, 126}, {1550, 864}, {2429, 2628}, {1048, 1460}, {1170, 2752}, {1484, 372}, {149, 571}, {983, 1184}, {2240, 2570}, {199, 625}, {2014, 975}, {1215, 730}, {1233, 2562}, {1589, 457}, {2392, 1592}, {2925, 882}, {25, 2940}, {261, 836}, {1806, 842}, {1828, 262}, {1776, 445}, {1047, 98}, {1196, 1624}, {1316, 2965}, {455, 1315}, {2324, 727}, {761, 1444}, {500, 852}, {1646, 2747}, {2339, 1074}, {919, 2837}, {2993, 1869}, {202, 2774}, {2857, 2298}, {1759, 605}, {1408, 1722}, {936, 1143}, {540, 1385}, {2770, 2769}, {2281, 1136}, {1672, 1138}, {2862, 1692}, {2729, 2369}, {643, 1205}, {911, 144}, {288, 618}, {2666, 2239}, {1589, 2969}, {2110, 655}, {2522, 2775}, {1227, 604}, {979, 2606}, {2797, 2395}, {2053, 1465}, {1765, 2688}, {2349, 62}, {866, 1912}, {2670, 2018}, {691, 683}, {2358, 308}, {1162, 2660}, {170, 2666}, {1641, 376}, {1696, 1367}, {231, 384}, {2915, 875}, {2675, 2499}, {212, 2735}, {1879, 1008}, {2958, 2053}, {544, 212}, {605, 2281}, {1390, 2014}, {1559, 1218}, {1194, 2222}, {1533, 1881}, {1863, 1736}, {2746, 473}, {2129, 481}, {2340, 1229}, {1394, 1040}, {1736, 2567}, {2442, 489}, {2388, 1453}, {226, 2027}, {1512, 1514}, {812, 2057}, {526, 992}, {1935, 1527}, {1167, 793}, {2236, 1763}, {740, 1623}, {619, 1777}, {2118, 763}, {2123, 19}, {2026, 32}, {1211, 697}, {73, 916}, {2847, 2372}, {559, 1180}, {2244, 1271}, {1632, 2730}, {1859, 2109}, {2701, 2367}, {2978, 1360}, {288, 1940}, {792, 356}, {1530, 936}, {891, 254}, {2943, 57}, {177, 1168}, {2095, 1910}, {1268, 2508}, {95, 23}, {2275, 838}, {1754, 1002}, {1634, 2616}, {69, 2285}, {1538, 1309}, {1135, 1559}, {2356, 1315}, {2568, 1483}, {2311, 2669}, {174, 943}, {669, 2851}, {1976, 2820}, {2526, 2874}, {656, 2117}, {227, 547}, {945, 508}, {2731, 355}, {443, 11}, {928, 1193}, {911, 1104}, {1760, 1363}, {897, 2599}, {810, 768}, {1911, 2502}, {1712, 2744}, {449, 1172}, {1189, 772}, {2195, 956}, {2938, 438}, {810, 915}, {2377, 2410}, {1437, 766}, {1804, 2212}, {2880, 1635}, {741, 2082}, {2180, 1998}, {635, 1480}, {2709, 909}, {1145, 2743}, {761, 2922}, {2117, 2770}, {477, 1934}, {838, 2609}, {1531, 1891}, {2553, 1820}, {2819, 463}, {1727, 637}, {1401, 1363}, {2652, 1603}, {942, 1435}, {1891, 2820}, {759, 274}, {796, 280}, {580, 1767}, {2929, 1671}, {1072, 2595}, {1952, 583}, {1346, 1962}, {2687, 2356}, {928, 2499}, {2427, 1071}, {1928, 447}, {2967, 1539}, {1901, 2602}, {2233, 1015}, {1104, 166}, {2696, 1837}, {2573, 2289}, {408, 881}, {1574, 1855}, {88, 25}, {1655, 160}, {2745, 2089}, {2155, 903}, {360, 1453}, {2967, 906}, {20, 1333}, {1765, 2285}, {2126, 1236}, {2466, 345}, {1427, 1245}, {1106, 1636}, {2459, 689}, {242, 566}, {2776, 1650}, {2890, 1202}, {1985, 2269}, {2563, 2833}, {2261, 1715}, {559, 2412}, {1371, 1338}, {1686, 708}, {1808, 29}, {944, 1019}, {852, 1661}, {2098, 2032}, {1787, 867}, {2477, 608}, {2098, 279}, {400, 1405}, {2170, 1705}, {1135, 836}, {870, 2092}, {1339, 2071}, {919, 1279}, {1704, 1879}, {686, 2482}, {1500, 701}, {239, 2148}, {2324, 2179}, {2560, 1501}, {117, 2421}, {298, 2481}, {2664, 769}, {752, 608}, {2647, 2313}, {1490, 734}, {48, 2751}, {1806, 1921}, {2756, 315}, {2105, 2459}, {1062, 813}, {364, 2856}, {2689, 407}, {1187, 2580}, {2023, 2663}, {682, 2132}, {1280, 2985}, {2356, 1307}, {1602, 48}, {173, 452}, {2337, 942}, {1106, 1318}, {597, 813}, {1446, 1973}, {2846, 904}, {730, 2033}, {950, 1240}, {620, 2804}, {395, 5}, {1212, 1391}, {1581, 2961}, {1317, 1943}, {2930, 1003}, {805, 707}, {277, 1946}, {775, 1007}, {2654, 2254}, {2885, 1105}, {1154, 20}, {578, 2710}, {2114, 2627}, {80, 1711}, {1320, 1423}, {1728, 288}, {1220, 388}, {2377, 2076}, {817, 415}, {797, 2945}, {1653, 874}, {1882, 352}, {632, 2710}, {1091, 369}, {1162, 1609}, {162, 688}, {1231, 1717}, {838, 1097}, {1270, 1538}, {1313, 2652}, {530, 2265}, {1408, 979}, {301, 1822}, {2754, 2531}, {904, 161}, {1201, 957}, {747, 1066}, {2248, 2394}, {2501, 420}, {2215, 762}, {2665, 2095}, {1311, 351}, {2478, 2880}, {708, 2645}, {24, 2614}, {1625, 828}, {1992, 717}, {889, 1540}, {241, 1249}, {971, 2074}, {550, 2130}, {2633, 1640}, {2726, 575}, {1962, 783}, {2649, 514}, {21, 1745}, {2892, 1890}, {2638, 330}, {1344, 2413}, {210, 2941}, {2819, 2809}, {2241, 2336}, {1592, 606}, {2706, 365}, {1046, 91}, {2342, 2389}, {224, 844}, {1159, 873}, {1584, 1361}, {2224, 1195}, {101, 235}, {2009, 1089}, {1812, 740}, {1002, 905}, {759, 2527}, {1879, 1692}, {1476, 145}, {791, 2624}, {992, 1226}, {1864, 1862}, {481, 843}, {2830, 609}, {2520, 2309}, {2716, 1251}, {867, 2325}, {1507, 296}, {1423, 1874}, {2869, 1931}, {1684, 978}, {2216, 633}, {381, 2781}, {2238, 1070}, {1048, 1743}, {2931, 239}, {960, 2239}, {2803, 2197}, {1065, 288}, {1263, 1613}, {1052, 2334}, {2594, 1928}, {1880, 731}, {919, 253}, {458, 2380}, {2995, 1935}, {1614, 1696}, {899, 113}, {2489, 2550}, {159, 1582}, {1597, 1843}, {2008, 459}, {2270, 1032}, {1522, 2951}, {1201, 2104}, {1626, 1259}, {943, 668}, {756, 2739}, {1954, 1445}, {1079, 1969}, {614, 1664}, {1437, 555}, {1169, 1237}, {1356, 701}, {2253, 1280}, {42, 703}, {36, 2410}, {887, 1361}, {1027, 2880}, {2835, 986}, {503, 1576}, {2059, 1524}, {1868, 1080}, {1371, 1693}, {2330, 2421}, {829, 1225}, {2836, 1127}, {2473, 1387}, {1431, 1792}, {1551, 1696}, {1354, 2738}, {2304, 1674}, {1707, 1536}, {382, 1494}, {99, 220}, {2832, 271}, {682, 1275}, {840, 1210}, {5, 2820}, {1615, 1184}, {959, 905}, {2853, 1400}, {863, 1659}, {50, 2878}, {1065, 2297}, {2818, 1648}, {2791, 2965}, {1049, 2712}, {1101, 1148}, {2140, 1633}, {1799, 1963}, {1361, 2671}, {2832, 1739}, {2468, 1363}, {2129, 2101}, {1631, 249}, {2953, 58}, {2888, 1897}, {504, 1586}, {2671, 318}, {160, 1234}, {245, 2858}, {227, 1050}, {180, 2340}, {980, 2320}, {1897, 2317}, {2735, 2166}, {2559, 2298}, {1767, 1904}, {1769, 2460}, {2953, 357}, {2548, 1841}, {1925, 2654}, {2888, 1166}, {814, 2826}, {1434, 2871}, {1949, 2558}, {1726, 2570}, {373, 1322}, {181, 108}, {79, 241}, {18, 1876}, {654, 512}, {2648, 105}, {1289, 1325}, {2679, 692}, {2139, 1187}, {1942, 629}, {1434, 788}, {2963, 2768}, {645, 2388}, {2841, 2644}, {693, 1025}, {89, 2428}, {2202, 893}, {2286, 2523}, {1492, 2991}, {1090, 59}, {806, 2384}, {1681, 644}, {156, 785}, {13, 184}, {775, 2648}, {353, 2961}, {2440, 596}, {1297, 2608}, {1829, 358}, {2893, 2372}, {2322, 1979}, {592, 1953}, {2543, 267}, {827, 970}, {1926, 2840}, {1041, 400}, {1098, 323}, {1281, 1141}, {40, 1635}, {588, 609}, {1067, 2920}, {2597, 1293}, {298, 2586}, {1450, 1495}, {1375, 2238}, {97, 459}, {742, 1641}, {2652, 2923}, {2343, 1798}, {796, 1344}, {1103, 2592}, {2314, 1836}, {388, 1371}, {1126, 1731}, {176, 949}, {115, 2420}, {435, 2632}, {2243, 515}, {2689, 1527}, {422, 569}, {406, 1426}, {2039, 2458}, {781, 1278}, {1557, 913}, {969, 1382}, {18, 1942}, {832, 575}, {2952, 2348}, {1162, 2023}, {2508, 651}, {2494, 1182}, {2626, 1962}, {1253, 2288}, {1840, 137}, {1709, 2308}, {121, 1287}, {371, 2745}, {2458, 178}, {1470, 1691}, {1994, 1061}, {288, 2325}},
		262,
	},

	// 可以有多个 testcase
}

func Test_removeStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, removeStones(tc.stones), "输入:%v", tc)
	}
}

func Benchmark_removeStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeStones(tc.stones)
		}
	}
}
