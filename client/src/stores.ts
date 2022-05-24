import { writable } from "svelte/store";
import type User from "./types/user";

export let userData = writable<User>({
    id: "",
    username: "",
    email: "",
});

export let isLogin = writable(false);
