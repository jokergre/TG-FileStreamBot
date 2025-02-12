package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
)

func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	dispatcher.AddHandler(handlers.NewCommand("start", start))
}

func start(ctx *ext.Context, u *ext.Update) error {
	chatId := u.EffectiveChat().GetID()
	peerChatId := ctx.PeerStorage.GetPeerById(chatId)
	if peerChatId.Type != int(storage.TypeUser) {
		return dispatcher.EndGroups
	}
	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
		ctx.Reply(u, "You are not allowed to use this bot.", nil)
		return dispatcher.EndGroups
	}
	ctx.Reply(u, "هلا وسهلا ، اتبع التعليمات ادناه لكي يعمل البوت عندك بصورة مستمرة :

✅|- اشترك بقناة البوت ( شبكة اوكسجين ) 👇🏻
@zezzez
@zezzez

ثم قم بأعادة توجيه مقطع الفيديو او ارساله الى البوت حتى تحصل على رابط المشاهدة ورابط للتحميل بصورة سريعة⚡️..", nil)
	return dispatcher.EndGroups
}
