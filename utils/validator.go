package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/utils/answer"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RegisterRule
//@description: 注册自定义规则方案建议在路由初始化层即注册
//@param: key string, rule Rules
//@return: err error

func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册,无法重复注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: NotEmpty
//@description: 非空 不能为其对应类型的0值
//@return: string

func NotEmpty() string {
	return "notEmpty"
}

//@author: [zooqkl](https://github.com/zooqkl)
//@function: RegexpMatch
//@description: 正则校验 校验输入项是否满足正则表达式
//@param:  rule string
//@return: string
func RegexpMatch(rule string) string {
	return "regexp=" + rule
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Lt
//@description: 小于入参(<) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Lt(mark string) string {
	return "lt=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Le
//@description: 小于等于入参(<=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Le(mark string) string {
	return "le=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Eq
//@description: 等于入参(==) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Eq(mark string) string {
	return "eq=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ne
//@description: 不等于入参(!=)  如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ne(mark string) string {
	return "ne=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ge
//@description: 大于等于入参(>=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ge(mark string) string {
	return "ge=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Gt
//@description: 大于入参(>) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Gt(mark string) string {
	return "gt=" + mark
}

// Verify 校验参数 校验失败就直接返回
func Verify(st any, roleMap Rules, c *gin.Context) {

	if err := c.ShouldBind(st); err != nil {
		panic(errors.Wrap(err, "请求参数结构体接收失败"))
	}

	if err := verify(reflect.ValueOf(st).Elem().Interface(), roleMap); err != nil {
		ZapErrorLog(err)
		answer.FailWithMessage(err.Error(), c)
	}
}

// verify 校验方法
func verify(st any, roleMap Rules) (err error) {
	if len(roleMap) == 0 {
		return nil
	}

	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	var structMap = map[string]reflect.Value{}

	// 深度遍历结构体的所有字段 构成map
	if err = depthStruckToMap(st, structMap); err != nil {
		return errors.New("expect struct")
	}

	for filedName, ruleList := range roleMap {
		if len(ruleList) <= 0 {
			continue
		}

		oneV, OK := structMap[filedName]
		if !OK {
			return errors.New(filedName + "字段不存在")
		}

		for _, v := range ruleList {
			switch {
			case v == "notEmpty":
				if isBlank(oneV) {
					return errors.New(filedName + "值不能为空")
				}
			case strings.Split(v, "=")[0] == "regexp":
				if !regexpMatch(strings.Split(v, "=")[1], oneV.String()) {
					return errors.New(filedName + "格式校验不通过")
				}
			case compareMap[strings.Split(v, "=")[0]]:
				if !compareVerify(oneV, v) {
					return errors.New(filedName + "长度或值不在合法范围," + v)
				}
			}
		}
	}
	return nil
}

// compareVerify 长度和数字的校验方法 根据类型自动校验
func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

// isBlank 非空校验
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// compare 比较函数
func compare(value any, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}

func depthStruckToMap(st any, structMap map[string]reflect.Value) error {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			if err := depthStruckToMap(val.Field(i).Interface(), structMap); err != nil {
				return err
			}
		} else {
			structMap[typ.Field(i).Name] = val.Field(i)
		}
	}

	return nil
}
