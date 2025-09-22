package handler

import (
	"bileygr/components"
	"fmt"
	"net/http"

	"github.com/darenliang/jikan-go"
	"github.com/labstack/echo"
)

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

func getMInfo(mangas *jikan.TopManga) []components.MInfo {
	var mangasStruct []components.MInfo
	for mangaIndex := range mangas.Data {
		mangasStruct = append(mangasStruct, components.MInfo{
			Title:  mangas.Data[mangaIndex].TitleEnglish,
			MType:  mangas.Data[mangaIndex].Type,
			Image:  mangas.Data[mangaIndex].Images.Jpg.LargeImageUrl,
			Status: getStatus(mangas.Data[mangaIndex].Status),
		})
	}
	return mangasStruct
}

func pPrintMInfo(mInfo []components.MInfo) {
	for idx := range mInfo {
		fmt.Println("M Name:", mInfo[idx].Title)
		fmt.Println("M Type", mInfo[idx].MType)
		fmt.Println("M Image:", mInfo[idx].Image)
		fmt.Println("M status:", mInfo[idx].Status)
		fmt.Println()
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

	return Render(c, http.StatusOK, components.Home(getMInfo(mangas), getMInfo(manhwas), getMInfo(manhuas)))
}
