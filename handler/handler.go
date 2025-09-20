package handler

import (
	"bileygr/components"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/darenliang/jikan-go"
	"github.com/labstack/echo"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

type MInfo struct {
	title  string
	mType  string
	image  string
	status string
}

func getStatus(status string) string {
	switch status {
	case "On Hiatus":
		return "Hiatus"
	case "Publishing":
		return "Ongoing"
	case "Finished":
		return "Completed"
	default:
		return "Undefined"
	}
}

func getMInfo(mangas *jikan.TopManga) []MInfo {
	var mangasStruct []MInfo
	for mangaIndex := range mangas.Data {
		mangasStruct = append(mangasStruct, MInfo{
			title:  mangas.Data[mangaIndex].TitleEnglish,
			mType:  mangas.Data[mangaIndex].Type,
			image:  mangas.Data[mangaIndex].Images.Jpg.LargeImageUrl,
			status: getStatus(mangas.Data[mangaIndex].Status),
		})
	}
	return mangasStruct
}

func pPrintMInfo(mInfo []MInfo) {
	for idx := range mInfo {
		fmt.Println("M Name:", mInfo[idx].title)
		fmt.Println("M Type", mInfo[idx].mType)
		fmt.Println("M Image:", mInfo[idx].image)
		fmt.Println("M status:", mInfo[idx].status)
	}
}

func HomeHandler(c echo.Context) error {
	mangas, errManga := jikan.GetTopManga(jikan.TopMangaTypeManga, jikan.TopMangaFilterByPopularity, 1)
	if errManga != nil {
		c.Logger().Warnf("failed to fetch Manga: %v", errManga)
		mangas = &jikan.TopManga{}
	}

	if len(mangas.Data) > 0 {
		mangaInfo := getMInfo(mangas)
		pPrintMInfo(mangaInfo)
	} else {
		fmt.Println("No manga data received")
	}

	manhwas, errManhwa := jikan.GetTopManga(jikan.TopMangaTypeManhwa, jikan.TopMangaFilterByPopularity, 1)
	if errManhwa != nil {
		c.Logger().Warnf("failed to fetch Manga: %v", errManhwa)
		manhwas = &jikan.TopManga{}
	}
	if len(manhwas.Data) > 0 {
		manhwasInfo := getMInfo(manhwas)
		pPrintMInfo(manhwasInfo)
	} else {
		fmt.Println("No manga data received")
	}

	manhuas, errManhua := jikan.GetTopManga(jikan.TopMangaTypeManhua, jikan.TopMangaFilterByPopularity, 1)
	if errManhua != nil {
		c.Logger().Warnf("failed to fetch Manga: %v", errManhua)
		manhuas = &jikan.TopManga{}
	}
	if len(manhuas.Data) > 0 {
		manhuasInfo := getMInfo(manhuas)
		pPrintMInfo(manhuasInfo)
	} else {
		fmt.Println("No manga data received")
	}

	return Render(c, http.StatusOK, components.Home())
}
