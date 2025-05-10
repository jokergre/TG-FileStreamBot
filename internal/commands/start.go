package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"
	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// التحقق من الاشتراك
func checkSubscription(bot *tgbotapi.BotAPI, chatId int64, channelId string) bool {
	member, err := bot.GetChatMember(channelId, chatId)
	if err != nil {
		// في حال حدوث خطأ أثناء التحقق
		return false
	}

	// التحقق من حالة الاشتراك (يجب أن يكون العضو "مشارك" أو "أعضاء")
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		return true
	}
	return false
}

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

	// إذا كانت القائمة المسموحة غير فارغة، تحقق من أن المستخدم موجود في القائمة
	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
		ctx.Reply(u, "You are not allowed to use this bot.", nil)
		return dispatcher.EndGroups
	}

	// إنشاء اتصال بوت جديد باستخدام API Key الخاص بك
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_API_KEY") // استخدم الـ API Key الخاص بك
	if err != nil {
		ctx.Reply(u, "حدث خطأ في الاتصال بالبوت.", nil)
		return dispatcher.EndGroups
	}

	// التحقق من الاشتراك في القناة
	channelId := "@zezzez" // استبدلها باسم القناة التي تريد التحقق من الاشتراك فيها
	if !checkSubscription(bot, chatId, channelId) {
		ctx.Reply(u, "من فضلك اشترك في قناتنا أولاً قبل استخدام البوت. القناة: @zezzez", nil)
		return dispatcher.EndGroups
	}

	// رسالة ترحيبية وتوجيه المستخدم
	ctx.Reply(u, `هلا وسهلا، 
اتبع التعليمات أدناه لكي يعمل البوت عندك بصورة مستمرة:

✅|- اشترك بقناة البوت (شبكة اوكسجين) 👇🏻
@zezzez

ثم قم بإعادة توجيه مقطع الفيديو أو إرساله إلى البوت حتى تحصل على رابط المشاهدة ورابط للتحميل بصورة سريعة⚡️.`, nil)

	return dispatcher.EndGroups
}
