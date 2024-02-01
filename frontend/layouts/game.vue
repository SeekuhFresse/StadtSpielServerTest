<script setup lang="ts">
const items = [{
    label: 'Tab1',
    content: 'This is the content shown for Tab1'
}, {
    label: 'Tab2',
    content: 'And, this is the content for Tab2'
}, {
    label: 'Tab3',
    content: 'Finally, this is the content for Tab3'
}]

</script>

<template>
    <div id="game-wrapper">
        <aside id="menu">
            <UTabs ui:="{ heigh: '1px'}" :items="items" />
        </aside>
        <main id="map">
            <NuxtPage />
        </main>
        <aside id="ressources">
jhgjgh
        </aside>
    </div>
</template>

<style lang="scss">
#game-wrapper {
    @apply flex h-screen;

    #menu {
        @apply w-[300px];

        #tabmenu {
            @apply h-[20px];
        }
    }

    #map {
        @apply flex-1 border;
    }

    #ressources {
        @apply w-[200px];
    }

}
</style>