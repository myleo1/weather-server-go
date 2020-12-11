package weatherdao

import (
	"github.com/mizuki1412/go-core-kit/library/httpkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"net/http"
)

var jsonCity = []byte(`[
    {
      "cities": [
        {
          "code": "1101",
          "name": "市辖区",
          "province": {
            "cities": null,
            "code": "11",
            "name": null
          }
        }
      ],
      "code": "11",
      "name": "北京市"
    },
    {
      "cities": [
        {
          "code": "1201",
          "name": "市辖区",
          "province": {
            "cities": null,
            "code": "12",
            "name": null
          }
        }
      ],
      "code": "12",
      "name": "天津市"
    },
    {
      "cities": [
        {
          "code": "1301",
          "name": "石家庄市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1302",
          "name": "唐山市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1303",
          "name": "秦皇岛市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1304",
          "name": "邯郸市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1305",
          "name": "邢台市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1306",
          "name": "保定市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1307",
          "name": "张家口市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1308",
          "name": "承德市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1309",
          "name": "沧州市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1310",
          "name": "廊坊市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        },
        {
          "code": "1311",
          "name": "衡水市",
          "province": {
            "cities": null,
            "code": "13",
            "name": null
          }
        }
      ],
      "code": "13",
      "name": "河北省"
    },
    {
      "cities": [
        {
          "code": "1401",
          "name": "太原市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1402",
          "name": "大同市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1403",
          "name": "阳泉市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1404",
          "name": "长治市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1405",
          "name": "晋城市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1406",
          "name": "朔州市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1407",
          "name": "晋中市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1408",
          "name": "运城市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1409",
          "name": "忻州市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1410",
          "name": "临汾市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        },
        {
          "code": "1411",
          "name": "吕梁市",
          "province": {
            "cities": null,
            "code": "14",
            "name": null
          }
        }
      ],
      "code": "14",
      "name": "山西省"
    },
    {
      "cities": [
        {
          "code": "1501",
          "name": "呼和浩特市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1502",
          "name": "包头市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1503",
          "name": "乌海市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1504",
          "name": "赤峰市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1505",
          "name": "通辽市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1506",
          "name": "鄂尔多斯市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1507",
          "name": "呼伦贝尔市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1508",
          "name": "巴彦淖尔市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1509",
          "name": "乌兰察布市",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1522",
          "name": "兴安盟",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1525",
          "name": "锡林郭勒盟",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        },
        {
          "code": "1529",
          "name": "阿拉善盟",
          "province": {
            "cities": null,
            "code": "15",
            "name": null
          }
        }
      ],
      "code": "15",
      "name": "内蒙古自治区"
    },
    {
      "cities": [
        {
          "code": "2101",
          "name": "沈阳市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2102",
          "name": "大连市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2103",
          "name": "鞍山市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2104",
          "name": "抚顺市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2105",
          "name": "本溪市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2106",
          "name": "丹东市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2107",
          "name": "锦州市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2108",
          "name": "营口市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2109",
          "name": "阜新市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2110",
          "name": "辽阳市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2111",
          "name": "盘锦市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2112",
          "name": "铁岭市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2113",
          "name": "朝阳市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        },
        {
          "code": "2114",
          "name": "葫芦岛市",
          "province": {
            "cities": null,
            "code": "21",
            "name": null
          }
        }
      ],
      "code": "21",
      "name": "辽宁省"
    },
    {
      "cities": [
        {
          "code": "2201",
          "name": "长春市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2202",
          "name": "吉林市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2203",
          "name": "四平市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2204",
          "name": "辽源市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2205",
          "name": "通化市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2206",
          "name": "白山市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2207",
          "name": "松原市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2208",
          "name": "白城市",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        },
        {
          "code": "2224",
          "name": "延边朝鲜族自治州",
          "province": {
            "cities": null,
            "code": "22",
            "name": null
          }
        }
      ],
      "code": "22",
      "name": "吉林省"
    },
    {
      "cities": [
        {
          "code": "2301",
          "name": "哈尔滨市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2302",
          "name": "齐齐哈尔市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2303",
          "name": "鸡西市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2304",
          "name": "鹤岗市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2305",
          "name": "双鸭山市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2306",
          "name": "大庆市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2307",
          "name": "伊春市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2308",
          "name": "佳木斯市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2309",
          "name": "七台河市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2310",
          "name": "牡丹江市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2311",
          "name": "黑河市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2312",
          "name": "绥化市",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        },
        {
          "code": "2327",
          "name": "大兴安岭地区",
          "province": {
            "cities": null,
            "code": "23",
            "name": null
          }
        }
      ],
      "code": "23",
      "name": "黑龙江省"
    },
    {
      "cities": [
        {
          "code": "3101",
          "name": "市辖区",
          "province": {
            "cities": null,
            "code": "31",
            "name": null
          }
        }
      ],
      "code": "31",
      "name": "上海市"
    },
    {
      "cities": [
        {
          "code": "3201",
          "name": "南京市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3202",
          "name": "无锡市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3203",
          "name": "徐州市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3204",
          "name": "常州市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3205",
          "name": "苏州市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3206",
          "name": "南通市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3207",
          "name": "连云港市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3208",
          "name": "淮安市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3209",
          "name": "盐城市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3210",
          "name": "扬州市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3211",
          "name": "镇江市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3212",
          "name": "泰州市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        },
        {
          "code": "3213",
          "name": "宿迁市",
          "province": {
            "cities": null,
            "code": "32",
            "name": null
          }
        }
      ],
      "code": "32",
      "name": "江苏省"
    },
    {
      "cities": [
        {
          "code": "3301",
          "name": "杭州市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3302",
          "name": "宁波市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3303",
          "name": "温州市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3304",
          "name": "嘉兴市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3305",
          "name": "湖州市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3306",
          "name": "绍兴市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3307",
          "name": "金华市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3308",
          "name": "衢州市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3309",
          "name": "舟山市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3310",
          "name": "台州市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        },
        {
          "code": "3311",
          "name": "丽水市",
          "province": {
            "cities": null,
            "code": "33",
            "name": null
          }
        }
      ],
      "code": "33",
      "name": "浙江省"
    },
    {
      "cities": [
        {
          "code": "3401",
          "name": "合肥市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3402",
          "name": "芜湖市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3403",
          "name": "蚌埠市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3404",
          "name": "淮南市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3405",
          "name": "马鞍山市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3406",
          "name": "淮北市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3407",
          "name": "铜陵市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3408",
          "name": "安庆市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3410",
          "name": "黄山市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3411",
          "name": "滁州市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3412",
          "name": "阜阳市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3413",
          "name": "宿州市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3415",
          "name": "六安市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3416",
          "name": "亳州市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3417",
          "name": "池州市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        },
        {
          "code": "3418",
          "name": "宣城市",
          "province": {
            "cities": null,
            "code": "34",
            "name": null
          }
        }
      ],
      "code": "34",
      "name": "安徽省"
    },
    {
      "cities": [
        {
          "code": "3501",
          "name": "福州市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3502",
          "name": "厦门市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3503",
          "name": "莆田市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3504",
          "name": "三明市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3505",
          "name": "泉州市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3506",
          "name": "漳州市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3507",
          "name": "南平市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3508",
          "name": "龙岩市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        },
        {
          "code": "3509",
          "name": "宁德市",
          "province": {
            "cities": null,
            "code": "35",
            "name": null
          }
        }
      ],
      "code": "35",
      "name": "福建省"
    },
    {
      "cities": [
        {
          "code": "3601",
          "name": "南昌市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3602",
          "name": "景德镇市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3603",
          "name": "萍乡市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3604",
          "name": "九江市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3605",
          "name": "新余市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3606",
          "name": "鹰潭市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3607",
          "name": "赣州市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3608",
          "name": "吉安市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3609",
          "name": "宜春市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3610",
          "name": "抚州市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        },
        {
          "code": "3611",
          "name": "上饶市",
          "province": {
            "cities": null,
            "code": "36",
            "name": null
          }
        }
      ],
      "code": "36",
      "name": "江西省"
    },
    {
      "cities": [
        {
          "code": "3701",
          "name": "济南市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3702",
          "name": "青岛市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3703",
          "name": "淄博市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3704",
          "name": "枣庄市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3705",
          "name": "东营市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3706",
          "name": "烟台市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3707",
          "name": "潍坊市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3708",
          "name": "济宁市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3709",
          "name": "泰安市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3710",
          "name": "威海市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3711",
          "name": "日照市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3712",
          "name": "莱芜市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3713",
          "name": "临沂市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3714",
          "name": "德州市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3715",
          "name": "聊城市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3716",
          "name": "滨州市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        },
        {
          "code": "3717",
          "name": "菏泽市",
          "province": {
            "cities": null,
            "code": "37",
            "name": null
          }
        }
      ],
      "code": "37",
      "name": "山东省"
    },
    {
      "cities": [
        {
          "code": "4101",
          "name": "郑州市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4102",
          "name": "开封市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4103",
          "name": "洛阳市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4104",
          "name": "平顶山市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4105",
          "name": "安阳市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4106",
          "name": "鹤壁市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4107",
          "name": "新乡市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4108",
          "name": "焦作市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4109",
          "name": "濮阳市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4110",
          "name": "许昌市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4111",
          "name": "漯河市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4112",
          "name": "三门峡市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4113",
          "name": "南阳市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4114",
          "name": "商丘市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4115",
          "name": "信阳市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4116",
          "name": "周口市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4117",
          "name": "驻马店市",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        },
        {
          "code": "4190",
          "name": "省直辖县级行政区划",
          "province": {
            "cities": null,
            "code": "41",
            "name": null
          }
        }
      ],
      "code": "41",
      "name": "河南省"
    },
    {
      "cities": [
        {
          "code": "4201",
          "name": "武汉市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4202",
          "name": "黄石市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4203",
          "name": "十堰市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4205",
          "name": "宜昌市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4206",
          "name": "襄阳市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4207",
          "name": "鄂州市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4208",
          "name": "荆门市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4209",
          "name": "孝感市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4210",
          "name": "荆州市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4211",
          "name": "黄冈市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4212",
          "name": "咸宁市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4213",
          "name": "随州市",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4228",
          "name": "恩施土家族苗族自治州",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        },
        {
          "code": "4290",
          "name": "省直辖县级行政区划",
          "province": {
            "cities": null,
            "code": "42",
            "name": null
          }
        }
      ],
      "code": "42",
      "name": "湖北省"
    },
    {
      "cities": [
        {
          "code": "4301",
          "name": "长沙市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4302",
          "name": "株洲市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4303",
          "name": "湘潭市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4304",
          "name": "衡阳市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4305",
          "name": "邵阳市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4306",
          "name": "岳阳市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4307",
          "name": "常德市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4308",
          "name": "张家界市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4309",
          "name": "益阳市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4310",
          "name": "郴州市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4311",
          "name": "永州市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4312",
          "name": "怀化市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4313",
          "name": "娄底市",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        },
        {
          "code": "4331",
          "name": "湘西土家族苗族自治州",
          "province": {
            "cities": null,
            "code": "43",
            "name": null
          }
        }
      ],
      "code": "43",
      "name": "湖南省"
    },
    {
      "cities": [
        {
          "code": "4401",
          "name": "广州市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4402",
          "name": "韶关市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4403",
          "name": "深圳市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4404",
          "name": "珠海市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4405",
          "name": "汕头市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4406",
          "name": "佛山市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4407",
          "name": "江门市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4408",
          "name": "湛江市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4409",
          "name": "茂名市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4412",
          "name": "肇庆市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4413",
          "name": "惠州市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4414",
          "name": "梅州市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4415",
          "name": "汕尾市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4416",
          "name": "河源市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4417",
          "name": "阳江市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4418",
          "name": "清远市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4419",
          "name": "东莞市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4420",
          "name": "中山市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4451",
          "name": "潮州市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4452",
          "name": "揭阳市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        },
        {
          "code": "4453",
          "name": "云浮市",
          "province": {
            "cities": null,
            "code": "44",
            "name": null
          }
        }
      ],
      "code": "44",
      "name": "广东省"
    },
    {
      "cities": [
        {
          "code": "4501",
          "name": "南宁市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4502",
          "name": "柳州市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4503",
          "name": "桂林市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4504",
          "name": "梧州市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4505",
          "name": "北海市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4506",
          "name": "防城港市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4507",
          "name": "钦州市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4508",
          "name": "贵港市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4509",
          "name": "玉林市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4510",
          "name": "百色市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4511",
          "name": "贺州市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4512",
          "name": "河池市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4513",
          "name": "来宾市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        },
        {
          "code": "4514",
          "name": "崇左市",
          "province": {
            "cities": null,
            "code": "45",
            "name": null
          }
        }
      ],
      "code": "45",
      "name": "广西壮族自治区"
    },
    {
      "cities": [
        {
          "code": "4601",
          "name": "海口市",
          "province": {
            "cities": null,
            "code": "46",
            "name": null
          }
        },
        {
          "code": "4602",
          "name": "三亚市",
          "province": {
            "cities": null,
            "code": "46",
            "name": null
          }
        },
        {
          "code": "4603",
          "name": "三沙市",
          "province": {
            "cities": null,
            "code": "46",
            "name": null
          }
        },
        {
          "code": "4604",
          "name": "儋州市",
          "province": {
            "cities": null,
            "code": "46",
            "name": null
          }
        },
        {
          "code": "4690",
          "name": "省直辖县级行政区划",
          "province": {
            "cities": null,
            "code": "46",
            "name": null
          }
        }
      ],
      "code": "46",
      "name": "海南省"
    },
    {
      "cities": [
        {
          "code": "5001",
          "name": "市辖区",
          "province": {
            "cities": null,
            "code": "50",
            "name": null
          }
        },
        {
          "code": "5002",
          "name": "县",
          "province": {
            "cities": null,
            "code": "50",
            "name": null
          }
        }
      ],
      "code": "50",
      "name": "重庆市"
    },
    {
      "cities": [
        {
          "code": "5101",
          "name": "成都市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5103",
          "name": "自贡市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5104",
          "name": "攀枝花市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5105",
          "name": "泸州市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5106",
          "name": "德阳市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5107",
          "name": "绵阳市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5108",
          "name": "广元市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5109",
          "name": "遂宁市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5110",
          "name": "内江市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5111",
          "name": "乐山市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5113",
          "name": "南充市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5114",
          "name": "眉山市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5115",
          "name": "宜宾市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5116",
          "name": "广安市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5117",
          "name": "达州市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5118",
          "name": "雅安市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5119",
          "name": "巴中市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5120",
          "name": "资阳市",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5132",
          "name": "阿坝藏族羌族自治州",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5133",
          "name": "甘孜藏族自治州",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        },
        {
          "code": "5134",
          "name": "凉山彝族自治州",
          "province": {
            "cities": null,
            "code": "51",
            "name": null
          }
        }
      ],
      "code": "51",
      "name": "四川省"
    },
    {
      "cities": [
        {
          "code": "5201",
          "name": "贵阳市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5202",
          "name": "六盘水市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5203",
          "name": "遵义市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5204",
          "name": "安顺市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5205",
          "name": "毕节市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5206",
          "name": "铜仁市",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5223",
          "name": "黔西南布依族苗族自治州",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5226",
          "name": "黔东南苗族侗族自治州",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        },
        {
          "code": "5227",
          "name": "黔南布依族苗族自治州",
          "province": {
            "cities": null,
            "code": "52",
            "name": null
          }
        }
      ],
      "code": "52",
      "name": "贵州省"
    },
    {
      "cities": [
        {
          "code": "5301",
          "name": "昆明市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5303",
          "name": "曲靖市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5304",
          "name": "玉溪市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5305",
          "name": "保山市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5306",
          "name": "昭通市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5307",
          "name": "丽江市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5308",
          "name": "普洱市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5309",
          "name": "临沧市",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5323",
          "name": "楚雄彝族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5325",
          "name": "红河哈尼族彝族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5326",
          "name": "文山壮族苗族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5328",
          "name": "西双版纳傣族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5329",
          "name": "大理白族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5331",
          "name": "德宏傣族景颇族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5333",
          "name": "怒江傈僳族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        },
        {
          "code": "5334",
          "name": "迪庆藏族自治州",
          "province": {
            "cities": null,
            "code": "53",
            "name": null
          }
        }
      ],
      "code": "53",
      "name": "云南省"
    },
    {
      "cities": [
        {
          "code": "5401",
          "name": "拉萨市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5402",
          "name": "日喀则市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5403",
          "name": "昌都市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5404",
          "name": "林芝市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5405",
          "name": "山南市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5406",
          "name": "那曲市",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        },
        {
          "code": "5425",
          "name": "阿里地区",
          "province": {
            "cities": null,
            "code": "54",
            "name": null
          }
        }
      ],
      "code": "54",
      "name": "西藏自治区"
    },
    {
      "cities": [
        {
          "code": "6101",
          "name": "西安市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6102",
          "name": "铜川市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6103",
          "name": "宝鸡市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6104",
          "name": "咸阳市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6105",
          "name": "渭南市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6106",
          "name": "延安市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6107",
          "name": "汉中市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6108",
          "name": "榆林市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6109",
          "name": "安康市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        },
        {
          "code": "6110",
          "name": "商洛市",
          "province": {
            "cities": null,
            "code": "61",
            "name": null
          }
        }
      ],
      "code": "61",
      "name": "陕西省"
    },
    {
      "cities": [
        {
          "code": "6201",
          "name": "兰州市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6202",
          "name": "嘉峪关市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6203",
          "name": "金昌市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6204",
          "name": "白银市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6205",
          "name": "天水市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6206",
          "name": "武威市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6207",
          "name": "张掖市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6208",
          "name": "平凉市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6209",
          "name": "酒泉市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6210",
          "name": "庆阳市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6211",
          "name": "定西市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6212",
          "name": "陇南市",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6229",
          "name": "临夏回族自治州",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        },
        {
          "code": "6230",
          "name": "甘南藏族自治州",
          "province": {
            "cities": null,
            "code": "62",
            "name": null
          }
        }
      ],
      "code": "62",
      "name": "甘肃省"
    },
    {
      "cities": [
        {
          "code": "6301",
          "name": "西宁市",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6302",
          "name": "海东市",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6322",
          "name": "海北藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6323",
          "name": "黄南藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6325",
          "name": "海南藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6326",
          "name": "果洛藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6327",
          "name": "玉树藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        },
        {
          "code": "6328",
          "name": "海西蒙古族藏族自治州",
          "province": {
            "cities": null,
            "code": "63",
            "name": null
          }
        }
      ],
      "code": "63",
      "name": "青海省"
    },
    {
      "cities": [
        {
          "code": "6401",
          "name": "银川市",
          "province": {
            "cities": null,
            "code": "64",
            "name": null
          }
        },
        {
          "code": "6402",
          "name": "石嘴山市",
          "province": {
            "cities": null,
            "code": "64",
            "name": null
          }
        },
        {
          "code": "6403",
          "name": "吴忠市",
          "province": {
            "cities": null,
            "code": "64",
            "name": null
          }
        },
        {
          "code": "6404",
          "name": "固原市",
          "province": {
            "cities": null,
            "code": "64",
            "name": null
          }
        },
        {
          "code": "6405",
          "name": "中卫市",
          "province": {
            "cities": null,
            "code": "64",
            "name": null
          }
        }
      ],
      "code": "64",
      "name": "宁夏回族自治区"
    },
    {
      "cities": [
        {
          "code": "6501",
          "name": "乌鲁木齐市",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6502",
          "name": "克拉玛依市",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6504",
          "name": "吐鲁番市",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6505",
          "name": "哈密市",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6523",
          "name": "昌吉回族自治州",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6527",
          "name": "博尔塔拉蒙古自治州",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6528",
          "name": "巴音郭楞蒙古自治州",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6529",
          "name": "阿克苏地区",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6530",
          "name": "克孜勒苏柯尔克孜自治州",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6531",
          "name": "喀什地区",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6532",
          "name": "和田地区",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6540",
          "name": "伊犁哈萨克自治州",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6542",
          "name": "塔城地区",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6543",
          "name": "阿勒泰地区",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        },
        {
          "code": "6590",
          "name": "自治区直辖县级行政区划",
          "province": {
            "cities": null,
            "code": "65",
            "name": null
          }
        }
      ],
      "code": "65",
      "name": "新疆维吾尔自治区"
    }
  ],`)

