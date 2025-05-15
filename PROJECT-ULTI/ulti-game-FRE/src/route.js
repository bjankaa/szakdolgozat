import { createRouter, createWebHistory } from "vue-router";

import FrontPage from "./pages/FrontPage.vue";
import AuthPage from "./pages/AuthPage.vue";
import Game from "./pages/Game.vue";
import Profile from "./pages/Profile.vue";
import Settings from "./pages/Settings.vue";
import NotFound from "./pages/NotFound.vue";


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', redirect: '/frontpage'},
        {path: '/frontpage', component: FrontPage},
        {path: '/auth', component: AuthPage},
        {path: '/game', component: Game},
        {path: '/profile', component: Profile},
        {path: '/settings', component: Settings},
        {path: '/:notFound(.*)', component: NotFound}
    ]
});

export default router;