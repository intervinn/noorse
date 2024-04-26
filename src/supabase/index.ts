import { createClient } from "@supabase/supabase-js";

function createSupabase() {
    return createClient(
        process.env.SUPABASE_URL!,
        process.env.SUPABASE_ANON_KEY!
    )
}

export default createSupabase