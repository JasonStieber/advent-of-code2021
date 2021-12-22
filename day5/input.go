package main

type cord struct {
	x, y int
}

type line struct {
	start, end cord
}

var example = []line{
	{cord{0, 9}, cord{5, 9}},
	{cord{8, 0}, cord{0, 8}},
	{cord{9, 4}, cord{3, 4}},
	{cord{2, 2}, cord{2, 1}},
	{cord{7, 0}, cord{7, 4}},
	{cord{6, 4}, cord{2, 0}},
	{cord{0, 9}, cord{2, 9}},
	{cord{3, 4}, cord{1, 4}},
	{cord{0, 0}, cord{8, 8}},
	{cord{5, 5}, cord{8, 2}},
}

var lines = []line{
	{cord{419, 207}, cord{419, 109}},
	{cord{300, 888}, cord{803, 385}},
	{cord{104, 959}, cord{457, 959}},
	{cord{987, 951}, cord{987, 385}},
	{cord{173, 602}, cord{919, 602}},
	{cord{173, 70}, cord{305, 70}},
	{cord{341, 19}, cord{486, 19}},
	{cord{128, 579}, cord{128, 100}},
	{cord{210, 867}, cord{969, 867}},
	{cord{880, 493}, cord{880, 58}},
	{cord{937, 831}, cord{131, 25}},
	{cord{520, 921}, cord{476, 965}},
	{cord{760, 147}, cord{461, 147}},
	{cord{646, 108}, cord{646, 27}},
	{cord{99, 906}, cord{99, 591}},
	{cord{19, 956}, cord{19, 273}},
	{cord{89, 201}, cord{326, 201}},
	{cord{275, 948}, cord{962, 261}},
	{cord{292, 489}, cord{689, 489}},
	{cord{674, 109}, cord{20, 763}},
	{cord{861, 861}, cord{529, 529}},
	{cord{155, 200}, cord{273, 200}},
	{cord{628, 803}, cord{209, 384}},
	{cord{654, 401}, cord{578, 325}},
	{cord{723, 625}, cord{828, 730}},
	{cord{137, 406}, cord{862, 406}},
	{cord{893, 45}, cord{41, 897}},
	{cord{631, 10}, cord{941, 320}},
	{cord{618, 435}, cord{537, 435}},
	{cord{939, 29}, cord{30, 938}},
	{cord{505, 796}, cord{505, 244}},
	{cord{799, 779}, cord{77, 779}},
	{cord{938, 576}, cord{427, 576}},
	{cord{522, 635}, cord{405, 518}},
	{cord{244, 89}, cord{946, 89}},
	{cord{447, 791}, cord{316, 660}},
	{cord{560, 731}, cord{687, 731}},
	{cord{16, 878}, cord{835, 59}},
	{cord{45, 707}, cord{45, 565}},
	{cord{767, 166}, cord{404, 529}},
	{cord{791, 260}, cord{791, 950}},
	{cord{373, 949}, cord{373, 156}},
	{cord{38, 774}, cord{38, 557}},
	{cord{445, 537}, cord{445, 370}},
	{cord{817, 756}, cord{959, 898}},
	{cord{472, 551}, cord{952, 71}},
	{cord{696, 381}, cord{657, 420}},
	{cord{43, 829}, cord{43, 190}},
	{cord{101, 635}, cord{728, 635}},
	{cord{197, 532}, cord{140, 532}},
	{cord{693, 368}, cord{299, 368}},
	{cord{433, 140}, cord{433, 610}},
	{cord{136, 58}, cord{136, 666}},
	{cord{472, 294}, cord{886, 294}},
	{cord{690, 883}, cord{671, 864}},
	{cord{141, 598}, cord{141, 118}},
	{cord{56, 651}, cord{56, 957}},
	{cord{747, 82}, cord{747, 91}},
	{cord{219, 455}, cord{55, 291}},
	{cord{444, 131}, cord{444, 802}},
	{cord{326, 459}, cord{528, 661}},
	{cord{245, 965}, cord{143, 965}},
	{cord{916, 316}, cord{630, 316}},
	{cord{263, 55}, cord{977, 769}},
	{cord{262, 451}, cord{587, 451}},
	{cord{960, 178}, cord{960, 564}},
	{cord{960, 88}, cord{476, 572}},
	{cord{314, 259}, cord{314, 169}},
	{cord{404, 742}, cord{429, 742}},
	{cord{830, 921}, cord{409, 921}},
	{cord{181, 396}, cord{463, 678}},
	{cord{338, 293}, cord{23, 608}},
	{cord{851, 667}, cord{851, 350}},
	{cord{181, 859}, cord{718, 322}},
	{cord{314, 240}, cord{870, 796}},
	{cord{778, 984}, cord{77, 283}},
	{cord{476, 178}, cord{440, 178}},
	{cord{935, 357}, cord{841, 263}},
	{cord{695, 683}, cord{414, 964}},
	{cord{760, 241}, cord{306, 241}},
	{cord{390, 355}, cord{791, 355}},
	{cord{460, 710}, cord{851, 710}},
	{cord{559, 448}, cord{870, 448}},
	{cord{161, 526}, cord{301, 386}},
	{cord{935, 495}, cord{633, 193}},
	{cord{205, 536}, cord{383, 536}},
	{cord{290, 626}, cord{290, 94}},
	{cord{55, 972}, cord{946, 81}},
	{cord{240, 531}, cord{631, 922}},
	{cord{189, 806}, cord{573, 806}},
	{cord{518, 827}, cord{866, 479}},
	{cord{239, 829}, cord{260, 829}},
	{cord{151, 51}, cord{849, 51}},
	{cord{301, 736}, cord{532, 736}},
	{cord{23, 889}, cord{336, 889}},
	{cord{284, 124}, cord{284, 933}},
	{cord{637, 610}, cord{67, 40}},
	{cord{610, 828}, cord{610, 159}},
	{cord{763, 590}, cord{763, 963}},
	{cord{804, 576}, cord{804, 694}},
	{cord{689, 872}, cord{82, 265}},
	{cord{440, 377}, cord{190, 127}},
	{cord{933, 330}, cord{310, 953}},
	{cord{873, 99}, cord{873, 328}},
	{cord{756, 808}, cord{860, 808}},
	{cord{119, 64}, cord{928, 873}},
	{cord{74, 144}, cord{489, 559}},
	{cord{957, 938}, cord{838, 938}},
	{cord{148, 320}, cord{932, 320}},
	{cord{386, 171}, cord{386, 985}},
	{cord{357, 171}, cord{494, 171}},
	{cord{254, 67}, cord{254, 95}},
	{cord{196, 910}, cord{827, 910}},
	{cord{107, 114}, cord{758, 114}},
	{cord{971, 40}, cord{801, 40}},
	{cord{504, 602}, cord{215, 891}},
	{cord{184, 310}, cord{720, 846}},
	{cord{280, 300}, cord{955, 975}},
	{cord{49, 637}, cord{49, 572}},
	{cord{352, 512}, cord{739, 899}},
	{cord{610, 123}, cord{585, 123}},
	{cord{808, 881}, cord{758, 881}},
	{cord{646, 980}, cord{818, 980}},
	{cord{948, 482}, cord{384, 482}},
	{cord{115, 144}, cord{852, 881}},
	{cord{506, 836}, cord{547, 836}},
	{cord{985, 369}, cord{374, 980}},
	{cord{883, 975}, cord{48, 975}},
	{cord{447, 664}, cord{312, 799}},
	{cord{24, 597}, cord{24, 331}},
	{cord{45, 19}, cord{979, 953}},
	{cord{210, 689}, cord{210, 430}},
	{cord{704, 806}, cord{704, 612}},
	{cord{985, 982}, cord{124, 121}},
	{cord{70, 174}, cord{550, 174}},
	{cord{463, 12}, cord{637, 12}},
	{cord{107, 97}, cord{716, 97}},
	{cord{935, 265}, cord{390, 810}},
	{cord{42, 223}, cord{42, 86}},
	{cord{60, 245}, cord{60, 473}},
	{cord{695, 735}, cord{208, 735}},
	{cord{547, 265}, cord{802, 265}},
	{cord{941, 667}, cord{941, 806}},
	{cord{250, 286}, cord{611, 286}},
	{cord{10, 64}, cord{630, 64}},
	{cord{482, 889}, cord{482, 150}},
	{cord{441, 820}, cord{776, 820}},
	{cord{529, 474}, cord{529, 265}},
	{cord{533, 465}, cord{217, 149}},
	{cord{242, 473}, cord{242, 830}},
	{cord{633, 160}, cord{476, 317}},
	{cord{942, 24}, cord{942, 784}},
	{cord{80, 313}, cord{92, 325}},
	{cord{295, 109}, cord{295, 712}},
	{cord{31, 964}, cord{857, 138}},
	{cord{285, 255}, cord{955, 925}},
	{cord{650, 610}, cord{650, 366}},
	{cord{722, 586}, cord{625, 586}},
	{cord{580, 384}, cord{580, 531}},
	{cord{78, 407}, cord{896, 407}},
	{cord{296, 310}, cord{730, 744}},
	{cord{717, 966}, cord{924, 966}},
	{cord{524, 551}, cord{524, 671}},
	{cord{44, 127}, cord{784, 867}},
	{cord{214, 849}, cord{238, 849}},
	{cord{749, 320}, cord{749, 241}},
	{cord{886, 146}, cord{336, 696}},
	{cord{889, 933}, cord{455, 499}},
	{cord{644, 232}, cord{79, 797}},
	{cord{400, 979}, cord{626, 979}},
	{cord{433, 681}, cord{433, 523}},
	{cord{447, 57}, cord{676, 57}},
	{cord{185, 416}, cord{659, 890}},
	{cord{849, 645}, cord{257, 53}},
	{cord{633, 721}, cord{633, 901}},
	{cord{766, 355}, cord{766, 56}},
	{cord{669, 393}, cord{669, 523}},
	{cord{833, 336}, cord{833, 58}},
	{cord{52, 114}, cord{52, 413}},
	{cord{699, 957}, cord{109, 957}},
	{cord{14, 953}, cord{945, 22}},
	{cord{641, 15}, cord{929, 303}},
	{cord{25, 874}, cord{866, 33}},
	{cord{856, 73}, cord{28, 901}},
	{cord{94, 892}, cord{592, 892}},
	{cord{256, 357}, cord{256, 700}},
	{cord{960, 579}, cord{31, 579}},
	{cord{940, 859}, cord{940, 987}},
	{cord{507, 673}, cord{820, 986}},
	{cord{164, 361}, cord{133, 361}},
	{cord{210, 424}, cord{876, 424}},
	{cord{28, 186}, cord{28, 376}},
	{cord{452, 149}, cord{531, 149}},
	{cord{142, 160}, cord{142, 435}},
	{cord{180, 801}, cord{180, 439}},
	{cord{681, 267}, cord{42, 267}},
	{cord{724, 414}, cord{786, 476}},
	{cord{762, 492}, cord{762, 427}},
	{cord{902, 808}, cord{227, 133}},
	{cord{70, 923}, cord{821, 172}},
	{cord{468, 12}, cord{457, 12}},
	{cord{208, 129}, cord{986, 907}},
	{cord{78, 786}, cord{78, 352}},
	{cord{573, 869}, cord{820, 869}},
	{cord{780, 680}, cord{520, 940}},
	{cord{276, 66}, cord{276, 244}},
	{cord{423, 629}, cord{592, 629}},
	{cord{888, 507}, cord{888, 139}},
	{cord{869, 878}, cord{869, 951}},
	{cord{274, 614}, cord{625, 965}},
	{cord{926, 289}, cord{982, 233}},
	{cord{102, 687}, cord{102, 214}},
	{cord{52, 264}, cord{52, 12}},
	{cord{904, 43}, cord{657, 43}},
	{cord{184, 685}, cord{184, 628}},
	{cord{506, 912}, cord{601, 817}},
	{cord{356, 524}, cord{87, 524}},
	{cord{202, 260}, cord{202, 276}},
	{cord{970, 63}, cord{83, 950}},
	{cord{402, 332}, cord{950, 880}},
	{cord{195, 666}, cord{843, 666}},
	{cord{13, 82}, cord{892, 961}},
	{cord{614, 28}, cord{614, 871}},
	{cord{892, 162}, cord{892, 101}},
	{cord{363, 665}, cord{59, 665}},
	{cord{768, 208}, cord{410, 208}},
	{cord{483, 300}, cord{295, 300}},
	{cord{590, 108}, cord{881, 108}},
	{cord{837, 967}, cord{837, 326}},
	{cord{368, 731}, cord{368, 913}},
	{cord{900, 921}, cord{873, 921}},
	{cord{896, 931}, cord{848, 979}},
	{cord{562, 939}, cord{857, 939}},
	{cord{85, 351}, cord{598, 351}},
	{cord{917, 30}, cord{624, 30}},
	{cord{605, 314}, cord{605, 303}},
	{cord{382, 655}, cord{340, 697}},
	{cord{949, 115}, cord{653, 115}},
	{cord{667, 311}, cord{370, 608}},
	{cord{971, 983}, cord{39, 51}},
	{cord{178, 687}, cord{178, 69}},
	{cord{906, 197}, cord{296, 807}},
	{cord{947, 886}, cord{383, 322}},
	{cord{551, 667}, cord{551, 238}},
	{cord{86, 65}, cord{916, 895}},
	{cord{589, 887}, cord{865, 611}},
	{cord{332, 53}, cord{84, 53}},
	{cord{361, 148}, cord{55, 148}},
	{cord{883, 205}, cord{661, 205}},
	{cord{415, 552}, cord{52, 552}},
	{cord{46, 42}, cord{46, 952}},
	{cord{955, 13}, cord{39, 929}},
	{cord{677, 482}, cord{208, 482}},
	{cord{414, 268}, cord{927, 268}},
	{cord{101, 509}, cord{101, 149}},
	{cord{946, 971}, cord{139, 164}},
	{cord{223, 597}, cord{223, 517}},
	{cord{805, 896}, cord{796, 896}},
	{cord{565, 875}, cord{878, 875}},
	{cord{472, 431}, cord{472, 732}},
	{cord{643, 15}, cord{643, 202}},
	{cord{618, 725}, cord{618, 284}},
	{cord{376, 532}, cord{376, 120}},
	{cord{807, 981}, cord{415, 981}},
	{cord{716, 401}, cord{61, 401}},
	{cord{893, 955}, cord{743, 805}},
	{cord{264, 935}, cord{264, 677}},
	{cord{586, 908}, cord{638, 908}},
	{cord{780, 277}, cord{780, 418}},
	{cord{234, 410}, cord{428, 410}},
	{cord{899, 214}, cord{899, 703}},
	{cord{948, 51}, cord{948, 509}},
	{cord{238, 664}, cord{879, 23}},
	{cord{20, 877}, cord{638, 877}},
	{cord{146, 738}, cord{109, 738}},
	{cord{670, 893}, cord{524, 893}},
	{cord{317, 423}, cord{27, 713}},
	{cord{91, 600}, cord{477, 986}},
	{cord{902, 63}, cord{902, 797}},
	{cord{647, 839}, cord{647, 667}},
	{cord{227, 358}, cord{236, 349}},
	{cord{985, 541}, cord{660, 866}},
	{cord{86, 562}, cord{86, 949}},
	{cord{368, 851}, cord{863, 356}},
	{cord{327, 905}, cord{57, 635}},
	{cord{561, 275}, cord{781, 495}},
	{cord{196, 65}, cord{626, 65}},
	{cord{110, 688}, cord{720, 78}},
	{cord{720, 472}, cord{115, 472}},
	{cord{817, 135}, cord{817, 876}},
	{cord{752, 387}, cord{752, 104}},
	{cord{78, 127}, cord{635, 684}},
	{cord{812, 170}, cord{155, 170}},
	{cord{606, 718}, cord{804, 916}},
	{cord{843, 494}, cord{979, 494}},
	{cord{919, 346}, cord{454, 346}},
	{cord{866, 828}, cord{818, 828}},
	{cord{114, 115}, cord{114, 250}},
	{cord{895, 308}, cord{370, 308}},
	{cord{665, 893}, cord{690, 893}},
	{cord{939, 275}, cord{741, 275}},
	{cord{290, 321}, cord{290, 910}},
	{cord{747, 327}, cord{107, 967}},
	{cord{734, 715}, cord{391, 372}},
	{cord{368, 497}, cord{506, 359}},
	{cord{773, 945}, cord{391, 563}},
	{cord{772, 537}, cord{733, 537}},
	{cord{271, 679}, cord{488, 679}},
	{cord{665, 745}, cord{665, 984}},
	{cord{143, 177}, cord{685, 719}},
	{cord{671, 860}, cord{147, 860}},
	{cord{674, 365}, cord{857, 182}},
	{cord{343, 74}, cord{985, 716}},
	{cord{284, 46}, cord{180, 46}},
	{cord{595, 800}, cord{20, 225}},
	{cord{57, 278}, cord{792, 278}},
	{cord{649, 285}, cord{165, 769}},
	{cord{600, 24}, cord{600, 116}},
	{cord{862, 939}, cord{862, 871}},
	{cord{153, 917}, cord{682, 388}},
	{cord{117, 884}, cord{257, 884}},
	{cord{726, 763}, cord{531, 763}},
	{cord{810, 985}, cord{899, 985}},
	{cord{718, 942}, cord{718, 466}},
	{cord{674, 19}, cord{674, 203}},
	{cord{117, 677}, cord{117, 918}},
	{cord{928, 261}, cord{928, 945}},
	{cord{719, 390}, cord{719, 321}},
	{cord{822, 601}, cord{484, 263}},
	{cord{725, 793}, cord{725, 111}},
	{cord{201, 745}, cord{588, 745}},
	{cord{404, 889}, cord{908, 385}},
	{cord{981, 39}, cord{610, 410}},
	{cord{148, 426}, cord{711, 989}},
	{cord{128, 260}, cord{319, 451}},
	{cord{325, 306}, cord{325, 585}},
	{cord{557, 415}, cord{557, 745}},
	{cord{915, 101}, cord{648, 101}},
	{cord{104, 636}, cord{104, 520}},
	{cord{93, 964}, cord{641, 416}},
	{cord{201, 709}, cord{201, 90}},
	{cord{921, 571}, cord{798, 571}},
	{cord{313, 624}, cord{313, 510}},
	{cord{343, 649}, cord{28, 649}},
	{cord{688, 246}, cord{24, 910}},
	{cord{696, 610}, cord{353, 610}},
	{cord{126, 310}, cord{126, 394}},
	{cord{457, 98}, cord{457, 981}},
	{cord{277, 707}, cord{277, 531}},
	{cord{943, 721}, cord{37, 721}},
	{cord{959, 295}, cord{702, 295}},
	{cord{23, 547}, cord{891, 547}},
	{cord{209, 114}, cord{931, 836}},
	{cord{737, 174}, cord{737, 195}},
	{cord{208, 890}, cord{115, 797}},
	{cord{170, 401}, cord{726, 401}},
	{cord{11, 218}, cord{11, 297}},
	{cord{989, 10}, cord{10, 989}},
	{cord{866, 86}, cord{487, 86}},
	{cord{867, 31}, cord{867, 334}},
	{cord{846, 414}, cord{861, 414}},
	{cord{478, 315}, cord{478, 697}},
	{cord{572, 843}, cord{731, 843}},
	{cord{657, 12}, cord{161, 508}},
	{cord{903, 194}, cord{142, 955}},
	{cord{612, 321}, cord{147, 786}},
	{cord{813, 920}, cord{259, 920}},
	{cord{834, 389}, cord{651, 206}},
	{cord{824, 153}, cord{824, 557}},
	{cord{399, 871}, cord{115, 871}},
	{cord{270, 785}, cord{270, 120}},
	{cord{469, 640}, cord{753, 640}},
	{cord{620, 132}, cord{620, 175}},
	{cord{620, 234}, cord{666, 234}},
	{cord{594, 409}, cord{948, 55}},
	{cord{670, 323}, cord{670, 89}},
	{cord{262, 65}, cord{262, 379}},
	{cord{879, 617}, cord{284, 22}},
	{cord{493, 423}, cord{761, 423}},
	{cord{17, 931}, cord{906, 42}},
	{cord{512, 494}, cord{473, 494}},
	{cord{122, 230}, cord{122, 87}},
	{cord{15, 207}, cord{533, 207}},
	{cord{216, 183}, cord{50, 183}},
	{cord{360, 107}, cord{280, 107}},
	{cord{403, 841}, cord{941, 841}},
	{cord{913, 442}, cord{500, 29}},
	{cord{864, 947}, cord{864, 85}},
	{cord{500, 516}, cord{634, 382}},
	{cord{283, 20}, cord{669, 20}},
	{cord{916, 770}, cord{176, 30}},
	{cord{966, 73}, cord{252, 787}},
	{cord{847, 841}, cord{171, 165}},
	{cord{163, 219}, cord{766, 219}},
	{cord{482, 515}, cord{275, 308}},
	{cord{528, 949}, cord{240, 949}},
	{cord{725, 574}, cord{847, 696}},
	{cord{109, 131}, cord{109, 538}},
	{cord{655, 837}, cord{476, 837}},
	{cord{803, 631}, cord{803, 51}},
	{cord{977, 83}, cord{149, 911}},
	{cord{207, 231}, cord{171, 231}},
	{cord{617, 29}, cord{617, 294}},
	{cord{838, 708}, cord{446, 708}},
	{cord{711, 597}, cord{612, 498}},
	{cord{975, 942}, cord{279, 246}},
	{cord{315, 128}, cord{315, 293}},
	{cord{146, 962}, cord{873, 235}},
	{cord{448, 180}, cord{54, 180}},
	{cord{177, 680}, cord{866, 680}},
	{cord{891, 265}, cord{741, 265}},
	{cord{656, 949}, cord{414, 949}},
	{cord{909, 456}, cord{196, 456}},
	{cord{574, 286}, cord{58, 286}},
	{cord{861, 691}, cord{861, 383}},
	{cord{779, 351}, cord{779, 827}},
	{cord{459, 989}, cord{459, 350}},
	{cord{936, 480}, cord{936, 699}},
	{cord{645, 309}, cord{348, 606}},
	{cord{861, 62}, cord{621, 302}},
	{cord{568, 324}, cord{568, 358}},
	{cord{889, 221}, cord{889, 335}},
	{cord{538, 759}, cord{538, 266}},
	{cord{780, 736}, cord{780, 827}},
	{cord{866, 518}, cord{983, 401}},
	{cord{67, 871}, cord{840, 98}},
	{cord{432, 664}, cord{664, 664}},
	{cord{146, 24}, cord{755, 24}},
	{cord{964, 585}, cord{964, 770}},
	{cord{372, 144}, cord{809, 144}},
	{cord{688, 827}, cord{867, 827}},
	{cord{137, 916}, cord{137, 942}},
	{cord{846, 131}, cord{846, 46}},
	{cord{764, 21}, cord{457, 328}},
	{cord{140, 66}, cord{799, 725}},
	{cord{703, 224}, cord{83, 844}},
	{cord{557, 67}, cord{557, 681}},
	{cord{355, 544}, cord{764, 135}},
	{cord{625, 893}, cord{126, 394}},
	{cord{842, 214}, cord{842, 322}},
	{cord{582, 778}, cord{582, 762}},
	{cord{341, 861}, cord{341, 859}},
	{cord{143, 767}, cord{52, 858}},
	{cord{114, 109}, cord{114, 200}},
	{cord{394, 210}, cord{396, 212}},
	{cord{861, 353}, cord{861, 652}},
	{cord{873, 553}, cord{62, 553}},
	{cord{44, 962}, cord{984, 22}},
	{cord{734, 56}, cord{734, 828}},
	{cord{798, 516}, cord{950, 516}},
	{cord{367, 755}, cord{367, 618}},
	{cord{868, 637}, cord{868, 780}},
	{cord{192, 952}, cord{192, 734}},
	{cord{603, 109}, cord{705, 211}},
	{cord{12, 17}, cord{984, 989}},
	{cord{910, 147}, cord{910, 620}},
	{cord{515, 749}, cord{515, 517}},
	{cord{775, 136}, cord{761, 150}},
	{cord{662, 636}, cord{662, 21}},
	{cord{894, 490}, cord{310, 490}},
	{cord{956, 732}, cord{297, 73}},
	{cord{514, 99}, cord{140, 99}},
	{cord{308, 419}, cord{691, 419}},
	{cord{485, 86}, cord{485, 187}},
	{cord{737, 783}, cord{979, 783}},
	{cord{90, 76}, cord{869, 855}},
	{cord{959, 112}, cord{84, 112}},
	{cord{879, 494}, cord{879, 257}},
	{cord{933, 425}, cord{933, 619}},
	{cord{64, 391}, cord{64, 21}},
	{cord{106, 305}, cord{253, 452}},
	{cord{324, 152}, cord{853, 152}},
	{cord{666, 225}, cord{39, 852}},
	{cord{370, 904}, cord{257, 791}},
	{cord{592, 845}, cord{592, 15}},
	{cord{936, 971}, cord{267, 302}},
	{cord{147, 210}, cord{62, 210}},
	{cord{308, 323}, cord{495, 323}},
	{cord{212, 918}, cord{110, 918}},
	{cord{229, 392}, cord{685, 848}},
	{cord{896, 132}, cord{326, 702}},
	{cord{483, 143}, cord{605, 265}},
	{cord{251, 317}, cord{130, 317}},
	{cord{758, 93}, cord{445, 93}},
	{cord{156, 286}, cord{458, 286}},
	{cord{401, 904}, cord{383, 904}},
	{cord{244, 256}, cord{851, 256}},
	{cord{928, 411}, cord{612, 411}},
	{cord{642, 920}, cord{642, 420}},
	{cord{494, 707}, cord{494, 225}},
	{cord{87, 112}, cord{87, 256}},
	{cord{972, 907}, cord{83, 18}},
	{cord{139, 104}, cord{139, 761}},
	{cord{493, 725}, cord{493, 529}},
	{cord{981, 145}, cord{459, 667}},
	{cord{390, 240}, cord{702, 240}},
	{cord{466, 982}, cord{807, 982}},
	{cord{320, 143}, cord{692, 515}},
	{cord{477, 649}, cord{477, 206}},
	{cord{456, 254}, cord{456, 578}},
}