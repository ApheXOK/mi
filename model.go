package mi

type MediaInfo struct {
	CreatingLibrary CreatingLibrary `json:"creatingLibrary,omitempty"`
	Media           Media           `json:"media,omitempty"`
}
type CreatingLibrary struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	URL     string `json:"url,omitempty"`
}
type Extra struct {
	Bsid            string `json:"bsid,omitempty"`
	Dialnorm        string `json:"dialnorm,omitempty"`
	Compr           string `json:"compr,omitempty"`
	Acmod           string `json:"acmod,omitempty"`
	Lfeon           string `json:"lfeon,omitempty"`
	DialnormAverage string `json:"dialnorm_Average,omitempty"`
	DialnormMinimum string `json:"dialnorm_Minimum,omitempty"`
	ComprAverage    string `json:"compr_Average,omitempty"`
	ComprMinimum    string `json:"compr_Minimum,omitempty"`
	ComprMaximum    string `json:"compr_Maximum,omitempty"`
	ComprCount      string `json:"compr_Count,omitempty"`
}

type Track struct {
	Type                           string `json:"@type,omitempty"`
	UniqueID                       string `json:"UniqueID,omitempty"`
	VideoCount                     string `json:"VideoCount,omitempty"`
	AudioCount                     string `json:"AudioCount,omitempty"`
	TextCount                      string `json:"TextCount,omitempty"`
	MenuCount                      string `json:"MenuCount,omitempty"`
	FileExtension                  string `json:"FileExtension,omitempty"`
	Format                         string `json:"Format,omitempty"`
	FormatVersion                  string `json:"Format_Version,omitempty"`
	FileSize                       string `json:"FileSize,omitempty"`
	Duration                       string `json:"Duration,omitempty"`
	OverallBitRateMode             string `json:"OverallBitRate_Mode,omitempty"`
	OverallBitRate                 string `json:"OverallBitRate,omitempty"`
	FrameRate                      string `json:"FrameRate,omitempty"`
	FrameCount                     string `json:"FrameCount,omitempty"`
	StreamSize                     string `json:"StreamSize,omitempty"`
	IsStreamable                   string `json:"IsStreamable,omitempty"`
	EncodedDate                    string `json:"Encoded_Date,omitempty"`
	FileModifiedDate               string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal          string `json:"File_Modified_Date_Local,omitempty"`
	EncodedApplication             string `json:"Encoded_Application,omitempty"`
	EncodedLibrary                 string `json:"Encoded_Library,omitempty"`
	StreamOrder                    string `json:"StreamOrder,omitempty"`
	ID                             string `json:"ID,omitempty"`
	FormatProfile                  string `json:"Format_Profile,omitempty"`
	FormatLevel                    string `json:"Format_Level,omitempty"`
	FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
	FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
	CodecID                        string `json:"CodecID,omitempty"`
	BitRateMode                    string `json:"BitRate_Mode,omitempty"`
	BitRate                        string `json:"BitRate,omitempty"`
	BitRateMaximum                 string `json:"BitRate_Maximum,omitempty"`
	Width                          string `json:"Width,omitempty"`
	Height                         string `json:"Height,omitempty"`
	StoredHeight                   string `json:"Stored_Height,omitempty"`
	SampledWidth                   string `json:"Sampled_Width,omitempty"`
	SampledHeight                  string `json:"Sampled_Height,omitempty"`
	PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
	DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
	FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
	FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
	FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
	ColorSpace                     string `json:"ColorSpace,omitempty"`
	ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
	BitDepth                       string `json:"BitDepth,omitempty"`
	ScanType                       string `json:"ScanType,omitempty"`
	Delay                          string `json:"Delay,omitempty"`
	DelaySource                    string `json:"Delay_Source,omitempty"`
	EncodedLibraryName             string `json:"Encoded_Library_Name,omitempty"`
	EncodedLibraryVersion          string `json:"Encoded_Library_Version,omitempty"`
	EncodedLibrarySettings         string `json:"Encoded_Library_Settings,omitempty"`
	Language                       string `json:"Language,omitempty"`
	Default                        string `json:"Default,omitempty"`
	Forced                         string `json:"Forced,omitempty"`
	BufferSize                     string `json:"BufferSize,omitempty"`
	ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
	ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
	ColourPrimaries                string `json:"colour_primaries,omitempty"`
	ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
	TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
	TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
	MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
	MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
	Typeorder                      string `json:"@typeorder,omitempty"`
	FormatCommercialIfAny          string `json:"Format_Commercial_IfAny,omitempty"`
	FormatSettingsEndianness       string `json:"Format_Settings_Endianness,omitempty"`
	Channels                       string `json:"Channels,omitempty"`
	ChannelPositions               string `json:"ChannelPositions,omitempty"`
	ChannelLayout                  string `json:"ChannelLayout,omitempty"`
	SamplesPerFrame                string `json:"SamplesPerFrame,omitempty"`
	SamplingRate                   string `json:"SamplingRate,omitempty"`
	SamplingCount                  string `json:"SamplingCount,omitempty"`
	CompressionMode                string `json:"Compression_Mode,omitempty"`
	VideoDelay                     string `json:"Video_Delay,omitempty"`
	ServiceKind                    string `json:"ServiceKind,omitempty"`
	Extra                          Extra  `json:"extra,omitempty"`
	ElementCount                   string `json:"ElementCount,omitempty"`
	Title                          string `json:"Title,omitempty"`
	Extra0                         string `json:"extra,omitempty"`
}

type Media struct {
	Ref   string  `json:"@ref,omitempty"`
	Track []Track `json:"track,omitempty"`
}
