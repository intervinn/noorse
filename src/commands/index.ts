import { Message, PermissionFlagsBits } from "discord.js"
import add from "./add"
import view from "./view"
import remove from "./remove"
import help from "./help"

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
    permissions?: {
        roleName?: string,
        permission?: bigint
    }
    run(ctx: Context): any
}

export type OptionValidationResult = {
    success: boolean,
    result: string | OptionMap
}

export type PermissionValidationResult = {
    success: boolean,
    error?: string
}


export function makeFormat(cmd: Command, prefix?: string): string {
    let formatmsg = `${prefix} see the format:\n .${cmd.name}`
    cmd.options.forEach(v => {
        formatmsg += ` [${v.name}]`
    })
    return formatmsg
}

export function validatePermissions(cmd: Command, msg: Message): PermissionValidationResult {

    let valid = true

    if ((cmd.permissions?.roleName 
        && !msg.member?.roles.cache.find(v => v.name === cmd.permissions?.roleName)
        || (cmd.permissions?.permission
        && !msg.member?.permissions.has(cmd.permissions.permission))
    )) valid = false

    if (valid) {
        return {
            success: true
        }
    } else {
        return {
            success: false,
            error: "lacking permissions"
        }
    }
}

export function validateOptions(cmd: Command, msg: Message): OptionValidationResult {
    const args = msg.content.split(" ").slice(1)

    if (args.length < cmd.options.length) {
        return {
            success: false,
            result: makeFormat(cmd, "lacking options.")
        }
    }

    let res: OptionMap = {}


    for (let i = 0; i < cmd.options.length; i++) {
        const v = cmd.options[i]

        if (v.type === "string" && v.greedy) {
            res[v.name] = args.slice(i)
            break
        } 

        let val: any = args[i]
        if (v.type === "int") {
            val = Number(val)

            if (isNaN(val)) return {
                success: false,
                result: makeFormat(cmd, "lacking options.")
            }
        }

        if (v.type == "user") {
            val = msg.mentions.users.first()
            if (!val) return {
                success: false,
                result: makeFormat(cmd, "lacking options.")
            }

            if (v.greedy) {
                val = msg.mentions.users.clone()
            } else {
                const fk = msg.mentions.users.firstKey()
                msg.mentions.users.sweep((v, k) => v.id != val.id && k != fk)
            }
        }
        res[v.name] = val
    }

    return {
        success: true,
        result: res
    }

}


export const commands = [
    add,
    view,
    remove,
    help
] as Command[]
export default commands