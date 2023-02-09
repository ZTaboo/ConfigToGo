<template>
    <div style="height: 100vh;">
        <va-sidebar hoverable minimizedWidth="64px">
            <va-sidebar-item
                v-for="item in menus"
                :key="item.title"
                :active="item.active"
            >
                <va-sidebar-item-content class="menu-item" @click="menuItemBtn(item.title)">
                    <va-icon :name="item.icon"/>
                    <va-sidebar-item-title>
                        <span class="menu-item" size="medium">{{ item.title }}</span>
                    </va-sidebar-item-title>
                </va-sidebar-item-content>
            </va-sidebar-item>
        </va-sidebar>
    </div>
</template>

<script lang="ts" setup>
import {defineProps, onMounted, Ref, ref} from "vue";

const emit = defineEmits(['editMenuItemValue'])
const props = defineProps(["menuItemValue"])
onMounted(() => {
})

const menus: Ref = ref([
    {
        title: 'json',
        icon: 'data_object',
        active: true
    },
    {
        title: 'yaml',
        icon: 'terminal',
        active: false
    },
    {
        title: 'toml',
        icon: 'link',
        active: false
    },
    {
        title: 'hcl',
        icon: 'code_off',
        active: false
    },
    {
        title: 'env',
        icon: 'settings_suggest',
        active: false
    },
    {
        title: 'ini',
        icon: 'recycling',
        active: false
    }
])

const menuItemBtn = (e: string) => {
    let copyList = [...menus.value]
    let tmpList: any = []
    copyList.map(item => {
        let tmpObj = {}
        if (item.title === e) {
            emit('editMenuItemValue', item.title)
            tmpObj = {
                title: item.title,
                icon: item.icon,
                active: true
            }
        } else {
            tmpObj = {
                title: item.title,
                icon: item.icon,
                active: false
            }
        }

        tmpList.push(tmpObj)
    })
    menus.value = tmpList
}

</script>

<style scoped>
.menu-item {
    cursor: pointer;
    user-select: none;
}
</style>