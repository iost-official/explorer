package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iost-official/explorer/backend/util/common"
	"net/http"
	"strconv"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
)

func SetProducerContributions(c echo.Context) (err error) {
	aInfo := new(db.ProducerLevelInfo)

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	s := buf.String() // Does a complete copy of the bytes in the buffer.
	fmt.Println(s)

	json.Unmarshal([]byte(s), &aInfo)
	fmt.Printf("haha %+v", aInfo)

	if err = db.SaveProducerLevelInfo(*aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	err = common.CalculateProducerContributions(*aInfo)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, FormatResponse(aInfo.Aid))
}

func SetVoteAwardInfo(c echo.Context) (err error) {
	aInfo := new(db.AwardInfo)
	if err = c.Bind(aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	aInfo.Aid = strconv.FormatInt(aInfo.StartTime, 10) + strconv.FormatInt(aInfo.EndTime, 10) + strconv.FormatInt(aInfo.TotalAmount, 10)

	if err = db.SaveAwardInfo(*aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	return CalculateAward(c, *aInfo)
}

func GetUserContributionAward(c echo.Context) error {
	id := c.QueryParam("aid")
	format := c.QueryParam("format")

	userAward, err := db.GetUserContributionAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	var userAwardsOutput []interface{}
	for _, ua := range userAward {
		userAwardOutput := []interface{}{ua.Username, ua.Pid, ua.Vote, ua.Award}
		userAwardsOutput = append(userAwardsOutput, userAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(userAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(userAward))
	}
}

func GetProducerContributionAward(c echo.Context) (err error) {
	id := c.QueryParam("aid")
	format := c.QueryParam("format")
	producerAward, err := db.GetProducerContributionAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	var producerAwardsOutput []interface{}
	for _, pa := range producerAward {
		producerAwardOutput := []interface{}{pa.Pid, pa.Vote, pa.Award}
		producerAwardsOutput = append(producerAwardsOutput, producerAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(producerAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(producerAward))
	}
}

func GetUserAward(c echo.Context) error {
	id := c.QueryParam("aid")
	format := c.QueryParam("format")

	userAward, err := db.GetUserAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	var userAwardsOutput []interface{}
	for _, ua := range userAward {
		userAwardOutput := []interface{}{ua.Username, ua.Pid, ua.Vote, ua.Award}
		userAwardsOutput = append(userAwardsOutput, userAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(userAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(userAward))
	}
}

func GetProducerAward(c echo.Context) (err error) {
	id := c.QueryParam("aid")
	format := c.QueryParam("format")
	producerAward, err := db.GetProducerAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	var producerAwardsOutput []interface{}
	for _, pa := range producerAward {
		producerAwardOutput := []interface{}{pa.Pid, pa.Vote, pa.Award}
		producerAwardsOutput = append(producerAwardsOutput, producerAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(producerAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(producerAward))
	}
}

func GetVoteAwardList(c echo.Context) (err error) {
	voteAwardList, err := db.GetVoteAwardList()
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	return c.JSON(http.StatusOK, FormatResponse(voteAwardList))
}

func GetVoteAwardInfo(c echo.Context) (err error) {
	id := c.QueryParam("aid")
	voteAward, err := db.GetVoteAwardInfo(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, FormatResponse(voteAward))
}
func CalculateAward(c echo.Context, ainfo db.AwardInfo) (err error) {
	err = common.CalculateAward(ainfo)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, FormatResponse(ainfo.Aid))
}


