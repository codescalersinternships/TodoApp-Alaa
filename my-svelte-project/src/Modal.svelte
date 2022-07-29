<script>
$:lists = []
axios.get(`http://localhost:8080/todo`).then(res => {
  lists = res.data;
  this.setState({lists});

})

let task =""

async function createTask(){
  const res = await fetch('http://localhost:8080/todo`',{
    method: 'POST',
    body: JSON.stringify({
      id,
      task,
      done: false

    })
  })

  const json = await res.json()
  result = JSON.stringify(json)
  lists.push(task)
  var node = document.createElement('li');
  node.appendChild(document.createTextNode('myUL'));
  document.querySelector('ul').appendChild(node);
}


</script>

<div clas ="backdrop">
    <dir class="modal">
        <div id="myDIV" class="header">
            <h2>My To Do List</h2>
            <input type="text" id="myInput" placeholder="Title...">
            <button id="Button" type="button" onclick="createTask()"> Add</button>
          </div>
          
          <ol id="myUL">
           
            {#each lists as list }
            <input type="checkbox" id="checkBox"> <li >{list.id} {list.task}</li> <span id="span">❌</span>
          
            {/each}
          </ol>
    </dir>

    <section>
        <script>

        </script>
    </section>

</div>


<style>
    
* {
  box-sizing: border-box;
}

ul {
  margin: 0;
  padding: 0;
}

ul li:nth-child(odd) {
  background: #f9f9f9;
}

/* Darker background-color on hover */
ul li:hover {
  background: #ddd;
}

.close {
  position: absolute;
  right: 0;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.close:hover {
  background-color: #f44336;
  color: white;
}

.header {
  background-color: #ca8c87;
  padding: 30px 40px;
  color: white;
  text-align: center;
}

.header:after {
  content: "";
  display: table;
  clear: both;
}

input {
  margin: 0;
  border: none;
  border-radius: 0;
  width: 75%;
  padding: 10px;
  float: left;
  font-size: 16px;
}

.addBtn {
  padding: 10px;
  width: 25%;
  background: #d9d9d9;
  color: rgb(128, 110, 110);
  float: left;
  text-align: center;
  font-size: 16px;
  cursor: pointer;
  transition: 0.3s;
  border-radius: 0;
}

li {
    list-style: none;
    padding: 10px 10px;
    border-bottom: 2px solid rgb(105, 67, 67);
  }

  #checkBox{
    position: absolute;
    margin-top: 12px;
  }

  #span{
    position: absolute;
    left: 100px;
    margin-top: -25px;
  }

  #Button{
    display: inline-block;
    font-size: 20px;
    border-radius: 0px;
    background-color: white;
    width: 18%;
    

  }

</style>