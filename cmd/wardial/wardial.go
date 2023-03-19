package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	_, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	areaCode, err := getAreaCode(os.Args)
	io.FatalIfError(err)

	// find all numbers in given area:
	phoneNumbers := network.FindPhoneNumbersWithPrefix(ctx, "+1 "+areaCode+"-")

	// filter out numbers with extensions:
	phoneNumbers = slices.Filter(phoneNumbers,
		func(x schema.PhoneNumber) bool {
			return !strings.Contains(string(x), "ext.")
		})

	targetNumber, err := slices.Choose(phoneNumbers)
	if err != nil {
		// create a fake number
		st := "+1 " + areaCode + "-"
		st += fmt.Sprintf("%d", rand.Intn(9)+1)

		for i := 0; i < 6; i++ {
			st += fmt.Sprintf("%d", rand.Intn(10))
		}

		targetNumber, err = schema.ParsePhoneNumber(st)
		io.FatalIfError(err)
	}

	// decrement the number a couple of times
	prob := 1.0
	for rand.Float64() < prob {
		targetNumber, _ = targetNumber.Prev()
		prob -= 0.1
	}

	for i := 0; i < 10; i++ {
		ok, err := altnet.Dial(ctx, targetNumber, network)
		io.FatalIfError(err)
		if ok {
			break
		}
		targetNumber, ok = targetNumber.Next()
		if !ok {
			break
		}
	}
}

func getAreaCode(args []string) (string, error) {

	help := ""
	help += "Alberta: 403  Alaska: 907  Alabama: 205,334,256  Arkansas: 501\n"
	help += "Arizona: 602,520  British Columbia: 604,250\n"
	help += "California: 415,408,619,714,805,818,916,310,213,510,209,707,562,323,858,909\n"
	help += "Colorado: 303,719  Connecticut: 203  District of Columbia: 202\n"
	help += "Delaware: 302  Dominican Republic: 809  Florida: 407,305,904,813,321,386\n"
	help += "Georgia: 404,912,706,478  Guam: 671  Hawaii: 808  Iowa: 319,515,712,641\n"
	help += "Idaho: 208  Illinois: 312,708,309,217,815,618  Indiana: 317,219,812\n"
	help += "Kansas: 913,316  Kentucky: 502,606  Louisiana: 504,318,337\n"
	help += "Massachusetts: 617,508,413,339,351  Midway Islands: 301,410  Maine: 207\n"
	help += "Michigan: 313,616,517,906,248  Minnesota: 612,218,507  Missouri: 314,816,417\n"
	help += "Mississippi: 601  Montana: 406  North Carolina: 919,704,336\n"
	help += "North Dakota: 701  Nebraska: 402,308  New Hampshire: 603\n"
	help += "New Jersey: 201,609,908,551,732  New Mexico: 505  Nova Scotia: 902\n"
	help += "Nevada: 702  New York: 212,914,716,516,315,718,518,607\n"
	help += "Ohio: 216,513,614,419  Oklahoma: 405,918,580  Ontario: 416,519,613,705,807\n"
	help += "Oregon: 503,541  Pennsylvania: 215,717,412,814,610  Rhode Island: 401\n"
	help += "South Carolina: 803  South Dakota: 605  Saskatchewan: 306\n"
	help += "Tennessee: 615,901,931\n"
	help += "Texas: 512,713,214,817,915,409,806,903,210,972,325,432,469  Utah: 801\n"
	help += "Virginia: 703,804,540,571  Vermont: 802  Washington: 206,509,360\n"
	help += "Wisconsin: 414,608,715  West Virginia: 304  Wyoming: 307\n"

	r := regexp.MustCompile(`(\d+)`)
	matches := r.FindAllStringSubmatch(help, -1)
	areaCodes := slices.Map(matches, func(item []string) string { return item[1] })

	if len(args) > 1 && slices.Contains(areaCodes, args[1]) {
		return args[1], nil
	}

	for {
		res, err := io.ReadNotEmpty[string]("enter area code (? for list): ")
		if err != nil {
			return "", err
		}

		if res == "?" {
			fmt.Println(help)
		} else if slices.Contains(areaCodes, res) {
			return res, nil
		}
	}
}
