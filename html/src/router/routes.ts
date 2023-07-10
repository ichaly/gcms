import HelloWorld from "~/components/HelloWorld.vue";

const routes: any[] = [
    {
        path: "/",
        name: "HelloWorld",
        component: HelloWorld,
    },
];

interface Route {
    path: string
    name: string
    meta: Record<string, string | boolean>

    component(): any
}

/** 当路由很多时，自动导入Modules下拆分的成小模块 **/
const routeModuleFiles = import.meta.globEager('./modules/*.ts')
Object.keys(routeModuleFiles).forEach((key: string) => {
    const module = routeModuleFiles[key].default
    module.forEach((route: Route) => {
        routes.push(route)
    })
})

export default routes