package service

import (
	"github.com/bwmarrin/discordgo"
	"sanjose/config"
	"sanjose/model"
	"sanjose/utils"
)

var Discord *discordgo.Session

func ConnectDiscord() {
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		utils.SugarLogger.Errorln("Error creating Discord session, ", err)
		return
	}
	Discord = dg
	_, err = Discord.ChannelMessageSend(config.DiscordChannel, ":white_check_mark: "+config.Service.Name+" v"+config.Version+" online! `[ENV = "+config.Env+"]`")
	if err != nil {
		utils.SugarLogger.Errorln("Error sending Discord message, ", err)
		return
	}
}

func DiscordLogNewUser(user model.User) {
	var embeds []*discordgo.MessageEmbed
	var fields []*discordgo.MessageEmbedField
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "ID",
		Value:  user.ID,
		Inline: false,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Email",
		Value:  user.Email,
		Inline: true,
	})
	//fields = append(fields, &discordgo.MessageEmbedField{
	//	Name:   "School",
	//	Value:  user.School.School.Name,
	//	Inline: true,
	//})
	embeds = append(embeds, &discordgo.MessageEmbed{
		Title: "New Account Created!",
		Color: 6609663,
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://app.pacificesports.org/u/" + user.ID,
			Name:    user.FirstName + " " + user.LastName,
			IconURL: user.ProfilePictureURL,
		},
		Fields: fields,
	})
	_, err := Discord.ChannelMessageSendEmbeds(config.DiscordChannel, embeds)
	if err != nil {
		utils.SugarLogger.Errorln(err.Error())
	}
}
