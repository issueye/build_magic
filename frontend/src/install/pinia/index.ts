import { createPinia } from "pinia"
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

export const piniaInstall = (app: any) => {
    const pinia = createPinia()
    pinia.use(piniaPluginPersistedstate)
    
    app.use(pinia)
}