package best5

type SyncTrack struct {
	Type       string // sync
	Ts         int64
	UserId     uint64 `json:"user_id"`
	Credit     int64
	After      int64
	Bet        bool
	TermNumber int64
	GameType   int32
}

type StrategySystem struct {
	Type            string //触发系统策略
	Ts              int64
	Room            int32    //房间
	Table           uint64   //桌子
	UserNum         int32    //用户数量
	UserAmount      int64    //用户投注金额
	Term            int64    // 期数
	HandCardsBefore []string // 作弊前的牌组
	HandCards       []string // 包括庄在内的手牌数组
	StrategyId      string   // 触发策略
	Match           bool
	Hit             bool
}

type StrategyUser struct {
	Type            string //触发系统策略
	Ts              int64
	Room            int32    //房间
	Table           uint64   //桌子
	UserId          uint64   `json:"user_id"`
	UserAmount      int64    //用户投注金额
	Term            int64    // 期数
	HandCardsBefore []string // 作弊前的牌组
	HandCards       []string // 包括庄在内的手牌数组
	StrategyId      string   // 触发策略
	Match           bool
	Hit             bool
}

type BetTrack struct {
	Type       string // bet
	Ts         int64
	GameType   int32
	UserId     uint64 `json:"user_id"`
	TermNumber int64  // room kind
	BetIndex   int32
	Amount     int64
	RoomIndex  int32
	TableId    int64
	After      int64
	RobotLevel int
}

type LoginTrack struct {
	Type     string // login
	Ts       int64
	UserId   uint64
	GameType int32
	Addr     string
}

type LogoutTrack struct {
	Type     string // logout
	Ts       int64
	UserId   uint64
	GameType int32
}

type NewTermTrack struct {
	Type       string // new_term
	Ts         int64
	GameType   int32
	TableId    int64
	RoomIndex  int32
	TermNumber int64
	UserCount  int
	RobotCount int
	Banker     uint64
}

type AwardTrack struct {
	Type       string // award
	Ts         int64
	GameType   int32
	UserId     uint64
	TableId    int64
	Info       map[int32]int64
	Total      int64
	Award      int64
	Bet        int64
	After      int64
	Tax        int64
	RoomIndex  int32
	TermNumber int64
	RobotLevel int
}

type SalaryTrack struct {
	Type   string // award
	Ts     int64
	Salary int64
}

type GameStreamTrack struct {
	Type        string // award
	Ts          int64
	Term        int64
	UserCount   int32
	BankerAward int64
	BetTotalNum int32
	BetArea     string
	Tax         int64
}

type BetInfo struct {
	Bet      int64
	BetNum   int32
	HandCard string
	Pattern  int32
	IsWin    bool
}

type UserStreamTrack struct {
	Type    string // award
	Ts      int64
	Term    int64
	UserId  uint64
	Bet     int64
	Award   int64
	Credit  int64
	Tax     int64
	BetArea string
}

type UserBetInfo struct {
	Bet      int64
	PayRate  int64
	IsWin    bool
	HandCard string
	Pattern  int32
}

type AnnounceTrack struct {
	Type        string // announce 结算
	Ts          int64
	GameType    int32
	TableId     int64
	TermNumber  int64
	RoomIndex   int32
	TotalTax    int64
	TotalBet    int64
	TotalAward  int64
	WinInfo     string
	IsAdjust    bool
	WaterPool   int64
	Banker      uint64
	BankerAward int64
	BankerTax   int64
	TrimRate    float64
	RefundRate  float64
}
