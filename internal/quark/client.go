package quark

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	HTTPClient *resty.Client
}

func New() *Client {
	return &Client{
		HTTPClient: resty.New(),
	}
}

func (c *Client) Folder(fid string) {
	c.HTTPClient.R().
		SetCookie(&http.Cookie{
			Name:  "__pus",
			Value: "bff922df356e0c0e6b45fdf55d9100c2AATbYpehiUFKlX8CxTXeSpZ5KkhoczSBz6ooOfc6+G9byTZyGZ+ldrGO+SiOE3rZ/PH0PChaJGi1wvgxZryQNYut",
		})
}

func (c *Client) Direct(fids []string) (*DirectResponse, error) {
	apiUrl := "https://drive-pc.quark.cn/1/clouddrive/file/download?pr=ucpro&fr=pc"

	dr := DirectResponse{}
	payload := map[string]interface{}{
		"fids": fids,
	}

	resp, err := c.HTTPClient.R().
		SetBody(payload).
		SetCookie(&http.Cookie{
			Name:  "__pus",
			Value: "",
		}).SetResult(dr).Post(apiUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// 检查响应状态码
	if resp.IsError() {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status())
	}
	return resp.Result().(*DirectResponse), nil

}

type DirectResponse struct {
	Status    int      `json:"status"`
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Timestamp int64    `json:"timestamp"`
	Data      []File   `json:"data"`
	Metadata  Metadata `json:"metadata"`
}

type File struct {
	FID                 string      `json:"fid"`
	FileName            string      `json:"file_name"`
	PDirFID             string      `json:"pdir_fid"`
	Category            int         `json:"category"`
	FileType            int         `json:"file_type"`
	ObjKey              string      `json:"obj_key"`
	Size                int64       `json:"size"`
	FormatType          string      `json:"format_type"`
	Status              int         `json:"status"`
	Tags                string      `json:"tags"`
	OwnerUCID           string      `json:"owner_ucid"`
	LCreatedAt          int64       `json:"l_created_at"`
	LUpdatedAt          int64       `json:"l_updated_at"`
	Source              string      `json:"source"`
	FileSource          string      `json:"file_source"`
	NameSpace           int         `json:"name_space"`
	LShotAt             int64       `json:"l_shot_at"`
	DownloadURL         string      `json:"download_url"`
	MD5                 string      `json:"md5"`
	SourceDisplay       string      `json:"source_display"`
	SeriesDir           bool        `json:"series_dir"`
	UploadCameraRootDir bool        `json:"upload_camera_root_dir"`
	FPS                 float64     `json:"fps"`
	Like                int         `json:"like"`
	OperatedAt          int64       `json:"operated_at"`
	RiskType            int         `json:"risk_type"`
	RangeSize           int64       `json:"range_size"`
	BackupSign          int         `json:"backup_sign"`
	ObjCategory         string      `json:"obj_category"`
	FileNameHlStart     int         `json:"file_name_hl_start"`
	FileNameHlEnd       int         `json:"file_name_hl_end"`
	FileStruct          FileStruct  `json:"file_struct"`
	Duration            int         `json:"duration"`
	LastPlayInfo        PlayInfo    `json:"last_play_info"`
	EventExtra          interface{} `json:"event_extra"`
	ScrapeStatus        int         `json:"scrape_status"`
	UpdateViewAt        int64       `json:"update_view_at"`
	Ban                 bool        `json:"ban"`
	BackupSource        bool        `json:"backup_source"`
	OfflineSource       bool        `json:"offline_source"`
	OwnerDriveType      int         `json:"owner_drive_type_or_default"`
	SaveAsSource        bool        `json:"save_as_source"`
	CurVersion          int         `json:"cur_version_or_default"`
	RawNameSpace        int         `json:"raw_name_space"`
	Dir                 bool        `json:"dir"`
	File                bool        `json:"file"`
	CreatedAt           int64       `json:"created_at"`
	UpdatedAt           int64       `json:"updated_at"`
	Extra               interface{} `json:"_extra"`
}

type FileStruct struct {
	PlatformSource string `json:"platform_source"`
}

type PlayInfo struct {
	Time int `json:"time"`
}

type Metadata struct {
	Acc2 string `json:"acc2"`
	Acc1 string `json:"acc1"`
}
