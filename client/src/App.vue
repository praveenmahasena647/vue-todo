<script setup lang="ts">
import axios from 'axios'
import { reactive, ref, onBeforeMount } from 'vue'
// type stuffs
interface todo {
    act: string
    done: boolean
    id: string
}
let todoCollections = reactive<todo[]>([])
let todoAct = ref<string>('')
// last done adding stuffs to db thats all
onBeforeMount(async function () {
    try {
        let { data } = await axios.get('http://localhost:42069')
        if (data) {
            todoCollections.push(...data)
        }
    } catch (error) {
        console.log(error)
    }
})

async function addtodo() {
    let todo: todo = {
        id: '',
        act: todoAct.value,
        done: false,
    }

    try {
        //@ts-ignore
        let { data } = await axios.post('http://localhost:42069', todo)
        todo.id = data
        todoCollections.push(todo)
        todoAct.value = ''
    } catch (error) {
        console.log(error)
    }
}
async function deleteEm(index: number) {
    // have to do that all array thing
    try {
        let { data } = await axios.delete(
            `http://localhost:42069/${todoCollections[index].id}`
        )
        todoCollections.splice(index,1)
    } catch (error) {
        console.log(error)
    }
}
async function deleteAll() {
    try {
        let { data } = await axios.delete('http://localhost:42069')
        console.log(data)
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
    <form @submit.prevent="addtodo">
        <input v-model="todoAct" required placeholder="Activity" />
        <button>submit</button>
    </form>
    <div>
        <div>
            <button @click="deleteAll">delete All</button>
        </div>
        <div v-for="(todos, index) in todoCollections" :key="index">
            <div v-if="todos.act">
                {{ todos.act }}
                <button @click="deleteEm(index)">DeleteEm</button>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss"></style>
