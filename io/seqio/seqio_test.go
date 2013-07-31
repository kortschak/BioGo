// Copyright ©2011-2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seqio_test

import (
	"bytes"
	"code.google.com/p/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/io/seqio/fastq"
	"code.google.com/p/biogo/seq/linear"

	check "launchpad.net/gocheck"
	"testing"
)

func TestSeqio(t *testing.T) {
	var (
		_ seqio.Reader = (*fasta.Reader)(nil)
		_ seqio.Reader = (*fastq.Reader)(nil)
		_ seqio.Writer = (*fasta.Writer)(nil)
		_ seqio.Writer = (*fastq.Writer)(nil)
	)
}

// Tests
func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

var (
	testaln0 = `>AK1H_ECOLI/114-431 DESCRIPTION HERE
CPDSINAALICRGEKMSIAIMAGVLEARGH-N--VTVIDPVEKLLAVG-HYLESTVDIAE
STRRIAASRIP------A-DHMVLMAGFTAGN-EKGELVVLGRNGSDYSAAVLAACLRAD
CCEIWTDVNGVYTCDP-------------RQVPDARLLKSMSYQEAMELSY--FGAKVLH
PRTITPIAQFQIPCLIKNTGNPQAPGTL-IG--ASRDEDELP----VKGISNLN------
NMAMFSVSGP-GMKGMVGMAARVFAAMS-------RARISVVLITQSSSEYSISFCVPQS
DCVRAERAMLEEFY-----LELKEGLLEPLAVAERLAIISV-VGDGLRTLRGISAKF---
---FAALARANINIVAIA
>AKH_HAEIN 114-431
-----------------VEDAVKATIDCRGEKLSIAMMKAWFEARGY-S--VHIVDPVKQ
LLAKG-GYLESSVEIEESTKRVDAANIA--K-DKVVLMAGF---TAGNEKGELVLLGRNG
SDYSAAC-----------------LAACLGASVCEIWTDVDGVYTCDP--RLVPDARLLP
TLSYREAMELSYFGAKVIHPRTIGPLLPQNIPCVIKNTGNPSAPGSI-ID--GNVKSESL
Q----VKGITNLDNLAMFNVSGPGMQGM---VGMASRVFSAMSGAGISVILITQSSSEYS
---ISFCVPVKSAEVAKTVLETEFA-----NELNEHQLEPIEVIKDLSIISV-VGDGMKQ
AKGIAARF------FSALAQANISIVAIA
>AKH1_MAIZE/117-440
-----------------ATESFSDFVVGHGELWSAQMLSYAIQKSGT-P--CSWMDTREV
LVVNPSGANQVDPDYLESEKRLEKWFSRC-P-AETIIATGF---IASTPENIPTTLKRDG
SDFSAAI-----------------IGSLVKARQVTIWTDVDGVFSADP--RKVSEAVILS
TLSYQEAWEMSYFGANVLHPRTIIPVMKYNIPIVIRNIFNTSAPGTM-IC--QQPANENG
DLEACVKAFATIDKLALVNVEGTGMAGV---PGTANAIFGAVKDVGANVIMISQASSEHS
---VCFAVPEKEVALVSAALHARFR-----EALAAGRLSKVEVIHNCSILAT-VGLRMAS
TPGVSATL------FDALAKANINVRAIA
>AK2H_ECOLI/112-431
-----------------INDAVYAEVVGHGEVWSARLMSAVLNQQG-----LPAAWLDAR
EFLRAERAAQPQVDEGLSYPLLQQLLVQH-P-GKRLVVTGF---ISRNNAGETVLLGRNG
SDYSATQ-----------------IGALAGVSRVTIWSDVAGVYSADP--RKVKDACLLP
LLRLDEASELARLAAPVLHARTLQPVSGSEIDLQLRCSYTPDQGSTRIERVLASGTGARI
VTSHDDVCLI-EFQVPASQDFKLAHKEI--DQILKRAQVRPLAVGVHNDRQLLQFCYTSE
VADSALKILDEAG---------LPGELRLRQGLALVAMVGAGVTRNPLHCHRFWQQLKGQ
PVEFTWQSDDGISLVAVL
>AK1_BACSU/66-374
-----------------ISPREQDLLLSCGETISSVVFTSMLLDNGVKA--AALTGAQAG
FLTNDQHTNAKIIEMKPER--LFSVLAN----HDAVVVAGF---QGATEKGDTTTIGRGG
SDTSAAA-----------------LGAAVDAEYIDIFTDVEGVMTADP--RVVENAKPLP
VVTYTEICNLAYQGAKVISPRAVEIAMQAKVPIRVRSTYS-NDKGTLVTSHHSSKVGSDV
FERLITGIAH-VKDVTQFKVPAKIGQYN-----VQTEVFKAMANAGISVDFFNITPSEIV
YTVAGNKTETAQR------------ILMDMGYDPMVTRNCAKVSAVGAGIMGVPGVTSKI
------VSALSEKEIPILQSA
>AK2_BACST/63-370
-----------------KRE--MDMLLSTGEQVSIALLAMSLHEKGYKA--VSLTGWQAG
ITTEEMHGNARIMNIDTT--RIRRCLDE----GAIVIVAGF---QGVTETGEITTLGRGG
SDTTAVA-----------------LAAALKAEKCDIYTDVTGVFTTDP--RYVKTARKIK
EISYDEMLELANLGAGVLHPRAVEFAKNYEVPLEVRSSME-NERGTMVK--EEVSMEQHL
IVRGIAFEDQ-VTRVTVVGIEKYLQSVA--------TIFTALANRGINVDIIIQNA----
----------------TNSETAS--VSFSIRTEDLPETLQVLQ-------------ALEG
ADVHYESGLAKVSI-VGSGMISNPGVAARV------FEVLADQGIEIKMVS
>AK2_BACSU/63-373
-----------------KRE--MDMLLATGEQVTISLLSMALQEKGYDA--VSYTGWQAG
IRTEAIHGNARITDIDTS--VLADQLEK----GKIVIVAGF---QGMTEDCEITTLGRGG
SDTTAVA-----------------LAAALKVDKCDIYTDVPGVFTTDP--RYVKSARKLE
GISYDEMLELANLGAGVLHPRAVEFAKNYQVPLEVRSSTE-TEAGTLIE--EESSMEQNL
IVRGIAFEDQ-ITRVTIYGLTSGLTTLS--------TIFTTLAKRNINVDIIIQTQ----
----------------AEDKTG---ISFSVKTEDADQTVAVLEEYK---------DALEF
EKIETESKLAKVSI-VGSGMVSNPGVAAEM------FAVLAQKNILIKMVS
>AKAB_CORFL/63-379
-----------------ARE--MDMLLTAGERISNALVAMAIESLGAEA--QSFTGSQAG
VLTTERHGNARIVDVTPG--RVREALDE----GKICIVAGF--QGVNKETRDVTTLGRGG
SDTTAVA-----------------LAAALNADVCEIYSDVDGVYTADP--RIVPNAQKLE
KLSFEEMLELAAVGSKILVLRSVEYARAFNVPLRVRSSYS-NDPGTLIAGSMEDIPVEEA
VLTGVATDKS-EAKVTVLGISDKPGEAA--------KVFRALADAEINIDMVLQNV----
----------------SSVEDGTTDITFTCPRADGRRAMEILKKLQ---------VQGNW
TNVLYDDQVDKVSL-VGAGMKSHPGVTAEF------MEALRDVNVNIELIS
>AKAB_MYCSM/63-379
-----------------PRE--MDMLLTAGERISNALVAMAIESLGAQA--RSFTGSQAG
VITTGTHGNAKIIDVTPG--RLRDALDE----GQIVLVAGF--QGVSQDSKDVTTLGRGG
SDTTAVA-----------------VAAALDADVCEIYTDVDGIFTADP--RIVPNARHLD
TVSFEEMLEMAACGAKVLMLRCVEYARRYNVPIHVRSSYS-DKPGTIVKGSIEDIPMEDA
ILTGVAHDRS-EAKVTVVGLPDVPGYAA--------KVFRAVAEADVNIDMVLQNI----
----------------SKIEDGKTDITFTCARDNGPRAVEKLSALK---------SEIGF
SQVLYDDHIGKVSL-IGAGMRSHPGVTATF------CEALAEAGINIDLIS
>AK3_ECOLI/106-407
-----------------TSPALTDELVSHGELMSTLLFVEILRERD--V--QAQWFDVRK
VMRTNDRFGRAEPDIAALAELAALQLLPR-LNEGLVITQGF---IGSENKGRTTTLGRGG
SDYTAAL-----------------LAEALHASRVDIWTDVPGIYTTDP--RVVSAAKRID
EIAFAEAAEMATFGAKVLHPATLLPAVRSDIPVFVGSSKDPRAGGTLVCNKTENPPLFRA
LAL--RRNQT-LLTLHSLNMLHSRGFLA--------EVFGILARHNISVDLITTSEVSVA
LTLDTTGSTSTG----------DTLLTQSLLMELSALCRVEVEEGLALVALIG-------
---NDLSKACGVGKEVF
>AK_YEAST/134-472 A COMMENT FOR YEAST
-----------------VSSRTVDLVMSCGEKLSCLFMTALCNDRGCKAKYVDLSHIVPS
DFSASALDNSFYTFLVQALKEKLAPFVSA-KERIVPVFTGF---FGLVPTGLLNGVGRGY
TDLCAAL-----------------IAVAVNADELQVWKEVDGIFTADP--RKVPEARLLD
SVTPEEASELTYYGSEVIHPFTMEQVIRAKIPIRIKNVQNPLGNGTIIYPDNVAKKGEST
PPHPPENLSS----SFYEKRKRGATAITTKN----DIFVINIHSNKKTLSHGFLAQIFTI
LDKYKLVVDLISTSEVHVSMALPIPDADS-LKSLRQAEEKLRILGSVDITKKLSIVSLVG
KHMKQYIGIAG---TMFTTLAEEGINIEMIS
`

	expectNfa = []string{
		"AK1H_ECOLI/114-431 DESCRIPTION HERE",
		"AKH_HAEIN 114-431",
		"AKH1_MAIZE/117-440",
		"AK2H_ECOLI/112-431",
		"AK1_BACSU/66-374",
		"AK2_BACST/63-370",
		"AK2_BACSU/63-373",
		"AKAB_CORFL/63-379",
		"AKAB_MYCSM/63-379",
		"AK3_ECOLI/106-407",
		"AK_YEAST/134-472 A COMMENT FOR YEAST",
	}

	expectSfa = [][]alphabet.Letter{
		[]alphabet.Letter("CPDSINAALICRGEKMSIAIMAGVLEARGH-N--VTVIDPVEKLLAVG-HYLESTVDIAESTRRIAASRIP------A-DHMVLMAGFTAGN-EKGELVVLGRNGSDYSAAVLAACLRADCCEIWTDVNGVYTCDP-------------RQVPDARLLKSMSYQEAMELSY--FGAKVLHPRTITPIAQFQIPCLIKNTGNPQAPGTL-IG--ASRDEDELP----VKGISNLN------NMAMFSVSGP-GMKGMVGMAARVFAAMS-------RARISVVLITQSSSEYSISFCVPQSDCVRAERAMLEEFY-----LELKEGLLEPLAVAERLAIISV-VGDGLRTLRGISAKF------FAALARANINIVAIA"),
		[]alphabet.Letter("-----------------VEDAVKATIDCRGEKLSIAMMKAWFEARGY-S--VHIVDPVKQLLAKG-GYLESSVEIEESTKRVDAANIA--K-DKVVLMAGF---TAGNEKGELVLLGRNGSDYSAAC-----------------LAACLGASVCEIWTDVDGVYTCDP--RLVPDARLLPTLSYREAMELSYFGAKVIHPRTIGPLLPQNIPCVIKNTGNPSAPGSI-ID--GNVKSESLQ----VKGITNLDNLAMFNVSGPGMQGM---VGMASRVFSAMSGAGISVILITQSSSEYS---ISFCVPVKSAEVAKTVLETEFA-----NELNEHQLEPIEVIKDLSIISV-VGDGMKQAKGIAARF------FSALAQANISIVAIA"),
		[]alphabet.Letter("-----------------ATESFSDFVVGHGELWSAQMLSYAIQKSGT-P--CSWMDTREVLVVNPSGANQVDPDYLESEKRLEKWFSRC-P-AETIIATGF---IASTPENIPTTLKRDGSDFSAAI-----------------IGSLVKARQVTIWTDVDGVFSADP--RKVSEAVILSTLSYQEAWEMSYFGANVLHPRTIIPVMKYNIPIVIRNIFNTSAPGTM-IC--QQPANENGDLEACVKAFATIDKLALVNVEGTGMAGV---PGTANAIFGAVKDVGANVIMISQASSEHS---VCFAVPEKEVALVSAALHARFR-----EALAAGRLSKVEVIHNCSILAT-VGLRMASTPGVSATL------FDALAKANINVRAIA"),
		[]alphabet.Letter("-----------------INDAVYAEVVGHGEVWSARLMSAVLNQQG-----LPAAWLDAREFLRAERAAQPQVDEGLSYPLLQQLLVQH-P-GKRLVVTGF---ISRNNAGETVLLGRNGSDYSATQ-----------------IGALAGVSRVTIWSDVAGVYSADP--RKVKDACLLPLLRLDEASELARLAAPVLHARTLQPVSGSEIDLQLRCSYTPDQGSTRIERVLASGTGARIVTSHDDVCLI-EFQVPASQDFKLAHKEI--DQILKRAQVRPLAVGVHNDRQLLQFCYTSEVADSALKILDEAG---------LPGELRLRQGLALVAMVGAGVTRNPLHCHRFWQQLKGQPVEFTWQSDDGISLVAVL"),
		[]alphabet.Letter("-----------------ISPREQDLLLSCGETISSVVFTSMLLDNGVKA--AALTGAQAGFLTNDQHTNAKIIEMKPER--LFSVLAN----HDAVVVAGF---QGATEKGDTTTIGRGGSDTSAAA-----------------LGAAVDAEYIDIFTDVEGVMTADP--RVVENAKPLPVVTYTEICNLAYQGAKVISPRAVEIAMQAKVPIRVRSTYS-NDKGTLVTSHHSSKVGSDVFERLITGIAH-VKDVTQFKVPAKIGQYN-----VQTEVFKAMANAGISVDFFNITPSEIVYTVAGNKTETAQR------------ILMDMGYDPMVTRNCAKVSAVGAGIMGVPGVTSKI------VSALSEKEIPILQSA"),
		[]alphabet.Letter("-----------------KRE--MDMLLSTGEQVSIALLAMSLHEKGYKA--VSLTGWQAGITTEEMHGNARIMNIDTT--RIRRCLDE----GAIVIVAGF---QGVTETGEITTLGRGGSDTTAVA-----------------LAAALKAEKCDIYTDVTGVFTTDP--RYVKTARKIKEISYDEMLELANLGAGVLHPRAVEFAKNYEVPLEVRSSME-NERGTMVK--EEVSMEQHLIVRGIAFEDQ-VTRVTVVGIEKYLQSVA--------TIFTALANRGINVDIIIQNA--------------------TNSETAS--VSFSIRTEDLPETLQVLQ-------------ALEGADVHYESGLAKVSI-VGSGMISNPGVAARV------FEVLADQGIEIKMVS"),
		[]alphabet.Letter("-----------------KRE--MDMLLATGEQVTISLLSMALQEKGYDA--VSYTGWQAGIRTEAIHGNARITDIDTS--VLADQLEK----GKIVIVAGF---QGMTEDCEITTLGRGGSDTTAVA-----------------LAAALKVDKCDIYTDVPGVFTTDP--RYVKSARKLEGISYDEMLELANLGAGVLHPRAVEFAKNYQVPLEVRSSTE-TEAGTLIE--EESSMEQNLIVRGIAFEDQ-ITRVTIYGLTSGLTTLS--------TIFTTLAKRNINVDIIIQTQ--------------------AEDKTG---ISFSVKTEDADQTVAVLEEYK---------DALEFEKIETESKLAKVSI-VGSGMVSNPGVAAEM------FAVLAQKNILIKMVS"),
		[]alphabet.Letter("-----------------ARE--MDMLLTAGERISNALVAMAIESLGAEA--QSFTGSQAGVLTTERHGNARIVDVTPG--RVREALDE----GKICIVAGF--QGVNKETRDVTTLGRGGSDTTAVA-----------------LAAALNADVCEIYSDVDGVYTADP--RIVPNAQKLEKLSFEEMLELAAVGSKILVLRSVEYARAFNVPLRVRSSYS-NDPGTLIAGSMEDIPVEEAVLTGVATDKS-EAKVTVLGISDKPGEAA--------KVFRALADAEINIDMVLQNV--------------------SSVEDGTTDITFTCPRADGRRAMEILKKLQ---------VQGNWTNVLYDDQVDKVSL-VGAGMKSHPGVTAEF------MEALRDVNVNIELIS"),
		[]alphabet.Letter("-----------------PRE--MDMLLTAGERISNALVAMAIESLGAQA--RSFTGSQAGVITTGTHGNAKIIDVTPG--RLRDALDE----GQIVLVAGF--QGVSQDSKDVTTLGRGGSDTTAVA-----------------VAAALDADVCEIYTDVDGIFTADP--RIVPNARHLDTVSFEEMLEMAACGAKVLMLRCVEYARRYNVPIHVRSSYS-DKPGTIVKGSIEDIPMEDAILTGVAHDRS-EAKVTVVGLPDVPGYAA--------KVFRAVAEADVNIDMVLQNI--------------------SKIEDGKTDITFTCARDNGPRAVEKLSALK---------SEIGFSQVLYDDHIGKVSL-IGAGMRSHPGVTATF------CEALAEAGINIDLIS"),
		[]alphabet.Letter("-----------------TSPALTDELVSHGELMSTLLFVEILRERD--V--QAQWFDVRKVMRTNDRFGRAEPDIAALAELAALQLLPR-LNEGLVITQGF---IGSENKGRTTTLGRGGSDYTAAL-----------------LAEALHASRVDIWTDVPGIYTTDP--RVVSAAKRIDEIAFAEAAEMATFGAKVLHPATLLPAVRSDIPVFVGSSKDPRAGGTLVCNKTENPPLFRALAL--RRNQT-LLTLHSLNMLHSRGFLA--------EVFGILARHNISVDLITTSEVSVALTLDTTGSTSTG----------DTLLTQSLLMELSALCRVEVEEGLALVALIG----------NDLSKACGVGKEVF"),
		[]alphabet.Letter("-----------------VSSRTVDLVMSCGEKLSCLFMTALCNDRGCKAKYVDLSHIVPSDFSASALDNSFYTFLVQALKEKLAPFVSA-KERIVPVFTGF---FGLVPTGLLNGVGRGYTDLCAAL-----------------IAVAVNADELQVWKEVDGIFTADP--RKVPEARLLDSVTPEEASELTYYGSEVIHPFTMEQVIRAKIPIRIKNVQNPLGNGTIIYPDNVAKKGESTPPHPPENLSS----SFYEKRKRGATAITTKN----DIFVINIHSNKKTLSHGFLAQIFTILDKYKLVVDLISTSEVHVSMALPIPDADS-LKSLRQAEEKLRILGSVDITKKLSIVSLVGKHMKQYIGIAG---TMFTTLAEEGINIEMIS"),
	}
)

