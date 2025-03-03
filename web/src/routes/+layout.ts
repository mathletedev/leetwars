import { redirect } from "@sveltejs/kit";
import { SERVER_URL } from "$lib/config";
import type { User } from "$types/user";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ url }) => {
    const res = await fetch(`${SERVER_URL}/api/me`, {
        credentials: "include",
    });

    if (res.status === 401 || url.pathname === "/init") {
        return {
            user: null,
        };
    }

    if (res.status === 418) {
        redirect(302, "/init");
    }

    const data = await res.json();

    return {
        user: data as User,
    };
};
