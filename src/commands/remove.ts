import { Collection, PermissionFlagsBits, User } from "discord.js";
import { Command, OptionType } from ".";
import createSupabase from "../supabase";
import { addPoints } from "../supabase/points";

const supabase = createSupabase()

export default {
    name: "add",
    options: [
        {
            name: "amount",
            type: "int"
        },
        {
            name: "member",
            type: "user",
            greedy: true
        },
        {
            name: "message",
            type: "string",
            greedy: true
        }
    ],
    permissions: {
        roleName: "Bot Manager"
    },
    run: async (ctx) => {
        const users = ctx.options.member as Collection<string, User>
        users.forEach(async (v) => {
            const result = await addPoints({
                user: v,
                amount: ctx.options.amount * -1,
                supabase: supabase,
                guildId: ctx.message.guildId!
            })

            if (!result?.success) {
                ctx.message.reply(`${result?.error}`)
                return
            }

            ctx.message.reply({
                embeds: [
                    {
                        title: "successfully gave points",
                        description: `user: ${v.displayName}`,
                        fields: [
                            {name: "message", value: ctx.options.message},
                            {name: "old points", value: result.oldPoints},
                            {name: "new points", value: result.oldPoints + ctx.options.amount}
                        ],
                        thumbnail: {
                            url: v.displayAvatarURL({size: 128})
                        }
                    }
                ]
            })
        })

    }
} as Command