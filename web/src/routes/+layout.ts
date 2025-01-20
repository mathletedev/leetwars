import { SERVER_URL } from "$lib/config";
import type { User } from "$types/user";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async () => {
    const res = await fetch(`${SERVER_URL}/api/me`, {
        credentials: "include",
    });

    if (res.status === 401) {
        return {
            user: null,
        };
    }

    const data = await res.json();

    return {
        user: data.RawData as User,
    };
};