func (s *S) TestReadFasta(c *check.C) {
	var (
		obtainNfa []string
		obtainSfa [][]alphabet.Letter
	)

	sc := seqio.NewScanner(
		fasta.NewReader(
			bytes.NewBufferString(testaln0),
			linear.NewSeq("", nil, alphabet.Protein),
		),
	)
	for sc.Scan() {
		t := sc.Seq().(*linear.Seq)
		header := t.Name()
		if desc := t.Description(); len(desc) > 0 {
			header += " " + desc
		}
		obtainNfa = append(obtainNfa, header)
		obtainSfa = append(obtainSfa, t.Slice().(alphabet.Letters))
	}
	c.Check(sc.Err(), check.Equals, nil)
	c.Check(obtainNfa, check.DeepEquals, expectNfa)
	for i := range obtainSfa {
		c.Check(len(obtainSfa[i]), check.Equals, len(expectSfa[i]))
		c.Check(obtainSfa[i], check.DeepEquals, expectSfa[i])
	}
}

// Helper
func constructQL(l [][]alphabet.Letter, q [][]alphabet.Qphred) (ql [][]alphabet.QLetter) {
	if len(l) != len(q) {
		panic("test data length mismatch")
	}
	ql = make([][]alphabet.QLetter, len(l))
	for i := range ql {
		if len(l[i]) != len(q[i]) {
			panic("test data length mismatch")
		}
		ql[i] = make([]alphabet.QLetter, len(l[i]))
		for j := range ql[i] {
			ql[i][j] = alphabet.QLetter{L: l[i][j], Q: q[i][j]}
		}
	}

	return
}

