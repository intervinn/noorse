import { User } from "discord.js";
import { Command, OptionType } from ".";

export default {
    name: "add",
    options: [
        {
            name: "member",
            type: "member"
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
        ctx.message.reply({
            embeds: [
                {
                    title: "successfully gave points",
                    description: `user: ${user.displayName}`,
                    fields: [
                        {name: "message", value: ctx.options.message}
                    ]
                }
            ]
        })
    }
} as Command