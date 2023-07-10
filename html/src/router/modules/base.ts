const baseRouter = [
    {
        path: "/about",
        name: "About",
        component: () => import("~/components/About.vue")
    }
]

export default baseRouter