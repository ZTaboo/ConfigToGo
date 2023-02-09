<template>
    <div style="display: flex;height: 100%">
        <Sidebar :menuItemValue="menuItemValue" @editMenuItemValue="editMenuItemValue"></Sidebar>
        <Content :key="menuItemValue" :menuItemValue="menuItemValue"></Content>
    </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import "@fontsource/material-icons"
import "./assets/css/icon.css"
import Sidebar from "./components/Sidebar.vue";
import Content from "./components/Content.vue";

const menuItemValue = ref('json')
const editMenuItemValue = (name: string) => {
    menuItemValue.value = name
}
</script>
<style scoped>
</style>
