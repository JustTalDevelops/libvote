package libvote

const (
	BaseUrl             = "https://minecraftpocket-servers.com"
	BaseApiUrl          = BaseUrl + "/api/"
	QueryApiUrl         = BaseApiUrl + "?"
	ServerUrl           = BaseUrl + "/server/%v/"
	ListServersEndpoint = QueryApiUrl + "object=servers&element=list"
	ServerInfoEndpoint  = QueryApiUrl + "object=servers&element=detail&key=%v"
	VotesListEndpoint   = QueryApiUrl + "object=servers&element=votes&key=%v&format=%s"
	VoteEndpoint        = ServerUrl + "vote/action/"
)
