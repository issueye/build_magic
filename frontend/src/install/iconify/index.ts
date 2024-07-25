import { addCollection } from '@iconify/vue'
import uimIcons from '@iconify/json/json/uim.json'
import lineMdIcons from '@iconify/json/json/line-md.json'
import wiIcons from '@iconify/json/json/wi.json'

export async function iconifyInstall() {
    addCollection(uimIcons)
    addCollection(lineMdIcons)
    addCollection(wiIcons)
}
