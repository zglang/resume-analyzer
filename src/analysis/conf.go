package analysis

func InitConf() {
	TagMap = make(TMap)
	GroupMap = make(TMap)
	TagList = make([]string, 100)
	TagFilters = make([]IFilter, 100)
	MaterialCorrections = make([]Correction,100)
	GroupProcessors = make([]GroupProcessor,100)
	initReplaceMap()
	initMap()
	initFilter()
	initWord()
	initGroup()
	initCorrections()
	initGroupProcessor()
	initPosition()
	initIndustry()
	initBrTags()
	initNotNameTag()
}
