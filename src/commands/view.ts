import { User } from "discord.js";
import { Command } from ".";
import createSupabase from "../supabase";

const supabase = createSupabase()

export default {
    name: "view",
    options: [
        {
            name: "user",
            type: "user"
        }
    ],
    run: async (ctx) => {
        const user = ctx.options.user as User
        const {data} = await supabase
            .from("users")
            .select()
            .eq("id", user.id)
            .eq("server", ctx.message.guildId)

        const dbuser = data?.at(0)
        if (!dbuser) {
            return ctx.message.reply(
                "this user isn't logged on this server"
            )
        }  

        ctx.message.reply({
            embeds: [
                {
                    title: user.displayName,
                    fields: [
                        {name: "points", value: dbuser.points || 0}
                    ],
                    thumbnail: {
                        url: user.displayAvatarURL({size: 128})
                    }
                }
            ]
        })

        
    }
} as Command