func CityRet() []map[string]interface{} {
	var rs = gjson.ParseBytes(jsonCity).Array()
	data := make([]map[string]interface{}, 0, len(rs))
	for _, v := range rs {
		tmp := v.Map()
		m := map[string]interface{}{}
		for k, v := range tmp {
			m[k] = v.Value()
		}
		data = append(data, m)
	}
	return data
}

func GetWeather(location string) interface{} {
	key := configkit.GetStringD("amap.key")
	ret1, _ := httpkit.Request(httpkit.Req{
		Method: http.MethodGet,
		Url:    "https://restapi.amap.com/v3/weather/weatherInfo?key=" + key + "&city=" + location + "&extensions=base",
	})
	rs1 := gjson.Parse(ret1).Map()
	var ret interface{}
	for k, v := range rs1 {
		if k == "lives" {
			s := cast.ToSlice(v.Value())
			m := cast.ToStringMap(s[0])
			weather := m["weather"]
			switch weather {
			case "阵雨", "雷阵雨", "雷阵雨并伴有冰雹", "阵雨夹雪", "冻雨", "雪", "阵雪", "小雪-中雪", "雨雪天气", "雨夹雪":
				m["rain"] = "30"
			case "小雨", "中雨", "大雨", "毛毛雨/细雨", "雨", "小雨-中雨", "中雨-大雨", "小雪", "中雪", "中雪-大雪":
				m["rain"] = "80"
			case "暴雨", "大暴雨", "特大暴雨", "强阵雨", "强雷阵雨", "大雨-暴雨", "暴雨-大暴雨", "大雪", "暴雪", "大雪-暴雪":
				m["rain"] = "90"
			case "极端降雨", "大暴雨-特大暴雨":
				m["rain"] = "100"
			default:
				m["rain"] = "0"
			}
			ret = m
		}
	}
	return ret
}

