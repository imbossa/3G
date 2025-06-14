package http

import (
	"encoding/json"
	"github.com/imbossa/3G/internal/cachekey"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imbossa/3G/internal/controller/http/example/request"
	"github.com/imbossa/3G/internal/entity"
)

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.TranslationHistory
// @Failure     500 {object} response.Error
// @Router      /translation/history [get]
func (r *Example) history(ctx *gin.Context) {
	// use cache
	cachedHistory, err := r.c.Client.Get(ctx.Request.Context(), cachekey.History).Result()
	if err == nil {
		var history entity.TranslationHistory
		if err := json.Unmarshal([]byte(cachedHistory), &history); err == nil {
			ctx.JSON(http.StatusOK, history)
			return
		}
		r.l.Error(err, "http - example - history - redis unmarshal error")
	}

	translationHistory, err := r.t.History(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - example - history")
		ctx.JSON(http.StatusInternalServerError, errorResponseGin("database problems"))
		return
	}

	// use cache
	cacheValue,_ := json.Marshal(translationHistory)
	if err := r.c.Client.Set(ctx.Request.Context(), cachekey.History, cacheValue, time.Duration(20)*time.Second).Err(); err != nil {
		r.l.Error(err, "http - example - history - redis set error")
	}

	ctx.JSON(http.StatusOK, translationHistory)
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body request.Translate true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /translation/do-translate [post]
func (r *Example) doTranslate(ctx *gin.Context) {
	var body request.Translate

	if err := ctx.ShouldBindJSON(&body); err != nil {
		r.l.Error(err, "http - example - doTranslate")
		ctx.JSON(http.StatusBadRequest, errorResponseGin("invalid request body"))
		return
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, "http - example - doTranslate")
		ctx.JSON(http.StatusBadRequest, errorResponseGin("invalid request body"))
		return
	}

	translation, err := r.t.Translate(
		ctx.Request.Context(),
		entity.Translation{
			Source:      body.Source,
			Destination: body.Destination,
			Original:    body.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - example - doTranslate")
		ctx.JSON(http.StatusInternalServerError, errorResponseGin("translation service problems"))
		return
	}

	ctx.JSON(http.StatusOK, translation)
}

func errorResponseGin(message string) gin.H {
	return gin.H{"error": message}
}
