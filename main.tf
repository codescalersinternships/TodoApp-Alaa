terraform {
  required_providers {
    todo = {
      source = "AlaaElattar/TodoApp/tree/review"
    }

  }
}

provider "todo" {
  host = "172.0.0.1"
  port = "8080"
  apipath = "/"
  schema = "http"
  
}

resource "todo" "task1" {
  index = 1
  task = "Github Actions"
  completed = false
  
}




