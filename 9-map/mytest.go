package main

import "fmt"

type Expose struct {
	IsExpose bool  `json:"is_expose" bson:"is_expose"`
	Ut       int64 `json:"ut" bson:"ut"`
}

type PlaylistExposeRecord struct {
	Mid int64 `json:"mid" bson:"mid"`
	Fid int64 `json:"fid" bson:"fid"`
	Ut  int64 `json:"ut" bson:"ut"`
}

func test() (history map[int64]*Expose) {
	history = make(map[int64]*Expose)
	var list []*PlaylistExposeRecord
	list = append(list, &PlaylistExposeRecord{Mid: 123, Fid: 123, Ut: 123})
	//list = append(list, &PlaylistExposeRecord{Mid: 1234, Ut: 1234444})
	for _, r := range list {
		history[r.Fid] = &Expose{Ut: r.Ut}
	}
	return
}

func test2() (ll map[int64]bool) {
	ll = make(map[int64]bool)
	//ll[123] = "hello"
	//ll[1234] = "world"
	//ll[12345] = "wwwww"

	return
}

type MediaIdFolderData struct {
	Folder    *MaterialFolderDo `json:"folder"`
	Type      int               `json:"type"`       // 1: 发帖上传关联播单 2: 添加关联播单 3:白名单播单
	FlagCount int               `json:"flag_count"` // 锦旗数量
	Order     int               `json:"order"`      // 当前视频在播单中排序
}
type MaterialFolderDo struct {
	Id  int64 `json:"id" bson:"id"`
	Mid int64 `json:"mid" bson:"mid"`
}

type GetItemsByMediaIdResp struct {
	Data GetItemsByMediaIdData `json:"data"`
}
type GetItemsByMediaIdData struct {
	MediaIdFolderMap map[int64]*MediaIdFolderData `json:"mediaId_folder_map"`
}

func main() {
	//var data GetItemsByMediaIdData
	var mediaIdFolderMap = make(map[int64]*MediaIdFolderData)
	if data, ok := mediaIdFolderMap[111]; ok {
		folder := data.Folder
		fmt.Println(folder)
	}
}

type MaterialGetIFloderParam struct {
	Pid      int64   `json:"pid"`
	MediaIds []int64 `json:"media_ids"`
}
