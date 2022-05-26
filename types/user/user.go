package user

type Package struct {
	Result	int64	`json:"result"`
	User	User	`json:"user"`
	TS		int64	`json:"ts"`
}

type User struct {
	Status           UserStatus       				`json:"status"`          
	Dungeon          Dungeon          				`json:"dungeon"`         
	Troop            Troop            				`json:"troop"`           
	PushFlags        PushFlags       				`json:"pushFlags"`       
	Skin             Skin             				`json:"skin"`            
	Shop             Shop         	  				`json:"shop"`            
	Mission          Mission      	  				`json:"mission"`         
	Social           Social      	  				`json:"social"`          
	Building         Building         				`json:"building"`        
	DexNav           DexNav          				`json:"dexNav"`          
	Crisis           Crisis          	 			`json:"crisis"`          
	Medal            Medal            				`json:"medal"`           
	Activity 		 map[string]map[string]Actor	`json:"activity"`
	Gacha            Gacha            				`json:"gacha"`           
	Avatar           UserAvatar       				`json:"avatar"`          
	Background       Background       				`json:"background"`      
	Storyreview      StoryReview      				`json:"storyreview"`     
	Inventory        map[string]int64 				`json:"inventory"`       
	CollectionReward map[string]map[string]int64 	`json:"collectionReward"`
	OpenServer       OpenServer       				`json:"openServer"`      
	Equipment        Equipment        				`json:"equipment"`       
	Roguelike        Roguelike        				`json:"roguelike"`       
	Charm            Charm            				`json:"charm"`           
	CheckMeta        CheckMeta        				`json:"checkMeta"`       
	CheckIn          CheckIn      					`json:"checkIn"`         
	Consumable		 map[string]map[string]Consumed	`json:"consumable"`  
	CampaignsV2      CampaignsV2     				`json:"campaignsV2"`     
	Carousel         Carousel         				`json:"carousel"`        
	Recruit          Recruit          				`json:"recruit"`         
	TShop			 map[string]TShop				`json:"tshop"`	
	Event			 Event							`json:"event"`      
	Ticket           interface{}      				`json:"ticket"`          
	Retro            Retro            				`json:"retro"`           
}