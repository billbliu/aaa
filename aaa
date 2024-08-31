package init_db

import (
	"encoding/json"
	"tarot-server/global"
	"tarot-server/models/tables"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

type TCountryPhoneCodeInit struct{}

func (t *TCountryPhoneCodeInit) Init() error {
	data := `{
		"A": [
		    {
			"chinese_name": "阿尔巴尼亚",
			"chinesePhoneticize": "aerbaniya",
			"english_name": "Albania",
			"disablePrefill": false,
			"displayName": "阿尔巴尼亚",
			"enable": true,
			"formal_name": "Republic of Albania",
			"country_code": "AL",
			
			"national_flag": "🇦🇱",
			"start_char": "A",
			"telephone_code": "355"
		    },
		    {
			"chinese_name": "阿尔及利亚",
			"chinesePhoneticize": "aerjiliya",
			"english_name": "Algeria",
			"disablePrefill": false,
			"displayName": "阿尔及利亚",
			"enable": true,
			"formal_name": "People's Democratic Republic of Algeria",
			"country_code": "DZ",
			
			"national_flag": "🇩🇿",
			"start_char": "A",
			"telephone_code": "213"
		    },
		    {
			"chinese_name": "阿富汗",
			"chinesePhoneticize": "afuhan",
			"english_name": "Afghanistan",
			"disablePrefill": false,
			"displayName": "阿富汗",
			"enable": true,
			"formal_name": "Islamic State of Afghanistan",
			"country_code": "AF",
			
			"national_flag": "🇦🇫",
			"start_char": "A",
			"telephone_code": "93"
		    },
		    {
			"chinese_name": "阿根廷",
			"chinesePhoneticize": "agenting",
			"english_name": "Argentina",
			"disablePrefill": false,
			"displayName": "阿根廷",
			"enable": true,
			"formal_name": "Argentine Republic",
			"country_code": "AR",
			
			"national_flag": "🇦🇷",
			"start_char": "A",
			"telephone_code": "54"
		    },
		    {
			"chinese_name": "爱尔兰",
			"chinesePhoneticize": "aierlan",
			"english_name": "Ireland",
			"disablePrefill": false,
			"displayName": "爱尔兰",
			"enable": true,
			"formal_name": "",
			"country_code": "IE",
			
			"national_flag": "🇮🇪",
			"start_char": "I",
			"telephone_code": "353"
		    },
		    {
			"chinese_name": "埃及",
			"chinesePhoneticize": "aiji",
			"english_name": "Egypt",
			"disablePrefill": false,
			"displayName": "埃及",
			"enable": true,
			"formal_name": "Arab Republic of Egypt",
			"country_code": "EG",
			
			"national_flag": "🇪🇬",
			"start_char": "E",
			"telephone_code": "20"
		    },
		    {
			"chinese_name": "埃塞俄比亚",
			"chinesePhoneticize": "aisaiebiya",
			"english_name": "Ethiopia",
			"disablePrefill": false,
			"displayName": "埃塞俄比亚",
			"enable": true,
			"formal_name": "Federal Democratic Republic of Ethiopia",
			"country_code": "ET",
			
			"national_flag": "🇪🇹",
			"start_char": "E",
			"telephone_code": "251"
		    },
		    {
			"chinese_name": "爱沙尼亚",
			"chinesePhoneticize": "aishaniya",
			"english_name": "Estonia",
			"disablePrefill": false,
			"displayName": "爱沙尼亚",
			"enable": true,
			"formal_name": "Republic of Estonia",
			"country_code": "EE",
			
			"national_flag": "🇪🇪",
			"start_char": "E",
			"telephone_code": "372"
		    },
		    {
			"chinese_name": "阿拉伯联合酋长国",
			"chinesePhoneticize": "alabolianheqiuzhangguo",
			"english_name": "United Arab Emirates",
			"disablePrefill": false,
			"displayName": "阿拉伯联合酋长国",
			"enable": true,
			"formal_name": "United Arab Emirates",
			"country_code": "AE",
			
			"national_flag": "🇦🇪",
			"start_char": "U",
			"telephone_code": "971"
		    },
		    {
			"chinese_name": "阿鲁巴",
			"chinesePhoneticize": "aluba",
			"english_name": "Aruba",
			"disablePrefill": false,
			"displayName": "阿鲁巴",
			"enable": true,
			"formal_name": "Aruba",
			"country_code": "AW",
			
			"national_flag": "🇦🇼",
			"start_char": "A",
			"telephone_code": "297"
		    },
		    {
			"chinese_name": "阿曼",
			"chinesePhoneticize": "aman",
			"english_name": "Oman",
			"disablePrefill": false,
			"displayName": "阿曼",
			"enable": true,
			"formal_name": "Sultanate of Oman",
			"country_code": "OM",
			
			"national_flag": "🇴🇲",
			"start_char": "O",
			"telephone_code": "968"
		    },
		    {
			"chinese_name": "安道尔",
			"chinesePhoneticize": "andaoer",
			"english_name": "Andorra",
			"disablePrefill": false,
			"displayName": "安道尔",
			"enable": true,
			"formal_name": "Principality of Andorra",
			"country_code": "AD",
			
			"national_flag": "🇦🇩",
			"start_char": "A",
			"telephone_code": "376"
		    },
		    {
			"chinese_name": "安哥拉",
			"chinesePhoneticize": "angela",
			"english_name": "Angola",
			"disablePrefill": false,
			"displayName": "安哥拉",
			"enable": true,
			"formal_name": "Republic of Angola",
			"country_code": "AO",
			
			"national_flag": "🇦🇴",
			"start_char": "A",
			"telephone_code": "244"
		    },
		    {
			"chinese_name": "安圭拉",
			"chinesePhoneticize": "anguila",
			"english_name": "Anguilla",
			"disablePrefill": false,
			"displayName": "安圭拉",
			"enable": true,
			"formal_name": "",
			"country_code": "AI",
			
			"national_flag": "🇦🇮",
			"start_char": "A",
			"telephone_code": "1264"
		    },
		    {
			"chinese_name": "安提瓜和巴布达",
			"chinesePhoneticize": "antiguahebabuda",
			"english_name": "Antigua and Barbuda",
			"disablePrefill": false,
			"displayName": "安提瓜和巴布达",
			"enable": true,
			"formal_name": "",
			"country_code": "AG",
			
			"national_flag": "🇦🇬",
			"start_char": "A",
			"telephone_code": "1268"
		    },
		    {
			"chinese_name": "澳大利亚",
			"chinesePhoneticize": "aodaliya",
			"english_name": "Australia",
			"disablePrefill": false,
			"displayName": "澳大利亚",
			"enable": true,
			"formal_name": "Commonwealth of Australia",
			"country_code": "AU",
			
			"national_flag": "🇦🇺",
			"start_char": "A",
			"telephone_code": "61"
		    },
		    {
			"chinese_name": "奥地利",
			"chinesePhoneticize": "aodili",
			"english_name": "Austria",
			"disablePrefill": false,
			"displayName": "奥地利",
			"enable": true,
			"formal_name": "Republic of Austria",
			"country_code": "AT",
			
			"national_flag": "🇦🇹",
			"start_char": "A",
			"telephone_code": "43"
		    },
		    {
			"chinese_name": "奥兰群岛",
			"chinesePhoneticize": "aolanqundao",
			"english_name": "Åland Islands",
			"disablePrefill": false,
			"displayName": "奥兰群岛",
			"enable": true,
			"formal_name": "Åland Islands",
			"country_code": "AX",
			
			"national_flag": "🇦🇽",
			"start_char": "Å",
			"telephone_code": "340"
		    },
		    {
			"chinese_name": "澳门",
			"chinesePhoneticize": "aomen",
			"english_name": "Macau",
			"disablePrefill": false,
			"displayName": "中国澳门",
			"enable": true,
			"formal_name": "Macau&nbsp;Special&nbsp;Administrative&nbsp;Region",
			"country_code": "MO",
			
			"national_flag": "🇲🇴",
			"start_char": "M",
			"telephone_code": "853"
		    },
		    {
			"chinese_name": "阿塞拜疆",
			"chinesePhoneticize": "asaibaijiang",
			"english_name": "Azerbaijan",
			"disablePrefill": false,
			"displayName": "阿塞拜疆",
			"enable": true,
			"formal_name": "Republic of Azerbaijan",
			"country_code": "AZ",
			
			"national_flag": "🇦🇿",
			"start_char": "A",
			"telephone_code": "994"
		    }
		],
		"B": [
		    {
			"chinese_name": "巴巴多斯",
			"chinesePhoneticize": "babaduosi",
			"english_name": "Barbados",
			"disablePrefill": false,
			"displayName": "巴巴多斯",
			"enable": true,
			"formal_name": "",
			"country_code": "BB",
			
			"national_flag": "🇧🇧",
			"start_char": "B",
			"telephone_code": "1246"
		    },
		    {
			"chinese_name": "巴布亚新几内亚",
			"chinesePhoneticize": "babuyaxinjineiya",
			"english_name": "Papua New Guinea",
			"disablePrefill": false,
			"displayName": "巴布亚新几内亚",
			"enable": true,
			"formal_name": "Independent State of Papua New Guinea",
			"country_code": "PG",
			
			"national_flag": "🇵🇬",
			"start_char": "P",
			"telephone_code": "675"
		    },
		    {
			"chinese_name": "巴哈马",
			"chinesePhoneticize": "bahama",
			"english_name": "Bahamas",
			"disablePrefill": false,
			"displayName": "巴哈马",
			"enable": true,
			"formal_name": "Commonwealth of The Bahamas",
			"country_code": "BS",
			
			"national_flag": "🇧🇸",
			"start_char": "B",
			"telephone_code": "1242"
		    },
		    {
			"chinese_name": "白俄罗斯",
			"chinesePhoneticize": "baieluosi",
			"english_name": "Belarus",
			"disablePrefill": false,
			"displayName": "白俄罗斯",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Belarus",
			"country_code": "BY",
			
			"national_flag": "🇧🇾",
			"start_char": "B",
			"telephone_code": "375"
		    },
		    {
			"chinese_name": "百慕大",
			"chinesePhoneticize": "baimuda",
			"english_name": "Bermuda",
			"disablePrefill": false,
			"displayName": "百慕大",
			"enable": true,
			"formal_name": "",
			"country_code": "BM",
			
			"national_flag": "🇧🇲",
			"start_char": "B",
			"telephone_code": "1441"
		    },
		    {
			"chinese_name": "巴基斯坦",
			"chinesePhoneticize": "bajisitan",
			"english_name": "Pakistan",
			"disablePrefill": false,
			"displayName": "巴基斯坦",
			"enable": true,
			"formal_name": "Islamic Republic of Pakistan",
			"country_code": "PK",
			
			"national_flag": "🇵🇰",
			"start_char": "P",
			"telephone_code": "92"
		    },
		    {
			"chinese_name": "巴拉圭",
			"chinesePhoneticize": "balagui",
			"english_name": "Paraguay",
			"disablePrefill": false,
			"displayName": "巴拉圭",
			"enable": true,
			"formal_name": "Republic of Paraguay",
			"country_code": "PY",
			
			"national_flag": "🇵🇾",
			"start_char": "P",
			"telephone_code": "595"
		    },
		    {
			"chinese_name": "巴勒斯坦",
			"chinesePhoneticize": "balesitan",
			"english_name": "Palestine",
			"disablePrefill": false,
			"displayName": "巴勒斯坦",
			"enable": true,
			"formal_name": "State of Palestine",
			"country_code": "PS",
			
			"national_flag": "🇵🇸",
			"start_char": "P",
			"telephone_code": "970"
		    },
		    {
			"chinese_name": "巴林",
			"chinesePhoneticize": "balin",
			"english_name": "Bahrain",
			"disablePrefill": false,
			"displayName": "巴林",
			"enable": true,
			"formal_name": "Kingdom of Bahrain",
			"country_code": "BH",
			
			"national_flag": "🇧🇭",
			"start_char": "B",
			"telephone_code": "973"
		    },
		    {
			"chinese_name": "巴拿马",
			"chinesePhoneticize": "banama",
			"english_name": "Panama",
			"disablePrefill": false,
			"displayName": "巴拿马",
			"enable": true,
			"formal_name": "Republic of Panama",
			"country_code": "PA",
			
			"national_flag": "🇵🇦",
			"start_char": "P",
			"telephone_code": "507"
		    },
		    {
			"chinese_name": "保加利亚",
			"chinesePhoneticize": "baojialiya",
			"english_name": "Bulgaria",
			"disablePrefill": false,
			"displayName": "保加利亚",
			"enable": true,
			"formal_name": "Republic of Bulgaria",
			"country_code": "BG",
			
			"national_flag": "🇧🇬",
			"start_char": "B",
			"telephone_code": "359"
		    },
		    {
			"chinese_name": "巴西",
			"chinesePhoneticize": "baxi",
			"english_name": "Brazil",
			"disablePrefill": false,
			"displayName": "巴西",
			"enable": true,
			"formal_name": "Federative Republic of Brazil",
			"country_code": "BR",
			
			"national_flag": "🇧🇷",
			"start_char": "B",
			"telephone_code": "55"
		    },
		    {
			"chinese_name": "北马里亚纳群岛",
			"chinesePhoneticize": "beimaliyanaqundao",
			"english_name": "Northern Mariana Islands",
			"disablePrefill": false,
			"displayName": "北马里亚纳群岛",
			"enable": true,
			"formal_name": "Northern Mariana Islands",
			"country_code": "MP",
			
			"national_flag": "🇲🇵",
			"start_char": "N",
			"telephone_code": "1-670"
		    },
		    {
			"chinese_name": "贝宁",
			"chinesePhoneticize": "beining",
			"english_name": "Benin",
			"disablePrefill": false,
			"displayName": "贝宁",
			"enable": true,
			"formal_name": "Republic of Benin",
			"country_code": "BJ",
			
			"national_flag": "🇧🇯",
			"start_char": "B",
			"telephone_code": "229"
		    },
		    {
			"chinese_name": "比利时",
			"chinesePhoneticize": "bilishi",
			"english_name": "Belgium",
			"disablePrefill": false,
			"displayName": "比利时",
			"enable": true,
			"formal_name": "Kingdom of Belgium",
			"country_code": "BE",
			
			"national_flag": "🇧🇪",
			"start_char": "B",
			"telephone_code": "32"
		    },
		    {
			"chinese_name": "冰岛",
			"chinesePhoneticize": "bingdao",
			"english_name": "Iceland",
			"disablePrefill": false,
			"displayName": "冰岛",
			"enable": true,
			"formal_name": "Republic of Iceland",
			"country_code": "IS",
			
			"national_flag": "🇮🇸",
			"start_char": "I",
			"telephone_code": "354"
		    },
		    {
			"chinese_name": "博茨瓦纳",
			"chinesePhoneticize": "bociwana",
			"english_name": "Botswana",
			"disablePrefill": false,
			"displayName": "博茨瓦纳",
			"enable": true,
			"formal_name": "Republic of Botswana",
			"country_code": "BW",
			
			"national_flag": "🇧🇼",
			"start_char": "B",
			"telephone_code": "267"
		    },
		    {
			"chinese_name": "波多黎各",
			"chinesePhoneticize": "boduolige",
			"english_name": "Puerto Rico",
			"disablePrefill": false,
			"displayName": "波多黎各",
			"enable": true,
			"formal_name": "Puerto Rico",
			"country_code": "PR",
			
			"national_flag": "🇵🇷",
			"start_char": "P",
			"telephone_code": "1"
		    },
		    {
			"chinese_name": "波兰",
			"chinesePhoneticize": "bolan",
			"english_name": "Poland",
			"disablePrefill": false,
			"displayName": "波兰",
			"enable": true,
			"formal_name": "Republic of Poland",
			"country_code": "PL",
			
			"national_flag": "🇵🇱",
			"start_char": "P",
			"telephone_code": "48"
		    },
		    {
			"chinese_name": "玻利维亚",
			"chinesePhoneticize": "boliweiya",
			"english_name": "Bolivia",
			"disablePrefill": false,
			"displayName": "玻利维亚",
			"enable": true,
			"formal_name": "Republic of Bolivia",
			"country_code": "BO",
			
			"national_flag": "🇧🇴",
			"start_char": "B",
			"telephone_code": "591"
		    },
		    {
			"chinese_name": "伯利兹",
			"chinesePhoneticize": "bolizi",
			"english_name": "Belize",
			"disablePrefill": false,
			"displayName": "伯利兹",
			"enable": true,
			"formal_name": "",
			"country_code": "BZ",
			
			"national_flag": "🇧🇿",
			"start_char": "B",
			"telephone_code": "501"
		    },
		    {
			"chinese_name": "波斯尼亚和黑塞哥维那 (波黑)",
			"chinesePhoneticize": "bosiniyaheheisaigeweina (bohei)",
			"english_name": "Bosnia and Herzegovina",
			"disablePrefill": false,
			"displayName": "波斯尼亚和黑塞哥维那 (波黑)",
			"enable": true,
			"formal_name": "Bosnia&nbsp;and&nbsp;Herzegovina",
			"country_code": "BA",
			
			"national_flag": "🇧🇦",
			"start_char": "B",
			"telephone_code": "387"
		    },
		    {
			"chinese_name": "不丹",
			"chinesePhoneticize": "budan",
			"english_name": "Bhutan",
			"disablePrefill": false,
			"displayName": "不丹",
			"enable": true,
			"formal_name": "Kingdom of Bhutan",
			"country_code": "BT",
			
			"national_flag": "🇧🇹",
			"start_char": "B",
			"telephone_code": "975"
		    },
		    {
			"chinese_name": "布基纳法索",
			"chinesePhoneticize": "bujinafasuo",
			"english_name": "Burkina Faso",
			"disablePrefill": false,
			"displayName": "布基纳法索",
			"enable": true,
			"formal_name": "",
			"country_code": "BF",
			
			"national_flag": "🇧🇫",
			"start_char": "B",
			"telephone_code": "226"
		    },
		    {
			"chinese_name": "布隆迪",
			"chinesePhoneticize": "bulongdi",
			"english_name": "Burundi",
			"disablePrefill": false,
			"displayName": "布隆迪",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Burundi",
			"country_code": "BI",
			
			"national_flag": "🇧🇮",
			"start_char": "B",
			"telephone_code": "257"
		    }
		],
		"C": [
		    {
			"chinese_name": "朝鲜",
			"chinesePhoneticize": "chaoxian",
			"english_name": "North Korea",
			"disablePrefill": false,
			"displayName": "朝鲜",
			"enable": false,
			"formal_name": "Democratic People's Republic of Korea",
			"country_code": "KP",
			
			"national_flag": "🇰🇵",
			"start_char": "N",
			"telephone_code": "850"
		    },
		    {
			"chinese_name": "赤道几内亚",
			"chinesePhoneticize": "chidaojineiya",
			"english_name": "Equatorial Guinea",
			"disablePrefill": false,
			"displayName": "赤道几内亚",
			"enable": true,
			"formal_name": "Republic of Equatorial Guinea",
			"country_code": "GQ",
			
			"national_flag": "🇬🇶",
			"start_char": "E",
			"telephone_code": "240"
		    }
		],
		"D": [
		    {
			"chinese_name": "丹麦",
			"chinesePhoneticize": "danmai",
			"english_name": "Denmark",
			"disablePrefill": false,
			"displayName": "丹麦",
			"enable": true,
			"formal_name": "Kingdom of Denmark",
			"country_code": "DK",
			
			"national_flag": "🇩🇰",
			"start_char": "D",
			"telephone_code": "45"
		    },
		    {
			"chinese_name": "德国",
			"chinesePhoneticize": "deguo",
			"english_name": "Germany",
			"disablePrefill": false,
			"displayName": "德国",
			"enable": true,
			"formal_name": "Federal Republic of Germany",
			"country_code": "DE",
			
			"national_flag": "🇩🇪",
			"start_char": "G",
			"telephone_code": "49"
		    },
		    {
			"chinese_name": "东帝汶",
			"chinesePhoneticize": "dongdiwen",
			"english_name": "Timor-Leste (East Timor)",
			"disablePrefill": false,
			"displayName": "东帝汶",
			"enable": true,
			"formal_name": "Democratic Republic of Timor-Leste",
			"country_code": "TL",
			
			"national_flag": "🇹🇱",
			"start_char": "T",
			"telephone_code": "670"
		    },
		    {
			"chinese_name": "多哥",
			"chinesePhoneticize": "duoge",
			"english_name": "Togo",
			"disablePrefill": false,
			"displayName": "多哥",
			"enable": true,
			"formal_name": "Togolese Republic",
			"country_code": "TG",
			
			"national_flag": "🇹🇬",
			"start_char": "T",
			"telephone_code": "228"
		    },
		    {
			"chinese_name": "多明尼加共和国",
			"chinesePhoneticize": "duomingnijiagongheguo",
			"english_name": "Dominican Republic",
			"disablePrefill": false,
			"displayName": "多明尼加共和国",
			"enable": true,
			"formal_name": "",
			"country_code": "DO",
			
			"national_flag": "🇩🇴",
			"start_char": "D",
			"telephone_code": "1829"
		    },
		    {
			"chinese_name": "多明尼加共和国",
			"chinesePhoneticize": "duomingnijiagongheguo",
			"english_name": "Dominican Republic",
			"disablePrefill": false,
			"displayName": "多明尼加共和国",
			"enable": true,
			"formal_name": "",
			"country_code": "DO",
			
			"national_flag": "🇩🇴",
			"start_char": "D",
			"telephone_code": "1809"
		    },
		    {
			"chinese_name": "多明尼加共和国",
			"chinesePhoneticize": "duomingnijiagongheguo",
			"english_name": "Dominican Republic",
			"disablePrefill": false,
			"displayName": "多明尼加共和国",
			"enable": true,
			"formal_name": "",
			"country_code": "DO",
			
			"national_flag": "🇩🇴",
			"start_char": "D",
			"telephone_code": "1849"
		    },
		    {
			"chinese_name": "多米尼克",
			"chinesePhoneticize": "duominike",
			"english_name": "Dominica",
			"disablePrefill": false,
			"displayName": "多米尼克",
			"enable": true,
			"formal_name": "Commonwealth of Dominica",
			"country_code": "DM",
			
			"national_flag": "🇩🇲",
			"start_char": "D",
			"telephone_code": "1767"
		    }
		],
		"E": [
		    {
			"chinese_name": "厄瓜多尔",
			"chinesePhoneticize": "eguaduoer",
			"english_name": "Ecuador",
			"disablePrefill": false,
			"displayName": "厄瓜多尔",
			"enable": true,
			"formal_name": "Republic of Ecuador",
			"country_code": "EC",
			
			"national_flag": "🇪🇨",
			"start_char": "E",
			"telephone_code": "593"
		    },
		    {
			"chinese_name": "厄立特里亚",
			"chinesePhoneticize": "eliteliya",
			"english_name": "Eritrea",
			"disablePrefill": false,
			"displayName": "厄立特里亚",
			"enable": true,
			"formal_name": "State of Eritrea",
			"country_code": "ER",
			
			"national_flag": "🇪🇷",
			"start_char": "E",
			"telephone_code": "291"
		    },
		    {
			"chinese_name": "俄国",
			"chinesePhoneticize": "eluosi",
			"english_name": "Russia",
			"disablePrefill": false,
			"displayName": "俄罗斯",
			"enable": true,
			"formal_name": "Russian Federation",
			"country_code": "RU",
			
			"national_flag": "🇷🇺",
			"start_char": "R",
			"telephone_code": "7"
		    }
		],
		"F": [
		    {
			"chinese_name": "法国",
			"chinesePhoneticize": "faguo",
			"english_name": "France",
			"disablePrefill": false,
			"displayName": "法国",
			"enable": true,
			"formal_name": "French Republic",
			"country_code": "FR",
			
			"national_flag": "🇫🇷",
			"start_char": "F",
			"telephone_code": "33"
		    },
		    {
			"chinese_name": "法罗群岛",
			"chinesePhoneticize": "faluoqundao",
			"english_name": "Faroe Islands",
			"disablePrefill": false,
			"displayName": "法罗群岛",
			"enable": true,
			"formal_name": "Faroe Islands (Autonomous Territory of Denmark)",
			"country_code": "FO",
			
			"national_flag": "🇫🇴",
			"start_char": "F",
			"telephone_code": "298"
		    },
		    {
			"chinese_name": "梵蒂冈城",
			"chinesePhoneticize": "fandigangcheng",
			"english_name": "Vatican City",
			"disablePrefill": false,
			"displayName": "梵蒂冈城",
			"enable": true,
			"formal_name": "State of the Vatican City",
			"country_code": "VA",
			
			"national_flag": "🇻🇦",
			"start_char": "V",
			"telephone_code": "379"
		    },
		    {
			"chinese_name": "法属圣马丁",
			"chinesePhoneticize": "fanshushengmading",
			"english_name": "Collectivity of Saint Martin",
			"disablePrefill": false,
			"displayName": "法属圣马丁",
			"enable": true,
			"formal_name": "Collectivity of Saint Martin",
			"country_code": "MF",
			
			"national_flag": "🇲🇫",
			"start_char": "C",
			"telephone_code": "590"
		    },
		    {
			"chinese_name": "法属波利尼西亚",
			"chinesePhoneticize": "fashubolinixiya",
			"english_name": "French Polynesia",
			"disablePrefill": false,
			"displayName": "法属波利尼西亚",
			"enable": true,
			"formal_name": "french polynesia",
			"country_code": "PF",
			
			"national_flag": "🇵🇫",
			"start_char": "F",
			"telephone_code": "689"
		    },
		    {
			"chinese_name": "斐济",
			"chinesePhoneticize": "feiji",
			"english_name": "Fiji",
			"disablePrefill": false,
			"displayName": "斐济",
			"enable": true,
			"formal_name": "Republic of the Fiji Islands",
			"country_code": "FJ",
			
			"national_flag": "🇫🇯",
			"start_char": "F",
			"telephone_code": "679"
		    },
		    {
			"chinese_name": "菲律宾",
			"chinesePhoneticize": "feilübin",
			"english_name": "Philippines",
			"disablePrefill": false,
			"displayName": "菲律宾",
			"enable": true,
			"formal_name": "Republic of the Philippines",
			"country_code": "PH",
			
			"national_flag": "🇵🇭",
			"start_char": "P",
			"telephone_code": "63"
		    },
		    {
			"chinese_name": "芬兰",
			"chinesePhoneticize": "fenlan",
			"english_name": "Finland",
			"disablePrefill": false,
			"displayName": "芬兰",
			"enable": true,
			"formal_name": "Republic of Finland",
			"country_code": "FI",
			
			"national_flag": "🇫🇮",
			"start_char": "F",
			"telephone_code": "358"
		    },
		    {
			"chinese_name": "佛得角",
			"chinesePhoneticize": "fodejiao",
			"english_name": "Cape Verde",
			"disablePrefill": false,
			"displayName": "佛得角",
			"enable": true,
			"formal_name": "Republic of Cape Verde",
			"country_code": "CV",
			
			"national_flag": "🇨🇻",
			"start_char": "C",
			"telephone_code": "238"
		    },
		    {
			"chinese_name": "福克兰群岛 (马尔维纳斯群岛)",
			"chinesePhoneticize": "fukelanqundao",
			"english_name": "Falkland Islands (Islas Malvinas)",
			"disablePrefill": false,
			"displayName": "福克兰群岛 (马尔维纳斯群岛)",
			"enable": true,
			"formal_name": "Falkland Islands (Islas Malvinas)",
			"country_code": "FK",
			
			"national_flag": "🇫🇰",
			"start_char": "F",
			"telephone_code": "500"
		    }
		],
		"G": [
		    {
			"chinese_name": "冈比亚",
			"chinesePhoneticize": "gangbiya",
			"english_name": "Gambia",
			"disablePrefill": false,
			"displayName": "冈比亚",
			"enable": true,
			"formal_name": "Republic of The Gambia",
			"country_code": "GM",
			
			"national_flag": "🇬🇲",
			"start_char": "G",
			"telephone_code": "220"
		    },
		    {
			"chinese_name": "刚果共和国",
			"chinesePhoneticize": "gangguogongheguo",
			"english_name": "Congo (Republic)",
			"disablePrefill": false,
			"displayName": "刚果共和国",
			"enable": true,
			"formal_name": "Republic of the Congo",
			"country_code": "CG",
			
			"national_flag": "🇨🇬",
			"start_char": "C",
			"telephone_code": "242"
		    },
		    {
			"chinese_name": "刚果民主共和国",
			"chinesePhoneticize": "gangguominzhugongheguo",
			"english_name": "Congo (Democratic Republic)",
			"disablePrefill": false,
			"displayName": "刚果民主共和国",
			"enable": true,
			"formal_name": "Democratic&nbsp;Republic&nbsp;of&nbsp;the&nbsp;Congo",
			"country_code": "CD",
			
			"national_flag": "🇨🇩",
			"start_char": "C",
			"telephone_code": "243"
		    },
		    {
			"chinese_name": "格陵兰岛",
			"chinesePhoneticize": "gelinglandao",
			"english_name": "Greenland",
			"disablePrefill": false,
			"displayName": "格陵兰岛",
			"enable": true,
			"formal_name": "Greenland (Autonomous Territory of Denmark)",
			"country_code": "GL",
			
			"national_flag": "🇬🇱",
			"start_char": "G",
			"telephone_code": "299"
		    },
		    {
			"chinese_name": "格林纳达",
			"chinesePhoneticize": "gelinnada",
			"english_name": "Grenada",
			"disablePrefill": false,
			"displayName": "格林纳达",
			"enable": true,
			"formal_name": "",
			"country_code": "GD",
			
			"national_flag": "🇬🇩",
			"start_char": "G",
			"telephone_code": "1473"
		    },
		    {
			"chinese_name": "格鲁吉亚",
			"chinesePhoneticize": "gelujiya",
			"english_name": "Georgia",
			"disablePrefill": false,
			"displayName": "格鲁吉亚",
			"enable": true,
			"formal_name": "Republic of Georgia",
			"country_code": "GE",
			
			"national_flag": "🇬🇪",
			"start_char": "G",
			"telephone_code": "995"
		    },
		    {
			"chinese_name": "哥伦比亚",
			"chinesePhoneticize": "gelunbiya",
			"english_name": "Colombia",
			"disablePrefill": false,
			"displayName": "哥伦比亚",
			"enable": true,
			"formal_name": "Republic of Colombia",
			"country_code": "CO",
			
			"national_flag": "🇨🇴",
			"start_char": "C",
			"telephone_code": "57"
		    },
		    {
			"chinese_name": "根西岛",
			"chinesePhoneticize": "genxidao",
			"english_name": "Guernsey",
			"disablePrefill": false,
			"displayName": "根西岛",
			"enable": true,
			"formal_name": "Bailiwick of Guernsey",
			"country_code": "GG",
			
			"national_flag": "🇬🇬",
			"start_char": "G",
			"telephone_code": "1481"
		    },
		    {
			"chinese_name": "哥斯达黎加",
			"chinesePhoneticize": "gesidalijia",
			"english_name": "Costa Rica",
			"disablePrefill": false,
			"displayName": "哥斯达黎加",
			"enable": true,
			"formal_name": "Republic of Costa Rica",
			"country_code": "CR",
			
			"national_flag": "🇨🇷",
			"start_char": "C",
			"telephone_code": "506"
		    },
		    {
			"chinese_name": "瓜德罗普",
			"chinesePhoneticize": "guadeluopu",
			"english_name": "Guadeloupe",
			"disablePrefill": false,
			"displayName": "瓜德罗普",
			"enable": true,
			"formal_name": "Guadeloupe",
			"country_code": "GP",
			
			"national_flag": "🇬🇵",
			"start_char": "G",
			"telephone_code": "590"
		    },
		    {
			"chinese_name": "关岛",
			"chinesePhoneticize": "guandao",
			"english_name": "Guam",
			"disablePrefill": false,
			"displayName": "关岛",
			"enable": true,
			"formal_name": "Guam",
			"country_code": "GU",
			
			"national_flag": "🇬🇺",
			"start_char": "G",
			"telephone_code": "1"
		    },
		    {
			"chinese_name": "古巴",
			"chinesePhoneticize": "guba",
			"english_name": "Cuba",
			"disablePrefill": false,
			"displayName": "古巴",
			"enable": false,
			"formal_name": "Republic of Cuba",
			"country_code": "CU",
			
			"national_flag": "🇨🇺",
			"start_char": "C",
			"telephone_code": "53"
		    },
		    {
			"chinese_name": "圭亚那",
			"chinesePhoneticize": "guiyana",
			"english_name": "Guyana",
			"disablePrefill": false,
			"displayName": "圭亚那",
			"enable": true,
			"formal_name": "Co-operative&nbsp;Republic&nbsp;of&nbsp;Guyana",
			"country_code": "GY",
			
			"national_flag": "🇬🇾",
			"start_char": "G",
			"telephone_code": "592"
		    }
		],
		"H": [
		    {
			"chinese_name": "海地",
			"chinesePhoneticize": "haidi",
			"english_name": "Haiti",
			"disablePrefill": false,
			"displayName": "海地",
			"enable": true,
			"formal_name": "Republic of Haiti",
			"country_code": "HT",
			
			"national_flag": "🇭🇹",
			"start_char": "H",
			"telephone_code": "509"
		    },
		    {
			"chinese_name": "韩国",
			"chinesePhoneticize": "hanguo",
			"english_name": "South Korea",
			"disablePrefill": false,
			"displayName": "韩国",
			"enable": true,
			"formal_name": "Republic of Korea",
			"country_code": "KR",
			
			"national_flag": "🇰🇷",
			"start_char": "S",
			"telephone_code": "82"
		    },
		    {
			"chinese_name": "哈萨克斯坦",
			"chinesePhoneticize": "hasakesitan",
			"english_name": "Kazakhstan",
			"disablePrefill": false,
			"displayName": "哈萨克斯坦",
			"enable": true,
			"formal_name": "Republic of Kazakhstan",
			"country_code": "KZ",
			
			"national_flag": "🇰🇿",
			"start_char": "K",
			"telephone_code": "7"
		    },
		    {
			"chinese_name": "黑山",
			"chinesePhoneticize": "heishan",
			"english_name": "Montenegro",
			"disablePrefill": false,
			"displayName": "黑山",
			"enable": true,
			"formal_name": "Republic of Montenegro",
			"country_code": "ME",
			
			"national_flag": "🇲🇪",
			"start_char": "M",
			"telephone_code": "382"
		    },
		    {
			"chinese_name": "荷兰",
			"chinesePhoneticize": "helan",
			"english_name": "Netherlands",
			"disablePrefill": false,
			"displayName": "荷兰",
			"enable": true,
			"formal_name": "Kingdom of the Netherlands",
			"country_code": "NL",
			
			"national_flag": "🇳🇱",
			"start_char": "N",
			"telephone_code": "31"
		    },
		    {
			"chinese_name": "荷属加勒比",
			"chinesePhoneticize": "heshujialebi",
			"english_name": "Caribbean Netherlands",
			"disablePrefill": false,
			"displayName": "荷属加勒比",
			"enable": true,
			"formal_name": "Caribbean Netherlands",
			"country_code": "BQ",
			
			"national_flag": "🇧🇶",
			"start_char": "C",
			"telephone_code": "599"
		    },
		    {
			"chinese_name": "荷属圣马丁",
			"chinesePhoneticize": "heshushengmading",
			"english_name": "Sint Maarten",
			"disablePrefill": false,
			"displayName": "荷属圣马丁",
			"enable": true,
			"formal_name": "Sint Maarten",
			"country_code": "SX",
			
			"national_flag": "🇸🇽",
			"start_char": "S",
			"telephone_code": "721"
		    },
		    {
			"chinese_name": "洪都拉斯",
			"chinesePhoneticize": "hongdulasi",
			"english_name": "Honduras",
			"disablePrefill": false,
			"displayName": "洪都拉斯",
			"enable": true,
			"formal_name": "Republic of Honduras",
			"country_code": "HN",
			
			"national_flag": "🇭🇳",
			"start_char": "H",
			"telephone_code": "504"
		    }
		],
		"J": [
		    {
			"chinese_name": "加纳",
			"chinesePhoneticize": "jiana",
			"english_name": "Ghana",
			"disablePrefill": false,
			"displayName": "加纳",
			"enable": true,
			"formal_name": "Republic of Ghana",
			"country_code": "GH",
			
			"national_flag": "🇬🇭",
			"start_char": "G",
			"telephone_code": "233"
		    },
		    {
			"chinese_name": "加拿大",
			"chinesePhoneticize": "jianada",
			"english_name": "Canada",
			"disablePrefill": false,
			"displayName": "加拿大",
			"enable": true,
			"formal_name": "",
			"country_code": "CA",
			
			"national_flag": "🇨🇦",
			"start_char": "C",
			"telephone_code": "1"
		    },
		    {
			"chinese_name": "柬埔寨",
			"chinesePhoneticize": "jianpuzhai",
			"english_name": "Cambodia",
			"disablePrefill": false,
			"displayName": "柬埔寨",
			"enable": true,
			"formal_name": "Kingdom of Cambodia",
			"country_code": "KH",
			
			"national_flag": "🇰🇭",
			"start_char": "C",
			"telephone_code": "855"
		    },
		    {
			"chinese_name": "加蓬",
			"chinesePhoneticize": "jiapeng",
			"english_name": "Gabon",
			"disablePrefill": false,
			"displayName": "加蓬",
			"enable": true,
			"formal_name": "Gabonese Republic",
			"country_code": "GA",
			
			"national_flag": "🇬🇦",
			"start_char": "G",
			"telephone_code": "241"
		    },
		    {
			"chinese_name": "吉布提",
			"chinesePhoneticize": "jibuti",
			"english_name": "Djibouti",
			"disablePrefill": false,
			"displayName": "吉布提",
			"enable": true,
			"formal_name": "Republic of Djibouti",
			"country_code": "DJ",
			
			"national_flag": "🇩🇯",
			"start_char": "D",
			"telephone_code": "253"
		    },
		    {
			"chinese_name": "捷克共和国",
			"chinesePhoneticize": "jiekegongheguo",
			"english_name": "Czech Republic",
			"disablePrefill": false,
			"displayName": "捷克共和国",
			"enable": true,
			"formal_name": "",
			"country_code": "CZ",
			
			"national_flag": "🇨🇿",
			"start_char": "C",
			"telephone_code": "420"
		    },
		    {
			"chinese_name": "吉尔吉斯斯坦",
			"chinesePhoneticize": "jierjisisitan",
			"english_name": "Kyrgyzstan",
			"disablePrefill": false,
			"displayName": "吉尔吉斯斯坦",
			"enable": true,
			"formal_name": "Kyrgyz&nbsp;Republic",
			"country_code": "KG",
			
			"national_flag": "🇰🇬",
			"start_char": "K",
			"telephone_code": "996"
		    },
		    {
			"chinese_name": "基里巴斯",
			"chinesePhoneticize": "jilibasi",
			"english_name": "Kiribati",
			"disablePrefill": false,
			"displayName": "基里巴斯",
			"enable": true,
			"formal_name": "Republic of Kiribati",
			"country_code": "KI",
			
			"national_flag": "🇰🇮",
			"start_char": "K",
			"telephone_code": "686"
		    },
		    {
			"chinese_name": "津巴布韦",
			"chinesePhoneticize": "jinbabuwei",
			"english_name": "Zimbabwe",
			"disablePrefill": false,
			"displayName": "津巴布韦",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Zimbabwe",
			"country_code": "ZW",
			
			"national_flag": "🇿🇼",
			"start_char": "Z",
			"telephone_code": "263"
		    },
		    {
			"chinese_name": "几内亚",
			"chinesePhoneticize": "jineiya",
			"english_name": "Guinea",
			"disablePrefill": false,
			"displayName": "几内亚",
			"enable": true,
			"formal_name": "Republic of Guinea",
			"country_code": "GN",
			
			"national_flag": "🇬🇳",
			"start_char": "G",
			"telephone_code": "224"
		    },
		    {
			"chinese_name": "几内亚比绍",
			"chinesePhoneticize": "jineiyabishao",
			"english_name": "Guinea-Bissau",
			"disablePrefill": false,
			"displayName": "几内亚比绍",
			"enable": true,
			"formal_name": "Republic of Guinea-Bissau",
			"country_code": "GW",
			
			"national_flag": "🇬🇼",
			"start_char": "G",
			"telephone_code": "245"
		    }
		],
		"K": [
		    {
			"chinese_name": "开曼群岛",
			"chinesePhoneticize": "kaimanqundao",
			"english_name": "Cayman Islands",
			"disablePrefill": false,
			"displayName": "开曼群岛",
			"enable": true,
			"formal_name": "",
			"country_code": "KY",
			
			"national_flag": "🇰🇾",
			"start_char": "C",
			"telephone_code": "1345"
		    },
		    {
			"chinese_name": "喀麦隆",
			"chinesePhoneticize": "kamailong",
			"english_name": "Cameroon",
			"disablePrefill": false,
			"displayName": "喀麦隆",
			"enable": true,
			"formal_name": "Republic of Cameroon",
			"country_code": "CM",
			
			"national_flag": "🇨🇲",
			"start_char": "C",
			"telephone_code": "237"
		    },
		    {
			"chinese_name": "卡塔尔",
			"chinesePhoneticize": "kataer",
			"english_name": "Qatar",
			"disablePrefill": false,
			"displayName": "卡塔尔",
			"enable": true,
			"formal_name": "State of Qatar",
			"country_code": "QA",
			
			"national_flag": "🇶🇦",
			"start_char": "Q",
			"telephone_code": "974"
		    },
		    {
			"chinese_name": "科科斯 (基林) 群岛",
			"chinesePhoneticize": "kekesiqundao",
			"english_name": "Cocos (Keeling) Islands",
			"disablePrefill": false,
			"displayName": "科科斯 (基林) 群岛",
			"enable": true,
			"formal_name": "Cocos (Keeling) Islands",
			"country_code": "CC",
			
			"national_flag": "🇨🇨",
			"start_char": "C",
			"telephone_code": "61"
		    },
		    {
			"chinese_name": "克罗地亚",
			"chinesePhoneticize": "keluodiya",
			"english_name": "Croatia",
			"disablePrefill": false,
			"displayName": "克罗地亚",
			"enable": true,
			"formal_name": "Republic of Croatia",
			"country_code": "HR",
			
			"national_flag": "🇭🇷",
			"start_char": "C",
			"telephone_code": "385"
		    },
		    {
			"chinese_name": "科摩罗",
			"chinesePhoneticize": "kemoluo",
			"english_name": "Comoros",
			"disablePrefill": false,
			"displayName": "科摩罗",
			"enable": true,
			"formal_name": "Union of Comoros",
			"country_code": "KM",
			
			"national_flag": "🇰🇲",
			"start_char": "C",
			"telephone_code": "269"
		    },
		    {
			"chinese_name": "肯尼亚",
			"chinesePhoneticize": "kenniya",
			"english_name": "Kenya",
			"disablePrefill": false,
			"displayName": "肯尼亚",
			"enable": true,
			"formal_name": "Republic of Kenya",
			"country_code": "KE",
			
			"national_flag": "🇰🇪",
			"start_char": "K",
			"telephone_code": "254"
		    },
		    {
			"chinese_name": "科索沃",
			"chinesePhoneticize": "kesuowo",
			"english_name": "Kosovo",
			"disablePrefill": false,
			"displayName": "科索沃",
			"enable": true,
			"formal_name": "Republic of Kosovo",
			"country_code": "XK",
			
			"national_flag": "",
			"start_char": "K",
			"telephone_code": "383"
		    },
		    {
			"chinese_name": "科威特",
			"chinesePhoneticize": "keweite",
			"english_name": "Kuwait",
			"disablePrefill": false,
			"displayName": "科威特",
			"enable": true,
			"formal_name": "State of Kuwait",
			"country_code": "KW",
			
			"national_flag": "🇰🇼",
			"start_char": "K",
			"telephone_code": "965"
		    },
		    {
			"chinese_name": "库克群岛",
			"chinesePhoneticize": "kukequndao",
			"english_name": "Cook Islands",
			"disablePrefill": false,
			"displayName": "库克群岛",
			"enable": true,
			"formal_name": "The Cook Islands",
			"country_code": "CK",
			
			"national_flag": "🇨🇰",
			"start_char": "C",
			"telephone_code": "682"
		    },
		    {
			"chinese_name": "库拉索",
			"chinesePhoneticize": "kulasuo",
			"english_name": "Curacao",
			"disablePrefill": false,
			"displayName": "库拉索",
			"enable": true,
			"formal_name": "Curacao",
			"country_code": "CW",
			
			"national_flag": "🇨🇼",
			"start_char": "C",
			"telephone_code": "599"
		    }
		],
		"L": [
		    {
			"chinese_name": "莱索托",
			"chinesePhoneticize": "laisuotuo",
			"english_name": "Lesotho",
			"disablePrefill": false,
			"displayName": "莱索托",
			"enable": true,
			"formal_name": "Kingdom of Lesotho",
			"country_code": "LS",
			
			"national_flag": "🇱🇸",
			"start_char": "L",
			"telephone_code": "266"
		    },
		    {
			"chinese_name": "老挝",
			"chinesePhoneticize": "laowo",
			"english_name": "Laos",
			"disablePrefill": false,
			"displayName": "老挝",
			"enable": true,
			"formal_name": "Lao&nbsp;People&acute;s&nbsp;Democratic&nbsp;Republic",
			"country_code": "LA",
			
			"national_flag": "🇱🇦",
			"start_char": "L",
			"telephone_code": "856"
		    },
		    {
			"chinese_name": "拉脱维亚",
			"chinesePhoneticize": "latuoweiya",
			"english_name": "Latvia",
			"disablePrefill": false,
			"displayName": "拉脱维亚",
			"enable": true,
			"formal_name": "Republic of Latvia",
			"country_code": "LV",
			
			"national_flag": "🇱🇻",
			"start_char": "L",
			"telephone_code": "371"
		    },
		    {
			"chinese_name": "黎巴嫩",
			"chinesePhoneticize": "libanen",
			"english_name": "Lebanon",
			"disablePrefill": false,
			"displayName": "黎巴嫩",
			"enable": true,
			"formal_name": "Lebanese&nbsp;Republic",
			"country_code": "LB",
			
			"national_flag": "🇱🇧",
			"start_char": "L",
			"telephone_code": "961"
		    },
		    {
			"chinese_name": "利比里亚",
			"chinesePhoneticize": "libiliya",
			"english_name": "Liberia",
			"disablePrefill": false,
			"displayName": "利比里亚",
			"enable": true,
			"formal_name": "Republic of Liberia",
			"country_code": "LR",
			
			"national_flag": "🇱🇷",
			"start_char": "L",
			"telephone_code": "231"
		    },
		    {
			"chinese_name": "利比亚",
			"chinesePhoneticize": "libiya",
			"english_name": "Libya",
			"disablePrefill": false,
			"displayName": "利比亚",
			"enable": true,
			"formal_name": "Great&nbsp;Socialist&nbsp;People&acute;s&nbsp;Libyan&nbsp;Arab&nbsp;Jamahiriya",
			"country_code": "LY",
			
			"national_flag": "🇱🇾",
			"start_char": "L",
			"telephone_code": "218"
		    },
		    {
			"chinese_name": "列支敦士登",
			"chinesePhoneticize": "liezhidunshideng",
			"english_name": "Liechtenstein",
			"disablePrefill": false,
			"displayName": "列支敦士登",
			"enable": true,
			"formal_name": "Principality of Liechtenstein",
			"country_code": "LI",
			
			"national_flag": "🇱🇮",
			"start_char": "L",
			"telephone_code": "423"
		    },
		    {
			"chinese_name": "立陶宛",
			"chinesePhoneticize": "litaowan",
			"english_name": "Lithuania",
			"disablePrefill": false,
			"displayName": "立陶宛",
			"enable": true,
			"formal_name": "Republic of Lithuania",
			"country_code": "LT",
			
			"national_flag": "🇱🇹",
			"start_char": "L",
			"telephone_code": "370"
		    },
		    {
			"chinese_name": "留尼旺",
			"chinesePhoneticize": "liuniwang",
			"english_name": "Reunion",
			"disablePrefill": false,
			"displayName": "留尼旺",
			"enable": true,
			"formal_name": "Reunion Island",
			"country_code": "RE",
			
			"national_flag": "🇷🇪",
			"start_char": "R",
			"telephone_code": "262"
		    },
		    {
			"chinese_name": "罗马尼亚",
			"chinesePhoneticize": "luomaniya",
			"english_name": "Romania",
			"disablePrefill": false,
			"displayName": "罗马尼亚",
			"enable": true,
			"formal_name": "",
			"country_code": "RO",
			
			"national_flag": "🇷🇴",
			"start_char": "R",
			"telephone_code": "40"
		    },
		    {
			"chinese_name": "卢森堡",
			"chinesePhoneticize": "lusenbao",
			"english_name": "Luxembourg",
			"disablePrefill": false,
			"displayName": "卢森堡",
			"enable": true,
			"formal_name": "Grand Duchy of Luxembourg",
			"country_code": "LU",
			
			"national_flag": "🇱🇺",
			"start_char": "L",
			"telephone_code": "352"
		    },
		    {
			"chinese_name": "卢旺达",
			"chinesePhoneticize": "luwangda",
			"english_name": "Rwanda",
			"disablePrefill": false,
			"displayName": "卢旺达",
			"enable": true,
			"formal_name": "Republic of Rwanda",
			"country_code": "RW",
			
			"national_flag": "🇷🇼",
			"start_char": "R",
			"telephone_code": "250"
		    }
		],
		"M": [
		    {
			"chinese_name": "马达加斯加",
			"chinesePhoneticize": "madajiasijia",
			"english_name": "Madagascar",
			"disablePrefill": false,
			"displayName": "马达加斯加",
			"enable": true,
			"formal_name": "Republic of Madagascar",
			"country_code": "MG",
			
			"national_flag": "🇲🇬",
			"start_char": "M",
			"telephone_code": "261"
		    },
		    {
			"chinese_name": "马尔代夫",
			"chinesePhoneticize": "maerdaifu",
			"english_name": "Maldives",
			"disablePrefill": false,
			"displayName": "马尔代夫",
			"enable": true,
			"formal_name": "Republic of Maldives",
			"country_code": "MV",
			
			"national_flag": "🇲🇻",
			"start_char": "M",
			"telephone_code": "960"
		    },
		    {
			"chinese_name": "马耳他",
			"chinesePhoneticize": "maerta",
			"english_name": "Malta",
			"disablePrefill": false,
			"displayName": "马耳他",
			"enable": true,
			"formal_name": "Republic of Malta",
			"country_code": "MT",
			
			"national_flag": "🇲🇹",
			"start_char": "M",
			"telephone_code": "356"
		    },
		    {
			"chinese_name": "马来西亚",
			"chinesePhoneticize": "malaixiya",
			"english_name": "Malaysia",
			"disablePrefill": false,
			"displayName": "马来西亚",
			"enable": true,
			"formal_name": "",
			"country_code": "MY",
			
			"national_flag": "🇲🇾",
			"start_char": "M",
			"telephone_code": "60"
		    },
		    {
			"chinese_name": "马拉维",
			"chinesePhoneticize": "malawei",
			"english_name": "Malawi",
			"disablePrefill": false,
			"displayName": "马拉维",
			"enable": true,
			"formal_name": "Republic of Malawi",
			"country_code": "MW",
			
			"national_flag": "🇲🇼",
			"start_char": "M",
			"telephone_code": "265"
		    },
		    {
			"chinese_name": "马里",
			"chinesePhoneticize": "mali",
			"english_name": "Mali",
			"disablePrefill": false,
			"displayName": "马里",
			"enable": true,
			"formal_name": "Republic of Mali",
			"country_code": "ML",
			
			"national_flag": "🇲🇱",
			"start_char": "M",
			"telephone_code": "223"
		    },
		    {
			"chinese_name": "曼岛",
			"chinesePhoneticize": "mandao",
			"english_name": "Isle of Man",
			"disablePrefill": false,
			"displayName": "马恩岛",
			"enable": true,
			"formal_name": "Isle of Man",
			"country_code": "IM",
			
			"national_flag": "🇮🇲",
			"start_char": "I",
			"telephone_code": "44"
		    },
		    {
			"chinese_name": "毛里求斯",
			"chinesePhoneticize": "maoliqiusi",
			"english_name": "Mauritius",
			"disablePrefill": false,
			"displayName": "毛里求斯",
			"enable": true,
			"formal_name": "Republic of Mauritius",
			"country_code": "MU",
			
			"national_flag": "🇲🇺",
			"start_char": "M",
			"telephone_code": "230"
		    },
		    {
			"chinese_name": "毛里塔尼亚",
			"chinesePhoneticize": "maolitaniya",
			"english_name": "Mauritania",
			"disablePrefill": false,
			"displayName": "毛里塔尼亚",
			"enable": true,
			"formal_name": "Islamic Republic of Mauritania",
			"country_code": "MR",
			
			"national_flag": "🇲🇷",
			"start_char": "M",
			"telephone_code": "222"
		    },
		    {
			"chinese_name": "马其顿",
			"chinesePhoneticize": "maqidun",
			"english_name": "Macedonia",
			"disablePrefill": false,
			"displayName": "马其顿",
			"enable": true,
			"formal_name": "Republic of Macedonia",
			"country_code": "MK",
			
			"national_flag": "🇲🇰",
			"start_char": "M",
			"telephone_code": "389"
		    },
		    {
			"chinese_name": "马绍尔群岛",
			"chinesePhoneticize": "mashaoerqundao",
			"english_name": "Marshall Islands",
			"disablePrefill": false,
			"displayName": "马绍尔群岛",
			"enable": true,
			"formal_name": "Republic of the Marshall Islands",
			"country_code": "MH",
			
			"national_flag": "🇲🇭",
			"start_char": "M",
			"telephone_code": "692"
		    },
		    {
			"chinese_name": "马提尼克",
			"chinesePhoneticize": "matinike",
			"english_name": "Martinique",
			"disablePrefill": false,
			"displayName": "马提尼克",
			"enable": true,
			"formal_name": "Martinique",
			"country_code": "MQ",
			
			"national_flag": "🇲🇶",
			"start_char": "M",
			"telephone_code": "596"
		    },
		    {
			"chinese_name": "马约特",
			"chinesePhoneticize": "mayuete",
			"english_name": "Mayotte",
			"disablePrefill": false,
			"displayName": "马约特",
			"enable": true,
			"formal_name": "Mayotte",
			"country_code": "YT",
			
			"national_flag": "🇾🇹",
			"start_char": "M",
			"telephone_code": "262"
		    },
		    {
			"chinese_name": "美国",
			"chinesePhoneticize": "meiguo",
			"english_name": "United States",
			"disablePrefill": false,
			"displayName": "美国",
			"enable": true,
			"formal_name": "United&nbsp;States&nbsp;of&nbsp;America",
			"country_code": "US",
			
			"national_flag": "🇺🇸",
			"start_char": "U",
			"telephone_code": "1"
		    },
		    {
			"chinese_name": "美属萨摩亚",
			"chinesePhoneticize": "meishusamoya",
			"english_name": "American Samoa",
			"disablePrefill": false,
			"displayName": "美属萨摩亚",
			"enable": true,
			"formal_name": "American Samoa",
			"country_code": "AS",
			
			"national_flag": "🇦🇸",
			"start_char": "A",
			"telephone_code": "1-684"
		    },
		    {
			"chinese_name": "美属维尔京群岛",
			"chinesePhoneticize": "meishuweierjingqundao",
			"english_name": "U.S. Virgin Islands",
			"disablePrefill": false,
			"displayName": "美属维尔京群岛",
			"enable": true,
			"formal_name": "United&nbsp;States&nbsp;Virgin&nbsp;Islands",
			"country_code": "VI",
			
			"national_flag": "🇻🇮",
			"start_char": "U",
			"telephone_code": "1340"
		    },
		    {
			"chinese_name": "蒙古",
			"chinesePhoneticize": "menggu",
			"english_name": "Mongolia",
			"disablePrefill": false,
			"displayName": "蒙古",
			"enable": true,
			"formal_name": "",
			"country_code": "MN",
			
			"national_flag": "🇲🇳",
			"start_char": "M",
			"telephone_code": "976"
		    },
		    {
			"chinese_name": "孟加拉国",
			"chinesePhoneticize": "mengjialaguo",
			"english_name": "Bangladesh",
			"disablePrefill": false,
			"displayName": "孟加拉国",
			"enable": true,
			"formal_name": "People&acute;s&nbsp;Republic&nbsp;of&nbsp;Bangladesh",
			"country_code": "BD",
			
			"national_flag": "🇧🇩",
			"start_char": "B",
			"telephone_code": "880"
		    },
		    {
			"chinese_name": "缅甸",
			"chinesePhoneticize": "miandian",
			"english_name": "Myanmar (Burma)",
			"disablePrefill": false,
			"displayName": "缅甸",
			"enable": true,
			"formal_name": "Union&nbsp;of&nbsp;Myanmar",
			"country_code": "MM",
			
			"national_flag": "🇲🇲",
			"start_char": "M",
			"telephone_code": "95"
		    },
		    {
			"chinese_name": "密克罗尼西亚",
			"chinesePhoneticize": "mikeluonixiya",
			"english_name": "Micronesia",
			"disablePrefill": false,
			"displayName": "密克罗尼西亚",
			"enable": true,
			"formal_name": "Federated States of Micronesia",
			"country_code": "FM",
			
			"national_flag": "🇫🇲",
			"start_char": "M",
			"telephone_code": "691"
		    },
		    {
			"chinese_name": "秘鲁",
			"chinesePhoneticize": "milu",
			"english_name": "Peru",
			"disablePrefill": false,
			"displayName": "秘鲁",
			"enable": true,
			"formal_name": "Republic of Peru",
			"country_code": "PE",
			
			"national_flag": "🇵🇪",
			"start_char": "P",
			"telephone_code": "51"
		    },
		    {
			"chinese_name": "摩尔多瓦",
			"chinesePhoneticize": "moerduowa",
			"english_name": "Moldova",
			"disablePrefill": false,
			"displayName": "摩尔多瓦",
			"enable": true,
			"formal_name": "Republic of Moldova",
			"country_code": "MD",
			
			"national_flag": "🇲🇩",
			"start_char": "M",
			"telephone_code": "373"
		    },
		    {
			"chinese_name": "摩洛哥",
			"chinesePhoneticize": "moluoge",
			"english_name": "Morocco",
			"disablePrefill": false,
			"displayName": "摩洛哥",
			"enable": true,
			"formal_name": "Kingdom of Morocco",
			"country_code": "MA",
			
			"national_flag": "🇲🇦",
			"start_char": "M",
			"telephone_code": "212"
		    },
		    {
			"chinese_name": "摩纳哥",
			"chinesePhoneticize": "monage",
			"english_name": "Monaco",
			"disablePrefill": false,
			"displayName": "摩纳哥",
			"enable": true,
			"formal_name": "Principality of Monaco",
			"country_code": "MC",
			
			"national_flag": "🇲🇨",
			"start_char": "M",
			"telephone_code": "377"
		    },
		    {
			"chinese_name": "莫桑比克",
			"chinesePhoneticize": "mosangbike",
			"english_name": "Mozambique",
			"disablePrefill": false,
			"displayName": "莫桑比克",
			"enable": true,
			"formal_name": "Republic of Mozambique",
			"country_code": "MZ",
			
			"national_flag": "🇲🇿",
			"start_char": "M",
			"telephone_code": "258"
		    },
		    {
			"chinese_name": "墨西哥",
			"chinesePhoneticize": "moxige",
			"english_name": "Mexico",
			"disablePrefill": false,
			"displayName": "墨西哥",
			"enable": true,
			"formal_name": "United Mexican States",
			"country_code": "MX",
			
			"national_flag": "🇲🇽",
			"start_char": "M",
			"telephone_code": "52"
		    }
		],
		"N": [
		    {
			"chinese_name": "纳米比亚",
			"chinesePhoneticize": "namibiya",
			"english_name": "Namibia",
			"disablePrefill": false,
			"displayName": "纳米比亚",
			"enable": true,
			"formal_name": "Republic of Namibia",
			"country_code": "NA",
			
			"national_flag": "🇳🇦",
			"start_char": "N",
			"telephone_code": "264"
		    },
		    {
			"chinese_name": "南非",
			"chinesePhoneticize": "nanfei",
			"english_name": "South Africa",
			"disablePrefill": false,
			"displayName": "南非",
			"enable": true,
			"formal_name": "Republic of South Africa",
			"country_code": "ZA",
			
			"national_flag": "🇿🇦",
			"start_char": "S",
			"telephone_code": "27"
		    },
		    {
			"chinese_name": "南苏丹",
			"chinesePhoneticize": "nansudan",
			"english_name": "South Sudan",
			"disablePrefill": false,
			"displayName": "南苏丹",
			"enable": true,
			"formal_name": "Republic of South Sudan",
			"country_code": "SS",
			
			"national_flag": "🇸🇸",
			"start_char": "S",
			"telephone_code": "211"
		    },
		    {
			"chinese_name": "瑙鲁",
			"chinesePhoneticize": "naolu",
			"english_name": "Nauru",
			"disablePrefill": false,
			"displayName": "瑙鲁",
			"enable": true,
			"formal_name": "Republic of Nauru",
			"country_code": "NR",
			
			"national_flag": "🇳🇷",
			"start_char": "N",
			"telephone_code": "674"
		    },
		    {
			"chinese_name": "尼泊尔",
			"chinesePhoneticize": "niboer",
			"english_name": "Nepal",
			"disablePrefill": false,
			"displayName": "尼泊尔",
			"enable": true,
			"formal_name": "",
			"country_code": "NP",
			
			"national_flag": "🇳🇵",
			"start_char": "N",
			"telephone_code": "977"
		    },
		    {
			"chinese_name": "尼加拉瓜",
			"chinesePhoneticize": "nijialagua",
			"english_name": "Nicaragua",
			"disablePrefill": false,
			"displayName": "尼加拉瓜",
			"enable": true,
			"formal_name": "Republic of Nicaragua",
			"country_code": "NI",
			
			"national_flag": "🇳🇮",
			"start_char": "N",
			"telephone_code": "505"
		    },
		    {
			"chinese_name": "尼日尔",
			"chinesePhoneticize": "nirier",
			"english_name": "Niger",
			"disablePrefill": false,
			"displayName": "尼日尔",
			"enable": true,
			"formal_name": "Republic of Niger",
			"country_code": "NE",
			
			"national_flag": "🇳🇪",
			"start_char": "N",
			"telephone_code": "227"
		    },
		    {
			"chinese_name": "尼日利亚",
			"chinesePhoneticize": "niriliya",
			"english_name": "Nigeria",
			"disablePrefill": false,
			"displayName": "尼日利亚",
			"enable": true,
			"formal_name": "Federal Republic of Nigeria",
			"country_code": "NG",
			
			"national_flag": "🇳🇬",
			"start_char": "N",
			"telephone_code": "234"
		    },
		    {
			"chinese_name": "纽埃",
			"chinesePhoneticize": "niuai",
			"english_name": "Niue",
			"disablePrefill": false,
			"displayName": "纽埃",
			"enable": true,
			"formal_name": "Niue",
			"country_code": "NU",
			
			"national_flag": "🇳🇺",
			"start_char": "N",
			"telephone_code": "683"
		    },
		    {
			"chinese_name": "诺福克岛",
			"chinesePhoneticize": "nuofukedao",
			"english_name": "Norfolk Island",
			"disablePrefill": false,
			"displayName": "诺福克岛",
			"enable": true,
			"formal_name": "Norfolk Island",
			"country_code": "NF",
			
			"national_flag": "🇳🇫",
			"start_char": "N",
			"telephone_code": "672"
		    },
		    {
			"chinese_name": "挪威",
			"chinesePhoneticize": "nuowei",
			"english_name": "Norway",
			"disablePrefill": false,
			"displayName": "挪威",
			"enable": true,
			"formal_name": "Kingdom of Norway",
			"country_code": "NO",
			
			"national_flag": "🇳🇴",
			"start_char": "N",
			"telephone_code": "47"
		    }
		],
		"P": [
		    {
			"chinese_name": "帕劳",
			"chinesePhoneticize": "palao",
			"english_name": "Palau",
			"disablePrefill": false,
			"displayName": "帕劳",
			"enable": true,
			"formal_name": "Republic of Palau",
			"country_code": "PW",
			
			"national_flag": "🇵🇼",
			"start_char": "P",
			"telephone_code": "680"
		    },
		    {
			"chinese_name": "皮特凯恩群岛",
			"chinesePhoneticize": "pitekaienqundao",
			"english_name": "Pitcairn Islands",
			"disablePrefill": false,
			"displayName": "皮特凯恩群岛",
			"enable": true,
			"formal_name": "Pitcairn Islands",
			"country_code": "PN",
			
			"national_flag": "🇵🇳",
			"start_char": "P",
			"telephone_code": "64"
		    },
		    {
			"chinese_name": "葡萄牙",
			"chinesePhoneticize": "putaoya",
			"english_name": "Portugal",
			"disablePrefill": false,
			"displayName": "葡萄牙",
			"enable": true,
			"formal_name": "Portuguese Republic",
			"country_code": "PT",
			
			"national_flag": "🇵🇹",
			"start_char": "P",
			"telephone_code": "351"
		    }
		],
		"R": [
		    {
			"chinese_name": "日本",
			"chinesePhoneticize": "riben",
			"english_name": "Japan",
			"disablePrefill": false,
			"displayName": "日本",
			"enable": true,
			"formal_name": "",
			"country_code": "JP",
			
			"national_flag": "🇯🇵",
			"start_char": "J",
			"telephone_code": "81"
		    },
		    {
			"chinese_name": "瑞典",
			"chinesePhoneticize": "ruidian",
			"english_name": "Sweden",
			"disablePrefill": false,
			"displayName": "瑞典",
			"enable": true,
			"formal_name": "Kingdom of Sweden",
			"country_code": "SE",
			
			"national_flag": "🇸🇪",
			"start_char": "S",
			"telephone_code": "46"
		    },
		    {
			"chinese_name": "瑞士",
			"chinesePhoneticize": "ruishi",
			"english_name": "Switzerland",
			"disablePrefill": false,
			"displayName": "瑞士",
			"enable": true,
			"formal_name": "Swiss Confederation",
			"country_code": "CH",
			
			"national_flag": "🇨🇭",
			"start_char": "S",
			"telephone_code": "41"
		    }
		],
		"S": [
		    {
			"chinese_name": "萨尔瓦多",
			"chinesePhoneticize": "saerwaduo",
			"english_name": "El Salvador",
			"disablePrefill": false,
			"displayName": "萨尔瓦多",
			"enable": true,
			"formal_name": "Republic of El Salvador",
			"country_code": "SV",
			
			"national_flag": "🇸🇻",
			"start_char": "E",
			"telephone_code": "503"
		    },
		    {
			"chinese_name": "塞尔维亚",
			"chinesePhoneticize": "saierweiya",
			"english_name": "Serbia",
			"disablePrefill": false,
			"displayName": "塞尔维亚",
			"enable": true,
			"formal_name": "Republic of Serbia",
			"country_code": "RS",
			
			"national_flag": "🇷🇸",
			"start_char": "S",
			"telephone_code": "381"
		    },
		    {
			"chinese_name": "塞拉利昂",
			"chinesePhoneticize": "sailaliang",
			"english_name": "Sierra Leone",
			"disablePrefill": false,
			"displayName": "塞拉利昂",
			"enable": true,
			"formal_name": "Republic of Sierra Leone",
			"country_code": "SL",
			
			"national_flag": "🇸🇱",
			"start_char": "S",
			"telephone_code": "232"
		    },
		    {
			"chinese_name": "塞内加尔",
			"chinesePhoneticize": "saineijiaer",
			"english_name": "Senegal",
			"disablePrefill": false,
			"displayName": "塞内加尔",
			"enable": true,
			"formal_name": "Republic of Senegal",
			"country_code": "SN",
			
			"national_flag": "🇸🇳",
			"start_char": "S",
			"telephone_code": "221"
		    },
		    {
			"chinese_name": "塞浦路斯",
			"chinesePhoneticize": "saipulusi",
			"english_name": "Cyprus",
			"disablePrefill": false,
			"displayName": "塞浦路斯",
			"enable": true,
			"formal_name": "Republic of Cyprus",
			"country_code": "CY",
			
			"national_flag": "🇨🇾",
			"start_char": "C",
			"telephone_code": "357"
		    },
		    {
			"chinese_name": "塞舌尔",
			"chinesePhoneticize": "saisheer",
			"english_name": "Seychelles",
			"disablePrefill": false,
			"displayName": "塞舌尔",
			"enable": true,
			"formal_name": "Republic of Seychelles",
			"country_code": "SC",
			
			"national_flag": "🇸🇨",
			"start_char": "S",
			"telephone_code": "248"
		    },
		    {
			"chinese_name": "萨摩亚",
			"chinesePhoneticize": "samoya",
			"english_name": "Samoa",
			"disablePrefill": false,
			"displayName": "萨摩亚",
			"enable": true,
			"formal_name": "Independent State of Samoa",
			"country_code": "WS",
			
			"national_flag": "🇼🇸",
			"start_char": "S",
			"telephone_code": "685"
		    },
		    {
			"chinese_name": "圣皮埃尔和密克隆",
			"chinesePhoneticize": "sehngpiaierhemikelong",
			"english_name": "Saint Pierre and Miquelon",
			"disablePrefill": false,
			"displayName": "圣皮埃尔和密克隆",
			"enable": true,
			"formal_name": "Saint Pierre and Miquelon",
			"country_code": "PM",
			
			"national_flag": "🇵🇲",
			"start_char": "S",
			"telephone_code": "508"
		    },
		    {
			"chinese_name": "沙特阿拉伯",
			"chinesePhoneticize": "shatealabo",
			"english_name": "Saudi Arabia",
			"disablePrefill": false,
			"displayName": "沙特阿拉伯",
			"enable": true,
			"formal_name": "Kingdom of Saudi Arabia",
			"country_code": "SA",
			
			"national_flag": "🇸🇦",
			"start_char": "S",
			"telephone_code": "966"
		    },
		    {
			"chinese_name": "圣诞岛",
			"chinesePhoneticize": "shengdandao",
			"english_name": "Christmas  as Island",
			"disablePrefill": false,
			"displayName": "圣诞岛",
			"enable": true,
			"formal_name": "Christmas Island",
			"country_code": "CX",
			
			"national_flag": "🇨🇽",
			"start_char": "C",
			"telephone_code": "61"
		    },
		    {
			"chinese_name": "圣多美和普林西比",
			"chinesePhoneticize": "shengduomeihepulinxibi",
			"english_name": "Sao Tome and Principe",
			"disablePrefill": false,
			"displayName": "圣多美和普林西比",
			"enable": true,
			"formal_name": "Democratic Republic of Sao Tome and Principe",
			"country_code": "ST",
			
			"national_flag": "🇸🇹",
			"start_char": "S",
			"telephone_code": "239"
		    },
		    {
			"chinese_name": "圣赫勒拿、阿森松岛和特里斯坦-达库尼亚",
			"chinesePhoneticize": "shenghelenaasensongdaohetelisitan",
			"english_name": "Saint Helena, Ascension and Tristan da Cunha",
			"disablePrefill": false,
			"displayName": "圣赫勒拿、阿森松岛和特里斯坦-达库尼亚",
			"enable": true,
			"formal_name": "Saint Helena, Ascension and Tristan da Cunha",
			"country_code": "SH",
			
			"national_flag": "🇸🇭",
			"start_char": "S",
			"telephone_code": "290"
		    },
		    {
			"chinese_name": "圣基茨和尼维斯",
			"chinesePhoneticize": "shengjiciheniweisi",
			"english_name": "Saint Kitts and Nevis",
			"disablePrefill": false,
			"displayName": "圣基茨和尼维斯",
			"enable": true,
			"formal_name": "Federation of Saint Kitts and Nevis",
			"country_code": "KN",
			
			"national_flag": "🇰🇳",
			"start_char": "S",
			"telephone_code": "1869"
		    },
		    {
			"chinese_name": "圣卢西亚",
			"chinesePhoneticize": "shengluxiya",
			"english_name": "Saint Lucia",
			"disablePrefill": false,
			"displayName": "圣卢西亚",
			"enable": true,
			"formal_name": "",
			"country_code": "LC",
			
			"national_flag": "🇱🇨",
			"start_char": "S",
			"telephone_code": "1758"
		    },
		    {
			"chinese_name": "圣马力诺",
			"chinesePhoneticize": "shengmalinuo",
			"english_name": "San Marino",
			"disablePrefill": false,
			"displayName": "圣马力诺",
			"enable": true,
			"formal_name": "Republic of San Marino",
			"country_code": "SM",
			
			"national_flag": "🇸🇲",
			"start_char": "S",
			"telephone_code": "378"
		    },
		    {
			"chinese_name": "圣文森特和格林纳丁斯",
			"chinesePhoneticize": "shengwensentehegelinnadingsi",
			"english_name": "Saint Vincent and the Grenadines",
			"disablePrefill": false,
			"displayName": "圣文森特和格林纳丁斯",
			"enable": true,
			"formal_name": "",
			"country_code": "VC",
			
			"national_flag": "🇻🇨",
			"start_char": "S",
			"telephone_code": "1784"
		    },
		    {
			"chinese_name": "斯里兰卡",
			"chinesePhoneticize": "sililanka",
			"english_name": "Sri Lanka",
			"disablePrefill": false,
			"displayName": "斯里兰卡",
			"enable": true,
			"formal_name": "Democratic Socialist Republic of Sri Lanka",
			"country_code": "LK",
			
			"national_flag": "🇱🇰",
			"start_char": "S",
			"telephone_code": "94"
		    },
		    {
			"chinese_name": "斯洛伐克",
			"chinesePhoneticize": "siluofake",
			"english_name": "Slovakia",
			"disablePrefill": false,
			"displayName": "斯洛伐克",
			"enable": true,
			"formal_name": "Slovak Republic",
			"country_code": "SK",
			
			"national_flag": "🇸🇰",
			"start_char": "S",
			"telephone_code": "421"
		    },
		    {
			"chinese_name": "斯洛文尼亚",
			"chinesePhoneticize": "siluowenniya",
			"english_name": "Slovenia",
			"disablePrefill": false,
			"displayName": "斯洛文尼亚",
			"enable": true,
			"formal_name": "Republic of Slovenia",
			"country_code": "SI",
			
			"national_flag": "🇸🇮",
			"start_char": "S",
			"telephone_code": "386"
		    },
		    {
			"chinese_name": "斯威士兰",
			"chinesePhoneticize": "siweishilan",
			"english_name": "Swaziland",
			"disablePrefill": false,
			"displayName": "斯威士兰",
			"enable": true,
			"formal_name": "Kingdom of Swaziland",
			"country_code": "SZ",
			
			"national_flag": "🇸🇿",
			"start_char": "S",
			"telephone_code": "268"
		    },
		    {
			"chinese_name": "苏丹",
			"chinesePhoneticize": "sudan",
			"english_name": "Sudan",
			"disablePrefill": false,
			"displayName": "苏丹",
			"enable": true,
			"formal_name": "Republic of the Sudan",
			"country_code": "SD",
			
			"national_flag": "🇸🇩",
			"start_char": "S",
			"telephone_code": "249"
		    },
		    {
			"chinese_name": "苏里南",
			"chinesePhoneticize": "sulinan",
			"english_name": "Suriname",
			"disablePrefill": false,
			"displayName": "苏里南",
			"enable": true,
			"formal_name": "Republic of Suriname",
			"country_code": "SR",
			
			"national_flag": "🇸🇷",
			"start_char": "S",
			"telephone_code": "597"
		    },
		    {
			"chinese_name": "所罗门群岛",
			"chinesePhoneticize": "suoluomenqundao",
			"english_name": "Solomon Islands",
			"disablePrefill": false,
			"displayName": "所罗门群岛",
			"enable": true,
			"formal_name": "",
			"country_code": "SB",
			
			"national_flag": "🇸🇧",
			"start_char": "S",
			"telephone_code": "677"
		    },
		    {
			"chinese_name": "索马里",
			"chinesePhoneticize": "suomali",
			"english_name": "Somalia",
			"disablePrefill": false,
			"displayName": "索马里",
			"enable": true,
			"formal_name": "null",
			"country_code": "SO",
			
			"national_flag": "🇸🇴",
			"start_char": "S",
			"telephone_code": "252"
		    }
		],
		"T": [
		    {
			"chinese_name": "泰国",
			"chinesePhoneticize": "taiguo",
			"english_name": "Thailand",
			"disablePrefill": false,
			"displayName": "泰国",
			"enable": true,
			"formal_name": "Kingdom of Thailand",
			"country_code": "TH",
			
			"national_flag": "🇹🇭",
			"start_char": "T",
			"telephone_code": "66"
		    },
		    {
			"chinese_name": "台湾",
			"chinesePhoneticize": "taiwan",
			"english_name": "Taiwan",
			"disablePrefill": false,
			"displayName": "台湾",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;China",
			"country_code": "TW",
			
			"national_flag": "",
			"start_char": "T",
			"telephone_code": "886"
		    },
		    {
			"chinese_name": "塔吉克斯坦",
			"chinesePhoneticize": "tajikesitan",
			"english_name": "Tajikistan",
			"disablePrefill": false,
			"displayName": "塔吉克斯坦",
			"enable": true,
			"formal_name": "Republic of Tajikistan",
			"country_code": "TJ",
			
			"national_flag": "🇹🇯",
			"start_char": "T",
			"telephone_code": "992"
		    },
		    {
			"chinese_name": "汤加",
			"chinesePhoneticize": "tangjia",
			"english_name": "Tonga",
			"disablePrefill": false,
			"displayName": "汤加",
			"enable": true,
			"formal_name": "Kingdom of Tonga",
			"country_code": "TO",
			
			"national_flag": "🇹🇴",
			"start_char": "T",
			"telephone_code": "676"
		    },
		    {
			"chinese_name": "坦桑尼亚",
			"chinesePhoneticize": "tansangniya",
			"english_name": "Tanzania",
			"disablePrefill": false,
			"displayName": "坦桑尼亚",
			"enable": true,
			"formal_name": "United Republic of Tanzania",
			"country_code": "TZ",
			
			"national_flag": "🇹🇿",
			"start_char": "T",
			"telephone_code": "255"
		    },
		    {
			"chinese_name": "特克斯和凯科斯群岛",
			"chinesePhoneticize": "tekesihekaikesiqundao",
			"english_name": "Turks and Caicos",
			"disablePrefill": false,
			"displayName": "特克斯和凯科斯群岛",
			"enable": true,
			"formal_name": "Turks and Caicos",
			"country_code": "TC",
			
			"national_flag": "🇹🇨",
			"start_char": "T",
			"telephone_code": "1"
		    },
		    {
			"chinese_name": "特里尼达和多巴哥",
			"chinesePhoneticize": "telinidaheduobage",
			"english_name": "Trinidad and Tobago",
			"disablePrefill": false,
			"displayName": "特里尼达和多巴哥",
			"enable": true,
			"formal_name": "Republic of Trinidad and Tobago",
			"country_code": "TT",
			
			"national_flag": "🇹🇹",
			"start_char": "T",
			"telephone_code": "1868"
		    },
		    {
			"chinese_name": "土耳其",
			"chinesePhoneticize": "tuerqi",
			"english_name": "Turkey",
			"disablePrefill": false,
			"displayName": "土耳其",
			"enable": true,
			"formal_name": "Republic of Turkey",
			"country_code": "TR",
			
			"national_flag": "🇹🇷",
			"start_char": "T",
			"telephone_code": "90"
		    },
		    {
			"chinese_name": "土库曼斯坦",
			"chinesePhoneticize": "tukumansitan",
			"english_name": "Turkmenistan",
			"disablePrefill": false,
			"displayName": "土库曼斯坦",
			"enable": true,
			"formal_name": "",
			"country_code": "TM",
			
			"national_flag": "🇹🇲",
			"start_char": "T",
			"telephone_code": "993"
		    },
		    {
			"chinese_name": "突尼斯",
			"chinesePhoneticize": "tunisi",
			"english_name": "Tunisia",
			"disablePrefill": false,
			"displayName": "突尼斯",
			"enable": true,
			"formal_name": "Tunisian Republic",
			"country_code": "TN",
			
			"national_flag": "🇹🇳",
			"start_char": "T",
			"telephone_code": "216"
		    },
		    {
			"chinese_name": "托克劳",
			"chinesePhoneticize": "tuokelao",
			"english_name": "Tokelau",
			"disablePrefill": false,
			"displayName": "托克劳",
			"enable": true,
			"formal_name": "Tokelau",
			"country_code": "TK",
			
			"national_flag": "🇹🇰",
			"start_char": "T",
			"telephone_code": "690"
		    },
		    {
			"chinese_name": "图瓦卢",
			"chinesePhoneticize": "tuwalu",
			"english_name": "Tuvalu",
			"disablePrefill": false,
			"displayName": "图瓦卢",
			"enable": true,
			"formal_name": "",
			"country_code": "TV",
			
			"national_flag": "🇹🇻",
			"start_char": "T",
			"telephone_code": "688"
		    }
		],
		"W": [
		    {
			"chinese_name": "瓦利斯和富图纳",
			"chinesePhoneticize": "walisihefutuna",
			"english_name": "Wallis and Futuna",
			"disablePrefill": false,
			"displayName": "瓦利斯和富图纳",
			"enable": true,
			"formal_name": "Wallis and Futuna",
			"country_code": "WF",
			
			"national_flag": "🇼🇫",
			"start_char": "W",
			"telephone_code": "681"
		    },
		    {
			"chinese_name": "瓦努阿图",
			"chinesePhoneticize": "wanuatu",
			"english_name": "Vanuatu",
			"disablePrefill": false,
			"displayName": "瓦努阿图",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Vanuatu",
			"country_code": "VU",
			
			"national_flag": "🇻🇺",
			"start_char": "V",
			"telephone_code": "678"
		    },
		    {
			"chinese_name": "危地马拉",
			"chinesePhoneticize": "weidimala",
			"english_name": "Guatemala",
			"disablePrefill": false,
			"displayName": "危地马拉",
			"enable": true,
			"formal_name": "Republic of Guatemala",
			"country_code": "GT",
			
			"national_flag": "🇬🇹",
			"start_char": "G",
			"telephone_code": "502"
		    },
		    {
			"chinese_name": "委内瑞拉",
			"chinesePhoneticize": "weineiruila",
			"english_name": "Venezuela",
			"disablePrefill": false,
			"displayName": "委内瑞拉",
			"enable": true,
			"formal_name": "Bolivarian&nbsp;Republic&nbsp;of&nbsp;Venezuela",
			"country_code": "VE",
			
			"national_flag": "🇻🇪",
			"start_char": "V",
			"telephone_code": "58"
		    },
		    {
			"chinese_name": "文莱",
			"chinesePhoneticize": "wenlai",
			"english_name": "Brunei",
			"disablePrefill": false,
			"displayName": "文莱",
			"enable": true,
			"formal_name": "Negara Brunei Darussalam",
			"country_code": "BN",
			
			"national_flag": "🇧🇳",
			"start_char": "B",
			"telephone_code": "673"
		    },
		    {
			"chinese_name": "乌干达",
			"chinesePhoneticize": "wuganda",
			"english_name": "Uganda",
			"disablePrefill": false,
			"displayName": "乌干达",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Uganda",
			"country_code": "UG",
			
			"national_flag": "🇺🇬",
			"start_char": "U",
			"telephone_code": "256"
		    },
		    {
			"chinese_name": "乌克兰",
			"chinesePhoneticize": "wukelan",
			"english_name": "Ukraine",
			"disablePrefill": false,
			"displayName": "乌克兰",
			"enable": true,
			"formal_name": "",
			"country_code": "UA",
			
			"national_flag": "🇺🇦",
			"start_char": "U",
			"telephone_code": "380"
		    },
		    {
			"chinese_name": "乌拉圭",
			"chinesePhoneticize": "wulagui",
			"english_name": "Uruguay",
			"disablePrefill": false,
			"displayName": "乌拉圭",
			"enable": true,
			"formal_name": "Oriental Republic of Uruguay",
			"country_code": "UY",
			
			"national_flag": "🇺🇾",
			"start_char": "U",
			"telephone_code": "598"
		    },
		    {
			"chinese_name": "乌兹别克斯坦",
			"chinesePhoneticize": "wuzibiekesitan",
			"english_name": "Uzbekistan",
			"disablePrefill": false,
			"displayName": "乌兹别克斯坦",
			"enable": true,
			"formal_name": "Republic of Uzbekistan",
			"country_code": "UZ",
			
			"national_flag": "🇺🇿",
			"start_char": "U",
			"telephone_code": "998"
		    }
		],
		"X": [
		    {
			"chinese_name": "香港",
			"chinesePhoneticize": "xianggang",
			"english_name": "Hong Kong",
			"disablePrefill": false,
			"displayName": "中国香港",
			"enable": true,
			"formal_name": "Hong&nbsp;Kong&nbsp;Special&nbsp;Administrative&nbsp;Region",
			"country_code": "HK",
			
			"national_flag": "🇭🇰",
			"start_char": "H",
			"telephone_code": "852"
		    },
		    {
			"chinese_name": "象牙海岸",
			"chinesePhoneticize": "xiangyahaian",
			"english_name": "Cote d´Ivoire (Ivory Coast)",
			"disablePrefill": false,
			"displayName": "象牙海岸",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Cote&nbsp;d&acute;Ivoire",
			"country_code": "CI",
			
			"national_flag": "🇨🇮",
			"start_char": "C",
			"telephone_code": "225"
		    },
		    {
			"chinese_name": "西班牙",
			"chinesePhoneticize": "xibanya",
			"english_name": "Spain",
			"disablePrefill": false,
			"displayName": "西班牙",
			"enable": true,
			"formal_name": "Kingdom of Spain",
			"country_code": "ES",
			
			"national_flag": "🇪🇸",
			"start_char": "S",
			"telephone_code": "34"
		    },
		    {
			"chinese_name": "希腊",
			"chinesePhoneticize": "xila",
			"english_name": "Greece",
			"disablePrefill": false,
			"displayName": "希腊",
			"enable": true,
			"formal_name": "Hellenic Republic",
			"country_code": "GR",
			
			"national_flag": "🇬🇷",
			"start_char": "G",
			"telephone_code": "30"
		    },
		    {
			"chinese_name": "新加坡",
			"chinesePhoneticize": "xinjiapo",
			"english_name": "Singapore",
			"disablePrefill": false,
			"displayName": "新加坡",
			"enable": true,
			"formal_name": "Republic of Singapore",
			"country_code": "SG",
			
			"national_flag": "🇸🇬",
			"start_char": "S",
			"telephone_code": "65"
		    },
		    {
			"chinese_name": "新喀里多尼亚",
			"chinesePhoneticize": "xinkaliduoniya",
			"english_name": "New Caledonia",
			"disablePrefill": false,
			"displayName": "新喀里多尼亚",
			"enable": true,
			"formal_name": "New Caledonia",
			"country_code": "NC",
			
			"national_flag": "🇳🇨",
			"start_char": "N",
			"telephone_code": "687"
		    },
		    {
			"chinese_name": "新西兰",
			"chinesePhoneticize": "xinxilan",
			"english_name": "New Zealand",
			"disablePrefill": false,
			"displayName": "新西兰",
			"enable": true,
			"formal_name": "",
			"country_code": "NZ",
			
			"national_flag": "🇳🇿",
			"start_char": "N",
			"telephone_code": "64"
		    },
		    {
			"chinese_name": "匈牙利",
			"chinesePhoneticize": "xiongyali",
			"english_name": "Hungary",
			"disablePrefill": false,
			"displayName": "匈牙利",
			"enable": true,
			"formal_name": "Republic of Hungary",
			"country_code": "HU",
			
			"national_flag": "🇭🇺",
			"start_char": "H",
			"telephone_code": "36"
		    },
		    {
			"chinese_name": "叙利亚",
			"chinesePhoneticize": "xuliya",
			"english_name": "Syria",
			"disablePrefill": false,
			"displayName": "叙利亚",
			"enable": false,
			"formal_name": "Syrian Arab Republic",
			"country_code": "SY",
			
			"national_flag": "🇸🇾",
			"start_char": "S",
			"telephone_code": "963"
		    }
		],
		"Y": [
		    {
			"chinese_name": "牙买加",
			"chinesePhoneticize": "yamaijia",
			"english_name": "Jamaica",
			"disablePrefill": false,
			"displayName": "牙买加",
			"enable": true,
			"formal_name": "",
			"country_code": "JM",
			
			"national_flag": "🇯🇲",
			"start_char": "J",
			"telephone_code": "1876"
		    },
		    {
			"chinese_name": "亚美尼亚",
			"chinesePhoneticize": "yameiniya",
			"english_name": "Armenia",
			"disablePrefill": false,
			"displayName": "亚美尼亚",
			"enable": true,
			"formal_name": "Republic of Armenia",
			"country_code": "AM",
			
			"national_flag": "🇦🇲",
			"start_char": "A",
			"telephone_code": "374"
		    },
		    {
			"chinese_name": "也门",
			"chinesePhoneticize": "yemen",
			"english_name": "Yemen",
			"disablePrefill": false,
			"displayName": "也门",
			"enable": true,
			"formal_name": "Republic of Yemen",
			"country_code": "YE",
			
			"national_flag": "🇾🇪",
			"start_char": "Y",
			"telephone_code": "967"
		    },
		    {
			"chinese_name": "意大利",
			"chinesePhoneticize": "yidali",
			"english_name": "Italy",
			"disablePrefill": false,
			"displayName": "意大利",
			"enable": true,
			"formal_name": "Italian Republic",
			"country_code": "IT",
			
			"national_flag": "🇮🇹",
			"start_char": "I",
			"telephone_code": "39"
		    },
		    {
			"chinese_name": "伊拉克",
			"chinesePhoneticize": "yilake",
			"english_name": "Iraq",
			"disablePrefill": false,
			"displayName": "伊拉克",
			"enable": true,
			"formal_name": "Republic&nbsp;of&nbsp;Iraq",
			"country_code": "IQ",
			
			"national_flag": "🇮🇶",
			"start_char": "I",
			"telephone_code": "964"
		    },
		    {
			"chinese_name": "伊朗",
			"chinesePhoneticize": "yilang",
			"english_name": "Iran",
			"disablePrefill": false,
			"displayName": "伊朗",
			"enable": false,
			"formal_name": "Islamic Republic of Iran",
			"country_code": "IR",
			
			"national_flag": "🇮🇷",
			"start_char": "I",
			"telephone_code": "98"
		    },
		    {
			"chinese_name": "印度",
			"chinesePhoneticize": "yindu",
			"english_name": "India",
			"disablePrefill": false,
			"displayName": "印度",
			"enable": true,
			"formal_name": "Republic of India",
			"country_code": "IN",
			
			"national_flag": "🇮🇳",
			"start_char": "I",
			"telephone_code": "91"
		    },
		    {
			"chinese_name": "印度尼西亚",
			"chinesePhoneticize": "yindunixiya",
			"english_name": "Indonesia",
			"disablePrefill": false,
			"displayName": "印度尼西亚",
			"enable": true,
			"formal_name": "Republic of Indonesia",
			"country_code": "ID",
			
			"national_flag": "🇮🇩",
			"start_char": "I",
			"telephone_code": "62"
		    },
		    {
			"chinese_name": "英国",
			"chinesePhoneticize": "yingguo",
			"english_name": "United Kingdom",
			"disablePrefill": false,
			"displayName": "英国",
			"enable": true,
			"formal_name": "United Kingdom of Great Britain and Northern Ireland",
			"country_code": "GB",
			
			"national_flag": "🇬🇧",
			"start_char": "U",
			"telephone_code": "44"
		    },
		    {
			"chinese_name": "英属维尔京群岛",
			"chinesePhoneticize": "yingshuweierjingqundao",
			"english_name": "British Virgin Islands",
			"disablePrefill": false,
			"displayName": "英属维尔京群岛",
			"enable": true,
			"formal_name": "",
			"country_code": "VG",
			
			"national_flag": "🇻🇬",
			"start_char": "B",
			"telephone_code": "1284"
		    },
		    {
			"chinese_name": "英属印度洋领地",
			"chinesePhoneticize": "yingshuyinduyanglingdi",
			"english_name": "British Indian Ocean Territory",
			"disablePrefill": false,
			"displayName": "英属印度洋领地",
			"enable": true,
			"formal_name": "British Indian Ocean Territory",
			"country_code": "IO",
			
			"national_flag": "🇮🇴",
			"start_char": "B",
			"telephone_code": "246"
		    },
		    {
			"chinese_name": "英属直布罗陀",
			"chinesePhoneticize": "yingshuzhibuluotuo",
			"english_name": "Gibraltar",
			"disablePrefill": false,
			"displayName": "英属直布罗陀",
			"enable": true,
			"formal_name": "The British Overseas Territory of Gibraltar",
			"country_code": "GI",
			
			"national_flag": "🇬🇮",
			"start_char": "G",
			"telephone_code": "350"
		    },
		    {
			"chinese_name": "以色列",
			"chinesePhoneticize": "yiselie",
			"english_name": "Israel",
			"disablePrefill": false,
			"displayName": "以色列",
			"enable": true,
			"formal_name": "State of Israel",
			"country_code": "IL",
			
			"national_flag": "🇮🇱",
			"start_char": "I",
			"telephone_code": "972"
		    },
		    {
			"chinese_name": "约旦",
			"chinesePhoneticize": "yuedan",
			"english_name": "Jordan",
			"disablePrefill": false,
			"displayName": "约旦",
			"enable": true,
			"formal_name": "Hashemite Kingdom of Jordan",
			"country_code": "JO",
			
			"national_flag": "🇯🇴",
			"start_char": "J",
			"telephone_code": "962"
		    },
		    {
			"chinese_name": "越南",
			"chinesePhoneticize": "yuenan",
			"english_name": "Vietnam",
			"disablePrefill": false,
			"displayName": "越南",
			"enable": true,
			"formal_name": "Socialist Republic of Vietnam",
			"country_code": "VN",
			
			"national_flag": "🇻🇳",
			"start_char": "V",
			"telephone_code": "84"
		    }
		],
		"Z": [
		    {
			"chinese_name": "赞比亚",
			"chinesePhoneticize": "zanbiya",
			"english_name": "Zambia",
			"disablePrefill": false,
			"displayName": "赞比亚",
			"enable": true,
			"formal_name": "Republic of Zambia",
			"country_code": "ZM",
			
			"national_flag": "🇿🇲",
			"start_char": "Z",
			"telephone_code": "260"
		    },
		    {
			"chinese_name": "泽西岛",
			"chinesePhoneticize": "zexidao",
			"english_name": "Jersey",
			"disablePrefill": false,
			"displayName": "泽西岛",
			"enable": true,
			"formal_name": "Bailiwick of Jersey",
			"country_code": "JE",
			
			"national_flag": "🇯🇪",
			"start_char": "J",
			"telephone_code": "44"
		    },
		    {
			"chinese_name": "乍得",
			"chinesePhoneticize": "zhade",
			"english_name": "Chad",
			"disablePrefill": false,
			"displayName": "乍得",
			"enable": true,
			"formal_name": "Republic of Chad",
			"country_code": "TD",
			
			"national_flag": "🇹🇩",
			"start_char": "C",
			"telephone_code": "235"
		    },
		    {
			"chinese_name": "智利",
			"chinesePhoneticize": "zhili",
			"english_name": "Chile",
			"disablePrefill": false,
			"displayName": "智利",
			"enable": true,
			"formal_name": "Republic of Chile",
			"country_code": "CL",
			
			"national_flag": "🇨🇱",
			"start_char": "C",
			"telephone_code": "56"
		    },
		    {
			"chinese_name": "中非共和国",
			"chinesePhoneticize": "zhongfeigongheguo",
			"english_name": "Central African Republic",
			"disablePrefill": false,
			"displayName": "中非共和国",
			"enable": true,
			"formal_name": "null",
			"country_code": "CF",
			
			"national_flag": "🇨🇫",
			"start_char": "C",
			"telephone_code": "236"
		    },
		    {
			"chinese_name": "中国",
			"chinesePhoneticize": "zhongguo",
			"english_name": "China",
			"disablePrefill": false,
			"displayName": "中国",
			"enable": true,
			"formal_name": "People´s Republic of China",
			"country_code": "CN",
			
			"national_flag": "🇨🇳",
			"start_char": "C",
			"telephone_code": "86"
		    }
		]
	    }`

	maps := new(map[string][]tables.TCountryPhoneCode)
	if err := json.Unmarshal([]byte(data), maps); err != nil {
		global.Logger.Error(err.Error())
	}
	// global.Logger.Error(maps)

	countrys := make([]tables.TCountryPhoneCode, 0)

	for _, cs := range *maps {
		countrys = append(countrys, cs...)
	}

	// global.Logger.Error(countrys)

	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("country_code = ?", "CN").Find(&tables.TCountryPhoneCode{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> t_country_phone_code 表的数据已存在!")
			return nil
		}
		if err := tx.Create(&countrys).Error; err != nil { // 遇到错误时回滚事务
			color.Error.Println(err.Error())
			return err
		}
		color.Info.Println("\n[Mysql] --> t_country_phone_code 表的数据创建成功!")
		return nil
	})
}
