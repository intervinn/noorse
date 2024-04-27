require("dotenv").config()

import { Client, IntentsBitField } from "discord.js"
import commands, { OptionMap, validateOptions, validatePermissions } from "./commands"

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

            const opts = validateOptions(cmd, msg)
            if (!opts.success) {
                msg.reply(`Invalid command:\n\`\`\`${opts.result}\`\`\``)
                return
            }
            
            
            const perms = validatePermissions(cmd, msg)
            if (!perms.success) {
                msg.reply(`Invalid command:\n\`\`\`${perms.error}\`\`\``)
                return
            }

            cmd.run({
                message: msg,
                options: opts.result as OptionMap
            })
        }
    }
})



client.login(process.env.TOKEN!)