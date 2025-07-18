package commands

import (
	"petpet/lib"
	"petpet/pet_maker"
	"petpet/utils"
	"slices"
)

var PetpetUserCtx = lib.Command{
	Type:        2,
	Name:        "Petpet this user",
	Description: "",
	CommandHandler: func(interaction *lib.CommandInteraction) {
		userId := interaction.Data.TargetID

		if slices.Contains(utils.BlacklistedUsers, userId) {
			interaction.SendSimpleReply("This user is blacklisted, sorry.", true)
			return
		}

		member := interaction.Data.Resolved.Members[userId]
		user := interaction.Data.Resolved.Users[userId]

		avatar := member.GuildAvatarURL()
		if avatar == "" {
			avatar = user.AvatarURL()
		}

		interaction.Defer(false)

		img := pet_maker.MakePetImage(avatar, 1, 128, 128)

		interaction.EditReply(lib.ResponseMessageData{}, false, []lib.DiscordFile{
			{
				Filename: "petpet.gif",
				Reader:   img,
			},
		})
	},
}
