import { SupabaseClient } from "@supabase/supabase-js"
import { GuildMember, User } from "discord.js"


export type AddPointOptions = {
    user: User | GuildMember, 
    guildId: string, 
    amount: number,
    supabase: SupabaseClient
}

export type AddPointsResult = {
    success: boolean,
    error?: string,
    oldPoints?: number
}

export async function addPoints({user, guildId, amount, supabase}: AddPointOptions) {
    const {data} = await supabase
    .from("users")
    .select()
    .eq("id", user.id)
    .eq("server", guildId)
    const dbuser = data?.at(0)

    if (!dbuser) {
        await supabase.from("users").upsert({
            id: user.id,
            points: amount,
            server: guildId
        })
    }

    const oldPoints = dbuser.points || 0

    const {error} = await supabase
        .from("users")
        .update({
            points: oldPoints + amount
        })
        .eq("id", user.id)
        .eq("server", guildId)
    
    if (error) {
        return {
            success: false,
            error: error.message
        }
    } else {
        return {
            success: true,
            oldPoints: oldPoints
        }
    }
}