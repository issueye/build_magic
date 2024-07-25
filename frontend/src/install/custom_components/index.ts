import BsHeader from '@/components/bs_header/index.vue';
import BsMain from '@/components/bs_main/index.vue';
import BsDialog from '@/components/bs_dialog/index.vue';


export const customComponentsInstall = (app: any) => {
    app.component('BsHeader', BsHeader)
    app.component('BsMain', BsMain)
    app.component('BsDialog', BsDialog)
}