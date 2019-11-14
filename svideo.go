package management

//SVideo seed video define
type SVideo struct {
	model        `xorm:"extends" json:"-"`
	FindNo       string   `json:"-"`                              //查找号
	Bangumi      string   `xorm:"bangumi" json:"bangumi"`         //番組
	Intro        string   `xorm:"varchar(2048)" json:"intro"`     //简介
	Alias        []string `xorm:"json" json:"alias"`              //别名，片名
	ThumbHash    string   `xorm:"thumb_hash" json:"thumb_hash"`   //缩略图
	PosterHash   string   `xorm:"poster_hash" json:"poster_hash"` //海报地址
	SourceHash   string   `xorm:"source_hash" json:"source_hash"` //原片地址
	M3U8Hash     string   `xorm:"m3u8_hash" json:"m3u8_hash"`     //切片地址
	Key          string   `json:"-"`                              //秘钥
	M3U8         string   `xorm:"m3u8" json:"-"`                  //M3U8名
	Role         []string `xorm:"json" json:"role"`               //主演
	Director     string   `json:"-"`                              //导演
	Systematics  string   `json:"-"`                              //分级
	Season       string   `json:"-"`                              //季
	TotalEpisode string   `json:"-"`                              //总集数
	Episode      string   `json:"-"`                              //集数
	Producer     string   `json:"-"`                              //生产商
	Publisher    string   `json:"-"`                              //发行商
	Type         string   `json:"-"`                              //类型：film，FanDrama
	Format       string   `json:"format"`                         //输出格式：3D，2D,VR(VR格式：Half-SBS：左右半宽,Half-OU：上下半高,SBS：左右全宽)
	Language     string   `json:"-"`                              //语言
	Caption      string   `json:"-"`                              //字幕
	Group        string   `json:"-"`                              //分组
	Index        string   `json:"-"`                              //索引
	Date         string   `json:"-"`                              //发行日期
	Sharpness    string   `json:"sharpness"`                      //清晰度
	Visit        uint64   `json:"-" xorm:"notnull default(0)"`    //访问数
	Series       string   `json:"series"`                         //系列
	Tags         []string `xorm:"json" json:"tags"`               //标签
	Length       string   `json:"length"`                         //时长
	MagnetLinks  []string `json:"-"`                              //磁链
	Uncensored   bool     `json:"uncensored"`                     //有码,无码
}

// TableName ...
func (s SVideo) TableName() string {
	return "video"
}

// SeedVideoToConversionVideo ...
func SeedVideoToConversionVideo(v SVideo) Video {
	vd := Video{
		No:           v.Bangumi,
		Intro:        v.Intro,
		Alias:        v.Alias,
		ThumbHash:    v.ThumbHash,
		PosterHash:   v.PosterHash,
		SourceHash:   v.SourceHash,
		M3U8Hash:     v.M3U8Hash,
		Key:          v.Key,
		M3U8:         v.M3U8,
		Role:         v.Role,
		Director:     v.Director,
		Systematics:  v.Systematics,
		Season:       v.Season,
		TotalEpisode: v.TotalEpisode,
		Episode:      v.Episode,
		Producer:     v.Producer,
		Publisher:    v.Publisher,
		Type:         v.Type,
		Format:       v.Format,
		Language:     v.Language,
		Caption:      v.Caption,
		Group:        v.Group,
		Index:        v.Index,
		Date:         v.Date,
		Sharpness:    v.Sharpness,
		Series:       v.Series,
		Tags:         v.Tags,
		Length:       v.Length,
		Sample:       []string{},
		Uncensored:   v.Uncensored,
	}
	return vd
}
