import { Message } from "discord.js"
import add from "./add"
import view from "./view"
import remove from "./remove"

export type OptionType = "string" | "int" | "user"

export type Options = {
    name: string,
    type: OptionType,
    greedy?: boolean
}

export type OptionMap = {[name: string]: any}

export type Context = {
    message: Message
    options: OptionMap
}

export type Command = {
    name: string,
    options: Options[],
    run(ctx: Context): any
}

export type OptionValidationResult = {
    success: boolean,
    result: string | OptionMap
}

export function validateOptions(cmd: Command, msg: Message): OptionValidationResult {
    const args = msg.content.split(" ").slice(1)

    let formatmsg = `lacking options, see the format:\n .${cmd.name}`
    cmd.options.forEach(v => {
        formatmsg += ` [${v.name}]`
    })

    if (args.length < cmd.options.length) {
        return {
            success: false,
            result: formatmsg
        }
    }

    let res: OptionMap = {}


    for (let i = 0; i < cmd.options.length; i++) {
        const v = cmd.options[i]

        if (v.greedy) {
            res[v.name] = args.slice(i).join(" ")
        } else {
            let val: any = args[i]
            if (v.type === "int") {
                val = Number(val)
                if (isNaN(val)) return {
                    success: false,
                    result: formatmsg
                }
            }

            if (v.type == "user") {
                val = msg.mentions.users.first()
                if (!val) return {
                    success: false,
                    result: formatmsg
                }
                
                const fk = msg.mentions.users.firstKey()
                msg.mentions.users.sweep((v, k) => v.id != val.id && k != fk)
            }
            
            
            res[v.name] = val
        }
    }

    return {
        success: true,
        result: res
    }

}


export default [
    add,
    view,
    remove
] as Command[]