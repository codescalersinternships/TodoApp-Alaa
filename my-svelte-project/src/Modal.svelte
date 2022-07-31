<script>
import { onMount } from "svelte/internal";

let todos = []
const baseURL = "http://localhost:8080/todo"

const getTodos = async () => {
  const response = await fetch(baseURL)
  const data = await response.json()
  todos = await data
}

onMount(() => {
  getTodos()
})


let Task
let ID
const createTodo = async event => {
  event.preventDefault() 
  await fetch(baseURL, {
    method: "post",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: ID,
      task: Task,
    }),
  })

  getTodos()
  ID = ""
  Task = ""
}

const deleteTodo = async (todo) => {
  event.preventDefault()
  await fetch(baseURL+"/"+todo.id,{
    method: "delete",
  })
  getTodos()
}

const markTask = async (todo) => {
  event.preventDefault()
  await fetch(baseURL + "/"+todo.id,{
    method: "patch",
  })
  // getTodos()
}


</script>

<main>
  <h2>TODO APP</h2>

  <form on:submit="{createTodo}">
    <input type="text" bind:value="{ID}" id="inputID">
    <input type="text" bind:value="{Task}" id="inputTask">
    <input type="submit" value="Submit" />
  </form>

  <hr  style="color: #BC3CBC" >
  <hr  style="color: #BC3CBC" >

  {#each todos as todo}
  <div>
    <input type="checkbox" on:checked={(e) => markTask(todo)} bind:value="{todo.Done}">
    <h3>{todo.id} {todo.task}</h3>
    <button class="btn" on:click={(e) => deleteTodo(todo)}><i calss="fa fa-trash">❌</i></button>
  </div>
  {/each}

</main>


<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>

  input[type=text]{
    margin: 8px 0;
    border-radius: 20px;

  }

  #inputID{
    width: 8%;
    box-sizing: border-box;
    padding: 12px 20px;
    border: 1px solid rgb(212, 130, 212);
  }

  #inputTask{
    width: 50%;
    box-sizing: border-box;
    padding: 12px 20px;
    border: 1px solid rgb(212, 130, 212);
  }

  input[type=submit]{
    background-color: rgb(243, 230, 230);
    border: none;
    color: purple;
    padding: 16px 20px;
    text-decoration: none;
    margin: 4px 4px;
    cursor: pointer;
    border-radius: 25px;
  }

  h2{
    color: purple;
    font-size: 35px;
    font-weight: 400;
    background-image: linear-gradient(to left, #fdfdff, #d1bce5);
  }

  .btn {
  background-color: white; 
  border: none; 
  color: white; 
  padding: 12px 16px; 
  font-size: 16px; 
  position: absolute;
}

button, h3{
  display: inline-block;
}

.btn:hover {
  background-color: rgb(255, 255, 255);
}

h3{
  font-weight: bold;
  color: purple;
  padding-right: 33px;
  font-style: italic;
  margin-left: 9px;
  
}
</style>