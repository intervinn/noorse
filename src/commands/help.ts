import { APIEmbedField } from "discord.js";
import commands, { Command, makeFormat } from ".";

export default {
    name: "view",
    run: async (ctx) => {

        const fields: APIEmbedField[] = []
        commands.forEach(v => {
            fields.push({
                name: v.name,
                value: makeFormat(v)
            })
        })

        ctx.message.reply({
            embeds: [
                {
                    title: "Noorse - a bot for your point needs",
                    description: "Prefix- `.`",
                    fields: fields
                }
            ]
        })
    }
} as Command