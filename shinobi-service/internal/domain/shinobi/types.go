package shinobi

import "time"

// ShinobiStatus — статус жизненного цикла шиноби в реестре.
type ShinobiStatus string

const (
	// StatusAvailable — доступен для назначения на миссию.
	StatusAvailable ShinobiStatus = "available"
	// StatusReserved — зарезервирован при сборе команды.
	StatusReserved ShinobiStatus = "reserved"
	// StatusOnMission — находится на миссии.
	StatusOnMission ShinobiStatus = "onMission"
	// StatusInjured — ранен, назначению не подлежит.
	StatusInjured ShinobiStatus = "injured"
	// StatusOnLeave — в отпуске.
	StatusOnLeave ShinobiStatus = "onLeave"
	// StatusDisabled — списан по состоянию здоровья.
	StatusDisabled ShinobiStatus = "disabled"
	// StatusDead — погиб.
	StatusDead ShinobiStatus = "dead"
	// StatusNukenin — дезертир; терминальный статус.
	StatusNukenin ShinobiStatus = "nukenin"
)

// ShinobiRank — ранг шиноби.
type ShinobiRank string

const (
	// Genin — младший ранг, выпускник Академии.
	Genin ShinobiRank = "genin"
	// Chuunin — средний ранг.
	Chuunin ShinobiRank = "chuunin"
	// Jounin — старший ранг, может вести отряд.
	Jounin ShinobiRank = "jounin"
	// Anbu — спецподразделение; данные ограниченного доступа.
	Anbu ShinobiRank = "anbu"
	// Kage — глава деревни.
	Kage ShinobiRank = "kage"
)

// JutsuLevel — уровень владения дзюцу, от D (низший) до S (высший).
type JutsuLevel int

const (
	// LevelS — высший уровень.
	LevelS JutsuLevel = 5
	// LevelA — продвинутый уровень.
	LevelA JutsuLevel = 4
	// LevelB — уровень выше среднего.
	LevelB JutsuLevel = 3
	// LevelC — средний уровень.
	LevelC JutsuLevel = 2
	// LevelD — базовый уровень.
	LevelD JutsuLevel = 1
)

// Gender — пол шиноби.
type Gender string

const (
	// Male — мужской.
	Male Gender = "male"
	// Female — женский.
	Female Gender = "female"
	// Unknown — не указан.
	Unknown Gender = "unknown"
)

// CompletedMissions — счётчики завершённых миссий шиноби по рангам.
type CompletedMissions struct {
	// MissionsRankS — число завершённых миссий ранга S.
	MissionsRankS int
	// MissionsRankA — число завершённых миссий ранга A.
	MissionsRankA int
	// MissionsRankB — число завершённых миссий ранга B.
	MissionsRankB int
	// MissionsRankC — число завершённых миссий ранга C.
	MissionsRankC int
	// MissionsRankD — число завершённых миссий ранга D.
	MissionsRankD int
}

// KekkeiGenkai — наследственная способность (особый тип способности).
type KekkeiGenkai string

const (
	// Sharingan — додзюцу клана Учиха.
	Sharingan KekkeiGenkai = "sharingan"
	// Byakugan — додзюцу клана Хьюга.
	Byakugan KekkeiGenkai = "byakugan"
	// MixElemental — комбинированная стихийная способность.
	MixElemental KekkeiGenkai = "mixElemental"
)

// SpecialSkill — особый навык шиноби; используется при подборе отряда (FR-CAP2).
type SpecialSkill string

const (
	// MedicineSkill — медицинские техники.
	MedicineSkill SpecialSkill = "medicine"
	// WeaponSkill — владение оружием.
	WeaponSkill SpecialSkill = "weaponUser"
	// Flying — способность к полёту.
	Flying SpecialSkill = "flying"
)

// Elemental — стихия чакры.
type Elemental string

const (
	// Katon — огонь.
	Katon Elemental = "katon"
	// Futon — ветер.
	Futon Elemental = "futon"
	// Raiton — молния.
	Raiton Elemental = "raiton"
	// Doton — земля.
	Doton Elemental = "doton"
	// Suiton — вода.
	Suiton Elemental = "suiton"
)

// Shinobi — карточка шиноби в реестре деревни (FR-S1).
type Shinobi struct {
	// ID — идентификатор шиноби.
	ID int
	// Name — имя.
	Name string
	// Age — возраст, полных лет.
	Age int
	// Gender — пол.
	Gender Gender
	// Clan — название клана.
	Clan string
	// SenseiID — идентификатор наставника; nil, если наставника нет.
	SenseiID *int
	// TeamID - принадлежность к постоянной команде; nil, если команды нет.
	TeamNumber *int

	// Rank — текущий ранг.
	Rank ShinobiRank
	// CompletedMissions — статистика завершённых миссий по рангам.
	CompletedMissions CompletedMissions

	// Status — текущий статус жизненного цикла.
	Status ShinobiStatus

	// TaijutsuLevel — уровень владения тайдзюцу.
	TaijutsuLevel JutsuLevel
	// NinjutsuLevel — уровень владения ниндзюцу.
	NinjutsuLevel JutsuLevel
	// GenjutsuLevel — уровень владения гендзюцу.
	GenjutsuLevel JutsuLevel

	// Elementals — освоенные стихии чакры.
	Elementals []Elemental
	// SpecialSkills — особые навыки.
	SpecialSkills []SpecialSkill
	// KekkeigGnkai — наследственная способность; nil, если нет.
	KekkeiGenkai *KekkeiGenkai

	// CreatedAt — когда запись создана.
	CreatedAt time.Time
	// StatusChangedAt — когда установлен текущий статус.
	StatusChangedAt time.Time
}
