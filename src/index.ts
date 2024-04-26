require("dotenv").config()

import { Client, IntentsBitField } from "discord.js"
import commands, { OptionMap, validateOptions } from "./commands"

const __token = process.env.TOKEN!

const client = new Client({
    intents: [
        IntentsBitField.Flags.Guilds,
        IntentsBitField.Flags.GuildMembers,
        IntentsBitField.Flags.GuildMessages,
        IntentsBitField.Flags.MessageContent
    ]
})

client.on("ready", () => {
    console.log("bot is ready")
})

client.on("messageCreate", (msg) => {
    if (msg.content.startsWith(".")) {
        const cmdName = msg.content.split(" ")[0].substring(1)
        const cmd = commands.find(v => v.name === cmdName)

        if (cmd) {

            let {result, success} = validateOptions(cmd, msg)
            if (!success) msg.reply(`Invalid command:\n\`\`\`${result}\`\`\``)
            else cmd.run({
                message: msg,
                options: result as OptionMap
            })
        }
    }
})



client.login(process.env.TOKEN!)