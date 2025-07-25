package petpetsubcommands

import (
	"petpet/lib"
	"petpet/pet_maker"
	"petpet/utils"
	"slices"
)

var PetpetUser = lib.Command{
	Name:        "user",
	Description: "Petpet someone's pfp",
	Options:     append(utils.PetpetCommandUserOptions, utils.PetpetCommandOptions...),
	CommandHandler: func(interaction *lib.CommandInteraction) {

		untypedUser, err := interaction.GetStringOptionValue("user_to_petpet", "")
		if err != nil {
			interaction.SendSimpleReply("Invalid user ID provided.", true)
			return
		}
		userId, err := lib.StringToSnowflake(untypedUser)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse user ID. You shouldn't see this.", true)
			return
		}

		if slices.Contains(utils.BlacklistedUsers, userId) {
			interaction.SendSimpleReply("This user is blacklisted, sorry.", true)
			return
		}

		member := interaction.Data.Resolved.Members[userId]
		user := interaction.Data.Resolved.Users[userId]

		useServerAvatar, err := interaction.GetBoolOptionValue("use_server_avatar", true)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse use_server_avatar option. You shouldn't see this.", true)
			return
		}

		avatar := member.GuildAvatarURL()
		if !useServerAvatar || avatar == "" {
			avatar = user.AvatarURL()
		}

		ephemeral, err := interaction.GetBoolOptionValue("ephemeral", false)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse ephemeral option. You shouldn't see this.", true)
			return
		}

		interaction.Defer(ephemeral)

		speed, err := interaction.GetFloatOptionValue("speed", 1.0)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse speed option. You shouldn't see this.", true)
			return
		}
		width, err := interaction.GetIntOptionValue("width", 128)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse width option. You shouldn't see this.", true)
			return
		}
		height, err := interaction.GetIntOptionValue("height", 128)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse height option. You shouldn't see this.", true)
			return
		}

		img := pet_maker.MakePetImage(avatar, speed, width, height)

		interaction.EditReply(lib.ResponseMessageData{}, ephemeral, []lib.DiscordFile{
			{
				Filename: "petpet.gif",
				Reader:   img,
			},
		})
	},
}
