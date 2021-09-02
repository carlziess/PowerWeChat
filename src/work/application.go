package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/work/agent"
	"github.com/ArtisanCloud/power-wechat/src/work/auth"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"github.com/ArtisanCloud/power-wechat/src/work/corpgroup"
	"github.com/ArtisanCloud/power-wechat/src/work/department"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact"
	"github.com/ArtisanCloud/power-wechat/src/work/groupRobot"
	"github.com/ArtisanCloud/power-wechat/src/work/invoice"
	"github.com/ArtisanCloud/power-wechat/src/work/media"
	"github.com/ArtisanCloud/power-wechat/src/work/menu"
	"github.com/ArtisanCloud/power-wechat/src/work/message"
	msgaudit "github.com/ArtisanCloud/power-wechat/src/work/msgAudit"
	"github.com/ArtisanCloud/power-wechat/src/work/oa"
	"github.com/ArtisanCloud/power-wechat/src/work/oauth"
	"github.com/ArtisanCloud/power-wechat/src/work/server"
	"github.com/ArtisanCloud/power-wechat/src/work/user"
	"net/http"
)

type Work struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config     *kernel.Config
	Department *department.Client

	Agent          *agent.Client
	AgentWorkbench *agent.WorkbenchClient

	Message  *message.Client
	Messager *message.Messager

	Encryptor *kernel.Encryptor
	Server    *server.Guard

	UserClient           *user.Client
	UserBatchJobsClient  *user.BatchJobsClient
	UserLinkedCorpClient *user.LinkedCorpClient
	UserTagClient        *user.TagClient

	ExternalContact                *externalContact.Client
	ExternalContactContactWay      *externalContact.ContactWayClient
	ExternalContactStatistics      *externalContact.StatisticsClient
	ExternalContactMessage         *externalContact.MessageClient
	ExternalContactSchool          *externalContact.SchoolClient
	ExternalContactMoment          *externalContact.MomentClient
	ExternalContactMessageTemplate *externalContact.MessageTemplateClient

	Media *media.Client
	Menu  *menu.Client

	OA *oa.Client

	MsgAudit *msgaudit.Client

	CorpGroup *corpgroup.Client

	Invoice *invoice.Client

	GroupRobot          *groupRobot.Client
	GroupRobotMessenger *groupRobot.Messager
}

type UserConfig struct {
	CorpID           string
	AgentID          int
	Secret           string
	Token            string
	AESKey           string
	AuthCallbackHost string

	ResponseType string
	Log          Log
	OAuth        OAuth
	HttpDebug    bool
	Debug        bool
}

type Log struct {
	Level string
	File  string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewWork(config *UserConfig) (*Work, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.AccessToken = auth.RegisterProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)

	//-------------- register agent --------------
	app.Agent,
		app.AgentWorkbench = agent.RegisterProvider(app)

	//-------------- register Department --------------
	app.Department = department.RegisterProvider(app)

	//-------------- register Message --------------
	app.Message, app.Messager = message.RegisterProvider(app)

	//-------------- register Encryptor --------------
	app.Encryptor, app.Server = server.RegisterProvider(app)

	//-------------- register user --------------
	app.UserClient,
		app.UserBatchJobsClient,
		app.UserLinkedCorpClient,
		app.UserTagClient = user.RegisterProvider(app)

	//-------------- register external contact --------------
	app.ExternalContact,
		app.ExternalContactContactWay,
		app.ExternalContactStatistics,
		app.ExternalContactMessage,
		app.ExternalContactSchool,
		app.ExternalContactMoment,
		app.ExternalContactMessageTemplate = externalContact.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

	//-------------- menu --------------
	app.Menu = menu.RegisterProvider(app)

	//-------------- oa --------------
	app.OA = oa.RegisterProvider(app)

	//-------------- msg audit --------------
	app.MsgAudit = msgaudit.RegisterProvider(app)

	//-------------- corp group --------------
	app.CorpGroup = corpgroup.RegisterProvider(app)

	//-------------- invoice --------------
	app.Invoice = invoice.RegisterProvider(app)

	app.GroupRobot, app.GroupRobotMessenger = groupRobot.RegisterProvider(app)

	return app, err
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Work) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *Work) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Work) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "OAuth":
		return app.OAuth
	case "Config":
		return app.Config
	case "Department":
		return app.Department

	case "Message":
		return app.Message
	case "Messager":
		return app.Messager

	case "Encryptor":
		return app.Encryptor
	case "Server":
		return app.Server

	case "UserClient":
		return app.UserClient
	case "UserBatchJobsClient":
		return app.UserBatchJobsClient
	case "UserLinkedCorpClient":
		return app.UserLinkedCorpClient
	case "UserTagClient":
		return app.UserTagClient

	case "ExternalContact":
		return app.ExternalContact
	case "ExternalContactContactWay":
		return app.ExternalContactContactWay
	case "ExternalContactStatistics":
		return app.ExternalContactStatistics
	case "ExternalContactMessage":
		return app.ExternalContactMessage
	case "ExternalContactSchool":
		return app.ExternalContactSchool
	case "ExternalContactMoment":
		return app.ExternalContactMoment
	case "ExternalContactMessageTemplate":
		return app.ExternalContactMessageTemplate

	case "Menu":
		return app.Menu
	case "OA":
		return app.OA
	case "MsgAudit":
		return app.MsgAudit
	case "CorpGroup":
		return app.CorpGroup
	case "Invoice":
		return app.Invoice

	case "GroupRobot":
		return app.GroupRobot
	case "GroupRobotMessenger":
		return app.GroupRobotMessenger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"corp_id":            userConfig.CorpID,
		"agent_id":           userConfig.AgentID,
		"secret":             userConfig.Secret,
		"token":              userConfig.Token,
		"aes_key":            userConfig.AESKey,
		"auth_callback_host": userConfig.AuthCallbackHost,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"oauth.callback": userConfig.OAuth.Callback,
		"oauth.scopes":   userConfig.OAuth.Scopes,
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
	}

	return config, nil

}
