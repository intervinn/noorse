import { User } from "discord.js";
import { Command, OptionType } from ".";
import createSupabase from "../supabase";

const supabase = createSupabase()

export default {
    name: "remove",
    options: [
        {
            name: "member",
            type: "user"
        },
        {
            name: "amount",
            type: "int"
        },
        {
            name: "message",
            type: "string",
            greedy: true
        }
    ],
    run: async (ctx) => {
        const user = ctx.options.member as User

        
        const {data} = await supabase
            .from("users")
            .select()
            .eq("id", user.id)
            .eq("server", ctx.message.guildId)
        const dbuser = data?.at(0)

        if (!dbuser) {
            await supabase.from("users").upsert({
                id: user.id,
                points: ctx.options.amount,
                server: ctx.message.guildId
            })
        }

        let oldPoints = dbuser.points || 0
        let newPoints = oldPoints - ctx.options.amount

        const {error} = await supabase
            .from("users")
            .update({
                points: newPoints
            })
            .eq("id", user.id)
            .eq("server", ctx.message.guildId)
        if (error) {
            ctx.message.reply(`there was en error on trying to update data: ${error}`)
        }

        ctx.message.reply({
            embeds: [
                {
                    title: "successfully gave points",
                    description: `user: ${user.displayName}`,
                    fields: [
                        {name: "message", value: ctx.options.message},
                        {name: "old points", value: oldPoints},
                        {name: "new points", value: newPoints}
                    ],
                    thumbnail: {
                        url: user.displayAvatarURL({size: 128})
                    }
                }
            ]
        })
    }
} as Command