func GetWeatherCast(location string) interface{} {
	key := configkit.GetStringD("amap.key")
	ret2, _ := httpkit.Request(httpkit.Req{
		Method: http.MethodGet,
		Url:    "https://restapi.amap.com/v3/weather/weatherInfo?key=" + key + "&city=" + location + "&extensions=all",
	})
	rs2 := gjson.Parse(ret2).Get("forecasts").Array()[0].Map()
	ret := make([]map[string]interface{}, 0, 5)
	for k, v := range rs2 {
		if k == "casts" {
			vv := cast.ToSlice(v.Value())
			for _, vvv := range vv {
				m := cast.ToStringMap(vvv)
				weather := m["dayweather"]
				switch weather {
				case "阵雨", "雷阵雨", "雷阵雨并伴有冰雹", "阵雨夹雪", "冻雨", "雪", "阵雪", "小雪-中雪", "雨雪天气", "雨夹雪":
					m["rain"] = "30"
				case "小雨", "中雨", "大雨", "毛毛雨/细雨", "雨", "小雨-中雨", "中雨-大雨", "小雪", "中雪", "中雪-大雪":
					m["rain"] = "80"
				case "暴雨", "大暴雨", "特大暴雨", "强阵雨", "强雷阵雨", "大雨-暴雨", "暴雨-大暴雨", "大雪", "暴雪", "大雪-暴雪":
					m["rain"] = "90"
				case "极端降雨", "大暴雨-特大暴雨":
					m["rain"] = "100"
				default:
					m["rain"] = "0"
				}
				ret = append(ret, m)
			}
		}
	}
	return ret
}