var (
	fq0 = `@FC12044_91407_8_200_406_24
GTTAGCTCCCACCTTAAGATGTTTA
+FC12044_91407_8_200_406_24
SXXTXXXXXXXXXTTSUXSSXKTMQ
@FC12044_91407_8_200_720_610
CTCTGTGGCACCCCATCCCTCACTT
+FC12044_91407_8_200_720_610
OXXXXXXXXXXXXXXXXXTSXQTXU
@FC12044_91407_8_200_345_133
GATTTTTTAACAATAAACGTACATA
+FC12044_91407_8_200_345_133
OQTOOSFORTFFFIIOFFFFFFFFF
@FC12044_91407_8_200_106_131
GTTGCCCAGGCTCGTCTTGAACTCC
+FC12044_91407_8_200_106_131
XXXXXXXXXXXXXXSXXXXISTXQS
@FC12044_91407_8_200_916_471
TGATTGAAGGTAGGGTAGCATACTG
+FC12044_91407_8_200_916_471
XXXXXXXXXXXXXXXUXXUSXXTXW
@FC12044_91407_8_200_57_85
GCTCCAATAGCGCAGAGGAAACCTG
+FC12044_91407_8_200_57_85
XFXMXSXXSXXXOSQROOSROFQIQ
@FC12044_91407_8_200_10_437
GCTGCTTGGGAGGCTGAGGCAGGAG
+FC12044_91407_8_200_10_437
USXSXXXXXXUXXXSXQXXUQXXKS
@FC12044_91407_8_200_154_436
AGACCTTTGGATACAATGAACGACT
+FC12044_91407_8_200_154_436
MKKMQTSRXMSQTOMRFOOIFFFFF
@FC12044_91407_8_200_336_64
AGGGAATTTTAGAGGAGGGCTGCCG
+FC12044_91407_8_200_336_64
STQMOSXSXSQXQXXKXXXKFXFFK
@FC12044_91407_8_200_620_233
TCTCCATGTTGGTCAGGCTGGTCTC
+FC12044_91407_8_200_620_233
XXXXXXXXXXXXXXXXXXXXXSXSW
@FC12044_91407_8_200_902_349
TGAACGTCGAGACGCAAGGCCCGCC
+FC12044_91407_8_200_902_349
XMXSSXMXXSXQSXTSQXFKSKTOF
@FC12044_91407_8_200_40_618
CTGTCCCCACGGCGGGGGGGCCTGG
+FC12044_91407_8_200_40_618
TXXXXSXXXXXXXXXXXXXRKFOXS
@FC12044_91407_8_200_83_511
GATGTACTCTTACACCCAGACTTTG
+FC12044_91407_8_200_83_511
SOXXXXXUXXXXXXQKQKKROOQSU
@FC12044_91407_8_200_76_246
TCAAGGGTGGATCTTGGCTCCCAGT
+FC12044_91407_8_200_76_246
XTXTUXXXXXRXXXTXXSUXSRFXQ
@FC12044_91407_8_200_303_427
TTGCGACAGAGTTTTGCTCTTGTCC
+FC12044_91407_8_200_303_427
XXQROXXXXIXFQXXXOIQSSXUFF
@FC12044_91407_8_200_31_299
TCTGCTCCAGCTCCAAGACGCCGCC
+FC12044_91407_8_200_31_299
XRXTSXXXRXXSXQQOXQTSQSXKQ
@FC12044_91407_8_200_553_135
TACGGAGCCGCGGGCGGGAAAGGCG
+FC12044_91407_8_200_553_135
XSQQXXXXXXXXXXSXXMFFQXTKU
@FC12044_91407_8_200_139_74
CCTCCCAGGTTCAAGCGATTATCCT
+FC12044_91407_8_200_139_74
RMXUSXTXXQXXQUXXXSQISISSO
@FC12044_91407_8_200_108_33
GTCATGGCGGCCCGCGCGGGGAGCG
+FC12044_91407_8_200_108_33
OOOSSXXSXXOMKMOFMKFOKFFFF
@FC12044_91407_8_200_980_965
ACAGTGGGTTCTTAAAGAAGAGTCG
+FC12044_91407_8_200_980_965
TOSSRXXXSSMSXMOMXIRXOXFFS
@FC12044_91407_8_200_981_857
AACGAGGGGCGCGACTTGACCTTGG
+FC12044_91407_8_200_981_857
RXMSSXXXXSXQXQXFSXQFQKMXS
@FC12044_91407_8_200_8_865
TTTCCCACCCCAGGAAGCCTTGGAC
+FC12044_91407_8_200_8_865
XXXFKOROMKOORMIMRIIKKORFF
@FC12044_91407_8_200_292_484
TCAGCCTCCGTGCCCAGCCCACTCC
+FC12044_91407_8_200_292_484
XQXOSXXXXXUXXXXIXXXXQTOXF
@FC12044_91407_8_200_675_16
CTCGGGAGGCTGAGGCAGGGGGGTT
+FC12044_91407_8_200_675_16
OXTXXXSXXQXXOXXKMXXMXOKQF
@FC12044_91407_8_200_285_136
CCAAATCTTGAATTGTAGCTCCCCT
+FC12044_91407_8_200_285_136
OSXOQXXXXXSXXUXXTXXXXTRMS
`

	expectNfq = []string{
		"FC12044_91407_8_200_406_24",
		"FC12044_91407_8_200_720_610",
		"FC12044_91407_8_200_345_133",
		"FC12044_91407_8_200_106_131",
		"FC12044_91407_8_200_916_471",
		"FC12044_91407_8_200_57_85",
		"FC12044_91407_8_200_10_437",
		"FC12044_91407_8_200_154_436",
		"FC12044_91407_8_200_336_64",
		"FC12044_91407_8_200_620_233",
		"FC12044_91407_8_200_902_349",
		"FC12044_91407_8_200_40_618",
		"FC12044_91407_8_200_83_511",
		"FC12044_91407_8_200_76_246",
		"FC12044_91407_8_200_303_427",
		"FC12044_91407_8_200_31_299",
		"FC12044_91407_8_200_553_135",
		"FC12044_91407_8_200_139_74",
		"FC12044_91407_8_200_108_33",
		"FC12044_91407_8_200_980_965",
		"FC12044_91407_8_200_981_857",
		"FC12044_91407_8_200_8_865",
		"FC12044_91407_8_200_292_484",
		"FC12044_91407_8_200_675_16",
		"FC12044_91407_8_200_285_136",
	}

	expectSfq = [][]alphabet.Letter{
		[]alphabet.Letter("GTTAGCTCCCACCTTAAGATGTTTA"),
		[]alphabet.Letter("CTCTGTGGCACCCCATCCCTCACTT"),
		[]alphabet.Letter("GATTTTTTAACAATAAACGTACATA"),
		[]alphabet.Letter("GTTGCCCAGGCTCGTCTTGAACTCC"),
		[]alphabet.Letter("TGATTGAAGGTAGGGTAGCATACTG"),
		[]alphabet.Letter("GCTCCAATAGCGCAGAGGAAACCTG"),
		[]alphabet.Letter("GCTGCTTGGGAGGCTGAGGCAGGAG"),
		[]alphabet.Letter("AGACCTTTGGATACAATGAACGACT"),
		[]alphabet.Letter("AGGGAATTTTAGAGGAGGGCTGCCG"),
		[]alphabet.Letter("TCTCCATGTTGGTCAGGCTGGTCTC"),
		[]alphabet.Letter("TGAACGTCGAGACGCAAGGCCCGCC"),
		[]alphabet.Letter("CTGTCCCCACGGCGGGGGGGCCTGG"),
		[]alphabet.Letter("GATGTACTCTTACACCCAGACTTTG"),
		[]alphabet.Letter("TCAAGGGTGGATCTTGGCTCCCAGT"),
		[]alphabet.Letter("TTGCGACAGAGTTTTGCTCTTGTCC"),
		[]alphabet.Letter("TCTGCTCCAGCTCCAAGACGCCGCC"),
		[]alphabet.Letter("TACGGAGCCGCGGGCGGGAAAGGCG"),
		[]alphabet.Letter("CCTCCCAGGTTCAAGCGATTATCCT"),
		[]alphabet.Letter("GTCATGGCGGCCCGCGCGGGGAGCG"),
		[]alphabet.Letter("ACAGTGGGTTCTTAAAGAAGAGTCG"),
		[]alphabet.Letter("AACGAGGGGCGCGACTTGACCTTGG"),
		[]alphabet.Letter("TTTCCCACCCCAGGAAGCCTTGGAC"),
		[]alphabet.Letter("TCAGCCTCCGTGCCCAGCCCACTCC"),
		[]alphabet.Letter("CTCGGGAGGCTGAGGCAGGGGGGTT"),
		[]alphabet.Letter("CCAAATCTTGAATTGTAGCTCCCCT"),
	}

	expectQ = [][]alphabet.Qphred{
		{50, 55, 55, 51, 55, 55, 55, 55, 55, 55, 55, 55, 55, 51, 51, 50, 52, 55, 50, 50, 55, 42, 51, 44, 48},
		{46, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 51, 50, 55, 48, 51, 55, 52},
		{46, 48, 51, 46, 46, 50, 37, 46, 49, 51, 37, 37, 37, 40, 40, 46, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 50, 55, 55, 55, 55, 40, 50, 51, 55, 48, 50},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 52, 55, 55, 52, 50, 55, 55, 51, 55, 54},
		{55, 37, 55, 44, 55, 50, 55, 55, 50, 55, 55, 55, 46, 50, 48, 49, 46, 46, 50, 49, 46, 37, 48, 40, 48},
		{52, 50, 55, 50, 55, 55, 55, 55, 55, 55, 52, 55, 55, 55, 50, 55, 48, 55, 55, 52, 48, 55, 55, 42, 50},
		{44, 42, 42, 44, 48, 51, 50, 49, 55, 44, 50, 48, 51, 46, 44, 49, 37, 46, 46, 40, 37, 37, 37, 37, 37},
		{50, 51, 48, 44, 46, 50, 55, 50, 55, 50, 48, 55, 48, 55, 55, 42, 55, 55, 55, 42, 37, 55, 37, 37, 42},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 50, 55, 50, 54},
		{55, 44, 55, 50, 50, 55, 44, 55, 55, 50, 55, 48, 50, 55, 51, 50, 48, 55, 37, 42, 50, 42, 51, 46, 37},
		{51, 55, 55, 55, 55, 50, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 49, 42, 37, 46, 55, 50},
		{50, 46, 55, 55, 55, 55, 55, 52, 55, 55, 55, 55, 55, 55, 48, 42, 48, 42, 42, 49, 46, 46, 48, 50, 52},
		{55, 51, 55, 51, 52, 55, 55, 55, 55, 55, 49, 55, 55, 55, 51, 55, 55, 50, 52, 55, 50, 49, 37, 55, 48},
		{55, 55, 48, 49, 46, 55, 55, 55, 55, 40, 55, 37, 48, 55, 55, 55, 46, 40, 48, 50, 50, 55, 52, 37, 37},
		{55, 49, 55, 51, 50, 55, 55, 55, 49, 55, 55, 50, 55, 48, 48, 46, 55, 48, 51, 50, 48, 50, 55, 42, 48},
		{55, 50, 48, 48, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 50, 55, 55, 44, 37, 37, 48, 55, 51, 42, 52},
		{49, 44, 55, 52, 50, 55, 51, 55, 55, 48, 55, 55, 48, 52, 55, 55, 55, 50, 48, 40, 50, 40, 50, 50, 46},
		{46, 46, 46, 50, 50, 55, 55, 50, 55, 55, 46, 44, 42, 44, 46, 37, 44, 42, 37, 46, 42, 37, 37, 37, 37},
		{51, 46, 50, 50, 49, 55, 55, 55, 50, 50, 44, 50, 55, 44, 46, 44, 55, 40, 49, 55, 46, 55, 37, 37, 50},
		{49, 55, 44, 50, 50, 55, 55, 55, 55, 50, 55, 48, 55, 48, 55, 37, 50, 55, 48, 37, 48, 42, 44, 55, 50},
		{55, 55, 55, 37, 42, 46, 49, 46, 44, 42, 46, 46, 49, 44, 40, 44, 49, 40, 40, 42, 42, 46, 49, 37, 37},
		{55, 48, 55, 46, 50, 55, 55, 55, 55, 55, 52, 55, 55, 55, 55, 40, 55, 55, 55, 55, 48, 51, 46, 55, 37},
		{46, 55, 51, 55, 55, 55, 50, 55, 55, 48, 55, 55, 46, 55, 55, 42, 44, 55, 55, 44, 55, 46, 42, 48, 37},
		{46, 50, 55, 46, 48, 55, 55, 55, 55, 55, 50, 55, 55, 52, 55, 55, 51, 55, 55, 55, 55, 51, 49, 44, 50},
	}

	expectQL = constructQL(expectSfq, expectQ)
)

func (s *S) TestReadFastq(c *check.C) {
	var (
		obtainNfq []string
		obtainQL  [][]alphabet.QLetter
	)

	sc := seqio.NewScanner(
		fastq.NewReader(
			bytes.NewBufferString(fq0),
			linear.NewQSeq("", nil, alphabet.DNA, alphabet.Sanger),
		),
	)
	for sc.Scan() {
		t := sc.Seq().(*linear.QSeq)
		header := t.Name()
		if desc := t.Description(); len(desc) > 0 {
			header += " " + desc
		}
		obtainNfq = append(obtainNfq, header)
		obtainQL = append(obtainQL, (t.Slice().(alphabet.QLetters)))
	}
	c.Check(sc.Err(), check.Equals, nil)
	c.Check(obtainNfq, check.DeepEquals, expectNfq)
	c.Check(obtainQL, check.DeepEquals, expectQL)
}
