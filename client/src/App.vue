<script setup lang="ts">
import { reactive, ref } from 'vue'
import axios from 'axios'
let todoVal = ref<string>('')
interface todo {
    act: string
    done?: boolean
    id?: string
}

let todos = reactive<todo[]>([])

async function addToDo() {
    let newTodo: todo = {
        act: todoVal.value,
    }
    todoVal.value = ''
    try {
        const { data } = await axios.post('http://localhost:42069', newTodo)
        newTodo.id = data
        newTodo.done = false
        todos.push(newTodo)
    } catch (error) {
        console.log('couldnt insert it')
    }
}
async function deleteElement(index:number){
console.log()
}
</script>

<template>
    <form @submit.prevent="addToDo">
        <input v-model="todoVal" placeholder="addtodo" required />
        <button>submit</button>
    </form>
    <ul>
        <li v-for="todo,index in todos" :key="index">
            {{ todo }}
            <button @click.prevent="deleteElement(index)">delete</button>
        </li>
    </ul>
</template>

<style scoped></style>
